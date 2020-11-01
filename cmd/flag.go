package cmd

import "github.com/spf13/cobra"

type flag struct {
}

func (f *flag) Init(cmd *cobra.Command) {
}

func (f *flag) Validate() error {
	return nil
}
