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

	Limiter struct {
		RPS     float64
		Burst   int
		Enabled bool
	}
}
