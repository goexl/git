package git

var (
	_           = Skip
	_           = Latest
	_           = Second
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

// Latest 最后
func Latest() *optionSkip {
	return Skip(0)
}

// Second 第二个
func Second() *optionSkip {
	return Skip(1)
}

func (s *optionSkip) applyTag(options *tagOptions) {
	options.skip = s.skip
}
