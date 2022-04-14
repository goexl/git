package git

import (
	"fmt"
)

var _ = Tag

func Tag(opts ...tagOption) (tag string, err error) {
	_options := defaultTagOptions()
	for _, opt := range opts {
		opt.applyTag(_options)
	}

	// 取得最后的提交
	// 命令：git rev-list --tags --skip=1 --max-count=1
	var latest string
	args := []interface{}{
		`rev-list`,
		`--tags`,
		fmt.Sprintf(`--skip=%d`, _options.skip),
		`--max-count=1`,
	}
	if err = execAddons(args, opts, String(&latest), DisablePwe()); nil != err || `` == latest {
		return
	}

	// 取得标签
	args = []interface{}{
		`describe`,
		`--abbrev=0`,
		`--tags`,
		latest,
	}
	if err = execAddons(args, opts, String(&latest), DisablePwe()); nil == err {
		tag = latest
	}

	return
}
