package goREST

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"
)

type TestMethodsHandler struct {
	DefaultEndpointMethodsHandler
}

var (
	testPrefix = "/api/events"

	logger = NewStdLogger()

	testMethodsHdlr = TestMethodsHandler{}
	testHttpClient  = &http.Client{}

	testEventRootUrl = "http://localhost:7878/api/events"
	testEventTypeUrl = "http://localhost:7878/api/events/alarm"
	testEventUrl     = "http://localhost:7878/api/events/alarm/some_alarm_id"
)

func get404Test(t *testing.T) {
	// GET 404 //
	httpResp, err := http.Get(fmt.Sprintf("%s/404test", testEventUrl))
	if err != nil {
		t.Errorf("%s", err)
	}
	if httpResp.StatusCode != 404 {
		t.Errorf("code mismatch: %d", httpResp.StatusCode)
	}
}

func pathMethodsTest(t *testing.T, path string) {
	var (
		httpResp *http.Response
		httpreq  *http.Request
		err      error
	)

	// GET //
	httpResp, err = http.Get(path)
	if err != nil {
		t.Errorf("%s", err)
	}
	if httpResp.StatusCode != 405 {
		t.Errorf("code mismatch: %d", httpResp.StatusCode)
	}
	// POST //
	httpResp, err = http.Post(path, "application/json", bytes.NewBuffer([]byte(`{"name":"test"}`)))
	if err != nil {
		t.Errorf("%s", err)
	}
	if httpResp.StatusCode != 405 {
		t.Errorf("code mismatch: %d", httpResp.StatusCode)
	}
	// PUT //
	httpreq, err = http.NewRequest("PUT", path, nil)
	if err != nil {
		t.Errorf("%s", err)
		t.FailNow()
	}
	httpResp, err = testHttpClient.Do(httpreq)
	if err != nil {
		t.Errorf("%s", err)
	}
	if httpResp.StatusCode != 405 {
		t.Errorf("code mismatch: %d", httpResp.StatusCode)
	}
	// PATCH //
	httpreq, err = http.NewRequest("PATCH", path, nil)
	if err != nil {
		t.Errorf("%s", err)
		t.FailNow()
	}
	httpResp, err = testHttpClient.Do(httpreq)
	if err != nil {
		t.Errorf("%s", err)
	}
	if httpResp.StatusCode != 405 {
		t.Errorf("code mismatch: %d", httpResp.StatusCode)
	}
	// DELETE //
	httpreq, err = http.NewRequest("DELETE", path, nil)
	if err != nil {
		t.Errorf("%s", err)
		t.FailNow()
	}
	httpResp, err = testHttpClient.Do(httpreq)
	if err != nil {
		t.Errorf("%s", err)
	}
	if httpResp.StatusCode != 405 {
		t.Errorf("code mismatch: %d", httpResp.StatusCode)
	}
}

func Test_NewRESTRouter(t *testing.T) {
	var rtr = NewRESTRouter(testPrefix, logger)

	rtr.Register("/", &testMethodsHdlr)
	rtr.Register("/:type", &testMethodsHdlr)
	rtr.Register("/:type/:_id", &testMethodsHdlr)

	go func() {
		http.ListenAndServe(":7878", nil)
	}()

	pathMethodsTest(t, testEventRootUrl)
	pathMethodsTest(t, testEventTypeUrl)
	pathMethodsTest(t, testEventUrl)

	get404Test(t)
}
