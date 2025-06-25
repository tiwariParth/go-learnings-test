# Go Concurrency MCQ Test

Answer the following questions to test your understanding of Go concurrency.

## Goroutines Basics

1. What is a goroutine?
   - [ ] A. A thread managed by the operating system
   - [ ] B. A lightweight thread managed by the Go runtime
   - [ ] C. A synchronization primitive in Go
   - [ ] D. A function that always runs in the background

2. How do you start a goroutine?
   - [ ] A. By calling the `thread()` function
   - [ ] B. By prefixing a function call with the `go` keyword
   - [ ] C. By importing the `goroutine` package
   - [ ] D. By calling the `async` function

3. What happens if the main function finishes execution while goroutines are still running?
   - [ ] A. The program waits for all goroutines to complete
   - [ ] B. The goroutines continue running in the background
   - [ ] C. The program terminates and all running goroutines are killed
   - [ ] D. The Go runtime throws an error

## Channels

4. What is the primary purpose of channels in Go?
   - [ ] A. To store and retrieve data like arrays
   - [ ] B. To communicate between goroutines
   - [ ] C. To format output to the console
   - [ ] D. To handle errors

5. How do you create an unbuffered channel that carries integer values?
   - [ ] A. `ch = new(chan int)`
   - [ ] B. `ch := make(chan int)`
   - [ ] C. `ch := channel(int)`
   - [ ] D. `ch := int channel()`

6. What happens when you send a value to an unbuffered channel with no receiver?
   - [ ] A. The value is queued until a receiver is available
   - [ ] B. The value is discarded
   - [ ] C. The sending goroutine blocks until a receiver is available
   - [ ] D. An error is thrown

7. What is the difference between a buffered and unbuffered channel?
   - [ ] A. Buffered channels are faster but less safe
   - [ ] B. Buffered channels can store multiple values before blocking
   - [ ] C. Unbuffered channels use less memory
   - [ ] D. Buffered channels automatically handle errors

8. How do you close a channel?
   - [ ] A. `ch.close()`
   - [ ] B. `close(ch)`
   - [ ] C. `ch.Close()`
   - [ ] D. `break ch`

## Select Statement

9. What does the `select` statement do in Go?
   - [ ] A. Executes only one case that's true, like an if-else statement
   - [ ] B. Waits for multiple channel operations and proceeds with the first one that's ready
   - [ ] C. Selects random goroutines to execute
   - [ ] D. Filters values from channels

10. What happens if no case in a `select` statement is ready and there is a `default` case?
    - [ ] A. The select statement blocks until a case is ready
    - [ ] B. The default case executes immediately
    - [ ] C. The select statement returns an error
    - [ ] D. The program exits

## Synchronization

11. What is the purpose of `sync.WaitGroup`?
    - [ ] A. To wait for a collection of goroutines to finish execution
    - [ ] B. To create a group of worker goroutines
    - [ ] C. To synchronize access to shared resources
    - [ ] D. To group channels together

12. How do you increment a `WaitGroup` counter?
    - [ ] A. `wg.Inc()`
    - [ ] B. `wg.Add(1)`
    - [ ] C. `wg++`
    - [ ] D. `wg.Increment()`

13. What is the purpose of `sync.Mutex`?
    - [ ] A. To prevent race conditions by ensuring only one goroutine accesses shared data at a time
    - [ ] B. To create mutual dependencies between goroutines
    - [ ] C. To block all goroutines at once
    - [ ] D. To count the number of running goroutines

14. Which method acquires a lock on a mutex?
    - [ ] A. `mutex.Acquire()`
    - [ ] B. `mutex.Lock()`
    - [ ] C. `mutex.Get()`
    - [ ] D. `mutex.Hold()`

## Concurrency Patterns

15. What is the fan-out/fan-in pattern?
    - [ ] A. A pattern where one goroutine distributes work to multiple goroutines and collects the results
    - [ ] B. A pattern for handling errors in concurrent code
    - [ ] C. A pattern for cancelling goroutines
    - [ ] D. A pattern for scheduling goroutines

16. What is a worker pool?
    - [ ] A. A group of goroutines that process jobs from a shared source
    - [ ] B. A pool of memory reserved for goroutines
    - [ ] C. A technique to limit the number of open files
    - [ ] D. A collection of channels

17. What is a race condition?
    - [ ] A. When two goroutines compete to finish first
    - [ ] B. When the output of a program depends on the timing of uncontrollable events
    - [ ] C. When a channel receives data too quickly
    - [ ] D. When a goroutine uses too much CPU

## Context Package

18. What is the main purpose of the `context` package?
    - [ ] A. To provide global variables across goroutines
    - [ ] B. To carry deadlines, cancellation signals, and other request-scoped values across API boundaries
    - [ ] C. To log execution context for debugging
    - [ ] D. To manage goroutine execution order

19. How do you create a context with a timeout?
    - [ ] A. `ctx, cancel := context.WithTimeout(parentCtx, duration)`
    - [ ] B. `ctx := context.Timeout(duration)`
    - [ ] C. `ctx := context.NewWithTimeout(duration)`
    - [ ] D. `ctx, _ := context.SetTimeout(parentCtx, duration)`

20. What should you do with the cancel function returned when creating a cancellable context?
    - [ ] A. Ignore it if you don't need cancellation
    - [ ] B. Always call it when you're done with the context, typically with defer
    - [ ] C. Pass it to all functions that use the context
    - [ ] D. Return it from your function

## Live Coding Simulation

For the live coding portion, implement a simple concurrent web scraper that:
1. Takes a list of URLs as input
2. Fetches each URL concurrently with appropriate error handling
3. Uses a worker pool to limit the number of concurrent requests
4. Collects and displays the results with response time metrics
5. Implements a timeout for each request
