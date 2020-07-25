package input

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"reflect"
)

// StructFromDefault takes a pointer to a struct and iterates over its fields,
// setting each one based on input from the in Reader, with prompts sent to the
// out Writer. Fields written to in in Reader are separated by newlines. If no
// bytes are written before a newline then the default value for that field will
// be used. Values are updated in-place on the provided struct.
func StructFromDefault(object interface{}, in io.Reader, out io.Writer) error {
	scanner := bufio.NewScanner(in)
	v := reflect.ValueOf(object).Elem()
	t := v.Type()

	if v.Kind() != reflect.Struct {
		return errors.New("input is not a pointer to a struct")
	}

	for i := 0; i < v.NumField(); i++ {
		fmt.Fprintf(out, "%s [%v]:", t.Field(i).Name, v.Field(i).String())
		scanner.Scan()
		text := scanner.Text()
		if text != "" {
			v.Field(i).SetString(text)
		}
	}
	return nil
}
