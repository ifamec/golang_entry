package processes

import "fmt"

// only have one, declare as global

var userMgr *UserMgr

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

func init()  {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

func (um *UserMgr) AddOnlineUser(user *UserProcess)  {
	(*um).onlineUsers[user.UserId] = user
}
func (um *UserMgr) RemoveOnlineUser(userId int)  {
	delete((*um).onlineUsers, userId)
}
func (um *UserMgr) GetAllOnlineUsers() map[int]*UserProcess {
	return (*um).onlineUsers
}
func (um *UserMgr) GetOnlineUserById(userId int) (user *UserProcess, err error) {
	user, isOnline := (*um).onlineUsers[userId]
	if !isOnline { // offline
		err = fmt.Errorf("User %d is offline.\n", userId)
	}
	return
}
