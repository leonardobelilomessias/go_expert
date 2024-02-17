package configs

import (
	"github.com/spf13/viper"
	"github.com/go-chi/jwtauth"
)


type conf struct{
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBHost string `mapstructure:"DB_HOST"`
	DBUser string `mapstructure:"DB_USER"`
	DBPort string `mapstructure:"DB_PORT"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName string `mapstructure:"DB_NAME"`
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
	JwtSecret string `mapstructure:"JWT_SECRET"`
	JwtExperesIn int `mapstructure:"JWT_EXPERESIN"`
	TokenAuth *jwtauth.JWTAuth 
}
func LoadConfig(path string)(*conf , error){
	var cfg *conf
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err:= viper.ReadInConfig()
	if err != nil{
		panic(err)
	}
	
	err = viper.Unmarshal(&cfg)
	if err != nil{
		panic(err)
	}
	cfg.TokenAuth = jwtauth.New("HS256",[]byte(cfg.JwtSecret),nil)
	return cfg, err
}