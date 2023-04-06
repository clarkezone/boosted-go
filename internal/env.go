// Package internal contains environment variables
package internal

import (
	"errors"
	"fmt"
	"os"
	"path"

	clarkezoneLog "github.com/clarkezone/boosted-go/log"
)

const (
	// LogLevelVar is name of environment variable containing loglevel
	LogLevelVar     = "loglevel"
	defaultLogLevel = "Warn"

	// TargetRepoVar is name of environment variable containing target repo URL
	TargetRepoVar = "targetrepo"

	// LocalDirVar is name of environment variable container local repo path
	LocalDirVar = "localdir"

	// KubeConfigPathVar is name of environment variable for kube config path
	KubeConfigPathVar = "kubeconfigpath"

	// NamespaceVar is name of environment variable for kube namespace
	NamespaceVar = "namespace"

	// InitialCloneVar is name of environment variable for the initial clone flag
	InitialCloneVar = "initialclone"

	// InitialBuildVar is name of environment variable for the initial clone flag
	InitialBuildVar = "initialbuild"

	// WebhookListenVar is name of environment variable for the webhook listen flag
	WebhookListenVar = "webhooklisten"

	// InitialBranchVar is the name environment variable for the webhook listen flag
	InitialBranchVar = "initialbranch"

	// StartupMode is the name of the environment variable for the startup mode
	StartupMode = "startupmode"

	// ServiceURLVar is the name of the environment variable for the service url
	ServiceURLVar = "serviceurl"
)

var (
	// LogLevel is read from env
	LogLevel string

	// TargetRepo Url to target repo
	TargetRepo string

	// LocalDir absolute path to local dir
	LocalDir string

	// KubeConfigPath is the path to a valid KubeConfig file
	KubeConfigPath string

	// Namespace is the kubernetes namespace to create resources in
	Namespace string

	// InitialClone indicates if an initial clone should be performed at startup time
	InitialClone bool

	// InitialBuild indicates if the source should be built at startup time
	InitialBuild bool

	// WebhookListen indicates if the webhook should listener should be run at startup time
	WebhookListen bool

	// InitialBranch holds the branch that should be cloned on startup
	InitialBranch string

	// ServiceURL is the url of the service
	ServiceURL string
)

func init() {
}

func getDefaultKubeConfig() string {
	dirName, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	p := path.Join(dirName, ".kube/config")

	if _, err := os.Stat(p); errors.Is(err, os.ErrNotExist) {
		clarkezoneLog.Debugf("getdefaultKubeConfig(): default not detected")
		return ""
	}

	clarkezoneLog.Debugf("getDefaultKubeConfig(): found default kube config:%v", p)
	return p
}

// ValidateEnv validates environment variables
func ValidateEnv() error {
	clarkezoneLog.Debugf("ValidateEnv called")
	if Port == 0 {
		clarkezoneLog.Debugf("ValudateEnv() error port == 0")
		return fmt.Errorf("bad port")
	}
	if MetricsPort == 0 {
		clarkezoneLog.Debugf("ValudateEnv() error etricsport == 0")
		return fmt.Errorf("bad port")
	}
	if TargetRepo == "" {
		clarkezoneLog.Errorf("TargetRepo empty")
		return fmt.Errorf("TargetRepo empty")
	}
	if LocalDir == "" {
		clarkezoneLog.Errorf("LocalDir empty")
		return fmt.Errorf("LocalDir empty")
	}
	return nil
}
