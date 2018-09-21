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
	supervisePort   []SupervisePort
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
func InitMonitor(writer Writer, serverName string, interval int, supervisePort []SupervisePort) (monitor *Monitor) {
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
		supervisePort:   supervisePort,
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
		go monitor.getHeartbeat()
		time.Sleep(time.Duration(monitor.monitorInterval) * time.Second)
	}
}

//getHeartbeat : 获取自身状态
func (monitor *Monitor) getHeartbeat() {
	monitorMessage := &MonitorMessage{
		Key:         "MonitorService",
		Value:       "1",
		Description: monitor.ServerName + "-MonitorService",
		InDateTime:  getTime(),
		ServerName:  monitor.ServerName,
		ServerIP:    monitor.IP,
		CreateTime:  monitor.CreateTime,
	}
	monitor.toMessageChanel(monitorMessage)
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
		monitor.toMessageChanel(monitorMessage)
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
		monitor.toMessageChanel(monitorMessage)
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
		monitor.toMessageChanel(monitorMessage)

	}
}

func (monitor *Monitor) getNetStat() {
	n, _ := gopsutilnet.Connections("tcp")
	established := 0
	timewait := 0
	closewait := 0
	for _, item := range n {
		// for _, i := range monitor.supervisePort {
		// 	if item.Laddr.Port == i.Value && item.Status == "LISTEN" {
		// 		i.Status = true
		// 	}
		// }
		for i := 0; i < len(monitor.supervisePort); i++ {
			if item.Laddr.Port == monitor.supervisePort[i].Value && item.Status == "LISTEN" {
				monitor.supervisePort[i].Status = true
			}
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
	go monitor.getServiceStat()
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
	monitor.toMessageChanel(monitorMessage)
}

func (monitor *Monitor) getServiceStat() {
	for _, item := range monitor.supervisePort {
		monitorMessage := &MonitorMessage{
			Key:         item.Key,
			Value:       "close",
			Description: item.Key,
			InDateTime:  getTime(),
			ServerName:  monitor.ServerName,
			ServerIP:    monitor.IP,
			CreateTime:  monitor.CreateTime,
		}
		if item.Status {
			monitorMessage.Value = "open"
		}
		monitor.toMessageChanel(monitorMessage)
	}

}

//GetTime : 获取当前时间字符串
func getTime() (dataTimeStr string) {
	dataTimeStr = time.Now().Format("2006-01-02 15:04:05")
	return
}

// float64ToString : float64转 string
func float64ToString(f float64) (str string) {
	str = strconv.FormatFloat(f, 'f', 2, 64)
	return
}

//toMessageChannel : 写入MessageChannel里
func (monitor *Monitor) toMessageChanel(monitorMessage *MonitorMessage) {
	select {
	case monitor.messageChanel <- monitorMessage:
	case <-time.After(1 * time.Second):
		fmt.Println("写入messageChanel超时")
	}
}

//SupervisePort : 监控服务
type SupervisePort struct {
	Key    string
	Value  uint32
	Status bool
}
