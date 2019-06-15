package handler

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"github.com/micro/go-grpc"
	"fmt"
	"context"
	GETAREA 		"sss/GetArea/proto/example"
	GETIMAGECD		"sss/GetImageCd/proto/example"
	GETSMSCD		"sss/GetSmscd/proto/example"
	POSTRET			"sss/PostRet/proto/example"
	"sss/IhomeWeb/models"
	"image"
	"github.com/afocus/captcha"
	"image/png"
	"regexp"
	"sss/IhomeWeb/utils"
)
//模板
//路由解析函数  （ 传出给前端  ，  传入 前端发送过来的  ， 解析路由请求参数  ）
func ExampleCall(w http.ResponseWriter, r *http.Request,ps httprouter.Params) {
	//// decode the incoming request as json
	////接受 前端发送过来数据的
	//var request map[string]interface{}
	//// 将前端 json 数据解析到 map当中
	//if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
	//	http.Error(w, err.Error(), 500)
	//	return
	//}
	//
	////创建 grpc 客户端
	//cli :=grpc.NewService()
	////客户端初始化
	//cli.Init()
	//
	//// call the backend service
	////通过protobuf 生成文件 创建 连接服务端 的客户端句柄
	//exampleClient := example.NewExampleService("go.micro.srv.template", cli.Client())
	////通过句柄调用服务端函数
	//rsp, err := exampleClient.Call(context.TODO(), &example.Request{
	//	Name: request["name"].(string),
	//})
	////判断是否成功
	//if err != nil {
	//	http.Error(w, err.Error(), 500)
	//	return
	//}
	//
	//// we want to augment the response
	////准备返回给前端的map
	//response := map[string]interface{}{
	//	"errno": utils.RECODE_MOBILEERR,
	//	"errmsg": utils.RecodeText(utils.RECODE_MOBILEERR),
	//}
	//// encode and write the response as json
	//	//设置返回数据的格式
	//	w.Header().Set("Content-Type","application/json")
	////将map转化为json 返回给前端
	//if err := json.NewEncoder(w).Encode(response); err != nil {
	//	http.Error(w, err.Error(), 500)
	//	return
	//}
}

func GetArea(w http.ResponseWriter, r *http.Request,ps httprouter.Params) {
	fmt.Println("获取地域信息服务   GetArea  /api/v1.0/areas")



	//1创建 grpc 客户端
	cli :=grpc.NewService()
	//2 客户端初始化
	cli.Init()

	// call the backend service
	//3通过protobuf 生成文件 创建 连接服务端 的客户端句柄
	exampleClient := GETAREA.NewExampleService("go.micro.srv.GetArea", cli.Client())
	//通过句柄调用服务端函数
	rsp, err := exampleClient.GetArea(context.TODO(), &GETAREA.Request{})
	//判断是否成功
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	//接受数据
	area_list := []models.Area{}

	for _, value := range rsp.Data {

		tmp:= models.Area{Id:int(value.Aid),Name:value.Aname}
		area_list=append(area_list ,tmp)

	}



	// we want to augment the response
	//准备返回给前端的map
	response := map[string]interface{}{
		"errno": rsp.Errno,
		"errmsg": rsp.Errmsg,
		"data":area_list,
	}

	//设置返回数据的格式
	w.Header().Set("Content-Type","application/json")

	// encode and write the response as json

	//将map转化为json 返回给前端
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}




func GetImageCd(w http.ResponseWriter, r *http.Request,ps httprouter.Params) {
	fmt.Println("获取验证码图片   GetImageCd  /api/v1.0/imagecode/:uuid")

	//获取url代参 的uuid
	uuid :=ps.ByName("uuid")

	//创建 grpc 客户端
	cli :=grpc.NewService()
	//客户端初始化
	cli.Init()

	// call the backend service
	//通过protobuf 生成文件 创建 连接服务端 的客户端句柄
	exampleClient := GETIMAGECD.NewExampleService("go.micro.srv.GetImageCd", cli.Client())
	//通过句柄调用服务端函数
	rsp, err := exampleClient.GetImageCd(context.TODO(), &GETIMAGECD.Request{
		Uuid:uuid,
	})
	//判断是否成功
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//判断返回值 如果不等于 0直接返回错误
	if rsp.Errno != "0"{
		// we want to augment the response
		//准备返回给前端的map
		response := map[string]interface{}{
			"errno": rsp.Errno,
			"errmsg": rsp.Errmsg,
		}
		// encode and write the response as json
		//设置返回数据的格式
		w.Header().Set("Content-Type","application/json")
		//将map转化为json 返回给前端
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}

	//	接受图片数据 拼接图片结构体 将图片发送给前端

//	接受图片变量
	var img  image.RGBA
	//	pix
	for _, value := range rsp.Pix {

		img.Pix =append(img.Pix , uint8(value))
	}
	//	Stride
	img.Stride = int(rsp.Stride)
	//  point
	img.Rect.Min.X = int(rsp.Min.X)
	img.Rect.Min.Y = int(rsp.Min.Y)
	img.Rect.Max.X = int(rsp.Max.X)
	img.Rect.Max.Y = int(rsp.Max.Y)


	var image captcha.Image

	image.RGBA  = &img
	//将图片发送给前端

	png.Encode(w ,image)

}


