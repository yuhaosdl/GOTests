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

	established := 0
	timewait := 0
	closewait := 0
	for _, item := range n {
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
	// almost every return value is a struct
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", m.Total, m.Free, m.UsedPercent)

	// convert to JSON. String() is also implemented
	fmt.Println(d)
	fmt.Println(n)
	fmt.Println(cp)
	fmt.Println(established)
	fmt.Println(timewait)
	fmt.Println(closewait)
}
