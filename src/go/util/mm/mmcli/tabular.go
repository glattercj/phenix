// Taken (almost) as-is from minimega/miniweb.

package mmcli

import (
	"phenix/util/plog"
	"strings"

	"github.com/activeshadow/libminimega/minicli"
)

type tabularToMapper func(*minicli.Response, []string) map[string]string

func tabularToMap(resp *minicli.Response, row []string) map[string]string {
	res := map[string]string{
		"host": resp.Host,
	}

	for i, header := range resp.Header {
		res[header] = row[i]
	}

	return res
}

func tabularToMapCols(columns []string) tabularToMapper {
	// create local copy of columns in case they get changed
	cols := make([]string, len(columns))
	copy(cols, columns)

	return func(resp *minicli.Response, row []string) map[string]string {
		res := map[string]string{}

		for _, column := range cols {
			if strings.Contains(column, "host") {
				res["host"] = resp.Host
				continue
			}

			for i, header := range resp.Header {
				if strings.Contains(column, header) {
					res[header] = row[i]
				}
			}
		}

		return res
	}
}

// RunTabular is used to run the given command when the response is expected to
// be in tabular form. A slice of maps is returned, with each map representing a
// row in the tabular response and each map key representing the column.
func RunTabular(cmd *Command) []map[string]string {
	// copy all fields in header order
	mapper := tabularToMap

	if len(cmd.Columns) > 0 {
		// replace mapper to only copy fields in column order
		mapper = tabularToMapCols(cmd.Columns)
	}

	res := []map[string]string{}

	for resps := range Run(cmd) {
		for _, resp := range resps.Resp {
			if resp.Error != "" {
				plog.Error(plog.TypeSystem, "error running mm cmd", "cmd", cmd.Command, "error", resp.Error)
				continue
			}

			for _, row := range resp.Tabular {
				res = append(res, mapper(resp, row))
			}
		}
	}

	return res
}
