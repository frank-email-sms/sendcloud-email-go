package sendcloud_email_go

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/url"
	"strconv"
)





func (e *EmailCommonFields) multipart(multipartWriter *multipart.Writer) error {

	var partWriter io.Writer

	var err error

	if e.From != "" {
		err = multipartWriter.WriteField("from", e.From)
		if err != nil {
			return err
		}
	}

	if e.To != "" {
		err =   multipartWriter.WriteField("to", e.To)
		if err != nil {
			return err
		}
	}

	if e.Subject != "" {
		err = multipartWriter.WriteField("subject", e.Subject)
        if err!= nil {
            return err
        }
	}


	if e.ContentSummary!= "" {
		err = multipartWriter.WriteField("contentSummary", e.ContentSummary)
        if err!= nil {
            return err
        }
	}

	if e.FromName != "" {
		err = multipartWriter.WriteField("fromName", e.FromName)
		if err != nil {
			return err
		}
	}

	if e.CC != "" {
		err = multipartWriter.WriteField("cc", e.CC)
		if err != nil {
			return err
		}
	}

	if e.BCC != "" {
		err = multipartWriter.WriteField("bcc", e.BCC)
		if err != nil {
			return err
		}
	}

	if e.ReplyTo != "" {
		err = multipartWriter.WriteField("replyTo", e.ReplyTo)
		if err != nil {
			return err
		}
	}

	if e.LabelName != "" {
		err = multipartWriter.WriteField("labelName", e.LabelName)
		if err != nil {
			return err
		}
	}

	if e.Headers != "" {
		err = multipartWriter.WriteField("headers", e.Headers)
		if err != nil {
			return err
		}
	}

	if e.Attachments != nil {
		for _, attachment := range e.Attachments {
			defer attachment.Close()
			partWriter, err = multipartWriter.CreateFormFile("attachments", attachment.Name())
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
		err = multipartWriter.WriteField("x-smtpapi", e.Xsmtpapi)
		if err != nil {
			return err
		}
	}

	if e.SendRequestID != "" {
		err = multipartWriter.WriteField("sendRequestId", e.SendRequestID)
		if err != nil {
			return err
		}
	}

	if e.RespEmailID {
		respEmailIDStr := strconv.FormatBool(e.RespEmailID)
		err = multipartWriter.WriteField("respEmailId", respEmailIDStr)
        if err!= nil {
            return err
        }
	}

	if e.UseNotification {
		notificationStr := strconv.FormatBool(e.UseNotification)
		err = multipartWriter.WriteField("useNotification", notificationStr)
		if err != nil {
			return err
		}
	}

	if e.UseAddressList {
		useAddressListStr := strconv.FormatBool(e.UseAddressList)
		err = multipartWriter.WriteField("useAddressList", useAddressListStr)
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

func (client *SendCloud) MarshalSendEmailArgs(e *SendEmailArgs) (*multipart.Writer,*bytes.Buffer, error) {
	buf := bytes.Buffer{}
	multipartWriter := multipart.NewWriter(&buf)
	var err error

	if client.apiUser != "" {
		err = multipartWriter.WriteField("apiUser", client.apiUser)
		if err!= nil {
			return multipartWriter,nil,err
		}
	}

	if client.apiKey != "" {
		err = multipartWriter.WriteField("apiKey", client.apiKey)
		if err!= nil {
			return multipartWriter,nil,err
		}
	}

	err = e.multipart(multipartWriter)
	if err != nil {
		return multipartWriter,nil, err
	}

	if e.Html!= "" {
		err = multipartWriter.WriteField("html", e.Html)
        if err!= nil {
            return multipartWriter,nil, err
        }
	}

	if e.Plain!= "" {
		err = multipartWriter.WriteField("plain", e.Plain)
        if err!= nil {
            return multipartWriter,nil, err
        }
	}

	multipartWriter.Close()
	return multipartWriter,&buf, nil
}