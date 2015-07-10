# GLETTER
Is a small package using `net/smtp` to send Auth SMTP e-mails in the ease way

## Usage
Import the package to your project and set the HTML forms and routes.

Setup your Form requests to fit the template

```go
  fname := r.FormValue("name")
	femail := r.FormValue("email")
	fphone := r.FormValue("phone")
	fmessage := r.FormValue("message")
```


and the basic HTML template(This is going to be send)

```go
  emailTmpl := "<html><body><h1>gletter Demo</h1>Name:" + fname + "<br />E-Mail:" + femail + "<br />Phone:" + fphone + "<br />Message:" + fmessage + "</body></html>"

```

To finish made a basic SMTP setup to send your mails.

```go
gletter.SendMail(
  "smtp.server.com",           // SMTP server
	"***************",           // Password
	"sender@email.com",          // E-mail Sender
	"receiver@email.com",        // E-mail Receiver
	"This mail send by gletter", // Message subject
	587,                         // SMTP port
	emailTmpl,                   // HTML template of the e-mail

)
```

For a more comprehensive demo, please check the `example` folder :)
