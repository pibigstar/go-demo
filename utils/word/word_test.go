package word

import "testing"

func TestReplaceWord(t *testing.T) {
	err := ReplaceWord("old.docx")
	if err != nil {
		t.Error(err)
	}
}

func TestReplaceWithStyle(t *testing.T) {
	err := ReplaceWithStyle("old.docx")
	if err != nil {
		t.Error(err)
	}
}
