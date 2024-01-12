package tufnotary

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TufTestSuite struct {
	suite.Suite
}

func (suite *TufTestSuite) SetupTest() {
}

func (suite *TufTestSuite) TestInit() {
	//TODO made then delete temp repo for testing
	err := Init("test_repo")
	assert.Nil(suite.T(), err)

	root, err := ioutil.ReadFile("test_repo/staged/root.json")
	assert.Nil(suite.T(), err)
	assert.True(suite.T(), strings.Contains(string(root), "\"_type\":\"root\""))
	assert.True(suite.T(), strings.Contains(string(root), "\"consistent_snapshot\":false"))

	targets, err := ioutil.ReadFile("test_repo/staged/targets.json")
	assert.Nil(suite.T(), err)
	assert.True(suite.T(), strings.Contains(string(targets), "\"_type\":\"targets\""))

	timestamp, err := ioutil.ReadFile("test_repo/staged/timestamp.json")
	assert.Nil(suite.T(), err)
	assert.True(suite.T(), strings.Contains(string(timestamp), "\"_type\":\"timestamp\""))

	snapshot, err := ioutil.ReadFile("test_repo/staged/snapshot.json")
	assert.Nil(suite.T(), err)
	assert.True(suite.T(), strings.Contains(string(snapshot), "\"_type\":\"snapshot\""))
}

func TestTufTestSuite(t *testing.T) {
	suite.Run(t, new(TufTestSuite))
}
