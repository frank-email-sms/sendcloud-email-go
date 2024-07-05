# sendcloud-email-go

This Go client provides an interface to interact with SendCloud's email delivery service. It exposes two main methods: `SendCommonEmail` for sending regular emails and `SendEmailTemplate` for sending emails using predefined templates.  

## Supported Go Versions  

This SDK supports Go 1.16 and above.

## SendCommonEmail  

The `SendCommonEmail` method allows you to send a regular email with a custom subject, body, and recipient list.  

### Usage

```go  
package main  
  
import (
	"context"
	"fmt"
	"github.com/frank-email-sms/sendcloud-email-go"
	"log"
	"time"
)

func main() {
	client, err := sendcloud.NewSendCloud("*", "*")
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	args :=  &sendcloud.CommonMail{
		Receiver: sendcloud.MailReceiver{
			To: "a@ifaxin.com;b@ifaxin.com",
		},
		Body: sendcloud.MailBody{
			From: "SendCloud@SendCloud.com",
			Subject: "Email from SendCloud SDK",
			FromName: "SendCloud",
		},
		Content: sendcloud.TextContent{
			Html:  "<p>This is an HTML email.</p>",
		},
	}
	result, err := client.SendCommonEmail(ctx, args)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
```

## SendEmailTemplate

The `SendEmailTemplate` method allows you to send an email using a predefined template. This is useful when you want to send emails with consistent design and layout.

### Usage

```go
package main  
  
import (
	"context"
	"fmt"
	"github.com/frank-email-sms/sendcloud-email-go"
	"log"
	"time"
)

func main() {
	client, err := sendcloud.NewSendCloud("*", "*")
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	args :=  &sendcloud.TemplateMail{
		Receiver: sendcloud.MailReceiver{
			To: "a@ifaxin.com;b@ifaxin.com",
		},
		Body: sendcloud.MailBody{
			From: "SendCloud@SendCloud.com",
			Subject: "Email from SendCloud SDK",
			FromName: "SendCloud",
		},
		Content: sendcloud.TemplateContent{
			TemplateInvokeName:  "test_template_active",
		},
	}
	result, err := client.SendEmailTemplate(ctx, args)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
```

## Code Examples  

All code examples for this project are located in the root directory in the `sendEmail_test.go` file. This file contains tests that demonstrate the usage of the email sending functionality and can also be used as examples for integration into your own code.  

## How to View the Examples  

To view the code examples, simply open the `sendEmail_test.go` file in the root directory of the project. You can also run the tests in this file to see the functionality in action.  

## Running the Tests  

To run the tests in `sendEmail_test.go`, navigate to the root directory of the project and run the following command:  

### Configuration

Before using the client, make sure to initialize it with the necessary configuration such as API keys or other credentials. The exact configuration steps may vary depending on the SendCloud client package you are using.

### Handling Errors

Always make sure to handle errors returned by the methods. They may indicate issues such as invalid credentials, API errors, or other problems that need to be addressed.

### Result Handling

The result returned by the methods contains information about the success or failure of the email sending operation. Make sure to check the status code and message to understand the outcome.



