package depdep_test

import (
	"testing"

	"github.com/gostaticanalysis/testutil"
	"github.com/shota3506/depdep"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	analyzer := depdep.Analyzer
	analyzer.Flags.Set("config", "testconfig.yaml")

	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.Run(t, testdata, analyzer, "foo/bar", "foo/baz")
}
