package main

/// TODO: remove this on the "master" branch, or even replace it
// with the "iris-mvc" (the new implementatioin is even faster, close to handlers version,
// with bindings or without).

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func main() {
	app := iris.New()
	mvc.New(app.Party("/api/values/{id}")).
		Handle(new(ValuesController))

	app.Run(iris.Addr(":5000"), iris.WithoutVersionChecker)
}

type ValuesController struct{
	Ctx iris.Context
}

// Get handles "GET" requests to "api/values/{id}".
func (c *ValuesController) Get() string {
	return "value"
}

// Put handles "PUT" requests to "api/values/{id}".
func (vc *ValuesController) Put() {}

// Delete handles "DELETE" requests to "api/values/{id}".
func (vc *ValuesController) Delete() {}

// +2MB/s faster than the previous implementation, 0.4MB/s difference from the raw handlers.
