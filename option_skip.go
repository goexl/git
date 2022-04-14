package git

var (
	_           = Skip
	_           = Sync
	_ tagOption = (*optionSkip)(nil)
)

type optionSkip struct {
	skip int
}

// Skip 跳跃数
func Skip(skip int) *optionSkip {
	return &optionSkip{
		skip: skip,
	}
}

func (s *optionSkip) applyTag(options *tagOptions) {
	options.skip = s.skip
}
