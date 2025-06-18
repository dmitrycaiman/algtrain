package main

const listSize = 50

var fibonacciNumbersList = []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, 17711, 28657, 46368, 75025, 121393, 196418, 317811, 514229, 832040, 1346269, 2178309, 3524578, 5702887, 9227465, 14930352, 24157817, 39088169, 63245986, 102334155, 165580141, 267914296, 433494437, 701408733, 1134903170, 1836311903, 2971215073, 4807526976, 7778742049}

var pellNumbersList = []int{0, 1, 2, 5, 12, 29, 70, 169, 408, 985, 2378, 5741, 13860, 33461, 80782, 195025, 470832, 1136689, 2744210, 6625109, 15994428, 38613965, 93222358, 225058681, 543339720, 1311738121, 3166815962, 7645370045, 18457556052, 44560482149, 107578520350, 259717522849, 627013566048, 1513744654945, 3654502875938, 8822750406821, 21300003689580, 51422757785981, 124145519261542, 299713796309065, 723573111879672, 1746860020068409, 4217293152016490, 10181446324101389, 24580185800219268, 59341817924539925, 143263821649299118, 345869461223138161, 835002744095575440, 2015874949414289041}

var padovanNumbersList = []int{1, 0, 0, 1, 0, 1, 1, 1, 2, 2, 3, 4, 5, 7, 9, 12, 16, 21, 28, 37, 49, 65, 86, 114, 151, 200, 265, 351, 465, 616, 816, 1081, 1432, 1897, 2513, 3329, 4410, 5842, 7739, 10252, 13581, 17991, 23833, 31572, 41824, 55405, 73396, 97229, 128801, 170625}

var jacobsthalNumbersList = []int{0, 1, 1, 3, 5, 11, 21, 43, 85, 171, 341, 683, 1365, 2731, 5461, 10923, 21845, 43691, 87381, 174763, 349525, 699051, 1398101, 2796203, 5592405, 11184811, 22369621, 44739243, 89478485, 178956971, 357913941, 715827883, 1431655765, 2863311531, 5726623061, 11453246123, 22906492245, 45812984491, 91625968981, 183251937963, 366503875925, 733007751851, 1466015503701, 2932031007403, 5864062014805, 11728124029611, 23456248059221, 46912496118443, 93824992236885, 187649984473771}

var tribonacciNumbersList = []int{0, 0, 1, 1, 2, 4, 7, 13, 24, 44, 81, 149, 274, 504, 927, 1705, 3136, 5768, 10609, 19513, 35890, 66012, 121415, 223317, 410744, 755476, 1389537, 2555757, 4700770, 8646064, 15902591, 29249425, 53798080, 98950096, 181997601, 334745777, 615693474, 1132436852, 2082876103, 3831006429, 7046319384, 12960201916, 23837527729, 43844049029, 80641778674, 148323355432, 272809183135, 501774317241, 922906855808, 1697490356184}

var tetranacciNumbersList = []int{0, 0, 0, 1, 1, 2, 4, 8, 15, 29, 56, 108, 208, 401, 773, 1490, 2872, 5536, 10671, 20569, 39648, 76424, 147312, 283953, 547337, 1055026, 2033628, 3919944, 7555935, 14564533, 28074040, 54114452, 104308960, 201061985, 387559437, 747044834, 1439975216, 2775641472, 5350220959, 10312882481, 19878720128, 38317465040, 73859288608, 142368356257, 274423830033, 528968939938, 1019620414836, 1965381541064, 3788394725871, 7302365621709}

// fibonacciNumbers represents Fibonacci numbers: F(n) = F(n-1) + F(n-2) with F(0) = 0 and F(1) = 1.
func fibonacciNumbers(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	n1, n2 := 0, 1
	for i := 2; i <= n; i++ {
		n1, n2 = n2, n1+n2
	}
	return n2
}

// fibonacciNumbersMemo рассчитывает последовательность Фибоначчи рекурсивно с мемоизацией.
func fibonacciNumbersMemo(n int) int {
	return fibonacciNumbersMemoResolver(n, make([]int, n+1))
}

func fibonacciNumbersMemoResolver(n int, cache []int) (result int) {
	if n == 0 || n == 1 {
		return n
	}
	if cache[n] != 0 {
		return cache[n]
	}
	defer func() { cache[n] = result }()
	return fibonacciNumbersMemoResolver(n-1, cache) + fibonacciNumbersMemoResolver(n-2, cache)
}

// padovanNumbers represents Padovan sequence: a(n) = a(n-2) + a(n-3) with a(0) = 1, a(1) = a(2) = 0.
func padovanNumbers(n int) int {
	switch n {
	case 0:
		return 1
	case 1, 2:
		return 0
	}

	n0, n1, n2, n3 := 0, 0, 0, 1
	for i := 3; i <= n; i++ {
		n0 = n2 + n3
		n1, n2, n3 = n0, n1, n2
	}

	return n0
}

// jacobsthalNumbers represents Jacobsthal sequence (or Jacobsthal numbers): a(n) = a(n-1) + 2*a(n-2), with a(0) = 0, a(1) = 1; also a(n) = nearest integer to 2^n/3.
func jacobsthalNumbers(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	n0, n1, n2 := 0, 1, 0
	for i := 2; i <= n; i++ {
		n0 = n1 + 2*n2
		n1, n2 = n0, n1
	}

	return n1
}

// pellNumbers represents Pell numbers: a(0) = 0, a(1) = 1; for n > 1, a(n) = 2*a(n-1) + a(n-2).
func pellNumbers(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	n0, n1, n2 := 0, 1, 0
	for i := 2; i <= n; i++ {
		n0 = 2*n1 + n2
		n1, n2 = n0, n1
	}

	return n0
}

// tribonacciNumbers represents Tribonacci numbers: a(n) = a(n-1) + a(n-2) + a(n-3) for n >= 3 with a(0) = a(1) = 0 and a(2) = 1.
func tribonacciNumbers(n int) int {
	switch n {
	case 0, 1:
		return 0
	case 2:
		return 1
	}

	n0, n1, n2, n3 := 0, 1, 0, 0
	for i := 3; i <= n; i++ {
		n0 = n1 + n2 + n3
		n1, n2, n3 = n0, n1, n2
	}

	return n0
}

// tetranacciNumbers represents Tetranacci numbers: a(n) = a(n-1) + a(n-2) + a(n-3) + a(n-4) for n >= 4 with a(0) = a(1) = a(2) = 0 and a(3) = 1.
func tetranacciNumbers(n int) int {
	switch n {
	case 0, 1, 2:
		return 0
	case 3:
		return 1
	}

	n0, n1, n2, n3, n4 := 0, 1, 0, 0, 0
	for i := 4; i <= n; i++ {
		n0 = n1 + n2 + n3 + n4
		n1, n2, n3, n4 = n0, n1, n2, n3
	}

	return n0
}
