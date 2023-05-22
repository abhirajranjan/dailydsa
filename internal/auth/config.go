package auth

type authHandler struct {
	dbbridge databasebridge
	AuthConfig
}

type AuthConfig struct {
	Key        string   `mapstructure:"key"`
	SigningAlg []string `mapstructure:"signing_alg"`
}
