//+build djavul

// Run in djavul.exe environment.

package inv

import (
	"unsafe"

	"github.com/sanctuary/djavul/internal/types"
)

func init() {
	// Init pointers to global variables of djavul.exe.
	ScreenPos = (*[73]types.Point)(unsafe.Pointer(uintptr(0x47AE60)))
	StartSlot2x2 = (*[10]int32)(unsafe.Pointer(uintptr(0x48E9A8)))
}
