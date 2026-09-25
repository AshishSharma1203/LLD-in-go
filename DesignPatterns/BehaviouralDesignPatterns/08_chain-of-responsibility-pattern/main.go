package main 

func main (){
	lead1:=NewLead(5);
	manager1:=NewManager(10)
	hr1:=NewHR(20)


	lead1.SetNextApprover(manager1)
	manager1.SetNextApprover(hr1)
	lead1.ApproveRequest(3);
	lead1.ApproveRequest(9);
	lead1.ApproveRequest(15);
	lead1.ApproveRequest(26);
}