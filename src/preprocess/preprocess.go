package preprocess

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func Preprocess(sourceFile string) string {
	fp, err := os.Open(sourceFile)
	if err != nil {
		// error
	}
	defer fp.Close()
	var source []string
	defmap := map[string]string{}
	includeReg := `include\s*"(\w+).v"`
	reg := "\\`(\\w+)"
	includeRegex := regexp.MustCompile(includeReg)
	regex := regexp.MustCompile(reg)
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		// Process line by line
		line := scanner.Text()
		if includeRegex.Match([]byte(line)) {
			// Handling of include files
			a := includeRegex.FindStringSubmatch(line)
			defmap = searchDefs(filepath.Dir(sourceFile) + "/" + a[1] + ".v")
		} else if regex.Match([]byte(line)) {
			// Substitution of define statement
			source = append(source, replaceStr(line, defmap))
		} else {
			source = append(source, line)
		}
	}
	return strings.Join(source, "\n")
}

func searchDefs(includeFile string) map[string]string {
	defmap := map[string]string{}

	// ファイルオープン
	fp, err := os.Open(includeFile)
	if err != nil {
		// error
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	reg := `define\s*(\w+)\s*(\w+)'(\w+)`
	regex := regexp.MustCompile(reg)

	for scanner.Scan() {
		// Process line by line
		if regex.Match([]byte(scanner.Text())) {
			a := regex.FindStringSubmatch(scanner.Text())
			defmap[a[1]] = a[2] + "'" + a[3]
		}
	}
	fmt.Println(defmap)
	return defmap
}

func replaceStr(line string, defmap map[string]string) string {
	slice := strings.Split(line, " ")
	// Look up words that contain "`"
	for i, str := range slice {
		if strings.Contains(str, "`") {
			if val, ok := defmap[str[1:]]; ok {
				slice[i] = val
			} else {
				slice[i] = defmap[str[1:len(str)-1]] + str[len(str)-1:]
			}
		}
	}
	line = strings.Join(slice, " ")
	return line
}
