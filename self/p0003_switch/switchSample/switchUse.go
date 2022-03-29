package switchSample

import "fmt"

const (
	Monday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

const (
	A = iota
	B
)

func calc() (int, int) {
	return 1, 2
}

func CheckSwitchUse() {
	fmt.Println(Monday)
	fmt.Println(Tuesday)
	fmt.Println(Wednesday)
	fmt.Println(Thursday)
	fmt.Println(Friday)
	fmt.Println(Saturday)
	fmt.Println(Sunday)
	fmt.Println(A)
	fmt.Println(B)

	test := true

	if test, test2 := calc(); test+test2 < 10 {
		fmt.Println("test:", test)
		fmt.Println("test2:", test2)
		fmt.Println("test + test2 < 10")
	}

	fmt.Println("test:", test)

	switch a := 1; {
	case a >= 0:
		fmt.Println("a is true")
		fallthrough
	case a == 1:
		fmt.Println("a == 1")
		fallthrough
	case a > 200:
		fmt.Println("a is false")
	}

	stringSlice1 := []string{"1", "2", "3", "4"}
	fmt.Println(stringSlice1[:2])
	fmt.Println(stringSlice1[2:])

	stringSlice2 := make([]string, 5, 10)
	copy(stringSlice2, stringSlice1)
	fmt.Println(stringSlice2[0])

	stringSlice2 = append(stringSlice2, "5", "6")
	fmt.Println(stringSlice2)
	fmt.Println(stringSlice2[5])
	fmt.Println(stringSlice2[6])
	stringSlice2 = append(stringSlice2, []string{"7", "8"}...)
	fmt.Println(stringSlice2)
	fmt.Println(stringSlice2[7])
	fmt.Println(stringSlice2[8])
}
