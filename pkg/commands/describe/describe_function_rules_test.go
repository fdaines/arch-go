package describe

import (
	"bytes"
	"github.com/fdaines/arch-go/old/config"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestDescribeFunctionRules(t *testing.T) {

	t.Run("function rules", func(t *testing.T) {
		outputBuffer := bytes.NewBufferString("")
		rules := []*config.FunctionsRule{
			{
				Package:                  "foobar",
				MaxLines:                 123,
				MaxParameters:            32,
				MaxPublicFunctionPerFile: 24,
				MaxReturnValues:          3,
			},
		}
		expectedOutput := `Function Rules
	* Packages that match pattern 'foobar' should comply with the following rules:
		* Functions should not have more than 123 lines
		* Functions should not have more than 32 parameters
		* Functions should not have more than 3 return values
		* Files should not have more than 24 public functions
`

		describeFunctionRules(rules, outputBuffer)

		outputBytes, _ := io.ReadAll(outputBuffer)

		assert.Equal(t, expectedOutput, string(outputBytes), "Output doesn't match expected values.")
	})

	t.Run("function rules with blanks - case 1", func(t *testing.T) {
		outputBuffer := bytes.NewBufferString("")
		rules := []*config.FunctionsRule{
			{
				Package:                  "foobar",
				MaxParameters:            32,
				MaxPublicFunctionPerFile: 24,
				MaxReturnValues:          3,
			},
		}
		expectedOutput := `Function Rules
	* Packages that match pattern 'foobar' should comply with the following rules:
		* Functions should not have more than 32 parameters
		* Functions should not have more than 3 return values
		* Files should not have more than 24 public functions
`

		describeFunctionRules(rules, outputBuffer)

		outputBytes, _ := io.ReadAll(outputBuffer)

		assert.Equal(t, expectedOutput, string(outputBytes), "Output doesn't match expected values.")
	})

	t.Run("function rules with blanks - case 2", func(t *testing.T) {
		outputBuffer := bytes.NewBufferString("")
		rules := []*config.FunctionsRule{
			{
				Package:                  "foobar",
				MaxLines:                 123,
				MaxPublicFunctionPerFile: 24,
				MaxReturnValues:          3,
			},
		}
		expectedOutput := `Function Rules
	* Packages that match pattern 'foobar' should comply with the following rules:
		* Functions should not have more than 123 lines
		* Functions should not have more than 3 return values
		* Files should not have more than 24 public functions
`

		describeFunctionRules(rules, outputBuffer)

		outputBytes, _ := io.ReadAll(outputBuffer)

		assert.Equal(t, expectedOutput, string(outputBytes), "Output doesn't match expected values.")
	})

	t.Run("function rules with blanks - case 3", func(t *testing.T) {
		outputBuffer := bytes.NewBufferString("")
		rules := []*config.FunctionsRule{
			{
				Package:         "foobar",
				MaxLines:        123,
				MaxParameters:   32,
				MaxReturnValues: 3,
			},
		}
		expectedOutput := `Function Rules
	* Packages that match pattern 'foobar' should comply with the following rules:
		* Functions should not have more than 123 lines
		* Functions should not have more than 32 parameters
		* Functions should not have more than 3 return values
`

		describeFunctionRules(rules, outputBuffer)

		outputBytes, _ := io.ReadAll(outputBuffer)

		assert.Equal(t, expectedOutput, string(outputBytes), "Output doesn't match expected values.")
	})

	t.Run("function rules with blanks - case 4", func(t *testing.T) {
		outputBuffer := bytes.NewBufferString("")
		rules := []*config.FunctionsRule{
			{
				Package:                  "foobar",
				MaxLines:                 123,
				MaxParameters:            32,
				MaxPublicFunctionPerFile: 24,
			},
		}
		expectedOutput := `Function Rules
	* Packages that match pattern 'foobar' should comply with the following rules:
		* Functions should not have more than 123 lines
		* Functions should not have more than 32 parameters
		* Files should not have more than 24 public functions
`

		describeFunctionRules(rules, outputBuffer)

		outputBytes, _ := io.ReadAll(outputBuffer)

		assert.Equal(t, expectedOutput, string(outputBytes), "Output doesn't match expected values.")
	})

	t.Run("empty rules", func(t *testing.T) {
		outputBuffer := bytes.NewBufferString("")
		var rules []*config.FunctionsRule
		expectedOutput := `Function Rules
	* No rules defined
`

		describeFunctionRules(rules, outputBuffer)

		outputBytes, _ := io.ReadAll(outputBuffer)

		assert.Equal(t, expectedOutput, string(outputBytes), "Output doesn't match expected values.")
	})
}