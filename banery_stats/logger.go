package banery_stats

import (
	"log"
	"os"
)

func Logger() func(string) {
	stdlogger := log.New(os.Stdout, "", 0)

	return func(msg string) {
		stdlogger.Printf(msg)
	}
}
