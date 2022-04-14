package git

var (
	_        = String
	_ option = (*optionString)(nil)
)

type optionString struct {
	tag
	file

	output *string
}

// String 字符串输出
func String(output *string) *optionString {
	return &optionString{
		output: output,
	}
}

func (s *optionString) apply(options *options) {
	options.collectors = append(options.collectors, newCollector(collectorModeString, s.output))
}
