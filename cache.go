package bookeeping

import (
    "github.com/bradfitz/gomemcache/memcache"
    "encoding/gob"
    "bytes"
    "log"
)
type Cache struct {
    client *memcache.Client

}

func newCache() (*Cache) {
    c := Cache{memcache.New("localhost:11211")}

    return &c
}

func (c *Cache) Get(key string,st interface{}) (error) {
    res,err:= c.client.Get(key)

    if err != nil {

        return err
    }

    r := bytes.NewReader(res.Value)
    de := gob.NewDecoder(r)

    return de.Decode(st)
}

func (c *Cache)Set(key string,value interface{},ttl int32) (bool) {
    var b bytes.Buffer
    en := gob.NewEncoder(&b)
    if err := en.Encode(value); err != nil {
        //log.SetFlags(log.LstdFlags | log.Lshortfile)
        log.Fatal(err)

        return false
    }
    c.client.Set(&memcache.Item{Key:key,Value:b.Bytes(),Expiration:ttl})

    return true
}
