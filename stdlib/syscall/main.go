package main

import (
	"fmt"
	"log"
	"syscall"
	"unsafe"
)

const OK = 0
const NULL = 0

type Sqlite3 struct {
	// not required
}

func main() {
	mod := syscall.NewLazyDLL("sqlite3.dll")

	sqlite3POpen := mod.NewProc("sqlite3_open")
	sqlite3Exec := mod.NewProc("sqlite3_exec")
	sqlite3Close := mod.NewProc("sqlite3_close")

	file := []byte("test.db")

	var db Sqlite3
	dbptr := unsafe.Pointer(&db)

	ret, _, err := sqlite3POpen.Call(
		uintptr(unsafe.Pointer(&file[0])), // const char *
		uintptr(unsafe.Pointer(&dbptr)),   // sqlite3 **
	)
	if ret != OK {
		log.Fatal(err)
	}

	defer sqlite3Close.Call(uintptr(dbptr)) // sqlite3*

	sql := []byte("CREATE TABLE test (id int);")
	ret, _, err = sqlite3Exec.Call(
		// sqlite3*
		uintptr(dbptr),
		// const char *sql
		uintptr(unsafe.Pointer(&sql[0])),
		// int (*callback)(void*,int,char**,char**)
		NULL,
		// void *
		NULL,
		// char **errmsg
		NULL,
	)
	if ret != OK {
		log.Fatal(err)
	}

	_, err = fmt.Println("OK.")
	if err != nil {
		log.Fatal(err)
	}
}
