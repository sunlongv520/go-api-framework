package AppInit
import(
	"github.com/gomodule/redigo/redis"
	"time"
)



var RedisDefaultPool *redis.Pool

func newPool(addr string) *redis.Pool{
	return &redis.Pool{
		MaxIdle: 10,
		IdleTimeout: 240 * time.Second,
		Dial: func () (redis.Conn, error) { return redis.Dial("tcp", addr) },
	}
}


func init(){
	RedisDefaultPool = newPool("192.168.1.188:6379")
}