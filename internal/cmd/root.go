package cmd

import (
	"github.com/fdaines/arch-go/internal/common"
	"github.com/fdaines/arch-go/internal/impl"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:     "arch-go",
	Version: common.Version,
	Short:   "Architecture checks for Go",
	Long: `Architecture checks for Go:
* Dependencies
* Package contents
* Cyclic dependencies
* Function rules
* Naming rules`,
	Run: runCommand,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&common.Verbose, "verbose", "v", false, "Verbose Output")
	rootCmd.PersistentFlags().BoolVar(&common.Html, "html", false, "Generate HTML Report")
}

func runCommand(cmd *cobra.Command, args []string) {
	success := impl.CheckArchitecture()
	if !success {
		os.Exit(1)
	}
}
