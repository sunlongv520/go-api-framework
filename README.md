# go-api-framework
go基于gin三层架构web框架

**三层架构模式**




```
func RegisterHandler(业务最终函数,怎么取参数,怎么处理业务结果) func(context *gin.Context) {

       xxxxxoooo


}

这个就是最终的结果

```


```
unc RegisterHandler(业务最终函数,怎么取参数,怎么处理业务结果) func(context *gin.Context) {

        参数:=怎么取参数() 
        业务结果:=业务最终函数(参数)
   
        
        怎么处理业务结果(业务结果)

}

```

**首先要定义原型**


```
业务最终函数

       type Endpoint func(ctx context.Context,request interface{}) (response interface{}, err error)


   一律使用interface{}  。这样可以处理不同的类型

```



```
怎么取参数 ：

      type EncodeRequestFunc func(*gin.Context, interface{}) (interface{}, error)

 
 怎么处理响应：
    
        type DecodeResponseFunc func(*gin.Context,  interface{}) error

```

**然后写成这样**

```
func RegisterHandler(endpoint Endpoint,encodeFunc EncodeRequestFunc,decodeFunc DecodeResponseFunc) func(context *gin.Context){
    return func(context *gin.Context) {
		req,err:=encodeFunc(context,nil)
		if err!=nil{
			context.JSON(500,gin.H{"error":"param err"+err.Error()})
			return
		}
		res,err:=endpoint(context,req)
		if err!=nil{
			context.JSON(500,gin.H{"error":err})
		}else{
			err:=decodeFunc(context,res)
			if err!=nil{
				context.JSON(500,gin.H{"error":err})
			}
		}

	}
}

```





**最终核心结构：**

```
package App

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Middleware func(Endpoint) Endpoint
//业务最终函数原型
type Endpoint func(ctx context.Context,request interface{}) (response interface{}, err error)

//怎么取参数
type EncodeRequestFunc func(*gin.Context) (interface{}, error)

//怎么处理业务结果
type DecodeResponseFunc func(*gin.Context, interface{}) error

func RegisterHandler(endpoint Endpoint,encodeFunc EncodeRequestFunc, decodeFunc DecodeResponseFunc) func(context *gin.Context){
	return func(context *gin.Context) {

		defer func() {
			if r:=recover();r!=nil{
				fmt.Fprintln(gin.DefaultWriter,fmt.Sprintf("fatal error:%s",r))
				context.JSON(500,gin.H{"error":fmt.Sprintf("fatal error:%s",r)})
				return
			}
		}()
		//参数:=怎么取参数(context)
		//业务结果,error:=业务最终函数(context,参数)
		//
		//
		//怎么处理业务结果(业务结果)
		req,err:=encodeFunc(context) //获取参数
		if err!=nil{
			context.JSON(400,gin.H{"error":"param error:"+err.Error()})
			return
		}
		rsp,err:=endpoint(context,req) //执行业务过程
		if err!=nil{
			fmt.Fprintln(gin.DefaultWriter,"response error:",err)
			context.JSON(400,gin.H{"error":"response error:"+err.Error()})
			return
		}
		err=decodeFunc(context,rsp) //处理 业务执行 结果
		if err!=nil{
			context.JSON(500,gin.H{"error":"server error:"+err.Error()})
			return
		}

	}
}

```














