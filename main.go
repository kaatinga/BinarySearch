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
		numbers[i] = numbers[i-1] + 1 + uint16(rand.Int31n(limit/sliceSize))
	}

	var aNumber uint16 = numbers[rand.Int31n(sliceSize)]

	log.Println("Chosen number is", aNumber)
	log.Println("Numbers:", numbers)

	iterationNumber, foundInIndex := search(0, numbers, aNumber)
	check(numbers, foundInIndex, aNumber, iterationNumber)

	iterationNumber, foundInIndex = search(0, numbers, 5)
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

func search(iterationToInput int32, numbers []uint16, aNumber uint16) (i int32, foundInIndex uint16) {

	i = iterationToInput + 1

	var whereToCut = uint16(len(numbers) / 2)

	switch {
	case len(numbers) == 1 && numbers[whereToCut] != aNumber:

		// это случается только в одном случае, если число не найдено
		return
	case numbers[whereToCut] == aNumber:

		log.Println("the number", numbers[whereToCut], "=", aNumber, "!")

		// число найдено
		foundInIndex = whereToCut

		return

	case numbers[whereToCut] > aNumber:

		log.Println("the number", numbers[whereToCut], ">", aNumber)

		// число в нижней части куска среза, отправляем на анализ повторно
		i, foundInIndex = search(i, numbers[:whereToCut], aNumber)

		log.Println("returned with index", foundInIndex)

	case numbers[whereToCut] < aNumber:

		log.Println("the number", numbers[whereToCut], "<", aNumber)

		// число в верхней части куска среза, отправляем на анализ повторно
		i, foundInIndex = search(i, numbers[whereToCut+1:], aNumber)

		// чтобы index был верный для исходного среза, нужно его считать
		// правильно (увеличиваем на длинну отрезанной слева части):
		foundInIndex = foundInIndex + (whereToCut + 1)

		log.Println("returned with index", foundInIndex)
	}

	return
}
