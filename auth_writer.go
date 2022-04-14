package git

type authWriter interface {
	write() error
}
