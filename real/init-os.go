package real

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// 初始化配置文件，需要一下四类
var rpmList = []string{
	"lz4-1.8.3-1.el7.x86_64.rpm",
	"libseccomp-2.3.1-4.el7.x86_64.rpm",
	"libtool-ltdl-2.4.2-22.el7_3.x86_64.rpm",
	"containerd.io-1.6.32-3.1.el7.x86_64.rpm",
}

var aptList = []string{
	"apt-transport-https_2.6.1_all.deb",
	"ca-certificates_20230311_all.deb",
	"lsb-release_12.0-1_all.deb",
}

// 初始化总接口，所有的系统操作类型的初始化接口
type initOS interface {
	InitOS()
}

// 基本RPM操作系统的初始化实现
type rpmOS struct {
	pkgDir string
	initOS initOS
}

func NewRpmOS() rpmOS {
	return rpmOS{
		pkgDir: "file/rpm/",
		initOS: nil,
	}
}

func (rpmOS *rpmOS) GetPkgDir() string {
	return rpmOS.pkgDir
}

func (rpmOS *rpmOS) InitOS() {
	for _, pkg := range rpmList {
		commandList := []string{"rpm", "-ivhU", rpmOS.GetPkgDir() + pkg, "--nodeps"}
		command := exec.Command(commandList[0], commandList[1], commandList[2], commandList[3])
		var stdout bytes.Buffer
		var stderr bytes.Buffer
		command.Stdout = &stdout
		command.Stderr = &stderr
		err := command.Run()
		if err != nil {
			if strings.Contains(stderr.String(), "already") {
				continue
			}
			fmt.Printf("%s\n", stderr.String())
			fmt.Printf("%s\n", "基本RPM操作系统环境部署失败")
			return
		}
		fmt.Printf("%s\n", stdout.String())
	}
	fmt.Printf("%s\n", "基本RPM操作系统环境部署成功")
}

// 基本APT操作系统的初始化实现
type aptOS struct {
	pkgDir string
	initOS initOS
}

func NewAptOS() aptOS {
	return aptOS{
		pkgDir: "./file/apt/",
		initOS: nil,
	}
}

func (aptOS *aptOS) GetPkgDir() string {
	return aptOS.pkgDir
}

func (aptOS *aptOS) InitOS() {
	for _, pkg := range aptList {
		commandList := []string{"apt-get", "install", "-y", aptOS.GetPkgDir() + pkg, "--allow-downgrades"}
		command := exec.Command(commandList[0], commandList[1], commandList[2], commandList[3], commandList[4])
		var stdout bytes.Buffer
		var stderr bytes.Buffer
		command.Stdout = &stdout
		command.Stderr = &stderr
		err := command.Run()
		if err != nil {
			if strings.Contains(stderr.String(), "already") {
				continue
			}
			fmt.Printf("%s\n", stderr.String())
			fmt.Printf("%s\n", "基本APT操作系统环境部署失败")
			return
		}
		fmt.Printf("%s\n", stdout.String())
	}
	fmt.Printf("%s\n", "基本APT操作系统环境部署成功")
}
