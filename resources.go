package patreon

// Address represents a patron's shipping address.
type AddressV2 struct {
	AddressAttributes
	// The user this address belongs to.
	User *UserV2
	// The campaigns that have access to the address.
	Campaigns []*CampaignV2
}

// Benefit represents a benefit added to the campaign, which can be added to a tier to be delivered to the patron.
type Benefit struct {
	BenefitAttributes
	// The Tiers the benefit has been added to.
	Tiers []*Tier
	// The Deliverables that have been generated by the Benefit
	Deliverables []*Deliverables
	// The Campaign the benefit belongs to
	Campaign *CampaignV2
}

// Campaign represents the creator's page, and the top-level object for accessing lists of members, tiers, etc.
type CampaignV2 struct {
	CampaignAttributes
	// The campaign's tiers.
	Tiers []*Tier
	// The campaign owner.
	Creator *UserV2
	// The campaign's benefits.
	Benefits []*Benefit
	// The campaign's goals.
	Goals []*GoalV2
}

// Deliverables represents the record of whether or not a patron has been delivered the benefitthey are owed
// because of their member tier.
type Deliverables struct {
	DeliverableAttributes
	// The Campaign the Deliverables were generated for.
	Campaign *CampaignV2
	// The Benefit the Deliverables were generated for.
	Benefit *Benefit
	// The member who has been granted the deliverable.
	Member *Member
	// The user who has been granted the deliverable. This user is the same as the member user.
	User *UserV2
}

// Goal represents a funding goal in USD set by a creator on a campaign.
type GoalV2 struct {
	GoalAttributes
	// The campaign trying to reach the goal
	Campaign *CampaignV2
}

// Media represents a file uploaded to patreon.com, usually an image.
type Media struct {
	MediaAttributes
}

// Member represents the record of a user's membership to a campaign. Remains consistent across months of pledging.
type Member struct {
	MemberAttributes
	// The member's shipping address that they entered for the campaign.Requires the campaign.members.address scope.
	Address *AddressV2
	// The campaign that the membership is for.
	Campaign *CampaignV2
	// The tiers that the member is entitled to. This includes a current pledge,
	// or payment that covers the current payment period.
	CurrentlyEntitledTiers []*Tier
	// The user who is pledging to the campaign.
	User *UserV2
}

// OAuthClient represents a client created by a developer, used for getting OAuth2 access tokens.
type OAuthClient struct {
	OAuthClientAttributes
	// The user who created the OAuth Client.
	User *UserV2
	// The campaign of the user who created the OAuth Client.
	Campaign *CampaignV2
	// The token of the user who created the client.
	CreatorToken string
}

// Tier represents a membership level on a campaign, which can have benefits attached to it.
type Tier struct {
	TierAttributes
	// The campaign the tier belongs to.
	Campaign *CampaignV2
	// The image file associated with the tier.
	TierImage *Media
	// The benefits attached to the tier, which are used for generating deliverables
	Benefits []*Benefit
}

// User represents the Patreon user, which can be both patron and creator.
type UserV2 struct {
	*UserAttributes
	ID string
	// Usually a zero or one-element array with the user's membership to the token creator's campaign,
	// if they are a member. With the identity.memberships scope, this returns memberships to ALL campaigns the user is
	// a member of.
	Memberships []*Member
	Campaign    *CampaignV2
}

// Webhook represents an event happening on a particular campaign.
type Webhook struct {
	WebhookAttributes
	// The client which created the webhook
	Client *OAuthClient
	// The campaign whose events trigger the webhook.
	Campaign *CampaignV2
}
