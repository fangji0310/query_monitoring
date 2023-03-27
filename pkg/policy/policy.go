package policy

import (
	"fmt"
	"log"
	"os"
	"time"

	"query_monitoring/pkg/db"

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

func (p MonitoringPolicy) Check(manager db.QueryManager) error {
	fmt.Println("-- Executing %s\n", p.Title)
	fmt.Println("-- offset: %d interval: %d\n", p.ScheduleOffsetMin, p.ScheduleIntervalMin)
	metrics, err := manager.ExecuteQuery(p.Db, p.Query)
	if err != nil {
		return err
	}
	fmt.Printf("name:%s count:%d\n", p.Title, metrics.Metrics)
	return nil
}

func LoadPolicy(filePath string) MonitoringPolicies {
	b, ioerr := os.ReadFile(filePath)
	if ioerr != nil {
		log.Panic(ioerr)
	}
	var r MonitoringPolicies
	yamlerr := yaml.Unmarshal(b, &r)
	if yamlerr != nil {
		log.Panic(yamlerr)
	}
	fmt.Printf("Found %d monitoring policies.\n", len(r))
	for _, val := range r {
		fmt.Printf("-- %s\n", val.Title)
		fmt.Printf("-- %s\n", val.Db)
	}
	return r
}
