package main

import (
    "fmt"
    "strings"
)

func main() {
    inputString := "This"

    // 文字列を空白で分割
    tokens := strings.Split(inputString, " ")

    if len(tokens) > 0 {
        // 最初のトークンを取得
        firstToken := tokens[0]

        // 最初のトークン以外の部分を取得
        otherTokens := strings.Join(tokens[1:], " ")

		fmt.Printf("元の文字列		: [%s]\n", inputString)
		fmt.Printf("最初のトークン		: [%s]\n", firstToken)
		fmt.Printf("その他のトークン	: [%s]\n", otherTokens)


    }
}

