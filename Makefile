run-server:
	go run cmd/server/main.go
gen:
	go generate gen.go
build-validation:
	docker build -f ./Dockerfile.test .
ci:
	docker build -f ./Dockerfile -t $(registry)/gex:$(env) .
	docker push $(registry)/gex:$(env)
setup:
	@prname='github.com/BetaLixT/gex';read -p "Enter Project Name ($$prname):" new; find . -type f  -not -path "./.git/*" -not -path "./pkg/app/server/static/swagger/*" -exec sed -i "" "s|$$prname|$$new|g" {} \;
	@prname='gex';read -p "Enter Service Name ($$prname):" new; mv proto/$$prname proto/$$new; find . -type f -not -path "./.git/*" -not -path "./pkg/app/server/static/swagger/*" -exec sed -i "" "s|$$prname|$$new|g" {} \;
setup-gen:
	go install github.com/google/wire/cmd/wire@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/BetaLixT/golang-tooling/protoc-gen-goconstrgen@latest
	go install github.com/BetaLixT/golang-tooling/protoc-gen-goblthttp@latest
	go install go install github.com/BetaLixT/golang-tooling/handlergen@latest
