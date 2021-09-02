package resultsCollector

import (
	"log"
	"sync"
	"time"
)

type result struct{}

func Run() {
	resultsGenerators := []func() []result{
		func() []result { return []result{{} } },          // generator1 returns 1 result
		func() []result { return []result{{},{} } },       // generator2 returns 2 result
		func() []result { return []result{{},{},{} } },    // generator3 returns 3 result
		func() []result { return []result{{},{},{},{} } }, // generator4 returns 4 result
	}

	allResults := collectResults5(resultsGenerators)
	log.Printf("combined result: %v", allResults)
}

// New purpose is to collect ALL result from ALL concurrent resultsGenerators: Expect the final result slice of length 10 = (1 + 2 + 3 + 4)
// Ideas:
//   Maybe a wait group would be useful ... somehow. (Luke)
//   Maybe use a buffered channel for the resultsStream. (Kingsuk)
// Result: We ran into a deadlock with this function.
func collectResults0(resultsGenerators []func() []result) []result {
	resultsStream := make(chan []result)

	for _, generator := range resultsGenerators {
		go func(generateResults func() []result) {
			select {
			case resultsStream <- generateResults():
			case <- time.After(1*time.Second):
			}
		}(generator)
	}

	var allResults []result
	for results := range resultsStream {
		allResults = append(allResults, results...)
	}

	return allResults
}

// Uses that wait group we talked about above
// Result: No deadlocks or race condition, but no results either.
func collectResults1(resultsGenerators []func() []result) []result {
	resultsStream := make(chan []result)

	var wg sync.WaitGroup
	wg.Add(len(resultsGenerators))

	for _, generator := range resultsGenerators {
		go func(generateResults func() []result) {
			defer wg.Done()

			select {
			case resultsStream <- generateResults():
			case <- time.After(1*time.Second):
			}
		}(generator)
	}

	wg.Wait() // blocking: block until the for-loop is over
	close(resultsStream)

	var allResults []result
	for results := range resultsStream {
		allResults = append(allResults, results...)
	}

	return allResults
}

// Uses the buffered channel we talked about above
// Result: So far, so good. This one is working :)
func collectResults2(resultsGenerators []func() []result) []result {
	resultsStream := make(chan []result, len(resultsGenerators))

	var wg sync.WaitGroup
	wg.Add(len(resultsGenerators))

	for _, generator := range resultsGenerators {
		go func(generateResults func() []result) {
			defer wg.Done()

			select {
			case resultsStream <- generateResults():
			case <- time.After(1*time.Second):
			}
		}(generator)
	}

	wg.Wait()
	close(resultsStream)

	var allResults []result
	for results := range resultsStream {
		allResults = append(allResults, results...)
	}

	return allResults
}

// This one has a data race: You get varied results if you run it a few times
func collectResults3a(resultsGenerators []func() []result) []result {
	resultsStream := make(chan []result)

	var wg sync.WaitGroup
	wg.Add(len(resultsGenerators))

	for _, generator := range resultsGenerators {
		go func(generateResults func() []result) {
      defer wg.Done()
      
			select {
			case resultsStream <- generateResults():
			case <- time.After(1*time.Second):
			}
		}(generator)
	}

	var allResults []result
	go func() {
		for results := range resultsStream {
			allResults = append(allResults, results...)
		}
	}()

	wg.Wait()
	close(resultsStream)

	return allResults
}

// This one fixes the data race from collectResults3a
func collectResults3b(resultsGenerators []func() []result) []result {
	resultsStream := make(chan []result)

	var wg sync.WaitGroup
	wg.Add(len(resultsGenerators))

	for _, generator := range resultsGenerators {
		go func(generateResults func() []result) {
			select {
			case resultsStream <- generateResults():
			case <- time.After(1*time.Second):
			}
		}(generator)
	}

	var allResults []result
	go func() {
		for results := range resultsStream {
			allResults = append(allResults, results...)
			wg.Done()
		}
	}()

	wg.Wait()
	close(resultsStream)

	return allResults
}

// Looks okay at first but will deadlock if you decrease the timeout interval enough `(1*time.Nanosecond)`
func collectResults4(resultsGenerators []func() []result) []result {
	resultsStream := make(chan []result)

	for _, generator := range resultsGenerators {
		go func(generateResults func() []result) {
			select {
			case resultsStream <- generateResults():
			case <-time.After(1*time.Second):
			}
		}(generator)
	}

	var allResults []result
	for i := 0; i < len(resultsGenerators); i++ {           // tempted to try a for-select, but don't need the open-ended iteration here
		allResults = append(allResults, <-resultsStream...)   // don't need a select stmt, because only one 'case'
	}

	return allResults
}

// This one works
func collectResults5(resultsGenerators []func() []result) []result {
	resultsStream := make(chan []result, len(resultsGenerators))

	for _, generator := range resultsGenerators {
		go func(generateResults func() []result) {
			select {
			case resultsStream <- generateResults():
			case <-time.After(1*time.Nanosecond):
			}
		}(generator)
	}

	var allResults []result
	for i := 0; i < len(resultsGenerators); i++ {
		allResults = append(allResults, <-resultsStream...)
	}

	return allResults
}
