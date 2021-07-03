package main

import (
  //"github.com/gin-gonic/gin"
  //"os"
  "fmt"
  "bookeeping/firebase"
  "bookeeping/account"
)

// key need to be upper for first char
type TestSt struct {
   Test string `json:"test,omitempty"`
   Test2 int64 `json:"test2,omitempty"`
}

func main() {
    firebase,_ := firebase.Init()
    account := account.Init(firebase)
    result :=account.Login("test","test2")
    fmt.Println(result)
    //x,_ := firebase.Get("account")
    //fmt.Println(x);
    //account := account.Init()
    //fmt.Println(account.GetByName("test"))
/*
    r := gin.Default()
	r.GET("/", func(c *gin.Context) {
        val := c.Request.Header.Get("XXX")
        c.Header("III","GGG")
		c.JSON(200, gin.H{
			"message": val,
            "xxx": gin.H{"yyy":"zzzz"},
		})
	})

	r.Run("0.0.0.0:"+os.Args[1])
*/
}


