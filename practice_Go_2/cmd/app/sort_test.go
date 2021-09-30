package main

import (
	//"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func TestSort(t *testing.T) {
	sampleSize := 10
	sortedArray := make([]int, sampleSize)
	randomArray := make([]int, sampleSize)

	for i := 0; i < sampleSize; i++ {
		sortedArray[i] = i
		randomArray[i] = i
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(randomArray), func(i, j int) { randomArray[i], randomArray[j] = randomArray[j], randomArray[i] })

	result := SortArray(randomArray)

	assert.Equal(t, sortedArray, result)
}

/*fmt.Println("Random")
	fmt.Println(randomArray)

	fmt.Println("Sorted")
	fmt.Println(sortedArray)

	fmt.Println("Result")
	fmt.Println(result)

}

*/
