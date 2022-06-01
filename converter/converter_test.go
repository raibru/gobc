package converter

import (
	"fmt"
	"testing"
)

func Test_NewConverterType_CreateHexType_Successful(t *testing.T) {
	//t.Fatal("Check Failure")
	//Given
	want := "*converter.HexType"

	// When
	o := NewConverterType[HexType](1)
	o.Init()
	got := fmt.Sprintf("%T", o)

	// Then
	if got != want {
		t.Errorf("Want '%s' but got '%s'", want, got)
	}
	if !o.Base.IsInit {
		t.Errorf("Want initialized '%s' but got IsInit=%v", got, o.Base.IsInit)
	}
}

func Test_NewConverterType_CreateDecType_Successful(t *testing.T) {
	//t.Fatal("Check Failure")
	//Given
	want := "*converter.DecType"

	// When
	o := NewConverterType[DecType](1)
	o.Init()
	got := fmt.Sprintf("%T", o)

	// Then
	if got != want {
		t.Errorf("Want '%s' but got '%s'", want, got)
	}
	if !o.Base.IsInit {
		t.Errorf("Want initialized '%s' but got IsInit=%v", got, o.Base.IsInit)
	}
}

func Test_NewConverterType_CreateOctType_Successful(t *testing.T) {
	//t.Fatal("Check Failure")
	//Given
	want := "*converter.OctType"

	// When
	o := NewConverterType[OctType](1)
	o.Init()
	got := fmt.Sprintf("%T", o)

	// Then
	if got != want {
		t.Errorf("Want '%s' but got '%s'", want, got)
	}
	if !o.Base.IsInit {
		t.Errorf("Want initialized '%s' but got IsInit=%v", got, o.Base.IsInit)
	}
}

func Test_NewConverterType_CreateBinType_Successful(t *testing.T) {
	//t.Fatal("Check Failure")
	//Given
	want := "*converter.BinType"

	// When
	o := NewConverterType[BinType](1)
	o.Init()
	got := fmt.Sprintf("%T", o)

	// Then
	if got != want {
		t.Errorf("Want '%s' but got '%s'", want, got)
	}
	if !o.Base.IsInit {
		t.Errorf("Want initialized '%s' but got IsInit=%v", got, o.Base.IsInit)
	}
}

func Test_ParseBaseValue_ByHexParameter_Successful(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	tables := []struct {
		input string
		want  uint64
	}{
		{"0", 0},
		{"1", 1},
		{"2", 2},
		{"a", 10},
		{"ff", 255},
		{"affe", 45054},
		{"deadbeaf", 3735928495},
		{"ffffffffffffffff", 18446744073709551615},
	}

	for _, table := range tables {
		p := Parameters{
			Hex: table.input,
		}
		// When
		v, err := ParseBaseValue(&p)

		// Then
		if err != nil {
			t.Fatalf("Want no error but got '%s'", err.Error())
		}
		if v != table.want {
			t.Errorf("Want %d but got %d", table.want, v)
		}
	}
}

func Test_ParseBaseValue_ByDecParameter_Successful(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	tables := []struct {
		input string
		want  uint64
	}{
		{"0", 0},
		{"1", 1},
		{"2", 2},
		{"10", 10},
		{"255", 255},
		{"45054", 45054},
		{"3735928495", 3735928495},
		{"18446744073709551615", 18446744073709551615},
	}

	for _, table := range tables {
		p := Parameters{
			Dec: table.input,
		}
		// When
		v, err := ParseBaseValue(&p)

		// Then
		if err != nil {
			t.Fatalf("Want no error but got '%s'", err.Error())
		}
		if v != table.want {
			t.Errorf("Want %d but got %d", table.want, v)
		}
	}
}

func Test_ParseBaseValue_ByOctParameter_Successful(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	tables := []struct {
		input string
		want  uint64
	}{
		{"0", 0},
		{"1", 1},
		{"2", 2},
		{"7", 7},
		{"10", 8},
		{"12", 10},
		{"377", 255},
		{"127776", 45054},
		{"1777777777777777777777", 18446744073709551615},
	}

	for _, table := range tables {
		p := Parameters{
			Oct: table.input,
		}
		// When
		v, err := ParseBaseValue(&p)

		// Then
		if err != nil {
			t.Fatalf("Want no error but got '%s'", err.Error())
		}
		if v != table.want {
			t.Errorf("Want %d but got %d", table.want, v)
		}
	}
}

func Test_ParseBaseValue_ByBinParameter_Successful(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	tables := []struct {
		input string
		want  uint64
	}{
		{"0", 0},
		{"1", 1},
		{"10", 2},
		{"11", 3},
		{"1010", 10},
		{"11111111", 255},
		{"1010111111111110", 45054},
		{"1111111111111111111111111111111111111111111111111111111111111111", 18446744073709551615},
	}

	for _, table := range tables {
		p := Parameters{
			Bin: table.input,
		}
		// When
		v, err := ParseBaseValue(&p)

		// Then
		if err != nil {
			t.Fatalf("Want no error but got '%s'", err.Error())
		}
		if v != table.want {
			t.Errorf("Want %d but got %d", table.want, v)
		}
	}
}

