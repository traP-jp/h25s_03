package main

import (
	"strings"

	"github.com/labstack/echo/v4"
)

type routerConfig struct {
	method      string
	path        string
	isGroup     bool
	middlewares []echo.MiddlewareFunc
}

type middlewareRouter struct {
	routerConfigs []routerConfig
}

func newMiddlewareRouter() *middlewareRouter {
	return &middlewareRouter{
		routerConfigs: []routerConfig{},
	}
}

func (mr *middlewareRouter) addRoute(method string, path string, middlewares ...echo.MiddlewareFunc) {
	mr.routerConfigs = append(mr.routerConfigs, routerConfig{
		method:      method,
		path:        path,
		isGroup:     false,
		middlewares: middlewares,
	})
}

func (mr *middlewareRouter) addGroup(method string, path string, middlewares ...echo.MiddlewareFunc) {
	mr.routerConfigs = append(mr.routerConfigs, routerConfig{
		method:      method,
		path:        path,
		isGroup:     true,
		middlewares: middlewares,
	})
}

func (mr *middlewareRouter) getMiddlewares(method string, path string) []echo.MiddlewareFunc {
	var middlewares []echo.MiddlewareFunc
	for _, config := range mr.routerConfigs {
		if config.method == method && (config.path == path || (config.isGroup && strings.HasPrefix(path, config.path+"/"))) {
			middlewares = append(middlewares, config.middlewares...)
		}
	}
	return middlewares
}

func (mr *middlewareRouter) registerMiddlewares(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		path := c.Path()
		method := c.Request().Method
		middlewares := mr.getMiddlewares(method, path)
		for i := len(middlewares) - 1; i >= 0; i-- {
			next = middlewares[i](next)
		}
		return next(c)
	}
}
