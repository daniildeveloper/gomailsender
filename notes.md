# Notes for beego. Tips and tricks
It is my first project with Golang framework beego.

## Routing
```go router.go
beego.Roter("/hello/:id([0-9]+)", &controllers.MainController{}, "get:Hello")
```

```go controller.go
c.Data["id"] = c.Ctx.Input.Param(":id") //import from context to Data
```

