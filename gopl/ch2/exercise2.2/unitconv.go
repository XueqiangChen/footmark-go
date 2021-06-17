package exercise2_2

import "fmt"

/**
练习 2.2：
写一个通用的单位转换程序，
(1)用类似cf程序的方式从命令行读取参数，如果缺省的话则是从标准输入读取参数，
(2)做类似Celsius和Fahrenheit的单位转换，长度单位可以对应英尺和米，重量单位可以对应磅和公斤等。

注：
1英尺=0.3048米
1米=3.2808398950131英尺

1磅=0.454千克
1千克=2.204623磅
*/

type Foot float64  //英尺
type Meter float64 //米

type Pound float64    //磅
type Kilogram float64 //千克

func (f Foot) String() string {
	return fmt.Sprintf("%gft", f)
}

func (m Meter) String() string {
	return fmt.Sprintf("%gm", m)
}

func (p Pound) String() string {
	return fmt.Sprintf("%glb", p)
}

func (k Kilogram) String() string {
	return fmt.Sprintf("%gkg", k)
}

