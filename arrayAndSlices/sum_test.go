package main

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestSum(t *testing.T) {

	t.Run("Collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})

	t.Run("Collection of any numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTrails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sum of some slices", func(t *testing.T) {
		got := SumAllTrails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTrails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)
	})

}

func TestTimeQuery(t *testing.T) {
	currentUTC := time.Now().UTC()
	startTime := time.Date(currentUTC.Year(), currentUTC.Month(), currentUTC.Day(), 0, 0, 0, 0, currentUTC.UTC().Location()).AddDate(0, 0, -60)
	endTime := startTime.AddDate(0, 0, 5)
	current := time.Date(currentUTC.Year(), currentUTC.Month(), currentUTC.Day(), 0, 0, 0, 0, currentUTC.UTC().Location())

	for startTime.Before(current) {
		query := fmt.Sprintf(`select from body where time between  '%s' and '%s'`, startTime.Format("2006-01-02 15:04:05"), endTime.Format("2006-01-02 15:04:05"))
		fmt.Println(query)
		startTime = endTime
		endTime = startTime.AddDate(0, 0, 5)
	}

}
