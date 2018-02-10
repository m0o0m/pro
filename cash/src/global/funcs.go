//[公共] 全局公共方法
package global

import (
	"os"
	"os/exec"
	"path/filepath"
)

func GetExecPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		GlobalLogger.Error("error:%s", err.Error())
		return "", err
	}
	p, err := filepath.Abs(file)
	if err != nil {
		GlobalLogger.Error("error:%s", err.Error())
		return "", err
	}
	p = filepath.Dir(p)
	return p, nil
}

func GetTemplatePath() string {
	binPwd, err := GetExecPath()
	if err != nil {
		return binPwd + "/template"
	}
	return "./template"
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
