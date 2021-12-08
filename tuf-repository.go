package tufnotary

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/theupdateframework/go-tuf"
	util "github.com/theupdateframework/go-tuf/util"
	"golang.org/x/crypto/ssh/terminal"
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

func Snapshot(repository string, passphrase bool) error {
	workingDir, err := os.Getwd()
	if err != nil {
		return err
	}

	dir := filepath.Join(workingDir, repository)

	var p util.PassphraseFunc
	if passphrase {
		p = getPassphrase
	}

	repo, err := tuf.NewRepo(tuf.FileSystemStore(dir, p))
	if err != nil {
		return err
	}

	repo.Snapshot()
	repo.Commit()
	return nil
}

func Timestamp(repository string, passphrase bool) error {
	workingDir, err := os.Getwd()
	if err != nil {
		return err
	}

	dir := filepath.Join(workingDir, repository)

	var p util.PassphraseFunc
	if passphrase {
		p = getPassphrase
	}

	repo, err := tuf.NewRepo(tuf.FileSystemStore(dir, p))
	if err != nil {
		return err
	}

	repo.Timestamp()
	repo.Commit()
	return nil
}

//from go-tuf/cmd/tuf/main.go
func getPassphrase(role string, confirm bool) ([]byte, error) {
	if pass := os.Getenv(fmt.Sprintf("TUF_%s_PASSPHRASE", strings.ToUpper(role))); pass != "" {
		return []byte(pass), nil
	}

	fmt.Printf("Enter %s keys passphrase: ", role)
	passphrase, err := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	if err != nil {
		return nil, err
	}

	if !confirm {
		return passphrase, nil
	}

	fmt.Printf("Repeat %s keys passphrase: ", role)
	confirmation, err := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	if err != nil {
		return nil, err
	}

	if !bytes.Equal(passphrase, confirmation) {
		return nil, errors.New("The entered passphrases do not match")
	}
	return passphrase, nil
}
