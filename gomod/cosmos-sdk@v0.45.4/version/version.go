



//

//



//





package version

import (
	"encoding/json"
	"fmt"
	"runtime"
	"runtime/debug"
)

var (

	Name = ""

	AppName = "<appd>"

	Version = ""

	Commit = ""

	BuildTags = ""
)

func getSDKVersion() string {
	deps, ok := debug.ReadBuildInfo()
	if !ok {
		return "unable to read deps"
	}
	var sdkVersion string
	for _, dep := range deps.Deps {
		if dep.Path == "github.com/cosmos/cosmos-sdk" {
			sdkVersion = dep.Version
		}
	}

	return sdkVersion
}


type Info struct {
	Name             string     `json:"name" yaml:"name"`
	AppName          string     `json:"server_name" yaml:"server_name"`
	Version          string     `json:"version" yaml:"version"`
	GitCommit        string     `json:"commit" yaml:"commit"`
	BuildTags        string     `json:"build_tags" yaml:"build_tags"`
	GoVersion        string     `json:"go" yaml:"go"`
	BuildDeps        []buildDep `json:"build_deps" yaml:"build_deps"`
	CosmosSdkVersion string     `json:"cosmos_sdk_version" yaml:"cosmos_sdk_version"`
}

func NewInfo() Info {
	sdkVersion := getSDKVersion()
	return Info{
		Name:             Name,
		AppName:          AppName,
		Version:          Version,
		GitCommit:        Commit,
		BuildTags:        BuildTags,
		GoVersion:        fmt.Sprintf("go version %s %s/%s", runtime.Version(), runtime.GOOS, runtime.GOARCH),
		BuildDeps:        depsFromBuildInfo(),
		CosmosSdkVersion: sdkVersion,
	}
}

func (vi Info) String() string {
	return fmt.Sprintf(`%s: %s
git commit: %s
build tags: %s
%s`,
		vi.Name, vi.Version, vi.GitCommit, vi.BuildTags, vi.GoVersion,
	)
}

func depsFromBuildInfo() (deps []buildDep) {
	buildInfo, ok := debug.ReadBuildInfo()
	if !ok {
		return nil
	}

	for _, dep := range buildInfo.Deps {
		deps = append(deps, buildDep{dep})
	}

	return
}

type buildDep struct {
	*debug.Module
}

func (d buildDep) String() string {
	if d.Replace != nil {
		return fmt.Sprintf("%s@%s => %s@%s", d.Path, d.Version, d.Replace.Path, d.Replace.Version)
	}

	return fmt.Sprintf("%s@%s", d.Path, d.Version)
}

func (d buildDep) MarshalJSON() ([]byte, error)      { return json.Marshal(d.String()) }
func (d buildDep) MarshalYAML() (interface{}, error) { return d.String(), nil }
