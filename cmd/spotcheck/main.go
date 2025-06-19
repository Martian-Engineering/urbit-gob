package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"
	"runtime"
	"sync"
	"sync/atomic"
	"time"

	"github.com/deelawn/urbit-gob/ob"
)

func main() {
	var numChecks int
	var workers int
	var verbose bool
	
	flag.IntVar(&numChecks, "n", 10000, "Number of spot checks to perform")
	flag.IntVar(&numChecks, "num-checks", 10000, "Number of spot checks to perform (alias)")
	flag.IntVar(&workers, "workers", runtime.NumCPU(), "Number of parallel workers")
	flag.BoolVar(&verbose, "v", false, "Verbose output")
	flag.BoolVar(&verbose, "verbose", false, "Verbose output (alias)")
	
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [flags]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\nSpot-check the bijectivity of the Urbit @p scrambler\n\n")
		fmt.Fprintf(os.Stderr, "Flags:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExample:\n")
		fmt.Fprintf(os.Stderr, "  %s -n 100000 -v\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s --num-checks 1000000 --workers 8\n", os.Args[0])
	}
	
	flag.Parse()
	
	if numChecks <= 0 {
		log.Fatal("Number of checks must be positive")
	}
	
	fmt.Printf("Starting spot-check with %d random values using %d workers...\n", numChecks, workers)
	start := time.Now()
	
	// Define test ranges
	ranges := []struct {
		name     string
		start    *big.Int
		end      *big.Int
		fraction float64 // fraction of total tests
	}{
		{"small values (< 0x10000)", big.NewInt(0), big.NewInt(0x10000), 0.1},
		{"low feistel (0x10000-0x1000000)", big.NewInt(0x10000), big.NewInt(0x1000000), 0.3},
		{"mid feistel (0x1000000-0x10000000)", big.NewInt(0x1000000), big.NewInt(0x10000000), 0.3},
		{"high feistel (0x10000000-0x100000000)", big.NewInt(0x10000000), big.NewInt(0x100000000), 0.3},
	}
	
	// Channel for work items
	type workItem struct {
		value    *big.Int
		rangeIdx int
	}
	
	workChan := make(chan workItem, workers*2)
	resultChan := make(chan error, workers)
	
	// Shared map for collision detection (with mutex)
	var mu sync.Mutex
	seen := make(map[string]string)
	
	// Progress tracking
	var processed atomic.Int64
	var collisions atomic.Int64
	
	// Start workers
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			
			for item := range workChan {
				// Encrypt
				encrypted, err := ob.Fein(item.value.String())
				if err != nil {
					resultChan <- fmt.Errorf("worker %d: Fein failed for %s: %v", workerID, item.value.String(), err)
					continue
				}
				
				encStr := encrypted.String()
				
				// Check for collisions (with lock)
				mu.Lock()
				if original, exists := seen[encStr]; exists {
					// Only report if it's a real collision (different inputs)
					if original != item.value.String() {
						mu.Unlock()
						collisions.Add(1)
						resultChan <- fmt.Errorf("COLLISION: %s and %s both encrypt to %s", original, item.value.String(), encStr)
						continue
					}
				} else {
					seen[encStr] = item.value.String()
				}
				mu.Unlock()
				
				// Verify round-trip
				decrypted, err := ob.Fynd(encrypted)
				if err != nil {
					resultChan <- fmt.Errorf("worker %d: Fynd failed for %s: %v", workerID, encrypted.String(), err)
					continue
				}
				
				if decrypted.Cmp(item.value) != 0 {
					resultChan <- fmt.Errorf("worker %d: Round-trip failed: %s -> %s -> %s", workerID, item.value.String(), encrypted.String(), decrypted.String())
					continue
				}
				
				count := processed.Add(1)
				if verbose && count%1000 == 0 {
					fmt.Printf("Progress: %d/%d (%.1f%%)\n", count, numChecks, float64(count)/float64(numChecks)*100)
				}
			}
		}(i)
	}
	
	// Generate work items
	go func() {
		for idx, r := range ranges {
			numTests := int(float64(numChecks) * r.fraction)
			rangeBig := new(big.Int).Sub(r.end, r.start)
			
			if verbose {
				fmt.Printf("Generating %d values for range %s\n", numTests, r.name)
			}
			
			for i := 0; i < numTests; i++ {
				// Generate cryptographically random value in range
				randOffset, err := rand.Int(rand.Reader, rangeBig)
				if err != nil {
					log.Fatalf("Failed to generate random number: %v", err)
				}
				
				testVal := new(big.Int).Add(r.start, randOffset)
				workChan <- workItem{value: testVal, rangeIdx: idx}
			}
		}
		close(workChan)
	}()
	
	// Collect errors in a separate goroutine
	var errors []error
	errDone := make(chan struct{})
	go func() {
		for err := range resultChan {
			errors = append(errors, err)
		}
		close(errDone)
	}()
	
	// Wait for workers to finish
	wg.Wait()
	close(resultChan)
	<-errDone
	
	elapsed := time.Since(start)
	
	// Print results
	fmt.Printf("\n=== Spot Check Results ===\n")
	fmt.Printf("Total values tested: %d\n", processed.Load())
	fmt.Printf("Time elapsed: %v\n", elapsed)
	fmt.Printf("Values per second: %.0f\n", float64(processed.Load())/elapsed.Seconds())
	fmt.Printf("Collisions found: %d\n", collisions.Load())
	
	if len(errors) > 0 {
		fmt.Printf("\n!!! ERRORS FOUND: %d !!!\n", len(errors))
		for i, err := range errors {
			fmt.Printf("Error %d: %v\n", i+1, err)
			if i >= 10 && len(errors) > 11 {
				fmt.Printf("... and %d more errors\n", len(errors)-11)
				break
			}
		}
		os.Exit(1)
	} else {
		fmt.Printf("\n✓ All spot checks passed! No collisions or round-trip failures detected.\n")
	}
}