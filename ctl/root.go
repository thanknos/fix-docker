package ctl

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "fix-docker",
	Short: "安装或卸载docker及docker-compose",
	Long:  "一键安装或卸载docker与docker-compose, 通过命令行直接使用",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func initRootCmd() {
	RootCmd.AddCommand(&cobra.Command{
		Use:    "completion",
		Hidden: true,
	})
	RootCmd.SetHelpCommand(&cobra.Command{
		Use:    "no-help",
		Hidden: true,
	})
	RootCmd.AddCommand(initOS)
	RootCmd.AddCommand(installCmd)
	RootCmd.AddCommand(uninstallCmd)
	RootCmd.AddCommand(version)
}

func init() {
	initOSCmd()
	initInstallCmd()
	initUninstallCmd()
	initRootCmd()
}
