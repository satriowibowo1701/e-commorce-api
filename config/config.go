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

}

var JWT_SECRET = ""
var JWT_EXPIRATION_DURATION time.Duration
var PORT = ""

var CONFIG_SMTP_HOST = ""
var CONFIG_SMTP_PORT = 0
var CONFIG_SENDER_NAME = ""
var CONFIG_AUTH_EMAIL = ""
var CONFIG_AUTH_PASSWORD = ""
var DATABASE_NAME = ""
var MAX_IDLE_CONNS = 0
var MAX_OPEN_CONNS = 0
var CONNMAXLIFETIME time.Duration
var CONNMAXIDLETIME time.Duration

var Mutex sync.Mutex
