package controllers

import (
	"github.com/birjemin/iris-structure/datasource"
	pb "github.com/birjemin/iris-structure/grpc/pb"
	"github.com/birjemin/iris-structure/models"
	"github.com/kataras/iris/v12"
)

type ClientController struct {
	Ctx     iris.Context
}


func NewClientController() *ClientController {
	return &ClientController{}
}

//GET http://localhost:8081/hello?name=ahh
func (c *ClientController) Get() (result models.Result) {
	r := c.Ctx.Request()
	name := r.FormValue("name")
	// Contact the server and print out its response.
	client := pb.NewGreeterClient(datasource.GetGRPC())
	res, err := client.SayHello(r.Context(), &pb.HelloRequest{Name: name})
	if err != nil {
		result.Msg = "错误"
		result.Code = -1
		result.Data = err
		return
	}
	result.Msg = ""
	result.Code = -1
	result.Data = res
	return
}