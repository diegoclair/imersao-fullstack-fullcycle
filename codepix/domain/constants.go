package domain

// Pix type values
const (
	PixKeytypeEmail string = "email"
	PixKeytypeCPF   string = "cpf"
)

// Pix status standard values
const (
	PixStatusActive   string = "active"
	PixStatusInactive string = "inactive"
)

//Transaction standard values
const (
	TransactionPending   string = "pending"
	TransactionCompleted string = "completed"
	TransactionError     string = "error"
	TransactionConfirmed string = "confirmed"
)
