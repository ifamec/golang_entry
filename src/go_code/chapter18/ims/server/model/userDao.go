package model

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

// init a userDao when server start
var ImsUserDao *UserDao


// user action

type UserDao struct {
	pool *redis.Pool
}

// factory to return a userDao object
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

// id -> user, err
func (ud *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {
	// search in redis
	rst, err := redis.String(conn.Do("HGET", "users", id))
	if err != nil{
		if err == redis.ErrNil { // not found
			err = ERROR_USER_NOTEXISTS
		}
		return
	}

	// deserialize rst
	user = &User{}
	err = json.Unmarshal([]byte(rst), user)
	if err != nil{
		fmt.Println("Server.UserDao : Unmarshall Error -", err)
	}

	return
}

// login validation
// 1. validate User with id & pwd return user object
// 2. if invalid return err msg

func (ud *UserDao) Login(userId int, userPwd string) (user *User, err error) {

	// get conn from pool
	conn := (*ud).pool.Get()
	defer conn.Close()

	user, err = (*ud).getUserById(conn, userId)
	if err != nil {
		return
	}

	// validate user
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD_INVALID
		return
	}

	return
}


// signup validation

func (ud *UserDao) Signup(user *User) (err error) {
	// get conn from pool
	conn := (*ud).pool.Get()
	defer conn.Close()

	// check if user exist
	user, err = (*ud).getUserById(conn, user.UserId)
	if err == nil {
		err = ERROR_USER_EXISTS
		return
	}

	// sign up
	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Server.UserDao : Marshall Error -", err)
		return
	}

	_, err = conn.Do("HSET", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("Server.UserDao : Signup Error -", err)
		return
	}


	return
}