package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type ccommon struct {
	Baseip     string `yaml:"baseip"`
	Port       string `yaml:"port"`
	Xgfaceaddr string `yaml:"xgfaceaddr"`
	Deadline   int    `yaml:"deadline"`
}
type ccase struct {
	Name  string `yaml:"name"`
	Mysql cmysql `yaml:"mysql"`
	Redis credis `yaml:"redis"`
	Staff cstaff `yaml:"staff"`
}
type cmysql struct {
	Host        string `yaml:"host"`
	User        string `yaml:"user"`
	Password    string `yaml:"password"`
	Database    string `yaml:"database"`
	GormLogMode bool   `yaml:"gormlogmode"`
}
type credis struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}
type cstaff struct {
	Defaultlibrary []string `yaml:"defaultlibrary"`
	Xgindexaddr    string   `yaml:"xgindexaddr"`
	Staffhold      float32  `yaml:"staffhold"`
	Imgbasepath    string   `yaml:"imgbasepath"`
	Accessbasepath string   `yaml:"accessbasepath"`
	Facefilepath   string   `yaml:"facefilepath"`
	Peopleimgpath  string   `yaml:"peopleimgpath"`
}
type Services struct {
	Common ccommon `yaml:"common"`
	Case   []ccase `yaml:"case"`
}
type Conf struct {
	Version  string   `yaml:"version"`
	Services Services `yaml:"services"`
}

var Config Conf

func loadConf() {
	yamlFile, err := ioutil.ReadFile("./conf.yml")
	if err != nil {
		log.Fatalf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		log.Fatalf("yamlFile.Unmarshal: %v", err)
	}
	fmt.Printf("%+v\n", Config.String())

	//
	//f, err := yaml.Marshal(Config)
	//ioutil.WriteFile("./t3.yml", f, 0666)
}
func (conf *Conf) String() string {
	b, err := json.Marshal(*conf)
	if err != nil {
		return fmt.Sprintf("%+v", *conf)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "    ")
	if err != nil {
		return fmt.Sprintf("%+v", *conf)
	}
	return out.String()
}
func main() {
	loadConf()
}
