package tufnotary

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/theupdateframework/go-tuf"
	"github.com/theupdateframework/go-tuf/data"
	"github.com/theupdateframework/go-tuf/pkg/keys"
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

func Delegate(repository string, delegatee string, keyfiles []string, threshold int) error {
	workingDir, err := os.Getwd()
	if err != nil {
		return err
	}

	dir := filepath.Join(workingDir, repository)

	//TODO: allow for this to be true
	// insecure := true
	var p util.PassphraseFunc
	//if !insecure {
	//p = getPassphrase
	//}

	repo, err := tuf.NewRepo(tuf.FileSystemStore(dir, p))
	if err != nil {
		return err
	}

	pubkeys := []*data.PublicKey{}
	privkeys := []keys.Signer{}
	keyids := []string{}
	if len(keyfiles) < 1 {
		key, err := keys.GenerateEd25519Key()
		if err != nil {
			return err
		}
		pubkeys = append(pubkeys, key.PublicData())
		privkeys = append(privkeys, key)
		fmt.Println(key.PublicData())
		for id, _ := range key.PublicData().IDs() {
			fmt.Println(id)
			//keyids = append(keyids, id)
		}
	} else {
		//TODO parse keys, multiple keys
		//also probably make this a helper function
		_, err := ioutil.ReadFile(filepath.Join(workingDir, keyfiles[0]))
		if err != nil {
			return err
		}
		// pubkeys = append(pubkeys, key)
		// keyids = append(keyids, keyid)
	}

	paths := []string{}
	paths = append(paths, delegatee+"/*")

	delegatedRole := data.DelegatedRole{
		Name:      delegatee,
		KeyIDs:    keyids,
		Paths:     paths,
		Threshold: threshold,
	}

	err = repo.AddTargetsDelegation("targets", delegatedRole, pubkeys)
	if err != nil {
		return err
	}

	//if keys were generated, store them
	// for k := range privkeys {
	// repo.local.SaveSigner(delegatee, k)
	// }

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
