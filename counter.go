package git

type counter struct {
	mode counterMode
	args interface{}
}

func newCounter(mode counterMode, args interface{}) *counter {
	return &counter{
		mode: mode,
		args: args,
	}
}
