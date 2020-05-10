package Services

import (
	"context"
	"product.jtthink.com/App"

	"product.jtthink.com/Models"
)

type BookListRequest struct {
	Size int
}
type BookListResponse struct {
	Result *Models.BookList
}

func BookListEndPoint(book *BookService)  App.Endpoint {
   return func(ctx context.Context, request interface{}) (response interface{}, err error) {
	   req:=request.(*BookListRequest)
	   return &BookListResponse{Result:book.LoadBookList(req)},nil
   }
}