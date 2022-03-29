package switchSample

var (
	i interface{}
)

func printType(i interface{}) {
	switch t := i.(type) {
	case int:
		println("i is interger", t)
	case string:
		println("i is string", t)
	case float64:
		println("i is float64", t)
	default:
		println("type not found")
	}
}

func CheckSwitchWithInterface() {
	i = 1000
	printType(i)
	i = float64(12.345)
	printType(i)
	i = "poo"
	printType(i)
	printType(float32(78.0))
}
