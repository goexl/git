package git

var _ = Push

func Push(opts ...pushOption) error {
	_options := newPushOptions(opts...)
	args := []interface{}{
		`push`,
		`--set-upstream`,
		_options.name,
		_options.branch,
	}

	return execAddons(args, opts)
}
