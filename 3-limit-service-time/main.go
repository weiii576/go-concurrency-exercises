//////////////////////////////////////////////////////////////////////
//
// Your video processing service has a freemium model. Everyone has 10
// sec of free processing time on your service. After that, the
// service will kill your process, unless you are a paid premium user.
//
// Beginner Level: 10s max per request
// Advanced Level: 10s max per user (accumulated)
//

package main

import (
	"sync"
	"time"
)

// User defines the UserModel. Use this to check whether a User is a
// Premium user or not
type User struct {
	ID        int
	IsPremium bool
	TimeUsed  int64 // in seconds
	Mutex     sync.Mutex
}

// HandleRequest runs the processes requested by users. Returns false
// if process had to be killed
func HandleRequest(process func(), u *User) bool {
	if !u.IsPremium {
		done := make(chan int)
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()

		go func() {
			process()
			close(done)
		}()

		for {
			select {
			case <-done:
				return true
			case <-ticker.C:
				if u.TimeUsed >= 10 {
					return false
				}

				u.Mutex.Lock()
				u.TimeUsed++
				u.Mutex.Unlock()
			}
		}
	}

	process()
	return true
}

func main() {
	RunMockServer()
}
