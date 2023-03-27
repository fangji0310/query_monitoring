package monitoring

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type MonitoringPolicy struct {
	Title string
	Db    string
	Query string
}
type MonitoringPolicies []MonitoringPolicy

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
