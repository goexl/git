package git

var _ = User

func User(author string, email string) (err error) {
	if err = Author(author); nil != err {
		return
	}
	err = Email(email)

	return
}

func Author(author string) error {
	return Exec(Args(`config`, `--global`, `user.name`, author))
}

func Email(email string) error {
	return Exec(Args(`config`, `--global`, `user.email`, email))
}
