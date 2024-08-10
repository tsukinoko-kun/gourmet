package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tsukinoko-kun/gourmet/internal/buildfile"
	"github.com/tsukinoko-kun/gourmet/internal/util"
	"os"
	"path/filepath"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run build scripts followed by the application",
	RunE: func(cmd *cobra.Command, args []string) error {
		buildConfig, err := buildfile.GetBuildConfig("run")
		if err != nil {
			return errors.Join(fmt.Errorf("failed to get build config"), err)
		}

		if len(args) == 0 {
			args = []string{buildConfig.ModRoot}
		}

		if buildConfig.HasPreBuildCmd {
			fmt.Println("Executing prebuild cmd")
			if err := util.Run(filepath.Join(buildConfig.ModRoot, "cmd", "prebuild"), true); err != nil {
				os.Exit(2)
			}
		}

		if buildConfig.HasPreDoCmd {
			fmt.Println("Executing prerun cmd")
			if err := util.Run(filepath.Join(buildConfig.ModRoot, "cmd", "prerun"), true); err != nil {
				os.Exit(2)
			}
		}

		if buildConfig.HasDoCmd {
			fmt.Println("Executing run cmd")
			if err := util.Run(filepath.Join(buildConfig.ModRoot, "cmd", "run"), true); err != nil {
				os.Exit(2)
			}
		} else {
			fmt.Println("Executing go run")
			if err := util.RunArgv(args); err != nil {
				os.Exit(2)
			}
		}

		if buildConfig.HasPostDoCmd {
			fmt.Println("Executing postrun cmd")
			if err := util.Run(filepath.Join(buildConfig.ModRoot, "cmd", "postrun"), true); err != nil {
				os.Exit(2)
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
