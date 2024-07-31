package co

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

// test helpers
// string2String: string to string
// string2Int: string to integer - TODO: is this needed?
// int2String: integer to string
// int2Int: integer to integer
// any2String: any type to string

type string2StringTestCase struct {
	in              string
	out             string
	expectedErrText string
}

type string2IntTestCase struct {
	in              string
	out             *big.Int
	expectedErrText string
}

type int2StringTestCase struct {
	in              *big.Int
	out             string
	expectedErrText string
}

type int2IntTestCase struct {
	in              *big.Int
	out             *big.Int
	expectedErrText string
}

type any2StringTestCase struct {
	in              interface{}
	out             string
	expectedErrText string
}

type string2StringCoFn func(string) (string, error)
type string2IntCoFn func(string) (*big.Int, error)
type int2StringCoFn func(*big.Int) (string, error)
type int2IntCoFn func(*big.Int) (*big.Int, error)
type any2StringCoFn func(interface{}) (string, error)

func string2StringTestRunner(t *testing.T, testCases []string2StringTestCase, f string2StringCoFn) {
	for _, tt := range testCases {
		t.Run(tt.in, func(t *testing.T) {

			actualOut, actualErr := f(tt.in)

			assert.Equal(t, tt.out, actualOut)
			if tt.expectedErrText == "" {
				assert.NoError(t, actualErr)
			} else {
				assert.Error(t, actualErr)
				if actualErr != nil {
					assert.Equal(t, tt.expectedErrText, actualErr.Error())
				}
			}
		})
	}
}

func string2IntTestRunner(t *testing.T, testCases []string2IntTestCase, f string2IntCoFn) {
	for _, tt := range testCases {
		t.Run(tt.in, func(t *testing.T) {

			actualOut, actualErr := f(tt.in)

			assert.Equal(t, tt.out, actualOut)
			if tt.expectedErrText == "" {
				assert.NoError(t, actualErr)
			} else {
				assert.Error(t, actualErr)
				if actualErr != nil {
					assert.Equal(t, tt.expectedErrText, actualErr.Error())
				}
			}
		})
	}
}

func int2StringTestRunner(t *testing.T, testCases []int2StringTestCase, f int2StringCoFn) {
	for _, tt := range testCases {
		t.Run(tt.in.String(), func(t *testing.T) {

			actualOut, actualErr := f(tt.in)

			assert.Equal(t, tt.out, actualOut)
			if tt.expectedErrText == "" {
				assert.NoError(t, actualErr)
			} else {
				assert.Error(t, actualErr)
				if actualErr != nil {
					assert.Equal(t, tt.expectedErrText, actualErr.Error())
				}
			}
		})
	}
}

func int2IntTestRunner(t *testing.T, testCases []int2IntTestCase, f int2IntCoFn) {
	for _, tt := range testCases {
		t.Run(tt.in.String(), func(t *testing.T) {

			actualOut, actualErr := f(tt.in)

			assert.Equal(t, tt.out, actualOut)
			if tt.expectedErrText == "" {
				assert.NoError(t, actualErr)
			} else {
				assert.Error(t, actualErr)
				if actualErr != nil {
					assert.Equal(t, tt.expectedErrText, actualErr.Error())
				}
			}
		})
	}
}

func any2StringTestRunner(t *testing.T, testCases []any2StringTestCase, f any2StringCoFn) {
	for _, tt := range testCases {
		var testName string
		switch v := tt.in.(type) {
		case string:
			testName = v
		case *big.Int:
			testName = v.String()
		default:
			testName = "unknown"
		}

		t.Run(testName, func(t *testing.T) {
			actualOut, actualErr := f(tt.in)

			assert.Equal(t, tt.out, actualOut)
			if tt.expectedErrText == "" {
				assert.NoError(t, actualErr)
			} else {
				assert.Error(t, actualErr)
				if actualErr != nil {
					assert.Equal(t, tt.expectedErrText, actualErr.Error())
				}
			}
		})
	}
}

// --- tests ---

