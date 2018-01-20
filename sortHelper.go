package sortHelper

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/alex/ds/sort"
)

/**
rangL<=rangR
**/
func GeneratorRandom(n int, rangL int, rangR int) []int {
	arr := make([]int, 0, n)
	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		arr = append(arr, rand.Int()%(rangR-rangL+1)+rangL)
	}
	return arr
}

func TestSort(sortName string, fun func(arr sort.Sortable), data sort.Sortable) {
	start := time.Now()
	fun(data)
	fmt.Printf("%s:%fs\n", sortName, float64(time.Since(start).Seconds()))
	fmt.Println(IsSorted(data))
}

func IsSorted(arr sort.Sortable) bool {
	type IsSorted interface {
		More(i, j int) bool
	}
	sArr := arr.(IsSorted)
	for i := 0; i < arr.Len()-1; i++ {
		if sArr.More(i, i+1) {
			return false
		}
	}
	return true

}

func CopyIntArray(datas []int) []int {
	arr := make([]int, len(datas))
	// arr = append(arr, datas...)
	copy(arr, datas)
	return arr
}

func GeneratorNearlyOrderlyRandom(n int, swapTimes int) []int {
	arr := make([]int, 0, n)
	for i := 0; i < n; i++ {
		arr = append(arr, i)
	}
	rand.Seed(time.Now().Unix())
	for i := 0; i < swapTimes; i++ {
		k := rand.Int() % n
		v := rand.Int() % n
		arr[k], arr[v] = arr[v], arr[k]
	}
	return arr
}
