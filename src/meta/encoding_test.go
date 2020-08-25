package meta

import (
	"strings"
	"testing"
)

func TestTypeEncoding(t *testing.T) {
	tests := []struct {
		wrapped     string
		stringified string
		unwrapCheck func(typ string) bool
	}{
		{`int`, `int`, func(typ string) bool { return typ == `int` }},
		{`\Foo`, `\Foo`, func(typ string) bool { return typ == `\Foo` }},

		{
			WrapArrayOf(`int`), `int[]`,
			func(typ string) bool { return unwrap1(typ) == `int` },
		},

		{
			WrapBaseMethodParam(0, `FooClass`, `barMethod`),
			`param(FooClass)::barMethod[0]`,
			func(typ string) bool {
				index, className, methodName := unwrap3(typ)
				return index == 0 && className == `FooClass` && methodName == `barMethod`
			},
		},

		{
			WrapBaseMethodParam('|', `FooClass`, `barMethod`),
			`param(FooClass)::barMethod[124]`,
			func(typ string) bool {
				index, className, methodName := unwrap3(typ)
				return index == '|' && className == `FooClass` && methodName == `barMethod`
			},
		},

		{
			WrapArrayOf(strings.Repeat(`a`, '|')),
			strings.Repeat(`a`, '|') + `[]`,
			func(typ string) bool {
				return unwrap1(typ) == strings.Repeat(`a`, '|')
			},
		},
	}

	for _, test := range tests {
		stringified := formatType(test.wrapped)
		if stringified != test.stringified {
			t.Errorf("formatType %q:\nhave: %q\nwant: %q",
				test.stringified, stringified, test.stringified)
			continue
		}

		if !test.unwrapCheck(test.wrapped) {
			t.Errorf("unwrap check failed for %q", test.stringified)
			continue
		}

		const typeSuffix = "|zzz|zzz[]"
		m := NewTypesMap(test.stringified + typeSuffix)
		if m.String() != test.stringified+typeSuffix {
			t.Errorf("type map string mismatched for %q:\nhave: %q\nwant: %q",
				test.stringified, m.String(), test.stringified+typeSuffix)
			continue
		}

		encoded, err := m.GobEncode()
		if err != nil {
			t.Errorf("failed to gob-encode %q: %v", test.stringified, err)
			continue
		}
		decoded := NewEmptyTypesMap(m.Len())
		if err := decoded.GobDecode(encoded); err != nil {
			t.Errorf("failed to gob-decode %q: %v", test.stringified, err)
			continue
		}
		if m.String() != decoded.String() {
			t.Errorf("decoded type string mismatched for %q:\nhave: %q\nwant: %q",
				test.stringified, decoded.String(), m.String())
		}

		// TODO(quasilyte): do something so we don't need to care about
		// encoded (wrapped) so much. For now, we need tests below to ensure
		// that nothing blows up because some <uint8> or type byte look like '|'
		// when printed as a string.
		if len(strings.Split(test.wrapped, `|`)) != len(strings.Split(test.stringified, `|`)) {
			t.Errorf("mismatched |-split parts count for %q:\nhave: %d\nwant: %d",
				test.stringified,
				len(strings.Split(test.wrapped, `|`)),
				len(strings.Split(test.stringified, `|`)))
		}
	}
}

func TestConstantValueDecodeEncode(t *testing.T) {
	testCases := []ConstValue{
		{Type: String, Value: "world"},
		{Type: Integer, Value: int64(5)},
		{Type: Float, Value: 5.56},
		{Type: String, Value: "hello"},
		{Type: Float, Value: 124.67},
		{Type: Integer, Value: int64(50000000)},
	}

	for _, testCase := range testCases {
		// encode this
		encoded, err := testCase.GobEncode()
		if err != nil {
			t.Errorf("unexpected error \"%s\"", err)
		}

		// decode this
		decoded := ConstValue{}
		err = decoded.GobDecode(encoded)
		if err != nil {
			t.Errorf("unexpected error \"%s\"", err)
		}

		// compare
		if decoded.Type != testCase.Type || decoded.Value != testCase.Value {
			t.Error("error decode, objects not equal")
		}
	}
}