func TestPatp(t *testing.T) {
	var testCases = []any2StringTestCase{
		{
			in:  "0",
			out: "~zod",
		},
		{
			in:  "255",
			out: "~fes",
		},
		{
			in:  "256",
			out: "~marzod",
		},
		{
			in:  "65535",
			out: "~fipfes",
		},
		{
			in:  "65536",
			out: "~dapnep-ronmyl",
		},
		{
			in:  "14287616",
			out: "~rosmur-hobrem",
		},
		{
			in:  "14287617",
			out: "~sallus-nodlut",
		},
		{
			in:  "14287618",
			out: "~marder-mopdur",
		},
		{
			in:  "14287619",
			out: "~laphec-savted",
		},
		{
			in:  "4294967295",
			out: "~dostec-risfen",
		},
		{
			in:  "4294967296",
			out: "~doznec-dozzod-dozzod",
		},
		{
			in:  big.NewInt(0),
			out: "~zod",
		},
		{
			in:  big.NewInt(255),
			out: "~fes",
		},
		{
			in:  big.NewInt(256),
			out: "~marzod",
		},
		{
			in:  big.NewInt(65535),
			out: "~fipfes",
		},
		{
			in:  big.NewInt(65536),
			out: "~dapnep-ronmyl",
		},
		{
			in:  big.NewInt(14287616),
			out: "~rosmur-hobrem",
		},
		{
			in:  big.NewInt(14287617),
			out: "~sallus-nodlut",
		},
		{
			in:  big.NewInt(14287618),
			out: "~marder-mopdur",
		},
		{
			in:  big.NewInt(14287619),
			out: "~laphec-savted",
		},
		{
			in:  big.NewInt(4294967295),
			out: "~dostec-risfen",
		},
		{
			in:  big.NewInt(4294967296),
			out: "~doznec-dozzod-dozzod",
		},
		{
			in:              "abcdefg",
			expectedErrText: "invalid integer string: abcdefg",
		},
	}

	any2StringTestRunner(t, testCases, Patp)
}

func TestPatq(t *testing.T) {
	var testCases = []any2StringTestCase{
		{
			in:  "0",
			out: "~zod",
		},
		{
			in:  "255",
			out: "~fes",
		},
		{
			in:  "256",
			out: "~marzod",
		},
		{
			in:  "65535",
			out: "~fipfes",
		},
		{
			in:  "65536",
			out: "~doznec-dozzod",
		},
		{
			in:  "14287616",
			out: "~dozler-wanzod",
		},
		{
			in:  "14287617",
			out: "~dozler-wannec",
		},
		{
			in:  "14287618",
			out: "~dozler-wanbud",
		},
		{
			in:  "14287619",
			out: "~dozler-wanwes",
		},
		{
			in:  "4294967295",
			out: "~fipfes-fipfes",
		},
		{
			in:  "4294967296",
			out: "~doznec-dozzod-dozzod",
		},
		{
			in:  big.NewInt(0),
			out: "~zod",
		},
		{
			in:  big.NewInt(255),
			out: "~fes",
		},
		{
			in:  big.NewInt(256),
			out: "~marzod",
		},
		{
			in:  big.NewInt(65535),
			out: "~fipfes",
		},
		{
			in:  big.NewInt(65536),
			out: "~doznec-dozzod",
		},
		{
			in:  big.NewInt(14287616),
			out: "~dozler-wanzod",
		},
		{
			in:  big.NewInt(14287617),
			out: "~dozler-wannec",
		},
		{
			in:  big.NewInt(14287618),
			out: "~dozler-wanbud",
		},
		{
			in:  big.NewInt(14287619),
			out: "~dozler-wanwes",
		},
		{
			in:  big.NewInt(4294967295),
			out: "~fipfes-fipfes",
		},
		{
			in:  big.NewInt(4294967296),
			out: "~doznec-dozzod-dozzod",
		},
		{
			in:              "abcdefg",
			expectedErrText: "invalid integer string: abcdefg",
		},
	}

	any2StringTestRunner(t, testCases, Patq)
}

