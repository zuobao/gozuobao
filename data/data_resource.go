package data

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/coopernurse/gorp"
	"strings"
	"sync"
)
type DataSource struct {
	DbClusters    map[string]*DbCluster
	RedisClusters map[string]*RedisCluster

	DbMapInitCallback func(dbMap *gorp.DbMap, readonly bool, zone string) error
}

type DbSetting struct {
	DriverName string
	Server     string
	Username   string
	Password   string
	Database   string
	Properties map[string]string
	Engine     string
	Charset    string

	dbmap *gorp.DbMap
	db    *sql.DB
	m    sync.Mutex
}

func (me *DbSetting) GetDbConnectionMethods() (driverName, dataSourceName string) {

	props := []string{}
	if len(me.Properties) > 0 {
		for k, v := range me.Properties {
			props = append(props, k+"="+v)
		}
	}
	query_string := ""
	if len(props) > 0 {
		query_string = "?" + strings.Join(props, "&")
	}

	server := strings.TrimSpace(me.Server)
	database := strings.TrimSpace(me.Database)

	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s%s", me.Username, me.Password, server, database, query_string)

	return me.DriverName, connectionString
}

func (me *DbSetting) GetDb() (*sql.DB, error) {
	var err error
	if me.db == nil {
		me.db, err = sql.Open(me.GetDbConnectionMethods())
	}
	return me.db, err
}

func (me *DbSetting) GetDbMap(callback func(dbMap *gorp.DbMap, readonly bool) error, readonly bool) (*gorp.DbMap, error) {
	var db *sql.DB
	var err error

	if me.dbmap == nil {
		me.m.Lock()
		defer me.m.Unlock()
		if me.dbmap == nil {
			db, err = me.GetDb()
			if err == nil && db != nil {
				dbmap := &gorp.DbMap{
					Db:      db,
					Dialect: gorp.MySQLDialect{me.Engine, me.Charset},
				}

				if callback != nil {
					err = callback(dbmap, readonly)
					if err == nil {
						me.dbmap = dbmap
					}
				}
				return dbmap, err
			}
		}
	}

	return me.dbmap, err
}

type DbCluster struct {
	Master *DbSetting
	Slave  []*DbSetting
}

type RedisSetting struct {
	Network string
	Address string
	Index   int
}

type RedisCluster struct {
	Servers []*RedisSetting
}

func (me *DataSource) GetDbCluster(zone string) *DbCluster {
	if len(me.DbClusters) > 0 {
		return me.DbClusters[zone]
	}
	return nil
}

func (me *DataSource) GetDb(zone string, readonly bool) (*sql.DB, error) {
	cluster := me.GetDbCluster(zone)
	if cluster != nil {
		if !readonly || len(cluster.Slave) == 0 {
			return cluster.Master.GetDb()
		} else if len(cluster.Slave) > 0 {
			return cluster.Slave[0].GetDb()
		}
	}

	return nil, errors.New("无可用的DbCluster")
}

func (me *DataSource) GetDbMap(zone string, readonly bool) (*gorp.DbMap, error) {
	cluster := me.GetDbCluster(zone)
	var dbmap *gorp.DbMap
	var err error
	if cluster != nil {
		callback_internal := func (dbMap *gorp.DbMap, readonly bool) error {
			if me.DbMapInitCallback != nil {
				return  me.DbMapInitCallback(dbMap, readonly, zone)
			}
			return  nil
		}

		if !readonly || len(cluster.Slave) == 0 {
			dbmap, err = cluster.Master.GetDbMap(callback_internal, false)

		} else if len(cluster.Slave) > 0 {
			// 从库不执行初始化方法
			dbmap, err = cluster.Slave[0].GetDbMap(callback_internal, true)
		} else {
			err = errors.New("无可用的DbSetting")
		}
	} else {
		err = errors.New("无可用的DbCluster")
	}

	return dbmap, err
}
