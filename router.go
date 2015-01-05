package goREST

import (
	"encoding/json"
	"fmt"
	"github.com/euforia/goREST/logging"
	"net/http"
	"strings"
)

type RESTRouter struct {
	Prefix     string
	handlerMap []EndpointMethodsHandler
	logger     *logging.Logger
}

func NewRESTRouter(prefix string, logger *logging.Logger) *RESTRouter {
	if logger == nil {
		logger = logging.NewStdLogger()
	}

	return &RESTRouter{
		Prefix:     prefix,
		handlerMap: make([]EndpointMethodsHandler, 0),
		logger:     logger,
	}
}

func (s *RESTRouter) writeHttpResponse(w http.ResponseWriter, headers map[string]string, data []byte, respCode int) {
	for k, v := range headers {
		w.Header().Set(k, v)
	}
	w.WriteHeader(respCode)
	if len(data) > 0 {
		w.Write(data)
	}
}

func (s *RESTRouter) writeJsonResponse(writer http.ResponseWriter, headers map[string]string, data interface{}, respCode int) int {
	var (
		bytes []byte
		code  int
		err   error
	)

	bytes, err = json.Marshal(&data)
	if err != nil {
		s.logger.Error.Println(err)
		bytes = []byte(fmt.Sprintf(`{"error": %s}`, err))
		code = 500
	} else {
		code = respCode
	}

	writer.Header().Set("Content-Type", "application/json; charset=utf8")
	s.writeHttpResponse(writer, headers, bytes, code)
	return code
}

func (s *RESTRouter) pathParts(path string) []string {
	parts := make([]string, 0)
	for _, v := range strings.Split(path, "/") {
		if v != "" {
			parts = append(parts, v)
		}
	}
	//s.logger.Trace.Printf("%v\n", parts)
	return parts
}

func (s *RESTRouter) Register(path string, hdlr EndpointMethodsHandler) {
	parts := s.pathParts(path)
	s.logger.Debug.Printf("Registering path: %s%s; Parts: %v\n", s.Prefix, path, parts)

	// 0 reservded for root path.
	if len(parts) == len(s.handlerMap) {
		s.handlerMap = append(s.handlerMap, hdlr)
	} else if len(parts) > len(s.handlerMap) {
		tmap := make([]EndpointMethodsHandler, len(parts)+1)
		for i, v := range s.handlerMap {
			tmap[i] = v
		}
		s.handlerMap = tmap
		s.handlerMap[len(parts)] = hdlr
	}
}

func (s *RESTRouter) runMethodHandler(r *http.Request, handlerIndex int, args ...string) (map[string]string, interface{}, int) {
	var (
		headers = map[string]string{}
		data    interface{}
		code    int
	)
	switch r.Method {
	case "GET":
		headers, data, code = s.handlerMap[handlerIndex].GET(r, args...)
		break
	case "POST":
		headers, data, code = s.handlerMap[handlerIndex].POST(r, args...)
		break
	case "PUT":
		headers, data, code = s.handlerMap[handlerIndex].PUT(r, args...)
		break
	case "DELETE":
		headers, data, code = s.handlerMap[handlerIndex].DELETE(r, args...)
		break
	case "OPTIONS":
		headers, data, code = s.handlerMap[handlerIndex].OPTIONS(r, args...)
		break
	case "PATCH":
		headers, data, code = s.handlerMap[handlerIndex].PATCH(r, args...)
		break
	default:
		data = map[string]string{"error": "Invalid method"}
		code = 405
		break
	}
	return headers, data, code
}

func (s *RESTRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// strip path before processing handler
	r.URL.Path = strings.TrimPrefix(r.URL.Path, s.Prefix)
	var (
		headers = map[string]string{}
		data    interface{}
		code    int
	)
	parts := s.pathParts(r.URL.Path)
	// account for root handler at index 0 //
	if len(parts) < 0 || len(parts) > (len(s.handlerMap)-1) {
		data = map[string]string{"error": "Not found!"}
		code = 404
	} else {
		if s.handlerMap[len(parts)] != nil {
			headers, data, code = s.runMethodHandler(r, len(parts), parts...)
		} else {
			data = map[string]string{"error": "Not found!"}
			code = 404
		}
	}
	s.writeJsonResponse(w, headers, data, code)
	s.logger.Info.Printf("%s %d %s %s\n", r.Method, code, r.RequestURI, r.RemoteAddr)
}
