package main

import (
	"fmt"

	"github.com/rizaramadan/slack-review/entity"
)

func main() {
	request := entity.NewRequest("[PCV-1] create voucher improvement")

	raymond := entity.NewApprover("Raymond")
	request.AddPending(raymond)

	fathir := entity.NewApprover("Fathir")
	request.AddPending(fathir)

	fmt.Println(fmt.Sprintf("pending of request: %s", request.GetName()))
	pending := request.GetPending()
	for _, each := range pending {
		fmt.Println(each.GetName())
	}
}
