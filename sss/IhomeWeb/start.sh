# 启动 redis  服务
redis-server  ./conf/redis.conf


# 启动fastdfs
fdfs_trackerd /home/itcast/go/src/sss/IhomeWeb/conf/tracker.conf    restart


fdfs_storaged /home/itcast/go/src/sss/IhomeWeb/conf/storage.conf    restart


# 启动 nginx

sudo nginx