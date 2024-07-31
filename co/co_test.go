package co

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

// test helpers
// s2s: string to string
// s2i: string to integer - TODO: is this needed?
// i2s: integer to string
// i2i: integer to integer
// a2s: any type to string

type s2sTestCase struct {
	in              string
	out             string
	expectedErrText string
}

type s2iTestCase struct {
	in              string
	out             *big.Int
	expectedErrText string
}

type i2sTestCase struct {
	in              *big.Int
	out             string
	expectedErrText string
}

type i2iTestCase struct {
	in              *big.Int
	out             *big.Int
	expectedErrText string
}

type a2sTestCase struct {
	in              interface{}
	out             string
	expectedErrText string
}

type s2sCoFn func(string) (string, error)
type s2iCoFn func(string) (*big.Int, error)
type i2sCoFn func(*big.Int) (string, error)
type i2iCoFn func(*big.Int) (*big.Int, error)
type a2sCoFn func(interface{}) (string, error)

func s2sTestRunner(t *testing.T, testCases []s2sTestCase, f s2sCoFn) {
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

func s2iTestRunner(t *testing.T, testCases []s2iTestCase, f s2iCoFn) {
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

func i2sTestRunner(t *testing.T, testCases []i2sTestCase, f i2sCoFn) {
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

func i2iTestRunner(t *testing.T, testCases []i2iTestCase, f i2iCoFn) {
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

func a2sTestRunner(t *testing.T, testCases []a2sTestCase, f a2sCoFn) {
	for _, tt := range testCases {
		t.Run(tt.in.(string), func(t *testing.T) {

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
	var testCases = []a2sTestCase{
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
			in:              "abcdefg",
			expectedErrText: "invalid integer string: abcdefg",
		},
	}

	a2sTestRunner(t, testCases, Patp)
}

func TestPatq(t *testing.T) {
	var testCases = []a2sTestCase{
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
			in:              "abcdefg",
			expectedErrText: "invalid integer string: abcdefg",
		},
	}

	a2sTestRunner(t, testCases, Patq)
}

func TestClan(t *testing.T) {
	var testCases = []s2sTestCase{
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

	s2sTestRunner(t, testCases, Clan)
}

func TestSein(t *testing.T) {
	var testCases = []s2sTestCase{
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

	s2sTestRunner(t, testCases, Sein)
}

func TestPatp2Dec(t *testing.T) {
	var testCases = []s2sTestCase{
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

	s2sTestRunner(t, testCases, Patp2Dec)
}

func TestPatq2Dec(t *testing.T) {
	var testCases = []s2sTestCase{
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

	s2sTestRunner(t, testCases, Patq2Dec)
}

func TestPatp2Hex(t *testing.T) {
	var testCases = []s2sTestCase{
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

	s2sTestRunner(t, testCases, Patp2Hex)
}

func TestPatq2Hex(t *testing.T) {
	var testCases = []s2sTestCase{
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

	s2sTestRunner(t, testCases, Patq2Hex)
}

func TestHex2Patp(t *testing.T) {
	var testCases = []s2sTestCase{
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

	s2sTestRunner(t, testCases, Hex2Patp)
}

func TestHex2Patq(t *testing.T) {
	var testCases = []s2sTestCase{
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

	s2sTestRunner(t, testCases, Hex2Patq)
}

func TestPatp2Point(t *testing.T) {
	var testCases = []s2iTestCase{
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

	s2iTestRunner(t, testCases, Patp2Point)
}

// Point2Patp
func TestPoint2Patp(t *testing.T) {
	var testCases = []i2sTestCase{
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

	i2sTestRunner(t, testCases, Point2Patp)
}

// Patq2Point
func TestPatq2Point(t *testing.T) {
	var testCases = []s2iTestCase{
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

	s2iTestRunner(t, testCases, Patq2Point)
}

// Point2Patq
func TestPoint2Patq(t *testing.T) {
	var testCases = []i2sTestCase{
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

	i2sTestRunner(t, testCases, Point2Patq)
}

// SeinPoint
func TestSeinPoint(t *testing.T) {
	var testCases = []i2iTestCase{
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

	i2iTestRunner(t, testCases, SeinPoint)
}

// ClanPoint
func TestClanPoint(t *testing.T) {
	var testCases = []i2sTestCase{
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

	i2sTestRunner(t, testCases, ClanPoint)
}

// TODO:
//
// Patp can handle string or int
// Patq can handle string or int
// string prepended / not with ~
