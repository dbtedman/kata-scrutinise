package report

import (
	"encoding/csv"
	"github.com/dbtedman/kata-scrutinise/internal/domain/check"
	"github.com/dbtedman/kata-scrutinise/internal/gateway/logger"
	"io"
	"os"
	"strconv"
)

func WriteResultsToFile(filePath string, linkStatusList []check.LinkStatus) {
	file, _ := os.Create(filePath)

	if file != nil {
		defer func() {
			_ = file.Close()
		}()
	}

	WriteResultsToWriter(file, linkStatusList)

	logger.InfoFormat("Results written to file \"%s\"", filePath)
}

func WriteResultsToWriter(out io.Writer, linkStatusList []check.LinkStatus) {
	writer := csv.NewWriter(out)

	if writer != nil {
		defer func() {
			writer.Flush()
		}()
	}

	_ = writer.Write([]string{
		"URL",
		"Success",
		"Redirects",
		"Code",
	})

	for _, value := range linkStatusList {
		_ = writer.Write([]string{
			value.Url,
			strconv.FormatBool(value.Success),
			strconv.FormatBool(value.Redirects),
			strconv.FormatUint(uint64(value.Code), 10),
		})
	}
}
