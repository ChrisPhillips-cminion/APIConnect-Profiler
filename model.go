package main

type topLevel struct {
	Name            string
	NoMembers       int
	NoMemberInvites int
	NoOauthP        int
	NoUserReg       int
	NoMailServers   int
	NoTLSProfile    int
	NoOrg           int
	Org             *[]organization
	Azs             *[]az
}

type organization struct {
	Name            string
	Catalog         *[]catalog
	NoMembers       int
	NoMemberInvites int
	NoDraftAPI      int
	AvgAPISize      int
	MaxAPISize      int
	NoDraftProduct  int
	NoTLSProfile    int
	NoOAuthProvider int
	UserRegistries  int
	NoKeyStore      int
	NoTrustStore    int
}
type az struct {
	Name         string
	NoPortal     int
	NoAnalytics  int
	NoV5cgateway int
	NoApigateway int
}

type catalog struct {
	Name            string
	NoMember        int
	NoMemberInvites int
	NoAPI           int
	NoProduct       int
	AvgAPISize      int
	MaxAPISize      int
	NoSpace         int
	NoConsumerOrg   int
	Portal          bool
	NoTLSProfile    int
	NoOAuthProvider int
	UserRegistries  int
	Applications    int
	Subscriptions   int
	Webhooks        *[]webhook
	Tasks           *[]task
}

type webhook struct {
	WebhookId        string
	OrganizationName string
	CatalogName      string
	Organization     string
	Catalog          string
	State            string
	Created_at       string
	Updated_at       string
	Level            string
	Title            string
}
type task struct {
	TaskId           string
	OrganizationName string
	CatalogName      string
	Organization     string
	Catalog          string
	State            string
	Created_at       string
	Updated_at       string
	Name             string
	Title            string
	Content          map[string]interface{}
}
