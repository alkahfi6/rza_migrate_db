package output

import (
	"fmt"
	"io"
)

// PrintFailureDetails writes every validation failure in a deterministic,
// self-contained format suitable for local terminals and CI job logs. The JSON
// report remains canonical, but users do not need artifact access to diagnose a
// failed gate.
func PrintFailureDetails(w io.Writer, report Report) {
	items := report.Fail.Items
	if len(items) == 0 {
		return
	}
	fmt.Fprintf(w, "[fail-summary] total=%d\n", len(items))
	for index, item := range items {
		prefix := fmt.Sprintf("[fail %d/%d]", index+1, len(items))
		fmt.Fprintf(w, "%s rule=%s severity=%d file=%q", prefix, item.Rule, item.Severity, item.File)
		if item.Object != "" {
			fmt.Fprintf(w, " object=%q", item.Object)
		}
		if item.Location != "" {
			fmt.Fprintf(w, " location=%q", item.Location)
		}
		if item.Reason != "" {
			fmt.Fprintf(w, " reason=%s", item.Reason)
		}
		if item.Action != "" {
			fmt.Fprintf(w, " action=%s", item.Action)
		}
		fmt.Fprintf(w, " message=%q", item.Message)
		if item.Suggestion != "" {
			fmt.Fprintf(w, " suggestion=%q", item.Suggestion)
		}
		fmt.Fprintln(w)
	}
}
