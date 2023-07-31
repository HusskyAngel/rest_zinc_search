package optionutils

import "runtime"

func numCores() int {
	return runtime.NumCPU()-1 
}
