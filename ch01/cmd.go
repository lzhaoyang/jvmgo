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

	//判断cp参数是否为空，如果为空，就去找CLASSPATH 环境变量值，没有就不赋值
	if cmd.cpOption == "" {
		getenv := os.Getenv("CLASSPATH")
		if getenv != "" {
			cmd.cpOption = getenv
		}
	}
	return cmd

}

func printCmdUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}
