package database

import "fmt"

// DBConnectStringGenerator 数据库连接字符串创建接口
type DBConnectStringGenerator interface {
	ConnString() (dialect string, connString string)
}

// Configure 数据库基础配置信息
type Configure struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DB       string `json:"db"`
	// Charset        string `json:"charset"`
	ConnectTimeout int `json:"connectTimeout"`
}

// PostgresConfigure pgsql 配置信息
type PostgresConfigure struct {
	Configure
	SSLMode string `json:"sslMode"` // 可选 verify-full, disable ...
}

// ConnString 生成连接字符串
func (c *PostgresConfigure) ConnString() (dialect string, connString string) {
	dialect = "postgres"
	if c.SSLMode == "" {
		c.SSLMode = "disable"
	}
	if c.ConnectTimeout == 0 {
		c.ConnectTimeout = 10
	}
	connString = fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s connect_timeout=%d",
		c.Host, c.Port, c.User, c.DB, c.Password, c.SSLMode, c.ConnectTimeout)
}
