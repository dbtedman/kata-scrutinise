package main

import (
	"flag"
	"github.com/dbtedman/kata-scrutinise/internal/domain/application"
	"github.com/dbtedman/kata-scrutinise/internal/domain/lang"
	"github.com/dbtedman/kata-scrutinise/internal/gateway/logger"
	"regexp"
)

func Run() {
	logger.ConfigureLogging()

	url := flag.String("url", "", lang.CliOptionUrl)
	outFile := flag.String("outFile", "results.csv", lang.CliOptionOutFile)

	flag.Parse()

	if *url == "" {
		logger.ErrorLine(lang.CliErrorMissingOptionUrl)
		return
	}

	const pattern = "^.+[.]csv$"
	match, _ := regexp.MatchString(pattern, *outFile)

	if !match {
		logger.ErrorFormat(lang.CliErrorInvalidOutFile, *outFile, pattern)
		return
	}

	config := application.ApplicationConfiguration{}
	config.BaseURL = url
	config.OutputFilePath = outFile

	application.Application(&config)
}
