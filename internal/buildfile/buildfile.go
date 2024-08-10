package buildfile

import (
	"errors"
	"github.com/tsukinoko-kun/gourmet/internal/util"
	"path/filepath"
)

type BildConfig struct {
	ModRoot string

	HasPreBuildCmd bool
	HasPreDoCmd    bool
	HasDoCmd       bool
	HasPostDoCmd   bool
}

func GetBuildConfig(do string) (*BildConfig, error) {
	findBuildFile := func(dir string) (*BildConfig, bool) {
		goModFile := filepath.Join(dir, "go.mod")
		if !util.Exists(goModFile) {
			return nil, false
		}

		config := &BildConfig{
			ModRoot: dir,
		}

		if util.Exists(filepath.Join(dir, "cmd", "prebuild")) {
			config.HasPreBuildCmd = true
		}

		if util.Exists(filepath.Join(dir, "cmd", "pre"+do)) {
			config.HasPreDoCmd = true
		}

		if util.Exists(filepath.Join(dir, "cmd", do)) {
			config.HasDoCmd = true
		}

		if util.Exists(filepath.Join(dir, "cmd", "post"+do)) {
			config.HasPostDoCmd = true
		}

		return config, true
	}

	buildConfig, err := util.Find(".", findBuildFile)
	if err != nil {
		return nil, errors.Join(errors.New("failed to find go module root"), err)
	}

	return buildConfig, nil
}
