package handler

import (
	"context"

	"fmt"
	example "sss/DeleteSession/proto/example"
	"sss/IhomeWeb/utils"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) DeleteSession(ctx context.Context, req *example.Request, rsp *example.Response) error {
	fmt.Println("退出登陆  DeleteSession   /api/v1.0/session")
	/*1 初始化返回值*/
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg =utils.RecodeText(rsp.Errno)

	/*2 获取sessionid*/
	sessionid :=req.Sessionid

	/*3 链接redis*/

	bm ,err :=utils.RedisOpen(utils.G_server_name ,utils.G_redis_addr,utils.G_redis_port,utils.G_redis_dbnum)
	if err !=nil{

		fmt.Println("redis连接失败",err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg =utils.RecodeText(rsp.Errno)

		return nil
	}

	/*4 拼接key删除 session登陆信息*/

	bm.Delete(sessionid +"name")
	bm.Delete(sessionid +"user_id")
	bm.Delete(sessionid +"mobile")


	return nil
}

