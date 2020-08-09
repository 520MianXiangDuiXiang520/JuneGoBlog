# JuneGoBlog

<center>
<a href="https://travis-ci.org/">
  <img src="https://travis-ci.com/520MianXiangDuiXiang520/JuneGoBlog.svg?token=7mqBvrpUUzHXp1nyitHA&branch=master">
</a>
<a href="https://gitmoji.carloscuesta.me">
  <img src="https://img.shields.io/badge/gitmoji-%20😜%20😍-FFDD67.svg?style=flat-square" alt="Gitmoji">
</a>
<a href='https://coveralls.io/github/520MianXiangDuiXiang520/JuneGoBlog?branch=master'><img src='https://coveralls.io/repos/github/520MianXiangDuiXiang520/JuneGoBlog/badge.svg?branch=master' alt='Coverage Status' /></a>

</center>

## 接口列表

|路径|描述|详情|
|----|----|----|
|api/friendship/list|获取所有友链列表|[api/friendship/list](#apifriendshiplist)|

### 接口详情

#### api/friendship/list

请求：

1. Method: POST
2. 请求参数： 无

响应：

```json
{
    "header": {
        "code": 200,
        "msg": "ok"
    },
    "total": 2,
    "friendShipList": [
        {
            "id": 1,
            "siteName": "DeepBlue的小站",
            "link": "http://dlddw.xyz/",
            "imgLink": "https://junebao.top/static/image/friends/dlddw.png",
            "intro": ""
        },
        {
            "id": 2,
            "siteName": "异国迷宫的十字路口",
            "link": "https://blog.fivezha.cn/",
            "imgLink": "https://blog.fivezha.cn/img/avatar.png",
            "intro": ""
        }
    ]
}
```
