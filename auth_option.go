package git

type (
	authOption interface {
		applyAuth(options *authOptions)
	}

	authOptions struct {
		writer authWriter
	}
)

func defaultAuthOptions() *authOptions {
	return &authOptions{}
}

func newAuthOptions(opts ...authOption) (_options *authOptions) {
	_options = defaultAuthOptions()
	for _, opt := range opts {
		opt.applyAuth(_options)
	}

	return
}
