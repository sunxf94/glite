package mysql

import "fmt"

type Config struct {
	Addr    string
	Port    int
	User    string
	Passwd  string
	DBName  string
	CharSet string

	DialTimeout string
	// I/O read timeout. The value must be a decimal number with a unit suffix ("ms", "s", "m", "h"), such as "30s", "0.5m" or "1m30s". default:0 不超时
	ReadTimeout string
	// I/O write timeout. The value must be a decimal number with a unit suffix ("ms", "s", "m", "h"), such as "30s", "0.5m" or "1m30s".  default:0 即不超时
	WriteTimeout string
}

func (c *Config) getDataSourceName() string {
	if c.DialTimeout == "" {
		c.DialTimeout = "3s"
	}
	if c.ReadTimeout == "" {
		c.ReadTimeout = "3s"
	}
	if c.WriteTimeout == "" {
		c.WriteTimeout = "3s"
	}
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local&timeout=%s&readTimeout=%s&writeTimeout=%s",
		c.User, c.Passwd, c.Addr, c.Port, c.DBName,
		c.DialTimeout, c.ReadTimeout, c.WriteTimeout,
	)
}
