package eventing

import "fmt"

func ExampleNewAgent() {
	a := newAgent(nil, nil, nil)

	fmt.Printf("test: newAgent() -> [%v]\n", a)

	//Output:
	//test: newAgent() -> [resiliency:agent/behavioral-ai/collective/eventing]

}
