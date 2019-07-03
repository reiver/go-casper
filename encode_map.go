package casper

import (
	"github.com/reiver/go-indent"

	"io"
	"reflect"
	"sort"
)

func encodeMap(writer io.Writer, source interface{}) error {
	reflectedSource := reflect.ValueOf(source)

	if reflect.Map != reflectedSource.Kind() {
		return unsupportedSource(source)
	}

	var keys []string = nil
	{
		var reflectedKeys []reflect.Value = reflectedSource.MapKeys()

		keys = make([]string, len(reflectedKeys))

		for i, reflectedKey := range reflectedKeys {
			keyInterface := reflectedKey.Interface()
			if nil == keyInterface {
				return errInternalError
			}

			key, casted := keyInterface.(string)
			if !casted {
				return unsupportedSource(source)
			}

			keys[i] = key
		}
	}

	sort.Strings(keys)

	for _, key := range keys {
		reflectedValue := reflectedSource.MapIndex(reflect.ValueOf(key))
		reflectedValueKind := reflectedValue.Kind()

		for reflect.Interface == reflectedValueKind {
			reflectedValue = reflectedValue.Elem()
			reflectedValueKind = reflectedValue.Kind()
		}

		switch reflectedValueKind {
		case reflect.Map:

			if err := encodeKey(writer, key); nil != err {
				return err
			}

			if err := encodeTab(writer); nil != err {
				return err
			}

			if err := encodeLeftCurlyBracket(writer); nil != err {
				return err
			}

			if err := encodeNewLine(writer); nil != err {
				return err
			}

			{
				var indenter io.Writer = &indent.Writer{
					Indentation: "\t",
					Writer: writer,
				}

				value := reflectedValue.Interface()

				if err := encodeMap(indenter, value); nil != err {
					return err
				}
			}

			if err := encodeRightCurlyBracket(writer); nil != err {
				return err
			}

			if err := encodeNewLine(writer); nil != err {
				return err
			}

		default:
			value := reflectedValue.Interface()

			if err := encodeKeyValue(writer, key, value); nil != err {
				return err
			}
		}
	}

	return nil
}
