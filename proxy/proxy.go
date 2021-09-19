package proxy

import (
	"fmt"
)

type User struct {
	ID int32
}

type UserFinder interface {
	FindUser(id int32) (User, error)
}

type UserList []User

func (t *UserList)FindUser(id int32)(User,error){
	for i:=0;i<len(*t);i++{
		if (*t)[i].ID==id{
			return (*t)[i],nil
		}
	}
	return User{},fmt.Errorf("user %v could not be found",id)
}

func (t *UserList) addUser(user User) {
	*t=append(*t,user)
}


type UserListProxy struct {
	MockedDatabase      *UserList
	StackCache          UserList
	StackSize           int
	LastSearchUserCache bool
}

func(u *UserListProxy)FindUser(id int32)(User,error){

	user,err:=u.StackCache.FindUser(id)
	if err==nil{
		fmt.Println("returning user from cache")
		u.LastSearchUserCache=true
		return user,nil
	}
	user,err=u.MockedDatabase.FindUser(id)
	if err !=nil{
		return User{},err
	}
	u.addUserToStack(user)
	return user,nil
}


func(u *UserListProxy)addUserToStack(user User){
	if len(u.StackCache)>=u.StackSize{
		u.StackCache=append(u.StackCache[1:],user)
	}else{
		u.StackCache.addUser(user)
	}
}