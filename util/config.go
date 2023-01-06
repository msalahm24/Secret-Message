package util

import "github.com/spf13/viper"

type Config struct {
	TWILIO_ACCOUNT_SID      string `mapstructure:"TWILIO_ACCOUNT_SID"`
	TWILIO_AUTH_TOKEN      string `mapstructure:"TWILIO_AUTH_TOKEN"`
	SreverAddress string `mapstructure:"SERVER_ADDRESS"`
}

func LoadConfig(path string)(config Config,err error){
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil{
		return
	}
	err = viper.Unmarshal(&config)
	if err != nil{
		return
	}
	return
}
