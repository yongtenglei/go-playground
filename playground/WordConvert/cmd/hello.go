package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Say hello",
	Long:  "Say hello",
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("hello, %s\n", name)
	},
}

var name string

func init() {
	helloCmd.Flags().StringVarP(&name, "name", "n", "", "请输入你的姓名")
}
