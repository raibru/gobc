package converter

import "testing"

func Test_Init_HexType_Successful(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	wantBaseId := "x"
	wantBaseNumber := 16
	wantValue := "NaN"
	wantLeadZeros := "0"
	wantIsInit := true

	b := &HexType{Base: BaseType{Number: 0, IsInit: false}}

	// When
	b.Init()

	// Then
	if b.Base.BaseId != wantBaseId {
		t.Errorf("Want BaseId '%s' but got '%s'", wantBaseId, b.Base.BaseId)
	}
	if b.Base.BaseNumber != wantBaseNumber {
		t.Errorf("Want BaseNumber '%d' but got '%d'", wantBaseNumber, b.Base.BaseNumber)
	}
	if b.Base.Value != wantValue {
		t.Errorf("Want Value '%s' but got '%s'", wantValue, b.Base.Value)
	}
	if b.Base.LeadZeros != wantLeadZeros {
		t.Errorf("Want LeadZeros '%s' but got '%s'", wantLeadZeros, b.Base.LeadZeros)
	}
	if b.Base.IsInit != wantIsInit {
		t.Errorf("Want IsInit '%v' but got '%v'", wantIsInit, b.Base.IsInit)
	}
}

func Test_Init_DecType_Successful(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	wantBaseId := "d"
	wantBaseNumber := 10
	wantValue := "NaN"
	wantLeadZeros := "0"
	wantIsInit := true

	b := &DecType{Base: BaseType{Number: 0, IsInit: false}}

	// When
	b.Init()

	// Then
	if b.Base.BaseId != wantBaseId {
		t.Errorf("Want BaseId '%s' but got '%s'", wantBaseId, b.Base.BaseId)
	}
	if b.Base.BaseNumber != wantBaseNumber {
		t.Errorf("Want BaseNumber '%d' but got '%d'", wantBaseNumber, b.Base.BaseNumber)
	}
	if b.Base.Value != wantValue {
		t.Errorf("Want Value '%s' but got '%s'", wantValue, b.Base.Value)
	}
	if b.Base.LeadZeros != wantLeadZeros {
		t.Errorf("Want LeadZeros '%s' but got '%s'", wantLeadZeros, b.Base.LeadZeros)
	}
	if b.Base.IsInit != wantIsInit {
		t.Errorf("Want IsInit '%v' but got '%v'", wantIsInit, b.Base.IsInit)
	}
}

func Test_Init_OctType_Successful(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	wantBaseId := "o"
	wantBaseNumber := 8
	wantValue := "NaN"
	wantLeadZeros := "0"
	wantIsInit := true

	b := &OctType{Base: BaseType{Number: 0, IsInit: false}}

	// When
	b.Init()

	// Then
	if b.Base.BaseId != wantBaseId {
		t.Errorf("Want BaseId '%s' but got '%s'", wantBaseId, b.Base.BaseId)
	}
	if b.Base.BaseNumber != wantBaseNumber {
		t.Errorf("Want BaseNumber '%d' but got '%d'", wantBaseNumber, b.Base.BaseNumber)
	}
	if b.Base.Value != wantValue {
		t.Errorf("Want Value '%s' but got '%s'", wantValue, b.Base.Value)
	}
	if b.Base.LeadZeros != wantLeadZeros {
		t.Errorf("Want LeadZeros '%s' but got '%s'", wantLeadZeros, b.Base.LeadZeros)
	}
	if b.Base.IsInit != wantIsInit {
		t.Errorf("Want IsInit '%v' but got '%v'", wantIsInit, b.Base.IsInit)
	}
}

func Test_Init_BinType_Successful(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	wantBaseId := "b"
	wantBaseNumber := 2
	wantValue := "NaN"
	wantLeadZeros := "0"
	wantIsInit := true

	b := &BinType{Base: BaseType{Number: 0, IsInit: false}}

	// When
	b.Init()

	// Then
	if b.Base.BaseId != wantBaseId {
		t.Errorf("Want BaseId '%s' but got '%s'", wantBaseId, b.Base.BaseId)
	}
	if b.Base.BaseNumber != wantBaseNumber {
		t.Errorf("Want BaseNumber '%d' but got '%d'", wantBaseNumber, b.Base.BaseNumber)
	}
	if b.Base.Value != wantValue {
		t.Errorf("Want Value '%s' but got '%s'", wantValue, b.Base.Value)
	}
	if b.Base.LeadZeros != wantLeadZeros {
		t.Errorf("Want LeadZeros '%s' but got '%s'", wantLeadZeros, b.Base.LeadZeros)
	}
	if b.Base.IsInit != wantIsInit {
		t.Errorf("Want IsInit '%v' but got '%v'", wantIsInit, b.Base.IsInit)
	}
}

