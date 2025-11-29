# tee-pattern
## Key features in this implementation:
* Context-aware cancellation for graceful shutdown
* Concurrent sending to prevent slow consumers from blocking others
* Buffered channels option for better performance
* Proper resource cleanup with deferred channel closing