package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

var (
	cfg *ini.File

	RunMode string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize  int
	JwtSecret string

	DBType         string
	DBName         string
	DBUser         string
	DBPwd          string
	DBHost         string
	DBTable_prefix string
)

func init() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	loadBase()
	loadServer()
	loadDB()
	loadApp()
}

func loadBase() {
	RunMode = cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func loadServer() {
	sec, err := cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	HTTPPort = sec.Key("HTTP_PORT").MustInt(5000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func loadDB() {
	sec, err := cfg.GetSection("database")
	if err != nil {
		log.Fatalf("Fail to get section 'database': %v", err)
	}

	DBType = sec.Key("TYPE").MustString("mariadb")
	DBName = sec.Key("NAME").MustString("blog")
	DBUser = sec.Key("USER").MustString("root")
	DBPwd = sec.Key("PASSWORD").MustString("Mariadb123")
	DBHost = sec.Key("HOST").MustString("127.0.0.1:3305")
	DBTable_prefix = sec.Key("TABLE_PREFIX").MustString("blog_")
}

func loadApp() {
	sec, err := cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	JwtSecret = sec.Key("JWT_SECRET").MustString("!@#$%^&*()")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}
