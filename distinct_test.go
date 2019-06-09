package underscore

import "testing"

func Test_Distinct(t *testing.T) {
	src := []int{1, 2, 1, 4, 1, 3}
	dst := make([]int, 0)
	Chain2(src).Distinct(func(n, _ int) (int, error) {
		return n % 2, nil
	}).Value(&dst)
	if len(dst) != 2 {
		t.Error(dst)
	}
}

func Test_Distinct_SelectorIsNil(t *testing.T) {
	src := []int{1, 2, 1, 4, 1, 3}
	dst := make([]int, 0)
	Chain2(src).Distinct(nil).Value(&dst)
	if len(dst) != 4 {
		t.Error(dst)
	}
}

func Test_DistinctBy(t *testing.T) {
	src := []testModel{
		{ID: 1, Name: "a"},
		{ID: 2, Name: "a"},
		{ID: 3, Name: "a"},
	}
	dst := make([]testModel, 0)
	Chain2(src).DistinctBy("name").Value(&dst)
	if len(dst) != 1 {
		t.Error(dst)
	}
}