func TestClan(t *testing.T) {
	var testCases = []string2StringTestCase{
		{
			in:  "~zod",
			out: ShipClassGalaxy,
		},
		{
			in:  "~fes",
			out: ShipClassGalaxy,
		},
		{
			in:  "~marzod",
			out: ShipClassStar,
		},
		{
			in:  "~fipfes",
			out: ShipClassStar,
		},
		{
			in:  "~dapnep-ronmyl",
			out: ShipClassPlanet,
		},
		{
			in:  "~rosmur-hobrem",
			out: ShipClassPlanet,
		},
		{
			in:  "~sallus-nodlut",
			out: ShipClassPlanet,
		},
		{
			in:  "~marder-mopdur",
			out: ShipClassPlanet,
		},
		{
			in:  "~laphec-savted",
			out: ShipClassPlanet,
		},
		{
			in:  "~dostec-risfen",
			out: ShipClassPlanet,
		},
		{
			in:  "~divrul-dalred-samhec-sidrex",
			out: ShipClassMoon,
		},
		{
			in:  "~dotmec-niblyd-tocdys-ravryg--panper-hilsug-nidnev-marzod",
			out: ShipClassComet,
		},
		{
			in:              "abcdefg",
			expectedErrText: "invalid @p: abcdefg",
		},
	}

	string2StringTestRunner(t, testCases, Clan)
}

func TestSein(t *testing.T) {
	var testCases = []string2StringTestCase{
		{
			in:  "~zod",
			out: "~zod",
		},
		{
			in:  "~fes",
			out: "~fes",
		},
		{
			in:  "~marzod",
			out: "~zod",
		},
		{
			in:  "~fipfes",
			out: "~fes",
		},
		{
			in:  "~dapnep-ronmyl",
			out: "~zod",
		},
		{
			in:  "~rosmur-hobrem",
			out: "~wanzod",
		},
		{
			in:  "~sallus-nodlut",
			out: "~wannec",
		},
		{
			in:  "~marder-mopdur",
			out: "~wanbud",
		},
		{
			in:  "~laphec-savted",
			out: "~wanwes",
		},
		{
			in:  "~dostec-risfen",
			out: "~fipfes",
		},
		{
			in:  "~divrul-dalred-samhec-sidrex",
			out: "~samhec-sidrex",
		},
		{
			in:  "~dotmec-niblyd-tocdys-ravryg--panper-hilsug-nidnev-marzod",
			out: "~zod",
		},
		{
			in:              "abcdefg",
			expectedErrText: "invalid @p: abcdefg",
		},
	}

	string2StringTestRunner(t, testCases, Sein)
}

func TestPatp2Dec(t *testing.T) {
	var testCases = []string2StringTestCase{
		{
			out: "0",
			in:  "~zod",
		},
		{
			out: "255",
			in:  "~fes",
		},
		{
			out: "256",
			in:  "~marzod",
		},
		{
			out: "65535",
			in:  "~fipfes",
		},
		{
			out: "65536",
			in:  "~dapnep-ronmyl",
		},
		{
			out: "14287616",
			in:  "~rosmur-hobrem",
		},
		{
			out: "14287617",
			in:  "~sallus-nodlut",
		},
		{
			out: "14287618",
			in:  "~marder-mopdur",
		},
		{
			out: "14287619",
			in:  "~laphec-savted",
		},
		{
			out: "4294967295",
			in:  "~dostec-risfen",
		},
		{
			out: "4294967296",
			in:  "~doznec-dozzod-dozzod",
		},
		{
			in:              "abcdefg",
			expectedErrText: "invalid @p: abcdefg",
		},
	}

	string2StringTestRunner(t, testCases, Patp2Dec)
}

