package main

import "fmt"

type Foo struct {
	Name string
}

func main() {
	var cId, pId uint64 = 4229, 5237
	date := "2024-07-10"
	key := fmt.Sprintf("mall:goods_total_stock_p:%d:cid:%d", pId, cId)
	fmt.Println(key)

	key2 := fmt.Sprintf("mall:"+"stock-ex"+":cid:%d"+"pid:%d"+":%s", cId, pId, date)

	fmt.Println(key2)

}
