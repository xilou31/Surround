# Auth 接口文档

## register 注册 POST

### 参数

| 字段  |类型 | 描述 |
| ---  | --- | --- |
| name | String | 用户名 |
| email | String | 电子邮箱 |
| password | String | 密码 |

### 返回值

``` json
HTTP/1.1 200
{
  "message": "",
  "status_code":0,
}
```

``` json
HTTP/1.1 422
{
  "message": "验证不通过",
  "errors": {
    "name": [
      "用户名不能为空"
    ],
    "email": [
      "邮箱不能为空"
    ],
    "password": [
      "密码不能为空"
    ]
  },
  "status_code": 422
}
```

## Login 登陆 POST

### 参数

| 字段  |类型 | 描述 |
| ---  | --- | --- |
| name | String | 用户名 |
| email | String | 电子邮箱 |
| password | String | 密码 |

``` json
HTTP/1.1 200
{
  "message": "",
  "status_code":0
}
```

``` json
HTTP/1.1 422
{
     "message": "验证不通过",
     "errors": {
         "name": ["用户名不能为空"],
         "email": ["邮箱不能为空"],
         "password": ["密码不能为空"],
         "geetest_challenge": ["请点击以滑动校验验证码"],
         "geetest_validate": ["请点击以滑动校验验证码"],
         "geetest_seccode": ["请点击以滑动校验验证码"],
         "geetest_status": ["请点击以滑动校验验证码"]
     },
     "status_code": 422
 }
 HTTP/1.1 422（极验验证不通过）
{
     "message": "验证不通过",
     "errors": {
         "error": ["验证失败请重试"],
     },
     "status_code": 422
 }
HTTP/1.1 400
{
     "message": "验证不通过",
     "errors": {
         "error": ["邮箱或密码不正确"],
     },
     "status_code": 400
 }
```

## refresh 刷新Token GET

### 参数

| 字段  |类型 | 描述 |
| ---  | --- | --- |
| name | String | 用户名 |

```json
{
  "message": "",
  "status_code": 200
}
```

```json
{
  "message": "用户没有登陆",
  "status_code": 400
}
```