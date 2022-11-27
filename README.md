1.1待更新：所有通道独立管理，使用时可以自己进行聚合，通过聚合通道进行请求，更加灵活。

![logo](./view/static/img/dobby.jpeg)

# dobby 多通道消息推送服务

界面：
![截图1](./doc/1.png)
![截图2](./doc/2.png)
![截图3](./doc/3.png)

流程：

![流程](./doc/process.svg)

支持推送通道：

    - 微信公众号/测试号
    - 企业微信webhook机器人
    - 钉钉webhook机器人
    - 飞书webhook机器人
    - email
    - 智能音箱（待定）

项目结构：

```
├── api //service api
├── common  //global util
│   └── sender  //all message sender
├── bootstrap  //projrct config init
├── doc 
├── model   //struce(like oop model)
│   ├── constant    //
│   └── r   //unified return model
├── router
├── service
├── view
│   ├── static
│   │   ├── css
│   │   ├── fonts
│   │   ├── img
│   │   ├── js
│   │   └── module
│   └── templates
├── config.ini //config file
├── main.go //starter
```

安装说明：
    
```
    
    1、配置go环境，clone本项目。
    2、打包。
    3、创建一个sqlite的.db文件。
    4、编辑config.ini配置文件，修改相应配置。
    5、将编辑好的config.ini文件，view文件夹，db文件与打包后的执行文件放置在同一目录。
    6、运行（如果配置无误，第一次运行时会自动建表）


```

config.ini配置文件说明：
```

[Web]
#监听地址和端口，gin原生配置。
port = 8082
addr =

[DB]
#sqlite .db文件路径
url = /Users/codecv/push.db
username =
password =

#默认邮件服务，不配置将无法登录
#一把情况下去你的电子邮件服务商的设置页面都可以查看并设置smtp服务
[Email]
host = smtp.163.com
port = 25
username : os7blue@163.com
password :
#切记 这是可以登录管理后台的邮箱，所有账户都具有相同权限，默认会将smtp自身的邮件地址加入。
admin:
[Setting]
#所有推送内容前置标题，会与实际推送标题进行拼接，也可用实现一些webhook通道的关键字认证
#如果调用推送api时 设置了 了title 那么会展现为为 pushTitle[你的title]：xxxxcontent 这种格式
pushTitle = dobby



```

其他说明：
```
    
    登录页面背景图是我自己拍的星空和领导手绘的图。
    文件有点大，可以自己替换自己喜欢的，或者直接删掉，设置了默认背景色，但是我觉得不好看。
    我也懒得在css上下太多功夫，就这样吧。
    
    只提供文字的推送的原因：
        1、本来写这个就是为了有个提醒，文字就够了。
        2、如果你需要更丰富的推送内容，你可以fork自己修改，我代码分的层级应该还行，另外个人认为，如果你的需求中需要推送复杂的内容，那你应该自己着手去定制。
        3、坦白说，如果不是为了有个可视化管理界面，写这个东西根本不需要这么麻烦。
        
    从我开始学go到做完这个东西用的时间加起来不超过半个月，如有不妥之处，那就不妥着吧。

```
