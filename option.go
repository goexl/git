package git

type (
	option interface {
		apply(options *options)
	}

	options struct {
		exe          string
		args         []interface{}
		environments []string
		dir          string
		pwe          bool
		async        bool

		collectors []*collector
		checkers   []*checker
		counters   []*counter

		debug   bool
		verbose bool
	}
)

func defaultOptions() *options {
	return &options{
		exe: exe,
	}
}

func newOptions(opts ...option) (_options *options) {
	_options = defaultOptions()
	for _, opt := range opts {
		opt.apply(_options)
	}

	return
}
