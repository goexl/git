package git

var (
	_        = Dir
	_ option = (*optionDir)(nil)
)

type optionDir struct {
	tag
	file

	dir string
}

// Dir 配置命令执行目录
func Dir(dir string) *optionDir {
	return &optionDir{
		dir: dir,
	}
}

func (d *optionDir) apply(options *options) {
	options.dir = d.dir
}
