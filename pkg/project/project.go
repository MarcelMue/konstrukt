package project

var (
	description = "Command line tool for generating konstruktive art."
	gitSHA      = "n/a"
	name        = "konstrukt"
	patternDesc = "Generated with https://github.com/marcelmue/konstrukt"
	source      = "https://github.com/marcelmue/konstrukt"
	version     = "0.1.0"
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

func PatternDesc() string {
	return patternDesc
}

func Source() string {
	return source
}

func Version() string {
	return version
}
