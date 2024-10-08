package github

import "time"

type Repository struct {
	Id       int    `json:"id"`
	NodeId   string `json:"node_id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Owner    struct {
		Login             string `json:"login"`
		Id                int    `json:"id"`
		NodeId            string `json:"node_id"`
		AvatarUrl         string `json:"avatar_url"`
		GravatarId        string `json:"gravatar_id"`
		Url               string `json:"url"`
		HtmlUrl           string `json:"html_url"`
		FollowersUrl      string `json:"followers_url"`
		FollowingUrl      string `json:"following_url"`
		GistsUrl          string `json:"gists_url"`
		StarredUrl        string `json:"starred_url"`
		SubscriptionsUrl  string `json:"subscriptions_url"`
		OrganizationsUrl  string `json:"organizations_url"`
		ReposUrl          string `json:"repos_url"`
		EventsUrl         string `json:"events_url"`
		ReceivedEventsUrl string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"owner"`
	Private          bool        `json:"private"`
	HtmlUrl          string      `json:"html_url"`
	Description      string      `json:"description"`
	Fork             bool        `json:"fork"`
	Url              string      `json:"url"`
	ArchiveUrl       string      `json:"archive_url"`
	AssigneesUrl     string      `json:"assignees_url"`
	BlobsUrl         string      `json:"blobs_url"`
	BranchesUrl      string      `json:"branches_url"`
	CollaboratorsUrl string      `json:"collaborators_url"`
	CommentsUrl      string      `json:"comments_url"`
	CommitsUrl       string      `json:"commits_url"`
	CompareUrl       string      `json:"compare_url"`
	ContentsUrl      string      `json:"contents_url"`
	ContributorsUrl  string      `json:"contributors_url"`
	DeploymentsUrl   string      `json:"deployments_url"`
	DownloadsUrl     string      `json:"downloads_url"`
	EventsUrl        string      `json:"events_url"`
	ForksUrl         string      `json:"forks_url"`
	GitCommitsUrl    string      `json:"git_commits_url"`
	GitRefsUrl       string      `json:"git_refs_url"`
	GitTagsUrl       string      `json:"git_tags_url"`
	GitUrl           string      `json:"git_url"`
	IssueCommentUrl  string      `json:"issue_comment_url"`
	IssueEventsUrl   string      `json:"issue_events_url"`
	IssuesUrl        string      `json:"issues_url"`
	KeysUrl          string      `json:"keys_url"`
	LabelsUrl        string      `json:"labels_url"`
	LanguagesUrl     string      `json:"languages_url"`
	MergesUrl        string      `json:"merges_url"`
	MilestonesUrl    string      `json:"milestones_url"`
	NotificationsUrl string      `json:"notifications_url"`
	PullsUrl         string      `json:"pulls_url"`
	ReleasesUrl      string      `json:"releases_url"`
	SshUrl           string      `json:"ssh_url"`
	StargazersUrl    string      `json:"stargazers_url"`
	StatusesUrl      string      `json:"statuses_url"`
	SubscribersUrl   string      `json:"subscribers_url"`
	SubscriptionUrl  string      `json:"subscription_url"`
	TagsUrl          string      `json:"tags_url"`
	TeamsUrl         string      `json:"teams_url"`
	TreesUrl         string      `json:"trees_url"`
	CloneUrl         string      `json:"clone_url"`
	MirrorUrl        string      `json:"mirror_url"`
	HooksUrl         string      `json:"hooks_url"`
	SvnUrl           string      `json:"svn_url"`
	Homepage         string      `json:"homepage"`
	Language         interface{} `json:"language"`
	ForksCount       int         `json:"forks_count"`
	StargazersCount  int         `json:"stargazers_count"`
	WatchersCount    int         `json:"watchers_count"`
	Size             int         `json:"size"`
	DefaultBranch    string      `json:"default_branch"`
	OpenIssuesCount  int         `json:"open_issues_count"`
	IsTemplate       bool        `json:"is_template"`
	Topics           []string    `json:"topics"`
	HasIssues        bool        `json:"has_issues"`
	HasProjects      bool        `json:"has_projects"`
	HasWiki          bool        `json:"has_wiki"`
	HasPages         bool        `json:"has_pages"`
	HasDownloads     bool        `json:"has_downloads"`
	Archived         bool        `json:"archived"`
	Disabled         bool        `json:"disabled"`
	Visibility       string      `json:"visibility"`
	PushedAt         time.Time   `json:"pushed_at"`
	CreatedAt        time.Time   `json:"created_at"`
	UpdatedAt        time.Time   `json:"updated_at"`
	Permissions      struct {
		Admin bool `json:"admin"`
		Push  bool `json:"push"`
		Pull  bool `json:"pull"`
	} `json:"permissions"`
	AllowRebaseMerge    bool        `json:"allow_rebase_merge"`
	TemplateRepository  interface{} `json:"template_repository"`
	TempCloneToken      string      `json:"temp_clone_token"`
	AllowSquashMerge    bool        `json:"allow_squash_merge"`
	AllowAutoMerge      bool        `json:"allow_auto_merge"`
	DeleteBranchOnMerge bool        `json:"delete_branch_on_merge"`
	AllowMergeCommit    bool        `json:"allow_merge_commit"`
	SubscribersCount    int         `json:"subscribers_count"`
	NetworkCount        int         `json:"network_count"`
	License             struct {
		Key     string `json:"key"`
		Name    string `json:"name"`
		Url     string `json:"url"`
		SpdxId  string `json:"spdx_id"`
		NodeId  string `json:"node_id"`
		HtmlUrl string `json:"html_url"`
	} `json:"license"`
	Forks      int `json:"forks"`
	OpenIssues int `json:"open_issues"`
	Watchers   int `json:"watchers"`
}

