package tools

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

/*判断文件是否是目录*/
func checkFileIsDir(filename string) (bool, error) {

	f, err := os.Stat(filename)
	if err != nil {
		return false, err
	}
	return f.IsDir(), err
}

/* 判断文件是否存在  存在返回 true 不存在返回false*/
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

/*写文件*/
func WriteFile(filePath, source string) {
	var f *os.File
	if checkFileIsExist(filePath) { //如果文件存在
		f, _ = os.OpenFile(filePath, os.O_APPEND, 0666) //打开文件
	} else {
		f, _ = os.Create(filePath) //创建文件
	}
	io.WriteString(f, source) //写入文件(字符串
	f.Close()
}

/*读文件*/
func ReadBufio(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(bytes), err
}

/*创建临时文件*/
func MkTmpDir() string {
	//fmt.Println(checkFileIsExist("./test_tmp"))
	if !checkFileIsExist("./test_tmp") {
		os.Mkdir("./test_tmp", os.ModePerm)
	}
	tmp_path := fmt.Sprintf("./test_tmp/%d", time.Now().UnixNano())
	os.Mkdir(tmp_path, os.ModePerm)
	return tmp_path
}

/*创建工作文件*/
func MkWorkDir() string {
	if !checkFileIsExist("./test_result") {
		os.Mkdir("./test_result", os.ModePerm)
	}
	tmp_path := fmt.Sprintf("./test_result/%d", time.Now().UnixNano())
	os.Mkdir(tmp_path, os.ModePerm)
	return tmp_path
}

/*复制文件*/
func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

/*删除目录*/
func DelDir(path string) error {
	if !checkFileIsExist(path) {
		return nil
	}
	err := os.RemoveAll(path)
	return err
}

/*遍历目录*/
func RangeDir(path string, fn func(path string, fi os.FileInfo, err error) error) error {

	if !checkFileIsExist(path) {
		return nil
	}
	err := filepath.Walk(path, fn)
	return err
}
