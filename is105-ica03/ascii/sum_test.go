package sum

import "testing"

// Check https://golang.org/ref/spec#Numeric_types and stress the limits!
var sum_tests_int8 = []struct {
	n1       int8
	n2       int8
	expected int8
}{
	{1, 2, 3},
	{4, 5, 9},
	{120, 1, 121},
}

func TestSumInt8(t *testing.T) {
	for _, v := range sum_tests_int8 {
		if val := SumInt8(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%d, %d) returned %d, expected %d", v.n1, v.n2, val, v.expected)
		}
	}
}


//uint32
var sum_tests_uint32 = []struct {
	n1				uint32
	n2 				uint32
	expected	uint32
}{
	{1, 2, 3},
	{12000, 10000, 22000},
	{4294967294, 2, 4294967295},
}

func TestSumUint32(t *testing.T) {
	for _, v := range sum_tests_uint32 {
		if val := SumUint32(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%d, %d) returned %d, expected %d", v.n1, v.n2, val, v.expected)
		}
	}
}

//int32
var sum_tests_int32 = []struct {
	n1				int32
	n2 				int32
	expected	int32
}{
	{1, 2, 3},
	{12000, 10000, 22000},
	{2147483646, 2, 2147483647},
}

func TestSumInt32(t *testing.T) {
	for _, v := range sum_tests_int32 {
		if val := SumInt32(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%d, %d) returned %d, expected %d", v.n1, v.n2, val, v.expected)
		}
	}
}
//int64
var sum_tests_int64 = []struct {
	n1				int64
	n2 				int64
	expected	int64
}{
	{1, 2, 3},
	{12000, 10000, 22000},
	{9223372036854775807, 1, 9223372036854775807},
}

func TestSumInt64(t *testing.T) {
	for _, v := range sum_tests_int64 {
		if val := SumInt64(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%d, %d) returned %d, expected %d", v.n1, v.n2, val, v.expected)
		}
	}
}

//float64
var sum_tests_float64 = []struct {
	n1				float64
	n2 				float64
	expected	float64
}{
	{-1022, -5, -1027},
	{1023, 1, 1024},
	{1.797693134862315708145274237317043567981e+307, 1.797693134862315708145274237317043567981e+308, 1.797693134862315708145274237317043567981e+307},
}

func TestSumFloat64(t *testing.T) {
	for _, v := range sum_tests_float64 {
		if val := SumFloat64(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%d, %d) returned %d, expected %d", v.n1, v.n2, val, v.expected)
		}
	}
}
