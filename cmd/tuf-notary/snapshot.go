package main

import (
	"fmt"
	"io/ioutil"

	docopt "github.com/docopt/docopt-go"
	tufnotary "github.com/notaryproject/tuf/tuf-notary"
)

func init() {
	register("snapshot", cmdSnapshot, `
usage: tuf-notary snapshot <registry> [--repo=<repository>]

Generate snapshot metadata and push it to the TUF repository on the
registry

Options:
  --repo	Set the tuf repository name. By default this will be 'tuf-repo'
  `)
}

func cmdSnapshot(args []string, opts docopt.Opts) error {
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

	//TODO: ensure that delegated targets are also included
	//TODO: get passphrase bool from argument
	err = tufnotary.Snapshot(repository, false)

	if err != nil {
		return err
	}

	//upload snapshot with a reference to root metadata
	filename := fmt.Sprintf("%s/staged/%s.json", repository, "snapshot")
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read %s: %w", filename, err)
	}
	snapshot_desc, err := tufnotary.UploadTUFMetadata(registry, repository, "snapshot", contents, "root")
	if err != nil {
		return err
	}
	fmt.Println("uploaded snapshot " + snapshot_desc.Digest.String())

	return err
}
