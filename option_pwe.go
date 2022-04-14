package git

var (
	_        = Pwe
	_        = Sync
	_        = DisablePwe
	_ option = (*optionPwe)(nil)
)

type optionPwe struct {
	tag
	file

	pwe bool
}

// Pwe 启用PWE
func Pwe() *optionPwe {
	return &optionPwe{
		pwe: true,
	}
}

// DisablePwe 禁用PWE
func DisablePwe() *optionPwe {
	return &optionPwe{
		pwe: false,
	}
}

func (p *optionPwe) apply(options *options) {
	options.pwe = p.pwe
}
