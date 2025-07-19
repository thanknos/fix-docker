package ctl

import (
	"fix-docker/real"
	"fmt"
	"github.com/spf13/cobra"
)

var initOS = &cobra.Command{
	Use:       "init",
	Short:     "初始化安装命令",
	Long:      "初始化安装命令",
	Args:      cobra.OnlyValidArgs,
	ValidArgs: []string{"rpm", "apt"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}
	},
}

var rpmCommand = &cobra.Command{
	Use:  "rpm",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("开始初始化rpm软件安装包")
		rpmOS := real.NewRpmOS()
		rpmOS.InitOS()
		fmt.Println("结束rpm软件包的安装")
	},
}

var aptCommand = &cobra.Command{
	Use:  "apt",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("开始初始化apt软件安装包")
		aptOS := real.NewAptOS()
		aptOS.InitOS()
		fmt.Println("结束apt软件包的安装")
	},
}

func initOSCmd() {
	initOS.AddCommand(rpmCommand)
	initOS.AddCommand(aptCommand)
}
