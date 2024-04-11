package functions

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/fdaines/arch-go/internal/model"
	"github.com/fdaines/arch-go/internal/utils/text"
	"github.com/fdaines/arch-go/pkg/config"
)

func CheckRules(moduleInfo model.ModuleInfo, functionRules []*config.FunctionsRule) RulesResult {
	result := RulesResult{
		Passes: true,
	}

	for _, it := range functionRules {
		result.Results = append(result.Results, CheckRule(moduleInfo, *it))
	}

	// Update result.Passes based on each rule result
	for _, r := range result.Results {
		result.Passes = result.Passes && r.Passes
	}

	return result
}

func CheckRule(moduleInfo model.ModuleInfo, functionRule config.FunctionsRule) *RuleResult {
	result := &RuleResult{
		Rule:        functionRule,
		Description: resolveDescription(functionRule),
		Passes:      true,
	}

	packageRegExp, _ := regexp.Compile(text.PreparePackageRegexp(functionRule.Package))
	for _, it := range moduleInfo.Packages {
		if it != nil && packageRegExp.MatchString(it.Path) {
			functions, _ := RetrieveFunctions(it, moduleInfo.MainPackage)
			pass, details := checkFunctionRule(functions, functionRule)
			result.Passes = result.Passes && pass
			result.Verifications = append(
				result.Verifications,
				Verification{
					Package: it.Path,
					Passes:  pass,
					Details: details,
				},
			)
		}
	}

	return result
}

func checkFunctionRule(functions []*FunctionDetails, functionRule config.FunctionsRule) (bool, []string) {
	pass1, details1 := checkMaxLines(functions, functionRule.MaxLines)
	pass2, details2 := checkMaxParameters(functions, functionRule.MaxParameters)
	pass3, details3 := checkMaxReturnValues(functions, functionRule.MaxReturnValues)
	pass4, details4 := checkMaxPublicFunctions(functions, functionRule.MaxPublicFunctionPerFile)

	return pass2 && pass1 && pass3 && pass4,
		append(details1, append(details2, append(details3, details4...)...)...)
}

func resolveDescription(r config.FunctionsRule) string {
	var ruleDescriptions []string
	if r.MaxParameters != nil {
		ruleDescriptions = append(ruleDescriptions, fmt.Sprintf("'at most %d parameters'", *r.MaxParameters))
	}
	if r.MaxReturnValues != nil {
		ruleDescriptions = append(ruleDescriptions, fmt.Sprintf("'at most %d return values'", *r.MaxReturnValues))
	}
	if r.MaxLines != nil {
		ruleDescriptions = append(ruleDescriptions, fmt.Sprintf("'at most %d lines'", *r.MaxLines))
	}
	if r.MaxPublicFunctionPerFile != nil {
		ruleDescriptions = append(ruleDescriptions, fmt.Sprintf("'no more than %d public functions per file'", *r.MaxPublicFunctionPerFile))
	}
	return fmt.Sprintf(
		"Functions in packages matching pattern '%s' should have [%s]",
		r.Package,
		strings.Join(ruleDescriptions, ","),
	)
}
