package operate

import (
	"errors"
	"io"
	"os"
	"strings"
)

func CopyDir(srcDir string, dstDir string) error {
	//检查目录是否正确
	if srcInfo, err := os.Stat(srcDir); err != nil {
		return err
	} else {
		if !srcInfo.IsDir() {
			return errors.New("源路径不是一个正确的目录！")
		}
	}

	if desInfo, err := os.Stat(dstDir); err != nil {
		return err
	} else {
		if !desInfo.IsDir() {
			return errors.New("目标路径不是一个正确的目录！")
		}
	}

	if strings.TrimSpace(srcDir) == strings.TrimSpace(dstDir) {
		return errors.New("源路径与目标路径不能相同！")
	}

	//开始复制目录
	dir, err := os.ReadDir(srcDir)
	if err != nil {
		return err
	}
	for _, entry := range dir {
		newDstName := dstDir + "/" + entry.Name()
		newSrcName := srcDir + "/" + entry.Name()
		if entry.IsDir() {
			err := MakeDir(newDstName)
			if err != nil {
				return errors.New("创建文件失败: " + newDstName)
			}
			errDir := CopyDir(newSrcName, newDstName)
			if errDir != nil {
				return errors.New("复制文件夹失败: " + newDstName)
			}
		} else {
			errFile := CopyFile(newSrcName, newDstName)
			if errFile != nil {
				return errors.New("复制文件失败: " + newDstName)
			}
		}
	}
	return err
}

func MakeDir(dir string) error {
	err := os.Mkdir(dir, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func DelDir(dir string) error {
	err := os.RemoveAll(dir)
	if err != nil {
		return err
	}
	return err
}

func CopyFile(srcFile string, dstFile string) error {
	readFile, err := os.OpenFile(srcFile, os.O_RDONLY|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	defer readFile.Close()

	makeFile, err := os.OpenFile(dstFile, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	defer makeFile.Close()

	io.Copy(makeFile, readFile)

	return nil
}

func DelFile(file string) error {
	err := os.Remove(file)
	if err != nil {
		return err
	}
	return nil
}
