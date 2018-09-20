package monitor

import (
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	gopsutilnet "github.com/shirou/gopsutil/net"
)

//MonitorMessage : 监控消息
type MonitorMessage struct {
	ServerName  string
	ServerIP    string
	InDateTime  string
	Key         string
	Value       string
	Description string
	CreateTime  string
}

//Monitor ：监控主体
type Monitor struct {
	IP              string
	ServerName      string
	messageChanel   chan *MonitorMessage
	writer          Writer
	CreateTime      string
	monitorInterval int
}

//Writer : 接口
type Writer interface {
	write(monitor *MonitorMessage) (err error)
}

//Write : 写入数据
func (monitor *Monitor) Write() {
	for {
		data := <-monitor.messageChanel
		go func() {
			if err := monitor.writer.write(data); err != nil {
				fmt.Println(err.Error())
			}
		}()
	}
}

//InitMonitor ： 初始化
func InitMonitor(writer Writer, serverName string, interval int) (monitor *Monitor) {
	var addr string
	if addrs, err := net.InterfaceAddrs(); err == nil {
		for _, address := range addrs {

			// 检查ip地址判断是否回环地址
			if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					addr = ipnet.IP.String()
					break
				}
			}
		}
	}
	monitor = &Monitor{
		messageChanel:   make(chan *MonitorMessage, 500),
		IP:              addr,
		ServerName:      serverName,
		writer:          writer,
		monitorInterval: interval,
	}
	return
}

//MonitorLoop : 循环监控 间隔X秒
func (monitor *Monitor) MonitorLoop() {
	for {
		monitor.CreateTime = getTime()
		go monitor.getCPUStat()
		go monitor.getDiskStat()
		go monitor.getMemStat()
		go monitor.getNetStat()
		time.Sleep(time.Duration(monitor.monitorInterval) * time.Second)
	}
}

//getMemStat : 循环获取内存状态
func (monitor *Monitor) getMemStat() {
	if m, err := mem.VirtualMemory(); err == nil {
		monitorMessage := &MonitorMessage{
			Key:         "RAM",
			Value:       float64ToString(m.UsedPercent),
			Description: "内存",
			InDateTime:  getTime(),
			ServerName:  monitor.ServerName,
			ServerIP:    monitor.IP,
			CreateTime:  monitor.CreateTime,
		}
		monitor.messageChanel <- monitorMessage
	}
}

//getDiskStat : 循环获取内存状态
func (monitor *Monitor) getDiskStat() {
	if d, err := disk.Usage("/"); err == nil {
		monitorMessage := &MonitorMessage{
			Key:         "DISK",
			Value:       float64ToString(d.UsedPercent),
			Description: "硬盘",
			InDateTime:  getTime(),
			ServerName:  monitor.ServerName,
			ServerIP:    monitor.IP,
			CreateTime:  monitor.CreateTime,
		}
		monitor.messageChanel <- monitorMessage
	}
}

//getCPUStat : 循环获取CPU状态
func (monitor *Monitor) getCPUStat() {
	if cp, err := cpu.Percent(time.Second, false); err == nil {
		monitorMessage := &MonitorMessage{
			Key:         "CPU",
			Value:       float64ToString(cp[0]),
			Description: "CPU",
			InDateTime:  getTime(),
			ServerName:  monitor.ServerName,
			ServerIP:    monitor.IP,
			CreateTime:  monitor.CreateTime,
		}
		select {
		case monitor.messageChanel <- monitorMessage:
		case <-time.After(1 * time.Second):
			fmt.Println("写入messageChanel超时")

		}

	}
}

func (monitor *Monitor) getNetStat() {
	n, _ := gopsutilnet.Connections("tcp")
	established := 0
	timewait := 0
	closewait := 0
	redis := false
	elasticsearch := false
	for _, item := range n {
		if item.Laddr.Port == 6379 && item.Status == "LISTEN" {
			redis = true
		}
		if item.Laddr.Port == 9200 && item.Status == "LISTEN" {
			elasticsearch = true
		}
		if item.Status == "ESTABLISHED" {
			established++
		}
		if item.Status == "TIME_WAIT" {
			timewait++
		}
		if item.Status == "CLOSE_WAIT" {
			closewait++
		}
	}
	go monitor.getRedisStat(redis)
	go monitor.getElasticsearchStat(elasticsearch)
	go monitor.getStat("Net_Established", strconv.Itoa(established))
	go monitor.getStat("Net_Timewait", strconv.Itoa(timewait))
	go monitor.getStat("Net_Closewait", strconv.Itoa(closewait))
}
func (monitor *Monitor) getStat(key string, value string) {
	monitorMessage := &MonitorMessage{
		Key:         key,
		Value:       value,
		Description: key,
		InDateTime:  getTime(),
		ServerName:  monitor.ServerName,
		ServerIP:    monitor.IP,
		CreateTime:  monitor.CreateTime,
	}
	monitor.messageChanel <- monitorMessage
}
func (monitor *Monitor) getRedisStat(n bool) {
	redisMonitorMessage := &MonitorMessage{
		Key:         "Service_redis",
		Value:       "close",
		Description: "redis服务",
		InDateTime:  getTime(),
		ServerName:  monitor.ServerName,
		ServerIP:    monitor.IP,
		CreateTime:  monitor.CreateTime,
	}
	if n {
		redisMonitorMessage.Value = "open"
	}
	monitor.messageChanel <- redisMonitorMessage
}

func (monitor *Monitor) getElasticsearchStat(n bool) {
	monitorMessage := &MonitorMessage{
		Key:         "Service_elasticsearch",
		Value:       "close",
		Description: "elasticsearch服务",
		InDateTime:  getTime(),
		ServerName:  monitor.ServerName,
		ServerIP:    monitor.IP,
		CreateTime:  monitor.CreateTime,
	}
	if n {
		monitorMessage.Value = "open"
	}
	monitor.messageChanel <- monitorMessage
}

//GetTime : 获取当前时间字符串
func getTime() (dataTimeStr string) {
	dataTimeStr = time.Now().Format("2006-01-02 15:04:05")
	return
}

// float64ToString : float64转 string
func float64ToString(f float64) (str string) {
	str = strconv.FormatFloat(f, 'f', 2, 64) + "%"
	return
}
