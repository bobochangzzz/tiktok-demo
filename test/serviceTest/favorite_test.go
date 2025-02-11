package serviceTest

import (
	"SimpleDouyin/dao"
	"SimpleDouyin/middleware/DBUtils"
	"SimpleDouyin/service"
	"fmt"
	"testing"
)

var name = []string{"1", "2", "3", "4"}

/*
Tips：
单元测试需要在方法名前加大写开头的Test
需要在方法中传入测试的固定参数t *testing.T
可以使用log

基准测试相当于简易版的性能测试，会显示接口的执行速度，执行次数的信息

建议手动点一下，看看控制台
*/
func TestGetVedioLikeList(t *testing.T) {
	DBUtils.InitRedisTemplete()
	if DBUtils.RDB == nil {
		fmt.Println("初始化失败")
	}
	list := service.GetVedioLikeCount("2")
	fmt.Print("list is that:", list)
	service.GetVedioLikeList("2")

}

func TestAdddata(t *testing.T) {
	DBUtils.InitRedisTemplete()
	service.Add("17", "26")
	service.Add("17", "556")

}

func BenchmarkDislikeVedio(b *testing.B) {
	DBUtils.InitRedisTemplete()
	result := service.DislikeVedio("11", "15")
	fmt.Println(result)
}

func TestTimeClock(t *testing.T) {
	DBUtils.InitMysqlTemplete()
	fmt.Println(DBUtils.DB)
	//vedioIds := dao.GetVedioIdWithLimit(0, 99)
	//count := dao.GetVedioCount()
	//fmt.Println(vedioIds) //1-100
	//fmt.Println(count)    //100
	result := dao.UpdateVedioLikeCount(2, 999) //true
	fmt.Println(result)
}

func Test1(t *testing.T) {
	DBUtils.Init()

}
