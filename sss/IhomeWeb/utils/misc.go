package utils

import (
	"encoding/json"
	"github.com/astaxie/beego/cache"
	"fmt"
	_ "github.com/astaxie/beego/cache/redis"
	_ "github.com/garyburd/redigo/redis"
	_ "github.com/gomodule/redigo/redis"
	"crypto/md5"
	"encoding/hex"
	"github.com/weilaihui/fdfs_client"
)

/* 将url加上 http://IP:PROT/  前缀 */
//http:// + 127.0.0.1 + ：+ 8080 + 请求
//https://img.alicdn.com/tps/i4/TB1L7lExXzqK1RjSZFoSuvfcXXa.jpg_q90_.webp
func AddDomain2Url(url string) (domain_url string) {
	domain_url = "http://" + G_fastdfs_addr + ":" + G_fastdfs_port + "/" + url

	return domain_url
}


//redis 连接函数

func RedisOpen( server_name,redis_addr,redis_port,redis_dbnum string)(bm  cache.Cache ,err error){
	redis_config_map := map[string]string{
		"key":server_name,
		"conn":redis_addr+":"+redis_port,
		"dbNum":redis_dbnum,
	}
	//将 配置信息的map 转化为json
	redis_config ,_:=json.Marshal(redis_config_map)

	//连接redis
	bm ,err =cache.NewCache("redis",string(redis_config))
	if err!=nil{
		fmt.Println("连接 redis错误",err)
		return nil,err
	}
	return  bm ,nil
}



//加密函数

func Getmd5string(s string)string{
	m :=md5.New()
	return  hex.EncodeToString(m.Sum([]byte(s)))
}


//fastdfs 上传 操作 (二进制 文件 ，后缀名)（fileid ，err）

func Uploadbybuf(file []byte ,Extname string)(fileid string ,err error) {
//	读取配置文件 创建fdfs句柄
	fdfsclient ,err :=fdfs_client.NewFdfsClient("/home/itcast/go/src/sss/IhomeWeb/conf/client.conf")
	if err!=nil{
		fmt.Println("fdfs_client.NewFdfsClient 创建失败",err)
		return  "",err
	}

// 	上传文件
	rsp  ,err :=fdfsclient.UploadByBuffer(file  ,Extname)
	if err!=nil{
		fmt.Println("fdfsclient.UploadByBuffer 上传 失败",err)

		return  "",err
	}

	fmt.Println("GroupName:",rsp.GroupName , "Fileid:",rsp.RemoteFileId)

//	返回fileid
	return  rsp.RemoteFileId,nil
}

// 后续还会添加