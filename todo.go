package to

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

// expire is a separate function variable to enable testing
var expire = func(msg string) {
	fmt.Printf("%s TODO Expired: - %s\n", callerInfo(), msg)
	os.Exit(1)
}

// shortForm is the format string expected by time.Parse
const shortForm = "2006-Jan-02"

// enabledKey is the string for the environment variable that must be set to activate todo
const enabledKey = "TODO_ENABLED"

// times caches the parsed time.Time arguments so subsequent calls will be quicker.
var times = map[string]time.Time{}

// Do checks to see if the by time has passed.
// If it has, "file:line: TODO Expired - msg" is printed and the program exits.
var Do = func(by, msg string) {
	var byDate time.Time
	ok := false
	var err error
	if byDate, ok = times[by]; !ok {
		byDate, err = time.Parse(shortForm, by)
		if err != nil {
			panic(fmt.Sprintf("Unable to parse Do date: %s", err))
		}
		times[by] = byDate
	}
	if time.Now().After(byDate) {
		expire(msg)
	}
}

func init() {
	if os.Getenv(enabledKey) == "" {
		Do = func(by, msg string) {
			// this line intentionally left blank
		}
	}
}

// callerInfo returns a string containing the file and line number of the Do call
// that expired.
func callerInfo() string {

	file := ""
	line := 0
	ok := false

	for i := 0; ; i++ {
		_, file, line, ok = runtime.Caller(i)
		if !ok {
			return ""
		}
		parts := strings.Split(file, "/")
		dir := parts[len(parts)-2]
		file = parts[len(parts)-1]
		if dir != "todo" || file == "todo_test.go" {
			break
		}
	}

	return fmt.Sprintf("%s:%d", file, line)
}
