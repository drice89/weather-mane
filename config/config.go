package config

import (
	"os"
    "log"

	"github.com/kylelemons/go-gypsy/yaml"
)

// This struct is really to document whats expected in the config.yml file
type awnApiKeys struct {
    apiKey string
    appKey string
}
type config struct {
    env string
    awnKeys awnApiKeys
}


func InitConfig() {
    f, err := yaml.ReadFile("config.yaml")
    if err != nil {
        log.Panicln("ERROR: No config.yml file found")
    }
    configureAppFromFile(f)
}

func configureAppFromFile(f *yaml.File) {
    conf := config{}

    env, err := f.Get("env")
    if err != nil || env == "" {
        log.Printf("Error: %+v", err)
        log.Println("WARNING: Running in development mode")
        env = "dev"
    }
    conf.env = env

    //awn api keys
    awnApiKey, _ := f.Get("awnApiKeys.apiKey")
    awnAppKey, _ := f.Get("awnApiKeys.appKey")


    if awnApiKey != "" && awnAppKey != "" {
        aak := awnApiKeys {
            appKey: awnApiKey,
            apiKey: awnAppKey,
        }
        conf.awnKeys = aak
    } else {
        log.Println("WARNING: Invalid AWN Keys set in config.yml")
    }
    
    conf.setEnvVars()
}

func (c config) setEnvVars() {
    // the keys will be set to a blank string if they are not set in the config yml file
    os.Setenv("AWN_APP_KEY", c.awnKeys.appKey)
    os.Setenv("AWN_API_KEY", c.awnKeys.apiKey)
    os.Setenv("ENV", c.env)
}