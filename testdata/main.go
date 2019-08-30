package p

var i = 0

var I = 1

var c = "constant un-exported"

const C = "constant exported"

type t struct{}

type T struct{}

func main() {
}

func unexport(s string) {
}

func Export(s, s2 string, data string, list ...int) (d string, e, e2 error) {
	return "", nil, nil
}

// 测试
// fasdfasd
func ExportWithComment(s string) {
	// TODO after
}
