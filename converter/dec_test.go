package converter

import (
	"bufio"
	"strconv"
	"strings"
	"testing"
)

func Test_DecApplyPipeInput_ExpectSuccess(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	var bctx BaseContext
	var dc DecConverter
	var vexp string

	// When
	vexp = "1"
	bctx.Dec = "."
	s := bufio.NewScanner(strings.NewReader(vexp))
	eerr := dc.ApplyPipeInput(s, &bctx)

	// Then
	if bctx.Dec != vexp {
		t.Errorf("Failure: ApplyPipeInput set Decimal BaseContext expect %v but get %v", vexp, bctx.Dec)
	}
	if eerr != nil {
		t.Errorf("Failure: ApplyPipeInput returns error: %v", eerr)
	}
}

func Test_DecApplyPipeInput_FailureByWrongParam(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	var bctx BaseContext
	var dc DecConverter
	var vexp string

	// When
	vexp = "1"
	bctx.Dec = "#"
	s := bufio.NewScanner(strings.NewReader(vexp))
	eerr := dc.ApplyPipeInput(s, &bctx)

	// Then
	if bctx.Dec == vexp {
		t.Errorf("Failure: ApplyPipeInput shall not apply pipe input expect %s but get %v", "#", bctx.Dec)
	}
	if eerr == nil {
		t.Error("Failure: ApplyPipeInput shall returns an error but is nil")
	}
}

func Test_DecApplyContext_ExpectSuccess(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	var bctx BaseContext
	var dc DecConverter
	var vexp string
	var v int64

	leadZero := "b:0"

	// When
	m, ferr := ParseLeadingZero("d:0")
	vexp = "1"
	v, perr := strconv.ParseInt(vexp, 10, 64)
	bctx.Dec = ""
	exp := dc.ApplyContext(v, &bctx, m)

	// Then
	if ferr != nil {
		t.Fatalf("Fatal: ParseLeadingZero (%s) failed: %v", leadZero, ferr)
	}
	if perr != nil {
		t.Fatalf("Fatal: expect no error ParseInt %v to base 10: %v", vexp, perr)
	}
	if bctx.Dec != vexp {
		t.Errorf("Failure: ApplyContext set BaseContext Decimal expect %v but get %v", vexp, bctx.Dec)
	}
	if exp.Dec != vexp {
		t.Errorf("Failure: ApplyContext returns expect %v but get %v", vexp, exp.Dec)
	}
}

func Test_DecParseInt64_ExpectSuccess(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	var bctx BaseContext
	var dc DecConverter
	// When
	bctx.Dec = "1"

	// Then
	r, err := dc.ParseInt64(&bctx)
	if err != nil {
		t.Errorf("Failure: expect %v but get %v with error %v", bctx.Dec, r, err)
	}
}

func Test_DecParseInt64_ExpectErrorByEmptyString(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	var bctx BaseContext
	var dc DecConverter
	errExpect := "strconv.ParseInt: parsing \"\": invalid syntax"
	// When
	bctx.Dec = ""

	// Then
	r, err := dc.ParseInt64(&bctx)
	if err == nil {
		t.Errorf("Failure: expect error 'error strconv.ParseInt: parsing \"\": invalid syntax' but get %v with parameter %v", r, bctx.Dec)
	}
	if errExpect != err.Error() {
		t.Errorf("Failure: expect 'strconv.ParseInt: parsing \"\": invalid syntax' but get '%v'", err)
	}
}

func Test_DecParseInt64_ExpectErrorByAlphaString(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	var bctx BaseContext
	var dc DecConverter
	errExpect := "strconv.ParseInt: parsing \"a\": invalid syntax"
	// When
	bctx.Dec = "a"

	// Then
	r, err := dc.ParseInt64(&bctx)
	if err == nil {
		t.Errorf("Failure: expect error 'error strconv.ParseInt: parsing \"a\": invalid syntax' but get %v with parameter %v", r, bctx.Dec)
	}
	if errExpect != err.Error() {
		t.Errorf("Failure: expect 'strconv.ParseInt: parsing \"a\": invalid syntax' but get '%v'", err)
	}
}

func Test_DecExists_ExpectSuccess(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	var bctx BaseContext
	var dc DecConverter
	// When
	bctx.Dec = "1"

	// Then
	r := dc.Exists(&bctx)
	if !r {
		t.Errorf("Failure: expect 'True' but get 'False' with context {%s}", bctx.Dec)
	}
}

func Test_DecExists_ExpectFailure(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	var bctx BaseContext
	var dc DecConverter
	// When
	bctx.Dec = ""

	// Then
	r := dc.Exists(&bctx)
	if r {
		t.Errorf("Failure: expect 'False' but get 'True' with context {%s}", bctx.Dec)
	}
}
