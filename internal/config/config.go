package config

type Config struct {
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
	SSLMode    string
	JWTKey     string
}
