package ob

import (
	"math/big"
	"testing"
)

// TestEdgeCases tests specific edge cases and legacy bug compatibility
func TestEdgeCases(t *testing.T) {
	// Test cases where arr == a (0xFFFF) in the Feistel cipher
	// These trigger the legacy bug compatibility code
	testCases := []struct {
		name  string
		input string
		desc  string
	}{
		// Values that might trigger arr == a condition
		{"arr_equals_a_1", "4294901760", "k value (0xFFFF * 0x10000)"},
		{"arr_equals_a_2", "4294836224", "Near k value"},
		{"boundary_low", "65536", "0x10000 - start of Feistel range"},
		{"boundary_high", "4294967295", "0xFFFFFFFF - end of range"},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Encrypt
			encrypted, err := Fein(tc.input)
			if err != nil {
				t.Fatalf("Fein failed for %s (%s): %v", tc.input, tc.desc, err)
			}
			
			// Decrypt
			decrypted, err := Fynd(encrypted)
			if err != nil {
				t.Fatalf("Fynd failed for encrypted value of %s: %v", tc.input, err)
			}
			
			// Verify round-trip
			inputBig, _ := new(big.Int).SetString(tc.input, 10)
			if decrypted.Cmp(inputBig) != 0 {
				t.Errorf("Round-trip failed for %s (%s): got %s", tc.input, tc.desc, decrypted.String())
			}
		})
	}
}

// TestSmallValues verifies that values < 0x10000 are returned unchanged
func TestSmallValues(t *testing.T) {
	testVals := []string{"0", "1", "255", "256", "65535"}
	
	for _, val := range testVals {
		t.Run("small_"+val, func(t *testing.T) {
			encrypted, err := Fein(val)
			if err != nil {
				t.Fatalf("Fein failed for %s: %v", val, err)
			}
			
			// Small values should be unchanged
			valBig, _ := new(big.Int).SetString(val, 10)
			if encrypted.Cmp(valBig) != 0 {
				t.Errorf("Small value %s was changed to %s", val, encrypted.String())
			}
			
			// Round-trip should still work
			decrypted, err := Fynd(encrypted)
			if err != nil {
				t.Fatalf("Fynd failed for %s: %v", val, err)
			}
			
			if decrypted.Cmp(valBig) != 0 {
				t.Errorf("Round-trip failed for small value %s: got %s", val, decrypted.String())
			}
		})
	}
}

// TestLargeValues tests values in the extended range
func TestLargeValues(t *testing.T) {
	testVals := []struct {
		val  string
		desc string
	}{
		{"4294967296", "0x100000000 (just outside 32-bit)"},
		{"18446744073709551615", "Max uint64 (0xFFFFFFFFFFFFFFFF)"},
		{"18446744073709551616", "Beyond uint64"},
	}
	
	for _, tc := range testVals {
		t.Run("large_"+tc.val[:8], func(t *testing.T) {
			encrypted, err := Fein(tc.val)
			if err != nil {
				t.Fatalf("Fein failed for %s (%s): %v", tc.val, tc.desc, err)
			}
			
			valBig, _ := new(big.Int).SetString(tc.val, 10)
			
			// Values > 0xFFFFFFFFFFFFFFFF should be unchanged
			// Values in [0x100000000, 0xFFFFFFFFFFFFFFFF] will be processed
			beyondUint64, _ := new(big.Int).SetString("18446744073709551616", 10)
			if valBig.Cmp(beyondUint64) >= 0 {
				// Should be unchanged
				if encrypted.Cmp(valBig) != 0 {
					t.Errorf("Value beyond uint64 %s was changed to %s", tc.val, encrypted.String())
				}
			}
			
			// Regardless, round-trip should always work
			decrypted, err := Fynd(encrypted)
			if err != nil {
				t.Fatalf("Fynd failed for %s: %v", tc.val, err)
			}
			
			if decrypted.Cmp(valBig) != 0 {
				t.Errorf("Round-trip failed for %s (%s): got %s", tc.val, tc.desc, decrypted.String())
			}
		})
	}
}

// TestSpecificCollisionPrevention tests values that might collide without proper implementation
func TestSpecificCollisionPrevention(t *testing.T) {
	// Test a large sample of sequential values to ensure no collisions
	seen := make(map[string]string)
	start := big.NewInt(3108299000) // Near one of the known collision values
	
	for i := 0; i < 1000; i++ {
		val := new(big.Int).Add(start, big.NewInt(int64(i)))
		encrypted, err := Fein(val.String())
		if err != nil {
			t.Fatalf("Fein failed for %s: %v", val.String(), err)
		}
		
		encStr := encrypted.String()
		if original, exists := seen[encStr]; exists {
			t.Fatalf("COLLISION: %s and %s both encrypt to %s", original, val.String(), encStr)
		}
		seen[encStr] = val.String()
	}
	
	t.Logf("Tested 1000 sequential values starting from %s with no collisions", start.String())
}