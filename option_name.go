package git

var (
	_              = Name
	_ remoteOption = (*optionName)(nil)
	_ pushOption   = (*optionName)(nil)
)

type optionName struct {
	name string
}

// Name 名称
func Name(name string) *optionName {
	return &optionName{
		name: name,
	}
}

func (n *optionName) applyRemote(options *remoteOptions) {
	options.name = n.name
}

func (n *optionName) applyPush(options *pushOptions) {
	options.name = n.name
}
