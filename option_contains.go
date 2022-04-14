package git

var (
	_        = Contains
	_ option = (*optionContains)(nil)
)

type optionContains struct {
	contains string
}

// Contains 检查是否包含字符串
func Contains(contains string) *optionContains {
	return &optionContains{
		contains: contains,
	}
}

func (c *optionContains) apply(options *options) {
	options.checkers = append(options.checkers, newChecker(checkerModeContains, c.contains))
}
