package git

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	sshConfigFormatter = `Host *
  IgnoreUnknown UseKeychain
  UseKeychain yes
  AddKeysToAgent yes
  StrictHostKeyChecking=no
  IdentityFile %s
`
	sshHome           = `.ssh`
	sshConfigFilename = `config`
	sshKeyFilename    = `id_rsa`
)

type ssh struct {
	dir string
	key string
}

func newSsh(dir string, key string) *ssh {
	return &ssh{
		dir: dir,
		key: key,
	}
}

func (s *ssh) write() (err error) {
	if err = os.MkdirAll(s.dir, os.ModePerm); nil != err {
		return
	}

	keyfile := filepath.Join(s.dir, sshKeyFilename)
	// 必须以换行符结束
	if !strings.HasSuffix(s.key, `\n`) {
		s.key = fmt.Sprintf(`%s\n`, s.key)
	}
	if err = ioutil.WriteFile(keyfile, []byte(s.key), defaultFilePerm); nil != err {
		return
	}

	configFile := filepath.Join(s.dir, sshConfigFilename)
	err = ioutil.WriteFile(configFile, []byte(fmt.Sprintf(sshConfigFormatter, keyfile)), defaultFilePerm)

	return
}
