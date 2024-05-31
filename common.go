package sendcloud_email_go

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/textproto"
	"os"
	"strconv"
	"strings"
)

type EmailCommonFields struct {
	APIUser             string
	APIKey              string
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
	XSMTPAPI            string
	SendRequestID       string
	RespEmailID         bool
	UseNotification     bool
	UseAddressList      bool
	boundary                string
}

type SendEmailArgs struct {
	EmailCommonFields
	HTML              string
	Plain             string
}

type SendEmailTemplateArgs struct {
	EmailCommonFields
	TemplateInvokeName string
}

func (e *EmailCommonFields) multipart(multipartWriter *multipart.Writer) error {

	var partWriter io.Writer

	var err error

	if e.APIUser != "" {
		err = writeMultipartText(multipartWriter, "apiUser", e.APIUser)
        if err!= nil {
            return err
        }
	}

	if e.APIKey != "" {
		err = writeMultipartText(multipartWriter, "apiKey", e.APIKey)
        if err!= nil {
            return err
        }
	}

	if e.From != "" {
		err = writeMultipartText(multipartWriter, "from", e.From)
		if err != nil {
			return err
		}
	}

	if e.To != "" {
		err = writeMultipartText(multipartWriter, "to", e.To)
		if err != nil {
			return err
		}
	}

	if e.Subject != "" {
		err = writeMultipartText(multipartWriter, "subject", e.Subject)
        if err!= nil {
            return err
        }
	}


	if e.ContentSummary!= "" {
		err = writeMultipartText(multipartWriter, "contentSummary", e.ContentSummary)
        if err!= nil {
            return err
        }
	}

	if e.FromName != "" {
		err = writeMultipartText(multipartWriter, "fromName", e.FromName)
		if err != nil {
			return err
		}
	}

	if e.CC != "" {
		err = writeMultipartText(multipartWriter, "cc", e.CC)
		if err != nil {
			return err
		}
	}

	if e.BCC != "" {
		err = writeMultipartText(multipartWriter, "bcc", e.BCC)
		if err != nil {
			return err
		}
	}

	if e.ReplyTo != "" {
		err = writeMultipartText(multipartWriter, "replyTo", e.ReplyTo)
		if err != nil {
			return err
		}
	}

	if e.LabelName != "" {
		err = writeMultipartText(multipartWriter, "labelName", e.LabelName)
		if err != nil {
			return err
		}
	}

	if e.Headers != "" {
		err = writeMultipartText(multipartWriter, "headers", e.Headers)
		if err != nil {
			return err
		}
	}

	if e.Attachments != nil {
		for _, attachment := range e.Attachments {
			defer attachment.Close()
			partWriter, err = multipartWriter.CreateFormFile("attachment", attachment.Name())
			if err != nil {
				return err
			}
			_, err = io.Copy(partWriter, attachment)
			if err != nil {
				return err
			}
		}
	}

	if e.XSMTPAPI != "" {
		err = writeMultipartText(multipartWriter, "xsmtpapi", e.XSMTPAPI)
		if err != nil {
			return err
		}
	}

	if e.SendRequestID != "" {
		err = writeMultipartText(multipartWriter, "sendRequestId", e.SendRequestID)
		if err != nil {
			return err
		}
	}

	if e.RespEmailID {
		respEmailIDStr := strconv.FormatBool(e.RespEmailID)
        err = writeMultipartText(multipartWriter, "respEmailId", respEmailIDStr)
        if err!= nil {
            return err
        }
	}

	if e.UseNotification {
		notificationStr := strconv.FormatBool(e.UseNotification)
		err = writeMultipartText(multipartWriter, "useNotification", notificationStr)
		if err != nil {
			return err
		}
	}

	if e.UseAddressList {
		useAddressListStr := strconv.FormatBool(e.UseAddressList)
		err = writeMultipartText(multipartWriter, "useAddressList", useAddressListStr)
		if err != nil {
			return err
		}
	}

	return nil
}

func (e *SendEmailArgs) Marshal() (*bytes.Buffer, error) {
	buf := bytes.Buffer{}
	multipartWriter := multipart.NewWriter(&buf)
	multipartWriter.Boundary()
	var err error

	err = e.multipart(multipartWriter)
	if err != nil {
		return nil, err
	}

	if e.HTML!= "" {
		err = writeMultipartText(multipartWriter, "html", e.HTML)
        if err!= nil {
            return nil, err
        }
	}

	if e.Plain!= "" {
		err = writeMultipartText(multipartWriter, "plain", e.Plain)
        if err!= nil {
            return nil, err
        }
	}

	multipartWriter.Close()
	e.boundary = multipartWriter.Boundary()
	return &buf, nil
}




type SendEmailResult struct {
	Result     bool        `json:"result"`
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Info       interface{} `json:"info"`
}



func escapeQuotes(s string) string {
	quoteEscaper := strings.NewReplacer("\\", "\\\\", `"`, "\\\"")
	return quoteEscaper.Replace(s)
}

func writeMultipart(writer *multipart.Writer, fieldName string, content []byte, contentType string) error {
	header := textproto.MIMEHeader{}
	header.Set("Content-Type", contentType)
	header.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"`, escapeQuotes(fieldName)))
	part, err := writer.CreatePart(header)
	if err != nil {
		return err
	}

	_, err = part.Write(content)
	if err != nil {
		return err
	}

	return nil
}


func writeMultipartText(writer *multipart.Writer, fieldName string, text string) error {
	return writeMultipart(writer, fieldName, []byte(text), "text/plain")
}
