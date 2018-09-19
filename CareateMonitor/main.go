package main

import "GOTests/CareateMonitor/monitor"

func main() {
	mongoWriter := monitor.InitMongoDBConnection("mongodb://CareateAdmin:Careate2016!@127.0.0.1:27017/CareateMonitorDB", "CareateMonitorDB", "test")
	a := monitor.InitMonitor(mongoWriter)
	go a.MonitorLoop()
	a.Write()
	// if session, err := mgo.Dial("mongodb://CareateAdmin:Careate2016!@127.0.0.1:27017/CareateMonitorDB"); err != nil {
	// 	fmt.Println(err.Error())
	// 	fmt.Println(1)
	// } else {
	// 	fmt.Println(session)
	// 	fmt.Println(1)
	// }

	// m, _ := mem.VirtualMemory()
	// d, _ := disk.Usage("/")
	//n, _ := net.Connections("tcp")

	// established := 0
	// timewait := 0
	// closewait := 0
	// for _, item := range n {
	// 	if item.Status == "ESTABLISHED" {
	// 		established++
	// 	}
	// 	if item.Status == "TIME_WAIT" {
	// 		timewait++
	// 	}
	// 	if item.Status == "CLOSE_WAIT" {
	// 		closewait++
	// 	}
	// }
	// almost every return value is a struct
	// fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", m.Total, m.Free, m.UsedPercent)

	// // convert to JSON. String() is also implemented
	// fmt.Println(d)
	// fmt.Println(n)
	// fmt.Println(established)
	// fmt.Println(timewait)
	// fmt.Println(closewait)
}
