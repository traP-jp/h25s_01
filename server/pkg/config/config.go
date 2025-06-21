package config

import (
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
)

type AWSConfig struct {
	Region     string
	BucketName string
	AccessKey  string
	SecretKey  string
	Endpoint   string
	PathStyle  bool
}

func getEnv(key, defaultValue string) string {
	v, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	return v
}

func AppAddr() string {
	return getEnv("APP_ADDR", ":8080")
}

func AWS() *AWSConfig {
	pathStyleStr := getEnv("AWS_S3_FORCE_PATH_STYLE", "false")
	pathStyle := pathStyleStr == "true"

	return &AWSConfig{
		Region:     getEnv("AWS_REGION", "ap-northeast-1"),
		BucketName: getEnv("AWS_S3_BUCKET", "my-app-bucket"),
		AccessKey:  getEnv("AWS_ACCESS_KEY_ID", ""),
		SecretKey:  getEnv("AWS_SECRET_ACCESS_KEY", ""),
		Endpoint:   getEnv("AWS_ENDPOINT_URL", ""),
		PathStyle:  pathStyle,
	}
}

func MySQL() *mysql.Config {
	c := mysql.NewConfig()

	c.User = getEnv("DB_USER", "root")
	c.Passwd = getEnv("DB_PASS", "pass")
	c.Net = getEnv("DB_NET", "tcp")
	c.Addr = fmt.Sprintf(
		"%s:%s",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_PORT", "3306"),
	)
	c.DBName = getEnv("DB_NAME", "app")

	// NeoShowcase対応
	c.User = getEnv("NS_MARIADB_USER", "root")
	c.Passwd = getEnv("NS_MARIADB_PASSWORD", "pass")
	c.Addr = fmt.Sprintf(
		"%s:%s",
		getEnv("NS_MARIADB_HOSTNAME", "localhost"),
		getEnv("NS_MARIADB_PORT", "3306"),
	)
	c.DBName = getEnv("NS_MARIADB_DATABASE", "app")

	c.Collation = "utf8mb4_general_ci"
	c.ParseTime = true
	c.AllowNativePasswords = true

	return c
}
