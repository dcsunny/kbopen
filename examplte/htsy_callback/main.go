package main

import (
	"fmt"
	"net/http"

	"github.com/dcsunny/kbopen/conf"
	"github.com/dcsunny/kbopen/htsy"
	"github.com/dcsunny/kbopen/htsy/callback"
)

func callbackFunc(w http.ResponseWriter, r *http.Request) {
	cfg := &conf.Config{
		Appid:            "",
		AppSecret:        "",
		AuthorizerUserId: "",
	}
	client := htsy.NewClient(cfg)
	c := client.CallbackByHttp(r, w)
	c.SetHandler(func(msg *callback.Message) {
		fmt.Println(msg.Type)
	})
	c.Serve()
}

func main() {
	http.HandleFunc("/callback", callbackFunc)

	fmt.Println("Server is listening on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
