package mapdemo

import "testing"

func TestEx9(t *testing.T) {
	Ex9()
}

func BenchmarkEx9(b *testing.B) {
	for i:=0; i<b.N; i++{
		Ex9()
	}
}
func BenchmarkEx10(b *testing.B) {
	for i:=0; i<b.N; i++{
		Ex10()
	}
}

