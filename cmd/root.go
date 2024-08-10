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

var (
	rootCmd = &cobra.Command{
		Use:   "gourmet",
		Short: "Write build scripts for your Go projects in Go",
		RunE: func(cmd *cobra.Command, args []string) error {
			buildConfig, err := buildfile.GetBuildConfig()
			if err != nil {
				return errors.Join(fmt.Errorf("failed to get build config"), err)
			}

			if len(args) == 0 {
				args = []string{buildConfig.ModRoot}
			}

			if buildConfig.HasPreBuild {
				fmt.Println("Executing prebuild cmd")
				if err := util.Run(filepath.Join(buildConfig.ModRoot, "cmd", "prebuild"), true); err != nil {
					os.Exit(2)
				}
			}

			if buildConfig.HasBuildCmd {
				fmt.Println("Executing build cmd")
				if err := util.Run(filepath.Join(buildConfig.ModRoot, "cmd", "build"), true); err != nil {
					os.Exit(2)
				}
			} else {
				fmt.Println("Executing go build")
				if err := util.BuildArgv(args); err != nil {
					os.Exit(2)
				}
			}

			if buildConfig.HasPostBuild {
				fmt.Println("Executing postbuild cmd")
				if err := util.Run(filepath.Join(buildConfig.ModRoot, "cmd", "postbuild"), true); err != nil {
					os.Exit(2)
				}
			}

			return nil
		},
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
