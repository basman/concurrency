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

# Issue 3: Pass the logger within a context

In many enterprise projects your packages will need a logger to properly generate log messages,
i.e. with a timestamp. If each package creates its own logger, the messages will not only look
different. They also might end up at different places.

Create a logger at only one place, which is `main()`, using the custom log package of this
project.

Use the `log.ContextWithValue()` method to enrich your context with that logger and pass the
context it returns to the `Run()` methods. They can then use `FromContext()` to get hold of
the logger and use it for any messages.

In this example constructors can not fail. But if they do in your projects, have them return 
errors, so `main()` can fail immediately.