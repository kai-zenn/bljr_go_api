package configs

import (
    "log"
    "os"
    "github.com/joho/godotenv"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
    err := godotenv.Load()
    if err != nil {
        err = godotenv.Load("../../.env")
        if err != nil {
            log.Println("Error loading .env file, menggunakan env bawaan sistem")
        }
    }

    var dsn string
    if os.Getenv("DATABASE_URL") != "" {
  		dsn = os.Getenv("DATABASE_URL")
  	} else {
  		dsn = "host=" + os.Getenv("DB_HOST") +
  			" user=" + os.Getenv("DB_USER") +
  			" password=" + os.Getenv("DB_PASSWORD") +
  			" dbname=" + os.Getenv("DB_NAME") +
  			" port=" + os.Getenv("DB_PORT") +
  			" sslmode=" + os.Getenv("SSL_MODE")
  	}

    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    log.Println("Database connection established")
}
