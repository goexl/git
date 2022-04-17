package git

import (
	"reflect"

	"github.com/goexl/gex"
)

var _ = Exec

func Exec(opts ...option) error {
	_options := defaultOptions()
	for _, opt := range opts {
		opt.apply(_options)
	}

	return exec(_options)
}

func execAddons(args []interface{}, opts interface{}, addons ...option) error {
	_options := defaultOptions()
	// 处理用户传进来的选项
	switch reflect.TypeOf(opts).Kind() {
	case reflect.Slice:
		slice := reflect.ValueOf(opts)
		for index := 0; index < slice.Len(); index++ {
			if _option, ok := slice.Index(index).Interface().(option); ok {
				_option.apply(_options)
			}
		}
	}

	// 处理附加选项
	for _, addon := range addons {
		addon.apply(_options)
	}

	// 处理参数
	_options.args = args

	return exec(_options)
}

func exec(options *options) (err error) {
	gexOptions := gex.NewOptions(gex.Args(options.args...))
	if `` != options.dir {
		gexOptions = append(gexOptions, gex.Dir(options.dir))
	}

	if 0 != len(options.environments) {
		gexOptions = append(gexOptions, gex.StringEnvs(options.environments...))
	}

	if options.async {
		gexOptions = append(gexOptions, gex.Async())
	} else {
		gexOptions = append(gexOptions, gex.Sync())
	}

	// 检查器
	for _, _checker := range options.checkers {
		switch _checker.mode {
		case checkerModeContains:
			gexOptions = append(gexOptions, gex.ContainsChecker(_checker.args.(string)))
		case checkerModeEqual:
			gexOptions = append(gexOptions, gex.EqualChecker(_checker.args.(string)))
		}
	}

	// 收集器
	for _, _collector := range options.collectors {
		switch _collector.mode {
		case collectorModeString:
			gexOptions = append(gexOptions, gex.StringCollector(_collector.args.(*string)))
		}
	}

	// 计数器
	for _, _counter := range options.counters {
		switch _counter.mode {
		case counterModeLine:
			gexOptions = append(gexOptions, gex.LineCounter(_counter.args.(*int64)))
		}
	}

	// PWE处理
	if !options.pwe {
		gexOptions = append(gexOptions, gex.DisablePwe())
	}

	if !options.debug {
		gexOptions = append(gexOptions, gex.Quiet())
	} else {
		gexOptions = append(gexOptions, gex.Terminal())
	}

	// 执行命令
	_, err = gex.Exec(options.exe, gexOptions...)

	return
}
