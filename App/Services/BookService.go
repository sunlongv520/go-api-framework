package Services

import (

	"product.jtthink.com/AppInit"
	"product.jtthink.com/Models"
)

type BookService struct {
}

func(this *BookService) LoadBookList(req  *BookListRequest) *Models.BookList {
	prods:=&Models.BookList{}
	AppInit.GetDB().Limit(req.Size).Order("book_id desc").Find(prods)
	return prods
}