func GetSmscd(w http.ResponseWriter, r *http.Request,ps httprouter.Params) {
	fmt.Println(" 获取短信验证码 GetSmscd  /api/v1.0/smscode/:mobile")

	//获取手机号 byname
	mobile := ps.ByName("mobile")
	//验证手机号
	myreg := regexp.MustCompile(`0?(13|14|15|17|18|19)[0-9]{9}`)

	bo :=myreg.MatchString(mobile)
	if bo ==false{
		//准备返回给前端的map
		response := map[string]interface{}{
			"errno": utils.RECODE_MOBILEERR,
			"errmsg": utils.RecodeText(utils.RECODE_MOBILEERR),
		}
		// encode and write the response as json
		//设置返回数据的格式
		w.Header().Set("Content-Type","application/json")
		//将map转化为json 返回给前端
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}



	//获取 图片验证码 与 uuid
	text :=r.URL.Query()["text"][0]
	uuid :=r.URL.Query()["id"][0]



	//创建 grpc 客户端
	cli :=grpc.NewService()
	//客户端初始化
	cli.Init()

	// call the backend service
	//通过protobuf 生成文件 创建 连接服务端 的客户端句柄
	exampleClient := GETSMSCD.NewExampleService("go.micro.srv.GetSmscd", cli.Client())
	//通过句柄调用服务端函数
	rsp, err := exampleClient.GetSmscd(context.TODO(), &GETSMSCD.Request{
		Mobile:mobile,
		Uuid:uuid,
		Text:text,
	})
	//判断是否成功
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	//准备返回给前端的map
	response := map[string]interface{}{
		"errno": rsp.Errno,
		"errmsg": rsp.Errmsg,
	}
	// encode and write the response as json
		//设置返回数据的格式
	w.Header().Set("Content-Type","application/json")
	//将map转化为json 返回给前端
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}


func PostRet(w http.ResponseWriter, r *http.Request,ps httprouter.Params) {

	fmt.Println(" 注册服务  PostRet  /api/v1.0/users")

	// decode the incoming request as json
	//接受 前端发送过来数据的
	var request map[string]interface{}
	// 将前端 json 数据解析到 map当中
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//判断数据
	/*    mobile: "123",
    password: "123",
    sms_code: "123"
*/
	if request["mobile"].(string) =="" ||request["password"].(string) =="" || request["sms_code"].(string) ==""{
		response := map[string]interface{}{
			"errno": utils.RECODE_DATAERR,
			"errmsg": utils.RecodeText(utils.RECODE_DATAERR),
		}
		// encode and write the response as json
		//设置返回数据的格式
		w.Header().Set("Content-Type","application/json")
		//将map转化为json 返回给前端
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}

	//创建 grpc 客户端
	cli :=grpc.NewService()
	//客户端初始化
	cli.Init()

	// call the backend service
	//通过protobuf 生成文件 创建 连接服务端 的客户端句柄
	exampleClient := POSTRET.NewExampleService("go.micro.srv.PostRet", cli.Client())
	//通过句柄调用服务端函数
	rsp, err := exampleClient.PostRet(context.TODO(), &POSTRET.Request{
		Mobile:request["mobile"].(string),
		Password:request["password"].(string),
		SmsCode:request["sms_code"].(string),
	})
	//判断是否成功
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	//准备返回给前端的map
	response := map[string]interface{}{
		"errno": rsp.Errno,
		"errmsg": rsp.Errmsg,
	}


	//设置cookie
	cookie ,err := r.Cookie("ihomelogin")

	if err!=nil || cookie.Value ==""{
		cookie := http.Cookie{Name:"ihomelogin",Value:rsp.Sessionid, MaxAge:600,Path:"/"}
		http.SetCookie(w,&cookie)
	}

	// encode and write the response as json
		//设置返回数据的格式
		w.Header().Set("Content-Type","application/json")
	//将map转化为json 返回给前端
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}




func GetSession(w http.ResponseWriter, r *http.Request,ps httprouter.Params) {

	// we want to augment the response
	//准备返回给前端的map
	response := map[string]interface{}{
		"errno": "4101",
		"errmsg": "用户未登录",
	}
	// encode and write the response as json
		//设置返回数据的格式
		w.Header().Set("Content-Type","application/json")
	//将map转化为json 返回给前端
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}


func GetIndex(w http.ResponseWriter, r *http.Request,ps httprouter.Params) {


	// we want to augment the response
	//准备返回给前端的map
	response := map[string]interface{}{
		"errno": "0",
		"errmsg": "ok",
	}
	// encode and write the response as json
		//设置返回数据的格式
		w.Header().Set("Content-Type","application/json")
	//将map转化为json 返回给前端
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}