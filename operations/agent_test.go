package operations

import "fmt"

func ExampleNewAgent() {
	a := newAgent()

	fmt.Printf("test: newAgent() -> [%v]\n", a)

	//Output:
	//test: newAgent() -> [common:core:agent/operations/collective]

}
