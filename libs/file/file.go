package file

import (
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/masashi-toda/go-packfile/libs/log"
)

type fileFilter func(string) bool

func WalkAndScan(baseDir string, filter fileFilter, listener scanListener) {
	log.Infof("start walk and scan... [%s]", baseDir)
	filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return err
		}
		var fileName = info.Name()
		if !filter(fileName) {
			return err
		}
		Scan(path, listener)
		return nil
	})
}

func Scan(path string, listener scanListener) {
	log.Infof("scan file [%s]", path)
	newReader(path).scan(listener)
}

func UsingWriter(output string, writeFunc func(Writer)) {
	src, err := os.Create(output)
	if err != nil {
		log.Panicf("failed to open file [%s] %s", output, err.Error())
	}
	defer src.Close()

	var out io.Writer = src
	if strings.HasSuffix(output, ".gz") {
		gzipWriter := gzip.NewWriter(src)
		defer gzipWriter.Close()
		out = gzipWriter
	}

	writeFunc(&writer{internal: out})
}

func GetCSVSeparator(path string) string {
	sep := ","
	if strings.Contains(path, "tsv") {
		sep = "\t"
	}
	return sep
}
