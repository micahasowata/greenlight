package config

type Config struct {
	Port int
	Env  string
	Db   struct {
		DSN          string
		MaxOpenConns int
		MaxIdleConns int
		MaxIdleTime  string
	}
}
