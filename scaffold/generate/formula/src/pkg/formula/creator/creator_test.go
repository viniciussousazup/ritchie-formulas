package creator

import (
	"formula/pkg/formula"
	"formula/pkg/formula/template"
	"formula/pkg/stream"
	"os"
	"path"
	"testing"
)

func TestCreateManager_Create(t *testing.T) {

	tplM := template.NewManagerCustom(
		"/home/vinicius.sousa/go/src/github.com/viniciussousazup/ritchie-formulas/scaffold/generate/formula/src/templates",
	)

	file := stream.NewFileManager()
	dir := stream.NewDirManager(file)

	outPutDir := path.Join(os.TempDir(), "test_generate_formula")
	_ = dir.Remove(outPutDir)
	_ = dir.Create(outPutDir)

	type fields struct {
		tplM template.Manager
		dir  stream.DirManager
		file stream.FileManager
	}
	type args struct {
		cf formula.Create
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Run with success",
			fields: fields{
				tplM: tplM,
				file: file,
				dir:  dir,
			},
			args: args{
				cf: formula.Create{
					FormulaCmd: "rit testing hello world",
					Lang:       "go",
					OutPutDir:  outPutDir,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CreateManager{
				tplM: tt.fields.tplM,
				dir:  tt.fields.dir,
				file: tt.fields.file,
			}
			if err := c.Create(tt.args.cf); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
