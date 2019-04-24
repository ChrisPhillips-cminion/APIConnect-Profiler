package main

type topLevel struct {
	name            string
	noMembers       int
	noMemberInvites int
	noOauthP        int
	noUserReg       int
	noMailServers   int
	noTLSProfile    int
	noOrg           int
	org             *[]organization
	azs             *[]az
}

type organization struct {
	name            string
	catalog         *[]catalog
	noMembers       int
	noMemberInvites int
	noDraftAPI      int
	avgAPISize      int
	maxAPISize      int
	noDraftProduct  int
	noTLSProfile    int
	noOAuthProvider int
	userRegistries  int
	noKeyStore      int
	noTrustStore    int
}
type az struct {
	name         string
	noPortal     int
	noAnalytics  int
	noV5cgateway int
	noApigateway int
}

type catalog struct {
	name            string
	noMember        int
	noMemberInvites int
	noAPI           int
	noProduct       int
	avgAPISize      int
	maxAPISize      int
	noSpace         int //*[]space
	noConsumerOrg   int //*[]consumerOrg
	portal          bool
	noTLSProfile    int
	noOAuthProvider int
	userRegistries  int
	applications    int
}

type space struct {
	name       string
	noMember   int
	noAPI      int
	noProduct  int
	avgAPISize int
	maxAPISize int
}

type consumerOrg struct {
	name           string
	noMember       int
	noApplications int
}
