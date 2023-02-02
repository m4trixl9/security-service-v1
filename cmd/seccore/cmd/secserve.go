/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"

	"github.com/device-security-v1/seccore/internal/app/seccore"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var http_port string

// secserveCmd represents the secserve command
var secserveCmd = &cobra.Command{
	Use:   "secserve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("secserve called")
		seccore.Serve(context.Background())
	},
}

func init() {
	rootCmd.AddCommand(secserveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// secserveCmd.PersistentFlags().String("foo", "", "A help for foo")
	secserveCmd.Flags().StringVarP(&http_port, "http_port", "p", "false", "port for http serve")
	secserveCmd.MarkFlagRequired("http_port") // 默认情况下，flag是optional（选填），若flag为必填，则需要做该设置

	// 无论是：全局配置，还是局部配置，都是通过 pflag 来实现的
	// 将 pflag 与 viper 绑定，顺理成章，则可以通过 viper.Get() 获取标志的值
	if err := viper.BindPFlag("config.port", secserveCmd.Flags().Lookup("http_port")); err != nil {
		panic(err)
	}

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// secserveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
