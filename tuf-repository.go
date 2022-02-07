package tufnotary

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/theupdateframework/go-tuf"
	util "github.com/theupdateframework/go-tuf/util"
)

func Init(repository string) error {
	var p util.PassphraseFunc
	//TODO: get passphrase
	repo, err := tuf.NewRepo(tuf.FileSystemStore(repository, p))
	if err != nil {
		return err
	}

	//not using consistent snapshots
	err = repo.Init(false)
	if err != nil {
		return err
	}

	//add root key
	_, err = repo.GenKey("root")
	if err != nil {
		return err
	}

	//add targets key
	_, err = repo.GenKey("targets")
	if err != nil {
		return err
	}

	//add snapshot key
	_, err = repo.GenKey("snapshot")
	if err != nil {
		return err
	}

	//add timestamp key
	_, err = repo.GenKey("timestamp")
	if err != nil {
		return err
	}

	//make empty targets metadata
	emptyTargets := []string{}
	err = repo.AddTargets(emptyTargets, nil)
	if err != nil {
		return err
	}

	err = repo.Snapshot()
	if err != nil {
		return err
	}

	err = repo.Timestamp()
	return err
}

func Sign(tufRepository string, signer string, digest string, length int64, descriptor json.RawMessage) error {

	workingDir, err := os.Getwd()
	if err != nil {
		return err
	}

	dir := filepath.Join(workingDir, tufRepository)

	var p util.PassphraseFunc
	// TODO passphase func (same as in delegations)

	repo, err := tuf.NewRepo(tuf.FileSystemStore(dir, p))
	if err != nil {
		return err
	}

	digestParts := strings.Split(digest, ":")

	path := fmt.Sprintf("%s/%s", signer, digest)
	err = repo.AddTargetsWithDigest(digestParts[1], digestParts[0], length, path, descriptor)
	if err != nil {
		return err
	}

	err = repo.Sign("targets")
	if err != nil {
		return err
	}

	err = repo.Snapshot()
	if err != nil {
		return err
	}

	err = repo.Timestamp()
	if err != nil {
		return err
	}

	err = repo.Commit()
	return err
}
