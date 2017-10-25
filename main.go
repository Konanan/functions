package main

import "./httpsvr"

func main() {
	svr := new(httpsvr.HttpServer)
	svr.StartHttpServer()
}

