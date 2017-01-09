package file

import (
	"github.com/awisu2/golib/log"
	"io/ioutil"
	"os"
)

// ファイル読み込み
func ReadFile(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	s := ""
	if err == nil {
		s = string(data)
	} else {
		log.Errorf(err)
	}

	return s, err
}

// ファイル書き込み
func WriteFile(filename string, data string, perm os.FileMode) error {
	err := ioutil.WriteFile(filename, []byte(data), perm)
	if err != nil {
		log.Errorf(err)
		return err
	}
	return nil
}

// 存在チェック
func Exists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

// ディレクトリ作成
func MkdirIfNotExists(name string, perm os.FileMode) (err error) {
	if !Exists(name) {
		log.Println("file not exists, try create dir. " + name)
		return os.Mkdir(name, perm)
	}
	return
}

// ディレクトリ配下のファイル一覧(ディレクトリは除く)
func FileInfos(dir string) (fileInfos []os.FileInfo, err error) {
	infos, err := ioutil.ReadDir(dir)
	if err != nil {
		return
	}

	fileInfos = make([]os.FileInfo, 0)
	for _, info := range infos {
		if info.IsDir() {
			continue
		}
		fileInfos = append(fileInfos, info)
	}
	return
}
