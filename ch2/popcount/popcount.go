package popcount

var pc [256]byte

func popCount1(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func popCount2(x uint64) int {
	var (
		num byte
		i   uint
	)
	for i = 0; i < 8; i++ {
		num += pc[byte(x>>(i*8))]
	}
	return int(num)
}

func popCount3(x uint64) int {
	var num int
	for i := byte(0); i < 64; i++ {
		if 1&(x>>i) != 0 {
			num++
		}
	}
	return num
}

func popCount4(x uint64) int {
	var num int
	for x != 0 {
		x &= (x - 1)
		num++
	}
	return num
}

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}
