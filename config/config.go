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
	JWT_EXPIRATION_DURATION = time.Hour * time.Duration(duration)
	JWT_SECRET = os.Getenv("JWT_SECRET")
	PORT = os.Getenv("PORT")
	DBURL = os.Getenv("DBURL")
}

var (
	DBURL                   string
	JWT_SECRET              = ""
	JWT_EXPIRATION_DURATION time.Duration
	PORT                    = ""
	MAX_IDLE_CONNS          = 0
	MAX_OPEN_CONNS          = 0
	CONNMAXLIFETIME         time.Duration
	CONNMAXIDLETIME         time.Duration
	Mutex                   sync.Mutex
)
