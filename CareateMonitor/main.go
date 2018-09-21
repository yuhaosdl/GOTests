package main

import (
	"GOTests/CareateMonitor/monitor"
	"encoding/json"
	"os"
	"path/filepath"
)

func main() {
	conf := GetConf("CareateMonitor.json")
	mongoWriter := monitor.InitMongoDBConnection(conf.ConnectionStr, conf.DBName, conf.CollectionName)
	defer mongoWriter.Close()
	a := monitor.InitMonitor(mongoWriter, conf.ServerName, conf.Interval, conf.SupervisePort)
	go a.MonitorLoop()
	a.Write()

}

//GetConf ： 获取配置文件
func GetConf(fileName string) (conf *Configuration) {
	execpath, err := os.Executable() // 获得程序路径
	if err != nil {
		panic("获取程序路径出错" + err.Error())
	}
	configfile := filepath.Join(filepath.Dir(execpath), "./"+fileName)
	file, err := os.Open(configfile)
	defer file.Close()
	if err != nil {
		panic("读取配置文件出错" + err.Error())
	}
	decoder := json.NewDecoder(file)
	conf = &Configuration{}
	err = decoder.Decode(&conf)
	if err != nil {
		panic("Decode配置文件出错" + err.Error())
	}
	return
}

//Configuration : 配置文件
type Configuration struct {
	ServerName     string
	ConnectionStr  string
	DBName         string
	CollectionName string                  //
	Interval       int                     //时间间隔秒
	SupervisePort  []monitor.SupervisePort //监控服务
}
