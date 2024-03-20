package basic

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func testBufferIoRead(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("close file failed, fileName:", filePath, err)
			return
		}
		fmt.Println("close file failed, fileName:", filePath, err)
	}(file)
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n') //注意是字符
		if err == io.EOF {
			if len(line) != 0 {
				fmt.Println(line)
			}
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		fmt.Print(line)
	}
}

func testBufferIoWriter(filePath string) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writeString, err := writer.WriteString("hello沙河\n")
		if err != nil {
			fmt.Println("write string failed. str:", writeString)
		} //将数据先写入缓存
	}
	err = writer.Flush()
	if err != nil {
		fmt.Println("close writer failed. filePath:", filePath)
		return
	}
}

func testCopyFile() {

}
