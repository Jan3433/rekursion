package lists

// Liefert eine Liste mit allen Elementen aus list, die kleiner oder gleich key sind.
// Verwenden Sie Rekursion und benutzen Sie NICHT die len-Funktion.
// Sie können die Hilfsfunktion Empty aus empty.go verwenden.
func FilterLess(list []int, key int) []int {

	n := []int{}

	if Empty(list) {
		return n
	}

	first := list[0]
	rest := list[1:]
	if list[0] <= key {
		return append([]int{first}, FilterLess(rest, key)...)
	}

	return FilterLess(list[1:], key)
}

// Liefert eine Liste mit allen Elementen aus list, die echt größer als key sind.
// Verwenden Sie Rekursion und benutzen Sie NICHT die len-Funktion.
// Sie können die Hilfsfunktion Empty aus empty.go verwenden.
func FilterGreater(list []int, key int) []int {
	// Gehen Sie analog zu FilterLess vor.

	// TODO
	return []int{}
}
