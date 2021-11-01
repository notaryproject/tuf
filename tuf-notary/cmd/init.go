package main

import (
	"fmt"

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

	registry := args["<registry>"].(string)

	err := tufnotary.Init(repository)

	if err != nil {
		return err
	}

	root_desc, err := tufnotary.UploadTUFMetadata(registry, repository, "root", "")
	if err != nil {
		return err
	}
	fmt.Println("uploaded root " + root_desc.Digest.String())

	targets_desc, err := tufnotary.UploadTUFMetadata(registry, repository, "targets", "")
	if err != nil {
		return err
	}
	fmt.Println("uploaded targets " + targets_desc.Digest.String())


	return err
}
