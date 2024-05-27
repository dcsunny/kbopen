## 酷班开放平台的golang版本的sdk

厦门酷班科技开放平台的sdk，利用旗下产品开放出来的能力开发微信公众号，企业微信

### 平台入口

[酷班开放平台](https://open.cpshelp.cn)

### 文档

[酷班开放平台文档](https://open.cpshelp.cn/document/server-docs/api-call-guide/calling-process/overview)

## 快速开始

```
import "github.com/dcsunny/kbopen"
```

---

## 海豚私域

[海豚私域](https://htsy.cpshelp.cn)

[海豚私域文档](https://china.feishu.cn/wiki/wikcnIpX7d79lzv5ip49GEnLg1c?fromScene=spaceOverview)

### 前提条件

1. 注册海豚私域
2. 将企微号授权到海豚私域内，可以使用hook的方式，也可以使用协议的方式
3. 注册酷班开放平台
4. 创建应用获取appid和secret
5. 授权海豚私域账号的权限

### 功能

1. 接受好友添加
2. 群管理
3. 发送消息给企微客户和群聊

### 示例

#### 获取账号信息

```
package main

import (
	"fmt"
	"log"

	"github.com/dcsunny/kbopen/conf"
	"github.com/dcsunny/kbopen/htsy"
)

func main() {
	cfg := &conf.Config{
		Appid:            "",
		AppSecret:        "",
		AuthorizerUserId: "",
	}
	client := htsy.NewClient(cfg)
	r, err := client.Account().Info()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(client.Ctx.HttpClient.HttpLastResult)
	fmt.Println(r)
}
```

#### 获取微号发送的消息

需要再应用配置回调地址

![htsy_callback.png](resources/htsy_callback.png)

```
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

```
