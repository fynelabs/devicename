//go:build android

package main

import "fyne.io/fyne/v2/driver"

/*
#include <stdlib.h>

const char *androidName(uintptr_t java_vm, uintptr_t jni_env, uintptr_t ctx);
*/
import "C"

func lookupName() (name string) {
	driver.RunNative(func(ctx interface{}) error {
		ac := ctx.(*driver.AndroidContext)

		str := C.androidName(C.uintptr_t(ac.VM), C.uintptr_t(ac.Env), C.uintptr_t(ac.Ctx))
		name = C.GoString(str)
		return nil
	})

	return name
}
