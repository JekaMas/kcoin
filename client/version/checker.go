package version

import (
	"github.com/blang/semver"
	"github.com/kowala-tech/kcoin/client/log"
	"runtime"
	"time"

	"github.com/kowala-tech/kcoin/client/params"
)

func Checker(repository string) {
	current, err := MakeSemver(params.Version)
	if err != nil {
		log.Error("error parsing current version, exiting checker", "err", err)
		return
	}

	checker := &checker{
		repository:    repository,
		current:       current,
		checkInterval: time.Minute,
	}
	go checker.check()
}

type checker struct {
	repository    string
	current       semver.Version
	latest        semver.Version
	checkInterval time.Duration
}

func (c *checker) check() {
	for {
		time.Sleep(c.checkInterval)
		if c.isNewVersionAvailable() {
			c.printNewVersionAvailable()
			c.checkInterval = c.checkInterval * 2
		}
	}
}

func (c *checker) isNewVersionAvailable() bool {
	log.Debug("Checking repository " + c.repository)

	latest, err := c.getLatest()
	if err != nil {
		log.Error("error checking latest version available", "err", err)
		return false
	}

	c.latest = latest

	return latest.GT(c.current)
}

func (c *checker) getLatest() (semver.Version, error) {
	finder := NewFinder(NewS3AssetRepository(c.repository))
	latest, err := finder.Latest(runtime.GOOS, runtime.GOARCH)
	if err != nil {
		return semver.Version{}, err
	}
	return latest.Semver(), nil
}

func (c *checker) printNewVersionAvailable() {
	log.Warn("New version is available: " + c.latest.String())
}
