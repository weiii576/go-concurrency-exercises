# Limit Service Time for Free-tier Users

Your video processing service has a freemium model. Everyone has 10
sec of free processing time on your service. After that, the
service will kill your process, unless you are a paid premium user.

Beginner Level: 10s max per request
Advanced Level: 10s max per user (accumulated)

# Beginner Level
10s max per request

## Wanted Result

```
Process 1 done.
Process 2 done.
Process 3 done.
Process 4 killed. (No quota left).
Process 5 done.
```

## Solution

Perform `process()` on another goroutine and using a channel named `done` to return the signal of `process()` finish.

Return false to kill the `process()` and

```go
func HandleRequest(process func(), u *User) bool {
	if !u.IsPremium {
		done := make(chan int)

		go func() {
			process()
			close(done)
		}()

		select {
		case <-done:
			return true
		case <-time.After(time.Second * 10):
			return false
		}
	}

	return true
}
```

### Result
```
UserID: 0 	Process 1 started.
UserID: 1 	Process 2 started.
UserID: 1 	Process 2 done. <--
UserID: 0 	Process 3 started.
UserID: 1 	Process 5 started.
UserID: 1 	Process 5 done. <--
UserID: 0 	Process 4 started.
UserID: 0 	Process 1 done. <--
UserID: 0 	Process 3 done. <--
UserID: 0 	Process 4 killed. (No quota left) <--
```





