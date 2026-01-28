package router

import (
	"net/http"
)

// Define the handler function signature.
type HandlerFunc func(http.ResponseWriter, *http.Request)

// ---------------------------------
// routes := {
// 	"/health": {
// 		"GET": HealthHandler,
// 	},
// 	"/users": {
// 		"GET":  GetUsers,
// 		"POST": CreateUser,
// 	},
// }
// ---------------------------------
type Router struct {
	// map[path]map[method]handler
	routes map[string]map[string]HandlerFunc
}

// This will initialize the router map and return a new ready to use router
func New() *Router {
	return &Router{
		routes: make(map[string]map[string]HandlerFunc),
	}
}

// This function is used for creating new routes.
// Usage: router.Handle("GET", "/health", Health)
func (r *Router) Handle(method, path string, handler HandlerFunc) {
	if r.routes[path] == nil {
		r.routes[path] = make(map[string]HandlerFunc)
	}
	r.routes[path][method] = handler
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if methods, ok := r.routes[req.URL.Path]; ok {
		if handler, ok := methods[req.Method]; ok {
			handler(w, req)
			return
		}
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.Error(w, "Not Found", http.StatusNotFound)
}
