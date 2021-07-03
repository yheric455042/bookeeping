package account

import (
  //"log"
  "fmt"
  "bookeeping/firebase"
)
const PATH string = "bookeeping/account/"

type Account struct{
    db *firebase.Firebase
}

// result need to be this data struct
type AuthData struct {
   TelegramId string `json:"telegram_id,omitempty"`
   Email string `json:"email,omitempty"`
   Token string `json:"token"`
}

func Init(db *firebase.Firebase) (*Account) {
    construct := Account{db}

    return &construct
}

func (c *Account) GetAll() (interface{}) {
    x := make(map[string]AuthData)
    res,_ := c.db.Get(PATH,x)

    return res
}

func (c *Account) GetByName(name string) (interface{}) {
    var a AuthData
    res,_ := c.db.Get(PATH + name,a)

    return res
}


func (c *Account) Login(user string, password string) (bool) {
    var a AuthData
    res,_ := c.db.Get(PATH + user,a)
    fmt.Println(a)
    fmt.Println(res)

    return true
}
