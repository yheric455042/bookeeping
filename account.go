package bookeeping

import (
  "github.com/golang-jwt/jwt"
)
var path string = "bookeeping/account/"
var key  = []byte("secret-key-bookeeping")

type Account struct{
    db *Firebase
}

// result need to be this data struct
type AuthData struct {
   TelegramId string `json:"telegram_id,omitempty"`
   Email string `json:"email,omitempty"`
   Token string `json:"token"`
}

type UserClaims struct{
    UserName string `json:"username"`
    Password string `json:"password"`
    jwt.StandardClaims
}

func NewAccount(db *Firebase) (*Account) {
    construct := Account{db}

    return &construct
}

func (c *Account) GetAll() (interface{}) {
    var mapResult map[string]AuthData
    ok,_ := c.db.Get(path,&mapResult)

    if ok {
        return mapResult
    } else {
        return nil;
    }
}

func (c *Account) GetByName(name string) (interface{}) {
    var data AuthData
    ok,_ := c.db.Get(path + name,&data)
    if ok {
        return data
    } else {
        return nil;
    }
}

func (c *Account) Login(user string, password string) (bool) {
    return c.auth(user,password)
}

func(c *Account) Register(user string,email string,password string)(*AuthData,error) {
    return c.setToDB(user,email,password)
}

//when register or reset password use this method
func (c *Account) setToDB(user string,email string,password string)(*AuthData,error) {
    authData := AuthData{
        Email:email,
    }
    claims := UserClaims{UserName:user,Password:password,}
    token := jwt.NewWithClaims(jwt.SigningMethodHS256,&claims)
    tokenString,err := token.SignedString(key)
    authData.Token = tokenString

    if c.db.Set(path + user,&authData) {
        return &authData,nil
    } else {
        return nil,err
    }
}

func (c *Account) auth(user string, password string) (bool) {
    var authData AuthData

    c.db.Get(path + user,&authData)

    if(authData.Token != "") {
        token,_ := jwt.ParseWithClaims(authData.Token,&UserClaims{},func(token *jwt.Token) (i interface{}, err error) {
            return key, nil
        })
        if claims,ok := token.Claims.(*UserClaims); ok && token.Valid {
            return password == claims.Password && user == claims.UserName
        } else {
            return false
        }
    }

    return false
}
