package main

import (
	"fmt"
	"io/ioutil"

	docopt "github.com/docopt/docopt-go"
	tufnotary "github.com/notaryproject/tuf/tuf-notary"
)

func init() {
	register("timestamp", cmdTimestamp, `
usage: tuf-notary timestamp <registry> [--repo=<repository>]

Generate timestamp metadata and push it to the TUF repository on the
registry

Options:
  --repo	Set the tuf repository name. By default this will be 'tuf-repo'
  `)
}

func cmdTimestamp(args []string, opts docopt.Opts) error {
	repository := "tuf-repo"
	if r := opts["--repo"]; r != nil {
		repository = r.(string)
	}

	registry := args[0]

	err := tufnotary.DownloadTUFMetadata(registry, repository, "root")
	if err != nil {
		return err
	}
	err = tufnotary.DownloadTUFMetadata(registry, repository, "targets")
	if err != nil {
		return err
	}
	//TODO: verify the snapshot before adding timestamp
	err = tufnotary.DownloadTUFMetadata(registry, repository, "snapshot")
	if err != nil {
		return err
	}

	//TODO: get passphrase bool from argument
	err = tufnotary.Timestamp(repository, false)

	if err != nil {
		return err
	}

	//upload timestamp with a reference to root metadata
	filename := fmt.Sprintf("%s/staged/%s.json", repository, "timestamp")
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read %s: %w", filename, err)
	}
	timestamp_desc, err := tufnotary.UploadTUFMetadata(registry, repository, "timestamp", contents, "root")
	if err != nil {
		return err
	}
	fmt.Println("uploaded timestamp " + timestamp_desc.Digest.String())

	return err
}
