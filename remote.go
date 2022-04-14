package git

var _ = Remote

func Remote(url string, opts ...remoteOption) (err error) {
	_options := newRemoteOptions(opts...)
	err = execAddons([]interface{}{`remote`, `add`, _options.name, url}, opts)

	return
}
