package git

var (
	_            = Netrc
	_            = Ssh
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

// Ssh 安全连接
func Ssh(key string, opts ...sshOption) *optionAuth {
	return &optionAuth{
		writer: newSsh(newSshOptions(opts...).dir, key),
	}
}

func (a *optionAuth) applyAuth(options *authOptions) {
	options.writer = a.writer
}
