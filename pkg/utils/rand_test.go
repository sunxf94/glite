package utils

import "testing"

func TestRandA2B(t *testing.T) {
	for i := 0; i < 100; i++ {
		n := RandA2B(1, 100)
		t.Log(i, n)
	}
}

func TestRandPercent(t *testing.T) {
	target, total := 0, 20
	for i := 0; i < total; i++ {
		res := RandPercent(95)
		//t.Log(i, res)

		if res {
			target++
		}
	}
	t.Log(target, total)
}
