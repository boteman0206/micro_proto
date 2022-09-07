# micro_proto
proto文件


其中的hello模块是学习测试rpc的模块


添加子模块
git submodule add https://gitlab.com/b-project

直接拉取含有子模块的项目
git clone --recurse-submodules xxx/micro_product.git




生成的proto文件的命令  cd到指定的proto的目录文件下面，然后执行
protoc --go_out=plugins=grpc:. ./*.proto 




第二版proto和第三版的区别
第二版的 Protobuf 有个默认值特性，可以为字符串或数值类型的成员定义默认值。 在第三版的 Protobuf 中不再支持默认值特性，但是我们可以通过扩展选项自己模拟默认值特性。