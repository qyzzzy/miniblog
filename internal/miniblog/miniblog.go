package miniblog

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewMiniBlogCommand() *cobra.Command {
	cmd := &cobra.Command{
		//指定名字
		Use: "miniblog",
		//命令的简短描述
		Short: "A good Go practical project",
		//命令的详细描述
		Long: `A good GO practical project,used to create user with basic information
		Find more miniblog information at:
		    https://github.com/qylearn/miniblog#readme`,
		//命令出错时，不打印帮助信息。不需要打印帮助信息，设置为true可以保持命令出错时一眼就能看到错误信息
		SilenceUsage: true,
		//指定调用cmd.Execute()时，执行的RUN函数，函数执行失败会返回错误信息
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()

		},
		//这里设置命令运行的时，不需要指定命令行参数
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any argument,got %q", cmd.CommandPath(), args)

				}
			}
			return nil

		},
	}
	return cmd

}
func run() error {
	fmt.Println("Hello Miniblog")
	return nil

}
