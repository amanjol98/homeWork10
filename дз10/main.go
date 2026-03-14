package main

import (
	"fmt"
	"sort"
)

func main() {
	// Задание 1

	// numbers := []int{}

	// var num int

	// for {
	// 	fmt.Print("Введите число: ")
	// 	fmt.Scan(&num)

	// 	if num == 0 {
	// 		break
	// 	}
	// 	numbers = append(numbers, num)

	// }

	// fmt.Println(numbers)

	// sum := 0

	// for _, val := range numbers {
	// 	sum += val
	// }

	// fmt.Println(sum)

	// Задание 2

	// var values []int
	// var evenValues []int
	// var num int

	// for {
	// 	fmt.Print("Введите число: ")
	// 	fmt.Scan(&num)
	// 	if num == 0 {
	// 		break
	// 	}
	// 	values = append(values, num)
	// }

	// for _, v := range values {
	// 	if v%2 == 0 {
	// 		evenValues = append(evenValues, v)
	// 	}
	// }

	// fmt.Println(values)
	// fmt.Println(evenValues)

	// Задание 3

	// data := []int{}

	// var num int

	// for {
	// 	fmt.Print("Введите число: ")
	// 	fmt.Scan(&num)
	// 	if num == 0 {
	// 		break
	// 	}
	// 	data = append(data, num)
	// }
	// fmt.Println(data)

	// n := 2

	// data = append(data[:n], data[n+1:]...)

	// fmt.Println(data)

	// Задание 4

	// temps := []int{}
	// var temp int

	// for {
	// 	fmt.Print("Введите температуру: ")
	// 	fmt.Scan(&temp)
	// 	if temp == 0 {
	// 		break
	// 	}
	// 	temps = append(temps, temp)
	// }

	// fmt.Println("Заданные температуры: ", temps)

	// min_temp := temp
	// max_temp := temp

	// for _, temp := range temps {
	// 	if temp > max_temp {
	// 		max_temp = temp
	// 	}

	// 	if temp < min_temp {
	// 		min_temp = temp
	// 	}
	// }

	// fmt.Printf("Максимальная температура: %d, минимальная температура: %d\n", max_temp, min_temp)

	// Задание 5

	// var words []string

	// var word string

	// for {
	// 	fmt.Print("Введите слово: ")
	// 	fmt.Scan(&word)
	// 	if word == "stop" {
	// 		break
	// 	}
	// 	words = append(words, word)
	// }

	// fmt.Println(words)

	// var reversed []string

	// for i := len(words) - 1; i >= 0; i-- {
	// 	reversed = append(reversed, words[i])
	// }

	// fmt.Println(reversed)

	// Задание 6

	var nums []int
	var num int

	for {
		fmt.Print("Введите число: ")
		fmt.Scan(&num)

		if num == 0 {
			break
		}

		nums = append(nums, num)
	}

	fmt.Println(nums)

	copyNums := make([]int, len(nums))

	copy(copyNums, nums)

	sort.Ints(copyNums)

	fmt.Println(copyNums)

	isSorted := true

	for i := range nums {
		if nums[i] != copyNums[i] {
			isSorted = false
			break
		}
	}

	fmt.Println(isSorted)

	// Задание 7

	// var myWishList []string
	// var friendsWishList []string

	// var myList string
	// var friendsList string

	// for {
	// 	fmt.Print("Введи свой список желании: ")
	// 	fmt.Scan(&myList)

	// 	if myList == "stop" {
	// 		break
	// 	}

	// 	myWishList = append(myWishList, myList)
	// }

	// fmt.Println(myWishList)

	// for {
	// 	fmt.Print("Введи список желаении своего друга: ")
	// 	fmt.Scan(&friendsList)

	// 	if friendsList == "stop" {
	// 		break
	// 	}

	// 	friendsWishList = append(friendsWishList, friendsList)
	// }

	// fmt.Println(friendsWishList)

	// registrationList := append(myWishList, friendsWishList...)

	// fmt.Println(registrationList)

	// Задание 8

	toyList := []string{"Car", "Doll", "Ball"}
	testToyList := make([]string, len(toyList))
	copy(testToyList, toyList)

	testToyList[1] = "Boat"

	fmt.Println(toyList)
	fmt.Println(testToyList)

}
