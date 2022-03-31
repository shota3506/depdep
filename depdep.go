// depdep is a package for static analysis
// to find (internal) dependencies to avoid.
//
// You can define the blocked dependencies in a yaml file.
//
//	blocked:
//		- from: example.com/example
//		- to:
//			- example.com/example/a
//			- example.com/example/b
package depdep

import (
	"go/ast"
	"os"
	"regexp"
	"strings"
	"sync"

	"github.com/goccy/go-yaml"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "depdep finds dependencies to avoid"

// Analyzer finds deprecated dependencies
// based on your custom configuration.
var Analyzer = &analysis.Analyzer{
	Name: "depdep",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

var (
	r    *rule
	file string

	mu       sync.Mutex
	once     sync.Once
	setupErr error
)

func init() {
	Analyzer.Flags.StringVar(&file, "config", "", "configuration file path")
}

func setupRule() (*rule, error) {
	mu.Lock()
	defer mu.Unlock()

	once.Do(func() {
		bytes, err := os.ReadFile(file)
		if err != nil {
			setupErr = err
			return
		}

		var v struct {
			Blocked []*struct {
				From string   `yaml:"from"`
				To   []string `yaml:"to"`
			} `yaml:"blocked"`
		}
		if err := yaml.Unmarshal(bytes, &v); err != nil {
			setupErr = err
			return
		}

		r = &rule{make(map[*regexp.Regexp][]*regexp.Regexp)}
		for _, b := range v.Blocked {
			if err := r.add(b.From, b.To...); err != nil {
				setupErr = err
				return
			}
		}
	})
	if setupErr != nil {
		return nil, setupErr
	}
	return r, nil
}

func run(pass *analysis.Pass) (interface{}, error) {
	rule, err := setupRule()
	if err != nil {
		return nil, err
	}

	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.ImportSpec)(nil),
	}

	pkgPath := pass.Pkg.Path()
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.ImportSpec:
			path := strings.Trim(n.Path.Value, "\"") // trim quotes
			if rule.check(pkgPath, path) {
				pass.Reportf(n.Pos(), "found dependency on %s in %s package", path, pkgPath)
			}
		}
	})

	return nil, nil
}
