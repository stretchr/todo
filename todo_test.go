package to

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTODOExpires(t *testing.T) {
	expired := false
	printed := ""
	expire = func(msg string) {
		expired = true
		printed = fmt.Sprintf("%s: TODO Expired: - %s", callerInfo(), msg)
	}

	Do(time.Now().Format(shortForm), "Testing expiration!")
	assert.True(t, expired, "Do function should have expired!")
	assert.Equal(t, "todo_test.go:16: TODO Expired: - Testing expiration!", printed)
}

func TestTODONotExpires(t *testing.T) {
	expired := false
	expire = func(msg string) {
		expired = true
	}
	Do(time.Now().AddDate(0, 0, 1).Format(shortForm), "Testing expiration!")
	assert.False(t, expired, "Do function should not have expired!")
}

func TestTODOCacheDate(t *testing.T) {
	expired := false
	expire = func(msg string) {
		expired = true
	}

	now := time.Now()
	parsed, _ := time.Parse(shortForm, now.Format(shortForm))

	Do(now.Format(shortForm), "Testing expiration!")
	assert.True(t, expired, "Do function should have expired!")
	assert.Equal(t, times[now.Format(shortForm)], parsed)

	Do(now.Format(shortForm), "Testing expiration!")
	assert.True(t, expired, "Do function should have expired!")

}

func BenchmarkTODOCachedParse(b *testing.B) {

	b.StopTimer()

	benchTimes := map[string]time.Time{}
	by := time.Now().Format(shortForm)

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		var byDate time.Time
		ok := false
		var err error
		if byDate, ok = benchTimes[by]; !ok {
			byDate, err = time.Parse(shortForm, by)
			if err != nil {
				panic(fmt.Sprintf("Unable to parse Do date: %s", err))
			}
			benchTimes[by] = byDate
		}
	}
}

func BenchmarkTODOParseAlways(b *testing.B) {

	b.StopTimer()

	by := time.Now().Format(shortForm)

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		byDate, err := time.Parse(shortForm, by)
		if err != nil {
			panic(fmt.Sprintf("Unable to parse Do date: %s", err))
		}
		_ = byDate
	}
}
