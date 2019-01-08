package ledger

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	es := make([]Entry, len(entries))
	copy(es, entries)

	if len(es) == 0 {
		if _, err := FormatLedger(currency, "en-US",
			[]Entry{{Date: "2014-01-01", Description: "", Change: 0}}); err != nil {
			return "", err
		}
	}

	sort.Slice(es, func(e1, e2 int) bool {
		if es[e1].Date == es[e2].Date && es[e1].Description == es[e2].Description {
			return es[e1].Change < es[e2].Change
		} else if es[e1].Date == es[e2].Date {
			return es[e1].Description < es[e2].Description
		} else {
			return es[e1].Date < es[e2].Date
		}
	})

	var s string
	if locale == "nl-NL" {
		s = fmt.Sprintf("%-10s | %-25s | %s\n", "Datum", "Omschrijving", "Verandering")
	} else if locale == "en-US" {
		s = fmt.Sprintf("%-10s | %-25s | %s\n", "Date", "Description", "Change")
	} else {
		return "", errors.New("")
	}

	for _, entry := range es {
		if len(entry.Date) != 10 {
			return "", errors.New("")
		}
		d1, d2, d3, d4, d5 := entry.Date[0:4], entry.Date[4], entry.Date[5:7], entry.Date[7], entry.Date[8:10]
		if d2 != '-' || d4 != '-' {
			return "", errors.New("")
		}
		de := entry.Description
		if len(de) > 25 {
			de = de[:22] + "..."
		} else {
			de = de + strings.Repeat(" ", 25-len(de))
		}
		var d string
		if locale == "nl-NL" {
			d = d5 + "-" + d3 + "-" + d1
		} else if locale == "en-US" {
			d = d3 + "/" + d5 + "/" + d1
		}
		negative := false
		cents := entry.Change
		if cents < 0 {
			cents = cents * -1
			negative = true
		}
		var a string
		if locale == "nl-NL" {
			if currency == "EUR" {
				a += "€ "
			} else if currency == "USD" {
				a += "$ "
			} else {
				return "", errors.New("")
			}
			centsStr := strconv.Itoa(cents)
			switch len(centsStr) {
			case 1:
				centsStr = "00" + centsStr
			case 2:
				centsStr = "0" + centsStr
			}
			rest := centsStr[:len(centsStr)-2]
			var parts []string
			for len(rest) > 3 {
				parts = append(parts, rest[len(rest)-3:])
				rest = rest[:len(rest)-3]
			}
			if len(rest) > 0 {
				parts = append(parts, rest)
			}
			for i := len(parts) - 1; i >= 0; i-- {
				a += parts[i] + "."
			}
			a = a[:len(a)-1]
			a += ","
			a += centsStr[len(centsStr)-2:]
			if negative {
				a += "-"
			} else {
				a += " "
			}
		} else if locale == "en-US" {
			if negative {
				a += "("
			}
			if currency == "EUR" {
				a += "€"
			} else if currency == "USD" {
				a += "$"
			} else {
				return "", errors.New("")
			}
			centsStr := strconv.Itoa(cents)
			switch len(centsStr) {
			case 1:
				centsStr = "00" + centsStr
			case 2:
				centsStr = "0" + centsStr
			}
			rest := centsStr[:len(centsStr)-2]
			var parts []string
			for len(rest) > 3 {
				parts = append(parts, rest[len(rest)-3:])
				rest = rest[:len(rest)-3]
			}
			if len(rest) > 0 {
				parts = append(parts, rest)
			}
			for i := len(parts) - 1; i >= 0; i-- {
				a += parts[i] + ","
			}
			a = a[:len(a)-1]
			a += "."
			a += centsStr[len(centsStr)-2:]
			if negative {
				a += ")"
			} else {
				a += " "
			}
		} else {
			return "", errors.New("")
		}
		s += fmt.Sprintf("%-10s | %s | %13s\n", d, de, a)
	}
	return s, nil
}
