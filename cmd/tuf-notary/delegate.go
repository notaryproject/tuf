package main

import (
	"fmt"

	docopt "github.com/docopt/docopt-go"
	tufnotary "github.com/notaryproject/tuf/tuf-notary/tuf-notary"
)

func init() {
	register("delegate", cmdDelegate, `
usage: tuf-notary delegate <registry> <delegateeName> [--repo=<repository> --keyfile=<path>] 

Add a delegation from the top-level targets role to delegatee and
push the updated targets metadata to the TUF reposistory on the registry.

Options:
  --repo	Set the tuf repository name. By default this will be 'tuf-repo'
  `)
}

func cmdDelegate(args []string, opts docopt.Opts) error {
	repository := "tuf-repo"
	if r := opts["--repo"]; r != nil {
		repository = r.(string)
	}

	//TODO support multiple keys
	keyfiles := []string{}
	if k := opts["--keyfile"]; k != nil {
		keyfiles = append(keyfiles, k.(string))
	}

	//TODO get from user
	threshold := 1

	registry := args[0]
	delegatee := args[1]

	//TODO: pull current targets metadata from the registry

	//add delegation
	err := tufnotary.Delegate(repository, delegatee, keyfiles, threshold)

	if err != nil {
		return err
	}
	fmt.Println("added delegation to " + delegatee)

	//upload targets with a reference to root metadata
	targets_desc, err := tufnotary.UploadTUFMetadata(registry, repository, "targets", "root")
	if err != nil {
		return err
	}
	fmt.Println("uploaded targets " + targets_desc.Digest.String())

	return err
}