func Test_String_HexType_Successful(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	want := "NaN"

	b := &HexType{Base: BaseType{Number: 0, IsInit: false}}
	b.Init()

	// When
	got := b.String()

	// Then
	if want != got {
		t.Errorf("Want '%s' but got '%s'", want, got)
	}
}

func Test_String_DecType_Successful(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	want := "NaN"

	b := &DecType{Base: BaseType{Number: 0, IsInit: false}}
	b.Init()

	// When
	got := b.String()

	// Then
	if want != got {
		t.Errorf("Want '%s' but got '%s'", want, got)
	}
}

func Test_String_OctType_Successful(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	want := "NaN"

	b := &OctType{Base: BaseType{Number: 0, IsInit: false}}
	b.Init()

	// When
	got := b.String()

	// Then
	if want != got {
		t.Errorf("Want '%s' but got '%s'", want, got)
	}
}

func Test_String_BinType_Successful(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	want := "NaN"

	b := &BinType{Base: BaseType{Number: 0, IsInit: false}}
	b.Init()

	// When
	got := b.String()

	// Then
	if want != got {
		t.Errorf("Want '%s' but got '%s'", want, got)
	}
}

func Test_SetupBy_HexType_Successful(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	wantBaseId := "x"
	wantBaseNumber := 16
	wantIsInit := true
	tables := []struct {
		inputNumber    uint64
		inputLeadZeros string
		inputUsePrefix bool
		wantLeadZeros  string
		wantValue      string
	}{
		{0, "", false, "0", "0"},
		{0, "x:1", false, "1", "0"},
		{1, "x:4", false, "4", "0001"},
		{1, "x:4", true, "4", "0x0001"},
		{255, "x:4", true, "4", "0x00ff"},
		{65535, "x:4", true, "4", "0xffff"},
	}

	for _, table := range tables {
		p := Parameters{
			LeadZeros: table.inputLeadZeros,
			UsePrefix: table.inputUsePrefix,
		}
		b := &HexType{Base: BaseType{Number: table.inputNumber, IsInit: false}}

		// When
		b.Init()
		b.SetupBy(&p)

		// Then
		if b.Base.BaseId != wantBaseId {
			t.Errorf("Want BaseId '%s' but got '%s'", wantBaseId, b.Base.BaseId)
		}
		if b.Base.BaseNumber != wantBaseNumber {
			t.Errorf("Want BaseNumber '%d' but got '%d'", wantBaseNumber, b.Base.BaseNumber)
		}
		if b.Base.Value != table.wantValue {
			t.Errorf("Want Value '%s' but got '%s'", table.wantValue, b.Base.Value)
		}
		if b.Base.LeadZeros != table.wantLeadZeros {
			t.Errorf("Want LeadZeros '%s' but got '%s'", table.wantLeadZeros, b.Base.LeadZeros)
		}
		if b.Base.IsInit != wantIsInit {
			t.Errorf("Want IsInit '%v' but got '%v'", wantIsInit, b.Base.IsInit)
		}
	}
}

func Test_SetupBy_DecType_Successful(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	wantBaseId := "d"
	wantBaseNumber := 10
	wantIsInit := true
	tables := []struct {
		inputNumber    uint64
		inputLeadZeros string
		inputUsePrefix bool
		wantLeadZeros  string
		wantValue      string
	}{
		{0, "", false, "0", "0"},
		{0, "d:1", false, "1", "0"},
		{1, "d:4", false, "4", "0001"},
		{1, "d:4", true, "4", "0d0001"},
		{255, "d:4", true, "4", "0d0255"},
		{65535, "d:5", true, "5", "0d65535"},
	}

	for _, table := range tables {
		p := Parameters{
			LeadZeros: table.inputLeadZeros,
			UsePrefix: table.inputUsePrefix,
		}
		b := &DecType{Base: BaseType{Number: table.inputNumber, IsInit: false}}

		// When
		b.Init()
		b.SetupBy(&p)

		// Then
		if b.Base.BaseId != wantBaseId {
			t.Errorf("Want BaseId '%s' but got '%s'", wantBaseId, b.Base.BaseId)
		}
		if b.Base.BaseNumber != wantBaseNumber {
			t.Errorf("Want BaseNumber '%d' but got '%d'", wantBaseNumber, b.Base.BaseNumber)
		}
		if b.Base.Value != table.wantValue {
			t.Errorf("Want Value '%s' but got '%s'", table.wantValue, b.Base.Value)
		}
		if b.Base.LeadZeros != table.wantLeadZeros {
			t.Errorf("Want LeadZeros '%s' but got '%s'", table.wantLeadZeros, b.Base.LeadZeros)
		}
		if b.Base.IsInit != wantIsInit {
			t.Errorf("Want IsInit '%v' but got '%v'", wantIsInit, b.Base.IsInit)
		}
	}
}

