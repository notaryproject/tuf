package main

import (
	"github.com/notaryproject/tuf/tuf-notary/tuf-notary"
)

func init() {
	register("init", cmdInit, `
usage: tuf-notary init <registry> [--repo=<repository>]

Initialize a new repository and push it to the TUF repository on the
registry

Options:
  --repo	Set the tuf repository name. By default this will be 'tuf-repo'
  `)
}

func cmdInit(args map[string]interface{}) error {
	repository := "tuf-repo"
	if r := args["--repo"]; r != nil {
		repository = r.(string)
	}
	err := tufnotary.Init(repository)

	//TODO upload to registry
	return err
}
