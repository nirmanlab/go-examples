package random

import "fmt"

// array
// single item --> function controll outer scope

func processArray(callback func(item int64) string) {
	arr := []int64{1, 2, 3, 4, 5, 6, 7, 8}
	for _, item := range arr {
		fmt.Println(callback(item))
	}
}

func printCallback1(item int64) string {
	return fmt.Sprintf(" @@@@@ %d @@@@@ ", item)
}

func printCallback2(item int64) string {
	return fmt.Sprintf(" ======= %d ====== ", item)
}

func StartCallbackDemo() {
	processArray(printCallback1)

	processArray(printCallback2)
}
