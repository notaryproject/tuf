package tufnotary

import (
	"context"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
	"time"

	//ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/distribution/distribution/v3/configuration"
	"github.com/distribution/distribution/v3/registry"
	_ "github.com/distribution/distribution/v3/registry/storage/driver/inmemory"
	"github.com/phayes/freeport"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RegistryTestSuite struct {
	suite.Suite
	RegistryHost string
}

func (suite *RegistryTestSuite) SetupTest() {
	// set up registry
	port, err := freeport.GetFreePort()
	if err != nil {
		suite.Nil(err, "no error finding free port for test registry")
	}
	config := &configuration.Configuration{}
	config.HTTP.Addr = fmt.Sprintf(":%d", port)
	config.HTTP.DrainTimeout = time.Duration(10) * time.Second
	config.Storage = map[string]configuration.Parameters{"inmemory": map[string]interface{}{}}
	suite.RegistryHost = fmt.Sprintf("localhost:%d", port)
	dockerRegistry, err := registry.NewRegistry(context.Background(), config)

	go dockerRegistry.ListenAndServe()
}

func (suite *RegistryTestSuite) TestUploadTUFMetadata() {
	contents, err := ioutil.ReadFile("test/tuf-repo/staged/root.json")
	assert.Nil(suite.T(), err)

	//good case
	desc, err := UploadTUFMetadata(suite.RegistryHost, "test-tuf-repo", "root", contents, "")
	assert.Nil(suite.T(), err)
	assert.True(suite.T(), strings.HasPrefix(desc.Digest.String(), "sha256"))

	//bad registry
	badHost := fmt.Sprintf("localhost:%d", 2)
	desc, err = UploadTUFMetadata(badHost, "test-tuf-repo", "root", contents, "")
	assert.NotNil(suite.T(), err)
}

func TestRegistryTestSuite(t *testing.T) {
	suite.Run(t, new(RegistryTestSuite))
}
