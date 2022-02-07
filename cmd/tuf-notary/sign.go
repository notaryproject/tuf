package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	docopt "github.com/docopt/docopt-go"
	tufnotary "github.com/notaryproject/tuf/tuf-notary"
	godigest "github.com/opencontainers/go-digest"
	artifactspec "github.com/oras-project/artifacts-spec/specs-go/v1"
)

func init() {
	register("sign", cmdSign, `
usage: tuf-notary sign <registry> <rolename> <digest> <length> <mediatype> [--repo=<repository>]

Sign digest and upload the signature alongside it to rolename repo on
the registry. This will add a tuf targets metadata file to the repository.

Options:
  --repo	Set the tuf repository name. By default this will be 'tuf-repo'
  `)
}

func cmdSign(args []string, opts docopt.Opts) error {
	repository := "tuf-repo"
	if r := opts["--repo"]; r != nil {
		repository = r.(string)
	}

	registry := args[0]
	signer := args[1]
	digest := args[2]
	length, err := strconv.ParseInt(args[3], 10, 64)
	if err != nil {
		return err
	}
	mediatype := args[4]

	//TODO verify these
	err = tufnotary.DownloadTUFMetadata(registry, repository, "root")
	if err != nil {
		return err
	}
	err = tufnotary.DownloadTUFMetadata(registry, repository, "targets")
	if err != nil {
		return err
	}

	//TODO add new delegation in the repo for this signature, add this there
	// blocking on the delegations pr in go-tuf

	descriptor := artifactspec.Descriptor{
		MediaType: mediatype,
		Digest:    godigest.FromString(digest),
		Size:      length,
	}
	bytes, _ := json.Marshal(&descriptor)
	fmt.Println(json.RawMessage(bytes))
	err = tufnotary.Sign(repository, signer, digest, length, json.RawMessage(bytes))

	if err != nil {
		return err
	}

	//TODO: once the signature is added to the correct delegated metadata,
	//upload the correct thing here

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
