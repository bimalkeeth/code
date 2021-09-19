package proxy

import (
	"math/rand"
	"testing"
)

func Test_UserListProxy(t *testing.T){
	mockDatabase :=UserList{}
	rand.Seed(2342342)

	for i:=0;i<1000000;i++ {
		n:=rand.Int31()
		mockDatabase=append(mockDatabase,User{ID:n})
	}

	proxy:=UserListProxy{
		MockedDatabase: &mockDatabase,
		StackCache: UserList{},
		StackSize: 2,
	}

	knowsIDs:=[3]int32{mockDatabase[3].ID,mockDatabase[4].ID,mockDatabase[5].ID}

	t.Run("FindUser - Empty Cache",func(t *testing.T){
		user,err:=proxy.FindUser(knowsIDs[0])
		if err !=nil{
			t.Fatal(err)
		}
		if user.ID !=knowsIDs[0]{
			t.Error("returned user name doesn't match with expected")
		}
		if len(proxy.StackCache)!=1{
			t.Error("after one successful search in an empty cache, the size of it must be pne")
		}
		if proxy.LastSearchUserCache==true{
			t.Error("no user can be returned from empty cache")
		}

	})

	t.Run("FindUser - overflowing the stack",func(t *testing.T){

		user1,err:=proxy.FindUser(knowsIDs[0])
		if err !=nil{
			t.Fatal(err)
		}

		user2,_:=proxy.FindUser(knowsIDs[1])
		if proxy.LastSearchUserCache{
			t.Error("the user wasn't stored on the proxy cache yet")
		}

		user3,_:=proxy.FindUser(knowsIDs[2])
		if proxy.LastSearchUserCache{
			t.Error("the user wasn't stored on the proxy cache yet")
		}

		for i:=0;i<len(proxy.StackCache);i++{
			if  proxy.StackCache[1].ID ==user1.ID{
				t.Error("user that should be gone was found")
			}
		}
		if len(proxy.StackCache)!=2{
			t.Error("after inserting 3 users the cache should not grow more than to two")
		}
		for _,v:=range proxy.StackCache{
			if v!=user2 && v!=user3{
				t.Error("a none expected user was found on the cache")
			}
		}

	})




}