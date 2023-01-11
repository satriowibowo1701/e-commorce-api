package config

import (
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load()

	duration, _ := strconv.Atoi(os.Getenv("JWT_EXPIRATION_DURATION"))
	Idlecons, _ := strconv.Atoi(os.Getenv("MAX_IDLE_CONNS"))
	maxopencons, _ := strconv.Atoi(os.Getenv("MAX_OPEN_CONNS"))
	conmaxlifetime, _ := strconv.Atoi(os.Getenv("CONNMAXLIFETIME"))
	conmaxidletime, _ := strconv.Atoi(os.Getenv("CONNMAXIDLETIME"))
	MAX_OPEN_CONNS = maxopencons
	MAX_IDLE_CONNS = Idlecons
	CONNMAXLIFETIME = time.Duration(conmaxlifetime) * time.Hour
	CONNMAXIDLETIME = time.Duration(conmaxidletime) * time.Minute
	DATABASE_NAME = os.Getenv("DATABASE_NAME")
	JWT_EXPIRATION_DURATION = time.Hour * time.Duration(duration)
	JWT_SECRET = os.Getenv("JWT_SECRET")
	PORT = os.Getenv("PORT")
	smptpport, _ := strconv.Atoi(os.Getenv("CONFIG_SMTP_PORT"))
	CONFIG_SMTP_HOST = os.Getenv("CONFIG_SMTP_HOST")
	CONFIG_SMTP_PORT = smptpport
	CONFIG_SENDER_NAME = os.Getenv("CONFIG_SENDER_NAME")
	CONFIG_AUTH_EMAIL = os.Getenv("CONFIG_AUTH_EMAIL")
	CONFIG_AUTH_PASSWORD = os.Getenv("CONFIG_AUTH_PASSWORD")
	DBUSER = os.Getenv("DB_USER")
	DBPASS = os.Getenv("DB_PASSWORD")
	DBHOST = os.Getenv("DB_HOST")
	DBNAME = os.Getenv("DB_NAME")
	DBPORT = os.Getenv("DB_PORT")
}

var (
	DBUSER                  string
	DBPASS                  string
	DBHOST                  string
	DBNAME                  string
	DBPORT                  string
	JWT_SECRET              = ""
	JWT_EXPIRATION_DURATION time.Duration
	PORT                    = ""
	CONFIG_SMTP_HOST        = ""
	CONFIG_SMTP_PORT        = 0
	CONFIG_SENDER_NAME      = ""
	CONFIG_AUTH_EMAIL       = ""
	CONFIG_AUTH_PASSWORD    = ""
	DATABASE_NAME           = ""
	MAX_IDLE_CONNS          = 0
	MAX_OPEN_CONNS          = 0
	CONNMAXLIFETIME         time.Duration
	CONNMAXIDLETIME         time.Duration
	Mutex                   sync.Mutex
)
