package version

var (
	// Version should be updated by hand at each release
	Version = "0.1.1"

	// GITCOMMIT will be overwritten automatically by the build system
	GitCommit = "HEAD"
)

func FullVersion() string {
	return Version + " (" + GitCommit + ")"
}
