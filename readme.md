# goctl

### 说明:
本项目为goctl定制化版本,原始版本见[官方](https://github.com/zeromicro/go-zero),主要定制内容:
* 生成的gozero api和rpc项目引入cobra命令行工具,解决项目有脚本开发的需求
* 生成的gozero api和rpc项目实现配置单例

## goctl 用途

* 定义api请求
* 根据定义的api自动生成golang(后端), java(iOS & Android), typescript(web & 晓程序)，dart(flutter)
* 生成MySQL CURD 详情见[goctl model模块](model/sql)

## goctl 使用说明

### goctl 参数说明

  `goctl api [go/java/ts] [-api user/user.api] [-dir ./src]`

  > api 后面接生成的语言，现支持go/java/typescript
  >
  > -api 自定义api所在路径
  >
  > -dir 自定义生成目录

#### API 语法说明

```golang
type int userType

type user {
	name string `json:"user"` // 用户姓名
}

type student {
	name string `json:"name"` // 学生姓名
}

type teacher {
}

type (
	address {
		city string `json:"city"` // 城市
	}

	innerType {
		image string `json:"image"`
	}

	createRequest {
		innerType
		name    string    `form:"name"`
		age     int       `form:"age,optional"`
		address []address `json:"address,optional"`
	}

	getRequest {
		name string `path:"name"`
		age  int    `form:"age,optional"`
	}

	getResponse {
		code    int     `json:"code"`
		desc    string  `json:"desc,omitempty"`
		address address `json:"address"`
		service int     `json:"service"`
	}
)

service user-api {
    @server(
        handler: GetUserHandler
        group: user
    )
    get /api/user/:name(getRequest) returns(getResponse)

    @server(
        handler: CreateUserHandler
        group: user
    )
    post /api/users/create(createRequest)
}

@server(
    jwt: Auth
    group: profile
)
service user-api {
    @handler GetProfileHandler
    get /api/profile/:name(getRequest) returns(getResponse)

    @handler CreateProfileHandler
    post /api/profile/create(createRequest)
}

service user-api {
    @handler PingHandler
    head /api/ping()
}
```

1. type部分：type类型声明和golang语法兼容。
3. service部分：service代表一组服务，一个服务可以由多组名称相同的service组成，可以针对每一组service配置group属性来指定service生成所在子目录。
   service里面包含api路由，比如上面第一组service的第一个路由，GetProfileHandler表示处理这个路由的handler，
   `get /api/profile/:name(getRequest) returns(getResponse)` 中get代表api的请求方式（get/post/put/delete）, `/api/profile/:name` 描述了路由path，`:name`通过
   请求getRequest里面的属性赋值，getResponse为返回的结构体，这两个类型都定义在2描述的类型中。

#### api vscode插件

开发者可以在vscode中搜索goctl的api插件，它提供了api语法高亮，语法检测和格式化相关功能。

 1. 支持语法高亮和类型导航。
 2. 语法检测，格式化api会自动检测api编写错误地方，用vscode默认的格式化快捷键(option+command+F)或者自定义的也可以。
 3. 格式化(option+command+F)，类似代码格式化，统一样式支持。

#### 根据定义好的api文件生成golang代码

  命令如下：  
  `goctl api go -api user/user.api -dir user`

  ```Plain Text
  .
  ├── cmd (新增)
  │   ├── root.go
  ├── etc
  │   ├── user.yaml
  ├── internal
  │   ├── config
  │   │   └── config.go (修改)
  │   ├── handler
  │   │   ├── pinghandler.go
  │   │   ├── profile
  │   │   │   ├── createprofilehandler.go
  │   │   │   └── getprofilehandler.go
  │   │   ├── routes.go
  │   │   └── user
  │   │       ├── createuserhandler.go
  │   │       └── getuserhandler.go
  │   ├── logic
  │   │   ├── pinglogic.go
  │   │   ├── profile
  │   │   │   ├── createprofilelogic.go
  │   │   │   └── getprofilelogic.go
  │   │   └── user
  │   │       ├── createuserlogic.go
  │   │       └── getuserlogic.go
  │   ├── svc
  │   │   └── servicecontext.go
  │   └── types
  │       └── types.go
  └── user.go (修改)
  ```
  *定制位置已经标注

  生成的代码可以直接跑，有几个地方需要改：

* 在`servicecontext.go`里面增加需要传递给logic的一些资源，比如mysql, redis，rpc等
* 在定义的get/post/put/delete等请求的handler和logic里增加处理业务逻辑的代码

#### 根据定义好的api文件生成java代码

```Plain Text
goctl api java -api user/user.api -dir ./src
```

#### 根据定义好的api文件生成typescript代码

```Plain Text
goctl api ts -api user/user.api -dir ./src -webapi ***
```

ts需要指定webapi所在目录

#### 根据定义好的api文件生成Dart代码

```Plain Text
goctl api dart -api user/user.api -dir ./src
```
