package git

var (
	_        = Args
	_ option = (*optionArgs)(nil)
)

type optionArgs struct {
	args []interface{}
}

// Args 参数
func Args(args ...interface{}) *optionArgs {
	return &optionArgs{
		args: args,
	}
}

func (a *optionArgs) apply(options *options) {
	options.args = a.args
}
