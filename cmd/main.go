package main

import (
	"flag"
	"fmt"
	"math/big"
	"os"
	"strconv"

	"github.com/deelawn/urbit-gob/co"
)

const (
	// Commands
	cmdPatp       string = "patp"
	cmdPatp2Dec   string = "patp2dec"
	cmdPatp2Hex   string = "patp2hex"
	cmdPatp2Point string = "patp2point"

	cmdPatq       string = "patq"
	cmdPatq2Dec   string = "patq2dec"
	cmdPatq2Hex   string = "patq2hex"
	cmdPatq2Point string = "patq2point"

	cmdPoint2Patp string = "point2patp"
	cmdPoint2Patq string = "point2patq"

	cmdHex2Patp string = "hex2patp"
	cmdHex2Patq string = "hex2patq"

	cmdClan      string = "clan"
	cmdClanPoint string = "clanpoint"
	cmdSein      string = "sein"
	cmdSeinPoint string = "seinpoint"
	cmdEqPatq    string = "eqpatq"

	cmdIsValidPat  string = "isvalidpat"
	cmdIsValidPatp string = "isvalidpatp"
	cmdIsValidPatq string = "isvalidpatq"

	// Exit codes
	codeInsufficientArguments int = 1
	codeInvalidCommand        int = 2
	codeErrorReturned         int = 3

	// Others
	minArgLen        int    = 2
	usageCmdFmtStr   string = "    %-20s: %s\n"
	errInvalidCmdStr string = "invalid command: %s\n"
	errMissingEqArg  string = "missing second argument for equality check"
)

func main() {

	flag.Usage = func() {
		fmt.Printf("Usage: %s COMMAND args...\n\n", os.Args[0])
		fmt.Printf("Valid commands:\n\n")
		fmt.Printf(usageCmdFmtStr, cmdPatp, "converts a number to a @p-encoded string\n")
		fmt.Printf(usageCmdFmtStr, cmdPatp2Dec, "converts a @p-encoded string to a decimal-encoded string\n")
		fmt.Printf(usageCmdFmtStr, cmdPatp2Hex, "converts a @p-encoded string to a hex-encoded string\n")
		fmt.Printf(usageCmdFmtStr, cmdPatp2Point, "converts a @p-encoded string to a big.Int\n")
		fmt.Printf(usageCmdFmtStr, cmdPatq, "converts a number to a @q-encoded string\n")
		fmt.Printf(usageCmdFmtStr, cmdPatq2Dec, "converts a @q-encoded string to a decimal-encoded string\n")
		fmt.Printf(usageCmdFmtStr, cmdPatq2Hex, "converts a @q-encoded string to a hex-encoded string\n")
		fmt.Printf(usageCmdFmtStr, cmdPatq2Point, "converts a @q-encoded string to a big.Int\n")
		fmt.Printf(usageCmdFmtStr, cmdPoint2Patp, "converts a big.Int to a @p-encoded string\n")
		fmt.Printf(usageCmdFmtStr, cmdPoint2Patq, "converts a big.Int to a @q-encoded string\n")
		fmt.Printf(usageCmdFmtStr, cmdHex2Patp, "converts a hex-encoded string to a @p-encoded string\n")
		fmt.Printf(usageCmdFmtStr, cmdHex2Patq, "converts a hex-encoded string to a @q-encoded string\n")
		fmt.Printf(usageCmdFmtStr, cmdClan, "determines the ship class of a @p value\n")
		fmt.Printf(usageCmdFmtStr, cmdClanPoint, "determines the ship class of a big.Int\n")
		fmt.Printf(usageCmdFmtStr, cmdSein, "determines the parent of a @p value\n")
		fmt.Printf(usageCmdFmtStr, cmdSeinPoint, "determines the parent of a big.Int\n")
		fmt.Printf(usageCmdFmtStr, cmdEqPatq, "performs an equality comparison on @q values\n")
		fmt.Printf(usageCmdFmtStr, cmdIsValidPat, "weakly checks if a string is a valid @p or @q value\n")
		fmt.Printf(usageCmdFmtStr, cmdIsValidPatp, "validates a @p string\n")
		fmt.Printf(usageCmdFmtStr, cmdIsValidPatq, "validates a @q string\n")
	}

	flag.Parse()

	var (
		result interface{}
		err    error
	)

	args := flag.Args()
	if len(args) < minArgLen {
		flag.Usage()
		os.Exit(codeInsufficientArguments)
	}

	switch args[0] {

	case cmdPatp:
		result, err = co.Patp(args[1])
	case cmdPatp2Dec:
		result, err = co.Patp2Dec(args[1])
	case cmdPatp2Hex:
		result, err = co.Patp2Hex(args[1])
	case cmdPatp2Point:
		result, err = co.Patp2Point(args[1])
	case cmdPatq:
		result, err = co.Patq(args[1])
	case cmdPatq2Dec:
		result, err = co.Patq2Dec(args[1])
	case cmdPatq2Hex:
		result, err = co.Patq2Hex(args[1])
	case cmdPatq2Point:
		result, err = co.Patq2Point(args[1])
	case cmdPoint2Patp:
		i, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			fmt.Println(err)
			os.Exit(codeErrorReturned)
		}
		result, err = co.Point2Patp(big.NewInt(i))
	case cmdPoint2Patq:
		i, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			fmt.Println(err)
			os.Exit(codeErrorReturned)
		}
		result, err = co.Point2Patq(big.NewInt(i))
	case cmdHex2Patp:
		result, err = co.Hex2Patp(args[1])
	case cmdHex2Patq:
		result, err = co.Hex2Patq(args[1])
	case cmdClan:
		result, err = co.Clan(args[1])
	case cmdClanPoint:
		i, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			fmt.Println(err)
			os.Exit(codeErrorReturned)
		}
		result, err = co.ClanPoint(big.NewInt(i))
	case cmdSein:
		result, err = co.Sein(args[1])
	case cmdSeinPoint:
		i, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			fmt.Println(err)
			os.Exit(codeErrorReturned)
		}
		result, err = co.SeinPoint(big.NewInt(i))
	case cmdIsValidPat:
		result = co.IsValidPat(args[1])
	case cmdIsValidPatp:
		result = co.IsValidPatp(args[1])
	case cmdIsValidPatq:
		result = co.IsValidPatq(args[1])
	case cmdEqPatq:
		if len(args) < 3 {
			fmt.Println(errMissingEqArg)
			flag.Usage()
			os.Exit(codeInsufficientArguments)
		}
		result, err = co.EqPatq(args[1], args[2])
	default:
		fmt.Printf(errInvalidCmdStr, args[0])
		flag.Usage()
		os.Exit(codeInvalidCommand)
	}

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(codeErrorReturned)
	}

	fmt.Println(result)
}
