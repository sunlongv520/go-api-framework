package main

import (
	"github.com/gin-gonic/gin"
	"product.jtthink.com/App"
	"product.jtthink.com/App/Services"
	. "product.jtthink.com/AppInit"
)

func main() {

	router:=gin.Default()
	v1:=router.Group("v1")
	{
		bookService  := &Services.BookService{}
		bookListHandler := App.RegisterHandler(
			Services.BookListEndPoint(bookService),//业务最终函数
			Services.CreateBookListRequest(),//怎么取参数
			Services.CreateBookListResponse(),//怎么处理响应
			)
		v1.Handle(HTTP_METHOD_GET,"/prods", bookListHandler)
	}
	router.Run(SERVER_ADDRESS)
}
