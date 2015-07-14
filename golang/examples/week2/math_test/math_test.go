package math

import "testing"

func TestAverage(t *testing.T) {
  var v float64
  v = Average([]float64{1,2})
  if v != 1.5 {
    t.Error("Expected 1.5, got ", v)
  }
}

func TestSum(t *testing.T) {
  var v = Sum(4, 6)
  if v != 10 {
    t.Error("Expected 10, got", v)
  }
}

func BenchmarkSum(b *testing.B) {
  // run the Sum function b.N times
  for n := 0; n < b.N; n++ {
    Sum(4, 6)
  }
}
