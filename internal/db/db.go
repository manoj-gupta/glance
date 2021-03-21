package db

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	// needed for "postgres" driver
	_ "github.com/lib/pq"

	"github.com/jinzhu/gorm"
)

// DB .. global variable accessed by models
var DB *gorm.DB

// Config ... represents db configuration
type Config struct {
	host     string
	port     int
	user     string
	password string
	dbname   string
}

// New returns a new Config struct
func newConfig() *Config {
	conf := Config{
		host:     getEnv("DB_HOST", "localhost"),
		port:     getEnvAsInt("DB_PORT", 5432),
		user:     getEnv("DB_USERNAME", "postgres"),
		password: getEnv("DB_PASSWORD", "postgres"),
		dbname:   getEnv("DB_NAME", "testdb"),
	}
	return &conf
}

// getURL ... returns DB connection URL
func getURL(dbConfig *Config) string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbConfig.host,
		dbConfig.port,
		dbConfig.user,
		dbConfig.password,
		dbConfig.dbname,
	)
}

// Init .. initialize database
func Init() (*gorm.DB, error) {
	cfg := newConfig()
	connString := getURL(cfg)
	db, err := gorm.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	err = db.DB().Ping()
	if err != nil {
		return nil, err
	}
	DB = db

	fmt.Println("db: successfully connected!")
	return db, err
}

// DeInit .. deinitialize database
func DeInit(db *gorm.DB) {
	DB = nil
	db.Close()
}

// getEnv ... read an environment variable or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

// getEnvAsInt ... read an environment variable into integer or return a default value
func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

// getEnvAsBool .. read an environment variable into a bool or return default value
func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}

// getEnvAsSlice ... read an environment variable into a string slice or return default value
func getEnvAsSlice(name string, defaultVal []string, sep string) []string {
	valStr := getEnv(name, "")

	if valStr == "" {
		return defaultVal
	}

	val := strings.Split(valStr, sep)
	return val
}
