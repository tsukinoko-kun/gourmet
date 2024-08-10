package buildfile

import (
	"errors"
	"github.com/tsukinoko-kun/gourmet/internal/util"
	"path/filepath"
)

type BildConfig struct {
	ModRoot      string
	HasPreBuild  bool
	HasBuildCmd  bool
	HasRunCmd    bool
	HasPostBuild bool
}

func findBuildFile(dir string) (*BildConfig, bool) {
	goModFile := filepath.Join(dir, "go.mod")
	if !util.Exists(goModFile) {
		return nil, false
	}

	config := &BildConfig{
		ModRoot: dir,
	}

	if util.Exists(filepath.Join(dir, "cmd", "prebuild")) {
		config.HasPreBuild = true
	}

	if util.Exists(filepath.Join(dir, "cmd", "build")) {
		config.HasBuildCmd = true
	}

	if util.Exists(filepath.Join(dir, "cmd", "run")) {
		config.HasRunCmd = true
	}

	if util.Exists(filepath.Join(dir, "cmd", "postbuild")) {
		config.HasPostBuild = true
	}

	return config, true
}

func GetBuildConfig() (*BildConfig, error) {
	buildConfig, err := util.Find(".", findBuildFile)
	if err != nil {
		return nil, errors.Join(errors.New("failed to find go module root"), err)
	}

	return buildConfig, nil
}
