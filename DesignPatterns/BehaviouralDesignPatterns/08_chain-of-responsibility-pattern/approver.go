package main 

type Approver interface {
	ApproveRequest(days int)
}