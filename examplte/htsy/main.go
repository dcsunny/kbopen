package main

import (
	"fmt"
	"log"

	"github.com/dcsunny/kbopen/conf"
	"github.com/dcsunny/kbopen/htsy"
)

func main() {
	cfg := &conf.Config{
		Appid:     "",
		AppSecret: "",
	}
	client := htsy.NewClient(cfg, "")
	r, err := client.Account().Info()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(client.Ctx.HttpClient.HttpLastResult)
	fmt.Println(r)
}
