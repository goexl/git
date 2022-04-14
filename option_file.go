package git

var (
	_            = Filenames
	_ fileOption = (*optionFile)(nil)
)

type optionFile struct {
	filenames []string
}

// Filenames 配置命令执行目录
func Filenames(filenames ...string) *optionFile {
	return &optionFile{
		filenames: filenames,
	}
}

func (f *optionFile) applyFile(options *fileOptions) {
	options.filenames = append(options.filenames, f.filenames...)
}
