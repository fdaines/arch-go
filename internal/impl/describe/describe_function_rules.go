package describe

import (
	"fmt"
	"github.com/fdaines/arch-go/internal/common"
	"github.com/fdaines/arch-go/internal/config"
	"io"
)

func describeFunctionRules(rules []*config.FunctionsRule, out io.Writer) {
	fmt.Fprintf(out, "Function Rules\n")
	if len(rules) == 0 {
		fmt.Fprintf(out, common.NoRulesDefined)
		return
	}
	for _,r := range rules {
		fmt.Fprintf(out, "\t* Packages that match pattern '%s' should comply with the following rules:\n", r.Package)
		if r.MaxLines > 0 {
			fmt.Fprintf(out, "\t\t* Functions should not have more than %d lines\n", r.MaxLines)
		}
		if r.MaxParameters > 0 {
			fmt.Fprintf(out, "\t\t* Functions should not have more than %d parameters\n", r.MaxParameters)
		}
		if r.MaxReturnValues > 0 {
			fmt.Fprintf(out, "\t\t* Functions should not have more than %d return values\n", r.MaxReturnValues)
		}
		if r.MaxPublicFunctionPerFile > 0 {
			fmt.Fprintf(out, "\t\t* Files should not have more than %d public functions\n", r.MaxPublicFunctionPerFile)
		}
	}
	fmt.Println()
}