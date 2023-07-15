//go:build ignore

// Contains instructions for required generators, invoke by entering
// go generate gen.go

package main

//go:generate protoc --go_out=paths=source_relative:./pkg/domain/ --goconstrgen_out=paths=source_relative:./pkg/domain/ proto/gosearch/contracts/models.proto
// protoc is sometimes just awful
//go:generate rm -rf pkg/domain/contracts/models.con.go pkg/domain/contracts/models.pb.go
//go:generate cp pkg/domain/proto/gosearch/contracts/models.con.go pkg/domain/proto/gosearch/contracts/models.pb.go pkg/domain/contracts
//go:generate rm -r pkg/domain/proto/

//go:generate protoc --go-grpc_out=. --goblthttp_out=. proto/gosearch/contracts/service.proto
//go:generate handlergen

//go:generate cp pkg/app/server/contracts/service.http.json pkg/app/server/static/swagger/swagger.json

//go:generate wire ./pkg/app/server
