package sess

import (
    "os"
    "net/url"
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
    // extract things from the parsed url
    host := u.Host
    password, err := u.User.Password()
    
    if err != nil {
        return nil, err
    }

    
	s, err := sessions.NewRedisStore(10, "tcp", host, password, []byte("secret"))
	core.CheckErr(err, "Cannot connect to RedisStore")

	//save connection
	return s

}
