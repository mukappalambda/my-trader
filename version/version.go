package version

import (
	"fmt"
	"os"
	"runtime"
	"text/tabwriter"
)

var (
	GitCommit  = "unknown"
	GitVersion = "unknown"
	BuildDate  = "unknown"
	RepoUrl    = "unknown"
)

func Version(name string) {
	fmt.Printf("%s%s\n\n", name, RepoUrl)

	type KV struct {
		K string
		V string
	}
	kvs := []KV{
		{K: "GitVersion", V: GitVersion},
		{K: "GitCommit", V: GitCommit},
		{K: "BuildDate", V: BuildDate},
		{K: "GoVersion", V: runtime.Version()},
		{K: "Compiler", V: runtime.Compiler},
		{K: "Platform", V: fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)},
	}

	tw := tabwriter.NewWriter(os.Stdout, 1, 8, 1, ' ', 0)
	defer tw.Flush()
	for _, kv := range kvs {
		fmt.Fprintf(tw, "%s:\t%s\n", kv.K, kv.V)
	}
}
