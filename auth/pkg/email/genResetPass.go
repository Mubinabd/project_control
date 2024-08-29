package email

import (
	"fmt"
	"math/rand"
)

func GenForgotPassword() string {
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}
