package git

type (
	option interface {
		apply(options *options)
	}

	options struct {
		exe          string
		name         string
		args         []interface{}
		environments []string
		dir          string
		pwe          bool
		async        bool

		collectors []*collector
		checkers   []*checker

		debug   bool
		verbose bool
	}
)

func defaultOptions() *options {
	return &options{
		exe: exe,
	}
}
