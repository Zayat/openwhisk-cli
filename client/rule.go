package client

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type RuleService struct {
	client *Client
}

type Rule struct {
	Namespace string `json:"namespace,omitempty"`
	Name      string `json:"name,omitempty"`
	Version   string `json:"version,omitempty"`
	Publish   bool   `json:"publish,omitempty"`

	Status  string `json:"status"`
	Trigger string `json:"trigger"`
	Rule    string `json:"rule"`
}

type RuleListOptions struct {
	Limit string `url:"limit,omitempty"`
	Skip  int    `url:"skip,omitempty"`
	Docs  bool   `url:"docs,omitempty"`
}

func (s *RuleService) List(options *RuleListOptions) ([]Rule, *http.Response, error) {
	route := fmt.Sprintf("rules")
	route, err := addRouteOptions(route, options)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", route, nil)
	if err != nil {
		return nil, nil, err
	}

	var rules []Rule
	resp, err := s.client.Do(req, &rules)
	if err != nil {
		return nil, resp, err
	}

	return rules, resp, err

}

func (s *RuleService) Create(rule *Rule, blocking bool) (*Rule, *http.Response, error) {
	route := fmt.Sprintf("rules?blocking=%s", blocking)

	req, err := s.client.NewRequest("POST", route, rule)
	if err != nil {
		return nil, nil, err
	}

	r := new(Rule)
	resp, err := s.client.Do(req, &r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil

}

func (s *RuleService) Fetch(ruleName string) (*Rule, *http.Response, error) {
	route := fmt.Sprintf("rules/%s", ruleName)

	req, err := s.client.NewRequest("GET", route, nil)
	if err != nil {
		return nil, nil, err
	}

	r := new(Rule)
	resp, err := s.client.Do(req, &r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil

}

func (s *RuleService) Delete(ruleName string) (*http.Response, error) {
	route := fmt.Sprintf("rules/%s", ruleName)

	req, err := s.client.NewRequest("DELETE", route, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (s *RuleService) Update(rule *Rule, overwrite bool) (*Rule, *http.Response, error) {
	route := fmt.Sprintf("rules/%s?overwrite=", rule.Name, overwrite)

	req, err := s.client.NewRequest("POST", route, rule)
	if err != nil {
		return nil, nil, err
	}

	r := new(Rule)
	resp, err := s.client.Do(req, &r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

func (s *RuleService) SetState(ruleName string, state string) (*Rule, *http.Response, error) {
	state = strings.ToLower(state)
	if state != "enable" && state != "disable" {
		return nil, nil, errors.New(fmt.Sprintf("Invalid state option %s.  Valid options are \"disabled\" and \"enabled\"."), state)
	}

	route := fmt.Sprintf("rules/%s?state=", ruleName, state)

	req, err := s.client.NewRequest("POST", route, nil)
	if err != nil {
		return nil, nil, err
	}

	r := new(Rule)
	resp, err := s.client.Do(req, &r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil

}
