package config

type Config struct {
	ServerPort int     `env:"SERVER_PORT"`
	DSN        string `env:"DSN"`
}

const (
	defaultServerPort = 8888
	defaultDSN = "host=127.0.0.1 port=5432 user=user dbname=tasks sslmode=disable"
)

func DefaultConfig() *Config {
	return &Config{
		ServerPort: defaultServerPort,
		DSN: defaultDSN,
	}
}

func TestConfig() *Config {
	return &Config{
		ServerPort: defaultServerPort,
		DSN:        "host=127.0.0.1 port=5432 user=user dbname=tasks_test sslmode=disable",
	}
}
