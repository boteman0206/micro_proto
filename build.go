package micro_proto

// 其他人写的工具
import "github.com/maybgit/pbgo"

func BuildProto() error {
	err := pbgo.Generate()
	return err
}
