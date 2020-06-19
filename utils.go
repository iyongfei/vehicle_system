package golinux

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
	"vsys_itai/common"
	"vsys_itai/gowrt"
	"vsys_itai/log"
	"vsys_itai/utils"
)

//old mac: 1122334455AA

const (
	Path_Cconfig = "/vehicle/cconfig.txt"
)

var (
	IpConnectChan chan string
	IpAddr        string
)

func init() {
	IpConnectChan = make(chan string)

	//IpAddr = GetIPAddr()

	//go listenIPAddr()
	go getNewIpAddr()
}

func getNewIpAddr() {
	for {
		time.Sleep(time.Second * 20)
		a := GetIPAddr()

		if len(a) > 0 {
			IpConnectChan <- a
		}

	}
}

func GetAllRxTxInfo() (name string, rx string, tx string, errRT error) {
	dir, _ := os.Getwd()

	pathSli := strings.Split(dir, "/")
	pa := ""
	for k, v := range pathSli {
		if k >= len(pathSli)-1 {
			break
		}
		if len(v) <= 0 {
			continue
		}
		pa = pa + "/"
		pa = pa + v
	}
	pa = pa + Path_Cconfig
	command := "sed -n 2p " + pa
	fmt.Println("path=", pa)
	bytes, k, err := gowrt.ExecCommand(command)
	if err != nil {
		fmt.Println(err)
		fmt.Println(k)
		return "", "1", "", err
	}
	out := string(bytes)
	fmt.Println("out=", out)
	out = strings.ReplaceAll(out, "\n", "")
	if len(out) > 0 {

	}

	commandR := "ifconfig " + out + "|grep \"(\""
	fmt.Println("cmd=", commandR)
	//ifconfig enp12s0|grep byte|awk '{print $2}'
	bytesR, kR, errR := gowrt.ExecCommand(commandR)

	if errR != nil {
		fmt.Println(errR, kR)
		errRT = err
		return "", "", "2", err
	}
	outR := string(bytesR)
	outT := ""
	fmt.Println("outr", outR)
	if strings.Contains(outR, ":") {
		sliR := strings.Split(outR, ":")
		if len(sliR) > 2 {
			outR = sliR[1]
			outT = sliR[2]

			rsli := strings.Split(outR, "(")
			tsli := strings.Split(outT, "(")
			r := strings.ReplaceAll(rsli[0], " ", "")
			t := strings.ReplaceAll(tsli[0], " ", "")
			outR = strings.ReplaceAll(r, "\n", "")
			outT = strings.ReplaceAll(t, "\n", "")
		}
	}

	return out, outR, outT, errRT
}

// GetBoardName as name
func GetBoardName() (string, error) {
	return "YinHe-QiLin", nil
}

func GetGUID() string {
	fmt.Println("guid===", GetLinuxUniqueGUID())
	return GetLinuxUniqueGUID()
}

func GetUpRouterIp() string {
	command := "ip route show|grep default"
	bytes, _, err := gowrt.ExecCommand(command)
	if err != nil {
		fmt.Println(err)
	}
	out := string(bytes)
	if len(out) <= 0 {
		return ""
	}
	outsli := strings.Split(out, " ")
	for _, v := range outsli {
		if strings.Contains(v, ".") {
			return v
		}
	}
	return ""
}

func GetLinuxUniqueGUID() string {
	ipmac, _ := GetLocalMac()
	cpuinfo, _ := GetCPUInfo()
	diskinfo, _ := GetDiskMac()
	mixd := ipmac + "," + cpuinfo + "," + diskinfo
	fmt.Println("Vgw utils define=", utils.Base64Encode(mixd))
	md5str := utils.Md5V(mixd)
	return md5str
	//command := "ls -l /dev/disk/by-uuid|grep /sda1|awk '{print $9}'"
	//bytes,_,err := gowrt.ExecCommand(command)
	//if err != nil{
	//	fmt.Println(err)
	//}
	//out := string(bytes)

	guid := ""
	if len(guid) <= 0 {
		fileName := "guid.txt"

		pwd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}

		var exist = true
		if _, err := os.Stat(pwd + "/" + fileName); os.IsNotExist(err) {
			exist = false
		}
		if exist {
			f, err := os.Open(pwd + "/" + fileName)
			fmt.Println(f.Name())
			if err != nil {

			}
			bytes, _ := ioutil.ReadAll(f)
			guid = strings.ReplaceAll(string(bytes), "\n", "")
			fmt.Println("read local guid", guid)
			return guid
		}

		command := "touch " + fileName
		_, _, err = gowrt.ExecCommand(command)
		if err != nil {
			fmt.Println(err)
		}
		if err != nil {
			fmt.Println(err)
		}

		guid = RandomString(32)

		logFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			log.Error("write2File open guid err:%s", err)
		}
		defer logFile.Close()
		logWriter := bufio.NewWriter(logFile)

		if _, err = logWriter.Write([]byte(guid)); err != nil {
			log.Error("guidWriter Write err:%s", err)
		}
		defer logWriter.Flush()
		guid = strings.ReplaceAll(guid, "\n", "")
		fmt.Println("generate guid", guid)
		return guid
	}
	guid = strings.ReplaceAll(guid, "\n", "")
	return guid
}

