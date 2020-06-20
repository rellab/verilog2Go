package builder

import (
	"log"
	"os"
)

var file *os.File

// CreateNewFile は新しいファイルを生成する
func CreateNewFile() {
	file, err := os.OpenFile("generated/"+ModuleName+".go", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		//エラー処理
		log.Fatal(err)
	}
	defer file.Close()
}

// Write はファイルにgeneratedGoを書き込む
func Write() {
	file, err := os.OpenFile("generated/"+ModuleName+".go", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		//エラー処理
		log.Fatal(err)
	}
	defer file.Close()
	// fmt.Fprintln(file, ModuleName) //書き込み
}
