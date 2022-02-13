package file

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"encoding/csv"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/masashi-toda/go-packfile/libs/log"
)

type reader struct {
	path             string
	separate         string
	comment          string
	lazyQuotes       bool
	trimLeadingSpace bool
	reuseRecord      bool
	isLTSV           bool
}

type scanListener func(Record)

func (r *reader) scan(listener scanListener) {
	src, err := os.Open(filepath.Clean(r.path))
	if err != nil {
		log.Panicf("failed to open file [%s]", r.path)
	}
	defer src.Close()

	var reader io.Reader = src
	if strings.HasSuffix(r.path, ".gz") {
		gzipReader, err := gzip.NewReader(reader)
		if err != nil {
			log.Panicf("failed to create gzip reader [%s]", err.Error())
		}
		defer gzipReader.Close()
		reader = gzipReader
	}

	headers := make([]string, 0)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		data := scanner.Bytes()
		if len(data) == 0 {
			continue
		}
		if bytes.HasPrefix(data, []byte(string(r.comment))) {
			continue
		}
		csvReader := csv.NewReader(bytes.NewBuffer(data))
		csvReader.Comma = []rune(r.separate)[0]
		csvReader.Comment = []rune(r.comment)[0]
		csvReader.LazyQuotes = r.lazyQuotes
		csvReader.TrimLeadingSpace = r.trimLeadingSpace
		csvReader.ReuseRecord = r.reuseRecord

		values, err := csvReader.Read()
		if err != nil {
			log.Panicf("failed to read csv file [%s]", err.Error())
		}
		if len(values) == 0 {
			continue
		}
		if len(headers) == 0 {
			for _, value := range values {
				if r.isLTSV {
					kv := strings.SplitN(value, ":", 2)
					value = kv[0]
				}
				headers = append(headers, value)
			}
			if !r.isLTSV {
				continue
			}
		}
		records := Record{
			headers: headers,
			values:  make([]string, 0),
		}
		for _, value := range values {
			if r.isLTSV {
				kv := strings.SplitN(value, ":", 2)
				value = kv[1]
			}
			records.values = append(records.values, value)
		}
		if len(records.headers) > 0 {
			log.Infof("%s", records.headers)
			listener(records)
		}
	}
}

func newReader(path string) *reader {
	return &reader{
		path:             path,
		separate:         GetCSVSeparator(path),
		comment:          "#",
		lazyQuotes:       true,
		trimLeadingSpace: true,
		reuseRecord:      true,
		isLTSV:           strings.Contains(path, "ltsv"),
	}
}
