module github.com/bensema/goadmin

go 1.15

replace library => ./library

require (
	entgo.io/ent v0.8.0
	github.com/BurntSushi/toml v0.3.1
	github.com/gin-gonic/gin v1.6.3
	github.com/go-redis/redis/v8 v8.8.2
	github.com/json-iterator/go v1.1.10
	github.com/mssola/user_agent v0.5.2
	github.com/sirupsen/logrus v1.2.0
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	golang.org/x/image v0.0.0-20200801110659-972c09e46d76 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	library v0.0.0-00010101000000-000000000000
)
