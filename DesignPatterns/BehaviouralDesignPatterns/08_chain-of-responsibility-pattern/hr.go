package main

import "fmt"

type HR struct {
	approvalLimit int
	nextApprover  Approver
}

func NewHR(days int) *HR {
	return &HR{
		approvalLimit: days,
	}
}

func (hr *HR) SetNextApprover(next Approver) {
	hr.nextApprover = next
}

func (hr *HR) ApproveRequest(days int) {
	fmt.Println("Approving leave of ", days, "from hr ")
	if days <= hr.approvalLimit {
		fmt.Println("Leave approved by HR as it is within ", hr.approvalLimit)
	} else if hr.nextApprover != nil {
		hr.nextApprover.ApproveRequest(days)

	} else {
		fmt.Println("We will get back to you ")
	}
}
