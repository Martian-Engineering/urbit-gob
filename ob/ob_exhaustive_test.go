package ob

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"testing"
)

// TestExhaustiveBijectivity verifies that the scrambler is fully bijective
// by checking that every input maps to a unique output and can be reversed.
// WARNING: This test is extremely expensive and should not be run in normal CI.
// Use TestSpotCheckBijectivity for regular testing.
func TestExhaustiveBijectivity(t *testing.T) {
	t.Skip("Skipping exhaustive test - use TestSpotCheckBijectivity or run explicitly with -run=TestExhaustiveBijectivity")

	// Test ranges based on the Urbit @p system
	testRanges := []struct {
		name  string
		start *big.Int
		end   *big.Int
	}{
		// Test small values that don't go through Feistel
		{"small values", big.NewInt(0), big.NewInt(0x10000)},
		// Test the full Feistel cipher range
		{"feistel range", big.NewInt(0x10000), big.NewInt(0x100000000)},
	}

	for _, tr := range testRanges {
		t.Run(tr.name, func(t *testing.T) {
			seen := make(map[string]string) // encrypted -> original
			count := big.NewInt(0)
			
			for i := new(big.Int).Set(tr.start); i.Cmp(tr.end) < 0; i.Add(i, big.NewInt(1)) {
				// Encrypt
				encrypted, err := Fein(i.String())
				if err != nil {
					t.Fatalf("Fein failed for %s: %v", i.String(), err)
				}
				
				encStr := encrypted.String()
				
				// Check for collisions
				if original, exists := seen[encStr]; exists {
					t.Fatalf("COLLISION DETECTED: %s and %s both encrypt to %s", original, i.String(), encStr)
				}
				seen[encStr] = i.String()
				
				// Verify round-trip
				decrypted, err := Fynd(encrypted)
				if err != nil {
					t.Fatalf("Fynd failed for %s: %v", encrypted.String(), err)
				}
				
				if decrypted.Cmp(i) != 0 {
					t.Fatalf("Round-trip failed: %s -> %s -> %s", i.String(), encrypted.String(), decrypted.String())
				}
				
				count.Add(count, big.NewInt(1))
				if count.Mod(count, big.NewInt(1000000)).Cmp(big.NewInt(0)) == 0 {
					fmt.Printf("Progress: tested %s values\n", count.String())
				}
			}
			
			fmt.Printf("Successfully tested %s values in range %s-%s\n", count.String(), tr.start.String(), tr.end.String())
		})
	}
}

// TestKnownCollisions tests the specific collision cases from issue #1105
func TestKnownCollisions(t *testing.T) {
	// These were the problematic values from the original implementation
	collisions := []struct {
		val1 string
		val2 string
		name string // What they incorrectly mapped to
	}{
		{"3108299008", "479733505", "morlyd-mogmev"},
		{"145391618", "1859915444", "fipfes-hossev/hossev-roppec"},
	}
	
	for _, c := range collisions {
		t.Run(fmt.Sprintf("collision_%s_vs_%s", c.val1, c.val2), func(t *testing.T) {
			// Encrypt both values
			enc1, err := Fein(c.val1)
			if err != nil {
				t.Fatalf("Failed to encrypt %s: %v", c.val1, err)
			}
			
			enc2, err := Fein(c.val2)
			if err != nil {
				t.Fatalf("Failed to encrypt %s: %v", c.val2, err)
			}
			
			// They should NOT be equal
			if enc1.Cmp(enc2) == 0 {
				t.Errorf("Values %s and %s both encrypt to %s (collision!)", c.val1, c.val2, enc1.String())
			}
			
			// Verify round trips
			dec1, err := Fynd(enc1)
			if err != nil {
				t.Fatalf("Failed to decrypt %s: %v", enc1.String(), err)
			}
			
			dec2, err := Fynd(enc2)
			if err != nil {
				t.Fatalf("Failed to decrypt %s: %v", enc2.String(), err)
			}
			
			val1Big, _ := new(big.Int).SetString(c.val1, 10)
			val2Big, _ := new(big.Int).SetString(c.val2, 10)
			
			if dec1.Cmp(val1Big) != 0 {
				t.Errorf("Round trip failed for %s: got %s", c.val1, dec1.String())
			}
			
			if dec2.Cmp(val2Big) != 0 {
				t.Errorf("Round trip failed for %s: got %s", c.val2, dec2.String())
			}
		})
	}
}

// TestSpotCheckBijectivity performs random spot checks for bijectivity
func TestSpotCheckBijectivity(t *testing.T) {
	// Default to 10000 spot checks
	numChecks := 10000
	
	// Can be overridden by build flag or environment variable
	// This will be implemented in the separate binary
	
	seen := make(map[string]string)
	
	// Test various ranges
	ranges := []struct {
		start    *big.Int
		end      *big.Int
		numTests int
	}{
		// Small values
		{big.NewInt(0), big.NewInt(0x10000), numChecks / 4},
		// Feistel range - lower part
		{big.NewInt(0x10000), big.NewInt(0x1000000), numChecks / 4},
		// Feistel range - middle part  
		{big.NewInt(0x1000000), big.NewInt(0x10000000), numChecks / 4},
		// Feistel range - upper part
		{big.NewInt(0x10000000), big.NewInt(0x100000000), numChecks / 4},
	}
	
	for _, r := range ranges {
		rangeBig := new(big.Int).Sub(r.end, r.start)
		
		for i := 0; i < r.numTests; i++ {
			// Generate random value in range
			randOffset, _ := randBigInt(rangeBig)
			testVal := new(big.Int).Add(r.start, randOffset)
			
			// Encrypt
			encrypted, err := Fein(testVal.String())
			if err != nil {
				t.Fatalf("Fein failed for %s: %v", testVal.String(), err)
			}
			
			encStr := encrypted.String()
			
			// Check for collisions
			if original, exists := seen[encStr]; exists {
				// If it's a real collision (not the same input)
				if original != testVal.String() {
					t.Fatalf("COLLISION: %s and %s both encrypt to %s", original, testVal.String(), encStr)
				}
			} else {
				seen[encStr] = testVal.String()
			}
			
			// Verify round-trip
			decrypted, err := Fynd(encrypted)
			if err != nil {
				t.Fatalf("Fynd failed for %s: %v", encrypted.String(), err)
			}
			
			if decrypted.Cmp(testVal) != 0 {
				t.Fatalf("Round-trip failed: %s -> %s -> %s", testVal.String(), encrypted.String(), decrypted.String())
			}
		}
	}
	
	t.Logf("Successfully spot-checked %d values across all ranges", numChecks)
}

// Helper function to generate random big.Int in range [0, max)
func randBigInt(max *big.Int) (*big.Int, error) {
	return rand.Int(rand.Reader, max)
}