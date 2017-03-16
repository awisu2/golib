package file

import (
	"io/ioutil"
	"os"
)

// ファイル読み込み
func ReadFile(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	s := string(data)
	return s, err
}

// ファイル書き込み
func WriteFile(filename string, data string, perm os.FileMode) error {
	return ioutil.WriteFile(filename, []byte(data), perm)
}

// ファイル書き込み
// if you want 0777 set os.ModePerm for os.FileMode
// windows os can't save other drive
func WriteFileByBytes(filename string, data []byte, perm os.FileMode) error {
	return ioutil.WriteFile(filename, data, perm)
}

// 存在チェック
func Exists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

// ディレクトリ作成
func MkdirIfNotExists(name string, perm os.FileMode) (err error) {
	if !Exists(name) {
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

// remove file
func Remove(path string) error {
	return os.Remove(path)
}

func Rename(oldpath, newpath string) error {
	return os.Rename(oldpath, newpath)
}
