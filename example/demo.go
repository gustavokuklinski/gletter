package main

import (
	"github.com/gustavokuklinski/gletter"
	"html/template"
	"net/http"
)

// Parse templates from folder templates/
var Tmpl = template.Must(template.ParseGlob("templates/*"))

// Index has the Form to be procesed
func Index(w http.ResponseWriter, r *http.Request) {
	Tmpl.ExecuteTemplate(w, "Index", nil)
}

// Send is called by Index function form to send e-mails
func Send(w http.ResponseWriter, r *http.Request) {

	// Request the Form method
	if r.Method == "POST" {

		// Request fields from HTML Form
		fname := r.FormValue("name")
		femail := r.FormValue("email")
		fphone := r.FormValue("phone")
		fmessage := r.FormValue("message")

		// Set the HTML template
		emailTmpl := "<html><body><h1>gletter Demo</h1>Name:" + fname + "<br />E-Mail:" + femail + "<br />Phone:" + fphone + "<br />Message:" + fmessage + "</body></html>"

		// Call gletter package. Read the parameters.
		// If you don't want to use all params, leave it as a
		// blank string
		gletter.SendMail(
			"smtp.server.com",           // SMTP server
			"***********",               // Password
			"sender@email.com",          // E-mail Sender
			"receiver@email.com",        // E-mail Receiver
			"This mail send by gletter", // Message subject
			587,       // SMTP port
			emailTmpl, // HTML template of the e-mail
		)

		// Render the success template
		mTmpl := template.New("Gletter Mail")
		mTmpl, _ = mTmpl.Parse("Succefull Send! :D")
		mTmpl.Execute(w, nil)

	}
}

func main() {

	http.HandleFunc("/", Index)
	http.HandleFunc("/send", Send)

	http.ListenAndServe(":9000", nil)

}
