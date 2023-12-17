package config

type Config struct {
	ApiKey           string `env:"API_KEY,required"`
	PostgresUser     string `env:"POSTGRES_USER,required"`
	PostgresPassword string `env:"POSTGRES_PASSWORD,required"`
	Database         string `env:"DB,required"`
	DatabaseHost     string `env:"DB_HOST,required"`
	DatabasePort     string `env:"DB_PORT,required"`
}
