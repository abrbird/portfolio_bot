package cache

import (
	"fmt"
	"github.com/abrbird/portfolio_bot/internal/domain"
	"github.com/abrbird/portfolio_bot/pkg/api"
	"strings"
	"sync"
	"time"
)

type IO func() error

type Cache struct {
	seconds int64
	data    map[domain.UserId]*UserCache
	m       *sync.RWMutex
}

type UserCache struct {
	User    *api.User
	Ts      time.Time
	Command string
}

type UserCacheRetrieve struct {
	Data  UserCache
	Error error
}

func New(seconds int64) *Cache {
	return &Cache{
		seconds: seconds,
		data:    make(map[domain.UserId]*UserCache),
		m:       new(sync.RWMutex),
	}
}

func (c *Cache) Get(userId domain.UserId) <-chan UserCacheRetrieve {
	channel := make(chan UserCacheRetrieve, 1)

	go func() {
		c.m.RLock()
		defer c.m.RUnlock()

		if data, ok := c.data[userId]; ok {
			differenceInMinutes := int64(time.Now().UTC().Sub(data.Ts.UTC()).Seconds())
			if differenceInMinutes < c.seconds {
				channel <- UserCacheRetrieve{Data: *data, Error: nil}
			}
		}
		channel <- UserCacheRetrieve{Error: fmt.Errorf("does not exist id=%d", userId)}
		close(channel)
	}()
	return channel
}

func (c *Cache) Set(userCache *UserCache) <-chan UserCacheRetrieve {
	channel := make(chan UserCacheRetrieve, 1)

	go func() {
		c.m.Lock()
		defer c.m.Unlock()

		userId := domain.UserId(userCache.User.Id)
		c.data[userId] = userCache
		channel <- UserCacheRetrieve{Data: *c.data[userId], Error: nil}

		close(channel)
	}()
	return channel
}

func (ud *UserCache) GetName() string {
	userName := "anonymous"
	if len(ud.User.FirstName) > 0 || len(ud.User.LastName) > 0 {
		userName = fmt.Sprintf("%s %s", ud.User.FirstName, ud.User.LastName)
	} else if len(ud.User.UserName) > 0 {
		userName = ud.User.UserName
	}
	userName = strings.TrimSpace(userName)

	return userName
}
