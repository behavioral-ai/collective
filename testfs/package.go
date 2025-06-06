package testfs

import (
	"embed"
	"github.com/behavioral-ai/core/iox"
)

const (
	BehavioralAIHtmlExample = "file:///f:/files/behavioral-ai/example.html"
	BehavioralAIJsonExample = "file:///f:/files/behavioral-ai/threshold.json"
	BehavioralAITextExample = "file:///f:/files/behavioral-ai/get-resp.txt"

	ResiliencyTrafficProfile1 = "file:///f:/files/metrics/traffic-profile-1.json"

	ResiliencyThreshold1 = "file:///f:/files/resiliency/threshold-1.json"
	ResiliencyInterpret1 = "file:///f:/files/resiliency/interpret-1.json"

	ResiliencyThreshold2 = "file:///f:/files/resiliency/threshold-2.json"
	ResiliencyInterpret2 = "file:///f:/files/resiliency/interpret-2.json"
)

//go:embed files
var f embed.FS

func init() {
	iox.Mount(f)
}

func ReadFile(name string) ([]byte, error) {
	return iox.ReadFile(name)
}
