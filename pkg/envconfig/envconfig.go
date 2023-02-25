package envconfig

import (
	"log"
	"time"

	"github.com/caarlos0/env/v6"
	_ "github.com/joho/godotenv/autoload"
)

type config struct {
	DB_HOST             string        `env:"DB_HOST,notEmpty"`
	DB_USERNAME         string        `env:"DB_USERNAME,notEmpty"`
	DB_PASSWORD         string        `env:"DB_PASSWORD,notEmpty"`
	DB_NAME             string        `env:"DB_NAME,notEmpty"`
	DB_PORT             int           `env:"DB_PORT,notEmpty"`
	APP_PORT            int           `env:"APP_PORT,notEmpty"`
	ALLOWED_ORIGIN      []string      `env:"ALLOWED_ORIGIN,notEmpty" envSeparator:","`
	PASETO_PRIVATE_KEY  string        `env:"PASETO_PRIVATE_KEY,notEmpty"`
	PASETO_EXPIRATION   time.Duration `env:"PASETO_EXPIRATION,notEmpty"`
	SUPABASE_SECRET_KEY string        `env:"SUPABASE_SECRET_KEY,notEmpty"`
	SUPABASE_BASE_URL   string        `env:"SUPABASE_BASE_URL,notEmpty"`
	SIGNED_BY           string        `env:"SIGNED_BY,notEmpty"`
}

var EnvVars config = config{}

func InitEnvVars() {

	if err := env.Parse(&EnvVars); err != nil {
		log.Fatalf("failed to parse EnvVars: %s", err)
	}
}
