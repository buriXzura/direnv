package cmd

var IgnoredEnvKeys = map[string]bool{
	"DIRENV_CONFIG":   true,
	"DIRENV_BASH":     true,
	"DIRENV_IN_ENVRC": true,
	"OLDPWD":          true,
	"PWD":             true,
	"PS1":             true,
	"SHELL":           true,
	"SHELLOPTS":       true,
	"SHLVL":           true,
	"COMP_WORDBREAK":  true,
	"_":               true,
}
