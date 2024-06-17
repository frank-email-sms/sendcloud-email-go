package sendcloud_email_go

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/url"
	"strconv"
)


func (client *SendCloud) PrepareReceiverParams(e *Receiver) url.Values {
	params := url.Values{}
	params.Set("apiUser", client.apiUser)
	params.Set("apiKey", client.apiKey)
	params.Set("to", e.To)
	if e.CC!= "" {
		params.Set("cc", e.CC)
	}
	if e.BCC!= "" {
		params.Set("bcc", e.BCC)
	}
	if e.UseAddressList {
		params.Set("useAddressList", strconv.FormatBool(e.UseAddressList))
	}
	return params
}

func (e *MailBody) PrepareMailBodyParams(params *url.Values){
	params.Set("from", e.From)

	params.Set("subject", e.Subject)
	if e.ContentSummary!= "" {
		params.Set("contentSummary", e.ContentSummary)
	}
	if e.FromName!= "" {
		params.Set("fromName", e.FromName)
	}

	if e.ReplyTo!= "" {
		params.Set("replyTo", e.ReplyTo)
	}
	if e.LabelName!= "" {
		params.Set("labelName", e.LabelName)
	}
	if len(e.Headers) > 0 {
		headers, _ := json.Marshal(e.Headers)
		params.Set("headers", string(headers))
	}
	if !e.Xsmtpapi.IsEmpty() {
		xsmtpapi, _ := json.Marshal(e.Xsmtpapi)
		params.Set("xsmtpapi", string(xsmtpapi))
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
}

func (client *SendCloud) PrepareSendCommonEmailParams(e *CommonMail) url.Values {
	params := client.PrepareReceiverParams(&e.To)
	e.body.PrepareMailBodyParams(&params)
	if e.content.Plain!= "" {
		params.Set("plain", e.content.Plain)
	}
	if e.content.Html!= "" {
		params.Set("html", e.content.Html)
	}
	return params
}

func (client *SendCloud) PrepareSendEmailTemplateParams (e *TemplateMail) url.Values {
	params := client.PrepareReceiverParams(&e.To)
	e.body.PrepareMailBodyParams(&params)
	params.Set("templateInvokeName", e.content.TemplateInvokeName)
	return params
}

func (e *Receiver) multipartReceiver(client *SendCloud,multipartWriter *multipart.Writer) error {

	var err error

	if client.apiUser != "" {
		err = multipartWriter.WriteField("apiUser", client.apiUser)
		if err!= nil {
			return err
		}
	}

	if client.apiKey != "" {
		err = multipartWriter.WriteField("apiKey", client.apiKey)
		if err!= nil {
			return err
		}
	}

	if e.To != "" {
		err =   multipartWriter.WriteField("to", e.To)
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

	if e.UseAddressList {
		useAddressListStr := strconv.FormatBool(e.UseAddressList)
		err = multipartWriter.WriteField("useAddressList", useAddressListStr)
		if err != nil {
			return err
		}
	}

	return nil
}



func (e *MailBody) multipartMailBody(multipartWriter *multipart.Writer) error {

	var err error

	var partWriter io.Writer

	if e.From != "" {
		err = multipartWriter.WriteField("from", e.From)
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

	if len(e.Headers) > 0 {
		headers, _ := json.Marshal(e.Headers)
		err = multipartWriter.WriteField("headers", string(headers))
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

	if !e.Xsmtpapi.IsEmpty() {
		xsmtpapi, err := json.Marshal(e.Xsmtpapi)
		if err != nil {
			return err
		}
		err = multipartWriter.WriteField("xsmtpapi", string(xsmtpapi))
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

	return nil
}

func (client *SendCloud) MultipartSendCommonMail(e *CommonMail) (*multipart.Writer,*bytes.Buffer, error) {
	buf := bytes.Buffer{}
	multipartWriter := multipart.NewWriter(&buf)
	var err error

	err = e.To.multipartReceiver(client,multipartWriter)
	if err != nil {
		return multipartWriter,nil, err
	}

	err = e.body.multipartMailBody(multipartWriter)
	if err != nil {
		return multipartWriter,nil, err
	}

	if e.content.Html!= "" {
		err = multipartWriter.WriteField("html", e.content.Html)
        if err!= nil {
            return multipartWriter,nil, err
        }
	}

	if e.content.Plain!= "" {
		err = multipartWriter.WriteField("plain", e.content.Plain)
        if err!= nil {
            return multipartWriter,nil, err
        }
	}

	multipartWriter.Close()
	return multipartWriter,&buf, nil
}

func (client *SendCloud) MultipartSendEmailTemplate(e *TemplateMail) (*multipart.Writer,*bytes.Buffer, error) {
	buf := bytes.Buffer{}
	multipartWriter := multipart.NewWriter(&buf)
	var err error

	err = e.To.multipartReceiver(client,multipartWriter)
	if err != nil {
		return multipartWriter,nil, err
	}

	err = e.body.multipartMailBody(multipartWriter)
	if err != nil {
		return multipartWriter,nil, err
	}

	if e.content.TemplateInvokeName!= "" {
		err = multipartWriter.WriteField("templateInvokeName", e.content.TemplateInvokeName)
		if err!= nil {
			return multipartWriter,nil, err
		}
	}
	multipartWriter.Close()
	return multipartWriter,&buf, nil
}