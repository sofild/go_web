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
    sort2(sms1_int_arr)
    t2 := time.Now().UnixNano()
	fmt.Println(sms1_int_arr)
    fmt.Println(t2-t1,"ns")
}

func sort(waitSort []int) {
	if len(waitSort) <= 1 {
		return
	}
	key := waitSort[0]
	i, j := 0, len(waitSort)-1
	for i < j {
		if waitSort[j] < key {
			waitSort[i], waitSort[j] = waitSort[j], waitSort[i]
			fmt.Println(i, j)
			for i < len(waitSort)-1 {
				if waitSort[i] > key {
					waitSort[i], waitSort[j] = waitSort[j], waitSort[i]
					fmt.Println(i, j)
					break
				} else {
					i++
				}
			}
		} else {
			j--
		}
	}
	fmt.Println(i, j, waitSort[:j+1], waitSort[j+1:])
	sort(waitSort[:i])
	sort(waitSort[j+1:])
}

func sort2(waitSort []int) {
	if len(waitSort) <= 1 {
		return
	}
	key := waitSort[0]
	i, j := 0, len(waitSort)-1
	for ; j > i; j-- {
		if key > waitSort[j] {
			waitSort[i], waitSort[j] = waitSort[j], waitSort[i]
			for ; i < len(waitSort)-1; i++ {
				if key < waitSort[i] {
					waitSort[i], waitSort[j] = waitSort[j], waitSort[i]
					break
				}
			}
		}
	}
	sort2(waitSort[:i])
	sort2(waitSort[j+1:])
}

func checkErr(err error, line ...int) {
	if err != nil {
		fmt.Println(line, err.Error())
	}
}

func in_array(arr []string, str string) bool {
	tag := false
	for _, key := range arr {
		if key == str {
			tag = true
		}
	}
	return tag
}

func contact(arr1 []int, arr2 []int) []int {
	arr := make([]int, len(arr1)+len(arr2))
	copy(arr, arr1)
	copy(arr[len(arr1):], arr2)
	return arr
}
