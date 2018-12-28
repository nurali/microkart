package config

import (
	"fmt"

	"github.com/spf13/cast"
)

const (
	logLevel = "log.level"
	httpPort = "http.port"

	postgresHost     = "postgres.host"
	postgresPort     = "postgres.port"
	postgresUser     = "postgres.user"
	postgresPasswd   = "postgres.passwd"
	postgresDatabase = "postgres.database"
	postgresSSLMode  = "postgres.sslmode"
)

type Config struct {
	config map[string]interface{}
}

func New() *Config {
	c := new(Config)
	c.config = make(map[string]interface{})
	c.setDefaults()
	return c
}

func (c *Config) SetDefault(key string, value interface{}) {
	c.config[key] = value
}

func (c *Config) setDefaults() {
	c.SetLogLevel("info")
	c.SetHttpPort(8080)

	c.SetDefault(postgresHost, "localhost")
	c.SetDefault(postgresPort, 5432)
	c.SetDefault(postgresPasswd, "mysecretpassword")
	c.SetDefault(postgresUser, "postgres")
	c.SetDefault(postgresDatabase, "postgres")
	c.SetDefault(postgresSSLMode, "disable")
}

func (c *Config) GetHttpPort() int {
	return cast.ToInt(c.config[httpPort])
}

func (c *Config) SetHttpPort(port int) {
	c.config[httpPort] = port
}

func (c *Config) GetLogLevel() string {
	return cast.ToString(c.config[logLevel])
}

func (c *Config) SetLogLevel(level string) {
	c.config[logLevel] = level
}

func (c *Config) GetPostgresHost() string {
	return cast.ToString(c.config[postgresHost])
}

func (c *Config) GetPostgresPort() int {
	return cast.ToInt(c.config[postgresPort])
}

func (c *Config) GetPostgresUser() string {
	return cast.ToString(c.config[postgresUser])
}

func (c *Config) GetPostgresPasswd() string {
	return cast.ToString(c.config[postgresPasswd])
}

func (c *Config) GetPostgresDatabase() string {
	return cast.ToString(c.config[postgresDatabase])
}

func (c *Config) GetPostgresSSLMode() string {
	return cast.ToString(c.config[postgresSSLMode])
}

func (c *Config) GetPostgresConfigString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.GetPostgresHost(),
		c.GetPostgresPort(),
		c.GetPostgresUser(),
		c.GetPostgresPasswd(),
		c.GetPostgresDatabase(),
		c.GetPostgresSSLMode(),
	)
}
