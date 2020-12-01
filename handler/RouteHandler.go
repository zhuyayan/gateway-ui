package handler

import (
	"fmt"
	"net/http"
)

type RouteHandler struct {

}
func (this *RouteHandler) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(w, "route method")
}
