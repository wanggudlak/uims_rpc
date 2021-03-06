package gredis

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	log "github.com/sirupsen/logrus"
	"uims/conf"
)

var connections = map[string]*redis.Client{}

func InitDef() (*redis.Client, error) {
	return initConn(conf.Database.Redis["default"])
}

func initConn(c conf.RedisConf) (*redis.Client, error) {
	var client *redis.Client
	client = newConn(fmt.Sprintf("%s:%d", c.Host, c.Port), c.Password, c.Database)
	_, err := client.Ping().Result()
	if err != nil {
		return client, err
	}
	return client, nil
}

func newConn(addr string, password string, db int) *redis.Client {
	var options = &redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	}
	return redis.NewClient(options)
}

// Get connection
func Conn(name string) *redis.Client {
	if client, ok := connections[name]; ok {
		return client
	}
	if c, ok := conf.Database.Redis[name]; ok {
		connections[name], _ = initConn(c)
		return connections[name]
	}
	log.Panicf("Not config redis connection: %s", name)
	return Def()
}

// Get default connection
func Def() *redis.Client {
	return Conn("default")
}

func Close() {
	for k, conn := range connections {
		if err := conn.Close(); err != nil {
			log.Printf("Close mysql conn %s err: %+v", k, err)
		}
	}
}
