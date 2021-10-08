package config

import (
	"log"

	"github.com/joeshaw/envdecode"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

// Config holds configuration for the project.
type Config struct {
	Port           string `env:"PORT,default=8080"`
	Env            string `env:"ENV,default=development"`
	Database       DatabaseConfig
	JWTConfig      JWTConfig
	InternalConfig InternalConfig
}

// DatabaseConfig holds configuration for database.
type DatabaseConfig struct {
	Host     string `env:"DATABASE_HOST,default=localhost"`
	Port     string `env:"DATABASE_PORT,default=5432"`
	Username string `env:"DATABASE_USERNAME,required"`
	Password string `env:"DATABASE_PASSWORD,required"`
	Name     string `env:"DATABASE_NAME,required"`
}

// JWTConfig holds configuration for JWT secret key
type JWTConfig struct {
	SecretKey string `env:"JWT_SECRET_KEY,required"`
}

// InternalConfig holds configuration for internal communication between microservices.
type InternalConfig struct {
	Username string `env:"SVC_USERNAME,required"`
	Password string `env:"SVC_PASSWORD,required"`
}

// NewConfig creates an instance of Config.
// It needs the path of the env file to be used.
func NewConfig(env string) (*Config, error) {
	var config Config
	if err := godotenv.Load(env); err != nil {
		log.Println(errors.Wrap(err, "[NewConfig] error reading .env file, defaulting to OS environment variables"))
	}

	if err := envdecode.Decode(&config); err != nil {
		return nil, errors.Wrap(err, "[NewConfig] error decoding env")
	}

	return &config, nil
}

// func CreateConnection() *gorm.DB {
// 	config, err := NewConfig(".env")
// 	if err != nil {
// 		panic(err)
// 	}

// 	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
// 		config.Database.Host,
// 		config.Database.Username,
// 		config.Database.Password,
// 		config.Database.Name,
// 		config.Database.Port)

// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("sukses konek")

// 	db.DB()

// 	return &gorm.DB{}
// }
