package environment

import (
	"errors"
	"flag"
	ascii_art "go-kit-base/infrastructure/configuration/ascii-art"
	"log"

	"github.com/spf13/viper"
)

//VARIABLE DEFINITION
//Endpoint mapping
//var Endpoints map[string]config.Uri

//Environment
var environment *Environment

func init() {
	//CONFIGURATION STARTUP
	//Flags definition
	basePath := flag.String("base_path", "infrastructure/configuration/static/", "The path to application base configurations")
	envPath := flag.String("env_path", "infrastructure/configuration/static/", "The path to application environment configurations")
	baseFile := flag.String("base_file_name", "application", "Base config file")
	flag.Parse()

	//Base and environment config read
	baseConf := loadBaseConfiguration(*basePath, *baseFile)
	envConf := loadEnvironmentConfiguration(*envPath, *baseFile, baseConf)

	//Environment loading
	environment = NewEnvironment(envConf)

	//Welcome message
	displayWelcomeMessage(baseConf)

	//Vault loading
	//TODO Set the vault configuration here

}

//CONFIGURATION LOADING

//Base config environment
func loadBaseConfiguration(basePath, baseFile string) *viper.Viper {
	baseConf := viper.New()
	baseConf.SetConfigName(baseFile)
	baseConf.SetConfigType("yml")
	baseConf.AddConfigPath(basePath)
	baseConf.SetDefault("server.port", "90")

	//Base config reading
	if err := baseConf.ReadInConfig(); err != nil {
		log.Fatalf("Not able to read base config, %s", err)
	}
	return baseConf
}

//Environment config environment
func loadEnvironmentConfiguration(basePath string, baseFile string, baseConf *viper.Viper) *viper.Viper {
	envConf := viper.New()
	envProfile := baseFile + "-" + baseConf.GetString("profiles.active")
	envConf.SetConfigName(envProfile)
	envConf.SetConfigType("yml")
	envConf.AddConfigPath(basePath)

	//Configuration reading
	if err := envConf.ReadInConfig(); err != nil {
		log.Fatalf("Not able to read environment config, %s", err)
	}
	return envConf
}

//WELCOME DISPLAY

//Displays the welcome message with specific artistic style font
func displayWelcomeMessage(baseConf *viper.Viper) {
	if (baseConf.GetString("log.startup-phrase.title.value") != "") &&
		(baseConf.GetString("log.startup-phrase.title.font") != "") {
		ascii_art.WelcomeTitle(
			baseConf.GetString("log.startup-phrase.title.value"),
			baseConf.GetString("log.startup-phrase.title.font"))
	}
	if (baseConf.GetString("log.startup-phrase.message.value") != "") &&
		(baseConf.GetString("log.startup-phrase.message.font") != "") {
		ascii_art.WelcomeMessage(
			baseConf.GetString("log.startup-phrase.message.value"),
			baseConf.GetString("log.startup-phrase.message.font"))
	}
}

//Getters
func New() (*Environment, error) {
	if environment != nil {
		return environment, nil
	} else {
		return environment, errors.New("The environment is empty!") // TODO Sacar error a constante
	}

}
