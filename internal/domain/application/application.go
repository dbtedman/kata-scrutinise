package application

import (
	"github.com/dbtedman/kata-scrutinise/internal/domain/check"
	"github.com/dbtedman/kata-scrutinise/internal/domain/fetch"
	"github.com/dbtedman/kata-scrutinise/internal/domain/report"
	"net/http"
)

type ApplicationConfiguration struct {
	BaseURL        *string
	OutputFilePath *string
}

func Application(config *ApplicationConfiguration) {
	urls := fetch.FetchLinksFromPage(*config.BaseURL, http.DefaultClient)

	linkStatusList := check.CheckAllLinkStatus(urls)

	report.WriteResultsToFile(*config.OutputFilePath, linkStatusList)
}
