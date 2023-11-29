# Issue 1: avoid unreliable sleep and introduce proper synchronisation

Similar to this exercise, we found sleeps in the main goroutine in production code in an attempt
to make the main goroutine wait before it terminates the program. This is bad practice.
The system load might be too high for the goroutines to complete their shutdown after receiving
the cancellation signal.

Find a better way to synchronise termination and remove the sleep in main().

# Issue 2: Introduce context for proper cancellation

[This article](https://www.digitalocean.com/community/tutorials/how-to-use-contexts-in-go) gives 
you an introduction how to use contexts in go.

Use the proper method to signal termination to the producers. You will have to change the Run() 
function signature in the producers.