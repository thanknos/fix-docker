package real

import (
	"bytes"
	"fix-docker/real/operate"
	"fmt"
	"os/exec"
)

// 需要执行的命令集合
var execDockerComposeList = [][]string{
	{"chmod", "+x", "/usr/bin/docker-compose"},
	{"docker-compose", "version"},
}

type dockerCompose struct {
	filePath string
}

func NewDockerCompose() *dockerCompose {
	return &dockerCompose{filePath: "./file/docker-compose/"}
}

func (dockerCompose *dockerCompose) GetFilePath() string {
	return dockerCompose.filePath
}

func (dockerCompose *dockerCompose) Install() {
	operate.CopyFile(dockerCompose.GetFilePath()+"docker-compose", "/usr/bin/docker-compose")
	for _, cmdList := range execDockerComposeList {
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

func (dockerCompose *dockerCompose) UnInstall() {
	operate.DelFile("/usr/bin/docker-compose")
}
