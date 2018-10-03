package execise

import (
	"bufio"
	"fmt"
	"learn8/stack"
	"os"
	"strconv"
	"strings"

	"github.com/urfave/cli"
)

// 从终端读取一行数据
func getInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	return reader.ReadString('\n')
}

// 转换为后缀表达式
func transPostExpress(express string) (postExpress []string, err error) {
	var opStack stack.Stack // 声明用来存储操作符的堆栈
	var i int
LABEL:
	for i < len(express) {
		switch {
		case express[i] >= '0' && express[i] <= '9': // 如果是数字
			var number []byte
			for ; i < len(express); i++ {
				if express[i] < '0' || express[i] > '9' {
					break
				}
				number = append(number, express[i])
			}
			postExpress = append(postExpress, string(number))
		case express[i] == '(':
			opStack.Push(fmt.Sprintf("%c", express[i]))
			i++
		case express[i] == ')':
			for !opStack.Emtyp() {
				data, _ := opStack.Pop()
				if data[0] == '(' {
					break
				}
				postExpress = append(postExpress, data)
			}
			i++
		case express[i] == '+' || express[i] == '-' || express[i] == '*' || express[i] == '/':
			if opStack.Emtyp() {
				opStack.Push(fmt.Sprintf("%c", express[i]))
				i++
				continue LABEL
			}
			data, _ := opStack.Top()
			if data[0] == '(' || data[0] == ')' {
				opStack.Push(fmt.Sprintf("%c", express[i]))
				i++
				continue LABEL
			}
			if (express[i] == '+' || express[i] == '-') ||
				((express[i] == '*' || express[i] == '/') && (data[0] == '*' || data[0] == '/')) {
				postExpress = append(postExpress, data)
				opStack.Pop()
				opStack.Push(fmt.Sprintf("%c", express[i]))
				i++
				continue LABEL
			}
			opStack.Push(fmt.Sprintf("%c", express[i]))
			i++
		default:
			err = fmt.Errorf("invalid express:%v", express[i])
			return
		}
	}

	for !opStack.Emtyp() {
		data, _ := opStack.Pop()
		if data[0] == '#' {
			break
		}

		postExpress = append(postExpress, data)
	}
	return
}

func calc(postExpress []string) (result int64, err error) {
	var n1, n2 string
	var s stack.Stack
	for i := 0; i < len(postExpress); i++ {
		var cur = postExpress[i]
		if cur[0] == '+' || cur[0] == '-' || cur[0] == '*' || cur[0] == '/' {
			n1, err = s.Pop()
			if err != nil {
				return
			}
			n2, err = s.Pop()
			if err != nil {
				return
			}

			num2, _ := strconv.Atoi(n1) // 字符串转数字
			num1, _ := strconv.Atoi(n2)
			var r1 int

			switch cur[0] {
			case '+':
				r1 = num1 + num2
			case '-':
				r1 = num1 - num2
			case '*':
				r1 = num1 * num2
			case '/':
				r1 = num1 / num2
			default:
				err = fmt.Errorf("invalid op")
				return
			}

			s.Push(fmt.Sprintf("%d", r1))
		} else {
			s.Push(cur)
		}
	}

	resultStr, err := s.Top()
	if err != nil {
		return
	}
	result, err = strconv.ParseInt(resultStr, 10, 64)
	return
}

func process(c *cli.Context) (err error) {
	for {
		// 获取用户输入
		express, err := getInput()
		if len(express) == 0 {
			continue
		}
		express = strings.TrimSpace(express)
		postExpress, errRes := transPostExpress(express)
		if errRes != nil {
			err = errRes
			fmt.Println(err)
			continue
		}

		result, errRes := calc(postExpress)
		if errRes != nil {
			fmt.Println("error: ", errRes)
			continue
		}
		fmt.Println("result: ", result)
	}
	return
}

func Test() {
	app := cli.NewApp()
	app.Name = "计算器"
	app.Usage = "实现简单计算器"

	app.Action = func(c *cli.Context) error {
		return process(c)
	}

	app.Run(os.Args)
}
