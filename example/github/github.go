package github

type GithubApi struct {
	GetUser         func() (*UserInfo, error)                                       `method:"GET" path:"/user" auth:"bearer"`
	GetRepositories func() ([]Repository, error)                                    `method:"GET" path:"/user/repos"`
	GetCommits      func(owner string, repo string, since string) ([]Commit, error) `method:"GET" path:"/repos/{owner}/{repo}/commits" query:"since"`
	GetIssues       func(owner, repo string) ([]Issue, error)                       `method:"GET" path:"/repos/{owner}/{repo}/issues"`
	GetLanguages    func(owner, repo string) (*Language, error)                     `method:"GET" path:"/repos/{owner}/{repo}/languages"`
	GetTraffic      func(owner, repo string) (*Traffic, error)                      `method:"GET" path:"/repos/{repo}/{owner}/traffic/views"`
	GetStargazers   func(owner, repo string) ([]Stargazer, error)                   `method:"GET" path:"/repos/{repo}/stargazers"`
	GetForks        func(owner, repo string) ([]Fork, error)                        `method:"GET" path:"/repos/{owner}/{repo}/forks"`
}
