package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var dataExportTests = []string{
	"123",
	"abc",
}

func TestDataExport(t *testing.T) {
	data := DataExport(dataExportTests)
	assert.Equal(t, dataImportTests, data)
}

var dataImportTests = `H4sIAAAAAAAA/4pWMjQyVtJRSkxKVooFBAAA//9iXM2zDQAAAA==`

func TestDataImport(t *testing.T) {
	data := DataImport(dataImportTests)
	assert.Equal(t, dataExportTests, data)
}
