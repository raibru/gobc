package converter

import (
	"bufio"
	"strconv"
	"strings"
	"testing"
)

func Test_BinApplyPipeInput_ExpectSuccess(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	var bctx BaseContext
	var bc BinConverter
	var vexp string

	// When
	vexp = "1"
	bctx.Bin = "."
	s := bufio.NewScanner(strings.NewReader(vexp))
	eerr := bc.ApplyPipeInput(s, &bctx)

	// Then
	if bctx.Bin != vexp {
		t.Errorf("Failure: ApplyPipeInput set Binary BaseContext expect %v but get %v", vexp, bctx.Bin)
	}
	if eerr != nil {
		t.Errorf("Failure: ApplyPipeInput returns error: %v", eerr)
	}
}

func Test_BinApplyPipeInput_FailureByWrongParam(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	var bctx BaseContext
	var bc BinConverter
	var vexp string

	// When
	vexp = "1"
	bctx.Bin = "#"
	s := bufio.NewScanner(strings.NewReader(vexp))
	eerr := bc.ApplyPipeInput(s, &bctx)

	// Then
	if bctx.Bin == vexp {
		t.Errorf("Failure: ApplyPipeInput shall not apply pipe input expect %s but get %v", "#", bctx.Bin)
	}
	if eerr == nil {
		t.Error("Failure: ApplyPipeInput shall returns an error but is nil")
	}
}

func Test_BinApplyContext_ExpectSuccess(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	var bctx BaseContext
	var bc BinConverter
	var vexp string
	var v int64

	// When
	vexp = "1"
	v, perr := strconv.ParseInt(vexp, 2, 64)
	bctx.Bin = ""
	exp := bc.ApplyContext(v, &bctx)

	// Then
	if perr != nil {
		t.Fatalf("Fatal: expect no error ParseInt %v to base 2: %v", vexp, perr)
	}
	if bctx.Bin != vexp {
		t.Errorf("Failure: ApplyContext set Binary BaseContext expect %v but get %v", vexp, bctx.Bin)
	}
	if exp.Bin != vexp {
		t.Errorf("Failure: ApplyContext returns expect %v but get %v", vexp, exp.Bin)
	}
}

func Test_BinParseInt64_ExpectSuccess(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	var bctx BaseContext
	var bc BinConverter
	// When
	bctx.Bin = "1"

	// Then
	r, err := bc.ParseInt64(&bctx)
	if err != nil {
		t.Errorf("Failure: expect %v but get %v with error %v", bctx.Bin, r, err)
	}
}

func Test_BinParseInt64_ExpectErrorByEmptyString(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	var bctx BaseContext
	var bc BinConverter
	errExpect := "strconv.ParseInt: parsing \"\": invalid syntax"
	// When
	bctx.Bin = ""

	// Then
	r, err := bc.ParseInt64(&bctx)
	if err == nil {
		t.Errorf("Failure: expect error 'error strconv.ParseInt: parsing \"\": invalid syntax' but get %v with parameter %v", r, bctx.Bin)
	}
	if errExpect != err.Error() {
		t.Errorf("Failure: expect 'strconv.ParseInt: parsing \"\": invalid syntax' but get '%v'", err)
	}
}

func Test_BinParseInt64_ExpectErrorByDec2String(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	var bctx BaseContext
	var bc BinConverter
	errExpect := "strconv.ParseInt: parsing \"2\": invalid syntax"
	// When
	bctx.Bin = "2"

	// Then
	r, err := bc.ParseInt64(&bctx)
	if err == nil {
		t.Errorf("Failure: expect error 'error strconv.ParseInt: parsing \"2\": invalid syntax' but get %v with parameter %v", r, bctx.Bin)
	}
	if errExpect != err.Error() {
		t.Errorf("Failure: expect 'strconv.ParseInt: parsing \"2\": invalid syntax' but get '%v'", err)
	}
}

func Test_BinParseInt64_ExpectErrorByAlphaString(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	var bctx BaseContext
	var bc BinConverter
	errExpect := "strconv.ParseInt: parsing \"a\": invalid syntax"
	// When
	bctx.Bin = "a"

	// Then
	r, err := bc.ParseInt64(&bctx)
	if err == nil {
		t.Errorf("Failure: expect error 'error strconv.ParseInt: parsing \"a\": invalid syntax' but get %v with parameter %v", r, bctx.Bin)
	}
	if errExpect != err.Error() {
		t.Errorf("Failure: expect 'strconv.ParseInt: parsing \"a\": invalid syntax' but get '%v'", err)
	}
}

func Test_BinExists_ExpectSuccess(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	var bctx BaseContext
	var bc BinConverter
	// When
	bctx.Bin = "1"

	// Then
	r := bc.Exists(&bctx)
	if !r {
		t.Errorf("Failure: expect 'True' but get 'False' with context {%s}", bctx.Bin)
	}
}

func Test_BinExists_ExpectFailure(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	var bctx BaseContext
	var bc BinConverter
	// When
	bctx.Bin = ""

	// Then
	r := bc.Exists(&bctx)
	if r {
		t.Errorf("Failure: expect 'False' but get 'True' with context {%s}", bctx.Bin)
	}
}
