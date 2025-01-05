package middleware

import (
	"bytes"
	"fmt"
	"net"
	"net/http"
	"os"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-nop/ginstarter/internal/utils/response"
	"github.com/go-nop/ginstarter/pkg/log"
)

// SkipFrames is the number of frames to skip when calling the stack function.
const SkipFrames = 3

// RecoveryOnPanic is a middleware that recovers from any panics and writes a 500 error.
func RecoveryOnPanic() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			err := recover()
			if err != nil {
				// check for a broken connection, as it is not really a condition that warrants a panic stack trace
				var brokenPipe bool
				if netErr, ok := err.(*net.OpError); ok {
					if sysErr, ok := netErr.Err.(*os.SyscallError); ok {
						if strings.EqualFold(sysErr.Error(), "broken pipe") || strings.Contains(strings.ToLower(sysErr.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				stacks := stackWithSkip(SkipFrames)

				if brokenPipe {
					log.Errorf("Broken pipe: %v", err)
				} else {
					log.Errorf("[Recovery] panic recovered: \n%v\n%v", err, stacks)
				}

				if brokenPipe {
					// If the connection is dead, we can't write a status to it
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
				} else {
					c.Status(http.StatusInternalServerError)
					response.ErrorResponse(c, http.StatusInternalServerError, "internal server error")
				}
			}
		}()
		c.Next()
	}
}

// stackWithSkip returns a formatted stack trace with skip frames skipped.
func stackWithSkip(skip int) string {
	buf := new(bytes.Buffer)

	for i := skip; ; i++ {
		pc, _, line, ok := runtime.Caller(i)
		if !ok {
			break
		}

		fmt.Fprintf(buf, "\t%s:%d\n", function(pc), line)
	}
	return buf.String()
}

// function returns the name of the function that called the function that called it.
func function(pc uintptr) []byte {
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return nil
	}

	name := []byte(fn.Name())
	return name
}
