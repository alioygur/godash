package godash

import (
	"fmt"
	"testing"
)

func TestToInt(t *testing.T) {
	tests := []string{"1000", "-123", "abcdef", "100000000000000000000000000000000000000000000"}
	expected := []int64{1000, -123, 0, 0}
	for i := 0; i < len(tests); i++ {
		result, _ := ToInt(tests[i])
		if result != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", result)
			t.FailNow()
		}
	}
}

func TestToBoolean(t *testing.T) {
	tests := []string{"true", "1", "True", "false", "0", "abcdef"}
	expected := []bool{true, true, true, false, false, false}
	for i := 0; i < len(tests); i++ {
		res, _ := ToBoolean(tests[i])
		if res != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", res)
			t.FailNow()
		}
	}
}

func ExampleToBoolean() {
	fmt.Println(ToBoolean("True"))
	fmt.Println(ToBoolean("true"))
	fmt.Println(ToBoolean("1"))
	fmt.Println(ToBoolean("False"))
	fmt.Println(ToBoolean("false"))
	fmt.Println(ToBoolean("0"))
	// Output:
	// true <nil>
	// true <nil>
	// true <nil>
	// false <nil>
	// false <nil>
	// false <nil>
}

func toString(t *testing.T, test interface{}, expected string) {
	res := ToString(test)
	if res != expected {
		t.Log("Case ToString: expected ", expected, " when result is ", res)
		t.FailNow()
	}
}

func TestToString(t *testing.T) {
	toString(t, "str123", "str123")
	toString(t, 123, "123")
	toString(t, 12.3, "12.3")
	toString(t, true, "true")
	toString(t, 1.5+10i, "(1.5+10i)")
	// Sprintf function not guarantee that maps with equal keys always will be equal in string  representation
	//toString(t, struct{ Keys map[int]int }{Keys: map[int]int{1: 2, 3: 4}}, "{map[1:2 3:4]}")
}

func TestToFloat(t *testing.T) {
	tests := []string{"", "123", "-.01", "10.", "string", "1.23e3", ".23e10"}
	expected := []float64{0, 123, -0.01, 10.0, 0, 1230, 0.23e10}
	for i := 0; i < len(tests); i++ {
		res, _ := ToFloat(tests[i])
		if res != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", res)
			t.FailNow()
		}
	}
}

func TestToCamelCase(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected string
	}{
		{"a_b_c", "ABC"},
		{"my_func", "MyFunc"},
		{"1ab_cd", "1abCd"},
		{"toUTFCamelCase", "ToUTFCamelCase"},
		{"camel case", "CamelCase"},
		{"  camel case  ", "CamelCase"},
		{"!!!camel case====", "CamelCase"},
		{"camel-case", "CamelCase"},
		{"camel_case", "CamelCase"},
	}
	for _, test := range tests {
		actual := ToCamelCase(test.param)
		if actual != test.expected {
			t.Errorf("Expected ToCamelCase(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func ExampleToCamelCase() {
	fmt.Println(ToCamelCase("camel case"))
	fmt.Println(ToCamelCase("  camel case  "))
	fmt.Println(ToCamelCase("!!!camel case===="))
	fmt.Println(ToCamelCase("camel-case"))
	fmt.Println(ToCamelCase("camel_case"))
	// Output:
	// CamelCase
	// CamelCase
	// CamelCase
	// CamelCase
	// CamelCase
}

func TestToSnakeCase(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected string
	}{
		{"a", "a"},
		{"snake", "snake"},
		{"A", "a"},
		{"ID", "id"},
		{"MOTD", "motd"},
		{"Snake", "snake"},
		{"SnakeTest", "snake_test"},
		{"SnakeID", "snake_id"},
		{"LinuxMOTD", "linux_motd"},
		{"snake case", "snake_case"},
		{"   ! snake----- [ case", "snake_case"},
		{"publicF", "public_f"},
		{"RESTful", "restful"},
		{"APIResponse", "apiresponse"},
		{"simpleXML", "simple_xml"},
	}
	for _, test := range tests {
		actual := ToSnakeCase(test.param)
		if actual != test.expected {
			t.Errorf("Expected ToSnakeCase(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func ExampleToSnakeCase() {
	fmt.Println(ToSnakeCase("SnakeCase"))
	fmt.Println(ToSnakeCase("snake case"))
	fmt.Println(ToSnakeCase("  snake case  "))
	fmt.Println(ToSnakeCase("!!!snake case===="))
	fmt.Println(ToSnakeCase("snake-case"))
	fmt.Println(ToSnakeCase("snake_case"))
	// Output:
	// snake_case
	// snake_case
	// snake_case
	// snake_case
	// snake_case
	// snake_case
}
