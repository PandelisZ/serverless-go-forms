package types

// ContactBasic is a basic form used when you only really need a message
type ContactBasic struct {
	From    string
	Message string
	Referer string
}
