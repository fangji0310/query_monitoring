package db

import (
	"fmt"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"gopkg.in/ini.v1"
)

type Metrics struct {
	Metrics int `db:"metrics"`
}

type dbConfiguration struct {
	sectionName string
	dbName      string
	host        string
	port        int
	user        string
	password    string
}

type QueryManager struct {
	dbConfiguration map[string]dbConfiguration
}

func (q QueryManager) ExecuteQuery(dbName string, query string) (Metrics, error) {
	db, connectionErr := q.getConnection(dbName)
	if connectionErr != nil {
		return Metrics{Metrics: -1}, connectionErr
	}
	defer db.Close()
	var metrics Metrics
	if err := db.Get(&metrics, query); err != nil {
		return Metrics{Metrics: -1}, fmt.Errorf("クエリの実行に失敗しました。DB[%s] Query[%s] Err[%w]", dbName, query, connectionErr)
	}
	return metrics, nil
}
func (q QueryManager) getConnection(dbName string) (*sqlx.DB, error) {
	d, ok := q.dbConfiguration[dbName]
	if !ok {
		return &sqlx.DB{}, fmt.Errorf("Not found db configuration %s", dbName)
	}
	return sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", d.user, d.password, d.host, d.port, d.dbName))
}

func InitializeFromIni(environment string, iniPath string) QueryManager {
	cfg, err := ini.Load(iniPath)
	if err != nil {
		log.Panic(err)
	}
	manager := QueryManager{dbConfiguration: make(map[string]dbConfiguration, 0)}
	for _, section := range cfg.Sections() {
		section_name := section.Name()
		prefix := environment + "-"
		if !strings.HasPrefix(section_name, environment+"-") {
			continue
		}
		new_section_name := strings.Replace(section_name, prefix, "", 1)
		config := dbConfiguration{
			sectionName: new_section_name,
			dbName:      section.Key("DBNAME").String(),
			host:        section.Key("HOST").String(),
			port:        section.Key("PORT").MustInt(3306),
			user:        section.Key("USER").String(),
			password:    section.Key("PASSWORD").String(),
		}
		manager.dbConfiguration[new_section_name] = config
	}
	return manager
}
