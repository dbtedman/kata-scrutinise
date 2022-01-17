package report_test

import (
	"bytes"
	"fmt"
	"github.com/dbtedman/kata-scrutinise/internal/domain/check"
	"github.com/dbtedman/kata-scrutinise/internal/domain/report"
	"testing"
)

func TestWriteResultsToWriter(t *testing.T) {
	t.Run("handles empty link statuses", func(t *testing.T) {
		buffer := bytes.Buffer{}
		linkStatuses := make([]check.LinkStatus, 0)

		report.WriteResultsToWriter(&buffer, linkStatuses)

		got := buffer.String()
		want := "URL,Success,Redirects,Code\n"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("handles a single link status", func(t *testing.T) {
		buffer := bytes.Buffer{}
		linkStatuses := make([]check.LinkStatus, 0)

		exampleLinkStatus := check.LinkStatus{
			Url:       "https://example.com",
			Success:   false,
			Redirects: true,
			Code:      5,
		}

		linkStatuses = append(linkStatuses, exampleLinkStatus)

		report.WriteResultsToWriter(&buffer, linkStatuses)

		got := buffer.String()
		want := fmt.Sprintf(
			"URL,Success,Redirects,Code\n%s,%t,%t,%d\n",
			exampleLinkStatus.Url,
			exampleLinkStatus.Success,
			exampleLinkStatus.Redirects,
			exampleLinkStatus.Code,
		)

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("handles many link statuses", func(t *testing.T) {
		t.Skip()
	})

	t.Run("handles error in writer", func(t *testing.T) {
		t.Skip()
	})
}
