package lctime

import (
	"bytes"
	"fmt"
	"strconv"
	"time"

	"golang.org/x/text/language"
)

// Strftime formats a time.Time. It's locale-aware, so make sure you call
// SetLocale if needed.
func Strftime(format string, t time.Time) string {
	return lc.Strftime(format, t)
}

// StrftimeLoc formats a time.Time. It's locale-aware, so make sure you call
// SetLocale if needed.
func StrftimeLoc(locale, format string, t time.Time) (string, error) {
	lc, err := loadLocale(locale)
	if err != nil {
		return "", err
	}
	return lc.Strftime(format, t), nil
}

// /convert charNumber to int
func CharToNumber(c rune) int {
	return int(c - '0')
}

// /to translate number
func TranslateNumber(num string) (string, error) {
	lc, err := loadLocale("ar_EG")
	if err != nil {
		fmt.Println("Failed to load locale. Err= ", err)
		return "", err
	}

	n, err := strconv.Atoi(num)
	if err != nil {
		fmt.Println("Failed to parse string. Err= ", err)
		return "", err
	}

	return lc.translateNumber(n), nil
}

// to translate time.duration
func Strfduration(duration time.Duration, locale language.Tag) (string, error) {
	switch locale {
	case language.Arabic:
		lc, err := loadLocale(locale.String() + "_EG")
		if err != nil {
			fmt.Println("Failed to load locale. Err= ", err)
			return "", err
		}
		return lc.translateNumber(int(duration.Minutes())), nil
	default:
		return fmt.Sprintf("%d", int32(duration.Minutes())), nil
	}
}

func (lc *localeData) Strftime(format string, t time.Time) string {
	if len(format) < 1 {
		return format
	}

	buf := new(bytes.Buffer)
	end := len(format)

	for i := 0; i < end; i++ {
		if format[i] == '%' && i+2 <= end {
			buf.WriteString(lc.parseDirective(format[i:i+2], t))
			i++
			continue
		}

		buf.WriteByte(format[i])
	}

	return buf.String()
}

func (lc *localeData) parseDirective(direc string, t time.Time) string {
	if len(direc) < 2 {
		return direc
	}

	switch direc[:2] {
	case "%a":
		return lc.pera(t)
	case "%A":
		return lc.perA(t)
	case "%b":
		return lc.perb(t)
	case "%B":
		return lc.perB(t)
	case "%c":
		return lc.perc(t)
	case "%C":
		return lc.perC(t)
	case "%d":
		return lc.perd(t)
	case "%D":
		return lc.perD(t)
	case "%e":
		return lc.pere(t)
	case "%F":
		return lc.perF(t)
	case "%g":
		return lc.perg(t)
	case "%G":
		return lc.perG(t)
	case "%H":
		return lc.perH(t)
	case "%I":
		return lc.perI(t)
	case "%j":
		return lc.perj(t)
	case "%m":
		return lc.perm(t)
	case "%M":
		return lc.perM(t)
	case "%n":
		return lc.pern(t)
	case "%p":
		return lc.perp(t)
	case "%r":
		return lc.perr(t)
	case "%R":
		return lc.perR(t)
	case "%S":
		return lc.perS(t)
	case "%t":
		return lc.pert(t)
	case "%T":
		return lc.perT(t)
	case "%u":
		return lc.peru(t)
	case "%U":
		return lc.perU(t)
	case "%V":
		return lc.perV(t)
	case "%w":
		return lc.perw(t)
	case "%W":
		return lc.perW(t)
	case "%x":
		return lc.perx(t)
	case "%X":
		return lc.perX(t)
	case "%y":
		return lc.pery(t)
	case "%Y":
		return lc.perY(t)
	case "%z":
		return lc.perz(t)
	case "%Z":
		return lc.perZ(t)
	case "%%":
		return lc.perper(t)
	}

	return direc
}
