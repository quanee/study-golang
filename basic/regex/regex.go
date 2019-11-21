package main

import (
	"fmt"
	"regexp"
)

const text = `
My email is ccmousse@gmail.com@abc
email1 is abc@def.org
email2 is kkk@qq.com
email3 is ddd@abc.com.cn
`
const test = `"basicInfo":["未婚","27岁","天秤座(09.23-10.22)","160cm","48kg","工作地:广安武胜","月收入:8千-1.2万","待业","大专"],"detailInfo":["汉族","籍贯:四川广安","体型:保密","不吸烟","不喝酒","和家人同住","未买车","没有小孩","是否想要孩子:想要孩子","何时结婚:时机成熟就结婚"],"educationString":"大专","emotionStatus":0,"gender":1,"genderString":"女士","hasIntroduce":true,"heightString":"160cm","hideVerifyModule":false,"introduceContent":"猫系女孩～～佛系青年～～\n性格懒散，万事不萦于心，不喜强求一切随缘。\n","introducePraiseCount":0,"isActive":false,"isFollowing":false,"isInBlackList":false,"isStar":false,"isZhenaiMail":true,"lastLoginTimeString":"2天前活跃","liveAudienceCount":0,"liveType":0,"marriageString":"未婚",`

func main() {
	re := regexp.MustCompile(`"basicInfo":\["([^"]+)","(\d+)岁","([^\(]+)\([^\)]+\)","(\d+)cm","(\d+)kg","工作地:[^"]+","月收入:([^"]+)","([^"]+)","([^"]+)"]`)
	//re := regexp.MustCompile(`"basicInfo":\["([^\)]+)","(\d+)岁","([^\(]+)\([^\)]+\)","(\d+)cm","(\d+)kg","工作地:[^"]+","月收入:([^"]+)","([^"]+)","([^"]+)"]`)
	//re := regexp.MustCompile(`([a-zA-Z0-9.]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9]+)`)
	match := re.FindSubmatch([]byte(test))
	for _, m := range match {
		fmt.Println(string(m))
	}
}
