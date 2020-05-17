package Services

import (
	"context"
	"encoding/json"
	"fmt"
	"product.jtthink.com/App"
	"product.jtthink.com/App/RedisUtil"
	"product.jtthink.com/AppInit"
)

type BookMetaRequest struct {
	Type string `json:"type"`
	UserID int `json:"uid"`
	BookID int `json:"bookid"`
}

///prods?size=
type BookListRequest struct {
	Size int `form:"size"`
	Type string `form:"t"`
}
type BookResponse struct {
	Result interface{} `json:"result"`
	Metas interface{} `json:"metas"`
}

//  /prods/300
type BookDetailRequest struct {
	BookID int `uri:"id" binding:"required,gt=0,max=70000"`
}


//图书列表相关的业务最终函数
func BookListEndPoint(book *BookService)  App.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req:=request.(*BookListRequest)
		result,err:=book.LoadBookList(req)
		return &BookResponse{Result:result},err
	}
}

//缓存中间件
func BookDetailCache() App.Middleware{
	return func(next App.Endpoint) App.Endpoint{
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			req:=request.(*BookDetailRequest)
			//假设缓存key = books123445
			cacheKey:=fmt.Sprintf("%s%d","books",req.BookID)
			redConn:=AppInit.RedisDefaultPool.Get()
			defer redConn.Close()
			strData:=RedisUtil.NewStringData()
			//return strData.Get(cacheKey).Then(func(rsp []byte) (interface{},error) {
			//	bookrsp:=&BookResponse{}
			//	if RedisUtil.JsonDecode(rsp,bookrsp){
			//		return bookrsp,nil
			//	}
			//	return nil,nil
			//}, func() (interface{},error) {  //缓存中没有
			//	rsp,err:=next(ctx,request)//从数据库取
			//	if rsp!=nil && err==nil{
			//		jsonRet:=RedisUtil.JsonEncode(rsp)
			//		if _,ok:=jsonRet.(string);ok{
			//			//redConn.Do("setex",cacheKey,20,jsonStr) //插入缓存
			//			strData.Set(cacheKey,rsp,20,false)
			//		}
			//	}
			//	return rsp,err
			//})
			strRet:=<-strData.Get(cacheKey)
			//b,redErr:= redis.Bytes(redConn.Do("get",cacheKey))
			if strRet==nil{ //缓存里没有
				rsp,err:=next(ctx,request)//从数据库取
				if rsp!=nil && err==nil{
					jsonRsp,_:=json.Marshal(rsp)
					redConn.Do("setex",cacheKey,20,jsonRsp) //插入缓存
				}
				return rsp,err
			}else { //
				bookrsp:=&BookResponse{}
				_=json.Unmarshal(strRet,bookrsp)
				return bookrsp,nil
			}
		}
	}
}



//图书详细
func BookDetailEndPoint(book *BookService)  App.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req:=request.(*BookDetailRequest)
		result,metas,err:=book.LoadBookDetail(req)
		return &BookResponse{Result:result,Metas:metas},err
	}
}
//收藏图书最终函数
func BookFavEndPoint(book *BookService)  App.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req:=request.(*BookMetaRequest)
		getError:=book.BookFav(req)
		if err!=nil{
			return &BookResponse{Result:"error"},getError
		}
		return &BookResponse{Result:"success"},nil
	}
}

