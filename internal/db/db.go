package db

import (
	"fmt"

	// needed for "postgres" driver
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"

	"github.com/manoj-gupta/glance/internal/utils"
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
		host:     utils.GetEnv("POSTGRES_HOST", "database"),
		port:     utils.GetEnvAsInt("POSTGRES_PORT", 5432),
		user:     utils.GetEnv("POSTGRES_USER", "postgres"),
		password: utils.GetEnv("POSTGRES_PASSWORD", "postgres"),
		dbname:   utils.GetEnv("POSTGRES_DB", "testdb"),
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
		fmt.Printf("Cannot connect to %s database! %s\n", cfg.dbname, connString)
		return nil, err
	}

	err = db.DB().Ping()
	if err != nil {
		return nil, err
	}
	DB = db

	fmt.Printf("Connected to %s database!\n", cfg.dbname)
	return db, err
}

// DeInit .. deinitialize database
func DeInit(db *gorm.DB) {
	DB = nil
	db.Close()
}
