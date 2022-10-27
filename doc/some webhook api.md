企业微信
```

{
	"msgtype": "text",
	"text": {
		"content": "广州今日天气：29度，大部分多云，降雨概率：60%",
		"mentioned_list":["wangqing", "@all"],
		"mentioned_mobile_list":["13800001111", "@all"]
	}
}

msgtype	是	消息类型，此时固定为text
content	是	文本内容，最长不超过2048个字节，必须是utf8编码
mentioned_list	否	userid的列表，提醒群中的指定成员(@某个成员)，@all表示提醒所有人，如果开发者获取不到userid，可以使用mentioned_mobile_list
mentioned_mobile_list	否	手机号列表，提醒手机号对应的群成员(@某个成员)，@all表示提醒所有人

企业微信群@个人需要获取单独获取id，本平台只涉及简单信息推送而已，没有必要越做越大，直接@所有人。
如果需要设置@的人，单独建个群又没有成本对吧。


*/

```

钉钉

```
{
    "at": {
        "atMobiles": [
            "180xxxxxx"
        ],
        "atUserIds": [
            "user123"
        ],
        "isAtAll": false
    },
    "text": {
        "content": "我就是我, @180xxxxxx 是不一样的烟火"
    },
    "msgtype": "text"
}
@个人会增加做这个项目的复杂度，因此不考虑@个人，本平台只涉及简单信息推送而已，没有必要越做越大，直接@所有人。
如果需要设置@的人，单独建个群又没有成本对吧。


*/

```

飞书

```

{
	"msg_type": "post",
	"content": {
		"post": {
			"zh_cn": {
				"title": "项目更新通知",
				"content": [
					[{
							"tag": "text",
							"text": "项目有更新: "
						},
						{
							"tag": "a",
							"text": "请查看",
							"href": "http://www.example.com/"
						},
						{
							"tag": "at",
							"user_id": "ou_18eac8********17ad4f02e8bbbb" all为@所有人
						}
					]
				]
			}
		}
	}
}

飞书群机器人@个人需要单独获取openid，因此不考虑@个人，本平台只涉及简单信息推送而已，没有必要越做越大，直接@所有人
如果需要设置@的人，单独建个群又没有成本对吧。

```
