package builder

import (
	"fmt"
	"log"
	"os"
)

var file *os.File

// Write はファイルにgeneratedGoを書き込む
func Write() {
	os.Mkdir("generated", 0777)
	file, err := os.OpenFile("generated/"+ModuleName+".go", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		//エラー処理
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Fprintln(file, Source) //書き込み
}
