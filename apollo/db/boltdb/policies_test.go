package boltdb

import (
	"io/ioutil"
	"os"
	"testing"

	"bitbucket.org/stack-rox/apollo/pkg/api/generated/api/v1"
	"github.com/stretchr/testify/suite"
)

func TestBoltPolicies(t *testing.T) {
	suite.Run(t, new(BoltPoliciesTestSuite))
}

type BoltPoliciesTestSuite struct {
	suite.Suite
	*BoltDB
}

func (suite *BoltPoliciesTestSuite) SetupSuite() {
	tmpDir, err := ioutil.TempDir("", "")
	if err != nil {
		suite.FailNow("Failed to get temporary directory", err.Error())
	}
	db, err := New(tmpDir)
	if err != nil {
		suite.FailNow("Failed to make BoltDB", err.Error())
	}
	suite.BoltDB = db
}

func (suite *BoltPoliciesTestSuite) TeardownSuite() {
	suite.Close()
	os.Remove(suite.Path())
}

func (suite *BoltPoliciesTestSuite) TestPolicies() {
	policy1 := &v1.Policy{
		Name:     "policy1",
		Severity: v1.Severity_LOW_SEVERITY,
	}
	err := suite.AddPolicy(policy1)
	suite.Nil(err)

	policy2 := &v1.Policy{
		Name:     "policy2",
		Severity: v1.Severity_HIGH_SEVERITY,
	}
	err = suite.AddPolicy(policy2)
	suite.Nil(err)
	// Get all alerts
	policies, err := suite.GetPolicies(&v1.GetPoliciesRequest{})
	suite.Nil(err)
	suite.Equal([]*v1.Policy{policy1, policy2}, policies)

	policy1.Severity = v1.Severity_HIGH_SEVERITY
	err = suite.UpdatePolicy(policy1)
	suite.Nil(err)
	policies, err = suite.GetPolicies(&v1.GetPoliciesRequest{})
	suite.Nil(err)
	suite.Equal([]*v1.Policy{policy1, policy2}, policies)

	err = suite.RemovePolicy(policy1.Name)
	suite.Nil(err)
	policies, err = suite.GetPolicies(&v1.GetPoliciesRequest{})
	suite.Nil(err)
	suite.Equal([]*v1.Policy{policy2}, policies)
}
