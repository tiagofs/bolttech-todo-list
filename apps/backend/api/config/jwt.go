package config

type JWTConfig struct {
	JwtSecret string
	Type      string
	ExpiresIn int
}
