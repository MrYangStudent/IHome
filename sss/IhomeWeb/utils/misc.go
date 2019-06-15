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



// 后续还会添加