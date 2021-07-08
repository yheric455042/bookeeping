package bookeeping

import (
    "github.com/allegro/bigcache"
    "encoding/gob"
    "time"
    "bytes"
)

func GetInCache(key string,st interface{}) (bool) {
    cache,_:=bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Second))
    res,_:=cache.Get(key)
    r := bytes.NewReader(res)
    de := gob.NewDecoder(r)
    if err := de.Decode(&st); err != nil {
        return false
    }

    return true
}

func SetInCache(key string,value interface{}) (bool) {
    var b bytes.Buffer
    en := gob.NewEncoder(&b)
    cache,_:=bigcache.NewBigCache(bigcache.DefaultConfig(500 * time.Second))
    if err := en.Encode(&value); err != nil {
        return false
    }

    cache.Set(key,[]byte(b.String()))

    return true
}
