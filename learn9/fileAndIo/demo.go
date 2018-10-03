package demo

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func openFileForOnlyRead() {
	file, err := os.Open("E:/gospace/src/learn9/fileAndIo/demo.go")
	if err != nil {
		fmt.Println("open file failed. err: ", err)
		return
	}
	defer file.Close()
}

func readFileByFile() {
	file, err := os.Open("E:/gospace/src/learn9/fileAndIo/demo.go")
	if err != nil {
		fmt.Println("open file failed. err: ", err)
		return
	}
	var data []byte
	var bytes [128]byte
	for {
		n, err := file.Read(bytes[:]) // 数组转切片
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read file failed. err: ", err)
			return
		}
		data = append(data, bytes[:n]...) // 切片追加数组

	}
	fmt.Println("读取数据为：", string(data))
	defer file.Close()
}

func readFileByIo() {
	file, err := os.Open("E:/gospace/src/learn9/fileAndIo/demo.go")
	if err != nil {
		fmt.Println("open file failed. err: ", err)
		return
	}
	defer file.Close()
	buf := bufio.NewReader(file)
	for {
		data, err := buf.ReadString('\n') // 读取一行数据
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read file failed. err: ", err)
			return
		}
		fmt.Println(data)
	}
}

func readFileByIoutil() {
	data, err := ioutil.ReadFile("E:/gospace/src/learn9/fileAndIo/demo.go")
	if err != nil {
		fmt.Println("read file failed. err: ", err)
	}
	fmt.Println(string(data))
}

func readGzipFile() {
	file, err := os.Open("E:/gospace/src/learn9/fileAndIo/demo.gzip")
	if err != nil {
		fmt.Println("open file failed. err: ", err)
		return
	}
	defer file.Close()

	gzip, err := gzip.NewReader(file)
	if err != nil {
		fmt.Println("gzip file failed. err: ", err)
		return
	}
	data, err := ioutil.ReadAll(gzip)
	if err != nil {
		fmt.Println("read file failed. err: ", err)
		return
	}
	fmt.Println(string(data))
}

func writeFileByFile() {
	file, err := os.OpenFile("E:/gospace/src/learn9/fileAndIo/test.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("openFile file failed. err: ", err)
		return
	}
	defer file.Close()

	str := "hello world...\n"
	file.Write([]byte(str))
	file.WriteString("你好吗？\n")
	str1 := "我很好...\n"
	file.WriteAt([]byte(str1), 1)
}

func writeFileByIo() {
	file, err := os.OpenFile("E:/gospace/src/learn9/fileAndIo/test.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("openFile file failed. err: ", err)
		return
	}
	defer file.Close()

	buf := bufio.NewWriter(file)
	str := "hello world...\n"
	buf.Write([]byte(str))
	for i := 0; i < 10; i++ {
		buf.WriteString(fmt.Sprintf("写入第%d条数据 \n", i))
	}

	buf.Flush()
}

func writeFileByIoutil() {
	str := "用 ioutil 包 写入文件..."
	err := ioutil.WriteFile("E:/gospace/src/learn9/fileAndIo/test.txt", []byte(str), 0111)
	if err != nil {
		fmt.Println("ioutil write file failed. err: ", err)
		return
	}
}

func copyFile() {
	file1, err := os.Open("E:/gospace/src/learn9/fileAndIo/test.txt")
	if err != nil {
		fmt.Println("openFile file failed. err: ", err)
		return
	}
	defer file1.Close()

	file2, err := os.OpenFile("E:/gospace/src/learn9/fileAndIo/test1.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("openFile file failed. err: ", err)
		return
	}
	defer file2.Close()

	_, err1 := io.Copy(file2, file1)
	if err1 != nil {
		fmt.Println("copy file failed. err: ", err)
		return
	}
}

func Test() {
	// openFileForOnlyRead()
	// readFileByFile()
	// readFileByIo()
	// readFileByIoutil()
	// readGzipFile()
	// writeFileByFile()
	// writeFileByIo()
	// writeFileByIoutil()
	copyFile()
}
