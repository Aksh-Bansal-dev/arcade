package player

import (
	"testing"
)

func TestGetFilename(t *testing.T){
	filepath := getFilepath()
	if filepath!="fail"{
		t.Errorf("Filepath: %s", filepath)
	}
}