/*
package main

import (
	"os"
	"syscall"
)

func main() {
	// 置き換えるプログラムのパス
	binary := "/bin/ls"
	
	// プログラムに渡す引数
	args := []string{"ls", "-l", "-a"}

	// 環境変数
	env := os.Environ()

	// execve システムコールを実行
	err := syscall.Exec(binary, args, env)
	if err != nil {
		panic(err)
	}
}
*/

/*
Go言語でシステムプログラミングを行う際に役立つ主要なパッケージの一覧を以下に示します。これらのパッケージは、ファイルシステムの操作、システムコールの実行、プロセスの管理、ネットワーキングといったシステムレベルのタスクに使用されます。

5. **`os/exec`**: 外部プロセスを実行し、その入出力を管理するための機能を提供します。

これらのパッケージはGo言語の標準ライブラリに含まれており、追加のインストールなしで使用できます。システムプログラミングにおいては、これらのパッケージを組み合わせて様々な低レベル操作を行うことが一般的です。
*/

package main

import (
    "os/exec"
    "fmt"
)

func main() {
	args := []string{"./"}
    //cmd := exec.Command("ls", "-a", "-l")
    cmd := exec.Command("ls", args...)

    output, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Printf("コマンド実行エラー: %v\n", err)
        return
    }

    fmt.Println(string(output))
}
