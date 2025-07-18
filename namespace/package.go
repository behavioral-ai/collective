package namespace

import (
	"fmt"
	"github.com/behavioral-ai/core/messaging"
	"sync/atomic"
)

// Namespace names format:
//  collective = domain-sub-domain
//  domain = domain-sub-domain
//
//  {collective}:{domain}:{type}/{path}#{fragment}
//
//  Example:  wikipedia-eng:resiliency-traffic:agent/rate-limiting/request/http#v1.2.3
//
// What would be a collective name in the root registry?
//     registry:{dns-name}:collective/{collective-name}
//
// TODO : add activity
//        Can we add frame resource version to activity?
//        Need author, frame, things, Accessor
//        Need counts, lists of aspects

const (
	Fragment = "#"
	Colon    = ":"
	Slash    = "/"
)

// For humans : thing -> understanding -> relation
//RelationKind = "relation" // Used for relating 2 resources
// What is a name? -> a word or phrase that constitutes the distinctive designation of a person or thing
//

// How to support temporal ordering - order, on a join??
// We could add an aspect to the join, which would create a one-to-one relationship between the two things
// and an aspect. To implement with the current architecture would require to create a unique aspect every time
// and ordering is created. Maybe too much overhead?? Also, the created unique order aspect would be ignored by all
// attempts/queries to determine relationships.
// If this is an infrequently used functionality, maybe the creation of a unique order aspect is not too much overhead.
// However, if there is other aspects that need to be used, then it becomes harder to justify the overhead.
// But, to join 2 things on an aspect other than order, the normal mechanism would be used.
// So, maybe there is an order join?
// The things in an order join would be constrained to: perceptions, observations, events.
// So, if a perception is an aspect will that work?
// Probably not limit thing1, and thing2, but the aspect would always be core:common:order

const (
	ThingKind = "thing" // A name
	JoinKind  = "join"  // Join 2, maybe 3 things
	FrameKind = "frame" // A container for things, similar to a dir entry on a file system

	// AspectKind - A way in which something can be viewed by the mind. Aspects vary from real
	//	and contextualized(LH), to artificial and de-contextualized(RH).
	//	A named understanding, often linked with things, enabling relations/connections/associations and generalization.
	//	For analogies, a "relation" is used instead of an aspect.
	//  Are analogy "relations" an understanding?  In  these cases, the relations we perceive only
	//  exist in the mind of the observer.
	AspectKind = "aspect"

	// CollectiveKind -
	CollectiveKind = "collective" // References a collective in the registry/root collective
	LinkKind       = "link"       // References another collective to link to
	HandlerKind    = "handler"    // Function used as a link in a Micro-REST chain
	UriKind        = "uri"        // Http URI

	AgentKind   = "agent"   // Used to define a thing that is empowered, agents are members of the collective, not just things.
	EventKind   = "event"   // Messaging events
	PersonKind  = "person"  // Used for authorization and ownership resources: self, info, instance
	ServiceKind = "service" // Service name
	TaskKind    = "task"    // Used for tracing agent activity. What is the agent tasked with
	TypeKind    = "type"    // Programming language types
	ViewKind    = "view"    // View names are for namespace retrievals.
	NetworkKind = "network" // A network of agents configuration

)

var (
	counter = new(atomic.Int64)
)

// Name -
type Name struct {
	Collective string `json:"collective"`
	Domain     string `json:"domain"`
	Kind       string `json:"kind"`
	Path       string `json:"path"`
	Fragment   string `json:"fragment"`
}

func ParseName(name string) Name {
	return parse(name)
}

func AddFragment(name, fragment string) string {
	return addFragment(name, fragment)
}

func Versioned(name string) string {
	return AddFragment(name, fmt.Sprintf("%v", counter.Add(1)))
}

func Kind(name string) string {
	return ParseName(name).Kind
}

// Accessor -
type Accessor struct {
	//Version string // returned on a Get
	Type    string // Content-Type
	Content any
}

// Adder - add
type Adder struct {
	Thing    func(name, cname, author string) *messaging.Status
	Relation func(name, cname, thing1, thing2, author string) *messaging.Status
	// What exactly are the results?
	// How to query+select/return generational information
	// Content can be captured if provided.
	ConnectThing  func(name, frame, author string, access Accessor) (results string, status *messaging.Status)
	ConnectAspect func(name []string, frame, author string, access Accessor) (results string, status *messaging.Status)
}

// Add -
var Add = func() *Adder {
	return &Adder{
		Thing: func(name, cname, author string) *messaging.Status {
			return agent.addThing(name, cname, author)
		},
		Relation: func(name, cname, thing1, thing2, author string) *messaging.Status {
			return agent.addRelation(name, cname, thing1, thing2, author)
		},
	}
}()