func Test_ParseBaseValue_WithMultiBaseParameter_HasErrorOnlyOneParameter(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	wantError := "only one parameter shall be used"
	tables := []struct {
		inputHex string
		inputDec string
		inputOct string
		inputBin string
	}{
		{"1", "1", "", ""},
		{"1", "", "1", ""},
		{"1", "", "", "1"},
		{"", "1", "1", ""},
		{"", "1", "", "1"},
		{"", "", "1", "1"},
		{"1", "1", "1", ""},
		{"1", "1", "", "1"},
		{"1", "", "1", "1"},
		{"", "1", "1", "1"},
		{"1", "1", "1", "1"},
	}

	for _, table := range tables {
		p := Parameters{
			Hex: table.inputHex,
			Dec: table.inputDec,
			Oct: table.inputOct,
			Bin: table.inputBin,
		}
		// When
		_, err := ParseBaseValue(&p)

		// Then
		if err == nil {
			t.Fatalf("When set hex='%s' dec='%s' oct='%s' bin='%s' then want error but got NONE", p.Hex, p.Dec, p.Oct, p.Bin)
		}
		if err.Error() != wantError {
			t.Errorf("Want error '%s' but got '%s'", wantError, err.Error())
		}
	}
}

func Test_ParseBaseValue_WithNoBaseParameter_HasErrorNeedOneParameter(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	wantError := "need one parameter to convert into bases"
	p := Parameters{Hex: "", Dec: "", Oct: "", Bin: ""}

	// When
	_, err := ParseBaseValue(&p)

	// Then
	if err == nil {
		t.Fatalf("When set hex='%s' dec='%s' oct='%s' bin='%s' then want error but got NONE", p.Hex, p.Dec, p.Oct, p.Bin)
	}
	if err.Error() != wantError {
		t.Errorf("Want error '%s' but got '%s'", wantError, err.Error())
	}
}

func Test_ParseBaseValue_WithWrongBaseParameter_HasErrorInvalidSyntax(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	wantError := "strconv.ParseUint: parsing \"z\": invalid syntax"
	tables := []struct {
		inputHex string
		inputDec string
		inputOct string
		inputBin string
	}{
		{"z", "", "", ""},
		{"", "z", "", ""},
		{"", "", "z", ""},
		{"", "", "", "z"},
	}

	for _, table := range tables {
		p := Parameters{
			Hex: table.inputHex,
			Dec: table.inputDec,
			Oct: table.inputOct,
			Bin: table.inputBin,
		}
		// When
		_, err := ParseBaseValue(&p)

		// Then
		if err == nil {
			t.Fatalf("When set hex='%s' dec='%s' oct='%s' bin='%s' then want error but got NONE", p.Hex, p.Dec, p.Oct, p.Bin)
		}
		if err.Error() != wantError {
			t.Errorf("Want error '%s' but got '%s'", wantError, err.Error())
		}
	}
}

func Test_ParseLeadingZeros_Successful(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	tables := []struct {
		input string
		key   string
		want  string
	}{
		{"x:0", "x", "0"},
		{"x:1", "x", "1"},
		{"x:2", "x", "2"},
		{"d:0", "d", "0"},
		{"d:1", "d", "1"},
		{"d:2", "d", "2"},
		{"o:0", "o", "0"},
		{"o:1", "o", "1"},
		{"o:2", "o", "2"},
		{"b:0", "b", "0"},
		{"b:1", "b", "1"},
		{"b:2", "b", "2"},
		{"x:4,d:5,o:6,b:7", "x", "4"},
		{"x:4,d:5,o:6,b:7", "d", "5"},
		{"x:4,d:5,o:6,b:7", "o", "6"},
		{"x:4,d:5,o:6,b:7", "b", "7"},
	}

	for _, table := range tables {

		// When
		got, err := ParseLeadingZeros(table.input, table.key)

		// Then
		if err != nil {
			t.Fatalf("Want no error but got '%s'", err.Error())
		}
		if got != table.want {
			t.Errorf("Want %s but got %s", table.want, got)
		}
	}
}

func Test_ParseLeadingZeros_HasError(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	tables := []struct {
		input string
		key   string
	}{
		{"y:0", "x"},
		{"x:y", "x"},
		{"d:y", "d"},
		{"o:y", "o"},
		{"b:y", "b"},
		{"x:1,d:y", "x"},
		{"d:1,o:y", "d"},
		{"o:1,b:y", "o"},
	}

	for _, table := range tables {

		// When
		_, err := ParseLeadingZeros(table.input, table.key)

		// Then
		if err == nil {
			t.Fatalf("Want error when set lead-zeros string '%s' but got NONE", table.input)
		}
	}
}

func Test_PrintBases_Successful(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	want := "NaN NaN NaN NaN\n"

	var bases = []Converter{
		NewConverterType[HexType](0),
		NewConverterType[DecType](0),
		NewConverterType[OctType](0),
		NewConverterType[BinType](0),
	}
	for _, b := range bases {
		b.Init()
	}

	writer := &mockWriter{}

	// When
	PrintBases(bases, writer)
	got := writer.String()

	// Then
	if got != want {
		t.Errorf("Want '%s' but got '%s'", want, got)
	}
}

//func Test_ApplyPipeInput_FailureByWrongParam(t *testing.T) {
//	//t.Fatal("Check Failure")
//	// Given
//	var co = []Convertable{}
//	bc := BaseContext{Dec: "#"}
//	vexp := "1"
//
//	co = append(co, &DecConverter{})
//	s := bufio.NewScanner(strings.NewReader(vexp))
//
//	// When
//	err := ApplyPipeInput(s, &co, &bc)
//
//	// Then
//	if err == nil {
//		t.Error("Failure: apply pipe input shall have an error but returns nil")
//	}
//}

type mockWriter struct {
	content []byte
}

func (w *mockWriter) Write(buf []byte) (int, error) {
	w.content = append(w.content, buf...)
	return len(buf), nil
}

func (w *mockWriter) String() string {
	return string(w.content)
}
