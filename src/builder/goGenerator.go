package builder

import (
	"fmt"
	"log"
	"os"
)

var file *os.File

// CreateNewFile は新しいファイルを生成する
func CreateNewFile() {
	file, err := os.OpenFile("generated/className.go", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		//エラー処理
		log.Fatal(err)
	}
	defer file.Close()
}

// Write はファイルにgeneratedGoを書き込む
func Write(generatedGo string) {
	fmt.Println(generatedGo)
	file, err := os.OpenFile("generated/className.go", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		//エラー処理
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Fprintln(file, ModuleName) //書き込み
}
