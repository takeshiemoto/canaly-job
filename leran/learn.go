package leran

import "fmt"

func Run() {
	array1 := [3]int{1, 2, 3}
	array2 := [...]int{1, 2, 3}
	array3 := [...]int{1: 3, 2: 2}

	slice1 := []int{1, 2, 3}
	map1 := map[int]string{1: "Hello", 2: "world"}

	fmt.Printf("%v+", array1)
	fmt.Printf("%v+", array2)
	fmt.Printf("%v+", array3)
	fmt.Printf("%v+", slice1)
	fmt.Printf("%v+", map1)
}

func RunPointer() {
	x := 10
	y := 20

	px := &x
	py := &y

	// xのポインタが出力される
	fmt.Printf("pointer px = %p\n", px)
	// yのポインタが出力される
	fmt.Printf("pointer py = %p\n", py)

	// デリファレンスすると参照されているアドレスに保存された値が取得できる
	// vxにはpxから取り出した値のコピーが入る
	vx := *px
	vy := *py

	fmt.Printf("pointer px = %d\n", vx)
	fmt.Printf("pointer py = %d\n", vy)

	// ポインタを渡す
	multiplyByTenMutable(px)
	fmt.Printf("pxは10倍されている -> %d\n", *px)

	// デリファレンスした値を渡す
	multiplyByTenImmutable(*py)
	fmt.Printf("pyは10倍されていない -> %d\n", *py)
}

// 元の値を10倍する関数（）
func multiplyByTenMutable(num *int) {
	*num = *num * 10
}

func multiplyByTenImmutable(num int) {
	// 値に10掛けているだけ
	num = num * 10
}
