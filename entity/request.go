package entity

//Request is the item to be approved by approver
type Request struct {
	name     string
	pending  []*Approver
	approved []Approver
}

//NewRequest create new request
func NewRequest(name string) *Request {
	result := Request{name: name}
	result.pending = nil
	result.approved = nil
	return &result
}

//GetName get name
func (r *Request) GetName() string { return r.name }

//AddPending add approver to pending list
func (r *Request) AddPending(approver *Approver) {
	r.pending = append(r.pending, approver)
}

//GetPending get pending approvals
func (r Request) GetPending() []*Approver {
	return r.pending
}
