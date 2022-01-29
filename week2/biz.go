package week2

import "fmt"

func someBiz() error {
	data, err := query()
	if err != nil {
		if IsNotFoundError(err) {
			fmt.Printf("it's a not found error")
		} else {
			fmt.Printf("it's a system error")
		}
	}
	for _, _ = range data {

	}
	return err
}
