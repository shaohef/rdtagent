# rdtagent
a daemon for Intel rdt

Run this daemon
$ cd rdtagent
$ make lib
$ make install-lib
$ GOPATH=`pwd` go get "github.com/emicklei/go-restful"
$ GOPATH=`pwd` go get "github.com/emicklei/go-restful-swagger12"
$ GOPATH=`pwd` go run src/rdtagent.go
