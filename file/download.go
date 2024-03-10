package file

import (
	"io"
	"net/http"
	"path/filepath"
)

/**
 * 下载文件到指定文件夹
 * 如果 newFilename 为空，则使用 url 中的文件名
 */
func DownloadFileToFolder(url string, folder string, newFilename string) (localPath string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	filename := filepath.Base(url)
	if newFilename != "" {
		filename = newFilename
	}
	localPath = filepath.Join(folder, filename)

	// 创建一个文件用于保存
	out, err := CreateFileWithFolder(localPath)
	if err != nil {
		return
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return
	}

	return
}

/**
 * 下载文件到指定文件
 */
func DownLoadFileToFile(url string, path string) (err error) {
	// 下载文件
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// 保存文件
	out, err := CreateFileWithFolder(path)
	if err != nil {
		return
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return
}
