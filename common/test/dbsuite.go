package test

import (
	"database/sql"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
)

type DBConfig interface {
	GetPostgresConfigString() string
}

type DBSuite struct {
	suite.Suite
	DB     *sql.DB
	config DBConfig
}

func NewDBSuite(config DBConfig) DBSuite {
	return DBSuite{
		config: config,
	}
}

func (s *DBSuite) SetupSuite() {
	var err error
	dbConfig := s.config.GetPostgresConfigString()
	log.Debugf("dbConfig:%s", dbConfig)
	s.DB, err = sql.Open("postgres", dbConfig)
	if err != nil {
		log.Panicf("failed to connect to databse, err:%v", err)
	}
}
