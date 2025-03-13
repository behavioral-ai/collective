package content

import (
	"fmt"
	"github.com/behavioral-ai/collective/testrsc"
	url2 "net/url"
)

func ExampleResolution_PutValue() {
	urn := "type/test"
	r := NewEphemeralResolver()

	status := r.PutValue("", "author", nil, 1)
	fmt.Printf("test: PutValue_Name() -> [status:%v]\n", status)

	status = r.PutValue(urn, "author", nil, 1)
	fmt.Printf("test: PutValue_Name() -> [status:%v]\n", status)

	var buff []byte
	status = r.PutValue(urn, "author", buff, 1)
	fmt.Printf("test: PutValue_Name() -> [status:%v]\n", status)

	url, _ := url2.Parse(testrsc.ResiliencyTrafficProfile1)
	status = r.PutValue(urn, "author", url, 1)
	fmt.Printf("test: PutValue() -> [status:%v]\n", status)

	buf, status1 := r.GetValue(urn, 1)
	fmt.Printf("test: GetValue() -> [status:%v] [%v]\n", status1, len(buf) > 0)

	//Output:
	//test: PutValue_Name() -> [status:Invalid Content [err:name is empty on call to PutValue()] [agent:resiliency:agent/domain/content/content]]
	//test: PutValue_Name() -> [status:Invalid Content [err:content is nil on call to PutValue() for name : type/test] [agent:resiliency:agent/domain/content/content]]
	//test: PutValue_Name() -> [status:Invalid Content [err:content is empty on call to PutValue() for name : type/test] [agent:resiliency:agent/domain/content/content]]
	//test: PutValue() -> [status:OK]
	//test: GetValue() -> [status:OK] [true]

}
