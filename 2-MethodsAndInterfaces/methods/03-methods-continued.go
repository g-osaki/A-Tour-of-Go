package main

type MyFloat float64

func (f MyFloat) anyTypeAbs() float64 { // 任意のtypeにもMethodを宣言できる
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

//func main() {
//	f := MyFloat303(-math.Sqrt2)
//	fmt.Println(f.Abs303())
//}
