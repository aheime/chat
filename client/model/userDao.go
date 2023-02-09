package model

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"gotest/chat/common"
	"strconv"
)

var (
	MyUserDao *UserDao
)

type UserDao struct {
	pool *redis.Client
}

func NewUserDao(pool *redis.Client) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}

	return
}

func (t *UserDao) GetUserById(userName string) (user *common.User, err error) {
	id, err := t.pool.HGet("user", userName).Result()
	if err != nil {
		if err == redis.Nil {
			err = ERROR_USER_NOT_EXITSTS
		}
		fmt.Println("redis err=", err)
		return
	}

	str, err := t.pool.Get("userid:" + id).Result()
	if err != nil {
		if err == redis.Nil {
			err = ERROR_USER_NOT_EXITSTS
		}
		fmt.Println("redis err=", err)
		return
	}
	idInt, err := strconv.Atoi(id)
	user = &common.User{
		UserId: idInt,
	}
	err = json.Unmarshal([]byte(str), user)
	if err != nil {
		fmt.Println("json err=", err)
		return
	}
	return
}

func (t *UserDao) Login(userName string, pwd string) (user *common.User, err error) {
	user, err = t.GetUserById(userName)
	if err != nil {
		return
	}

	if user.Password != pwd {
		err = ERROR_USER_PWD
		return
	}
	return
}

func (t *UserDao) Register(user *common.User) (err error) {

	_, err = t.GetUserById(user.UserName)
	if err == nil {
		err = ERROR_USER_EXITSTS
		return
	}

	user.UserId = common.Inc.Id()
	fmt.Println("user=", user)
	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("json err=", err)
		return
	}
	_, err = t.pool.Set("userid:"+strconv.Itoa(user.UserId), string(data), 0).Result()
	if err != nil {
		fmt.Println("保存注册数据err=", err)
		return
	}
	_, err = t.pool.HSet("user", user.UserName, user.UserId).Result()
	if err != nil {
		fmt.Println("保存注册数据err=", err)
		return
	}
	return
}
