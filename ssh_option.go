package git

import (
	"os"
	"path/filepath"
)

type (
	sshOption interface {
		applySsh(options *sshOptions)
	}

	sshOptions struct {
		dir string
	}
)

func defaultSshOptions() *sshOptions {
	return &sshOptions{
		dir: filepath.Join(os.Getenv(homeEnv), sshHome),
	}
}

func newSshOptions(opts ...sshOption) (_options *sshOptions) {
	_options = defaultSshOptions()
	for _, opt := range opts {
		opt.applySsh(_options)
	}

	return
}
