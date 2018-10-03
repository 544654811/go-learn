package main

import (
	"fmt"
	// "learn10/interface_demo"
	// "learn10/empty_interface"
	"logger"
)

func InitLogger(name, filePath, fileName, level string) {
	config := make(map[string]string, 8)
	config["log_level"] = level
	config["log_path"] = filePath
	config["log_name"] = fileName
	err := logger.InitLogger(name, config)
	if err != nil {
		fmt.Printf("init logger failed")
	}
}

func main() {
	// demo.Test()
	InitLogger("file", "E:/gospace/src/logger", "myLogger1", "debug")
	logger.Debug("this is debug log...")
	logger.Trace("tihs is trace log...")
	logger.Info("this is info log...")
	logger.Warn("this is warn log...")
	logger.Error("this is error log...")
	logger.Fatal("this is fatal log...")
}
