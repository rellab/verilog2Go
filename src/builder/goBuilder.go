package builder

import "fmt"

var ModuleName string

func StartModule(moduleName string) {
	fmt.Println(moduleName)
	ModuleName = moduleName
}
