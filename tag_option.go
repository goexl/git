package git

var _ tagOption = (*tag)(nil)

type (
	tagOption interface {
		applyTag(options *tagOptions)
	}

	tagOptions struct {
		skip int
	}

	tag struct{}
)

func defaultTagOptions() *tagOptions {
	return &tagOptions{
		skip: 0,
	}
}

func newTagOptions(opts ...tagOption) (_options *tagOptions) {
	_options = defaultTagOptions()
	for _, opt := range opts {
		opt.applyTag(_options)
	}

	return
}

func (t tag) applyTag(_ *tagOptions) {}
