// Code generated by GoRetro; DO NOT EDIT.
package goretro

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/AlterNayte/go-retro/example/github"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

type GithubApiClient struct {
	BaseURL        string
	HTTPClient     *http.Client
	AuthType       AuthType
	Username       string
	Password       string
	APIKey         string
	BearerToken    string
	CustomAuthFunc func() (string, error)
	CustomHeaders  map[string]string
	Timeout        time.Duration
	MaxRetries     int
}

func NewGithubApiClient(baseURL string) *GithubApiClient {
	return &GithubApiClient{
		BaseURL:       baseURL,
		HTTPClient:    &http.Client{},
		AuthType:      AuthNone,
		CustomHeaders: make(map[string]string),
		Timeout:       30 * time.Second,
		MaxRetries:    3,
	}
}

func (c *GithubApiClient) SetBasicAuth(username, password string) {
	c.AuthType = AuthBasic
	c.Username = username
	c.Password = password
}

func (c *GithubApiClient) SetAPIKeyAuth(apiKey string) {
	c.AuthType = AuthAPIKey
	c.APIKey = apiKey
}

func (c *GithubApiClient) SetBearerAuth(token string) {
	c.AuthType = AuthBearer
	c.BearerToken = token
}

func (c *GithubApiClient) SetCustomAuthFunc(authFunc func() (string, error)) {
	c.CustomAuthFunc = authFunc
}

func (c *GithubApiClient) SetCustomHeader(key, value string) {
	c.CustomHeaders[key] = value
}

func (c *GithubApiClient) doRequest(req *http.Request) (*http.Response, error) {
	for k, v := range c.CustomHeaders {
		req.Header.Set(k, v)
	}
	req.Header.Set("Content-Type", "application/json")

	for i := 0; i < c.MaxRetries; i++ {
		c.HTTPClient.Timeout = c.Timeout
		resp, err := c.HTTPClient.Do(req)
		if err != nil {
			if i == c.MaxRetries-1 {
				return nil, err
			}
			time.Sleep(2 * time.Second)
			continue
		}

		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			return resp, nil
		}

		if i == c.MaxRetries-1 {
			bodyBytes, _ := ioutil.ReadAll(resp.Body)
			return nil, fmt.Errorf("request failed with status %d: %s", resp.StatusCode, string(bodyBytes))
		}

		time.Sleep(2 * time.Second)
	}

	return nil, errors.New("max retries exceeded")
}

func (c *GithubApiClient) GetUser() (*github.UserInfo, error) {
	reqURL := c.BaseURL + "/user"

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return nil, err
	}

	if c.CustomAuthFunc != nil {
		token, err := c.CustomAuthFunc()
		if err != nil {
			return nil, err
		}
		req.Header.Set("Authorization", "Bearer "+token)
	} else {
		req.Header.Set("Authorization", "Bearer "+c.BearerToken)
	}

	log.Printf("Making request to %s", req.URL)

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result *github.UserInfo
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

func (c *GithubApiClient) GetRepositories() ([]github.Repository, error) {
	reqURL := c.BaseURL + "/user/repos"

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return nil, err
	}

	log.Printf("Making request to %s", req.URL)

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []github.Repository
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

func (c *GithubApiClient) GetCommits(owner string, repo string, since string) ([]github.Commit, error) {
	reqURL := c.BaseURL + "/repos/%s/%s/commits"

	reqURL = fmt.Sprintf(reqURL, owner, repo)

	query := url.Values{}

	query.Add("since", fmt.Sprintf("%v", since))

	reqURL += "?" + query.Encode()

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return nil, err
	}

	log.Printf("Making request to %s", req.URL)

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []github.Commit
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

func (c *GithubApiClient) GetIssues(owner string, repo string) ([]github.Issue, error) {
	reqURL := c.BaseURL + "/repos/%s/%s/issues"

	reqURL = fmt.Sprintf(reqURL, owner, repo)

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return nil, err
	}

	log.Printf("Making request to %s", req.URL)

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []github.Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

func (c *GithubApiClient) GetLanguages(owner string, repo string) (*github.Language, error) {
	reqURL := c.BaseURL + "/repos/%s/%s/languages"

	reqURL = fmt.Sprintf(reqURL, owner, repo)

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return nil, err
	}

	log.Printf("Making request to %s", req.URL)

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result *github.Language
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

func (c *GithubApiClient) GetTraffic(owner string, repo string) (*github.Traffic, error) {
	reqURL := c.BaseURL + "/repos/%s/%s/traffic/views"

	reqURL = fmt.Sprintf(reqURL, owner, repo)

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return nil, err
	}

	log.Printf("Making request to %s", req.URL)

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result *github.Traffic
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

func (c *GithubApiClient) GetStargazers(owner string, repo string) ([]github.Stargazer, error) {
	reqURL := c.BaseURL + "/repos/%s/stargazers"

	reqURL = fmt.Sprintf(reqURL, repo)

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return nil, err
	}

	log.Printf("Making request to %s", req.URL)

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []github.Stargazer
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

func (c *GithubApiClient) GetForks(owner string, repo string) ([]github.Fork, error) {
	reqURL := c.BaseURL + "/repos/%s/%s/forks"

	reqURL = fmt.Sprintf(reqURL, owner, repo)

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return nil, err
	}

	log.Printf("Making request to %s", req.URL)

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []github.Fork
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}