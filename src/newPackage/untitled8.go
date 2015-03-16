package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	const response = `
	Dear {{.Honor}} {{.Name}},{{if .Attend}}
Thank you very much for attending our fundraiser.{{else}}
Im sorry that you not able to attend our fundraiser.{{end}}
{{if .Donate}}Thank you for your genrous donation to our cause.{{else}}{{end}}
Be sure to save dates for these upcoming events on your calender!
{{range .Events}} • {{.}}
{{end}}sincerely,
-Tim Schwartz, Event Coordinator
`
	type Recipient struct {
		Name, Honor    string
		Attend, Donate bool
		Events         []string
	}
	var UpEvents = []string{
		"May 24th, Park BBQ, grill the gay away!",
		"June 10th, Join us in throwing rocks at abortion doctors!",
		"June 18th, Cannabis kills conference, help us fight marijuana legalization",
	}
	var Recipients = []Recipient{
		{"Boehner", "Mr.", true, true, UpEvents},
		{"Palin", "Mrs.", false, true, UpEvents},
		{"Cruz", "Mr.", true, false, UpEvents},
		{"Clinton", "Mrs.", false, false, UpEvents},
	}

	t := template.Must(template.New("response").Parse(response))

	for _, r := range Recipients {
		err := t.Execute(os.Stdout, r)
		if err != nil {
			log.Println("executing template:", err)
		}
	}

}
