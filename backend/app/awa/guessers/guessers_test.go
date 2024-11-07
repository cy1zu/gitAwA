package guessers

import (
	"fmt"
	"testing"
)

func TestGuessNation(t *testing.T) {
	head, err := Init()
	if err != nil {
		t.Fatal(err)
	}
	query := map[string]interface{}{
		"url":      "https://github.com/tiger1103",
		"name":     "",
		"company":  "",
		"blog":     "",
		"location": "",
		"email":    "",
		"comments": []string{
			"新版本验证码多了个返回参数，已修复",
			"link就是数据库配置，驱动:用户名:密码@tcp(ip:端\ufffd\ufffd",
			"支持，用户信息是存的redis，存的是jwt信息，conte",
			"https://goframe.org/pages/viewpage.action?pageId=57183742\r\nhttp:",
			"参考文档：https://goframe.org/pages/viewpage.action?pageId=",
			"右上角点那个垃圾桶图标清下缓存，然后f12清下l",
			"项目数据库文件 resource/data/gfast-v32.sql 创建数据\ufffd\ufffd",
			"文档已上线，请查看readme中文档地址",
		},
	}
	res, err := GuessNation(head, query)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", res)
}

/*
{
	"name": "David J Wu",
	"company": "",
	"blog": "",
	"location": "",
	"email": "",
	"comments": [
		"Do you know anyone who would be interested to implement and main",
		"Thanks for the discussion. I'll investigate a little and conside",
		"I reviewed probably around 50 different 19x19 training games and",
		"@kaorahi - Very cool. Is this phenomenon consistent across multi",
		"Thanks for posting. Would you be able to include an SGF file?",
		"Thanks!",
		"Thanks for reporting! This will probably also join into a pile o",
		"Thanks for posting this. I will add more variations of this kind"
	]
}

*/
