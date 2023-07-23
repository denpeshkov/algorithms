package combinatorial

// Floyd returns an index of the first element in cycle and the length of the cycle in a sequence of iterated function values.
func FloydCycle(f func(int) int, i int) (ind int, len int) {
	slow, fast := f(i), f(f(i))

	for slow != fast {
		slow = f(slow)
		fast = f(f(fast))
	}

	slow = i
	for slow != fast {
		slow = f(slow)
		fast = f(fast)
		ind++
	}

	len = 1
	fast = f(slow)
	for slow != fast {
		fast = f(fast)
		len++
	}

	return
}
