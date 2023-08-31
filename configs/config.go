package configs

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"log"
)
import "github.com/spf13/viper"

type Server struct {
	AppMode  string `yaml:"AppMode"`
	HttpPort string `yaml:"HttpPort"`
}

type Config struct {
	Server Server
}

var Conf Config

func InitConfig() {
	viper.SetConfigName("config") //找寻文件的名字
	viper.SetConfigType("yaml")   // 找寻文件的类型
	//
	viper.AddConfigPath("configs/dev") //从当前目录下的哪个文件夹找寻，
	//viper.AddConfigPath(".")        //.代表当前文件夹找寻，可以多个目录找寻，生成数组
	err := viper.ReadInConfig() //读取配置文件
	if err != nil {
		if v, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(v)
		} else {
			panic(fmt.Errorf("read config error:%s", err))
		}
	}

	//1.序列化配置
	err = viper.Unmarshal(&Conf)
	if err != nil {
		panic(fmt.Errorf("config unmarshal error:%s", err))
	}
	//打印配置信息
	byteConfig, _ := jsoniter.Marshal(Conf)
	log.Println(string(byteConfig))

}
