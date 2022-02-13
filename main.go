package main

import (
	"flag"
	"strings"
	"time"

	"github.com/masashi-toda/go-packfile/libs/file"
	"github.com/masashi-toda/go-packfile/libs/log"
)

func main() {
	log.Info("Start go-packfile...")
	var (
		begin        = time.Now()
		flgBase      = flag.String("base", "", "file scan base directory")
		flgFileNames = flag.String("filter", "", "filter file names")
		flgHeaders   = flag.String("header", "", "filter header names")
		flgOutput    = flag.String("out", "./result.csv.gz", "output file path")
	)
	flag.Parse()

	var (
		baseDir       = *flgBase
		filterNames   = strings.Split(*flgFileNames, ",")
		targetHeaders = strings.Split(*flgHeaders, ",")
		output        = *flgOutput
	)
	// setup file filter
	fileFilter := func(path string) bool {
		for _, name := range filterNames {
			if strings.Contains(path, name) {
				return true
			}
		}
		return false
	}

	file.UsingWriter(output, func(writer file.Writer) {
		var (
			outputSep     = file.GetCSVSeparator(output)
			outputHeaders = targetHeaders
			isWriteHeader = false
		)
		// scan all files
		file.WalkAndScan(baseDir, fileFilter, func(records file.Record) {
			if len(outputHeaders) == 0 {
				outputHeaders = records.Headers()
			}
			if !isWriteHeader {
				writer.WriteStrings(outputHeaders, outputSep).WriteNewLine()
				isWriteHeader = true
			}
			writer.WriteStrings(records.TargetValues(outputHeaders...), outputSep).WriteNewLine()
		})
	})
	log.Infof("Finished. [%s]", time.Now().Sub(begin))
}
