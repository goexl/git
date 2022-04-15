package git

var (
	_ = Init
	_ = Commit
	_ = Add
)

func Init(opts ...option) error {
	return Exec(append(opts, Args(`init`))...)
}

func Commit(message string, opts ...fileOption) (err error) {
	_options := newFileOptions(opts...)
	if err = add(_options.filenames, opts...); nil != err {
		return
	}

	args := []interface{}{
		`commit`,
		`--message`, message,
		`--allow-empty`,
	}
	for _, filename := range _options.filenames {
		args = append(args, filename)
	}
	err = execAddons(args, opts)

	return
}

func Add(opts ...fileOption) error {
	return add(newFileOptions(opts...).filenames, opts...)
}

func add(filenames []string, opts ...fileOption) (err error) {
	for _, filename := range filenames {
		if err = execAddons([]interface{}{`add`, filename}, opts); nil != err {
			return
		}
	}

	return
}
