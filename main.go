package main

import (
	"mymodule/Routes"
	"net/http"
)

func main() {
	Routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
