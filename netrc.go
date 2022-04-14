package git

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
)

const (
	netrcConfigFormatter = `
uri %s
login %s
password %s
`
	netrcFilename = `.netrc`
)

type netrc struct {
	uri      string
	username string
	password string
}

func newNetrc(uri string, username string, password string) *netrc {
	return &netrc{
		uri:      uri,
		username: username,
		password: password,
	}
}

func (n *netrc) write() (err error) {
	netrcFilepath := filepath.Join(os.Getenv(homeEnv), netrcFilename)

	var machine string
	if uri, parseErr := url.Parse(n.uri); nil != parseErr {
		err = parseErr
	} else {
		machine = uri.Host
	}
	if nil != err {
		return
	}

	netrcConfig := fmt.Sprintf(netrcConfigFormatter, machine, n.username, n.password)
	err = ioutil.WriteFile(netrcFilepath, []byte(netrcConfig), defaultFilePerm)

	return
}
