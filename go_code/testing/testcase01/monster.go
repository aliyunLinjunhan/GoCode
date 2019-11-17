package monster
import (
	"encoding/json"
	"io/ioutil"
	"fmt"
)

type Monster struct {

	Name string
	Age int
	Skill string
}

//给Monster绑定方法store，可以将一个Monster变量（对象），序列化后保存到文件中
func (this *Monster) Store() bool{

	// 先序列化
	data, err := json.Marshal(this)
	if err != nil {
		fmt.Println("marshal err =", err)
		return false
	}

	// 保存文件
	filePath := "monster.ser"
	err = ioutil.WriteFile(filePath, data, 0666)
	if err != nil{
		fmt.Println(err)
		return false
	}
	return true
}

// 给Monster绑定方法Restore，可以将一个序列化
func (this *Monster) ReStore() bool{

	filePath := "monster.ser"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("ReadFile err = ", err)
		return false
	}

	err = json.Unmarshal(data, this)
	if err != nil {
		fmt.Println("UnMArshal err = ", err)
		return false
	}
	return true
}