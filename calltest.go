package main

import (
	"runtime"
)

func main() {
	Foo()
}
func Bar() {
	// fmt.Printf("我是 %s, %s 又在调用我!\n", printMyName(), printCallerName())
	trace2()
}
func trace() {
	pc := make([]uintptr, 10) // at least 1 entry needed, 几位，表示打印几层
	// fmt.Printf("pc is %v\n", pc) //pc is [0 0 0 0 0 0 0 0 0 0]
	n := runtime.Callers(0, pc)
	// fmt.Printf("new pc is %v\n", pc) //new pc is [165856 926944 926592 928192 926208 357248 578177 0 0 0]
	// fmt.Printf("n is %v\n", n)
	for i := 0; i < n; i++ {
		f := runtime.FuncForPC(pc[i])
		file, line := f.FileLine(pc[i])
		// fmt.Printf("%s:%d %s\n", file, line, f.Name())
	}

}

func trace2() {
	pc := make([]uintptr, 10) // at least 1 entry needed
	n := runtime.Callers(0, pc)
	// fmt.Printf("n is %v\n", n)
	frames := runtime.CallersFrames(pc[:n])
	for {
		frame, more := frames.Next()
		// fmt.Printf("%s:%d %s\n", frame.File, frame.Line, frame.Function)
		if !more {
			break
		}
	}
}

func Foo() {
	// fmt.Printf("我是 %s, %s 在调用我!\n", printMyName(), printCallerName())
	Bar()
}

func printMyName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}
func printCallerName() string {
	pc, _, _, _ := runtime.Caller(2)
	return runtime.FuncForPC(pc).Name()
}
