package builder

import (
	"fmt"
	"log"
	"os"
)

// CreateNewFile は新しいファイルを生成する
func CreateNewFile() {
	file, err := os.OpenFile("generated/className.go", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		//エラー処理
		log.Fatal(err)
	}
	defer file.Close()
	print("書き込み")
	fmt.Fprintln(file, "書き込みテスト") //書き込み
}
