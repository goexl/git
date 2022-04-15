package git

var (
	_        = Dir
	_        = Folder
	_ option = (*optionDir)(nil)
)

type optionDir struct {
	tag
	file
	remote

	dir string
}

// Dir 目录
func Dir(dir string) *optionDir {
	return &optionDir{
		dir: dir,
	}
}

// Folder 目录
func Folder(folder string) *optionDir {
	return Dir(folder)
}

func (d *optionDir) apply(options *options) {
	options.dir = d.dir
}

func (d *optionDir) applyClean(options *cleanOptions) {
	options.dir = d.dir
}

func (d *optionDir) applySsh(options *sshOptions) {
	options.dir = d.dir
}
