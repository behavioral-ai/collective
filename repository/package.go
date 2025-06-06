package repository

import "github.com/behavioral-ai/core/messaging"

var (
	exchange = messaging.NewExchange()
	ctor     = newCtorMap[string, messaging.NewAgent]()
	msg      = newMessageMap[string, *messaging.Message]()
)

// Register - register an agent
func Register(a messaging.Agent) error {
	return exchange.Register(a)
}

// Agent - get an agent
func Agent(name string) messaging.Agent {
	agent := exchange.Get(name)
	if agent != nil {
		return agent
	}
	agent = Constructor(name)
	if agent == nil {
		return nil
	}
	Register(agent)
	return agent
}

// Message - message an agent
func Message(m *messaging.Message) bool {
	return exchange.Message(m)
}

// Broadcast - broadcast a message
func Broadcast(m *messaging.Message) {
	exchange.Broadcast(m)
}

// RegisterConstructor - register a new agent function
func RegisterConstructor(name string, fn messaging.NewAgent) {
	if name == "" || fn == nil {
		return
	}
	ctor.store(name, fn)
}

// Constructor - construct a new agent
func Constructor(name string) messaging.Agent {
	fn := ctor.get(name)
	if fn != nil {
		return fn()
	}
	return nil
}

func Exists(name string) bool {
	return ctor.get(name) != nil
}

// GetMessage - get a message
func GetMessage(name string) *messaging.Message {
	return msg.get(name)
}

// StoreMessage - store a message
func StoreMessage(m *messaging.Message) {
	msg.store(m.Name, m)
}

/*
// ModifyMessage - modify a message
func ModifyMessage(m *messaging.Message) {
	msg.modify(m)
}

// DeleteMessage - delete a message
func DeleteMessage(name string) {
	msg.delete(name)
}


*/
