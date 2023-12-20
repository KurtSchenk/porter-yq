//go:build mage

package main

import (
	"fmt"
	"os"

	"get.porter.sh/magefiles/mixins"
	"get.porter.sh/magefiles/releases"
	"github.com/carolynvs/magex/shx"
	"github.com/magefile/mage/sh"
)

const (
	mixinName    = "yq"
	mixinPackage = "github.com/getporter/yq"
	mixinBin     = "bin/mixins/" + mixinName
)

var magefile = mixins.NewMagefile(mixinPackage, mixinName, mixinBin)

var must = shx.CommandBuilder{StopOnError: true}

// ConfigureAgent sets up the CI server with mage and GO
func ConfigureAgent() {
	magefile.ConfigureAgent()
}

// Build the mixin
func Build() {
	magefile.Build()
}

// XBuildAll cross-compiles the mixin before a release
func XBuildAll() {
	magefile.XBuildAll()
}

// TestUnit runs the unit tests
func TestUnit() {
	magefile.TestUnit()
}

// Test runs all types of tests
func Test() {
	magefile.Test()
}

func GitVersion() (string, error) {
	var version string
	out, err := sh.Output("git", "describe", "--tags")
	if err != nil {
		return "", fmt.Errorf("failed to run git describe: %w", err)
	}
	version = out
	return version, nil
}

func ClientPlatform() (string, error) {
	var platform string
	out, err := sh.Output("go", "env", "GOOS")
	if err != nil {
		return "", fmt.Errorf("failed to run go env GOOS: %w", err)
	}
	platform = out
	return platform, nil
}

func ClientArch() (string, error) {
	var arch string
	out, err := sh.Output("go", "env", "GOARCH")
	if err != nil {
		return "", fmt.Errorf("failed to run go env GOARCH: %w", err)
	}
	arch = out
	return arch, nil
}

func TestIntegration() {
	//	cp $(BINDIR)/$(VERSION)/$(MIXIN)-$(CLIENT_PLATFORM)-$(CLIENT_ARCH)$(FILE_EXT) $(BINDIR)/$(MIXIN)$(FILE_EXT)
	//GO111MODULE=on go test -tags=integration ./tests/...

	/* TODO: TestIntegration now yet working like it does in Makefile
	bindir := mixinBin
	version, err := GitVersion()
	mixin := mixinName
	client_platform, err := ClientPlatform()
	client_arch, err := ClientArch()
	runtime_platform := "linux"
	file_ext := "" // .exe for windows runtime platform
	*/
	must.RunV("go", "test", "-tags=integration", "./tests/...")
}

// Publish the mixin to GitHub
func Publish() {
	// You can test out publishing locally by overriding PORTER_RELEASE_REPOSITORY and PORTER_PACKAGES_REMOTE
	if _, overridden := os.LookupEnv(releases.ReleaseRepository); !overridden {
		os.Setenv(releases.ReleaseRepository, "github.com/Kurt Schenk/YOURREPO")
	}
	magefile.PublishBinaries()

	// TODO: uncomment out the lines below to publish a mixin feed
	// Set PORTER_PACKAGES_REMOTE to a repository that will contain your mixin feed, similar to github.com/getporter/packages
	//if _, overridden := os.LookupEnv(releases.PackagesRemote); !overridden {
	//	os.Setenv("PORTER_PACKAGES_REMOTE", "git@github.com:Kurt Schenk/YOUR_PACKAGES_REPOSITORY")
	//}
	//magefile.PublishMixinFeed()
}

// TestPublish publishes the project to the specified GitHub username.
// If your mixin is official hosted in a repository under your username, you will need to manually
// override PORTER_RELEASE_REPOSITORY and PORTER_PACKAGES_REMOTE to test out publishing safely.
func TestPublish(username string) {
	magefile.TestPublish(username)
}

// Install the mixin
func Install() {
	magefile.Install()
}

// Clean removes generated build files
func Clean() {
	magefile.Clean()
}
