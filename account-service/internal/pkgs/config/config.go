package config

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

var (
	EnvFilePath = ".env"
	Config      *GlobalConfig
)

type Environment string

type GlobalConfig struct {
	SERVER      ServerConfig
	DATABASE    DatabaseConfig
	LOGGER      LoggerConfig
	SECRETS     SecretConfig
	ENVIRONMENT Environment
}

type ServerConfig struct {
	HOST     string
	PORT     string
	MODE     string
	BASE_URL string
}

type DatabaseConfig struct {
	HOST     string
	PORT     string
	USER     string
	PASSWORD string
	NAME     string
}

type LoggerConfig struct {
	LOG_LEVEL           string
	LOG_FILE_PATH_ERROR string
	LOG_FILE_PATH_INFO  string
}

type SecretConfig struct {
	JWT_SECRET                    string
	BCRYPT_SECRET                 string
	SESSION_COOKIE_SECRET         string
	SESSION_ID_TOKEN_NAME         string
	SESSION_ID_TOKEN_EXP_TIME     int
	REFRESH_TOKEN_EXP_TIME        int
	VERIFICATION_TOKEN_EXP_TIME   string
	PASSWORD_RESET_TOKEN_EXP_TIME int
	TWO_FA_EXP_TIME               int
	TWO_FA_ISSUER                 string
	PRIVATE_KEY                   *rsa.PrivateKey
	PUBLIC_KEY                    *rsa.PublicKey
}

func LoadConfig() (*GlobalConfig, error) {
	viper.SetConfigFile(EnvFilePath)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Warning: Could not read .env file: %v", err)
	}

	// privateKey, err := loadPrivateKey("private_key.pem")
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to load private key: %v", err)
	// }

	// publicKey, err := loadPublicKey("public_key.pem")
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to load public key: %v", err)
	// }

	Config = &GlobalConfig{
		SERVER: ServerConfig{
			HOST:     viper.GetString("SERVER_HOST"),
			PORT:     viper.GetString("SERVER_PORT"),
			MODE:     viper.GetString("SERVER_MODE"),
			BASE_URL: viper.GetString("BASE_URL"),
		},
		DATABASE: DatabaseConfig{
			HOST:     viper.GetString("DB_HOST"),
			PORT:     viper.GetString("DB_PORT"),
			USER:     viper.GetString("DB_USER"),
			PASSWORD: viper.GetString("DB_PASSWORD"),
			NAME:     viper.GetString("DB_NAME"),
		},
		LOGGER: LoggerConfig{
			LOG_LEVEL:           viper.GetString("LOG_LEVEL"),
			LOG_FILE_PATH_ERROR: viper.GetString("LOG_FILE_PATH_ERROR"),
			LOG_FILE_PATH_INFO:  viper.GetString("LOG_FILE_PATH_INFO"),
		},
		SECRETS: SecretConfig{
			JWT_SECRET:                    viper.GetString("JWT_SECRET"),
			BCRYPT_SECRET:                 viper.GetString("BCRYPT_SECRET"),
			SESSION_COOKIE_SECRET:         viper.GetString("SESSION_COOKIE_SECRET"),
			SESSION_ID_TOKEN_NAME:         viper.GetString("SESSION_ID_TOKEN_NAME"),
			SESSION_ID_TOKEN_EXP_TIME:     viper.GetInt("SESSION_ID_TOKEN_EXP_TIME"),
			REFRESH_TOKEN_EXP_TIME:        viper.GetInt("REFRESH_TOKEN_EXP_TIME"),
			VERIFICATION_TOKEN_EXP_TIME:   viper.GetString("VERIFICATION_TOKEN_EXP_TIME"),
			PASSWORD_RESET_TOKEN_EXP_TIME: viper.GetInt("PASSWORD_RESET_TOKEN_EXP_TIME"),
			TWO_FA_EXP_TIME:               viper.GetInt("TWO_FA_EXP_TIME"),
			TWO_FA_ISSUER:                 viper.GetString("TWO_FA_ISSUER"),
			// PRIVATE_KEY:                   privateKey,
			// PUBLIC_KEY:                    publicKey,
		},
		ENVIRONMENT: Environment(viper.GetString("ENVIRONMENT")),
	}

	return Config, nil
}
func loadPrivateKey(filePath string) (*rsa.PrivateKey, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read private key: %w", err)
	}
	block, _ := pem.Decode(data)
	if block == nil {
		return nil, fmt.Errorf("invalid private key PEM")
	}
	if key, err := x509.ParsePKCS1PrivateKey(block.Bytes); err == nil {
		return key, nil
	}
	if key, err := x509.ParsePKCS8PrivateKey(block.Bytes); err == nil {
		if rsaKey, ok := key.(*rsa.PrivateKey); ok {
			return rsaKey, nil
		}
	}
	return nil, fmt.Errorf("unsupported private key format")
}

func loadPublicKey(filePath string) (*rsa.PublicKey, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read public key: %w", err)
	}
	block, _ := pem.Decode(data)
	if block == nil {
		return nil, fmt.Errorf("invalid public key PEM")
	}
	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("parse public key error: %w", err)
	}
	if rsaKey, ok := key.(*rsa.PublicKey); ok {
		return rsaKey, nil
	}
	return nil, fmt.Errorf("not an RSA public key")
}
