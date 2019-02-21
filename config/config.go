package config

import (
	"fmt"
	"github.com/pelletier/go-toml"
	log "github.com/sirupsen/logrus"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
)

var globalConfig *Configure

// Config ...
func Config() *Configure {
	if globalConfig == nil {
		return DefaultConfig("")
	}
	return globalConfig
}

// REST ...
type REST struct {
	Enable bool   `toml:"enable"`
	Type   string `toml:"type"`
	Path   string `toml:"path"`
	Port   string `toml:"port"`
}

// REST ...
type GRPC struct {
	Enable bool   `toml:"enable"`
	Type   string `toml:"type"`
	Path   string `toml:"path"`
	Port   string `toml:"port"`
}

// Database ...
type Database struct {
	ShowSQL  bool   `toml:"show_sql"`
	Type     string `toml:"type"`
	Addr     string `toml:"addr"`
	Port     string `toml:"port"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	Schema   string `toml:"schema"`
	Location string `toml:"location"`
	Charset  string `toml:"charset"`
}

// Source ...
func (d *Database) Source() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?loc=%s&charset=%s&parseTime=true",
		d.Username, d.Password, d.Addr, d.Port, d.Schema, d.Location, d.Charset)
}

// WebToken ...
type WebToken struct {
	Key string `toml:"key"`
}

// Configure ...
type Configure struct {
	WebToken WebToken `toml:"web_token"`
	Database Database `toml:"database"`
	REST     REST     `toml:"rest"`
	GRPC     GRPC     `toml:"grpc"`
}

// Initialize ...
func Initialize(runPath string, configPath ...string) error {
	log.Info(runPath, configPath)
	s, e := filepath.Abs(filepath.Dir(runPath))
	if e != nil {
		s = ""
	}
	log.Info(s)
	globalConfig = DefaultConfig(s)

	globalConfig.LoadConfig(configPath[0])
	//globalConfig.ConfigPath, globalConfig.ConfigName = filepath.Split(configPath[0])
	return nil
}

// LoadConfig ...
func (c *Configure) LoadConfig(filePath string) *Configure {
	openFile, err := os.OpenFile(filePath, os.O_RDONLY|os.O_SYNC, os.ModePerm)
	if err != nil {
		log.Error("config open:", err)
		return c
	}
	defer openFile.Close()
	decoder := toml.NewDecoder(openFile)
	err = decoder.Decode(c)
	if err != nil {
		log.Error("config decode:", err)
		return c
	}
	log.Debugf("config: %+v", c)
	return c
}

// DefaultConfig ...
func DefaultConfig(runPath string) *Configure {
	return &Configure{
		WebToken: WebToken{
			Key: "im-godcong-yelion",
		},
		Database: Database{
			ShowSQL:  true,
			Type:     "mysql",
			Addr:     "localhost",
			Port:     "3306",
			Username: "root",
			Password: "111111",
			Schema:   "auth",
			Location: url.QueryEscape("Asia/Shanghai"),
			Charset:  "utf8mb4",
		},
		REST: REST{},
	}
}

// MustString ...
func MustString(v, def string) string {
	if v == "" {
		return def
	}
	return v
}

// MustInt ...
func MustInt(v string, def int) int {
	i, err := strconv.Atoi(v)
	if err == nil {
		return i
	}
	return def
}
