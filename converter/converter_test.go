package converter

import (
	"bufio"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func Test_NewConverterType_CreateHexType_Successful(t *testing.T) {
	//t.Fatal("Check Failure")
	//Given
	want := "converter.HexType"

	// When
	o := NewConverterType[HexType]()
	got := fmt.Sprintf("%T", o)

	// Then
	if got != want {
		t.Errorf("Want '%s' but got '%s'", want, got)
	}
}

func Test_NewConverterType_CreateDecType_Successful(t *testing.T) {
	//t.Fatal("Check Failure")
	//Given
	want := "converter.DecType"

	// When
	o := NewConverterType[DecType]()
	got := fmt.Sprintf("%T", o)

	// Then
	if got != want {
		t.Errorf("Want '%s' but got '%s'", want, got)
	}
}

func Test_NewConverterType_CreateOctType_Successful(t *testing.T) {
	//t.Fatal("Check Failure")
	//Given
	want := "converter.OctType"

	// When
	o := NewConverterType[OctType]()
	got := fmt.Sprintf("%T", o)

	// Then
	if got != want {
		t.Errorf("Want '%s' but got '%s'", want, got)
	}
}

func Test_NewConverterType_CreateBinType_Successful(t *testing.T) {
	//t.Fatal("Check Failure")
	//Given
	want := "converter.BinType"

	// When
	o := NewConverterType[BinType]()
	got := fmt.Sprintf("%T", o)

	// Then
	if got != want {
		t.Errorf("Want '%s' but got '%s'", want, got)
	}
}

func Test_ApplyPipeInput_ExpectSuccess(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	var co = []Convertable{}
	bc := BaseContext{Dec: "."}
	vexp := "1"

	co = append(co, &DecConverter{})
	s := bufio.NewScanner(strings.NewReader(vexp))

	// When
	err := ApplyPipeInput(s, &co, &bc)

	// Then
	if err != nil {
		t.Errorf("Failure: apply pipe input for decimal '1' fails by error: %v", err)
	}
}

func Test_ApplyPipeInput_FailureByWrongParam(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	var co = []Convertable{}
	bc := BaseContext{Dec: "#"}
	vexp := "1"

	co = append(co, &DecConverter{})
	s := bufio.NewScanner(strings.NewReader(vexp))

	// When
	err := ApplyPipeInput(s, &co, &bc)

	// Then
	if err == nil {
		t.Error("Failure: apply pipe input shall have an error but returns nil")
	}
}

func Test_CreateConverter_ExpectSuccess(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	var co = []Convertable{}
	bc := BaseContext{Dec: "1"}

	co = append(co, &DecConverter{})

	// When
	conv, err := CreateConverter(&co, &bc)

	// Then
	if err != nil {
		t.Errorf("Failure: create converter for decimal '1' fails by error: %v", err)
	}
	if conv == nil {
		t.Error("Failure: create converter should not return nil")
	}
}

func Test_CreateConverter_ExpectFailure_2_BaseParameters(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	expVal := "1"
	bc := BaseContext{
		Dec: expVal,
		Hex: expVal}

	var co = []Convertable{}
	co = append(co, &DecConverter{})
	co = append(co, &HexConverter{})

	// When
	result, err := CreateConverter(&co, &bc)

	// Then
	if err == nil {
		t.Error("Failure: expect error with 2 base parameters but get none")
	}
	if result != nil {
		t.Errorf("Failure: expect no result with 2 base parameters but get %v", result)
	}
}

func Test_CreateConverter_ExpectFailure_None_BaseParameters(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	var co = []Convertable{}
	bc := BaseContext{}

	// When
	result, err := CreateConverter(&co, &bc)

	// Then
	if err == nil {
		t.Error("Failure: expect error with none base parameters but get none")
	}
	if result != nil {
		t.Errorf("Failure: expect no result with none base parameters but get %v", result)
	}
}

func Test_ParseBaseValue_ExpectSuccess(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	expVal := "1"
	bc := BaseContext{Dec: expVal}

	var co = []Convertable{}
	co = append(co, &DecConverter{})

	conv, err := CreateConverter(&co, &bc)

	if err != nil || conv == nil {
		t.Fatal("Fatal: can not create converter for decimal '1' parameter")
	}

	// When
	result, perr := ParseBaseValue(conv, &bc)

	// Then
	if perr != nil {
		t.Errorf("Failure: parse base value expect '%v' but get error: %v", expVal, perr)
	}
	if expVal != strconv.FormatInt(result, 10) {
		t.Errorf("Failure: parse base value expect '%v' but get '%v'", expVal, result)
	}

}

func Test_ParseBaseValue_ExpectFailure(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	expVal := "z"
	bc := BaseContext{Dec: expVal}

	var co = []Convertable{}
	co = append(co, &DecConverter{})

	conv, err := CreateConverter(&co, &bc)

	if err != nil || conv == nil {
		t.Fatal("Fatal: can not create converter for decimal '1' parameter")
	}

	// When
	_, perr := ParseBaseValue(conv, &bc)

	// Then
	if perr == nil {
		t.Errorf("Failure: parse base value '%v' expect error but get none", expVal)
	}

}

func Test_ApplyBaseContext_ExpectSuccess(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	expVal := "1"
	bc := BaseContext{
		Dec: expVal}

	var co = []Convertable{}
	co = append(co, &DecConverter{})
	co = append(co, &HexConverter{})
	co = append(co, &OctConverter{})
	co = append(co, &BinConverter{})

	conv, err := CreateConverter(&co, &bc)

	if err != nil || conv == nil {
		t.Fatal("Fatal: can not create converter")
	}

	i, perr := ParseBaseValue(conv, &bc)

	if perr != nil {
		t.Fatalf("Fatal: can not parse base value '%v'", expVal)
	}

	// When
	result := ApplyBaseContext(&co, i, &bc)

	// Then
	if reflect.TypeOf(result).Name() != reflect.TypeOf(bc).Name() {
		t.Errorf("Failure: ApplyBaseContext expect a BaseContext as return but get %s", reflect.TypeOf(result).Name())
	}
	if result.Dec == "" {
		t.Errorf("Failure: BaseContext Decimal expect %s  but get empty string", expVal)
	}
	if result.Hex == "" {
		t.Errorf("Failure: BaseContext Hexadecimal expect %s  but get empty string", expVal)
	}
	if result.Oct == "" {
		t.Errorf("Failure: BaseContext Octal expect %s  but get empty string", expVal)
	}
	if result.Bin == "" {
		t.Errorf("Failure: BaseContext Binary expect %s  but get empty string", expVal)
	}
}

func Test_ParseLeadingZerosSingle(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	tables := []struct {
		input         string
		expectedKey   string
		expectedValue string
	}{
		{"x:0", "x", "0"},
		{"x:1", "x", "1"},
		{"d:0", "d", "0"},
		{"d:1", "d", "1"},
		{"o:0", "o", "0"},
		{"o:1", "o", "1"},
		{"b:0", "b", "0"},
		{"b:1", "b", "1"},
	}

	for _, table := range tables {

		// When
		bc := BaseContext{
			Bin: "1",
		}

		var co = []Convertable{}
		co = append(co, &HexConverter{})
		co = append(co, &DecConverter{})
		co = append(co, &OctConverter{})
		co = append(co, &BinConverter{})

		conv, err := CreateConverter(&co, &bc)

		if err != nil || conv == nil {
			t.Fatal("Fatal: can not create converter")
		}

		m, err := ParseLeadingZero(table.input)

		// Then
		if err != nil {
			t.Fatalf("Fatal: can not parse leading zeros for '%s': '%v'", table.input, err)
		}

		output, ok := (*m)[table.expectedKey]
		if !ok {
			t.Fatalf("Fatal: key '%s' not found for: '%s'", table.expectedKey, table.input)
		}

		if table.expectedValue != output {
			t.Errorf("Failure: Leading zero expect '%s' but get '%s'", table.expectedValue, output)
		}
	}
}
