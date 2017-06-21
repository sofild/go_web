package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	//tcpAddr := new(net.TCPAddr)
	//ip, _, _ := net.ParseCIDR("192.168.143.56")
	//tcpAddr.IP = ip
	//tcpAddr.Port = 6403

	//tcpListener, _ := net.ListenTCP("tcp", tcpAddr)
	//conn, _ := tcpListener.Accept()

	conn, err0 := net.Dial("tcp", "192.168.143.56:6403")
	if err0 != nil {
		fmt.Println("server connect failed.", err0.Error())
		return
	}
	fmt.Println("server has connected.")
	conn.SetReadDeadline(time.Now().Add(time.Second * 60))
	//fmt.Println(conn)
	/*
			sms := make([]byte, 128)
			fmt.Print("please input command:")
			_,err1 := fmt.Scan(&sms)
		    if err1!=nil{
		        fmt.Println("param id abnormal.", err1.Error())
		    }
			conn.Write(sms)

			buf := make([]byte, 128)
			c, err := conn.Read(buf)
			if err != nil {
				fmt.Println("err:", err.Error())
			}
			fmt.Println(string(buf[0:c]))
	*/
	defer conn.Close()
	exec(conn)
}

func exec(conn net.Conn) {
	//var sms string
	for {
		fmt.Print("please input command:")
		/*
			        _, err1 := fmt.Sscan(sms)
					if err1 != nil {
						fmt.Println("param id abnormal.", err1.Error())
					}
		*/
		reader := bufio.NewReader(os.Stdin)
		sms, _, _ := reader.ReadLine()
		//fmt.Println(sms)
		data := strings.Split(string(sms), " ")
		var datas []string
		dataLen := len(data)
		data1 := fmt.Sprintf("*%s", strconv.Itoa(dataLen))
		datas = append(datas, data1)
		var data2 string
		for _, dat := range data {
			len2 := len(dat)
			data2 = fmt.Sprintf("$%s", strconv.Itoa(len2))
			datas = append(datas, data2)
			datas = append(datas, dat)
		}
		//fmt.Println(dataLen, datas)
		nowsms := strings.Join(datas, "\r\n")
		nowsms = fmt.Sprintf("%s\r\n", nowsms)
		//fmt.Printf("now sms:%s\n", nowsms)
		writer := bufio.NewWriter(conn)
		smsByte := []byte(nowsms)
		//fmt.Println(smsByte)
		_, err2 := writer.Write(smsByte)
		//n, err2 := conn.Write(smsByte)
		if err2 != nil {
			fmt.Println("write failed.", err2.Error())
		}
		writer.Flush()
		//fmt.Printf("%s\n", nowsms)
		//fmt.Printf("write %d\n", n)

		buf := make([]byte, 1024*1000*1000)
		//fmt.Println("Ready to read...\n")

		//c, err := conn.Read(buf)
		c, err := bufio.NewReader(conn).Read(buf)

		//fmt.Println("Start to read...\n")
		if err != nil {
			fmt.Println("err:", err.Error())
		}
		//fmt.Printf("read %d", c)
		//fmt.Println(string(buf[0:c]))
        result(string(buf[0:c]))
	}
}

func result(value string){
    tag := value[0:1]
    switch tag {
        case "+":
            fmt.Printf("%s", value[1:])
        case "-":
            fmt.Printf("%s", value[1:])
        case ":":
            fmt.Printf("%s", value[1:])
        case "$":
            content := value[1:]
            index := strings.Index(content, "\n")
            index += 1
            fmt.Printf("%s", content[index:])
        case "*":
            mainCnt := value[1:]
            index := strings.Index(mainCnt, "\n")
            index += 1
            content := mainCnt[index:] 
            contArr := strings.Split(content, "\r\n")
            var data []string
            for _, v := range contArr {
                if !strings.Contains(v, "$") {
                    data = append(data, v)
                }
            }
            prt := strings.Join(data, "\n")
            fmt.Printf("%s", prt)
        default:
            fmt.Printf("%s", value[1:])
    }
}
