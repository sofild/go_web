package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
    "time"
)

func main() {
	fmt.Print("请输入要排序的数字，请用空格分隔：")
	reader1 := bufio.NewReader(os.Stdin)
	sms1, _, err1 := reader1.ReadLine()
	checkErr(err1)
	sms1_str := string(sms1)
	sms1_arr := strings.Split(sms1_str, " ")
	if len(sms1_arr) == 0 {
		fmt.Print("没有要排序的值")
		return
	}
	sms1_int_arr := make([]int, len(sms1_arr))
	for i, key := range sms1_arr {
		if key != " " {
			sms1_int_arr[i], _ = strconv.Atoi(key)
		}
	}
	//fmt.Println(sms1_int_arr)
	t1 := time.Now().UnixNano()
    sort(sms1_int_arr)
    t2 := time.Now().UnixNano()
	fmt.Println(sms1_int_arr)
    fmt.Println(t2-t1, "ns")
}

func sort(waitSort []int) {
	if len(waitSort) <= 1 {
		return
	}
	for i := len(waitSort) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if waitSort[j] > waitSort[j+1] {
				waitSort[j], waitSort[j+1] = waitSort[j+1], waitSort[j]
			    //fmt.Println(waitSort)
            }
		}
       // fmt.Println(i,".")
	}
}

func checkErr(err error, line ...int) {
	if err != nil {
		fmt.Println(line, err.Error())
	}
}
