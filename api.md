# API

1. 所有请求方式均为 POST
2. 所有请求和响应格式均为 JSON
3. 所有响应都包含 header

## 接口列表

|路径|描述|详情|权限|
|----|----|----|-----|
|api/friendship/list|获取所有友链列表|[api/friendship/list](#apifriendshiplist)|无|
|api/friendship/application| 申请添加友链 |[api/friendship/application](#apifriendshipapplication)     |无|
|api/friendship/delete| 删除友链| [api/friendship/delete](#apifriendshipdelete)|管理员权限|
|api/friendship/unshow|审核中和之前审核失败的友链列表|[unShow](#apifriendshipunshow)     |管理员权限|
|api/friendship/approval|审核友链（修改状态）|[approval](#apifriendshipapproval)|管理员权限
|api/tag/list| 获取所有标签 | [api/tag/list](#apitaglist)|无|
|api/tag/add|添加标签 | [api/tag/add](#apitagadd)|管理员权限|

## 接口详情

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

#### api/friendship/application

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

#### api/friendship/delete

请求

```json
{
    "id": 13
}
```

响应

```json
{
    "header": {
        "code": 200,
        "msg": "ok"
    }
}
```

#### api/friendship/unshow

请求

```json
{
    "status" : 1,
}
```

* 1: 查看请求中的友链
* 3：查看审核失败的友链
* 不传：查看 1， 3 状态的所以友链

响应

* 同 [api/friendship/list](#apifriendshiplist)

#### api/friendship/approval

请求：

```json
{
    "id": 1,
    "result": 2
}
```

* id: 要操作的友链 ID
* result: 操作结果：
  * 2： 审核通过
  * 3： 审核不通过

响应：

```json
{
    "header": {
        "code": 200,
        "msg": "ok"
    }
}
```

### api/tag

#### api/tag/list

请求

```json
{}
```

响应

```json
{
    "header": {
        "code": 200,
        "msg": "ok"
    },
    "total": 3,
    "tags": [
        {
            "id": 3,
            "name": "Java",
            "create_time": "2020-08-01T23:49:59+08:00",
            "article_total": 10
        },
        {
            "id": 1,
            "name": "Golang",
            "create_time": "2020-08-12T23:46:13+08:00",
            "article_total": 10
        },
        {
            "id": 2,
            "name": "Python",
            "create_time": "2020-08-12T23:46:41+08:00",
            "article_total": 10
        },
    ]
}
```

#### api/tag/add

请求

```json
{
    "name": "设计模式"
}
```

响应

```json
{
    "header": {
        "code": 200,
        "msg": "ok"
    }
}
```
