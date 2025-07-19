package ctl

import (
	"fix-docker/real"
	"fmt"
	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "卸载命令，指定卸载docker或docker-compose",
	Long:  "卸载命令，指定卸载docker或docker-compose",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var uninstallDocker = &cobra.Command{
	Use:  "docker",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("开始卸载docker命令")
		docker := real.NewDocker()
		docker.UnInstall()
		fmt.Println("卸载docker命令完成")
	},
}

var uninstallDockerCompose = &cobra.Command{
	Use:  "docker-compose",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("开始卸载docker-compose命令")
		compose := real.NewDockerCompose()
		compose.UnInstall()
		fmt.Println("卸载docker-compose命令完成")
	},
}

func initUninstallCmd() {
	uninstallCmd.AddCommand(uninstallDocker)
	uninstallCmd.AddCommand(uninstallDockerCompose)
}
