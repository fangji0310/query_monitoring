package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"query_monitoring/db"
	"query_monitoring/monitoring"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type QueryMonitoringStruct struct {
	Type               string
	IniPath            string
	MonitoringYamlPath string
}

func main() {
	environment := loadEnvironment()
	q := loadQueryMonitoringSetting()
	fmt.Printf("loading configuration manager. type: %s, ini: %s, yaml: %s\n", q.Type, q.IniPath, q.MonitoringYamlPath)
	_ = db.InitializeFromIni(environment, q.IniPath)
	fmt.Printf("loading monitoring yaml on %s\n", q.MonitoringYamlPath)
	_ = monitoring.LoadPolicy(q.MonitoringYamlPath)
	// TODO
}

func loadQueryMonitoringSetting() QueryMonitoringStruct {
	var q QueryMonitoringStruct
	err := envconfig.Process("QUERY_MONITORING", &q)
	if err != nil {
		log.Panic(err.Error())
	}
	if q.Type != "ini" {
		log.Panic("Not support")
	}
	if strings.HasPrefix(q.IniPath, "./") {
		executableFileDir := getExecutableFileDir()
		q.IniPath = strings.Replace(q.IniPath, "./", executableFileDir+"/", 1)
	}
	if strings.HasPrefix(q.MonitoringYamlPath, "./") {
		executableFileDir := getExecutableFileDir()
		q.MonitoringYamlPath = strings.Replace(q.MonitoringYamlPath, "./", executableFileDir+"/", 1)
	}
	return q
}

// QUERY_MONITORING_TOOL_ENVを読み込み、該当のenvファイルを読み込む
func loadEnvironment() string {
	environment, ok := os.LookupEnv("QUERY_MONITORING_TOOL_ENV")
	if !ok {
		log.Panicln("Failed to get env.QUERY_MONITORING_TOOL_ENV")
	}
	if environment != "test" {
		log.Panicf("env.QUERY_MONITORING_TOOL_ENV is invalid.[%s]", environment)
	}
	executableFileDir := getExecutableFileDir()
	fmt.Printf("loading environment from %s/config/%s.env\n", executableFileDir, environment)
	err := godotenv.Load(fmt.Sprintf("%s/config/%s.env", executableFileDir, environment))
	if err != nil {
		log.Panicln("Failed to load environment")
	}
	return environment
}

// 実行ファイルパスを取得する
func getExecutableFileDir() string {
	exePath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return filepath.Dir(exePath)
}
