package git

import (
	"io/fs"
	"io/ioutil"
	"os"
	"path"
)

var (
	_ = Init
	_ = Clean
)

func Init(opts ...option) error {
	return Exec(append(opts, Args(`init`))...)
}

func Clean(opts ...cleanOption) (err error) {
	_options := newCleanOptions(opts...)

	var fis []fs.FileInfo
	if fis, err = ioutil.ReadDir(_options.dir); nil != err {
		return
	}

	// 删除所有
	for _, fi := range fis {
		if err = os.RemoveAll(path.Join(_options.dir, fi.Name())); nil != err {
			return
		}
	}

	return
}