func TestPatq2Dec(t *testing.T) {
	var testCases = []string2StringTestCase{
		{
			out: "0",
			in:  "~zod",
		},
		{
			out: "255",
			in:  "~fes",
		},
		{
			out: "256",
			in:  "~marzod",
		},
		{
			out: "65535",
			in:  "~fipfes",
		},
		{
			out: "65536",
			in:  "~doznec-dozzod",
		},
		{
			out: "14287616",
			in:  "~dozler-wanzod",
		},
		{
			out: "14287617",
			in:  "~dozler-wannec",
		},
		{
			out: "14287618",
			in:  "~dozler-wanbud",
		},
		{
			out: "14287619",
			in:  "~dozler-wanwes",
		},
		{
			out: "4294967295",
			in:  "~fipfes-fipfes",
		},
		{
			out: "4294967296",
			in:  "~doznec-dozzod-dozzod",
		},
		{
			in:              "abcdefg",
			expectedErrText: "invalid @q: abcdefg",
		},
	}

	string2StringTestRunner(t, testCases, Patq2Dec)
}

func TestPatp2Hex(t *testing.T) {
	var testCases = []string2StringTestCase{
		{
			out: "00",
			in:  "~zod",
		},
		{
			out: "ff",
			in:  "~fes",
		},
		{
			out: "0100",
			in:  "~marzod",
		},
		{
			out: "ffff",
			in:  "~fipfes",
		},
		{
			out: "010000",
			in:  "~dapnep-ronmyl",
		},
		{
			out: "da0300",
			in:  "~rosmur-hobrem",
		},
		{
			out: "da0301",
			in:  "~sallus-nodlut",
		},
		{
			out: "da0302",
			in:  "~marder-mopdur",
		},
		{
			out: "da0303",
			in:  "~laphec-savted",
		},
		{
			out: "ffffffff",
			in:  "~dostec-risfen",
		},
		{
			out: "0100000000",
			in:  "~doznec-dozzod-dozzod",
		},
		{
			in:              "abcdefg",
			expectedErrText: "invalid @p: abcdefg",
		},
	}

	string2StringTestRunner(t, testCases, Patp2Hex)
}

func TestPatq2Hex(t *testing.T) {
	var testCases = []string2StringTestCase{
		{
			out: "00",
			in:  "~zod",
		},
		{
			out: "ff",
			in:  "~fes",
		},
		{
			out: "0100",
			in:  "~marzod",
		},
		{
			out: "ffff",
			in:  "~fipfes",
		},
		{
			out: "00010000",
			in:  "~doznec-dozzod",
		},
		// TODO: Look into these leading zeroes; does the same in JS implementation
		{
			out: "00da0300",
			in:  "~dozler-wanzod",
		},
		{
			out: "00da0301",
			in:  "~dozler-wannec",
		},
		{
			out: "00da0302",
			in:  "~dozler-wanbud",
		},
		{
			out: "00da0303",
			in:  "~dozler-wanwes",
		},
		{
			out: "ffffffff",
			in:  "~fipfes-fipfes",
		},
		{
			out: "000100000000",
			in:  "~doznec-dozzod-dozzod",
		},
		{
			in:              "abcdefg",
			expectedErrText: "invalid @q: abcdefg",
		},
	}

	string2StringTestRunner(t, testCases, Patq2Hex)
}

func TestHex2Patp(t *testing.T) {
	var testCases = []string2StringTestCase{
		{
			in:  "00",
			out: "~zod",
		},
		{
			in:  "ff",
			out: "~fes",
		},
		{
			in:  "0100",
			out: "~marzod",
		},
		{
			in:  "ffff",
			out: "~fipfes",
		},
		{
			in:  "010000",
			out: "~dapnep-ronmyl",
		},
		{
			in:  "da0300",
			out: "~rosmur-hobrem",
		},
		{
			in:  "da0301",
			out: "~sallus-nodlut",
		},
		{
			in:  "da0302",
			out: "~marder-mopdur",
		},
		{
			in:  "da0303",
			out: "~laphec-savted",
		},
		{
			in:  "ffffffff",
			out: "~dostec-risfen",
		},
		{
			in:  "0100000000",
			out: "~doznec-dozzod-dozzod",
		},
		{
			in:              "abcdefg",
			expectedErrText: "invalid hexadecimal string: abcdefg",
		},
	}

	string2StringTestRunner(t, testCases, Hex2Patp)
}

