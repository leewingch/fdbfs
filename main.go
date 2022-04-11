package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"commands"
)

type Context struct {
	CurrentPath string
}

func NewContext() *Context {
	return &Context{CurrentPath: "/"}
}

func main() {
	fmt.Println("hello world")

	context := NewContext();
	var parser commands.CmdParser

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("fdbfs> ");
		cmdStr, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		cmdStr = strings.TrimSuffix(cmd, "\n")

		cmd, err := parser.parse(cmdStr);
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		cmd.Execute(context);
	}
}

func printHelp() {
	fmt.Println("fdbfs ls | cd | cat | quit");
}
