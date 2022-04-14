package git

var (
	_ = User
	_ = SafeDirectory
)

func User(author string, email string, opts ...option) (err error) {
	if err = Author(author, opts...); nil != err {
		return
	}
	err = Email(email, opts...)

	return
}

func Author(author string, opts ...option) error {
	return globalConfig([]interface{}{`user.name`, author}, opts...)
}

func Email(email string, opts ...option) error {
	return globalConfig([]interface{}{`user.email`, email}, opts...)
}

func SafeDirectory(dir string, opts ...option) error {
	return globalConfig([]interface{}{`--add`, `safe.directory`, dir}, opts...)
}

func globalConfig(args []interface{}, opts ...option) error {
	global := []interface{}{
		`config`,
		`--global`,
	}

	return Exec(append(opts, Args(append(global, args...)))...)
}
