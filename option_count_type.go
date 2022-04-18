package git

var (
	_             = Tagging
	_ countOption = (*optionCountType)(nil)
)

type optionCountType struct {
	typ countType
}

// Tagging 标签
func Tagging() *optionCountType {
	return &optionCountType{
		typ: countTypeTag,
	}
}

func (ct *optionCountType) applyCount(options *countOptions) {
	options.typ = ct.typ
}
