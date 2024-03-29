//+build djavul

package main

// #include <stdio.h>
//
// static void djavul_cinit() {
//    void (**djavul_cpp_init_funcs)(void) = (void *) 0x483000;
//    for (int i = 1; i < 34; i++) {
//       void (*f)(void) = djavul_cpp_init_funcs[i];
//       if (f == NULL) {
//          break;
//       }
//       f();
//    }
// }
import "C"

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/sanctuary/djavul/d1/diablo"
	"github.com/sanctuary/djavul/d1/engine"
	"github.com/sanctuary/djavul/d1/sound"
	"github.com/sanctuary/djavul/internal/proto"
	// TODO: update to github.com/AllenDang/w32 when
	// https://github.com/AllenDang/w32/pull/81 is merged.
	"github.com/TheTitanrain/w32"
)

// useFrontend specifies whether the Djavul frontend is enabled. When enabled,
// djavul.exe connect to djavul-frontend through IPC (either TCP or named
// pipes), and sends rendering and audio playback events to the frontend.
const useFrontend = false

//export Start
func Start() {
	// Store standard output in djavul.log. For trouble-shooting on Windows.
	//f, err := os.Create("djavul.log")
	//if err != nil {
	//	log.Fatalf("%+v", err)
	//}
	//defer f.Close()
	//os.Stdout = f
	//os.Stderr = f
	//log.SetOutput(f)

	fmt.Println("djavul.Start: entry point in Go")
	cinit()

	// Parse command line flags.
	//var (
	//	start, end int64
	//)
	//flag.Int64Var(&start, "start", 0, "first seed")
	//flag.Int64Var(&end, "end", 256, "last seed")
	//flag.Parse()

	//engine.UseGUI = false
	//sound.UseSound = false
	//if err := compareL1(start, end); err != nil {
	//	log.Fatalf("%+v", err)
	//}
	//os.Exit(0)

	if useFrontend {
		// Parse command line flags.
		var (
			// frontend IP-address.
			ip string
			// npipe specifies whether to use named pipes for IPC.
			npipe bool
		)
		flag.StringVar(&ip, "ip", "localhost", "frontend IP-address")
		flag.BoolVar(&npipe, "npipe", false, "use named pipes for IPC")
		flag.Parse()

		var (
			stable   proto.IPC
			unstable proto.IPC
		)
		if npipe {
			stable = proto.NewStableNamedPipe(ip)
			unstable = proto.NewUnstableNamedPipe(ip)
		} else {
			stable = proto.NewStableTCP(ip)
			unstable = proto.NewUnstableTCP(ip)
		}

		engine.UseGUI = true
		sound.UseSound = false
		if err := initFrontConn(stable, unstable); err != nil {
			log.Fatalf("%+v", err)
		}
	} else {
		engine.UseGUI = false
		sound.UseSound = false
	}
	winGUI()

	//l1.UseGo = false
	//dumpL1Maps()

	//if err := checkL1Regular(); err != nil {
	//	log.Fatalf("%+v", err)
	//}
	//if err := checkL1Quest(); err != nil {
	//	log.Fatalf("%+v", err)
	//}
	os.Exit(0)
}

// winGUI initializes the Windows GUI.
func winGUI() {
	inst := w32.GetModuleHandle("")
	// Parse arguments from command line.
	var s int64
	flag.Int64Var(&s, "r", 0, "initial signed 32-bit seed for dungeon generation")
	flag.Parse()
	switch {
	case s >= -2147483648 && s <= 2147483647:
		*diablo.FlagRSeed = int32(s)
	default:
		panic(fmt.Errorf("invalid seed; expected >= -2147483648 and <= 2147483647; got %d", s))
	}
	args := strings.Join(flag.Args(), " ")
	fmt.Println("args:", args)
	show := w32.SW_SHOWDEFAULT
	diablo.WinMain(inst, 0, args, show)
}

// cinit invokes cpp initialiation functions.
func cinit() {
	C.djavul_cinit()
}
