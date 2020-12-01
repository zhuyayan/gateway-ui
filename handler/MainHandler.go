package handler

import (
	"fmt"
	"net/http"
)

type MainHandler struct {

}

func (this *MainHandler) ServeHTTP(w http.ResponseWriter, request *http.Request)  {
	switch request.Method {
	case http.MethodGet:
		fmt.Fprintln(w, "123456")
		break
	case http.MethodPost:
		fmt.Fprintln(w, "This interface only support GET Method")
		break
	default:
		fmt.Fprintln(w, "Error Request")
		break
	}
}
