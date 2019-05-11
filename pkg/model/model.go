package model

type TopLevel struct {
	Name            string
	NoMembers       int
	NoMemberInvites int
	NoOauthP        int
	NoUserReg       int
	NoMailServers   int
	NoTLSProfile    int
	NoOrg           int
	Org             *[]Organization
	Azs             *[]Az
}

type Organization struct {
	Name            string
	Catalog         *[]Catalog
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
	Error           string
}
type Az struct {
	Name         string
	NoPortal     int
	NoAnalytics  int
	NoV5cgateway int
	NoApigateway int
}

type Catalog struct {
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
	Webhooks        *[]Webhook
	Tasks           *[]Task
}

type Webhook struct {
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
type Task struct {
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

type UserCreds struct {
	Username string
	Password string
	Realm    string
}

type GlobalVariables struct {
	Server         string
	Orgs           []string
	UserDetails    UserCreds
	UserDetailsOrg UserCreds
	Token          string
	Debug          bool
	Output         string
}
