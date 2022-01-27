package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	docopt "github.com/docopt/docopt-go"
	tufnotary "github.com/notaryproject/tuf/tuf-notary"
)

func init() {
	register("delegate", cmdDelegate, `
usage: tuf-notary delegate <registry> <delegateeName> [--repo=<repository> --keyfiles=<names> --threshold=<threshold> --no-passphrase]

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

	passphrase := true
	if p := opts["--no-passphrase"]; p != nil {
		passphrase = !p.(bool)
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
	err = tufnotary.Delegate(repository, delegatee, keyfiles, threshold, passphrase)

	if err != nil {
		return err
	}
	fmt.Println("added delegation to " + delegatee)

	//upload targets with a reference to root metadata
	filename := fmt.Sprintf("%s/staged/%s.json", repository, "targets")
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read %s: %w", filename, err)
	}
	targets_desc, err := tufnotary.UploadTUFMetadata(registry, repository, "targets", contents, "root")
	if err != nil {
		return err
	}
	fmt.Println("uploaded targets " + targets_desc.Digest.String())

	return err
}
