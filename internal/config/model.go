package config

import (
	"github.com/abhirajranjan/dailydsa/internal/auth"
	"github.com/abhirajranjan/dailydsa/internal/database"
)

type config struct {
	Database database.DatabaseConfig `mapstructure:"database"`
	Auth     auth.AuthConfig         `mapstructure:"auth"`
}
