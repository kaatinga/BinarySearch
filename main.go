package main

import (
	"log"
	"math/rand"
	"time"
)

const (
	sliceSize int32 = 100 // размер среза
	limit     int32 = 65535
)

func main() {

	// рандомизация рандома
	rand.Seed(time.Now().UnixNano())

	// Инициализируем срез с числами
	var numbers []uint16
	numbers = make([]uint16, sliceSize)
	numbers[0] = 1

	// Добавляем числа
	for i := 1; i < int(sliceSize); i++ {
		numbers[i] = numbers[i-1] + uint16(rand.Int31n(limit/sliceSize))
	}

	var aNumber uint16 = numbers[rand.Int31n(sliceSize)]

	log.Println("Chosen number is", aNumber)
	log.Println("Numbers:", numbers)

	iterationNumber, foundInIndex := search(&numbers, aNumber)
	check(numbers, foundInIndex, aNumber, iterationNumber)

	iterationNumber, foundInIndex = search(&numbers, 5)
	check(numbers, foundInIndex, 5, iterationNumber)
}

func check(numbers []uint16, foundInIndex uint16, aNumber uint16, iterationNumber int32) {
	if numbers[foundInIndex] == aNumber {
		log.Println("== Check index", foundInIndex, ": numbers[", foundInIndex, "] =", numbers[foundInIndex])
		log.Println("== function complexity is O(", iterationNumber, "/", sliceSize, ")")
		log.Println("== The number", aNumber, "was found for", iterationNumber, "iterations")
	} else {
		log.Println("The number", aNumber, "was NOT found")
	}
}

func search(numbers *[]uint16, aNumber uint16) (i int32, foundInIndex uint16) {

	var start, end, currentIndex uint16
	end = uint16(len(*numbers) - 1)

	for {

		time.Sleep(2 * time.Second)

		i++

		currentIndex = start + ((end - start) / 2)
		log.Println("current index is", currentIndex)
		log.Println("start is", start)
		log.Println("end is", end)

		switch {
		case start-end == 0 && (*numbers)[start] != aNumber:

			// это случается только в одном случае, если число не найдено
			return

		case (*numbers)[currentIndex] > aNumber:

			log.Println("the number", (*numbers)[currentIndex], ">", aNumber)

			// число в нижней части куска среза, сжимаем область поиска справа
			end = currentIndex

		case (*numbers)[currentIndex] < aNumber:

			log.Println("the number", (*numbers)[currentIndex], "<", aNumber)

			// число в верхней части куска среза, сжимаем область поиска слева
			start = currentIndex + 1

			log.Println("returned with index", foundInIndex)
		default:

			log.Println("the number", (*numbers)[currentIndex], "=", aNumber, "!")

			// число найдено
			foundInIndex = currentIndex

			return
		}
	}
}
