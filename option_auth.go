package git

var (
	_            = Netrc
	_ authOption = (*optionAuth)(nil)
)

type optionAuth struct {
	writer authWriter
}

// Netrc 用户名密码
func Netrc(uri string, username string, password string) *optionAuth {
	return &optionAuth{
		writer: newNetrc(uri, username, password),
	}
}

func (a *optionAuth) applyAuth(options *authOptions) {
	options.writer = a.writer
}
