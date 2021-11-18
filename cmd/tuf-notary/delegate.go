package main

import (
	"fmt"
	"strings"

	docopt "github.com/docopt/docopt-go"
	tufnotary "github.com/notaryproject/tuf/tuf-notary/tuf-notary"
)

func init() {
	register("delegate", cmdDelegate, `
usage: tuf-notary delegate <registry> <delegateeName> [--repo=<repository> --keyfiles=<namess> --threshold=<threshold>]

Add a delegation from the top-level targets role to delegatee and
push the updated targets metadata to the TUF reposistory on the registry.

Options:
  --repo		Set the tuf repository name. By default this will be 'tuf-repo'
  --keyfiles	Comma separaged names of public key files stored in tuf-repo/keys that will be used to sign this delegated role. If none are supplied, a keypair will be generated and written to tuf-repo/keys/<delegate>
  --threshold	The threshold for the delegation. By default this will be 1.
  `)
}

func cmdDelegate(args []string, opts docopt.Opts) error {
	repository := "tuf-repo"
	if r := opts["--repo"]; r != nil {
		repository = r.(string)
	}

	threshold := 1
	if t := opts["-threshold"]; t != nil {
		threshold = t.(int)
	}

	keyfiles := []string{}
	if k := opts["--keyfiles"]; k != nil {
		ks := k.(string)
		splitKeys := strings.Split(ks, ",")
		for _, key := range splitKeys {
			keyfiles = append(keyfiles, key)
		}
	}

	registry := args[0]
	delegatee := args[1]

	err := tufnotary.DownloadTUFMetadata(registry, repository, "root")
	if err != nil {
		return err
	}
	err = tufnotary.DownloadTUFMetadata(registry, repository, "targets")
	if err != nil {
		return err
	}

	//add delegation
	err = tufnotary.Delegate(repository, delegatee, keyfiles, threshold)

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
