package exchange

import (
	"fmt"
	"github.com/behavioral-ai/core/host"
	"github.com/behavioral-ai/core/rest"
)

func ExampleNewLinkMap() {
	m := host.NewSyncMap[string, rest.ExchangeLink]()
	name := ""
	t := m.Load("")
	fmt.Printf("test:  Load(\"%v\") -> %v\n", name, t)

	name = "common:core:ctor/test"
	m.Store(name, nil)
	//fmt.Printf("test:  store(\"%v\") -> %v\n", name, t)

	m.Store(name, func(next rest.Exchange) rest.Exchange { return nil })
	t = m.Load(name)
	fmt.Printf("test:  Load(\"%v\") -> %v\n", name, t != nil)

	//Output:
	//test:  Load("") -> <nil>
	//test:  Load("common:core:ctor/test") -> true

}
