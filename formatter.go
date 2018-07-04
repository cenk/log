package log

import "fmt"

// Formatter formats a record.
type Formatter interface {
	// Format the record and return a message.
	Format(*Record) (message string)
}

type defaultFormatter struct{}

// Format outputs a message like "2014-02-28 18:15:57 [example] INFO     something happened"
func (f defaultFormatter) Format(rec *Record) string {
	return fmt.Sprintf("%s [%s] %-8s %s", fmt.Sprint(rec.Time)[:19], rec.LoggerName, rec.Level, rec.Message)
}
