package converter

import (
	"bufio"
	"strconv"
	"strings"
	"testing"
)

func Test_OctApplyPipeInput_ExpectSuccess(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	var bctx BaseContext
	var oc OctConverter
	var vexp string

	// When
	vexp = "1"
	bctx.Oct = "."
	s := bufio.NewScanner(strings.NewReader(vexp))
	eerr := oc.ApplyPipeInput(s, &bctx)

	// Then
	if bctx.Oct != vexp {
		t.Errorf("Failure: ApplyPipeInput set Octal BaseContext expect %v but get %v", vexp, bctx.Oct)
	}
	if eerr != nil {
		t.Errorf("Failure: ApplyPipeInput returns error: %v", eerr)
	}
}

func Test_OctApplyPipeInput_FailureByWrongParam(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	var bctx BaseContext
	var oc OctConverter
	var vexp string

	// When
	vexp = "1"
	bctx.Oct = "#"
	s := bufio.NewScanner(strings.NewReader(vexp))
	eerr := oc.ApplyPipeInput(s, &bctx)

	// Then
	if bctx.Oct == vexp {
		t.Errorf("Failure: ApplyPipeInput shall not apply pipe input expect %s but get %v", "#", bctx.Oct)
	}
	if eerr == nil {
		t.Error("Failure: ApplyPipeInput shall returns an error but is nil")
	}
}

func Test_OctApplyContext_ExpectSuccess(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	var bctx BaseContext
	var oc OctConverter
	var vexp string
	var v int64

	leadZero := "o:0"

	// When
	m, ferr := ParseLeadingZero(leadZero)
	vexp = "1"
	v, perr := strconv.ParseInt(vexp, 8, 64)
	bctx.Oct = ""
	exp := oc.ApplyContext(v, &bctx, m)

	// Then
	if ferr != nil {
		t.Fatalf("Fatal: ParseLeadingZero (%s) failed: %v", leadZero, ferr)
	}

	if perr != nil {
		t.Fatalf("Fatal: expect no error ParseInt %v to base 8: %v", vexp, perr)
	}
	if bctx.Oct != vexp {
		t.Errorf("Failure: ApplyContext set BaseContext Octal expect %v but get %v", vexp, bctx.Oct)
	}
	if exp.Oct != vexp {
		t.Errorf("Failure: ApplyContext returns expect %v but get %v", vexp, exp.Oct)
	}
}

func Test_OctParseInt64_ExpectSuccess(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	var bctx BaseContext
	var oc OctConverter
	// When
	bctx.Oct = "1"

	// Then
	r, err := oc.ParseInt64(&bctx)
	if err != nil {
		t.Errorf("Failure: expect %v but get %v with error %v", bctx.Oct, r, err)
	}
}

func Test_OctParseInt64_ExpectErrorByEmptyString(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	var bctx BaseContext
	var oc OctConverter
	errExpect := "strconv.ParseInt: parsing \"\": invalid syntax"
	// When
	bctx.Oct = ""

	// Then
	r, err := oc.ParseInt64(&bctx)
	if err == nil {
		t.Errorf("Failure: expect error 'error strconv.ParseInt: parsing \"\": invalid syntax' but get %v with parameter %v", r, bctx.Oct)
	}
	if errExpect != err.Error() {
		t.Errorf("Failure: expect 'strconv.ParseInt: parsing \"\": invalid syntax' but get '%v'", err)
	}
}

func Test_OctParseInt64_ExpectErrorByAlphaString(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	var bctx BaseContext
	var oc OctConverter
	errExpect := "strconv.ParseInt: parsing \"a\": invalid syntax"
	// When
	bctx.Oct = "a"

	// Then
	r, err := oc.ParseInt64(&bctx)
	if err == nil {
		t.Errorf("Failure: expect error 'error strconv.ParseInt: parsing \"a\": invalid syntax' but get %v with parameter %v", r, bctx.Oct)
	}
	if errExpect != err.Error() {
		t.Errorf("Failure: expect 'strconv.ParseInt: parsing \"a\": invalid syntax' but get '%v'", err)
	}
}

func Test_OctExists_ExpectSuccess(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	var bctx BaseContext
	var oc OctConverter
	// When
	bctx.Oct = "1"

	// Then
	r := oc.Exists(&bctx)
	if !r {
		t.Errorf("Failure: expect 'True' but get 'False' with context {%s}", bctx.Oct)
	}
}

func Test_OctExists_ExpectFailure(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	var bctx BaseContext
	var oc OctConverter
	// When
	bctx.Oct = ""

	// Then
	r := oc.Exists(&bctx)
	if r {
		t.Errorf("Failure: expect 'False' but get 'True' with context {%s}", bctx.Oct)
	}
}
