package entity

//Approver  of request
type Approver struct {
	name string
}

//NewApprover is constructor of Approver
func NewApprover(name string) *Approver {
	result := Approver{name: name}
	return &result
}

//GetName return approver name
func (a *Approver) GetName() string {
	return a.name
}
