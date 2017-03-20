package main

import (
	"fmt"
	"runtime"
	"sync"
	"wiki.ruokuai.com/ApiDemo_Go.ashx/rkdama"
)

var wg = sync.WaitGroup{}

func CreateMultiThread(username string, password string, typeid string, timeout string, softid string, softkey string, filename string) error {
	creatresult, err := rkdama.RKCreate(username, password, typeid, timeout, softid, softkey, filename)
	if err == nil {
		fmt.Println("答题结果:" + creatresult.Result)
		fmt.Println("结果ID:" + creatresult.Id)
	} else {
		fmt.Println(err)
	}
	defer wg.Done()
	return nil
}

func CreateUrlMultiThread(username string, password string, typeid string, timeout string, softid string, softkey string, url string) error {
	creatreurlsult, err := rkdama.RKCreateUrl(username, password, typeid, timeout, softid, softkey, url)
	if err == nil {
		fmt.Println("答题结果:" + creatreurlsult.Result)
		fmt.Println("结果ID:" + creatreurlsult.Id)
	} else {
		fmt.Println(err)
	}
	defer wg.Done()
	return nil
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	wg.Add(200)
	//查询接口调用
	inforesult, err := rkdama.RKQueryInfo("用户名", "密码")
	if err == nil {
		fmt.Println("查询信息")
		fmt.Println("剩余快豆:" + inforesult.Score)
		fmt.Println("历史快豆:" + inforesult.HistoryScore)
		fmt.Println("使用快豆:" + inforesult.TotalScore)
		fmt.Println("答题总数:" + inforesult.TotalScore)
	} else {
		fmt.Println(err)
	}
	//上传本地验证码答题接口调用
	filename := "./1.png"
	for i := 0; i < 1; i++ {
		go CreateMultiThread("用户名", "密码", "2040", "60", "1", "b40ffbee5c1cf4e38028c197eb2fc751", filename)
	}
	wg.Wait()
	//上传URL验证码答题接口调用
	//	fileurl := "http://captcha.qq.com/getimage"
	//	for i := 0; i < 2; i++ {
	//		go CreateUrlMultiThread("用户名", "密码", "2040", "60", "1", "b40ffbee5c1cf4e38028c197eb2fc751", fileurl)
	//	}
	//	wg.Wait()
	//报错接口调用
	//	reportresult, err := rkdama.RKReportError("用户名", "密码", "1", "b40ffbee5c1cf4e38028c197eb2fc751", "验证码结果ID")
	//	if err == nil {
	//		fmt.Println("报错结果:" + reportresult.Result)
	//	} else {
	//		fmt.Println(err)
	//	}
	//注册接口调用
	//	registerresult, err := rkdama.RKRegister("用户名", "密码", "邮箱")
	//	if err == nil {
	//		fmt.Println("注册结果:" + registerresult.Result)
	//	} else {
	//		fmt.Println(err)
	//	}
	//充值接口调用
	//	rechargeresult, err := rkdama.RKRecharge("用户名", "卡号", "卡密")
	//	if err == nil {
	//		fmt.Println("注册结果:" + rechargeresult.Result)
	//	} else {
	//		fmt.Println(err)
	//	}
}
