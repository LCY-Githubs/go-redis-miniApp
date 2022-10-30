## 本系统的目录
|-app  
    |-apis   
    |-controller  
    |-model#与数据库表对应的映射类  
    |-routers#记录所有的路由地址并发布  
    |-service#接口发布后，调用时的处理方法


|-app2  
|-controller#处理接口与service的映射关系  
|-model#与数据库表对应的映射类  
|-routers#记录所有的路由地址并发布  
|-service#接口发布后，调用时的处理方法


##相关依赖
```CMD
go get github.com/swaggo/swag/cmd/swag

go get github.com/swaggo/gin-swagger

go get github.com/swaggo/files

go get github.com/alecthomas/template

```

##集成Swagger
```cmd
方式一
$ go get -u github.com/swaggo/swag/cmd/swag

方式二
# 1.16 or newer
$ go install github.com/swaggo/swag/cmd/swag@latest

```

##本项目的目的：做一个具有一些小功能组件的web前后端应用
> 1.便签功能
> * 能够新增待做项
> * 能够修改待做项的内容
> * 能标记待做项为已完成
> * 能按照一定规则对待做的东西进行整理归纳