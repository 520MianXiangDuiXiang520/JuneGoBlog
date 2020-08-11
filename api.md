# API

1. 所有请求方式均为 POST
2. 所有请求和响应格式均为 JSON
3. 所有响应都包含 header

## 接口列表

|路径|描述|详情|
|----|----|----|
|api/friendship/list|获取所有友链列表|[api/friendship/list](#apifriendshiplist)|
|api/friendship/add| 添加友链 |[api/friendship/add](#apifriendshipadd)     |

### 接口详情

### api/friendship

#### api/friendship/list

请求：


```json
{}
```

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

#### api/friendship/add

请求：

```json
{
    "siteName": "万般皆下品，唯有读书高",    // 必填
    "intro": "小生的学习笔记",
    "imgLink": "https://wanghao15536870732.github.io/uploads/icon.jpg",
    "siteLink": "https://wanghao15536870732.github.io/"   // 必填
}
```

响应：

```json
{
    "header": {
        "code": 200,
        "msg": "ok"
    }
}
```