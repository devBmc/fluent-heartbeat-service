package config

import(
	"github.com/spf13/viper"
	"sync"
	"fmt"
)
var once sync.Once
var CONFIG Config

type Config struct {
    FileStorePath      string `mapstructure:"FILESTORE_PATH"`
    FileStoreName      string `mapstructure:"FILESTORE_NAME"`
    LogPath       string `mapstructure:"LOG_FILE_PATH"`
	PrimaryHost       string `mapstructure:"END_POINT_PRIMARY_HOST"`
	Port           string `mapstructure:"ENG_POINT_PORT"`
	SecondaryHosts      []string `mapstructure:"END_POINT_SECONDARY_HOSTS"`
	ScheduleTimeInteval int `mapstructure:"TIME_INTERVAL_MINUTES"`
	SSLKeyPath     string `mapstructure:"SSL_KEY_PATH"`
	SSLCERTPath     string `mapstructure:"SSL_CERT_PATH"`
}

func LoadConfig(path string) (config Config, err error) {
    viper.AddConfigPath(path)
    viper.SetConfigName("app")
    //viper.SetConfigType("env")	
	//viper.SetConfigType("properties")
	viper.SetConfigFile("C:\\git\\DSOM-ADE\\fluent-heartbeat-service\\app.properties")
fmt.Printf("Using config: %s\n", viper.ConfigFileUsed())
    viper.AutomaticEnv()
    err = viper.ReadInConfig()
    if err != nil {
        return
    }
    err = viper.Unmarshal(&config)
	if err==nil{
		CONFIG=config
	}
    return CONFIG,err
}

func GetInstance(path string)(config Config, err error){
	once.Do(func() {
		LoadConfig(path)
		 
	})
	return
}

func GetConfig() Config{
	return CONFIG
}