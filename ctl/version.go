package ctl

import (
	"fmt"
	"github.com/spf13/cobra"
)

var version = &cobra.Command{
	Use:   "version",
	Short: "打印fix-docker版本信息",
	Long:  "打印fix-docker版本信息",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("the fix-docker version is 1.0.0")
	},
}
