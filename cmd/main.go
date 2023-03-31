package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"query_monitoring/pkg/db"
	"query_monitoring/pkg/policy"
	"query_monitoring/pkg/util"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

type QueryMonitoringStruct struct {
	Type               string
	IniPath            string
	MonitoringYamlPath string
}

func main() {
	time.Local = time.FixedZone("JST", 9*60*60)
	environment := loadEnvironment()
	logger := util.NewLogger(environment)
	startTime := time.Now()
	q := loadQueryMonitoringSetting()
	logger.Debug(fmt.Sprintf("loading configuration manager type: %s, ini: %s, yaml: %s\n", q.Type, q.IniPath, q.MonitoringYamlPath))
	queryManager := db.InitializeFromIni(environment, q.IniPath)
	logger.Debug(fmt.Sprintf("loading policy yaml on %s\n", q.MonitoringYamlPath))
	monitoringPolicies := policy.LoadPolicy(logger, q.MonitoringYamlPath)
	for _, p := range monitoringPolicies {
		if !p.IsExecute(startTime) {
			logger.Debug(fmt.Sprintf("skip %s\n", p.Title))
			continue
		}
		metrics, err := p.Check(queryManager)
		if err != nil {
			logger.Warn(fmt.Sprintf("%w\n", err))
			continue
		}
		logger.Info(p.Title, zap.String("name", p.Title), zap.Int("count", metrics.Metrics))
	}
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
	if environment != "local" && environment != "test" && environment != "production" {
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
