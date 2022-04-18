package git

var (
	_        = Exe
	_        = Bin
	_ option = (*optionExe)(nil)
)

type optionExe struct {
	exe string
}

// Exe 执行程序
func Exe(exe string) *optionExe {
	return &optionExe{
		exe: exe,
	}
}

// Bin 执行程序
func Bin(bin string) *optionExe {
	return Exe(bin)
}

func (e *optionExe) apply(options *options) {
	options.exe = e.exe
}
