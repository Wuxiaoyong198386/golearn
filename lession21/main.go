/**
 * 版本控制器
 *
 * @desc 健身馆家
 * ---------------------------------------------------------------------
 * @versionor		wu <vip120@126.com>
 * @date		2018-08-27
 * @copyright	ppfs-api-gym
 * ---------------------------------------------------------------------
 */

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/henrylee2cn/ini"
	"time"
)

type RedisPool struct {
	pool *redis.Pool
}

//连接redis函数
func newPool(addr string, pwd string) (*RedisPool, error) {
	pool := &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", addr)
			if err != nil {
				return nil, err
			}
			if pwd != "" {
				if _, err := conn.Do("AUTH", ""); err != nil {
					return nil, err
				}
			}
			return conn, nil
		},
	}

	return &RedisPool{pool}, nil
}

var (
	config_file   = "/home/wwwroot/golearn/lession21/config.ini"
	redisServer   string
	redisPassword string
	cfg           *ini.File
)

func main() {

	//读取ini文件,得到配置配置项
	cfg, err := ini.Load(config_file)
	if err != nil {
		fmt.Println("Error:", err.Error())
	}
	if Section, err := cfg.GetSection("redisPool"); err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		redisServer = Section.Key("redisServer").String()
		redisPassword = Section.Key("redisPassword").String()
	}

	//连接redis
	//var pool RedisPool
	pool, err := newPool(redisServer, redisPassword)
	//生成队列
	pool.Enqueue()
	//消费队列
	pool.GetJob()
	//取得队列中的长度
	pool.GetJoblen()

	//conn := pool
	//defer conn.Close()
	//myname := "wuxiaoyong"
	////写数据
	//if w, e := conn.Do("SET", "myname", myname); e != nil {
	//	fmt.Println("Error:", e.Error())
	//} else {
	//	fmt.Println(w)
	//}
	////取数据
	//if r, e := redis.String(conn.Do("GET", "myname")); e != nil {
	//	fmt.Println("Error:", e.Error())
	//} else {
	//	fmt.Println(r)  //写入成功，返回OK
	//}
	////删除数据
	//if d,e:=conn.Do("DEL","myname");e!=nil{
	//	fmt.Println(e.Error())
	//}else{
	//	fmt.Println(d)  //如果删除返回，返回1
	//}

}

//结构体中的key首字母一定要大写
type students struct {
	Name string
	Age  int
}

//写入redis队列
func (r *RedisPool) Enqueue() {
	conn := r.pool.Get()
	defer conn.Close()
	//var tmp_data []*students
	s1 := students{Name: "wuxiaoyong2", Age: 30}
	//tmp_data=append(tmp_data,&s1)
	//fmt.Println(tmp_data)
	//s2 := &students{"小华", 40}
	s1json, err := json.Marshal(s1)
	if err != nil {
		fmt.Println("json marshal error:", err.Error())
	}
	//s2json,err:=json.Marshal(s2)
	//if err!=nil{
	//	fmt.Println("json marshal error:",err.Error())
	//}
	//fmt.Println(s1json)
	conn.Do("RPUSH", "students", s1json)
	//conn.Do("rpush", "students", s2json)
}

//读取redis队列
func (r *RedisPool) GetJob() {
	conn := r.pool.Get()
	defer conn.Close()
	count, err := r.GetJoblen()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(count)
	for i := 0; i < count; i++ {
		s, err := redis.String(conn.Do("LPOP", "students"))
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(s)
		}
	}
}

//求队列长度
func (r *RedisPool) GetJoblen() (int, error) {
	conn := r.pool.Get()
	defer conn.Close()
	joblen, err := conn.Do("llen", "students")
	if err != nil {
		return 0, err
	}
	count, ok := joblen.(int64)
	if !ok {
		return 0, errors.New("类型转换失败")
	}
	return int(count), nil
}
