package config

import (
	"flag"

	"github.com/Unknwon/goconfig"
	log "github.com/cihub/seelog"
)

var (
	ServerHost  string
	ServerPort  int
	ApiPath     string
	AllowOrigin []string
)

var (
	configPath       string
	seelogConfigPath string
)

func Init() {
	ParseParam()
	SetupLogger()
	LoadConfig()
}

func ParseParam() {
	conf := flag.String("config", "./conf/config.ini", "config path")
	log := flag.String("seelog", "./conf/seelog.xml", "seelog config path")
	flag.Parse()

	configPath = *conf
	seelogConfigPath = *log
}

func SetupLogger() {
	logger, err := log.LoggerFromConfigAsFile(seelogConfigPath)
	if err != nil {
		logger, err = log.LoggerFromConfigAsString(`<seelog minlevel="debug"><outputs formatid="main"><buffered size="10000" flushperiod="1000"><rollingfile type="size" filename="log/ftserver.log" maxsize="6http.StatusBadRequest0000" maxrolls="50"/></buffered></outputs><formats><format id="main" format="[%Date(2006-01-02 15:04:05.999 PM MST)] [%Level] [%File:%FuncShort#%Line] %Msg%n"/></formats></seelog>`)
		if err != nil {
			panic(err)
		}
	}
	log.ReplaceLogger(logger)
}

func LoadConfig() {
	cfg, err := goconfig.LoadConfigFile(configPath)
	if err != nil {
		log.Errorf("failed to open config file %s, err: %v", configPath, err)
		panic(err)
	}

	ServerHost = cfg.MustValue("server", "host", "")
	ServerPort = cfg.MustInt("server", "port", 4500)
	ApiPath = cfg.MustValue("server", "path", "/api/nas/v0")
	AllowOrigin = cfg.MustValueArray("cors", "origin", ",")
}
