package git

type (
	cleanOption interface {
		applyClean(options *cleanOptions)
	}

	cleanOptions struct {
		dir string
	}
)

func defaultCleanOptions() *cleanOptions {
	return &cleanOptions{
		dir: currentDir,
	}
}

func newCleanOptions(opts ...cleanOption) (_options *cleanOptions) {
	_options = defaultCleanOptions()
	for _, opt := range opts {
		opt.applyClean(_options)
	}

	return
}