type Commit struct {
	Url         string `json:"url"`
	Sha         string `json:"sha"`
	NodeId      string `json:"node_id"`
	HtmlUrl     string `json:"html_url"`
	CommentsUrl string `json:"comments_url"`
	Commit      struct {
		Url    string `json:"url"`
		Author struct {
			Name  string    `json:"name"`
			Email string    `json:"email"`
			Date  time.Time `json:"date"`
		} `json:"author"`
		Committer struct {
			Name  string    `json:"name"`
			Email string    `json:"email"`
			Date  time.Time `json:"date"`
		} `json:"committer"`
		Message string `json:"message"`
		Tree    struct {
			Url string `json:"url"`
			Sha string `json:"sha"`
		} `json:"tree"`
		CommentCount int `json:"comment_count"`
		Verification struct {
			Verified  bool        `json:"verified"`
			Reason    string      `json:"reason"`
			Signature interface{} `json:"signature"`
			Payload   interface{} `json:"payload"`
		} `json:"verification"`
	} `json:"commit"`
	Author struct {
		Login             string `json:"login"`
		Id                int    `json:"id"`
		NodeId            string `json:"node_id"`
		AvatarUrl         string `json:"avatar_url"`
		GravatarId        string `json:"gravatar_id"`
		Url               string `json:"url"`
		HtmlUrl           string `json:"html_url"`
		FollowersUrl      string `json:"followers_url"`
		FollowingUrl      string `json:"following_url"`
		GistsUrl          string `json:"gists_url"`
		StarredUrl        string `json:"starred_url"`
		SubscriptionsUrl  string `json:"subscriptions_url"`
		OrganizationsUrl  string `json:"organizations_url"`
		ReposUrl          string `json:"repos_url"`
		EventsUrl         string `json:"events_url"`
		ReceivedEventsUrl string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"author"`
	Committer struct {
		Login             string `json:"login"`
		Id                int    `json:"id"`
		NodeId            string `json:"node_id"`
		AvatarUrl         string `json:"avatar_url"`
		GravatarId        string `json:"gravatar_id"`
		Url               string `json:"url"`
		HtmlUrl           string `json:"html_url"`
		FollowersUrl      string `json:"followers_url"`
		FollowingUrl      string `json:"following_url"`
		GistsUrl          string `json:"gists_url"`
		StarredUrl        string `json:"starred_url"`
		SubscriptionsUrl  string `json:"subscriptions_url"`
		OrganizationsUrl  string `json:"organizations_url"`
		ReposUrl          string `json:"repos_url"`
		EventsUrl         string `json:"events_url"`
		ReceivedEventsUrl string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"committer"`
	Parents []struct {
		Url string `json:"url"`
		Sha string `json:"sha"`
	} `json:"parents"`
}

type UserInfo struct {
	Login                   string    `json:"login"`
	Id                      int       `json:"id"`
	NodeId                  string    `json:"node_id"`
	AvatarUrl               string    `json:"avatar_url"`
	GravatarId              string    `json:"gravatar_id"`
	Url                     string    `json:"url"`
	HtmlUrl                 string    `json:"html_url"`
	FollowersUrl            string    `json:"followers_url"`
	FollowingUrl            string    `json:"following_url"`
	GistsUrl                string    `json:"gists_url"`
	StarredUrl              string    `json:"starred_url"`
	SubscriptionsUrl        string    `json:"subscriptions_url"`
	OrganizationsUrl        string    `json:"organizations_url"`
	ReposUrl                string    `json:"repos_url"`
	EventsUrl               string    `json:"events_url"`
	ReceivedEventsUrl       string    `json:"received_events_url"`
	Type                    string    `json:"type"`
	SiteAdmin               bool      `json:"site_admin"`
	Name                    string    `json:"name"`
	Company                 string    `json:"company"`
	Blog                    string    `json:"blog"`
	Location                string    `json:"location"`
	Email                   string    `json:"email"`
	Hireable                bool      `json:"hireable"`
	Bio                     string    `json:"bio"`
	TwitterUsername         string    `json:"twitter_username"`
	PublicRepos             int       `json:"public_repos"`
	PublicGists             int       `json:"public_gists"`
	Followers               int       `json:"followers"`
	Following               int       `json:"following"`
	CreatedAt               time.Time `json:"created_at"`
	UpdatedAt               time.Time `json:"updated_at"`
	PrivateGists            int       `json:"private_gists"`
	TotalPrivateRepos       int       `json:"total_private_repos"`
	OwnedPrivateRepos       int       `json:"owned_private_repos"`
	DiskUsage               int       `json:"disk_usage"`
	Collaborators           int       `json:"collaborators"`
	TwoFactorAuthentication bool      `json:"two_factor_authentication"`
	Plan                    struct {
		Name          string `json:"name"`
		Space         int    `json:"space"`
		PrivateRepos  int    `json:"private_repos"`
		Collaborators int    `json:"collaborators"`
	} `json:"plan"`
}

type Issue struct {
	Id            int    `json:"id"`
	NodeId        string `json:"node_id"`
	Url           string `json:"url"`
	RepositoryUrl string `json:"repository_url"`
	LabelsUrl     string `json:"labels_url"`
	CommentsUrl   string `json:"comments_url"`
	EventsUrl     string `json:"events_url"`
	HtmlUrl       string `json:"html_url"`
	Number        int    `json:"number"`
	State         string `json:"state"`
	Title         string `json:"title"`
	Body          string `json:"body"`
	User          struct {
		Login             string `json:"login"`
		Id                int    `json:"id"`
		NodeId            string `json:"node_id"`
		AvatarUrl         string `json:"avatar_url"`
		GravatarId        string `json:"gravatar_id"`
		Url               string `json:"url"`
		HtmlUrl           string `json:"html_url"`
		FollowersUrl      string `json:"followers_url"`
		FollowingUrl      string `json:"following_url"`
		GistsUrl          string `json:"gists_url"`
		StarredUrl        string `json:"starred_url"`
		SubscriptionsUrl  string `json:"subscriptions_url"`
		OrganizationsUrl  string `json:"organizations_url"`
		ReposUrl          string `json:"repos_url"`
		EventsUrl         string `json:"events_url"`
		ReceivedEventsUrl string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"user"`
	Labels []struct {
		Id          int    `json:"id"`
		NodeId      string `json:"node_id"`
		Url         string `json:"url"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Color       string `json:"color"`
		Default     bool   `json:"default"`
	} `json:"labels"`
	Assignee struct {
		Login             string `json:"login"`
		Id                int    `json:"id"`
		NodeId            string `json:"node_id"`
		AvatarUrl         string `json:"avatar_url"`
		GravatarId        string `json:"gravatar_id"`
		Url               string `json:"url"`
		HtmlUrl           string `json:"html_url"`
		FollowersUrl      string `json:"followers_url"`
		FollowingUrl      string `json:"following_url"`
		GistsUrl          string `json:"gists_url"`
		StarredUrl        string `json:"starred_url"`
		SubscriptionsUrl  string `json:"subscriptions_url"`
		OrganizationsUrl  string `json:"organizations_url"`
		ReposUrl          string `json:"repos_url"`
		EventsUrl         string `json:"events_url"`
		ReceivedEventsUrl string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"assignee"`
	Assignees []struct {
		Login             string `json:"login"`
		Id                int    `json:"id"`
		NodeId            string `json:"node_id"`
		AvatarUrl         string `json:"avatar_url"`
		GravatarId        string `json:"gravatar_id"`
		Url               string `json:"url"`
		HtmlUrl           string `json:"html_url"`
		FollowersUrl      string `json:"followers_url"`
		FollowingUrl      string `json:"following_url"`
		GistsUrl          string `json:"gists_url"`
		StarredUrl        string `json:"starred_url"`
		SubscriptionsUrl  string `json:"subscriptions_url"`
		OrganizationsUrl  string `json:"organizations_url"`
		ReposUrl          string `json:"repos_url"`
		EventsUrl         string `json:"events_url"`
		ReceivedEventsUrl string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"assignees"`
	Milestone struct {
		Url         string `json:"url"`
		HtmlUrl     string `json:"html_url"`
		LabelsUrl   string `json:"labels_url"`
		Id          int    `json:"id"`
		NodeId      string `json:"node_id"`
		Number      int    `json:"number"`
		State       string `json:"state"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Creator     struct {
			Login             string `json:"login"`
			Id                int    `json:"id"`
			NodeId            string `json:"node_id"`
			AvatarUrl         string `json:"avatar_url"`
			GravatarId        string `json:"gravatar_id"`
			Url               string `json:"url"`
			HtmlUrl           string `json:"html_url"`
			FollowersUrl      string `json:"followers_url"`
			FollowingUrl      string `json:"following_url"`
			GistsUrl          string `json:"gists_url"`
			StarredUrl        string `json:"starred_url"`
			SubscriptionsUrl  string `json:"subscriptions_url"`
			OrganizationsUrl  string `json:"organizations_url"`
			ReposUrl          string `json:"repos_url"`
			EventsUrl         string `json:"events_url"`
			ReceivedEventsUrl string `json:"received_events_url"`
			Type              string `json:"type"`
			SiteAdmin         bool   `json:"site_admin"`
		} `json:"creator"`
		OpenIssues   int       `json:"open_issues"`
		ClosedIssues int       `json:"closed_issues"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
		ClosedAt     time.Time `json:"closed_at"`
		DueOn        time.Time `json:"due_on"`
	} `json:"milestone"`
	Locked           bool   `json:"locked"`
	ActiveLockReason string `json:"active_lock_reason"`
	Comments         int    `json:"comments"`
	PullRequest      struct {
		Url      string `json:"url"`
		HtmlUrl  string `json:"html_url"`
		DiffUrl  string `json:"diff_url"`
		PatchUrl string `json:"patch_url"`
	} `json:"pull_request"`
	ClosedAt  interface{} `json:"closed_at"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	ClosedBy  struct {
		Login             string `json:"login"`
		Id                int    `json:"id"`
		NodeId            string `json:"node_id"`
		AvatarUrl         string `json:"avatar_url"`
		GravatarId        string `json:"gravatar_id"`
		Url               string `json:"url"`
		HtmlUrl           string `json:"html_url"`
		FollowersUrl      string `json:"followers_url"`
		FollowingUrl      string `json:"following_url"`
		GistsUrl          string `json:"gists_url"`
		StarredUrl        string `json:"starred_url"`
		SubscriptionsUrl  string `json:"subscriptions_url"`
		OrganizationsUrl  string `json:"organizations_url"`
		ReposUrl          string `json:"repos_url"`
		EventsUrl         string `json:"events_url"`
		ReceivedEventsUrl string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"closed_by"`
	AuthorAssociation string `json:"author_association"`
	StateReason       string `json:"state_reason"`
}

type Language struct {
	C      int `json:"C"`
	Python int `json:"Python"`
}

type Traffic struct {
	Count   int `json:"count"`
	Uniques int `json:"uniques"`
	Clones  []struct {
		Timestamp time.Time `json:"timestamp"`
		Count     int       `json:"count"`
		Uniques   int       `json:"uniques"`
	} `json:"clones"`
}

type Stargazer struct {
	Login             string `json:"login"`
	Id                int    `json:"id"`
	NodeId            string `json:"node_id"`
	AvatarUrl         string `json:"avatar_url"`
	GravatarId        string `json:"gravatar_id"`
	Url               string `json:"url"`
	HtmlUrl           string `json:"html_url"`
	FollowersUrl      string `json:"followers_url"`
	FollowingUrl      string `json:"following_url"`
	GistsUrl          string `json:"gists_url"`
	StarredUrl        string `json:"starred_url"`
	SubscriptionsUrl  string `json:"subscriptions_url"`
	OrganizationsUrl  string `json:"organizations_url"`
	ReposUrl          string `json:"repos_url"`
	EventsUrl         string `json:"events_url"`
	ReceivedEventsUrl string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

type Fork struct {
	Id       int    `json:"id"`
	NodeId   string `json:"node_id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Owner    struct {
		Login             string `json:"login"`
		Id                int    `json:"id"`
		NodeId            string `json:"node_id"`
		AvatarUrl         string `json:"avatar_url"`
		GravatarId        string `json:"gravatar_id"`
		Url               string `json:"url"`
		HtmlUrl           string `json:"html_url"`
		FollowersUrl      string `json:"followers_url"`
		FollowingUrl      string `json:"following_url"`
		GistsUrl          string `json:"gists_url"`
		StarredUrl        string `json:"starred_url"`
		SubscriptionsUrl  string `json:"subscriptions_url"`
		OrganizationsUrl  string `json:"organizations_url"`
		ReposUrl          string `json:"repos_url"`
		EventsUrl         string `json:"events_url"`
		ReceivedEventsUrl string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"owner"`
	Private          bool        `json:"private"`
	HtmlUrl          string      `json:"html_url"`
	Description      string      `json:"description"`
	Fork             bool        `json:"fork"`
	Url              string      `json:"url"`
	ArchiveUrl       string      `json:"archive_url"`
	AssigneesUrl     string      `json:"assignees_url"`
	BlobsUrl         string      `json:"blobs_url"`
	BranchesUrl      string      `json:"branches_url"`
	CollaboratorsUrl string      `json:"collaborators_url"`
	CommentsUrl      string      `json:"comments_url"`
	CommitsUrl       string      `json:"commits_url"`
	CompareUrl       string      `json:"compare_url"`
	ContentsUrl      string      `json:"contents_url"`
	ContributorsUrl  string      `json:"contributors_url"`
	DeploymentsUrl   string      `json:"deployments_url"`
	DownloadsUrl     string      `json:"downloads_url"`
	EventsUrl        string      `json:"events_url"`
	ForksUrl         string      `json:"forks_url"`
	GitCommitsUrl    string      `json:"git_commits_url"`
	GitRefsUrl       string      `json:"git_refs_url"`
	GitTagsUrl       string      `json:"git_tags_url"`
	GitUrl           string      `json:"git_url"`
	IssueCommentUrl  string      `json:"issue_comment_url"`
	IssueEventsUrl   string      `json:"issue_events_url"`
	IssuesUrl        string      `json:"issues_url"`
	KeysUrl          string      `json:"keys_url"`
	LabelsUrl        string      `json:"labels_url"`
	LanguagesUrl     string      `json:"languages_url"`
	MergesUrl        string      `json:"merges_url"`
	MilestonesUrl    string      `json:"milestones_url"`
	NotificationsUrl string      `json:"notifications_url"`
	PullsUrl         string      `json:"pulls_url"`
	ReleasesUrl      string      `json:"releases_url"`
	SshUrl           string      `json:"ssh_url"`
	StargazersUrl    string      `json:"stargazers_url"`
	StatusesUrl      string      `json:"statuses_url"`
	SubscribersUrl   string      `json:"subscribers_url"`
	SubscriptionUrl  string      `json:"subscription_url"`
	TagsUrl          string      `json:"tags_url"`
	TeamsUrl         string      `json:"teams_url"`
	TreesUrl         string      `json:"trees_url"`
	CloneUrl         string      `json:"clone_url"`
	MirrorUrl        string      `json:"mirror_url"`
	HooksUrl         string      `json:"hooks_url"`
	SvnUrl           string      `json:"svn_url"`
	Homepage         string      `json:"homepage"`
	Language         interface{} `json:"language"`
	ForksCount       int         `json:"forks_count"`
	StargazersCount  int         `json:"stargazers_count"`
	WatchersCount    int         `json:"watchers_count"`
	Size             int         `json:"size"`
	DefaultBranch    string      `json:"default_branch"`
	OpenIssuesCount  int         `json:"open_issues_count"`
	IsTemplate       bool        `json:"is_template"`
	Topics           []string    `json:"topics"`
	HasIssues        bool        `json:"has_issues"`
	HasProjects      bool        `json:"has_projects"`
	HasWiki          bool        `json:"has_wiki"`
	HasPages         bool        `json:"has_pages"`
	HasDownloads     bool        `json:"has_downloads"`
	Archived         bool        `json:"archived"`
	Disabled         bool        `json:"disabled"`
	Visibility       string      `json:"visibility"`
	PushedAt         time.Time   `json:"pushed_at"`
	CreatedAt        time.Time   `json:"created_at"`
	UpdatedAt        time.Time   `json:"updated_at"`
	Permissions      struct {
		Admin bool `json:"admin"`
		Push  bool `json:"push"`
		Pull  bool `json:"pull"`
	} `json:"permissions"`
	TempCloneToken      string `json:"temp_clone_token"`
	DeleteBranchOnMerge bool   `json:"delete_branch_on_merge"`
	SubscribersCount    int    `json:"subscribers_count"`
	NetworkCount        int    `json:"network_count"`
	License             struct {
		Key    string `json:"key"`
		Name   string `json:"name"`
		SpdxId string `json:"spdx_id"`
		Url    string `json:"url"`
		NodeId string `json:"node_id"`
	} `json:"license"`
}
