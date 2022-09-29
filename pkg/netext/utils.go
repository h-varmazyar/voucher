package netext

import "fmt"

func LocalAddress(port Port) string {
	return fmt.Sprintf(":%v", port)
}