func Test_SetupBy_OctType_Successful(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	wantBaseId := "o"
	wantBaseNumber := 8
	wantIsInit := true
	tables := []struct {
		inputNumber    uint64
		inputLeadZeros string
		inputUsePrefix bool
		wantLeadZeros  string
		wantValue      string
	}{
		{0, "", false, "0", "0"},
		{0, "o:1", false, "1", "0"},
		{1, "o:4", false, "4", "0001"},
		{1, "o:4", true, "4", "0o0001"},
		{255, "o:4", true, "4", "0o0377"},
		{65535, "o:6", true, "6", "0o177777"},
	}

	for _, table := range tables {
		p := Parameters{
			LeadZeros: table.inputLeadZeros,
			UsePrefix: table.inputUsePrefix,
		}
		b := &OctType{Base: BaseType{Number: table.inputNumber, IsInit: false}}

		// When
		b.Init()
		b.SetupBy(&p)

		// Then
		if b.Base.BaseId != wantBaseId {
			t.Errorf("Want BaseId '%s' but got '%s'", wantBaseId, b.Base.BaseId)
		}
		if b.Base.BaseNumber != wantBaseNumber {
			t.Errorf("Want BaseNumber '%d' but got '%d'", wantBaseNumber, b.Base.BaseNumber)
		}
		if b.Base.Value != table.wantValue {
			t.Errorf("Want Value '%s' but got '%s'", table.wantValue, b.Base.Value)
		}
		if b.Base.LeadZeros != table.wantLeadZeros {
			t.Errorf("Want LeadZeros '%s' but got '%s'", table.wantLeadZeros, b.Base.LeadZeros)
		}
		if b.Base.IsInit != wantIsInit {
			t.Errorf("Want IsInit '%v' but got '%v'", wantIsInit, b.Base.IsInit)
		}
	}
}

func Test_SetupBy_BinType_Successful(t *testing.T) {
	//t.Fatal("Check Failure")
	// Given
	wantBaseId := "b"
	wantBaseNumber := 2
	wantIsInit := true
	tables := []struct {
		inputNumber    uint64
		inputLeadZeros string
		inputUsePrefix bool
		wantLeadZeros  string
		wantValue      string
	}{
		{0, "", false, "0", "0"},
		{0, "b:1", false, "1", "0"},
		{1, "b:4", false, "4", "0001"},
		{1, "b:4", true, "4", "0b0001"},
		{255, "b:16", true, "16", "0b0000000011111111"},
		{65535, "b:16", true, "16", "0b1111111111111111"},
	}

	for _, table := range tables {
		p := Parameters{
			LeadZeros: table.inputLeadZeros,
			UsePrefix: table.inputUsePrefix,
		}
		b := &BinType{Base: BaseType{Number: table.inputNumber, IsInit: false}}

		// When
		b.Init()
		b.SetupBy(&p)

		// Then
		if b.Base.BaseId != wantBaseId {
			t.Errorf("Want BaseId '%s' but got '%s'", wantBaseId, b.Base.BaseId)
		}
		if b.Base.BaseNumber != wantBaseNumber {
			t.Errorf("Want BaseNumber '%d' but got '%d'", wantBaseNumber, b.Base.BaseNumber)
		}
		if b.Base.Value != table.wantValue {
			t.Errorf("Want Value '%s' but got '%s'", table.wantValue, b.Base.Value)
		}
		if b.Base.LeadZeros != table.wantLeadZeros {
			t.Errorf("Want LeadZeros '%s' but got '%s'", table.wantLeadZeros, b.Base.LeadZeros)
		}
		if b.Base.IsInit != wantIsInit {
			t.Errorf("Want IsInit '%v' but got '%v'", wantIsInit, b.Base.IsInit)
		}
	}
}
