package git

var _ cleanOption = (*clean)(nil)

type (
	cleanOption interface {
		applyClean(options *cleanOptions)
	}

	cleanOptions struct {
		dir string
	}

	clean struct{}
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

func (c *clean) applyClean(_ *cleanOptions) {}
