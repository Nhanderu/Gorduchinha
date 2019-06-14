package service

import (
	"bytes"
	"regexp"
	"strconv"

	"github.com/Nhanderu/gorduchinha/src/domain"
	"github.com/Nhanderu/gorduchinha/src/domain/entity"
	"github.com/andybalholm/cascadia"
	"github.com/pkg/errors"
	"golang.org/x/net/html"
)

type scraperService struct {
	svc *Service
}

func newScraperService(svc *Service) scraperService {
	return scraperService{
		svc: svc,
	}
}

func (s scraperService) ScrapeAndUpdate() error {

	teams, err := s.scrapeAll()
	if err != nil {
		return errors.WithStack(err)
	}

	tx, err := s.svc.db.Begin()
	if err != nil {
		return errors.WithStack(err)
	}
	defer tx.Rollback()

	err = tx.Trophy().DeleteAll()
	if err != nil {
		return errors.WithStack(err)
	}

	for _, team := range teams {
		for _, trophy := range team.Trophies {
			err := tx.Trophy().Insert(team.ID, trophy)
			if err != nil {
				return errors.WithStack(err)
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

type scraperFn func() (map[string][]int, error)

func (s scraperService) scrapeAll() ([]entity.Team, error) {

	scrapers := map[string]scraperFn{
		domain.ChampSlugNationalLeague1Div:    s.scrapeNationalLeague1Div,
		domain.ChampSlugNationalLeague2Div:    s.scrapeNationalLeague2Div,
		domain.ChampSlugNationalCup:           s.scrapeNationalCup,
		domain.ChampSlugWorldCup:              s.scrapeWorldCup,
		domain.ChampSlugIntercontinentalCup:   s.scrapeIntercontinentalCup,
		domain.ChampSlugSouthAmericanCupA:     s.scrapeSouthAmericanCupA,
		domain.ChampSlugSouthAmericanCupB:     s.scrapeSouthAmericanCupB,
		domain.ChampSlugSouthAmericanSupercup: s.scrapeSouthAmericanSupercup,
		domain.ChampSlugSPStateCup:            s.scrapeSPStateCup,
		domain.ChampSlugRJStateCup:            s.scrapeRJStateCup,
		domain.ChampSlugRSStateCup:            s.scrapeRSStateCup,
		domain.ChampSlugMGStateCup:            s.scrapeMGStateCup,
	}

	allTrophies := make(map[string][]entity.Trophy)
	for champSlug, scraper := range scrapers {

		champ, err := s.svc.Champ.FindBySlug(champSlug)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		champTrophies, err := scraper()
		if err != nil {
			return nil, errors.WithStack(err)
		}

		for team, years := range champTrophies {
			for _, year := range years {
				allTrophies[team] = append(allTrophies[team], entity.Trophy{
					Year:  year,
					Champ: champ,
				})
			}
		}

	}

	teams := make([]entity.Team, 0)
	for teamAbbr, trophies := range allTrophies {

		team, err := s.svc.Team.FindByAbbr(teamAbbr)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		team.Trophies = trophies
		teams = append(teams, team)
	}

	return teams, nil
}

var (
	reYears  = regexp.MustCompile(`\d{4}`)
	mapTeams = map[string]string{
		"Corinthians":      domain.TeamAbbrSCCP,
		"Palmeiras":        domain.TeamAbbrSEP,
		"São Paulo":        domain.TeamAbbrSPFC,
		"Santos":           domain.TeamAbbrSFC,
		"Flamengo":         domain.TeamAbbrCRF,
		"Vasco da Gama":    domain.TeamAbbrCRVG,
		"Vasco":            domain.TeamAbbrCRVG,
		"Fluminense":       domain.TeamAbbrFFC,
		"Botafogo":         domain.TeamAbbrBFR,
		"Atlético Mineiro": domain.TeamAbbrCAM,
		"Cruzeiro":         domain.TeamAbbrCEC,
		"Grêmio":           domain.TeamAbbrGFBPA,
		"Internacional":    domain.TeamAbbrIEC,
	}
)

func (s scraperService) scrape(champSlug string, url string, linesSel, teamSel, yearsSel cascadia.Selector) (map[string][]int, error) {

	trophies := make(map[string][]int)

	res, err := s.svc.httpClient.Get(url)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	doc, err := html.Parse(res.Body)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	lines := linesSel.MatchAll(doc)
	for _, line := range lines {

		rawTeam := teamSel.MatchFirst(line)
		if rawTeam == nil {
			continue
		}

		teamName := innerText(rawTeam)
		teamAbbr, found := mapTeams[teamName]
		if !found {
			continue
		}

		rawYears := innerText(yearsSel.MatchFirst(line))
		years := reYears.FindAllString(rawYears, -1)
		if len(years) < 1 {
			continue
		}

		teamTrophies := make([]int, len(years))
		for i := range teamTrophies {
			teamTrophies[i], err = strconv.Atoi(years[i])
			if err != nil {
				s.svc.log.Errorf("Error scraping title %s for %s: %s.",
					champSlug,
					teamAbbr,
					err.Error(),
				)
				continue
			}
		}

		trophies[teamAbbr] = teamTrophies
	}

	return trophies, nil
}

func (s scraperService) scrapeNationalLeague1Div() (map[string][]int, error) {

	const (
		url = "https://pt.wikipedia.org/wiki/Lista_de_campe%C3%B5es_do_Campeonato_Brasileiro_de_Futebol"
	)

	var (
		lines = cascadia.MustCompile("h3:has(#Por_clube) + table > tbody > tr:not(:first-child)")
		team  = cascadia.MustCompile("td:first-child > span > a:last-child")
		years = cascadia.MustCompile("td:nth-child(6)")
	)

	trophies, err := s.scrape(domain.ChampSlugNationalLeague1Div, url, lines, team, years)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return trophies, nil
}

func (s scraperService) scrapeNationalLeague2Div() (map[string][]int, error) {

	const (
		url = "https://pt.wikipedia.org/wiki/Campeonato_Brasileiro_de_Futebol_-_S%C3%A9rie_B"
	)

	var (
		lines = cascadia.MustCompile("h3:has(#Títulos_por_clube) + table > tbody > tr")
		team  = cascadia.MustCompile("td:first-child > a:last-child")
		years = cascadia.MustCompile("td:nth-child(2)")
	)

	trophies, err := s.scrape(domain.ChampSlugNationalLeague1Div, url, lines, team, years)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return trophies, nil
}

func (s scraperService) scrapeNationalCup() (map[string][]int, error) {

	const (
		url = "https://pt.wikipedia.org/wiki/Lista_de_campe%C3%B5es_da_Copa_do_Brasil_de_Futebol"
	)

	var (
		lines = cascadia.MustCompile("h2:has(#Resultados_por_clube) + table > tbody > tr")
		team  = cascadia.MustCompile("td:first-child > a:last-child")
		years = cascadia.MustCompile("td:nth-child(4)")
	)

	trophies, err := s.scrape(domain.ChampSlugNationalLeague1Div, url, lines, team, years)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return trophies, nil
}

func (s scraperService) scrapeWorldCup() (map[string][]int, error) {

	const (
		url = "https://pt.wikipedia.org/wiki/Copa_do_Mundo_de_Clubes_da_FIFA"
	)

	var (
		lines = cascadia.MustCompile("h3:has(#Por_clube) ~ table > tbody > tr")
		team  = cascadia.MustCompile("td:first-child > a:last-child")
		years = cascadia.MustCompile("td:nth-child(2)")
	)

	trophies, err := s.scrape(domain.ChampSlugNationalLeague1Div, url, lines, team, years)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return trophies, nil
}

func (s scraperService) scrapeIntercontinentalCup() (map[string][]int, error) {

	const (
		url = "https://pt.wikipedia.org/wiki/Copa_Intercontinental"
	)

	var (
		lines = cascadia.MustCompile("h3:has(#Por_clube) + table > tbody > tr")
		team  = cascadia.MustCompile("td:first-child > a:last-child")
		years = cascadia.MustCompile("td:nth-child(2)")
	)

	trophies, err := s.scrape(domain.ChampSlugNationalLeague1Div, url, lines, team, years)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return trophies, nil
}

func (s scraperService) scrapeSouthAmericanCupA() (map[string][]int, error) {

	const (
		url = "https://pt.wikipedia.org/wiki/Lista_de_campe%C3%B5es_da_Copa_Libertadores_da_Am%C3%A9rica"
	)

	var (
		lines = cascadia.MustCompile("h2:has(#Títulos_e_vice_por_equipe) + table > tbody > tr")
		team  = cascadia.MustCompile("td:first-child > a")
		years = cascadia.MustCompile("td:nth-child(3)")
	)

	trophies, err := s.scrape(domain.ChampSlugNationalLeague1Div, url, lines, team, years)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return trophies, nil
}

func (s scraperService) scrapeSouthAmericanCupB() (map[string][]int, error) {

	const (
		url = "https://pt.wikipedia.org/wiki/Copa_Sul-Americana"
	)

	var (
		lines = cascadia.MustCompile("h3:has(#Títulos_por_clube) ~ table > tbody > tr")
		team  = cascadia.MustCompile("td:first-child > a")
		years = cascadia.MustCompile("td:nth-child(3)")
	)

	trophies, err := s.scrape(domain.ChampSlugNationalLeague1Div, url, lines, team, years)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return trophies, nil
}

func (s scraperService) scrapeSouthAmericanSupercup() (map[string][]int, error) {

	const (
		url = "https://pt.wikipedia.org/wiki/Recopa_Sul-Americana"
	)

	var (
		lines = cascadia.MustCompile("h3:has(#Por_equipe) + table > tbody > tr")
		team  = cascadia.MustCompile("td:first-child > a")
		years = cascadia.MustCompile("td:nth-child(3)")
	)

	trophies, err := s.scrape(domain.ChampSlugNationalLeague1Div, url, lines, team, years)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return trophies, nil
}

func (s scraperService) scrapeSPStateCup() (map[string][]int, error) {

	const (
		url = "https://pt.wikipedia.org/wiki/Campeonato_Paulista_de_Futebol"
	)

	var (
		lines = cascadia.MustCompile("h3:has(#Títulos_por_clube) + table > tbody > tr:not(:first-child)")
		team  = cascadia.MustCompile("td:first-child > b > a:last-child")
		years = cascadia.MustCompile("td:nth-child(3)")
	)

	trophies, err := s.scrape(domain.ChampSlugNationalLeague1Div, url, lines, team, years)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return trophies, nil
}

func (s scraperService) scrapeRJStateCup() (map[string][]int, error) {

	const (
		url = "https://pt.wikipedia.org/wiki/Campeonato_Carioca_de_Futebol"
	)

	var (
		lines = cascadia.MustCompile("h3:has(#Títulos_por_clube) ~ table > tbody > tr")
		team  = cascadia.MustCompile("td:first-child > a")
		years = cascadia.MustCompile("td:nth-child(2)")
	)

	trophies, err := s.scrape(domain.ChampSlugNationalLeague1Div, url, lines, team, years)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return trophies, nil
}

func (s scraperService) scrapeRSStateCup() (map[string][]int, error) {

	const (
		url = "https://pt.wikipedia.org/wiki/Campeonato_Ga%C3%BAcho_de_Futebol"
	)

	var (
		lines = cascadia.MustCompile("h3:has(#Por_equipe) + table > tbody > tr:not(:first-child)")
		team  = cascadia.MustCompile("td:first-child > b > a:last-child")
		years = cascadia.MustCompile("td:nth-child(3)")
	)

	trophies, err := s.scrape(domain.ChampSlugNationalLeague1Div, url, lines, team, years)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return trophies, nil
}

func (s scraperService) scrapeMGStateCup() (map[string][]int, error) {

	const (
		url = "https://pt.wikipedia.org/wiki/Campeonato_Mineiro_de_Futebol"
	)

	var (
		lines = cascadia.MustCompile("h3:has(#Por_equipe) + table > tbody > tr")
		team  = cascadia.MustCompile("td:first-child > b > a")
		years = cascadia.MustCompile("td:nth-child(3)")
	)

	trophies, err := s.scrape(domain.ChampSlugNationalLeague1Div, url, lines, team, years)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return trophies, nil
}

func innerText(node *html.Node) string {

	if node.FirstChild == nil {
		return node.Data
	}

	buffer := bytes.NewBufferString("")
	child := node.FirstChild
	for {

		if child.FirstChild == nil {
			buffer.WriteString(child.Data)
		} else {
			buffer.WriteString(innerText(child))
		}

		if child.NextSibling == nil || child == node.LastChild {
			break
		}

		child = child.NextSibling
	}

	return buffer.String()
}
