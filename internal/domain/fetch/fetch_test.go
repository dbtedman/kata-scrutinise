package fetch_test

import (
	"github.com/dbtedman/kata-scrutinise/internal/domain/fetch"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFetchLinksFromPage(t *testing.T) {
	t.Skip("TestFetchLinksFromPage")
}

func TestMakeURLAbsolute(t *testing.T) {

	baseURL := "https://example.com"
	firstPath := "/something/other"
	secondPath := "other/something"
	fullPath := "https://example.com/abcdef"

	assert.Equal(t, baseURL+firstPath, fetch.MakeURLAbsolute(firstPath, baseURL))
	assert.Equal(t, baseURL+"/"+secondPath, fetch.MakeURLAbsolute(secondPath, baseURL))
	assert.Equal(t, fullPath, fetch.MakeURLAbsolute(fullPath, baseURL))
}
