package git

var _ = User

func User(author string, email string, opts ...option) (err error) {
	if err = Author(author, opts...); nil != err {
		return
	}
	err = Email(email, opts...)

	return
}

func Author(author string, opts ...option) error {
	return Exec(append(opts, Args(`config`, `--global`, `user.name`, author))...)
}

func Email(email string, opts ...option) error {
	return Exec(append(opts, Args(`config`, `--global`, `user.email`, email))...)
}
