# message-push 多通道消息推送服务

流程：

![alt 属性文本](./doc/process.svg)

支持推送通道：

    - 微信公众号/测试号
    - 企业微信
    - 钉钉群机器人
    - 飞书群机器人
    - email
    - 短信
    - 智能音箱（待定）

项目结构：

```
├── api //service api
├── common  //global util
│   └── sender  //all message sender
├── config  //projrct config init
├── doc 
├── model   //struce(like oop model)
│   ├── constant    //
│   └── r   //unified return model
├── router
├── service
├── view
│   ├── static
│   │   ├── css│
│   │   ├── fonts
│   │   ├── img
│   │   └── js
│   └── templates
├── config.ini //config file
├── main.go //starter
```