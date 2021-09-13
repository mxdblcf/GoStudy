package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	httpque()
}

func Goto() {
	//
	roomId := ""
	//8
	start := "08"
	//12
	end := "12"

	//2021-09-05
	data := "2021-09-05"
	//座位号： 002
	seatNum := "002"
	//
	token1 := "b2e6c1ad07834aa88d44d0cb922130c4"

	//clint := &http.Client{}
	todyDate := time.Now()
	fmt.Println(todyDate)
	nextDate := time.Now().AddDate(0, 0, 1)
	fmt.Println("---1.")
	fmt.Println(nextDate)
	url1 := seturl(roomId, start, end, data, seatNum, token1)
	//开始循环
	fmt.Println(url1)
	//确定今天是哪一天，并再开启一个循环，不停的转

	//在今天的下午15：00开始抢座     ，抢的是明天的上午8：00-12：00的座位
	timeLayout := "2016-09-04 15:00:00"
	unix := time.Now().Unix()
	fmt.Println("--1")
	fmt.Println(unix)

	format := time.Unix(0, 0).Format(timeLayout)
	fmt.Println("--2")
	fmt.Println(format)
	//抢到之后，发送邮件给我  ；抢不到也发送一个邮件给我，继续优化，结束小循环
	//在外面的大循环再进行判断，如果今天拿到票，就休眠23小时

}
func seturl(roomId, start, end, data, seatNum, token1 string) string {
	s1 := "https://office.chaoxing.com/data/apps/seat/submit" +
		"?"
	s2 := "roomId=" + roomId + "&"

	s3 := "startTime=" + start + "%3A00&"

	s4 := "endTime=" + end + "%3A00&"

	s5 := "day=" + data + "&"
	//座位号
	s6 := "seatNum=" + seatNum + "&"

	s7 := "token=" + token1

	//s8:= " HTTP/1.1"

	return s1 + s2 + s3 + s4 + s5 + s6 + s7
}

func httpque() {
	client := &http.Client{}
	//生成要访问的url
	url := "http://office.chaoxing.com/data/apps/seat/submit?roomId=2582&startTim" +
		"e=08%3A00&endTime=12%3A00&day=2021-09-06&seatNum=002&token=fa5dab673ed44757be85b40d9dd23386"
	//提交请求
	request, err := http.NewRequest("GET", url, nil)
	cookie := "lv=4; fid=38691; _uid=83079908; UID=83079908;" +
		" vc=4E9DD59C6E2CADA40B52AC036C1151F3; vc2=0A06E52A765591A4475E83E55E3483DD;" +
		" xxtenc=c96178d0877c8eef1c37bf6c92747dea;" +
		" KI4SO_SERVER_EC=RERFSWdRQWdsckNwckxwM2pXeFFSYnlTYWJPaTNOTzRPYitoVDhVVG1" +
		"BU1dpd2RDWWYzaXdhNjE3%0Ack9rL2xJSkNPdGdxMXhwZmNPcApyTHAzald4UVJRSG9zQVQ4aDJj" +
		"Y08yaWxYeUxzSTd2clJ1a0ZW%0AMy9nZ01FRXdUUWt5Q3dmTGJSVzlhQUtjTmRVdmpjVGZ2STd3VEh" +
		"1Ck5EcWZVcTVkUUk1RzczdDFI%0AZkwwWG9jL1QrcUl5UTlkOW5PLytCa0FMNXdmTWpUU1lpMFhsb1g" +
		"rbG16YjVGL1VKYS9IZldLLzlv%0AMkIKeW1ZVkZ3WjJuQ2FhbjRxelFlUENzakVOdzlNZExVSDRFNzN0" +
		"RlpGRWxidWRBcXhLMHVkVkxF%0AUzNib2JiNDhLeU1RM0QweDFlYS9jMwpqN0hBeUpzQm9Kbis2Y3p2RV" +
		"lEaU5ndnM0aHJqd3JJeERj%0AUFRIWU52OEl4bS80Uk9vZzNxZVhKVEFUWkZpSlRSSEhucE5nSk9sdWlw" +
		"CkZLSHRJU2c5TEZsODkv%0AWkhmaXlkc3BmUmR0TDIzdWkrNGdwVFMwOEJrSDNyNmpTY3FTZEVFanZSTWdu" +
		"MDh2VEgzbHJuP2Fw%0AcElkPTEma2V5SWQ9MQ%3D%3D; _tid=59029490; sso_puid=83079908; _indu" +
		"stry=5; fidsCount=5; uf=94ffe74515793f365e4eb4d2f88e9162066dc8e3979931e2f905cbda4d4cb" +
		"b357be45317daa2847375690cf76fcddd5d913b662843f1f4ad6d92e371d7fdf64419c2e6aa68aac736fd" +
		"68be96b6183b1a0e69b5ac742f5f1d25f10702387a2da231dc80f7279e6ead; _d=1630823240796; vc3=" +
		"VHfyypmfHmHKWrjmgwI%2FPL2KTCJbIqqjXsdy%2FhZ9suzUPtcEqS72QD1%2BXg%2BcnwhmuTsRZbJm9n0gaIL" +
		"x3kLBneODm1wX9EiNEuvKcz5UJ1VtcRJm7QKXlbBHWrLf3VBPxKcESuoVguJyc4IqMrWkXBfjXmSEmP5ORMRhOY" +
		"49np8%3Db8d0bbcaa024bb4f65d94094956aef49; DSSTASH_LOG=C_38-UN_863-US_83079908-T_1630823" +
		"240798; oa_uid=83079908; oa_name=%E6%9E%97%E7%81%BF%E5%81%A5; JSESSIONID=9C0B77F53E52908" +
		"C98301424A05CFF41.reserve_web_125; oa_deptid=126519; oa_enc=beea12ae4626e23d843da8bc6569ad37; route=2107b3b209f27d6acae2395f23ff4722"
	referer := "http://office.chaoxing.com/front/apps/seat/select?id=2582&day=2021-09-06&backLevel=2&pageToken=d033ebeec0e547739e15ea6c7089ca69"
	//增加header选项
	request.Header.Add("Cookie", cookie)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 10; YAL-AL50 Build/HUAWEIYAL-AL50; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/83.0.4103.106 Mobile Safari/537.36 (device:YAL-AL50) Language/zh_CN_#Hans com.chaoxing.mobile/ChaoXingStudy_3_5.0.3_android_phone_609_67 (@Kalimdor)_b33b67d000704ce0abcac7065e9fad34")
	request.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7")
	request.Header.Add("Referer", referer)
	request.Header.Add("Accept-Encoding", "gzip, deflate")
	request.Header.Add("Host", "office.chaoxing.com")
	request.Header.Add("Connection", "keep-alive")
	request.Header.Add("X-Requested-With", "XMLHttpRequest")
	request.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")

	if err != nil {
		panic(err)
	}
	//处理返回结果
	response, _ := client.Do(request)
	fmt.Println("%v", response)
	defer response.Body.Close()
}
