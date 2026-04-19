package main

// Export/Import of data in the form of string arrays.
// This can be used to unload the conditions of a problem (input)
// in a compressed form into the debug console and unpack it in the IDE.

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
)

// DataExport serializes and compresses a slice of strings,
// returning a base64 encoded string.
func DataExport(data []string) string {
	jsonData, _ := json.Marshal(data)
	var gzBuf bytes.Buffer
	gz := gzip.NewWriter(&gzBuf)
	_, _ = gz.Write(jsonData)
	_ = gz.Close()

	return base64.StdEncoding.EncodeToString(gzBuf.Bytes())
}

// DataImport decodes a base64 string, decompresses it,
// and deserializes the JSON data into a slice of strings.
func DataImport(encodedData string) []string {
	gzData, _ := base64.StdEncoding.DecodeString(encodedData)
	gz, _ := gzip.NewReader(bytes.NewBuffer(gzData))
	_ = gz.Close()
	var jsonData bytes.Buffer
	_, _ = jsonData.ReadFrom(gz)
	var data []string
	_ = json.Unmarshal(jsonData.Bytes(), &data)

	return data
}
