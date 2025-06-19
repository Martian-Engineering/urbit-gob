# Spot Check Tool

This tool performs random spot checks to verify the bijectivity of the Urbit @p scrambler implementation.

## Usage

```bash
# Build the tool
go build -o spotcheck ./cmd/spotcheck

# Run with default settings (10,000 checks)
./spotcheck

# Run with custom number of checks
./spotcheck -n 100000

# Run with verbose output
./spotcheck -n 50000 -v

# Run with custom worker count
./spotcheck --num-checks 1000000 --workers 16
```

## Options

- `-n`, `--num-checks`: Number of random values to test (default: 10000)
- `-v`, `--verbose`: Show progress during testing
- `--workers`: Number of parallel workers (default: number of CPUs)

## What it tests

The tool verifies that:
1. No two different inputs produce the same output (no collisions)
2. Every value can be encrypted and then decrypted back to the original (round-trip property)

It tests values across different ranges:
- Small values (< 0x10000) - 10% of tests
- Low Feistel range (0x10000-0x1000000) - 30% of tests  
- Mid Feistel range (0x1000000-0x10000000) - 30% of tests
- High Feistel range (0x10000000-0x100000000) - 30% of tests

## Exit codes

- 0: All tests passed
- 1: Collisions or round-trip failures detected