package create

// Flags represents all the flags that can be set via the command line
type Flags struct {
	SomeFlag string
}

// Validate checks the flags
func (f *Flags) Validate() error {

	return nil
}
