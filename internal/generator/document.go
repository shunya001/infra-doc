package generator

import (
	"os"
	"text/template"
)



type DocumentData struct {
	Project string
	OS      string
	CPU     string
	Memory  string
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