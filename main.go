package main

import (
	"fmt"
)

type person struct {
	name    string
	age     int
	favFood []string
} //구조체로 대부분 다룸

func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
		fmt.Println(number)
	}
	return total
} //for range / range는 key, value 두 개를 return

func canIDrink(age int) bool {
	if korean_age := age + 2; korean_age > 19 {
		return true
	}
	switch age {
	case 15:
		return true
	case 16:
		return true
	default:
		return false
	}
} //if, switch / 조건문에 변수 새로 설정 가능

func main() {
	fmt.Println(superAdd(1, 2, 3, 4, 5, 6))
	fmt.Println(canIDrink(13))

	names := []string{"nico", "lynn"}
	names = append(names, "flyn") //go에서는 array, slice(크기 미정) append는 새로 할당해서 return 됨
	fmt.Println(names)
	var number = [4]int{4, 3, 2, 1}
	fmt.Println(number)

	nico := map[string]string{"name": "nico", "age": "12"}
	for key, value := range nico {
		fmt.Println(key, value)
	} //map / 새로운 자료구조

	favFood := []string{"kimchi", "ramen"}
	kyungbin := person{name: "kyungbin", age: 21, favFood: favFood} //구조체를 사용할 때 member 변수 이름도 같이 써주면 좋음
	fmt.Println(kyungbin)
}
