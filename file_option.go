package git

var _ fileOption = (*file)(nil)

type (
	fileOption interface {
		applyFile(options *fileOptions)
	}

	fileOptions struct {
		filenames []string
	}

	file struct{}
)

func defaultFileOptions() *fileOptions {
	return &fileOptions{
		filenames: make([]string, 0),
	}
}

func newFileOptions(opts ...fileOption) (_options *fileOptions) {
	_options = defaultFileOptions()
	for _, opt := range opts {
		opt.applyFile(_options)
	}

	return
}

func (f *file) applyFile(_ *fileOptions) {}
