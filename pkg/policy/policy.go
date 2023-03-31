package policy

import (
	"fmt"
	"os"
	"time"

	"query_monitoring/pkg/db"

	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

type MonitoringPolicy struct {
	Title               string `yaml:"title"`
	ScheduleOffsetMin   int    `yaml:"schedule_offset_min"`
	ScheduleIntervalMin int    `yaml:"schedule_interval_min"`
	Db                  string `yaml:"db"`
	Query               string `yaml:"query"`
}
type MonitoringPolicies []MonitoringPolicy

func (p MonitoringPolicy) IsExecute(batchStartTime time.Time) bool {
	current := batchStartTime.Hour()*60 + batchStartTime.Minute()
	if current < p.ScheduleOffsetMin {
		return false
	}
	return (current-p.ScheduleOffsetMin)%p.ScheduleIntervalMin == 0
}

func (p MonitoringPolicy) Check(manager db.QueryManager) (db.Metrics, error) {
	return manager.ExecuteQuery(p.Db, p.Query)
}

func LoadPolicy(logger *zap.Logger, filePath string) MonitoringPolicies {
	b, ioerr := os.ReadFile(filePath)
	if ioerr != nil {
		logger.Panic(ioerr.Error())
	}
	var r MonitoringPolicies
	yamlerr := yaml.Unmarshal(b, &r)
	if yamlerr != nil {
		logger.Panic(yamlerr.Error())
	}
	logger.Debug(fmt.Sprintf("Found %d monitoring policies.\n", len(r)))
	for _, val := range r {
		logger.Debug(fmt.Sprintf("-- %s\n", val.Title))
		logger.Debug(fmt.Sprintf("-- %s\n", val.Db))
	}
	return r
}
