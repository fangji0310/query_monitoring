package db

import (
	"log"
	"strings"

	"gopkg.in/ini.v1"
)

type dbConfiguration struct {
	sectionName string
	host        string
	user        string
	password    string
}

type QueryManager struct {
	dbConfiguration map[string]dbConfiguration
}

func InitializeFromIni(environment string, iniPath string) QueryManager {
	cfg, err := ini.Load(iniPath)
	if err != nil {
		log.Panic(err)
	}
	manager := QueryManager{dbConfiguration: make(map[string]dbConfiguration, 0)}
	for _, section := range cfg.Sections() {
		section_name := section.Name()
		if !strings.HasPrefix(section_name, environment+"-") {
			continue
		}
		config := dbConfiguration{
			sectionName: section_name,
			host:        section.Key("HOST").String(),
			user:        section.Key("USER").String(),
			password:    section.Key("PASSWORD").String(),
		}
		manager.dbConfiguration[section_name] = config
	}
	return manager
}
