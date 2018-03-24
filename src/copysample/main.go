package main

import (
	"flag"
	"os"
	"fmt"
	"bufio"
	"strings"
	"io"
)

// 判断文件是否存在
func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

// 拷贝文件的行为
func copyFileAction(src, dst string, showProgress, force bool) {
	// 如果不是强制执行
	if !force {
		// 如果文件存在
		if fileExists(dst) {
			// 提示用户文件存在，并且读入命令行
			fmt.Printf("%s exists override? y/n \n", dst)
			reader := bufio.NewReader(os.Stdin)
			data, _, _ := reader.ReadLine()

			if strings.TrimSpace(string(data)) != "y" {
				return
			}
		}
	}
	copyFile(src, dst)

	if showProgress {
		fmt.Printf("'%s' -> '%s'\n ", src, dst)
	}

}

// copy file
func copyFile(src, dst string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dstFile.Close()

	return io.Copy(dstFile, srcFile)
}

func main() {
	var showProgress, force bool
	flag.BoolVar(&showProgress, "f", false, "force copy when existing")
	flag.BoolVar(&force, "v", false, "explain what is being done")

	flag.Parse()

	if flag.NArg() < 2 {
		flag.Usage()
		return
	}
	copyFileAction(flag.Arg(0), flag.Arg(1), showProgress, force)

}
