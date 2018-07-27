package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/raggaer/bison/lua"
	glua "github.com/yuin/gopher-lua"
)

// Route defines a lua route
type Route struct {
	Path   string
	Method string
	File   string
}

func loadRoutes() ([]*Route, error) {
	routerState := glua.NewState()

	defer routerState.Close()
	if err := routerState.DoFile("router.lua"); err != nil {
		return nil, err
	}
	routerTable := routerState.Get(-1)

	// Check if returned value is table
	if !lua.IsValueTable(routerTable) {
		return nil, fmt.Errorf("Invalid router.lua returned data. Expected table but got %s", routerTable.Type().String())
	}

	routerMap := lua.TableToMap(routerTable.(*glua.LTable))
	return createRoutes(routerMap), nil
}

func createRoutes(m map[string]interface{}) []*Route {
	dst := []*Route{}
	for path, route := range m {
		methodMap, ok := route.(map[string]interface{})
		if !ok {
			continue
		}
		getPath, ok := methodMap["get"].(string)
		if ok {
			dst = append(dst, &Route{
				Path:   path,
				Method: http.MethodGet,
				File:   getPath,
			})
		}
		postPath, ok := methodMap["post"].(string)
		if ok {
			dst = append(dst, &Route{
				Path:   path,
				Method: http.MethodPost,
				File:   postPath,
			})
		}
	}
	return dst
}

func retrieveCurrentRoute(params map[string]string, method, uri string, routes []*Route) *Route {
	for _, route := range routes {
		// Build the current route
		n := ""
		parts := strings.Split(route.Path, "/")
		for i, part := range parts {
			if strings.HasPrefix(part, ":") {
				v, ok := params[strings.TrimPrefix(part, ":")]
				if ok {
					n += v
				}
			} else {
				n += part
			}
			if i < len(parts)-1 {
				n += "/"
			}
		}

		// Check for route
		if n == uri && method == route.Method {
			return route
		}
	}
	return nil
}
