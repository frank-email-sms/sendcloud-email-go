package sendcloud_email_go

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/textproto"
	"net/url"
	"strconv"
	"strings"
)





func (e *EmailCommonFields) multipart(multipartWriter *multipart.Writer) error {

	var partWriter io.Writer

	var err error

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

	if e.Xsmtpapi != "" {
		err = writeMultipartText(multipartWriter, "xsmtpapi", e.Xsmtpapi)
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

func (client *SendCloud) PrepareSendCommonEmailParams(e *SendEmailArgs) (url.Values, error) {
	params := url.Values{}
	params.Set("apiUser", client.apiUser)
	params.Set("apiKey", client.apiKey)
	params.Set("from", e.From)
	if e.To!= "" {
		params.Set("to", e.To)
	}
	params.Set("subject", e.Subject)
	if e.Html!= "" {
        params.Set("html", e.Html)
    }
	if e.ContentSummary!= "" {
        params.Set("contentSummary", e.ContentSummary)
    }
    if e.FromName!= "" {
        params.Set("fromName", e.FromName)
    }
    if e.CC!= "" {
        params.Set("cc", e.CC)
    }
    if e.BCC!= "" {
        params.Set("bcc", e.BCC)
    }
    if e.ReplyTo!= "" {
        params.Set("replyTo", e.ReplyTo)
    }
    if e.LabelName!= "" {
        params.Set("labelName", e.LabelName)
    }
    if e.Headers!= "" {
    	params.Set("headers", e.Headers)
	}
	if e.Xsmtpapi!= "" {
		params.Set("xsmtpapi", e.Xsmtpapi)
	}
	if e.Plain!= "" {
		params.Set("plain", e.Plain)
	}
	if e.SendRequestID!= "" {
        params.Set("sendRequestId", e.SendRequestID)
    }
    if e.RespEmailID {
		params.Set("respEmailId", strconv.FormatBool(e.RespEmailID))
    }
    if e.UseNotification {
		params.Set("useNotification", strconv.FormatBool(e.UseNotification))
	}
	if e.UseAddressList {
        params.Set("useAddressList", strconv.FormatBool(e.UseAddressList))
    }
	return params, nil
}

func (client *SendCloud) MarshalSendEmailArgs(e *SendEmailArgs) (*bytes.Buffer, error) {
	buf := bytes.Buffer{}
	multipartWriter := multipart.NewWriter(&buf)
	multipartWriter.Boundary()
	var err error

	if client.apiUser != "" {
		err = writeMultipartText(multipartWriter, "apiUser", client.apiUser)
		if err!= nil {
			return nil,err
		}
	}

	if client.apiKey != "" {
		err = writeMultipartText(multipartWriter, "apiKey", client.apiKey)
		if err!= nil {
			return nil,err
		}
	}

	err = e.multipart(multipartWriter)
	if err != nil {
		return nil, err
	}

	if e.Html!= "" {
		err = writeMultipartText(multipartWriter, "html", e.Html)
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
