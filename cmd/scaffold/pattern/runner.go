package pattern

import (
	"context"
	"fmt"
	"html/template"
	"io"
	"os"
	"path"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"

	scaffoldfiles "github.com/marcelmue/konstrukt/pkg/scaffold"
)

type runner struct {
	flag   *flag
	logger micrologger.Logger
	stdout io.Writer
	stderr io.Writer
}

type input struct {
	// Path is the absolute path of the file to be generated from this
	// Input.
	Path string
	// TemplateBody is the Go text template from which the file is
	// generated.
	TemplateBody string
	// TemplateData defines data for the template defined in TemplateBody.
	TemplateData interface{}
}

func (r *runner) Run(cmd *cobra.Command, args []string) error {
	ctx := context.Background()

	err := r.flag.Validate()
	if err != nil {
		return microerror.Mask(err)
	}

	err = r.run(ctx, cmd, args)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}

func (r *runner) run(ctx context.Context, cmd *cobra.Command, args []string) error {
	files := []input{
		{
			Path:         fmt.Sprintf("./cmd/gen/%s/command.go", r.flag.Name),
			TemplateBody: scaffoldfiles.PatternCommandTemplate,
			TemplateData: *r.flag,
		},
		{
			Path:         fmt.Sprintf("./cmd/gen/%s/command_test.go", r.flag.Name),
			TemplateBody: scaffoldfiles.PatternCommandTestTemplate,
			TemplateData: *r.flag,
		},
		{
			Path:         fmt.Sprintf("./cmd/gen/%s/error.go", r.flag.Name),
			TemplateBody: scaffoldfiles.PatternErrorTemplate,
			TemplateData: *r.flag,
		},
		{
			Path:         fmt.Sprintf("./cmd/gen/%s/flag.go", r.flag.Name),
			TemplateBody: scaffoldfiles.PatternFlagTemplate,
			TemplateData: *r.flag,
		},
		{
			Path:         fmt.Sprintf("./cmd/gen/%s/runner.go", r.flag.Name),
			TemplateBody: scaffoldfiles.PatternRunnerTemplate,
			TemplateData: *r.flag,
		},
	}

	for _, f := range files {
		err := execute(ctx, f)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	return nil
}

func execute(ctx context.Context, file input) error {
	// Create the file's directory if it doesn't exist.
	{
		dir := path.Dir(file.Path)
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	w, err := os.OpenFile(file.Path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return microerror.Mask(err)
	}
	defer w.Close()

	err = executeTemplate(ctx, w, file)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}

func executeTemplate(ctx context.Context, w io.Writer, f input) error {
	var err error

	tmpl := template.New(fmt.Sprintf("%T", f))

	tmpl, err = tmpl.Parse(f.TemplateBody)
	if err != nil {
		return microerror.Mask(err)
	}

	err = tmpl.Execute(w, f.TemplateData)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}
