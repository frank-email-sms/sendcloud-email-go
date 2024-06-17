package sendcloud_email_go

import (
	"os"
	"reflect"
)

type CommonMail struct {
	To                  Receiver
	body 				MailBody
	content				TextContent
}

type TemplateMail struct {
	To                  Receiver
	body 				MailBody
	content				TemplateContent
}

type TextContent struct {
	Html              string
	Plain             string
}

type TemplateContent struct {
	TemplateInvokeName string
}


// SetHTML - Set the html content of the email, required if not using a template.
func (e *TextContent) SetHTML(html string) {
	e.Html = html
}

// SetPlain - Set the plain content of the email, required if not using a template.
func (e *TextContent) SetPlain(plain string) {
	e.Plain = plain
}

// SetTemplateInvokeName - Set the template invoke name.
func (e *TemplateContent) SetTemplateInvokeName(name string) {
	e.TemplateInvokeName = name
}

type MailBody struct {
	From                string
	Subject             string
	ContentSummary      string
	FromName            string
	ReplyTo             string
	LabelName           string
	Headers             map[string]string
	Attachments         []*os.File
	Xsmtpapi           XSMTPAPI
	SendRequestID       string
	RespEmailID         bool
	UseNotification     bool
}

type XSMTPAPI struct {
	To        []string         `json:"to,omitempty"`
	Sub       map[string][]any `json:"sub,omitempty"`
	Pubsub    map[string]any   `json:"pubsub,omitempty"`
	Filters      Filter    `json:"filters,omitempty"`
	Settings      Settings      `json:"settings,omitempty"`
}

func (a XSMTPAPI) IsEmpty() bool {
	return reflect.DeepEqual(a, XSMTPAPI{})
}

type FilterSettings struct {
	Enable string `json:"enable"`
}

type TrackingFilter struct {
	Settings FilterSettings `json:"settings"`
}

type Filter struct {
	SubscriptionTracking TrackingFilter `json:"subscription_tracking"`
	OpenTracking         TrackingFilter `json:"open_tracking"`
	ClickTracking        TrackingFilter `json:"click_tracking"`
}

// UnsubscribeSettings 表示退订设置的结构体
type UnsubscribeSettings struct {
	PageID []int `json:"page_id"`
}

// Settings 表示设置的结构体
type Settings struct {
	Unsubscribe UnsubscribeSettings `json:"unsubscribe"`
}

// SetFrom - Set the from address.
func (e *MailBody) SetFrom(from string) {
	e.From = from
}

// SetContentSummary - Set the content summary of the email.
func (e *MailBody) SetContentSummary(contentSummary string) {
	e.ContentSummary = contentSummary
}

// SetFromName - Set the from name of the email.
func (e *MailBody) SetFromName(fromName string) {
	e.FromName = fromName
}


// SetReplyTo - Set the reply to address.
func (e *MailBody) SetReplyTo(replyTo string) {
	e.ReplyTo = replyTo
}


// SetLabelName - Set the label name of the email.
func (e *MailBody) SetLabelName(labelName string) {
	e.LabelName = labelName
}

// AddHeaders - Add the headers of the email.
func (e *MailBody) AddHeaders(headers map[string]string) {
	e.Headers = headers
}

// AddAttachment - Add an attachment content.
func (e *MailBody) AddAttachment(attachment *os.File) {
	e.Attachments = append(e.Attachments, attachment)
}

// SetXsmtpapi - Set the xsmtpapi of the email.
func (e *MailBody) SetXsmtpapi(xsmtpapi XSMTPAPI) {
	e.Xsmtpapi = xsmtpapi
}


type Receiver struct {
	To                  string
	CC                  string
	BCC                 string
	UseAddressList      bool
}

type SendEmailResult struct {
	Result     bool        `json:"result"`
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Info       interface{} `json:"info"`
}


