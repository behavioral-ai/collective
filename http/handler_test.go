package http

import (
	"fmt"
	"github.com/behavioral-ai/core/httpx"
	"github.com/behavioral-ai/core/iox"
	"net/http"
	"net/http/httptest"
)

func ExampleHandler_Text() {
	exchange(textResource)

	//Output:
	//test: Handler() -> [status:200] [ct:text/plain; charset=utf-8] [len:1649]
	//test: Handler() -> [len:1649] [err:<nil>]

}

func ExampleHandler_Html() {
	exchange(htmlResource)

	//Output:
	//test: Handler() -> [status:200] [ct:text/html; charset=utf-8] [len:108]
	//test: Handler() -> [len:108] [err:<nil>]

}

func ExampleHandler_Json() {
	exchange(jsonResource)

	//Output:
	//test: Handler() -> [status:200] [ct:text/plain; charset=utf-8] [len:303]
	//test: Handler() -> [len:303] [err:<nil>]

}

func exchange(rsc string) {
	req, _ := http.NewRequest(http.MethodGet, "https://localhost:8081"+rsc, nil)
	rec := httptest.NewRecorder()

	Handler(rec, req)
	rec.Flush()
	fmt.Printf("test: Handler() -> [status:%v] [ct:%v] [len:%v]\n", rec.Result().StatusCode, rec.Result().Header.Get(httpx.ContentType), rec.Result().Header.Get(httpx.ContentLength))

	buf, err := iox.ReadAll(rec.Result().Body, nil)
	fmt.Printf("test: Handler() -> [len:%v] [err:%v]\n", len(buf), err)

}
