package main

import "fmt"

type Manager struct {
	approvalLimit int
	nextApprover  Approver
}

func NewManager(days int) *Manager {
	return &Manager{
		approvalLimit: days,
	}
}

func (manager *Manager) SetNextApprover(next Approver) {
	manager.nextApprover = next
}

func (manager *Manager) ApproveRequest(days int) {
	fmt.Println("Approving leave of ", days, "from manager  ")
	if days <= manager.approvalLimit {
		fmt.Println("Leave approved by Manager as it is within ", manager.approvalLimit)
	} else if manager.nextApprover != nil {
		manager.nextApprover.ApproveRequest(days)

	} else {
		fmt.Println("Request denied by Manager not withing power")
	}
}
