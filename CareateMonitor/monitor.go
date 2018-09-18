package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

func main() {
	m, _ := mem.VirtualMemory()
	d, _ := disk.Usage("/")
	cp, _ := cpu.Percent(time.Second, false)
	n, _ := net.Connections("tcp")

	// almost every return value is a struct
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", m.Total, m.Free, m.UsedPercent)

	// convert to JSON. String() is also implemented
	fmt.Println(d)
	fmt.Println(n)
	fmt.Println(cp)
}
