# weather-mane
Simple web server for personal weather information from AWN



Uses the following technology
- GORM
- sqlite

## Installation
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
go get github.com/lrosenman/ambient     //Ambient Libray
go get github.com/kylelemons/go-gypsy   //YAML parser

## config.yaml file

You must create a config yaml file with 
```
env: [DEV, PROD]
awnApiKeys:
  apiKey: string
  appKey: string
```