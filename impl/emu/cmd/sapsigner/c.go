package main

// #cgo         LDFLAGS: -L../../lib
// #cgo !darwin LDFLAGS: -Wl,-rpath,$ORIGIN
// #cgo !darwin LDFLAGS: -Wl,-rpath,$ORIGIN/lib
// #cgo !darwin LDFLAGS: -Wl,-rpath,$ORIGIN/../lib
// #cgo  darwin LDFLAGS: -Wl,-rpath,@loader_path
// #cgo  darwin LDFLAGS: -Wl,-rpath,@loader_path/lib
// #cgo  darwin LDFLAGS: -Wl,-rpath,@loader_path/../lib
// #cgo   linux LDFLAGS: -Wl,-rpath,/home/linuxbrew/.linuxbrew/lib
// #cgo  darwin LDFLAGS: -Wl,-rpath,/opt/homebrew/lib
// #cgo         LDFLAGS: -Wl,-rpath,/usr/local/lib
import "C"