func TestHex2Patq(t *testing.T) {
	var testCases = []string2StringTestCase{
		{
			in:  "00",
			out: "~zod",
		},
		{
			in:  "ff",
			out: "~fes",
		},
		{
			in:  "0100",
			out: "~marzod",
		},
		{
			in:  "ffff",
			out: "~fipfes",
		},
		{
			in:  "00010000",
			out: "~doznec-dozzod",
		},
		{
			in:  "00da0300",
			out: "~dozler-wanzod",
		},
		{
			in:  "00da0301",
			out: "~dozler-wannec",
		},
		{
			in:  "00da0302",
			out: "~dozler-wanbud",
		},
		{
			in:  "00da0303",
			out: "~dozler-wanwes",
		},
		{
			in:  "ffffffff",
			out: "~fipfes-fipfes",
		},
		{
			in:  "000100000000",
			out: "~doznec-dozzod-dozzod",
		},
		{
			in:              "abcdefg",
			expectedErrText: "invalid hexadecimal string: abcdefg",
		},
	}

	string2StringTestRunner(t, testCases, Hex2Patq)
}

func TestPatp2Point(t *testing.T) {
	var testCases = []string2IntTestCase{
		{
			in:  "~zod",
			out: big.NewInt(0),
		},
		{
			in:  "~fes",
			out: big.NewInt(255),
		},
		{
			in:  "~marzod",
			out: big.NewInt(256),
		},
		{
			in:  "~fipfes",
			out: big.NewInt(65535),
		},
		{
			in:  "~dapnep-ronmyl",
			out: big.NewInt(65536),
		},
		{
			in:  "~rosmur-hobrem",
			out: big.NewInt(14287616),
		},
		{
			in:  "~sallus-nodlut",
			out: big.NewInt(14287617),
		},
		{
			in:  "~marder-mopdur",
			out: big.NewInt(14287618),
		},
		{
			in:  "~laphec-savted",
			out: big.NewInt(14287619),
		},
		{
			in:  "~dostec-risfen",
			out: big.NewInt(4294967295),
		},
		{
			in:  "~doznec-dozzod-dozzod",
			out: big.NewInt(4294967296),
		},
		{
			in:              "~abcdefg",
			expectedErrText: "invalid @p: ~abcdefg",
		},
	}

	string2IntTestRunner(t, testCases, Patp2Point)
}

// Point2Patp
func TestPoint2Patp(t *testing.T) {
	var testCases = []int2StringTestCase{
		{
			in:  big.NewInt(0),
			out: "~zod",
		},
		{
			in:  big.NewInt(255),
			out: "~fes",
		},
		{
			in:  big.NewInt(256),
			out: "~marzod",
		},
		{
			in:  big.NewInt(65535),
			out: "~fipfes",
		},
		{
			in:  big.NewInt(65536),
			out: "~dapnep-ronmyl",
		},
		{
			in:  big.NewInt(14287616),
			out: "~rosmur-hobrem",
		},
		{
			in:  big.NewInt(14287617),
			out: "~sallus-nodlut",
		},
		{
			in:  big.NewInt(14287618),
			out: "~marder-mopdur",
		},
		{
			in:  big.NewInt(14287619),
			out: "~laphec-savted",
		},
		{
			in:  big.NewInt(4294967295),
			out: "~dostec-risfen",
		},
		{
			in:  big.NewInt(4294967296),
			out: "~doznec-dozzod-dozzod",
		},
	}

	int2StringTestRunner(t, testCases, Point2Patp)
}

// Patq2Point
func TestPatq2Point(t *testing.T) {
	var testCases = []string2IntTestCase{
		{
			in:  "~zod",
			out: big.NewInt(0),
		},
		{
			in:  "~fes",
			out: big.NewInt(255),
		},
		{
			in:  "~marzod",
			out: big.NewInt(256),
		},
		{
			in:  "~fipfes",
			out: big.NewInt(65535),
		},
		{
			in:  "~doznec-dozzod",
			out: big.NewInt(65536),
		},
		{
			in:  "~dozler-wanzod",
			out: big.NewInt(14287616),
		},
		{
			in:  "~dozler-wannec",
			out: big.NewInt(14287617),
		},
		{
			in:  "~dozler-wanbud",
			out: big.NewInt(14287618),
		},
		{
			in:  "~dozler-wanwes",
			out: big.NewInt(14287619),
		},
		{
			in:  "~fipfes-fipfes",
			out: big.NewInt(4294967295),
		},
		{
			in:  "~doznec-dozzod-dozzod",
			out: big.NewInt(4294967296),
		},
		{
			in:              "abcdefg",
			expectedErrText: "invalid @q: abcdefg",
		},
	}

	string2IntTestRunner(t, testCases, Patq2Point)
}

