/**
 * @Author: sunrise
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2020/7/19 19:44
 */

package main

import "fmt"

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printCmdUsage()
	} else {
		startJvmGo(cmd)
	}
}

func startJvmGo(cmd *Cmd) {
	fmt.Printf("classpath: %s, class: %s args: %v \n", cmd.cpOption, cmd.class, cmd.args)
}
