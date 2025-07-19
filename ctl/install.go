package ctl

import (
	"fix-docker/real"
	"fmt"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "安装命令，指定安装docker或docker-compose",
	Long:  "安装命令，指定安装docker或docker-compose",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var installDocker = &cobra.Command{
	Use:  "docker",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("开始安装docker命令")
		docker := real.NewDocker()
		docker.Install()
		fmt.Println("安装docker命令完成")
	},
}

var installDockerCompose = &cobra.Command{
	Use:  "docker-compose",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("开始安装docker-compose命令")
		compose := real.NewDockerCompose()
		compose.Install()
		fmt.Println("安装docker-compose命令完成")
	},
}

func initInstallCmd() {
	installCmd.AddCommand(installDocker)
	installCmd.AddCommand(installDockerCompose)
}
