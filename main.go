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
func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

/** Range **/
var ppow = []int{1, 2, 4, 8, 16, 32, 64, 128}

/** Excercise : Slices **/
func Pic(dx, dy int) [][]uint8 {

	picture := make([][]uint8, dy)

	for i := 0; i < dy; i++ {
		picture[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			picture[i][j] = uint8((i - j) * (i + j))
		}
	}

	return picture
}

/** Maps **/
type Vertexx struct {
	Lat, Long float64
}

var m map[string]Vertexx

var mm = map[string]Vertexx{
	"Bell Labs": Vertexx{
		40.68433, -74.39967,
	},
	"Google": Vertexx{
		37.42202, -122.08408,
	},

	"Bell Labss": {
		40.68433, -74.39967,
	},
	"Googlee": {
		37.42202, -122.08408,
	},
}

/** Function closures **/
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

/** Method **/
type Vertexxx struct {
	X, Y float64
}

func (v *Vertexxx) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertexxx) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

/** Interface **/
type Abser interface {
	Abs() float64
}

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
	//}	/** Range **/
	//for i, v := range ppow {a
	//	fmt.Printf("2**%d = %d\n", i, v)
	//}

	/** pic **/
	//pic.Show(Pic)

	/** Maps **/
	//m = make(map[string]Vertexx)
	//m["Bell Labs"] = Vertexx{
	//	40.68433, -74.39967,
	//}
	//fmt.Println(m["Bell Labs"])
	//fmt.Println(mm)

	/** Mutating Maps **/
	//zm := make(map[string]int)
	//
	//zm["Answer"] = 42
	//fmt.Println("The Value:" , zm["Answer"])
	//
	//zm["Answer"] = 48
	//fmt.Println("The Value:" , zm["Answer"])
	//
	//delete(zm, "Answer")
	//fmt.Println("The Value:" , zm["Answer"])
	//
	//v, ok := zm["Answer"]
	//fmt.Println("The Value:" , v, "Present?", ok)

	/** Function values **/
	//hypot := func(x, y float64) float64 {
	//	return math.Sqrt(x*x + y*y)
	//}
	//
	//fmt.Println(hypot)

	/** Functino closures **/
	//pos, neg := adder(), adder()
	//for i := 0; i < 10; i++ {
	//	fmt.Println(
	//		pos(i),
	//		neg(-2*i),
	//		)
	//}

	/** Switch **/
	//fmt.Print("Go runs on ")
	//
	//switch os := runtime.GOOS; os {
	//case "darwin":
	//	fmt.Println("OS X.")
	//case "linux":
	//	fmt.Println("Linux.")
	//default:
	//	fmt.Printf("%s", os)
	//}

	/** Methods **/
	//vv := &Vertexxx{3,  4 }
	//fmt.Println(vv.Abs())

	//ff := MyFloat(-math.Sqrt2)
	//fmt.Println(ff.Abs())

	//ss := &Vertexxx{3, 4}
	//ss.Scale(5)
	//fmt.Println(ss, ss.Abs())

	/** Interface **/
	//var a Abser
	//f := MyFloat(-math.Sqrt2)
	//v := Vertexxx{3,4}
	//
	//a = f
	//a = &v
	//a = v
	//
	//fmt.Println(a.Abs())

}
