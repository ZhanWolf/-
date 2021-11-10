package main

import "fmt"

func main() {
	fmt.Println("选择模板")
    var c int
    fmt.Scanln(&c)
	switch c {
	case 1:ReleaseSkill("龙卷风摧毁停车场",func(skillName string){
		fmt.Println("尝尝我的厉害吧",skillName)
	    })
	case 2:
		ReleaseSkill("龙卷风摧毁停车场",func(skillName string){
			fmt.Println("你已经是一具尸体了",skillName)
		})
	}

}


func ReleaseSkill(skillNames string, releaseSkillFunc1 func(string)) {
	fmt.Println("输入技能名")
	var s string
	fmt.Scanln(&s)
	if s =="龙卷风摧毁停车场"{
	releaseSkillFunc1(skillNames)}

}



