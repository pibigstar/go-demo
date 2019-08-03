package word

import (
	"github.com/nguyenthenguyen/docx"
	"github.com/unidoc/unioffice/document"
)

// https://github.com/nguyenthenguyen/docx
func ReplaceWord(file string) error {
	r, err := docx.ReadDocxFile(file)
	if err != nil {
		return err
	}
	docx := r.Editable()
	docx.Replace("{name}", "派大星", -1)
	docx.Replace("{age}", "20", -1)
	docx.WriteToFile("new.docx")
	r.Close()

	return nil
}

// https://github.com/unidoc/unioffice
func ReplaceWithStyle(file string) error {

	doc, err := document.Open(file)
	if err != nil {
		return err
	}
	paragraphs := []document.Paragraph{}
	for _, p := range doc.Paragraphs() {
		paragraphs = append(paragraphs, p)
	}

	for _, sdt := range doc.StructuredDocumentTags() {
		for _, p := range sdt.Paragraphs() {
			paragraphs = append(paragraphs, p)
		}
	}

	for _, p := range paragraphs {
		for _, r := range p.Runs() {
			switch r.Text() {
			case "{name}":
				r.ClearContent()
				r.AddText("派大星")
			case "{age}":
				r.ClearContent()
				r.AddText("20")
			default:
				continue
			}
		}
	}

	doc.SaveToFile("new.docx")

	return nil
}
