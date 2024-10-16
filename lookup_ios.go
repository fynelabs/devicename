//go:build ios

package main

import "fyne.io/fyne/v2/driver"

/*
const char *iosName();
*/
import "C"

func lookupName() (name string) {
	driver.RunNative(func(ctx interface{}) error {
		str := C.iosName()
		name = C.GoString(str)
		return nil
	})

	return name
}
