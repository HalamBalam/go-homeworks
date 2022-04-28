package hw9

import "io"

func printStrings(w io.Writer, objects ...any) error {
	for _, obj := range objects {
		switch str := obj.(type) {
		case string:
			_, err := w.Write([]byte(str))
			if err != nil {
				return err
			}
		}
	}
	return nil
}
