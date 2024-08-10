package cmd

import (
	"errors"
	"fmt"
	"github.com/tsukinoko-kun/gourmet/internal/buildfile"
	"github.com/tsukinoko-kun/gourmet/internal/util"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Run build scripts followed by the tests",
	RunE: func(cmd *cobra.Command, args []string) error {
		buildConfig, err := buildfile.GetBuildConfig("test")
		if err != nil {
			return errors.Join(fmt.Errorf("failed to get build config"), err)
		}

		if len(args) == 0 {
			args = []string{"./..."}
		}

		if buildConfig.HasPreBuildCmd {
			fmt.Println("Executing prebuild cmd")
			if err := util.Run(filepath.Join(buildConfig.ModRoot, "cmd", "prebuild"), true); err != nil {
				os.Exit(2)
			}
		}

		if buildConfig.HasPreDoCmd {
			fmt.Println("Executing pretest cmd")
			if err := util.Run(filepath.Join(buildConfig.ModRoot, "cmd", "pretest"), true); err != nil {
				os.Exit(2)
			}
		}

		fmt.Println("Executing go test")
		if err := util.TestArgv(args); err != nil {
			os.Exit(2)
		}

		if buildConfig.HasPostDoCmd {
			fmt.Println("Executing posttest cmd")
			if err := util.Run(filepath.Join(buildConfig.ModRoot, "cmd", "posttest"), true); err != nil {
				os.Exit(2)
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}
