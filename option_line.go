package git

var (
	_        = Line
	_ option = (*optionLine)(nil)
)

type optionLine struct {
	line *int64
}

// Line 行
func Line(line *int64) *optionLine {
	return &optionLine{
		line: line,
	}
}

func (c *optionLine) apply(options *options) {
	options.counters = append(options.counters, newCounter(counterModeLine, c.line))
}
