package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/chzyer/readline"
)

type GoShell struct {
	envMap map[string]string
}

func lineExport(line string) string {
	line = strings.TrimSpace(line)
	return line
}

//func envMapCreate()  {
func envMapCreate() map[string]string {

	m := make(map[string]string)

	for _, env := range os.Environ() {
		// 環境変数をキーと値に分割
		pair := strings.SplitN(env, "=", 2)
		m[pair[0]] = pair[1]
	}

	return m
}

func cmdExec(line string) bool {
	Tokens := strings.Split(line, " ")
	if len(Tokens) == 0 {
		return true
	}
	fmt.Println(len(Tokens))
	cmd := exec.Command(Tokens[0], Tokens[1:]...)
	//cmd.Stdin = os.Stdin

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("コマンド実行エラー: %v\n", err)
		return true
	}
	fmt.Println(string(output))
	return true
}

func main() {
	rl, err := readline.New("> ")
	if err != nil {
		fmt.Println("readlineの初期化に失敗しました:", err)
		os.Exit(1)
	}
	defer rl.Close()

	shell := GoShell{
		envMap: envMapCreate(),
	}

	for key, value := range shell.envMap {
		fmt.Printf("[%s]=[%s]\n", key, value)
	}

	for {
		line, err := rl.Readline()
		if err != nil {
			break
		}
		line = lineExport(line)
		if line == "" {
			continue
		}

		if cmdExec(line) == false {
			continue
		}

		// if line == "exit" {
		// 	os.Exit(0)
		// }
		//fmt.Println("入力:", line)
	}
}
