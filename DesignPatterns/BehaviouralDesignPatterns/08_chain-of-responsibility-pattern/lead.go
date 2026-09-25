package main

import "fmt"

type Lead struct {
	approvalLimit int
	nextApprover  Approver
}

func NewLead(days int) *Lead {
	return &Lead{
		approvalLimit: days,
	}
}

func (lead *Lead) SetNextApprover(next Approver) {
	lead.nextApprover = next
}

func (lead *Lead) ApproveRequest(days int) {
	fmt.Println("Approving leave of ", days, "from lead ")
	if days <= lead.approvalLimit {
		fmt.Println("Leave approved by Lead as it is within ", lead.approvalLimit)
	} else if lead.nextApprover != nil {
		lead.nextApprover.ApproveRequest(days)

	} else {
		fmt.Println("Request denied by Lead not withing power")
	}

}
