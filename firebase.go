package bookeeping

import (
    "context"
    "log"
    firebase "firebase.google.com/go"
    "google.golang.org/api/option"
    db "firebase.google.com/go/db"
)

type Firebase struct {
    client *db.Client
    ctx context.Context
}

func newFirebase(accountPath string,projectName string) (*Firebase,error) {
    ctx := context.Background()
    conf := &firebase.Config{
        //DatabaseURL: "https://euphoric-anchor-238602.firebaseio.com",
        DatabaseURL: "https://"+projectName+".firebaseio.com",
    }

    opt := option.WithCredentialsFile(accountPath)
    app, err := firebase.NewApp(ctx, conf, opt)

    if err != nil {
        log.Fatalln("Error initializing database client:", err)

        return nil,err
    }

    client, err := app.Database(ctx)

    if err != nil {
        log.Fatalln("Error initializing database client:", err)

        return nil,err
    }
    construct := Firebase{client: client,ctx:ctx}

    return &construct,nil
}

//method for client object
func (c *Firebase) Set(path string,data interface{}) (bool) {
    ref := c.client.NewRef(path)
    if err := ref.Set(c.ctx, data); err != nil {
        log.Fatalln("Error initializing database client:", err)

        return false
    }

    return true
}

func (c *Firebase) Push(path string,data interface{}) (string,error) {
    ref := c.client.NewRef(path)
    newRef,err := ref.Push(c.ctx, data);
    if err != nil {
        log.Fatalln("Error initializing database client:", err)

        return "",err
    }

    return newRef.Key,nil
}

func (c *Firebase) Get(path string,st interface{}) (bool,error) {
    ref := c.client.NewRef(path)

    if err := ref.Get(c.ctx, st); err != nil {
        log.Fatalln("Error initializing database client:", err)

        return false,err
    }

   return true,nil
}

