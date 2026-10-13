package main

import "fmt"

func main() {
	xgcd(300, 320)
}

func extendedEuclideanAlgoritm(a int, b int) (int, int, int) {
	if a < 0 {
		d, x, y := extendedEuclideanAlgoritm(-a, b)
		return d, -x, y
	}
	if b < 0 {
		d, x, y := extendedEuclideanAlgoritm(a, -b)
		return d, x, -y
	}

	if a == 0 && b == 0 {
		return 0, 1, 0
	}

	oldR, r := a, b
	oldX, x := 1, 0
	oldY, y := 0, 1

	for r != 0 {
		q := oldR / r
		oldR, r = r, oldR-q*r
		oldX, x = x, oldX-q*x
		oldY, y = y, oldY-q*y
	}

	return oldR, oldX, oldY
}

func xgcd(a int, b int) int {
	d, x, y := extendedEuclideanAlgoritm(a, b)
	fmt.Printf("x = %d, y = %d\n", x, y)
	return d
}
