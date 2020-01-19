package svc

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type (
	// Currency cache
	CurrencyRates struct {
		Rates map[string]float64
		Base  string
		Date  string
	}

	CurrencyConversor struct {
		sync.Mutex
		Rates   map[string]float64
		current string
		base    float64
		eur     float64
		usd     float64
		pln     float64
		czk     float64
		rub     float64
	}
)

var ()

const (
	// USD, PLN, CZK, RUB
	// rates API URL: https://api.exchangeratesapi.io/latest
	ratesAPIURL = "https://api.exchangeratesapi.io/latest"
)

func (s *Service) UpdateRates() {
	// Currency rates cache asynchronous update
	// NOTE: We can later add some safety mechanisms
	// * Retry after failed requests
	// * Preserve historical data in order to havei
	// a reference vale on failed requests after startup.
	go s.doUpdateRates()
}

func (s *Service) doUpdateRates() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req, err := http.NewRequest("GET", ratesAPIURL, nil)
	if err != nil {
		s.Log.Error(err)
		return err
	}

	clt := http.DefaultClient
	res, err := clt.Do(req.WithContext(ctx))
	if err != nil {
		s.Log.Error(err)
		return err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		s.Log.Error(err)
		return err
	}

	var c CurrencyRates
	err = json.Unmarshal(body, &c)
	if err != nil {
		s.Log.Error(err)
	}

	s.Rates = c

	s.Log.Info("Currency rates updated", "base", c.Base, "date", c.Date)

	return nil
}

// CurrencyConversor methods
func NewCurrencyConversor(rates map[string]float64) *CurrencyConversor {
	return &CurrencyConversor{
		Rates: rates,
	}
}

func (cc *CurrencyConversor) SetAmount(value float64, currency string) {
	if currency == "EUR" {
		cc.setEUR(value)
	}

	if currency == "USD" {
		cc.setUSD(value)
	}

	if currency == "PLN" {
		cc.setPLN(value)
	}

	if currency == "CZK" {
		cc.setCZK(value)
	}

	if currency == "RUB" {
		cc.setRUB(value)
	}
}

func (cc *CurrencyConversor) setEUR(eurValue float64) {
	cc.clear()
	cc.eur = eurValue
	cc.current = "EUR"
}

func (cc *CurrencyConversor) setUSD(usdValue float64) {
	cc.clear()
	cc.usd = usdValue
	cc.current = "USD"
}

func (cc *CurrencyConversor) setPLN(plnValue float64) {
	cc.clear()
	cc.pln = plnValue
	cc.current = "PLN"
}

func (cc *CurrencyConversor) setCZK(czkValue float64) {
	cc.clear()
	cc.czk = czkValue
	cc.current = "CZK"
}

func (cc *CurrencyConversor) setRUB(rubValue float64) {
	cc.clear()
	cc.rub = rubValue
	cc.current = "RUB"
}

func (cc *CurrencyConversor) clear() {
	cc.base = 0
	cc.eur = 0
	cc.usd = 0
	cc.pln = 0
	cc.czk = 0
	cc.rub = 0
	cc.current = ""
}

func (cc *CurrencyConversor) CalculateF32() (result map[string]float32, err error) {
	result = map[string]float32{}

	r, err := cc.Calculate()
	if err != nil {
		return result, err
	}

	for k, v := range r {
		result[k] = float32(v)
	}

	return result, nil
}

func (cc *CurrencyConversor) Calculate() (result map[string]float64, err error) {
	cc.Lock()
	defer cc.Unlock()

	err = cc.checkRates()
	if err != nil {
		return result, err
	}

	usdRate := cc.Rates["USD"]
	plnRate := cc.Rates["PLN"]
	czkRate := cc.Rates["CZK"]
	rubRate := cc.Rates["RUB"]

	switch cc.current {
	case "EUR":
		cc.base = cc.eur

	case "USD":
		cc.base = cc.usd * usdRate

	case "PLN":
		cc.base = cc.pln * plnRate

	case "CZK":
		cc.base = cc.czk * czkRate

	case "RUB":
		cc.base = cc.rub * rubRate

	default:
		return result, errors.New("not a valid currency input")
	}

	return map[string]float64{
		"EUR": cc.base,
		"USD": cc.base * usdRate,
		"PLN": cc.base * plnRate,
		"CZK": cc.base * czkRate,
		"RUB": cc.base * rubRate,
	}, nil

}

func (cc *CurrencyConversor) checkRates() error {
	_, ok0 := cc.Rates["USD"]
	_, ok1 := cc.Rates["PLN"]
	_, ok2 := cc.Rates["CZK"]
	_, ok3 := cc.Rates["RUB"]

	if !(ok0 && ok1 && ok2 && ok3) {
		return errors.New("no currency rates available")
	}

	return nil
}
