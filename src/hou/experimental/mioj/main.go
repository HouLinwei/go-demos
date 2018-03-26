package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
)

func solution(line string) string {
	return line
}

func solution4(line string) string {
	raw := strings.Split(line, ",")
	if len(raw) == 0{
		return line
	}
	if len(raw) == 1{
		return  raw[0]
	}

	var i int
	var rot bool
	for i = 1; i< len(raw); i++{
		current, _ := strconv.Atoi(raw[i])
		pre, _ := strconv.Atoi(raw[i-1])
		if current == pre +1{
			continue
		}else{
			rot = true
			break
		}
	}
	if rot{
		//fmt.Println(i)
		idx := i-1 + len(raw)/2+1
		//fmt.Println(idx)
		if idx >= len(raw){
			return raw[len(raw)-idx]
		}else{
			return raw[idx]
		}
	}else{
		return raw[len(raw)/2]
	}

}

func QS(list []int, start, end int){
	if start < end{
		left := start
		right := end
		key :=  list[left]
		for left < right {
			for left < right && list[right] > key{
				right -= 1
			}
			list[left] = list[right]
			for left < right && list[left] <= key{
				left += 1
			}
			list[right] = list[left]
		}
		list[left] = key
		QS(list, start, left-1)
		QS(list, left+1, end)
	}
}

func solution3(line string) string {
	raw := strings.Split(line, ",")
	if len(raw) == 0{
		return "0"
	}
	if len(raw) == 1{
		return "1"
	}
	var list []int
	for _, v := range raw{
		i, _ := strconv.Atoi(v)
		list = append(list, i)
	}
	QS(list, 0, len(list)-1)

	//fmt.Println(list)

	var max int
	max = 1
	var t int
	t = 1

	for i := 1; i< len(list); i++{
		//fmt.Println(max)
		if list[i] == list[i-1] +1{
			t += 1
		}else{
			t = 1
		}
		if max < t {
			max = t
		}
	}

	return strconv.Itoa(max)
}

func solution2(line string) string {
	// 在此处理单行数据
	// 返回处理后的结果
	res := strings.Split(line, " ")
	if len(res)!= 3{
		return line
	}
	if res[1] != "-"{
		return line
	}
	if res[2] == "0"{
		return res[1]
	}

	a := res[0]
	b := res[2]
	la := len(a)
	lb := len(b)
	var ret []uint8
	var i int
	var pre uint8
	for i = 0; i< la && i< lb; i++{
		ta := a[la-1-i]
		tb := b[lb-1-i]
		ta = ta - pre
		if ta >= tb {
			ret = append([]uint8{ta -tb + 48},ret...)
			pre = 0
		}else{
			ret = append([]uint8{ta+10 -tb +48},ret...)
			pre = 1
		}
	}
	for j:= i; j< la; j++{
		if pre > 0{
			ret = append([]uint8{a[la-1-j]-1}, ret...)
			pre = 0
		}else{
			ret = append([]uint8{a[la-1-j]}, ret...)
		}

	}
	return string(ret)
}

func main() {
	r := bufio.NewReader(os.Stdin)
	for line, _, err := r.ReadLine(); err == nil; line, _, err = r.ReadLine() {
		fmt.Println(solution(string(line)))
	}
}
