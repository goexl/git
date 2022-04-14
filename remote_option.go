package git

var _ remoteOption = (*remote)(nil)

type (
	remoteOption interface {
		applyRemote(options *remoteOptions)
	}

	remoteOptions struct {
		name string
	}

	remote struct{}
)

func defaultRemoteOptions() *remoteOptions {
	return &remoteOptions{
		name: defaultRemoteName,
	}
}

func newRemoteOptions(opts ...remoteOption) (_options *remoteOptions) {
	_options = defaultRemoteOptions()
	for _, opt := range opts {
		opt.applyRemote(_options)
	}

	return
}

func (r remote) applyRemote(_ *remoteOptions) {}
