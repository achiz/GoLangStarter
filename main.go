package main

import (
	"fmt"
	"math"
)

/** Func **/
func add(x, y int) int {
	return x + y
}

/** Mutiple Results **/
func swap(x, y string) (string, string) {
	return y, x
}

/** Named Results **/
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

/** Variables **/
var x, y, z int = 1, 2, 3
var c, python, java = true, false, "no!"
//var c, python, java bool

/** Constants **/
const Pi = 3.14

/** Numeric Constants **/
const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int           { return x*10 + 1 }
func needFloat(x float64) float64 { return x * 0.1 }

/** If **/
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		fmt.Println(v)
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

/** Structs **/
type Vertex struct {
	X int
	Y int
}

/** Struct Literals **/
var (
	p = Vertex{1, 2}
	q = &Vertex{1, 2}
	r = Vertex{X: 1}
	s = Vertex{}
)


/** Make Slice **/
func printSlice( s string, x []int ) {
	fmt.Printf("%s len=%d cap=%d %v\n" , s, len(x), cap(x), x )
}


/** Range **/
var ppow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	/** func **/
	//fmt.Print(add(42,11))

	/** Mutiple Result **/
	//a, b := swap("hello", "world")
	//fmt.Println(a, b)

	/** Named Result **/
	//fmt.Print(split(17))

	/** Variables **/
	//fmt.Print(x, y, z, c, python, java)
	//var x, y, z int = 1, 2, 3
	//c, python, java := true, false, "no!"
	//fmt.Print(x, y, z, c, python, java)

	/** Constants **/
	//const World = "안녕"
	//fmt.Println("Hello", World)
	//fmt.Println("Happy", Pi, "Day")
	//
	//const Truth = true
	//fmt.Println("Go rules?", Truth)

	/** Numeric Constants **/
	//fmt.Println(needInt(Small))
	//fmt.Println(needFloat(Small))
	//fmt.Println(needFloat(Big))

	/** For **/
	//sum := 0
	//for i := 0; i < 10; i++ {
	//	sum += i
	//}
	//fmt.Println(sum)

	/** If **/
	//fmt.Println(sqrt(2), sqrt(-4))
	//fmt.Println(
	//	pow(3, 2, 10),
	//	pow(3, 3, 20),
	//	)

	/** Structs **/
	//fmt.Println(Vertex{1,2})
	//v := Vertex{1, 2}
	//v.X = 4
	//fmt.Println(v.X)

	/** Pointers **/
	//p := Vertex{1, 2}
	//q := &p
	//q.X = 1e9
	//fmt.Println(p)

	/** Struct Literals **/
	//fmt.Println(p, q, r, s)

	/** new **/
	//v := new(Vertex)
	//fmt.Println(v)
	//v.X, v.Y = 11, 9
	//fmt.Println(v)

	/** Slices **/
	//p := []int{2, 3, 5, 7, 11, 13}
	//fmt.Println("p ==", p)

	//for i := 0; i < len(p); i++ {
	//	fmt.Printf("p[%d] == %d\n",i ,p[i])
	//}


	/** Slices **/
	//p := []int{2, 3, 5, 7, 11, 13}
	//fmt.Println("p == " , p )
	//fmt.Println("p[1:4] ==", p[1:4])
	//
	//fmt.Println("p[:3] == ", p[:3])
	//fmt.Println("p[4:] == ", p[4:])


	/** Make Slice **/
	//a := make([]int, 5)
	//printSlice("a", a)
	//b := make([]int, 0, 5)
	//printSlice("b", b)
	//c := b[:2]
	//printSlice("c", c)

	/** Empty Slice **/
	//var z []int
	//fmt.Println(z, len(z), cap(z))
	//if z == nil {
	//	fmt.Println("nil!")
	//}

	/** Range **/
	//for i, v := range ppow {
	//	fmt.Printf("2**%d = %d\n", i, v)
	//}

	xpow := make([]int, 10)
	for i := range xpow {
		xpow[i] = 1 << uint(i)
	}

	for _, value := range xpow {
		fmt.Printf("%d\n", value)
	}





}
