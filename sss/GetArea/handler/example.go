package handler

import (
	"context"
	"fmt"
	example "sss/GetArea/proto/example"
	"sss/IhomeWeb/utils"
	"github.com/astaxie/beego/orm"
	"sss/IhomeWeb/models"
	"encoding/json"
	"time"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) GetArea(ctx context.Context, req *example.Request, rsp *example.Response) error {
	fmt.Println("获取地域信息服务   GetArea  /api/v1.0/areas")
	//1 初始化返回值
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)
	// 1.5 连接reids   获取数据

	//bm, err := cache.NewCache("memory", `{"interval":60}`)
	//redis配置信息
	/*redis_config_map := map[string]string{
		"key":utils.G_server_name,
		"conn":utils.G_redis_addr+":"+utils.G_redis_port,
		"dbNum":utils.G_redis_dbnum,
	}
	//将 配置信息的map 转化为json
	redis_config ,_:=json.Marshal(redis_config_map)

	//连接redis
	bm ,err :=cache.NewCache("redis",string(redis_config))
	if err!=nil{
		fmt.Println("连接 redis错误",err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}*/

	bm ,err :=utils.RedisOpen(utils.G_server_name,utils.G_redis_addr,utils.G_redis_port,utils.G_redis_dbnum)
	if err!=nil{
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	//redis key
	key := "area_info"

	//获取数据

	area_info_value := bm.Get(key)
	//接收数据
	var areas []models.Area
	//判断数据是否为空
	if area_info_value !=nil{
		fmt.Println("获取到数据准备发送给web")

		//解码
		err =json.Unmarshal(area_info_value.([]byte) ,&areas)


	//	循环转换数据发送给web

		for key , value := range areas {
			fmt.Println(key , value)

			//	结构体 ---》proto
			area :=example.ResponseAddress{Aid: int32(value.Id),Aname:string(value.Name)}
			//4 数据返回
			rsp.Data =append(rsp.Data ,&area)
		}

		return nil
	}



	//2 查询数据库
	o :=orm.NewOrm()


	//设查询条件
	qs :=o.QueryTable("Area")

	//查询全部
	num ,err := qs.All(&areas)
	if err!=nil{
		fmt.Println("查询数据库错误",err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	if num == 0{
		fmt.Println("无数数据",err)
		rsp.Errno = utils.RECODE_NODATA
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}


	//2.5  存入数据
	//json编码
	area_info_json,_ :=json.Marshal(areas)

	//存入数据
	err =bm.Put(key,area_info_json,time.Second*7200)
	if err!=nil{
		fmt.Println("redis存入数据失败",err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	//3  将查询到的数据转化类型
	for key , value := range areas {
		fmt.Println(key , value)

	//	结构体 ---》proto
		area :=example.ResponseAddress{Aid: int32(value.Id),Aname:string(value.Name)}
		//4 数据返回
		rsp.Data =append(rsp.Data ,&area)
	}


	return nil
}

