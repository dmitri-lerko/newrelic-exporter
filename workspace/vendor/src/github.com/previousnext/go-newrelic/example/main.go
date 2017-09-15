package main

import (
	"fmt"

	"github.com/alecthomas/kingpin"
	"github.com/previousnext/go-newrelic"
)

var (
	key     = kingpin.Flag("key", "The API key").Required().String()
	app     = kingpin.Arg("app", "The application name").Required().String()
	version = kingpin.Arg("version", "The version number").Required().String()
	user    = kingpin.Arg("user", "The user who deployed").Required().String()
)

func main() {
	kingpin.Parse()

	nr := newrelic.New(*key)

	id, err := nr.NameToApplicationID(*app)
	if err != nil {
		panic(err)
	}

	fmt.Println("Found ID:", id)

	err = nr.Deployment(id, *version, *user)
	if err != nil {
		panic(err)
	}
}
