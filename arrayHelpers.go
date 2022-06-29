package main

import (
	"math"
	"sort"
	"strconv"
)

func sliceAtoi(sa []string) ([]int, error) {
	si := make([]int, 0, len(sa))
	for _, a := range sa {
		i, err := strconv.Atoi(a)
		if err != nil {
			return si, err
		}
		si = append(si, i)
	}
	return si, nil
}

func Min(numbers []int) int {
	var min int = numbers[0]
	for _, value := range numbers {
		if min > value {
			min = value
		}
	}

	return min
}

func Max(numbers []int) int {
	var max int = numbers[0]
	for _, value := range numbers {
		if max < value {
			max = value
		}
	}
	return max
}

func Avg(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total = total + number
	}

	return total / len(numbers)
}

func Median(numbers []int) int {
	numbersCopy := make([]int, len(numbers))
	copy(numbersCopy, numbers)

	sort.Ints(numbersCopy)

	var median int
	length := len(numbersCopy)
	if length == 0 {
		return 0
	} else if length%2 == 0 {
		median = (numbersCopy[length/2-1] + numbersCopy[length/2]) / 2
	} else {
		median = numbersCopy[length/2]
	}

	return median
}

func PercentileNearestRank(input []int, percent float64) int {
	numbersCopy := make([]int, len(input))
	copy(numbersCopy, input)
	sort.Ints(numbersCopy)

	inputLength := len(input)
	if percent == 100.0 {
		return numbersCopy[inputLength-1]
	}

	ordinalRanking := int(math.Ceil(float64(inputLength) * percent / 100))

	if ordinalRanking == 0 {
		return numbersCopy[0]
	}
	return numbersCopy[ordinalRanking-1]

}
