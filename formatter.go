package filtered

import (
	"regexp"
	"strings"

	"github.com/sirupsen/logrus"
)

// Formatter formats logs with filtered fields
type Formatter struct {
	formatter logrus.Formatter
	filter    *regexp.Regexp
}

// New returns an instance of logrus formatter
func New(fields []string, formatter logrus.Formatter) *Formatter {
	return &Formatter{
		formatter: formatter,
		filter:    regexp.MustCompile(`"(` + strings.Join(fields, "|") + `)":\s*"*(\d*|true|false|([^"]*)")"*(,\s*|\s*\n?\s*})`),
	}
}

// Format renders a single log entry
func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	data, err := f.formatter.Format(entry)
	data = f.filter.ReplaceAll(data, []byte(`"$1":[FILTERED]$4`))
	return data, err
}
