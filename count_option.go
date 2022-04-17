package git

type (
	countOption interface {
		applyCount(options *countOptions)
	}

	countOptions struct {
		typ countType
	}
)

func defaultCountOptions() *countOptions {
	return &countOptions{
		typ: countTypeTag,
	}
}

func newCountOptions(opts ...countOption) (_options *countOptions) {
	_options = defaultCountOptions()
	for _, opt := range opts {
		opt.applyCount(_options)
	}

	return
}
