package main

import (
    "os/exec"
    "fmt"
    "os"
)

func main() {
    // ファイルのパス（開きたいファイル）を指定
    filePath := "ファイルのパスをここに入力"

    // vimコマンドを実行してファイルを開く
    cmd := exec.Command("vim", filePath)

    // 標準入出力をターミナルに接続
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    // コマンドを実行
    if err := cmd.Run(); err != nil {
        fmt.Printf("コマンド実行エラー: %v\n", err)
    }
}

