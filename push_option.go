package git

var _ pushOption = (*push)(nil)

type (
	pushOption interface {
		applyPush(options *pushOptions)
	}

	pushOptions struct {
		name   string
		branch string
	}

	push struct{}
)

func defaultPushOptions() *pushOptions {
	return &pushOptions{
		name:   defaultRemoteName,
		branch: defaultBranch,
	}
}

func newPushOptions(opts ...pushOption) (_options *pushOptions) {
	_options = defaultPushOptions()
	for _, opt := range opts {
		opt.applyPush(_options)
	}

	return
}

func (r push) applyPush(_ *pushOptions) {}
