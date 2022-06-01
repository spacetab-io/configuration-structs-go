package cfgstructs_test

import (
	"errors"
	"testing"

	cfgstructs "github.com/spacetab-io/configuration-structs-go/v2"
	"github.com/spacetab-io/configuration-structs-go/v2/contracts"
	"github.com/stretchr/testify/assert"
)

func TestNewConfigValidator(t *testing.T) {
	t.Parallel()

	cv := cfgstructs.NewConfigValidator()

	assert.Implements(t, new(contracts.ConfigValidatorInterface), cv)
}

func TestConfigValidator_Add(t *testing.T) {
	t.Parallel()

	cv := cfgstructs.NewConfigValidator()
	tc := testConfig{}
	cv.Add(tc)

	assert.Equal(t, []contracts.ValidatableInterface{tc}, cv.GetConfigs())
}

type testConfig struct {
	RequiredParam string `valid:"required"`
}

func (tc testConfig) String() string {
	return "test"
}

func (tc testConfig) Validate() (bool, error) {
	return contracts.ConfigValidate(tc)
}

type testConfigAlter struct {
	RequiredParam string
}

var errEmptyReqParam = errors.New("empty required param")

func (tca testConfigAlter) String() string {
	return "testAlter"
}

func (tca testConfigAlter) Validate() (bool, error) {
	if tca.RequiredParam == "" {
		return false, errEmptyReqParam
	}

	return true, nil
}

func TestConfigValidator_Validate(t *testing.T) {
	type testCase struct {
		name     string
		in       contracts.ValidatableInterface
		hasError bool
	}

	tcs := []testCase{
		{
			name:     "has validating errors",
			in:       testConfig{},
			hasError: true,
		},
		{
			name:     "no validating errors",
			in:       testConfig{RequiredParam: "exists"},
			hasError: false,
		},
		{
			name:     "alter config with custom validate func",
			in:       testConfigAlter{},
			hasError: true,
		},
	}

	t.Parallel()

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			cv := cfgstructs.NewConfigValidator()
			cv.Add(tc.in)
			err := cv.Validate()

			if tc.hasError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
