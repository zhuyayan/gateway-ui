package main

import (
	"gateway/task/net"
)

// 网络基本配置
type NetSystemConfig struct {

}

//启动任务
type StartUp struct {

}
func main()  {
	//mainHandler := &handler.MainHandler{}
	//routeHandler := &handler.RouteHandler{}
	//http.Handle("/", mainHandler)
	//http.Handle("/route", routeHandler)
	//http.ListenAndServe(":3000",nil)
	var interfacs = net.ActiveWanInterface{}
	interfacs.Active()
}