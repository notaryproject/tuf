package tufnotary

import (
	"fmt"

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

	fmt.Println("init repo")
	//not using consistent snapshots
	err = repo.Init(false)
	if err != nil {
		return err
	}

	fmt.Println("root key")
	//add root key
	_, err = repo.GenKey("root")
	if err != nil {
		return err
	}

	fmt.Println("targets key")
	//add targets key
	_, err = repo.GenKey("targets")
	if err != nil {
		return err
	}

	fmt.Println("snapshot key")
	//add snapshot key
	_, err = repo.GenKey("snapshot")
	if err != nil {
		return err
	}

	fmt.Println("timetamp key")
	//add timestamp key
	_, err = repo.GenKey("timestamp")
	if err != nil {
		return err
	}

	fmt.Println("empty targets")
	//make empty targets metadata
	emptyTargets := []string{}
	err = repo.AddTargets(emptyTargets, nil)
	if err != nil {
		return err
	}

	fmt.Println("snapshot")
	err = repo.Snapshot()
	if err != nil {
		return err
	}

	fmt.Println("timestamp")
	err = repo.Timestamp()
	return err
}
