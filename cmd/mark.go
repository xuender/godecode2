package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("create: data.go")
	fout, err := os.Create("../data.go")
	defer fout.Close()
	if err != nil {
		panic(err)
	}
	fout.WriteString("package godecode2\n")
	fout.WriteString("var cache = make(map[rune][]string)\n")
	fout.WriteString("func init(){\n")
	filepath.Walk("../../godecode/data",
		func(path string, fi os.FileInfo, err error) error {
			if fi == nil {
				return err
			}
			if fi.IsDir() {
				return nil
			}
			num, err := strconv.ParseInt(path[len(path)-2:], 16, 0)
			if err != nil {
				return err
			}
			fout.WriteString("\tcache[" + strconv.FormatInt(num, 10) + "] = []string {\n")
			f, err := os.OpenFile(path, os.O_RDONLY, 0660)
			if err != nil {
				panic(err)
			}
			v := read(bufio.NewScanner(f))
			fout.WriteString("\t\t" + strings.Join(v, ",\n\t\t") + ",\n")
			fout.WriteString("\t}\n")
			return nil
		})
	fout.WriteString("}\n")
}
func read(scanner *bufio.Scanner) []string {
	var ret []string
	for scanner.Scan() {
		ret = append(ret, "\""+scanner.Text()+"\"")
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}
	return ret
}
