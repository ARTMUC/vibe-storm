package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	App      AppConfig      `mapstructure:"app"`
	JWT      JWTConfig      `mapstructure:"jwt"`
}

type ServerConfig struct {
	Port            string `mapstructure:"port"`
	ReadTimeout     string `mapstructure:"read_timeout"`
	WriteTimeout    string `mapstructure:"write_timeout"`
	ShutdownTimeout string `mapstructure:"shutdown_timeout"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
	SSLMode  string `mapstructure:"ssl_mode"`
}

type AppConfig struct {
	Name    string `mapstructure:"name"`
	Version string `mapstructure:"version"`
	Env     string `mapstructure:"env"`
}

type JWTConfig struct {
	Secret          string `mapstructure:"secret"`
	TokenDuration   string `mapstructure:"token_duration"`
	RefreshDuration string `mapstructure:"refresh_duration"`
}

func Load() *Config {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Set default values
	setDefaults()

	viper.AutomaticEnv()

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("Failed to unmarshal config:", err)
	}

	return &config
}

func setDefaults() {
	viper.SetDefault("server.port", "8080")
	viper.SetDefault("server.read_timeout", "10s")
	viper.SetDefault("server.write_timeout", "10s")
	viper.SetDefault("server.shutdown_timeout", "30s")

	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", 3306)
	viper.SetDefault("database.username", "root")
	viper.SetDefault("database.password", "password")
	viper.SetDefault("database.database", "vibe_storm")
	viper.SetDefault("database.ssl_mode", "false")

	viper.SetDefault("app.name", "VibeStorm")
	viper.SetDefault("app.version", "1.0.0")
	viper.SetDefault("app.env", getEnvOrDefault("GO_ENV", "development"))

	viper.SetDefault("jwt.secret", getEnvOrDefault("JWT_SECRET", "your-super-secret-jwt-key-change-this-in-production"))
	viper.SetDefault("jwt.token_duration", "24h")
	viper.SetDefault("jwt.refresh_duration", "168h") // 7 days
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// GetDSN returns the database connection string
func (c *DatabaseConfig) GetDSN() string {
	return c.Username + ":" + c.Password + "@tcp(" + c.Host + ":" + strconv.Itoa(c.Port) + ")/" + c.Database + "?parseTime=true&charset=utf8mb4&collation=utf8mb4_unicode_ci"
}