// Point2Patq
func TestPoint2Patq(t *testing.T) {
	var testCases = []int2StringTestCase{
		{
			in:  big.NewInt(0),
			out: "~zod",
		},
		{
			in:  big.NewInt(255),
			out: "~fes",
		},
		{
			in:  big.NewInt(256),
			out: "~marzod",
		},
		{
			in:  big.NewInt(65535),
			out: "~fipfes",
		},
		{
			in:  big.NewInt(65536),
			out: "~doznec-dozzod",
		},
		{
			in:  big.NewInt(14287616),
			out: "~dozler-wanzod",
		},
		{
			in:  big.NewInt(14287617),
			out: "~dozler-wannec",
		},
		{
			in:  big.NewInt(14287618),
			out: "~dozler-wanbud",
		},
		{
			in:  big.NewInt(14287619),
			out: "~dozler-wanwes",
		},
		{
			in:  big.NewInt(4294967295),
			out: "~fipfes-fipfes",
		},
		{
			in:  big.NewInt(4294967296),
			out: "~doznec-dozzod-dozzod",
		},
	}

	int2StringTestRunner(t, testCases, Point2Patq)
}

// SeinPoint
func TestSeinPoint(t *testing.T) {
	var testCases = []int2IntTestCase{
		{
			in:  big.NewInt(0),
			out: big.NewInt(0),
		},
		{
			in:  big.NewInt(255),
			out: big.NewInt(255),
		},
		{
			in:  big.NewInt(256),
			out: big.NewInt(0),
		},
		{
			in:  big.NewInt(65535),
			out: big.NewInt(255),
		},
		{
			in:  big.NewInt(65536),
			out: big.NewInt(0),
		},
		{
			in:  big.NewInt(14287616),
			out: big.NewInt(768),
		},
		{
			in:  big.NewInt(14287617),
			out: big.NewInt(769),
		},
		{
			in:  big.NewInt(14287618),
			out: big.NewInt(770),
		},
		{
			in:  big.NewInt(14287619),
			out: big.NewInt(771),
		},
		{
			in:  big.NewInt(4294967295),
			out: big.NewInt(65535),
		},
		{
			in:  big.NewInt(4294967296),
			out: big.NewInt(0),
		},
	}

	int2IntTestRunner(t, testCases, SeinPoint)
}

// ClanPoint
func TestClanPoint(t *testing.T) {
	var testCases = []int2StringTestCase{
		{
			in:  big.NewInt(0),
			out: ShipClassGalaxy,
		},
		{
			in:  big.NewInt(255),
			out: ShipClassGalaxy,
		},
		{
			in:  big.NewInt(256),
			out: ShipClassStar,
		},
		{
			in:  big.NewInt(65535),
			out: ShipClassStar,
		},
		{
			in:  big.NewInt(65536),
			out: ShipClassPlanet,
		},
		{
			in:  big.NewInt(14287616),
			out: ShipClassPlanet,
		},
		{
			in:  big.NewInt(14287617),
			out: ShipClassPlanet,
		},
		{
			in:  big.NewInt(14287618),
			out: ShipClassPlanet,
		},
		{
			in:  big.NewInt(14287619),
			out: ShipClassPlanet,
		},
		{
			in:  big.NewInt(4294967295),
			out: ShipClassPlanet,
		},
		{
			in:  big.NewInt(4294967296),
			out: ShipClassMoon,
		},
	}

	int2StringTestRunner(t, testCases, ClanPoint)
}

// TODO:
//
// string prepended / not with ~
