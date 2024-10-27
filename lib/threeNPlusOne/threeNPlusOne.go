package threeNPlusOneLib

/* program inspiration:
 * https://youtu.be/094y1Z2wpJg?si=J5aw_yDn-MmA9nJY
 * https://en.wikipedia.org/wiki/Collatz_conjecture
 */

func GenThreeNPlusOne(num int64) []int64 {
	var out []int64

	if num == 1 {
		return []int64{num}
	}

	for i := num; i != 1; i = GetNextThreeNplusOne(i) {
		out = append(out, i)
	}

	out = append(out, 1)

	return out
}

func GetNextThreeNplusOne(num int64) int64 {
	var out int64 = 0

	even := num%2 == 0
	if even {
		out = num / 2
	} else {
		out = (num * 3) + 1
	}

	return out
}
