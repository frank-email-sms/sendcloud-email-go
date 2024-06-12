package sendcloud_email_go

import "os"

type EmailCommonFields struct {
	From                string
	To                  string
	Subject             string
	ContentSummary      string
	FromName            string
	CC                  string
	BCC                 string
	ReplyTo             string
	LabelName           string
	Headers             string
	Attachments         []*os.File
	Xsmtpapi            string
	SendRequestID       string
	RespEmailID         bool
	UseNotification     bool
	UseAddressList      bool
	boundary                string
}

type SendEmailArgs struct {
	EmailCommonFields
	Html              string
	Plain             string
}

type SendEmailTemplateArgs struct {
	EmailCommonFields
	TemplateInvokeName string
}

type SendEmailResult struct {
	Result     bool        `json:"result"`
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Info       interface{} `json:"info"`
}