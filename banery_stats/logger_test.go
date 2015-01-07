package banery_stats

import (
  "testing"
  "bytes"
)

func TestLogger(t *testing.T) {
  out_buffer := &bytes.Buffer{}
  InitLoggerWriter(out_buffer)

  logger := Logger()
  logger("XXX")

  if (out_buffer.String() != "XXX\n") {
    t.Error("Expected output 'XXX\\n', got ", out_buffer.String())
  }
}
