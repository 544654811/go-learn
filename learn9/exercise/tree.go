package tree

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/urfave/cli"
)

func buildTree(dir string, deep int) (err error) {
	dirs, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println("readDir is error, err: ", err)
		return
	}
	if deep == 1 {
		fmt.Println("|----", filepath.Base(dir))
	}
	sep := string(os.PathSeparator)
	for _, fn := range dirs {
		if fn.IsDir() {
			fmt.Printf("|")
			for i := 0; i < deep; i++ {
				fmt.Printf("    |")
			}
			fmt.Println("----", fn.Name())
			buildTree(dir+sep+fn.Name(), deep+1)
			continue
		}
		fmt.Printf("|")
		for i := 0; i < deep; i++ {
			fmt.Printf("    |")
		}
		fmt.Println("----", fn.Name())
	}
	return
}

func Test() {
	app := cli.NewApp()
	app.Name = "文件树状图"
	app.Usage = "输入文件目录，展示其树状图"

	app.Action = func(c *cli.Context) error {
		dir := "."
		if c.NArg() > 0 {
			dir = c.Args()[0]
		}
		buildTree(dir, 1)
		return nil
	}

	app.Run(os.Args)
}
