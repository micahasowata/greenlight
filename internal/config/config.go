package config

import (
	"flag"
	"os"
)

type Config struct {
	Port int
	Env  string
	Db   struct {
		DSN          string
		MaxOpenConns int
		MaxIdleConns int
		MaxIdleTime  string
	}

	Limiter struct {
		RPS     float64
		Burst   int
		Enabled bool
	}
	SMTP struct {
		Host     string
		Port     int
		Username string
		Password string
		Sender   string
	}
}

func (c *Config) Parse() {
	flag.IntVar(&c.Port, "port", 4000, "API server port")
	flag.StringVar(&c.Env, "env", "development", "development|staging|production")
	flag.StringVar(&c.Db.DSN, "db-dsn", os.Getenv("DB_DSN"), "PostgresQL dsn")
	flag.IntVar(&c.Db.MaxOpenConns, "db-max-open-conns", 25, "PostgresQL max open connections")
	flag.IntVar(&c.Db.MaxIdleConns, "db-max-idle-conns", 25, "PostgresQL max idle connections")
	flag.StringVar(&c.Db.MaxIdleTime, "db-max-idle-time", "15m", "PostgresQL max connection idle time")
	flag.Float64Var(&c.Limiter.RPS, "limiter-rps", 2, "Rate limiter maximum request per second")
	flag.IntVar(&c.Limiter.Burst, "limiter-burst", 4, "Rate limiter maximum burst")
	flag.BoolVar(&c.Limiter.Enabled, "limiter-enabled", true, "Enable rate limiter")
	flag.StringVar(&c.SMTP.Host, "smtp-host", os.Getenv("SMTP_HOST"), "SMTP host")
	flag.IntVar(&c.SMTP.Port, "smtp-port", 25, "SMTP port")
	flag.StringVar(&c.SMTP.Username, "smtp-username", os.Getenv("SMTP_NAME"), "SMTP username")
	flag.StringVar(&c.SMTP.Password, "smtp-password", os.Getenv("SMTP_PASSWORD"), "SMTP password")
	flag.StringVar(&c.SMTP.Sender, "smtp-sender", os.Getenv("SMTP_SENDER"), "SMTP sender")
	flag.Parse()
}
