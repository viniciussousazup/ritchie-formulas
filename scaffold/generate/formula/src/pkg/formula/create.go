package formula

import (
	"path"
	"strings"
)

type Create struct {
	FormulaCmd string
	Lang       string
	OutPutDir  string
}

func (c Create) FormulaPath() string {
	cc := strings.Split(c.FormulaCmd, " ")
	formulaPath := strings.Join(cc[1:], "/")
	return path.Join(c.OutPutDir, formulaPath)
}

// FormulaName remove rit from formulaCmd
func (c Create) FormulaCmdName() string {
	d := strings.Split(c.FormulaCmd, " ")
	return strings.Join(d[1:], " ")
}
