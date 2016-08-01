package data

import (
	"errors"
	"github.com/garyburd/redigo/redis"
	"time"
)

func (this *RedisSetting) GetRedis() (redis.Conn, error) {
	c , err := redis.DialTimeout(this.Network,
		this.Address,
		time.Second*15,
		time.Second*15,
		time.Second*15)
	if err == nil && c != nil {
		if this.Index > 0 {
			//选择上指定索引的数据库
			c.Do("select", this.Index)
		}
	}
	return c, err
}

func (me *DataSource) GetRedisCluster(zone string) *RedisCluster {
	if len(me.RedisClusters) > 0 {
		return me.RedisClusters[zone]
	}
	return nil
}

func (me *DataSource) GetDefaultRedis() (redis.Conn, error) {
	cluster := me.GetRedisCluster("default")
	if cluster != nil && len(cluster.Servers) > 0 {
		return cluster.Servers[0].GetRedis()
	} else {
		return nil, errors.New("redis未找到")
	}
}
