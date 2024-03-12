#### go mod init blog-api

#### go get -u github.com/gin-gonic/gin

#### go get -u github.com/go-swagger/go-swagger/cmd/swagger

##### 安装swagger可能会出现包引用错误，解决方案

`go mod edit -replace="github.com/imdario/mergo=github.com/imdario/mergo@v0.3.16"`

##### 生成映射文件
`swagger generate spec -o ./swagger.json`

##### 启动swagger UI
`swagger serve -F=swagger swagger.json`