package router

import (
	"println/httprouter"
)

type Router struct {
	Port   string
	Router *httprouter.Router
}

const PORT = "8080"

//添加路由...
func Add(rootPath, method string) {

}

// func (r *Router) ListenPort(port string) {
// 	portInt64, err := strconv.ParseInt(port, 10, 64)
// 	if err != nil {
// 		log.Fatal("类型转失败")
// 	}
// 	if portInt64 <= 0 {
// 		port = PORT
// 	}
// 	port = ":" + port
// 	err := http.ListenAndServe(port, nil)
// 	if err != nil {

// 	}
// }

var routerTmp Router

func init() {
	router := httprouter.New()
	routerTmp.Router = router
}
