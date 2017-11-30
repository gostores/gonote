package notepad

import (
	"bytes"
	"testing"

	"github.com/gostores/require"
)

func TestNotepad(t *testing.T) {
	var logHandle, outHandle bytes.Buffer

	n := NewNotepad(LevelCritical, LevelError, &outHandle, &logHandle, "TestNotePad", 0)

	require.Equal(t, LevelCritical, n.GetStdoutThreshold())
	require.Equal(t, LevelError, n.GetLogThreshold())

	n.DEBUG.Println("Some debug")
	n.ERROR.Println("Some error")
	n.CRITICAL.Println("Some critical error")

	require.Contains(t, logHandle.String(), "[TestNotePad] ERROR Some error")
	require.NotContains(t, logHandle.String(), "Some debug")
	require.NotContains(t, outHandle.String(), "Some error")
	require.Contains(t, outHandle.String(), "Some critical error")

	require.Equal(t, n.LogCountForLevel(LevelError), uint64(1))
	require.Equal(t, n.LogCountForLevel(LevelDebug), uint64(1))
	require.Equal(t, n.LogCountForLevel(LevelTrace), uint64(0))
}

func TestThresholdString(t *testing.T) {
	require.Equal(t, LevelError.String(), "ERROR")
	require.Equal(t, LevelTrace.String(), "TRACE")
}

func BenchmarkLogPrintOnlyToCounter(b *testing.B) {
	var logHandle, outHandle bytes.Buffer
	n := NewNotepad(LevelCritical, LevelCritical, &outHandle, &logHandle, "TestNotePad", 0)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n.INFO.Print("Test")
	}
}
