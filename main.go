package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/chzyer/readline"
)

func main() {
	rl, err := readline.New("> ")
	if err != nil {
		fmt.Println("readlineの初期化に失敗しました:", err)
		os.Exit(1)
	}
	defer rl.Close()
	//envp := os.Environ()
	for {
		line, err := rl.Readline()
		if err != nil {
			break
		}
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if line == "exit" {
			os.Exit(0)
		}
		fmt.Println("入力:", line)
	}
}
