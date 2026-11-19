package generator

import (
	"os"
	"text/template"
)



type DocumentData struct {
	Project string
	Hostname   string
	OS      string
	CPU     string
	Memory  string
	Disk       string
	NetworkInterface string
	IPAddress  string
	Middleware string
	Procedure  string
}

func GenerateFile(
	templatePath string,
	outputPath string,
	data DocumentData,
) error {
	tmpl, err := template.ParseFiles(templatePath)
	
	if err != nil {
		return err
	}

	file, err := os.Create(outputPath)
	
	if err != nil {
		return err
	}

	defer file.Close()

	err = tmpl.Execute(file, data)

	if err != nil {
		return err
	}

	return nil
}

func GetProcedure(middleware string) string {
	switch middleware {
		case "nginx":
			return `
sudo dnf install nginx -y
sudo systemctl enable nginx
sudo systemctl start nginx
	`
		case "apache":
			return `
sudo dnf install httpd -y
sudo systemctl enable httpd
sudo systemctl start httpd
	`
		default:
			return "manual installation"
	}
}