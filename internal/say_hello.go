package internal

import (
	"net/http"
	"strings"
)

// SayHello test
func SayHello(writer http.ResponseWriter, request *http.Request) {
	message := request.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message

	writer.Write([]byte(message))
}
