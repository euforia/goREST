package goREST

import (
	"net/http"
)

type EndpointMethodsHandler interface {
	DELETE(r *http.Request, args ...string) (map[string]string, interface{}, int)
	GET(r *http.Request, args ...string) (map[string]string, interface{}, int)
	OPTIONS(r *http.Request, args ...string) (map[string]string, interface{}, int)
	PATCH(r *http.Request, args ...string) (map[string]string, interface{}, int)
	POST(r *http.Request, args ...string) (map[string]string, interface{}, int)
	PUT(r *http.Request, args ...string) (map[string]string, interface{}, int)
}

type DefaultEndpointMethodsHandler struct{}

func (d *DefaultEndpointMethodsHandler) DELETE(r *http.Request, args ...string) (map[string]string, interface{}, int) {
	return nil, map[string]string{"error": "Invalid method"}, 405
}
func (d *DefaultEndpointMethodsHandler) GET(r *http.Request, args ...string) (map[string]string, interface{}, int) {
	return nil, map[string]string{"error": "Invalid method"}, 405
}
func (d *DefaultEndpointMethodsHandler) OPTIONS(r *http.Request, args ...string) (map[string]string, interface{}, int) {
	return nil, map[string]string{"error": "Invalid method"}, 405
}
func (d *DefaultEndpointMethodsHandler) PATCH(r *http.Request, args ...string) (map[string]string, interface{}, int) {
	return nil, map[string]string{"error": "Invalid method"}, 405
}
func (d *DefaultEndpointMethodsHandler) POST(r *http.Request, args ...string) (map[string]string, interface{}, int) {
	return nil, map[string]string{"error": "Invalid method"}, 405
}
func (d *DefaultEndpointMethodsHandler) PUT(r *http.Request, args ...string) (map[string]string, interface{}, int) {
	return nil, map[string]string{"error": "Invalid method"}, 405
}
