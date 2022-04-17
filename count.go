package git

var _ = Count

func Count(opts ...countOption) (total int64, err error) {
	_options := newCountOptions(opts...)
	var args []interface{}
	switch _options.typ {
	case countTypeTag:
		args = []interface{}{`tag`, `--list`}
	}
	err = execAddons(args, opts, Line(&total))

	return
}
