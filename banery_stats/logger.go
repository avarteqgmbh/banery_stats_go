package banery_stats

import (
	"log"
	"os"
  "io"
)

var out io.Writer = os.Stdout

func Logger() func(string) {
	stdlogger := log.New(out, "", 0)

	return func(msg string) {
		stdlogger.Printf(msg)
	}
}

func InitLoggerWriter(writer io.Writer) {
  out = writer
}
