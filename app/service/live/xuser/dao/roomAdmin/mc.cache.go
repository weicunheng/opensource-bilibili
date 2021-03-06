// Code generated by $GOPATH/src/go-common/app/tool/cache/mc. DO NOT EDIT.

/*
  Package roomAdmin is a generated mc cache package.
  It is generated from:
  type _mc interface {
		// 获取主播的房管列表
		// mc: -key=KeyRoom
		CacheRoomAdminRoom(c context.Context, anchor int64) ([]*model.RoomAdmin, error)
		// 获取用户的房管列表
		// mc: -key=KeyUser
		CacheRoomAdminUser(c context.Context, user int64) ([]*model.RoomAdmin, error)

		// mc: -key=KeyRoom -expire=d.RoomAdminExpire -encode=json|gzip
		AddCacheKeyAnchorRoom(c context.Context, anchor int64, value []*model.RoomAdmin) error
		// mc: -key=KeyUser -expire=d.RoomAdminExpire -encode=gob
		AddCacheRoomAdminUser(c context.Context, user int64, value []*model.RoomAdmin) error

		// mc: -key=KeyRoom
		DelCacheKeyAnchorRoom(c context.Context, anchor int64) error
		// mc: -key=KeyUser
		DelCacheRoomAdminUser(c context.Context, user int64) error
	}
*/

package roomAdmin

import (
	"context"
	"fmt"

	"go-common/app/service/live/xuser/model"
	"go-common/library/cache/memcache"
	"go-common/library/log"
	"go-common/library/stat/prom"
)

var _ _mc

// CacheRoomAdminRoom 获取主播的房管列表
func (d *Dao) CacheRoomAdminRoom(c context.Context, id int64) (res []*model.RoomAdmin, err error) {
	conn := d.mc.Get(c)
	defer conn.Close()
	key := KeyRoom(id)
	reply, err := conn.Get(key)
	if err != nil {
		if err == memcache.ErrNotFound {
			err = nil
			return
		}
		prom.BusinessErrCount.Incr("mc:CacheRoomAdminRoom")
		log.Errorv(c, log.KV("CacheRoomAdminRoom", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	res = []*model.RoomAdmin{}
	err = conn.Scan(reply, &res)
	if err != nil {
		prom.BusinessErrCount.Incr("mc:CacheRoomAdminRoom")
		log.Errorv(c, log.KV("CacheRoomAdminRoom", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	return
}

// CacheRoomAdminUser 获取用户的房管列表
func (d *Dao) CacheRoomAdminUser(c context.Context, id int64) (res []*model.RoomAdmin, err error) {
	conn := d.mc.Get(c)
	defer conn.Close()
	key := KeyUser(id)
	reply, err := conn.Get(key)
	if err != nil {
		if err == memcache.ErrNotFound {
			err = nil
			return
		}
		prom.BusinessErrCount.Incr("mc:CacheRoomAdminUser")
		log.Errorv(c, log.KV("CacheRoomAdminUser", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	res = []*model.RoomAdmin{}
	err = conn.Scan(reply, &res)
	if err != nil {
		prom.BusinessErrCount.Incr("mc:CacheRoomAdminUser")
		log.Errorv(c, log.KV("CacheRoomAdminUser", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	return
}

// AddCacheKeyAnchorRoom Set data to mc
func (d *Dao) AddCacheKeyAnchorRoom(c context.Context, id int64, val []*model.RoomAdmin) (err error) {
	if len(val) == 0 {
		return
	}
	conn := d.mc.Get(c)
	defer conn.Close()
	key := KeyRoom(id)
	item := &memcache.Item{Key: key, Object: val, Expiration: d.RoomAdminExpire, Flags: memcache.FlagJSON | memcache.FlagGzip}
	if err = conn.Set(item); err != nil {
		prom.BusinessErrCount.Incr("mc:AddCacheKeyAnchorRoom")
		log.Errorv(c, log.KV("AddCacheKeyAnchorRoom", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	return
}

// AddCacheRoomAdminUser Set data to mc
func (d *Dao) AddCacheRoomAdminUser(c context.Context, id int64, val []*model.RoomAdmin) (err error) {
	if len(val) == 0 {
		return
	}
	conn := d.mc.Get(c)
	defer conn.Close()
	key := KeyUser(id)
	item := &memcache.Item{Key: key, Object: val, Expiration: d.RoomAdminExpire, Flags: memcache.FlagGOB}
	if err = conn.Set(item); err != nil {
		prom.BusinessErrCount.Incr("mc:AddCacheRoomAdminUser")
		log.Errorv(c, log.KV("AddCacheRoomAdminUser", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	return
}

// DelCacheKeyAnchorRoom delete data from mc
func (d *Dao) DelCacheKeyAnchorRoom(c context.Context, id int64) (err error) {
	conn := d.mc.Get(c)
	defer conn.Close()
	key := KeyRoom(id)
	if err = conn.Delete(key); err != nil {
		if err == memcache.ErrNotFound {
			err = nil
			return
		}
		prom.BusinessErrCount.Incr("mc:DelCacheKeyAnchorRoom")
		log.Errorv(c, log.KV("DelCacheKeyAnchorRoom", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	return
}

// DelCacheRoomAdminUser delete data from mc
func (d *Dao) DelCacheRoomAdminUser(c context.Context, id int64) (err error) {
	conn := d.mc.Get(c)
	defer conn.Close()
	key := KeyUser(id)
	if err = conn.Delete(key); err != nil {
		if err == memcache.ErrNotFound {
			err = nil
			return
		}
		prom.BusinessErrCount.Incr("mc:DelCacheRoomAdminUser")
		log.Errorv(c, log.KV("DelCacheRoomAdminUser", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	return
}
