package impl

import (
	"fmt"
	"github.com/fdaines/arch-go/internal/config"
	"github.com/fdaines/arch-go/internal/impl/contents"
	"github.com/fdaines/arch-go/internal/impl/cycles"
	"github.com/fdaines/arch-go/internal/impl/dependencies"
	"github.com/fdaines/arch-go/internal/impl/functions"
	"github.com/fdaines/arch-go/internal/impl/model"
	"github.com/fdaines/arch-go/internal/impl/naming"
	baseModel "github.com/fdaines/arch-go/internal/model"
	"github.com/fdaines/arch-go/internal/model/result"
	"github.com/fdaines/arch-go/internal/utils"
	"github.com/fdaines/arch-go/internal/utils/packages"
	"github.com/fdaines/arch-go/internal/utils/text"
	"os"
	"regexp"
)

func CheckArchitecture() bool {
	returnValue := true
	utils.ExecuteWithTimer(func() {
		configuration, err := config.LoadConfig("arch-go.yml")
		if err != nil {
			fmt.Printf("Error: %+v\n", err)
			os.Exit(1)
		} else {
			mainPackage, _ := packages.GetMainPackage()
			pkgs, _ := packages.GetBasicPackagesInfo()
			moduleInfo := &baseModel.ModuleInfo{
				MainPackage: mainPackage,
				Packages:    pkgs,
			}

			verifications := resolveVerifications(configuration, moduleInfo)
			summary := result.ResultSummary{
				Rules: len(verifications),
			}
			for _, v := range verifications {
				result := v.Verify()
				if result {
					summary.Succeeded++
				} else {
					summary.Failed++
				}
			}
			for _, v := range verifications {
				v.PrintResults()
			}
			summary.Print()
			returnValue = summary.Status()
		}
	})
	return returnValue
}

func resolveVerifications(configuration *config.Config, moduleInfo *baseModel.ModuleInfo) []model.RuleVerification {
	var verifications []model.RuleVerification
	verifications = resolveDependencyRules(configuration, moduleInfo, verifications)
	verifications = resolveFunctionRules(configuration, moduleInfo, verifications)
	verifications = resolveContentRules(configuration, moduleInfo, verifications)
	verifications = resolveCycleRules(configuration, moduleInfo, verifications)
	verifications = resolveNamingRules(configuration, moduleInfo, verifications)

	return verifications
}

func resolveCycleRules(configuration *config.Config, moduleInfo *baseModel.ModuleInfo, verifications []model.RuleVerification) []model.RuleVerification {
	for _, cycleRule := range configuration.CyclesRules {
		verificationInstance := cycles.NewCyclesRuleVerification(moduleInfo.MainPackage, moduleInfo.Packages, cycleRule)
		packageRegExp, _ := regexp.Compile(text.PreparePackageRegexp(cycleRule.Package))
		for _, pkg := range moduleInfo.Packages {
			if packageRegExp.MatchString(pkg.Path) {
				verificationInstance.PackageDetails = append(verificationInstance.PackageDetails, model.PackageVerification{
					Package: pkg,
					Passes:  false,
				})
			}
		}
		verifications = append(verifications, verificationInstance)
	}
	return verifications
}

func resolveContentRules(configuration *config.Config, moduleInfo *baseModel.ModuleInfo, verifications []model.RuleVerification) []model.RuleVerification {
	for _, contentRule := range configuration.ContentRules {
		verificationInstance := contents.NewContentsRuleVerification(moduleInfo.MainPackage, contentRule)
		packageRegExp, _ := regexp.Compile(text.PreparePackageRegexp(contentRule.Package))
		for _, pkg := range moduleInfo.Packages {
			if packageRegExp.MatchString(pkg.Path) {
				verificationInstance.PackageDetails = append(verificationInstance.PackageDetails, model.PackageVerification{
					Package: pkg,
					Passes:  false,
				})
			}
		}
		verifications = append(verifications, verificationInstance)
	}
	return verifications
}

func resolveFunctionRules(configuration *config.Config, moduleInfo *baseModel.ModuleInfo, verifications []model.RuleVerification) []model.RuleVerification {
	for _, functionRule := range configuration.FunctionsRules {
		verificationInstance := functions.NewFunctionsRuleVerification(moduleInfo.MainPackage, functionRule)
		packageRegExp, _ := regexp.Compile(text.PreparePackageRegexp(functionRule.Package))
		for _, pkg := range moduleInfo.Packages {
			if packageRegExp.MatchString(pkg.Path) {
				verificationInstance.PackageDetails = append(verificationInstance.PackageDetails, model.PackageVerification{
					Package: pkg,
					Passes:  false,
				})
			}
		}
		verifications = append(verifications, verificationInstance)
	}
	return verifications
}

func resolveDependencyRules(configuration *config.Config, moduleInfo *baseModel.ModuleInfo, verifications []model.RuleVerification) []model.RuleVerification {
	for _, dependencyRule := range configuration.DependenciesRules {
		verificationInstance := dependencies.NewDependencyRuleVerification(moduleInfo.MainPackage, dependencyRule)
		packageRegExp, _ := regexp.Compile(text.PreparePackageRegexp(dependencyRule.Package))
		for _, pkg := range moduleInfo.Packages {
			if packageRegExp.MatchString(pkg.Path) {
				verificationInstance.PackageDetails = append(verificationInstance.PackageDetails, model.PackageVerification{
					Package: pkg,
					Passes:  false,
				})
			}
		}
		verifications = append(verifications, verificationInstance)
	}
	return verifications
}

func resolveNamingRules(configuration *config.Config, moduleInfo *baseModel.ModuleInfo, verifications []model.RuleVerification) []model.RuleVerification {
	for _, namingRule := range configuration.NamingRules {
		verificationInstance := naming.NewNamingRuleVerification(moduleInfo.MainPackage, namingRule)
		packageRegExp, _ := regexp.Compile(text.PreparePackageRegexp(namingRule.Package))
		for _, pkg := range moduleInfo.Packages {
			if packageRegExp.MatchString(pkg.Path) {
				verificationInstance.PackageDetails = append(verificationInstance.PackageDetails, model.PackageVerification{
					Package: pkg,
					Passes:  false,
				})
			}
		}
		verifications = append(verifications, verificationInstance)
	}
	return verifications
}
