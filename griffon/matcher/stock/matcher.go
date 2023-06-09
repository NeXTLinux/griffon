package stock

import (
	"github.com/nextlinux/griffon/griffon/distro"
	"github.com/nextlinux/griffon/griffon/match"
	"github.com/nextlinux/griffon/griffon/pkg"
	"github.com/nextlinux/griffon/griffon/search"
	"github.com/nextlinux/griffon/griffon/vulnerability"

	syftPkg "github.com/anchore/syft/syft/pkg"
)

type Matcher struct {
	cfg MatcherConfig
}

type MatcherConfig struct {
	UseCPEs bool
}

func NewStockMatcher(cfg MatcherConfig) *Matcher {
	return &Matcher{
		cfg: cfg,
	}
}

func (m *Matcher) PackageTypes() []syftPkg.Type {
	return nil
}

func (m *Matcher) Type() match.MatcherType {
	return match.StockMatcher
}

func (m *Matcher) Match(store vulnerability.Provider, d *distro.Distro, p pkg.Package) ([]match.Match, error) {
	criteria := search.CommonCriteria
	if m.cfg.UseCPEs {
		criteria = append(criteria, search.ByCPE)
	}
	return search.ByCriteria(store, d, p, m.Type(), criteria...)
}
