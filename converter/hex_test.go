package converter

import (
	"bufio"
	"strconv"
	"strings"
	"testing"
)

func Test_HexApplyPipeInput_ExpectSuccess(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	var bctx BaseContext
	var hc HexConverter
	var vexp string

	// When
	vexp = "1"
	bctx.Hex = "."
	s := bufio.NewScanner(strings.NewReader(vexp))
	eerr := hc.ApplyPipeInput(s, &bctx)

	// Then
	if bctx.Hex != vexp {
		t.Errorf("Failure: ApplyPipeInput set Hexadecimal BaseContext expect %v but get %v", vexp, bctx.Hex)
	}
	if eerr != nil {
		t.Errorf("Failure: ApplyPipeInput returns error: %v", eerr)
	}
}

func Test_HexApplyPipeInput_FailureByWrongParam(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	var bctx BaseContext
	var hc HexConverter
	var vexp string

	// When
	vexp = "1"
	bctx.Hex = "#"
	s := bufio.NewScanner(strings.NewReader(vexp))
	eerr := hc.ApplyPipeInput(s, &bctx)

	// Then
	if bctx.Hex == vexp {
		t.Errorf("Failure: ApplyPipeInput shall not apply pipe input expect %s but get %v", "#", bctx.Hex)
	}
	if eerr == nil {
		t.Error("Failure: ApplyPipeInput shall returns an error but is nil")
	}
}

func Test_HexApplyContext_ExpectSuccess(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	var bctx BaseContext
	var hc HexConverter
	var vexp string
	var v int64

	leadZero := "x:0"

	// When
	m, ferr := ParseLeadingZero(leadZero)
	vexp = "1"
	v, perr := strconv.ParseInt(vexp, 16, 64)
	bctx.Hex = ""
	exp := hc.ApplyContext(v, &bctx, m)

	// Then
	if ferr != nil {
		t.Fatalf("Fatal: ParseLeadingZero (%s) failed: %v", leadZero, ferr)
	}

	if perr != nil {
		t.Fatalf("Fatal: expect no error ParseInt %v to base 16: %v", vexp, perr)
	}
	if bctx.Hex != vexp {
		t.Errorf("Failure: ApplyContext set BaseContext Hexadecimal expect %v but get %v", vexp, bctx.Hex)
	}
	if exp.Hex != vexp {
		t.Errorf("Failure: ApplyContext returns expect %v but get %v", vexp, exp.Hex)
	}
}

func Test_HexParseInt64_ExpectSuccess(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	var bctx BaseContext
	var hc HexConverter
	// When
	bctx.Hex = "1"

	// Then
	r, err := hc.ParseInt64(&bctx)
	if err != nil {
		t.Errorf("Failure: expect %v but get %v with error %v", bctx.Hex, r, err)
	}
}

func Test_HexParseInt64_ExpectErrorByEmptyString(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	var bctx BaseContext
	var hc HexConverter
	errExpect := "strconv.ParseInt: parsing \"\": invalid syntax"
	// When
	bctx.Hex = ""

	// Then
	r, err := hc.ParseInt64(&bctx)
	if err == nil {
		t.Errorf("Failure: expect error 'error strconv.ParseInt: parsing \"\": invalid syntax' but get %v with parameter %v", r, bctx.Hex)
	}
	if errExpect != err.Error() {
		t.Errorf("Failure: expect 'strconv.ParseInt: parsing \"\": invalid syntax' but get '%v'", err)
	}
}

func Test_HexParseInt64_ExpectErrorByAlphaString(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	var bctx BaseContext
	var hc HexConverter
	errExpect := "strconv.ParseInt: parsing \"z\": invalid syntax"
	// When
	bctx.Hex = "z"

	// Then
	r, err := hc.ParseInt64(&bctx)
	if err == nil {
		t.Errorf("Failure: expect error 'error strconv.ParseInt: parsing \"z\": invalid syntax' but get %v with parameter %v", r, bctx.Hex)
	}
	if errExpect != err.Error() {
		t.Errorf("Failure: expect 'strconv.ParseInt: parsing \"z\": invalid syntax' but get '%v'", err)
	}
}

func Test_HexExists_ExpectSuccess(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	var bctx BaseContext
	var hc HexConverter
	// When
	bctx.Hex = "1"

	// Then
	r := hc.Exists(&bctx)
	if !r {
		t.Errorf("Failure: expect 'True' but get 'False' with context {%s}", bctx.Hex)
	}
}

func Test_HexExists_ExpectFailure(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	var bctx BaseContext
	var hc HexConverter
	// When
	bctx.Hex = ""

	// Then
	r := hc.Exists(&bctx)
	if r {
		t.Errorf("Failure: expect 'False' but get 'True' with context {%s}", bctx.Hex)
	}
}
