package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
)

func main() {
	loadEnvironment()
}

// QUERY_MONITORING_TOOL_ENVを読み込み、該当のenvファイルを読み込む
func loadEnvironment() {
	environment, ok := os.LookupEnv("QUERY_MONITORING_TOOL_ENV")
	if !ok {
		log.Panicln("Failed to get env.QUERY_MONITORING_TOOL_ENV")
	}
	if environment != "test" {
		log.Panicf("env.QUERY_MONITORING_TOOL_ENV is invalid.[%s]", environment)
	}
	executableFileDir := getExecutableFileDir()
	err := godotenv.Load(fmt.Sprintf("%s/config/%s.env", executableFileDir, environment))
	if err != nil {
		log.Panicln("Failed to load environment")
	}
}

// 実行ファイルパスを取得する
func getExecutableFileDir() string {
	exePath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return filepath.Dir(exePath)
}
