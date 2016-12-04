package sess

import (
	"github.com/vodstv/core"

	"github.com/gin-gonic/contrib/sessions"
	_ "github.com/jinzhu/gorm/dialects/postgres" //import postgres
)

const (
	//sessHost ...
	sessHost = "localdocker"
	//sessUser ...
	sessUser = "esvodsapi"
	//sessPassword ...
	sessPassword = "esvodsapi"
	//sessName ...
	sessName = "esvods"
)

//Init ...
func Init() sessions.RedisStore {

	//connect to redis
    u, err := url.Parse(os.Getenv("REDIS_URL"))
    if err != nil {
        panic(err)
    }
    
	s, err := sessions.NewRedisStore(10, "tcp", u.Host(), u.User.Password(), []byte("secret"))
	core.CheckErr(err, "Cannot connect to RedisStore")

	//save connection
	return s

}
