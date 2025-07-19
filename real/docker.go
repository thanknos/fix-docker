package real

import (
	"bytes"
	"fix-docker/real/operate"
	"fmt"
	"os/exec"
)

// 需要执行的命令集合
var execDockerList = [][]string{
	{"chmod", "+x", "/usr/bin/docker"},
	{"chmod", "+x", "/usr/bin/docker-init"},
	{"chmod", "+x", "/usr/bin/docker-proxy"},
	{"chmod", "+x", "/usr/bin/dockerd"},
	{"chmod", "644", "/etc/systemd/system/docker.service"},
	{"systemctl", "enable", "docker"},
	{"systemctl", "daemon-reload"},
	{"systemctl", "start", "docker"},
	{"docker", "version"},
}

var stopDockerList = [][]string{
	{"systemctl", "stop", "docker"},
	{"systemctl", "disable", "docker"},
	{"systemctl", "daemon-reload"},
}

// docker命令安装与卸载
type docker struct {
	filePath string
}

func NewDocker() *docker {
	return &docker{filePath: "./file/docker/"}
}

func (docker *docker) GetFilePath() string {
	return docker.filePath
}

func (docker *docker) Install() {
	//复制完整目录与文件到操作系统
	err := operate.CopyDir(docker.GetFilePath()+"bin", "/usr/bin")
	if err != nil {
		fmt.Println(err)
	}

	err = operate.MakeDir("/etc/docker")
	if err != nil {
		fmt.Println(err)
	}

	err = operate.CopyFile(docker.GetFilePath()+"service/daemon.json", "/etc/docker/daemon.json")
	if err != nil {
		fmt.Println(err)
	}

	err = operate.CopyFile(docker.GetFilePath()+"service/docker.service", "/etc/systemd/system/docker.service")
	if err != nil {
		fmt.Println(err)
	}

	for _, cmdList := range execDockerList {
		command := exec.Command(cmdList[0], cmdList[1:]...)
		var stdout bytes.Buffer
		var stderr bytes.Buffer
		command.Stdout = &stdout
		command.Stderr = &stderr
		err := command.Run()
		if err != nil {
			fmt.Printf("%s\n", stderr.String())
			return
		}
		if cmdList[1] == "version" {
			fmt.Println(stdout.String())
		}
	}
}

func (docker *docker) UnInstall() {
	//关闭docker的所有服务
	for _, cmdList := range execDockerList {
		command := exec.Command(cmdList[0], cmdList[1:]...)
		var stdout bytes.Buffer
		var stderr bytes.Buffer
		command.Stdout = &stdout
		command.Stderr = &stderr
		err := command.Run()
		if err != nil {
			fmt.Printf("%s\n", stderr.String())
			return
		}
	}
	//删除docker已安装的源文件
	operate.DelDir("/etc/docker")
	operate.DelFile("/usr/bin/ctr")
	operate.DelFile("/usr/bin/runc")
	operate.DelFile("/usr/bin/docker")
	operate.DelFile("/usr/bin/dockerd")
	operate.DelFile("/usr/bin/docker-init")
	operate.DelFile("/usr/bin/docker-proxy")
	operate.DelFile("/etc/docker/daemon.json")
	operate.DelFile("/etc/docker/daemon.json")
	operate.DelFile("/usr/bin/containerd")
	operate.DelFile("/usr/bin/containerd-shim-runc-v2")
	operate.DelFile("/etc/systemd/system/docker.service")
}
