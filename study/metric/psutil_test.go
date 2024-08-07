package metric

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	gnet "github.com/shirou/gopsutil/net"
	"net"
	"testing"
	"time"
)

func TestCpuInfo(t *testing.T) {
	cpuInfos, err := cpu.Info()
	if err != nil {
		fmt.Printf("get cpuInfo failed.%+v", err)
		return
	}
	for _, ci := range cpuInfos {
		fmt.Println(ci)
	}
	// CPU使用率
	for {
		percent, _ := cpu.Percent(time.Second, false)
		fmt.Printf("cpu percent:%v\n", percent)
	}
}

func TestCpuLoad(t *testing.T) {
	info, _ := load.Avg()
	fmt.Printf("%v\n", info)
}

func TestMemoInfo(t *testing.T) {
	memInfo, _ := mem.VirtualMemory()
	fmt.Printf("mem info:%v\n", memInfo)
}

func TestHostInfo(t *testing.T) {
	hInfo, _ := host.Info()
	fmt.Printf("host info:%v uptime:%v boottime:%v\n", hInfo, hInfo.Uptime, hInfo.BootTime)
}

func TestNetIo(t *testing.T) {
	info, _ := gnet.IOCounters(true)
	for index, v := range info {
		fmt.Printf("%v:%v send:%v recv:%v\n", index, v, v.BytesSent, v.BytesRecv)
	}
}

func TestDiskInfo(t *testing.T) {
	parts, err := disk.Partitions(true)
	if err != nil {
		fmt.Printf("get Partitions failed, err:%v\n", err)
		return
	}
	for _, part := range parts {
		fmt.Printf("part:%v\n", part.String())
		diskInfo, _ := disk.Usage(part.Mountpoint)
		fmt.Printf("disk info:used:%v free:%v\n", diskInfo.UsedPercent, diskInfo.Free)
	}

	ioStat, _ := disk.IOCounters()
	for k, v := range ioStat {
		fmt.Printf("%v:%v\n", k, v)
	}
}

func TestLocalIp(t *testing.T) {
	localIP, err := getLocalIP()
	if err != nil {
		fmt.Printf("get localIp failed. %+v", err)
		return
	}
	fmt.Println(localIP)
}

func getLocalIP() (ip string, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}
	for _, addr := range addrs {
		ipAddr, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if ipAddr.IP.IsLoopback() {
			continue
		}
		if !ipAddr.IP.IsGlobalUnicast() {
			continue
		}
		return ipAddr.IP.String(), nil
	}
	return
}
