package git

var _ = Auth

func Auth(opts ...authOption) error {
	return newAuthOptions(opts...).writer.write()
}
