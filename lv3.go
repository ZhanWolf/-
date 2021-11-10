package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

var account =make(map[string]string)
var(
	s1 string
	s2 string
	s3 string
	c string
)


func main(){
	 b:=make([]byte,1024)
     f,err:=os.OpenFile("lv3-4-2021213619/users.data",os.O_RDWR|os.O_APPEND, 0777)
	 defer f.Close()
	 if err !=nil{
		 fmt.Println("err:",err)
	 }

	 _,err=f.Read(b)
	 index:=bytes.IndexByte(b,0)
	 if err !=nil{
		fmt.Println("err:",err)
	 }
     a:=b[0:index]
	 err=json.Unmarshal(a,&account)
	 if err !=nil{
		fmt.Println("err:",err)
	 }

	 fmt.Println("注册or登录")
	 fmt.Scanln(&c)
	switch c {
	case "注册":register()
		os.Truncate("lv3-4-2021213619/users.data",0)
		if err!=nil {
			fmt.Println(err)
		}
		k,_ :=json.Marshal(account)
		_,err =f.Write(k)
		if err !=nil{
			fmt.Println("err:",err)
		}
	case "登录":login()

	}
}

func register() {
	fmt.Println("请输入要注册的账号")
	fmt.Scanln(&s1)
	_,exist:=account[s1]
	if exist == true{
		fmt.Println("此账户已存在")
		goto end
	}
	for s2==s3 {
		fmt.Println("请输入密码")
		fmt.Scanln(&s2)
		fmt.Println("请重复输入密码")
		fmt.Scanln(&s3)
		if s2==s3{
			account[s1]=s2
			break
		}else {
			fmt.Println("密码不相同")
		}
	}
	end:
}


func login(){
	for {
		fmt.Println("请输入账号")
		fmt.Scanln(&s1)
		_, exist := account[s1]
		if exist == false {
			fmt.Println("此账号不存在")
			break
		}
		fmt.Println("请输入密码")
		fmt.Scanln(&s2)
		if account[s1] == s2 {
			fmt.Println("登录成功")
			break
		} else {
			fmt.Println("密码错误")
			break
		}
	}

}