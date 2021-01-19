package main

import (
	"fmt"
)

func main() {
	var want int
	var first int
	var second int
	sum := 0

	fmt.Println("안녕하세요. 태엽 계산기입니다. 어떤 연산을 도와드릴까요?")
	fmt.Println("1. 덧셈 2. 뺄셈 3. 곱셈 4. 나눗셈")
	fmt.Println("원하시는 연산의 번호를 입력해주세요")
	fmt.Scanln(&want)

	switch want {
	case 1:
		fmt.Println("덧셈 하려는 두 수를 입력해주세요")
		fmt.Scanln(&first, &second)
		sum = first + second
		fmt.Println(sum)

	case 2:
		fmt.Println("뺄셈 하려는 두 수를  입력해주세요 ")
		fmt.Scanln(&first, &second)
		sum = first - second
		fmt.Println(sum)
	case 3:
		fmt.Println("곱셈 하려는 두 수를  입력해주세요 ")
		fmt.Scanln(&first, &second)
		sum = first * second
		fmt.Println(sum)
	case 4:
		fmt.Println("나눗셈 하려는 두 수를 입력해주세요 ")
		fmt.Scanln(&first, &second)
		sum = first / second
		fmt.Println(sum)
	}
}
