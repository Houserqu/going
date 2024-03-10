package file

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strconv"

	"github.com/dlclark/regexp2"
)

/**
 * 创建文件，如果文件夹不存在则创建文件夹
 */
func CreateFileWithFolder(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.Create(p)
}

/**
 * 移动文件
 */
func MoveFile(sourcePath, destPath string) (err error) {
	inputFile, err := os.Open(sourcePath)
	if err != nil {
		return
	}
	outputFile, err := os.Create(destPath)
	if err != nil {
		inputFile.Close()
		return
	}
	defer outputFile.Close()

	_, err = io.Copy(outputFile, inputFile)
	inputFile.Close()
	if err != nil {
		return
	}
	// The copy was successful, so now delete the original file
	err = os.Remove(sourcePath)
	if err != nil {
		return
	}
	return
}

/**
 * 复制文件
 */
func Copy(srcFile, dstFile string) error {
	out, err := os.Create(dstFile)
	if err != nil {
		return err
	}

	defer out.Close()

	in, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer in.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	return nil
}

/**
 * 复制文件夹
 */
func CopyDirectory(scrDir, dest string) error {
	entries, err := os.ReadDir(scrDir)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		sourcePath := filepath.Join(scrDir, entry.Name())
		destPath := filepath.Join(dest, entry.Name())

		fileInfo, err := os.Stat(sourcePath)
		if err != nil {
			return err
		}

		// stat, ok := fileInfo.Sys().(*syscall.Stat_t)
		// if !ok {
		// 	return fmt.Errorf("failed to get raw syscall.Stat_t data for '%s'", sourcePath)
		// }

		switch fileInfo.Mode() & os.ModeType {
		case os.ModeDir:
			if err := CreateFolder(destPath, 0770); err != nil {
				return err
			}
			if err := CopyDirectory(sourcePath, destPath); err != nil {
				return err
			}
		case os.ModeSymlink:
			if err := CopySymLink(sourcePath, destPath); err != nil {
				return err
			}
		default:
			if err := Copy(sourcePath, destPath); err != nil {
				return err
			}
		}

		// if err := os.Lchown(destPath, int(stat.Uid), int(stat.Gid)); err != nil {
		// 	return err
		// }

		fInfo, err := entry.Info()
		if err != nil {
			return err
		}

		isSymlink := fInfo.Mode()&os.ModeSymlink != 0
		if !isSymlink {
			if err := os.Chmod(destPath, fInfo.Mode()); err != nil {
				return err
			}
		}
	}
	return nil
}

/**
 * 判断文件是否存在
 */
func Exists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}

	return true
}

/**
 * 创建目录
 */
func CreateFolder(dir string, perm os.FileMode) error {
	if Exists(dir) {
		return nil
	}

	if err := os.MkdirAll(dir, perm); err != nil {
		return fmt.Errorf("failed to create directory: '%s', error: '%s'", dir, err.Error())
	}

	return nil
}

/**
 * 复制软连接
 */
func CopySymLink(source, dest string) error {
	link, err := os.Readlink(source)
	if err != nil {
		return err
	}
	return os.Symlink(link, dest)
}

/**
 * 读取目录中文件并根据末尾的数字进行排序
 */
func ReadFolderWithLastNumberSort(path string) (files []fs.DirEntry, err error) {
	// 读取目录
	files, err = os.ReadDir(path)
	if err != nil {
		return
	}

	// 根据文件名进行排序
	sort.Slice(files, func(i, j int) bool {
		iNumber, iNumberErr := GetFileNumber(files[i].Name())
		if iNumberErr != nil {
			err = iNumberErr
		}

		jNumber, jNumberErr := GetFileNumber(files[j].Name())
		if jNumberErr != nil {
			err = jNumberErr
		}

		return iNumber < jNumber
	})

	return
}

// 从文件名中解析出数字
func GetFileNumber(fileName string) (number int, err error) {
	re := regexp2.MustCompile("\\d+(?=\\.\\w+$)", 0)
	m, _ := re.FindStringMatch(fileName)

	if m == nil { // 匹配失败
		err = errors.New("get filename number fail: " + fileName)
		return
	}

	value := m.String()
	number, err = strconv.Atoi(value)
	return
}
