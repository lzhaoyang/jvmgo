/** 命令行结构体
 * @Author: sunrise
 * @Description:
 * @File:  cmd
 * @Version: 1.0.0
 * @Date: 2020/7/19 19:39
 */

package main

import (
	"flag"
	"fmt"
	"os"
)

//命令行结构体表示
type Cmd struct {
	//帮助
	helpFlag bool
	//版本
	versionFlag bool
	//classpath参数
	cpOption string
	//类名
	class string
	//args main方法参数
	args []string
}

//解析命令行参数到结构体
func parseCmd() *Cmd {
	//初始化结构体
	cmd := &Cmd{}
	//默认使用方法打印
	flag.Usage = printCmdUsage
	//帮助提示
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	//版本提示
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	//classpath参数
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	//编译处理参数
	//if !flag.Parsed() {
	flag.Parse()
	//}
	//获取参数
	args := flag.Args()
	if len(args) > 0 {
		//类名 java Test
		cmd.class = args[0]
		cmd.args = args[1:]
	}
	return cmd

}

//环境变量获取classpath参数，CLASSPATH环境变量可以被命令行参数覆盖
//即cli args的优先级更高
func getEnvValueByKey(key string) (string, error) {

}

func printCmdUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}
