// Package popcount 用于返回一个数字中含二进制1bit的个数
package popcount

// pc[i] is the population count of i.
var pc [256]byte

// 它使用init初始化函数来生成辅助表格pc，pc表格用于处理每个8bit宽度的数字含二进制的1bit的bit个数，
// 这样的话在处理64bit宽度的数字时就没有必要循环64次，只需要8次查表就可以了。
func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}