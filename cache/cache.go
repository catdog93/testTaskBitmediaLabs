package cache

import (
	"fmt"
	"sync"
	"testTaskBitmediaLabs/entity"
)

type CacheUserBodyService interface {
	Add(docs ...entity.UserBody) (err error)
	FindId(id uint64) (UserBody entity.UserBody, isCreated bool, err error)
	ReplaceId(UserBody *entity.UserBody) (err error)
	//UpsertId(UserBody *entity.UserBody) (err error)
	DeleteId(id uint64) (err error)
}

type UserCache struct { // UserBody.ID equals key of map's element
	usersCache map[string]entity.User
	sync.RWMutex
}

var Cache = UserCache{
	usersCache: make(map[string]entity.User),
}

func (c *UserCache) Add(docs ...entity.User) (err error) {
	errString := ""
	if docs != nil {
		c.Lock()
		switch {
		case c == nil:
			errString += fmt.Sprintf("cache Add error: UserCache can't be nil\n")
		case c.usersCache == nil:
			errString += fmt.Sprintf("cache Add error: cache can't be nil\n")
		default:
			for _, v := range docs {
				c.usersCache[v.ID.String()] = v
			}
			c.Unlock()
		}
	} else {
		errString += fmt.Sprintf("cache Add error: impossible to add nil value\n")
	}
	if errString != "" {
		return fmt.Errorf(errString)
	}
	return
}

//func (c *UserCache) FindId(id uint64) (UserBody entity.UserBody, isCreated bool, err error) {
//	errString := ""
//	if id != 0 {
//		c.RLock()
//		switch {
//		case c == nil:
//			errString += fmt.Sprintf("cache FindId error: UserCache can't be nil\n")
//		case c.usersCache == nil:
//			errString += fmt.Sprintf("cache FindId error: cache can't be nil\n")
//		default:
//			if _, ok := c.usersCache[id]; ok {
//				UserBody = c.usersCache[id]
//				isCreated = true
//			}
//		}
//		c.RUnlock()
//	} else {
//		errString += fmt.Sprintf("cache FindId error: id can't be == 0\n")
//	}
//	if errString != "" {
//		err = fmt.Errorf(errString)
//	}
//	return
//}
//
//func (c *UserCache) ReplaceId(id uint64, UserBody *entity.UserBody) (err error) {
//	errString := ""
//	isCreated := false
//	if UserBody != nil {
//		c.RLock()
//		switch {
//		case c == nil:
//			c.RUnlock()
//			errString += fmt.Sprintf("cache ReplaceId error: UserCache can't be nil\n")
//		case c.usersCache == nil:
//			c.RUnlock()
//			errString += fmt.Sprintf("cache ReplaceId error: cache can't be nil\n")
//		default:
//			c.RUnlock()
//			if _, isCreated, err = c.FindId(id); err != nil {
//				return
//			} else {
//				if isCreated {
//					c.Lock()
//					c.usersCache[id] = *UserBody
//					c.Unlock()
//				} else {
//					errString += fmt.Sprintf("cache ReplaceId error: element doesn't exist with ID = %d\n", id)
//				}
//			}
//		}
//	} else {
//		errString += fmt.Sprintf("cache ReplaceId error: impossible to replace UserBody with nil value\n")
//	}
//	if errString != "" {
//		err = fmt.Errorf(errString)
//	}
//	return
//}

//
//func (c *UserCache) UpsertId(UserBody *entity.UserBody) (err error) {
//	isCreated := false
//	if UserBody != nil {
//		if _, isCreated, err = c.FindId(UserBody.ID); err != nil {
//			return
//		} else {
//			if isCreated {
//				err = c.ReplaceId(UserBody)
//			} else {
//				err = c.Add(*UserBody)
//			}
//		}
//	} else {
//		fmt.Errorf("cache UpsertId error: impossible to upsert UserBody with nil value\n")
//	}
//	return
//}

//func (c *UserCache) DeleteId(id uint64) (err error) {
//	isCreated := false
//	if _, isCreated, err = c.FindId(id); err != nil {
//		return
//	} else {
//		if isCreated {
//			c.Lock()
//			delete(c.usersCache, id)
//			c.Unlock()
//		} else {
//			err = fmt.Errorf("cache DeleteId error: impossible to delete element with id = %v because it doesn't exist\n", id)
//		}
//	}
//	return
//}