func RandomString(ln int) string {
	letters := []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, ln)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}

	return string(b)
}

// GetFirminfo get firminfo by uci
func GetFirminfo() (*common.Firminfo, error) {
	fi := &common.Firminfo{Version: "1.0", GitMD5: "1258456s", SupplyID: "6.0"}
	return fi, nil
}

func NetWorkStatus() bool {
	cmd := exec.Command("ping", "baidu.com", "-c", "1", "-W", "5")
	err := cmd.Run()
	if err != nil {
		log.Error("no network")
		return false
	} else {
		log.Info("network is well")
	}
	return true
}

func GetAllIpAddrs() (ips []string) {
	interfaceAddr, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Printf("fail to get net interface addrs: %v", err)
		return ips
	}

	for _, address := range interfaceAddr {
		ipNet, isValidIPNet := address.(*net.IPNet)
		if isValidIPNet && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ips = append(ips, ipNet.IP.String())
			}
		}
	}
	return ips
}

//获取本机ip，外部可以访问的IP
func GetIPAddr() string {
	MasterIP := "127.0.0.1"
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {

		//get ether name && ether ip
		interfaces, _ := net.Interfaces()
		//fmt.Println(interfaces)
		for _, inter := range interfaces {
			if addrs, err := inter.Addrs(); err == nil {
				for _, addr := range addrs {
					if addr.(*net.IPNet).IP.To4() != nil &&
						addr.(*net.IPNet).IP.String() != "127.0.0.1" {
						//fmt.Println(inter.Name, "->", addr.(*net.IPNet).IP.String())
						if len(inter.Name) >= 1 && string(inter.Name[0]) == "e" {
							MasterIP = addr.(*net.IPNet).IP.String()
							//fmt.Println(inter.Name, "->", addr.(*net.IPNet).IP.String())
							break
						}

					}

				}
			}

		}
		fmt.Println(MasterIP)
		//os.Exit(0)
		return ""
	}
	defer conn.Close()
	MasterIP = strings.Split(conn.LocalAddr().String(), ":")[0]
	fmt.Println(MasterIP)
	return MasterIP
}

func GetDiskMac() (string, error) {
	command := "ls -l /dev/disk/by-uuid|grep /sda1|awk '{print $9}'"
	bytes, _, err := gowrt.ExecCommand(command)
	if err != nil {
		fmt.Println(err)
	}
	out := string(bytes)
	return out, nil
}

func GetCPUInfo() (string, error) {
	command := "dmidecode -t 4|grep ID"
	bytes, _, err := gowrt.ExecCommand(command)
	if err != nil {
		fmt.Println(err)
	}
	out := string(bytes)
	//ID: D4 06 00 00 FF FB 8B 0F
	//ID: D4 06 00 00 FF FB 8B 0F
	//ID: D4 06 00 00 FF FB 8B 0F
	//ID: D4 06 00 00 FF FB 8B 0F
	mac := strings.Split(out, "ID:")
	if len(mac) > 0 && strings.Contains(out, "ID:") {
		id1 := mac[1]
		id1 = strings.ReplaceAll(id1, "\n", "")
		id1 = strings.ReplaceAll(id1, " ", "")
		return id1, nil
	}

	return out, nil
}

//获取本机Mac
func GetLocalMac() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Get loacl Mac failed")
		return "", nil
	}
	combin := []string{}
	for _, inter := range interfaces {
		mac := inter.HardwareAddr
		if mac.String() != "" {
			mac := strings.ReplaceAll(mac.String(), ":", "")
			combin = append(combin, mac)
		}
	}
	if len(combin) > 1 {
		return combin[0] + combin[1], nil
	} else if len(combin) == 1 {
		return combin[0], nil
	}
	return "", nil
}

func StringIpToInt(ipstring string) int64 {
	ipSegs := strings.Split(ipstring, ".")
	var ipInt int64 = 0
	var pos uint = 24
	for _, ipSeg := range ipSegs {
		//tempInt, _ := strconv.Atoi(ipSeg)
		tempInt, _ := strconv.ParseInt(ipSeg, 10, 64)
		tempInt = tempInt << pos
		ipInt = ipInt | tempInt
		pos -= 8
	}
	return ipInt
}
