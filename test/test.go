package main

import (
	"fmt"
	"github.com/bobtfish/go-nsone-api"
)

func main() {
	if k := os.Getenv("NSONE_APIKEY"); k == "" {
		fmt.Println("NSONE_APIKEY environment variable is not set, giving up")
	}
	apiu := nsone.New(k)
	fmt.Println(n.GetZones())

	z := nsone.NewZone("foo.com")
	api.CreateZone(z)

	r := nsone.NewRecord("foo.com", "www.foo.com", "A")
	r.Answers = []nsone.Answer{nsone.Answer{Answer: "1.1.1.1"}}
	api.CreateRecord(r)

	api.DeleteZone("foo.com")
}
