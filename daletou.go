package main

import (
	"math/rand"
	"fmt"
	"time"
	"os"
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

func main() {
	//

	cancel := make(chan int,1)

	go func() {
		for {
			var a  = make([]int,0,5)
			var b  = make([]int,0,5)
			goOn := true
			for i := 0;i<5;i++ {
				if i == 0 {
					fmt.Println("请输入前五位数字（以数字结束，超过5个数取前五个数）：")
				}
				var t int

				_, err := fmt.Scanf("%d", &t)
				if err != nil {
					fmt.Println(err)
					goOn = false
					break
				}
				if t <= 0 || t > 35 {
					fmt.Println("超过1-35范围，数字输入错误，请重新输入")
					goOn = false
					break
				}
				if !contain(a, t) {
					a = append(a, t)
				} else {
					goOn = false
					fmt.Println("数字重复，重新输入")
					break
				}
			}
			if len(a) > 5 {
				fmt.Println("超过5位长度")
				break
			}
			if !goOn {
				break
			}
			time.Sleep(50*time.Millisecond)

			for i:=0; i<2;i++ {
				if i ==0{
					fmt.Println("请输入后两位数字（以数字结束，超过2个数取前2个数）：")
				}
				var t int
				_,err := fmt.Scanf("%d",&t)
				if err != nil {
					fmt.Println(err)
					goOn = false
					break
				}
				if t<=0 || t>12 {
					fmt.Println("超过1-12范围，请重新输入")
					goOn = false
					break
				}
				if !contain(b,t){
					b = append(b,t)
				} else {
					goOn = false
					fmt.Println("数字重复，重新输入")
					break
				}
			}
			if goOn {
				fmt.Println("请输入要生成的号码个数：")
				var n int
				_,err := fmt.Scanf("%d",&n)
				if err != nil {
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

		}
	}()

	<-cancel

}