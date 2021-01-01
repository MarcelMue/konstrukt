package pattern

import (
	"github.com/giantswarm/microerror"
	"github.com/spf13/cobra"
)

const (
	flagAuthor = "author"
	flagName   = "name"
)

type flag struct {
	Author string
	Name   string
}

func (f *flag) Init(cmd *cobra.Command) {
	cmd.Flags().StringVar(&f.Author, flagAuthor, "MarcelMue", `Name of the author to scaffold.`)
	cmd.Flags().StringVar(&f.Name, flagName, "", `Name of the pattern to scaffold.`)
}

func (f *flag) Validate() error {
	if f.Name == "" {
		return microerror.Mask(invalidFlagError)
	}

	return nil
}
