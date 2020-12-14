package main

func main() {
	purchase := 3333_33
	percent := 1
	limit := 100

	const oneHundredPercent = 100_00

	bonus := (purchase / oneHundredPercent) * percent
	if bonus > limit {
		bonus = limit
	}
	println(bonus) // должно быть 33
}
