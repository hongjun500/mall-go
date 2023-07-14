# 编译项目
build:
	go build -o myapp main.go

# 清理编译产物
clean:
	go clean
	rm -f myapp

# 运行项目
run:
	go run main.go

# 运行测试
test:
	go test ./...

# 安装依赖
deps:
	go mod download

.PHONY: build clean run test deps
