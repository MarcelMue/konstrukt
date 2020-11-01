package project

var (
	description = "Command line tool for generating konstruktive art."
	gitSHA      = "n/a"
	name        = "konstrukt"
	source      = "https://github.com/marcelmue/konstrukt"
	version     = "0.0.1"
)

func Description() string {
	return description
}

func GitSHA() string {
	return gitSHA
}

func Name() string {
	return name
}

func Source() string {
	return source
}

func Version() string {
	return version
}
