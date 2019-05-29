package main

import (
	"math/rand"
	"fmt"
	"time"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func contain(a []int,i int) bool {
	for _,v := range a {
		if v == i {
			return true
		}
	}
	return false
}

func getNumbers(a []int,b[]int)  {
	rand.Seed(time.Now().Unix())

	//随机替换3个位置的数
	for changed := 0;changed<3;{
		for i,_ := range a {
			if changed == 3 {
				break
			}
			p := rand.Float64()
			if p<=0.5 {
				rn := rand.Intn(34)+1
				if !contain(a,rn){
					a[i] = rn
					changed += 1
				}
			}
		}
	}

	//随机替换1个位置的数
	for changed := 0;changed<1;{
		for i,_ := range b {
			if changed == 1 {
				break
			}
			p := rand.Float64()
			if p<=0.5 {
				rn := rand.Intn(11)+1
				if !contain(b,rn){
					b[i] = rn
					changed += 1
				}
			}
		}
	}

	fmt.Println(a,b)
	time.Sleep(1*time.Second)
}

func numbers(s string) []int {
	var n []int
	for _, f := range strings.Fields(s) {
		i, err := strconv.Atoi(f)
		if err == nil {
			n = append(n, i)
		}
	}
	return n
}


func main() {
	//

	cancel := make(chan int,1)

	go func() {
		for {
			var a  = make([]int,0,5)
			var b  = make([]int,0,5)

			scanner := bufio.NewScanner(os.Stdin)
			fmt.Println("请输入前5个数字")
			if scanner.Scan(){
				a = numbers(scanner.Text())
			}
			fmt.Println("请输入后两个数字")
			if scanner.Scan(){
				b = numbers(scanner.Text())
			}

			fmt.Println("请输入要生成的号码个数：")
			var n int
			_,err := fmt.Scanf("%d",&n)
			if err != nil {
				fmt.Println("输入有误，退出")
				cancel <- 0
			}

			for j := 0;j<n;j++{
				var ac  = make([]int,5,5)
				var bc  = make([]int,2,2)
				copy(ac,a)
				copy(bc,b)
				getNumbers(ac,bc)
			}
			os.Exit(0)

		}
	}()

	<-cancel

}