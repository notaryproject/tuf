package tufnotary

import (
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
	return repo.Init(false)
}
