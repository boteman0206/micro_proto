# micro_proto
proto文件



添加子模块
git submodule add https://gitlab.com/b-project



生成的proto文件的命令  cd到指定的proto的目录文件下面，然后执行
protoc --go_out=plugins=grpc:. ./*.proto 

