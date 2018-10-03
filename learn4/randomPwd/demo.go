package randomPwd

import (
	"flag" // 命令行工具包
	"fmt"
	"math/rand"
	"time"
)

var (
	length  int
	charset string
)

const (
	numCharset     = "0123456789"
	charCharset    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	mixCharset     = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	advanceCharset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ~`@#$%^&*-_=+|?/()<>[]{},.;'!"
)

func generate() string {
	password := make([]byte, length, length)
	sourceCharset := numCharset
	switch charset {
	case "num":
		sourceCharset = numCharset
	case "char":
		sourceCharset = charCharset
	case "mix":
		sourceCharset = mixCharset
	case "advance":
		sourceCharset = advanceCharset
	}
	for i := 0; i < length; i++ {
		index := rand.Intn(len(sourceCharset))
		password[i] = sourceCharset[index]
	}
	return string(password)
}

func initArgs() {
	flag.IntVar(&length, "l", 16, "-l 生成密码的长度")
	flag.StringVar(&charset, "t", "numCharset",
		`-t num 生成全数字的密码,
		-t char 包含全英文字符的密码,
		-t mix 包含数字和英文的密码,
		-t advance 包含数字、英文以及特殊字符的密码`)
	flag.Parse()
}

func Test() {
	rand.Seed(time.Now().UnixNano()) // 定义随机数种子
	initArgs()
	// fmt.Printf("length: %d, charset: %s\n", length, charset)
	password := generate()
	fmt.Println(password)
}
