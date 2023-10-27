package vm

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"unsafe"

	"github.com/IBAX-io/needle/compile"

	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
)

// StateName checks the name of the contract and modifies it to @[state]name if it is necessary.
func StateName(state uint32, name string) string {
	if !strings.HasPrefix(name, `@`) {
		return fmt.Sprintf(`@%d%s`, state, name)
	} else if len(name) > 1 && (name[1] < '0' || name[1] > '9') {
		name = `@1` + name[1:]
	}
	return name
}

// ParseName gets a state identifier and the name of the contract or table
// from the full name like @[id]name
func ParseName(in string) (id int64, name string) {
	re := regexp.MustCompile(`(?is)^@(\d+)(\w[_\w]*)$`)
	ret := re.FindStringSubmatch(in)
	if len(ret) == 3 {
		id = StrToInt64(ret[1])
		name = ret[2]
	}
	return
}

// StrToInt64 converts string to int64
func StrToInt64(s string) int64 {
	ret, _ := strconv.ParseInt(s, 10, 64)
	return ret
}

// ValueToInt converts interface (string or int64) to int64
func ValueToInt(v any) (ret int64, err error) {
	switch val := v.(type) {
	case float64:
		ret = int64(val)
	case int64:
		ret = val
	case string:
		if len(val) == 0 {
			return 0, nil
		}
		ret, err = strconv.ParseInt(val, 10, 64)
		if err != nil {
			errText := err.Error()
			if strings.Contains(errText, `:`) {
				errText = errText[strings.LastIndexByte(errText, ':'):]
			} else {
				errText = ``
			}
			err = fmt.Errorf(`%s is not a valid integer %s`, val, errText)
		}
	case decimal.Decimal:
		ret = val.IntPart()
	case json.Number:
		ret, err = val.Int64()
	default:
		if v == nil {
			return 0, nil
		}
		err = fmt.Errorf(`%v is not a valid integer`, val)
	}
	if err != nil {
		log.WithFields(log.Fields{"type": ConversionError, "error": err,
			"value": fmt.Sprint(v)}).Error("converting value to int")
	}
	return
}

// ValueToFloat converts interface (string, float64 or int64) to float64
func ValueToFloat(v any) (ret float64) {
	var err error
	switch val := v.(type) {
	case float64:
		ret = val
	case int64:
		ret = float64(val)
	case string:
		ret, err = strconv.ParseFloat(val, 64)
		if err != nil {
			log.WithFields(log.Fields{"type": ConversionError, "error": err, "value": val}).Error("converting value from string to float")
		}
	case decimal.Decimal:
		ret = val.InexactFloat64()
	}
	return
}

// ValueToDecimal converts interface (string, float64, Decimal or int64) to Decimal
func ValueToDecimal(v any) (ret decimal.Decimal, err error) {
	switch val := v.(type) {
	case float64:
		ret = decimal.NewFromFloat(val)
	case string:
		ret, err = decimal.NewFromString(val)
	case int64:
		ret = decimal.New(val, 0)
	case decimal.Decimal:
		ret = val
	}
	return
}

func valueToBool(v any) bool {
	switch val := v.(type) {
	case int:
		if val != 0 {
			return true
		}
	case int64:
		if val != 0 {
			return true
		}
	case float64:
		if val != 0.0 {
			return true
		}
	case bool:
		return val
	case string:
		return len(val) > 0
	case []uint8:
		return len(val) > 0
	case []any:
		return val != nil && len(val) > 0
	case map[string]any:
		return val != nil && len(val) > 0
	case map[string]string:
		return val != nil && len(val) > 0
	case *compile.Map:
		return val != nil && val.Size() > 0
	default:
		dec, _ := decimal.NewFromString(fmt.Sprintf(`%v`, val))
		return dec.Cmp(decimal.Zero) != 0
	}
	return false
}

func isSelfAssignment(dest, value any) bool {
	if _, ok := value.([]any); !ok {
		if _, ok = value.(*compile.Map); !ok {
			return false
		}
	}
	if reflect.ValueOf(dest).Pointer() == reflect.ValueOf(value).Pointer() {
		return true
	}
	switch v := value.(type) {
	case []any:
		for _, item := range v {

			if isSelfAssignment(dest, item) {
				return true
			}
		}
	case *compile.Map:
		for _, item := range v.Values() {
			if isSelfAssignment(dest, item) {
				return true
			}
		}
	}
	return false
}

func calcMem(v any) (mem int64) {
	rv := reflect.ValueOf(v)

	switch rv.Kind() {
	case reflect.Bool:
		mem = 1
	case reflect.Int8, reflect.Uint8:
		mem = 1
	case reflect.Int16, reflect.Uint16:
		mem = 2
	case reflect.Int32, reflect.Uint32:
		mem = 4
	case reflect.Int64, reflect.Uint64, reflect.Int, reflect.Uint:
		mem = 8
	case reflect.Float32:
		mem = 4
	case reflect.Float64:
		mem = 8
	case reflect.String:
		mem += int64(rv.Len())
	case reflect.Slice, reflect.Array:
		mem = 12
		for i := 0; i < rv.Len(); i++ {
			mem += calcMem(rv.Index(i).Interface())
		}
	case reflect.Map:
		mem = 4
		for _, k := range rv.MapKeys() {
			mem += calcMem(k.Interface())
			mem += calcMem(rv.MapIndex(k).Interface())
		}
	default:
		mem = int64(unsafe.Sizeof(v))
	}

	return
}
