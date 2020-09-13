/*
 * MIT License
 *
 * Copyright (c) 2020 Beate Ottenwälder
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package core

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/contrib/ginrus"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/ottenwbe/go-cook/utils"
)

const (
	addressCfg    = "html.address"
	corsOriginCfg = "html.cors.origin"

	baseAPIPath = "api"
)

var (
	defaultAddress string
	corsOrigin     string
)

// init configures the router for api calls when the core package is initialized
func init() {
	utils.Config.SetDefault(addressCfg, ":8080")
	utils.Config.SetDefault(corsOriginCfg, "*")
	defaultAddress = utils.Config.GetString(addressCfg)
	corsOrigin = utils.Config.GetString(corsOriginCfg)
}

//Routes is managing a set of API endpoints.
//Routes implementation(s) call handler function to perform typical CRUD operations (GET, POST, PATCH, ...).
type Routes interface {
	//Route is created to a specific set of endpoints
	Route(string) Routes
	//GET endpoint is added to the routes set and registers a corresponding handler
	GET(string, func(c *APICallContext))
	//Path returns the base path
	Path() string
	//PATCH endpoint is added to the routes set and registers a corresponding handler
	PATCH(string, func(c *APICallContext))
	//POST endpoint is added to the routes set and registers a corresponding handler
	POST(string, func(c *APICallContext))
}

//Router is a facade for a HTTP router and can be implemented by a concrete router like gin.
type Router interface {
	API(version int16) Routes
	http.Handler
}

//APICallContext is a facade for any concrete Context, e.g. gins
type APICallContext = gin.Context

//NewRouter creates a router for API calls with a pre-configured ADDRESS
func NewRouter() Router {
	router := &ginRouter{
		gin.New(),
		make(map[string]Routes),
	}
	router.configure()
	router.prepareDefaultRoutes()
	return router
}

type ginRouter struct {
	router       *gin.Engine
	routerGroups map[string]Routes
}

func (g *ginRouter) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	g.router.ServeHTTP(writer, request)
}

func (g *ginRouter) addSubGroup(groupName string, subGroupName string) Routes {
	rg, ok := g.routerGroups[groupName]
	if !ok {
		// we create the missing group if it cannot be found
		rg = g.route(groupName)
		g.routerGroups[groupName] = rg
	}
	return rg.Route(subGroupName)
}

func v(version int16) string {
	return fmt.Sprintf("v%v", version)
}

//API registers the endpoint /api/v<version> and returns a group of endpoints under /api/v<version>
func (g *ginRouter) API(version int16) Routes {
	rg, ok := g.routerGroups[v(version)]
	if !ok {
		rg = g.addSubGroup(baseAPIPath, v(version))
		g.routerGroups[v(version)] = rg
	}
	return rg
}

func (g *ginRouter) route(route string) Routes {
	return &ginRoutes{g.router.Group(route)}
}

// configure the default middleware with a logger and recovery (crash-free) middleware
func (g *ginRouter) configure() {
	g.router.Use(ginrus.Ginrus(log.StandardLogger(), time.RFC3339, true))
	g.router.Use(g.corsMiddleware())
	// Return 500 if there was a panic.
	g.router.Use(gin.Recovery())
}

func (g *ginRouter) prepareDefaultRoutes() {
	g.router.GET("/version", func(c *gin.Context) {
		c.JSON(200, AppVersion())
	})
}

type ginRoutes struct {
	rg *gin.RouterGroup
}

func (g *ginRoutes) Route(path string) Routes {
	return &ginRoutes{g.rg.Group(path)}
}

//GET endpoint for a specific path and a corresponding handler
func (g *ginRoutes) GET(path string, handler func(c *APICallContext)) {
	g.rg.GET(path, handler)
}

//PATCH endpoint for a specific path and a corresponding handler
func (g *ginRoutes) PATCH(path string, handler func(c *APICallContext)) {
	g.rg.PATCH(path, handler)
}

//POST endpoint for a specific path and a corresponding handler
func (g *ginRoutes) POST(path string, handler func(c *APICallContext)) {
	g.rg.POST(path, handler)
}

//PATH of the given route
func (g *ginRoutes) Path() string {
	return g.rg.BasePath()
}

func (g *ginRouter) corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", corsOrigin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, PATCH, POST")

		if c.Request.Method == "OPTIONS" || c.Request.Method == "PUT" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// Server interface
type Server struct {
	Address string
	server  *http.Server
}

//NewServerA creates a new server with default address
func NewServerA(addr string) Server {
	return Server{
		Address: addr,
		server: &http.Server{
			Addr:    addr,
			Handler: NewRouter(),
		}}
}

//NewServerH creates a new server with default address
func NewServerH(handler http.Handler) Server {
	return Server{server: &http.Server{
		Addr:    defaultAddress,
		Handler: handler,
	}}
}

//NewServer creates a new server to listen on the defaultAddress
func NewServer() Server {
	return NewServerA(defaultAddress)
}

//Run the server for the api
func (s Server) Run() {
	go func() {
		if err := s.server.ListenAndServe(); err != nil {
			log.Errorf("Server's running: %s\n", err)
		}
	}()
}

//Close the server
func (s Server) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := s.server.Shutdown(ctx)
	return err
}
