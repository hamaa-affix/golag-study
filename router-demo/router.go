package router_demo

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	// "github.com/go-chi/chi/v5/middleware"
)

func Do() {
	//r := boot() -> ここでrouting登録を行いたかったが、うまく登録ができなかった。
	r := chi.NewRouter()

	r.Get("todo", func(rw http.ResponseWriter, r *http.Request) {

	})

	http.ListenAndServe(":8080", r)
}

//package router

// import (
// 	"encoding/json"
// 	"net/http"
// 	"regexp"

// )

// type Route struct {
// 	Method string
// 	Pattern string
// 	Handler http.Handler
// }

// type Router struct {
// 	routes []Route
// }

// func NewRouter() *Router {
// 	return &Router{}
// }

// // routerへの登録
// func (r *Router) AddRoute(method, path string, handler Handler) {
// 	r.routes = append(r.routes, Route{
// 		Method: method,
// 		Pattern: path,
// 		Handler: handler,
// 	})
// }

// func (r *Router) GetRoute(method string, path string) http.Handler {
// 	for _, router := range r.routes {
// 		if router.Method == method && router.Pattern == path {
// 			return router.Handler
// 		}
// 	}

// 	return http.NotFoundHandler()
// }

// func (r *Router) GET(path string, handler Handler) {
// 	r.AddRoute("GET",path, handler)
// }

// func (r *Router) POST(path string, handler Handler) {
// 	r.AddRoute("POST", path, handler)
// }

// func (r *Router) PUT(path string, handler Handler) {
// 	r.AddRoute("PUT", path, handler)
// }

// func (r *Router) DELETE(path string, handler Handler) {
// 	r.AddRoute("DELETE", path, handler)
// }

// func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
// 	path := req.URL.Path
// 	method := req.Method

// 	handler := r.getHandler(method, path)

// 	// handler middlewares go here

// 	handler.ServeHTTP(w, req)
// }

// func (r *Router) getHandler(method, path string) http.Handler {
// 	for _, route := range r.routes {
// 		re := regexp.MustCompile(route.Pattern)
// 		if route.Method == method && re.MatchString(path) {
// 			return route.Handler
// 		}
// 	}

// 	return http.NotFoundHandler()
// }

// type Handler func(r *http.Request) (statusCode int, data map[string]interface{})

// func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	statusCode, data := h(r)
// 	w.WriteHeader(statusCode)
// 	json.NewEncoder(w).Encode(data)
// }


// func New() {

// }

// func Do() {
// 	// http serverの立ち上げ
// 	// http method, route pathでHandlerの取得
// 	http.ListenAndServe(":8080", nil)
// }
