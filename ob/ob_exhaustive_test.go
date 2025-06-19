package ob

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"runtime"
	"sync"
	"sync/atomic"
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

	numWorkers := runtime.NumCPU()
	
	for _, tr := range testRanges {
		t.Run(tr.name, func(t *testing.T) {
			// Use sync.Map for thread-safe collision detection
			seen := &sync.Map{} // encrypted -> original
			
			// Atomic counter for progress tracking
			var count atomic.Int64
			var collisions atomic.Int64
			
			// Calculate total work
			rangeSize := new(big.Int).Sub(tr.end, tr.start)
			
			// Work channel and wait group
			workChan := make(chan *big.Int, numWorkers*10)
			var wg sync.WaitGroup
			
			// Error channel for collecting failures
			errChan := make(chan error, numWorkers)
			
			// Start workers
			for w := 0; w < numWorkers; w++ {
				wg.Add(1)
				go func(workerID int) {
					defer wg.Done()
					
					for i := range workChan {
						// Encrypt
						encrypted, err := Fein(i.String())
						if err != nil {
							errChan <- fmt.Errorf("worker %d: Fein failed for %s: %v", workerID, i.String(), err)
							continue
						}
						
						encStr := encrypted.String()
						
						// Check for collisions
						if original, loaded := seen.LoadOrStore(encStr, i.String()); loaded {
							origStr := original.(string)
							if origStr != i.String() {
								collisions.Add(1)
								errChan <- fmt.Errorf("COLLISION DETECTED: %s and %s both encrypt to %s", origStr, i.String(), encStr)
							}
							// If it's the same input, that's fine (can happen due to range overlap)
							continue
						}
						
						// Verify round-trip
						decrypted, err := Fynd(encrypted)
						if err != nil {
							errChan <- fmt.Errorf("worker %d: Fynd failed for %s: %v", workerID, encrypted.String(), err)
							continue
						}
						
						if decrypted.Cmp(i) != 0 {
							errChan <- fmt.Errorf("worker %d: Round-trip failed: %s -> %s -> %s", workerID, i.String(), encrypted.String(), decrypted.String())
							continue
						}
						
						// Update progress
						currentCount := count.Add(1)
						if currentCount%1000000 == 0 {
							progress := float64(currentCount) / float64(rangeSize.Int64()) * 100
							fmt.Printf("Progress: tested %d values (%.1f%%)\n", currentCount, progress)
						}
					}
				}(w)
			}
			
			// Generate work
			go func() {
				for i := new(big.Int).Set(tr.start); i.Cmp(tr.end) < 0; i.Add(i, big.NewInt(1)) {
					workChan <- new(big.Int).Set(i) // Copy i to avoid race
				}
				close(workChan)
			}()
			
			// Collect errors in a separate goroutine
			var errors []error
			errDone := make(chan struct{})
			go func() {
				for err := range errChan {
					errors = append(errors, err)
					// Stop on first collision
					if collisions.Load() > 0 {
						break
					}
				}
				close(errDone)
			}()
			
			// Wait for all workers to finish
			wg.Wait()
			close(errChan)
			<-errDone
			
			// Check for errors
			if len(errors) > 0 {
				for i, err := range errors {
					t.Errorf("Error %d: %v", i+1, err)
					if i >= 10 && len(errors) > 11 {
						t.Errorf("... and %d more errors", len(errors)-11)
						break
					}
				}
				t.FailNow()
			}
			
			finalCount := count.Load()
			fmt.Printf("Successfully tested %d values in range %s-%s using %d workers\n", 
				finalCount, tr.start.String(), tr.end.String(), numWorkers)
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