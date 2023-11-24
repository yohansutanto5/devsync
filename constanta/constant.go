package constanta

const (
	InternalServerErrorMessage = "Please call our Customer service and inform the transaction IDs. Transaction ID is: "
	SuccessMessage             = "Success"
)

// Status and Error Code
const (
	DebugCode               = "APP-DBG"
	DebugStatus             = 999
	ErrorStatus             = 500
	FailToConnectCode       = "APP-CONN"
	InternalServerErrorCode = "APP-SVR"
	StatusOK                = 200
	CodeOK                  = "APP-SUCCESS"
	CodeErrorService        = "APP-SVR"
)

// Release Status
const (
	UATReady        = "UAT Deployment Ready"        // UAT Trigger Deployment
	UATDeploy       = "UAT Deployment In Progress"  // deployment Failed / Redeploy
	UATVerify       = "UAT Deployment Verified"     // Verify UAT deployment
	ApprovalPending = "Pending Production Approval" // Approve / Reject
	PRDReady        = "PRD Deployment Ready"        // Trigger Prd Deployment
	PRDDeploy       = "PRD Deployment In Progress"  // PRD Deployment Failed /Redeploy
	PRDVerify       = "PRD Deployment Verified"     // Verify PRD Deployment
	Closed          = "Closed"                      // End
	Failed          = "Deployment Failed"           // click redeploy/Rollback
	Rejected        = "Rejected"                    // End
)
