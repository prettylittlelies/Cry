package discord

import (
	"encoding/json"
	"github.com/prettylittlelies/cry/snowflake"
	"time"
)

type AcceptActivityInviteArguments struct {
	Type          ActivityActionType `json:"type"`
	UserID        snowflake.ID       `json:"user_id"`
	SessionID     string             `json:"session_id"`
	ChannelID     snowflake.ID       `json:"channel_id"`
	MessageID     snowflake.ID       `json:"message_id"`
	ApplicationID snowflake.ID       `json:"application_id,omitzero"`
}

type AcceptApplicationTesterInvitationQuery struct {
	Token string `json:"token"`
}

type AcceptInviteParams struct {
	SessionID string `json:"session_id,omitzero"`
}

type AcceptQuestParams struct {
	Location              QuestContentType `json:"location"`
	MetadataSealed        *string          `json:"metadata_sealed,omitempty"`
	TrafficMetadataSealed *string          `json:"traffic_metadata_sealed,omitempty"`
}

type AcceptTeamInviteParams struct {
	Token string `json:"token"`
}

type AccessToken struct {
	TokenType    string  `json:"token_type"`
	AccessToken  string  `json:"access_token"`
	IDToken      string  `json:"id_token,omitzero"`
	Scope        string  `json:"scope"`
	ExpiresIn    int64   `json:"expires_in"`
	RefreshToken string  `json:"refresh_token,omitzero"`
	Guild        Guild   `json:"guild,omitzero"`
	Webhook      Webhook `json:"webhook,omitzero"`
}

type AccountStanding struct {
	State AccountStandingState `json:"state"`
}

type AcknowledgeApplicationDisclosuresParams struct {
	Disclosures []int64 `json:"disclosures"`
}

type AcknowledgeApplicationDisclosuresResponse struct {
	Disclosures []int64 `json:"disclosures"`
}

type AcknowledgeGuildFeatureParams struct {
	Token *string `json:"token,omitempty"`
}

type AcknowledgeGuildFeatureResponse struct {
	Token *string `json:"token"`
}

type AcknowledgeMessageParams struct {
	Token        *string        `json:"token,omitempty"`
	Manual       bool           `json:"manual,omitzero"`
	MentionCount int64          `json:"mention_count,omitzero"`
	Flags        ReadStateFlags `json:"flags,omitzero"`
	LastViewed   int64          `json:"last_viewed,omitzero"`
}

type AcknowledgeMessageResponse struct {
	Token *string `json:"token"`
}

type AcknowledgeSafetyWarningsParams struct {
	WarningIDs []string `json:"warning_ids"`
}

type AcknowledgeUserFeatureParams struct {
	Token *string `json:"token,omitempty"`
}

type AcknowledgeUserFeatureResponse struct {
	Token *string `json:"token"`
}

type AcknowledgeUserOfferParams struct {
	UserTrialOfferID    snowflake.ID `json:"user_trial_offer_id,omitzero"`
	UserDiscountOfferID snowflake.ID `json:"user_discount_offer_id,omitzero"`
}

type AcknowledgeUserOfferResponse struct {
	UserTrialOffer    *UserTrialOffer    `json:"user_trial_offer"`
	UserDiscountOffer *UserDiscountOffer `json:"user_discount_offer,omitempty"`
	UserDiscount      *UserDiscountOffer `json:"user_discount,omitempty"`
}

type Action struct {
	EventID     snowflake.ID          `json:"event_id"`
	UserID      snowflake.ID          `json:"user_id"`
	EntityID    snowflake.ID          `json:"entity_id"`
	DisplayType TeenActionDisplayType `json:"display_type"`
}

type ActionGuildJoinRequestParams struct {
	Action          GuildJoinRequestStatus `json:"action"`
	RejectionReason *string                `json:"rejection_reason,omitempty"`
}

type ActionRow struct {
	Type       ComponentType `json:"type"`
	ID         int64         `json:"id,omitzero"`
	Components Components    `json:"components"`
}

type Activity struct {
	ID                  string               `json:"id"`
	Name                string               `json:"name"`
	Type                ActivityType         `json:"type"`
	URL                 *string              `json:"url,omitempty"`
	CreatedAt           int64                `json:"created_at"`
	SessionID           *string              `json:"session_id,omitempty"`
	Platform            ActivityPlatformType `json:"platform,omitzero"`
	SupportedPlatforms  []string             `json:"supported_platforms,omitempty"`
	Timestamps          ActivityTimestamps   `json:"timestamps,omitzero"`
	ApplicationID       snowflake.ID         `json:"application_id,omitzero"`
	ParentApplicationID snowflake.ID         `json:"parent_application_id,omitzero"`
	StatusDisplayType   *StatusDisplayType   `json:"status_display_type,omitempty"`
	Details             *string              `json:"details,omitempty"`
	DetailsURL          *string              `json:"details_url,omitempty"`
	State               *string              `json:"state,omitempty"`
	StateURL            *string              `json:"state_url,omitempty"`
	SyncID              string               `json:"sync_id,omitzero"`
	Flags               ActivityFlags        `json:"flags,omitzero"`
	Buttons             []string             `json:"buttons,omitempty"`
	Emoji               *ActivityEmoji       `json:"emoji,omitempty"`
	Party               ActivityParty        `json:"party,omitzero"`
	Assets              ActivityAssets       `json:"assets,omitzero"`
	Secrets             ActivitySecrets      `json:"secrets,omitzero"`
	Metadata            ActivityMetadata     `json:"metadata,omitzero"`
}

type ActivityAssets struct {
	LargeImage       string `json:"large_image,omitzero"`
	LargeText        string `json:"large_text,omitzero"`
	LargeURL         string `json:"large_url,omitzero"`
	SmallImage       string `json:"small_image,omitzero"`
	SmallText        string `json:"small_text,omitzero"`
	SmallURL         string `json:"small_url,omitzero"`
	InviteCoverImage string `json:"invite_cover_image,omitzero"`
}

type ActivityEmoji struct {
	Name     string       `json:"name"`
	ID       snowflake.ID `json:"id,omitzero"`
	Animated bool         `json:"animated,omitzero"`
}

type ActivityInstanceParticipantsUpdate struct {
	Participants []RPCActivityParticipant `json:"participants"`
}

type ActivityInvite struct {
	User      RPCUser            `json:"user"`
	Activity  Activity           `json:"activity"`
	Type      ActivityActionType `json:"type"`
	ChannelID snowflake.ID       `json:"channel_id"`
	MessageID snowflake.ID       `json:"message_id"`
}

type ActivityInviteCreate struct {
	MessageID   snowflake.ID           `json:"message_id,omitzero"`
	ChannelID   snowflake.ID           `json:"channel_id,omitzero"`
	Author      PartialUser            `json:"author,omitzero"`
	Application IntegrationApplication `json:"application,omitzero"`
	Activity    MessageActivity        `json:"activity,omitzero"`
}

type ActivityJoin struct {
	Secret string     `json:"secret"`
	Intent JoinIntent `json:"intent,omitzero"`
}

type ActivityJoinRequest struct {
	User      RPCUser            `json:"user"`
	Activity  Activity           `json:"activity"`
	Type      ActivityActionType `json:"type"`
	ChannelID snowflake.ID       `json:"channel_id"`
	MessageID snowflake.ID       `json:"message_id"`
}

type ActivityLayoutModeUpdate struct {
	LayoutMode LayoutMode `json:"layout_mode"`
}

type ActivityLink struct {
	ApplicationID snowflake.ID `json:"application_id"`
	LinkID        string       `json:"link_id"`
	AssetPath     string       `json:"asset_path,omitzero"`
	AssetID       snowflake.ID `json:"asset_id,omitzero"`
	Title         string       `json:"title"`
	Description   string       `json:"description"`
	CustomID      *string      `json:"custom_id,omitempty"`
	PrimaryCta    *string      `json:"primary_cta,omitempty"`
}

type ActivityMetadata struct {
	ButtonURLs []string `json:"button_urls,omitempty"`
	ArtistIDs  []string `json:"artist_ids,omitempty"`
	AlbumID    string   `json:"album_id,omitzero"`
	ContextURI string   `json:"context_uri,omitzero"`
	Type       string   `json:"type,omitzero"`
}

type ActivityParty struct {
	ID   string  `json:"id,omitzero"`
	Size []int64 `json:"size,omitempty"`
}

type ActivitySecrets struct {
	Join     string `json:"join,omitzero"`
	Spectate string `json:"spectate,omitzero"`
	Match    string `json:"match,omitzero"`
}

type ActivitySpectate struct {
	Secret string `json:"secret"`
}

type ActivityStart struct {
	ActivityType       ActivityType `json:"activity_type"`
	ActivityName       string       `json:"activity_name"`
	ActivityInstanceID string       `json:"activity_instance_id,omitzero"`
	ApplicationID      snowflake.ID `json:"application_id,omitzero"`
	ApplicationName    string       `json:"application_name,omitzero"`
	ApplicationIcon    *string      `json:"application_icon,omitempty"`
	UserID             snowflake.ID `json:"user_id,omitzero"`
	UserUsername       string       `json:"user_username,omitzero"`
	UserAvatar         *string      `json:"user_avatar,omitempty"`
}

type ActivitySubscription struct {
	UserID        snowflake.ID `json:"user_id"`
	ApplicationID snowflake.ID `json:"application_id"`
	PartyID       snowflake.ID `json:"party_id"`
	MessageID     snowflake.ID `json:"message_id"`
	ChannelID     snowflake.ID `json:"channel_id"`
}

type ActivityTimestamps struct {
	Start string `json:"start,omitzero"`
	End   string `json:"end,omitzero"`
}

type AddApplicationTesterParams struct {
	Username      string  `json:"username"`
	Discriminator *string `json:"discriminator,omitempty"`
}

type AddChannelRecipientParams struct {
	AccessToken string `json:"access_token,omitzero"`
	Nick        string `json:"nick,omitzero"`
}

type AddGuildMemberParams struct {
	AccessToken string           `json:"access_token"`
	Nick        *string          `json:"nick,omitempty"`
	Roles       []snowflake.ID   `json:"roles,omitempty"`
	Mute        bool             `json:"mute,omitzero"`
	Deaf        bool             `json:"deaf,omitzero"`
	Flags       GuildMemberFlags `json:"flags,omitzero"`
}

type AddGuildPowerupParams struct {
	GameServerName   string `json:"game_server_name,omitzero"`
	GameServerRegion string `json:"game_server_region,omitzero"`
}

type AddGuildRoleMembersParams struct {
	MemberIDs []snowflake.ID `json:"member_ids"`
}

type AddLobbyMemberParams struct {
	Metadata map[string]string `json:"metadata,omitempty"`
	Flags    LobbyMemberFlags  `json:"flags,omitzero"`
}

type AddPhoneNumberParams struct {
	Phone             string            `json:"phone,omitzero"`
	PhoneToken        string            `json:"phone_token,omitzero"`
	Password          string            `json:"password,omitzero"`
	ChangePhoneReason ChangePhoneReason `json:"change_phone_reason,omitzero"`
}

type AddSafetyWarningParams struct {
	SafetyWarningType SafetyWarningType `json:"safety_warning_type"`
}

type AddTeamMemberParams struct {
	Username      string  `json:"username"`
	Discriminator *string `json:"discriminator,omitempty"`
	Role          string  `json:"role"`
}

type AddThreadMemberQuery struct {
	Location string `json:"location,omitzero"`
}

type AdsSettings struct {
	AlwaysDeliver bool `json:"always_deliver"`
}

type AgeVerificationSystemMessageEmbed struct {
	Ctas        string `json:"ctas,omitzero"`
	ContentType string `json:"content_type"`
}

type AllApplicationSettings struct {
	AppSettings map[string]ApplicationSettings `json:"app_settings"`
}

type AllGuildSettings struct {
	Guilds               map[string]GuildSettings `json:"guilds"`
	LeaderboardsDisabled bool                     `json:"leaderboards_disabled"`
}

type AllowedMentions struct {
	Parse       []string       `json:"parse,omitempty"`
	Roles       []snowflake.ID `json:"roles,omitempty"`
	Users       []snowflake.ID `json:"users,omitempty"`
	RepliedUser bool           `json:"replied_user,omitzero"`
}

type AndroidQuestCtaConfig struct {
	AndroidAppID string `json:"android_app_id"`
}

type ApexExperimentAssignment struct {
	HashedName       int64                `json:"hashed_name"`
	VariantID        int64                `json:"variant_id"`
	Flags            *ApexExperimentFlags `json:"flags"`
	Revision         int64                `json:"revision"`
	TrackedVariantID int64                `json:"tracked_variant_id,omitzero"`
}

type ApexExperimentAssignments struct {
	EvaluationID string                     `json:"evaluation_id"`
	Assignments  []ApexExperimentAssignment `json:"assignments"`
}

type ApexExperimentMetadata struct {
	ID       int64                   `json:"id"`
	Name     string                  `json:"name"`
	Title    string                  `json:"title"`
	Revision int64                   `json:"revision"`
	UnitType ApexExperimentUnitType  `json:"unit_type"`
	Variants []ApexExperimentVariant `json:"variants"`
}

type ApexExperimentVariant struct {
	ID    int64                     `json:"id"`
	Label string                    `json:"label"`
	Type  ApexExperimentVariantType `json:"type"`
}

type ApexExperiments struct {
	Assignments  map[int64]map[snowflake.ID]ApexExperimentAssignments `json:"assignments"`
	Installation string                                               `json:"installation,omitzero"`
}

type AppVersionSettings struct {
	IsUsingOutdatedMobileVersion bool `json:"is_using_outdated_mobile_version"`
}

type AppealStatus struct {
	Status SafetyHubAppealStatus `json:"status"`
}

type AppearanceSettings struct {
	Theme                         json.RawMessage     `json:"theme"`
	DeveloperMode                 bool                `json:"developer_mode"`
	ClientThemeSettings           ClientThemeSettings `json:"client_theme_settings,omitzero"`
	MobileRedesignDisabled        bool                `json:"mobile_redesign_disabled"`
	ChannelListLayout             string              `json:"channel_list_layout,omitzero"`
	MessagePreviews               string              `json:"message_previews,omitzero"`
	SearchResultExactCountEnabled bool                `json:"search_result_exact_count_enabled,omitzero"`
	TimestampHourCycle            json.RawMessage     `json:"timestamp_hour_cycle"`
	HappeningNowCardsDisabled     bool                `json:"happening_now_cards_disabled,omitzero"`
	LaunchPadMode                 json.RawMessage     `json:"launch_pad_mode"`
	UIDensity                     json.RawMessage     `json:"ui_density"`
	SwipeRightToLeftMode          json.RawMessage     `json:"swipe_right_to_left_mode"`
	DefaultGuildThemePreference   json.RawMessage     `json:"default_guild_theme_preference"`
	DarkSidebar                   bool                `json:"dark_sidebar"`
}

type Application struct {
	ID                                snowflake.ID                                       `json:"id"`
	Name                              string                                             `json:"name"`
	Description                       string                                             `json:"description"`
	Icon                              *string                                            `json:"icon"`
	CoverImage                        string                                             `json:"cover_image,omitzero"`
	Splash                            string                                             `json:"splash,omitzero"`
	Type                              *ApplicationType                                   `json:"type"`
	Flags                             ApplicationFlags                                   `json:"flags"`
	FlagsNew                          string                                             `json:"flags_new"`
	PrimarySKUID                      snowflake.ID                                       `json:"primary_sku_id,omitzero"`
	VerifyKey                         string                                             `json:"verify_key"`
	GuildID                           snowflake.ID                                       `json:"guild_id,omitzero"`
	EulaID                            snowflake.ID                                       `json:"eula_id,omitzero"`
	Slug                              string                                             `json:"slug,omitzero"`
	Aliases                           []string                                           `json:"aliases,omitempty"`
	Executables                       []ApplicationExecutable                            `json:"executables,omitempty"`
	ThirdPartySKUs                    []ApplicationSKU                                   `json:"third_party_skus,omitempty"`
	Hook                              bool                                               `json:"hook"`
	Overlay                           bool                                               `json:"overlay,omitzero"`
	OverlayMethods                    OverlayMethodFlags                                 `json:"overlay_methods,omitzero"`
	OverlayWarn                       bool                                               `json:"overlay_warn,omitzero"`
	OverlayCompatibilityHook          bool                                               `json:"overlay_compatibility_hook,omitzero"`
	Bot                               PartialUser                                        `json:"bot,omitzero"`
	Owner                             PartialUser                                        `json:"owner"`
	Team                              *Team                                              `json:"team,omitempty"`
	Developers                        []Company                                          `json:"developers,omitempty"`
	Publishers                        []Company                                          `json:"publishers,omitempty"`
	RPCOrigins                        []string                                           `json:"rpc_origins,omitempty"`
	RedirectUris                      []string                                           `json:"redirect_uris"`
	DeeplinkURI                       string                                             `json:"deeplink_uri,omitzero"`
	IntegrationPublic                 bool                                               `json:"integration_public"`
	IntegrationRequireCodeGrant       bool                                               `json:"integration_require_code_grant"`
	BotPublic                         bool                                               `json:"bot_public,omitzero"`
	BotRequireCodeGrant               bool                                               `json:"bot_require_code_grant,omitzero"`
	BotDisabled                       bool                                               `json:"bot_disabled,omitzero"`
	BotQuarantined                    bool                                               `json:"bot_quarantined,omitzero"`
	BotApproximateGuildCount          int64                                              `json:"bot_approximate_guild_count,omitzero"`
	ApproximateGuildCount             int64                                              `json:"approximate_guild_count,omitzero"`
	ApproximateUserInstallCount       int64                                              `json:"approximate_user_install_count"`
	ApproximateUserAuthorizationCount int64                                              `json:"approximate_user_authorization_count"`
	InternalGuildRestriction          InternalGuildRestriction                           `json:"internal_guild_restriction"`
	TermsOfServiceURL                 string                                             `json:"terms_of_service_url,omitzero"`
	PrivacyPolicyURL                  string                                             `json:"privacy_policy_url,omitzero"`
	RoleConnectionsVerificationURL    *string                                            `json:"role_connections_verification_url"`
	InteractionsEndpointURL           string                                             `json:"interactions_endpoint_url"`
	InteractionsVersion               ApplicationInteractionsVersion                     `json:"interactions_version"`
	InteractionsEventTypes            []string                                           `json:"interactions_event_types"`
	EventWebhooksStatus               EventWebhooksStatus                                `json:"event_webhooks_status,omitzero"`
	EventWebhooksURL                  string                                             `json:"event_webhooks_url,omitzero"`
	EventWebhooksTypes                []string                                           `json:"event_webhooks_types,omitempty"`
	ExplicitContentFilter             ExplicitContentFilterLevel                         `json:"explicit_content_filter"`
	Tags                              []string                                           `json:"tags,omitempty"`
	InstallParams                     ApplicationInstallParams                           `json:"install_params,omitzero"`
	CustomInstallURL                  string                                             `json:"custom_install_url,omitzero"`
	IntegrationTypesConfig            map[int64]*ApplicationIntegrationTypeConfiguration `json:"integration_types_config,omitempty"`
	ConnectionEntrypointURL           string                                             `json:"connection_entrypoint_url,omitzero"`
	IsVerified                        bool                                               `json:"is_verified"`
	VerificationState                 ApplicationVerificationState                       `json:"verification_state"`
	StoreApplicationState             StoreApplicationState                              `json:"store_application_state"`
	RPCApplicationState               RPCApplicationState                                `json:"rpc_application_state"`
	CreatorMonetizationState          int64                                              `json:"creator_monetization_state"`
	IsDiscoverable                    bool                                               `json:"is_discoverable"`
	DiscoverabilityState              ApplicationDiscoverabilityState                    `json:"discoverability_state"`
	DiscoveryEligibilityFlags         ApplicationDiscoveryEligibilityFlags               `json:"discovery_eligibility_flags"`
	IsMonetized                       bool                                               `json:"is_monetized"`
	StorefrontAvailable               bool                                               `json:"storefront_available"`
	MonetizationState                 ApplicationMonetizationState                       `json:"monetization_state"`
	MonetizationEligibilityFlags      ApplicationMonetizationEligibilityFlags            `json:"monetization_eligibility_flags,omitzero"`
	MaxParticipants                   int64                                              `json:"max_participants,omitzero"`
	EmbeddedActivityConfig            EmbeddedActivityConfig                             `json:"embedded_activity_config,omitzero"`
	ApprovedConsoles                  []int64                                            `json:"approved_consoles"`
	PricingLocalizationStrategy       PricingLocalizationStrategy                        `json:"pricing_localization_strategy"`
}

type ApplicationActivityStatistics struct {
	UserID        string    `json:"user_id"`
	LastPlayedAt  time.Time `json:"last_played_at"`
	TotalDuration int64     `json:"total_duration"`
}

type ApplicationAsset struct {
	AssetID    snowflake.ID               `json:"asset_id"`
	AssetType  ApplicationAssetType       `json:"asset_type"`
	Key        string                     `json:"key"`
	Metadata   ApplicationAssetMetadata   `json:"metadata"`
	UpdatedAt  time.Time                  `json:"updated_at"`
	Visibility ApplicationAssetVisibility `json:"visibility"`
}

type ApplicationAssetMetadata struct {
	ContentType string `json:"content_type"`
	Height      int64  `json:"height"`
	Width       int64  `json:"width"`
	IsAnimated  bool   `json:"is_animated"`
}

type ApplicationCommand struct {
	ID                       snowflake.ID                       `json:"id,omitzero"`
	Type                     int64                              `json:"type,omitzero"`
	ApplicationID            snowflake.ID                       `json:"application_id"`
	GuildID                  snowflake.ID                       `json:"guild_id,omitzero"`
	Name                     string                             `json:"name"`
	NameLocalizations        map[string]string                  `json:"name_localizations,omitempty"`
	NameLocalized            *string                            `json:"name_localized,omitempty"`
	NameDefault              string                             `json:"name_default,omitzero"`
	Description              string                             `json:"description"`
	DescriptionLocalizations map[string]string                  `json:"description_localizations,omitempty"`
	DescriptionLocalized     *string                            `json:"description_localized,omitempty"`
	DescriptionDefault       string                             `json:"description_default,omitzero"`
	Options                  []ApplicationCommandOption         `json:"options,omitempty"`
	DefaultMemberPermissions *string                            `json:"default_member_permissions"`
	DMPermission             bool                               `json:"dm_permission,omitzero"`
	Permissions              ApplicationCommandIndexPermissions `json:"permissions,omitzero"`
	NSFW                     bool                               `json:"nsfw,omitzero"`
	IntegrationTypes         []int64                            `json:"integration_types,omitempty"`
	GlobalPopularityRank     int64                              `json:"global_popularity_rank,omitzero"`
	Contexts                 []int64                            `json:"contexts,omitempty"`
	Version                  snowflake.ID                       `json:"version"`
	Handler                  ApplicationCommandHandlerType      `json:"handler,omitzero"`
}

type ApplicationCommandAutocompleteResponse struct {
	Choices []ApplicationCommandOptionChoice `json:"choices"`
	Nonce   *string                          `json:"nonce"`
}

type ApplicationCommandData struct {
	ID       snowflake.ID                   `json:"id"`
	Name     string                         `json:"name"`
	Type     ApplicationCommandType         `json:"type"`
	Resolved ResolvedData                   `json:"resolved,omitzero"`
	Options  []ApplicationCommandDataOption `json:"options,omitempty"`
	GuildID  snowflake.ID                   `json:"guild_id,omitzero"`
	TargetID snowflake.ID                   `json:"target_id,omitzero"`
}

type ApplicationCommandDataOption struct {
	Type    ApplicationCommandOptionType   `json:"type"`
	Name    string                         `json:"name"`
	Value   json.RawMessage                `json:"value,omitempty"`
	Options []ApplicationCommandDataOption `json:"options,omitempty"`
	Focused bool                           `json:"focused,omitzero"`
}

type ApplicationCommandFrecency struct {
	ApplicationCommands map[string]FrecencyItem `json:"application_commands"`
}

type ApplicationCommandIndex struct {
	Applications        []CommandIndexApplication `json:"applications"`
	ApplicationCommands []ApplicationCommand      `json:"application_commands"`
	Version             snowflake.ID              `json:"version"`
}

type ApplicationCommandIndexPermissions struct {
	User     bool                  `json:"user,omitzero"`
	Roles    map[snowflake.ID]bool `json:"roles,omitempty"`
	Channels map[snowflake.ID]bool `json:"channels,omitempty"`
}

type ApplicationCommandOption struct {
	Type                     ApplicationCommandOptionType     `json:"type"`
	Name                     string                           `json:"name"`
	NameLocalizations        map[string]string                `json:"name_localizations,omitempty"`
	NameLocalized            *string                          `json:"name_localized,omitempty"`
	Description              string                           `json:"description"`
	DescriptionLocalizations map[string]string                `json:"description_localizations,omitempty"`
	DescriptionLocalized     *string                          `json:"description_localized,omitempty"`
	Required                 bool                             `json:"required,omitzero"`
	Choices                  []ApplicationCommandOptionChoice `json:"choices,omitempty"`
	Options                  []ApplicationCommandOption       `json:"options,omitempty"`
	ChannelTypes             []int64                          `json:"channel_types,omitempty"`
	MinValue                 json.RawMessage                  `json:"min_value,omitempty"`
	MaxValue                 json.RawMessage                  `json:"max_value,omitempty"`
	MinLength                ApplicationCommandOptionType     `json:"min_length,omitzero"`
	MaxLength                ApplicationCommandOptionType     `json:"max_length,omitzero"`
	Autocomplete             bool                             `json:"autocomplete,omitzero"`
	FileTypes                []json.RawMessage                `json:"file_types,omitempty"`
}

type ApplicationCommandOptionChoice struct {
	Name              string            `json:"name"`
	NameLocalizations map[string]string `json:"name_localizations,omitempty"`
	NameLocalized     *string           `json:"name_localized,omitempty"`
	Value             json.RawMessage   `json:"value"`
}

type ApplicationCommandPermissions struct {
	ID         snowflake.ID                     `json:"id"`
	Type       ApplicationCommandPermissionType `json:"type"`
	Permission bool                             `json:"permission"`
}

type ApplicationCommandsGetUserApplicationCommandIndexParams struct {
	Name                     string                         `json:"name,omitzero"`
	NameLocalizations        map[string]string              `json:"name_localizations,omitempty"`
	Description              string                         `json:"description,omitzero"`
	DescriptionLocalizations map[string]string              `json:"description_localizations,omitempty"`
	Options                  []ApplicationCommandOption     `json:"options,omitempty"`
	DefaultMemberPermissions *string                        `json:"default_member_permissions"`
	DMPermission             *bool                          `json:"dm_permission,omitempty"`
	DefaultPermission        bool                           `json:"default_permission,omitzero"`
	NSFW                     *bool                          `json:"nsfw,omitempty"`
	IntegrationTypes         []int64                        `json:"integration_types,omitempty"`
	Contexts                 []int64                        `json:"contexts,omitempty"`
	Handler                  *ApplicationCommandHandlerType `json:"handler,omitempty"`
}

type ApplicationCommandsGetUserApplicationCommandIndexQuery struct {
	WithLocalizations bool `json:"with_localizations,omitzero"`
}

type ApplicationDMSettings struct {
	AllowMobilePush bool `json:"allow_mobile_push,omitzero"`
}

type ApplicationDirectoryCarouselItem struct {
	Type     ApplicationDirectoryCarouselItemType `json:"type"`
	URL      string                               `json:"url"`
	ProxyURL string                               `json:"proxy_url"`
}

type ApplicationDirectoryCategory struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type ApplicationDirectoryCollection struct {
	ID                                  string                               `json:"id"`
	Type                                ApplicationDirectoryItemType         `json:"type"`
	Position                            int64                                `json:"position"`
	Platforms                           ApplicationDirectoryPlatformFlags    `json:"platforms"`
	ActiveState                         ApplicationDirectoryActiveStateType  `json:"active_state"`
	Flags                               ApplicationFlags                     `json:"flags"`
	Title                               string                               `json:"title"`
	Description                         string                               `json:"description"`
	ApplicationDirectoryCollectionItems []ApplicationDirectoryCollectionItem `json:"application_directory_collection_items"`
}

type ApplicationDirectoryCollectionItem struct {
	ID          snowflake.ID                 `json:"id"`
	Type        ApplicationDirectoryItemType `json:"type"`
	Position    int64                        `json:"position"`
	Flags       ApplicationFlags             `json:"flags"`
	ImageHash   *string                      `json:"image_hash"`
	Application Application                  `json:"application"`
}

type ApplicationDirectoryEntry struct {
	GuildCount                       int64                `json:"guild_count"`
	CarouselItems                    []json.RawMessage    `json:"carousel_items,omitempty"`
	SupportedLocales                 []string             `json:"supported_locales"`
	ExternalURLs                     []json.RawMessage    `json:"external_urls,omitempty"`
	PopularApplicationCommandIDs     []snowflake.ID       `json:"popular_application_command_ids,omitempty"`
	PopularApplicationCommands       []ApplicationCommand `json:"popular_application_commands,omitempty"`
	DetailedDescription              string               `json:"detailed_description,omitzero"`
	ShortDescription                 string               `json:"short_description"`
	ShortDescriptionLocalizations    map[string]string    `json:"short_description_localizations,omitempty"`
	DetailedDescriptionLocalizations map[string]string    `json:"detailed_description_localizations,omitempty"`
}

type ApplicationDirectoryExternalURL struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type ApplicationDirectorySearchResults struct {
	Type ApplicationDirectoryItemType `json:"type"`
	Data Application                  `json:"data"`
}

type ApplicationDistributor struct {
	Distributor DistributorType `json:"distributor"`
	SKU         string          `json:"sku,omitzero"`
}

type ApplicationExecutable struct {
	OS         PresenceOperatingSystemType `json:"os"`
	Name       string                      `json:"name"`
	IsLauncher bool                        `json:"is_launcher"`
}

type ApplicationExternalIdentityProviderClient struct {
	ID            string  `json:"id"`
	OidcIssuerURL *string `json:"oidc_issuer_url"`
	Description   *string `json:"description,omitempty"`
	Environment   *string `json:"environment,omitempty"`
}

type ApplicationExternalIdentityProviderConfiguration struct {
	ApplicationID snowflake.ID                                `json:"application_id"`
	ProviderType  int64                                       `json:"provider_type"`
	ClientID      string                                      `json:"client_id"`
	OidcIssuerURL *string                                     `json:"oidc_issuer_url"`
	Clients       []ApplicationExternalIdentityProviderClient `json:"clients"`
}

type ApplicationFrecency struct {
	Applications map[string]FrecencyItem `json:"applications"`
}

type ApplicationInstallParams struct {
	Scopes      []string    `json:"scopes"`
	Permissions Permissions `json:"permissions"`
}

type ApplicationIntegrationTypeConfiguration struct {
	OAuth2InstallParams ApplicationInstallParams `json:"oauth2_install_params,omitzero"`
}

type ApplicationLaunchEmbeddedActivityParams struct {
	SessionID string `json:"session_id"`
}

type ApplicationLibraryInstallComplete struct {
	ApplicationID   snowflake.ID `json:"application_id"`
	ApplicationName string       `json:"application_name"`
	ApplicationIcon *string      `json:"application_icon,omitempty"`
}

type ApplicationListEmbeddedActivityInstancesResponse struct {
	Instances []ChannelActivityInstance `json:"instances"`
}

type ApplicationOAuth2Asset struct {
	ID   string `json:"id"`
	Type int64  `json:"type"`
	Name string `json:"name"`
}

type ApplicationProxyConfig struct {
	URLMap []ApplicationProxyMapping `json:"url_map"`
}

type ApplicationProxyMapping struct {
	Prefix string `json:"prefix"`
	Target string `json:"target"`
}

type ApplicationRoleConnection struct {
	PlatformName        *string                             `json:"platform_name"`
	PlatformUsername    *string                             `json:"platform_username"`
	Metadata            json.RawMessage                     `json:"metadata"`
	Application         IntegrationApplication              `json:"application,omitzero"`
	ApplicationMetadata []ApplicationRoleConnectionMetadata `json:"application_metadata,omitempty"`
}

type ApplicationRoleConnectionMetadata struct {
	Type                     RoleConnectionOperatorType `json:"type"`
	Key                      string                     `json:"key"`
	Name                     string                     `json:"name"`
	NameLocalizations        map[string]json.RawMessage `json:"name_localizations,omitempty"`
	Description              string                     `json:"description"`
	DescriptionLocalizations map[string]json.RawMessage `json:"description_localizations,omitempty"`
}

type ApplicationSKU struct {
	ID          *string         `json:"id"`
	SKU         *string         `json:"sku"`
	Distributor DistributorType `json:"distributor"`
}

type ApplicationSettings struct {
	AppDMSettings      UserSettingsProtoApplicationDMSettings `json:"app_dm_settings,omitzero"`
	AppSharingSettings ApplicationSharingSettings             `json:"app_sharing_settings,omitzero"`
}

type ApplicationSharingSettings struct {
	DisableApplicationActivitySharing bool `json:"disable_application_activity_sharing"`
}

type ApplyConfettiConsumableParams struct {
	ChannelID snowflake.ID `json:"channel_id"`
	MessageID snowflake.ID `json:"message_id"`
	EmojiName string       `json:"emoji_name"`
}

type ApplyHdStreamingConsumableParams struct {
	ChannelID snowflake.ID `json:"channel_id"`
}

type AssetConfig struct {
	Animated *string `json:"animated"`
	Static   string  `json:"static"`
}

type Attachment struct {
	ID                 snowflake.ID        `json:"id"`
	Filename           string              `json:"filename"`
	Title              string              `json:"title,omitzero"`
	UploadedFilename   string              `json:"uploaded_filename,omitzero"`
	Description        string              `json:"description,omitzero"`
	ContentType        string              `json:"content_type,omitzero"`
	Size               int64               `json:"size"`
	URL                string              `json:"url"`
	ProxyURL           string              `json:"proxy_url"`
	Height             *int64              `json:"height,omitempty"`
	Width              *int64              `json:"width,omitempty"`
	ContentScanVersion int64               `json:"content_scan_version,omitzero"`
	PlaceholderVersion int64               `json:"placeholder_version,omitzero"`
	Placeholder        string              `json:"placeholder,omitzero"`
	Ephemeral          bool                `json:"ephemeral,omitzero"`
	DurationSecs       float64             `json:"duration_secs,omitzero"`
	Waveform           string              `json:"waveform,omitzero"`
	Flags              AttachmentFlags     `json:"flags,omitzero"`
	IsClip             bool                `json:"is_clip,omitzero"`
	IsThumbnail        bool                `json:"is_thumbnail,omitzero"`
	IsRemix            bool                `json:"is_remix,omitzero"`
	IsSpoiler          bool                `json:"is_spoiler,omitzero"`
	ClipCreatedAt      time.Time           `json:"clip_created_at,omitzero"`
	ClipParticipantIDs []snowflake.ID      `json:"clip_participant_ids,omitempty"`
	ClipParticipants   []PartialUser       `json:"clip_participants,omitempty"`
	ApplicationID      snowflake.ID        `json:"application_id,omitzero"`
	Application        *PartialApplication `json:"application,omitempty"`
}

type AudienceNewMembersByDiscordTenure struct {
	DayPt      string `json:"day_pt"`
	Tenure     string `json:"tenure"`
	NewMembers int64  `json:"new_members"`
}

type AudienceParticipatorsByGuildTenure struct {
	DayPt         string `json:"day_pt"`
	Tenure        string `json:"tenure"`
	Participators int64  `json:"participators"`
}

type AudienceParticipatorsByPlatform struct {
	DayPt         string `json:"day_pt"`
	Platform      string `json:"platform"`
	Participators int64  `json:"participators"`
}

type AudienceParticipatorsByRegistrationCountry struct {
	DayPt         string `json:"day_pt"`
	CountryCode   string `json:"country_code"`
	Participators int64  `json:"participators"`
}

type AudioContext struct {
	Muted           bool    `json:"muted"`
	Volume          float64 `json:"volume"`
	SoundboardMuted bool    `json:"soundboard_muted"`
}

type AudioSettings struct {
	User   map[string]UserSettingsProtoAudioContext `json:"user"`
	Stream map[string]UserSettingsProtoAudioContext `json:"stream"`
}

type AudioSettingsUpdate struct {
	User   map[snowflake.ID]json.RawMessage `json:"user"`
	Stream map[snowflake.ID]json.RawMessage `json:"stream"`
}

type AuditLogChange struct {
	NewValue json.RawMessage `json:"new_value,omitempty"`
	OldValue json.RawMessage `json:"old_value,omitempty"`
	Key      string          `json:"key"`
}

type AuditLogEntry struct {
	TargetID   *snowflake.ID      `json:"target_id"`
	Changes    []AuditLogChange   `json:"changes,omitempty"`
	UserID     *snowflake.ID      `json:"user_id"`
	ID         snowflake.ID       `json:"id"`
	ActionType AuditLogActionType `json:"action_type"`
	Options    json.RawMessage    `json:"options,omitempty"`
	Reason     string             `json:"reason,omitzero"`
}

type AuthSession struct {
	IDHash             string                `json:"id_hash"`
	ApproxLastUsedTime time.Time             `json:"approx_last_used_time"`
	ClientInfo         AuthSessionClientInfo `json:"client_info"`
}

type AuthSessionChange struct {
	AuthSessionIDHash string `json:"auth_session_id_hash"`
}

type AuthSessionClientInfo struct {
	OS       *string `json:"os"`
	Platform *string `json:"platform"`
	Location *string `json:"location"`
}

type AuthenticateArguments struct {
	AccessToken string `json:"access_token"`
}

type AuthenticateResponse struct {
	AccessToken string `json:"access_token"`
}

type AuthenticationGetUniqueUsernameEligibilityParams struct {
	Username string `json:"username"`
}

type AuthenticationGetUniqueUsernameEligibilityResponse struct {
	Taken *bool `json:"taken"`
}

type AuthenticationGetUniqueUsernameSuggestionsResponse struct {
	Username string `json:"username"`
}

type Authenticator struct {
	ID       string     `json:"id"`
	Type     string     `json:"type"`
	Name     string     `json:"name"`
	LastUsed *time.Time `json:"last_used"`
	CredID   string     `json:"cred_id,omitzero"`
}

type AuthenticatorDelete struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type AuthorizationURL struct {
	ClientID            snowflake.ID               `json:"client_id"`
	ResponseType        ResponseType               `json:"response_type,omitzero"`
	Scope               string                     `json:"scope,omitzero"`
	RedirectURI         string                     `json:"redirect_uri,omitzero"`
	Prompt              PromptBehavior             `json:"prompt,omitzero"`
	State               string                     `json:"state,omitzero"`
	Nonce               string                     `json:"nonce,omitzero"`
	CodeChallenge       string                     `json:"code_challenge,omitzero"`
	CodeChallengeMethod string                     `json:"code_challenge_method,omitzero"`
	IntegrationType     ApplicationIntegrationType `json:"integration_type,omitzero"`
	Permissions         int64                      `json:"permissions,omitzero"`
	GuildID             snowflake.ID               `json:"guild_id,omitzero"`
	ChannelID           snowflake.ID               `json:"channel_id,omitzero"`
	DisableGuildSelect  bool                       `json:"disable_guild_select,omitzero"`
}

type AuthorizeArguments struct {
	ClientID            snowflake.ID               `json:"client_id"`
	ResponseType        ResponseType               `json:"response_type,omitzero"`
	RedirectURI         string                     `json:"redirect_uri,omitzero"`
	Scopes              []string                   `json:"scopes,omitempty"`
	CodeChallenge       string                     `json:"code_challenge,omitzero"`
	CodeChallengeMethod string                     `json:"code_challenge_method,omitzero"`
	State               string                     `json:"state,omitzero"`
	Nonce               string                     `json:"nonce,omitzero"`
	Permissions         Permissions                `json:"permissions,omitzero"`
	GuildID             snowflake.ID               `json:"guild_id,omitzero"`
	ChannelID           snowflake.ID               `json:"channel_id,omitzero"`
	Prompt              PromptBehavior             `json:"prompt,omitzero"`
	DisableGuildSelect  bool                       `json:"disable_guild_select,omitzero"`
	IntegrationType     ApplicationIntegrationType `json:"integration_type,omitzero"`
	Pid                 int64                      `json:"pid,omitzero"`
}

type AuthorizeIPAddressParams struct {
	Token string `json:"token"`
}

type AuthorizeResponse struct {
	Code string `json:"code"`
}

type AuthorizeUserConnectionQuery struct {
	TwoWayLinkType *TwoWayLinkType `json:"two_way_link_type,omitempty"`
	TwoWayUserCode *string         `json:"two_way_user_code,omitempty"`
	Continuation   bool            `json:"continuation,omitzero"`
}

type AuthorizeUserConnectionResponse struct {
	URL string `json:"url"`
}

type AutoModerationActionExecution struct {
	GuildID              snowflake.ID  `json:"guild_id"`
	Action               AutomodAction `json:"action"`
	RuleID               snowflake.ID  `json:"rule_id"`
	RuleTriggerType      int64         `json:"rule_trigger_type"`
	UserID               snowflake.ID  `json:"user_id"`
	ChannelID            snowflake.ID  `json:"channel_id,omitzero"`
	MessageID            snowflake.ID  `json:"message_id,omitzero"`
	AlertSystemMessageID snowflake.ID  `json:"alert_system_message_id,omitzero"`
	Content              string        `json:"content"`
	MatchedKeyword       *string       `json:"matched_keyword"`
	MatchedContent       *string       `json:"matched_content"`
}

type AutoModerationMentionRaidDetection struct {
	GuildID                        snowflake.ID `json:"guild_id"`
	DecisionID                     string       `json:"decision_id"`
	SuspiciousMentionActivityUntil time.Time    `json:"suspicious_mention_activity_until"`
}

type AutoModerationValidateGuildAutomodRuleParams struct {
	Name            string           `json:"name,omitzero"`
	EventType       AutomodEventType `json:"event_type,omitzero"`
	TriggerMetadata json.RawMessage  `json:"trigger_metadata,omitempty"`
	Actions         []AutomodAction  `json:"actions,omitempty"`
	Enabled         bool             `json:"enabled,omitzero"`
	ExemptRoles     []snowflake.ID   `json:"exempt_roles,omitempty"`
	ExemptChannels  []snowflake.ID   `json:"exempt_channels,omitempty"`
}

type AutomodAction struct {
	Type     AutomodActionType `json:"type"`
	Metadata json.RawMessage   `json:"metadata,omitempty"`
}

type AutomodAlertAction struct {
	Actor snowflake.ID `json:"actor"`
	Ts    time.Time    `json:"ts"`
}

type AutomodAlertActionsExecution struct {
	V       int64                      `json:"v"`
	Actions map[string]json.RawMessage `json:"actions"`
}

type AutomodAlertEmbed struct {
	RuleName                  string                         `json:"rule_name"`
	DecisionID                string                         `json:"decision_id"`
	DecisionReason            string                         `json:"decision_reason,omitzero"`
	DecisionOutcome           AutomodDecisionOutcome         `json:"decision_outcome"`
	ChannelID                 snowflake.ID                   `json:"channel_id,omitzero"`
	FlaggedMessageID          snowflake.ID                   `json:"flagged_message_id,omitzero"`
	Keyword                   string                         `json:"keyword"`
	KeywordMatchedContent     string                         `json:"keyword_matched_content"`
	BlockProfileUpdateType    AutomodProfileUpdateType       `json:"block_profile_update_type,omitzero"`
	QuarantineUser            AutomodQuarantineUserReason    `json:"quarantine_user,omitzero"`
	QuarantineUserAction      AutomodQuarantineUserAction    `json:"quarantine_user_action,omitzero"`
	QuarantineEvent           AutomodQuarantineEventType     `json:"quarantine_event,omitzero"`
	VoiceChannelStatusOutcome AutomodDecisionOutcome         `json:"voice_channel_status_outcome,omitzero"`
	ApplicationName           string                         `json:"application_name,omitzero"`
	InteractionUserID         snowflake.ID                   `json:"interaction_user_id,omitzero"`
	InteractionCallbackType   AutomodInteractionCallbackType `json:"interaction_callback_type,omitzero"`
	TimeoutDuration           int64                          `json:"timeout_duration,omitzero"`
	AlertActionsExecution     json.RawMessage                `json:"alert_actions_execution,omitempty"`
}

type AutomodIncidentNotificationEmbed struct {
	NotificationType               AutomodIncidentNotificationType `json:"notification_type,omitzero"`
	DecisionID                     string                          `json:"decision_id,omitzero"`
	ActionByUserID                 snowflake.ID                    `json:"action_by_user_id,omitzero"`
	RaidType                       AutomodRaidType                 `json:"raid_type,omitzero"`
	RaidDatetime                   time.Time                       `json:"raid_datetime,omitzero"`
	JoinAttempts                   int64                           `json:"join_attempts,omitzero"`
	DMsSent                        int64                           `json:"dms_sent,omitzero"`
	SuspiciousMentionActivityUntil time.Time                       `json:"suspicious_mention_activity_until,omitzero"`
	ResolvedReason                 AutomodRaidResolutionReason     `json:"resolved_reason,omitzero"`
}

type AutomodIncidentsData struct {
	RaidDetectedAt       *time.Time `json:"raid_detected_at"`
	DMSpamDetectedAt     *time.Time `json:"dm_spam_detected_at"`
	InvitesDisabledUntil *time.Time `json:"invites_disabled_until"`
	DMsDisabledUntil     *time.Time `json:"dms_disabled_until"`
}

type AutomodRule struct {
	ID              snowflake.ID       `json:"id"`
	GuildID         snowflake.ID       `json:"guild_id"`
	Name            string             `json:"name"`
	CreatorID       snowflake.ID       `json:"creator_id"`
	EventType       AutomodEventType   `json:"event_type"`
	TriggerType     AutomodTriggerType `json:"trigger_type"`
	TriggerMetadata json.RawMessage    `json:"trigger_metadata"`
	Actions         []Action           `json:"actions"`
	Enabled         bool               `json:"enabled"`
	ExemptRoles     []snowflake.ID     `json:"exempt_roles"`
	ExemptChannels  []snowflake.ID     `json:"exempt_channels"`
}

type AvailableDevice struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Avatar struct {
	ID          string  `json:"id"`
	StorageHash string  `json:"storage_hash"`
	Description *string `json:"description"`
}

type AvatarDecorationCollectibleItem struct {
	Type   CollectibleProductType `json:"type"`
	SKUID  snowflake.ID           `json:"sku_id"`
	Asset  string                 `json:"asset"`
	Assets ItemAssets             `json:"assets"`
	Label  string                 `json:"label"`
}

type AvatarDecorationData struct {
	Asset     string       `json:"asset"`
	SKUID     snowflake.ID `json:"sku_id"`
	ExpiresAt *int64       `json:"expires_at"`
}

type BackupCode struct {
	UserID   snowflake.ID `json:"user_id"`
	Code     string       `json:"code"`
	Consumed bool         `json:"consumed"`
}

type BadgeCollectiblesMarketing struct {
	Type                CollectiblesMarketingType `json:"type"`
	Version             int64                     `json:"version"`
	DismissibleContent  int64                     `json:"dismissible_content"`
	RefTargetBackground TargetBackground          `json:"ref_target_background,omitzero"`
	BadgeIcon           string                    `json:"badge_icon,omitzero"`
	BadgeText           string                    `json:"badge_text,omitzero"`
	ShowHoverGradient   bool                      `json:"show_hover_gradient,omitzero"`
}

type Ban struct {
	User   PartialUser `json:"user"`
	Reason *string     `json:"reason"`
}

type BannerAsset struct {
	BackgroundStyle  *string `json:"background_style"`
	DesktopMaxHeight *int64  `json:"desktop_max_height,omitempty"`
	MobileMaxHeight  *int64  `json:"mobile_max_height,omitempty"`
	Responsive       *bool   `json:"responsive"`
}

type BannerCollectiblesMarketing struct {
	Type            CollectiblesMarketingType `json:"type"`
	Version         int64                     `json:"version"`
	Title           string                    `json:"title"`
	Body            string                    `json:"body"`
	Asset           string                    `json:"asset"`
	PopoutAsset     string                    `json:"popout_asset"`
	RevertTextColor bool                      `json:"revert_text_color,omitzero"`
}

type BatchRejectMessageRequestsParams struct {
	ChannelIDs []snowflake.ID `json:"channel_ids"`
}

type BillingAddress struct {
	Line1      string  `json:"line_1"`
	Line2      *string `json:"line_2,omitempty"`
	Name       string  `json:"name"`
	PostalCode string  `json:"postal_code,omitzero"`
	City       string  `json:"city"`
	State      string  `json:"state,omitzero"`
	Country    string  `json:"country"`
}

type BillingPopupBridgeCallback struct {
	PaymentSourceType PaymentSourceType `json:"payment_source_type,omitzero"`
	State             string            `json:"state,omitzero"`
	Path              string            `json:"path,omitzero"`
	Query             map[string]string `json:"query,omitempty"`
}

type BillingPopupBridgeCallbackArguments struct {
	State             string            `json:"state"`
	Path              string            `json:"path"`
	Query             map[string]string `json:"query,omitempty"`
	PaymentSourceType PaymentSourceType `json:"payment_source_type"`
}

type BillingPopupBridgeCallbackResponse struct {
	Ok      bool              `json:"ok"`
	Headers map[string]string `json:"headers"`
	Body    json.RawMessage   `json:"body"`
	Text    string            `json:"text"`
	Status  int64             `json:"status"`
}

type BroadcastSettings struct {
	AllowFriends    bool              `json:"allow_friends,omitzero"`
	AllowedGuildIDs []json.RawMessage `json:"allowed_guild_ids"`
	AllowedUserIDs  []json.RawMessage `json:"allowed_user_ids"`
	AutoBroadcast   bool              `json:"auto_broadcast,omitzero"`
}

type BrowserHandoffArguments struct {
	Handofftoken string `json:"handoffToken"`
	Fingerprint  string `json:"fingerprint"`
}

type BulkActionGuildJoinRequestsParams struct {
	Action GuildJoinRequestStatus `json:"action"`
}

type BulkAddRelationshipsParams struct {
	UserIDs []snowflake.ID `json:"user_ids"`
	Token   string         `json:"token"`
}

type BulkAddRelationshipsResponse struct {
	FailedRequests     []snowflake.ID `json:"failed_requests"`
	SuccessfulRequests []snowflake.ID `json:"successful_requests"`
}

type BulkDeleteMessagesParams struct {
	Messages []snowflake.ID `json:"messages"`
}

type BulkGuildBanParams struct {
	UserIDs              []snowflake.ID `json:"user_ids"`
	DeleteMessageSeconds int64          `json:"delete_message_seconds,omitzero"`
}

type BulkGuildBanResponse struct {
	BannedUsers []snowflake.ID `json:"banned_users"`
	FailedUsers []snowflake.ID `json:"failed_users"`
}

type BulkGuildBanV2Params struct {
	UserIDs              []snowflake.ID `json:"user_ids"`
	DeleteMessageSeconds int64          `json:"delete_message_seconds,omitzero"`
}

type BulkLeaveGroupDMsParams struct {
	ChannelIDs []snowflake.ID `json:"channel_ids"`
}

type BulkModifyUserGuildSettingsParams struct {
	Guilds map[snowflake.ID]PartialUserGuildSettings `json:"guilds"`
}

type BulkRemoveRelationshipsParams struct {
	Filters []int64 `json:"filters,omitempty"`
}

type BulkRemoveRelationshipsQuery struct {
	RelationshipType RelationshipType `json:"relationship_type,omitzero"`
	OnlySpam         bool             `json:"only_spam,omitzero"`
}

type BulkScanExplicitMediaParams struct {
	Messages []BulkScanMessage `json:"messages"`
}

type BulkScanMessage struct {
	ChannelID snowflake.ID `json:"channel_id"`
	MessageID snowflake.ID `json:"message_id"`
}

type BulkUpdateReadStatesParams struct {
	ReadStates []ReadStateUpdate `json:"read_states"`
}

type Button struct {
	Type     ComponentType `json:"type"`
	ID       int64         `json:"id,omitzero"`
	Style    ButtonStyle   `json:"style"`
	Label    string        `json:"label,omitzero"`
	Emoji    Emoji         `json:"emoji,omitzero"`
	CustomID string        `json:"custom_id"`
	SKUID    snowflake.ID  `json:"sku_id,omitzero"`
	URL      string        `json:"url,omitzero"`
	Disabled bool          `json:"disabled,omitzero"`
}

type CallConnect struct {
	ChannelID          snowflake.ID `json:"channel_id"`
	GuildID            snowflake.ID `json:"guild_id,omitzero"`
	UserID             snowflake.ID `json:"user_id,omitzero"`
	IsFullscreenCallUI bool         `json:"is_fullscreen_call_ui,omitzero"`
}

type CallCreate struct {
	ChannelID   snowflake.ID   `json:"channel_id"`
	MessageID   snowflake.ID   `json:"message_id"`
	Region      string         `json:"region"`
	Ringing     []snowflake.ID `json:"ringing"`
	VoiceStates []VoiceState   `json:"voice_states"`
}

type CallDelete struct {
	ChannelID   snowflake.ID `json:"channel_id"`
	Unavailable bool         `json:"unavailable,omitzero"`
}

type CallRing struct {
	ChannelID    snowflake.ID `json:"channel_id"`
	ChannelType  ChannelType  `json:"channel_type"`
	ChannelName  string       `json:"channel_name"`
	ChannelIcon  *string      `json:"channel_icon,omitempty"`
	RTCRegion    *string      `json:"rtc_region,omitempty"`
	UserID       snowflake.ID `json:"user_id"`
	UserUsername string       `json:"user_username"`
	UserAvatar   *string      `json:"user_avatar,omitempty"`
}

type CallUpdate struct {
	ChannelID snowflake.ID    `json:"channel_id"`
	MessageID snowflake.ID    `json:"message_id"`
	Region    string          `json:"region"`
	Ringing   json.RawMessage `json:"ringing"`
}

type CancelRemoteAuthParams struct {
	HandshakeToken string `json:"handshake_token"`
}

type CaptchaResponse struct {
	CaptchaKey           []string       `json:"captcha_key"`
	CaptchaService       CaptchaService `json:"captcha_service"`
	CaptchaSitekey       *string        `json:"captcha_sitekey"`
	CaptchaSessionID     string         `json:"captcha_session_id,omitzero"`
	CaptchaRqdata        string         `json:"captcha_rqdata,omitzero"`
	CaptchaRqtoken       string         `json:"captcha_rqtoken,omitzero"`
	ShouldServeInvisible bool           `json:"should_serve_invisible,omitzero"`
}

type CaptureLogArguments struct {
	Level   LogLevel `json:"level"`
	Message string   `json:"message"`
}

type CertifiedDevice struct {
	Type                 CertifiedDeviceType   `json:"type"`
	ID                   string                `json:"id"`
	Vendor               CertifiedDeviceVendor `json:"vendor"`
	Model                CertifiedDeviceModel  `json:"model"`
	Related              []string              `json:"related,omitempty"`
	EchoCancellation     bool                  `json:"echo_cancellation,omitzero"`
	NoiseSuppression     bool                  `json:"noise_suppression,omitzero"`
	AutomaticGainControl bool                  `json:"automatic_gain_control,omitzero"`
	HardwareMute         bool                  `json:"hardware_mute,omitzero"`
}

type CertifiedDeviceModel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CertifiedDeviceVendor struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Channel struct {
	ID                                 snowflake.ID          `json:"id"`
	Type                               ChannelType           `json:"type"`
	GuildID                            snowflake.ID          `json:"guild_id,omitzero"`
	Position                           int64                 `json:"position,omitzero"`
	PermissionOverwrites               []PermissionOverwrite `json:"permission_overwrites,omitempty"`
	Name                               *string               `json:"name,omitempty"`
	Topic                              *string               `json:"topic,omitempty"`
	NSFW                               bool                  `json:"nsfw,omitzero"`
	LastMessageID                      *snowflake.ID         `json:"last_message_id,omitempty"`
	LastPinTimestamp                   *time.Time            `json:"last_pin_timestamp,omitempty"`
	Bitrate                            int64                 `json:"bitrate,omitzero"`
	UserLimit                          int64                 `json:"user_limit,omitzero"`
	RateLimitPerUser                   int64                 `json:"rate_limit_per_user,omitzero"`
	Recipients                         []PartialUser         `json:"recipients,omitempty"`
	RecipientFlags                     RecipientFlags        `json:"recipient_flags,omitzero"`
	Icon                               *string               `json:"icon,omitempty"`
	Nicks                              []ChannelNick         `json:"nicks,omitempty"`
	Managed                            bool                  `json:"managed,omitzero"`
	BlockedUserWarningDismissed        bool                  `json:"blocked_user_warning_dismissed,omitzero"`
	SafetyWarnings                     []SafetyWarning       `json:"safety_warnings,omitempty"`
	ApplicationID                      snowflake.ID          `json:"application_id,omitzero"`
	OwnerID                            snowflake.ID          `json:"owner_id,omitzero"`
	Owner                              *GuildMember          `json:"owner,omitempty"`
	ParentID                           *snowflake.ID         `json:"parent_id,omitempty"`
	RTCRegion                          *string               `json:"rtc_region,omitempty"`
	VideoQualityMode                   VideoQualityMode      `json:"video_quality_mode,omitzero"`
	TotalMessageSent                   int64                 `json:"total_message_sent,omitzero"`
	MessageCount                       int64                 `json:"message_count,omitzero"`
	MemberCount                        int64                 `json:"member_count,omitzero"`
	MemberIDsPreview                   []snowflake.ID        `json:"member_ids_preview,omitempty"`
	ThreadMetadata                     ThreadMetadata        `json:"thread_metadata,omitzero"`
	Member                             ThreadMember          `json:"member,omitzero"`
	DefaultAutoArchiveDuration         *int64                `json:"default_auto_archive_duration,omitempty"`
	DefaultThreadRateLimitPerUser      int64                 `json:"default_thread_rate_limit_per_user,omitzero"`
	Permissions                        Permissions           `json:"permissions,omitzero"`
	Flags                              ChannelFlags          `json:"flags,omitzero"`
	AvailableTags                      []json.RawMessage     `json:"available_tags,omitempty"`
	AppliedTags                        []snowflake.ID        `json:"applied_tags,omitempty"`
	DefaultReactionEmoji               *DefaultReaction      `json:"default_reaction_emoji,omitempty"`
	DefaultForumLayout                 ForumLayoutType       `json:"default_forum_layout,omitzero"`
	DefaultSortOrder                   *SortOrderType        `json:"default_sort_order,omitempty"`
	DefaultTagSetting                  SearchTagSetting      `json:"default_tag_setting,omitzero"`
	IconEmoji                          *IconEmoji            `json:"icon_emoji,omitempty"`
	ThemeColor                         *int64                `json:"theme_color,omitempty"`
	IsMessageRequest                   bool                  `json:"is_message_request,omitzero"`
	IsMessageRequestTimestamp          *time.Time            `json:"is_message_request_timestamp,omitempty"`
	IsSpam                             bool                  `json:"is_spam,omitzero"`
	Status                             *string               `json:"status,omitempty"`
	HdStreamingUntil                   *time.Time            `json:"hd_streaming_until,omitempty"`
	HdStreamingBuyerID                 *snowflake.ID         `json:"hd_streaming_buyer_id,omitempty"`
	LinkedLobby                        *LinkedLobby          `json:"linked_lobby,omitempty"`
	IsLinkable                         bool                  `json:"is_linkable,omitzero"`
	IsViewableAndWriteableByAllMembers bool                  `json:"is_viewable_and_writeable_by_all_members,omitzero"`
	Template                           string                `json:"template,omitzero"`
	Version                            string                `json:"version,omitzero"`
}

type ChannelActivityInstance struct {
	ApplicationID snowflake.ID   `json:"application_id"`
	InstanceID    string         `json:"instance_id"`
	ChannelID     snowflake.ID   `json:"channel_id"`
	GuildID       snowflake.ID   `json:"guild_id,omitzero"`
	Users         []snowflake.ID `json:"users"`
}

type ChannelAffinity struct {
	ChannelID snowflake.ID `json:"channel_id"`
	Affinity  float64      `json:"affinity"`
}

type ChannelCreateExtraFields struct {
	OriginChannelID snowflake.ID `json:"origin_channel_id,omitzero"`
}

type ChannelFollowingByChannel struct {
	DayPt                time.Time    `json:"day_pt"`
	ChannelID            snowflake.ID `json:"channel_id"`
	TotalGuildsFollowing int64        `json:"total_guilds_following"`
	NewGuildsFollowing   int64        `json:"new_guilds_following"`
	GuildsUnfollowed     int64        `json:"guilds_unfollowed"`
}

type ChannelFollowingGuildSize struct {
	DayPt                time.Time `json:"day_pt"`
	GuildSize            string    `json:"guild_size"`
	TotalGuildsFollowing int64     `json:"total_guilds_following"`
}

type ChannelFollowingGuildSizeByChannel struct {
	DayPt                time.Time    `json:"day_pt"`
	ChannelID            snowflake.ID `json:"channel_id"`
	GuildSize            string       `json:"guild_size"`
	TotalGuildsFollowing int64        `json:"total_guilds_following"`
}

type ChannelFollowingOverview struct {
	DayPt                time.Time `json:"day_pt"`
	TotalGuildsFollowing int64     `json:"total_guilds_following"`
	NewGuildsFollowing   int64     `json:"new_guilds_following"`
	GuildsUnfollowed     int64     `json:"guilds_unfollowed"`
}

type ChannelFollowingReach struct {
	DayPt              time.Time    `json:"day_pt"`
	ChannelID          snowflake.ID `json:"channel_id"`
	ChannelName        string       `json:"channel_name"`
	ReferenceMessageID snowflake.ID `json:"reference_message_id"`
	GuildsReached      int64        `json:"guilds_reached"`
}

type ChannelIconEmoji struct {
	ID    uint64 `json:"id,omitzero"`
	Name  string `json:"name,omitzero"`
	Color uint64 `json:"color,omitzero"`
}

type ChannelInfo struct {
	GuildID  snowflake.ID      `json:"guild_id"`
	Channels []ChannelMetadata `json:"channels"`
}

type ChannelListSettings struct {
	Layout          string `json:"layout,omitzero"`
	MessagePreviews string `json:"message_previews,omitzero"`
}

type ChannelMemberCountUpdate struct {
	GuildID       snowflake.ID `json:"guild_id"`
	ChannelID     snowflake.ID `json:"channel_id"`
	MemberCount   int64        `json:"member_count"`
	PresenceCount int64        `json:"presence_count"`
}

type ChannelMetadata struct {
	ID             snowflake.ID `json:"id"`
	Status         *string      `json:"status,omitempty"`
	VoiceStartTime *time.Time   `json:"voice_start_time,omitempty"`
}

type ChannelNick struct {
	ID   snowflake.ID `json:"id"`
	Nick string       `json:"nick"`
}

type ChannelOverride struct {
	ChannelID            snowflake.ID             `json:"channel_id"`
	Collapsed            bool                     `json:"collapsed"`
	Flags                ChannelOverrideFlags     `json:"flags,omitzero"`
	MessageNotifications MessageNotificationLevel `json:"message_notifications"`
	Muted                bool                     `json:"muted"`
	MuteConfig           *MuteConfig              `json:"mute_config"`
}

type ChannelPinsAck struct {
	ChannelID snowflake.ID `json:"channel_id"`
	Timestamp time.Time    `json:"timestamp"`
	Version   int64        `json:"version"`
}

type ChannelPinsUpdate struct {
	GuildID          snowflake.ID `json:"guild_id,omitzero"`
	ChannelID        snowflake.ID `json:"channel_id"`
	LastPinTimestamp *time.Time   `json:"last_pin_timestamp,omitempty"`
}

type ChannelRecipientAdd struct {
	ChannelID snowflake.ID `json:"channel_id"`
	User      PartialUser  `json:"user"`
	Nick      string       `json:"nick,omitzero"`
}

type ChannelRecipientRemove struct {
	ChannelID snowflake.ID `json:"channel_id"`
	User      PartialUser  `json:"user"`
}

type ChannelSelect struct {
	Type          ComponentType     `json:"type"`
	ID            int64             `json:"id,omitzero"`
	CustomID      string            `json:"custom_id"`
	ChannelTypes  []int64           `json:"channel_types,omitempty"`
	Placeholder   string            `json:"placeholder,omitzero"`
	DefaultValues []json.RawMessage `json:"default_values,omitempty"`
	MinValues     int64             `json:"min_values,omitzero"`
	MaxValues     int64             `json:"max_values,omitzero"`
	Required      bool              `json:"required,omitzero"`
	Disabled      bool              `json:"disabled,omitzero"`
}

type ChannelSettings struct {
	CollapsedInInbox              bool                          `json:"collapsed_in_inbox"`
	IconEmoji                     ChannelIconEmoji              `json:"icon_emoji,omitzero"`
	CustomNotificationSoundConfig CustomNotificationSoundConfig `json:"custom_notification_sound_config,omitzero"`
}

type ChannelStatus struct {
	ID     snowflake.ID `json:"id"`
	Status string       `json:"status"`
}

type ChannelStatuses struct {
	GuildID  snowflake.ID    `json:"guild_id"`
	Channels []ChannelStatus `json:"channels"`
}

type ChannelSync struct {
	GuildID        snowflake.ID `json:"guild_id"`
	Channels       []Channel    `json:"channels"`
	IntegrityCheck bool         `json:"integrity_check"`
}

type ChannelUnread struct {
	ID               snowflake.ID  `json:"id"`
	LastMessageID    *snowflake.ID `json:"last_message_id"`
	LastPinTimestamp *time.Time    `json:"last_pin_timestamp,omitempty"`
}

type ChannelUnreadUpdate struct {
	GuildID              snowflake.ID    `json:"guild_id"`
	ChannelUnreadUpdates []ChannelUnread `json:"channel_unread_updates"`
}

type CheckAgeVerificationForSuspendedUserParams struct {
	Token string `json:"token"`
}

type CheckAgeVerificationForSuspendedUserResponse struct {
	Success bool `json:"success"`
}

type CheckSocialLayerSKUPurchaseEligibilityResponse struct {
	InteractionID snowflake.ID `json:"interaction_id"`
}

type Checkbox struct {
	Type     ComponentType `json:"type"`
	ID       int64         `json:"id,omitzero"`
	CustomID string        `json:"custom_id"`
	Default  bool          `json:"default,omitzero"`
}

type CheckboxGroup struct {
	Type      ComponentType         `json:"type"`
	ID        int64                 `json:"id,omitzero"`
	CustomID  string                `json:"custom_id"`
	Options   []CheckboxGroupOption `json:"options"`
	MinValues int64                 `json:"min_values,omitzero"`
	MaxValues int64                 `json:"max_values,omitzero"`
	Required  bool                  `json:"required,omitzero"`
}

type CheckboxGroupOption struct {
	Value       string `json:"value"`
	Label       string `json:"label"`
	Description string `json:"description,omitzero"`
	Default     bool   `json:"default,omitzero"`
}

type CheckpointCard struct {
	Type           ComponentType  `json:"type"`
	ID             int64          `json:"id,omitzero"`
	CheckpointData CheckpointData `json:"checkpoint_data"`
}

type CheckpointCardEmoji struct {
	EmojiID   *snowflake.ID `json:"emoji_id"`
	EmojiName string        `json:"emoji_name"`
}

type CheckpointCardGame struct {
	ApplicationID      snowflake.ID `json:"application_id"`
	ApplicationName    string       `json:"application_name"`
	ApplicationImageID *string      `json:"application_image_id"`
}

type CheckpointCardGuild struct {
	GuildID   snowflake.ID `json:"guild_id"`
	GuildName string       `json:"guild_name"`
	GuildIcon *string      `json:"guild_icon"`
}

type CheckpointData struct {
	Version              int64               `json:"version"`
	CardID               CheckpointCardType  `json:"card_id"`
	PowerLevel           float64             `json:"power_level"`
	PowerLevelPercentile float64             `json:"power_level_percentile"`
	NumMessagesSent      int64               `json:"num_messages_sent"`
	TotalVoiceMinutes    float64             `json:"total_voice_minutes"`
	NumEmojisSent        int64               `json:"num_emojis_sent"`
	TopGuild             CheckpointCardGuild `json:"top_guild,omitzero"`
	TopEmoji             CheckpointCardEmoji `json:"top_emoji,omitzero"`
	TopGame              CheckpointCardGame  `json:"top_game,omitzero"`
}

type CheckpointEmojiStatistics struct {
	NumEmojisSent int64   `json:"num_emojis_sent"`
	Emojis        []Emoji `json:"emojis"`
}

type CheckpointGameStatistics struct {
	NumSessions int64              `json:"num_sessions"`
	Game        PartialApplication `json:"game"`
}

type CheckpointGamesStatistics struct {
	TotalGamesPlayed int64                      `json:"total_games_played"`
	Applications     []CheckpointGameStatistics `json:"applications"`
}

type CheckpointGuildStatistics struct {
	NumMessagesSent int64        `json:"num_messages_sent,omitzero"`
	NumVoiceMinutes float64      `json:"num_voice_minutes,omitzero"`
	Guild           PartialGuild `json:"guild"`
}

type CheckpointGuildsStatistics struct {
	NumGuildsJoined int64                       `json:"num_guilds_joined"`
	Guilds          []CheckpointGuildStatistics `json:"guilds"`
}

type CheckpointMessageStatistics struct {
	NumMessagesSent           int64                                `json:"num_messages_sent"`
	NumMessagesSentPercentile *float64                             `json:"num_messages_sent_percentile"`
	TopMonth                  *CheckpointMessageTopMonthStatistics `json:"top_month"`
}

type CheckpointMessageTopMonthStatistics struct {
	Month           int64 `json:"month"`
	NumMessagesSent int64 `json:"num_messages_sent"`
}

type CheckpointQuestStatistics struct {
	NumCompleted int64 `json:"num_completed"`
	NumOrbs      int64 `json:"num_orbs"`
}

type CheckpointSidekick struct {
	User            PartialUser `json:"user"`
	NumMessagesSent int64       `json:"num_messages_sent"`
	NumVoiceMinutes float64     `json:"num_voice_minutes"`
}

type CheckpointUser struct {
	User PartialUser `json:"user"`
}

type CheckpointVoiceStatistics struct {
	TotalVoiceMinutes           float64                            `json:"total_voice_minutes"`
	TotalVoiceMinutesPercentile *float64                           `json:"total_voice_minutes_percentile"`
	TopMonth                    *CheckpointVoiceTopMonthStatistics `json:"top_month"`
}

type CheckpointVoiceTopMonthStatistics struct {
	Month             int64   `json:"month"`
	NumMinutesInVoice float64 `json:"num_minutes_in_voice"`
}

type ClaimPremiumCollectiblesProductParams struct {
	SKUID snowflake.ID `json:"sku_id"`
}

type ClaimQuestCreativeRewardParams struct {
	DecisionMetadataSealed   *string `json:"decision_metadata_sealed,omitempty"`
	TrafficMetadataSealed    *string `json:"traffic_metadata_sealed,omitempty"`
	ClientHeartbeatSessionID string  `json:"client_heartbeat_session_id,omitzero"`
	ClientAdSessionID        string  `json:"client_ad_session_id,omitzero"`
}

type ClaimQuestRewardParams struct {
	Location              QuestContentType  `json:"location"`
	Platform              QuestPlatformType `json:"platform"`
	MetadataSealed        *string           `json:"metadata_sealed,omitempty"`
	TrafficMetadataSealed *string           `json:"traffic_metadata_sealed,omitempty"`
}

type ClaimQuestRewardResponse struct {
	ClaimedAt                     *time.Time                                     `json:"claimed_at"`
	EntitlementExpirationMetadata map[snowflake.ID]EntitlementExpirationMetadata `json:"entitlement_expiration_metadata"`
	Entitlements                  []Entitlement                                  `json:"entitlements"`
	Errors                        []json.RawMessage                              `json:"errors"`
}

type ClaimRewardCategoryProductParams struct {
	CategoryID snowflake.ID `json:"category_id"`
}

type ClaimSubscriptionPromotionRewardResponse struct {
	Eligible    bool                                         `json:"eligible"`
	Reason      SubscriptionPromotionRewardEligibilityReason `json:"reason"`
	Entitlement Entitlement                                  `json:"entitlement,omitzero"`
}

type ClaimedPromotion struct {
	Code      string       `json:"code"`
	UserID    snowflake.ID `json:"user_id"`
	ClaimedAt time.Time    `json:"claimed_at"`
	Promotion Promotion    `json:"promotion"`
}

type ClaimedQuest struct {
	ID         snowflake.ID       `json:"id"`
	Config     ClaimedQuestConfig `json:"config"`
	UserStatus QuestUserStatus    `json:"user_status"`
}

type ClaimedQuestConfig struct {
	ID        snowflake.ID         `json:"id"`
	StartsAt  time.Time            `json:"starts_at"`
	ExpiresAt time.Time            `json:"expires_at"`
	Features  []int64              `json:"features"`
	Colors    QuestGradient        `json:"colors"`
	Assets    QuestAssets          `json:"assets"`
	Messages  QuestMessages        `json:"messages"`
	Rewards   []ClaimedQuestReward `json:"rewards"`
}

type ClaimedQuestReward struct {
	Type               QuestRewardType `json:"type"`
	SKUID              snowflake.ID    `json:"sku_id"`
	Name               string          `json:"name"`
	NameWithArticle    string          `json:"name_with_article"`
	Asset              string          `json:"asset"`
	AssetVideo         *string         `json:"asset_video"`
	OrbQuantity        *int64          `json:"orb_quantity"`
	CollectibleProduct Collectibles    `json:"collectible_product,omitzero"`
}

type Classification struct {
	ID                  snowflake.ID        `json:"id"`
	ClassificationType  ClassificationType  `json:"classification_type"`
	Description         string              `json:"description"`
	ExplainerLink       string              `json:"explainer_link"`
	Actions             []SafetyHubAction   `json:"actions"`
	MaxExpirationTime   *time.Time          `json:"max_expiration_time"`
	FlaggedContent      []FlaggedContent    `json:"flagged_content"`
	AppealStatus        AppealStatus        `json:"appeal_status,omitzero"`
	IsCoppa             bool                `json:"is_coppa"`
	IsSpam              bool                `json:"is_spam"`
	AppealIngestionType AppealIngestionType `json:"appeal_ingestion_type"`
	GuildMetadata       GuildMetadata       `json:"guild_metadata,omitzero"`
}

type ClientDisconnect struct {
	UserID snowflake.ID `json:"user_id"`
}

type ClientEnvironmentConfig struct {
	CDNHost     string `json:"cdn_host,omitzero"`
	APIEndpoint string `json:"api_endpoint"`
	Environment string `json:"environment"`
}

type ClientFlags struct {
	UserID snowflake.ID `json:"user_id"`
	Flags  *VoiceFlags  `json:"flags"`
}

type ClientInfo struct {
	Client  ClientType                  `json:"client"`
	OS      PresenceOperatingSystemType `json:"os"`
	Version int64                       `json:"version"`
}

type ClientPlatform struct {
	UserID   snowflake.ID   `json:"user_id"`
	Platform *VoicePlatform `json:"platform"`
}

type ClientProperties struct {
	OS                          string                    `json:"os"`
	OSVersion                   string                    `json:"os_version,omitzero"`
	OSSdkVersion                string                    `json:"os_sdk_version,omitzero"`
	OSArch                      string                    `json:"os_arch,omitzero"`
	AppArch                     string                    `json:"app_arch,omitzero"`
	Browser                     string                    `json:"browser"`
	BrowserUserAgent            string                    `json:"browser_user_agent"`
	BrowserVersion              string                    `json:"browser_version"`
	ClientBuildNumber           int64                     `json:"client_build_number"`
	NativeBuildNumber           *int64                    `json:"native_build_number,omitempty"`
	ClientVersion               string                    `json:"client_version,omitzero"`
	ClientEventSource           *ClientEventSource        `json:"client_event_source,omitempty"`
	ClientAppState              ClientAppState            `json:"client_app_state,omitzero"`
	ClientLaunchID              string                    `json:"client_launch_id,omitzero"`
	ClientHeartbeatSessionID    *string                   `json:"client_heartbeat_session_id,omitempty"`
	ClientPerformanceCpu        int64                     `json:"client_performance_cpu,omitzero"`
	ClientPerformanceMemory     int64                     `json:"client_performance_memory,omitzero"`
	CpuCoreCount                int64                     `json:"cpu_core_count,omitzero"`
	ReleaseChannel              string                    `json:"release_channel"`
	SystemLocale                string                    `json:"system_locale"`
	Device                      string                    `json:"device,omitzero"`
	DeviceVendorID              string                    `json:"device_vendor_id,omitzero"`
	DeviceAdvertiserID          string                    `json:"device_advertiser_id,omitzero"`
	DesignID                    DesignID                  `json:"design_id,omitzero"`
	AccessibilitySupportEnabled bool                      `json:"accessibility_support_enabled,omitzero"`
	AccessibilityFeatures       AccessibilityFeatureFlags `json:"accessibility_features,omitzero"`
	WindowManager               string                    `json:"window_manager,omitzero"`
	Distro                      string                    `json:"distro,omitzero"`
	RuntimeEnvironment          RuntimeEnvironment        `json:"runtime_environment,omitzero"`
	DisplayServer               string                    `json:"display_server,omitzero"`
	Referrer                    string                    `json:"referrer,omitzero"`
	ReferrerCurrent             string                    `json:"referrer_current,omitzero"`
	ReferringDomain             string                    `json:"referring_domain,omitzero"`
	ReferringDomainCurrent      string                    `json:"referring_domain_current,omitzero"`
	SearchEngine                SearchEngine              `json:"search_engine,omitzero"`
	SearchEngineCurrent         string                    `json:"search_engine_current,omitzero"`
	MpKeyword                   string                    `json:"mp_keyword,omitzero"`
	MpKeywordCurrent            string                    `json:"mp_keyword_current,omitzero"`
	UtmCampaign                 string                    `json:"utm_campaign,omitzero"`
	UtmCampaignCurrent          string                    `json:"utm_campaign_current,omitzero"`
	UtmContent                  string                    `json:"utm_content,omitzero"`
	UtmContentCurrent           string                    `json:"utm_content_current,omitzero"`
	UtmMedium                   string                    `json:"utm_medium,omitzero"`
	UtmMediumCurrent            string                    `json:"utm_medium_current,omitzero"`
	UtmSource                   string                    `json:"utm_source,omitzero"`
	UtmSourceCurrent            string                    `json:"utm_source_current,omitzero"`
	UtmTerm                     string                    `json:"utm_term,omitzero"`
	UtmTermCurrent              string                    `json:"utm_term_current,omitzero"`
	HasClientMods               bool                      `json:"has_client_mods,omitzero"`
	LaunchSignature             string                    `json:"launch_signature,omitzero"`
	InstallationID              string                    `json:"installation_id,omitzero"`
	IsFastConnect               bool                      `json:"is_fast_connect,omitzero"`
	Version                     string                    `json:"version,omitzero"`
}

type ClientState struct {
	GuildVersions            map[snowflake.ID]int64 `json:"guild_versions"`
	HighestLastMessageID     snowflake.ID           `json:"highest_last_message_id,omitzero"`
	ReadStateVersion         int64                  `json:"read_state_version,omitzero"`
	UserGuildSettingsVersion int64                  `json:"user_guild_settings_version,omitzero"`
	UserSettingsVersion      int64                  `json:"user_settings_version,omitzero"`
	PrivateChannelsVersion   snowflake.ID           `json:"private_channels_version,omitzero"`
	APICodeVersion           int64                  `json:"api_code_version,omitzero"`
	InitialGuildID           snowflake.ID           `json:"initial_guild_id,omitzero"`
}

type ClientStatus struct {
	Desktop  StatusType `json:"desktop,omitzero"`
	Mobile   StatusType `json:"mobile,omitzero"`
	Web      StatusType `json:"web,omitzero"`
	Embedded StatusType `json:"embedded,omitzero"`
	VR       StatusType `json:"vr,omitzero"`
}

type ClientThemeSettings struct {
	PrimaryColor               int64                   `json:"primary_color,omitzero"`
	BackgroundGradientPresetID int64                   `json:"background_gradient_preset_id,omitzero"`
	BackgroundGradientAngle    float64                 `json:"background_gradient_angle,omitzero"`
	CustomUserThemeSettings    CustomUserThemeSettings `json:"custom_user_theme_settings,omitzero"`
}

type ClientsConnect struct {
	UserIDs []snowflake.ID `json:"user_ids"`
}

type ClipsSettings struct {
	AllowVoiceRecording bool `json:"allow_voice_recording,omitzero"`
}

type Close struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type CloseActivityJoinRequestArguments struct {
	UserID snowflake.ID `json:"user_id"`
}

type CloudAttachment struct {
	ID             *snowflake.ID `json:"id"`
	UploadURL      string        `json:"upload_url"`
	UploadFilename string        `json:"upload_filename"`
}

type CoachmarkCollectiblesMarketing struct {
	Type                CollectiblesMarketingType `json:"type"`
	Title               string                    `json:"title"`
	Body                string                    `json:"body"`
	AssetDark           string                    `json:"asset_dark"`
	AssetLight          string                    `json:"asset_light"`
	Version             int64                     `json:"version"`
	RefTargetBackground json.RawMessage           `json:"ref_target_background"`
	BadgeIcon           string                    `json:"badge_icon,omitzero"`
	BadgeText           string                    `json:"badge_text,omitzero"`
	ButtonLabel         string                    `json:"button_label,omitzero"`
}

type CoachtipCollectiblesMarketing struct {
	Type                CollectiblesMarketingType `json:"type"`
	Version             int64                     `json:"version"`
	Title               string                    `json:"title"`
	Body                string                    `json:"body"`
	Avatar              string                    `json:"avatar"`
	Decorations         []string                  `json:"decorations"`
	DismissibleContent  int64                     `json:"dismissible_content"`
	RefTargetBackground TargetBackground          `json:"ref_target_background"`
}

type Codec struct {
	Name           string `json:"name"`
	Type           string `json:"type"`
	Priority       int64  `json:"priority"`
	PayloadType    int64  `json:"payload_type"`
	RtxPayloadType int64  `json:"rtx_payload_type,omitzero"`
	Encode         bool   `json:"encode,omitzero"`
	Decode         bool   `json:"decode,omitzero"`
}

type CollectibleCategory struct {
	SKUID                    snowflake.ID         `json:"sku_id"`
	StoreListingID           snowflake.ID         `json:"store_listing_id"`
	BannerTextColor          string               `json:"banner_text_color,omitzero"`
	CatalogBannerAsset       BannerAsset          `json:"catalog_banner_asset,omitzero"`
	CatalogBannerAnimatedURL string               `json:"catalog_banner_animated_url,omitzero"`
	CatalogBannerRiveURL     string               `json:"catalog_banner_rive_url,omitzero"`
	FeaturedBlockBody        string               `json:"featured_block_body,omitzero"`
	FeaturedBlockURL         string               `json:"featured_block_url,omitzero"`
	HeroBannerAsset          BannerAsset          `json:"hero_banner_asset,omitzero"`
	HeroBannerAnimatedURL    string               `json:"hero_banner_animated_url,omitzero"`
	HeroBannerDisplayConfig  AssetConfig          `json:"hero_banner_display_config,omitzero"`
	HeroBlockTitle           string               `json:"hero_block_title,omitzero"`
	HeroLogoDisplayConfig    AssetConfig          `json:"hero_logo_display_config,omitzero"`
	HeroLogoURL              string               `json:"hero_logo_url,omitzero"`
	IsOrbsExclusive          bool                 `json:"is_orbs_exclusive,omitzero"`
	HeroRanking              []snowflake.ID       `json:"hero_ranking"`
	HeroRiveURL              string               `json:"hero_rive_url,omitzero"`
	LogoURL                  string               `json:"logo_url"`
	MobileBannerURL          string               `json:"mobile_banner_url,omitzero"`
	MobileBgURL              string               `json:"mobile_bg_url,omitzero"`
	MobileHeroBlockTitle     string               `json:"mobile_hero_block_title,omitzero"`
	MobileProductsTitle      string               `json:"mobile_products_title,omitzero"`
	MobileSummary            string               `json:"mobile_summary"`
	Name                     string               `json:"name"`
	PdpBgURL                 string               `json:"pdp_bg_url"`
	Products                 []CollectibleProduct `json:"products"`
	Styles                   CollectibleStyle     `json:"styles"`
	Summary                  string               `json:"summary"`
	UnpublishedAt            *time.Time           `json:"unpublished_at"`
	WideBannerAsset          AssetConfig          `json:"wide_banner_asset,omitzero"`
	WideBannerBody           string               `json:"wide_banner_body,omitzero"`
	WideBannerTitle          string               `json:"wide_banner_title,omitzero"`
}

type CollectiblePreviewAssets struct {
	FgStatic   string `json:"fg_static,omitzero"`
	FgAnimated string `json:"fg_animated,omitzero"`
	BgStatic   string `json:"bg_static,omitzero"`
	BgAnimated string `json:"bg_animated,omitzero"`
}

type CollectibleProduct struct {
	SKUID            snowflake.ID                 `json:"sku_id"`
	StoreListingID   snowflake.ID                 `json:"store_listing_id"`
	Type             CollectibleProductType       `json:"type"`
	BundledProducts  []CollectibleProduct         `json:"bundled_products,omitempty"`
	CategorySKUID    snowflake.ID                 `json:"category_sku_id"`
	GoogleSKUIDs     map[int64]string             `json:"google_sku_ids,omitempty"`
	EligibleOffers   []snowflake.ID               `json:"eligible_offers,omitempty"`
	Items            []json.RawMessage            `json:"items"`
	Name             string                       `json:"name"`
	PremiumType      int64                        `json:"premium_type"`
	Prices           map[int64]SubscriptionPrices `json:"prices,omitempty"`
	PreviewAssets    CollectiblePreviewAssets     `json:"preview_assets,omitzero"`
	Styles           CollectibleStyle             `json:"styles"`
	Summary          string                       `json:"summary"`
	UnpublishedAt    *time.Time                   `json:"unpublished_at"`
	BaseVariantName  string                       `json:"base_variant_name,omitzero"`
	BaseVariantSKUID snowflake.ID                 `json:"base_variant_sku_id,omitzero"`
	ExpiresAt        *time.Time                   `json:"expires_at,omitempty"`
	PurchaseType     CollectiblePurchaseType      `json:"purchase_type,omitzero"`
	PurchasedAt      time.Time                    `json:"purchased_at,omitzero"`
	BadgeOverride    *string                      `json:"badge_override,omitempty"`
	HideBadge        bool                         `json:"hide_badge,omitzero"`
	VariantLabel     string                       `json:"variant_label,omitzero"`
	VariantValue     string                       `json:"variant_value,omitzero"`
	Variants         []CollectibleProduct         `json:"variants,omitempty"`
}

type CollectibleStyle struct {
	BackgroundColors []int64 `json:"background_colors"`
	ButtonColors     []int64 `json:"button_colors"`
	ConfettiColors   []int64 `json:"confetti_colors"`
}

type CollectibleUserDiscount struct {
	Amount     int64        `json:"amount"`
	DiscountID snowflake.ID `json:"discount_id"`
	ExpiresAt  *time.Time   `json:"expires_at"`
}

type Collectibles struct {
	Nameplate *NameplateData `json:"nameplate"`
}

type CommandIndexApplication struct {
	ID                     snowflake.ID                       `json:"id"`
	Name                   string                             `json:"name"`
	Description            string                             `json:"description"`
	Icon                   *string                            `json:"icon"`
	Permissions            ApplicationCommandIndexPermissions `json:"permissions,omitzero"`
	Bot                    PartialUser                        `json:"bot,omitzero"`
	BotID                  snowflake.ID                       `json:"bot_id,omitzero"`
	Flags                  string                             `json:"flags"`
	EmbeddedActivityConfig EmbeddedActivityConfig             `json:"embedded_activity_config,omitzero"`
}

type CommunitiesSettings struct {
	DisableHomeAutoNav bool `json:"disable_home_auto_nav,omitzero"`
}

type Company struct {
	Name  string  `json:"name"`
	Roles []int64 `json:"roles"`
}

type ComponentsContentInventoryEntry struct {
	ID           string                      `json:"id"`
	AuthorID     snowflake.ID                `json:"author_id"`
	AuthorType   ContentInventoryAuthorType  `json:"author_type"`
	ContentType  ContentInventoryContentType `json:"content_type"`
	Traits       []ContentInventoryTrait     `json:"traits"`
	Extra        ContentInventoryMetadata    `json:"extra"`
	Participants []snowflake.ID              `json:"participants,omitempty"`
	ExpiresAt    time.Time                   `json:"expires_at,omitzero"`
	EndedAt      time.Time                   `json:"ended_at,omitzero"`
	StartedAt    time.Time                   `json:"started_at,omitzero"`
	OriginalID   snowflake.ID                `json:"original_id,omitzero"`
	GuildID      snowflake.ID                `json:"guild_id,omitzero"`
	ChannelID    snowflake.ID                `json:"channel_id,omitzero"`
	SessionID    snowflake.ID                `json:"session_id,omitzero"`
	Signature    ContentInventorySignature   `json:"signature"`
}

type ConfettiPotion struct {
	MessageEmoji Emoji `json:"message_emoji"`
}

type ConnectRequestProperties struct {
	HandoffType ConsoleHandoffType `json:"handoff_type"`
}

type ConnectedAccount struct {
	UserID      snowflake.ID `json:"user_id"`
	ProviderIDs []string     `json:"provider_ids"`
}

type Connection struct {
	ID                 string                  `json:"id"`
	Type               ConnectionType          `json:"type"`
	Name               string                  `json:"name"`
	Verified           bool                    `json:"verified"`
	Metadata           json.RawMessage         `json:"metadata,omitempty"`
	MetadataVisibility VisibilityType          `json:"metadata_visibility"`
	Revoked            bool                    `json:"revoked"`
	Integrations       []ConnectionIntegration `json:"integrations"`
	FriendSync         bool                    `json:"friend_sync"`
	ShowActivity       bool                    `json:"show_activity"`
	TwoWayLink         bool                    `json:"two_way_link"`
	Visibility         VisibilityType          `json:"visibility"`
	AccessToken        string                  `json:"access_token,omitzero"`
}

type ConnectionIntegration struct {
	ID      snowflake.ID     `json:"id"`
	Type    IntegrationType  `json:"type"`
	Account json.RawMessage  `json:"account"`
	Guild   IntegrationGuild `json:"guild"`
}

type ConnectionsCallbackArguments struct {
	Providertype ConnectionType  `json:"providerType"`
	Code         string          `json:"code"`
	OpenidParams json.RawMessage `json:"openid_params,omitempty"`
	Iss          string          `json:"iss,omitzero"`
	State        string          `json:"state"`
}

type ConsentStatus struct {
	Consented bool `json:"consented"`
}

type ConsoleCommandUpdate struct {
	ID     snowflake.ID    `json:"id"`
	Result string          `json:"result"`
	Error  json.RawMessage `json:"error"`
}

type ConsoleConnectionRequestData struct {
	AnalyticsProperties ConnectRequestProperties `json:"analytics_properties,omitzero"`
}

type ConsoleDevice struct {
	ID       snowflake.ID   `json:"id"`
	Name     string         `json:"name"`
	Platform ConnectionType `json:"platform"`
}

type ContactSyncSettingsResponse struct {
	AllowedInSuggestions ContactSyncSuggestionsSetting `json:"allowed_in_suggestions"`
}

type Container struct {
	Type        ComponentType `json:"type"`
	ID          int64         `json:"id,omitzero"`
	Components  Components    `json:"components"`
	AccentColor *int64        `json:"accent_color,omitempty"`
	Spoiler     bool          `json:"spoiler,omitzero"`
}

type ContentInventoryArtist struct {
	ExternalID string `json:"external_id"`
	Name       string `json:"name"`
}

type ContentInventoryEntry struct {
	Type                  ComponentType          `json:"type"`
	ID                    int64                  `json:"id,omitzero"`
	ContentInventoryEntry *ContentInventoryEntry `json:"content_inventory_entry"`
}

type ContentInventoryMetadata struct {
	Type                  ContentInventoryMetadataType `json:"type"`
	GameName              string                       `json:"game_name,omitzero"`
	ApplicationID         snowflake.ID                 `json:"application_id,omitzero"`
	Platform              ContentInventoryPlatformType `json:"platform,omitzero"`
	LastUpdate            time.Time                    `json:"last_update,omitzero"`
	Entries               []json.RawMessage            `json:"entries,omitempty"`
	Media                 json.RawMessage              `json:"media,omitempty"`
	Provider              ContentInventoryProviderType `json:"provider,omitzero"`
	MediaType             ContentInventoryMediaType    `json:"media_type,omitzero"`
	ParentTitle           string                       `json:"parent_title,omitzero"`
	Title                 string                       `json:"title,omitzero"`
	ImageURL              string                       `json:"image_url,omitzero"`
	Artist                json.RawMessage              `json:"artist,omitempty"`
	Artists               []json.RawMessage            `json:"artists,omitempty"`
	ExternalID            string                       `json:"external_id,omitzero"`
	ExternalParentID      string                       `json:"external_parent_id,omitzero"`
	MediaAssetsLargeImage string                       `json:"media_assets_large_image,omitzero"`
	MediaAssetsLargeText  string                       `json:"media_assets_large_text,omitzero"`
	MediaAssetsSmallImage string                       `json:"media_assets_small_image,omitzero"`
	MediaAssetsSmallText  string                       `json:"media_assets_small_text,omitzero"`
	MediaTitle            string                       `json:"media_title,omitzero"`
	MediaSubtitle         string                       `json:"media_subtitle,omitzero"`
	URL                   string                       `json:"url,omitzero"`
	ActivityName          string                       `json:"activity_name,omitzero"`
}

type ContentInventoryMetadataEntry struct {
	Media             json.RawMessage `json:"media,omitempty"`
	VerificationState int64           `json:"verification_state,omitzero"`
	RepeatCount       int64           `json:"repeat_count,omitzero"`
}

type ContentInventorySignature struct {
	Signature string `json:"signature"`
	Kid       string `json:"kid"`
	Version   int64  `json:"version"`
}

type ContentInventoryTrait struct {
	Type                  ContentInventoryTraitType          `json:"type"`
	FirstTime             bool                               `json:"first_time,omitzero"`
	DurationSeconds       int64                              `json:"duration_seconds,omitzero"`
	IsLive                bool                               `json:"is_live,omitzero"`
	Range                 ContentInventoryAggregateRangeType `json:"range,omitzero"`
	ResurrectedLastPlayed time.Time                          `json:"resurrected_last_played,omitzero"`
	Marathon              bool                               `json:"marathon,omitzero"`
	StreakCountDays       int64                              `json:"streak_count_days,omitzero"`
	Trending              ContentInventoryTrendingType       `json:"trending,omitzero"`
	Count                 int64                              `json:"count,omitzero"`
}

type ContentRating struct {
	Rating      ContentRatingAgency `json:"rating"`
	Descriptors []int64             `json:"descriptors"`
}

type ContentScanMetadata struct {
	Flags   ContentScanFlags `json:"flags"`
	Version int64            `json:"version"`
}

type ConversationSummary struct {
	ID         snowflake.ID   `json:"id"`
	Topic      string         `json:"topic"`
	SummShort  string         `json:"summ_short"`
	MessageIDs []snowflake.ID `json:"message_ids"`
	People     []snowflake.ID `json:"people"`
	Unsafe     bool           `json:"unsafe"`
	StartID    snowflake.ID   `json:"start_id"`
	EndID      snowflake.ID   `json:"end_id"`
	Count      int64          `json:"count"`
	Source     SummarySource  `json:"source"`
	Type       SummaryType    `json:"type"`
}

type ConversationSummaryUpdate struct {
	GuildID   snowflake.ID          `json:"guild_id"`
	ChannelID snowflake.ID          `json:"channel_id"`
	Summaries []ConversationSummary `json:"summaries"`
}

type CountdownTimerShopBlock struct {
	Type      ShopBlockType `json:"type"`
	Title     string        `json:"title"`
	Body      string        `json:"body"`
	BannerURL string        `json:"banner_url"`
	EndTime   time.Time     `json:"end_time"`
	TextColor string        `json:"text_color"`
}

type CountryPrices struct {
	CountryCode string      `json:"country_code"`
	Prices      []UnitPrice `json:"prices"`
}

type CreateApplicationAssetParams struct {
	Key              string  `json:"key"`
	UploadedFilename int64   `json:"uploaded_filename"`
	Visibility       *string `json:"visibility,omitempty"`
}

type CreateApplicationAttachmentParams struct {
	File json.RawMessage `json:"file"`
}

type CreateApplicationAttachmentResponse struct {
	Attachment Attachment `json:"attachment"`
}

type CreateApplicationBotResponse struct {
	Token *string `json:"token"`
}

type CreateApplicationEmojiParams struct {
	Name  string    `json:"name"`
	Image ImageData `json:"image"`
}

type CreateApplicationEntitlementParams struct {
	SKUID     snowflake.ID         `json:"sku_id"`
	OwnerID   snowflake.ID         `json:"owner_id"`
	OwnerType EntitlementOwnerType `json:"owner_type"`
}

type CreateApplicationGameClaimQuery struct {
	VerificationCode       string `json:"verification_code,omitzero"`
	ResetVerification      bool   `json:"reset_verification,omitzero"`
	ResendVerificationCode bool   `json:"resend_verification_code,omitzero"`
	FullName               string `json:"full_name,omitzero"`
	BusinessEmail          string `json:"business_email,omitzero"`
	BusinessRole           string `json:"business_role,omitzero"`
	BusinessWebsite        string `json:"business_website,omitzero"`
}

type CreateApplicationGiftCodeBatchParams struct {
	SKUID               snowflake.ID   `json:"sku_id"`
	Amount              int64          `json:"amount"`
	Description         string         `json:"description"`
	EntitlementBranches []snowflake.ID `json:"entitlement_branches,omitempty"`
	EntitlementStartsAt time.Time      `json:"entitlement_starts_at,omitzero"`
	EntitlementEndsAt   time.Time      `json:"entitlement_ends_at,omitzero"`
}

type CreateApplicationManagedLinkParams struct {
	CustomID    *string   `json:"custom_id,omitempty"`
	Description string    `json:"description"`
	Image       ImageData `json:"image"`
	Title       string    `json:"title"`
}

type CreateApplicationParams struct {
	Name         string           `json:"name"`
	Type         ApplicationType  `json:"type,omitzero"`
	TeamID       snowflake.ID     `json:"team_id,omitzero"`
	Description  *string          `json:"description,omitempty"`
	Icon         *ImageData       `json:"icon,omitempty"`
	CoverImage   *ImageData       `json:"cover_image,omitempty"`
	Flags        ApplicationFlags `json:"flags,omitzero"`
	GuildID      *snowflake.ID    `json:"guild_id,omitempty"`
	RedirectUris []string         `json:"redirect_uris,omitempty"`
	DeeplinkURI  *string          `json:"deeplink_uri,omitempty"`
}

type CreateApplicationQuickLinkParams struct {
	CustomID    *string   `json:"custom_id,omitempty"`
	Description string    `json:"description"`
	Image       ImageData `json:"image"`
	Title       string    `json:"title"`
}

type CreateApplicationWidgetConfigParams struct {
	DisplayName string                   `json:"display_name"`
	Surfaces    map[string]WidgetSurface `json:"surfaces"`
}

type CreateBillingPopupBridgeCallbackParams struct {
	State    string            `json:"state"`
	Path     string            `json:"path"`
	Query    map[string]string `json:"query,omitempty"`
	Insecure bool              `json:"insecure,omitzero"`
}

type CreateBillingPopupBridgeResponse struct {
	State string `json:"state"`
}

type CreateChannelInviteArguments struct {
	ChannelID snowflake.ID `json:"channel_id"`
}

type CreateChannelInviteParams struct {
	Flags               InviteFlags      `json:"flags,omitzero"`
	MaxAge              int64            `json:"max_age,omitzero"`
	MaxUses             int64            `json:"max_uses,omitzero"`
	Temporary           bool             `json:"temporary,omitzero"`
	Unique              bool             `json:"unique,omitzero"`
	TargetType          InviteTargetType `json:"target_type,omitzero"`
	TargetUserID        snowflake.ID     `json:"target_user_id,omitzero"`
	TargetApplicationID snowflake.ID     `json:"target_application_id,omitzero"`
	TargetUsersFile     json.RawMessage  `json:"target_users_file,omitempty"`
	RoleIDs             []snowflake.ID   `json:"role_ids,omitempty"`
}

type CreateChannelTagParams struct {
	Name      string        `json:"name"`
	Moderated bool          `json:"moderated,omitzero"`
	EmojiID   *snowflake.ID `json:"emoji_id,omitempty"`
	EmojiName *string       `json:"emoji_name,omitempty"`
}

type CreateChurnUserOfferResponse struct {
	Offer UserDiscountOffer `json:"offer"`
}

type CreateCompanyParams struct {
	Name string `json:"name"`
}

type CreateConsoleConnectionParams struct {
	AnalyticsProperties ConnectRequestProperties `json:"analytics_properties,omitzero"`
}

type CreateConsoleConnectionResponse struct {
	Nonce string `json:"nonce"`
}

type CreateContactSyncConnectionParams struct {
	Name       string `json:"name"`
	FriendSync bool   `json:"friend_sync,omitzero"`
}

type CreateDirectoryEntryParams struct {
	Type              DirectoryEntryType `json:"type,omitzero"`
	PrimaryCategoryID DirectoryCategory  `json:"primary_category_id,omitzero"`
	Description       *string            `json:"description,omitempty"`
}

type CreateFingerprintResponse struct {
	Fingerprint string `json:"fingerprint"`
}

type CreateGameInviteParams struct {
	RecipientID      snowflake.ID `json:"recipient_id"`
	LaunchParameters string       `json:"launch_parameters"`
	ApplicationAsset string       `json:"application_asset"`
	ApplicationName  string       `json:"application_name"`
	FallbackURL      *string      `json:"fallback_url,omitempty"`
	TTL              *int64       `json:"ttl,omitempty"`
}

type CreateGameInviteResponse struct {
	InviteID snowflake.ID `json:"invite_id"`
}

type CreateGravityAttachmentsParams struct {
	Files []UploadAttachment `json:"files"`
}

type CreateGreetMessageParams struct {
	StickerIDs       []snowflake.ID   `json:"sticker_ids"`
	AllowedMentions  AllowedMentions  `json:"allowed_mentions,omitzero"`
	MessageReference MessageReference `json:"message_reference,omitzero"`
}

type CreateGroupDMShellParams struct {
	RecipientID snowflake.ID `json:"recipient_id"`
}

type CreateGuildAutomodRuleParams struct {
	Name            string             `json:"name"`
	EventType       AutomodEventType   `json:"event_type"`
	TriggerType     AutomodTriggerType `json:"trigger_type"`
	TriggerMetadata json.RawMessage    `json:"trigger_metadata,omitempty"`
	Actions         []AutomodAction    `json:"actions"`
	Enabled         bool               `json:"enabled,omitzero"`
	ExemptRoles     []snowflake.ID     `json:"exempt_roles,omitempty"`
	ExemptChannels  []snowflake.ID     `json:"exempt_channels,omitempty"`
}

type CreateGuildBanParams struct {
	DeleteMessageDays    int64 `json:"delete_message_days,omitzero"`
	DeleteMessageSeconds int64 `json:"delete_message_seconds,omitzero"`
}

type CreateGuildChannelParams struct {
	Name                          string                `json:"name"`
	Position                      *int64                `json:"position,omitempty"`
	Type                          *ChannelType          `json:"type,omitempty"`
	Topic                         *string               `json:"topic,omitempty"`
	NSFW                          *bool                 `json:"nsfw,omitempty"`
	RateLimitPerUser              *int64                `json:"rate_limit_per_user,omitempty"`
	Bitrate                       *int64                `json:"bitrate,omitempty"`
	UserLimit                     *int64                `json:"user_limit,omitempty"`
	PermissionOverwrites          []PermissionOverwrite `json:"permission_overwrites,omitempty"`
	ParentID                      *snowflake.ID         `json:"parent_id,omitempty"`
	RTCRegion                     *string               `json:"rtc_region,omitempty"`
	VideoQualityMode              *VideoQualityMode     `json:"video_quality_mode,omitempty"`
	SKUID                         snowflake.ID          `json:"sku_id"`
	BranchID                      *snowflake.ID         `json:"branch_id,omitempty"`
	DefaultAutoArchiveDuration    *int64                `json:"default_auto_archive_duration,omitempty"`
	DefaultThreadRateLimitPerUser *int64                `json:"default_thread_rate_limit_per_user,omitempty"`
	AvailableTags                 []ForumTag            `json:"available_tags,omitempty"`
	DefaultReactionEmoji          *DefaultReaction      `json:"default_reaction_emoji,omitempty"`
	DefaultForumLayout            *ForumLayoutType      `json:"default_forum_layout,omitempty"`
	DefaultSortOrder              *SortOrderType        `json:"default_sort_order,omitempty"`
	DefaultTagSetting             *SearchTagSetting     `json:"default_tag_setting,omitempty"`
	Flags                         *ChannelFlags         `json:"flags,omitempty"`
}

type CreateGuildEmojiParams struct {
	Name  string         `json:"name"`
	Image ImageData      `json:"image"`
	Roles []snowflake.ID `json:"roles,omitempty"`
}

type CreateGuildIntegrationParams struct {
	Type IntegrationType `json:"type"`
	ID   string          `json:"id"`
}

type CreateGuildJoinRequestParams struct {
	FormFields []MemberVerificationFormField `json:"form_fields"`
	Version    *time.Time                    `json:"version"`
}

type CreateGuildOnboardingResponsesParams struct {
	OnboardingResponses     []snowflake.ID         `json:"onboarding_responses,omitempty"`
	OnboardingPromptsSeen   map[snowflake.ID]int64 `json:"onboarding_prompts_seen,omitempty"`
	OnboardingResponsesSeen map[snowflake.ID]int64 `json:"onboarding_responses_seen,omitempty"`
}

type CreateGuildParams struct {
	Name                        string                           `json:"name"`
	Description                 *string                          `json:"description,omitempty"`
	Region                      *string                          `json:"region,omitempty"`
	Icon                        *ImageData                       `json:"icon,omitempty"`
	VerificationLevel           *VerificationLevel               `json:"verification_level,omitempty"`
	DefaultMessageNotifications *MessageNotificationLevel        `json:"default_message_notifications,omitempty"`
	ExplicitContentFilter       *GuildExplicitContentFilterLevel `json:"explicit_content_filter,omitempty"`
	PreferredLocale             *string                          `json:"preferred_locale,omitempty"`
	Roles                       []PartialRole                    `json:"roles,omitempty"`
	Channels                    []PartialChannel                 `json:"channels,omitempty"`
	AFKChannelID                *snowflake.ID                    `json:"afk_channel_id,omitempty"`
	AFKTimeout                  *int64                           `json:"afk_timeout,omitempty"`
	SystemChannelID             *snowflake.ID                    `json:"system_channel_id,omitempty"`
	SystemChannelFlags          *SystemChannelFlags              `json:"system_channel_flags,omitempty"`
	GuildTemplateCode           *string                          `json:"guild_template_code,omitempty"`
	StaffOnly                   bool                             `json:"staff_only,omitzero"`
}

type CreateGuildProductAttachmentsParams struct {
	Files []UploadAttachment `json:"files"`
}

type CreateGuildProductListingParams struct {
	Name          string       `json:"name"`
	Description   string       `json:"description,omitzero"`
	PriceTier     int64        `json:"price_tier,omitzero"`
	CreateNewRole bool         `json:"create_new_role,omitzero"`
	Image         ImageData    `json:"image"`
	ImageName     string       `json:"image_name"`
	Attachments   []Attachment `json:"attachments,omitempty"`
}

type CreateGuildRoleParams struct {
	Name         *string      `json:"name,omitempty"`
	Description  *string      `json:"description,omitempty"`
	Color        *int64       `json:"color,omitempty"`
	Colors       *RoleColors  `json:"colors,omitempty"`
	Hoist        *bool        `json:"hoist,omitempty"`
	Icon         *ImageData   `json:"icon,omitempty"`
	UnicodeEmoji *string      `json:"unicode_emoji,omitempty"`
	Permissions  *Permissions `json:"permissions,omitempty"`
	Mentionable  *bool        `json:"mentionable,omitempty"`
}

type CreateGuildRoleSubscriptionGroupListingParams struct {
	Description string `json:"description,omitzero"`
}

type CreateGuildScheduledEventParams struct {
	ChannelID          snowflake.ID                  `json:"channel_id,omitzero"`
	EntityMetadata     json.RawMessage               `json:"entity_metadata,omitempty"`
	Name               string                        `json:"name"`
	PrivacyLevel       PrivacyLevel                  `json:"privacy_level"`
	ScheduledStartTime time.Time                     `json:"scheduled_start_time"`
	ScheduledEndTime   time.Time                     `json:"scheduled_end_time,omitzero"`
	Description        string                        `json:"description,omitzero"`
	EntityType         GuildScheduledEventEntityType `json:"entity_type"`
	Image              ImageData                     `json:"image,omitzero"`
	RecurrenceRule     json.RawMessage               `json:"recurrence_rule,omitempty"`
}

type CreateGuildScheduledEventParamsGuildScheduledEvent struct {
	OriginalScheduledStartTime time.Time  `json:"original_scheduled_start_time"`
	IsCanceled                 *bool      `json:"is_canceled,omitempty"`
	ScheduledStartTime         *time.Time `json:"scheduled_start_time,omitempty"`
	ScheduledEndTime           *time.Time `json:"scheduled_end_time,omitempty"`
}

type CreateGuildScheduledEventParamsGuildScheduledEventXxx struct {
	ChannelID          *snowflake.ID                 `json:"channel_id,omitempty"`
	EntityMetadata     json.RawMessage               `json:"entity_metadata,omitempty"`
	Name               string                        `json:"name,omitzero"`
	PrivacyLevel       PrivacyLevel                  `json:"privacy_level,omitzero"`
	ScheduledStartTime time.Time                     `json:"scheduled_start_time,omitzero"`
	ScheduledEndTime   time.Time                     `json:"scheduled_end_time,omitzero"`
	Description        *string                       `json:"description,omitempty"`
	EntityType         GuildScheduledEventEntityType `json:"entity_type,omitzero"`
	Status             GuildScheduledEventStatus     `json:"status,omitzero"`
	Image              ImageData                     `json:"image,omitzero"`
	RecurrenceRule     json.RawMessage               `json:"recurrence_rule,omitempty"`
}

type CreateGuildSoundboardSoundParams struct {
	Name      string          `json:"name"`
	Sound     json.RawMessage `json:"sound"`
	Volume    *float64        `json:"volume,omitempty"`
	EmojiID   *snowflake.ID   `json:"emoji_id,omitempty"`
	EmojiName *string         `json:"emoji_name,omitempty"`
}

type CreateGuildStickerParams struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Tags        string          `json:"tags"`
	File        json.RawMessage `json:"file"`
}

type CreateGuildTemplateParams struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

type CreateHandoffTokenParams struct {
	Key string `json:"key"`
}

type CreateHandoffTokenResponse struct {
	HandoffToken string `json:"handoff_token"`
}

type CreateHeadlessSessionParams struct {
	Activities []Activity `json:"activities"`
	Token      string     `json:"token,omitzero"`
}

type CreateHeadlessSessionResponse struct {
	Activities []Activity `json:"activities"`
	Token      string     `json:"token"`
}

type CreateInteractionParams struct {
	Type              InteractionType                `json:"type"`
	ApplicationID     snowflake.ID                   `json:"application_id"`
	GuildID           snowflake.ID                   `json:"guild_id,omitzero"`
	ChannelID         snowflake.ID                   `json:"channel_id"`
	MessageID         snowflake.ID                   `json:"message_id,omitzero"`
	MessageFlags      MessageFlags                   `json:"message_flags,omitzero"`
	SessionID         string                         `json:"session_id,omitzero"`
	Data              json.RawMessage                `json:"data"`
	FilesN            json.RawMessage                `json:"files[n],omitempty"`
	Nonce             string                         `json:"nonce,omitzero"`
	AnalyticsLocation InteractionCreationLocation    `json:"analytics_location,omitzero"`
	SectionName       InteractionCreationSectionName `json:"section_name,omitzero"`
	Source            InteractionCreationSourceType  `json:"source,omitzero"`
}

type CreateInteractionResponseQuery struct {
	WithResponse bool `json:"with_response,omitzero"`
}

type CreateInteractionResponseResponse struct {
	Interaction InteractionCallback         `json:"interaction"`
	Resource    InteractionCallbackResource `json:"resource,omitzero"`
}

type CreateLinkedUserRequestParams struct {
	RecipientID snowflake.ID `json:"recipient_id"`
	Code        string       `json:"code"`
}

type CreateLobbyInviteForCurrentUserResponse struct {
	Code string `json:"code"`
}

type CreateLobbyInviteResponse struct {
	Code string `json:"code"`
}

type CreateLobbyParams struct {
	Metadata           map[string]string `json:"metadata,omitempty"`
	Members            []LobbyMember     `json:"members,omitempty"`
	IdleTimeoutSeconds int64             `json:"idle_timeout_seconds,omitzero"`
	Flags              LobbyFlags        `json:"flags,omitzero"`
}

type CreateMessageAttachmentsParams struct {
	Files []UploadAttachment `json:"files"`
}

type CreateMessageParams struct {
	Content           string            `json:"content,omitzero"`
	TTS               bool              `json:"tts,omitzero"`
	Embeds            []Embed           `json:"embeds,omitempty"`
	Nonce             json.RawMessage   `json:"nonce,omitempty"`
	AllowedMentions   AllowedMentions   `json:"allowed_mentions,omitzero"`
	MessageReference  MessageReference  `json:"message_reference,omitzero"`
	Components        Components        `json:"components,omitzero"`
	StickerIDs        []snowflake.ID    `json:"sticker_ids,omitempty"`
	Activity          MessageActivity   `json:"activity,omitzero"`
	ApplicationID     snowflake.ID      `json:"application_id,omitzero"`
	Flags             MessageFlags      `json:"flags,omitzero"`
	FilesN            json.RawMessage   `json:"files[n],omitempty"`
	Attachments       []Attachment      `json:"attachments,omitempty"`
	Poll              PollCreate        `json:"poll,omitzero"`
	ConfettiPotion    ConfettiPotion    `json:"confetti_potion,omitzero"`
	SharedClientTheme SharedClientTheme `json:"shared_client_theme,omitzero"`
	WithCheckpoint    bool              `json:"with_checkpoint,omitzero"`
}

type CreateNotificationSettingsSnapshotParams struct {
	Label *string `json:"label,omitempty"`
}

type CreateOAuth2ApplicationAssetParams struct {
	Name  string    `json:"name"`
	Type  int64     `json:"type"`
	Image ImageData `json:"image"`
}

type CreateOAuth2AuthorizationParams struct {
	Authorize        bool                       `json:"authorize,omitzero"`
	IntegrationType  ApplicationIntegrationType `json:"integration_type,omitzero"`
	Permissions      Permissions                `json:"permissions,omitzero"`
	GuildID          snowflake.ID               `json:"guild_id,omitzero"`
	WebhookChannelID snowflake.ID               `json:"webhook_channel_id,omitzero"`
	DMSettings       ApplicationDMSettings      `json:"dm_settings,omitzero"`
	LocationContext  OAuth2LocationContext      `json:"location_context,omitzero"`
}

type CreateOAuth2AuthorizationQuery struct {
	ClientID            snowflake.ID `json:"client_id"`
	ResponseType        ResponseType `json:"response_type,omitzero"`
	Scope               string       `json:"scope"`
	RedirectURI         string       `json:"redirect_uri,omitzero"`
	State               string       `json:"state,omitzero"`
	Nonce               string       `json:"nonce,omitzero"`
	CodeChallenge       string       `json:"code_challenge,omitzero"`
	CodeChallengeMethod string       `json:"code_challenge_method,omitzero"`
}

type CreateOAuth2AuthorizationResponse struct {
	URL string `json:"url"`
}

type CreatePaymentSourceParams struct {
	Token               string             `json:"token"`
	PaymentGateway      PaymentGateway     `json:"payment_gateway"`
	BillingAddress      BillingAddress     `json:"billing_address"`
	BillingAddressToken string             `json:"billing_address_token,omitzero"`
	ReturnURL           string             `json:"return_url,omitzero"`
	Bank                string             `json:"bank,omitzero"`
	Default             bool               `json:"default,omitzero"`
	Pix                 PixPaymentMetadata `json:"pix,omitzero"`
}

type CreatePaymentSourceQuery struct {
	Location string `json:"location,omitzero"`
}

type CreatePaypalBillingAgreementTokenParams struct {
	ReturnURL string `json:"return_url"`
	CancelURL string `json:"cancel_url"`
}

type CreatePaypalBillingAgreementTokenResponse struct {
	Token string `json:"token"`
}

type CreatePollVoteParams struct {
	AnswerIDs []int64 `json:"answer_ids"`
}

type CreatePremiumGuildSubscriptionsParams struct {
	UserPremiumGuildSubscriptionSlotIDs []snowflake.ID `json:"user_premium_guild_subscription_slot_ids"`
	DisablePowerupAutoApply             bool           `json:"disable_powerup_auto_apply,omitzero"`
}

type CreatePrivateChannelParams struct {
	RecipientID  snowflake.ID            `json:"recipient_id,omitzero"`
	Recipients   []snowflake.ID          `json:"recipients,omitempty"`
	AccessTokens []string                `json:"access_tokens,omitempty"`
	Nicks        map[snowflake.ID]string `json:"nicks,omitempty"`
}

type CreateReactionQuery struct {
	Type ReactionType `json:"type,omitzero"`
}

type CreateRelationshipParams struct {
	Type                   RelationshipType `json:"type,omitzero"`
	FromFriendSuggestion   bool             `json:"from_friend_suggestion,omitzero"`
	ConfirmStrangerRequest bool             `json:"confirm_stranger_request,omitzero"`
	Note                   *string          `json:"note,omitempty"`
}

type CreateRemoteAuthSessionParams struct {
	Fingerprint string `json:"fingerprint"`
}

type CreateRemoteAuthSessionResponse struct {
	HandshakeToken string `json:"handshake_token"`
}

type CreateReportParams struct {
	MessageID snowflake.ID `json:"message_id"`
	ChannelID snowflake.ID `json:"channel_id"`
	UserID    snowflake.ID `json:"user_id"`
	Reason    int64        `json:"reason"`
}

type CreateReportResponse struct {
	ID snowflake.ID `json:"id"`
}

type CreateSKUParams struct {
	Type               SKUType                      `json:"type"`
	ApplicationID      snowflake.ID                 `json:"application_id"`
	Name               json.RawMessage              `json:"name"`
	Flags              SKUFlags                     `json:"flags,omitzero"`
	LegalNotice        json.RawMessage              `json:"legal_notice,omitempty"`
	DependentSKUID     snowflake.ID                 `json:"dependent_sku_id,omitzero"`
	BundledSKUs        []snowflake.ID               `json:"bundled_skus,omitempty"`
	AccessType         SKUAccessType                `json:"access_type,omitzero"`
	ManifestLabels     []snowflake.ID               `json:"manifest_labels,omitempty"`
	Features           []int64                      `json:"features,omitempty"`
	Locales            []string                     `json:"locales,omitempty"`
	Genres             []int64                      `json:"genres,omitempty"`
	ContentRatings     map[int64]ContentRating      `json:"content_ratings,omitempty"`
	SystemRequirements map[int64]SystemRequirements `json:"system_requirements,omitempty"`
	PriceTier          int64                        `json:"price_tier,omitzero"`
	Price              map[string]int64             `json:"price,omitempty"`
	SalePriceTier      int64                        `json:"sale_price_tier,omitzero"`
	SalePrice          map[string]int64             `json:"sale_price,omitempty"`
	ReleaseDate        string                       `json:"release_date,omitzero"`
}

type CreateSKUPurchaseParams struct {
	PaymentSourceID        snowflake.ID            `json:"payment_source_id,omitzero"`
	PaymentSourceToken     *string                 `json:"payment_source_token,omitempty"`
	ReturnURL              *string                 `json:"return_url,omitempty"`
	CountryCode            string                  `json:"country_code,omitzero"`
	Currency               string                  `json:"currency,omitzero"`
	SKUSubscriptionPlanID  snowflake.ID            `json:"sku_subscription_plan_id,omitzero"`
	Gift                   bool                    `json:"gift,omitzero"`
	GiftInfoOptions        GiftInfoOptions         `json:"gift_info_options,omitzero"`
	TestMode               bool                    `json:"test_mode,omitzero"`
	ExpectedAmount         int64                   `json:"expected_amount,omitzero"`
	ExpectedCurrency       string                  `json:"expected_currency,omitzero"`
	PurchaseToken          string                  `json:"purchase_token"`
	GatewayCheckoutContext *GatewayCheckoutContext `json:"gateway_checkout_context,omitempty"`
	LoadID                 string                  `json:"load_id"`
}

type CreateSKUPurchaseResponse struct {
	Entitlements         []Entitlement       `json:"entitlements"`
	LibraryApplications  []json.RawMessage   `json:"library_applications,omitempty"`
	AppliedUserDiscounts []UserDiscountOffer `json:"applied_user_discounts,omitempty"`
	GiftCode             string              `json:"gift_code,omitzero"`
}

type CreateSocialLayerSKUParams struct {
	Name      string `json:"name"`
	PriceTier int64  `json:"price_tier"`
}

type CreateSpeedTest struct {
	PreferredRegion *string `json:"preferred_region,omitempty"`
}

type CreateStageInstanceParams struct {
	ChannelID             snowflake.ID `json:"channel_id"`
	Topic                 string       `json:"topic"`
	PrivacyLevel          PrivacyLevel `json:"privacy_level,omitzero"`
	GuildScheduledEventID snowflake.ID `json:"guild_scheduled_event_id,omitzero"`
	SendStartNotification bool         `json:"send_start_notification,omitzero"`
}

type CreateStagedReportParams struct {
	Token         string `json:"token"`
	ReportType    string `json:"report_type"`
	ReportSubtype string `json:"report_subtype,omitzero"`
	Subject       string `json:"subject"`
	Description   string `json:"description"`
}

type CreateStoreListingParams struct {
	ApplicationID               snowflake.ID        `json:"application_id"`
	SKUID                       snowflake.ID        `json:"sku_id"`
	ChildSKUIDs                 []snowflake.ID      `json:"child_sku_ids,omitempty"`
	Summary                     json.RawMessage     `json:"summary"`
	Description                 json.RawMessage     `json:"description"`
	Tagline                     json.RawMessage     `json:"tagline,omitempty"`
	Published                   bool                `json:"published,omitzero"`
	CarouselItems               []StoreCarouselItem `json:"carousel_items,omitempty"`
	GuildID                     snowflake.ID        `json:"guild_id,omitzero"`
	ThumbnailAssetID            snowflake.ID        `json:"thumbnail_asset_id,omitzero"`
	PreviewVideoAssetID         snowflake.ID        `json:"preview_video_asset_id,omitzero"`
	HeaderBackgroundAssetID     snowflake.ID        `json:"header_background_asset_id,omitzero"`
	HeaderLogoDarkThemeAssetID  snowflake.ID        `json:"header_logo_dark_theme_asset_id,omitzero"`
	HeaderLogoLightThemeAssetID snowflake.ID        `json:"header_logo_light_theme_asset_id,omitzero"`
	BoxArtAssetID               snowflake.ID        `json:"box_art_asset_id,omitzero"`
	HeroBackgroundAssetID       snowflake.ID        `json:"hero_background_asset_id,omitzero"`
	HeroVideoAssetID            snowflake.ID        `json:"hero_video_asset_id,omitzero"`
}

type CreateStoreListingQuery struct {
	CountryCode string `json:"country_code,omitzero"`
	Localize    bool   `json:"localize,omitzero"`
}

type CreateStream struct {
	Type            StreamType    `json:"type"`
	GuildID         *snowflake.ID `json:"guild_id,omitempty"`
	ChannelID       snowflake.ID  `json:"channel_id"`
	PreferredRegion *string       `json:"preferred_region,omitempty"`
}

type CreateSubscriptionParams struct {
	Items                  []SubscriptionItem          `json:"items"`
	PaymentSourceID        snowflake.ID                `json:"payment_source_id,omitzero"`
	PaymentSourceToken     *string                     `json:"payment_source_token,omitempty"`
	ReturnURL              *string                     `json:"return_url,omitempty"`
	Currency               string                      `json:"currency,omitzero"`
	TrialID                *snowflake.ID               `json:"trial_id,omitempty"`
	ExpectedInvoicePrice   json.RawMessage             `json:"expected_invoice_price,omitempty"`
	ExpectedRenewalPrice   json.RawMessage             `json:"expected_renewal_price,omitempty"`
	PurchaseToken          string                      `json:"purchase_token"`
	GatewayCheckoutContext *GatewayCheckoutContext     `json:"gateway_checkout_context,omitempty"`
	Code                   string                      `json:"code,omitzero"`
	Metadata               SubscriptionMetadataRequest `json:"metadata,omitzero"`
	LoadID                 string                      `json:"load_id,omitzero"`
}

type CreateSubscriptionPreviewParams struct {
	Items             []SubscriptionItem         `json:"items"`
	Currency          string                     `json:"currency,omitzero"`
	PaymentSourceID   *snowflake.ID              `json:"payment_source_id,omitempty"`
	TrialID           *snowflake.ID              `json:"trial_id,omitempty"`
	ApplyEntitlements bool                       `json:"apply_entitlements,omitzero"`
	Renewal           bool                       `json:"renewal,omitzero"`
	Code              string                     `json:"code,omitzero"`
	Metadata          map[string]json.RawMessage `json:"metadata,omitempty"`
}

type CreateTeamIdentityVerificationParams struct {
	ReturnURL string `json:"return_url"`
}

type CreateTeamParams struct {
	Name string `json:"name"`
}

type CreateThreadFromMessageParams struct {
	Name                string `json:"name"`
	AutoArchiveDuration int64  `json:"auto_archive_duration,omitzero"`
	RateLimitPerUser    int64  `json:"rate_limit_per_user,omitzero"`
	Location            string `json:"location,omitzero"`
}

type CreateThreadParams struct {
	Name                string                   `json:"name"`
	AutoArchiveDuration int64                    `json:"auto_archive_duration,omitzero"`
	RateLimitPerUser    int64                    `json:"rate_limit_per_user,omitzero"`
	Location            string                   `json:"location,omitzero"`
	Type                ChannelType              `json:"type,omitzero"`
	Invitable           bool                     `json:"invitable,omitzero"`
	AppliedTags         []snowflake.ID           `json:"applied_tags,omitempty"`
	Message             ThreadOnlyChannelMessage `json:"message,omitzero"`
}

type CreateUniqueUsernameParams struct {
	Username string `json:"username"`
}

type CreateUserConnectionCallbackParams struct {
	Code           string          `json:"code"`
	State          string          `json:"state"`
	TwoWayLinkCode string          `json:"two_way_link_code,omitzero"`
	Insecure       bool            `json:"insecure,omitzero"`
	FriendSync     bool            `json:"friend_sync,omitzero"`
	OpenidParams   json.RawMessage `json:"openid_params,omitempty"`
}

type CreateUserGiftCodeParams struct {
	SKUID              snowflake.ID `json:"sku_id"`
	SubscriptionPlanID snowflake.ID `json:"subscription_plan_id,omitzero"`
	GiftStyle          GiftStyle    `json:"gift_style,omitzero"`
}

type CreateUserHarvestParams struct {
	Backends []string `json:"backends,omitempty"`
	Email    string   `json:"email"`
}

type CreateUserIdentityVerificationParams struct {
	ReturnURL string `json:"return_url"`
}

type CreateUserInviteParams struct {
	Code string `json:"code,omitzero"`
}

type CreateVideoFilterAssetParams struct {
	Type  VideoFilterType `json:"type"`
	Asset ImageData       `json:"asset"`
}

type CreateWebauthnAuthenticatorParams struct {
	Name       string `json:"name,omitzero"`
	Ticket     string `json:"ticket,omitzero"`
	Credential string `json:"credential,omitzero"`
}

type CreateWebauthnAuthenticatorResponse struct {
	Ticket      string       `json:"ticket"`
	Challenge   string       `json:"challenge"`
	ID          string       `json:"id"`
	Type        string       `json:"type"`
	Name        string       `json:"name"`
	BackupCodes []BackupCode `json:"backup_codes"`
}

type CreateWebhookParams struct {
	Name   string     `json:"name"`
	Avatar *ImageData `json:"avatar,omitempty"`
}

type CreatorMonetizationMarketingOnboardingSection struct {
	Type     CreatorMonetizationMarketingOnboardingSectionType     `json:"type"`
	Creators []CreatorMonetizationMarketingOnboardingSectionReview `json:"creators,omitempty"`
}

type CreatorMonetizationMarketingOnboardingSectionReview struct {
	GuildID               snowflake.ID `json:"guild_id"`
	Quote                 string       `json:"quote"`
	QuoteAttribution      string       `json:"quote_attribution"`
	QuoteAttributionTitle string       `json:"quote_attribution_title"`
}

type CreatorMonetizationRejection struct {
	CanReapplyAt *time.Time `json:"can_reapply_at,omitempty"`
}

type CreatorMonetizationRequest struct {
	ID               snowflake.ID                    `json:"id"`
	State            CreatorMonetizationRequestState `json:"state"`
	ActionedAt       *time.Time                      `json:"actioned_at"`
	RequesterID      snowflake.ID                    `json:"requester_id"`
	RequesterAckedAt *time.Time                      `json:"requester_acked_at"`
}

type CurrentGuildMemberUpdateSubscriptionArguments struct {
	GuildID snowflake.ID `json:"guild_id"`
}

type CustomCallSound struct {
	SoundID json.RawMessage `json:"sound_id"`
	GuildID json.RawMessage `json:"guild_id"`
}

type CustomNotificationSoundConfig struct {
	NotificationSoundPackID NotificationSoundPack `json:"notification_sound_pack_id,omitzero"`
}

type CustomStatus struct {
	Text      *string       `json:"text"`
	EmojiID   *snowflake.ID `json:"emoji_id"`
	EmojiName *string       `json:"emoji_name"`
	ExpiresAt *time.Time    `json:"expires_at"`
}

type CustomUserThemeSettings struct {
	Colors             []string        `json:"colors"`
	GradientColorStops []float64       `json:"gradient_color_stops"`
	GradientAngle      json.RawMessage `json:"gradient_angle"`
	BaseMix            json.RawMessage `json:"base_mix"`
}

type DMSettingsUpsellShow struct {
	GuildID snowflake.ID `json:"guild_id"`
}

type DebouncedReaction struct {
	Users []snowflake.ID `json:"users"`
	Emoji Emoji          `json:"emoji"`
}

type DebugSettings struct {
	RTCPanelShowVoiceStates bool `json:"rtc_panel_show_voice_states,omitzero"`
}

type DeepLinkArguments struct {
	Type   DeepLinkType   `json:"type"`
	Params DeepLinkParams `json:"params"`
}

type DeepLinkParams struct {
	Section       string              `json:"section,omitzero"`
	Subsection    string              `json:"subsection,omitzero"`
	Search        string              `json:"search,omitzero"`
	Fingerprint   *string             `json:"fingerprint,omitempty"`
	Date          string              `json:"date,omitzero"`
	Query         string              `json:"query,omitzero"`
	Skuid         snowflake.ID        `json:"skuId,omitzero"`
	Slug          string              `json:"slug,omitzero"`
	Guildid       snowflake.ID        `json:"guildId,omitzero"`
	Channelid     snowflake.ID        `json:"channelId,omitzero"`
	Messageid     snowflake.ID        `json:"messageId,omitzero"`
	Pageindex     int64               `json:"pageIndex,omitzero"`
	Sort          QuestHomeSortType   `json:"sort,omitzero"`
	Filter        QuestHomeFilterType `json:"filter,omitzero"`
	Tab           QuestHomeTab        `json:"tab,omitzero"`
	Questid       snowflake.ID        `json:"questId,omitzero"`
	Gameid        snowflake.ID        `json:"gameId,omitzero"`
	Token         string              `json:"token,omitzero"`
	Path          string              `json:"path,omitzero"`
	Attemptid     string              `json:"attemptId,omitzero"`
	Applicationid snowflake.ID        `json:"applicationId,omitzero"`
	URL           string              `json:"url,omitzero"`
}

type DefaultReaction struct {
	EmojiID   *snowflake.ID `json:"emoji_id"`
	EmojiName *string       `json:"emoji_name"`
}

type DefaultReactionEmoji struct {
	EmojiID          uint64 `json:"emoji_id,omitzero"`
	EmojiName        string `json:"emoji_name,omitzero"`
	Animated         bool   `json:"animated,omitzero"`
	DisableDoubleTap bool   `json:"disable_double_tap,omitzero"`
}

type DeleteChannelQuery struct {
	Silent bool `json:"silent,omitzero"`
}

type DeleteGuildProductListingResponse struct {
	URL string `json:"url"`
}

type DeleteGuildRoleSubscriptionGroupListingParams struct {
	Name                 string                    `json:"name,omitzero"`
	Description          string                    `json:"description,omitzero"`
	Image                *ImageData                `json:"image,omitempty"`
	PriceTier            int64                     `json:"price_tier,omitzero"`
	Benefits             []RoleSubscriptionBenefit `json:"benefits,omitempty"`
	CanAccessAllChannels bool                      `json:"can_access_all_channels,omitzero"`
}

type DeleteGuildScheduledEventUserParams struct {
	Response GuildScheduledEventUserResponse `json:"response"`
}

type DeleteGuildScheduledEventUserQuery struct {
	Before     snowflake.ID `json:"before,omitzero"`
	After      snowflake.ID `json:"after,omitzero"`
	Limit      float64      `json:"limit,omitzero"`
	WithMember bool         `json:"with_member,omitzero"`
}

type DeleteHeadlessSessionParams struct {
	Token string `json:"token"`
}

type DeleteNotificationCenterItemParams struct {
	ItemType NotificationCenterItemDeletionType `json:"item_type,omitzero"`
}

type DeleteReadStateParams struct {
	ReadStateType ReadStateType `json:"read_state_type,omitzero"`
	Version       int64         `json:"version,omitzero"`
}

type DeleteStream struct {
	StreamKey string `json:"stream_key"`
}

type DeleteSubscriptionQuery struct {
	Location      string `json:"location,omitzero"`
	LocationStack string `json:"location_stack,omitzero"`
}

type DeleteUserAccountParams struct {
	Password *string `json:"password"`
}

type DeletedEntityIDs struct {
	GuildID  snowflake.ID   `json:"guild_id"`
	Stickers []snowflake.ID `json:"stickers,omitempty"`
	Roles    []snowflake.ID `json:"roles,omitempty"`
	Emojis   []snowflake.ID `json:"emojis,omitempty"`
	Channels []snowflake.ID `json:"channels,omitempty"`
}

type DesktopHello struct {
	HeartbeatInterval int64 `json:"heartbeat_interval"`
	TimeoutMs         int64 `json:"timeout_ms"`
}

type DetectableApplication struct {
	ID                       snowflake.ID            `json:"id"`
	Name                     string                  `json:"name"`
	IconHash                 *string                 `json:"icon_hash,omitempty"`
	CoverImageHash           *string                 `json:"cover_image_hash,omitempty"`
	Aliases                  []string                `json:"aliases"`
	Executables              []ApplicationExecutable `json:"executables"`
	Themes                   []string                `json:"themes"`
	Hook                     bool                    `json:"hook"`
	Overlay                  bool                    `json:"overlay"`
	OverlayMethods           *OverlayMethodFlags     `json:"overlay_methods"`
	OverlayWarn              bool                    `json:"overlay_warn"`
	OverlayCompatibilityHook bool                    `json:"overlay_compatibility_hook"`
	LinkedApplications       []LinkedApplication     `json:"linked_applications,omitempty"`
}

type DirectoryEntry struct {
	Type                DirectoryEntryType           `json:"type"`
	DirectoryChannelID  snowflake.ID                 `json:"directory_channel_id"`
	EntityID            snowflake.ID                 `json:"entity_id"`
	CreatedAt           string                       `json:"created_at"`
	PrimaryCategoryID   DirectoryCategory            `json:"primary_category_id,omitzero"`
	Description         *string                      `json:"description"`
	AuthorID            snowflake.ID                 `json:"author_id"`
	Guild               DirectoryGuild               `json:"guild,omitzero"`
	GuildScheduledEvent DirectoryGuildScheduledEvent `json:"guild_scheduled_event,omitzero"`
}

type DirectoryGuild struct {
	FeaturableInDirectory bool `json:"featurable_in_directory"`
}

type DirectoryGuildScheduledEvent struct {
	Guild    PartialGuild `json:"guild"`
	UserRsvp bool         `json:"user_rsvp,omitzero"`
}

type DisableSMSMFAParams struct {
	Password string `json:"password"`
}

type DisableTotpMFAResponse struct {
	Token string `json:"token"`
}

type DisableUserAccountParams struct {
	Password string `json:"password"`
}

type Discount struct {
	ID                          snowflake.ID         `json:"id"`
	Amount                      int64                `json:"amount"`
	StartsAt                    *time.Time           `json:"starts_at,omitempty"`
	EndsAt                      *time.Time           `json:"ends_at,omitempty"`
	Status                      int64                `json:"status"`
	CreatedAt                   time.Time            `json:"created_at"`
	SKUIDs                      []snowflake.ID       `json:"sku_ids"`
	SKUGroupIDs                 []snowflake.ID       `json:"sku_group_ids"`
	PlanIDs                     []snowflake.ID       `json:"plan_ids"`
	UserUsageLimitInterval      SubscriptionInterval `json:"user_usage_limit_interval"`
	UserUsageLimitIntervalCount int64                `json:"user_usage_limit_interval_count"`
	UserUsageLimit              int64                `json:"user_usage_limit"`
}

type DiscoverableGuild struct {
	ID                       snowflake.ID             `json:"id"`
	Name                     string                   `json:"name"`
	Icon                     *string                  `json:"icon"`
	Description              *string                  `json:"description"`
	Banner                   *string                  `json:"banner"`
	Splash                   *string                  `json:"splash"`
	DiscoverySplash          *string                  `json:"discovery_splash"`
	Features                 []string                 `json:"features"`
	VanityURLCode            *string                  `json:"vanity_url_code"`
	PreferredLocale          string                   `json:"preferred_locale"`
	PremiumSubscriptionCount int64                    `json:"premium_subscription_count"`
	ApproximateMemberCount   int64                    `json:"approximate_member_count"`
	ApproximatePresenceCount int64                    `json:"approximate_presence_count"`
	Emojis                   []Emoji                  `json:"emojis,omitempty"`
	EmojiCount               int64                    `json:"emoji_count,omitzero"`
	Stickers                 []Sticker                `json:"stickers,omitempty"`
	StickerCount             int64                    `json:"sticker_count,omitzero"`
	AutoRemoved              bool                     `json:"auto_removed"`
	PrimaryCategoryID        int64                    `json:"primary_category_id"`
	PrimaryCategory          DiscoveryCategory        `json:"primary_category,omitzero"`
	Keywords                 []string                 `json:"keywords"`
	IsPublished              bool                     `json:"is_published"`
	ReasonsToJoin            []DiscoveryReason        `json:"reasons_to_join,omitempty"`
	SocialLinks              []string                 `json:"social_links,omitempty"`
	About                    *string                  `json:"about,omitempty"`
	CategoryIDs              []snowflake.ID           `json:"category_ids,omitempty"`
	Categories               []DiscoveryCategory      `json:"categories,omitempty"`
	CreatedAt                time.Time                `json:"created_at,omitzero"`
	NSFWProperties           *DiscoveryNSFWProperties `json:"nsfw_properties,omitempty"`
}

type DiscoveryCategory struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	IsPrimary bool   `json:"is_primary"`
}

type DiscoveryGameActivity struct {
	ActivityLevel int64 `json:"activity_level"`
	ActivityScore int64 `json:"activity_score"`
}

type DiscoveryHealthScore struct {
	AvgNonnewCommunicators *string  `json:"avg_nonnew_communicators"`
	AvgNonnewParticipators *string  `json:"avg_nonnew_participators"`
	NumIntentfulJoiners    *string  `json:"num_intentful_joiners"`
	PercRetW1Intentful     *float64 `json:"perc_ret_w1_intentful"`
}

type DiscoveryMetadata struct {
	GuildID                     snowflake.ID      `json:"guild_id"`
	PrimaryCategoryID           int64             `json:"primary_category_id"`
	Keywords                    []string          `json:"keywords"`
	EmojiDiscoverabilityEnabled bool              `json:"emoji_discoverability_enabled"`
	PartnerActionedTimestamp    *time.Time        `json:"partner_actioned_timestamp"`
	PartnerApplicationTimestamp *time.Time        `json:"partner_application_timestamp"`
	IsPublished                 bool              `json:"is_published"`
	ReasonsToJoin               []json.RawMessage `json:"reasons_to_join"`
	SocialLinks                 []string          `json:"social_links"`
	About                       *string           `json:"about"`
	CategoryIDs                 []int64           `json:"category_ids"`
}

type DiscoveryNSFWProperties struct {
	Channels                  []snowflake.ID            `json:"channels,omitempty"`
	ChannelBannedKeywords     map[snowflake.ID][]string `json:"channel_banned_keywords,omitempty"`
	Name                      string                    `json:"name,omitzero"`
	NameBannedKeywords        []string                  `json:"name_banned_keywords,omitempty"`
	Description               string                    `json:"description,omitzero"`
	DescriptionBannedKeywords []string                  `json:"description_banned_keywords,omitempty"`
}

type DiscoveryReason struct {
	Reason    string        `json:"reason"`
	EmojiID   *snowflake.ID `json:"emoji_id"`
	EmojiName *string       `json:"emoji_name"`
}

type DiscoveryRequirements struct {
	GuildID                      snowflake.ID            `json:"guild_id,omitzero"`
	SafeEnvironment              bool                    `json:"safe_environment,omitzero"`
	Healthy                      bool                    `json:"healthy,omitzero"`
	HealthScorePending           bool                    `json:"health_score_pending,omitzero"`
	Size                         bool                    `json:"size,omitzero"`
	NSFWProperties               DiscoveryNSFWProperties `json:"nsfw_properties,omitzero"`
	Protected                    bool                    `json:"protected,omitzero"`
	Sufficient                   bool                    `json:"sufficient"`
	SufficientWithoutGracePeriod bool                    `json:"sufficient_without_grace_period"`
	ValidRulesChannel            bool                    `json:"valid_rules_channel,omitzero"`
	RetentionHealthy             bool                    `json:"retention_healthy,omitzero"`
	EngagementHealthy            bool                    `json:"engagement_healthy,omitzero"`
	Age                          bool                    `json:"age,omitzero"`
	MinimumAge                   *int64                  `json:"minimum_age,omitempty"`
	HealthScore                  DiscoveryHealthScore    `json:"health_score,omitzero"`
	MinimumSize                  *int64                  `json:"minimum_size,omitempty"`
	GracePeriodEndDate           time.Time               `json:"grace_period_end_date,omitzero"`
}

type DisplayNameStyle struct {
	FontID   DisplayNameFont   `json:"font_id"`
	EffectID DisplayNameEffect `json:"effect_id"`
	Colors   []int64           `json:"colors"`
}

type DropdownElement struct {
	Title   string           `json:"title"`
	Options []DropdownOption `json:"options"`
}

type DropdownOption struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

type DynamicImageData struct {
	Name  string          `json:"name"`
	Value json.RawMessage `json:"value"`
	Type  int64           `json:"type"`
}

type DynamicNumberData struct {
	Name  string `json:"name"`
	Value int64  `json:"value"`
	Type  int64  `json:"type"`
}

type DynamicStringData struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Type  int64  `json:"type"`
}

type EditDMMessageParams struct {
	Content *string `json:"content,omitempty"`
}

type EditMessageParams struct {
	Content         *string          `json:"content,omitempty"`
	Embeds          []Embed          `json:"embeds,omitempty"`
	AllowedMentions *AllowedMentions `json:"allowed_mentions,omitempty"`
	Components      *Components      `json:"components,omitempty"`
	Flags           *MessageFlags    `json:"flags,omitempty"`
	FilesN          json.RawMessage  `json:"files[n]"`
	Attachments     []Attachment     `json:"attachments"`
}

type EditWebhookMessageParams struct {
	Content         string          `json:"content,omitzero"`
	TTS             bool            `json:"tts,omitzero"`
	Embeds          []Embed         `json:"embeds,omitempty"`
	AllowedMentions AllowedMentions `json:"allowed_mentions,omitzero"`
	Components      Components      `json:"components,omitzero"`
	Flags           MessageFlags    `json:"flags,omitzero"`
	FilesN          json.RawMessage `json:"files[n],omitempty"`
	Attachments     []Attachment    `json:"attachments,omitempty"`
}

type EmailSettings struct {
	Initialized bool            `json:"initialized"`
	Categories  map[string]bool `json:"categories"`
}

type Embed struct {
	Title              string        `json:"title,omitzero"`
	Type               string        `json:"type,omitzero"`
	Description        string        `json:"description,omitzero"`
	URL                string        `json:"url,omitzero"`
	Timestamp          time.Time     `json:"timestamp,omitzero"`
	Color              int64         `json:"color,omitzero"`
	Footer             EmbedFooter   `json:"footer,omitzero"`
	Image              EmbedMedia    `json:"image,omitzero"`
	Thumbnail          EmbedMedia    `json:"thumbnail,omitzero"`
	Video              EmbedMedia    `json:"video,omitzero"`
	Provider           EmbedProvider `json:"provider,omitzero"`
	Author             EmbedAuthor   `json:"author,omitzero"`
	Fields             []EmbedField  `json:"fields,omitempty"`
	ReferenceID        snowflake.ID  `json:"reference_id,omitzero"`
	ContentScanVersion int64         `json:"content_scan_version,omitzero"`
	Flags              EmbedFlags    `json:"flags,omitzero"`
}

type EmbedAuthor struct {
	Name         string `json:"name"`
	URL          string `json:"url,omitzero"`
	IconURL      string `json:"icon_url,omitzero"`
	ProxyIconURL string `json:"proxy_icon_url,omitzero"`
}

type EmbedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline,omitzero"`
}

type EmbedFooter struct {
	Text         string `json:"text"`
	IconURL      string `json:"icon_url,omitzero"`
	ProxyIconURL string `json:"proxy_icon_url,omitzero"`
}

type EmbedMedia struct {
	URL                 string              `json:"url"`
	ProxyURL            string              `json:"proxy_url,omitzero"`
	Height              int64               `json:"height,omitzero"`
	Width               int64               `json:"width,omitzero"`
	Flags               AttachmentFlags     `json:"flags,omitzero"`
	Description         string              `json:"description,omitzero"`
	ContentType         string              `json:"content_type,omitzero"`
	ContentScanMetadata ContentScanMetadata `json:"content_scan_metadata,omitzero"`
	PlaceholderVersion  int64               `json:"placeholder_version,omitzero"`
	Placeholder         string              `json:"placeholder,omitzero"`
}

type EmbedProvider struct {
	Name string `json:"name,omitzero"`
	URL  string `json:"url,omitzero"`
}

type EmbedResponse struct {
	Headers map[string]string `json:"headers"`
	Body    string            `json:"body"`
}

type EmbeddedActivityConfig struct {
	ApplicationID                     snowflake.ID                              `json:"application_id,omitzero"`
	ActivityPreviewVideoAssetID       *snowflake.ID                             `json:"activity_preview_video_asset_id"`
	SupportedPlatforms                []string                                  `json:"supported_platforms"`
	DefaultOrientationLockState       EmbeddedActivityOrientationLockStateType  `json:"default_orientation_lock_state"`
	TabletDefaultOrientationLockState EmbeddedActivityOrientationLockStateType  `json:"tablet_default_orientation_lock_state"`
	RequiresAgeGate                   bool                                      `json:"requires_age_gate"`
	LegacyResponsiveAspectRatio       bool                                      `json:"legacy_responsive_aspect_ratio"`
	PremiumTierRequirement            *PremiumTier                              `json:"premium_tier_requirement"`
	FreePeriodStartsAt                *time.Time                                `json:"free_period_starts_at"`
	FreePeriodEndsAt                  *time.Time                                `json:"free_period_ends_at"`
	ClientPlatformConfig              map[string]EmbeddedActivityPlatformConfig `json:"client_platform_config"`
	ShelfRank                         int64                                     `json:"shelf_rank"`
	HasCspException                   bool                                      `json:"has_csp_exception"`
	DisplaysAdvertisements            bool                                      `json:"displays_advertisements"`
	SupportedLocales                  []string                                  `json:"supported_locales"`
	BlockedLocales                    []string                                  `json:"blocked_locales"`
}

type EmbeddedActivityInstance struct {
	ApplicationID       snowflake.ID                  `json:"application_id"`
	CompositeInstanceID string                        `json:"composite_instance_id"`
	InstanceID          snowflake.ID                  `json:"instance_id"`
	LaunchID            snowflake.ID                  `json:"launch_id"`
	Location            EmbeddedActivityLocation      `json:"location"`
	Participants        []EmbeddedActivityParticipant `json:"participants"`
}

type EmbeddedActivityLocation struct {
	ID        string                       `json:"id"`
	Kind      EmbeddedActivityLocationType `json:"kind"`
	ChannelID snowflake.ID                 `json:"channel_id"`
	GuildID   *snowflake.ID                `json:"guild_id,omitempty"`
}

type EmbeddedActivityParticipant struct {
	UserID    snowflake.ID `json:"user_id"`
	Member    GuildMember  `json:"member,omitzero"`
	SessionID string       `json:"session_id"`
	Nonce     *string      `json:"nonce"`
}

type EmbeddedActivityPlatformConfig struct {
	LabelType             EmbeddedActivityLabelType    `json:"label_type"`
	LabelUntil            *time.Time                   `json:"label_until,omitempty"`
	ReleasePhase          EmbeddedActivityReleasePhase `json:"release_phase"`
	OmitBadgeFromSurfaces []string                     `json:"omit_badge_from_surfaces"`
}

type Emoji struct {
	ID            *snowflake.ID  `json:"id"`
	Name          string         `json:"name"`
	Roles         []snowflake.ID `json:"roles,omitempty"`
	User          PartialUser    `json:"user,omitzero"`
	RequireColons bool           `json:"require_colons,omitzero"`
	Managed       bool           `json:"managed,omitzero"`
	Animated      bool           `json:"animated,omitzero"`
	Available     bool           `json:"available,omitzero"`
	Version       string         `json:"version,omitzero"`
}

type EmojiApplication struct {
	ID   snowflake.ID `json:"id"`
	Name string       `json:"name"`
}

type EmojiFrecency struct {
	Emojis map[string]FrecencyItem `json:"emojis"`
}

type EmojiGuild struct {
	ID                       snowflake.ID `json:"id"`
	Name                     string       `json:"name"`
	Icon                     *string      `json:"icon"`
	Description              *string      `json:"description"`
	Features                 []string     `json:"features"`
	Emojis                   []Emoji      `json:"emojis"`
	PremiumTier              PremiumTier  `json:"premium_tier"`
	PremiumSubscriptionCount int64        `json:"premium_subscription_count"`
	ApproximateMemberCount   int64        `json:"approximate_member_count"`
	ApproximatePresenceCount int64        `json:"approximate_presence_count"`
}

type EmojiReactionFrecency struct {
	Emojis map[string]FrecencyItem `json:"emojis"`
}

type EnableSMSMFAParams struct {
	Password string `json:"password"`
}

type EnableSocialSdkParams struct {
	Name                string `json:"name"`
	BusinessEmail       string `json:"business_email"`
	GameOrStudioName    string `json:"game_or_studio_name,omitzero"`
	GameOrStudioURL     string `json:"game_or_studio_url,omitzero"`
	EmailUpdatesConsent bool   `json:"email_updates_consent,omitzero"`
}

type EnableTotpMFAParams struct {
	Password string `json:"password"`
	Secret   string `json:"secret,omitzero"`
	Code     string `json:"code,omitzero"`
}

type EnableTotpMFAResponse struct {
	Token       string       `json:"token"`
	BackupCodes []BackupCode `json:"backup_codes"`
}

type EngagementBase struct {
	DayPt                   time.Time `json:"day_pt"`
	Visitors                int64     `json:"visitors"`
	Communicators           int64     `json:"communicators"`
	PctCommunicators        float64   `json:"pct_communicators,omitzero"`
	Messages                int64     `json:"messages"`
	MessagesPerCommunicator float64   `json:"messages_per_communicator"`
	SpeakingMinutes         int64     `json:"speaking_minutes"`
}

type EngagementMuters struct {
	DayPt       time.Time `json:"day_pt"`
	DaysInGuild string    `json:"days_in_guild"`
	Muters      int64     `json:"muters"`
}

type EngagementOverview struct {
	DayPt           time.Time `json:"day_pt"`
	Visitors        int64     `json:"visitors"`
	Communicators   int64     `json:"communicators"`
	Messages        int64     `json:"messages"`
	SpeakingMinutes int64     `json:"speaking_minutes"`
}

type EngagementPruneableMembers struct {
	DayPt    time.Time `json:"day_pt"`
	Inactive int64     `json:"inactive"`
}

type Entitlement struct {
	ID                snowflake.ID                 `json:"id"`
	Type              EntitlementType              `json:"type"`
	SKUID             snowflake.ID                 `json:"sku_id"`
	ApplicationID     snowflake.ID                 `json:"application_id"`
	UserID            snowflake.ID                 `json:"user_id"`
	User              PartialUser                  `json:"user,omitzero"`
	GuildID           snowflake.ID                 `json:"guild_id,omitzero"`
	ParentID          snowflake.ID                 `json:"parent_id,omitzero"`
	Deleted           bool                         `json:"deleted"`
	Consumed          bool                         `json:"consumed,omitzero"`
	Branches          []snowflake.ID               `json:"branches,omitempty"`
	StartsAt          *time.Time                   `json:"starts_at"`
	EndsAt            *time.Time                   `json:"ends_at"`
	PromotionID       *snowflake.ID                `json:"promotion_id"`
	SubscriptionID    snowflake.ID                 `json:"subscription_id,omitzero"`
	GiftCodeFlags     GiftCodeFlags                `json:"gift_code_flags"`
	GiftCodeBatchID   snowflake.ID                 `json:"gift_code_batch_id,omitzero"`
	GifterUserID      snowflake.ID                 `json:"gifter_user_id,omitzero"`
	GiftStyle         GiftStyle                    `json:"gift_style,omitzero"`
	FulfillmentStatus EntitlementFulfillmentStatus `json:"fulfillment_status,omitzero"`
	FulfilledAt       time.Time                    `json:"fulfilled_at,omitzero"`
	SourceType        EntitlementSourceType        `json:"source_type,omitzero"`
	TenantMetadata    TenantMetadata               `json:"tenant_metadata,omitzero"`
	SKU               SKU                          `json:"sku,omitzero"`
	SubscriptionPlan  PartialSubscriptionPlan      `json:"subscription_plan,omitzero"`
}

type EntitlementExpirationMetadata struct {
	Extended   bool `json:"extended"`
	Extendable bool `json:"extendable"`
}

type EntitlementListGuildEntitlementsQuery struct {
	WithSKU         bool            `json:"with_sku,omitzero"`
	WithApplication bool            `json:"with_application,omitzero"`
	ExcludeEnded    bool            `json:"exclude_ended,omitzero"`
	ExcludeDeleted  bool            `json:"exclude_deleted,omitzero"`
	EntitlementType EntitlementType `json:"entitlement_type,omitzero"`
}

type Error struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type Eula struct {
	ID      snowflake.ID `json:"id"`
	Name    string       `json:"name"`
	Content string       `json:"content"`
}

type ExchangeHandoffTokenParams struct {
	Key          string `json:"key"`
	HandoffToken string `json:"handoff_token"`
}

type ExchangeHandoffTokenResponse struct {
	Token string      `json:"token"`
	User  PartialUser `json:"user"`
}

type ExchangeProvisionalAccountChildTokenParams struct {
	ChildApplicationID snowflake.ID `json:"child_application_id"`
	ParentAccessToken  string       `json:"parent_access_token"`
}

type ExchangeRemoteAuthTicketParams struct {
	Ticket string `json:"ticket"`
}

type ExchangeRemoteAuthTicketResponse struct {
	EncryptedToken string `json:"encrypted_token"`
}

type ExecuteAutomodAlertActionParams struct {
	ChannelID       snowflake.ID           `json:"channel_id"`
	MessageID       snowflake.ID           `json:"message_id"`
	AlertActionType AutomodAlertActionType `json:"alert_action_type"`
}

type ExecuteServiceWebhookQuery struct {
	Wait     bool         `json:"wait,omitzero"`
	ThreadID snowflake.ID `json:"thread_id,omitzero"`
}

type ExecuteWebhookParams struct {
	Username        string          `json:"username,omitzero"`
	AvatarURL       string          `json:"avatar_url,omitzero"`
	ThreadName      string          `json:"thread_name,omitzero"`
	AppliedTags     []snowflake.ID  `json:"applied_tags,omitempty"`
	Content         string          `json:"content,omitzero"`
	TTS             bool            `json:"tts,omitzero"`
	Embeds          []Embed         `json:"embeds,omitempty"`
	AllowedMentions AllowedMentions `json:"allowed_mentions,omitzero"`
	Components      Components      `json:"components,omitzero"`
	Flags           MessageFlags    `json:"flags,omitzero"`
	FilesN          json.RawMessage `json:"files[n],omitempty"`
	Attachments     []Attachment    `json:"attachments,omitempty"`
	Poll            json.RawMessage `json:"poll,omitempty"`
}

type ExecuteWebhookQuery struct {
	Wait     bool         `json:"wait,omitzero"`
	ThreadID snowflake.ID `json:"thread_id,omitzero"`
}

type ExpectedPrice struct {
	Currency string `json:"currency"`
	Amount   int64  `json:"amount"`
}

type ExperimentBucketOverride struct {
	B int64          `json:"b"`
	K []snowflake.ID `json:"k"`
}

type ExperimentPopulation struct {
	Ranges  []ExperimentPopulationRange `json:"ranges"`
	Filters ExperimentPopulationFilters `json:"filters"`
}

type ExperimentPopulationFilters struct {
	GuildHasFeature       ExperimentPopulationGuildFeatureFilter `json:"guild_has_feature,omitzero"`
	GuildIDRange          ExperimentPopulationRangeFilter        `json:"guild_id_range,omitzero"`
	GuildAgeRangeDays     ExperimentPopulationRangeFilter        `json:"guild_age_range_days,omitzero"`
	GuildMemberCountRange ExperimentPopulationRangeFilter        `json:"guild_member_count_range,omitzero"`
	GuildIDs              ExperimentPopulationIDFilter           `json:"guild_ids,omitzero"`
	GuildHubTypes         ExperimentPopulationHubTypeFilter      `json:"guild_hub_types,omitzero"`
	GuildHasVanityURL     ExperimentPopulationVanityURLFilter    `json:"guild_has_vanity_url,omitzero"`
	GuildInRangeByHash    ExperimentPopulationRangeByHashFilter  `json:"guild_in_range_by_hash,omitzero"`
}

type ExperimentPopulationGuildFeatureFilter struct {
	GuildFeatures []string `json:"guild_features"`
}

type ExperimentPopulationHubTypeFilter struct {
	GuildHubTypes []int64 `json:"guild_hub_types"`
}

type ExperimentPopulationIDFilter struct {
	GuildIDs []snowflake.ID `json:"guild_ids"`
}

type ExperimentPopulationRange struct {
	Bucket  int64                         `json:"bucket"`
	Rollout []ExperimentPopulationRollout `json:"rollout"`
}

type ExperimentPopulationRangeByHashFilter struct {
	HashKey int64 `json:"hash_key"`
	Target  int64 `json:"target"`
}

type ExperimentPopulationRangeFilter struct {
	MinID *snowflake.ID `json:"min_id"`
	MaxID *snowflake.ID `json:"max_id"`
}

type ExperimentPopulationRollout struct {
	S int64 `json:"s"`
	E int64 `json:"e"`
}

type ExperimentPopulationVanityURLFilter struct {
	GuildHasVanityURL bool `json:"guild_has_vanity_url"`
}

type ExperimentSessionOverrideCreate struct {
	ExperimentName string `json:"experiment_name"`
	VariantID      int64  `json:"variant_id"`
}

type ExperimentSessionOverrideDelete struct {
	ExperimentName string `json:"experiment_name"`
}

type ExplicitContentSettings struct {
	ExplicitContentGuilds      json.RawMessage `json:"explicit_content_guilds"`
	ExplicitContentFriendDM    json.RawMessage `json:"explicit_content_friend_dm"`
	ExplicitContentNonFriendDM json.RawMessage `json:"explicit_content_non_friend_dm"`
}

type ExternalAsset struct {
	URL               string `json:"url"`
	ExternalAssetPath string `json:"external_asset_path"`
}

type ExternalLinkElement struct {
	URL             string `json:"url"`
	LinkText        string `json:"link_text"`
	LinkDescription string `json:"link_description,omitzero"`
	IsLocalized     bool   `json:"is_localized"`
}

type ExternalSKUStrategy struct {
	Type     ExternalSKUStrategyType `json:"type"`
	Metadata map[string]string       `json:"metadata,omitempty"`
}

type FamilyCenter struct {
	LinkedUsers  []LinkedUser  `json:"linked_users"`
	TeenAuditLog TeenAuditLog  `json:"teen_audit_log"`
	Users        []PartialUser `json:"users"`
}

type FavoriteChannel struct {
	Nickname    string          `json:"nickname"`
	Type        json.RawMessage `json:"type"`
	Position    json.RawMessage `json:"position"`
	ParentID    json.RawMessage `json:"parent_id"`
	ChannelType ChannelType     `json:"channel_type,omitzero"`
	Collapsed   bool            `json:"collapsed"`
}

type FavoriteEmojis struct {
	Emojis []string `json:"emojis"`
}

type FavoriteGIF struct {
	Format json.RawMessage `json:"format"`
	Src    string          `json:"src"`
	Width  json.RawMessage `json:"width"`
	Height json.RawMessage `json:"height"`
	Order  json.RawMessage `json:"order"`
}

type FavoriteGifs struct {
	Gifs        map[string]FavoriteGIF `json:"gifs"`
	HideTooltip bool                   `json:"hide_tooltip"`
}

type FavoriteSoundboardSounds struct {
	SoundIDs []json.RawMessage `json:"sound_ids"`
}

type FavoriteStickers struct {
	StickerIDs []json.RawMessage `json:"sticker_ids"`
}

type Favorites struct {
	FavoriteChannels map[string]FavoriteChannel `json:"favorite_channels"`
	Muted            bool                       `json:"muted"`
	GuildVisible     bool                       `json:"guild_visible,omitzero"`
}

type FeaturedShopBlock struct {
	Subblocks json.RawMessage `json:"subblocks"`
	Type      ShopBlockType   `json:"type"`
}

type FeedShopBlock struct {
	RankedSKUIDs []snowflake.ID `json:"ranked_sku_ids"`
	SortedSKUIDs SortedSKUIDs   `json:"sorted_sku_ids"`
	Type         ShopBlockType  `json:"type"`
}

type File struct {
	Type    ComponentType   `json:"type"`
	ID      int64           `json:"id,omitzero"`
	File    json.RawMessage `json:"file"`
	Spoiler bool            `json:"spoiler,omitzero"`
	Name    string          `json:"name"`
	Size    int64           `json:"size"`
}

type FileUpload struct {
	Type      ComponentType     `json:"type"`
	ID        int64             `json:"id,omitzero"`
	CustomID  string            `json:"custom_id"`
	MinValues int64             `json:"min_values,omitzero"`
	MaxValues int64             `json:"max_values,omitzero"`
	Required  bool              `json:"required,omitzero"`
	FileTypes []json.RawMessage `json:"file_types,omitempty"`
}

type FinishOAuth2DeviceFlowParams struct {
	UserCode string                 `json:"user_code"`
	Result   OAuth2DeviceFlowStatus `json:"result"`
}

type FinishPasswordlessLoginParams struct {
	Ticket        string       `json:"ticket"`
	Credential    string       `json:"credential"`
	LoginSource   *LoginSource `json:"login_source,omitempty"`
	GiftCodeSKUID *string      `json:"gift_code_sku_id,omitempty"`
}

type FinishPasswordlessLoginResponse struct {
	UserID          snowflake.ID  `json:"user_id"`
	Token           string        `json:"token"`
	UserSettings    LoginSettings `json:"user_settings"`
	RequiredActions []string      `json:"required_actions,omitempty"`
}

type FinishRemoteAuthParams struct {
	HandshakeToken string `json:"handshake_token"`
	TemporaryToken bool   `json:"temporary_token,omitzero"`
}

type FixGrammarParams struct {
	Content string `json:"content"`
}

type FixGrammarResponse struct {
	Content string `json:"content"`
}

type FlaggedContent struct {
	Type        string       `json:"type"`
	ID          snowflake.ID `json:"id"`
	Content     string       `json:"content"`
	Attachments []Attachment `json:"attachments"`
}

type FollowChannelParams struct {
	WebhookChannelID snowflake.ID `json:"webhook_channel_id"`
}

type FollowedChannel struct {
	ChannelID snowflake.ID `json:"channel_id"`
	WebhookID snowflake.ID `json:"webhook_id"`
}

type ForLaterSettings struct {
	CurrentTab json.RawMessage `json:"current_tab"`
}

type ForgotPasswordParams struct {
	Login string `json:"login"`
}

type ForgotPasswordResponse struct {
	Method PasswordResetMethod `json:"method"`
}

type ForumTag struct {
	ID        snowflake.ID  `json:"id"`
	Name      string        `json:"name"`
	Moderated bool          `json:"moderated"`
	EmojiID   *snowflake.ID `json:"emoji_id"`
	EmojiName *string       `json:"emoji_name"`
}

type ForumThreadCreated struct {
	ChannelID         snowflake.ID `json:"channel_id"`
	ChannelName       string       `json:"channel_name"`
	ParentID          snowflake.ID `json:"parent_id"`
	ParentName        string       `json:"parent_name"`
	UserID            snowflake.ID `json:"user_id"`
	UserUsername      string       `json:"user_username"`
	UserAvatar        *string      `json:"user_avatar,omitempty"`
	UserGuildAvatar   *string      `json:"user_guild_avatar,omitempty"`
	UserDiscriminator *int64       `json:"user_discriminator,omitempty"`
	GuildID           snowflake.ID `json:"guild_id,omitzero"`
	GuildName         string       `json:"guild_name,omitzero"`
	GuildIcon         *string      `json:"guild_icon,omitempty"`
}

type FramesBannerShopBlock struct {
	Type                   ShopBlockType `json:"type"`
	Title                  string        `json:"title"`
	Body                   string        `json:"body"`
	MobileBackgroundImage  string        `json:"mobile_background_image,omitzero"`
	MobileForegroundImage  string        `json:"mobile_foreground_image,omitzero"`
	DesktopBackgroundImage string        `json:"desktop_background_image,omitzero"`
}

type FramesProductShelfShopBlock struct {
	Type                   ShopBlockType  `json:"type"`
	Title                  string         `json:"title"`
	CategorySKUID          snowflake.ID   `json:"category_sku_id"`
	CategoryStoreListingID snowflake.ID   `json:"category_store_listing_id"`
	RankedSKUIDs           []snowflake.ID `json:"ranked_sku_ids"`
	BackgroundImage        string         `json:"background_image,omitzero"`
	DesktopBackgroundImage string         `json:"desktop_background_image,omitzero"`
	MobileBackgroundImage  string         `json:"mobile_background_image,omitzero"`
	ButtonText             string         `json:"button_text,omitzero"`
}

type FrecencyItem struct {
	TotalUses  json.RawMessage   `json:"total_uses"`
	RecentUses []json.RawMessage `json:"recent_uses"`
	Frecency   json.RawMessage   `json:"frecency"`
	Score      json.RawMessage   `json:"score"`
}

type FrecencyUserSettings struct {
	Versions                   Versions                   `json:"versions"`
	FavoriteGifs               FavoriteGifs               `json:"favorite_gifs"`
	FavoriteStickers           FavoriteStickers           `json:"favorite_stickers"`
	StickerFrecency            StickerFrecency            `json:"sticker_frecency"`
	FavoriteEmojis             FavoriteEmojis             `json:"favorite_emojis"`
	EmojiFrecency              EmojiFrecency              `json:"emoji_frecency"`
	ApplicationCommandFrecency ApplicationCommandFrecency `json:"application_command_frecency"`
	FavoriteSoundboardSounds   FavoriteSoundboardSounds   `json:"favorite_soundboard_sounds"`
	ApplicationFrecency        ApplicationFrecency        `json:"application_frecency"`
	HeardSoundFrecency         HeardSoundFrecency         `json:"heard_sound_frecency"`
	PlayedSoundFrecency        PlayedSoundFrecency        `json:"played_sound_frecency"`
	GuildAndChannelFrecency    GuildAndChannelFrecency    `json:"guild_and_channel_frecency"`
	EmojiReactionFrecency      EmojiReactionFrecency      `json:"emoji_reaction_frecency"`
}

type FreeTextElement struct {
	Title          string `json:"title,omitzero"`
	Subtitle       string `json:"subtitle,omitzero"`
	Placeholder    string `json:"placeholder,omitzero"`
	Rows           int64  `json:"rows"`
	CharacterLimit int64  `json:"character_limit"`
	Pattern        string `json:"pattern,omitzero"`
}

type FriendListEntry struct {
	FriendID string `json:"friend_id"`
}

type FriendSourceFlags struct {
	All           bool `json:"all,omitzero"`
	MutualFriends bool `json:"mutual_friends,omitzero"`
	MutualGuilds  bool `json:"mutual_guilds,omitzero"`
}

type FriendSuggestion struct {
	SuggestedUser             PartialUser              `json:"suggested_user"`
	Reasons                   []FriendSuggestionReason `json:"reasons"`
	FromSuggestedUserContacts bool                     `json:"from_suggested_user_contacts,omitzero"`
	MutualFriendsCount        int64                    `json:"mutual_friends_count,omitzero"`
	ContactNames              []string                 `json:"contact_names,omitempty"`
}

type FriendSuggestionCreate struct {
	UserID               snowflake.ID   `json:"user_id"`
	UserUsername         string         `json:"user_username"`
	UserAvatar           *string        `json:"user_avatar,omitempty"`
	PlatformType         ConnectionType `json:"platform_type"`
	PlatformName         *string        `json:"platform_name,omitempty"`
	PlatformUserUsername string         `json:"platform_user_username"`
}

type FriendSuggestionDelete struct {
	SuggestedUserID snowflake.ID `json:"suggested_user_id"`
}

type FriendSuggestionReason struct {
	Type     FriendSuggestionReasonType `json:"type"`
	Platform ConnectionType             `json:"platform"`
	Name     string                     `json:"name"`
}

type GIF struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	URL     string `json:"url"`
	Src     string `json:"src"`
	GIFSrc  string `json:"gif_src"`
	Preview string `json:"preview"`
	Width   int64  `json:"width"`
	Height  int64  `json:"height"`
}

type GIFCategory struct {
	Name string `json:"name"`
	Src  string `json:"src"`
}

type Game struct {
	SupplementalGameData GameData      `json:"supplemental_game_data,omitzero"`
	Genres               []int64       `json:"genres,omitempty"`
	Platforms            []int64       `json:"platforms,omitempty"`
	Websites             []GameWebsite `json:"websites,omitempty"`
	Companies            []Company     `json:"companies"`
	ScreenshotHashes     []string      `json:"screenshot_hashes,omitempty"`
	ScreenshotURLs       []string      `json:"screenshot_urls,omitempty"`
	Trailers             []StoreAsset  `json:"trailers,omitempty"`
	L30Rank              int64         `json:"l30_rank,omitzero"`
}

type GameActivity struct {
	GameApplicationID snowflake.ID `json:"game_application_id"`
	ActivityLevel     int64        `json:"activity_level"`
	ActivityScore     int64        `json:"activity_score"`
}

type GameClaim struct {
	GameID         snowflake.ID    `json:"game_id"`
	ApplicationID  snowflake.ID    `json:"application_id"`
	ClaimStatus    GameClaimStatus `json:"claim_status"`
	ThirdPartyURLs []ThirdPartyURL `json:"third_party_urls,omitempty"`
	UpdatedAt      time.Time       `json:"updated_at"`
	InviteProof    string          `json:"invite_proof"`
}

type GameData struct {
	ApplicationID          snowflake.ID       `json:"application_id"`
	IgdbGameID             *string            `json:"igdb_game_id"`
	Name                   string             `json:"name"`
	Summary                *string            `json:"summary,omitempty"`
	SummaryLocalized       *string            `json:"summary_localized,omitempty"`
	Websites               []GameWebsite      `json:"websites,omitempty"`
	Themes                 []int64            `json:"themes,omitempty"`
	Genres                 []int64            `json:"genres,omitempty"`
	Platforms              []int64            `json:"platforms,omitempty"`
	ArtworkURLs            []string           `json:"artwork_urls,omitempty"`
	ScreenshotURLs         []string           `json:"screenshot_urls,omitempty"`
	IconHash               *string            `json:"icon_hash,omitempty"`
	CoverImageURL          *string            `json:"cover_image_url,omitempty"`
	FirstReleaseDate       *time.Time         `json:"first_release_date,omitempty"`
	PublisherNames         []string           `json:"publisher_names,omitempty"`
	DeveloperNames         []string           `json:"developer_names,omitempty"`
	Trailers               []StoreAsset       `json:"trailers,omitempty"`
	ShopCollectionIDs      []snowflake.ID     `json:"shop_collection_ids,omitempty"`
	SteamReleaseStatus     SteamReleaseStatus `json:"steam_release_status,omitzero"`
	Reviews                GameDataReviews    `json:"reviews,omitzero"`
	OpencriticURL          string             `json:"opencritic_url,omitzero"`
	SteamID                string             `json:"steam_id,omitzero"`
	AnnouncementsChannelID snowflake.ID       `json:"announcements_channel_id,omitzero"`
	L30Rank                int64              `json:"l30_rank"`
	GameFlags              GameFlags          `json:"game_flags,omitzero"`
}

type GameDataOpencriticReviews struct {
	TopCriticRating      *int64          `json:"top_critic_rating"`
	TopCriticRatingCount *int64          `json:"top_critic_rating_count"`
	Tier                 *OpencriticTier `json:"tier"`
}

type GameDataOverrides struct {
	Description       *string        `json:"description,omitempty"`
	Genres            []int64        `json:"genres,omitempty"`
	Platforms         []int64        `json:"platforms,omitempty"`
	CoverHash         *string        `json:"cover_hash,omitempty"`
	ScreenshotHashes  []string       `json:"screenshot_hashes,omitempty"`
	TrailerAssetIDs   []snowflake.ID `json:"trailer_asset_ids,omitempty"`
	Websites          []GameWebsite  `json:"websites,omitempty"`
	Companies         []Company      `json:"companies"`
	ShopCollectionIDs []snowflake.ID `json:"shop_collection_ids,omitempty"`
	BannerHash        *string        `json:"banner_hash,omitempty"`
	IconHash          *string        `json:"icon_hash,omitempty"`
	GameFlags         GameFlags      `json:"game_flags,omitzero"`
}

type GameDataReviews struct {
	Steam      GameDataSteamReviews      `json:"steam,omitzero"`
	Opencritic GameDataOpencriticReviews `json:"opencritic,omitzero"`
}

type GameDataSteamReviews struct {
	Rating               float64 `json:"rating,omitzero"`
	RatingCount          int64   `json:"rating_count,omitzero"`
	RecentRating         float64 `json:"recent_rating,omitzero"`
	RecentRatingCount    int64   `json:"recent_rating_count,omitzero"`
	LocalizedRating      float64 `json:"localized_rating,omitzero"`
	LocalizedRatingCount int64   `json:"localized_rating_count,omitzero"`
}

type GameInvite struct {
	InviteID         snowflake.ID   `json:"invite_id"`
	CreatedAt        time.Time      `json:"created_at"`
	TTL              int64          `json:"ttl"`
	InviterID        snowflake.ID   `json:"inviter_id"`
	RecipientID      snowflake.ID   `json:"recipient_id"`
	PlatformType     ConnectionType `json:"platform_type"`
	LaunchParameters string         `json:"launch_parameters"`
	Installed        bool           `json:"installed,omitzero"`
	Joinable         bool           `json:"joinable,omitzero"`
	FallbackURL      *string        `json:"fallback_url"`
	ApplicationAsset string         `json:"application_asset"`
	ApplicationName  string         `json:"application_name"`
}

type GameInviteDelete struct {
	InviteID snowflake.ID `json:"invite_id"`
}

type GameInviteDeleteMany struct {
	InviteIDs []snowflake.ID `json:"invite_ids"`
}

type GameLibrarySettings struct {
	InstallShortcutDesktop   bool `json:"install_shortcut_desktop,omitzero"`
	InstallShortcutStartMenu bool `json:"install_shortcut_start_menu,omitzero"`
	DisableGamesTab          bool `json:"disable_games_tab,omitzero"`
}

type GameRelationship struct {
	ID            string       `json:"id"`
	ApplicationID snowflake.ID `json:"application_id"`
	Type          int64        `json:"type"`
	User          PartialUser  `json:"user"`
	Since         time.Time    `json:"since"`
	DMAccessType  int64        `json:"dm_access_type"`
	UserID        snowflake.ID `json:"user_id"`
}

type GameRelationshipRemove struct {
	ID            string       `json:"id"`
	ApplicationID snowflake.ID `json:"application_id"`
	Type          int64        `json:"type"`
	Since         time.Time    `json:"since"`
	DMAccessType  int64        `json:"dm_access_type"`
	UserID        snowflake.ID `json:"user_id"`
}

type GameServer struct {
	ID              snowflake.ID           `json:"id"`
	Name            string                 `json:"name"`
	RegionID        *string                `json:"region_id,omitempty"`
	RegionName      *string                `json:"region_name,omitempty"`
	SKUID           snowflake.ID           `json:"sku_id"`
	PlanName        string                 `json:"plan_name"`
	PlayersCount    *int64                 `json:"players_count,omitempty"`
	MaxPlayersCount *int64                 `json:"max_players_count,omitempty"`
	IP              string                 `json:"ip,omitzero"`
	Port            string                 `json:"port,omitzero"`
	EntitlementID   snowflake.ID           `json:"entitlement_id"`
	ProviderType    GameServerProviderType `json:"provider_type"`
	ProviderURL     *string                `json:"provider_url,omitempty"`
	Status          int64                  `json:"status"`
	GameID          snowflake.ID           `json:"game_id"`
	GameConfig      GameServerConfig       `json:"game_config,omitzero"`
}

type GameServerConfig struct {
	Type    string `json:"type"`
	Version string `json:"version"`
}

type GameServerCreate struct {
	GuildID    snowflake.ID `json:"guild_id"`
	GameServer GameServer   `json:"game_server"`
}

type GameServerDelete struct {
	GuildID      snowflake.ID `json:"guild_id"`
	GameServerID snowflake.ID `json:"game_server_id"`
}

type GameServerInstructions struct {
	PC []string `json:"pc"`
}

type GameServerPowerupMetadata struct {
	BoostPrice         int64                    `json:"boost_price"`
	PurchaseLimit      int64                    `json:"purchase_limit"`
	GuildFeatures      GuildPremiumFeatures     `json:"guild_features"`
	CategoryType       GuildPowerupCategoryType `json:"category_type"`
	AvailableProviders []string                 `json:"available_providers"`
	Memory             int64                    `json:"memory"`
	Cpu                int64                    `json:"cpu"`
	Storage            int64                    `json:"storage"`
	MaxSlots           int64                    `json:"max_slots"`
	MemoryString       string                   `json:"memory_string"`
	PlayerString       string                   `json:"player_string"`
}

type GameServerPowerupProductMetadata struct {
	Instructions                   json.RawMessage        `json:"instructions"`
	DeactivationCooldownPeriodDays int64                  `json:"deactivation_cooldown_period_days"`
	GameApplicationID              snowflake.ID           `json:"game_application_id"`
	Provider                       GameServerProviderType `json:"provider"`
	Disabled                       bool                   `json:"disabled"`
	EarlyAccess                    bool                   `json:"early_access"`
	CanMarket                      bool                   `json:"can_market"`
}

type GameServerRegion struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	CountryCode string `json:"country_code"`
	PingURL     string `json:"ping_url"`
}

type GameServerUpdate struct {
	GuildID    snowflake.ID `json:"guild_id"`
	GameServer GameServer   `json:"game_server"`
}

type GameWebsite struct {
	URL      string              `json:"url"`
	Category GameWebsiteCategory `json:"category"`
}

type GameWidget struct {
	Data      GameWidgetData `json:"data"`
	ID        snowflake.ID   `json:"id"`
	UpdatedAt time.Time      `json:"updated_at"`
}

type GameWidgetData struct {
	Type          GameWidgetType   `json:"type"`
	Games         []GameWidgetGame `json:"games,omitempty"`
	ApplicationID snowflake.ID     `json:"application_id,omitzero"`
}

type GameWidgetGame struct {
	GameID  snowflake.ID `json:"game_id"`
	Comment *string      `json:"comment,omitempty"`
	Tags    []string     `json:"tags,omitempty"`
}

type GatewayApplication struct {
	ID       snowflake.ID     `json:"id"`
	Flags    ApplicationFlags `json:"flags"`
	FlagsNew string           `json:"flags_new"`
	Name     string           `json:"name,omitzero"`
	ParentID snowflake.ID     `json:"parent_id,omitzero"`
}

type GatewayAuth struct {
	AuthenticatorTypes []int64 `json:"authenticator_types,omitempty"`
}

type GatewayCheckoutContext struct {
	BraintreeDeviceData *string `json:"braintree_device_data,omitempty"`
}

type GatewayEventsMessageDelete struct {
	ID        snowflake.ID `json:"id"`
	ChannelID snowflake.ID `json:"channel_id"`
	GuildID   snowflake.ID `json:"guild_id,omitzero"`
}

type GatewayEventsQoSHeartbeat struct {
	InitializationTimestamp int64  `json:"initialization_timestamp"`
	SessionID               string `json:"session_id"`
	ClientLaunchID          string `json:"client_launch_id"`
}

type GatewayEventsUserApplicationRemove struct {
	ApplicationID snowflake.ID `json:"application_id"`
}

type GatewayFeatureFlags struct {
	DisabledFunctions     []string `json:"disabled_functions"`
	DisabledGatewayEvents []string `json:"disabled_gateway_events"`
}

type GatewayGiftCode struct {
	Code      string       `json:"code"`
	SKUID     snowflake.ID `json:"sku_id"`
	Uses      int64        `json:"uses"`
	ChannelID snowflake.ID `json:"channel_id,omitzero"`
	GuildID   snowflake.ID `json:"guild_id,omitzero"`
}

type GatewayGuild struct {
	JoinedAt                 time.Time                  `json:"joined_at"`
	Large                    bool                       `json:"large"`
	Unavailable              bool                       `json:"unavailable,omitzero"`
	GeoRestricted            bool                       `json:"geo_restricted,omitzero"`
	MemberCount              int64                      `json:"member_count"`
	Members                  []GuildMember              `json:"members"`
	Channels                 []Channel                  `json:"channels"`
	Threads                  []Channel                  `json:"threads"`
	Presences                []Presence                 `json:"presences"`
	VoiceStates              []VoiceState               `json:"voice_states"`
	ActivityInstances        []EmbeddedActivityInstance `json:"activity_instances"`
	StageInstances           []StageInstance            `json:"stage_instances"`
	GuildScheduledEvents     []GuildScheduledEvent      `json:"guild_scheduled_events"`
	DataMode                 DataMode                   `json:"data_mode"`
	PartialUpdates           PartialGuildUpdates        `json:"partial_updates,omitzero"`
	ChannelUpdates           []ChannelUnread            `json:"channel_updates,omitempty"`
	UnableToSyncDeletes      bool                       `json:"unable_to_sync_deletes,omitzero"`
	Properties               PartialGuild               `json:"properties"`
	Stickers                 []Sticker                  `json:"stickers"`
	Roles                    []Role                     `json:"roles"`
	Emojis                   []Emoji                    `json:"emojis"`
	SoundboardSounds         []SoundboardSound          `json:"soundboard_sounds"`
	PremiumSubscriptionCount int64                      `json:"premium_subscription_count"`
	ID                       snowflake.ID               `json:"id"`
}

type GatewayLobbyExtraFields struct {
	VoiceStates []json.RawMessage `json:"voice_states"`
	Region      string            `json:"region"`
	Metadata    map[string]string `json:"metadata"`
}

type GatewayPayload struct {
	Op int64           `json:"op"`
	D  json.RawMessage `json:"d"`
	S  *int64          `json:"s,omitempty"`
	T  *string         `json:"t,omitempty"`
}

type GenericPushNotification struct {
	NotificationType     string        `json:"notification_type,omitzero"`
	Deeplink             string        `json:"deeplink,omitzero"`
	NotificationID       string        `json:"notification_id,omitzero"`
	NotificationCenterID snowflake.ID  `json:"notification_center_id,omitzero"`
	GuildID              snowflake.ID  `json:"guild_id,omitzero"`
	ChannelID            snowflake.ID  `json:"channel_id,omitzero"`
	MessageID            snowflake.ID  `json:"message_id,omitzero"`
	UserID               snowflake.ID  `json:"user_id,omitzero"`
	StatusText           string        `json:"status_text,omitzero"`
	StatusEmojiID        *snowflake.ID `json:"status_emoji_id,omitempty"`
	StatusEmojiName      *string       `json:"status_emoji_name,omitempty"`
	StatusEmojiAnimated  bool          `json:"status_emoji_animated,omitzero"`
}

type GetActivityInstanceConnectedParticipantsResponse struct {
	Participants []RPCActivityParticipant `json:"participants"`
}

type GetActivityMetadataQuery struct {
	ChannelID snowflake.ID `json:"channel_id,omitzero"`
	MessageID snowflake.ID `json:"message_id,omitzero"`
}

type GetActivityMetadataResponse struct {
	Secret string `json:"secret"`
}

type GetAdminCommunityEligibilityResponse struct {
	EligibleForAdminServer bool `json:"eligible_for_admin_server"`
}

type GetApplicationDirectoryApplicationEmbedQuery struct {
	WithLocalizations bool `json:"with_localizations,omitzero"`
}

type GetApplicationDirectoryApplicationQuery struct {
	Locale            string `json:"locale,omitzero"`
	Nocache           bool   `json:"nocache,omitzero"`
	WithLocalizations bool   `json:"with_localizations,omitzero"`
}

type GetApplicationDiscoverabilityStateResponse struct {
	DiscoverabilityState      ApplicationDiscoverabilityState      `json:"discoverability_state"`
	DiscoveryEligibilityFlags ApplicationDiscoveryEligibilityFlags `json:"discovery_eligibility_flags"`
	BadCommands               []ApplicationCommand                 `json:"bad_commands"`
}

type GetApplicationPrimaryStoreListingQuery struct {
	CountryCode string `json:"country_code,omitzero"`
	Localize    bool   `json:"localize,omitzero"`
}

type GetApplicationSKUStorefrontResponse struct {
	StoreListing       StoreListing       `json:"store_listing"`
	Assets             []StoreAsset       `json:"assets"`
	StorefrontMetadata StorefrontMetadata `json:"storefront_metadata"`
	StorefrontPricing  StorefrontPrices   `json:"storefront_pricing"`
}

type GetApplicationStoreLayoutResponse struct {
	Subscriptions     []StoreListing     `json:"subscriptions"`
	Otps              []StoreListing     `json:"otps"`
	SubscriptionPlans []SubscriptionPlan `json:"subscription_plans"`
}

type GetApplicationStorefrontForPremiumButtonResponse struct {
	SKUs              []SKU              `json:"skus,omitempty"`
	StoreListings     []StoreListing     `json:"store_listings"`
	SubscriptionPlans []SubscriptionPlan `json:"subscription_plans,omitempty"`
}

type GetApplicationTicketResponse struct {
	Ticket string `json:"ticket"`
}

type GetApplicationUndeletableReasonResponse struct {
	Deletable *bool                         `json:"deletable,omitempty"`
	Reason    *ApplicationUndeletableReason `json:"reason,omitempty"`
}

type GetApplicationUpdatesQuery struct {
	Platform DesktopPlatformType `json:"platform,omitzero"`
}

type GetApplicationUpdatesResponse struct {
	Name    string    `json:"name"`
	PubDate time.Time `json:"pub_date"`
	URL     string    `json:"url,omitzero"`
	Notes   string    `json:"notes,omitzero"`
}

type GetAuthSessionsResponse struct {
	UserSessions []AuthSession `json:"user_sessions"`
}

type GetBackupCodesParams struct {
	Key        string `json:"key"`
	Nonce      string `json:"nonce"`
	Regenerate bool   `json:"regenerate"`
}

type GetBackupCodesResponse struct {
	BackupCodes []BackupCode `json:"backup_codes"`
}

type GetBillingCountryCodeResponse struct {
	CountryCode string `json:"country_code"`
}

type GetBillingLocationInfoResponse struct {
	CountryCode     string  `json:"country_code"`
	SubdivisionCode *string `json:"subdivision_code"`
}

type GetCallEligibilityResponse struct {
	Ringable bool `json:"ringable"`
}

type GetChannelArguments struct {
	ChannelID snowflake.ID `json:"channel_id"`
}

type GetChannelLinkedAccountsQuery struct {
	UserIDs []snowflake.ID `json:"user_ids,omitempty"`
}

type GetChannelLinkedAccountsResponse struct {
	LinkedAccounts map[snowflake.ID][]LinkedAccount `json:"linked_accounts"`
}

type GetChannelMediaPreviewResponse struct {
	Media MediaPreview `json:"media"`
}

type GetChannelPermissionsResponse struct {
	Permissions Permissions `json:"permissions"`
}

type GetChannelPostDataParams struct {
	ThreadIDs []snowflake.ID `json:"thread_ids"`
}

type GetChannelPostDataResponse struct {
	Threads json.RawMessage `json:"threads"`
}

type GetChannelsArguments struct {
	GuildID snowflake.ID `json:"guild_id"`
}

type GetChannelsResponse struct {
	Channels []PartialRPCChannel `json:"channels"`
}

type GetCheckoutRecoveryResponse struct {
	IsEligible bool `json:"is_eligible,omitzero"`
}

type GetCheckpointResponse struct {
	CardID               CheckpointCardType              `json:"card_id,omitzero"`
	AvatarDecoration     AvatarDecorationCollectibleItem `json:"avatar_decoration"`
	PowerLevel           float64                         `json:"power_level,omitzero"`
	PowerLevelPercentile float64                         `json:"power_level_percentile,omitzero"`
	Messages             CheckpointMessageStatistics     `json:"messages,omitzero"`
	Emojis               CheckpointEmojiStatistics       `json:"emojis,omitzero"`
	Voice                CheckpointVoiceStatistics       `json:"voice,omitzero"`
	Guilds               CheckpointGuildsStatistics      `json:"guilds,omitzero"`
	Sidekick             CheckpointSidekick              `json:"sidekick,omitzero"`
	Users                []CheckpointUser                `json:"users,omitempty"`
	Applications         CheckpointGamesStatistics       `json:"applications,omitzero"`
	Quests               CheckpointQuestStatistics       `json:"quests,omitzero"`
}

type GetChurnUserOfferResponse struct {
	Offer UserDiscountOffer `json:"offer"`
}

type GetCollectiblesMarketingQuery struct {
	Platform CollectiblesMarketingPlatformType `json:"platform"`
	Release  CollectiblesMarketingReleaseType  `json:"release,omitzero"`
}

type GetCollectiblesMarketingResponse struct {
	Marketings map[int64]json.RawMessage `json:"marketings"`
}

type GetCollectiblesProductQuery struct {
	CountryCode         string              `json:"country_code,omitzero"`
	IncludeBundles      bool                `json:"include_bundles,omitzero"`
	VariantsReturnStyle VariantsReturnStyle `json:"variants_return_style,omitzero"`
}

type GetCollectiblesShopQuery struct {
	CountryCode          string                  `json:"country_code,omitzero"`
	Tab                  CollectiblesShopHomeTab `json:"tab,omitzero"`
	ShopHomeConfig       string                  `json:"shop_home_config,omitzero"`
	IncludeBundles       bool                    `json:"include_bundles,omitzero"`
	IncludeDynamicBlocks bool                    `json:"include_dynamic_blocks,omitzero"`
	IncludeUnpublished   bool                    `json:"include_unpublished,omitzero"`
	NoCache              bool                    `json:"no_cache,omitzero"`
	PaymentGateway       PaymentGateway          `json:"payment_gateway,omitzero"`
	SkipNumCategories    int64                   `json:"skip_num_categories,omitzero"`
	VariantsReturnStyle  VariantsReturnStyle     `json:"variants_return_style,omitzero"`
}

type GetCollectiblesShopResponse struct {
	Categories    []CollectibleCategory     `json:"categories"`
	ShopBlocks    []json.RawMessage         `json:"shop_blocks"`
	UserDiscounts []CollectibleUserDiscount `json:"user_discounts,omitempty"`
}

type GetCollectiblesShopTabLayoutResponse struct {
	LayoutID string `json:"layout_id"`
}

type GetConfettiConsumableResponse struct {
	Entitlement *Entitlement `json:"entitlement"`
	NumPotions  int64        `json:"num_potions"`
}

type GetConsumableSKUPricingResponse struct {
	Price SKUPrice `json:"price"`
}

type GetCreatorMonetizationEligibilityResponse struct {
	GuildID                snowflake.ID                 `json:"guild_id"`
	CanApply               bool                         `json:"can_apply,omitzero"`
	Sufficient             bool                         `json:"sufficient,omitzero"`
	Country                bool                         `json:"country,omitzero"`
	MFA                    bool                         `json:"mfa,omitzero"`
	Size                   bool                         `json:"size,omitzero"`
	SafeEnvironment        bool                         `json:"safe_environment,omitzero"`
	HealthScore            *DiscoveryHealthScore        `json:"health_score,omitempty"`
	HealthScorePending     bool                         `json:"health_score_pending,omitzero"`
	NSFWProperties         DiscoveryNSFWProperties      `json:"nsfw_properties,omitzero"`
	RetentionHealthy       bool                         `json:"retention_healthy,omitzero"`
	EngagementHealthy      bool                         `json:"engagement_healthy,omitzero"`
	Age                    bool                         `json:"age,omitzero"`
	OwnerAge               bool                         `json:"owner_age,omitzero"`
	MinimumAgeInDays       *int64                       `json:"minimum_age_in_days,omitempty"`
	MinimumOwnerAgeInYears *int64                       `json:"minimum_owner_age_in_years,omitempty"`
	MinimumSize            *int64                       `json:"minimum_size,omitempty"`
	LatestRequest          CreatorMonetizationRequest   `json:"latest_request,omitzero"`
	Rejection              CreatorMonetizationRejection `json:"rejection,omitzero"`
}

type GetCreatorMonetizationMarketingOnboardingResponse struct {
	Sections []CreatorMonetizationMarketingOnboardingSection `json:"sections"`
}

type GetCreatorMonetizationRestrictionsResponse struct {
	Restrictions []int64 `json:"restrictions"`
}

type GetCurrentAuthorizationInformationResponse struct {
	Application PartialApplication `json:"application"`
	Scopes      []string           `json:"scopes"`
	Expires     time.Time          `json:"expires"`
	User        PartialUser        `json:"user,omitzero"`
}

type GetCurrentUserQuery struct {
	WithAnalyticsToken bool `json:"with_analytics_token,omitzero"`
}

type GetDeletedEntityIDsNotMatchingHash struct {
	GuildID        snowflake.ID `json:"guild_id"`
	ChannelIDsHash string       `json:"channel_ids_hash"`
	RoleIDsHash    string       `json:"role_ids_hash"`
	EmojiIDsHash   string       `json:"emoji_ids_hash"`
	StickerIDsHash string       `json:"sticker_ids_hash"`
}

type GetDeveloperApplicationWidgetConfigsResponse struct {
	Configs map[snowflake.ID][]WidgetConfig `json:"configs"`
}

type GetDeviceSyncTokenResponse struct {
	Token string `json:"token"`
}

type GetDirectoryBroadcastInfoResponse struct {
	CanBroadcast bool `json:"can_broadcast"`
	HasBroadcast bool `json:"has_broadcast,omitzero"`
}

type GetDiscoverySlugResponse struct {
	Slug      string                `json:"slug"`
	Guild     DiscoverableGuild     `json:"guild,omitzero"`
	StorePage MonetizationStorePage `json:"store_page,omitzero"`
}

type GetEarnedQuestPlacementQuery struct {
	QuestIDs                 []string         `json:"quest_ids"`
	Content                  QuestContentType `json:"content"`
	ClientHeartbeatSessionID string           `json:"client_heartbeat_session_id,omitzero"`
}

type GetEarnedQuestPlacementResponse struct {
	Quests             map[snowflake.ID]*Quest `json:"quests"`
	ResponseTTLSeconds int64                   `json:"response_ttl_seconds,omitzero"`
	RequestID          *string                 `json:"request_id,omitempty"`
}

type GetEmojiSourceResponse struct {
	Type        EmojiSourceType   `json:"type"`
	Guild       *EmojiGuild       `json:"guild"`
	Application *EmojiApplication `json:"application"`
}

type GetEntitlementTicketResponse struct {
	Ticket string `json:"ticket"`
}

type GetEntitlementsEmbeddedResponse struct {
	Entitlements []Entitlement `json:"entitlements"`
}

type GetExperimentAssignmentsResponse struct {
	Fingerprint      string            `json:"fingerprint,omitzero"`
	Assignments      []json.RawMessage `json:"assignments"`
	GuildExperiments []GuildExperiment `json:"guild_experiments,omitempty"`
}

type GetGameQuery struct {
	WithSupplementalData bool `json:"with_supplemental_data,omitzero"`
}

type GetGatewayBotResponse struct {
	URL               string            `json:"url"`
	Shards            int64             `json:"shards"`
	SessionStartLimit SessionStartLimit `json:"session_start_limit"`
}

type GetGatewayResponse struct {
	URL string `json:"url"`
}

type GetGiftCodeQuery struct {
	WithApplication      bool `json:"with_application,omitzero"`
	WithSubscriptionPlan bool `json:"with_subscription_plan,omitzero"`
}

type GetGuildApplicationSKUStorefrontResponse struct {
	StoreListing       StoreListing       `json:"store_listing"`
	Assets             []StoreAsset       `json:"assets"`
	StorefrontMetadata StorefrontMetadata `json:"storefront_metadata"`
	StorefrontPricing  StorefrontPrices   `json:"storefront_pricing"`
}

type GetGuildApplicationStorefrontAnnouncementResponse struct {
	InteractionID snowflake.ID `json:"interaction_id"`
}

type GetGuildArguments struct {
	GuildID snowflake.ID `json:"guild_id"`
	Timeout int64        `json:"timeout"`
}

type GetGuildJoinRequestCooldownResponse struct {
	Cooldown int64 `json:"cooldown"`
}

type GetGuildMemberVerificationQuery struct {
	WithGuild  bool   `json:"with_guild,omitzero"`
	InviteCode string `json:"invite_code,omitzero"`
}

type GetGuildPruneQuery struct {
	Days         int64          `json:"days,omitzero"`
	IncludeRoles []snowflake.ID `json:"include_roles,omitempty"`
}

type GetGuildPruneResponse struct {
	Pruned int64 `json:"pruned"`
}

type GetGuildQuery struct {
	WithCounts bool `json:"with_counts,omitzero"`
}

type GetGuildScheduledEventQuery struct {
	WithUserCount bool `json:"with_user_count,omitzero"`
}

type GetGuildScheduledEventUserCountQuery struct {
	GuildScheduledEventExceptionIDs []snowflake.ID `json:"guild_scheduled_event_exception_ids,omitempty"`
}

type GetGuildScheduledEventUserCountResponse struct {
	GuildScheduledEventCount           int64                  `json:"guild_scheduled_event_count"`
	GuildScheduledEventExceptionCounts map[snowflake.ID]int64 `json:"guild_scheduled_event_exception_counts"`
}

type GetGuildVanityInviteResponse struct {
	Code *string `json:"code"`
	Uses int64   `json:"uses"`
}

type GetGuildWidgetImageQuery struct {
	Style GuildWidgetImageStyleOption `json:"style,omitzero"`
}

type GetHdStreamingConsumableResponse struct {
	Entitlement *Entitlement `json:"entitlement"`
}

type GetImageArguments struct {
	Type   ImageType    `json:"type"`
	ID     snowflake.ID `json:"id"`
	Format string       `json:"format"`
	Size   int64        `json:"size"`
}

type GetImageResponse struct {
	DataURL string `json:"data_url"`
}

type GetInviteQuery struct {
	WithCounts            bool         `json:"with_counts,omitzero"`
	WithPermissions       bool         `json:"with_permissions,omitzero"`
	GuildScheduledEventID snowflake.ID `json:"guild_scheduled_event_id,omitzero"`
}

type GetLatestApplicationInstallerQuery struct {
	Platform DesktopPlatformType     `json:"platform"`
	Format   DesktopExecutableFormat `json:"format,omitzero"`
}

type GetLatestDistributedApplicationInstallerQuery struct {
	Channel  DesktopReleaseChannel   `json:"channel"`
	Platform DesktopPlatformType     `json:"platform"`
	Arch     DesktopArchitectureType `json:"arch"`
}

type GetLatestDistributedApplicationManifestQuery struct {
	InstallID       string                  `json:"install_id,omitzero"`
	Channel         DesktopReleaseChannel   `json:"channel"`
	Platform        DesktopPlatformType     `json:"platform"`
	Arch            DesktopArchitectureType `json:"arch"`
	PlatformVersion string                  `json:"platform_version,omitzero"`
	ClientVersion   string                  `json:"client_version,omitzero"`
}

type GetLatestDistributedApplicationManifestResponse struct {
	Full            ManifestPackageVersion     `json:"full"`
	Deltas          []ManifestPackageVersion   `json:"deltas"`
	Modules         map[string]ManifestPackage `json:"modules"`
	RequiredModules []string                   `json:"required_modules"`
	MetadataVersion *int64                     `json:"metadata_version"`
	RequiredUpdate  bool                       `json:"required_update"`
}

type GetLayoutDefinitionsResponse struct {
	Configs map[snowflake.ID][]WidgetConfig `json:"configs"`
}

type GetLinkCodeResponse struct {
	LinkCode  string `json:"link_code"`
	ExpiresAt int64  `json:"expires_at"`
}

type GetLocalizedPricingPromoResponse struct {
	CountryCode           string                 `json:"country_code"`
	LocalizedPricingPromo *LocalizedPricingPromo `json:"localized_pricing_promo"`
}

type GetLocationMetadataResponse struct {
	CountryCode           string                   `json:"country_code"`
	ConsentRequired       bool                     `json:"consent_required"`
	PromotionalEmailOptIn PromotionalEmailMetadata `json:"promotional_email_opt_in"`
}

type GetMessageInteractionDataResponse struct {
	ID                 snowflake.ID                   `json:"id"`
	Type               InteractionType                `json:"type"`
	Name               string                         `json:"name"`
	ApplicationCommand ApplicationCommand             `json:"application_command"`
	Options            []ApplicationCommandDataOption `json:"options,omitempty"`
}

type GetMetadataForApexExperimentsResponse struct {
	Experiments []ApexExperimentMetadata `json:"experiments"`
}

type GetNativeModuleQuery struct {
	Platform    DesktopPlatformType `json:"platform,omitzero"`
	HostVersion string              `json:"host_version,omitzero"`
}

type GetNativeModuleVersionsQuery struct {
	Platform    DesktopPlatformType `json:"platform,omitzero"`
	HostVersion string              `json:"host_version,omitzero"`
}

type GetNetworkingConfigResponse struct {
	Address string `json:"address"`
	Token   string `json:"token"`
}

type GetOAuth2AuthorizationsQuery struct {
	ApplicationIDs []snowflake.ID `json:"application_ids"`
}

type GetOAuth2DeviceCodeParams struct {
	ClientID     snowflake.ID `json:"client_id,omitzero"`
	ClientSecret string       `json:"client_secret,omitzero"`
	Scope        string       `json:"scope,omitzero"`
}

type GetOAuth2DeviceCodeResponse struct {
	DeviceCode              string `json:"device_code"`
	UserCode                string `json:"user_code"`
	VerificationURI         string `json:"verification_uri"`
	VerificationURIComplete string `json:"verification_uri_complete"`
	ExpiresIn               int64  `json:"expires_in"`
	Interval                int64  `json:"interval"`
}

type GetOAuth2DeviceFlowParams struct {
	UserCode string `json:"user_code"`
}

type GetOAuth2DeviceFlowResponse struct {
	Scopes         []string               `json:"scopes"`
	ClientID       snowflake.ID           `json:"client_id"`
	TwoWayLinkCode *string                `json:"two_way_link_code"`
	Type           OAuth2DeviceFlowStatus `json:"type"`
}

type GetOAuth2TokenParams struct {
	GrantType         OAuth2GrantType                    `json:"grant_type"`
	ClientID          snowflake.ID                       `json:"client_id,omitzero"`
	ClientSecret      string                             `json:"client_secret,omitzero"`
	Code              string                             `json:"code,omitzero"`
	CodeVerifier      string                             `json:"code_verifier,omitzero"`
	RedirectURI       string                             `json:"redirect_uri,omitzero"`
	RefreshToken      string                             `json:"refresh_token,omitzero"`
	DeviceCode        string                             `json:"device_code,omitzero"`
	Scope             string                             `json:"scope,omitzero"`
	ExternalAuthType  ExternalProviderAuthenticationType `json:"external_auth_type,omitzero"`
	ExternalAuthToken string                             `json:"external_auth_token,omitzero"`
}

type GetOpenidUserInformationResponse struct {
	Sub               string  `json:"sub"`
	Email             *string `json:"email"`
	EmailVerified     bool    `json:"email_verified"`
	PreferredUsername string  `json:"preferred_username"`
	Nickname          *string `json:"nickname"`
	Picture           string  `json:"picture"`
	Locale            string  `json:"locale"`
}

type GetPartialApplicationQuery struct {
	WithGuild bool `json:"with_guild,omitzero"`
}

type GetPasswordStrengthParams struct {
	Password string `json:"password"`
}

type GetPasswordStrengthResponse struct {
	Valid            bool  `json:"valid"`
	PasswordStrength int64 `json:"password_strength"`
}

type GetPaymentInvoiceBreakdownQuery struct {
	PaymentID snowflake.ID `json:"payment_id"`
}

type GetPaymentInvoiceBreakdownResponse struct {
	Invoicelink        string   `json:"invoiceLink,omitzero"`
	Refundinvoicelinks []string `json:"refundInvoiceLinks,omitempty"`
}

type GetPaymentSourceCreationContextResponse struct {
	StoreCountry                   *string  `json:"store_country"`
	AllowedPaymentSourceTypes      []int64  `json:"allowed_payment_source_types"`
	AllowedBillingAddressCountries []string `json:"allowed_billing_address_countries"`
}

type GetPlatformBehaviorsResponse struct {
	Ioskeyboardresizesview bool `json:"iosKeyboardResizesView,omitzero"`
}

type GetPremiumGuildSubscriptionCooldownResponse struct {
	EndsAt    time.Time `json:"ends_at"`
	Limit     int64     `json:"limit"`
	Remaining int64     `json:"remaining"`
}

type GetPremiumReferralIncentiveEligibilityResponse struct {
	IsEligibleForIncentive bool `json:"is_eligible_for_incentive"`
}

type GetProfileWidgetsSuggestedGamesResponse struct {
	SuggestedGames         []snowflake.ID `json:"suggested_games"`
	SuggestedWishlistGames []snowflake.ID `json:"suggested_wishlist_games"`
}

type GetProviderAccessTokenArguments struct {
	Provider           ConnectionType `json:"provider"`
	ConnectionRedirect string         `json:"connection_redirect,omitzero"`
}

type GetProviderAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
}

type GetProvisionalAccountTokenParams struct {
	ClientID          snowflake.ID                       `json:"client_id"`
	ClientSecret      string                             `json:"client_secret,omitzero"`
	ExternalAuthType  ExternalProviderAuthenticationType `json:"external_auth_type"`
	ExternalAuthToken string                             `json:"external_auth_token"`
}

type GetProvisionalAccountTokenWithBotParams struct {
	ExternalUserID      string `json:"external_user_id"`
	PreferredGlobalName string `json:"preferred_global_name,omitzero"`
}

type GetQuestCreativePreviewQuery struct {
	AdCreativeIDs []snowflake.ID `json:"ad_creative_ids"`
}

type GetQuestCreativePreviewResponse struct {
	RequestID string            `json:"request_id"`
	Decisions []QuestAdDecision `json:"decisions"`
}

type GetQuestDecisionsQuery struct {
	Placement                QuestPlacementArea `json:"placement"`
	NumDecisionsRequested    int64              `json:"num_decisions_requested"`
	ClientHeartbeatSessionID string             `json:"client_heartbeat_session_id,omitzero"`
	ClientAdSessionID        string             `json:"client_ad_session_id,omitzero"`
}

type GetQuestDecisionsResponse struct {
	RequestID string            `json:"request_id"`
	Decisions []QuestAdDecision `json:"decisions"`
}

type GetQuestEnrollmentStatusArguments struct {
	QuestID snowflake.ID `json:"quest_id"`
}

type GetQuestEnrollmentStatusResponse struct {
	QuestID    snowflake.ID `json:"quest_id"`
	IsEnrolled bool         `json:"is_enrolled"`
	EnrolledAt *time.Time   `json:"enrolled_at"`
}

type GetQuestPlacementQuery struct {
	Placement                QuestPlacementArea `json:"placement"`
	ClientHeartbeatSessionID string             `json:"client_heartbeat_session_id,omitzero"`
	ClientAdSessionID        string             `json:"client_ad_session_id,omitzero"`
	VisibleGuildIDs          []snowflake.ID     `json:"visible_guild_ids,omitempty"`
}

type GetQuestPlacementResponse struct {
	RequestID             string              `json:"request_id"`
	Quest                 *Quest              `json:"quest"`
	Creative              *QuestAdCreative    `json:"creative"`
	AdIdentifiers         *QuestAdIdentifiers `json:"ad_identifiers"`
	AdContext             *QuestAdContext     `json:"ad_context"`
	ResponseTTLSeconds    int64               `json:"response_ttl_seconds,omitzero"`
	MetadataSealed        *string             `json:"metadata_sealed,omitempty"`
	TrafficMetadataSealed *string             `json:"traffic_metadata_sealed,omitempty"`
}

type GetRelationshipsResponse struct {
	Relationships []RPCRelationship `json:"relationships"`
}

type GetReportMenuQuery struct {
	Variant string `json:"variant,omitzero"`
}

type GetReportReasonsQuery struct {
	MessageID snowflake.ID `json:"message_id"`
	ChannelID snowflake.ID `json:"channel_id"`
	UserID    snowflake.ID `json:"user_id"`
}

type GetSKUPublishedStoreListingQuery struct {
	CountryCode string `json:"country_code,omitzero"`
	Localize    bool   `json:"localize,omitzero"`
}

type GetSKUPurchasePreviewQuery struct {
	PaymentSourceID       snowflake.ID `json:"payment_source_id,omitzero"`
	Currency              string       `json:"currency,omitzero"`
	SKUSubscriptionPlanID snowflake.ID `json:"sku_subscription_plan_id,omitzero"`
	Gift                  bool         `json:"gift,omitzero"`
	TestMode              bool         `json:"test_mode,omitzero"`
	LoadID                string       `json:"load_id,omitzero"`
}

type GetSKUQuery struct {
	CountryCode string `json:"country_code,omitzero"`
	Localize    bool   `json:"localize,omitzero"`
}

type GetSKUsEmbeddedResponse struct {
	SKUs []SKU `json:"skus"`
}

type GetSocialLayerStorefrontConfigResponse struct {
	PromotionalSKUIDs    []snowflake.ID     `json:"promotional_sku_ids"`
	PromotionEndDatetime *time.Time         `json:"promotion_end_datetime"`
	Storefronts          []StorefrontConfig `json:"storefronts"`
}

type GetStorefrontCollectionQuery struct {
	CountryCode                  string       `json:"country_code,omitzero"`
	GuildID                      snowflake.ID `json:"guild_id,omitzero"`
	IncludeUnpublishedProducts   bool         `json:"include_unpublished_products,omitzero"`
	IncludeUnpublishedCollection bool         `json:"include_unpublished_collection,omitzero"`
}

type GetStorefrontCollectionResponse struct {
	Collection StorefrontCollection `json:"collection"`
	Products   []StorefrontProduct  `json:"products"`
}

type GetStreamPreviewResponse struct {
	URL string `json:"url"`
}

type GetStripeSetupIntentSecretParams struct {
	RegionalPaymentElementSourceTypes []int64 `json:"regional_payment_element_source_types,omitempty"`
}

type GetStripeSetupIntentSecretResponse struct {
	ClientSecret string `json:"client_secret"`
}

type GetSubscriptionGroupListingBySubscriptionPlanResponse struct {
	ID                      snowflake.ID          `json:"id"`
	ApplicationID           *snowflake.ID         `json:"application_id"`
	SKUFlags                SKUFlags              `json:"sku_flags"`
	Published               bool                  `json:"published"`
	Name                    string                `json:"name,omitzero"`
	Description             string                `json:"description,omitzero"`
	SubscriptionListingsIDs []snowflake.ID        `json:"subscription_listings_ids"`
	SubscriptionListings    []SubscriptionListing `json:"subscription_listings"`
}

type GetSupplementalMessageRequestDataQuery struct {
	ChannelIDs []snowflake.ID `json:"channel_ids"`
}

type GetSuspendedUserSafetyHubParams struct {
	Token string `json:"token"`
}

type GetTeamPayoutOnboardingResponse struct {
	URL string `json:"url"`
}

type GetTeamPayoutReportQuery struct {
	Type TeamPayoutReportType `json:"type"`
}

type GetTeamStripeConnectURLParams struct {
	CountryCode string `json:"country_code,omitzero"`
}

type GetTeamStripeConnectURLResponse struct {
	StripeConnectRedirectURL string `json:"stripe_connect_redirect_url"`
}

type GetThreadMemberQuery struct {
	WithMember bool `json:"with_member,omitzero"`
}

type GetUnauthenticatedReportCapabilitiesResponse struct {
	Capabilities []string `json:"capabilities"`
}

type GetUnauthenticatedReportMenuQuery struct {
	Variant string `json:"variant,omitzero"`
}

type GetUnauthenticatedReportVerificationCodeParams struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type GetUniqueUsernameEligibilityParams struct {
	Username string `json:"username"`
}

type GetUniqueUsernameEligibilityResponse struct {
	Taken *bool `json:"taken"`
}

type GetUniqueUsernameSuggestionsQuery struct {
	GlobalName string `json:"global_name,omitzero"`
}

type GetUniqueUsernameSuggestionsResponse struct {
	Username string `json:"username"`
}

type GetUserApplicationCommandIndexParams struct {
	Type                     int64                      `json:"type,omitzero"`
	Name                     string                     `json:"name"`
	NameLocalizations        map[string]string          `json:"name_localizations,omitempty"`
	Description              string                     `json:"description"`
	DescriptionLocalizations map[string]string          `json:"description_localizations,omitempty"`
	Options                  []ApplicationCommandOption `json:"options,omitempty"`
	DefaultMemberPermissions *string                    `json:"default_member_permissions,omitempty"`
	NSFW                     bool                       `json:"nsfw,omitzero"`
}

type GetUserApplicationCommandIndexParamsApplicationCommands struct {
	Type                     int64                         `json:"type,omitzero"`
	Name                     string                        `json:"name"`
	NameLocalizations        map[string]string             `json:"name_localizations,omitempty"`
	Description              string                        `json:"description,omitzero"`
	DescriptionLocalizations map[string]string             `json:"description_localizations,omitempty"`
	Options                  []ApplicationCommandOption    `json:"options,omitempty"`
	DefaultMemberPermissions *string                       `json:"default_member_permissions,omitempty"`
	DMPermission             bool                          `json:"dm_permission,omitzero"`
	NSFW                     bool                          `json:"nsfw,omitzero"`
	IntegrationTypes         []int64                       `json:"integration_types,omitempty"`
	Contexts                 []int64                       `json:"contexts,omitempty"`
	Handler                  ApplicationCommandHandlerType `json:"handler,omitzero"`
}

type GetUserApplicationCommandIndexParamsApplicationCommandsXxx struct {
	Name                     string                     `json:"name,omitzero"`
	NameLocalizations        map[string]string          `json:"name_localizations,omitempty"`
	Description              string                     `json:"description,omitzero"`
	DescriptionLocalizations map[string]string          `json:"description_localizations,omitempty"`
	Options                  []ApplicationCommandOption `json:"options,omitempty"`
	DefaultMemberPermissions *string                    `json:"default_member_permissions,omitempty"`
	NSFW                     *bool                      `json:"nsfw,omitempty"`
}

type GetUserApplicationCommandIndexParamsApplicationCommandsXxxx struct {
	Permissions []ApplicationCommandPermissions `json:"permissions"`
}

type GetUserApplicationCommandIndexQuery struct {
	WithLocalizations bool `json:"with_localizations,omitzero"`
}

type GetUserConnectionAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
}

type GetUserNoteResponse struct {
	Note       string       `json:"note"`
	NoteUserID snowflake.ID `json:"note_user_id"`
	UserID     snowflake.ID `json:"user_id"`
}

type GetUserOfferParams struct {
	PaymentGateway PaymentGateway `json:"payment_gateway,omitzero"`
	OfferID        snowflake.ID   `json:"offer_id,omitzero"`
}

type GetUserOfferResponse struct {
	UserTrialOffer    *UserTrialOffer    `json:"user_trial_offer"`
	UserDiscountOffer *UserDiscountOffer `json:"user_discount_offer"`
	UserDiscount      *UserDiscountOffer `json:"user_discount"`
}

type GetUserPremiumUsageResponse struct {
	NitroStickerSends   PremiumUsage `json:"nitro_sticker_sends"`
	TotalAnimatedEmojis PremiumUsage `json:"total_animated_emojis"`
	TotalGlobalEmojis   PremiumUsage `json:"total_global_emojis"`
	TotalLargeUploads   PremiumUsage `json:"total_large_uploads"`
	TotalHdStreams      PremiumUsage `json:"total_hd_streams"`
	HdHoursStreamed     PremiumUsage `json:"hd_hours_streamed"`
}

type GetUserProfileQuery struct {
	WithMutualGuilds       bool               `json:"with_mutual_guilds,omitzero"`
	WithMutualFriends      bool               `json:"with_mutual_friends,omitzero"`
	WithMutualFriendsCount bool               `json:"with_mutual_friends_count,omitzero"`
	GuildID                snowflake.ID       `json:"guild_id,omitzero"`
	ConnectionsRoleID      snowflake.ID       `json:"connections_role_id,omitzero"`
	JoinRequestID          snowflake.ID       `json:"join_request_id,omitzero"`
	Type                   ProfileRequestType `json:"type,omitzero"`
}

type GetUserProfileResponse struct {
	Application                ProfileApplication          `json:"application,omitzero"`
	User                       PartialUser                 `json:"user"`
	UserProfile                ProfileMetadata             `json:"user_profile"`
	Badges                     []ProfileBadge              `json:"badges"`
	GuildMember                GuildMember                 `json:"guild_member,omitzero"`
	GuildMemberProfile         ProfileMetadata             `json:"guild_member_profile,omitzero"`
	GuildBadges                []ProfileBadge              `json:"guild_badges"`
	Widgets                    []GameWidget                `json:"widgets,omitempty"`
	LegacyUsername             *string                     `json:"legacy_username,omitempty"`
	MutualGuilds               []MutualGuild               `json:"mutual_guilds,omitempty"`
	MutualFriends              []PartialUser               `json:"mutual_friends,omitempty"`
	MutualFriendsCount         int64                       `json:"mutual_friends_count,omitzero"`
	ConnectedAccounts          []PartialConnection         `json:"connected_accounts"`
	ApplicationRoleConnections []ApplicationRoleConnection `json:"application_role_connections,omitempty"`
	PremiumType                *PremiumType                `json:"premium_type"`
	PremiumSince               *time.Time                  `json:"premium_since"`
	PremiumGuildSince          *time.Time                  `json:"premium_guild_since"`
	Private                    bool                        `json:"private,omitzero"`
}

type GetUserSettingsProtoResponse struct {
	Settings string `json:"settings"`
}

type GetUserSurveyQuery struct {
	DisableAutoSeen bool         `json:"disable_auto_seen,omitzero"`
	SurveyOverride  snowflake.ID `json:"survey_override,omitzero"`
}

type GetUserSurveyResponse struct {
	Survey *UserSurvey `json:"survey"`
}

type GetValidCollectiblesGiftRecipientQuery struct {
	RecipientID snowflake.ID `json:"recipient_id"`
	SKUID       snowflake.ID `json:"sku_id"`
}

type GetValidCollectiblesGiftRecipientsBatchQuery struct {
	RecipientID snowflake.ID   `json:"recipient_id"`
	SKUIDs      []snowflake.ID `json:"sku_ids"`
}

type GetVirtualCurrencyBalanceResponse struct {
	Balance int64 `json:"balance"`
}

type GiftCode struct {
	Code                string            `json:"code"`
	SKUID               snowflake.ID      `json:"sku_id"`
	ApplicationID       snowflake.ID      `json:"application_id"`
	Flags               GiftCodeFlags     `json:"flags,omitzero"`
	Uses                int64             `json:"uses"`
	MaxUses             int64             `json:"max_uses"`
	Redeemed            bool              `json:"redeemed"`
	ExpiresAt           *time.Time        `json:"expires_at"`
	BatchID             snowflake.ID      `json:"batch_id,omitzero"`
	EntitlementBranches []snowflake.ID    `json:"entitlement_branches,omitempty"`
	GiftStyle           *GiftStyle        `json:"gift_style,omitempty"`
	User                PartialUser       `json:"user,omitzero"`
	StoreListing        StoreListing      `json:"store_listing,omitzero"`
	SubscriptionPlanID  snowflake.ID      `json:"subscription_plan_id,omitzero"`
	SubscriptionPlan    SubscriptionPlan  `json:"subscription_plan,omitzero"`
	SubscriptionTrial   SubscriptionTrial `json:"subscription_trial,omitzero"`
	Promotion           Promotion         `json:"promotion,omitzero"`
}

type GiftCodeBatch struct {
	ID                  snowflake.ID   `json:"id"`
	SKUID               snowflake.ID   `json:"sku_id"`
	Amount              int64          `json:"amount"`
	Description         string         `json:"description,omitzero"`
	EntitlementBranches []snowflake.ID `json:"entitlement_branches,omitempty"`
	EntitlementStartsAt time.Time      `json:"entitlement_starts_at,omitzero"`
	EntitlementEndsAt   time.Time      `json:"entitlement_ends_at,omitzero"`
}

type GiftCodeBrowserArguments struct {
	Code string `json:"code"`
}

type GiftCodeBrowserResponse struct {
	Giftcode GiftCode `json:"giftCode"`
}

type GiftEligibility struct {
	Valid bool `json:"valid"`
}

type GiftInfoOptions struct {
	GiftStyle     *GiftStyle     `json:"gift_style,omitempty"`
	RecipientID   snowflake.ID   `json:"recipient_id,omitzero"`
	CustomMessage *string        `json:"custom_message,omitempty"`
	EmojiID       *snowflake.ID  `json:"emoji_id,omitempty"`
	EmojiName     *string        `json:"emoji_name,omitempty"`
	SoundID       *snowflake.ID  `json:"sound_id,omitempty"`
	RewardSKUIDs  []snowflake.ID `json:"reward_sku_ids,omitempty"`
}

type GlobalActivityStatistics struct {
	UserID        string             `json:"user_id"`
	User          PartialUser        `json:"user,omitzero"`
	ApplicationID string             `json:"application_id"`
	Application   PartialApplication `json:"application,omitzero"`
	UpdatedAt     time.Time          `json:"updated_at"`
	Duration      int64              `json:"duration"`
}

type GoreContentSettings struct {
	GoreContentGuilds      json.RawMessage `json:"gore_content_guilds"`
	GoreContentFriendDM    json.RawMessage `json:"gore_content_friend_dm"`
	GoreContentNonFriendDM json.RawMessage `json:"gore_content_non_friend_dm"`
}

type GrowthActivationJoins struct {
	DayPt        time.Time `json:"day_pt"`
	Joins        int64     `json:"joins"`
	AllTimeJoins int64     `json:"all_time_joins,omitzero"`
}

type GrowthActivationJoinsByInvite struct {
	DayPt      time.Time `json:"day_pt"`
	InviteLink string    `json:"invite_link"`
	Joins      int64     `json:"joins"`
}

type GrowthActivationJoinsByReferrer struct {
	DayPt           time.Time `json:"day_pt"`
	ReferringDomain string    `json:"referring_domain"`
	Joins           int64     `json:"joins"`
}

type GrowthActivationJoinsBySource struct {
	DayPt            time.Time `json:"day_pt"`
	DiscoveryJoins   int64     `json:"discovery_joins"`
	Invites          int64     `json:"invites"`
	VanityJoins      int64     `json:"vanity_joins"`
	HubsJoins        int64     `json:"hubs_joins"`
	BotJoins         int64     `json:"bot_joins"`
	IntegrationJoins int64     `json:"integration_joins"`
	OtherJoins       int64     `json:"other_joins"`
	TotalJoins       int64     `json:"total_joins"`
}

type GrowthActivationLeavers struct {
	DayPt       time.Time `json:"day_pt"`
	DaysInGuild string    `json:"days_in_guild"`
	Leavers     int64     `json:"leavers"`
}

type GrowthActivationMembership struct {
	DayPt           time.Time `json:"day_pt"`
	TotalMembership int64     `json:"total_membership"`
}

type GrowthActivationOverview struct {
	DayPt            time.Time `json:"day_pt"`
	NewMembers       int64     `json:"new_members"`
	NewCommunicators int64     `json:"new_communicators"`
}

type GrowthActivationPercentage struct {
	DayPt             time.Time `json:"day_pt"`
	NewMembers        int64     `json:"new_members"`
	PctCommunicated   float64   `json:"pct_communicated,omitzero"`
	PctOpenedChannels float64   `json:"pct_opened_channels,omitzero"`
}

type GrowthActivationRetention struct {
	DayPt       time.Time `json:"day_pt"`
	NewMembers  int64     `json:"new_members"`
	PctRetained float64   `json:"pct_retained,omitzero"`
}

type Guild struct {
	ID                                     snowflake.ID                    `json:"id"`
	Name                                   string                          `json:"name"`
	Icon                                   *string                         `json:"icon"`
	Banner                                 *string                         `json:"banner"`
	HomeHeader                             *string                         `json:"home_header"`
	Splash                                 *string                         `json:"splash"`
	DiscoverySplash                        *string                         `json:"discovery_splash"`
	OwnerID                                snowflake.ID                    `json:"owner_id"`
	ApplicationID                          *snowflake.ID                   `json:"application_id"`
	Description                            *string                         `json:"description"`
	Region                                 *string                         `json:"region,omitempty"`
	AFKChannelID                           *snowflake.ID                   `json:"afk_channel_id"`
	AFKTimeout                             int64                           `json:"afk_timeout"`
	WidgetEnabled                          bool                            `json:"widget_enabled,omitzero"`
	WidgetChannelID                        *snowflake.ID                   `json:"widget_channel_id,omitempty"`
	VerificationLevel                      VerificationLevel               `json:"verification_level"`
	DefaultMessageNotifications            MessageNotificationLevel        `json:"default_message_notifications"`
	ExplicitContentFilter                  GuildExplicitContentFilterLevel `json:"explicit_content_filter"`
	Features                               []string                        `json:"features"`
	Roles                                  []Role                          `json:"roles"`
	Emojis                                 []Emoji                         `json:"emojis"`
	Stickers                               []Sticker                       `json:"stickers"`
	MFALevel                               MFALevel                        `json:"mfa_level"`
	SystemChannelID                        *snowflake.ID                   `json:"system_channel_id"`
	SystemChannelFlags                     SystemChannelFlags              `json:"system_channel_flags"`
	RulesChannelID                         *snowflake.ID                   `json:"rules_channel_id"`
	PublicUpdatesChannelID                 *snowflake.ID                   `json:"public_updates_channel_id"`
	SafetyAlertsChannelID                  *snowflake.ID                   `json:"safety_alerts_channel_id"`
	MaxPresences                           *int64                          `json:"max_presences,omitempty"`
	MaxMembers                             int64                           `json:"max_members,omitzero"`
	VanityURLCode                          *string                         `json:"vanity_url_code"`
	PremiumTier                            PremiumTier                     `json:"premium_tier"`
	PremiumSubscriptionCount               int64                           `json:"premium_subscription_count"`
	PreferredLocale                        string                          `json:"preferred_locale"`
	MaxVideoChannelUsers                   int64                           `json:"max_video_channel_users,omitzero"`
	MaxStageVideoChannelUsers              int64                           `json:"max_stage_video_channel_users,omitzero"`
	NSFW                                   bool                            `json:"nsfw"`
	NSFWLevel                              NSFWLevel                       `json:"nsfw_level"`
	OwnerConfiguredContentLevel            *NSFWLevel                      `json:"owner_configured_content_level"`
	HubType                                *HubType                        `json:"hub_type"`
	PremiumProgressBarEnabled              bool                            `json:"premium_progress_bar_enabled"`
	PremiumProgressBarEnabledUserUpdatedAt *time.Time                      `json:"premium_progress_bar_enabled_user_updated_at"`
	LatestOnboardingQuestionID             *snowflake.ID                   `json:"latest_onboarding_question_id"`
	IncidentsData                          *AutomodIncidentsData           `json:"incidents_data"`
	InventorySettings                      *GuildInventorySettings         `json:"inventory_settings"`
	PremiumFeatures                        *GuildPremiumFeatures           `json:"premium_features"`
	Profile                                *GuildIdentity                  `json:"profile"`
	ApproximateMemberCount                 int64                           `json:"approximate_member_count,omitzero"`
	ApproximatePresenceCount               int64                           `json:"approximate_presence_count,omitzero"`
	OfficialMessageColor                   *int64                          `json:"official_message_color"`
	Version                                string                          `json:"version,omitzero"`
}

type GuildAffinity struct {
	GuildID  snowflake.ID `json:"guild_id"`
	Affinity float64      `json:"affinity"`
}

type GuildAndChannelFrecency struct {
	GuildAndChannels map[string]FrecencyItem `json:"guild_and_channels"`
}

type GuildApplicationCommandIndexUpdate struct {
	GuildID snowflake.ID `json:"guild_id"`
}

type GuildApplicationCommandPermissions struct {
	ID            snowflake.ID                    `json:"id"`
	ApplicationID snowflake.ID                    `json:"application_id"`
	GuildID       snowflake.ID                    `json:"guild_id"`
	Permissions   []ApplicationCommandPermissions `json:"permissions"`
}

type GuildBanAdd struct {
	GuildID           snowflake.ID `json:"guild_id"`
	User              PartialUser  `json:"user"`
	DeleteMessageSecs int64        `json:"delete_message_secs"`
}

type GuildBanRemove struct {
	GuildID snowflake.ID `json:"guild_id"`
	User    PartialUser  `json:"user"`
}

type GuildBulkBan struct {
	BannedUsers []snowflake.ID `json:"banned_users"`
	FailedUsers []snowflake.ID `json:"failed_users"`
}

type GuildBulkBanUpdate struct {
	GuildID snowflake.ID `json:"guild_id"`
	BulkBan GuildBulkBan `json:"bulk_ban"`
}

type GuildCollectibles struct {
	Nameplate *GuildNameplateData `json:"nameplate,omitempty"`
}

type GuildDeleteGuildRoleSubscriptionGroupListingParams struct {
	Name                 string                    `json:"name"`
	Description          string                    `json:"description"`
	Image                ImageData                 `json:"image"`
	PriceTier            int64                     `json:"price_tier"`
	Benefits             []RoleSubscriptionBenefit `json:"benefits"`
	CanAccessAllChannels bool                      `json:"can_access_all_channels,omitzero"`
}

type GuildDirectoryEntryDelete struct {
	Type               DirectoryEntryType `json:"type"`
	DirectoryChannelID snowflake.ID       `json:"directory_channel_id"`
	GuildID            snowflake.ID       `json:"guild_id"`
	EntityID           snowflake.ID       `json:"entity_id"`
	CreatedAt          string             `json:"created_at"`
	PrimaryCategoryID  DirectoryCategory  `json:"primary_category_id,omitzero"`
	Description        *string            `json:"description"`
	AuthorID           snowflake.ID       `json:"author_id"`
}

type GuildDismissibleContentState struct {
	Dismissed             bool            `json:"dismissed"`
	LastDismissedVersion  json.RawMessage `json:"last_dismissed_version"`
	LastDismissedAtMs     json.RawMessage `json:"last_dismissed_at_ms"`
	LastDismissedObjectID json.RawMessage `json:"last_dismissed_object_id"`
	NumTimesDismissed     json.RawMessage `json:"num_times_dismissed"`
}

type GuildEmojisUpdate struct {
	GuildID snowflake.ID `json:"guild_id"`
	Emojis  []Emoji      `json:"emojis"`
}

type GuildExperiment struct {
	Hash               int64                      `json:"hash"`
	HashKey            *string                    `json:"hash_key"`
	Revision           int64                      `json:"revision"`
	Populations        []ExperimentPopulation     `json:"populations"`
	Overrides          []ExperimentBucketOverride `json:"overrides"`
	OverridesFormatted [][]ExperimentPopulation   `json:"overrides_formatted"`
	HoldoutName        *string                    `json:"holdout_name"`
	HoldoutBucket      *int64                     `json:"holdout_bucket"`
	AaMode             int64                      `json:"aa_mode"`
	TriggerDebugging   int64                      `json:"trigger_debugging"`
}

type GuildFeatureAck struct {
	AckType    ReadStateType `json:"ack_type"`
	ResourceID snowflake.ID  `json:"resource_id"`
	EntityID   snowflake.ID  `json:"entity_id"`
	Version    int64         `json:"version"`
}

type GuildFolder struct {
	Color    *int64         `json:"color"`
	GuildIDs []snowflake.ID `json:"guild_ids"`
	ID       *int64         `json:"id"`
	Name     *string        `json:"name"`
}

type GuildFolders struct {
	Folders        []UserSettingsProtoGuildFolder `json:"folders"`
	GuildPositions []json.RawMessage              `json:"guild_positions"`
}

type GuildIdentity struct {
	Tag   string `json:"tag"`
	Badge string `json:"badge"`
}

type GuildIntegrationsUpdate struct {
	GuildID snowflake.ID `json:"guild_id"`
}

type GuildInventorySettings struct {
	IsEmojiPackCollectible bool `json:"is_emoji_pack_collectible,omitzero"`
}

type GuildJoinRequest struct {
	ID                 snowflake.ID                  `json:"id"`
	JoinRequestID      snowflake.ID                  `json:"join_request_id"`
	CreatedAt          time.Time                     `json:"created_at"`
	ApplicationStatus  GuildJoinRequestStatus        `json:"application_status"`
	GuildID            snowflake.ID                  `json:"guild_id"`
	FormResponses      []MemberVerificationFormField `json:"form_responses,omitempty"`
	LastSeen           *time.Time                    `json:"last_seen"`
	ActionedAt         snowflake.ID                  `json:"actioned_at,omitzero"`
	ReviewedAt         time.Time                     `json:"reviewed_at,omitzero"`
	ActionedByUser     PartialUser                   `json:"actioned_by_user,omitzero"`
	RejectionReason    *string                       `json:"rejection_reason"`
	UserID             snowflake.ID                  `json:"user_id"`
	User               PartialUser                   `json:"user,omitzero"`
	InterviewChannelID *snowflake.ID                 `json:"interview_channel_id"`
}

type GuildJoinRequestCreate struct {
	GuildID snowflake.ID           `json:"guild_id"`
	Request GuildJoinRequest       `json:"request"`
	Status  GuildJoinRequestStatus `json:"status"`
}

type GuildJoinRequestDelete struct {
	GuildID snowflake.ID `json:"guild_id"`
	ID      snowflake.ID `json:"id"`
	UserID  snowflake.ID `json:"user_id"`
}

type GuildJoinRequestUpdate struct {
	GuildID snowflake.ID           `json:"guild_id"`
	Request GuildJoinRequest       `json:"request"`
	Status  GuildJoinRequestStatus `json:"status"`
}

type GuildListGuildRoleSubscriptionListingTemplatesResponse struct {
	Templates []RoleSubscriptionListingTemplate `json:"templates"`
}

type GuildMember struct {
	User                       PartialUser           `json:"user"`
	Nick                       *string               `json:"nick,omitempty"`
	Avatar                     *string               `json:"avatar,omitempty"`
	AvatarDecorationData       *AvatarDecorationData `json:"avatar_decoration_data,omitempty"`
	Collectibles               *GuildCollectibles    `json:"collectibles,omitempty"`
	DisplayNameStyles          *DisplayNameStyle     `json:"display_name_styles,omitempty"`
	Banner                     *string               `json:"banner,omitempty"`
	Bio                        *string               `json:"bio,omitempty"`
	Roles                      []snowflake.ID        `json:"roles"`
	JoinedAt                   time.Time             `json:"joined_at"`
	PremiumSince               *time.Time            `json:"premium_since,omitempty"`
	Deaf                       bool                  `json:"deaf,omitzero"`
	Mute                       bool                  `json:"mute,omitzero"`
	Pending                    bool                  `json:"pending,omitzero"`
	CommunicationDisabledUntil *time.Time            `json:"communication_disabled_until,omitempty"`
	UnusualDMActivityUntil     *time.Time            `json:"unusual_dm_activity_until,omitempty"`
	Flags                      GuildMemberFlags      `json:"flags"`
	Permissions                Permissions           `json:"permissions,omitzero"`
	UserID                     snowflake.ID          `json:"user_id,omitzero"`
}

type GuildMemberAddExtraFields struct {
	GuildID snowflake.ID `json:"guild_id"`
}

type GuildMemberRemove struct {
	GuildID snowflake.ID `json:"guild_id"`
	User    User         `json:"user"`
}

type GuildMemberUnusualDMActivity struct {
	UserID                 snowflake.ID `json:"user_id"`
	GuildID                snowflake.ID `json:"guild_id"`
	UnusualDMActivityUntil time.Time    `json:"unusual_dm_activity_until"`
}

type GuildMemberUpdate struct {
	GuildID                    snowflake.ID          `json:"guild_id"`
	User                       PartialUser           `json:"user"`
	Nick                       *string               `json:"nick,omitempty"`
	Avatar                     *string               `json:"avatar"`
	AvatarDecorationData       *AvatarDecorationData `json:"avatar_decoration_data,omitempty"`
	Roles                      []snowflake.ID        `json:"roles"`
	JoinedAt                   time.Time             `json:"joined_at"`
	PremiumSince               *time.Time            `json:"premium_since"`
	Deaf                       bool                  `json:"deaf,omitzero"`
	Mute                       bool                  `json:"mute,omitzero"`
	Pending                    bool                  `json:"pending,omitzero"`
	CommunicationDisabledUntil *time.Time            `json:"communication_disabled_until"`
	Flags                      GuildMemberFlags      `json:"flags"`
}

type GuildMembersChunk struct {
	GuildID    snowflake.ID    `json:"guild_id"`
	Members    []GuildMember   `json:"members"`
	ChunkIndex int64           `json:"chunk_index"`
	ChunkCount int64           `json:"chunk_count"`
	NotFound   json.RawMessage `json:"not_found,omitempty"`
	Presences  []Presence      `json:"presences,omitempty"`
	Nonce      string          `json:"nonce,omitzero"`
}

type GuildMetadata struct {
	Name       string          `json:"name"`
	Icon       *string         `json:"icon"`
	MemberType GuildMemberType `json:"member_type"`
}

type GuildMonetizationMetadata struct {
	Powerup    GuildPowerupMetadata      `json:"powerup,omitzero"`
	GameServer GameServerPowerupMetadata `json:"game_server,omitzero"`
}

type GuildMonetizationProductMetadata struct {
	GameServer GameServerPowerupProductMetadata `json:"game_server,omitzero"`
}

type GuildNameplateData struct {
	SKUID snowflake.ID `json:"sku_id,omitzero"`
}

type GuildOfficialGameApplicationsUpdate struct {
	GuildID            snowflake.ID   `json:"guild_id"`
	GameApplicationIDs []snowflake.ID `json:"game_application_ids"`
}

type GuildPowerupEntitlementsCreate struct {
	GuildID      snowflake.ID  `json:"guild_id"`
	Entitlements []Entitlement `json:"entitlements"`
}

type GuildPowerupEntitlementsDelete struct {
	GuildID      snowflake.ID  `json:"guild_id"`
	Entitlements []Entitlement `json:"entitlements"`
}

type GuildPowerupMetadata struct {
	BoostPrice       int64                    `json:"boost_price"`
	PurchaseLimit    int64                    `json:"purchase_limit"`
	GuildFeatures    GuildPremiumFeatures     `json:"guild_features"`
	CategoryType     GuildPowerupCategoryType `json:"category_type"`
	StaticImageURL   string                   `json:"static_image_url"`
	AnimatedImageURL string                   `json:"animated_image_url"`
	StoreRemovalDate json.RawMessage          `json:"store_removal_date,omitempty"`
}

type GuildPremiumFeatures struct {
	Features               []string `json:"features"`
	AdditionalEmojiSlots   int64    `json:"additional_emoji_slots"`
	AdditionalStickerSlots int64    `json:"additional_sticker_slots"`
	AdditionalSoundSlots   int64    `json:"additional_sound_slots"`
}

type GuildProductListing struct {
	ID               snowflake.ID `json:"id"`
	ApplicationID    snowflake.ID `json:"application_id"`
	GuildID          snowflake.ID `json:"guild_id"`
	Name             string       `json:"name"`
	Description      string       `json:"description"`
	ImageAsset       *StoreAsset  `json:"image_asset"`
	RoleID           snowflake.ID `json:"role_id"`
	Published        bool         `json:"published"`
	HasEntitlement   bool         `json:"has_entitlement"`
	AttachmentsCount int64        `json:"attachments_count"`
	PublishedAt      time.Time    `json:"published_at"`
	PriceTier        int64        `json:"price_tier"`
	Attachments      []Attachment `json:"attachments"`
	Price            UnitPrice    `json:"price"`
}

type GuildProductPurchase struct {
	ListingID   snowflake.ID `json:"listing_id"`
	ProductName string       `json:"product_name"`
}

type GuildProfile struct {
	ID                       snowflake.ID                           `json:"id"`
	Name                     string                                 `json:"name"`
	IconHash                 *string                                `json:"icon_hash"`
	MemberCount              int64                                  `json:"member_count"`
	OnlineCount              int64                                  `json:"online_count"`
	Description              string                                 `json:"description"`
	BrandColorPrimary        string                                 `json:"brand_color_primary,omitzero"`
	BannerHash               *string                                `json:"banner_hash"`
	GameApplicationIDs       []snowflake.ID                         `json:"game_application_ids"`
	GameActivity             map[snowflake.ID]DiscoveryGameActivity `json:"game_activity"`
	Tag                      *string                                `json:"tag"`
	Badge                    GuildBadgeType                         `json:"badge"`
	BadgeColorPrimary        string                                 `json:"badge_color_primary"`
	BadgeColorSecondary      string                                 `json:"badge_color_secondary"`
	BadgeHash                string                                 `json:"badge_hash"`
	Traits                   []GuildTrait                           `json:"traits"`
	Features                 []string                               `json:"features"`
	Visibility               GuildVisibility                        `json:"visibility"`
	CustomBannerHash         *string                                `json:"custom_banner_hash"`
	PremiumSubscriptionCount int64                                  `json:"premium_subscription_count"`
	PremiumTier              PremiumTier                            `json:"premium_tier"`
}

type GuildRoleCreate struct {
	GuildID snowflake.ID `json:"guild_id"`
	Role    Role         `json:"role"`
}

type GuildRoleDelete struct {
	GuildID snowflake.ID `json:"guild_id"`
	RoleID  snowflake.ID `json:"role_id"`
}

type GuildRoleUpdate struct {
	GuildID snowflake.ID `json:"guild_id"`
	Role    Role         `json:"role"`
}

type GuildScheduledEvent struct {
	ID                            snowflake.ID                  `json:"id"`
	GuildID                       snowflake.ID                  `json:"guild_id"`
	ChannelID                     *snowflake.ID                 `json:"channel_id"`
	CreatorID                     *snowflake.ID                 `json:"creator_id,omitempty"`
	Creator                       PartialUser                   `json:"creator,omitzero"`
	Name                          string                        `json:"name"`
	Description                   *string                       `json:"description,omitempty"`
	ScheduledStartTime            time.Time                     `json:"scheduled_start_time"`
	ScheduledEndTime              *time.Time                    `json:"scheduled_end_time"`
	AutoStart                     bool                          `json:"auto_start,omitzero"`
	PrivacyLevel                  PrivacyLevel                  `json:"privacy_level"`
	Status                        GuildScheduledEventStatus     `json:"status"`
	EntityType                    GuildScheduledEventEntityType `json:"entity_type"`
	EntityID                      *snowflake.ID                 `json:"entity_id"`
	EntityMetadata                json.RawMessage               `json:"entity_metadata"`
	UserCount                     int64                         `json:"user_count,omitzero"`
	Image                         *string                       `json:"image,omitempty"`
	RecurrenceRule                json.RawMessage               `json:"recurrence_rule"`
	GuildScheduledEventExceptions []json.RawMessage             `json:"guild_scheduled_event_exceptions"`
}

type GuildScheduledEventCreateGuildScheduledEventParams struct {
	IsCanceled         *bool      `json:"is_canceled,omitempty"`
	ScheduledStartTime *time.Time `json:"scheduled_start_time,omitempty"`
	ScheduledEndTime   *time.Time `json:"scheduled_end_time,omitempty"`
}

type GuildScheduledEventException struct {
	EventID            snowflake.ID `json:"event_id"`
	EventExceptionID   snowflake.ID `json:"event_exception_id"`
	IsCanceled         bool         `json:"is_canceled"`
	ScheduledStartTime *time.Time   `json:"scheduled_start_time"`
	ScheduledEndTime   *time.Time   `json:"scheduled_end_time"`
}

type GuildScheduledEventExceptionsDelete struct {
	GuildID snowflake.ID `json:"guild_id"`
	EventID snowflake.ID `json:"event_id"`
}

type GuildScheduledEventRecurrenceRule struct {
	Start      time.Time                                  `json:"start"`
	End        *time.Time                                 `json:"end"`
	Frequency  GuildScheduledEventRecurrenceRuleFrequency `json:"frequency"`
	Interval   int64                                      `json:"interval"`
	ByWeekday  []int64                                    `json:"by_weekday"`
	ByNWeekday []json.RawMessage                          `json:"by_n_weekday"`
	ByMonth    []int64                                    `json:"by_month"`
	ByMonthDay []int64                                    `json:"by_month_day"`
	ByYearDay  []int64                                    `json:"by_year_day"`
	Count      *int64                                     `json:"count"`
}

type GuildScheduledEventRecurrenceRuleNWeekday struct {
	N   int64                                    `json:"n"`
	Day GuildScheduledEventRecurrenceRuleWeekday `json:"day"`
}

type GuildScheduledEventUpdate struct {
	GuildScheduledEventID         snowflake.ID                  `json:"guild_scheduled_event_id"`
	GuildScheduledEventEntityType GuildScheduledEventEntityType `json:"guild_scheduled_event_entity_type"`
	GuildID                       snowflake.ID                  `json:"guild_id"`
	ChannelID                     snowflake.ID                  `json:"channel_id,omitzero"`
	ChannelName                   string                        `json:"channel_name,omitzero"`
	ChannelType                   ChannelType                   `json:"channel_type,omitzero"`
	GuildScheduledEventEntityID   snowflake.ID                  `json:"guild_scheduled_event_entity_id,omitzero"`
	GuildScheduledEventEntityName string                        `json:"guild_scheduled_event_entity_name,omitzero"`
	GuildName                     string                        `json:"guild_name,omitzero"`
	GuildIcon                     *string                       `json:"guild_icon,omitempty"`
}

type GuildScheduledEventUser struct {
	GuildScheduledEventID          snowflake.ID                    `json:"guild_scheduled_event_id"`
	GuildScheduledEventExceptionID snowflake.ID                    `json:"guild_scheduled_event_exception_id,omitzero"`
	Response                       GuildScheduledEventUserResponse `json:"response"`
	UserID                         snowflake.ID                    `json:"user_id"`
	User                           PartialUser                     `json:"user,omitzero"`
	Member                         GuildMember                     `json:"member,omitzero"`
}

type GuildSettings struct {
	Channels                          map[string]ChannelSettings              `json:"channels"`
	HubProgress                       json.RawMessage                         `json:"hub_progress"`
	GuildOnboardingProgress           json.RawMessage                         `json:"guild_onboarding_progress"`
	GuildRecentsDismissedAt           json.RawMessage                         `json:"guild_recents_dismissed_at,omitempty"`
	DismissedGuildContent             json.RawMessage                         `json:"dismissed_guild_content"`
	JoinSound                         CustomCallSound                         `json:"join_sound,omitzero"`
	MobileRedesignChannelListSettings ChannelListSettings                     `json:"mobile_redesign_channel_list_settings,omitzero"`
	DisableRaidAlertPush              bool                                    `json:"disable_raid_alert_push"`
	DisableRaidAlertNag               bool                                    `json:"disable_raid_alert_nag"`
	CustomNotificationSoundConfig     CustomNotificationSoundConfig           `json:"custom_notification_sound_config,omitzero"`
	LeaderboardsDisabled              bool                                    `json:"leaderboards_disabled"`
	GuildDismissibleContentStates     map[string]GuildDismissibleContentState `json:"guild_dismissible_content_states"`
	GuildThemeSourcePreference        json.RawMessage                         `json:"guild_theme_source_preference"`
}

type GuildSoundboardSoundDelete struct {
	GuildID snowflake.ID `json:"guild_id"`
	SoundID snowflake.ID `json:"sound_id"`
}

type GuildSoundboardSoundsUpdate struct {
	GuildID          snowflake.ID      `json:"guild_id"`
	SoundboardSounds []SoundboardSound `json:"soundboard_sounds"`
}

type GuildStatus struct {
	Guild RPCGuild `json:"guild"`
}

type GuildStatusSubscriptionArguments struct {
	GuildID snowflake.ID `json:"guild_id"`
}

type GuildStickersUpdate struct {
	GuildID  snowflake.ID `json:"guild_id"`
	Stickers []Sticker    `json:"stickers"`
}

type GuildStreamStart struct {
	ChannelID snowflake.ID `json:"channel_id"`
	GuildID   snowflake.ID `json:"guild_id"`
	UserID    snowflake.ID `json:"user_id"`
}

type GuildTemplate struct {
	Code                  string       `json:"code"`
	Name                  string       `json:"name"`
	Description           *string      `json:"description"`
	UsageCount            int64        `json:"usage_count"`
	CreatorID             snowflake.ID `json:"creator_id"`
	Creator               PartialUser  `json:"creator"`
	CreatedAt             time.Time    `json:"created_at"`
	UpdatedAt             time.Time    `json:"updated_at"`
	SourceGuildID         snowflake.ID `json:"source_guild_id"`
	SerializedSourceGuild PartialGuild `json:"serialized_source_guild"`
	IsDirty               *bool        `json:"is_dirty"`
}

type GuildTemplateBrowserArguments struct {
	Code string `json:"code"`
}

type GuildTemplateBrowserResponse struct {
	Guildtemplate RPCGuildTemplate `json:"guildTemplate"`
	Code          string           `json:"code"`
}

type GuildTrait struct {
	EmojiID       *snowflake.ID `json:"emoji_id"`
	EmojiName     *string       `json:"emoji_name"`
	EmojiAnimated bool          `json:"emoji_animated"`
	Label         string        `json:"label"`
	Position      int64         `json:"position"`
}

type GuildWidget struct {
	ID            snowflake.ID      `json:"id"`
	Name          string            `json:"name"`
	InstantInvite *string           `json:"instant_invite"`
	PresenceCount int64             `json:"presence_count"`
	Channels      []json.RawMessage `json:"channels"`
	Members       []json.RawMessage `json:"members"`
}

type GuildWidgetChannel struct {
	ID       snowflake.ID `json:"id"`
	Name     string       `json:"name"`
	Position int64        `json:"position"`
}

type GuildWidgetMember struct {
	ID        snowflake.ID    `json:"id"`
	Username  string          `json:"username"`
	AvatarURL string          `json:"avatar_url"`
	Status    StatusType      `json:"status"`
	Activity  json.RawMessage `json:"activity,omitempty"`
	ChannelID snowflake.ID    `json:"channel_id,omitzero"`
	Deaf      bool            `json:"deaf,omitzero"`
	Mute      bool            `json:"mute,omitzero"`
	SelfDeaf  bool            `json:"self_deaf,omitzero"`
	SelfMute  bool            `json:"self_mute,omitzero"`
	Suppress  bool            `json:"suppress,omitzero"`
}

type GuildWidgetMemberActivity struct {
	Name string `json:"name"`
}

type GuildWidgetSettings struct {
	Enabled   bool          `json:"enabled"`
	ChannelID *snowflake.ID `json:"channel_id"`
}

type Handshake struct {
	V        int64  `json:"v"`
	ClientID string `json:"client_id"`
}

type Harvest struct {
	HarvestID       snowflake.ID      `json:"harvest_id"`
	UserID          snowflake.ID      `json:"user_id"`
	Email           string            `json:"email"`
	State           HarvestState      `json:"state"`
	Status          HarvestStatus     `json:"status"`
	CreatedAt       time.Time         `json:"created_at"`
	CompletedAt     *time.Time        `json:"completed_at"`
	PolledAt        *time.Time        `json:"polled_at"`
	Backends        map[string]string `json:"backends"`
	UpdatedAt       time.Time         `json:"updated_at"`
	ShadowRun       bool              `json:"shadow_run"`
	HarvestMetadata HarvestMetadata   `json:"harvest_metadata"`
}

type HarvestMetadata struct {
	UserIsStaff     bool             `json:"user_is_staff"`
	SlaEmailSent    bool             `json:"sla_email_sent"`
	BypassCooldown  bool             `json:"bypass_cooldown"`
	IsProvisional   bool             `json:"is_provisional,omitzero"`
	BackendAttempts map[string]int64 `json:"backend_attempts,omitempty"`
}

type HeardSoundFrecency struct {
	HeardSounds map[string]FrecencyItem `json:"heard_sounds"`
}

type Heartbeat struct {
	T      int64 `json:"t"`
	SeqAck int64 `json:"seq_ack,omitzero"`
}

type Hello struct {
	Trace             []string `json:"_trace"`
	HeartbeatInterval int64    `json:"heartbeat_interval"`
}

type HeroShopBlock struct {
	Type                   ShopBlockType  `json:"type"`
	CategorySKUID          snowflake.ID   `json:"category_sku_id"`
	CategoryStoreListingID snowflake.ID   `json:"category_store_listing_id"`
	BannerAsset            AssetConfig    `json:"banner_asset"`
	BannerDisplayConfig    BannerAsset    `json:"banner_display_config"`
	BannerTextColor        string         `json:"banner_text_color,omitzero"`
	HeroLogoURL            string         `json:"hero_logo_url"`
	HeroRiveURL            string         `json:"hero_rive_url,omitzero"`
	LogoDisplayConfig      BannerAsset    `json:"logo_display_config"`
	LogoURL                string         `json:"logo_url"`
	Name                   string         `json:"name"`
	MobileTitle            string         `json:"mobile_title,omitzero"`
	MobileSummary          string         `json:"mobile_summary,omitzero"`
	MobileProductsTitle    string         `json:"mobile_products_title,omitzero"`
	MobileHeroURL          string         `json:"mobile_hero_url,omitzero"`
	MobileHeroAnimatedURL  string         `json:"mobile_hero_animated_url,omitzero"`
	RankedSKUIDs           []snowflake.ID `json:"ranked_sku_ids"`
	Summary                string         `json:"summary"`
	UnpublishedAt          *time.Time     `json:"unpublished_at"`
}

type IconEmoji struct {
	ID   *snowflake.ID `json:"id"`
	Name *string       `json:"name"`
}

type IcymiSettings struct {
	FeedGeneratedAt json.RawMessage `json:"feed_generated_at"`
}

type Identify struct {
	Token          string           `json:"token"`
	Properties     ClientProperties `json:"properties"`
	Compress       bool             `json:"compress,omitzero"`
	LargeThreshold int64            `json:"large_threshold,omitzero"`
	Shard          []int64          `json:"shard,omitempty"`
	Presence       UpdatePresence   `json:"presence,omitzero"`
	Intents        int64            `json:"intents,omitzero"`
	Capabilities   int64            `json:"capabilities,omitzero"`
	ClientState    ClientState      `json:"client_state,omitzero"`
}

type ImmersiveBannerShopBlock struct {
	Type          ShopBlockType `json:"type"`
	Title         string        `json:"title"`
	Body          string        `json:"body"`
	HelpCenterURL string        `json:"help_center_url"`
	TextColor     string        `json:"text_color"`
	EndTime       *time.Time    `json:"end_time"`
	BannerAsset   AssetConfig   `json:"banner_asset"`
}

type InAppFeedbackSettings struct {
	InAppFeedbackStates map[string]InAppFeedbackState `json:"in_app_feedback_states"`
}

type InAppFeedbackState struct {
	LastImpressionTime uint64 `json:"last_impression_time,omitzero"`
	OptOutExpiryTime   uint64 `json:"opt_out_expiry_time,omitzero"`
}

type InboxSettings struct {
	CurrentTab     json.RawMessage `json:"current_tab"`
	ViewedTutorial bool            `json:"viewed_tutorial"`
}

type IncomingPayload struct {
	Cmd   string          `json:"cmd"`
	Data  json.RawMessage `json:"data"`
	Nonce *string         `json:"nonce"`
	Evt   *string         `json:"evt"`
}

type Init struct {
	EncodedPublicKey string `json:"encoded_public_key"`
}

type InitiateImageUploadResponse struct {
	ImageURL string `json:"image_url"`
}

type Integration struct {
	ID                      snowflake.ID                        `json:"id"`
	Name                    string                              `json:"name"`
	Type                    IntegrationType                     `json:"type"`
	Enabled                 bool                                `json:"enabled"`
	Account                 IntegrationAccount                  `json:"account"`
	Syncing                 bool                                `json:"syncing,omitzero"`
	RoleID                  *snowflake.ID                       `json:"role_id,omitempty"`
	EnableEmoticons         bool                                `json:"enable_emoticons,omitzero"`
	ExpireBehavior          IntegrationExpireBehavior           `json:"expire_behavior,omitzero"`
	ExpireGracePeriod       int64                               `json:"expire_grace_period,omitzero"`
	SyncedAt                time.Time                           `json:"synced_at,omitzero"`
	SubscriberCount         int64                               `json:"subscriber_count,omitzero"`
	Revoked                 bool                                `json:"revoked,omitzero"`
	Application             IntegrationApplication              `json:"application,omitzero"`
	Scopes                  []string                            `json:"scopes,omitempty"`
	RoleConnectionsMetadata []ApplicationRoleConnectionMetadata `json:"role_connections_metadata"`
	User                    PartialUser                         `json:"user,omitzero"`
}

type IntegrationAccount struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type IntegrationApplication struct {
	ID                             snowflake.ID     `json:"id"`
	Name                           string           `json:"name"`
	Description                    string           `json:"description"`
	Icon                           *string          `json:"icon"`
	CoverImage                     string           `json:"cover_image,omitzero"`
	Splash                         string           `json:"splash,omitzero"`
	Type                           *ApplicationType `json:"type"`
	PrimarySKUID                   snowflake.ID     `json:"primary_sku_id,omitzero"`
	Bot                            PartialUser      `json:"bot,omitzero"`
	DeeplinkURI                    *string          `json:"deeplink_uri,omitempty"`
	ThirdPartySKUs                 []ApplicationSKU `json:"third_party_skus,omitempty"`
	RoleConnectionsVerificationURL *string          `json:"role_connections_verification_url,omitempty"`
	ConnectionEntrypointURL        string           `json:"connection_entrypoint_url,omitzero"`
	IsVerified                     bool             `json:"is_verified"`
	IsDiscoverable                 bool             `json:"is_discoverable"`
	IsMonetized                    bool             `json:"is_monetized"`
	ParentID                       snowflake.ID     `json:"parent_id,omitzero"`
}

type IntegrationCreateExtraFields struct {
	GuildID snowflake.ID `json:"guild_id"`
}

type IntegrationDelete struct {
	ID            snowflake.ID `json:"id"`
	GuildID       snowflake.ID `json:"guild_id"`
	ApplicationID snowflake.ID `json:"application_id,omitzero"`
}

type IntegrationGuild struct {
	ID   snowflake.ID `json:"id"`
	Name string       `json:"name"`
	Icon *string      `json:"icon"`
}

type IntegrationUpdateExtraFields struct {
	GuildID snowflake.ID `json:"guild_id"`
}

type Interaction struct {
	Version                      ApplicationInteractionsVersion `json:"version"`
	ID                           snowflake.ID                   `json:"id"`
	ApplicationID                snowflake.ID                   `json:"application_id"`
	Type                         InteractionType                `json:"type"`
	Token                        string                         `json:"token"`
	Data                         json.RawMessage                `json:"data,omitempty"`
	Guild                        InteractionGuild               `json:"guild,omitzero"`
	GuildID                      snowflake.ID                   `json:"guild_id,omitzero"`
	GuildLocale                  string                         `json:"guild_locale,omitzero"`
	Channel                      Channel                        `json:"channel,omitzero"`
	ChannelID                    snowflake.ID                   `json:"channel_id,omitzero"`
	Member                       GuildMember                    `json:"member,omitzero"`
	User                         PartialUser                    `json:"user,omitzero"`
	Locale                       string                         `json:"locale,omitzero"`
	Message                      Message                        `json:"message,omitzero"`
	AppPermissions               string                         `json:"app_permissions"`
	Entitlements                 []Entitlement                  `json:"entitlements"`
	EntitlementSKUIDs            []snowflake.ID                 `json:"entitlement_sku_ids,omitempty"`
	AuthorizingIntegrationOwners map[int64]snowflake.ID         `json:"authorizing_integration_owners"`
	Context                      InteractionContextType         `json:"context,omitzero"`
	AttachmentSizeLimit          int64                          `json:"attachment_size_limit"`
}

type InteractionCallback struct {
	ID                       snowflake.ID    `json:"id"`
	Type                     InteractionType `json:"type"`
	ResponseMessageID        snowflake.ID    `json:"response_message_id,omitzero"`
	ResponseMessageLoading   bool            `json:"response_message_loading,omitzero"`
	ResponseMessageEphemeral bool            `json:"response_message_ephemeral,omitzero"`
	ActivityInstanceID       string          `json:"activity_instance_id,omitzero"`
	ChannelID                snowflake.ID    `json:"channel_id,omitzero"`
	GuildID                  snowflake.ID    `json:"guild_id,omitzero"`
}

type InteractionCallbackActivityInstanceResource struct {
	ID string `json:"id"`
}

type InteractionCallbackAutocompleteData struct {
	Choices []ApplicationCommandOptionChoice `json:"choices"`
}

type InteractionCallbackIframeModalData struct {
	IframePath string          `json:"iframe_path"`
	Title      string          `json:"title"`
	ModalSize  IframeModalSize `json:"modal_size"`
	CustomID   string          `json:"custom_id"`
}

type InteractionCallbackMessageData struct {
	Content         string          `json:"content,omitzero"`
	TTS             bool            `json:"tts,omitzero"`
	Embeds          []Embed         `json:"embeds,omitempty"`
	AllowedMentions AllowedMentions `json:"allowed_mentions,omitzero"`
	Components      Components      `json:"components,omitzero"`
	Flags           MessageFlags    `json:"flags,omitzero"`
	Attachments     []Attachment    `json:"attachments,omitempty"`
	Poll            PollCreate      `json:"poll,omitzero"`
}

type InteractionCallbackModalData struct {
	CustomID   string     `json:"custom_id"`
	Title      string     `json:"title"`
	Components Components `json:"components"`
}

type InteractionCallbackResource struct {
	Type             InteractionCallbackType                     `json:"type"`
	ActivityInstance InteractionCallbackActivityInstanceResource `json:"activity_instance,omitzero"`
	Message          Message                                     `json:"message,omitzero"`
}

type InteractionCallbackSocialLayerSKUPurchaseEligibilityData struct {
	Eligible bool `json:"eligible"`
}

type InteractionCreate struct {
	ID    snowflake.ID `json:"id"`
	Nonce string       `json:"nonce,omitzero"`
}

type InteractionFailure struct {
	ID         snowflake.ID             `json:"id"`
	Nonce      string                   `json:"nonce,omitzero"`
	ReasonCode InteractionFailureReason `json:"reason_code"`
}

type InteractionGuild struct {
	ID       snowflake.ID `json:"id"`
	Features []string     `json:"features"`
	Locale   string       `json:"locale"`
}

type InteractionIframeModalCreate struct {
	ID          snowflake.ID           `json:"id"`
	ChannelID   snowflake.ID           `json:"channel_id"`
	CustomID    string                 `json:"custom_id"`
	Application IntegrationApplication `json:"application"`
	Title       string                 `json:"title"`
	IframePath  string                 `json:"iframe_path"`
	ModalSize   IframeModalSize        `json:"modal_size"`
	Nonce       *string                `json:"nonce"`
}

type InteractionModalCreate struct {
	ID          snowflake.ID           `json:"id"`
	ChannelID   snowflake.ID           `json:"channel_id"`
	CustomID    string                 `json:"custom_id"`
	Application IntegrationApplication `json:"application"`
	Title       string                 `json:"title"`
	Components  Components             `json:"components"`
	Nonce       *string                `json:"nonce"`
	Resolved    ResolvedData           `json:"resolved,omitzero"`
}

type InteractionOption struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type InteractionResponse struct {
	Type InteractionCallbackType `json:"type"`
	Data json.RawMessage         `json:"data,omitempty"`
}

type InteractionResponsePreviewImage struct {
	Height int64  `json:"height"`
	URL    string `json:"url"`
	Width  int64  `json:"width"`
}

type InteractionSuccess struct {
	ID    snowflake.ID `json:"id"`
	Nonce *string      `json:"nonce"`
}

type Invite struct {
	Code                     string               `json:"code"`
	Type                     InviteType           `json:"type"`
	Channel                  *PartialChannel      `json:"channel"`
	GuildID                  snowflake.ID         `json:"guild_id,omitzero"`
	Guild                    InviteGuild          `json:"guild,omitzero"`
	Profile                  GuildProfile         `json:"profile,omitzero"`
	Inviter                  PartialUser          `json:"inviter,omitzero"`
	Flags                    InviteFlags          `json:"flags,omitzero"`
	TargetType               InviteTargetType     `json:"target_type,omitzero"`
	TargetUser               PartialUser          `json:"target_user,omitzero"`
	TargetApplication        PartialApplication   `json:"target_application,omitzero"`
	Roles                    []PartialRole        `json:"roles,omitempty"`
	ApproximateMemberCount   int64                `json:"approximate_member_count,omitzero"`
	ApproximatePresenceCount int64                `json:"approximate_presence_count,omitzero"`
	ExpiresAt                *time.Time           `json:"expires_at"`
	GuildScheduledEvent      GuildScheduledEvent  `json:"guild_scheduled_event,omitzero"`
	NewMember                bool                 `json:"new_member,omitzero"`
	ShowVerificationForm     bool                 `json:"show_verification_form,omitzero"`
	IsNicknameChangeable     bool                 `json:"is_nickname_changeable,omitzero"`
	TargetUsersJobStatus     InviteTargetUsersJob `json:"target_users_job_status,omitzero"`
}

type InviteBrowserArguments struct {
	Code string `json:"code"`
}

type InviteBrowserResponse struct {
	Invite Invite `json:"invite"`
	Code   string `json:"code"`
}

type InviteCreate struct {
	Code              string             `json:"code"`
	Type              InviteType         `json:"type"`
	ChannelID         snowflake.ID       `json:"channel_id"`
	GuildID           snowflake.ID       `json:"guild_id"`
	Inviter           PartialUser        `json:"inviter,omitzero"`
	TargetType        InviteTargetType   `json:"target_type,omitzero"`
	TargetUser        PartialUser        `json:"target_user,omitzero"`
	TargetApplication PartialApplication `json:"target_application,omitzero"`
	RoleIDs           []snowflake.ID     `json:"role_ids,omitempty"`
	ExpiresAt         *time.Time         `json:"expires_at"`
	CreatedAt         time.Time          `json:"created_at"`
	Uses              int64              `json:"uses"`
	MaxUses           int64              `json:"max_uses"`
	MaxAge            int64              `json:"max_age"`
	Temporary         bool               `json:"temporary"`
}

type InviteDelete struct {
	Code      string       `json:"code"`
	ChannelID snowflake.ID `json:"channel_id"`
	GuildID   snowflake.ID `json:"guild_id"`
}

type InviteGuild struct {
	ID                       snowflake.ID      `json:"id"`
	Name                     string            `json:"name"`
	Icon                     *string           `json:"icon"`
	Description              *string           `json:"description"`
	Banner                   *string           `json:"banner"`
	Splash                   *string           `json:"splash"`
	VerificationLevel        VerificationLevel `json:"verification_level"`
	Features                 []string          `json:"features"`
	VanityURLCode            *string           `json:"vanity_url_code"`
	PremiumSubscriptionCount int64             `json:"premium_subscription_count,omitzero"`
	PremiumTier              PremiumTier       `json:"premium_tier"`
	NSFW                     bool              `json:"nsfw"`
	NSFWLevel                NSFWLevel         `json:"nsfw_level"`
}

type InviteMetadata struct {
	Uses      int64     `json:"uses,omitzero"`
	MaxUses   int64     `json:"max_uses,omitzero"`
	MaxAge    int64     `json:"max_age,omitzero"`
	Temporary bool      `json:"temporary,omitzero"`
	CreatedAt time.Time `json:"created_at"`
}

type InviteTargetUsersJob struct {
	Status         InviteTargetUsersJobStatus `json:"status"`
	TotalUsers     int64                      `json:"total_users"`
	ProcessedUsers int64                      `json:"processed_users"`
	CreatedAt      *time.Time                 `json:"created_at"`
	CompletedAt    *time.Time                 `json:"completed_at"`
	ErrorMessage   *string                    `json:"error_message"`
}

type InviteUserEmbeddedArguments struct {
	UserID  snowflake.ID `json:"user_id"`
	Content string       `json:"content,omitzero"`
}

type Invoice struct {
	ID                      snowflake.ID                `json:"id"`
	Status                  InvoiceStatus               `json:"status,omitzero"`
	Currency                string                      `json:"currency"`
	Subtotal                int64                       `json:"subtotal"`
	Tax                     int64                       `json:"tax"`
	Total                   int64                       `json:"total"`
	TaxInclusive            bool                        `json:"tax_inclusive"`
	Items                   []InvoiceItem               `json:"items"`
	SubscriptionPeriodStart *time.Time                  `json:"subscription_period_start"`
	SubscriptionPeriodEnd   *time.Time                  `json:"subscription_period_end"`
	AppliedDiscountIDs      []snowflake.ID              `json:"applied_discount_ids,omitempty"`
	AppliedUserDiscounts    map[snowflake.ID]*time.Time `json:"applied_user_discounts,omitempty"`
	OrbsReward              int64                       `json:"orbs_reward,omitzero"`
}

type InvoiceDiscount struct {
	Type   InvoiceDiscountType `json:"type"`
	Amount int64               `json:"amount"`
}

type InvoiceItem struct {
	ID                    snowflake.ID               `json:"id"`
	Quantity              int64                      `json:"quantity"`
	Amount                int64                      `json:"amount"`
	Proration             bool                       `json:"proration"`
	SubscriptionPlanID    *snowflake.ID              `json:"subscription_plan_id"`
	SubscriptionPlanPrice *int64                     `json:"subscription_plan_price"`
	Discounts             []InvoiceDiscount          `json:"discounts"`
	SKUID                 *snowflake.ID              `json:"sku_id"`
	UnitPrice             *UnitPrice                 `json:"unit_price"`
	TenantMetadata        map[string]json.RawMessage `json:"tenant_metadata,omitempty"`
}

type IosQuestCtaConfig struct {
	IosAppID string `json:"ios_app_id"`
}

type ItemAssets struct {
	StaticImageURL   string `json:"static_image_url"`
	AnimatedImageURL string `json:"animated_image_url"`
	VideoURL         string `json:"video_url,omitzero"`
}

type JoinGuildQuery struct {
	Lurker               bool   `json:"lurker,omitzero"`
	SessionID            string `json:"session_id,omitzero"`
	Location             string `json:"location,omitzero"`
	RecommendationLoadID string `json:"recommendation_load_id,omitzero"`
}

type JoinHypesquadOnlineParams struct {
	HouseID HypesquadHouse `json:"house_id"`
}

type JoinOrCreateLobbyParams struct {
	Secret             string            `json:"secret"`
	LobbyMetadata      map[string]string `json:"lobby_metadata,omitempty"`
	MemberMetadata     map[string]string `json:"member_metadata,omitempty"`
	IdleTimeoutSeconds int64             `json:"idle_timeout_seconds,omitzero"`
	Flags              LobbyFlags        `json:"flags,omitzero"`
}

type JoinStudentHubParams struct {
	Email   string       `json:"email"`
	GuildID snowflake.ID `json:"guild_id"`
	Code    string       `json:"code"`
}

type JoinStudentHubResponse struct {
	Joined bool  `json:"joined"`
	Guild  Guild `json:"guild,omitzero"`
}

type JoinStudentHubWaitlistParams struct {
	Email  string `json:"email"`
	School string `json:"school"`
}

type JoinStudentHubWaitlistResponse struct {
	Email       string       `json:"email"`
	EmailDomain string       `json:"email_domain"`
	School      string       `json:"school"`
	UserID      snowflake.ID `json:"user_id"`
}

type JoinThreadQuery struct {
	Location string `json:"location,omitzero"`
}

type KeywordFilterSettings struct {
	Profanity     bool `json:"profanity,omitzero"`
	SexualContent bool `json:"sexual_content,omitzero"`
	Slurs         bool `json:"slurs,omitzero"`
}

type Label struct {
	Type        ComponentType `json:"type"`
	ID          int64         `json:"id,omitzero"`
	Label       string        `json:"label"`
	Description string        `json:"description,omitzero"`
	Component   AnyComponent  `json:"component"`
}

type LastMessages struct {
	GuildID  snowflake.ID `json:"guild_id"`
	Messages []Message    `json:"messages"`
}

type LaunchEmbeddedActivityParams struct {
	SessionID string `json:"session_id"`
}

type LeaveGuildParams struct {
	Lurking bool `json:"lurking,omitzero"`
}

type LeaveThreadQuery struct {
	Location string `json:"location,omitzero"`
}

type LinkedAccount struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type LinkedApplication struct {
	ID   snowflake.ID          `json:"id"`
	Type LinkedApplicationType `json:"type"`
}

type LinkedLobby struct {
	ApplicationID                   snowflake.ID `json:"application_id"`
	LobbyID                         snowflake.ID `json:"lobby_id"`
	LinkedBy                        snowflake.ID `json:"linked_by"`
	LinkedAt                        time.Time    `json:"linked_at"`
	RequireApplicationAuthorization bool         `json:"require_application_authorization"`
}

type LinkedUser struct {
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	LinkStatus  LinkStatus   `json:"link_status"`
	LinkType    LinkType     `json:"link_type"`
	RequestorID snowflake.ID `json:"requestor_id"`
	UserID      snowflake.ID `json:"user_id"`
}

type ListActiveThreadsResponse struct {
	Threads []Channel      `json:"threads"`
	Members []ThreadMember `json:"members"`
}

type ListApplicationDirectoryCategoriesQuery struct {
	Locale string `json:"locale,omitzero"`
}

type ListApplicationDirectoryCollectionsQuery struct {
	Surface     ApplicationDirectoryRequestSurface  `json:"surface,omitzero"`
	ActiveState ApplicationDirectoryActiveStateType `json:"active_state,omitzero"`
	Platform    ApplicationDirectoryPlatformFlags   `json:"platform,omitzero"`
	Locale      string                              `json:"locale,omitzero"`
	Cache       bool                                `json:"cache,omitzero"`
}

type ListApplicationDirectorySimilarApplicationsQuery struct {
	GuildID string `json:"guild_id,omitzero"`
	Page    int64  `json:"page,omitzero"`
	Locale  string `json:"locale,omitzero"`
}

type ListApplicationDirectorySimilarApplicationsResponse struct {
	Applications []Application `json:"applications"`
	NumPages     int64         `json:"num_pages"`
	LoadID       string        `json:"load_id"`
}

type ListApplicationDisclosuresResponse struct {
	Disclosures      []int64 `json:"disclosures"`
	AckedDisclosures []int64 `json:"acked_disclosures"`
	AllAcked         bool    `json:"all_acked"`
}

type ListApplicationEmojisResponse struct {
	Items []Emoji `json:"items"`
}

type ListApplicationExternalIdentityProviderConfigurationsParams struct {
	Clients       []ApplicationExternalIdentityProviderClient `json:"clients"`
	OidcIssuerURL *string                                     `json:"oidc_issuer_url"`
}

type ListApplicationPublishedStoreListingsQuery struct {
	ApplicationID snowflake.ID `json:"application_id"`
	GuildID       snowflake.ID `json:"guild_id,omitzero"`
	CountryCode   string       `json:"country_code,omitzero"`
	Localize      bool         `json:"localize,omitzero"`
}

type ListApplicationsQuery struct {
	WithTeamApplications bool `json:"with_team_applications,omitzero"`
}

type ListApplicationsWithAssetsQuery struct {
	WithTeamApplications bool `json:"with_team_applications,omitzero"`
}

type ListApplicationsWithAssetsResponse struct {
	Applications []Application                       `json:"applications"`
	Assets       map[snowflake.ID][]ApplicationAsset `json:"assets"`
}

type ListAppliedPremiumGuildSubscriptionsQuery struct {
	Paused bool `json:"paused,omitzero"`
}

type ListBogoPromotionsQuery struct {
	Locale string `json:"locale,omitzero"`
}

type ListBulkApplicationIdentitiesParams struct {
	UserIDs []snowflake.ID `json:"user_ids"`
}

type ListBulkApplicationPrimaryStoreListingQuery struct {
	ApplicationIDs []snowflake.ID `json:"application_ids"`
	CountryCode    string         `json:"country_code,omitzero"`
	Localize       bool           `json:"localize,omitzero"`
}

type ListBulkPublishedSubscriptionPlansQuery struct {
	SKUIDs             []snowflake.ID `json:"sku_ids"`
	IncludeUnpublished bool           `json:"include_unpublished,omitzero"`
	RevenueSurface     RevenueSurface `json:"revenue_surface,omitzero"`
	CountryCode        string         `json:"country_code,omitzero"`
	PaymentSourceID    snowflake.ID   `json:"payment_source_id,omitzero"`
}

type ListChannelAffinitiesResponse struct {
	ChannelAffinities []ChannelAffinity `json:"channel_affinities"`
}

type ListClaimedPromotionsQuery struct {
	Locale string `json:"locale,omitzero"`
}

type ListClaimedQuestsResponse struct {
	Quests []ClaimedQuest `json:"quests"`
}

type ListCollectiblesCategoriesQuery struct {
	CountryCode          string                     `json:"country_code,omitzero"`
	IncludeBundles       bool                       `json:"include_bundles,omitzero"`
	IncludeDynamicBlocks bool                       `json:"include_dynamic_blocks,omitzero"`
	IncludeUnpublished   bool                       `json:"include_unpublished,omitzero"`
	NoCache              bool                       `json:"no_cache,omitzero"`
	PaymentGateway       PaymentGateway             `json:"payment_gateway,omitzero"`
	ShopHomeConfig       CollectiblesShopHomeConfig `json:"shop_home_config,omitzero"`
	SkipNumCategories    int64                      `json:"skip_num_categories,omitzero"`
	Tab                  CollectiblesShopHomeTab    `json:"tab,omitzero"`
	VariantsReturnStyle  VariantsReturnStyle        `json:"variants_return_style,omitzero"`
}

type ListCollectiblesCategoriesV2Query struct {
	CountryCode          string                     `json:"country_code,omitzero"`
	IncludeBundles       bool                       `json:"include_bundles,omitzero"`
	IncludeDynamicBlocks bool                       `json:"include_dynamic_blocks,omitzero"`
	IncludeUnpublished   bool                       `json:"include_unpublished,omitzero"`
	NoCache              bool                       `json:"no_cache,omitzero"`
	PaymentGateway       PaymentGateway             `json:"payment_gateway,omitzero"`
	ShopHomeConfig       CollectiblesShopHomeConfig `json:"shop_home_config,omitzero"`
	SkipNumCategories    int64                      `json:"skip_num_categories,omitzero"`
	Tab                  CollectiblesShopHomeTab    `json:"tab,omitzero"`
	VariantsReturnStyle  VariantsReturnStyle        `json:"variants_return_style,omitzero"`
}

type ListCollectiblesCategoriesV2Response struct {
	Categories    []CollectibleCategory     `json:"categories"`
	Collections   []StorefrontCollection    `json:"collections"`
	UserDiscounts []CollectibleUserDiscount `json:"user_discounts,omitempty"`
}

type ListConsoleDevicesResponse struct {
	Devices []ConsoleDevice `json:"devices"`
}

type ListConversationSummariesResponse struct {
	Summaries []ConversationSummary `json:"summaries"`
}

type ListCurrentUserQuestsResponse struct {
	Quests                      []Quest        `json:"quests"`
	ExcludedQuests              []PartialQuest `json:"excluded_quests"`
	QuestEnrollmentBlockedUntil *time.Time     `json:"quest_enrollment_blocked_until"`
}

type ListDMMessagesQuery struct {
	Limit int64 `json:"limit"`
}

type ListDetectableGameExclusionsResponse struct {
	Executables []string `json:"executables"`
	Patterns    []string `json:"patterns"`
}

type ListDiscoverableGuildsQuery struct {
	GuildIDs       []snowflake.ID `json:"guild_ids,omitempty"`
	ApplicationIDs []snowflake.ID `json:"application_ids,omitempty"`
	Categories     []int64        `json:"categories,omitempty"`
	Limit          int64          `json:"limit,omitzero"`
	Offset         int64          `json:"offset,omitzero"`
}

type ListDiscoverableGuildsResponse struct {
	Guilds []DiscoverableGuild `json:"guilds"`
	Total  int64               `json:"total"`
	Limit  int64               `json:"limit"`
	Offset int64               `json:"offset"`
}

type ListDiscoveryCategoriesQuery struct {
	Locale      string `json:"locale,omitzero"`
	PrimaryOnly bool   `json:"primary_only,omitzero"`
}

type ListEligibleApplicationSubscriptionGuildsQuery struct {
	ApplicationID snowflake.ID `json:"application_id"`
	SKUID         snowflake.ID `json:"sku_id,omitzero"`
}

type ListEmbeddedActivitiesQuery struct {
	GuildID snowflake.ID `json:"guild_id,omitzero"`
}

type ListEmbeddedActivitiesResponse struct {
	Activities   []EmbeddedActivityConfig            `json:"activities"`
	Applications []PartialApplication                `json:"applications"`
	Assets       map[snowflake.ID][]ApplicationAsset `json:"assets"`
}

type ListEmbeddedActivityInstancesResponse struct {
	ApplicationID snowflake.ID             `json:"application_id"`
	InstanceID    string                   `json:"instance_id"`
	LaunchID      string                   `json:"launch_id"`
	Location      EmbeddedActivityLocation `json:"location"`
	Users         []snowflake.ID           `json:"users"`
}

type ListFeaturedApplicationWidgetConfigsResponse struct {
	ApplicationIDs []snowflake.ID                  `json:"application_ids"`
	Configs        map[snowflake.ID][]WidgetConfig `json:"configs"`
}

type ListGameAnnouncementsQuery struct {
	Limit int64 `json:"limit"`
}

type ListGameAnnouncementsResponse struct {
	GuildID   *snowflake.ID `json:"guild_id"`
	ChannelID *snowflake.ID `json:"channel_id"`
	Messages  []Message     `json:"messages"`
}

type ListGamesQuery struct {
	GameIDs              []snowflake.ID `json:"game_ids"`
	WithSupplementalData bool           `json:"with_supplemental_data,omitzero"`
}

type ListGlobalActivityStatisticsQuery struct {
	WithUsers        bool `json:"with_users,omitzero"`
	WithApplications bool `json:"with_applications,omitzero"`
}

type ListGuildActiveThreadsResponse struct {
	Threads []Channel      `json:"threads"`
	Members []ThreadMember `json:"members"`
}

type ListGuildAffinitiesResponse struct {
	GuildAffinities []GuildAffinity `json:"guild_affinities"`
}

type ListGuildApplicationsQuery struct {
	Type        ApplicationType `json:"type,omitzero"`
	IncludeTeam bool            `json:"include_team,omitzero"`
	ChannelID   snowflake.ID    `json:"channel_id,omitzero"`
}

type ListGuildAuditLogEntriesQuery struct {
	Before     snowflake.ID       `json:"before,omitzero"`
	After      snowflake.ID       `json:"after,omitzero"`
	Limit      int64              `json:"limit,omitzero"`
	UserID     snowflake.ID       `json:"user_id,omitzero"`
	TargetID   snowflake.ID       `json:"target_id,omitzero"`
	ActionType AuditLogActionType `json:"action_type,omitzero"`
}

type ListGuildAuditLogEntriesResponse struct {
	AuditLogEntries      []AuditLogEntry       `json:"audit_log_entries"`
	ApplicationCommands  []ApplicationCommand  `json:"application_commands"`
	AutoModerationRules  []AutomodRule         `json:"auto_moderation_rules"`
	GuildScheduledEvents []GuildScheduledEvent `json:"guild_scheduled_events"`
	Integrations         []PartialIntegration  `json:"integrations"`
	Threads              []Channel             `json:"threads"`
	Users                []PartialUser         `json:"users"`
	Webhooks             []Webhook             `json:"webhooks"`
}

type ListGuildBansQuery struct {
	Before snowflake.ID `json:"before,omitzero"`
	After  snowflake.ID `json:"after,omitzero"`
	Limit  float64      `json:"limit,omitzero"`
}

type ListGuildChannelFollowingGuildSizeByChannelQuery struct {
	Start     time.Time           `json:"start"`
	End       time.Time           `json:"end"`
	Interval  AggregationInterval `json:"interval"`
	ChannelID snowflake.ID        `json:"channel_id"`
}

type ListGuildChannelsQuery struct {
	Permissions      bool `json:"permissions,omitzero"`
	WithCanLinkLobby bool `json:"with_can_link_lobby,omitzero"`
}

type ListGuildEntitlementsQuery struct {
	UserID         snowflake.ID   `json:"user_id,omitzero"`
	SKUIDs         []snowflake.ID `json:"sku_ids,omitempty"`
	GuildID        snowflake.ID   `json:"guild_id,omitzero"`
	ExcludeEnded   bool           `json:"exclude_ended,omitzero"`
	ExcludeDeleted bool           `json:"exclude_deleted,omitzero"`
	Before         snowflake.ID   `json:"before,omitzero"`
	After          snowflake.ID   `json:"after,omitzero"`
	Limit          int64          `json:"limit,omitzero"`
}

type ListGuildJoinRequestsQuery struct {
	Status GuildJoinRequestStatus `json:"status,omitzero"`
	Limit  int64                  `json:"limit,omitzero"`
	Before snowflake.ID           `json:"before,omitzero"`
	After  snowflake.ID           `json:"after,omitzero"`
}

type ListGuildJoinRequestsResponse struct {
	GuildJoinRequests []GuildJoinRequest `json:"guild_join_requests"`
	Total             int64              `json:"total,omitzero"`
}

type ListGuildMembersQuery struct {
	Limit int64        `json:"limit,omitzero"`
	After snowflake.ID `json:"after,omitzero"`
}

type ListGuildMembersSupplementalParams struct {
	UserIDs []snowflake.ID `json:"user_ids"`
}

type ListGuildMembersWithUnusualDMActivityQuery struct {
	Limit int64        `json:"limit,omitzero"`
	After snowflake.ID `json:"after,omitzero"`
}

type ListGuildOnboardingAllowedApplicationsResponse struct {
	ApplicationIDs []snowflake.ID `json:"application_ids"`
}

type ListGuildPowerupsQuery struct {
	IncludeEndsAt bool `json:"include_ends_at,omitzero"`
}

type ListGuildProductListingsResponse struct {
	Listings []GuildProductListing `json:"listings"`
}

type ListGuildRoleSubscriptionListingTemplatesResponse struct {
	IsEligible bool `json:"is_eligible"`
}

type ListGuildRoleSubscriptionTrialsParams struct {
	Trial                  PartialSubscriptionTrial `json:"trial,omitzero"`
	MaxNumActiveTrialUsers *int64                   `json:"max_num_active_trial_users,omitempty"`
}

type ListGuildScheduledEventUsersQuery struct {
	Before              snowflake.ID `json:"before,omitzero"`
	After               snowflake.ID `json:"after,omitzero"`
	Limit               float64      `json:"limit,omitzero"`
	WithMember          bool         `json:"with_member,omitzero"`
	UpgradeResponseType bool         `json:"upgrade_response_type,omitzero"`
}

type ListGuildScheduledEventUsersResponse struct {
	Users []PartialUser `json:"users"`
}

type ListGuildScheduledEventsQuery struct {
	WithUserCount bool `json:"with_user_count,omitzero"`
}

type ListGuildSoundboardSoundsResponse struct {
	Items []SoundboardSound `json:"items"`
}

type ListGuildTopEmojisResponse struct {
	Items []TopEmoji `json:"items"`
}

type ListGuildTopGamesResponse struct {
	TopGames []GameActivity `json:"top_games"`
}

type ListInviteFriendMembersResponse struct {
	FriendMemberIDs []snowflake.ID `json:"friend_member_ids"`
}

type ListJoinedPrivateArchivedThreadsQuery struct {
	Before snowflake.ID `json:"before,omitzero"`
	Limit  int64        `json:"limit,omitzero"`
}

type ListJoinedPrivateArchivedThreadsResponse struct {
	Threads []Channel      `json:"threads"`
	Members []ThreadMember `json:"members"`
	HasMore bool           `json:"has_more"`
}

type ListLobbyMessagesQuery struct {
	Limit int64 `json:"limit,omitzero"`
}

type ListMessagePinsQuery struct {
	Before time.Time `json:"before,omitzero"`
	Limit  int64     `json:"limit,omitzero"`
}

type ListMessagePinsResponse struct {
	Items   []MessagePin `json:"items"`
	HasMore bool         `json:"has_more"`
}

type ListMessagesQuery struct {
	Around snowflake.ID `json:"around,omitzero"`
	Before snowflake.ID `json:"before,omitzero"`
	After  snowflake.ID `json:"after,omitzero"`
	Limit  int64        `json:"limit"`
}

type ListNotificationCenterItemsQuery struct {
	After          snowflake.ID `json:"after,omitzero"`
	WithMentions   bool         `json:"with_mentions,omitzero"`
	RolesFilter    bool         `json:"roles_filter,omitzero"`
	EveryoneFilter bool         `json:"everyone_filter,omitzero"`
	Limit          int64        `json:"limit,omitzero"`
}

type ListNotificationCenterItemsResponse struct {
	Limit   int64                    `json:"limit"`
	Items   []NotificationCenterItem `json:"items"`
	Cursor  *snowflake.ID            `json:"cursor"`
	HasMore bool                     `json:"has_more"`
}

type ListOAuth2ApplicationAssetsQuery struct {
	Nocache bool `json:"nocache,omitzero"`
}

type ListOutboundPromotionsQuery struct {
	Locale   string                `json:"locale,omitzero"`
	Platform PromotionPlatformType `json:"platform,omitzero"`
}

type ListPartialApplicationsQuery struct {
	ApplicationIDs []snowflake.ID `json:"application_ids"`
}

type ListPollAnswerVotersQuery struct {
	After snowflake.ID `json:"after,omitzero"`
	Limit int64        `json:"limit,omitzero"`
}

type ListPollAnswerVotersResponse struct {
	Users []PartialUser `json:"users"`
}

type ListPremiumGuildSubscriptionsQuery struct {
	IncludeEnded bool `json:"include_ended,omitzero"`
}

type ListPremiumReferralEligibleUsersParams struct {
	Index       int64  `json:"index"`
	Limit       int64  `json:"limit,omitzero"`
	SearchQuery string `json:"search_query,omitzero"`
}

type ListPremiumReferralEligibleUsersResponse struct {
	Users     []PartialUser `json:"users"`
	NextIndex int64         `json:"next_index,omitzero"`
}

type ListPresencesForXboxResponse struct {
	Guilds              []VoiceGuild         `json:"guilds"`
	Presences           []Presence           `json:"presences"`
	Applications        []PartialApplication `json:"applications"`
	ConnectedAccountIDs []ConnectedAccount   `json:"connected_account_ids"`
}

type ListPresencesResponse struct {
	Guilds       []VoiceGuild         `json:"guilds"`
	Presences    []Presence           `json:"presences"`
	Applications []PartialApplication `json:"applications"`
}

type ListPrivateArchivedThreadsQuery struct {
	Before time.Time `json:"before,omitzero"`
	Limit  int64     `json:"limit,omitzero"`
}

type ListPrivateArchivedThreadsResponse struct {
	Threads []Channel      `json:"threads"`
	Members []ThreadMember `json:"members"`
	HasMore bool           `json:"has_more"`
}

type ListPromotionsQuery struct {
	Locale   string                `json:"locale,omitzero"`
	Platform PromotionPlatformType `json:"platform,omitzero"`
}

type ListPublicArchivedThreadsQuery struct {
	Before time.Time `json:"before,omitzero"`
	Limit  int64     `json:"limit,omitzero"`
}

type ListPublicArchivedThreadsResponse struct {
	Threads []Channel      `json:"threads"`
	Members []ThreadMember `json:"members"`
	HasMore bool           `json:"has_more"`
}

type ListPublishedSubscriptionPlansQuery struct {
	IncludeUnpublished bool           `json:"include_unpublished,omitzero"`
	RevenueSurface     RevenueSurface `json:"revenue_surface,omitzero"`
	CountryCode        string         `json:"country_code,omitzero"`
	PaymentSourceID    snowflake.ID   `json:"payment_source_id,omitzero"`
}

type ListReactionsQuery struct {
	After snowflake.ID `json:"after,omitzero"`
	Limit int64        `json:"limit,omitzero"`
	Type  ReactionType `json:"type,omitzero"`
}

type ListRecentAvatarsResponse struct {
	Avatars []Avatar `json:"avatars"`
}

type ListRecentMentionsQuery struct {
	Before   snowflake.ID `json:"before,omitzero"`
	Limit    int64        `json:"limit,omitzero"`
	GuildID  snowflake.ID `json:"guild_id,omitzero"`
	Roles    bool         `json:"roles,omitzero"`
	Everyone bool         `json:"everyone,omitzero"`
}

type ListSKUStoreListingsQuery struct {
	CountryCode string `json:"country_code,omitzero"`
	Localize    bool   `json:"localize,omitzero"`
}

type ListSavedMessagesResponse struct {
	Results []SavedMessage `json:"results"`
}

type ListSocialSdkReleasesResponse struct {
	Releases      []SocialSdkRelease `json:"releases"`
	LatestVersion string             `json:"latest_version"`
}

type ListStickerPacksResponse struct {
	StickerPacks []StickerPack `json:"sticker_packs"`
}

type ListStorePriceTiersQuery struct {
	PriceTierType PriceTierType `json:"price_tier_type,omitzero"`
}

type ListStorefrontProductsBySKUIDQuery struct {
	SKUIDs []snowflake.ID `json:"sku_ids"`
}

type ListStorefrontProductsBySKUIDResponse struct {
	Products []StorefrontProduct `json:"products"`
}

type ListStorefrontSKUPricesQuery struct {
	SKUIDs []snowflake.ID `json:"sku_ids"`
}

type ListStorefrontSKUPricesResponse struct {
	SKUPrices map[snowflake.ID]SubscriptionPrices `json:"sku_prices"`
}

type ListSubscriptionInvoicesParams struct {
	PaymentSourceID    *snowflake.ID `json:"payment_source_id"`
	PaymentSourceToken string        `json:"payment_source_token,omitzero"`
	Currency           string        `json:"currency,omitzero"`
	ReturnURL          string        `json:"return_url,omitzero"`
}

type ListSubscriptionsQuery struct {
	IncludeInactive       bool                  `json:"include_inactive,omitzero"`
	Limit                 int64                 `json:"limit,omitzero"`
	ExcludeUnpaidStatuses bool                  `json:"exclude_unpaid_statuses,omitzero"`
	SubscriptionType      SubscriptionType      `json:"subscription_type,omitzero"`
	SyncLevel             UserLazyPerkSyncLevel `json:"sync_level,omitzero"`
}

type ListTeamPayoutsQuery struct {
	Limit float64      `json:"limit,omitzero"`
	After snowflake.ID `json:"after,omitzero"`
}

type ListTeamsQuery struct {
	IncludePayoutAccountStatus bool `json:"include_payout_account_status,omitzero"`
}

type ListThreadMembersQuery struct {
	WithMember bool         `json:"with_member,omitzero"`
	After      snowflake.ID `json:"after,omitzero"`
	Limit      int64        `json:"limit,omitzero"`
}

type ListTrendingGIFCategoriesResponse struct {
	Categories []GIFCategory `json:"categories"`
	Gifs       []GIF         `json:"gifs"`
}

type ListUserAffinitiesResponse struct {
	UserAffinities []UserAffinity `json:"user_affinities"`
}

type ListUserAffinitiesV2Response struct {
	UserAffinities []UserAffinityV2 `json:"user_affinities"`
}

type ListUserApplicationEntitlementsQuery struct {
	SKUIDs          []snowflake.ID `json:"sku_ids,omitempty"`
	ExcludeConsumed bool           `json:"exclude_consumed"`
}

type ListUserApplicationIdentitiesQuery struct {
	WithProfiles bool `json:"with_profiles,omitzero"`
}

type ListUserApplicationIdentitiesResponse struct {
	Identities []UserApplicationIdentity `json:"identities"`
}

type ListUserApplicationRoleConnectionsParams struct {
	PlatformName     string            `json:"platform_name,omitzero"`
	PlatformUsername string            `json:"platform_username,omitzero"`
	Metadata         map[string]string `json:"metadata,omitempty"`
}

type ListUserEntitlementsQuery struct {
	WithSKU         bool            `json:"with_sku,omitzero"`
	WithApplication bool            `json:"with_application,omitzero"`
	ExcludeEnded    bool            `json:"exclude_ended,omitzero"`
	EntitlementType EntitlementType `json:"entitlement_type,omitzero"`
}

type ListUserGiftCodesQuery struct {
	SKUIDs             []snowflake.ID `json:"sku_ids,omitempty"`
	SubscriptionPlanID snowflake.ID   `json:"subscription_plan_id,omitzero"`
}

type ListUserGiftableEntitlementsQuery struct {
	CountryCode string `json:"country_code,omitzero"`
}

type ListUserGuildScheduledEventsQuery struct {
	GuildIDs snowflake.ID `json:"guild_ids"`
}

type ListUserGuildsQuery struct {
	Before     snowflake.ID `json:"before,omitzero"`
	After      snowflake.ID `json:"after,omitzero"`
	Limit      int64        `json:"limit,omitzero"`
	WithCounts bool         `json:"with_counts,omitzero"`
}

type ListUserPurchasedCollectiblesQuery struct {
	VariantsReturnStyle VariantsReturnStyle `json:"variants_return_style,omitzero"`
}

type Lobby struct {
	ID            snowflake.ID      `json:"id"`
	ApplicationID snowflake.ID      `json:"application_id"`
	Metadata      map[string]string `json:"metadata"`
	Members       []LobbyMember     `json:"members"`
	Flags         LobbyFlags        `json:"flags"`
	LinkedChannel Channel           `json:"linked_channel,omitzero"`
}

type LobbyDelete struct {
	ID     snowflake.ID        `json:"id"`
	Reason LobbyDeletionReason `json:"reason"`
}

type LobbyMember struct {
	ID        snowflake.ID      `json:"id"`
	Metadata  map[string]string `json:"metadata,omitempty"`
	Flags     LobbyMemberFlags  `json:"flags,omitzero"`
	Connected bool              `json:"connected"`
}

type LobbyMemberAdd struct {
	Member        LobbyMember  `json:"member"`
	LobbyID       snowflake.ID `json:"lobby_id"`
	ApplicationID snowflake.ID `json:"application_id"`
}

type LobbyMemberConnect struct {
	Member  LobbyMember  `json:"member"`
	LobbyID snowflake.ID `json:"lobby_id"`
}

type LobbyMemberDisconnect struct {
	Member  LobbyMember  `json:"member"`
	LobbyID snowflake.ID `json:"lobby_id"`
}

type LobbyMemberRemove struct {
	Member        LobbyMember  `json:"member"`
	LobbyID       snowflake.ID `json:"lobby_id"`
	ApplicationID snowflake.ID `json:"application_id"`
}

type LobbyMemberUpdate struct {
	Member        LobbyMember  `json:"member"`
	LobbyID       snowflake.ID `json:"lobby_id"`
	ApplicationID snowflake.ID `json:"application_id"`
}

type LobbyVoiceServerUpdate struct {
	Token    string       `json:"token"`
	LobbyID  snowflake.ID `json:"lobby_id"`
	Endpoint *string      `json:"endpoint"`
}

type LocalizationSettings struct {
	Locale         string `json:"locale,omitzero"`
	TimezoneOffset int64  `json:"timezone_offset,omitzero"`
	TimezoneName   string `json:"timezone_name,omitzero"`
}

type LocalizedPrice struct {
	Currency string `json:"currency"`
	Amount   int64  `json:"amount"`
}

type LocalizedPricingPromo struct {
	PlanID             snowflake.ID   `json:"plan_id"`
	CountryCode        string         `json:"country_code"`
	PaymentSourceTypes []int64        `json:"payment_source_types"`
	Price              LocalizedPrice `json:"price"`
}

type LoginAccountParams struct {
	Login         string       `json:"login"`
	Password      string       `json:"password"`
	Undelete      bool         `json:"undelete,omitzero"`
	LoginSource   *LoginSource `json:"login_source,omitempty"`
	GiftCodeSKUID *string      `json:"gift_code_sku_id,omitempty"`
}

type LoginAccountResponse struct {
	UserID          snowflake.ID  `json:"user_id"`
	Token           string        `json:"token,omitzero"`
	UserSettings    LoginSettings `json:"user_settings,omitzero"`
	RequiredActions []string      `json:"required_actions,omitempty"`
	Ticket          string        `json:"ticket,omitzero"`
	LoginInstanceID string        `json:"login_instance_id,omitzero"`
	MFA             bool          `json:"mfa,omitzero"`
	Totp            bool          `json:"totp,omitzero"`
	SMS             bool          `json:"sms,omitzero"`
	Backup          bool          `json:"backup,omitzero"`
	Webauthn        *string       `json:"webauthn,omitempty"`
}

type LoginSettings struct {
	Locale string `json:"locale"`
	Theme  Theme  `json:"theme"`
}

type LogoutAuthSessionsParams struct {
	SessionIDHashes []string `json:"session_id_hashes"`
}

type LogoutParams struct {
	Provider     PushNotificationProvider `json:"provider,omitzero"`
	Token        string                   `json:"token,omitzero"`
	VoipProvider PushNotificationProvider `json:"voip_provider,omitzero"`
	VoipToken    string                   `json:"voip_token,omitzero"`
}

type MFAMethod struct {
	Type               AuthenticatorType `json:"type"`
	Challenge          string            `json:"challenge,omitzero"`
	BackupCodesAllowed bool              `json:"backup_codes_allowed,omitzero"`
}

type MFARequiredResponse struct {
	Message string                 `json:"message"`
	Code    int64                  `json:"code"`
	MFA     MFAVerificationRequest `json:"mfa"`
}

type MFAVerificationRequest struct {
	Ticket  string          `json:"ticket"`
	Methods json.RawMessage `json:"methods"`
}

type ManifestPackage struct {
	Full   ManifestPackageVersion   `json:"full"`
	Deltas []ManifestPackageVersion `json:"deltas"`
}

type ManifestPackageVersion struct {
	HostVersion   []int64 `json:"host_version"`
	ModuleVersion int64   `json:"module_version,omitzero"`
	PackageSha256 string  `json:"package_sha256"`
	URL           string  `json:"url"`
}

type MaybeGetProviderAccessTokenArguments struct {
	Provider ConnectionType `json:"provider"`
}

type MaybeGetProviderAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
}

type MediaGallery struct {
	Type  ComponentType      `json:"type"`
	ID    int64              `json:"id,omitzero"`
	Items []MediaGalleryItem `json:"items"`
}

type MediaGalleryItem struct {
	Media       json.RawMessage `json:"media"`
	Description *string         `json:"description,omitempty"`
	Spoiler     bool            `json:"spoiler,omitzero"`
}

type MediaPreview struct {
	GuildID            snowflake.ID `json:"guild_id"`
	GuildName          string       `json:"guild_name"`
	GuildIcon          string       `json:"guild_icon"`
	ChannelID          snowflake.ID `json:"channel_id"`
	ParentChannelID    snowflake.ID `json:"parent_channel_id"`
	MessageID          snowflake.ID `json:"message_id"`
	AuthorID           snowflake.ID `json:"author_id"`
	Title              string       `json:"title"`
	Description        string       `json:"description"`
	HasMediaAttachment bool         `json:"has_media_attachment"`
	Thumbnail          Attachment   `json:"thumbnail,omitzero"`
}

type MediaSinkWants struct {
	SSRC        int64           `json:"{ssrc},omitzero"`
	Any         int64           `json:"any,omitzero"`
	Pixelcounts json.RawMessage `json:"pixelCounts,omitempty"`
}

type MemberFilter struct {
	UserID           json.RawMessage `json:"user_id,omitempty"`
	Usernames        json.RawMessage `json:"usernames,omitempty"`
	RoleIDs          json.RawMessage `json:"role_ids,omitempty"`
	GuildJoinedAt    json.RawMessage `json:"guild_joined_at,omitempty"`
	SafetySignals    SafetySignals   `json:"safety_signals,omitzero"`
	IsPending        bool            `json:"is_pending,omitzero"`
	DidRejoin        bool            `json:"did_rejoin,omitzero"`
	JoinSourceType   json.RawMessage `json:"join_source_type,omitempty"`
	SourceInviteCode json.RawMessage `json:"source_invite_code,omitempty"`
}

type MemberPaginationFilter struct {
	UserID        snowflake.ID `json:"user_id"`
	GuildJoinedAt int64        `json:"guild_joined_at"`
}

type MemberVerification struct {
	Version     *time.Time                    `json:"version"`
	FormFields  []MemberVerificationFormField `json:"form_fields"`
	Description *string                       `json:"description"`
	Guild       *MemberVerificationGuild      `json:"guild"`
	Profile     GuildProfile                  `json:"profile"`
}

type MemberVerificationFormField struct {
	FieldType   MemberVerificationFormFieldType `json:"field_type"`
	Label       string                          `json:"label"`
	Choices     []string                        `json:"choices,omitempty"`
	Values      []string                        `json:"values,omitempty"`
	Response    json.RawMessage                 `json:"response,omitempty"`
	Required    bool                            `json:"required"`
	Description *string                         `json:"description"`
	Automations []string                        `json:"automations"`
	Placeholder *string                         `json:"placeholder,omitempty"`
}

type MemberVerificationGuild struct {
	ID                       snowflake.ID      `json:"id"`
	Name                     string            `json:"name"`
	Icon                     *string           `json:"icon"`
	Description              *string           `json:"description"`
	Splash                   *string           `json:"splash"`
	DiscoverySplash          *string           `json:"discovery_splash"`
	HomeHeader               *string           `json:"home_header"`
	VerificationLevel        VerificationLevel `json:"verification_level"`
	Features                 []string          `json:"features"`
	Emojis                   []Emoji           `json:"emojis"`
	ApproximateMemberCount   int64             `json:"approximate_member_count"`
	ApproximatePresenceCount int64             `json:"approximate_presence_count"`
}

type MentionableSelect struct {
	Type          ComponentType     `json:"type"`
	ID            int64             `json:"id,omitzero"`
	CustomID      string            `json:"custom_id"`
	Placeholder   string            `json:"placeholder,omitzero"`
	DefaultValues []json.RawMessage `json:"default_values,omitempty"`
	MinValues     int64             `json:"min_values,omitzero"`
	MaxValues     int64             `json:"max_values,omitzero"`
	Required      bool              `json:"required,omitzero"`
	Disabled      bool              `json:"disabled,omitzero"`
}

type MergedPresences struct {
	Friends []Presence   `json:"friends"`
	Guilds  [][]Presence `json:"guilds"`
}

type Message struct {
	ID                   snowflake.ID                `json:"id"`
	ChannelID            snowflake.ID                `json:"channel_id"`
	LobbyID              snowflake.ID                `json:"lobby_id,omitzero"`
	Author               PartialUser                 `json:"author"`
	Content              string                      `json:"content"`
	Timestamp            time.Time                   `json:"timestamp"`
	EditedTimestamp      *time.Time                  `json:"edited_timestamp"`
	TTS                  bool                        `json:"tts"`
	MentionEveryone      bool                        `json:"mention_everyone"`
	Mentions             []PartialUser               `json:"mentions"`
	MentionRoles         []snowflake.ID              `json:"mention_roles"`
	MentionChannels      []PartialChannel            `json:"mention_channels,omitempty"`
	Attachments          []Attachment                `json:"attachments"`
	Embeds               []Embed                     `json:"embeds"`
	Reactions            []Reaction                  `json:"reactions,omitempty"`
	Nonce                json.RawMessage             `json:"nonce,omitempty"`
	Pinned               bool                        `json:"pinned"`
	WebhookID            snowflake.ID                `json:"webhook_id,omitzero"`
	Type                 MessageType                 `json:"type"`
	Activity             MessageActivity             `json:"activity,omitzero"`
	Application          IntegrationApplication      `json:"application,omitzero"`
	ApplicationID        snowflake.ID                `json:"application_id,omitzero"`
	Flags                MessageFlags                `json:"flags"`
	MessageReference     MessageReference            `json:"message_reference,omitzero"`
	ReferencedMessage    *Message                    `json:"referenced_message,omitempty"`
	MessageSnapshots     []MessageSnapshot           `json:"message_snapshots,omitempty"`
	Call                 MessageCall                 `json:"call,omitzero"`
	Interaction          PartialMessageInteraction   `json:"interaction,omitzero"`
	InteractionMetadata  MessageInteraction          `json:"interaction_metadata,omitzero"`
	Resolved             ResolvedData                `json:"resolved,omitzero"`
	Thread               Channel                     `json:"thread,omitzero"`
	RoleSubscriptionData MessageRoleSubscription     `json:"role_subscription_data,omitzero"`
	PurchaseNotification MessagePurchaseNotification `json:"purchase_notification,omitzero"`
	GiftInfo             MessageGiftInfo             `json:"gift_info,omitzero"`
	Components           Components                  `json:"components"`
	StickerItems         []StickerItem               `json:"sticker_items,omitempty"`
	Stickers             []Sticker                   `json:"stickers,omitempty"`
	Poll                 Poll                        `json:"poll,omitzero"`
	ChangelogID          snowflake.ID                `json:"changelog_id,omitzero"`
	SoundboardSounds     []SoundboardSound           `json:"soundboard_sounds,omitempty"`
	Potions              []Potion                    `json:"potions,omitempty"`
	SharedClientTheme    SharedClientTheme           `json:"shared_client_theme,omitzero"`
}

type MessageAck struct {
	AckType      ReadStateType   `json:"ack_type,omitzero"`
	ChannelID    snowflake.ID    `json:"channel_id"`
	MessageID    snowflake.ID    `json:"message_id"`
	Manual       bool            `json:"manual,omitzero"`
	MentionCount int64           `json:"mention_count,omitzero"`
	Flags        *ReadStateFlags `json:"flags,omitempty"`
	LastViewed   *int64          `json:"last_viewed,omitempty"`
	Version      int64           `json:"version"`
}

type MessageActivity struct {
	Type         ActivityActionType `json:"type"`
	SessionID    string             `json:"session_id"`
	PartyID      string             `json:"party_id,omitzero"`
	NameOverride string             `json:"name_override,omitzero"`
	IconOverride string             `json:"icon_override,omitzero"`
}

type MessageCall struct {
	Participants   []snowflake.ID `json:"participants"`
	EndedTimestamp *time.Time     `json:"ended_timestamp"`
}

type MessageComponentData struct {
	CustomID      string          `json:"custom_id"`
	ComponentType ComponentType   `json:"component_type"`
	Values        json.RawMessage `json:"values"`
	Resolved      ResolvedData    `json:"resolved,omitzero"`
}

type MessageCreate struct {
	Message                PartialMessage        `json:"message,omitzero"`
	MessageID              snowflake.ID          `json:"message_id"`
	MessageType            MessageType           `json:"message_type_"`
	MessageContent         string                `json:"message_content"`
	MessageFlags           MessageFlags          `json:"message_flags,omitzero"`
	MessageReferenceType   MessageReferenceType  `json:"message_reference_type,omitzero"`
	MessageActivityType    ActivityActionType    `json:"message_activity_type,omitzero"`
	MessageApplicationName string                `json:"message_application_name,omitzero"`
	ChannelID              snowflake.ID          `json:"channel_id"`
	ChannelType            ChannelType           `json:"channel_type"`
	ChannelName            string                `json:"channel_name"`
	ChannelIcon            *string               `json:"channel_icon,omitempty"`
	IsSpoilerChannel       bool                  `json:"is_spoiler_channel,omitzero"`
	AppDM                  bool                  `json:"app_dm,omitzero"`
	GuildID                snowflake.ID          `json:"guild_id,omitzero"`
	GuildName              string                `json:"guild_name,omitzero"`
	GuildIcon              *string               `json:"guild_icon,omitempty"`
	UserID                 snowflake.ID          `json:"user_id"`
	UserUsername           string                `json:"user_username"`
	UserGlobalName         *string               `json:"user_global_name,omitempty"`
	UserDiscriminator      *int64                `json:"user_discriminator,omitempty"`
	UserAvatar             *string               `json:"user_avatar,omitempty"`
	UserGuildAvatar        *string               `json:"user_guild_avatar,omitempty"`
	IsFromCurrentUser      bool                  `json:"is_from_current_user,omitzero"`
	Category               string                `json:"__category,omitzero"`
	PollQuestion           string                `json:"poll_question,omitzero"`
	ImageURL               string                `json:"image_url,omitzero"`
	ImageURLErgoAndroid    string                `json:"image_url_ergo_android,omitzero"`
	ImageCount             int64                 `json:"image_count,omitzero"`
	VideoCount             int64                 `json:"video_count,omitzero"`
	AttachmentTextVariant  AttachmentTextVariant `json:"attachment_text_variant,omitzero"`
	InviteGuildName        string                `json:"invite_guild_name,omitzero"`
	InviteChannelName      string                `json:"invite_channel_name,omitzero"`
	InviteTitleVariant     InviteTitleVariant    `json:"invite_title_variant,omitzero"`
}

type MessageCreateSubscriptionArguments struct {
	ChannelID snowflake.ID `json:"channel_id"`
}

type MessageDelete struct {
	ID      snowflake.ID `json:"id"`
	LobbyID snowflake.ID `json:"lobby_id"`
}

type MessageDeleteBulk struct {
	IDs       []snowflake.ID `json:"ids"`
	ChannelID snowflake.ID   `json:"channel_id"`
	GuildID   snowflake.ID   `json:"guild_id,omitzero"`
}

type MessageDeleteSubscriptionArguments struct {
	ChannelID snowflake.ID `json:"channel_id"`
}

type MessageForwardOnly struct {
	EmbedIndices  []int64        `json:"embed_indices,omitempty"`
	AttachmentIDs []snowflake.ID `json:"attachment_ids,omitempty"`
}

type MessageGiftInfo struct {
	Emoji Emoji                  `json:"emoji,omitzero"`
	Sound MessageSoundboardSound `json:"sound,omitzero"`
}

type MessageInteraction struct {
	ID                            snowflake.ID           `json:"id"`
	Type                          InteractionType        `json:"type"`
	Name                          string                 `json:"name,omitzero"`
	CommandType                   ApplicationCommandType `json:"command_type,omitzero"`
	EphemeralityReason            EphemeralityReason     `json:"ephemerality_reason,omitzero"`
	User                          PartialUser            `json:"user"`
	AuthorizingIntegrationOwners  map[int64]snowflake.ID `json:"authorizing_integration_owners"`
	OriginalResponseMessageID     snowflake.ID           `json:"original_response_message_id,omitzero"`
	InteractedMessageID           snowflake.ID           `json:"interacted_message_id,omitzero"`
	TriggeringInteractionMetadata *MessageInteraction    `json:"triggering_interaction_metadata,omitempty"`
	TargetUser                    PartialUser            `json:"target_user,omitzero"`
	TargetMessageID               snowflake.ID           `json:"target_message_id,omitzero"`
}

type MessageObjectExtraFields struct {
	ChannelType        ChannelType       `json:"channel_type"`
	GuildID            snowflake.ID      `json:"guild_id,omitzero"`
	Member             GuildMember       `json:"member,omitzero"`
	Mentions           []PartialUser     `json:"mentions"`
	Metadata           map[string]string `json:"metadata,omitempty"`
	ModerationMetadata map[string]string `json:"moderation_metadata,omitempty"`
}

type MessagePin struct {
	PinnedAt time.Time `json:"pinned_at"`
	Message  Message   `json:"message"`
}

type MessagePollVoteAdd struct {
	UserID    snowflake.ID `json:"user_id"`
	ChannelID snowflake.ID `json:"channel_id"`
	MessageID snowflake.ID `json:"message_id"`
	GuildID   snowflake.ID `json:"guild_id,omitzero"`
	AnswerID  int64        `json:"answer_id"`
}

type MessagePollVoteRemove struct {
	UserID    snowflake.ID `json:"user_id"`
	ChannelID snowflake.ID `json:"channel_id"`
	MessageID snowflake.ID `json:"message_id"`
	GuildID   snowflake.ID `json:"guild_id,omitzero"`
	AnswerID  int64        `json:"answer_id"`
}

type MessagePurchaseNotification struct {
	Type                 MessagePurchaseNotificationType `json:"type"`
	GuildProductPurchase *GuildProductPurchase           `json:"guild_product_purchase"`
}

type MessageReactionAdd struct {
	UserID          snowflake.ID    `json:"user_id"`
	ChannelID       snowflake.ID    `json:"channel_id"`
	MessageID       snowflake.ID    `json:"message_id"`
	MessageAuthorID snowflake.ID    `json:"message_author_id"`
	GuildID         snowflake.ID    `json:"guild_id,omitzero"`
	Member          json.RawMessage `json:"member,omitempty"`
	Emoji           Emoji           `json:"emoji"`
	Type            ReactionType    `json:"type"`
	BurstColors     []string        `json:"burst_colors,omitempty"`
}

type MessageReactionAddMany struct {
	ChannelID snowflake.ID        `json:"channel_id"`
	MessageID snowflake.ID        `json:"message_id"`
	GuildID   snowflake.ID        `json:"guild_id,omitzero"`
	Reactions []DebouncedReaction `json:"reactions"`
}

type MessageReactionRemove struct {
	UserID    snowflake.ID    `json:"user_id"`
	ChannelID snowflake.ID    `json:"channel_id"`
	MessageID snowflake.ID    `json:"message_id"`
	GuildID   snowflake.ID    `json:"guild_id,omitzero"`
	Emoji     json.RawMessage `json:"emoji"`
	Type      ReactionType    `json:"type"`
}

type MessageReactionRemoveAll struct {
	ChannelID snowflake.ID `json:"channel_id"`
	MessageID snowflake.ID `json:"message_id"`
	GuildID   snowflake.ID `json:"guild_id,omitzero"`
}

type MessageReactionRemoveEmoji struct {
	ChannelID snowflake.ID `json:"channel_id"`
	MessageID snowflake.ID `json:"message_id"`
	GuildID   snowflake.ID `json:"guild_id,omitzero"`
	Emoji     Emoji        `json:"emoji"`
}

type MessageReference struct {
	Type            MessageReferenceType `json:"type,omitzero"`
	MessageID       snowflake.ID         `json:"message_id,omitzero"`
	ChannelID       snowflake.ID         `json:"channel_id"`
	GuildID         snowflake.ID         `json:"guild_id,omitzero"`
	FailIfNotExists bool                 `json:"fail_if_not_exists,omitzero"`
	ForwardOnly     MessageForwardOnly   `json:"forward_only,omitzero"`
}

type MessageRoleSubscription struct {
	RoleSubscriptionListingID snowflake.ID `json:"role_subscription_listing_id"`
	TierName                  string       `json:"tier_name"`
	TotalMonthsSubscribed     int64        `json:"total_months_subscribed"`
	IsRenewal                 bool         `json:"is_renewal"`
}

type MessageSnapshot struct {
	Message SnapshotMessage `json:"message"`
}

type MessageSoundboardSound struct {
	ID string `json:"id"`
}

type MessageUpdate struct {
	ChannelID snowflake.ID `json:"channel_id"`
	Message   RPCMessage   `json:"message"`
}

type MessageUpdateSubscriptionArguments struct {
	ChannelID snowflake.ID `json:"channel_id"`
}

type MigrateGuildCommandScopeResponse struct {
	IntegrationIDsWithAppCommands []snowflake.ID `json:"integration_ids_with_app_commands"`
}

type ModalSubmitComponentData struct {
	ID         int64                      `json:"id"`
	Type       ComponentType              `json:"type"`
	CustomID   string                     `json:"custom_id"`
	Value      json.RawMessage            `json:"value"`
	Values     json.RawMessage            `json:"values"`
	Component  *ModalSubmitComponentData  `json:"component"`
	Components []ModalSubmitComponentData `json:"components"`
}

type ModalSubmitData struct {
	CustomID   string                     `json:"custom_id"`
	Components []ModalSubmitComponentData `json:"components"`
	Resolved   ResolvedData               `json:"resolved,omitzero"`
}

type ModifyApplicationBotParams struct {
	Username string     `json:"username,omitzero"`
	Avatar   *ImageData `json:"avatar,omitempty"`
	Banner   *ImageData `json:"banner,omitempty"`
}

type ModifyApplicationEmojiParams struct {
	Name string `json:"name,omitzero"`
}

type ModifyApplicationGameClaimParams struct {
	AnnouncementsChannelID *snowflake.ID `json:"announcements_channel_id,omitempty"`
}

type ModifyApplicationParams struct {
	Name                           string                                             `json:"name,omitzero"`
	Description                    *string                                            `json:"description,omitempty"`
	Icon                           *ImageData                                         `json:"icon,omitempty"`
	CoverImage                     *ImageData                                         `json:"cover_image,omitempty"`
	Flags                          ApplicationFlags                                   `json:"flags,omitzero"`
	GuildID                        *snowflake.ID                                      `json:"guild_id,omitempty"`
	DeveloperIDs                   []snowflake.ID                                     `json:"developer_ids,omitempty"`
	PublisherIDs                   []snowflake.ID                                     `json:"publisher_ids,omitempty"`
	RPCOrigins                     []string                                           `json:"rpc_origins,omitempty"`
	RedirectUris                   []string                                           `json:"redirect_uris,omitempty"`
	DeeplinkURI                    *string                                            `json:"deeplink_uri,omitempty"`
	IntegrationPublic              bool                                               `json:"integration_public,omitzero"`
	IntegrationRequireCodeGrant    bool                                               `json:"integration_require_code_grant,omitzero"`
	BotPublic                      bool                                               `json:"bot_public,omitzero"`
	BotRequireCodeGrant            bool                                               `json:"bot_require_code_grant,omitzero"`
	TermsOfServiceURL              *string                                            `json:"terms_of_service_url,omitempty"`
	PrivacyPolicyURL               *string                                            `json:"privacy_policy_url,omitempty"`
	RoleConnectionsVerificationURL *string                                            `json:"role_connections_verification_url"`
	InteractionsEndpointURL        string                                             `json:"interactions_endpoint_url,omitzero"`
	InteractionsVersion            ApplicationInteractionsVersion                     `json:"interactions_version,omitzero"`
	InteractionsEventTypes         []string                                           `json:"interactions_event_types,omitempty"`
	EventWebhooksStatus            EventWebhooksStatus                                `json:"event_webhooks_status,omitzero"`
	EventWebhooksURL               string                                             `json:"event_webhooks_url,omitzero"`
	EventWebhooksTypes             []string                                           `json:"event_webhooks_types,omitempty"`
	ExplicitContentFilter          ExplicitContentFilterLevel                         `json:"explicit_content_filter,omitzero"`
	Tags                           []string                                           `json:"tags,omitempty"`
	InstallParams                  *ApplicationInstallParams                          `json:"install_params,omitempty"`
	CustomInstallURL               *string                                            `json:"custom_install_url,omitempty"`
	IntegrationTypesConfig         map[int64]*ApplicationIntegrationTypeConfiguration `json:"integration_types_config,omitempty"`
	ConnectionEntrypointURL        string                                             `json:"connection_entrypoint_url,omitzero"`
	DiscoverabilityState           ApplicationDiscoverabilityState                    `json:"discoverability_state,omitzero"`
	MonetizationState              ApplicationMonetizationState                       `json:"monetization_state,omitzero"`
	MaxParticipants                *int64                                             `json:"max_participants,omitempty"`
}

type ModifyApplicationProxyConfigParams struct {
	URLMap []ApplicationProxyMapping `json:"url_map"`
}

type ModifyApplicationStorefrontParams struct {
	Title                 string           `json:"title"`
	LogoAssetID           *snowflake.ID    `json:"logo_asset_id,omitempty"`
	LightThemeLogoAssetID *snowflake.ID    `json:"light_theme_logo_asset_id,omitempty"`
	Pages                 []StorefrontPage `json:"pages"`
}

type ModifyApplicationStorefrontPublishStatesParams struct {
	PublishState StorefrontPublishState `json:"publish_state"`
	SKUIDs       []snowflake.ID         `json:"sku_ids"`
}

type ModifyApplicationStorefrontPublishStatesResponse struct {
	StoreLayoutSKUIDs []snowflake.ID `json:"store_layout_sku_ids,omitempty"`
	StoreListings     []StoreListing `json:"store_listings,omitempty"`
}

type ModifyApplicationWidgetConfigParams struct {
	DisplayName string                   `json:"display_name,omitzero"`
	Surfaces    map[string]WidgetSurface `json:"surfaces"`
}

type ModifyAudioSettingsParams struct {
	Muted           bool    `json:"muted,omitzero"`
	Volume          float64 `json:"volume,omitzero"`
	SoundboardMuted bool    `json:"soundboard_muted,omitzero"`
}

type ModifyAutomodIncidentActionsParams struct {
	InvitesDisabledUntil time.Time `json:"invites_disabled_until,omitzero"`
	DMsDisabledUntil     time.Time `json:"dms_disabled_until,omitzero"`
}

type ModifyAutomodIncidentActionsResponse struct {
	InvitesDisabledUntil time.Time `json:"invites_disabled_until"`
	DMsDisabledUntil     time.Time `json:"dms_disabled_until"`
}

type ModifyCallParams struct {
	Region string `json:"region,omitzero"`
}

type ModifyChannelParams struct {
	Name                          string                `json:"name,omitzero"`
	Type                          ChannelType           `json:"type,omitzero"`
	Position                      *int64                `json:"position,omitempty"`
	Topic                         *string               `json:"topic,omitempty"`
	Icon                          ImageData             `json:"icon,omitzero"`
	NSFW                          *bool                 `json:"nsfw,omitempty"`
	RateLimitPerUser              *int64                `json:"rate_limit_per_user,omitempty"`
	Bitrate                       *int64                `json:"bitrate,omitempty"`
	UserLimit                     *int64                `json:"user_limit,omitempty"`
	PermissionOverwrites          []PermissionOverwrite `json:"permission_overwrites,omitempty"`
	ParentID                      *snowflake.ID         `json:"parent_id,omitempty"`
	Owner                         *snowflake.ID         `json:"owner,omitempty"`
	RTCRegion                     *string               `json:"rtc_region,omitempty"`
	VideoQualityMode              *VideoQualityMode     `json:"video_quality_mode,omitempty"`
	DefaultAutoArchiveDuration    *int64                `json:"default_auto_archive_duration,omitempty"`
	DefaultThreadRateLimitPerUser int64                 `json:"default_thread_rate_limit_per_user,omitzero"`
	AutoArchiveDuration           int64                 `json:"auto_archive_duration,omitzero"`
	Archived                      *bool                 `json:"archived,omitempty"`
	Locked                        *bool                 `json:"locked,omitempty"`
	Invitable                     *bool                 `json:"invitable,omitempty"`
	Flags                         ChannelFlags          `json:"flags,omitzero"`
	AvailableTags                 []ForumTag            `json:"available_tags,omitempty"`
	AppliedTags                   []snowflake.ID        `json:"applied_tags,omitempty"`
	DefaultReactionEmoji          *DefaultReaction      `json:"default_reaction_emoji,omitempty"`
	DefaultForumLayout            *ForumLayoutType      `json:"default_forum_layout,omitempty"`
	DefaultSortOrder              *SortOrderType        `json:"default_sort_order,omitempty"`
	DefaultTagSetting             *SearchTagSetting     `json:"default_tag_setting,omitempty"`
	IconEmoji                     *IconEmoji            `json:"icon_emoji,omitempty"`
	ThemeColor                    *int64                `json:"theme_color,omitempty"`
	Template                      string                `json:"template,omitzero"`
}

type ModifyChannelPermissionsParams struct {
	Type  int64       `json:"type"`
	Allow Permissions `json:"allow,omitzero"`
	Deny  Permissions `json:"deny,omitzero"`
}

type ModifyChannelStatusParams struct {
	Status *string `json:"status"`
}

type ModifyChannelTagParams struct {
	Name      string        `json:"name"`
	Moderated bool          `json:"moderated,omitzero"`
	EmojiID   *snowflake.ID `json:"emoji_id,omitempty"`
	EmojiName *string       `json:"emoji_name,omitempty"`
}

type ModifyCurrentApplicationParams struct {
	Description                    *string                                            `json:"description,omitempty"`
	Icon                           *ImageData                                         `json:"icon,omitempty"`
	CoverImage                     *ImageData                                         `json:"cover_image,omitempty"`
	Flags                          ApplicationFlags                                   `json:"flags,omitzero"`
	RPCOrigins                     []string                                           `json:"rpc_origins,omitempty"`
	DeeplinkURI                    *string                                            `json:"deeplink_uri,omitempty"`
	RoleConnectionsVerificationURL *string                                            `json:"role_connections_verification_url,omitempty"`
	InteractionsEndpointURL        *string                                            `json:"interactions_endpoint_url,omitempty"`
	InteractionsVersion            ApplicationInteractionsVersion                     `json:"interactions_version,omitzero"`
	InteractionsEventTypes         []string                                           `json:"interactions_event_types,omitempty"`
	ExplicitContentFilter          ExplicitContentFilterLevel                         `json:"explicit_content_filter,omitzero"`
	Tags                           []string                                           `json:"tags,omitempty"`
	InstallParams                  *ApplicationInstallParams                          `json:"install_params,omitempty"`
	CustomInstallURL               *string                                            `json:"custom_install_url,omitempty"`
	IntegrationTypesConfig         map[int64]*ApplicationIntegrationTypeConfiguration `json:"integration_types_config,omitempty"`
	ConnectionEntrypointURL        string                                             `json:"connection_entrypoint_url,omitzero"`
	MaxParticipants                *int64                                             `json:"max_participants,omitempty"`
}

type ModifyCurrentGuildMemberParams struct {
	Nick                  *string            `json:"nick,omitempty"`
	Avatar                *ImageData         `json:"avatar,omitempty"`
	AvatarDescription     *string            `json:"avatar_description,omitempty"`
	AvatarID              string             `json:"avatar_id,omitzero"`
	AvatarDecorationSKUID *snowflake.ID      `json:"avatar_decoration_sku_id,omitempty"`
	Collectibles          *GuildCollectibles `json:"collectibles,omitempty"`
	DisplayNameFontID     *DisplayNameFont   `json:"display_name_font_id,omitempty"`
	DisplayNameEffectID   *DisplayNameEffect `json:"display_name_effect_id,omitempty"`
	DisplayNameColors     []int64            `json:"display_name_colors,omitempty"`
	Pronouns              *string            `json:"pronouns,omitempty"`
	Bio                   *string            `json:"bio,omitempty"`
	Banner                *ImageData         `json:"banner,omitempty"`
}

type ModifyCurrentUserAccountParams struct {
	GlobalName *string `json:"global_name,omitempty"`
}

type ModifyCurrentUserParams struct {
	Username              string                   `json:"username,omitzero"`
	Discriminator         string                   `json:"discriminator,omitzero"`
	GlobalName            *string                  `json:"global_name,omitempty"`
	Avatar                *ImageData               `json:"avatar,omitempty"`
	AvatarDescription     *string                  `json:"avatar_description,omitempty"`
	AvatarID              string                   `json:"avatar_id,omitzero"`
	AvatarDecorationSKUID *snowflake.ID            `json:"avatar_decoration_sku_id,omitempty"`
	NameplateSKUID        *snowflake.ID            `json:"nameplate_sku_id,omitempty"`
	DisplayNameFontID     *DisplayNameFont         `json:"display_name_font_id,omitempty"`
	DisplayNameEffectID   *DisplayNameEffect       `json:"display_name_effect_id,omitempty"`
	DisplayNameColors     []int64                  `json:"display_name_colors,omitempty"`
	Email                 string                   `json:"email,omitzero"`
	EmailToken            string                   `json:"email_token,omitzero"`
	Pronouns              *string                  `json:"pronouns,omitempty"`
	Bio                   *string                  `json:"bio,omitempty"`
	Banner                *ImageData               `json:"banner,omitempty"`
	AccentColor           *int64                   `json:"accent_color,omitempty"`
	Flags                 UserFlags                `json:"flags,omitzero"`
	DateOfBirth           time.Time                `json:"date_of_birth,omitzero"`
	Password              string                   `json:"password,omitzero"`
	NewPassword           string                   `json:"new_password,omitzero"`
	PushProvider          PushNotificationProvider `json:"push_provider,omitzero"`
	PushToken             string                   `json:"push_token,omitzero"`
	PushVoipProvider      PushNotificationProvider `json:"push_voip_provider,omitzero"`
	PushVoipToken         string                   `json:"push_voip_token,omitzero"`
}

type ModifyCurrentUserVoiceStateParams struct {
	ChannelID               snowflake.ID `json:"channel_id,omitzero"`
	Suppress                bool         `json:"suppress,omitzero"`
	RequestToSpeakTimestamp *time.Time   `json:"request_to_speak_timestamp,omitempty"`
	Silent                  bool         `json:"silent,omitzero"`
}

type ModifyDirectoryEntryParams struct {
	PrimaryCategoryID DirectoryCategory `json:"primary_category_id,omitzero"`
	Description       string            `json:"description,omitzero"`
}

type ModifyEmailSettingsParams struct {
	Settings EmailSettings `json:"settings"`
}

type ModifyEmbeddedActivityConfigParams struct {
	ActivityPreviewVideoAssetID       *snowflake.ID                             `json:"activity_preview_video_asset_id,omitempty"`
	SupportedPlatforms                []string                                  `json:"supported_platforms,omitempty"`
	DefaultOrientationLockState       EmbeddedActivityOrientationLockStateType  `json:"default_orientation_lock_state,omitzero"`
	TabletDefaultOrientationLockState EmbeddedActivityOrientationLockStateType  `json:"tablet_default_orientation_lock_state,omitzero"`
	RequiresAgeGate                   bool                                      `json:"requires_age_gate,omitzero"`
	FreePeriodStartsAt                *time.Time                                `json:"free_period_starts_at,omitempty"`
	FreePeriodEndsAt                  *time.Time                                `json:"free_period_ends_at,omitempty"`
	ClientPlatformConfig              map[string]EmbeddedActivityPlatformConfig `json:"client_platform_config,omitempty"`
	ShelfRank                         int64                                     `json:"shelf_rank,omitzero"`
}

type ModifyGuildChannelPositionsParams struct {
	ID              snowflake.ID  `json:"id"`
	Position        *int64        `json:"position,omitempty"`
	LockPermissions *bool         `json:"lock_permissions,omitempty"`
	ParentID        *snowflake.ID `json:"parent_id,omitempty"`
}

type ModifyGuildDiscoveryMetadataParams struct {
	PrimaryCategoryID           *int64            `json:"primary_category_id,omitempty"`
	Keywords                    []string          `json:"keywords,omitempty"`
	EmojiDiscoverabilityEnabled *bool             `json:"emoji_discoverability_enabled,omitempty"`
	IsPublished                 *bool             `json:"is_published,omitempty"`
	ReasonsToJoin               []DiscoveryReason `json:"reasons_to_join,omitempty"`
	SocialLinks                 []string          `json:"social_links,omitempty"`
	About                       *string           `json:"about,omitempty"`
}

type ModifyGuildEmojiParams struct {
	Name  string         `json:"name,omitzero"`
	Roles []snowflake.ID `json:"roles,omitempty"`
}

type ModifyGuildIntegrationParams struct {
	ExpireBehavior    IntegrationExpireBehavior `json:"expire_behavior,omitzero"`
	ExpireGracePeriod int64                     `json:"expire_grace_period,omitzero"`
	EnableEmoticons   bool                      `json:"enable_emoticons,omitzero"`
}

type ModifyGuildMemberParams struct {
	Nick                       *string          `json:"nick,omitempty"`
	Roles                      []snowflake.ID   `json:"roles,omitempty"`
	Mute                       bool             `json:"mute,omitzero"`
	Deaf                       bool             `json:"deaf,omitzero"`
	ChannelID                  snowflake.ID     `json:"channel_id,omitzero"`
	CommunicationDisabledUntil *time.Time       `json:"communication_disabled_until,omitempty"`
	Flags                      GuildMemberFlags `json:"flags,omitzero"`
}

type ModifyGuildMemberProfileParams struct {
	Pronouns                    *string       `json:"pronouns,omitempty"`
	Bio                         *string       `json:"bio,omitempty"`
	Banner                      *ImageData    `json:"banner,omitempty"`
	AccentColor                 *int64        `json:"accent_color,omitempty"`
	ThemeColors                 []int64       `json:"theme_colors,omitempty"`
	PopoutAnimationParticleType *snowflake.ID `json:"popout_animation_particle_type,omitempty"`
	EmojiID                     *snowflake.ID `json:"emoji_id,omitempty"`
	ProfileEffectID             *snowflake.ID `json:"profile_effect_id,omitempty"`
}

type ModifyGuildMemberVerificationParams struct {
	Enabled     bool                          `json:"enabled,omitzero"`
	FormFields  []MemberVerificationFormField `json:"form_fields,omitempty"`
	Description *string                       `json:"description,omitempty"`
	BulkAction  GuildJoinRequestStatus        `json:"bulk_action,omitzero"`
}

type ModifyGuildNewMemberActionParams struct {
	Icon *ImageData `json:"icon,omitempty"`
}

type ModifyGuildNewMemberWelcomeParams struct {
	Enabled          bool                    `json:"enabled,omitzero"`
	WelcomeMessage   NewMemberWelcomeMessage `json:"welcome_message,omitzero"`
	NewMemberActions []NewMemberAction       `json:"new_member_actions,omitempty"`
	ResourceChannels []ResourceChannel       `json:"resource_channels,omitempty"`
}

type ModifyGuildOnboardingParams struct {
	Prompts           []OnboardingPrompt `json:"prompts,omitempty"`
	DefaultChannelIDs []snowflake.ID     `json:"default_channel_ids,omitempty"`
	Enabled           bool               `json:"enabled,omitzero"`
	Mode              OnboardingMode     `json:"mode,omitzero"`
}

type ModifyGuildOnboardingResponsesParams struct {
	OnboardingResponses     []snowflake.ID         `json:"onboarding_responses,omitempty"`
	OnboardingPromptsSeen   map[snowflake.ID]int64 `json:"onboarding_prompts_seen,omitempty"`
	OnboardingResponsesSeen map[snowflake.ID]int64 `json:"onboarding_responses_seen,omitempty"`
}

type ModifyGuildParams struct {
	Name                        string                          `json:"name,omitzero"`
	Icon                        *ImageData                      `json:"icon,omitempty"`
	Banner                      *ImageData                      `json:"banner,omitempty"`
	HomeHeader                  *ImageData                      `json:"home_header,omitempty"`
	Splash                      *ImageData                      `json:"splash,omitempty"`
	DiscoverySplash             *ImageData                      `json:"discovery_splash,omitempty"`
	OwnerID                     snowflake.ID                    `json:"owner_id,omitzero"`
	Code                        string                          `json:"code,omitzero"`
	Description                 *string                         `json:"description,omitempty"`
	Region                      *string                         `json:"region,omitempty"`
	AFKChannelID                *snowflake.ID                   `json:"afk_channel_id,omitempty"`
	AFKTimeout                  int64                           `json:"afk_timeout,omitzero"`
	VerificationLevel           VerificationLevel               `json:"verification_level,omitzero"`
	DefaultMessageNotifications MessageNotificationLevel        `json:"default_message_notifications,omitzero"`
	ExplicitContentFilter       GuildExplicitContentFilterLevel `json:"explicit_content_filter,omitzero"`
	Features                    []string                        `json:"features,omitempty"`
	SystemChannelID             *snowflake.ID                   `json:"system_channel_id,omitempty"`
	SystemChannelFlags          SystemChannelFlags              `json:"system_channel_flags,omitzero"`
	RulesChannelID              *snowflake.ID                   `json:"rules_channel_id,omitempty"`
	PublicUpdatesChannelID      *snowflake.ID                   `json:"public_updates_channel_id,omitempty"`
	SafetyAlertsChannelID       *snowflake.ID                   `json:"safety_alerts_channel_id,omitempty"`
	PreferredLocale             string                          `json:"preferred_locale,omitzero"`
	OwnerConfiguredContentLevel NSFWLevel                       `json:"owner_configured_content_level,omitzero"`
	PremiumProgressBarEnabled   bool                            `json:"premium_progress_bar_enabled,omitzero"`
}

type ModifyGuildPowerupParams struct {
	GameServerName string       `json:"game_server_name,omitzero"`
	SKUID          snowflake.ID `json:"sku_id,omitzero"`
}

type ModifyGuildProductListingParams struct {
	Name          string       `json:"name,omitzero"`
	Description   string       `json:"description,omitzero"`
	PriceTier     int64        `json:"price_tier,omitzero"`
	CreateNewRole bool         `json:"create_new_role,omitzero"`
	Image         *ImageData   `json:"image,omitempty"`
	ImageName     string       `json:"image_name,omitzero"`
	UnlinkRole    bool         `json:"unlink_role,omitzero"`
	Attachments   []Attachment `json:"attachments,omitempty"`
}

type ModifyGuildProfileParams struct {
	Name                string          `json:"name,omitzero"`
	Icon                *ImageData      `json:"icon,omitempty"`
	Description         *string         `json:"description,omitempty"`
	BrandColorPrimary   string          `json:"brand_color_primary,omitzero"`
	GameApplicationIDs  []snowflake.ID  `json:"game_application_ids,omitempty"`
	Tag                 *string         `json:"tag,omitempty"`
	Badge               GuildBadgeType  `json:"badge"`
	BadgeColorPrimary   *string         `json:"badge_color_primary"`
	BadgeColorSecondary *string         `json:"badge_color_secondary"`
	Traits              []GuildTrait    `json:"traits,omitempty"`
	Visibility          GuildVisibility `json:"visibility,omitzero"`
	CustomBanner        *ImageData      `json:"custom_banner"`
}

type ModifyGuildResourceChannelParams struct {
	Icon *ImageData `json:"icon,omitempty"`
}

type ModifyGuildRoleParams struct {
	Name         *string      `json:"name,omitempty"`
	Description  *string      `json:"description,omitempty"`
	Color        *int64       `json:"color,omitempty"`
	Colors       *RoleColors  `json:"colors,omitempty"`
	Hoist        *bool        `json:"hoist,omitempty"`
	Icon         *ImageData   `json:"icon,omitempty"`
	UnicodeEmoji *string      `json:"unicode_emoji,omitempty"`
	Permissions  *Permissions `json:"permissions,omitempty"`
	Mentionable  *bool        `json:"mentionable,omitempty"`
}

type ModifyGuildRolePositionsParams struct {
	ID       snowflake.ID `json:"id"`
	Position *int64       `json:"position,omitempty"`
}

type ModifyGuildRoleSubscriptionGroupListingParams struct {
	Description string `json:"description,omitzero"`
}

type ModifyGuildRoleSubscriptionsSettingsParams struct {
	FullServerGate                    bool                          `json:"full_server_gate,omitzero"`
	Description                       *string                       `json:"description,omitempty"`
	ServerShopTabOrder                GuildProductShopTabsOrderType `json:"server_shop_tab_order,omitzero"`
	StorePageEnabled                  bool                          `json:"store_page_enabled,omitzero"`
	StorePagePrimaryColor             *int64                        `json:"store_page_primary_color,omitempty"`
	StorePageTrailerURL               *string                       `json:"store_page_trailer_url,omitempty"`
	StorePageShowSubscriberCount      bool                          `json:"store_page_show_subscriber_count,omitzero"`
	StorePageGuildProductsDefaultSort GuildProductSortOrderType     `json:"store_page_guild_products_default_sort,omitzero"`
}

type ModifyGuildSoundboardSoundParams struct {
	Name      string        `json:"name,omitzero"`
	Volume    *float64      `json:"volume,omitempty"`
	EmojiID   *snowflake.ID `json:"emoji_id,omitempty"`
	EmojiName *string       `json:"emoji_name,omitempty"`
}

type ModifyGuildStickerParams struct {
	Name        string  `json:"name,omitzero"`
	Description *string `json:"description,omitempty"`
	Tags        string  `json:"tags,omitzero"`
}

type ModifyGuildTemplateParams struct {
	Name        string  `json:"name,omitzero"`
	Description *string `json:"description,omitempty"`
}

type ModifyGuildVanityInviteParams struct {
	Code *string `json:"code"`
}

type ModifyGuildVanityInviteResponse struct {
	Code *string `json:"code"`
	Uses int64   `json:"uses"`
}

type ModifyGuildWelcomeScreenParams struct {
	Enabled         *bool                  `json:"enabled,omitempty"`
	Description     *string                `json:"description,omitempty"`
	WelcomeChannels []WelcomeScreenChannel `json:"welcome_channels,omitempty"`
}

type ModifyGuildWidgetParams struct {
	Enabled   bool          `json:"enabled,omitzero"`
	ChannelID *snowflake.ID `json:"channel_id,omitempty"`
}

type ModifyLinkedUserParams struct {
	LinkStatus   LinkStatus   `json:"link_status"`
	LinkedUserID snowflake.ID `json:"linked_user_id"`
}

type ModifyLobbyLinkedChannelParams struct {
	ChannelID *snowflake.ID `json:"channel_id,omitempty"`
}

type ModifyLobbyParams struct {
	Metadata           map[string]string `json:"metadata,omitempty"`
	Members            []LobbyMember     `json:"members,omitempty"`
	IdleTimeoutSeconds int64             `json:"idle_timeout_seconds,omitzero"`
	Flags              LobbyFlags        `json:"flags,omitzero"`
}

type ModifyNotificationSettingsParams struct {
	Flags *NotificationSettingsFlags `json:"flags,omitempty"`
}

type ModifyPaymentSourceParams struct {
	BillingAddress BillingAddress `json:"billing_address,omitzero"`
	Default        bool           `json:"default,omitzero"`
	ExpiresMonth   int64          `json:"expires_month,omitzero"`
	ExpiresYear    int64          `json:"expires_year,omitzero"`
}

type ModifyProfileWidgetsParams struct {
	Widgets []GameWidget `json:"widgets"`
}

type ModifyRelationshipParams struct {
	Nickname string `json:"nickname,omitzero"`
}

type ModifySKUParams struct {
	Name               json.RawMessage              `json:"name,omitempty"`
	Flags              SKUFlags                     `json:"flags,omitzero"`
	LegalNotice        json.RawMessage              `json:"legal_notice,omitempty"`
	DependentSKUID     snowflake.ID                 `json:"dependent_sku_id,omitzero"`
	BundledSKUs        []snowflake.ID               `json:"bundled_skus,omitempty"`
	AccessType         SKUAccessType                `json:"access_type,omitzero"`
	ManifestLabels     []snowflake.ID               `json:"manifest_labels,omitempty"`
	Features           []int64                      `json:"features,omitempty"`
	Locales            []string                     `json:"locales,omitempty"`
	Genres             []int64                      `json:"genres,omitempty"`
	ContentRatings     map[int64]ContentRating      `json:"content_ratings,omitempty"`
	SystemRequirements map[int64]SystemRequirements `json:"system_requirements,omitempty"`
	PriceTier          int64                        `json:"price_tier,omitzero"`
	Price              map[string]int64             `json:"price,omitempty"`
	SalePriceTier      int64                        `json:"sale_price_tier,omitzero"`
	SalePrice          map[string]int64             `json:"sale_price,omitempty"`
	ReleaseDate        string                       `json:"release_date,omitzero"`
}

type ModifyStageInstanceParams struct {
	Topic        string       `json:"topic,omitzero"`
	PrivacyLevel PrivacyLevel `json:"privacy_level,omitzero"`
}

type ModifyStreamParams struct {
	Region string `json:"region,omitzero"`
}

type ModifySubscriptionParams struct {
	Items                  []SubscriptionItem      `json:"items,omitempty"`
	PaymentSourceID        snowflake.ID            `json:"payment_source_id,omitzero"`
	PaymentSourceToken     string                  `json:"payment_source_token,omitzero"`
	ReturnURL              *string                 `json:"return_url,omitempty"`
	Currency               string                  `json:"currency,omitzero"`
	Status                 SubscriptionStatus      `json:"status,omitzero"`
	PauseDuration          int64                   `json:"pause_duration,omitzero"`
	ExpectedInvoicePrice   json.RawMessage         `json:"expected_invoice_price,omitempty"`
	ExpectedRenewalPrice   json.RawMessage         `json:"expected_renewal_price,omitempty"`
	PurchaseToken          string                  `json:"purchase_token,omitzero"`
	GatewayCheckoutContext *GatewayCheckoutContext `json:"gateway_checkout_context,omitempty"`
	LoadID                 string                  `json:"load_id,omitzero"`
}

type ModifySubscriptionPreviewParams struct {
	Items               []SubscriptionItem `json:"items,omitempty"`
	PaymentSourceID     snowflake.ID       `json:"payment_source_id,omitzero"`
	Currency            string             `json:"currency,omitzero"`
	ApplyEntitlements   bool               `json:"apply_entitlements,omitzero"`
	Renewal             bool               `json:"renewal,omitzero"`
	UserDiscountOfferID snowflake.ID       `json:"user_discount_offer_id,omitzero"`
}

type ModifySubscriptionQuery struct {
	Location      string `json:"location,omitzero"`
	LocationStack string `json:"location_stack,omitzero"`
}

type ModifyTeamMemberParams struct {
	Role string `json:"role,omitzero"`
}

type ModifyTeamParams struct {
	Name        string       `json:"name,omitzero"`
	Icon        *ImageData   `json:"icon,omitempty"`
	OwnerUserID snowflake.ID `json:"owner_user_id,omitzero"`
}

type ModifyThreadSettingsParams struct {
	Flags      int64       `json:"flags,omitzero"`
	Muted      bool        `json:"muted,omitzero"`
	MuteConfig *MuteConfig `json:"mute_config,omitempty"`
}

type ModifyUserAgreementsParams struct {
	Terms   bool `json:"terms,omitzero"`
	Privacy bool `json:"privacy,omitzero"`
}

type ModifyUserApplicationProfileConfigParams struct {
	ConnectionVisible string `json:"connection_visible,omitzero"`
}

type ModifyUserConnectionParams struct {
	Name               string         `json:"name"`
	ShowActivity       bool           `json:"show_activity"`
	FriendSync         bool           `json:"friend_sync"`
	MetadataVisibility VisibilityType `json:"metadata_visibility"`
	Visibility         VisibilityType `json:"visibility"`
}

type ModifyUserConsentsParams struct {
	Grant  []string `json:"grant"`
	Revoke []string `json:"revoke"`
}

type ModifyUserNoteParams struct {
	Note *string `json:"note"`
}

type ModifyUserProfileParams struct {
	Pronouns                    *string       `json:"pronouns,omitempty"`
	Bio                         *string       `json:"bio,omitempty"`
	Banner                      *ImageData    `json:"banner,omitempty"`
	AccentColor                 *int64        `json:"accent_color,omitempty"`
	ThemeColors                 []int64       `json:"theme_colors,omitempty"`
	PopoutAnimationParticleType *snowflake.ID `json:"popout_animation_particle_type,omitempty"`
	EmojiID                     *snowflake.ID `json:"emoji_id,omitempty"`
	ProfileEffectID             *snowflake.ID `json:"profile_effect_id,omitempty"`
}

type ModifyUserSettingsParams struct {
	ActivityJoiningRestrictedGuildIDs []snowflake.ID            `json:"activity_joining_restricted_guild_ids,omitempty"`
	ActivityRestrictedGuildIDs        []snowflake.ID            `json:"activity_restricted_guild_ids,omitempty"`
	AFKTimeout                        int64                     `json:"afk_timeout,omitzero"`
	AllowAccessibilityDetection       bool                      `json:"allow_accessibility_detection,omitzero"`
	AnimateEmoji                      bool                      `json:"animate_emoji,omitzero"`
	AnimateStickers                   StickerAnimationOption    `json:"animate_stickers,omitzero"`
	ContactSyncEnabled                bool                      `json:"contact_sync_enabled,omitzero"`
	ConvertEmoticons                  bool                      `json:"convert_emoticons,omitzero"`
	CustomStatus                      *CustomStatus             `json:"custom_status,omitempty"`
	DefaultGuildsRestricted           bool                      `json:"default_guilds_restricted,omitzero"`
	DetectPlatformAccounts            bool                      `json:"detect_platform_accounts,omitzero"`
	DeveloperMode                     bool                      `json:"developer_mode,omitzero"`
	DisableGamesTab                   bool                      `json:"disable_games_tab,omitzero"`
	EnableTTSCommand                  bool                      `json:"enable_tts_command,omitzero"`
	ExplicitContentFilter             ExplicitContentFilter     `json:"explicit_content_filter,omitzero"`
	FriendDiscoveryFlags              FriendDiscoveryFlags      `json:"friend_discovery_flags,omitzero"`
	FriendSourceFlags                 FriendSourceFlags         `json:"friend_source_flags,omitzero"`
	GIFAutoPlay                       bool                      `json:"gif_auto_play,omitzero"`
	GuildFolders                      []GuildFolder             `json:"guild_folders,omitempty"`
	InlineAttachmentMedia             bool                      `json:"inline_attachment_media,omitzero"`
	InlineEmbedMedia                  bool                      `json:"inline_embed_media,omitzero"`
	Locale                            string                    `json:"locale,omitzero"`
	MessageDisplayCompact             bool                      `json:"message_display_compact,omitzero"`
	NativePhoneIntegrationEnabled     bool                      `json:"native_phone_integration_enabled,omitzero"`
	Passwordless                      bool                      `json:"passwordless,omitzero"`
	RenderEmbeds                      bool                      `json:"render_embeds,omitzero"`
	RenderReactions                   bool                      `json:"render_reactions,omitzero"`
	RestrictedGuilds                  []snowflake.ID            `json:"restricted_guilds,omitempty"`
	ShowCurrentGame                   bool                      `json:"show_current_game,omitzero"`
	SlayerSdkReceiveDMsInGame         SlayerSdkReceiveInGameDMs `json:"slayer_sdk_receive_dms_in_game,omitzero"`
	Status                            StatusType                `json:"status,omitzero"`
	StreamNotificationsEnabled        bool                      `json:"stream_notifications_enabled,omitzero"`
	Theme                             Theme                     `json:"theme,omitzero"`
	TimezoneOffset                    int64                     `json:"timezone_offset,omitzero"`
	ViewNSFWGuilds                    bool                      `json:"view_nsfw_guilds,omitzero"`
}

type ModifyUserSettingsProtoParams struct {
	Settings            string `json:"settings"`
	RequiredDataVersion int64  `json:"required_data_version,omitzero"`
}

type ModifyUserSettingsProtoResponse struct {
	Settings  string `json:"settings"`
	OutOfDate bool   `json:"out_of_date,omitzero"`
}

type ModifyUserVoiceStateParams struct {
	ChannelID snowflake.ID `json:"channel_id"`
	Suppress  bool         `json:"suppress,omitzero"`
}

type ModifyWebauthnAuthenticatorParams struct {
	Name string `json:"name,omitzero"`
}

type ModifyWebhookParams struct {
	Name      string       `json:"name,omitzero"`
	Avatar    *ImageData   `json:"avatar,omitempty"`
	ChannelID snowflake.ID `json:"channel_id,omitzero"`
}

type ModifyWebhookWithTokenParams struct {
	Name   string     `json:"name,omitzero"`
	Avatar *ImageData `json:"avatar,omitempty"`
}

type MonetizationStorePage struct {
	Guild            StorePageGuild            `json:"guild"`
	RoleSubscription StorePageRoleSubscription `json:"role_subscription"`
}

type MuteConfig struct {
	EndTime            *time.Time `json:"end_time,omitempty"`
	SelectedTimeWindow int64      `json:"selected_time_window,omitzero"`
}

type MutualGuild struct {
	ID   snowflake.ID `json:"id"`
	Nick *string      `json:"nick"`
}

type NameplateCollectibleItem struct {
	Type    CollectibleProductType `json:"type"`
	SKUID   snowflake.ID           `json:"sku_id"`
	Palette NameplateColorPalette  `json:"palette"`
	Asset   string                 `json:"asset"`
	Assets  ItemAssets             `json:"assets"`
	Label   string                 `json:"label"`
}

type NameplateData struct {
	Asset     string                `json:"asset"`
	SKUID     snowflake.ID          `json:"sku_id"`
	Label     string                `json:"label"`
	Palette   NameplateColorPalette `json:"palette"`
	ExpiresAt *int64                `json:"expires_at"`
}

type NetworkingCreateTokenResponse struct {
	Address string `json:"address"`
	Token   string `json:"token"`
}

type NewMemberAction struct {
	ChannelID   snowflake.ID        `json:"channel_id"`
	ActionType  NewMemberActionType `json:"action_type"`
	Title       string              `json:"title"`
	Description string              `json:"description"`
	Emoji       Emoji               `json:"emoji,omitzero"`
	Icon        string              `json:"icon,omitzero"`
}

type NewMemberActionProgress struct {
	Completed bool `json:"completed"`
}

type NewMemberActionsProgress struct {
	GuildID        snowflake.ID                             `json:"guild_id"`
	UserID         snowflake.ID                             `json:"user_id"`
	ChannelActions map[snowflake.ID]NewMemberActionProgress `json:"channel_actions"`
}

type NewMemberWelcome struct {
	GuildID          snowflake.ID            `json:"guild_id"`
	Enabled          bool                    `json:"enabled"`
	WelcomeMessage   NewMemberWelcomeMessage `json:"welcome_message"`
	NewMemberActions []NewMemberAction       `json:"new_member_actions"`
	ResourceChannels []ResourceChannel       `json:"resource_channels"`
}

type NewMemberWelcomeMessage struct {
	AuthorIDs []snowflake.ID `json:"author_ids"`
	Message   string         `json:"message"`
}

type NotificationCenterItem struct {
	ID                     snowflake.ID                `json:"id"`
	BundleID               string                      `json:"bundle_id"`
	Type                   NotificationCenterItemType  `json:"type"`
	ItemEnum               *NotificationCenterItemEnum `json:"item_enum"`
	Body                   string                      `json:"body"`
	Acked                  bool                        `json:"acked"`
	Deeplink               string                      `json:"deeplink"`
	IconURL                *string                     `json:"icon_url"`
	IconName               *string                     `json:"icon_name"`
	OtherUser              *PartialUser                `json:"other_user"`
	Message                *Message                    `json:"message"`
	Completed              bool                        `json:"completed"`
	GuildID                *snowflake.ID               `json:"guild_id"`
	MessageID              *snowflake.ID               `json:"message_id"`
	MessageChannelID       *snowflake.ID               `json:"message_channel_id"`
	GuildScheduledEventID  *snowflake.ID               `json:"guild_scheduled_event_id"`
	DisableAction          bool                        `json:"disable_action"`
	Callout                *string                     `json:"callout"`
	Application            *PartialApplication         `json:"application"`
	EmojiID                *snowflake.ID               `json:"emoji_id"`
	EmojiName              *string                     `json:"emoji_name"`
	MessageContent         *string                     `json:"message_content"`
	MessageEmbedCount      *int64                      `json:"message_embed_count"`
	MessageAttachmentCount *int64                      `json:"message_attachment_count"`
	MessageStickerCount    *int64                      `json:"message_sticker_count"`
	IsVoiceMessage         bool                        `json:"is_voice_message"`
}

type NotificationCenterItemCompleted struct {
	ItemEnum NotificationCenterItemEnum `json:"item_enum"`
}

type NotificationCenterItemDelete struct {
	ID     snowflake.ID `json:"id"`
	UserID snowflake.ID `json:"user_id"`
}

type NotificationCenterItemsAck struct {
	ID snowflake.ID `json:"id"`
}

type NotificationCreate struct {
	ChannelID snowflake.ID `json:"channel_id"`
	Message   RPCMessage   `json:"message"`
	IconURL   *string      `json:"icon_url"`
	Title     string       `json:"title"`
	Body      string       `json:"body"`
}

type NotificationSettings struct {
	Flags NotificationSettingsFlags `json:"flags"`
}

type NotificationSettingsSnapshot struct {
	ID         snowflake.ID `json:"id"`
	Label      *string      `json:"label"`
	RecordedAt time.Time    `json:"recorded_at"`
	Length     int64        `json:"length"`
}

type OAuth2Authorization struct {
	ID          snowflake.ID       `json:"id"`
	Scopes      []string           `json:"scopes"`
	Application PartialApplication `json:"application"`
	Disclosures []int64            `json:"disclosures,omitempty"`
}

type OAuth2Guild struct {
	ID          snowflake.ID `json:"id"`
	Name        string       `json:"name"`
	Icon        *string      `json:"icon"`
	MFALevel    MFALevel     `json:"mfa_level"`
	Permissions Permissions  `json:"permissions"`
}

type OAuth2LocationContext struct {
	GuildID     snowflake.ID `json:"guild_id,omitzero"`
	ChannelID   snowflake.ID `json:"channel_id,omitzero"`
	ChannelType ChannelType  `json:"channel_type,omitzero"`
}

type OAuth2TokenDelete struct {
	ID            snowflake.ID `json:"id"`
	ApplicationID snowflake.ID `json:"application_id"`
}

type OAuth2TokenRevoke struct {
	AccessToken   string       `json:"access_token"`
	ApplicationID snowflake.ID `json:"application_id,omitzero"`
}

type Onboarding struct {
	GuildID                 snowflake.ID           `json:"guild_id"`
	Prompts                 []OnboardingPrompt     `json:"prompts"`
	DefaultChannelIDs       []snowflake.ID         `json:"default_channel_ids"`
	Enabled                 bool                   `json:"enabled"`
	BelowRequirements       bool                   `json:"below_requirements"`
	Mode                    OnboardingMode         `json:"mode"`
	Connections             []OnboardingConnection `json:"connections,omitempty"`
	Responses               []snowflake.ID         `json:"responses,omitempty"`
	OnboardingPromptsSeen   map[snowflake.ID]int64 `json:"onboarding_prompts_seen,omitempty"`
	OnboardingResponsesSeen map[snowflake.ID]int64 `json:"onboarding_responses_seen,omitempty"`
}

type OnboardingConnection struct {
	ConnectionType OnboardingConnectionType `json:"connection_type"`
	ApplicationID  *snowflake.ID            `json:"application_id"`
	ProviderID     *ConnectionType          `json:"provider_id"`
	Description    string                   `json:"description"`
}

type OnboardingPrompt struct {
	ID           snowflake.ID             `json:"id"`
	Type         OnboardingPromptType     `json:"type"`
	Options      []OnboardingPromptOption `json:"options"`
	Title        string                   `json:"title"`
	SingleSelect bool                     `json:"single_select"`
	Required     bool                     `json:"required"`
	InOnboarding bool                     `json:"in_onboarding"`
}

type OnboardingPromptOption struct {
	ID          snowflake.ID   `json:"id"`
	ChannelIDs  []snowflake.ID `json:"channel_ids"`
	RoleIDs     []snowflake.ID `json:"role_ids"`
	Emoji       Emoji          `json:"emoji"`
	Title       string         `json:"title"`
	Description *string        `json:"description"`
}

type OnboardingResponses struct {
	GuildID                 snowflake.ID           `json:"guild_id"`
	UserID                  snowflake.ID           `json:"user_id"`
	OnboardingResponses     []snowflake.ID         `json:"onboarding_responses"`
	OnboardingPromptsSeen   map[snowflake.ID]int64 `json:"onboarding_prompts_seen"`
	OnboardingResponsesSeen map[snowflake.ID]int64 `json:"onboarding_responses_seen"`
}

type OneTimeLoginParams struct {
	Ticket string `json:"ticket"`
}

type OneTimeLoginResponse struct {
	Token string `json:"token"`
}

type OpenExternalLinkArguments struct {
	URL string `json:"url"`
}

type OpenMessageArguments struct {
	GuildID   *snowflake.ID `json:"guild_id,omitempty"`
	ChannelID snowflake.ID  `json:"channel_id"`
	MessageID snowflake.ID  `json:"message_id"`
	Pid       int64         `json:"pid"`
}

type OpenOverlayActivityInviteArguments struct {
	Type ActivityActionType `json:"type"`
	Pid  int64              `json:"pid"`
}

type OpenOverlayGuildInviteArguments struct {
	Code string `json:"code"`
	Pid  int64  `json:"pid"`
}

type OpenOverlayVoiceSettingsArguments struct {
	Pid int64 `json:"pid"`
}

type OpenShareMomentDialogArguments struct {
	Mediaurl string `json:"mediaUrl"`
}

type OrientationUpdate struct {
	ScreenOrientation OrientationState `json:"screen_orientation"`
}

type OutgoingPayload struct {
	Cmd   string          `json:"cmd"`
	Args  json.RawMessage `json:"args"`
	Nonce string          `json:"nonce"`
	Evt   string          `json:"evt,omitzero"`
}

type OverlayUpdate struct {
	Enabled bool `json:"enabled"`
	Locked  bool `json:"locked"`
}

type OverlayUpdateSubscriptionArguments struct {
	Pid int64 `json:"pid"`
}

type Packet struct {
	Opcode  json.RawMessage `json:"Opcode"`
	Length  json.RawMessage `json:"Length"`
	Payload json.RawMessage `json:"Payload"`
}

type Pagination struct {
	Offset  int64 `json:"offset"`
	Limit   int64 `json:"limit"`
	Total   int64 `json:"total"`
	HasMore bool  `json:"has_more"`
}

type Pan struct {
	Left  float64 `json:"left"`
	Right float64 `json:"right"`
}

type ParaphraseThreadTitleParams struct {
	Content string `json:"content"`
}

type ParaphraseThreadTitleResponse struct {
	Title string `json:"title"`
}

type PartialApplication struct {
	ID                          snowflake.ID                                       `json:"id"`
	Name                        string                                             `json:"name"`
	Description                 string                                             `json:"description"`
	Icon                        *string                                            `json:"icon"`
	CoverImage                  string                                             `json:"cover_image,omitzero"`
	Splash                      string                                             `json:"splash,omitzero"`
	Type                        *ApplicationType                                   `json:"type"`
	Flags                       ApplicationFlags                                   `json:"flags"`
	FlagsNew                    string                                             `json:"flags_new"`
	PrimarySKU                  SKU                                                `json:"primary_sku,omitzero"`
	PrimarySKUID                snowflake.ID                                       `json:"primary_sku_id,omitzero"`
	VerifyKey                   string                                             `json:"verify_key"`
	GuildID                     snowflake.ID                                       `json:"guild_id,omitzero"`
	Guild                       PartialGuild                                       `json:"guild,omitzero"`
	EulaID                      snowflake.ID                                       `json:"eula_id,omitzero"`
	Slug                        string                                             `json:"slug,omitzero"`
	Aliases                     []string                                           `json:"aliases,omitempty"`
	Executables                 []ApplicationExecutable                            `json:"executables,omitempty"`
	ThirdPartySKUs              []ApplicationSKU                                   `json:"third_party_skus,omitempty"`
	Hook                        bool                                               `json:"hook"`
	Overlay                     bool                                               `json:"overlay,omitzero"`
	OverlayMethods              OverlayMethodFlags                                 `json:"overlay_methods,omitzero"`
	OverlayWarn                 bool                                               `json:"overlay_warn,omitzero"`
	OverlayCompatibilityHook    bool                                               `json:"overlay_compatibility_hook,omitzero"`
	Bot                         PartialUser                                        `json:"bot,omitzero"`
	Team                        *Team                                              `json:"team,omitempty"`
	Developers                  []Company                                          `json:"developers,omitempty"`
	Publishers                  []Company                                          `json:"publishers,omitempty"`
	RPCOrigins                  []string                                           `json:"rpc_origins,omitempty"`
	DeeplinkURI                 string                                             `json:"deeplink_uri,omitzero"`
	IntegrationPublic           bool                                               `json:"integration_public,omitzero"`
	IntegrationRequireCodeGrant bool                                               `json:"integration_require_code_grant,omitzero"`
	BotPublic                   bool                                               `json:"bot_public,omitzero"`
	BotRequireCodeGrant         bool                                               `json:"bot_require_code_grant,omitzero"`
	TermsOfServiceURL           *string                                            `json:"terms_of_service_url,omitempty"`
	PrivacyPolicyURL            *string                                            `json:"privacy_policy_url,omitempty"`
	Tags                        []string                                           `json:"tags,omitempty"`
	InstallParams               ApplicationInstallParams                           `json:"install_params,omitzero"`
	CustomInstallURL            string                                             `json:"custom_install_url,omitzero"`
	IntegrationTypesConfig      map[int64]*ApplicationIntegrationTypeConfiguration `json:"integration_types_config,omitempty"`
	ConnectionEntrypointURL     string                                             `json:"connection_entrypoint_url,omitzero"`
	IsVerified                  bool                                               `json:"is_verified"`
	IsDiscoverable              bool                                               `json:"is_discoverable"`
	IsMonetized                 bool                                               `json:"is_monetized"`
	StorefrontAvailable         bool                                               `json:"storefront_available"`
	MaxParticipants             int64                                              `json:"max_participants,omitzero"`
	EmbeddedActivityConfig      EmbeddedActivityConfig                             `json:"embedded_activity_config,omitzero"`
	ParentID                    snowflake.ID                                       `json:"parent_id,omitzero"`
	Categories                  json.RawMessage                                    `json:"categories,omitempty"`
	DirectoryEntry              DirectoryEntry                                     `json:"directory_entry,omitzero"`
	Position                    int64                                              `json:"position,omitzero"`
	GameDataOverrides           GameDataOverrides                                  `json:"game_data_overrides,omitzero"`
}

type PartialApplicationIdentity struct {
	UserID         snowflake.ID `json:"user_id"`
	ExternalUserID string       `json:"external_user_id"`
}

type PartialChannel struct {
	ID         snowflake.ID  `json:"id"`
	Type       ChannelType   `json:"type"`
	Name       *string       `json:"name"`
	Recipients []PartialUser `json:"recipients,omitempty"`
	Icon       *string       `json:"icon,omitempty"`
	GuildID    *snowflake.ID `json:"guild_id,omitempty"`
}

type PartialConnection struct {
	ID       string          `json:"id"`
	Type     ConnectionType  `json:"type"`
	Name     string          `json:"name"`
	Verified bool            `json:"verified"`
	Metadata json.RawMessage `json:"metadata,omitempty"`
}

type PartialGuild struct {
	ID                       snowflake.ID `json:"id"`
	Name                     string       `json:"name"`
	Icon                     *string      `json:"icon"`
	Description              *string      `json:"description"`
	Splash                   *string      `json:"splash"`
	DiscoverySplash          *string      `json:"discovery_splash"`
	HomeHeader               *string      `json:"home_header"`
	Features                 []string     `json:"features"`
	Emojis                   []Emoji      `json:"emojis,omitempty"`
	Stickers                 []Sticker    `json:"stickers,omitempty"`
	ApproximateMemberCount   int64        `json:"approximate_member_count,omitzero"`
	ApproximatePresenceCount int64        `json:"approximate_presence_count,omitzero"`
	OwnerID                  snowflake.ID `json:"owner_id,omitzero"`
}

type PartialGuildUpdates struct {
	Channels          []Channel      `json:"channels,omitempty"`
	DeletedChannelIDs []snowflake.ID `json:"deleted_channel_ids,omitempty"`
	Emojis            []Emoji        `json:"emojis,omitempty"`
	DeletedEmojiIDs   []snowflake.ID `json:"deleted_emoji_ids,omitempty"`
	Roles             []Role         `json:"roles,omitempty"`
	DeletedRoleIDs    []snowflake.ID `json:"deleted_role_ids,omitempty"`
	Stickers          []Sticker      `json:"stickers,omitempty"`
	DeletedStickerIDs []snowflake.ID `json:"deleted_sticker_ids,omitempty"`
}

type PartialIntegration struct {
	ID            snowflake.ID    `json:"id"`
	Name          string          `json:"name"`
	Type          IntegrationType `json:"type"`
	Account       json.RawMessage `json:"account"`
	ApplicationID snowflake.ID    `json:"application_id,omitzero"`
}

type PartialMessage struct {
	ID                  snowflake.ID           `json:"id"`
	LobbyID             snowflake.ID           `json:"lobby_id,omitzero"`
	ChannelID           snowflake.ID           `json:"channel_id"`
	Type                MessageType            `json:"type,omitzero"`
	Content             string                 `json:"content"`
	Author              PartialUser            `json:"author"`
	Activity            MessageActivity        `json:"activity,omitzero"`
	Application         IntegrationApplication `json:"application,omitzero"`
	ApplicationID       snowflake.ID           `json:"application_id,omitzero"`
	ParentApplicationID snowflake.ID           `json:"parent_application_id,omitzero"`
	Flags               MessageFlags           `json:"flags,omitzero"`
	Channel             json.RawMessage        `json:"channel,omitempty"`
	RecipientID         snowflake.ID           `json:"recipient_id,omitzero"`
	ModerationMetadata  map[string]string      `json:"moderation_metadata,omitempty"`
}

type PartialMessageInteraction struct {
	ID   snowflake.ID    `json:"id"`
	Type InteractionType `json:"type"`
	Name string          `json:"name"`
	User PartialUser     `json:"user"`
}

type PartialQuest struct {
	ID            snowflake.ID  `json:"id"`
	ReplacementID *snowflake.ID `json:"replacement_id,omitempty"`
}

type PartialRPCChannel struct {
	ID   snowflake.ID `json:"id"`
	Name string       `json:"name"`
	Type ChannelType  `json:"type"`
}

type PartialRPCGuildMember struct {
	User     RPCUser    `json:"user"`
	Nick     *string    `json:"nick"`
	Status   StatusType `json:"status"`
	Activity *Activity  `json:"activity"`
}

type PartialRPCMessage struct {
	ID snowflake.ID `json:"id"`
}

type PartialRPCVoiceSettingsMode struct {
	Type          RPCVoiceSettingsModeType `json:"type,omitzero"`
	AutoThreshold bool                     `json:"auto_threshold,omitzero"`
	Threshold     int64                    `json:"threshold,omitzero"`
	Shortcut      []ShortcutKeyCombo       `json:"shortcut,omitempty"`
	Delay         int64                    `json:"delay,omitzero"`
}

type PartialRelationship struct {
	ID              snowflake.ID     `json:"id"`
	Type            RelationshipType `json:"type"`
	Nickname        *string          `json:"nickname"`
	Since           time.Time        `json:"since,omitzero"`
	StrangerRequest bool             `json:"stranger_request,omitzero"`
	UserIgnored     bool             `json:"user_ignored"`
}

type PartialRole struct {
	ID   snowflake.ID `json:"id"`
	Name string       `json:"name"`
}

type PartialSubscription struct {
	ID                           snowflake.ID       `json:"id"`
	Type                         SubscriptionType   `json:"type"`
	PaymentGateway               *PaymentGateway    `json:"payment_gateway"`
	Currency                     string             `json:"currency"`
	Items                        []SubscriptionItem `json:"items"`
	PaymentGatewayPlanID         *string            `json:"payment_gateway_plan_id"`
	PaymentGatewaySubscriptionID *string            `json:"payment_gateway_subscription_id,omitempty"`
	CurrentPeriodStart           time.Time          `json:"current_period_start"`
	CurrentPeriodEnd             time.Time          `json:"current_period_end"`
	StreakStartedAt              time.Time          `json:"streak_started_at,omitzero"`
}

type PartialSubscriptionPlan struct {
	ID            snowflake.ID         `json:"id"`
	Name          string               `json:"name"`
	SKUID         snowflake.ID         `json:"sku_id"`
	Interval      SubscriptionInterval `json:"interval"`
	IntervalCount int64                `json:"interval_count"`
	TaxInclusive  bool                 `json:"tax_inclusive"`
}

type PartialSubscriptionTrial struct {
	Interval      SubscriptionInterval `json:"interval"`
	IntervalCount int64                `json:"interval_count"`
}

type PartialUser struct {
	ID                   snowflake.ID          `json:"id"`
	Username             string                `json:"username"`
	Discriminator        string                `json:"discriminator"`
	GlobalName           *string               `json:"global_name,omitempty"`
	Avatar               *string               `json:"avatar"`
	AvatarDecorationData *AvatarDecorationData `json:"avatar_decoration_data,omitempty"`
	Collectibles         *Collectibles         `json:"collectibles,omitempty"`
	DisplayNameStyles    *DisplayNameStyle     `json:"display_name_styles,omitempty"`
	PrimaryGuild         *PrimaryGuild         `json:"primary_guild,omitempty"`
	Bot                  bool                  `json:"bot,omitzero"`
	System               bool                  `json:"system,omitzero"`
	Banner               *string               `json:"banner,omitempty"`
	AccentColor          *int64                `json:"accent_color,omitempty"`
	PublicFlags          UserFlags             `json:"public_flags,omitzero"`
}

type PartialUserApplicationProfile struct {
	Username          *string                    `json:"username"`
	Metadata          *string                    `json:"metadata"`
	Data              UserApplicationProfileData `json:"data,omitzero"`
	DataTrusted       bool                       `json:"data_trusted,omitzero"`
	ConnectionVisible bool                       `json:"connection_visible"`
}

type PartialUserGuildSettings struct {
	ChannelOverrides     map[snowflake.ID]ChannelOverride `json:"channel_overrides,omitempty"`
	Flags                UserGuildSettingsFlags           `json:"flags,omitzero"`
	HideMutedChannels    bool                             `json:"hide_muted_channels,omitzero"`
	MessageNotifications MessageNotificationLevel         `json:"message_notifications,omitzero"`
	MobilePush           bool                             `json:"mobile_push,omitzero"`
	MuteScheduledEvents  bool                             `json:"mute_scheduled_events,omitzero"`
	Muted                bool                             `json:"muted,omitzero"`
	MuteConfig           *MuteConfig                      `json:"mute_config,omitempty"`
	NotifyHighlights     HighlightLevel                   `json:"notify_highlights,omitzero"`
	SuppressEveryone     bool                             `json:"suppress_everyone,omitzero"`
	SuppressRoles        bool                             `json:"suppress_roles,omitzero"`
}

type PassiveUpdateV1 struct {
	GuildID     snowflake.ID    `json:"guild_id"`
	Channels    []ChannelUnread `json:"channels"`
	VoiceStates []VoiceState    `json:"voice_states,omitempty"`
	Members     []GuildMember   `json:"members,omitempty"`
}

type PassiveUpdateV2 struct {
	GuildID            snowflake.ID    `json:"guild_id"`
	UpdatedChannels    []ChannelUnread `json:"updated_channels"`
	UpdatedVoiceStates []VoiceState    `json:"updated_voice_states"`
	RemovedVoiceStates []snowflake.ID  `json:"removed_voice_states"`
	UpdatedMembers     []GuildMember   `json:"updated_members"`
}

type Payment struct {
	ID                            snowflake.ID        `json:"id"`
	Amount                        int64               `json:"amount"`
	Tax                           int64               `json:"tax"`
	TaxInclusive                  bool                `json:"tax_inclusive"`
	Currency                      string              `json:"currency"`
	AmountRefunded                int64               `json:"amount_refunded"`
	Description                   string              `json:"description"`
	Status                        PaymentStatus       `json:"status"`
	CreatedAt                     time.Time           `json:"created_at"`
	SKUID                         snowflake.ID        `json:"sku_id,omitzero"`
	SKUPrice                      int64               `json:"sku_price,omitzero"`
	SKUSubscriptionPlanID         snowflake.ID        `json:"sku_subscription_plan_id,omitzero"`
	PaymentGateway                PaymentGateway      `json:"payment_gateway,omitzero"`
	PaymentGatewayPaymentID       string              `json:"payment_gateway_payment_id,omitzero"`
	HasInvoiceURL                 bool                `json:"has_invoice_url,omitzero"`
	HasRefundInvoiceURLs          bool                `json:"has_refund_invoice_urls,omitzero"`
	DownloadableInvoice           string              `json:"downloadable_invoice,omitzero"`
	DownloadableRefundInvoices    []string            `json:"downloadable_refund_invoices,omitempty"`
	RefundDisqualificationReasons []int64             `json:"refund_disqualification_reasons,omitempty"`
	Flags                         PaymentFlags        `json:"flags"`
	SKU                           SKU                 `json:"sku,omitzero"`
	PaymentSource                 PaymentSource       `json:"payment_source,omitzero"`
	Subscription                  PartialSubscription `json:"subscription,omitzero"`
	Metadata                      PaymentMetadata     `json:"metadata"`
}

type PaymentMetadata struct {
	BillingErrorCode *int64 `json:"billing_error_code"`
}

type PaymentSource struct {
	ID                     snowflake.ID       `json:"id"`
	Brand                  string             `json:"brand,omitzero"`
	Country                string             `json:"country,omitzero"`
	Last4                  string             `json:"last_4,omitzero"`
	BillingAddress         BillingAddress     `json:"billing_address"`
	Type                   PaymentSourceType  `json:"type"`
	PaymentGateway         PaymentGateway     `json:"payment_gateway"`
	PaymentGatewaySourceID string             `json:"payment_gateway_source_id,omitzero"`
	Default                bool               `json:"default"`
	Invalid                bool               `json:"invalid"`
	Flags                  PaymentSourceFlags `json:"flags"`
	DeletedAt              *time.Time         `json:"deleted_at"`
	ExpiresMonth           int64              `json:"expires_month,omitzero"`
	ExpiresYear            int64              `json:"expires_year,omitzero"`
	Email                  string             `json:"email,omitzero"`
	Bank                   string             `json:"bank,omitzero"`
	Username               string             `json:"username,omitzero"`
}

type PendingLogin struct {
	Ticket string `json:"ticket"`
}

type PendingRemoteInit struct {
	Fingerprint string `json:"fingerprint"`
}

type PendingTicket struct {
	EncryptedUserPayload string `json:"encrypted_user_payload"`
}

type PermissionOverwrite struct {
	ID    snowflake.ID            `json:"id"`
	Type  PermissionOverwriteType `json:"type"`
	Allow Permissions             `json:"allow"`
	Deny  Permissions             `json:"deny"`
}

type PingLobbyVoiceServer struct {
	LobbyID snowflake.ID `json:"lobby_id"`
}

type PingStreamServer struct {
	StreamKey string `json:"stream_key"`
}

type PixPaymentMetadata struct {
	TaxID string `json:"tax_id,omitzero"`
}

type PlaySoundboardSoundArguments struct {
	GuildID snowflake.ID `json:"guild_id,omitzero"`
	SoundID snowflake.ID `json:"sound_id,omitzero"`
}

type PlayedSoundFrecency struct {
	PlayedSounds map[string]FrecencyItem `json:"played_sounds"`
}

type Poll struct {
	Question         PollMedia       `json:"question"`
	Answers          []PollAnswer    `json:"answers"`
	Expiry           json.RawMessage `json:"expiry"`
	AllowMultiselect bool            `json:"allow_multiselect"`
	LayoutType       PollLayoutType  `json:"layout_type"`
	Results          PollResults     `json:"results,omitzero"`
}

type PollAnswer struct {
	AnswerID  int64     `json:"answer_id"`
	PollMedia PollMedia `json:"poll_media"`
}

type PollAnswerCount struct {
	ID      int64 `json:"id"`
	Count   int64 `json:"count"`
	MeVoted bool  `json:"me_voted"`
}

type PollCreate struct {
	Question         PollMedia      `json:"question"`
	Answers          []PollAnswer   `json:"answers"`
	Duration         int64          `json:"duration"`
	AllowMultiselect bool           `json:"allow_multiselect,omitzero"`
	LayoutType       PollLayoutType `json:"layout_type,omitzero"`
}

type PollMedia struct {
	Text  string `json:"text,omitzero"`
	Emoji Emoji  `json:"emoji,omitzero"`
}

type PollResultEmbed struct {
	PollQuestionText          string       `json:"poll_question_text"`
	TotalVotes                int64        `json:"total_votes"`
	VictorAnswerID            int64        `json:"victor_answer_id,omitzero"`
	VictorAnswerText          string       `json:"victor_answer_text,omitzero"`
	VictorAnswerEmojiID       snowflake.ID `json:"victor_answer_emoji_id,omitzero"`
	VictorAnswerEmojiName     string       `json:"victor_answer_emoji_name,omitzero"`
	VictorAnswerEmojiAnimated bool         `json:"victor_answer_emoji_animated,omitzero"`
	VictorAnswerVotes         int64        `json:"victor_answer_votes"`
}

type PollResults struct {
	IsFinalized  bool              `json:"is_finalized"`
	AnswerCounts []PollAnswerCount `json:"answer_counts"`
}

type Potion struct {
	UsedBy    snowflake.ID `json:"used_by"`
	Type      PotionType   `json:"type"`
	Emoji     []Emoji      `json:"emoji"`
	CreatedAt time.Time    `json:"created_at"`
}

type PreloadedUserSettings struct {
	Versions              Versions                              `json:"versions"`
	Inbox                 InboxSettings                         `json:"inbox"`
	Guilds                AllGuildSettings                      `json:"guilds"`
	UserContent           UserContentSettings                   `json:"user_content"`
	VoiceAndVideo         VoiceAndVideoSettings                 `json:"voice_and_video"`
	TextAndImages         TextAndImagesSettings                 `json:"text_and_images"`
	Notifications         UserSettingsProtoNotificationSettings `json:"notifications"`
	Privacy               PrivacySettings                       `json:"privacy"`
	Debug                 DebugSettings                         `json:"debug"`
	GameLibrary           GameLibrarySettings                   `json:"game_library"`
	Status                StatusSettings                        `json:"status"`
	Localization          LocalizationSettings                  `json:"localization"`
	Appearance            AppearanceSettings                    `json:"appearance"`
	GuildFolders          GuildFolders                          `json:"guild_folders"`
	Favorites             Favorites                             `json:"favorites"`
	AudioContextSettings  AudioSettings                         `json:"audio_context_settings"`
	Communities           CommunitiesSettings                   `json:"communities"`
	Broadcast             BroadcastSettings                     `json:"broadcast"`
	Clips                 ClipsSettings                         `json:"clips"`
	ForLater              ForLaterSettings                      `json:"for_later"`
	SafetySettings        SafetySettings                        `json:"safety_settings"`
	IcymiSettings         IcymiSettings                         `json:"icymi_settings"`
	Applications          AllApplicationSettings                `json:"applications"`
	Ads                   AdsSettings                           `json:"ads"`
	InAppFeedbackSettings InAppFeedbackSettings                 `json:"in_app_feedback_settings"`
	AppVersionSettings    AppVersionSettings                    `json:"app_version_settings"`
}

type PremiumGuildSubscription struct {
	ID          snowflake.ID `json:"id"`
	GuildID     snowflake.ID `json:"guild_id"`
	UserID      snowflake.ID `json:"user_id"`
	Ended       bool         `json:"ended"`
	EndsAt      time.Time    `json:"ends_at,omitzero"`
	PauseEndsAt *time.Time   `json:"pause_ends_at"`
	User        PartialUser  `json:"user"`
}

type PremiumGuildSubscriptionSlot struct {
	ID                       snowflake.ID              `json:"id"`
	SubscriptionID           snowflake.ID              `json:"subscription_id"`
	Canceled                 bool                      `json:"canceled"`
	CooldownEndsAt           *string                   `json:"cooldown_ends_at"`
	PremiumGuildSubscription *PremiumGuildSubscription `json:"premium_guild_subscription"`
}

type PremiumPrice struct {
	Amount     int64 `json:"amount"`
	Percentage int64 `json:"percentage"`
}

type PremiumReferral struct {
	ID                snowflake.ID      `json:"id"`
	UserID            snowflake.ID      `json:"user_id"`
	TrialID           snowflake.ID      `json:"trial_id"`
	SubscriptionTrial SubscriptionTrial `json:"subscription_trial"`
	ExpiresAt         time.Time         `json:"expires_at"`
	ReferrerID        snowflake.ID      `json:"referrer_id"`
	Referrer          PartialUser       `json:"referrer"`
	RedeemedAt        time.Time         `json:"redeemed_at,omitzero"`
}

type PremiumReferralEligibility struct {
	ReferralsRemaining      int64                          `json:"referrals_remaining"`
	SentUserIDs             []snowflake.ID                 `json:"sent_user_ids"`
	RefreshAt               *time.Time                     `json:"refresh_at"`
	HasEligibleFriends      bool                           `json:"has_eligible_friends"`
	RecipientStatus         map[snowflake.ID]int64         `json:"recipient_status"`
	IsEligibleForIncentive  bool                           `json:"is_eligible_for_incentive"`
	IsQualifiedForIncentive bool                           `json:"is_qualified_for_incentive"`
	ReferralIncentiveStatus PremiumReferralIncentiveStatus `json:"referral_incentive_status"`
}

type PremiumState struct {
	PremiumSource                PremiumSource                `json:"premium_source"`
	PremiumSubscriptionType      PremiumSubscriptionType      `json:"premium_subscription_type"`
	PremiumSubscriptionGroupRole PremiumSubscriptionGroupRole `json:"premium_subscription_group_role,omitzero"`
}

type PremiumUsage struct {
	Value int64 `json:"value"`
}

type Presence struct {
	User             PartialUser  `json:"user"`
	GuildID          snowflake.ID `json:"guild_id,omitzero"`
	Status           StatusType   `json:"status"`
	Activities       []Activity   `json:"activities"`
	HiddenActivities []Activity   `json:"hidden_activities,omitempty"`
	ClientStatus     ClientStatus `json:"client_status"`
	HasPlayedGame    bool         `json:"has_played_game,omitzero"`
}

type PreviewOAuth2AuthorizationQuery struct {
	ClientID            snowflake.ID               `json:"client_id"`
	ResponseType        ResponseType               `json:"response_type,omitzero"`
	Scope               string                     `json:"scope"`
	RedirectURI         string                     `json:"redirect_uri,omitzero"`
	State               string                     `json:"state,omitzero"`
	Nonce               string                     `json:"nonce,omitzero"`
	CodeChallenge       string                     `json:"code_challenge,omitzero"`
	CodeChallengeMethod string                     `json:"code_challenge_method,omitzero"`
	IntegrationType     ApplicationIntegrationType `json:"integration_type,omitzero"`
}

type PreviewOAuth2AuthorizationResponse struct {
	Application     PartialApplication         `json:"application"`
	User            PartialUser                `json:"user"`
	Authorized      bool                       `json:"authorized"`
	IntegrationType ApplicationIntegrationType `json:"integration_type"`
	RedirectURI     *string                    `json:"redirect_uri,omitempty"`
	Bot             PartialUser                `json:"bot,omitzero"`
	Guilds          []OAuth2Guild              `json:"guilds,omitempty"`
}

type PrimaryGuild struct {
	IdentityEnabled *bool         `json:"identity_enabled"`
	IdentityGuildID *snowflake.ID `json:"identity_guild_id"`
	Tag             *string       `json:"tag"`
	Badge           *string       `json:"badge"`
}

type PrivacySettings struct {
	AllowActivityPartyPrivacyFriends      bool                               `json:"allow_activity_party_privacy_friends,omitzero"`
	AllowActivityPartyPrivacyVoiceChannel bool                               `json:"allow_activity_party_privacy_voice_channel,omitzero"`
	RestrictedGuildIDs                    []json.RawMessage                  `json:"restricted_guild_ids"`
	DefaultGuildsRestricted               bool                               `json:"default_guilds_restricted"`
	AllowAccessibilityDetection           bool                               `json:"allow_accessibility_detection"`
	DetectPlatformAccounts                bool                               `json:"detect_platform_accounts,omitzero"`
	Passwordless                          bool                               `json:"passwordless,omitzero"`
	ContactSyncEnabled                    bool                               `json:"contact_sync_enabled,omitzero"`
	FriendSourceFlags                     UserSettingsProtoFriendSourceFlags `json:"friend_source_flags,omitzero"`
	FriendDiscoveryFlags                  FriendDiscoveryFlags               `json:"friend_discovery_flags,omitzero"`
	ActivityRestrictedGuildIDs            []json.RawMessage                  `json:"activity_restricted_guild_ids"`
	DefaultGuildsActivityRestricted       json.RawMessage                    `json:"default_guilds_activity_restricted,omitempty"`
	ActivityJoiningRestrictedGuildIDs     []json.RawMessage                  `json:"activity_joining_restricted_guild_ids"`
	MessageRequestRestrictedGuildIDs      []json.RawMessage                  `json:"message_request_restricted_guild_ids"`
	DefaultMessageRequestRestricted       bool                               `json:"default_message_request_restricted,omitzero"`
	DropsOptedOut                         bool                               `json:"drops_opted_out,omitzero"`
	NonSpamRetrainingOptIn                bool                               `json:"non_spam_retraining_opt_in,omitzero"`
	FamilyCenterEnabled                   bool                               `json:"family_center_enabled,omitzero"`
	FamilyCenterEnabledV2                 bool                               `json:"family_center_enabled_v2,omitzero"`
	HideLegacyUsername                    bool                               `json:"hide_legacy_username,omitzero"`
	InappropriateConversationWarnings     bool                               `json:"inappropriate_conversation_warnings,omitzero"`
	RecentGamesEnabled                    bool                               `json:"recent_games_enabled,omitzero"`
	GuildsLeaderboardOptOutDefault        json.RawMessage                    `json:"guilds_leaderboard_opt_out_default"`
	AllowGameFriendDMsInDiscord           bool                               `json:"allow_game_friend_dms_in_discord,omitzero"`
	DefaultGuildsRestrictedV2             bool                               `json:"default_guilds_restricted_v2,omitzero"`
	SlayerSdkReceiveDMsInGame             json.RawMessage                    `json:"slayer_sdk_receive_dms_in_game"`
	DefaultGuildsActivityRestrictedV2     json.RawMessage                    `json:"default_guilds_activity_restricted_v2"`
	Quests3pDataOptedOut                  bool                               `json:"quests_3p_data_opted_out,omitzero"`
	ShowLocalTime                         bool                               `json:"show_local_time,omitzero"`
	ProfileVisibility                     json.RawMessage                    `json:"profile_visibility"`
	HideFriendRequestNotes                bool                               `json:"hide_friend_request_notes,omitzero"`
}

type ProductOption struct {
	Name         string   `json:"name"`
	OptionValues []string `json:"option_values"`
}

type ProductSKU struct {
	ID               snowflake.ID             `json:"id"`
	Type             SKUType                  `json:"type"`
	ProductLine      SKUProductLine           `json:"product_line"`
	ApplicationID    snowflake.ID             `json:"application_id"`
	Name             string                   `json:"name"`
	ThumbnailAssetID *snowflake.ID            `json:"thumbnail_asset_id"`
	Slug             string                   `json:"slug"`
	Premium          bool                     `json:"premium"`
	SelectedOptions  []ProductSKUOption       `json:"selected_options"`
	ProductID        snowflake.ID             `json:"product_id"`
	Position         int64                    `json:"position"`
	TenantMetadata   ProductSKUTenantMetadata `json:"tenant_metadata"`
}

type ProductSKUOption struct {
	OptionName  string `json:"option_name"`
	OptionValue string `json:"option_value"`
}

type ProductSKUPlanFeature struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ProductSKUTenantMetadata struct {
	BoostPrice    int64                    `json:"boost_price"`
	PurchaseLimit int64                    `json:"purchase_limit"`
	CategoryType  GuildPowerupCategoryType `json:"category_type"`
	PlanFeatures  []ProductSKUPlanFeature  `json:"plan_features"`
}

type ProductTenantMetadata struct {
	GuildMonetization GuildMonetizationProductMetadata `json:"guild_monetization,omitzero"`
}

type ProfileApplication struct {
	ID                           snowflake.ID                                       `json:"id"`
	Flags                        ApplicationFlags                                   `json:"flags"`
	Verified                     bool                                               `json:"verified"`
	StorefrontAvailable          bool                                               `json:"storefront_available"`
	PrimarySKUID                 snowflake.ID                                       `json:"primary_sku_id,omitzero"`
	InstallParams                ApplicationInstallParams                           `json:"install_params,omitzero"`
	IntegrationTypesConfig       map[int64]*ApplicationIntegrationTypeConfiguration `json:"integration_types_config,omitempty"`
	PopularApplicationCommandIDs []snowflake.ID                                     `json:"popular_application_command_ids,omitempty"`
	CustomInstallURL             string                                             `json:"custom_install_url,omitzero"`
}

type ProfileBadge struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	Link        string `json:"link,omitzero"`
}

type ProfileEffect struct {
	ID        snowflake.ID `json:"id"`
	ExpiresAt *int64       `json:"expires_at"`
}

type ProfileEffectAnimation struct {
	Src               string                `json:"src"`
	Loop              bool                  `json:"loop"`
	Height            int64                 `json:"height"`
	Width             int64                 `json:"width"`
	Duration          int64                 `json:"duration"`
	Start             int64                 `json:"start"`
	Loopdelay         int64                 `json:"loopDelay"`
	Position          ProfileEffectPosition `json:"position"`
	Zindex            int64                 `json:"zIndex"`
	Randomizedsources []ProfileEffectSource `json:"randomizedSources"`
}

type ProfileEffectCollectibleItem struct {
	Type                CollectibleProductType     `json:"type"`
	SKUID               snowflake.ID               `json:"sku_id"`
	Title               string                     `json:"title"`
	Description         string                     `json:"description"`
	Accessibilitylabel  string                     `json:"accessibilityLabel"`
	Animationtype       ProfileEffectAnimationType `json:"animationType"`
	Thumbnailpreviewsrc string                     `json:"thumbnailPreviewSrc"`
	Reducedmotionsrc    string                     `json:"reducedMotionSrc"`
	Staticframesrc      string                     `json:"staticFrameSrc,omitzero"`
	Effects             []ProfileEffectAnimation   `json:"effects"`
}

type ProfileEffectPosition struct {
	X int64 `json:"x"`
	Y int64 `json:"y"`
}

type ProfileEffectSource struct {
	Src string `json:"src"`
}

type ProfileFrameCollectibleItem struct {
	Type               CollectibleProductType `json:"type"`
	SKUID              snowflake.ID           `json:"sku_id"`
	Label              string                 `json:"label"`
	Layers             []ProfileFrameLayer    `json:"layers"`
	InnerWidth         int64                  `json:"inner_width"`
	OverflowTop        int64                  `json:"overflow_top"`
	OverflowBottom     int64                  `json:"overflow_bottom"`
	OverflowHorizontal int64                  `json:"overflow_horizontal"`
}

type ProfileFrameLayer struct {
	ID         snowflake.ID `json:"id"`
	Type       int64        `json:"type"`
	Order      int64        `json:"order"`
	Anchor     int64        `json:"anchor"`
	Responsive bool         `json:"responsive,omitzero"`
}

type ProfileMetadata struct {
	GuildID                     snowflake.ID   `json:"guild_id,omitzero"`
	Pronouns                    string         `json:"pronouns"`
	Bio                         string         `json:"bio,omitzero"`
	Banner                      *string        `json:"banner,omitempty"`
	AccentColor                 *int64         `json:"accent_color,omitempty"`
	ThemeColors                 []int64        `json:"theme_colors,omitempty"`
	PopoutAnimationParticleType *snowflake.ID  `json:"popout_animation_particle_type,omitempty"`
	Emoji                       *Emoji         `json:"emoji,omitempty"`
	ProfileEffect               *ProfileEffect `json:"profile_effect,omitempty"`
}

type Promotion struct {
	ID                          snowflake.ID                  `json:"id"`
	TrialID                     snowflake.ID                  `json:"trial_id,omitzero"`
	StartDate                   time.Time                     `json:"start_date"`
	EndDate                     time.Time                     `json:"end_date"`
	OutboundRedemptionEndDate   time.Time                     `json:"outbound_redemption_end_date,omitzero"`
	InboundHeaderText           string                        `json:"inbound_header_text,omitzero"`
	InboundBodyText             string                        `json:"inbound_body_text,omitzero"`
	InboundHelpCenterLink       string                        `json:"inbound_help_center_link,omitzero"`
	OutboundTitle               string                        `json:"outbound_title,omitzero"`
	OutboundRedemptionModalBody string                        `json:"outbound_redemption_modal_body,omitzero"`
	OutboundTermsAndConditions  string                        `json:"outbound_terms_and_conditions,omitzero"`
	OutboundRedemptionPageLink  string                        `json:"outbound_redemption_page_link,omitzero"`
	OutboundRedemptionURLFormat string                        `json:"outbound_redemption_url_format,omitzero"`
	Flags                       PromotionFlags                `json:"flags,omitzero"`
	InboundRestrictedCountries  []string                      `json:"inbound_restricted_countries,omitempty"`
	OutboundRestrictedCountries []string                      `json:"outbound_restricted_countries,omitempty"`
	PromotionType               PromotionType                 `json:"promotion_type"`
	PartnerID                   PromotionPartnerType          `json:"partner_id,omitzero"`
	MarketingComponents         []PromotionMarketingComponent `json:"marketing_components,omitempty"`
}

type PromotionMarketingComponent struct {
	ComponentType PromotionMarketingComponentType `json:"component_type"`
	ID            int64                           `json:"id"`
	PromotionID   snowflake.ID                    `json:"promotion_id"`
	Properties    string                          `json:"properties"`
}

type PromotionalEmailMetadata struct {
	Required   bool `json:"required"`
	PreChecked bool `json:"pre_checked"`
}

type ProtocolData struct {
	Address string                  `json:"address"`
	Port    int64                   `json:"port"`
	Mode    TransportEncryptionMode `json:"mode"`
}

type ProxyApplicationAssetsParams struct {
	URLs []string `json:"urls"`
}

type PruneGuildParams struct {
	Days              int64          `json:"days,omitzero"`
	ComputePruneCount bool           `json:"compute_prune_count,omitzero"`
	IncludeRoles      []snowflake.ID `json:"include_roles,omitempty"`
}

type PruneGuildResponse struct {
	Pruned *int64 `json:"pruned"`
}

type PushNotification struct {
	Type                PushNotificationType `json:"type"`
	ReceivingUserID     snowflake.ID         `json:"receiving_user_id,omitzero"`
	NotifTypeID         NotificationType     `json:"notif_type_id,omitzero"`
	NotifInstanceID     string               `json:"notif_instance_id,omitzero"`
	TrackingType        TrackingType         `json:"tracking_type,omitzero"`
	NotificationChannel string               `json:"notification_channel,omitzero"`
	Title               string               `json:"title,omitzero"`
	Subtitle            string               `json:"subtitle,omitzero"`
	ExpandSubtitle      bool                 `json:"expand_subtitle,omitzero"`
	IconURL             string               `json:"icon_url,omitzero"`
	Silent              bool                 `json:"silent,omitzero"`
	SentAtMs            int64                `json:"sent_at_ms,omitzero"`
	ChannelIDs          []snowflake.ID       `json:"channel_ids,omitempty"`
	MentionType         string               `json:"mention_type,omitzero"`
	JoinID              string               `json:"join_id,omitzero"`
}

type PushToTalkArguments struct {
	Active bool `json:"active"`
}

type QoSHeartbeat struct {
	Seq *int64     `json:"seq"`
	QoS QoSPayload `json:"qos"`
}

type QoSPayload struct {
	Ver     int64    `json:"ver"`
	Active  bool     `json:"active"`
	Reasons []string `json:"reasons"`
}

type Query struct {
	OrQuery  json.RawMessage `json:"or_query,omitempty"`
	AndQuery json.RawMessage `json:"and_query,omitempty"`
	Range    RangeQuery      `json:"range,omitzero"`
}

type QueryGuildMembersQuery struct {
	Query string `json:"query"`
	Limit int64  `json:"limit,omitzero"`
}

type QueryStudentHubsParams struct {
	Email               string       `json:"email"`
	UseVerificationCode bool         `json:"use_verification_code,omitzero"`
	AllowMultipleGuilds bool         `json:"allow_multiple_guilds,omitzero"`
	GuildID             snowflake.ID `json:"guild_id,omitzero"`
}

type QueryStudentHubsResponse struct {
	HasMatchingGuild bool              `json:"has_matching_guild"`
	GuildsInfo       []StudentHubGuild `json:"guilds_info,omitempty"`
}

type Quest struct {
	ID                    snowflake.ID     `json:"id"`
	Config                QuestConfig      `json:"config"`
	UserStatus            *QuestUserStatus `json:"user_status"`
	TargetedContent       []int64          `json:"targeted_content"`
	Preview               bool             `json:"preview"`
	TrafficMetadataSealed string           `json:"traffic_metadata_sealed,omitzero"`
}

type QuestAdContext struct {
	IsCampaignIasEnabled bool `json:"is_campaign_ias_enabled"`
}

type QuestAdCreative struct {
	CreativeType    AdCreativeType  `json:"creative_type"`
	CreativeContent json.RawMessage `json:"creative_content"`
	StartsAt        time.Time       `json:"starts_at"`
	EndsAt          time.Time       `json:"ends_at"`
}

type QuestAdDecision struct {
	Creative              *QuestAdCreative   `json:"creative"`
	AdIdentifiers         QuestAdIdentifiers `json:"ad_identifiers,omitzero"`
	AdContext             QuestAdContext     `json:"ad_context,omitzero"`
	ResponseTTLSeconds    int64              `json:"response_ttl_seconds,omitzero"`
	MetadataSealed        *string            `json:"metadata_sealed,omitempty"`
	TrafficMetadataSealed *string            `json:"traffic_metadata_sealed,omitempty"`
}

type QuestAdIdentifiers struct {
	CampaignID   snowflake.ID   `json:"campaign_id,omitzero"`
	AdSetID      snowflake.ID   `json:"ad_set_id,omitzero"`
	AdID         snowflake.ID   `json:"ad_id,omitzero"`
	CreativeID   snowflake.ID   `json:"creative_id,omitzero"`
	CreativeType AdCreativeType `json:"creative_type,omitzero"`
	IsTargeted   bool           `json:"is_targeted,omitzero"`
	AdContentID  snowflake.ID   `json:"ad_content_id,omitzero"`
}

type QuestApplication struct {
	ID   snowflake.ID `json:"id"`
	Name string       `json:"name"`
	Link string       `json:"link"`
}

type QuestAssets struct {
	Hero                 string  `json:"hero"`
	HeroVideo            *string `json:"hero_video"`
	QuestBarHero         string  `json:"quest_bar_hero"`
	QuestBarHeroBlurhash *string `json:"quest_bar_hero_blurhash,omitempty"`
	QuestBarHeroVideo    *string `json:"quest_bar_hero_video"`
	GameTile             string  `json:"game_tile"`
	GameTileLight        string  `json:"game_tile_light,omitzero"`
	GameTileDark         string  `json:"game_tile_dark,omitzero"`
	Logotype             string  `json:"logotype"`
	LogotypeLight        string  `json:"logotype_light,omitzero"`
	LogotypeDark         string  `json:"logotype_dark,omitzero"`
}

type QuestBountyCreative struct {
	ID                 snowflake.ID     `json:"id"`
	AdvertiserName     string           `json:"advertiser_name"`
	ProductName        string           `json:"product_name,omitzero"`
	ProductIcon        string           `json:"product_icon,omitzero"`
	VideoPreview       string           `json:"video_preview,omitzero"`
	ImagePreview       string           `json:"image_preview,omitzero"`
	VideoHls           string           `json:"video_hls"`
	Cta                QuestCreativeCta `json:"cta"`
	RewardTimerSeconds int64            `json:"reward_timer_seconds,omitzero"`
}

type QuestConfig struct {
	ID                snowflake.ID           `json:"id"`
	ConfigVersion     int64                  `json:"config_version"`
	StartsAt          time.Time              `json:"starts_at"`
	ExpiresAt         time.Time              `json:"expires_at"`
	Features          []int64                `json:"features"`
	Application       QuestApplication       `json:"application"`
	Assets            QuestAssets            `json:"assets"`
	Colors            QuestGradient          `json:"colors"`
	Messages          QuestMessages          `json:"messages"`
	TaskConfigV2      QuestTaskConfig        `json:"task_config_v2"`
	RewardsConfig     QuestRewardsConfig     `json:"rewards_config"`
	CosponsorMetadata QuestCosponsorMetadata `json:"cosponsor_metadata,omitzero"`
	SharePolicy       QuestSharePolicy       `json:"share_policy"`
	CtaConfig         QuestCtaConfig         `json:"cta_config"`
}

type QuestCosponsorMetadata struct {
	Name                   string `json:"name"`
	Logotype               string `json:"logotype"`
	LogotypeLight          string `json:"logotype_light,omitzero"`
	LogotypeDark           string `json:"logotype_dark,omitzero"`
	RedemptionInstructions string `json:"redemption_instructions"`
}

type QuestCreativeCta struct {
	URL         string                `json:"url"`
	ButtonLabel string                `json:"button_label"`
	Android     AndroidQuestCtaConfig `json:"android,omitzero"`
	Ios         IosQuestCtaConfig     `json:"ios,omitzero"`
}

type QuestCtaConfig struct {
	Link        string                `json:"link"`
	ButtonLabel string                `json:"button_label"`
	Subtitle    string                `json:"subtitle,omitzero"`
	Android     AndroidQuestCtaConfig `json:"android,omitzero"`
	Ios         IosQuestCtaConfig     `json:"ios,omitzero"`
}

type QuestEnrollmentStatusUpdate struct {
	QuestID    snowflake.ID `json:"quest_id"`
	IsEnrolled bool         `json:"is_enrolled"`
	EnrolledAt *time.Time   `json:"enrolled_at"`
}

type QuestErrorHint struct {
	Type                 string       `json:"type"`
	Message              string       `json:"message"`
	ConnectedAccountID   snowflake.ID `json:"connected_account_id"`
	ConnectedAccountType string       `json:"connected_account_type"`
}

type QuestGradient struct {
	Primary   string `json:"primary"`
	Secondary string `json:"secondary"`
}

type QuestHomeEntrypoint struct {
	LinearGradient  QuestHomeEntrypointGradient `json:"linear_gradient,omitzero"`
	RadialGradient  QuestHomeEntrypointGradient `json:"radial_gradient,omitzero"`
	GradientPreset  int64                       `json:"gradient_preset,omitzero"`
	Image           string                      `json:"image,omitzero"`
	TooltipImage    string                      `json:"tooltip_image,omitzero"`
	TooltipTitle    string                      `json:"tooltip_title,omitzero"`
	TooltipSubtitle string                      `json:"tooltip_subtitle,omitzero"`
}

type QuestHomeEntrypointGradient struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type QuestHomeHeroCreative struct {
	ID                  snowflake.ID        `json:"id"`
	LabelTitle          string              `json:"label_title"`
	LabelSubtitle       string              `json:"label_subtitle,omitzero"`
	HeroImage           string              `json:"hero_image"`
	HeroVideo           string              `json:"hero_video,omitzero"`
	SponsorImage        string              `json:"sponsor_image,omitzero"`
	Cta                 QuestCreativeCta    `json:"cta"`
	QuestIDs            []snowflake.ID      `json:"quest_ids,omitempty"`
	QuestHomeEntrypoint QuestHomeEntrypoint `json:"quest_home_entrypoint,omitzero"`
	ShelfImage          string              `json:"shelf_image,omitzero"`
	ShelfVideo          string              `json:"shelf_video,omitzero"`
}

type QuestMessages struct {
	QuestName     string `json:"quest_name"`
	GameTitle     string `json:"game_title"`
	GamePublisher string `json:"game_publisher"`
}

type QuestReward struct {
	Type               QuestRewardType           `json:"type"`
	SKUID              snowflake.ID              `json:"sku_id"`
	Asset              *string                   `json:"asset,omitempty"`
	AssetVideo         *string                   `json:"asset_video,omitempty"`
	Messages           QuestRewardMessages       `json:"messages"`
	ApproximateCount   *int64                    `json:"approximate_count,omitempty"`
	RedemptionLink     *string                   `json:"redemption_link,omitempty"`
	ExpiresAt          *time.Time                `json:"expires_at,omitempty"`
	ExpiresAtPremium   *time.Time                `json:"expires_at_premium,omitempty"`
	ExpirationMode     QuestRewardExpirationMode `json:"expiration_mode,omitzero"`
	OrbQuantity        int64                     `json:"orb_quantity,omitzero"`
	PremiumOrbQuantity int64                     `json:"premium_orb_quantity,omitzero"`
	Quantity           int64                     `json:"quantity,omitzero"`
}

type QuestRewardCode struct {
	QuestID   snowflake.ID                 `json:"quest_id"`
	Code      string                       `json:"code"`
	Platform  QuestPlatformType            `json:"platform"`
	UserID    snowflake.ID                 `json:"user_id"`
	ClaimedAt time.Time                    `json:"claimed_at"`
	Tier      *QuestRewardAssignmentMethod `json:"tier"`
}

type QuestRewardMessages struct {
	Name                             string           `json:"name"`
	NameWithArticle                  string           `json:"name_with_article"`
	RedemptionInstructionsByPlatform map[int64]string `json:"redemption_instructions_by_platform,omitempty"`
}

type QuestRewardsConfig struct {
	AssignmentMethod QuestRewardAssignmentMethod `json:"assignment_method"`
	Rewards          []QuestReward               `json:"rewards"`
	RewardsExpireAt  *time.Time                  `json:"rewards_expire_at"`
	Platforms        []int64                     `json:"platforms"`
}

type QuestRewardsMetadata struct {
	Tag        QuestRewardType `json:"tag"`
	RewardCode QuestRewardCode `json:"reward_code,omitzero"`
}

type QuestStartTimerArguments struct {
	QuestID snowflake.ID `json:"quest_id"`
}

type QuestStartTimerResponse struct {
	Success bool `json:"success"`
}

type QuestTask struct {
	Type                    QuestTaskType          `json:"type"`
	Target                  int64                  `json:"target"`
	Applications            []QuestTaskApplication `json:"applications,omitempty"`
	ExternalIDs             []string               `json:"external_ids,omitempty"`
	Assets                  QuestVideoAssets       `json:"assets,omitzero"`
	Messages                QuestTaskMessages      `json:"messages,omitzero"`
	EventName               string                 `json:"event_name,omitzero"`
	AccountLinkInstructions string                 `json:"account_link_instructions,omitzero"`
}

type QuestTaskApplication struct {
	ID snowflake.ID `json:"id"`
}

type QuestTaskConfig struct {
	JoinOperator string               `json:"join_operator"`
	Tasks        map[string]QuestTask `json:"tasks"`
}

type QuestTaskHeartbeat struct {
	LastBeatAt time.Time  `json:"last_beat_at"`
	ExpiresAt  *time.Time `json:"expires_at"`
}

type QuestTaskMessages struct {
	VideoTitle      string `json:"video_title,omitzero"`
	TaskTitle       string `json:"task_title,omitzero"`
	TaskDescription string `json:"task_description,omitzero"`
}

type QuestTaskProgress struct {
	EventName   QuestTaskType       `json:"event_name"`
	Value       int64               `json:"value"`
	UpdatedAt   time.Time           `json:"updated_at"`
	CompletedAt *time.Time          `json:"completed_at"`
	Heartbeat   *QuestTaskHeartbeat `json:"heartbeat,omitempty"`
}

type QuestUserStatus struct {
	UserID                snowflake.ID                 `json:"user_id"`
	QuestID               snowflake.ID                 `json:"quest_id,omitzero"`
	EnrolledAt            *time.Time                   `json:"enrolled_at"`
	CompletedAt           *time.Time                   `json:"completed_at"`
	ClaimedAt             *time.Time                   `json:"claimed_at"`
	ClaimedTier           *QuestRewardAssignmentMethod `json:"claimed_tier,omitempty"`
	OrbQuantityClaimed    *int64                       `json:"orb_quantity_claimed,omitempty"`
	LastStreamHeartbeatAt *time.Time                   `json:"last_stream_heartbeat_at,omitempty"`
	StreamProgressSeconds int64                        `json:"stream_progress_seconds,omitzero"`
	DismissedQuestContent DismissibleQuestContentFlags `json:"dismissed_quest_content,omitzero"`
	Progress              map[string]QuestTaskProgress `json:"progress"`
}

type QuestVideoAsset struct {
	URL        string `json:"url"`
	Width      int64  `json:"width,omitzero"`
	Height     int64  `json:"height,omitzero"`
	Thumbnail  string `json:"thumbnail,omitzero"`
	Caption    string `json:"caption,omitzero"`
	Transcript string `json:"transcript,omitzero"`
}

type QuestVideoAssets struct {
	Video       QuestVideoAsset `json:"video"`
	VideoLowRes QuestVideoAsset `json:"video_low_res,omitzero"`
	VideoHls    QuestVideoAsset `json:"video_hls,omitzero"`
}

type QuestsUserCompletionUpdate struct {
	QuestEnrollmentBlockedUntil *time.Time `json:"quest_enrollment_blocked_until"`
}

type QuestsUserStatusUpdate struct {
	UserStatus QuestUserStatus `json:"user_status"`
}

type RPCActivity struct {
	Name               string              `json:"name"`
	State              *string             `json:"state,omitempty"`
	StateURL           *string             `json:"state_url,omitempty"`
	Details            *string             `json:"details,omitempty"`
	DetailsURL         *string             `json:"details_url,omitempty"`
	Timestamps         ActivityTimestamps  `json:"timestamps,omitzero"`
	Assets             ActivityAssets      `json:"assets,omitzero"`
	Party              RPCActivityParty    `json:"party,omitzero"`
	Secrets            ActivitySecrets     `json:"secrets,omitzero"`
	Buttons            []RPCActivityButton `json:"buttons,omitempty"`
	Instance           bool                `json:"instance,omitzero"`
	SupportedPlatforms []string            `json:"supported_platforms,omitempty"`
	Type               ActivityType        `json:"type"`
	StatusDisplayType  *StatusDisplayType  `json:"status_display_type,omitempty"`
}

type RPCActivityButton struct {
	Label string `json:"label"`
	URL   string `json:"url"`
}

type RPCActivityParticipant struct {
	Nickname string `json:"nickname,omitzero"`
}

type RPCActivityParty struct {
	ID      string               `json:"id,omitzero"`
	Size    []int64              `json:"size,omitempty"`
	Privacy ActivityPartyPrivacy `json:"privacy,omitzero"`
}

type RPCChannel struct {
	ID          snowflake.ID    `json:"id"`
	Name        string          `json:"name"`
	Type        ChannelType     `json:"type"`
	Topic       string          `json:"topic"`
	Bitrate     int64           `json:"bitrate"`
	UserLimit   int64           `json:"user_limit"`
	GuildID     *snowflake.ID   `json:"guild_id"`
	Position    int64           `json:"position"`
	Messages    []RPCMessage    `json:"messages"`
	VoiceStates []RPCVoiceState `json:"voice_states"`
}

type RPCGuild struct {
	ID            snowflake.ID `json:"id"`
	Name          string       `json:"name"`
	IconURL       *string      `json:"icon_url"`
	VanityURLCode *string      `json:"vanity_url_code"`
}

type RPCGuildMember struct {
	UserID               snowflake.ID          `json:"user_id"`
	Nick                 *string               `json:"nick"`
	GuildID              snowflake.ID          `json:"guild_id"`
	Avatar               *string               `json:"avatar"`
	AvatarDecorationData *AvatarDecorationData `json:"avatar_decoration_data"`
	Banner               *string               `json:"banner,omitempty"`
	Bio                  *string               `json:"bio,omitempty"`
	Pronouns             *string               `json:"pronouns,omitempty"`
	ColorString          *string               `json:"color_string,omitempty"`
}

type RPCGuildTemplate struct {
	Code                  string                `json:"code"`
	State                 RPCGuildTemplateState `json:"state"`
	Name                  string                `json:"name"`
	Description           string                `json:"description"`
	Creatorid             snowflake.ID          `json:"creatorId"`
	Creator               PartialUser           `json:"creator"`
	Createdat             time.Time             `json:"createdAt"`
	Updatedat             time.Time             `json:"updatedAt"`
	Sourceguildid         snowflake.ID          `json:"sourceGuildId"`
	Serializedsourceguild PartialGuild          `json:"serializedSourceGuild"`
	Usagecount            int64                 `json:"usageCount"`
	Isdirty               *bool                 `json:"isDirty"`
}

type RPCHandshake struct {
	V          int64  `json:"v"`
	ClientID   string `json:"client_id"`
	FrameID    string `json:"frame_id"`
	SdkVersion string `json:"sdk_version,omitzero"`
}

type RPCMessage struct {
	ID              snowflake.ID      `json:"id"`
	Blocked         bool              `json:"blocked,omitzero"`
	Bot             bool              `json:"bot,omitzero"`
	Content         string            `json:"content"`
	ContentParsed   []json.RawMessage `json:"content_parsed,omitempty"`
	Nick            string            `json:"nick,omitzero"`
	AuthorColor     string            `json:"author_color,omitzero"`
	EditedTimestamp *time.Time        `json:"edited_timestamp"`
	Timestamp       time.Time         `json:"timestamp"`
	TTS             bool              `json:"tts"`
	Mentions        []RPCUser         `json:"mentions"`
	MentionEveryone bool              `json:"mention_everyone"`
	MentionRoles    []snowflake.ID    `json:"mention_roles"`
	Embeds          []Embed           `json:"embeds"`
	Attachments     []Attachment      `json:"attachments"`
	Author          *RPCUser          `json:"author,omitempty"`
	Pinned          bool              `json:"pinned"`
	Type            MessageType       `json:"type"`
}

type RPCMessageCreate struct {
	ChannelID snowflake.ID `json:"channel_id"`
	Message   RPCMessage   `json:"message"`
}

type RPCMessageDelete struct {
	ChannelID snowflake.ID      `json:"channel_id"`
	Message   PartialRPCMessage `json:"message"`
}

type RPCRPCVoiceInputMode struct {
	Type     RPCVoiceSettingsModeType `json:"type"`
	Shortcut string                   `json:"shortcut"`
}

type RPCReady struct {
	V      int64                   `json:"v"`
	Config ClientEnvironmentConfig `json:"config"`
	User   RPCUser                 `json:"user,omitzero"`
}

type RPCRelationship struct {
	Type     RelationshipType        `json:"type"`
	User     RPCUser                 `json:"user"`
	Presence RPCRelationshipPresence `json:"presence"`
}

type RPCRelationshipPresence struct {
	Status   StatusType `json:"status"`
	Activity *Activity  `json:"activity"`
}

type RPCRemoteVoiceState struct {
	Mute     bool `json:"mute"`
	Deaf     bool `json:"deaf"`
	SelfMute bool `json:"self_mute"`
	SelfDeaf bool `json:"self_deaf"`
	Suppress bool `json:"suppress"`
}

type RPCSendActivityJoinInviteArguments struct {
	UserID snowflake.ID `json:"user_id"`
	Pid    int64        `json:"pid"`
}

type RPCSetUserVoiceSettingsArguments struct {
	UserID snowflake.ID `json:"user_id"`
	Pan    Pan          `json:"pan,omitzero"`
	Volume int64        `json:"volume,omitzero"`
	Mute   bool         `json:"mute,omitzero"`
}

type RPCUser struct {
	ID                   snowflake.ID          `json:"id"`
	Username             string                `json:"username"`
	Discriminator        string                `json:"discriminator"`
	GlobalName           *string               `json:"global_name"`
	Avatar               *string               `json:"avatar"`
	AvatarDecorationData *AvatarDecorationData `json:"avatar_decoration_data"`
	Bot                  bool                  `json:"bot"`
	Flags                UserFlags             `json:"flags"`
	PremiumType          PremiumType           `json:"premium_type"`
}

type RPCVoiceInputMode struct {
	Type     RPCVoiceSettingsModeType `json:"type"`
	Shortcut string                   `json:"shortcut"`
}

type RPCVoiceIoSettings struct {
	AvailableDevices []AvailableDevice `json:"available_devices"`
	DeviceID         string            `json:"device_id"`
	Volume           float64           `json:"volume"`
}

type RPCVoiceSettings struct {
	Input                RPCVoiceIoSettings   `json:"input"`
	Output               RPCVoiceIoSettings   `json:"output"`
	Mode                 RPCVoiceSettingsMode `json:"mode"`
	AutomaticGainControl bool                 `json:"automatic_gain_control"`
	EchoCancellation     bool                 `json:"echo_cancellation"`
	NoiseSuppression     bool                 `json:"noise_suppression"`
	QoS                  bool                 `json:"qos"`
	SilenceWarning       bool                 `json:"silence_warning"`
	Deaf                 bool                 `json:"deaf"`
	Mute                 bool                 `json:"mute"`
}

type RPCVoiceSettingsMode struct {
	Type          RPCVoiceSettingsModeType `json:"type"`
	AutoThreshold bool                     `json:"auto_threshold"`
	Threshold     int64                    `json:"threshold"`
	Shortcut      []ShortcutKeyCombo       `json:"shortcut"`
	Delay         int64                    `json:"delay"`
}

type RPCVoiceState struct {
	Nick       string              `json:"nick"`
	Mute       bool                `json:"mute"`
	Volume     float64             `json:"volume"`
	Pan        Pan                 `json:"pan"`
	VoiceState RPCRemoteVoiceState `json:"voice_state"`
	User       RPCUser             `json:"user"`
}

type RadioGroup struct {
	Type     ComponentType      `json:"type"`
	ID       int64              `json:"id,omitzero"`
	CustomID string             `json:"custom_id"`
	Options  []RadioGroupOption `json:"options"`
	Required bool               `json:"required,omitzero"`
}

type RadioGroupOption struct {
	Value       string `json:"value"`
	Label       string `json:"label"`
	Description string `json:"description,omitzero"`
	Default     bool   `json:"default,omitzero"`
}

type RangeQuery struct {
	Gte json.RawMessage `json:"gte,omitempty"`
	Lte json.RawMessage `json:"lte,omitempty"`
}

type RateLimitMetadata struct {
	GuildID snowflake.ID `json:"guild_id"`
	Nonce   string       `json:"nonce,omitzero"`
}

type RateLimitResponse struct {
	Message    string  `json:"message"`
	RetryAfter float64 `json:"retry_after"`
	Global     bool    `json:"global"`
	Code       int64   `json:"code,omitzero"`
}

type RateLimited struct {
	Opcode     int64             `json:"opcode"`
	RetryAfter float64           `json:"retry_after"`
	Meta       RateLimitMetadata `json:"meta"`
}

type Reaction struct {
	Count        int64                `json:"count"`
	CountDetails ReactionCountDetails `json:"count_details"`
	Me           bool                 `json:"me"`
	MeBurst      bool                 `json:"me_burst"`
	Emoji        Emoji                `json:"emoji"`
	BurstColors  []string             `json:"burst_colors"`
}

type ReactionCountDetails struct {
	Normal int64 `json:"normal"`
	Burst  int64 `json:"burst"`
}

type ReactionNotificationSent struct {
	Message       Message      `json:"message"`
	ReactorUserID snowflake.ID `json:"reactor_user_id"`
	Emoji         Emoji        `json:"emoji"`
	Title         string       `json:"title"`
	Body          string       `json:"body"`
	Icon          string       `json:"icon"`
	Route         string       `json:"route"`
	TrackingType  string       `json:"tracking_type"`
}

type ReadState struct {
	ID               snowflake.ID   `json:"id"`
	ReadStateType    ReadStateType  `json:"read_state_type,omitzero"`
	LastMessageID    snowflake.ID   `json:"last_message_id,omitzero"`
	LastAckedID      snowflake.ID   `json:"last_acked_id,omitzero"`
	MentionCount     int64          `json:"mention_count,omitzero"`
	BadgeCount       int64          `json:"badge_count,omitzero"`
	LastPinTimestamp time.Time      `json:"last_pin_timestamp,omitzero"`
	Flags            ReadStateFlags `json:"flags,omitzero"`
	LastViewed       *int64         `json:"last_viewed,omitempty"`
}

type ReadStateUpdate struct {
	ReadStateType ReadStateType `json:"read_state_type,omitzero"`
	ChannelID     snowflake.ID  `json:"channel_id"`
	MessageID     snowflake.ID  `json:"message_id"`
}

type Ready struct {
	Trace                            []string                                  `json:"_trace"`
	V                                int64                                     `json:"v"`
	User                             User                                      `json:"user"`
	UserSettings                     UserSettings                              `json:"user_settings"`
	UserSettingsProto                string                                    `json:"user_settings_proto"`
	NotificationSettings             NotificationSettings                      `json:"notification_settings"`
	UserGuildSettings                json.RawMessage                           `json:"user_guild_settings"`
	ReadState                        json.RawMessage                           `json:"read_state"`
	Guilds                           []GatewayGuild                            `json:"guilds"`
	GuildJoinRequests                []json.RawMessage                         `json:"guild_join_requests"`
	Relationships                    []Relationship                            `json:"relationships"`
	GameRelationships                []GameRelationship                        `json:"game_relationships"`
	FriendSuggestionCount            int64                                     `json:"friend_suggestion_count,omitzero"`
	PrivateChannels                  []Channel                                 `json:"private_channels"`
	ConnectedAccounts                []Connection                              `json:"connected_accounts"`
	Notes                            map[snowflake.ID]string                   `json:"notes"`
	Presences                        []Presence                                `json:"presences"`
	MergedPresences                  MergedPresences                           `json:"merged_presences"`
	MergedMembers                    [][]GuildMember                           `json:"merged_members"`
	Users                            []PartialUser                             `json:"users"`
	LinkedUsers                      []LinkedUser                              `json:"linked_users"`
	Application                      GatewayApplication                        `json:"application,omitzero"`
	Scopes                           []string                                  `json:"scopes,omitempty"`
	SessionID                        string                                    `json:"session_id"`
	SessionType                      SessionType                               `json:"session_type"`
	Sessions                         []Session                                 `json:"sessions"`
	StaticClientSessionID            string                                    `json:"static_client_session_id"`
	AuthSessionIDHash                string                                    `json:"auth_session_id_hash"`
	AuthToken                        string                                    `json:"auth_token,omitzero"`
	AnalyticsToken                   string                                    `json:"analytics_token"`
	Auth                             GatewayAuth                               `json:"auth"`
	RequiredAction                   RequiredActionType                        `json:"required_action,omitzero"`
	CountryCode                      string                                    `json:"country_code"`
	GeoOrderedRTCRegions             []string                                  `json:"geo_ordered_rtc_regions"`
	Consents                         json.RawMessage                           `json:"consents"`
	Tutorial                         *Tutorial                                 `json:"tutorial"`
	Shard                            []int64                                   `json:"shard,omitempty"`
	ResumeGatewayURL                 string                                    `json:"resume_gateway_url"`
	APICodeVersion                   int64                                     `json:"api_code_version"`
	Experiments                      []UserExperiment                          `json:"experiments"`
	GuildExperiments                 []GuildExperiment                         `json:"guild_experiments"`
	ApexExperiments                  ApexExperiments                           `json:"apex_experiments,omitzero"`
	ExplicitContentScanVersion       int64                                     `json:"explicit_content_scan_version"`
	PendingPayments                  []Payment                                 `json:"pending_payments,omitempty"`
	AvSfProtocolFloor                int64                                     `json:"av_sf_protocol_floor,omitzero"`
	FeatureFlags                     GatewayFeatureFlags                       `json:"feature_flags,omitzero"`
	Lobbies                          []Lobby                                   `json:"lobbies,omitempty"`
	UserApplicationProfiles          map[snowflake.ID][]UserApplicationProfile `json:"user_application_profiles,omitempty"`
	ConnectionRequestData            ConsoleConnectionRequestData              `json:"connection_request_data,omitzero"`
	AdPersonalizationTogglesDisabled bool                                      `json:"ad_personalization_toggles_disabled"`
	BroadcasterUserIDs               []snowflake.ID                            `json:"broadcaster_user_ids"`
	RegionalFeatureConfig            RegionalFeatureConfig                     `json:"regional_feature_config"`
}

type ReadySupplemental struct {
	Guilds              []SupplementalGuild `json:"guilds"`
	MergedMembers       [][]json.RawMessage `json:"merged_members"`
	MergedPresences     MergedPresences     `json:"merged_presences"`
	LazyPrivateChannels []Channel           `json:"lazy_private_channels"`
	Disclose            []string            `json:"disclose"`
	GameInvites         []GameInvite        `json:"game_invites"`
}

type RecentMentionDelete struct {
	MessageID snowflake.ID `json:"message_id"`
}

type RecurringDismissibleContentState struct {
	LastDismissedVersion  json.RawMessage `json:"last_dismissed_version"`
	LastDismissedAtMs     json.RawMessage `json:"last_dismissed_at_ms"`
	LastDismissedObjectID json.RawMessage `json:"last_dismissed_object_id"`
	NumTimesDismissed     json.RawMessage `json:"num_times_dismissed"`
}

type RedeemGiftCodeParams struct {
	PaymentSourceID        *string                 `json:"payment_source_id,omitempty"`
	ChannelID              *snowflake.ID           `json:"channel_id,omitempty"`
	GatewayCheckoutContext *GatewayCheckoutContext `json:"gateway_checkout_context,omitempty"`
}

type RedeemUserOfferParams struct {
	UserDiscountOfferID snowflake.ID `json:"user_discount_offer_id"`
}

type RedeemVirtualCurrencyParams struct {
	CheckoutSessionID string `json:"checkout_session_id"`
}

type RefreshAttachmentURLsParams struct {
	AttachmentURLs []string `json:"attachment_urls"`
}

type RefreshAttachmentURLsResponse struct {
	RefreshedURLs []RefreshedAttachment `json:"refreshed_urls"`
}

type RefreshedAttachment struct {
	Original  string `json:"original"`
	Refreshed string `json:"refreshed"`
}

type RegionalFeatureConfig struct {
	AgeGatedFeatures      float64 `json:"age_gated_features"`
	TeenByDefaultSettings float64 `json:"teen_by_default_settings"`
}

type RegisterAccountParams struct {
	Username              string  `json:"username,omitzero"`
	GlobalName            *string `json:"global_name,omitempty"`
	Email                 string  `json:"email,omitzero"`
	PhoneToken            string  `json:"phone_token,omitzero"`
	Password              string  `json:"password,omitzero"`
	DateOfBirth           string  `json:"date_of_birth,omitzero"`
	Fingerprint           string  `json:"fingerprint,omitzero"`
	Invite                *string `json:"invite,omitempty"`
	GuildTemplateCode     *string `json:"guild_template_code,omitempty"`
	GiftCodeSKUID         *string `json:"gift_code_sku_id,omitempty"`
	Consent               bool    `json:"consent,omitzero"`
	PromotionalEmailOptIn bool    `json:"promotional_email_opt_in,omitzero"`
}

type RegisterAccountResponse struct {
	Token                string `json:"token"`
	ShowVerificationForm bool   `json:"show_verification_form,omitzero"`
}

type RegisterAccountWithPhoneNumberParams struct {
	Phone string `json:"phone"`
}

type RegisterDeviceParams struct {
	Provider                        PushNotificationProvider `json:"provider"`
	Token                           string                   `json:"token"`
	VoipProvider                    PushNotificationProvider `json:"voip_provider,omitzero"`
	VoipToken                       string                   `json:"voip_token,omitzero"`
	BypassServerThrottlingSupported bool                     `json:"bypass_server_throttling_supported,omitzero"`
	BundleID                        string                   `json:"bundle_id,omitzero"`
}

type Relationship struct {
	ID                  string           `json:"id"`
	Type                RelationshipType `json:"type"`
	User                PartialUser      `json:"user"`
	Nickname            *string          `json:"nickname"`
	IsSpamRequest       bool             `json:"is_spam_request,omitzero"`
	StrangerRequest     bool             `json:"stranger_request,omitzero"`
	UserIgnored         bool             `json:"user_ignored"`
	OriginApplicationID *snowflake.ID    `json:"origin_application_id,omitempty"`
	Since               time.Time        `json:"since,omitzero"`
	HasPlayedGame       bool             `json:"has_played_game,omitzero"`
	Note                string           `json:"note,omitzero"`
}

type RelationshipAdd struct {
	RelType          RelationshipType `json:"rel_type"`
	UserID           snowflake.ID     `json:"user_id"`
	UserUsername     string           `json:"user_username"`
	UserGlobalName   *string          `json:"user_global_name,omitempty"`
	UserAvatar       *string          `json:"user_avatar,omitempty"`
	NotificationType TrackingType     `json:"notification_type,omitzero"`
	Since            time.Time        `json:"since,omitzero"`
}

type RelationshipAddExtraFields struct {
	ShouldNotify bool `json:"should_notify,omitzero"`
}

type RemoteCommand struct {
	TargetSessionID string          `json:"target_session_id"`
	Payload         json.RawMessage `json:"payload"`
}

type RemoveGuildPowerupQuery struct {
	EntitlementID snowflake.ID `json:"entitlement_id,omitzero"`
}

type RemoveLinkedUserParams struct {
	LinkedUserID snowflake.ID `json:"linked_user_id"`
}

type RemovePhoneNumberParams struct {
	Password          string            `json:"password"`
	ChangePhoneReason ChangePhoneReason `json:"change_phone_reason"`
}

type RemoveThreadMemberQuery struct {
	Location string `json:"location,omitzero"`
}

type ReportButton struct {
	Type   ReportButtonType `json:"type"`
	Target *int64           `json:"target"`
}

type ReportElement struct {
	Name              string            `json:"name"`
	Type              ReportElementType `json:"type"`
	Data              json.RawMessage   `json:"data"`
	ShouldSubmitData  bool              `json:"should_submit_data"`
	SkipIfUnlocalized bool              `json:"skip_if_unlocalized"`
	IsLocalized       bool              `json:"is_localized"`
}

type ReportExplicitContentFalsePositiveParams struct {
	ChannelID     snowflake.ID   `json:"channel_id"`
	MessageID     snowflake.ID   `json:"message_id"`
	AttachmentIDs []snowflake.ID `json:"attachment_ids"`
	EmbedIDs      []string       `json:"embed_ids"`
}

type ReportMenu struct {
	Name          ReportMenuType       `json:"name"`
	Version       string               `json:"version"`
	Variant       string               `json:"variant"`
	PostbackURL   string               `json:"postback_url"`
	Language      string               `json:"language,omitzero"`
	RootNodeID    int64                `json:"root_node_id"`
	SuccessNodeID int64                `json:"success_node_id"`
	FailNodeID    int64                `json:"fail_node_id"`
	Nodes         map[int64]ReportNode `json:"nodes"`
}

type ReportNode struct {
	ID                    int64             `json:"id"`
	ReportType            string            `json:"report_type,omitzero"`
	Key                   string            `json:"key"`
	Header                string            `json:"header"`
	Subheader             string            `json:"subheader,omitzero"`
	Info                  string            `json:"info,omitzero"`
	Children              []ReportNodeChild `json:"children"`
	Elements              []ReportElement   `json:"elements"`
	Button                ReportButton      `json:"button,omitzero"`
	IsMultiSelectRequired bool              `json:"is_multi_select_required"`
	IsAutoSubmit          bool              `json:"is_auto_submit"`
}

type ReportNodeChild struct {
	Name         string `json:"name"`
	TargetNodeID int64  `json:"target_node_id"`
}

type ReportOption struct {
	Value       string                `json:"value"`
	Label       string                `json:"label"`
	Description string                `json:"description"`
	SubQuestion string                `json:"sub_question,omitzero"`
	SubTypes    []ReportOptionSubType `json:"sub_types,omitempty"`
}

type ReportOptionSubType struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

type ReportReason struct {
	Reason      int64  `json:"reason"`
	Label       string `json:"label"`
	Description string `json:"description"`
}

type ReportSentExplicitContentFalsePositiveParams struct {
	ChannelID     snowflake.ID   `json:"channel_id"`
	MessageID     snowflake.ID   `json:"message_id"`
	AttachmentIDs []snowflake.ID `json:"attachment_ids"`
	Filenames     []string       `json:"filenames"`
}

type ReportUnverifiedApplicationParams struct {
	ReportVersion          int64                  `json:"report_version,omitzero"`
	Name                   string                 `json:"name"`
	Icon                   string                 `json:"icon"`
	OS                     OperatingSystem        `json:"os"`
	Executable             string                 `json:"executable,omitzero"`
	Publisher              string                 `json:"publisher,omitzero"`
	DistributorApplication ApplicationDistributor `json:"distributor_application,omitzero"`
}

type ReportUnverifiedApplicationResponse struct {
	Name        string   `json:"name"`
	Hash        string   `json:"hash"`
	MissingData []string `json:"missing_data"`
}

type RequestAgeVerificationForSuspendedUserParams struct {
	Token                string        `json:"token"`
	FromClassificationID *snowflake.ID `json:"from_classification_id,omitempty"`
}

type RequestAgeVerificationForSuspendedUserResponse struct {
	VerificationRequestID  string `json:"verification_request_id"`
	VerificationVendorName string `json:"verification_vendor_name"`
	VerificationWebviewURL string `json:"verification_webview_url"`
}

type RequestApplicationGatewayIntentsParams struct {
	ApplicationDescription                                             string           `json:"application_description"`
	IntentsFlagsRequested                                              ApplicationFlags `json:"intents_flags_requested,omitzero"`
	IntentsGatewayPresenceUseCaseDescription                           *string          `json:"intents_gateway_presence_use_case_description,omitempty"`
	IntentsGatewayPresenceUseCaseSupplementalMaterialDescription       *string          `json:"intents_gateway_presence_use_case_supplemental_material_description,omitempty"`
	IntentsGatewayPresenceStoreOffPlatform                             *bool            `json:"intents_gateway_presence_store_off_platform,omitempty"`
	IntentsGatewayPresenceRetention                                    *bool            `json:"intents_gateway_presence_retention,omitempty"`
	IntentsGatewayPresenceEncrypted                                    *bool            `json:"intents_gateway_presence_encrypted,omitempty"`
	IntentsGatewayPresenceOptOutStored                                 *bool            `json:"intents_gateway_presence_opt_out_stored,omitempty"`
	IntentsGatewayPresenceContactDeletion                              *string          `json:"intents_gateway_presence_contact_deletion,omitempty"`
	IntentsGatewayGuildMembersUseCaseDescription                       *string          `json:"intents_gateway_guild_members_use_case_description,omitempty"`
	IntentsGatewayGuildMembersUseCaseSupplementalMaterialDescription   *string          `json:"intents_gateway_guild_members_use_case_supplemental_material_description,omitempty"`
	IntentsGatewayGuildMembersStoreOffPlatform                         *bool            `json:"intents_gateway_guild_members_store_off_platform,omitempty"`
	IntentsGatewayGuildMembersRetention                                *bool            `json:"intents_gateway_guild_members_retention,omitempty"`
	IntentsGatewayGuildMembersEncrypted                                *bool            `json:"intents_gateway_guild_members_encrypted,omitempty"`
	IntentsGatewayGuildMembersContactDeletion                          *string          `json:"intents_gateway_guild_members_contact_deletion,omitempty"`
	IntentsGatewayMessageContentUseCaseDescription                     *string          `json:"intents_gateway_message_content_use_case_description,omitempty"`
	IntentsGatewayMessageContentUseCaseSupplementalMaterialDescription *string          `json:"intents_gateway_message_content_use_case_supplemental_material_description,omitempty"`
	IntentsGatewayMessageContentStoreOffPlatform                       *bool            `json:"intents_gateway_message_content_store_off_platform,omitempty"`
	IntentsGatewayMessageContentRetention                              *bool            `json:"intents_gateway_message_content_retention,omitempty"`
	IntentsGatewayMessageContentEncrypted                              *bool            `json:"intents_gateway_message_content_encrypted,omitempty"`
	IntentsGatewayMessageContentOptOutStored                           *bool            `json:"intents_gateway_message_content_opt_out_stored,omitempty"`
	IntentsGatewayMessageContentAiTraining                             *bool            `json:"intents_gateway_message_content_ai_training,omitempty"`
	IntentsGatewayMessageContentPrivacyPolicyPublic                    *bool            `json:"intents_gateway_message_content_privacy_policy_public,omitempty"`
	IntentsGatewayMessageContentPrivacyPolicyLocation                  *string          `json:"intents_gateway_message_content_privacy_policy_location,omitempty"`
	IntentsGatewayMessageContentPrivacyPolicyExample                   *string          `json:"intents_gateway_message_content_privacy_policy_example,omitempty"`
	IntentsGatewayMessageContentContactDeletion                        *string          `json:"intents_gateway_message_content_contact_deletion,omitempty"`
}

type RequestCallConnect struct {
	ChannelID snowflake.ID `json:"channel_id"`
}

type RequestChannelInfo struct {
	GuildID snowflake.ID `json:"guild_id"`
	Fields  []string     `json:"fields"`
}

type RequestChannelMemberCount struct {
	GuildID   snowflake.ID `json:"guild_id"`
	ChannelID snowflake.ID `json:"channel_id"`
}

type RequestChannelStatuses struct {
	GuildID snowflake.ID `json:"guild_id"`
}

type RequestClassificationReviewForSuspendedUserParams struct {
	Token     string                `json:"token"`
	Signal    AppealIngestionSignal `json:"signal"`
	UserInput string                `json:"user_input"`
}

type RequestClassificationReviewForSuspendedUserResponse struct {
	AppealID snowflake.ID `json:"appeal_id"`
}

type RequestClassificationReviewParams struct {
	Signal    AppealIngestionSignal `json:"signal"`
	UserInput string                `json:"user_input"`
}

type RequestClassificationReviewResponse struct {
	AppealID snowflake.ID `json:"appeal_id"`
}

type RequestGuildMembers struct {
	GuildID   json.RawMessage `json:"guild_id"`
	Query     string          `json:"query,omitzero"`
	Limit     int64           `json:"limit,omitzero"`
	Presences bool            `json:"presences,omitzero"`
	UserIDs   json.RawMessage `json:"user_ids,omitempty"`
	Nonce     string          `json:"nonce,omitzero"`
}

type RequestLastMessages struct {
	GuildID    snowflake.ID   `json:"guild_id"`
	ChannelIDs []snowflake.ID `json:"channel_ids"`
}

type RequestProxyTicketResponse struct {
	Ticket string `json:"ticket"`
}

type RequestReportReviewParams struct {
	Token string `json:"token"`
}

type RequestSoundboardSounds struct {
	GuildIDs []snowflake.ID `json:"guild_ids"`
}

type ResendPaymentVerificationEmailParams struct {
	PurchaseToken string `json:"purchase_token"`
}

type ResendVerificationCodeParams struct {
	Phone string `json:"phone"`
}

type ResetApplicationBotTokenResponse struct {
	Token string `json:"token"`
}

type ResetApplicationSecretResponse struct {
	Secret string `json:"secret"`
}

type ResetPasswordParams struct {
	Token            string                   `json:"token"`
	Password         string                   `json:"password"`
	Source           string                   `json:"source,omitzero"`
	Method           AuthenticatorType        `json:"method,omitzero"`
	Ticket           string                   `json:"ticket,omitzero"`
	Code             string                   `json:"code,omitzero"`
	PushProvider     PushNotificationProvider `json:"push_provider,omitzero"`
	PushToken        string                   `json:"push_token,omitzero"`
	PushVoipProvider PushNotificationProvider `json:"push_voip_provider,omitzero"`
	PushVoipToken    string                   `json:"push_voip_token,omitzero"`
}

type ResetPasswordResponse struct {
	Token    string       `json:"token,omitzero"`
	UserID   snowflake.ID `json:"user_id,omitzero"`
	Ticket   string       `json:"ticket,omitzero"`
	MFA      bool         `json:"mfa,omitzero"`
	Totp     bool         `json:"totp,omitzero"`
	SMS      bool         `json:"sms,omitzero"`
	Backup   bool         `json:"backup,omitzero"`
	Webauthn *string      `json:"webauthn,omitempty"`
}

type ResolveAutomodIncidentParams struct {
	AlertMessageID snowflake.ID                `json:"alert_message_id"`
	Reason         AutomodRaidResolutionReason `json:"reason"`
}

type ResolvedData struct {
	Users       map[snowflake.ID]PartialUser     `json:"users,omitempty"`
	Members     map[snowflake.ID]json.RawMessage `json:"members,omitempty"`
	Roles       map[snowflake.ID]Role            `json:"roles,omitempty"`
	Channels    map[snowflake.ID]PartialChannel  `json:"channels,omitempty"`
	Messages    map[snowflake.ID]Message         `json:"messages,omitempty"`
	Attachments map[snowflake.ID]Attachment      `json:"attachments,omitempty"`
}

type ResourceChannel struct {
	ChannelID   snowflake.ID `json:"channel_id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Emoji       Emoji        `json:"emoji,omitzero"`
	Icon        string       `json:"icon,omitzero"`
}

type ResponseBodyExtraFields struct {
	ShowVerificationForm bool          `json:"show_verification_form,omitzero"`
	WelcomeScreen        WelcomeScreen `json:"welcome_screen,omitzero"`
}

type Resume struct {
	Token     string `json:"token"`
	SessionID string `json:"session_id"`
	Seq       int64  `json:"seq"`
}

type Resumed struct {
	Trace []string `json:"_trace"`
}

type ResyncGuildChannels struct {
	GuildID              snowflake.ID   `json:"guild_id"`
	ObfuscatedChannelIDs []snowflake.ID `json:"obfuscated_channel_ids"`
}

type ReverifyPhoneNumberParams struct {
	Phone             string            `json:"phone,omitzero"`
	PhoneToken        string            `json:"phone_token,omitzero"`
	Password          string            `json:"password,omitzero"`
	ChangePhoneReason ChangePhoneReason `json:"change_phone_reason,omitzero"`
}

type RevertAccountParams struct {
	Token    string `json:"token"`
	Password string `json:"password"`
}

type RevertAccountResponse struct {
	Email string `json:"email"`
}

type RevokeOAuth2TokenParams struct {
	Token        string       `json:"token"`
	ClientID     snowflake.ID `json:"client_id,omitzero"`
	ClientSecret string       `json:"client_secret,omitzero"`
}

type RewardHeroShopBlock struct {
	Type                   ShopBlockType  `json:"type"`
	CategorySKUID          snowflake.ID   `json:"category_sku_id,omitzero"`
	CategoryStoreListingID snowflake.ID   `json:"category_store_listing_id"`
	Name                   string         `json:"name"`
	Summary                string         `json:"summary"`
	BannerAsset            AssetConfig    `json:"banner_asset"`
	LogoURL                string         `json:"logo_url"`
	Title                  string         `json:"title"`
	BannerTextColor        string         `json:"banner_text_color,omitzero"`
	BannerDisplayConfig    BannerAsset    `json:"banner_display_config"`
	HeroRiveURL            string         `json:"hero_rive_url,omitzero"`
	LogoDisplayConfig      BannerAsset    `json:"logo_display_config"`
	MobileTitle            string         `json:"mobile_title,omitzero"`
	MobileSummary          string         `json:"mobile_summary,omitzero"`
	MobileProductsTitle    string         `json:"mobile_products_title,omitzero"`
	HeroBannerURL          string         `json:"hero_banner_url"`
	HeroBannerAnimatedURL  string         `json:"hero_banner_animated_url"`
	HeroLogoURL            string         `json:"hero_logo_url"`
	MobileHeroURL          string         `json:"mobile_hero_url,omitzero"`
	MobileHeroAnimatedURL  string         `json:"mobile_hero_animated_url,omitzero"`
	RankedSKUIDs           []snowflake.ID `json:"ranked_sku_ids"`
	UnpublishedAt          *time.Time     `json:"unpublished_at"`
	RewardSKUID            snowflake.ID   `json:"reward_sku_id"`
}

type RingChannelRecipientsParams struct {
	Recipients []snowflake.ID `json:"recipients,omitempty"`
}

type Role struct {
	ID           snowflake.ID `json:"id"`
	Name         string       `json:"name"`
	Description  *string      `json:"description"`
	Color        int64        `json:"color"`
	Colors       RoleColors   `json:"colors"`
	Hoist        bool         `json:"hoist"`
	Icon         *string      `json:"icon,omitempty"`
	UnicodeEmoji *string      `json:"unicode_emoji,omitempty"`
	Position     int64        `json:"position"`
	Permissions  Permissions  `json:"permissions"`
	Managed      bool         `json:"managed"`
	Mentionable  bool         `json:"mentionable"`
	Flags        RoleFlags    `json:"flags,omitzero"`
	Tags         RoleTags     `json:"tags,omitzero"`
	Version      string       `json:"version,omitzero"`
}

type RoleColors struct {
	PrimaryColor   int64  `json:"primary_color"`
	SecondaryColor *int64 `json:"secondary_color"`
	TertiaryColor  *int64 `json:"tertiary_color"`
}

type RoleConnectionRequirement struct {
	ConnectionType          ConnectionType              `json:"connection_type"`
	ConnectionMetadataField *string                     `json:"connection_metadata_field,omitempty"`
	Operator                *RoleConnectionOperatorType `json:"operator,omitempty"`
	Value                   *string                     `json:"value,omitempty"`
	ApplicationID           snowflake.ID                `json:"application_id,omitzero"`
	Application             IntegrationApplication      `json:"application,omitzero"`
	Name                    string                      `json:"name,omitzero"`
	Description             string                      `json:"description,omitzero"`
	Result                  bool                        `json:"result,omitzero"`
}

type RoleConnectionRule struct {
	RoleID       snowflake.ID                            `json:"role_id"`
	Rules        json.RawMessage                         `json:"rules"`
	Applications map[snowflake.ID]IntegrationApplication `json:"applications"`
}

type RoleSelect struct {
	Type          ComponentType     `json:"type"`
	ID            int64             `json:"id,omitzero"`
	CustomID      string            `json:"custom_id"`
	Placeholder   string            `json:"placeholder,omitzero"`
	DefaultValues []json.RawMessage `json:"default_values,omitempty"`
	MinValues     int64             `json:"min_values,omitzero"`
	MaxValues     int64             `json:"max_values,omitzero"`
	Required      bool              `json:"required,omitzero"`
	Disabled      bool              `json:"disabled,omitzero"`
}

type RoleSubscriptionBenefit struct {
	EmojiID     *snowflake.ID               `json:"emoji_id"`
	EmojiName   *string                     `json:"emoji_name"`
	Name        string                      `json:"name"`
	Description string                      `json:"description"`
	RefType     RoleSubscriptionBenefitType `json:"ref_type"`
	RefID       *snowflake.ID               `json:"ref_id"`
}

type RoleSubscriptionBenefitChannel struct {
	ID   snowflake.ID `json:"id"`
	Name string       `json:"name"`
}

type RoleSubscriptionBenefits struct {
	SKUID    snowflake.ID              `json:"sku_id"`
	Benefits []RoleSubscriptionBenefit `json:"benefits"`
}

type RoleSubscriptionGroupListing struct {
	ID                      snowflake.ID                     `json:"id"`
	GuildID                 snowflake.ID                     `json:"guild_id"`
	ApplicationID           snowflake.ID                     `json:"application_id"`
	Name                    string                           `json:"name"`
	Description             string                           `json:"description"`
	SubscriptionListingsIDs []snowflake.ID                   `json:"subscription_listings_ids,omitempty"`
	SubscriptionListings    []RoleSubscriptionListing        `json:"subscription_listings,omitempty"`
	BenefitChannels         []RoleSubscriptionBenefitChannel `json:"benefit_channels,omitempty"`
}

type RoleSubscriptionListing struct {
	ID                snowflake.ID             `json:"id"`
	GuildID           snowflake.ID             `json:"guild_id"`
	ApplicationID     snowflake.ID             `json:"application_id"`
	Name              string                   `json:"name"`
	Description       string                   `json:"description"`
	ImageAsset        StoreAsset               `json:"image_asset"`
	SubscriptionPlans []SubscriptionPlan       `json:"subscription_plans"`
	RoleBenefits      RoleSubscriptionBenefits `json:"role_benefits"`
	RoleID            snowflake.ID             `json:"role_id"`
	Published         bool                     `json:"published"`
	SoftDeleted       bool                     `json:"soft_deleted"`
	Archived          bool                     `json:"archived"`
}

type RoleSubscriptionListingTemplate struct {
	Category                          string                            `json:"category"`
	CategoryImage                     string                            `json:"category_image"`
	UnselectedDarkThemeCategoryImage  string                            `json:"unselected_dark_theme_category_image"`
	UnselectedLightThemeCategoryImage string                            `json:"unselected_light_theme_category_image"`
	Listings                          []TemplateRoleSubscriptionListing `json:"listings"`
}

type RoleSubscriptionSettings struct {
	GuildID                           snowflake.ID                  `json:"guild_id"`
	FullServerGate                    bool                          `json:"full_server_gate"`
	Description                       *string                       `json:"description,omitempty"`
	ServerShopTabOrder                GuildProductShopTabsOrderType `json:"server_shop_tab_order"`
	StorePagePrimaryColor             *int64                        `json:"store_page_primary_color,omitempty"`
	StorePageTrailerURL               *string                       `json:"store_page_trailer_url,omitempty"`
	StorePageShowSubscriberCount      bool                          `json:"store_page_show_subscriber_count,omitzero"`
	StorePageGuildProductsDefaultSort GuildProductSortOrderType     `json:"store_page_guild_products_default_sort,omitzero"`
	CoverImageAsset                   StoreAsset                    `json:"cover_image_asset,omitzero"`
	StorePageSlug                     string                        `json:"store_page_slug"`
}

type RoleSubscriptionTrial struct {
	ID                     snowflake.ID      `json:"id"`
	NumActiveTrialUsers    int64             `json:"num_active_trial_users"`
	MaxNumActiveTrialUsers *int64            `json:"max_num_active_trial_users"`
	ActiveTrial            SubscriptionTrial `json:"active_trial,omitzero"`
}

type RoleTags struct {
	BotID                 snowflake.ID    `json:"bot_id,omitzero"`
	IntegrationID         snowflake.ID    `json:"integration_id,omitzero"`
	PremiumSubscriber     json.RawMessage `json:"premium_subscriber,omitempty"`
	SubscriptionListingID snowflake.ID    `json:"subscription_listing_id,omitzero"`
	AvailableForPurchase  json.RawMessage `json:"available_for_purchase,omitempty"`
	GuildConnections      json.RawMessage `json:"guild_connections,omitempty"`
}

type RtpPacket struct {
	Sequence  json.RawMessage `json:"Sequence"`
	Timestamp json.RawMessage `json:"Timestamp"`
	SSRC      json.RawMessage `json:"SSRC"`
	Csrcs     []int64         `json:"CSRCs,omitempty"`
	Extension json.RawMessage `json:"Extension,omitempty"`
	Payload   json.RawMessage `json:"Payload"`
}

type SKU struct {
	ID                             snowflake.ID                  `json:"id"`
	Type                           SKUType                       `json:"type"`
	ApplicationID                  snowflake.ID                  `json:"application_id"`
	Application                    *PartialApplication           `json:"application,omitempty"`
	ProductLine                    *SKUProductLine               `json:"product_line"`
	ProductID                      snowflake.ID                  `json:"product_id,omitzero"`
	Flags                          SKUFlags                      `json:"flags"`
	Name                           json.RawMessage               `json:"name"`
	Summary                        json.RawMessage               `json:"summary,omitempty"`
	Description                    json.RawMessage               `json:"description,omitempty"`
	LegalNotice                    json.RawMessage               `json:"legal_notice,omitempty"`
	Slug                           string                        `json:"slug"`
	ThumbnailAssetID               snowflake.ID                  `json:"thumbnail_asset_id,omitzero"`
	DependentSKUID                 *snowflake.ID                 `json:"dependent_sku_id,omitempty"`
	BundledSKUs                    []SKU                         `json:"bundled_skus,omitempty"`
	BundledSKUIDs                  []snowflake.ID                `json:"bundled_sku_ids,omitempty"`
	AccessType                     SKUAccessType                 `json:"access_type"`
	ManifestLabels                 []snowflake.ID                `json:"manifest_labels,omitempty"`
	Features                       []int64                       `json:"features"`
	Locales                        []string                      `json:"locales,omitempty"`
	Genres                         []int64                       `json:"genres,omitempty"`
	AvailableRegions               []string                      `json:"available_regions,omitempty"`
	ContentRating                  ContentRating                 `json:"content_rating,omitzero"`
	ContentRatingAgency            ContentRatingAgency           `json:"content_rating_agency,omitzero"`
	ContentRatings                 map[int64]ContentRating       `json:"content_ratings,omitempty"`
	SystemRequirements             map[int64]SystemRequirements  `json:"system_requirements,omitempty"`
	Price                          SKUPrice                      `json:"price,omitzero"`
	PriceTier                      int64                         `json:"price_tier,omitzero"`
	SalePriceTier                  int64                         `json:"sale_price_tier,omitzero"`
	SalePrice                      map[string]int64              `json:"sale_price,omitempty"`
	CreatedAt                      json.RawMessage               `json:"created_at"`
	UpdatedAt                      json.RawMessage               `json:"updated_at"`
	ReleaseDate                    string                        `json:"release_date,omitzero"`
	PreorderApproximateReleaseDate string                        `json:"preorder_approximate_release_date,omitzero"`
	PreorderReleasedAt             json.RawMessage               `json:"preorder_released_at,omitempty"`
	ExternalPurchaseURL            string                        `json:"external_purchase_url,omitzero"`
	ExternalSKUStrategies          map[int64]ExternalSKUStrategy `json:"external_sku_strategies,omitempty"`
	EligiblePaymentGateways        []int64                       `json:"eligible_payment_gateways,omitempty"`
	Premium                        bool                          `json:"premium"`
	ShowAgeGate                    bool                          `json:"show_age_gate"`
	Restricted                     bool                          `json:"restricted,omitzero"`
	Exclusive                      bool                          `json:"exclusive,omitzero"`
	Deleted                        bool                          `json:"deleted,omitzero"`
	TenantMetadata                 StoreTenantMetadata           `json:"tenant_metadata,omitzero"`
	PowerupMetadata                GuildPowerupMetadata          `json:"powerup_metadata,omitzero"`
	OrbsReward                     int64                         `json:"orbs_reward,omitzero"`
}

type SKUPrice struct {
	Currency         string                 `json:"currency"`
	CurrencyExponent int64                  `json:"currency_exponent"`
	Amount           int64                  `json:"amount"`
	SaleAmount       int64                  `json:"sale_amount,omitzero"`
	SalePercentage   int64                  `json:"sale_percentage,omitzero"`
	Premium          map[int64]PremiumPrice `json:"premium,omitempty"`
}

type SafetyHub struct {
	Classifications      []Classification `json:"classifications"`
	GuildClassifications []Classification `json:"guild_classifications"`
	AccountStanding      AccountStanding  `json:"account_standing"`
	IsDsaEligible        bool             `json:"is_dsa_eligible"`
	IsAppealEligible     bool             `json:"is_appeal_eligible"`
	Username             string           `json:"username"`
	AppealEligibility    []int64          `json:"appeal_eligibility"`
}

type SafetyHubAction struct {
	ID           snowflake.ID `json:"id"`
	ActionType   ActionType   `json:"action_type"`
	Descriptions []string     `json:"descriptions"`
}

type SafetyPolicyNoticeEmbed struct {
	ClientVersionMessage string       `json:"client_version_message"`
	ClassificationID     snowflake.ID `json:"classification_id"`
	IncidentTime         float64      `json:"incident_time"`
}

type SafetySettings struct {
	SafetySettingsPreset           json.RawMessage       `json:"safety_settings_preset"`
	IgnoreProfileSpeedbumpDisabled bool                  `json:"ignore_profile_speedbump_disabled"`
	SpendingLimitSettings          SpendingLimitSettings `json:"spending_limit_settings,omitzero"`
}

type SafetySignals struct {
	UnusualDMActivityUntil     json.RawMessage `json:"unusual_dm_activity_until,omitempty"`
	CommunicationDisabledUntil json.RawMessage `json:"communication_disabled_until,omitempty"`
	UnusualAccountActivity     bool            `json:"unusual_account_activity,omitzero"`
	AutomodQuarantinedUsername bool            `json:"automod_quarantined_username,omitzero"`
}

type SafetySystemNotificationEmbed struct {
	ClientVersionMessage string                          `json:"client_version_message"`
	ClassificationID     snowflake.ID                    `json:"classification_id"`
	Body                 string                          `json:"body"`
	IconType             string                          `json:"icon_type"`
	Header               string                          `json:"header"`
	Timestamp            float64                         `json:"timestamp"`
	Theme                string                          `json:"theme"`
	Ctas                 SafetySystemNotificationCtaType `json:"ctas"`
	LearnMoreLink        string                          `json:"learn_more_link,omitzero"`
}

type SafetyWarning struct {
	ID               string            `json:"id"`
	Type             SafetyWarningType `json:"type"`
	Expiry           time.Time         `json:"expiry"`
	DismissTimestamp *time.Time        `json:"dismiss_timestamp"`
}

type SaveData struct {
	ChannelID      snowflake.ID `json:"channel_id"`
	MessageID      snowflake.ID `json:"message_id"`
	GuildID        snowflake.ID `json:"guild_id,omitzero"`
	SavedAt        time.Time    `json:"saved_at"`
	AuthorSummary  string       `json:"author_summary"`
	ChannelSummary string       `json:"channel_summary"`
	MessageSummary string       `json:"message_summary"`
	Notes          string       `json:"notes"`
	DueAt          *time.Time   `json:"due_at"`
}

type SaveMessageParams struct {
	DueAt *time.Time `json:"due_at,omitempty"`
}

type SavedMessage struct {
	Message  *Message `json:"message"`
	SaveData SaveData `json:"save_data"`
}

type SavedMessageDelete struct {
	ChannelID snowflake.ID `json:"channel_id"`
	MessageID snowflake.ID `json:"message_id"`
}

type ScanExplicitMediaParams struct {
	MessageIDs []snowflake.ID `json:"message_ids"`
}

type ScreenshareApplication struct {
	Name string `json:"name"`
}

type ScreenshareStateUpdate struct {
	Active      bool                    `json:"active"`
	Pid         *int64                  `json:"pid"`
	Application *ScreenshareApplication `json:"application"`
}

type SearchApplicationsDirectoryQuery struct {
	Query                                                 string                             `json:"query,omitzero"`
	GuildID                                               string                             `json:"guild_id,omitzero"`
	Page                                                  int64                              `json:"page,omitzero"`
	PageSize                                              int64                              `json:"page_size,omitzero"`
	CategoryID                                            int64                              `json:"category_id,omitzero"`
	Locale                                                string                             `json:"locale,omitzero"`
	MinUserInstallCommandCount                            int64                              `json:"min_user_install_command_count,omitzero"`
	ExcludeAppsWithCustomInstallURL                       bool                               `json:"exclude_apps_with_custom_install_url,omitzero"`
	ExcludeNonEmbeddedApps                                bool                               `json:"exclude_non_embedded_apps,omitzero"`
	ExcludeEmbeddedAppsWithoutPrimaryEntryPointAppCommand bool                               `json:"exclude_embedded_apps_without_primary_entry_point_app_command,omitzero"`
	Source                                                ApplicationDirectoryRequestSurface `json:"source,omitzero"`
}

type SearchApplicationsDirectoryResponse struct {
	Results          json.RawMessage              `json:"results"`
	NumPages         int64                        `json:"num_pages"`
	CountsByCategory map[int64]int64              `json:"counts_by_category"`
	Type             ApplicationDirectoryItemType `json:"type"`
	LoadID           string                       `json:"load_id"`
}

type SearchCollectiblesQuery struct {
	ItemTypes     []string     `json:"item_types,omitempty"`
	Colors        []string     `json:"colors,omitempty"`
	Themes        []string     `json:"themes,omitempty"`
	OrbsEligible  bool         `json:"orbs_eligible,omitzero"`
	Offset        int64        `json:"offset,omitzero"`
	Limit         int64        `json:"limit,omitzero"`
	SortType      ShopSortType `json:"sort_type,omitzero"`
	SortDirection string       `json:"sort_direction,omitzero"`
	Search        string       `json:"search,omitzero"`
}

type SearchCollectiblesResponse struct {
	Pagination Pagination     `json:"pagination"`
	SKUs       []snowflake.ID `json:"skus"`
}

type SearchCompaniesQuery struct {
	Name string `json:"name,omitzero"`
}

type SearchCursor struct {
	Type      SearchCursorType  `json:"type"`
	Timestamp snowflake.ID      `json:"timestamp"`
	Score     SearchCursorScore `json:"score,omitzero"`
}

type SearchCursorScore struct {
	Timestamp snowflake.ID `json:"timestamp"`
	Score     float64      `json:"score,omitzero"`
}

type SearchDiscoverableGuildsQuery struct {
	Query      string `json:"query"`
	Limit      int64  `json:"limit,omitzero"`
	Offset     int64  `json:"offset,omitzero"`
	CategoryID int64  `json:"category_id,omitzero"`
}

type SearchGuildBansQuery struct {
	Query   string         `json:"query,omitzero"`
	UserIDs []snowflake.ID `json:"user_ids,omitempty"`
	Limit   int64          `json:"limit,omitzero"`
}

type SearchGuildMembersParams struct {
	Limit    int64                  `json:"limit,omitzero"`
	Sort     MemberSortType         `json:"sort,omitzero"`
	OrQuery  MemberFilter           `json:"or_query,omitzero"`
	AndQuery MemberFilter           `json:"and_query,omitzero"`
	Before   MemberPaginationFilter `json:"before,omitzero"`
	After    MemberPaginationFilter `json:"after,omitzero"`
}

type SearchGuildMembersResponse struct {
	GuildID          snowflake.ID              `json:"guild_id"`
	Members          []SupplementalGuildMember `json:"members"`
	PageResultCount  int64                     `json:"page_result_count"`
	TotalResultCount int64                     `json:"total_result_count"`
}

type SearchGuildMessagesByTabParams struct {
	Tabs                map[string]SearchTab `json:"tabs"`
	ChannelIDs          []snowflake.ID       `json:"channel_ids,omitempty"`
	IncludeNSFW         bool                 `json:"include_nsfw,omitzero"`
	TrackExactTotalHits bool                 `json:"track_exact_total_hits,omitzero"`
}

type SearchGuildMessagesByTabResponse struct {
	AnalyticsID              string          `json:"analytics_id"`
	DoingDeepHistoricalIndex bool            `json:"doing_deep_historical_index"`
	DocumentsIndexed         int64           `json:"documents_indexed,omitzero"`
	Tabs                     json.RawMessage `json:"tabs"`
}

type SearchGuildMessagesQuery struct {
	Limit               int64          `json:"limit,omitzero"`
	Offset              int64          `json:"offset,omitzero"`
	MaxID               snowflake.ID   `json:"max_id,omitzero"`
	MinID               snowflake.ID   `json:"min_id,omitzero"`
	Slop                int64          `json:"slop,omitzero"`
	Content             string         `json:"content,omitzero"`
	Contents            []string       `json:"contents,omitempty"`
	ChannelID           []snowflake.ID `json:"channel_id,omitempty"`
	AuthorType          []string       `json:"author_type,omitempty"`
	AuthorID            []snowflake.ID `json:"author_id,omitempty"`
	Mentions            []snowflake.ID `json:"mentions,omitempty"`
	MentionsRoleID      []snowflake.ID `json:"mentions_role_id,omitempty"`
	MentionEveryone     bool           `json:"mention_everyone,omitzero"`
	RepliedToUserID     []snowflake.ID `json:"replied_to_user_id,omitempty"`
	RepliedToMessageID  []snowflake.ID `json:"replied_to_message_id,omitempty"`
	Pinned              bool           `json:"pinned,omitzero"`
	Has                 []string       `json:"has,omitempty"`
	EmbedType           []string       `json:"embed_type,omitempty"`
	EmbedProvider       []string       `json:"embed_provider,omitempty"`
	LinkHostname        []string       `json:"link_hostname,omitempty"`
	AttachmentFilename  []string       `json:"attachment_filename,omitempty"`
	AttachmentExtension []string       `json:"attachment_extension,omitempty"`
	CommandID           snowflake.ID   `json:"command_id,omitzero"`
	CommandName         string         `json:"command_name,omitzero"`
	SortBy              SearchSortMode `json:"sort_by,omitzero"`
	SortOrder           string         `json:"sort_order,omitzero"`
	IncludeNSFW         bool           `json:"include_nsfw,omitzero"`
}

type SearchGuildMessagesResponse struct {
	AnalyticsID              string         `json:"analytics_id"`
	DoingDeepHistoricalIndex bool           `json:"doing_deep_historical_index"`
	DocumentsIndexed         int64          `json:"documents_indexed,omitzero"`
	TotalResults             int64          `json:"total_results"`
	Messages                 [][]Message    `json:"messages"`
	Channels                 []Channel      `json:"channels,omitempty"`
	Threads                  []Channel      `json:"threads,omitempty"`
	Members                  []ThreadMember `json:"members,omitempty"`
}

type SearchPublishedGuildsQuery struct {
	Query  string `json:"query,omitzero"`
	Limit  int64  `json:"limit,omitzero"`
	Offset int64  `json:"offset,omitzero"`
}

type SearchRecentMembers struct {
	GuildID           snowflake.ID  `json:"guild_id"`
	Query             string        `json:"query"`
	ContinuationToken *snowflake.ID `json:"continuation_token"`
	Nonce             string        `json:"nonce,omitzero"`
}

type SearchTab struct {
	Limit               int64          `json:"limit,omitzero"`
	Cursor              *SearchCursor  `json:"cursor,omitempty"`
	Offset              int64          `json:"offset,omitzero"`
	MaxID               snowflake.ID   `json:"max_id,omitzero"`
	MinID               snowflake.ID   `json:"min_id,omitzero"`
	Slop                int64          `json:"slop,omitzero"`
	Content             string         `json:"content,omitzero"`
	Contents            []string       `json:"contents,omitempty"`
	AuthorType          []string       `json:"author_type,omitempty"`
	AuthorID            []snowflake.ID `json:"author_id,omitempty"`
	Mentions            []snowflake.ID `json:"mentions,omitempty"`
	MentionsRoleID      []snowflake.ID `json:"mentions_role_id,omitempty"`
	MentionEveryone     bool           `json:"mention_everyone,omitzero"`
	RepliedToUserID     []snowflake.ID `json:"replied_to_user_id,omitempty"`
	RepliedToMessageID  []snowflake.ID `json:"replied_to_message_id,omitempty"`
	Pinned              bool           `json:"pinned,omitzero"`
	Has                 []string       `json:"has,omitempty"`
	EmbedType           []string       `json:"embed_type,omitempty"`
	EmbedProvider       []string       `json:"embed_provider,omitempty"`
	LinkHostname        []string       `json:"link_hostname,omitempty"`
	AttachmentFilename  []string       `json:"attachment_filename,omitempty"`
	AttachmentExtension []string       `json:"attachment_extension,omitempty"`
	CommandID           snowflake.ID   `json:"command_id,omitzero"`
	CommandName         string         `json:"command_name,omitzero"`
	SortBy              SearchSortMode `json:"sort_by,omitzero"`
	SortOrder           string         `json:"sort_order,omitzero"`
}

type SearchTabResults struct {
	Messages     [][]Message    `json:"messages"`
	Channels     []Channel      `json:"channels"`
	Threads      []Channel      `json:"threads,omitempty"`
	Members      []ThreadMember `json:"members,omitempty"`
	Cursor       SearchCursor   `json:"cursor"`
	TotalResults int64          `json:"total_results"`
	TimeSpentMs  int64          `json:"time_spent_ms"`
}

type SearchThreadsQuery struct {
	Name       string           `json:"name,omitzero"`
	Slop       int64            `json:"slop,omitzero"`
	Tag        []snowflake.ID   `json:"tag,omitempty"`
	TagSetting SearchTagSetting `json:"tag_setting,omitzero"`
	Archived   bool             `json:"archived,omitzero"`
	SortBy     ThreadSortMode   `json:"sort_by,omitzero"`
	SortOrder  string           `json:"sort_order,omitzero"`
	Limit      int64            `json:"limit,omitzero"`
	Offset     int64            `json:"offset,omitzero"`
	MaxID      snowflake.ID     `json:"max_id,omitzero"`
	MinID      snowflake.ID     `json:"min_id,omitzero"`
}

type SearchThreadsResponse struct {
	Threads       []Channel      `json:"threads"`
	Members       []ThreadMember `json:"members"`
	HasMore       bool           `json:"has_more"`
	TotalResults  int64          `json:"total_results"`
	FirstMessages []Message      `json:"first_messages,omitempty"`
}

type Section struct {
	Type       ComponentType `json:"type"`
	ID         int64         `json:"id,omitzero"`
	Components Components    `json:"components"`
	Accessory  AnyComponent  `json:"accessory"`
}

type SelectDefaultValue struct {
	ID   snowflake.ID `json:"id"`
	Type string       `json:"type"`
}

type SelectOption struct {
	Label       string `json:"label"`
	Value       string `json:"value"`
	Description string `json:"description,omitzero"`
	Emoji       Emoji  `json:"emoji,omitzero"`
	Default     bool   `json:"default,omitzero"`
}

type SelectProtocol struct {
	Protocol        ProtocolType    `json:"protocol"`
	Data            json.RawMessage `json:"data"`
	RTCConnectionID string          `json:"rtc_connection_id,omitzero"`
	Codecs          []Codec         `json:"codecs,omitempty"`
	Experiments     []string        `json:"experiments,omitempty"`
}

type SelectTextChannelArguments struct {
	ChannelID *snowflake.ID `json:"channel_id"`
	Timeout   int64         `json:"timeout,omitzero"`
}

type SelectVoiceChannelArguments struct {
	ChannelID *snowflake.ID `json:"channel_id"`
	Timeout   int64         `json:"timeout,omitzero"`
	Force     bool          `json:"force,omitzero"`
	Navigate  bool          `json:"navigate,omitzero"`
}

type SelfHarmContentSettings struct {
	SelfHarmContentGuilds      json.RawMessage `json:"self_harm_content_guilds"`
	SelfHarmContentFriendDM    json.RawMessage `json:"self_harm_content_friend_dm"`
	SelfHarmContentNonFriendDM json.RawMessage `json:"self_harm_content_non_friend_dm"`
}

type SendActivityJoinInviteArguments struct {
	UserID  snowflake.ID       `json:"user_id"`
	Type    ActivityActionType `json:"type"`
	Content string             `json:"content,omitzero"`
	Pid     int64              `json:"pid"`
}

type SendAnalyticsEventArguments struct {
	EventName       string          `json:"event_name"`
	EventProperties json.RawMessage `json:"event_properties"`
}

type SendBackupCodesChallengeParams struct {
	Password string `json:"password"`
}

type SendBackupCodesChallengeResponse struct {
	Nonce           string `json:"nonce"`
	RegenerateNonce string `json:"regenerate_nonce"`
}

type SendConsoleCommandParams struct {
	Command   ConsoleCommandType `json:"command"`
	ChannelID snowflake.ID       `json:"channel_id"`
	GuildID   snowflake.ID       `json:"guild_id,omitzero"`
	Nonce     string             `json:"nonce,omitzero"`
}

type SendConsoleCommandResponse struct {
	ID snowflake.ID `json:"id"`
}

type SendCustomCallSoundParams struct {
	AnimationType VoiceChannelEffectAnimationType `json:"animation_type"`
	AnimationID   int64                           `json:"animation_id"`
}

type SendFriendRequestParams struct {
	Username      string  `json:"username"`
	Discriminator *string `json:"discriminator"`
	Note          *string `json:"note,omitempty"`
}

type SendGameFriendRequestParams struct {
	Username string `json:"username"`
}

type SendMFASMSParams struct {
	Ticket string `json:"ticket"`
}

type SendMFASMSResponse struct {
	Phone string `json:"phone"`
}

type SendQuestHeartbeatParams struct {
	StreamKey             string       `json:"stream_key,omitzero"`
	ApplicationID         snowflake.ID `json:"application_id,omitzero"`
	Terminal              bool         `json:"terminal,omitzero"`
	ExecutablePath        *string      `json:"executable_path,omitempty"`
	ExecutableFingerprint *string      `json:"executable_fingerprint,omitempty"`
}

type SendQuestVideoProgressParams struct {
	Timestamp int64 `json:"timestamp"`
}

type SendSoundboardSoundParams struct {
	SoundID       snowflake.ID  `json:"sound_id"`
	SourceGuildID *snowflake.ID `json:"source_guild_id,omitempty"`
}

type SendVoiceChannelEffectParams struct {
	AnimationType *VoiceChannelEffectAnimationType `json:"animation_type,omitempty"`
	AnimationID   *int64                           `json:"animation_id,omitempty"`
	EmojiID       *snowflake.ID                    `json:"emoji_id,omitempty"`
	EmojiName     *string                          `json:"emoji_name,omitempty"`
}

type SendableApplicationCommandData struct {
	ID                 snowflake.ID                   `json:"id"`
	Type               ApplicationCommandType         `json:"type,omitzero"`
	Name               string                         `json:"name"`
	Version            snowflake.ID                   `json:"version"`
	ApplicationCommand ApplicationCommand             `json:"application_command,omitzero"`
	Options            []ApplicationCommandDataOption `json:"options,omitempty"`
	TargetID           snowflake.ID                   `json:"target_id,omitzero"`
	Attachments        []Attachment                   `json:"attachments,omitempty"`
}

type SendableMessageComponentData struct {
	ComponentType ComponentType   `json:"component_type,omitzero"`
	Type          ComponentType   `json:"type,omitzero"`
	CustomID      string          `json:"custom_id"`
	Values        json.RawMessage `json:"values,omitempty"`
}

type SendableModalSubmitData struct {
	ID          snowflake.ID               `json:"id"`
	CustomID    string                     `json:"custom_id"`
	Components  []ModalSubmitComponentData `json:"components"`
	Attachments []Attachment               `json:"attachments,omitempty"`
}

type Separator struct {
	Type    ComponentType        `json:"type"`
	ID      int64                `json:"id,omitzero"`
	Divider bool                 `json:"divider,omitzero"`
	Spacing SeparatorSpacingType `json:"spacing,omitzero"`
}

type Session struct {
	SessionID        string     `json:"session_id"`
	ClientInfo       ClientInfo `json:"client_info"`
	Status           StatusType `json:"status"`
	Activities       []Activity `json:"activities"`
	HiddenActivities []Activity `json:"hidden_activities"`
	Active           bool       `json:"active,omitzero"`
}

type SessionDescription struct {
	AudioCodec                    string                  `json:"audio_codec"`
	VideoCodec                    string                  `json:"video_codec"`
	MediaSessionID                string                  `json:"media_session_id"`
	Mode                          TransportEncryptionMode `json:"mode,omitzero"`
	SecretKey                     []int64                 `json:"secret_key,omitempty"`
	Sdp                           string                  `json:"sdp,omitzero"`
	KeyframeInterval              int64                   `json:"keyframe_interval,omitzero"`
	BandwidthEstimationExperiment string                  `json:"bandwidth_estimation_experiment,omitzero"`
	DaveProtocolVersion           int64                   `json:"dave_protocol_version,omitzero"`
}

type SessionStartLimit struct {
	Total          int64 `json:"total"`
	Remaining      int64 `json:"remaining"`
	ResetAfter     int64 `json:"reset_after"`
	MaxConcurrency int64 `json:"max_concurrency"`
}

type SetActivityArguments struct {
	Pid      int64        `json:"pid,omitzero"`
	Activity *RPCActivity `json:"activity"`
}

type SetApplicationEmbeddabilityParams struct {
	Embedded bool `json:"embedded"`
}

type SetCertifiedDevicesArguments struct {
	Devices []CertifiedDevice `json:"devices"`
}

type SetGuildIdentityParams struct {
	IdentityEnabled *bool         `json:"identity_enabled,omitempty"`
	IdentityGuildID *snowflake.ID `json:"identity_guild_id,omitempty"`
}

type SetOrientationLockStateArguments struct {
	LockState                 OrientationState  `json:"lock_state"`
	PictureInPictureLockState *OrientationState `json:"picture_in_picture_lock_state,omitempty"`
	GridLockState             *OrientationState `json:"grid_lock_state,omitempty"`
}

type SetOverlayLockedArguments struct {
	Locked bool  `json:"locked"`
	Pid    int64 `json:"pid"`
}

type SetStreamPaused struct {
	StreamKey string `json:"stream_key"`
	Paused    bool   `json:"paused"`
}

type SetUserVoiceSettingsArguments struct {
	UserID snowflake.ID `json:"user_id"`
	Volume int64        `json:"volume,omitzero"`
	Mute   bool         `json:"mute,omitzero"`
}

type SetVoiceSettings2Arguments struct {
	InputMode RPCVoiceInputMode `json:"input_mode,omitzero"`
	SelfMute  bool              `json:"self_mute,omitzero"`
	SelfDeaf  bool              `json:"self_deaf,omitzero"`
}

type SetVoiceSettingsArguments struct {
	Input                RPCVoiceIoSettings   `json:"input,omitzero"`
	Output               RPCVoiceIoSettings   `json:"output,omitzero"`
	Mode                 RPCVoiceSettingsMode `json:"mode,omitzero"`
	AutomaticGainControl bool                 `json:"automatic_gain_control,omitzero"`
	EchoCancellation     bool                 `json:"echo_cancellation,omitzero"`
	NoiseSuppression     bool                 `json:"noise_suppression,omitzero"`
	QoS                  bool                 `json:"qos,omitzero"`
	SilenceWarning       bool                 `json:"silence_warning,omitzero"`
	Deaf                 bool                 `json:"deaf,omitzero"`
	Mute                 bool                 `json:"mute,omitzero"`
}

type ShareInteractionArguments struct {
	Command              string                          `json:"command"`
	Options              []InteractionOption             `json:"options,omitempty"`
	Content              string                          `json:"content,omitzero"`
	RequireLaunchChannel bool                            `json:"require_launch_channel,omitzero"`
	PreviewImage         InteractionResponsePreviewImage `json:"preview_image,omitzero"`
	Components           []ActionRow                     `json:"components,omitempty"`
	Pid                  int64                           `json:"pid,omitzero"`
}

type ShareInteractionResponse struct {
	Success bool `json:"success"`
}

type ShareLinkArguments struct {
	CustomID string `json:"custom_id,omitzero"`
	Message  string `json:"message"`
	LinkID   string `json:"link_id,omitzero"`
}

type ShareLinkResponse struct {
	Success        bool `json:"success"`
	Didcopylink    bool `json:"didCopyLink"`
	Didsendmessage bool `json:"didSendMessage"`
}

type SharedClientTheme struct {
	Colors        []string                `json:"colors"`
	GradientAngle int64                   `json:"gradient_angle"`
	BaseMix       int64                   `json:"base_mix"`
	BaseTheme     *UserSettingsProtoTheme `json:"base_theme,omitempty"`
}

type ShelfShopBlock struct {
	Type          ShopBlockType  `json:"type"`
	CategorySKUID snowflake.ID   `json:"category_sku_id,omitzero"`
	Name          string         `json:"name"`
	RankedSKUIDs  []snowflake.ID `json:"ranked_sku_ids"`
}

type ShortcutKeyCombo struct {
	Type ShortcutKeyComboType `json:"type"`
	Code int64                `json:"code"`
	Name string               `json:"name,omitzero"`
}

type SnapshotMessage struct {
	Content          string            `json:"content"`
	Timestamp        time.Time         `json:"timestamp"`
	EditedTimestamp  *time.Time        `json:"edited_timestamp"`
	Mentions         []PartialUser     `json:"mentions"`
	MentionRoles     []snowflake.ID    `json:"mention_roles"`
	Attachments      []Attachment      `json:"attachments"`
	Embeds           []Embed           `json:"embeds"`
	Type             MessageType       `json:"type"`
	Flags            MessageFlags      `json:"flags"`
	Components       Components        `json:"components,omitzero"`
	Resolved         ResolvedData      `json:"resolved,omitzero"`
	StickerItems     []StickerItem     `json:"sticker_items,omitempty"`
	SoundboardSounds []SoundboardSound `json:"soundboard_sounds,omitempty"`
}

type SocialLayerMetadata struct {
	CarouselItems              []StoreCarouselItem `json:"carousel_items"`
	Label                      string              `json:"label"`
	ExpiresAt                  json.RawMessage     `json:"expires_at"`
	CardImageAssetID           snowflake.ID        `json:"card_image_asset_id,omitzero"`
	CardBackgroundImageAssetID snowflake.ID        `json:"card_background_image_asset_id,omitzero"`
	PriceTier                  int64               `json:"price_tier,omitzero"`
}

type SocialLayerSKUPurchaseEligibilityData struct {
	SKUID snowflake.ID `json:"sku_id"`
}

type SocialLayerSKUPurchaseEligibilityResponse struct {
	InteractionID snowflake.ID `json:"interaction_id"`
	ApplicationID snowflake.ID `json:"application_id"`
	SKUID         snowflake.ID `json:"sku_id"`
	RecipientID   snowflake.ID `json:"recipient_id"`
	Eligible      snowflake.ID `json:"eligible"`
}

type SocialLayerStorefrontPromotionalBannerShopBlock struct {
	Type           ShopBlockType  `json:"type"`
	ApplicationID  snowflake.ID   `json:"application_id"`
	HeaderText     string         `json:"header_text"`
	GradientColors []string       `json:"gradient_colors"`
	GradientAngle  int64          `json:"gradient_angle"`
	SKUIDs         []snowflake.ID `json:"sku_ids"`
	EndTime        time.Time      `json:"end_time"`
	CtaType        string         `json:"cta_type,omitzero"`
	LogoURL        string         `json:"logo_url,omitzero"`
}

type SocialSdkRelease struct {
	Version         string                     `json:"version"`
	ReleaseDateTime time.Time                  `json:"release_date_time"`
	Artifacts       []SocialSdkReleaseArtifact `json:"artifacts"`
}

type SocialSdkReleaseArtifact struct {
	DownloadURL string `json:"download_url"`
	Filename    string `json:"filename"`
	SizeBytes   int64  `json:"size_bytes"`
}

type SortedSKUIDs struct {
	Popular     []snowflake.ID `json:"popular"`
	Recommended []snowflake.ID `json:"recommended"`
}

type SoundboardSettings struct {
	Volume float64 `json:"volume"`
}

type SoundboardSound struct {
	SoundID   snowflake.ID  `json:"sound_id"`
	Name      string        `json:"name"`
	Volume    float64       `json:"volume"`
	EmojiID   *snowflake.ID `json:"emoji_id"`
	EmojiName *string       `json:"emoji_name"`
	GuildID   snowflake.ID  `json:"guild_id,omitzero"`
	Available bool          `json:"available"`
	User      PartialUser   `json:"user,omitzero"`
	UserID    snowflake.ID  `json:"user_id,omitzero"`
}

type SoundboardSoundsEvent struct {
	GuildID          snowflake.ID      `json:"guild_id"`
	SoundboardSounds []SoundboardSound `json:"soundboard_sounds"`
}

type Speaking struct {
	Speaking SpeakingFlags `json:"speaking"`
	SSRC     int64         `json:"ssrc"`
	UserID   snowflake.ID  `json:"user_id"`
	Delay    int64         `json:"delay,omitzero"`
}

type SpeakingStart struct {
	ChannelID snowflake.ID `json:"channel_id"`
	UserID    snowflake.ID `json:"user_id"`
}

type SpeakingStartSubscriptionArguments struct {
	ChannelID *snowflake.ID `json:"channel_id"`
}

type SpeakingStop struct {
	ChannelID snowflake.ID `json:"channel_id"`
	UserID    snowflake.ID `json:"user_id"`
}

type SpeakingStopSubscriptionArguments struct {
	ChannelID *snowflake.ID `json:"channel_id"`
}

type SpendingLimit struct {
	Amount   json.RawMessage `json:"amount"`
	Currency string          `json:"currency"`
}

type SpendingLimitSettings struct {
	OneTimePurchaseLimit SpendingLimit `json:"one_time_purchase_limit,omitzero"`
}

type StageInstance struct {
	ID                    snowflake.ID  `json:"id"`
	GuildID               snowflake.ID  `json:"guild_id"`
	ChannelID             snowflake.ID  `json:"channel_id"`
	Topic                 string        `json:"topic"`
	PrivacyLevel          PrivacyLevel  `json:"privacy_level"`
	InviteCode            *string       `json:"invite_code"`
	DiscoverableDisabled  bool          `json:"discoverable_disabled"`
	GuildScheduledEventID *snowflake.ID `json:"guild_scheduled_event_id"`
}

type StageInstanceCreate struct {
	ChannelID          snowflake.ID `json:"channel_id"`
	GuildID            snowflake.ID `json:"guild_id"`
	StageInstanceTopic string       `json:"stage_instance_topic,omitzero"`
}

type StageReportResponse struct {
	Token string `json:"token"`
}

type StartConditionalLoginResponse struct {
	Ticket    string `json:"ticket"`
	Challenge string `json:"challenge"`
}

type StartConsoleQuestQuery struct {
	Preview bool `json:"preview,omitzero"`
}

type StartConsoleQuestResponse struct {
	Started         bool             `json:"started"`
	QuestUserStatus *QuestUserStatus `json:"quest_user_status"`
	ErrorHints      []string         `json:"error_hints"`
	ErrorHintsV2    []QuestErrorHint `json:"error_hints_v2"`
}

type StartPasswordlessLoginResponse struct {
	Ticket    string `json:"ticket"`
	Challenge string `json:"challenge"`
}

type StartPremiumPurchaseArguments struct {
	Pid int64 `json:"pid,omitzero"`
}

type StartPurchaseArguments struct {
	SKUID snowflake.ID `json:"sku_id"`
	Pid   int64        `json:"pid,omitzero"`
}

type StatusSettings struct {
	Status            StatusType                    `json:"status,omitzero"`
	CustomStatus      UserSettingsProtoCustomStatus `json:"custom_status,omitzero"`
	ShowCurrentGame   bool                          `json:"show_current_game,omitzero"`
	StatusExpiresAtMs json.RawMessage               `json:"status_expires_at_ms"`
	StatusCreatedAtMs uint64                        `json:"status_created_at_ms,omitzero"`
}

type Sticker struct {
	ID          snowflake.ID       `json:"id"`
	PackID      snowflake.ID       `json:"pack_id,omitzero"`
	Name        string             `json:"name"`
	Description *string            `json:"description"`
	Tags        string             `json:"tags"`
	Type        StickerTypes       `json:"type"`
	FormatType  StickerFormatTypes `json:"format_type"`
	Available   bool               `json:"available,omitzero"`
	GuildID     snowflake.ID       `json:"guild_id,omitzero"`
	User        PartialUser        `json:"user,omitzero"`
	SortValue   int64              `json:"sort_value,omitzero"`
	Version     string             `json:"version,omitzero"`
}

type StickerFrecency struct {
	Stickers map[string]FrecencyItem `json:"stickers"`
}

type StickerItem struct {
	ID         snowflake.ID       `json:"id"`
	Name       string             `json:"name"`
	FormatType StickerFormatTypes `json:"format_type"`
}

type StickerPack struct {
	ID             snowflake.ID `json:"id"`
	Stickers       []Sticker    `json:"stickers"`
	Name           string       `json:"name"`
	SKUID          snowflake.ID `json:"sku_id"`
	CoverStickerID snowflake.ID `json:"cover_sticker_id,omitzero"`
	Description    string       `json:"description"`
	BannerAssetID  snowflake.ID `json:"banner_asset_id,omitzero"`
}

type StopRingingChannelRecipientsParams struct {
	Recipients []snowflake.ID `json:"recipients,omitempty"`
}

type StoreAsset struct {
	ID            snowflake.ID `json:"id"`
	ApplicationID snowflake.ID `json:"application_id,omitzero"`
	Size          int64        `json:"size"`
	MimeType      string       `json:"mime_type"`
	Filename      string       `json:"filename"`
	Width         int64        `json:"width"`
	Height        int64        `json:"height"`
}

type StoreCarouselItem struct {
	YoutubeVideoID    *string       `json:"youtube_video_id,omitempty"`
	AssetID           snowflake.ID  `json:"asset_id,omitzero"`
	ThumbnailAssetID  snowflake.ID  `json:"thumbnail_asset_id,omitzero"`
	BackgroundAssetID *snowflake.ID `json:"background_asset_id,omitempty"`
	Label             string        `json:"label,omitzero"`
	LabelIconAssetID  snowflake.ID  `json:"label_icon_asset_id,omitzero"`
}

type StoreCreateSKUPurchaseParams struct {
	Files json.RawMessage `json:"files"`
}

type StoreCreateStoreListingParams struct {
	ChildSKUIDs                 []snowflake.ID      `json:"child_sku_ids,omitempty"`
	Summary                     json.RawMessage     `json:"summary,omitempty"`
	Description                 json.RawMessage     `json:"description,omitempty"`
	Tagline                     json.RawMessage     `json:"tagline,omitempty"`
	Published                   bool                `json:"published,omitzero"`
	CarouselItems               []StoreCarouselItem `json:"carousel_items,omitempty"`
	GuildID                     snowflake.ID        `json:"guild_id,omitzero"`
	ThumbnailAssetID            snowflake.ID        `json:"thumbnail_asset_id,omitzero"`
	PreviewVideoAssetID         snowflake.ID        `json:"preview_video_asset_id,omitzero"`
	HeaderBackgroundAssetID     snowflake.ID        `json:"header_background_asset_id,omitzero"`
	HeaderLogoDarkThemeAssetID  snowflake.ID        `json:"header_logo_dark_theme_asset_id,omitzero"`
	HeaderLogoLightThemeAssetID snowflake.ID        `json:"header_logo_light_theme_asset_id,omitzero"`
	BoxArtAssetID               snowflake.ID        `json:"box_art_asset_id,omitzero"`
	HeroBackgroundAssetID       snowflake.ID        `json:"hero_background_asset_id,omitzero"`
	HeroVideoAssetID            snowflake.ID        `json:"hero_video_asset_id,omitzero"`
}

type StoreGetGuildApplicationStorefrontAnnouncementResponse struct {
	ID                     snowflake.ID `json:"id"`
	ApplicationID          snowflake.ID `json:"application_id"`
	ApplicationName        string       `json:"application_name"`
	AssetID                snowflake.ID `json:"asset_id"`
	BackgroundImageAssetID snowflake.ID `json:"background_image_asset_id"`
}

type StoreListing struct {
	ID                   snowflake.ID          `json:"id"`
	SKU                  SKU                   `json:"sku"`
	ChildSKUs            []SKU                 `json:"child_skus,omitempty"`
	AlternativeSKUs      []SKU                 `json:"alternative_skus,omitempty"`
	Summary              json.RawMessage       `json:"summary"`
	Description          json.RawMessage       `json:"description,omitempty"`
	Tagline              json.RawMessage       `json:"tagline,omitempty"`
	FlavorText           *string               `json:"flavor_text,omitempty"`
	Benefits             []StoreListingBenefit `json:"benefits,omitempty"`
	Published            bool                  `json:"published,omitzero"`
	CarouselItems        []StoreCarouselItem   `json:"carousel_items,omitempty"`
	StaffNotes           StoreNote             `json:"staff_notes,omitzero"`
	Guild                *PartialGuild         `json:"guild,omitempty"`
	Assets               []StoreAsset          `json:"assets,omitempty"`
	Thumbnail            StoreAsset            `json:"thumbnail,omitzero"`
	PreviewVideo         StoreAsset            `json:"preview_video,omitzero"`
	HeaderBackground     StoreAsset            `json:"header_background,omitzero"`
	HeaderLogoDarkTheme  StoreAsset            `json:"header_logo_dark_theme,omitzero"`
	HeaderLogoLightTheme StoreAsset            `json:"header_logo_light_theme,omitzero"`
	BoxArt               StoreAsset            `json:"box_art,omitzero"`
	HeroBackground       StoreAsset            `json:"hero_background,omitzero"`
	HeroVideo            StoreAsset            `json:"hero_video,omitzero"`
	EntitlementBranchID  *snowflake.ID         `json:"entitlement_branch_id,omitempty"`
	PublishedAt          json.RawMessage       `json:"published_at,omitempty"`
	UnpublishedAt        json.RawMessage       `json:"unpublished_at,omitempty"`
	PowerupMetadata      GuildPowerupMetadata  `json:"powerup_metadata,omitzero"`
}

type StoreListingBenefit struct {
	ID          snowflake.ID     `json:"id"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Icon        StoreListingIcon `json:"icon"`
}

type StoreListingIcon struct {
	Type         StoreListingIconType `json:"type"`
	StoreAssetID snowflake.ID         `json:"store_asset_id"`
	Emoji        string               `json:"emoji"`
}

type StoreNote struct {
	User    *PartialUser `json:"user"`
	Content string       `json:"content"`
}

type StorePageGuild struct {
	ID                       snowflake.ID `json:"id"`
	Name                     string       `json:"name"`
	IconHash                 *string      `json:"icon_hash"`
	ApproximateMemberCount   int64        `json:"approximate_member_count"`
	ApproximatePresenceCount int64        `json:"approximate_presence_count"`
	LockedServer             bool         `json:"locked_server"`
	Invite                   *Invite      `json:"invite"`
}

type StorePageRoleSubscription struct {
	Settings           RoleSubscriptionSettings       `json:"settings"`
	GroupListings      []RoleSubscriptionGroupListing `json:"group_listings"`
	Trials             []RoleSubscriptionTrial        `json:"trials"`
	SubscriberCount    *int64                         `json:"subscriber_count"`
	BenefitChannels    []PartialChannel               `json:"benefit_channels"`
	BenefitEmojis      []Emoji                        `json:"benefit_emojis"`
	PurchasePageInvite *Invite                        `json:"purchase_page_invite"`
}

type StoreTenantMetadata struct {
	GuildMonetization GuildMonetizationMetadata `json:"guild_monetization,omitzero"`
	SocialLayer       SocialLayerMetadata       `json:"social_layer,omitzero"`
}

type Storefront struct {
	ApplicationID         snowflake.ID       `json:"application_id"`
	Application           PartialApplication `json:"application,omitzero"`
	Title                 string             `json:"title"`
	LogoAssetID           snowflake.ID       `json:"logo_asset_id,omitzero"`
	LightThemeLogoAssetID snowflake.ID       `json:"light_theme_logo_asset_id,omitzero"`
	Pages                 []StorefrontPage   `json:"pages"`
	StoreListings         []StoreListing     `json:"store_listings"`
	Assets                []StoreAsset       `json:"assets"`
	StorefrontPricing     StorefrontPrices   `json:"storefront_pricing,omitzero"`
}

type StorefrontCollection struct {
	ID             snowflake.ID    `json:"id"`
	ApplicationID  snowflake.ID    `json:"application_id"`
	Name           string          `json:"name"`
	Description    string          `json:"description"`
	ProductIDs     []snowflake.ID  `json:"product_ids"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
	TenantMetadata json.RawMessage `json:"tenant_metadata"`
}

type StorefrontConfig struct {
	GuildID                           snowflake.ID `json:"guild_id"`
	ApplicationID                     snowflake.ID `json:"application_id"`
	GameID                            snowflake.ID `json:"game_id"`
	CollectiblesShopNavigationEnabled bool         `json:"collectibles_shop_navigation_enabled"`
	ExcludedPlatforms                 []string     `json:"excluded_platforms"`
	DisableMobileAccountLinking       bool         `json:"disable_mobile_account_linking"`
}

type StorefrontEligibility struct {
	IsEligible bool `json:"is_eligible"`
}

type StorefrontLeaderboard struct {
	Title                  string       `json:"title,omitzero"`
	Description            string       `json:"description,omitzero"`
	BackgroundImageAssetID snowflake.ID `json:"background_image_asset_id,omitzero"`
}

type StorefrontMetadata struct {
	LogoAssetID           snowflake.ID `json:"logo_asset_id,omitzero"`
	LightThemeLogoAssetID snowflake.ID `json:"light_theme_logo_asset_id,omitzero"`
}

type StorefrontPage struct {
	Title        string                  `json:"title,omitzero"`
	Leaderboard  StorefrontLeaderboard   `json:"leaderboard,omitzero"`
	SKUIDs       []snowflake.ID          `json:"sku_ids"`
	CustomSKUIDs []string                `json:"custom_sku_ids"`
	Sections     []StorefrontPageSection `json:"sections,omitempty"`
}

type StorefrontPageSection struct {
	Title        string         `json:"title,omitzero"`
	SKUIDs       []snowflake.ID `json:"sku_ids"`
	CustomSKUIDs []string       `json:"custom_sku_ids"`
}

type StorefrontPrices struct {
	SKUPriceMap        map[snowflake.ID]StorefrontSKUPricing                `json:"sku_price_map"`
	PricingResultIDMap map[snowflake.ID]map[int64]StorefrontPricingResult   `json:"pricing_result_id_map"`
	RewardResultIDMap  map[snowflake.ID]map[int64]StorefrontPromotionReward `json:"reward_result_id_map"`
}

type StorefrontPricingResult struct {
	UserPrice []UnitPrice                     `json:"user_price"`
	Prices    map[int64]map[int64][]UnitPrice `json:"prices"`
}

type StorefrontProduct struct {
	ID             snowflake.ID          `json:"id"`
	ApplicationID  snowflake.ID          `json:"application_id"`
	SKUIDs         []snowflake.ID        `json:"sku_ids"`
	Name           string                `json:"name"`
	Options        []ProductOption       `json:"options"`
	CreatedAt      time.Time             `json:"created_at"`
	UpdatedAt      time.Time             `json:"updated_at"`
	TenantMetadata ProductTenantMetadata `json:"tenant_metadata"`
	SKUs           []ProductSKU          `json:"skus"`
}

type StorefrontPromotionReward struct {
	Type   StorefrontPromotionRewardType `json:"type"`
	Amount int64                         `json:"amount"`
}

type StorefrontSKUPricing struct {
	PricingResultID        snowflake.ID   `json:"pricing_result_id"`
	StorefrontPromotionIDs []snowflake.ID `json:"storefront_promotion_ids"`
	RewardResultIDs        []snowflake.ID `json:"reward_result_ids"`
}

type Stream struct {
	StreamKey    string          `json:"stream_key"`
	RTCServerID  snowflake.ID    `json:"rtc_server_id"`
	RTCChannelID snowflake.ID    `json:"rtc_channel_id"`
	Region       string          `json:"region"`
	ViewerIDs    json.RawMessage `json:"viewer_ids"`
	Paused       bool            `json:"paused"`
}

type StreamDelete struct {
	StreamKey   string             `json:"stream_key"`
	Reason      StreamDeleteReason `json:"reason"`
	Unavailable bool               `json:"unavailable,omitzero"`
}

type StreamKey struct {
	Type      StreamType   `json:"type"`
	GuildID   snowflake.ID `json:"guild_id,omitzero"`
	ChannelID snowflake.ID `json:"channel_id,omitzero"`
	OwnerID   snowflake.ID `json:"owner_id"`
}

type StreamResolution struct {
	Type   ResolutionType `json:"type"`
	Width  float64        `json:"width"`
	Height float64        `json:"height"`
}

type StreamServerUpdate struct {
	Token     string  `json:"token"`
	StreamKey string  `json:"stream_key"`
	Endpoint  *string `json:"endpoint"`
}

type StringSelect struct {
	Type        ComponentType  `json:"type"`
	ID          int64          `json:"id,omitzero"`
	CustomID    string         `json:"custom_id"`
	Options     []SelectOption `json:"options"`
	Placeholder string         `json:"placeholder,omitzero"`
	MinValues   int64          `json:"min_values,omitzero"`
	MaxValues   int64          `json:"max_values,omitzero"`
	Required    bool           `json:"required,omitzero"`
	Disabled    bool           `json:"disabled,omitzero"`
}

type StudentHubGuild struct {
	ID   snowflake.ID `json:"id"`
	Name string       `json:"name"`
	Icon *string      `json:"icon"`
}

type Subblock struct {
	Type                   SubblockType `json:"type"`
	CategoryStoreListingID snowflake.ID `json:"category_store_listing_id"`
	AssetURL               string       `json:"asset_url"`
	BannerTextColor        *string      `json:"banner_text_color"`
	BannerURL              string       `json:"banner_url"`
	BodyText               *string      `json:"body_text"`
	Name                   string       `json:"name"`
	UnpublishedAt          *time.Time   `json:"unpublished_at"`
}

type SubmitDeveloperPortalCsatSurveyParams struct {
	UserID       snowflake.ID `json:"user_id"`
	CsatResponse int64        `json:"csat_response"`
}

type SubmitReportMenuParams struct {
	Version               string              `json:"version"`
	Variant               string              `json:"variant"`
	Name                  ReportMenuType      `json:"name"`
	Language              string              `json:"language"`
	Breadcrumbs           []int64             `json:"breadcrumbs"`
	Elements              map[string][]string `json:"elements,omitempty"`
	ChannelID             snowflake.ID        `json:"channel_id,omitzero"`
	MessageID             snowflake.ID        `json:"message_id,omitzero"`
	GuildID               snowflake.ID        `json:"guild_id,omitzero"`
	StageInstanceID       snowflake.ID        `json:"stage_instance_id,omitzero"`
	GuildScheduledEventID snowflake.ID        `json:"guild_scheduled_event_id,omitzero"`
	ReportedUserID        snowflake.ID        `json:"reported_user_id,omitzero"`
	ApplicationID         snowflake.ID        `json:"application_id,omitzero"`
	UserID                snowflake.ID        `json:"user_id,omitzero"`
	WidgetID              snowflake.ID        `json:"widget_id,omitzero"`
}

type SubmitReportMenuResponse struct {
	ReportID snowflake.ID `json:"report_id"`
}

type SubmitUnauthenticatedReportMenuParams struct {
	Version     string              `json:"version"`
	Variant     string              `json:"variant"`
	Name        ReportMenuType      `json:"name"`
	Language    string              `json:"language"`
	Breadcrumbs []int64             `json:"breadcrumbs"`
	Elements    map[string][]string `json:"elements,omitempty"`
	EmailToken  string              `json:"email_token"`
}

type Subreddit struct {
	ID          string `json:"id"`
	Subscribers int64  `json:"subscribers"`
	URL         string `json:"url"`
}

type SubscribeResponse struct {
	Evt string `json:"evt"`
}

type Subscription struct {
	ID                           snowflake.ID                 `json:"id"`
	Type                         SubscriptionType             `json:"type"`
	PaymentGateway               *PaymentGateway              `json:"payment_gateway"`
	Currency                     string                       `json:"currency"`
	Items                        []SubscriptionItem           `json:"items"`
	PaymentGatewayPlanID         *string                      `json:"payment_gateway_plan_id"`
	PaymentGatewaySubscriptionID *string                      `json:"payment_gateway_subscription_id,omitempty"`
	CurrentPeriodStart           time.Time                    `json:"current_period_start"`
	CurrentPeriodEnd             time.Time                    `json:"current_period_end"`
	StreakStartedAt              time.Time                    `json:"streak_started_at,omitzero"`
	Status                       SubscriptionStatus           `json:"status"`
	RenewalMutations             SubscriptionRenewalMutations `json:"renewal_mutations,omitzero"`
	TrialID                      snowflake.ID                 `json:"trial_id,omitzero"`
	PaymentSourceID              *snowflake.ID                `json:"payment_source_id"`
	CreatedAt                    time.Time                    `json:"created_at"`
	CanceledAt                   time.Time                    `json:"canceled_at,omitzero"`
	CountryCode                  *string                      `json:"country_code"`
	TrialEndsAt                  time.Time                    `json:"trial_ends_at,omitzero"`
	Metadata                     SubscriptionMetadata         `json:"metadata,omitzero"`
	LatestInvoice                json.RawMessage              `json:"latest_invoice,omitempty"`
	UseStorekitResubscribe       bool                         `json:"use_storekit_resubscribe"`
	Price                        *int64                       `json:"price"`
	Entitlements                 []Entitlement                `json:"entitlements,omitempty"`
}

type SubscriptionItem struct {
	ID       snowflake.ID `json:"id"`
	Quantity int64        `json:"quantity"`
	PlanID   snowflake.ID `json:"plan_id"`
}

type SubscriptionListing struct {
	ID                   snowflake.ID          `json:"id"`
	ApplicationID        *snowflake.ID         `json:"application_id"`
	Published            bool                  `json:"published"`
	SoftDeleted          bool                  `json:"soft_deleted"`
	SKUFlags             SKUFlags              `json:"sku_flags"`
	ImageAsset           StoreAsset            `json:"image_asset,omitzero"`
	SubscriptionPlans    []SubscriptionPlan    `json:"subscription_plans"`
	StoreListingBenefits []StoreListingBenefit `json:"store_listing_benefits"`
}

type SubscriptionMetadata struct {
	IsEgs                          bool         `json:"is_egs,omitzero"`
	IsHolidayPromotion2021         bool         `json:"is_holiday_promotion_2021,omitzero"`
	EndedAt                        string       `json:"ended_at,omitzero"`
	GuildID                        snowflake.ID `json:"guild_id,omitzero"`
	ApplicationSubscriptionGuildID snowflake.ID `json:"application_subscription_guild_id,omitzero"`
	GracePeriodExpiresDate         time.Time    `json:"grace_period_expires_date,omitzero"`
	AppleGracePeriodExpiresDate    time.Time    `json:"apple_grace_period_expires_date,omitzero"`
	GoogleGracePeriodExpiresDate   time.Time    `json:"google_grace_period_expires_date,omitzero"`
	GoogleOriginalExpiresDate      string       `json:"google_original_expires_date,omitzero"`
	UserTrialOfferID               snowflake.ID `json:"user_trial_offer_id,omitzero"`
	UserDiscountOfferID            snowflake.ID `json:"user_discount_offer_id,omitzero"`
	ActiveDiscountID               snowflake.ID `json:"active_discount_id,omitzero"`
	ActiveDiscountExpiresAt        time.Time    `json:"active_discount_expires_at,omitzero"`
}

type SubscriptionMetadataRequest struct {
	GuildID snowflake.ID `json:"guild_id,omitzero"`
}

type SubscriptionPlan struct {
	ID            snowflake.ID                 `json:"id"`
	Name          string                       `json:"name"`
	SKUID         snowflake.ID                 `json:"sku_id"`
	Interval      SubscriptionInterval         `json:"interval"`
	IntervalCount int64                        `json:"interval_count"`
	TaxInclusive  bool                         `json:"tax_inclusive"`
	Price         int64                        `json:"price"`
	Currency      string                       `json:"currency"`
	Prices        map[int64]SubscriptionPrices `json:"prices"`
}

type SubscriptionPrices struct {
	CountryPrices       CountryPrices                `json:"country_prices"`
	PaymentSourcePrices map[snowflake.ID][]UnitPrice `json:"payment_source_prices"`
}

type SubscriptionRenewalMutations struct {
	PaymentGatewayPlanID *string            `json:"payment_gateway_plan_id,omitempty"`
	Items                []SubscriptionItem `json:"items,omitempty"`
}

type SubscriptionTrial struct {
	ID            snowflake.ID         `json:"id"`
	Interval      SubscriptionInterval `json:"interval"`
	IntervalCount int64                `json:"interval_count"`
	SKUID         snowflake.ID         `json:"sku_id"`
}

type SummarizeThreadParams struct {
	Ephemeral bool `json:"ephemeral,omitzero"`
}

type SupplementalGuild struct {
	ID                snowflake.ID               `json:"id"`
	VoiceStates       []VoiceState               `json:"voice_states"`
	ActivityInstances []EmbeddedActivityInstance `json:"activity_instances"`
}

type SupplementalGuildMember struct {
	UserID                  snowflake.ID     `json:"user_id,omitzero"`
	Member                  GuildMember      `json:"member,omitzero"`
	JoinSourceType          JoinSourceType   `json:"join_source_type,omitzero"`
	SourceInviteCode        *string          `json:"source_invite_code,omitempty"`
	InviterID               *snowflake.ID    `json:"inviter_id,omitempty"`
	IntegrationType         *IntegrationType `json:"integration_type,omitempty"`
	JoinSourceApplicationID *snowflake.ID    `json:"join_source_application_id,omitempty"`
	JoinSourceChannelID     *snowflake.ID    `json:"join_source_channel_id,omitempty"`
}

type SupplementalMessageRequest struct {
	ChannelID      snowflake.ID `json:"channel_id"`
	MessagePreview Message      `json:"message_preview"`
}

type SyncDevicesParams struct {
	Provider       PushNotificationProvider `json:"provider"`
	Token          string                   `json:"token"`
	PushSyncTokens []string                 `json:"push_sync_tokens"`
}

type SyncDevicesResponse struct {
	InvalidPushSyncTokens []string `json:"invalid_push_sync_tokens"`
}

type SystemRequirement struct {
	Ram                    int64           `json:"ram,omitzero"`
	Disk                   int64           `json:"disk,omitzero"`
	OperatingSystemVersion json.RawMessage `json:"operating_system_version,omitempty"`
	Cpu                    json.RawMessage `json:"cpu,omitempty"`
	Gpu                    json.RawMessage `json:"gpu,omitempty"`
	SoundCard              json.RawMessage `json:"sound_card,omitempty"`
	Directx                json.RawMessage `json:"directx,omitempty"`
	Network                json.RawMessage `json:"network,omitempty"`
	Notes                  json.RawMessage `json:"notes,omitempty"`
}

type SystemRequirements struct {
	Minimum     SystemRequirement `json:"minimum,omitzero"`
	Recommended SystemRequirement `json:"recommended,omitzero"`
}

type TabTooltipCollectiblesMarketing struct {
	Type                CollectiblesMarketingType `json:"type"`
	Title               string                    `json:"title"`
	Body                string                    `json:"body,omitzero"`
	Asset               string                    `json:"asset"`
	DismissibleContent  int64                     `json:"dismissible_content"`
	Version             int64                     `json:"version"`
	RefTargetBackground TargetBackground          `json:"ref_target_background,omitzero"`
	BadgeIcon           string                    `json:"badge_icon,omitzero"`
	BadgeText           string                    `json:"badge_text,omitzero"`
	ShowHoverGradient   bool                      `json:"show_hover_gradient,omitzero"`
}

type TargetBackground struct {
	Style TargetBackgroundInformation `json:"style"`
	Asset TargetBackgroundInformation `json:"asset"`
}

type TargetBackgroundInformation struct {
	Resting TargetBackgroundState `json:"resting"`
	Hovered TargetBackgroundState `json:"hovered"`
}

type TargetBackgroundState struct {
	Light *string `json:"light"`
	Dark  *string `json:"dark"`
}

type Team struct {
	ID                     snowflake.ID             `json:"id"`
	Name                   string                   `json:"name"`
	Icon                   *string                  `json:"icon"`
	OwnerUserID            snowflake.ID             `json:"owner_user_id"`
	Members                []TeamMember             `json:"members,omitempty"`
	PayoutAccountStatus    *TeamPayoutAccountStatus `json:"payout_account_status,omitempty"`
	PayoutAccountStatuses  []TeamPayoutAccount      `json:"payout_account_statuses,omitempty"`
	StripeConnectAccountID string                   `json:"stripe_connect_account_id,omitzero"`
}

type TeamCompany struct {
	ID   snowflake.ID `json:"id"`
	Name string       `json:"name"`
}

type TeamMember struct {
	User            PartialUser         `json:"user"`
	TeamID          snowflake.ID        `json:"team_id"`
	MembershipState TeamMembershipState `json:"membership_state"`
	Role            string              `json:"role"`
}

type TeamPayout struct {
	ID                              snowflake.ID     `json:"id"`
	UserID                          snowflake.ID     `json:"user_id"`
	Amount                          int64            `json:"amount"`
	Status                          TeamPayoutStatus `json:"status"`
	PeriodStart                     time.Time        `json:"period_start"`
	PeriodEnd                       *time.Time       `json:"period_end"`
	PayoutDate                      *time.Time       `json:"payout_date"`
	LatestTipaltiSubmissionResponse json.RawMessage  `json:"latest_tipalti_submission_response,omitempty"`
}

type TeamPayoutAccount struct {
	Gateway TeamPayoutGateway       `json:"gateway"`
	Status  TeamPayoutAccountStatus `json:"status"`
}

type TeenAuditLog struct {
	TeenUserID   *snowflake.ID   `json:"teen_user_id"`
	RangeStartID *snowflake.ID   `json:"range_start_id"`
	Actions      []Action        `json:"actions"`
	Users        []PartialUser   `json:"users"`
	Guilds       []Guild         `json:"guilds"`
	Totals       map[int64]int64 `json:"totals"`
}

type TemplateRoleSubscriptionListing struct {
	Name            string                                   `json:"name"`
	Description     string                                   `json:"description"`
	PriceTier       int64                                    `json:"price_tier"`
	Image           ImageData                                `json:"image"`
	RoleColor       int64                                    `json:"role_color"`
	Channels        []TemplateRoleSubscriptionListingChannel `json:"channels"`
	AdditionalPerks []RoleSubscriptionBenefit                `json:"additional_perks"`
}

type TemplateRoleSubscriptionListingChannel struct {
	ID          snowflake.ID `json:"id"`
	Type        ChannelType  `json:"type"`
	Name        string       `json:"name"`
	Tagline     string       `json:"tagline"`
	Description string       `json:"description"`
	Topic       string       `json:"topic"`
	EmojiName   string       `json:"emoji_name"`
}

type TenantMetadata struct {
	QuestRewards QuestRewardsMetadata `json:"quest_rewards"`
}

type TextAndImagesSettings struct {
	DiversitySurrogate                string                  `json:"diversity_surrogate,omitzero"`
	UseRichChatInput                  bool                    `json:"use_rich_chat_input,omitzero"`
	UseThreadSidebar                  bool                    `json:"use_thread_sidebar,omitzero"`
	RenderSpoilers                    SpoilerRenderOptions    `json:"render_spoilers,omitzero"`
	EmojiPickerCollapsedSections      []string                `json:"emoji_picker_collapsed_sections"`
	StickerPickerCollapsedSections    []string                `json:"sticker_picker_collapsed_sections"`
	ViewImageDescriptions             bool                    `json:"view_image_descriptions,omitzero"`
	ShowCommandSuggestions            bool                    `json:"show_command_suggestions,omitzero"`
	InlineAttachmentMedia             bool                    `json:"inline_attachment_media,omitzero"`
	InlineEmbedMedia                  bool                    `json:"inline_embed_media,omitzero"`
	GIFAutoPlay                       bool                    `json:"gif_auto_play,omitzero"`
	RenderEmbeds                      bool                    `json:"render_embeds,omitzero"`
	RenderReactions                   bool                    `json:"render_reactions,omitzero"`
	AnimateEmoji                      bool                    `json:"animate_emoji,omitzero"`
	AnimateStickers                   StickerAnimationOption  `json:"animate_stickers,omitzero"`
	EnableTTSCommand                  bool                    `json:"enable_tts_command,omitzero"`
	MessageDisplayCompact             bool                    `json:"message_display_compact,omitzero"`
	ExplicitContentFilter             ExplicitContentFilter   `json:"explicit_content_filter,omitzero"`
	ViewNSFWGuilds                    bool                    `json:"view_nsfw_guilds,omitzero"`
	ConvertEmoticons                  bool                    `json:"convert_emoticons,omitzero"`
	ExpressionSuggestionsEnabled      bool                    `json:"expression_suggestions_enabled,omitzero"`
	ViewNSFWCommands                  bool                    `json:"view_nsfw_commands,omitzero"`
	UseLegacyChatInput                bool                    `json:"use_legacy_chat_input,omitzero"`
	SoundboardPickerCollapsedSections []string                `json:"soundboard_picker_collapsed_sections"`
	DMSpamFilter                      int64                   `json:"dm_spam_filter,omitzero"`
	DMSpamFilterV2                    json.RawMessage         `json:"dm_spam_filter_v2"`
	IncludeStickersInAutocomplete     bool                    `json:"include_stickers_in_autocomplete,omitzero"`
	ExplicitContentSettings           ExplicitContentSettings `json:"explicit_content_settings,omitzero"`
	KeywordFilterSettings             KeywordFilterSettings   `json:"keyword_filter_settings,omitzero"`
	IncludeSoundmojiInAutocomplete    bool                    `json:"include_soundmoji_in_autocomplete,omitzero"`
	GoreContentSettings               GoreContentSettings     `json:"gore_content_settings,omitzero"`
	DefaultReactionEmoji              DefaultReactionEmoji    `json:"default_reaction_emoji,omitzero"`
	ShowMentionSuggestions            bool                    `json:"show_mention_suggestions,omitzero"`
	SelfHarmContentSettings           SelfHarmContentSettings `json:"self_harm_content_settings,omitzero"`
	IsCrossDMSearchEnabled            bool                    `json:"is_cross_dm_search_enabled,omitzero"`
	SearchProvider                    json.RawMessage         `json:"search_provider"`
	CustomSearchURL                   string                  `json:"custom_search_url,omitzero"`
}

type TextChannelEngagement struct {
	DayPt                    time.Time    `json:"day_pt"`
	ChannelName              string       `json:"channel_name,omitzero"`
	ChannelID                snowflake.ID `json:"channel_id"`
	Participators            int64        `json:"participators"`
	Communicators            int64        `json:"communicators"`
	MessagesSent             int64        `json:"messages_sent"`
	PctParticipatedInChannel float64      `json:"pct_participated_in_channel,omitzero"`
	PctCommunicatedInChannel float64      `json:"pct_communicated_in_channel,omitzero"`
}

type TextDisplay struct {
	Type    ComponentType `json:"type"`
	ID      int64         `json:"id,omitzero"`
	Content string        `json:"content"`
}

type TextInput struct {
	Type        ComponentType  `json:"type"`
	ID          int64          `json:"id,omitzero"`
	CustomID    string         `json:"custom_id"`
	Style       TextInputStyle `json:"style"`
	Label       string         `json:"label"`
	MinLength   int64          `json:"min_length,omitzero"`
	MaxLength   int64          `json:"max_length,omitzero"`
	Required    bool           `json:"required,omitzero"`
	Value       string         `json:"value,omitzero"`
	Placeholder string         `json:"placeholder,omitzero"`
}

type TextLineResourceElement struct {
	Title       string `json:"title"`
	Body        string `json:"body"`
	SMS         string `json:"sms"`
	SMSBody     string `json:"sms_body,omitzero"`
	IsLocalized bool   `json:"is_localized"`
}

type TextResourceElement struct {
	Header      string `json:"header"`
	Body        string `json:"body"`
	IsLocalized bool   `json:"is_localized"`
}

type ThermalStateUpdate struct {
	ThermalState ThermalState `json:"thermal_state"`
}

type ThirdPartyURL struct {
	Distributor DistributorType `json:"distributor"`
	URL         string          `json:"url"`
}

type ThreadCreateExtraFields struct {
	NewlyCreated bool `json:"newly_created,omitzero"`
}

type ThreadListSync struct {
	GuildID    snowflake.ID   `json:"guild_id"`
	ChannelIDs []snowflake.ID `json:"channel_ids,omitempty"`
	Threads    []Channel      `json:"threads"`
	Members    []ThreadMember `json:"members"`
}

type ThreadMember struct {
	ID            snowflake.ID      `json:"id,omitzero"`
	UserID        snowflake.ID      `json:"user_id,omitzero"`
	JoinTimestamp time.Time         `json:"join_timestamp"`
	Flags         ThreadMemberFlags `json:"flags"`
	Muted         bool              `json:"muted,omitzero"`
	MuteConfig    *MuteConfig       `json:"mute_config,omitempty"`
	Member        GuildMember       `json:"member,omitzero"`
}

type ThreadMemberUpdateExtraFields struct {
	GuildID snowflake.ID `json:"guild_id"`
}

type ThreadMembersUpdate struct {
	ID               snowflake.ID    `json:"id"`
	GuildID          snowflake.ID    `json:"guild_id"`
	MemberCount      int64           `json:"member_count"`
	AddedMembers     []ThreadMember  `json:"added_members,omitempty"`
	RemovedMemberIDs json.RawMessage `json:"removed_member_ids,omitempty"`
}

type ThreadMetadata struct {
	Archived            bool       `json:"archived"`
	AutoArchiveDuration int64      `json:"auto_archive_duration"`
	ArchiveTimestamp    time.Time  `json:"archive_timestamp"`
	Locked              bool       `json:"locked"`
	Invitable           bool       `json:"invitable,omitzero"`
	CreateTimestamp     *time.Time `json:"create_timestamp,omitempty"`
}

type ThreadOnlyChannelMessage struct {
	Content         string          `json:"content,omitzero"`
	Embeds          []Embed         `json:"embeds,omitempty"`
	AllowedMentions AllowedMentions `json:"allowed_mentions,omitzero"`
	Components      Components      `json:"components,omitzero"`
	StickerIDs      []snowflake.ID  `json:"sticker_ids,omitempty"`
	Activity        MessageActivity `json:"activity,omitzero"`
	ApplicationID   snowflake.ID    `json:"application_id,omitzero"`
	Flags           MessageFlags    `json:"flags,omitzero"`
	FilesN          json.RawMessage `json:"files[n],omitempty"`
	Attachments     []Attachment    `json:"attachments,omitempty"`
}

type ThreadPostData struct {
	Owner        *GuildMember `json:"owner"`
	FirstMessage *Message     `json:"first_message"`
}

type Thumbnail struct {
	Type        ComponentType   `json:"type"`
	ID          int64           `json:"id,omitzero"`
	Media       json.RawMessage `json:"media"`
	Description *string         `json:"description,omitempty"`
	Spoiler     bool            `json:"spoiler,omitzero"`
}

type ToggleScreenshareArguments struct {
	Pid int64 `json:"pid,omitzero"`
}

type TopEmoji struct {
	EmojiID   snowflake.ID `json:"emoji_id"`
	EmojiRank int64        `json:"emoji_rank"`
}

type TrackSelectedGIFParams struct {
	ID string `json:"id"`
	Q  string `json:"q"`
}

type TransferApplicationParams struct {
	TeamID snowflake.ID `json:"team_id"`
}

type TranslateMessageParams struct {
	Content string `json:"content"`
	Locale  string `json:"locale"`
}

type TranslateMessageResponse struct {
	Content string `json:"content"`
}

type TriggerTypingIndicatorResponse struct {
	MessageSendCooldownMs  float64 `json:"message_send_cooldown_ms,omitzero"`
	ThreadCreateCooldownMs float64 `json:"thread_create_cooldown_ms,omitzero"`
}

type Tutorial struct {
	IndicatorsSuppressed bool     `json:"indicators_suppressed"`
	IndicatorsConfirmed  []string `json:"indicators_confirmed"`
}

type TypingStart struct {
	ChannelID snowflake.ID    `json:"channel_id"`
	GuildID   snowflake.ID    `json:"guild_id,omitzero"`
	UserID    snowflake.ID    `json:"user_id"`
	Timestamp int64           `json:"timestamp"`
	Member    json.RawMessage `json:"member,omitempty"`
}

type UdpPing struct {
	Magic    json.RawMessage `json:"Magic"`
	Sequence json.RawMessage `json:"Sequence"`
}

type UnavailableGuild struct {
	ID            snowflake.ID `json:"id"`
	Unavailable   bool         `json:"unavailable,omitzero"`
	GeoRestricted bool         `json:"geo_restricted,omitzero"`
	Name          string       `json:"name,omitzero"`
	Icon          *string      `json:"icon,omitempty"`
}

type UnavailableResourceResponse struct {
	Message    string  `json:"message"`
	Code       int64   `json:"code"`
	RetryAfter float64 `json:"retry_after,omitzero"`
}

type UnfurlEmbedParams struct {
	URL string `json:"url"`
}

type UnfurlEmbedResponse struct {
	Response      EmbedResponse `json:"response,omitzero"`
	Error         string        `json:"error,omitzero"`
	ContextErrors []string      `json:"context_errors"`
	Embeds        []Embed       `json:"embeds,omitempty"`
}

type UnfurlEmbedsParams struct {
	URLs []string `json:"urls"`
}

type UnfurlEmbedsResponse struct {
	Embeds []Embed `json:"embeds"`
}

type UnfurledMediaItemObject struct {
	ID                  snowflake.ID                  `json:"id,omitzero"`
	URL                 string                        `json:"url"`
	ProxyURL            string                        `json:"proxy_url,omitzero"`
	Height              *int64                        `json:"height,omitempty"`
	Width               *int64                        `json:"width,omitempty"`
	Flags               AttachmentFlags               `json:"flags,omitzero"`
	ContentType         string                        `json:"content_type,omitzero"`
	ContentScanMetadata ContentScanMetadata           `json:"content_scan_metadata,omitzero"`
	PlaceholderVersion  int64                         `json:"placeholder_version,omitzero"`
	Placeholder         string                        `json:"placeholder,omitzero"`
	LoadingState        UnfurledMediaItemLoadingState `json:"loading_state,omitzero"`
	AttachmentID        snowflake.ID                  `json:"attachment_id,omitzero"`
}

type UnitPrice struct {
	Currency string `json:"currency"`
	Amount   int64  `json:"amount"`
	Exponent int64  `json:"exponent"`
}

type UnmergeProvisionalAccountParams struct {
	ClientID          snowflake.ID                       `json:"client_id"`
	ClientSecret      string                             `json:"client_secret,omitzero"`
	ExternalAuthType  ExternalProviderAuthenticationType `json:"external_auth_type"`
	ExternalAuthToken string                             `json:"external_auth_token"`
}

type UnmergeProvisionalAccountWithBotParams struct {
	ExternalUserID string `json:"external_user_id"`
}

type UnregisterDeviceParams struct {
	Provider PushNotificationProvider `json:"provider"`
	Token    string                   `json:"token"`
}

type UnsubscribeResponse struct {
	Evt string `json:"evt"`
}

type UpdateActivitySessionParams struct {
	Token          string          `json:"token,omitzero"`
	ApplicationID  snowflake.ID    `json:"application_id"`
	Duration       int64           `json:"duration,omitzero"`
	ShareActivity  bool            `json:"share_activity,omitzero"`
	Distributor    DistributorType `json:"distributor,omitzero"`
	ExePath        string          `json:"exe_path,omitzero"`
	VoiceChannelID snowflake.ID    `json:"voice_channel_id,omitzero"`
	SessionID      string          `json:"session_id,omitzero"`
	MediaSessionID string          `json:"media_session_id,omitzero"`
	Closed         bool            `json:"closed,omitzero"`
}

type UpdateActivitySessionResponse struct {
	Token string `json:"token"`
}

type UpdateActivitySubscriptionsParams struct {
	Subscriptions []ActivitySubscription `json:"subscriptions"`
}

type UpdateApplicationAssetParams struct {
	UploadedFilename string  `json:"uploaded_filename,omitzero"`
	Visibility       *string `json:"visibility,omitempty"`
}

type UpdateCreatorMonetizationTeamParams struct {
	TeamID snowflake.ID `json:"team_id"`
}

type UpdateCreatorMonetizationTeamResponse struct {
	Application PartialApplication `json:"application"`
}

type UpdateExternalFriendListEntriesParams struct {
	FriendListEntries           []FriendListEntry             `json:"friend_list_entries"`
	Background                  bool                          `json:"background"`
	AllowedInSuggestions        ContactSyncSuggestionsSetting `json:"allowed_in_suggestions"`
	IncludeMutualFriendsCount   bool                          `json:"include_mutual_friends_count"`
	AddReverseFriendSuggestions bool                          `json:"add_reverse_friend_suggestions,omitzero"`
}

type UpdateExternalFriendListEntriesResponse struct {
	BulkAddToken      *string            `json:"bulk_add_token"`
	FriendSuggestions []FriendSuggestion `json:"friend_suggestions"`
}

type UpdateInviteTargetUsersParams struct {
	TargetUsersFile json.RawMessage `json:"target_users_file"`
}

type UpdateLobbyMember struct {
	ID           snowflake.ID      `json:"id"`
	Metadata     map[string]string `json:"metadata,omitempty"`
	Flags        LobbyMemberFlags  `json:"flags,omitzero"`
	RemoveMember bool              `json:"remove_member,omitzero"`
}

type UpdateLobbyVoiceState struct {
	LobbyID          snowflake.ID `json:"lobby_id"`
	SelfMute         bool         `json:"self_mute"`
	SelfDeaf         bool         `json:"self_deaf"`
	SelfVideo        bool         `json:"self_video,omitzero"`
	PreferredRegion  *string      `json:"preferred_region,omitempty"`
	PreferredRegions []string     `json:"preferred_regions,omitempty"`
}

type UpdateMessageRequestParams struct {
	ConsentStatus ChannelConsentStatus `json:"consent_status"`
}

type UpdatePresence struct {
	Activities []Activity `json:"activities"`
	Status     StatusType `json:"status"`
	Since      int64      `json:"since"`
	AFK        bool       `json:"afk"`
}

type UpdatePresenceParams struct {
	PackageName string             `json:"package_name"`
	Update      PresenceUpdateType `json:"update,omitzero"`
}

type UpdateVoiceState struct {
	GuildID          *snowflake.ID `json:"guild_id"`
	ChannelID        *snowflake.ID `json:"channel_id"`
	SelfMute         bool          `json:"self_mute"`
	SelfDeaf         bool          `json:"self_deaf"`
	SelfVideo        bool          `json:"self_video,omitzero"`
	PreferredRegion  *string       `json:"preferred_region,omitempty"`
	PreferredRegions []string      `json:"preferred_regions,omitempty"`
	Flags            VoiceFlags    `json:"flags,omitzero"`
}

type UploadApplicationAssetParams struct {
	Filename string `json:"filename"`
	FileSize int64  `json:"file_size"`
}

type UploadAttachment struct {
	ID                  *snowflake.ID `json:"id,omitempty"`
	Filename            string        `json:"filename"`
	FileSize            int64         `json:"file_size"`
	IsClip              bool          `json:"is_clip,omitzero"`
	OriginalContentType string        `json:"original_content_type,omitzero"`
}

type UploadStreamPreviewParams struct {
	Thumbnail ImageData `json:"thumbnail"`
}

type UploadUnverifiedApplicationIconParams struct {
	ApplicationName string    `json:"application_name"`
	ApplicationHash string    `json:"application_hash"`
	Icon            ImageData `json:"icon"`
}

type UploadVideoStreamPreviewParams struct {
	File json.RawMessage `json:"file"`
}

type UseGuildTemplateParams struct {
	Name string    `json:"name"`
	Icon ImageData `json:"icon,omitzero"`
}

type User struct {
	ID                    snowflake.ID          `json:"id"`
	Username              string                `json:"username"`
	Discriminator         string                `json:"discriminator"`
	GlobalName            *string               `json:"global_name"`
	Avatar                *string               `json:"avatar"`
	AvatarDecorationData  *AvatarDecorationData `json:"avatar_decoration_data"`
	Collectibles          *Collectibles         `json:"collectibles,omitempty"`
	DisplayNameStyles     *DisplayNameStyle     `json:"display_name_styles,omitempty"`
	PrimaryGuild          *PrimaryGuild         `json:"primary_guild,omitempty"`
	LinkedUsers           []LinkedUser          `json:"linked_users"`
	Bot                   bool                  `json:"bot,omitzero"`
	System                bool                  `json:"system,omitzero"`
	MFAEnabled            bool                  `json:"mfa_enabled"`
	NSFWAllowed           *bool                 `json:"nsfw_allowed,omitempty"`
	AgeVerificationStatus AgeVerificationStatus `json:"age_verification_status"`
	Pronouns              string                `json:"pronouns,omitzero"`
	Bio                   string                `json:"bio"`
	Banner                *string               `json:"banner"`
	AccentColor           *int64                `json:"accent_color"`
	Locale                string                `json:"locale,omitzero"`
	Verified              bool                  `json:"verified"`
	Email                 *string               `json:"email"`
	Phone                 *string               `json:"phone,omitempty"`
	Premium               bool                  `json:"premium"`
	PremiumType           PremiumType           `json:"premium_type"`
	PremiumState          PremiumState          `json:"premium_state,omitzero"`
	PersonalConnectionID  snowflake.ID          `json:"personal_connection_id,omitzero"`
	Flags                 UserFlags             `json:"flags"`
	PublicFlags           UserFlags             `json:"public_flags,omitzero"`
	PurchasedFlags        PurchasedFlags        `json:"purchased_flags,omitzero"`
	PremiumUsageFlags     PremiumUsageFlags     `json:"premium_usage_flags,omitzero"`
	Desktop               bool                  `json:"desktop,omitzero"`
	Mobile                bool                  `json:"mobile,omitzero"`
	HasBouncedEmail       bool                  `json:"has_bounced_email,omitzero"`
	AuthenticatorTypes    []int64               `json:"authenticator_types,omitempty"`
	AnalyticsToken        string                `json:"analytics_token"`
}

type UserAffinity struct {
	UserID   snowflake.ID `json:"user_id"`
	Affinity float64      `json:"affinity"`
}

type UserAffinityV2 struct {
	OtherUserID              snowflake.ID    `json:"other_user_id"`
	UserSegment              UserSegmentType `json:"user_segment"`
	OtherUserSegment         UserSegmentType `json:"other_user_segment"`
	IsFriend                 bool            `json:"is_friend"`
	DMProbability            float64         `json:"dm_probability"`
	DMRank                   int64           `json:"dm_rank"`
	VcProbability            float64         `json:"vc_probability"`
	VcRank                   int64           `json:"vc_rank"`
	ServerMessageProbability float64         `json:"server_message_probability"`
	ServerMessageRank        int64           `json:"server_message_rank"`
	CommunicationProbability float64         `json:"communication_probability"`
	CommunicationRank        int64           `json:"communication_rank"`
}

type UserApplicationActivityStatistics struct {
	ApplicationID           string     `json:"application_id"`
	LastPlayedAt            time.Time  `json:"last_played_at"`
	FirstPlayedAt           *time.Time `json:"first_played_at"`
	TotalDuration           int64      `json:"total_duration"`
	TotalDiscordSKUDuration int64      `json:"total_discord_sku_duration"`
}

type UserApplicationIdentity struct {
	ApplicationID        snowflake.ID                    `json:"application_id"`
	ProviderIssuedUserID string                          `json:"provider_issued_user_id"`
	Profile              PartialUserApplicationProfile   `json:"profile,omitzero"`
	Profiles             []PartialUserApplicationProfile `json:"profiles,omitempty"`
}

type UserApplicationIdentityRemove struct {
	UserID        snowflake.ID `json:"user_id"`
	ApplicationID snowflake.ID `json:"application_id"`
}

type UserApplicationProfile struct {
	Username    *string                          `json:"username"`
	Metadata    *string                          `json:"metadata"`
	Data        *UserApplicationProfileData      `json:"data,omitempty"`
	DataTrusted *bool                            `json:"data_trusted,omitempty"`
	ExternalID  UserApplicationProfileExternalID `json:"external_id"`
	AvatarHash  *string                          `json:"avatar_hash"`
}

type UserApplicationProfileData struct {
	Primary UserApplicationProfilePrimaryData `json:"primary,omitzero"`
	Dynamic []json.RawMessage                 `json:"dynamic,omitempty"`
}

type UserApplicationProfileExternalID struct {
	ProviderType         ApplicationIdentityProviderType `json:"provider_type"`
	ProviderIssuedUserID string                          `json:"provider_issued_user_id"`
	ProviderID           *string                         `json:"provider_id"`
	PreferredGlobalName  *string                         `json:"preferred_global_name"`
}

type UserApplicationProfilePrimaryData struct {
	Season                       string          `json:"season,omitzero"`
	RankName                     string          `json:"rank_name,omitzero"`
	HighestRank                  string          `json:"highest_rank,omitzero"`
	FeaturedPlayedCharacter      string          `json:"featured_played_character,omitzero"`
	FeaturedPlayedCharacterImage json.RawMessage `json:"featured_played_character_image,omitempty"`
	PlaytimeHours                float64         `json:"playtime_hours,omitzero"`
	TotalWins                    int64           `json:"total_wins,omitzero"`
	CurrentPeriodWins            int64           `json:"current_period_wins,omitzero"`
	TotalGames                   int64           `json:"total_games,omitzero"`
	CurrentPeriodGames           int64           `json:"current_period_games,omitzero"`
	TotalKills                   int64           `json:"total_kills,omitzero"`
	CurrentPeriodKills           int64           `json:"current_period_kills,omitzero"`
	TotalAssists                 int64           `json:"total_assists,omitzero"`
	CurrentPeriodAssists         int64           `json:"current_period_assists,omitzero"`
	TotalDeaths                  int64           `json:"total_deaths,omitzero"`
	CurrentPeriodDeaths          int64           `json:"current_period_deaths,omitzero"`
	ServerName                   string          `json:"server_name,omitzero"`
	UserID                       string          `json:"user_id,omitzero"`
	UnionLevel                   string          `json:"union_level,omitzero"`
	TotalResonators              int64           `json:"total_resonators,omitzero"`
	TotalAchievements            int64           `json:"total_achievements,omitzero"`
	TotalEchoes                  int64           `json:"total_echoes,omitzero"`
	LoginDays                    int64           `json:"login_days,omitzero"`
	DataBankLevel                string          `json:"data_bank_level,omitzero"`
}

type UserApplicationRemove struct {
	ApplicationID snowflake.ID `json:"application_id"`
}

type UserConnectionsUpdate struct {
	UserID snowflake.ID `json:"user_id"`
}

type UserContentSettings struct {
	DismissedContents                       json.RawMessage                             `json:"dismissed_contents"`
	LastDismissedOutboundPromotionStartDate string                                      `json:"last_dismissed_outbound_promotion_start_date,omitzero"`
	PremiumTier0ModalDismissedAt            json.RawMessage                             `json:"premium_tier_0_modal_dismissed_at,omitempty"`
	GuildOnboardingUpsellDismissedAt        json.RawMessage                             `json:"guild_onboarding_upsell_dismissed_at,omitempty"`
	SafetyUserSentimentNoticeDismissedAt    json.RawMessage                             `json:"safety_user_sentiment_notice_dismissed_at,omitempty"`
	LastReceivedChangelogID                 json.RawMessage                             `json:"last_received_changelog_id"`
	RecurringDismissibleContentStates       map[string]RecurringDismissibleContentState `json:"recurring_dismissible_content_states"`
	LastGiftIntentDismissedAtMs             json.RawMessage                             `json:"last_gift_intent_dismissed_at_ms"`
}

type UserDiscountOffer struct {
	ID         snowflake.ID  `json:"id"`
	DiscountID snowflake.ID  `json:"discount_id"`
	UserID     snowflake.ID  `json:"user_id"`
	InvoiceID  *snowflake.ID `json:"invoice_id,omitempty"`
	CreatedAt  *time.Time    `json:"created_at,omitempty"`
	AppliedAt  *time.Time    `json:"applied_at"`
	DeletedAt  *time.Time    `json:"deleted_at,omitempty"`
	ExpiresAt  *time.Time    `json:"expires_at"`
	Discount   Discount      `json:"discount"`
}

type UserExperiment struct {
	Hash             int64   `json:"hash"`
	Revision         int64   `json:"revision"`
	Bucket           int64   `json:"bucket"`
	Override         int64   `json:"override"`
	Population       int64   `json:"population"`
	HashResult       int64   `json:"hash_result"`
	AaMode           int64   `json:"aa_mode"`
	TriggerDebugging int64   `json:"trigger_debugging"`
	HoldoutName      *string `json:"holdout_name"`
	HoldoutRevision  *int64  `json:"holdout_revision"`
	HoldoutBucket    *int64  `json:"holdout_bucket"`
}

type UserGuild struct {
	ID                       snowflake.ID `json:"id"`
	Name                     string       `json:"name"`
	Icon                     *string      `json:"icon"`
	Banner                   *string      `json:"banner"`
	Owner                    bool         `json:"owner"`
	Features                 []string     `json:"features"`
	Permissions              Permissions  `json:"permissions"`
	ApproximateMemberCount   int64        `json:"approximate_member_count,omitzero"`
	ApproximatePresenceCount int64        `json:"approximate_presence_count,omitzero"`
}

type UserGuildSettings struct {
	ChannelOverrides     []ChannelOverride        `json:"channel_overrides"`
	Flags                UserGuildSettingsFlags   `json:"flags"`
	GuildID              *snowflake.ID            `json:"guild_id"`
	HideMutedChannels    bool                     `json:"hide_muted_channels"`
	MessageNotifications MessageNotificationLevel `json:"message_notifications"`
	MobilePush           bool                     `json:"mobile_push"`
	MuteScheduledEvents  bool                     `json:"mute_scheduled_events"`
	Muted                bool                     `json:"muted"`
	MuteConfig           *MuteConfig              `json:"mute_config"`
	NotifyHighlights     HighlightLevel           `json:"notify_highlights"`
	SuppressEveryone     bool                     `json:"suppress_everyone"`
	SuppressRoles        bool                     `json:"suppress_roles"`
	Version              int64                    `json:"version"`
}

type UserIdentityVerification struct {
	ID          snowflake.ID `json:"id"`
	Status      float64      `json:"status"`
	LastError   *float64     `json:"last_error"`
	RedirectURL string       `json:"redirect_url"`
}

type UserMergeOperationCompleted struct {
	MergeOperationID snowflake.ID `json:"merge_operation_id"`
	SourceUserID     snowflake.ID `json:"source_user_id"`
}

type UserMessageSummary struct {
	UserID        snowflake.ID `json:"user_id"`
	LastMessageID snowflake.ID `json:"last_message_id"`
}

type UserNonChannelAck struct {
	AckType    ReadStateType `json:"ack_type"`
	ResourceID snowflake.ID  `json:"resource_id"`
	EntityID   snowflake.ID  `json:"entity_id"`
	Version    int64         `json:"version"`
}

type UserNoteUpdate struct {
	ID   snowflake.ID `json:"id"`
	Note string       `json:"note"`
}

type UserPayload struct {
	ID            snowflake.ID `json:"id"`
	Discriminator string       `json:"discriminator"`
	Avatar        string       `json:"avatar"`
	Username      string       `json:"username"`
}

type UserPaymentBrowserCheckoutDone struct {
	LoadID                *string       `json:"load_id"`
	SKUID                 *snowflake.ID `json:"sku_id"`
	SKUSubscriptionPlanID *snowflake.ID `json:"sku_subscription_plan_id"`
}

type UserPaymentClientAdd struct {
	PurchaseTokenHash string    `json:"purchase_token_hash"`
	ExpiresAt         time.Time `json:"expires_at"`
}

type UserRequiredActionUpdate struct {
	RequiredAction *RequiredActionType `json:"required_action"`
}

type UserSelect struct {
	Type          ComponentType     `json:"type"`
	ID            int64             `json:"id,omitzero"`
	CustomID      string            `json:"custom_id"`
	Placeholder   string            `json:"placeholder,omitzero"`
	DefaultValues []json.RawMessage `json:"default_values,omitempty"`
	MinValues     int64             `json:"min_values,omitzero"`
	MaxValues     int64             `json:"max_values,omitzero"`
	Required      bool              `json:"required,omitzero"`
	Disabled      bool              `json:"disabled,omitzero"`
}

type UserSettings struct {
	ActivityRestrictedGuildIDs            []snowflake.ID            `json:"activity_restricted_guild_ids"`
	ActivityJoiningRestrictedGuildIDs     []snowflake.ID            `json:"activity_joining_restricted_guild_ids"`
	AFKTimeout                            int64                     `json:"afk_timeout"`
	AllowAccessibilityDetection           bool                      `json:"allow_accessibility_detection"`
	AllowActivityPartyPrivacyFriends      bool                      `json:"allow_activity_party_privacy_friends"`
	AllowActivityPartyPrivacyVoiceChannel bool                      `json:"allow_activity_party_privacy_voice_channel"`
	AnimateEmoji                          bool                      `json:"animate_emoji"`
	AnimateStickers                       StickerAnimationOption    `json:"animate_stickers"`
	ContactSyncEnabled                    bool                      `json:"contact_sync_enabled"`
	ConvertEmoticons                      bool                      `json:"convert_emoticons"`
	CustomStatus                          *CustomStatus             `json:"custom_status"`
	DefaultGuildsRestricted               bool                      `json:"default_guilds_restricted"`
	DetectPlatformAccounts                bool                      `json:"detect_platform_accounts"`
	DeveloperMode                         bool                      `json:"developer_mode"`
	DisableGamesTab                       bool                      `json:"disable_games_tab"`
	EnableTTSCommand                      bool                      `json:"enable_tts_command"`
	ExplicitContentFilter                 ExplicitContentFilter     `json:"explicit_content_filter"`
	FriendDiscoveryFlags                  FriendDiscoveryFlags      `json:"friend_discovery_flags"`
	FriendSourceFlags                     *FriendSourceFlags        `json:"friend_source_flags"`
	GIFAutoPlay                           bool                      `json:"gif_auto_play"`
	GuildFolders                          []GuildFolder             `json:"guild_folders"`
	InlineAttachmentMedia                 bool                      `json:"inline_attachment_media"`
	InlineEmbedMedia                      bool                      `json:"inline_embed_media"`
	Locale                                string                    `json:"locale"`
	MessageDisplayCompact                 bool                      `json:"message_display_compact"`
	NativePhoneIntegrationEnabled         bool                      `json:"native_phone_integration_enabled"`
	Passwordless                          bool                      `json:"passwordless"`
	RenderEmbeds                          bool                      `json:"render_embeds"`
	RenderReactions                       bool                      `json:"render_reactions"`
	RestrictedGuilds                      []snowflake.ID            `json:"restricted_guilds"`
	ShowCurrentGame                       bool                      `json:"show_current_game"`
	SlayerSdkReceiveDMsInGame             SlayerSdkReceiveInGameDMs `json:"slayer_sdk_receive_dms_in_game"`
	SoundboardVolume                      float64                   `json:"soundboard_volume"`
	Status                                StatusType                `json:"status"`
	StreamNotificationsEnabled            bool                      `json:"stream_notifications_enabled"`
	Theme                                 Theme                     `json:"theme"`
	TimezoneOffset                        int64                     `json:"timezone_offset"`
	ViewNSFWCommands                      bool                      `json:"view_nsfw_commands"`
	ViewNSFWGuilds                        bool                      `json:"view_nsfw_guilds"`
}

type UserSettingsGetLocaleResponse struct {
	Locale string `json:"locale"`
}

type UserSettingsProto struct {
	Type  UserSettingsProtoType `json:"type"`
	Proto string                `json:"proto"`
}

type UserSettingsProtoApplicationDMSettings struct {
	DMDisabled      bool `json:"dm_disabled"`
	AllowMobilePush bool `json:"allow_mobile_push"`
}

type UserSettingsProtoAudioContext struct {
	Muted           bool            `json:"muted"`
	Volume          float64         `json:"volume"`
	ModifiedAt      json.RawMessage `json:"modified_at"`
	SoundboardMuted bool            `json:"soundboard_muted"`
}

type UserSettingsProtoCustomStatus struct {
	Text        string          `json:"text"`
	EmojiID     json.RawMessage `json:"emoji_id"`
	EmojiName   string          `json:"emoji_name"`
	ExpiresAtMs json.RawMessage `json:"expires_at_ms"`
	CreatedAtMs json.RawMessage `json:"created_at_ms"`
	Label       string          `json:"label"`
}

type UserSettingsProtoGuildFolder struct {
	GuildIDs []json.RawMessage `json:"guild_ids"`
	ID       int64             `json:"id,omitzero"`
	Name     string            `json:"name,omitzero"`
	Color    uint64            `json:"color,omitzero"`
}

type UserSettingsProtoNotificationSettings struct {
	ShowInAppNotifications                    bool            `json:"show_in_app_notifications,omitzero"`
	NotifyFriendsOnGoLive                     bool            `json:"notify_friends_on_go_live,omitzero"`
	NotificationCenterAckedBeforeID           json.RawMessage `json:"notification_center_acked_before_id"`
	EnableBurstReactionNotifications          bool            `json:"enable_burst_reaction_notifications,omitzero"`
	QuietMode                                 bool            `json:"quiet_mode,omitzero"`
	FocusModeExpiresAtMs                      json.RawMessage `json:"focus_mode_expires_at_ms"`
	ReactionNotifications                     json.RawMessage `json:"reaction_notifications"`
	GameActivityNotifications                 json.RawMessage `json:"game_activity_notifications"`
	CustomStatusPushNotifications             json.RawMessage `json:"custom_status_push_notifications"`
	GameActivityExcludeSteamNotifications     bool            `json:"game_activity_exclude_steam_notifications,omitzero"`
	EnableVoiceActivityNotifications          bool            `json:"enable_voice_activity_notifications,omitzero"`
	EnableFriendOnlineNotifications           bool            `json:"enable_friend_online_notifications,omitzero"`
	EnableUserResurrectionNotifications       bool            `json:"enable_user_resurrection_notifications,omitzero"`
	EnableFriendAnniversaryNotifications      bool            `json:"enable_friend_anniversary_notifications,omitzero"`
	EnableGameUpdateNotifications             bool            `json:"enable_game_update_notifications,omitzero"`
	EnableProfileUpdatesNotifications         bool            `json:"enable_profile_updates_notifications,omitzero"`
	EnableServerTrendingNotifications         bool            `json:"enable_server_trending_notifications,omitzero"`
	EnableDMReplyNudgeReminders               bool            `json:"enable_dm_reply_nudge_reminders,omitzero"`
	EnableSummaryReminderNotifications        bool            `json:"enable_summary_reminder_notifications,omitzero"`
	EnableGdmAllReactionNotifications         bool            `json:"enable_gdm_all_reaction_notifications,omitzero"`
	EnableFriendGamingActivityNotifications   bool            `json:"enable_friend_gaming_activity_notifications,omitzero"`
	EnableUpcomingServerEventNotifications    bool            `json:"enable_upcoming_server_event_notifications,omitzero"`
	EnableScreenDowntimeScheduleNotifications bool            `json:"enable_screen_downtime_schedule_notifications,omitzero"`
	NotifyFriendsOnProfileUpdate              bool            `json:"notify_friends_on_profile_update,omitzero"`
	NotifyFriendsOnComeOnline                 bool            `json:"notify_friends_on_come_online,omitzero"`
}

type UserSettingsProtoUpdate struct {
	Settings UserSettingsProto `json:"settings"`
	Partial  bool              `json:"partial"`
}

type UserSettingsProtoVideoFilterAsset struct {
	ID        json.RawMessage `json:"id"`
	AssetHash string          `json:"asset_hash"`
}

type UserSurvey struct {
	ID                snowflake.ID `json:"id"`
	Key               snowflake.ID `json:"key"`
	Prompt            string       `json:"prompt"`
	Cta               string       `json:"cta"`
	URL               string       `json:"url"`
	GuildRequirements []string     `json:"guild_requirements"`
	GuildSize         []*int64     `json:"guild_size"`
	GuildPermissions  []string     `json:"guild_permissions"`
}

type UserTrialOffer struct {
	ID                snowflake.ID      `json:"id"`
	UserID            snowflake.ID      `json:"user_id"`
	TrialID           snowflake.ID      `json:"trial_id"`
	ExpiresAt         *time.Time        `json:"expires_at"`
	SubscriptionTrial SubscriptionTrial `json:"subscription_trial"`
}

type ValidateBillingAddressResponse struct {
	Token string `json:"token"`
}

type ValidateDiscoverySearchTermQuery struct {
	Term string `json:"term"`
}

type ValidateDiscoverySearchTermResponse struct {
	Valid bool `json:"valid"`
}

type ValidateGuildAutomodRuleParams struct {
	TriggerMetadata json.RawMessage `json:"trigger_metadata"`
}

type ValidateGuildAutomodRuleResponse struct {
	TriggerMetadata json.RawMessage `json:"trigger_metadata"`
}

type VerifyAgeResponse struct {
	VerificationRequestID  string `json:"verification_request_id"`
	VerificationVendorName string `json:"verification_vendor_name"`
	VerificationWebviewURL string `json:"verification_webview_url"`
}

type VerifyEmailAddressParams struct {
	Token string `json:"token"`
}

type VerifyEmailAddressResponse struct {
	UserID string  `json:"user_id"`
	Token  *string `json:"token"`
}

type VerifyEmailChangeCodeParams struct {
	Code string `json:"code"`
}

type VerifyEmailChangeCodeResponse struct {
	Token string `json:"token"`
}

type VerifyMFALoginParams struct {
	Ticket          string       `json:"ticket"`
	LoginInstanceID string       `json:"login_instance_id,omitzero"`
	Code            string       `json:"code"`
	LoginSource     *LoginSource `json:"login_source,omitempty"`
	GiftCodeSKUID   *string      `json:"gift_code_sku_id,omitempty"`
}

type VerifyMFALoginResponse struct {
	Token        string        `json:"token"`
	UserSettings LoginSettings `json:"user_settings"`
}

type VerifyMFAParams struct {
	Ticket  string            `json:"ticket"`
	MFAType AuthenticatorType `json:"mfa_type"`
	Data    string            `json:"data"`
}

type VerifyMFAResponse struct {
	Token string `json:"token"`
}

type VerifyPhoneNumberParams struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}

type VerifyPhoneNumberResponse struct {
	Token string `json:"token"`
}

type VerifyPurchaseRequestParams struct {
	Token string `json:"token"`
}

type VerifyUnauthenticatedReportParams struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Code  string `json:"code"`
}

type VerifyUnauthenticatedReportResponse struct {
	Token string `json:"token"`
}

type VerifyUserCaptchaParams struct {
	CaptchaKey string `json:"captcha_key"`
}

type VerifyVoicePublicKeyParams struct {
	KeyVersion int64           `json:"key_version"`
	PublicKey  json.RawMessage `json:"public_key"`
}

type VerifyVoicePublicKeyResponse struct {
	IsMatch bool `json:"is_match"`
}

type Versioned struct {
	Entries []json.RawMessage `json:"entries"`
	Partial bool              `json:"partial"`
	Version int64             `json:"version"`
}

type Versions struct {
	ClientVersion json.RawMessage `json:"client_version"`
	ServerVersion json.RawMessage `json:"server_version"`
	DataVersion   json.RawMessage `json:"data_version"`
}

type Video struct {
	AudioSSRC int64                    `json:"audio_ssrc"`
	VideoSSRC int64                    `json:"video_ssrc"`
	RtxSSRC   int64                    `json:"rtx_ssrc"`
	Streams   []VoiceConnectionsStream `json:"streams"`
	UserID    snowflake.ID             `json:"user_id"`
}

type VideoFilterAsset struct {
	ID       snowflake.ID    `json:"id"`
	Type     VideoFilterType `json:"type"`
	UserID   snowflake.ID    `json:"user_id"`
	Asset    string          `json:"asset"`
	LastUsed int64           `json:"last_used,omitzero"`
}

type VideoFilterBackgroundBlur struct {
	UseBlur bool `json:"use_blur"`
}

type VideoStateUpdate struct {
	Active bool `json:"active"`
}

type VirtualCurrencyBalanceUpdate struct {
	Balance int64 `json:"balance"`
}

type VoiceAndVideoSettings struct {
	Blur                          VideoFilterBackgroundBlur         `json:"blur"`
	PresetOption                  json.RawMessage                   `json:"preset_option"`
	CustomAsset                   UserSettingsProtoVideoFilterAsset `json:"custom_asset"`
	AlwaysPreviewVideo            bool                              `json:"always_preview_video,omitzero"`
	AFKTimeout                    int64                             `json:"afk_timeout,omitzero"`
	StreamNotificationsEnabled    bool                              `json:"stream_notifications_enabled,omitzero"`
	NativePhoneIntegrationEnabled bool                              `json:"native_phone_integration_enabled,omitzero"`
	SoundboardSettings            SoundboardSettings                `json:"soundboard_settings,omitzero"`
	DisableStreamPreviews         bool                              `json:"disable_stream_previews,omitzero"`
	SoundmojiVolume               float64                           `json:"soundmoji_volume,omitzero"`
}

type VoiceBackendVersion struct {
	Voice     string `json:"voice"`
	RTCWorker string `json:"rtc_worker"`
}

type VoiceChannel struct {
	ChannelID   snowflake.ID   `json:"channel_id"`
	ChannelName string         `json:"channel_name"`
	Users       []snowflake.ID `json:"users"`
	Streams     []VoiceStream  `json:"streams,omitempty"`
}

type VoiceChannelEffectSend struct {
	ChannelID     snowflake.ID                     `json:"channel_id"`
	GuildID       snowflake.ID                     `json:"guild_id"`
	UserID        snowflake.ID                     `json:"user_id"`
	AnimationType *VoiceChannelEffectAnimationType `json:"animation_type"`
	AnimationID   int64                            `json:"animation_id"`
	Emoji         *Emoji                           `json:"emoji"`
	SoundID       snowflake.ID                     `json:"sound_id,omitzero"`
	SoundVolume   float64                          `json:"sound_volume,omitzero"`
}

type VoiceChannelEngagement struct {
	DayPt                    time.Time    `json:"day_pt"`
	ChannelName              string       `json:"channel_name"`
	ChannelID                snowflake.ID `json:"channel_id"`
	Participators            int64        `json:"participators"`
	Communicators            int64        `json:"communicators"`
	MessagesSent             int64        `json:"messages_sent"`
	PctParticipatedInChannel float64      `json:"pct_participated_in_channel,omitzero"`
	PctCommunicatedInChannel float64      `json:"pct_communicated_in_channel,omitzero"`
}

type VoiceChannelSelect struct {
	ChannelID *snowflake.ID `json:"channel_id"`
	GuildID   *snowflake.ID `json:"guild_id,omitempty"`
}

type VoiceChannelStartTimeUpdate struct {
	ID             snowflake.ID `json:"id"`
	GuildID        snowflake.ID `json:"guild_id"`
	VoiceStartTime *int64       `json:"voice_start_time"`
}

type VoiceChannelStatusUpdate struct {
	ID      snowflake.ID `json:"id"`
	GuildID snowflake.ID `json:"guild_id"`
	Status  *string      `json:"status"`
}

type VoiceConnectionPing struct {
	Time  int64 `json:"time"`
	Value int64 `json:"value"`
}

type VoiceConnectionStatus struct {
	State       VoiceConnectionState  `json:"state"`
	Hostname    string                `json:"hostname"`
	Pings       []VoiceConnectionPing `json:"pings"`
	AveragePing int64                 `json:"average_ping,omitzero"`
	LastPing    int64                 `json:"last_ping,omitzero"`
}

type VoiceConnectionsHello struct {
	V                 int64 `json:"v"`
	HeartbeatInterval int64 `json:"heartbeat_interval"`
}

type VoiceConnectionsIdentify struct {
	ServerID               snowflake.ID             `json:"server_id"`
	ChannelID              snowflake.ID             `json:"channel_id"`
	UserID                 snowflake.ID             `json:"user_id"`
	SessionID              string                   `json:"session_id"`
	Token                  string                   `json:"token"`
	Video                  bool                     `json:"video,omitzero"`
	Streams                []VoiceConnectionsStream `json:"streams,omitempty"`
	MaxDaveProtocolVersion int64                    `json:"max_dave_protocol_version,omitzero"`
}

type VoiceConnectionsReady struct {
	SSRC        int64                    `json:"ssrc"`
	IP          string                   `json:"ip"`
	Port        int64                    `json:"port"`
	Modes       []string                 `json:"modes"`
	Experiments []string                 `json:"experiments"`
	Streams     []VoiceConnectionsStream `json:"streams"`
}

type VoiceConnectionsResume struct {
	ServerID  snowflake.ID `json:"server_id"`
	ChannelID snowflake.ID `json:"channel_id"`
	SessionID string       `json:"session_id"`
	Token     string       `json:"token"`
	SeqAck    int64        `json:"seq_ack,omitzero"`
}

type VoiceConnectionsStream struct {
	Type          MediaType        `json:"type"`
	Rid           string           `json:"rid"`
	Quality       int64            `json:"quality,omitzero"`
	Active        bool             `json:"active,omitzero"`
	MaxBitrate    int64            `json:"max_bitrate,omitzero"`
	MaxFramerate  int64            `json:"max_framerate,omitzero"`
	MaxResolution StreamResolution `json:"max_resolution,omitzero"`
	SSRC          int64            `json:"ssrc,omitzero"`
	RtxSSRC       int64            `json:"rtx_ssrc,omitzero"`
}

type VoiceGuild struct {
	GuildID       snowflake.ID   `json:"guild_id"`
	GuildName     string         `json:"guild_name"`
	GuildIcon     *string        `json:"guild_icon"`
	VoiceChannels []VoiceChannel `json:"voice_channels"`
}

type VoiceRegion struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Optimal    bool   `json:"optimal"`
	Deprecated bool   `json:"deprecated"`
	Custom     bool   `json:"custom"`
}

type VoiceServerUpdate struct {
	Token     string        `json:"token"`
	GuildID   *snowflake.ID `json:"guild_id"`
	ChannelID snowflake.ID  `json:"channel_id,omitzero"`
	Endpoint  *string       `json:"endpoint"`
}

type VoiceSettingsUpdate2 struct {
	InputMode    RPCVoiceInputMode        `json:"input_mode"`
	LocalMutes   []snowflake.ID           `json:"local_mutes"`
	LocalVolumes map[snowflake.ID]float64 `json:"local_volumes"`
	SelfMute     bool                     `json:"self_mute"`
	SelfDeaf     bool                     `json:"self_deaf"`
}

type VoiceState struct {
	GuildID                 *snowflake.ID `json:"guild_id,omitempty"`
	ChannelID               *snowflake.ID `json:"channel_id"`
	LobbyID                 snowflake.ID  `json:"lobby_id,omitzero"`
	UserID                  snowflake.ID  `json:"user_id"`
	Member                  GuildMember   `json:"member,omitzero"`
	SessionID               string        `json:"session_id"`
	ConnectedAt             int64         `json:"connected_at,omitzero"`
	Deaf                    bool          `json:"deaf"`
	Mute                    bool          `json:"mute"`
	SelfDeaf                bool          `json:"self_deaf"`
	SelfMute                bool          `json:"self_mute"`
	SelfStream              bool          `json:"self_stream,omitzero"`
	SelfVideo               bool          `json:"self_video"`
	Suppress                bool          `json:"suppress"`
	RequestToSpeakTimestamp *time.Time    `json:"request_to_speak_timestamp"`
	Discoverable            bool          `json:"discoverable,omitzero"`
	UserVolume              float64       `json:"user_volume,omitzero"`
}

type VoiceStateCreateSubscriptionArguments struct {
	ChannelID snowflake.ID `json:"channel_id"`
}

type VoiceStateDeleteSubscriptionArguments struct {
	ChannelID snowflake.ID `json:"channel_id"`
}

type VoiceStateUpdateSubscriptionArguments struct {
	ChannelID snowflake.ID `json:"channel_id"`
}

type VoiceStream struct {
	UserID snowflake.ID `json:"user_id"`
}

type VoipPushNotification struct {
	ID              string                   `json:"id"`
	Type            VoipPushNotificationType `json:"type"`
	ChannelID       snowflake.ID             `json:"channel_id"`
	UserID          snowflake.ID             `json:"user_id"`
	NativePhoneName string                   `json:"native_phone_name"`
}

type WatchStream struct {
	StreamKey string `json:"stream_key"`
}

type Webhook struct {
	ID            snowflake.ID     `json:"id"`
	Type          WebhookTypes     `json:"type"`
	GuildID       *snowflake.ID    `json:"guild_id,omitempty"`
	ChannelID     *snowflake.ID    `json:"channel_id"`
	User          *PartialUser     `json:"user,omitempty"`
	Name          *string          `json:"name"`
	Avatar        *string          `json:"avatar"`
	Token         string           `json:"token,omitzero"`
	ApplicationID *snowflake.ID    `json:"application_id"`
	SourceGuild   IntegrationGuild `json:"source_guild,omitzero"`
	SourceChannel WebhookChannel   `json:"source_channel,omitzero"`
	URL           string           `json:"url,omitzero"`
}

type WebhookChannel struct {
	ID   snowflake.ID `json:"id"`
	Name string       `json:"name"`
}

type WebhooksUpdate struct {
	GuildID   snowflake.ID `json:"guild_id"`
	ChannelID snowflake.ID `json:"channel_id"`
}

type WelcomeScreen struct {
	Description     *string                `json:"description"`
	WelcomeChannels []WelcomeScreenChannel `json:"welcome_channels"`
}

type WelcomeScreenChannel struct {
	ChannelID   snowflake.ID  `json:"channel_id"`
	Description string        `json:"description"`
	EmojiID     *snowflake.ID `json:"emoji_id"`
	EmojiName   *string       `json:"emoji_name"`
}

type WelcomeScreenFunnel struct {
	DayPt                    time.Time    `json:"day_pt"`
	OptionChannelID          snowflake.ID `json:"option_channel_id"`
	OptionSelected           string       `json:"option_selected"`
	UsersViewedWelcomeScreen int64        `json:"users_viewed_welcome_screen"`
	UsersClickedAnyOption    int64        `json:"users_clicked_any_option"`
	UsersClickedOption       int64        `json:"users_clicked_option"`
	UsersSentMessage         int64        `json:"users_sent_message"`
	PctClickedOption         float64      `json:"pct_clicked_option,omitzero"`
	PctSentMessage           float64      `json:"pct_sent_message,omitzero"`
}

type WelcomeScreenUsers struct {
	DayPt                    time.Time `json:"day_pt"`
	UsersViewedWelcomeScreen int64     `json:"users_viewed_welcome_screen"`
}

type WhitelistedUser struct {
	User  PartialUser                `json:"user"`
	State ApplicationMembershipState `json:"state"`
}

type WideBannerShopBlock struct {
	Type                      ShopBlockType `json:"type"`
	CategoryStoreListingID    snowflake.ID  `json:"category_store_listing_id"`
	BannerAsset               AssetConfig   `json:"banner_asset"`
	LogoURL                   string        `json:"logo_url"`
	Title                     string        `json:"title"`
	Body                      string        `json:"body"`
	BannerTextColor           *string       `json:"banner_text_color"`
	DisableCta                bool          `json:"disable_cta"`
	CtaText                   string        `json:"cta_text"`
	CtaRoute                  string        `json:"cta_route"`
	IsDismissible             bool          `json:"is_dismissible"`
	DismissibleContentVersion int64         `json:"dismissible_content_version"`
	WideBannerURL             string        `json:"wide_banner_url"`
	WideBannerAnimatedURL     string        `json:"wide_banner_animated_url,omitzero"`
}

type WidgetConfig struct {
	ApplicationID  snowflake.ID               `json:"application_id"`
	ConfigID       snowflake.ID               `json:"config_id"`
	DisplayName    string                     `json:"display_name"`
	Surfaces       map[string]json.RawMessage `json:"surfaces"`
	Status         WidgetConfigStatus         `json:"status"`
	ResolvedAssets []ApplicationAsset         `json:"resolved_assets,omitempty"`
	PublishedAt    *time.Time                 `json:"published_at"`
	UpdatedAt      time.Time                  `json:"updated_at"`
}

type WidgetLayoutComponent struct {
	DisplayName string                       `json:"display_name"`
	Required    bool                         `json:"required"`
	Fields      map[string]WidgetLayoutField `json:"fields"`
}

type WidgetLayoutDefinition struct {
	Surface     string                           `json:"surface"`
	Components  map[string]WidgetLayoutComponent `json:"components"`
	Key         LayoutDefinitionKey              `json:"key"`
	DisplayName string                           `json:"display_name"`
}

type WidgetLayoutField struct {
	DisplayName              string   `json:"display_name"`
	Required                 bool     `json:"required"`
	AllowedPresentationTypes []string `json:"allowed_presentation_types"`
}

type WidgetSurface struct {
	Layout     string                  `json:"layout"`
	Components map[string]AnyComponent `json:"components"`
}

type WidgetSurfaceComponent struct {
	Fields map[string]json.RawMessage `json:"fields"`
}

type WidgetSurfaceComponentField struct {
	ValueType        WidgetSurfaceValueType `json:"value_type"`
	PresentationType PresentationType       `json:"presentation_type"`
	Value            string                 `json:"value"`
	Fallback         json.RawMessage        `json:"fallback,omitempty"`
}

type WidgetsListBulkApplicationIdentitiesParams struct {
	Username string                     `json:"username,omitzero"`
	Metadata json.RawMessage            `json:"metadata,omitempty"`
	Data     UserApplicationProfileData `json:"data,omitzero"`
}

type IdentifyEvent struct {
	Token          string           `json:"token"`
	Properties     ClientProperties `json:"properties"`
	Compress       bool             `json:"compress,omitzero"`
	LargeThreshold int64            `json:"large_threshold,omitzero"`
	Shard          []int64          `json:"shard,omitempty"`
	Presence       UpdatePresence   `json:"presence,omitzero"`
	Intents        int64            `json:"intents,omitzero"`
	Capabilities   int64            `json:"capabilities,omitzero"`
	ClientState    ClientState      `json:"client_state,omitzero"`
}

func (e *IdentifyEvent) AsIdentify() Identify {
	return Identify{
		Token:          e.Token,
		Properties:     e.Properties,
		Compress:       e.Compress,
		LargeThreshold: e.LargeThreshold,
		Shard:          e.Shard,
		Presence:       e.Presence,
		Intents:        e.Intents,
		Capabilities:   e.Capabilities,
		ClientState:    e.ClientState,
	}
}

type ResumeEvent struct {
	Token     string `json:"token"`
	SessionID string `json:"session_id"`
	Seq       int64  `json:"seq"`
}

func (e *ResumeEvent) AsResume() Resume {
	return Resume{
		Token:     e.Token,
		SessionID: e.SessionID,
		Seq:       e.Seq,
	}
}

type QoSHeartbeatEvent struct {
	Seq *int64     `json:"seq"`
	QoS QoSPayload `json:"qos"`
}

func (e *QoSHeartbeatEvent) AsQoSHeartbeat() QoSHeartbeat {
	return QoSHeartbeat{
		Seq: e.Seq,
		QoS: e.QoS,
	}
}

type UpdatePresenceEvent struct {
	Activities []Activity `json:"activities"`
	Status     StatusType `json:"status"`
	Since      int64      `json:"since"`
	AFK        bool       `json:"afk"`
}

func (e *UpdatePresenceEvent) AsUpdatePresence() UpdatePresence {
	return UpdatePresence{
		Activities: e.Activities,
		Status:     e.Status,
		Since:      e.Since,
		AFK:        e.AFK,
	}
}

type UpdateVoiceStateEvent struct {
	GuildID          *snowflake.ID `json:"guild_id"`
	ChannelID        *snowflake.ID `json:"channel_id"`
	SelfMute         bool          `json:"self_mute"`
	SelfDeaf         bool          `json:"self_deaf"`
	SelfVideo        bool          `json:"self_video,omitzero"`
	PreferredRegion  *string       `json:"preferred_region,omitempty"`
	PreferredRegions []string      `json:"preferred_regions,omitempty"`
	Flags            VoiceFlags    `json:"flags,omitzero"`
}

func (e *UpdateVoiceStateEvent) AsUpdateVoiceState() UpdateVoiceState {
	return UpdateVoiceState{
		GuildID:          e.GuildID,
		ChannelID:        e.ChannelID,
		SelfMute:         e.SelfMute,
		SelfDeaf:         e.SelfDeaf,
		SelfVideo:        e.SelfVideo,
		PreferredRegion:  e.PreferredRegion,
		PreferredRegions: e.PreferredRegions,
		Flags:            e.Flags,
	}
}

type CreateStreamEvent struct {
	Type            StreamType    `json:"type"`
	GuildID         *snowflake.ID `json:"guild_id,omitempty"`
	ChannelID       snowflake.ID  `json:"channel_id"`
	PreferredRegion *string       `json:"preferred_region,omitempty"`
}

func (e *CreateStreamEvent) AsCreateStream() CreateStream {
	return CreateStream{
		Type:            e.Type,
		GuildID:         e.GuildID,
		ChannelID:       e.ChannelID,
		PreferredRegion: e.PreferredRegion,
	}
}

type SetStreamPausedEvent struct {
	StreamKey string `json:"stream_key"`
	Paused    bool   `json:"paused"`
}

func (e *SetStreamPausedEvent) AsSetStreamPaused() SetStreamPaused {
	return SetStreamPaused{
		StreamKey: e.StreamKey,
		Paused:    e.Paused,
	}
}

type DeleteStreamEvent struct {
	StreamKey string `json:"stream_key"`
}

func (e *DeleteStreamEvent) AsDeleteStream() DeleteStream {
	return DeleteStream{
		StreamKey: e.StreamKey,
	}
}

type PingStreamServerEvent struct {
	StreamKey string `json:"stream_key"`
}

func (e *PingStreamServerEvent) AsPingStreamServer() PingStreamServer {
	return PingStreamServer{
		StreamKey: e.StreamKey,
	}
}

type PingLobbyVoiceServerEvent struct {
	LobbyID snowflake.ID `json:"lobby_id"`
}

func (e *PingLobbyVoiceServerEvent) AsPingLobbyVoiceServer() PingLobbyVoiceServer {
	return PingLobbyVoiceServer{
		LobbyID: e.LobbyID,
	}
}

type RequestGuildMembersEvent struct {
	GuildID   json.RawMessage `json:"guild_id"`
	Query     string          `json:"query,omitzero"`
	Limit     int64           `json:"limit,omitzero"`
	Presences bool            `json:"presences,omitzero"`
	UserIDs   json.RawMessage `json:"user_ids,omitempty"`
	Nonce     string          `json:"nonce,omitzero"`
}

func (e *RequestGuildMembersEvent) AsRequestGuildMembers() RequestGuildMembers {
	return RequestGuildMembers{
		GuildID:   e.GuildID,
		Query:     e.Query,
		Limit:     e.Limit,
		Presences: e.Presences,
		UserIDs:   e.UserIDs,
		Nonce:     e.Nonce,
	}
}

type RequestCallConnectEvent struct {
	ChannelID snowflake.ID `json:"channel_id"`
}

func (e *RequestCallConnectEvent) AsRequestCallConnect() RequestCallConnect {
	return RequestCallConnect{
		ChannelID: e.ChannelID,
	}
}

type GetDeletedEntityIDsNotMatchingHashEvent struct {
	GuildID        snowflake.ID `json:"guild_id"`
	ChannelIDsHash string       `json:"channel_ids_hash"`
	RoleIDsHash    string       `json:"role_ids_hash"`
	EmojiIDsHash   string       `json:"emoji_ids_hash"`
	StickerIDsHash string       `json:"sticker_ids_hash"`
}

func (e *GetDeletedEntityIDsNotMatchingHashEvent) AsGetDeletedEntityIDsNotMatchingHash() GetDeletedEntityIDsNotMatchingHash {
	return GetDeletedEntityIDsNotMatchingHash{
		GuildID:        e.GuildID,
		ChannelIDsHash: e.ChannelIDsHash,
		RoleIDsHash:    e.RoleIDsHash,
		EmojiIDsHash:   e.EmojiIDsHash,
		StickerIDsHash: e.StickerIDsHash,
	}
}

type RequestSoundboardSoundsEvent struct {
	GuildIDs []snowflake.ID `json:"guild_ids"`
}

func (e *RequestSoundboardSoundsEvent) AsRequestSoundboardSounds() RequestSoundboardSounds {
	return RequestSoundboardSounds{
		GuildIDs: e.GuildIDs,
	}
}

type CreateSpeedTestEvent struct {
	PreferredRegion *string `json:"preferred_region,omitempty"`
}

func (e *CreateSpeedTestEvent) AsCreateSpeedTest() CreateSpeedTest {
	return CreateSpeedTest{
		PreferredRegion: e.PreferredRegion,
	}
}

type RequestLastMessagesEvent struct {
	GuildID    snowflake.ID   `json:"guild_id"`
	ChannelIDs []snowflake.ID `json:"channel_ids"`
}

func (e *RequestLastMessagesEvent) AsRequestLastMessages() RequestLastMessages {
	return RequestLastMessages{
		GuildID:    e.GuildID,
		ChannelIDs: e.ChannelIDs,
	}
}

type SearchRecentMembersEvent struct {
	GuildID           snowflake.ID  `json:"guild_id"`
	Query             string        `json:"query"`
	ContinuationToken *snowflake.ID `json:"continuation_token"`
	Nonce             string        `json:"nonce,omitzero"`
}

func (e *SearchRecentMembersEvent) AsSearchRecentMembers() SearchRecentMembers {
	return SearchRecentMembers{
		GuildID:           e.GuildID,
		Query:             e.Query,
		ContinuationToken: e.ContinuationToken,
		Nonce:             e.Nonce,
	}
}

type ResyncGuildChannelsEvent struct {
	GuildID              snowflake.ID   `json:"guild_id"`
	ObfuscatedChannelIDs []snowflake.ID `json:"obfuscated_channel_ids"`
}

func (e *ResyncGuildChannelsEvent) AsResyncGuildChannels() ResyncGuildChannels {
	return ResyncGuildChannels{
		GuildID:              e.GuildID,
		ObfuscatedChannelIDs: e.ObfuscatedChannelIDs,
	}
}

type RequestChannelStatusesEvent struct {
	GuildID snowflake.ID `json:"guild_id"`
}

func (e *RequestChannelStatusesEvent) AsRequestChannelStatuses() RequestChannelStatuses {
	return RequestChannelStatuses{
		GuildID: e.GuildID,
	}
}

type RequestChannelMemberCountEvent struct {
	GuildID   snowflake.ID `json:"guild_id"`
	ChannelID snowflake.ID `json:"channel_id"`
}

func (e *RequestChannelMemberCountEvent) AsRequestChannelMemberCount() RequestChannelMemberCount {
	return RequestChannelMemberCount{
		GuildID:   e.GuildID,
		ChannelID: e.ChannelID,
	}
}

type RequestChannelInfoEvent struct {
	GuildID snowflake.ID `json:"guild_id"`
	Fields  []string     `json:"fields"`
}

func (e *RequestChannelInfoEvent) AsRequestChannelInfo() RequestChannelInfo {
	return RequestChannelInfo{
		GuildID: e.GuildID,
		Fields:  e.Fields,
	}
}

type HelloEvent struct {
	Trace             []string `json:"_trace"`
	HeartbeatInterval int64    `json:"heartbeat_interval"`
}

func (e *HelloEvent) AsHello() Hello {
	return Hello{
		Trace:             e.Trace,
		HeartbeatInterval: e.HeartbeatInterval,
	}
}

type InitialGuildEvent struct {
	JoinedAt                 time.Time                  `json:"joined_at"`
	Large                    bool                       `json:"large"`
	Unavailable              bool                       `json:"unavailable,omitzero"`
	GeoRestricted            bool                       `json:"geo_restricted,omitzero"`
	MemberCount              int64                      `json:"member_count"`
	Members                  []GuildMember              `json:"members"`
	Channels                 []Channel                  `json:"channels"`
	Threads                  []Channel                  `json:"threads"`
	Presences                []Presence                 `json:"presences"`
	VoiceStates              []VoiceState               `json:"voice_states"`
	ActivityInstances        []EmbeddedActivityInstance `json:"activity_instances"`
	StageInstances           []StageInstance            `json:"stage_instances"`
	GuildScheduledEvents     []GuildScheduledEvent      `json:"guild_scheduled_events"`
	DataMode                 DataMode                   `json:"data_mode"`
	PartialUpdates           PartialGuildUpdates        `json:"partial_updates,omitzero"`
	ChannelUpdates           []ChannelUnread            `json:"channel_updates,omitempty"`
	UnableToSyncDeletes      bool                       `json:"unable_to_sync_deletes,omitzero"`
	Properties               PartialGuild               `json:"properties"`
	Stickers                 []Sticker                  `json:"stickers"`
	Roles                    []Role                     `json:"roles"`
	Emojis                   []Emoji                    `json:"emojis"`
	SoundboardSounds         []SoundboardSound          `json:"soundboard_sounds"`
	PremiumSubscriptionCount int64                      `json:"premium_subscription_count"`
	ID                       snowflake.ID               `json:"id"`
}

func (e *InitialGuildEvent) AsGatewayGuild() GatewayGuild {
	return GatewayGuild{
		JoinedAt:                 e.JoinedAt,
		Large:                    e.Large,
		Unavailable:              e.Unavailable,
		GeoRestricted:            e.GeoRestricted,
		MemberCount:              e.MemberCount,
		Members:                  e.Members,
		Channels:                 e.Channels,
		Threads:                  e.Threads,
		Presences:                e.Presences,
		VoiceStates:              e.VoiceStates,
		ActivityInstances:        e.ActivityInstances,
		StageInstances:           e.StageInstances,
		GuildScheduledEvents:     e.GuildScheduledEvents,
		DataMode:                 e.DataMode,
		PartialUpdates:           e.PartialUpdates,
		ChannelUpdates:           e.ChannelUpdates,
		UnableToSyncDeletes:      e.UnableToSyncDeletes,
		Properties:               e.Properties,
		Stickers:                 e.Stickers,
		Roles:                    e.Roles,
		Emojis:                   e.Emojis,
		SoundboardSounds:         e.SoundboardSounds,
		PremiumSubscriptionCount: e.PremiumSubscriptionCount,
		ID:                       e.ID,
	}
}

type ReadyEvent struct {
	Trace                            []string                                  `json:"_trace"`
	V                                int64                                     `json:"v"`
	User                             User                                      `json:"user"`
	UserSettings                     UserSettings                              `json:"user_settings"`
	UserSettingsProto                string                                    `json:"user_settings_proto"`
	NotificationSettings             NotificationSettings                      `json:"notification_settings"`
	UserGuildSettings                json.RawMessage                           `json:"user_guild_settings"`
	ReadState                        json.RawMessage                           `json:"read_state"`
	Guilds                           []GatewayGuild                            `json:"guilds"`
	GuildJoinRequests                []json.RawMessage                         `json:"guild_join_requests"`
	Relationships                    []Relationship                            `json:"relationships"`
	GameRelationships                []GameRelationship                        `json:"game_relationships"`
	FriendSuggestionCount            int64                                     `json:"friend_suggestion_count,omitzero"`
	PrivateChannels                  []Channel                                 `json:"private_channels"`
	ConnectedAccounts                []Connection                              `json:"connected_accounts"`
	Notes                            map[snowflake.ID]string                   `json:"notes"`
	Presences                        []Presence                                `json:"presences"`
	MergedPresences                  MergedPresences                           `json:"merged_presences"`
	MergedMembers                    [][]GuildMember                           `json:"merged_members"`
	Users                            []PartialUser                             `json:"users"`
	LinkedUsers                      []LinkedUser                              `json:"linked_users"`
	Application                      GatewayApplication                        `json:"application,omitzero"`
	Scopes                           []string                                  `json:"scopes,omitempty"`
	SessionID                        string                                    `json:"session_id"`
	SessionType                      SessionType                               `json:"session_type"`
	Sessions                         []Session                                 `json:"sessions"`
	StaticClientSessionID            string                                    `json:"static_client_session_id"`
	AuthSessionIDHash                string                                    `json:"auth_session_id_hash"`
	AuthToken                        string                                    `json:"auth_token,omitzero"`
	AnalyticsToken                   string                                    `json:"analytics_token"`
	Auth                             GatewayAuth                               `json:"auth"`
	RequiredAction                   RequiredActionType                        `json:"required_action,omitzero"`
	CountryCode                      string                                    `json:"country_code"`
	GeoOrderedRTCRegions             []string                                  `json:"geo_ordered_rtc_regions"`
	Consents                         json.RawMessage                           `json:"consents"`
	Tutorial                         *Tutorial                                 `json:"tutorial"`
	Shard                            []int64                                   `json:"shard,omitempty"`
	ResumeGatewayURL                 string                                    `json:"resume_gateway_url"`
	APICodeVersion                   int64                                     `json:"api_code_version"`
	Experiments                      []UserExperiment                          `json:"experiments"`
	GuildExperiments                 []GuildExperiment                         `json:"guild_experiments"`
	ApexExperiments                  ApexExperiments                           `json:"apex_experiments,omitzero"`
	ExplicitContentScanVersion       int64                                     `json:"explicit_content_scan_version"`
	PendingPayments                  []Payment                                 `json:"pending_payments,omitempty"`
	AvSfProtocolFloor                int64                                     `json:"av_sf_protocol_floor,omitzero"`
	FeatureFlags                     GatewayFeatureFlags                       `json:"feature_flags,omitzero"`
	Lobbies                          []Lobby                                   `json:"lobbies,omitempty"`
	UserApplicationProfiles          map[snowflake.ID][]UserApplicationProfile `json:"user_application_profiles,omitempty"`
	ConnectionRequestData            ConsoleConnectionRequestData              `json:"connection_request_data,omitzero"`
	AdPersonalizationTogglesDisabled bool                                      `json:"ad_personalization_toggles_disabled"`
	BroadcasterUserIDs               []snowflake.ID                            `json:"broadcaster_user_ids"`
	RegionalFeatureConfig            RegionalFeatureConfig                     `json:"regional_feature_config"`
}

func (e *ReadyEvent) AsReady() Ready {
	return Ready{
		Trace:                            e.Trace,
		V:                                e.V,
		User:                             e.User,
		UserSettings:                     e.UserSettings,
		UserSettingsProto:                e.UserSettingsProto,
		NotificationSettings:             e.NotificationSettings,
		UserGuildSettings:                e.UserGuildSettings,
		ReadState:                        e.ReadState,
		Guilds:                           e.Guilds,
		GuildJoinRequests:                e.GuildJoinRequests,
		Relationships:                    e.Relationships,
		GameRelationships:                e.GameRelationships,
		FriendSuggestionCount:            e.FriendSuggestionCount,
		PrivateChannels:                  e.PrivateChannels,
		ConnectedAccounts:                e.ConnectedAccounts,
		Notes:                            e.Notes,
		Presences:                        e.Presences,
		MergedPresences:                  e.MergedPresences,
		MergedMembers:                    e.MergedMembers,
		Users:                            e.Users,
		LinkedUsers:                      e.LinkedUsers,
		Application:                      e.Application,
		Scopes:                           e.Scopes,
		SessionID:                        e.SessionID,
		SessionType:                      e.SessionType,
		Sessions:                         e.Sessions,
		StaticClientSessionID:            e.StaticClientSessionID,
		AuthSessionIDHash:                e.AuthSessionIDHash,
		AuthToken:                        e.AuthToken,
		AnalyticsToken:                   e.AnalyticsToken,
		Auth:                             e.Auth,
		RequiredAction:                   e.RequiredAction,
		CountryCode:                      e.CountryCode,
		GeoOrderedRTCRegions:             e.GeoOrderedRTCRegions,
		Consents:                         e.Consents,
		Tutorial:                         e.Tutorial,
		Shard:                            e.Shard,
		ResumeGatewayURL:                 e.ResumeGatewayURL,
		APICodeVersion:                   e.APICodeVersion,
		Experiments:                      e.Experiments,
		GuildExperiments:                 e.GuildExperiments,
		ApexExperiments:                  e.ApexExperiments,
		ExplicitContentScanVersion:       e.ExplicitContentScanVersion,
		PendingPayments:                  e.PendingPayments,
		AvSfProtocolFloor:                e.AvSfProtocolFloor,
		FeatureFlags:                     e.FeatureFlags,
		Lobbies:                          e.Lobbies,
		UserApplicationProfiles:          e.UserApplicationProfiles,
		ConnectionRequestData:            e.ConnectionRequestData,
		AdPersonalizationTogglesDisabled: e.AdPersonalizationTogglesDisabled,
		BroadcasterUserIDs:               e.BroadcasterUserIDs,
		RegionalFeatureConfig:            e.RegionalFeatureConfig,
	}
}

type ReadySupplementalEvent struct {
	Guilds              []SupplementalGuild `json:"guilds"`
	MergedMembers       [][]json.RawMessage `json:"merged_members"`
	MergedPresences     MergedPresences     `json:"merged_presences"`
	LazyPrivateChannels []Channel           `json:"lazy_private_channels"`
	Disclose            []string            `json:"disclose"`
	GameInvites         []GameInvite        `json:"game_invites"`
}

func (e *ReadySupplementalEvent) AsReadySupplemental() ReadySupplemental {
	return ReadySupplemental{
		Guilds:              e.Guilds,
		MergedMembers:       e.MergedMembers,
		MergedPresences:     e.MergedPresences,
		LazyPrivateChannels: e.LazyPrivateChannels,
		Disclose:            e.Disclose,
		GameInvites:         e.GameInvites,
	}
}

type ResumedEvent struct {
	Trace []string `json:"_trace"`
}

func (e *ResumedEvent) AsResumed() Resumed {
	return Resumed{
		Trace: e.Trace,
	}
}

type RateLimitedEvent struct {
	Opcode     int64             `json:"opcode"`
	RetryAfter float64           `json:"retry_after"`
	Meta       RateLimitMetadata `json:"meta"`
}

func (e *RateLimitedEvent) AsRateLimited() RateLimited {
	return RateLimited{
		Opcode:     e.Opcode,
		RetryAfter: e.RetryAfter,
		Meta:       e.Meta,
	}
}

type ActivityInviteCreateEvent struct {
	MessageID   snowflake.ID           `json:"message_id,omitzero"`
	ChannelID   snowflake.ID           `json:"channel_id,omitzero"`
	Author      PartialUser            `json:"author,omitzero"`
	Application IntegrationApplication `json:"application,omitzero"`
	Activity    MessageActivity        `json:"activity,omitzero"`
}

func (e *ActivityInviteCreateEvent) AsActivityInviteCreate() ActivityInviteCreate {
	return ActivityInviteCreate{
		MessageID:   e.MessageID,
		ChannelID:   e.ChannelID,
		Author:      e.Author,
		Application: e.Application,
		Activity:    e.Activity,
	}
}

type AuthSessionChangeEvent struct {
	AuthSessionIDHash string `json:"auth_session_id_hash"`
}

func (e *AuthSessionChangeEvent) AsAuthSessionChange() AuthSessionChange {
	return AuthSessionChange{
		AuthSessionIDHash: e.AuthSessionIDHash,
	}
}

type AuthenticatorCreateEvent struct {
	ID       string     `json:"id"`
	Type     string     `json:"type"`
	Name     string     `json:"name"`
	LastUsed *time.Time `json:"last_used"`
	CredID   string     `json:"cred_id,omitzero"`
}

func (e *AuthenticatorCreateEvent) AsAuthenticator() Authenticator {
	return Authenticator{
		ID:       e.ID,
		Type:     e.Type,
		Name:     e.Name,
		LastUsed: e.LastUsed,
		CredID:   e.CredID,
	}
}

type AuthenticatorUpdateEvent struct {
	ID       string     `json:"id"`
	Type     string     `json:"type"`
	Name     string     `json:"name"`
	LastUsed *time.Time `json:"last_used"`
	CredID   string     `json:"cred_id,omitzero"`
}

func (e *AuthenticatorUpdateEvent) AsAuthenticator() Authenticator {
	return Authenticator{
		ID:       e.ID,
		Type:     e.Type,
		Name:     e.Name,
		LastUsed: e.LastUsed,
		CredID:   e.CredID,
	}
}

type AuthenticatorDeleteEvent struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

func (e *AuthenticatorDeleteEvent) AsAuthenticatorDelete() AuthenticatorDelete {
	return AuthenticatorDelete{
		ID:   e.ID,
		Type: e.Type,
	}
}

type ApplicationCommandPermissionsUpdateEvent struct {
	ID            snowflake.ID                    `json:"id"`
	ApplicationID snowflake.ID                    `json:"application_id"`
	GuildID       snowflake.ID                    `json:"guild_id"`
	Permissions   []ApplicationCommandPermissions `json:"permissions"`
}

func (e *ApplicationCommandPermissionsUpdateEvent) AsGuildApplicationCommandPermissions() GuildApplicationCommandPermissions {
	return GuildApplicationCommandPermissions{
		ID:            e.ID,
		ApplicationID: e.ApplicationID,
		GuildID:       e.GuildID,
		Permissions:   e.Permissions,
	}
}

type AutoModerationRuleCreateEvent struct {
	ID              snowflake.ID       `json:"id"`
	GuildID         snowflake.ID       `json:"guild_id"`
	Name            string             `json:"name"`
	CreatorID       snowflake.ID       `json:"creator_id"`
	EventType       AutomodEventType   `json:"event_type"`
	TriggerType     AutomodTriggerType `json:"trigger_type"`
	TriggerMetadata json.RawMessage    `json:"trigger_metadata"`
	Actions         []Action           `json:"actions"`
	Enabled         bool               `json:"enabled"`
	ExemptRoles     []snowflake.ID     `json:"exempt_roles"`
	ExemptChannels  []snowflake.ID     `json:"exempt_channels"`
}

func (e *AutoModerationRuleCreateEvent) AsAutomodRule() AutomodRule {
	return AutomodRule{
		ID:              e.ID,
		GuildID:         e.GuildID,
		Name:            e.Name,
		CreatorID:       e.CreatorID,
		EventType:       e.EventType,
		TriggerType:     e.TriggerType,
		TriggerMetadata: e.TriggerMetadata,
		Actions:         e.Actions,
		Enabled:         e.Enabled,
		ExemptRoles:     e.ExemptRoles,
		ExemptChannels:  e.ExemptChannels,
	}
}

type AutoModerationRuleUpdateEvent struct {
	ID              snowflake.ID       `json:"id"`
	GuildID         snowflake.ID       `json:"guild_id"`
	Name            string             `json:"name"`
	CreatorID       snowflake.ID       `json:"creator_id"`
	EventType       AutomodEventType   `json:"event_type"`
	TriggerType     AutomodTriggerType `json:"trigger_type"`
	TriggerMetadata json.RawMessage    `json:"trigger_metadata"`
	Actions         []Action           `json:"actions"`
	Enabled         bool               `json:"enabled"`
	ExemptRoles     []snowflake.ID     `json:"exempt_roles"`
	ExemptChannels  []snowflake.ID     `json:"exempt_channels"`
}

func (e *AutoModerationRuleUpdateEvent) AsAutomodRule() AutomodRule {
	return AutomodRule{
		ID:              e.ID,
		GuildID:         e.GuildID,
		Name:            e.Name,
		CreatorID:       e.CreatorID,
		EventType:       e.EventType,
		TriggerType:     e.TriggerType,
		TriggerMetadata: e.TriggerMetadata,
		Actions:         e.Actions,
		Enabled:         e.Enabled,
		ExemptRoles:     e.ExemptRoles,
		ExemptChannels:  e.ExemptChannels,
	}
}

type AutoModerationRuleDeleteEvent struct {
	ID              snowflake.ID       `json:"id"`
	GuildID         snowflake.ID       `json:"guild_id"`
	Name            string             `json:"name"`
	CreatorID       snowflake.ID       `json:"creator_id"`
	EventType       AutomodEventType   `json:"event_type"`
	TriggerType     AutomodTriggerType `json:"trigger_type"`
	TriggerMetadata json.RawMessage    `json:"trigger_metadata"`
	Actions         []Action           `json:"actions"`
	Enabled         bool               `json:"enabled"`
	ExemptRoles     []snowflake.ID     `json:"exempt_roles"`
	ExemptChannels  []snowflake.ID     `json:"exempt_channels"`
}

func (e *AutoModerationRuleDeleteEvent) AsAutomodRule() AutomodRule {
	return AutomodRule{
		ID:              e.ID,
		GuildID:         e.GuildID,
		Name:            e.Name,
		CreatorID:       e.CreatorID,
		EventType:       e.EventType,
		TriggerType:     e.TriggerType,
		TriggerMetadata: e.TriggerMetadata,
		Actions:         e.Actions,
		Enabled:         e.Enabled,
		ExemptRoles:     e.ExemptRoles,
		ExemptChannels:  e.ExemptChannels,
	}
}

type AutoModerationActionExecutionEvent struct {
	GuildID              snowflake.ID  `json:"guild_id"`
	Action               AutomodAction `json:"action"`
	RuleID               snowflake.ID  `json:"rule_id"`
	RuleTriggerType      int64         `json:"rule_trigger_type"`
	UserID               snowflake.ID  `json:"user_id"`
	ChannelID            snowflake.ID  `json:"channel_id,omitzero"`
	MessageID            snowflake.ID  `json:"message_id,omitzero"`
	AlertSystemMessageID snowflake.ID  `json:"alert_system_message_id,omitzero"`
	Content              string        `json:"content"`
	MatchedKeyword       *string       `json:"matched_keyword"`
	MatchedContent       *string       `json:"matched_content"`
}

func (e *AutoModerationActionExecutionEvent) AsAutoModerationActionExecution() AutoModerationActionExecution {
	return AutoModerationActionExecution{
		GuildID:              e.GuildID,
		Action:               e.Action,
		RuleID:               e.RuleID,
		RuleTriggerType:      e.RuleTriggerType,
		UserID:               e.UserID,
		ChannelID:            e.ChannelID,
		MessageID:            e.MessageID,
		AlertSystemMessageID: e.AlertSystemMessageID,
		Content:              e.Content,
		MatchedKeyword:       e.MatchedKeyword,
		MatchedContent:       e.MatchedContent,
	}
}

type AutoModerationMentionRaidDetectionEvent struct {
	GuildID                        snowflake.ID `json:"guild_id"`
	DecisionID                     string       `json:"decision_id"`
	SuspiciousMentionActivityUntil time.Time    `json:"suspicious_mention_activity_until"`
}

func (e *AutoModerationMentionRaidDetectionEvent) AsAutoModerationMentionRaidDetection() AutoModerationMentionRaidDetection {
	return AutoModerationMentionRaidDetection{
		GuildID:                        e.GuildID,
		DecisionID:                     e.DecisionID,
		SuspiciousMentionActivityUntil: e.SuspiciousMentionActivityUntil,
	}
}

type BillingPopupBridgeCallbackEvent struct {
	PaymentSourceType PaymentSourceType `json:"payment_source_type,omitzero"`
	State             string            `json:"state,omitzero"`
	Path              string            `json:"path,omitzero"`
	Query             map[string]string `json:"query,omitempty"`
}

func (e *BillingPopupBridgeCallbackEvent) AsBillingPopupBridgeCallback() BillingPopupBridgeCallback {
	return BillingPopupBridgeCallback{
		PaymentSourceType: e.PaymentSourceType,
		State:             e.State,
		Path:              e.Path,
		Query:             e.Query,
	}
}

type CallCreateEvent struct {
	ChannelID   snowflake.ID   `json:"channel_id"`
	MessageID   snowflake.ID   `json:"message_id"`
	Region      string         `json:"region"`
	Ringing     []snowflake.ID `json:"ringing"`
	VoiceStates []VoiceState   `json:"voice_states"`
}

func (e *CallCreateEvent) AsCallCreate() CallCreate {
	return CallCreate{
		ChannelID:   e.ChannelID,
		MessageID:   e.MessageID,
		Region:      e.Region,
		Ringing:     e.Ringing,
		VoiceStates: e.VoiceStates,
	}
}

type CallUpdateEvent struct {
	ChannelID snowflake.ID    `json:"channel_id"`
	MessageID snowflake.ID    `json:"message_id"`
	Region    string          `json:"region"`
	Ringing   json.RawMessage `json:"ringing"`
}

func (e *CallUpdateEvent) AsCallUpdate() CallUpdate {
	return CallUpdate{
		ChannelID: e.ChannelID,
		MessageID: e.MessageID,
		Region:    e.Region,
		Ringing:   e.Ringing,
	}
}

type CallDeleteEvent struct {
	ChannelID   snowflake.ID `json:"channel_id"`
	Unavailable bool         `json:"unavailable,omitzero"`
}

func (e *CallDeleteEvent) AsCallDelete() CallDelete {
	return CallDelete{
		ChannelID:   e.ChannelID,
		Unavailable: e.Unavailable,
	}
}

type ChannelCreateEvent struct {
	ID                                 snowflake.ID          `json:"id"`
	Type                               ChannelType           `json:"type"`
	GuildID                            snowflake.ID          `json:"guild_id,omitzero"`
	Position                           int64                 `json:"position,omitzero"`
	PermissionOverwrites               []PermissionOverwrite `json:"permission_overwrites,omitempty"`
	Name                               *string               `json:"name,omitempty"`
	Topic                              *string               `json:"topic,omitempty"`
	NSFW                               bool                  `json:"nsfw,omitzero"`
	LastMessageID                      *snowflake.ID         `json:"last_message_id,omitempty"`
	LastPinTimestamp                   *time.Time            `json:"last_pin_timestamp,omitempty"`
	Bitrate                            int64                 `json:"bitrate,omitzero"`
	UserLimit                          int64                 `json:"user_limit,omitzero"`
	RateLimitPerUser                   int64                 `json:"rate_limit_per_user,omitzero"`
	Recipients                         []PartialUser         `json:"recipients,omitempty"`
	RecipientFlags                     RecipientFlags        `json:"recipient_flags,omitzero"`
	Icon                               *string               `json:"icon,omitempty"`
	Nicks                              []ChannelNick         `json:"nicks,omitempty"`
	Managed                            bool                  `json:"managed,omitzero"`
	BlockedUserWarningDismissed        bool                  `json:"blocked_user_warning_dismissed,omitzero"`
	SafetyWarnings                     []SafetyWarning       `json:"safety_warnings,omitempty"`
	ApplicationID                      snowflake.ID          `json:"application_id,omitzero"`
	OwnerID                            snowflake.ID          `json:"owner_id,omitzero"`
	Owner                              *GuildMember          `json:"owner,omitempty"`
	ParentID                           *snowflake.ID         `json:"parent_id,omitempty"`
	RTCRegion                          *string               `json:"rtc_region,omitempty"`
	VideoQualityMode                   VideoQualityMode      `json:"video_quality_mode,omitzero"`
	TotalMessageSent                   int64                 `json:"total_message_sent,omitzero"`
	MessageCount                       int64                 `json:"message_count,omitzero"`
	MemberCount                        int64                 `json:"member_count,omitzero"`
	MemberIDsPreview                   []snowflake.ID        `json:"member_ids_preview,omitempty"`
	ThreadMetadata                     ThreadMetadata        `json:"thread_metadata,omitzero"`
	Member                             ThreadMember          `json:"member,omitzero"`
	DefaultAutoArchiveDuration         *int64                `json:"default_auto_archive_duration,omitempty"`
	DefaultThreadRateLimitPerUser      int64                 `json:"default_thread_rate_limit_per_user,omitzero"`
	Permissions                        Permissions           `json:"permissions,omitzero"`
	Flags                              ChannelFlags          `json:"flags,omitzero"`
	AvailableTags                      []json.RawMessage     `json:"available_tags,omitempty"`
	AppliedTags                        []snowflake.ID        `json:"applied_tags,omitempty"`
	DefaultReactionEmoji               *DefaultReaction      `json:"default_reaction_emoji,omitempty"`
	DefaultForumLayout                 ForumLayoutType       `json:"default_forum_layout,omitzero"`
	DefaultSortOrder                   *SortOrderType        `json:"default_sort_order,omitempty"`
	DefaultTagSetting                  SearchTagSetting      `json:"default_tag_setting,omitzero"`
	IconEmoji                          *IconEmoji            `json:"icon_emoji,omitempty"`
	ThemeColor                         *int64                `json:"theme_color,omitempty"`
	IsMessageRequest                   bool                  `json:"is_message_request,omitzero"`
	IsMessageRequestTimestamp          *time.Time            `json:"is_message_request_timestamp,omitempty"`
	IsSpam                             bool                  `json:"is_spam,omitzero"`
	Status                             *string               `json:"status,omitempty"`
	HdStreamingUntil                   *time.Time            `json:"hd_streaming_until,omitempty"`
	HdStreamingBuyerID                 *snowflake.ID         `json:"hd_streaming_buyer_id,omitempty"`
	LinkedLobby                        *LinkedLobby          `json:"linked_lobby,omitempty"`
	IsLinkable                         bool                  `json:"is_linkable,omitzero"`
	IsViewableAndWriteableByAllMembers bool                  `json:"is_viewable_and_writeable_by_all_members,omitzero"`
	Template                           string                `json:"template,omitzero"`
	Version                            string                `json:"version,omitzero"`
	OriginChannelID                    snowflake.ID          `json:"origin_channel_id,omitzero"`
}

func (e *ChannelCreateEvent) AsChannel() Channel {
	return Channel{
		ID:                                 e.ID,
		Type:                               e.Type,
		GuildID:                            e.GuildID,
		Position:                           e.Position,
		PermissionOverwrites:               e.PermissionOverwrites,
		Name:                               e.Name,
		Topic:                              e.Topic,
		NSFW:                               e.NSFW,
		LastMessageID:                      e.LastMessageID,
		LastPinTimestamp:                   e.LastPinTimestamp,
		Bitrate:                            e.Bitrate,
		UserLimit:                          e.UserLimit,
		RateLimitPerUser:                   e.RateLimitPerUser,
		Recipients:                         e.Recipients,
		RecipientFlags:                     e.RecipientFlags,
		Icon:                               e.Icon,
		Nicks:                              e.Nicks,
		Managed:                            e.Managed,
		BlockedUserWarningDismissed:        e.BlockedUserWarningDismissed,
		SafetyWarnings:                     e.SafetyWarnings,
		ApplicationID:                      e.ApplicationID,
		OwnerID:                            e.OwnerID,
		Owner:                              e.Owner,
		ParentID:                           e.ParentID,
		RTCRegion:                          e.RTCRegion,
		VideoQualityMode:                   e.VideoQualityMode,
		TotalMessageSent:                   e.TotalMessageSent,
		MessageCount:                       e.MessageCount,
		MemberCount:                        e.MemberCount,
		MemberIDsPreview:                   e.MemberIDsPreview,
		ThreadMetadata:                     e.ThreadMetadata,
		Member:                             e.Member,
		DefaultAutoArchiveDuration:         e.DefaultAutoArchiveDuration,
		DefaultThreadRateLimitPerUser:      e.DefaultThreadRateLimitPerUser,
		Permissions:                        e.Permissions,
		Flags:                              e.Flags,
		AvailableTags:                      e.AvailableTags,
		AppliedTags:                        e.AppliedTags,
		DefaultReactionEmoji:               e.DefaultReactionEmoji,
		DefaultForumLayout:                 e.DefaultForumLayout,
		DefaultSortOrder:                   e.DefaultSortOrder,
		DefaultTagSetting:                  e.DefaultTagSetting,
		IconEmoji:                          e.IconEmoji,
		ThemeColor:                         e.ThemeColor,
		IsMessageRequest:                   e.IsMessageRequest,
		IsMessageRequestTimestamp:          e.IsMessageRequestTimestamp,
		IsSpam:                             e.IsSpam,
		Status:                             e.Status,
		HdStreamingUntil:                   e.HdStreamingUntil,
		HdStreamingBuyerID:                 e.HdStreamingBuyerID,
		LinkedLobby:                        e.LinkedLobby,
		IsLinkable:                         e.IsLinkable,
		IsViewableAndWriteableByAllMembers: e.IsViewableAndWriteableByAllMembers,
		Template:                           e.Template,
		Version:                            e.Version,
	}
}

type ChannelUpdateEvent struct {
	ID                                 snowflake.ID          `json:"id"`
	Type                               ChannelType           `json:"type"`
	GuildID                            snowflake.ID          `json:"guild_id,omitzero"`
	Position                           int64                 `json:"position,omitzero"`
	PermissionOverwrites               []PermissionOverwrite `json:"permission_overwrites,omitempty"`
	Name                               *string               `json:"name,omitempty"`
	Topic                              *string               `json:"topic,omitempty"`
	NSFW                               bool                  `json:"nsfw,omitzero"`
	LastMessageID                      *snowflake.ID         `json:"last_message_id,omitempty"`
	LastPinTimestamp                   *time.Time            `json:"last_pin_timestamp,omitempty"`
	Bitrate                            int64                 `json:"bitrate,omitzero"`
	UserLimit                          int64                 `json:"user_limit,omitzero"`
	RateLimitPerUser                   int64                 `json:"rate_limit_per_user,omitzero"`
	Recipients                         []PartialUser         `json:"recipients,omitempty"`
	RecipientFlags                     RecipientFlags        `json:"recipient_flags,omitzero"`
	Icon                               *string               `json:"icon,omitempty"`
	Nicks                              []ChannelNick         `json:"nicks,omitempty"`
	Managed                            bool                  `json:"managed,omitzero"`
	BlockedUserWarningDismissed        bool                  `json:"blocked_user_warning_dismissed,omitzero"`
	SafetyWarnings                     []SafetyWarning       `json:"safety_warnings,omitempty"`
	ApplicationID                      snowflake.ID          `json:"application_id,omitzero"`
	OwnerID                            snowflake.ID          `json:"owner_id,omitzero"`
	Owner                              *GuildMember          `json:"owner,omitempty"`
	ParentID                           *snowflake.ID         `json:"parent_id,omitempty"`
	RTCRegion                          *string               `json:"rtc_region,omitempty"`
	VideoQualityMode                   VideoQualityMode      `json:"video_quality_mode,omitzero"`
	TotalMessageSent                   int64                 `json:"total_message_sent,omitzero"`
	MessageCount                       int64                 `json:"message_count,omitzero"`
	MemberCount                        int64                 `json:"member_count,omitzero"`
	MemberIDsPreview                   []snowflake.ID        `json:"member_ids_preview,omitempty"`
	ThreadMetadata                     ThreadMetadata        `json:"thread_metadata,omitzero"`
	Member                             ThreadMember          `json:"member,omitzero"`
	DefaultAutoArchiveDuration         *int64                `json:"default_auto_archive_duration,omitempty"`
	DefaultThreadRateLimitPerUser      int64                 `json:"default_thread_rate_limit_per_user,omitzero"`
	Permissions                        Permissions           `json:"permissions,omitzero"`
	Flags                              ChannelFlags          `json:"flags,omitzero"`
	AvailableTags                      []json.RawMessage     `json:"available_tags,omitempty"`
	AppliedTags                        []snowflake.ID        `json:"applied_tags,omitempty"`
	DefaultReactionEmoji               *DefaultReaction      `json:"default_reaction_emoji,omitempty"`
	DefaultForumLayout                 ForumLayoutType       `json:"default_forum_layout,omitzero"`
	DefaultSortOrder                   *SortOrderType        `json:"default_sort_order,omitempty"`
	DefaultTagSetting                  SearchTagSetting      `json:"default_tag_setting,omitzero"`
	IconEmoji                          *IconEmoji            `json:"icon_emoji,omitempty"`
	ThemeColor                         *int64                `json:"theme_color,omitempty"`
	IsMessageRequest                   bool                  `json:"is_message_request,omitzero"`
	IsMessageRequestTimestamp          *time.Time            `json:"is_message_request_timestamp,omitempty"`
	IsSpam                             bool                  `json:"is_spam,omitzero"`
	Status                             *string               `json:"status,omitempty"`
	HdStreamingUntil                   *time.Time            `json:"hd_streaming_until,omitempty"`
	HdStreamingBuyerID                 *snowflake.ID         `json:"hd_streaming_buyer_id,omitempty"`
	LinkedLobby                        *LinkedLobby          `json:"linked_lobby,omitempty"`
	IsLinkable                         bool                  `json:"is_linkable,omitzero"`
	IsViewableAndWriteableByAllMembers bool                  `json:"is_viewable_and_writeable_by_all_members,omitzero"`
	Template                           string                `json:"template,omitzero"`
	Version                            string                `json:"version,omitzero"`
}

func (e *ChannelUpdateEvent) AsChannel() Channel {
	return Channel{
		ID:                                 e.ID,
		Type:                               e.Type,
		GuildID:                            e.GuildID,
		Position:                           e.Position,
		PermissionOverwrites:               e.PermissionOverwrites,
		Name:                               e.Name,
		Topic:                              e.Topic,
		NSFW:                               e.NSFW,
		LastMessageID:                      e.LastMessageID,
		LastPinTimestamp:                   e.LastPinTimestamp,
		Bitrate:                            e.Bitrate,
		UserLimit:                          e.UserLimit,
		RateLimitPerUser:                   e.RateLimitPerUser,
		Recipients:                         e.Recipients,
		RecipientFlags:                     e.RecipientFlags,
		Icon:                               e.Icon,
		Nicks:                              e.Nicks,
		Managed:                            e.Managed,
		BlockedUserWarningDismissed:        e.BlockedUserWarningDismissed,
		SafetyWarnings:                     e.SafetyWarnings,
		ApplicationID:                      e.ApplicationID,
		OwnerID:                            e.OwnerID,
		Owner:                              e.Owner,
		ParentID:                           e.ParentID,
		RTCRegion:                          e.RTCRegion,
		VideoQualityMode:                   e.VideoQualityMode,
		TotalMessageSent:                   e.TotalMessageSent,
		MessageCount:                       e.MessageCount,
		MemberCount:                        e.MemberCount,
		MemberIDsPreview:                   e.MemberIDsPreview,
		ThreadMetadata:                     e.ThreadMetadata,
		Member:                             e.Member,
		DefaultAutoArchiveDuration:         e.DefaultAutoArchiveDuration,
		DefaultThreadRateLimitPerUser:      e.DefaultThreadRateLimitPerUser,
		Permissions:                        e.Permissions,
		Flags:                              e.Flags,
		AvailableTags:                      e.AvailableTags,
		AppliedTags:                        e.AppliedTags,
		DefaultReactionEmoji:               e.DefaultReactionEmoji,
		DefaultForumLayout:                 e.DefaultForumLayout,
		DefaultSortOrder:                   e.DefaultSortOrder,
		DefaultTagSetting:                  e.DefaultTagSetting,
		IconEmoji:                          e.IconEmoji,
		ThemeColor:                         e.ThemeColor,
		IsMessageRequest:                   e.IsMessageRequest,
		IsMessageRequestTimestamp:          e.IsMessageRequestTimestamp,
		IsSpam:                             e.IsSpam,
		Status:                             e.Status,
		HdStreamingUntil:                   e.HdStreamingUntil,
		HdStreamingBuyerID:                 e.HdStreamingBuyerID,
		LinkedLobby:                        e.LinkedLobby,
		IsLinkable:                         e.IsLinkable,
		IsViewableAndWriteableByAllMembers: e.IsViewableAndWriteableByAllMembers,
		Template:                           e.Template,
		Version:                            e.Version,
	}
}

type ChannelDeleteEvent struct {
	ID                                 snowflake.ID          `json:"id"`
	Type                               ChannelType           `json:"type"`
	GuildID                            snowflake.ID          `json:"guild_id,omitzero"`
	Position                           int64                 `json:"position,omitzero"`
	PermissionOverwrites               []PermissionOverwrite `json:"permission_overwrites,omitempty"`
	Name                               *string               `json:"name,omitempty"`
	Topic                              *string               `json:"topic,omitempty"`
	NSFW                               bool                  `json:"nsfw,omitzero"`
	LastMessageID                      *snowflake.ID         `json:"last_message_id,omitempty"`
	LastPinTimestamp                   *time.Time            `json:"last_pin_timestamp,omitempty"`
	Bitrate                            int64                 `json:"bitrate,omitzero"`
	UserLimit                          int64                 `json:"user_limit,omitzero"`
	RateLimitPerUser                   int64                 `json:"rate_limit_per_user,omitzero"`
	Recipients                         []PartialUser         `json:"recipients,omitempty"`
	RecipientFlags                     RecipientFlags        `json:"recipient_flags,omitzero"`
	Icon                               *string               `json:"icon,omitempty"`
	Nicks                              []ChannelNick         `json:"nicks,omitempty"`
	Managed                            bool                  `json:"managed,omitzero"`
	BlockedUserWarningDismissed        bool                  `json:"blocked_user_warning_dismissed,omitzero"`
	SafetyWarnings                     []SafetyWarning       `json:"safety_warnings,omitempty"`
	ApplicationID                      snowflake.ID          `json:"application_id,omitzero"`
	OwnerID                            snowflake.ID          `json:"owner_id,omitzero"`
	Owner                              *GuildMember          `json:"owner,omitempty"`
	ParentID                           *snowflake.ID         `json:"parent_id,omitempty"`
	RTCRegion                          *string               `json:"rtc_region,omitempty"`
	VideoQualityMode                   VideoQualityMode      `json:"video_quality_mode,omitzero"`
	TotalMessageSent                   int64                 `json:"total_message_sent,omitzero"`
	MessageCount                       int64                 `json:"message_count,omitzero"`
	MemberCount                        int64                 `json:"member_count,omitzero"`
	MemberIDsPreview                   []snowflake.ID        `json:"member_ids_preview,omitempty"`
	ThreadMetadata                     ThreadMetadata        `json:"thread_metadata,omitzero"`
	Member                             ThreadMember          `json:"member,omitzero"`
	DefaultAutoArchiveDuration         *int64                `json:"default_auto_archive_duration,omitempty"`
	DefaultThreadRateLimitPerUser      int64                 `json:"default_thread_rate_limit_per_user,omitzero"`
	Permissions                        Permissions           `json:"permissions,omitzero"`
	Flags                              ChannelFlags          `json:"flags,omitzero"`
	AvailableTags                      []json.RawMessage     `json:"available_tags,omitempty"`
	AppliedTags                        []snowflake.ID        `json:"applied_tags,omitempty"`
	DefaultReactionEmoji               *DefaultReaction      `json:"default_reaction_emoji,omitempty"`
	DefaultForumLayout                 ForumLayoutType       `json:"default_forum_layout,omitzero"`
	DefaultSortOrder                   *SortOrderType        `json:"default_sort_order,omitempty"`
	DefaultTagSetting                  SearchTagSetting      `json:"default_tag_setting,omitzero"`
	IconEmoji                          *IconEmoji            `json:"icon_emoji,omitempty"`
	ThemeColor                         *int64                `json:"theme_color,omitempty"`
	IsMessageRequest                   bool                  `json:"is_message_request,omitzero"`
	IsMessageRequestTimestamp          *time.Time            `json:"is_message_request_timestamp,omitempty"`
	IsSpam                             bool                  `json:"is_spam,omitzero"`
	Status                             *string               `json:"status,omitempty"`
	HdStreamingUntil                   *time.Time            `json:"hd_streaming_until,omitempty"`
	HdStreamingBuyerID                 *snowflake.ID         `json:"hd_streaming_buyer_id,omitempty"`
	LinkedLobby                        *LinkedLobby          `json:"linked_lobby,omitempty"`
	IsLinkable                         bool                  `json:"is_linkable,omitzero"`
	IsViewableAndWriteableByAllMembers bool                  `json:"is_viewable_and_writeable_by_all_members,omitzero"`
	Template                           string                `json:"template,omitzero"`
	Version                            string                `json:"version,omitzero"`
}

func (e *ChannelDeleteEvent) AsChannel() Channel {
	return Channel{
		ID:                                 e.ID,
		Type:                               e.Type,
		GuildID:                            e.GuildID,
		Position:                           e.Position,
		PermissionOverwrites:               e.PermissionOverwrites,
		Name:                               e.Name,
		Topic:                              e.Topic,
		NSFW:                               e.NSFW,
		LastMessageID:                      e.LastMessageID,
		LastPinTimestamp:                   e.LastPinTimestamp,
		Bitrate:                            e.Bitrate,
		UserLimit:                          e.UserLimit,
		RateLimitPerUser:                   e.RateLimitPerUser,
		Recipients:                         e.Recipients,
		RecipientFlags:                     e.RecipientFlags,
		Icon:                               e.Icon,
		Nicks:                              e.Nicks,
		Managed:                            e.Managed,
		BlockedUserWarningDismissed:        e.BlockedUserWarningDismissed,
		SafetyWarnings:                     e.SafetyWarnings,
		ApplicationID:                      e.ApplicationID,
		OwnerID:                            e.OwnerID,
		Owner:                              e.Owner,
		ParentID:                           e.ParentID,
		RTCRegion:                          e.RTCRegion,
		VideoQualityMode:                   e.VideoQualityMode,
		TotalMessageSent:                   e.TotalMessageSent,
		MessageCount:                       e.MessageCount,
		MemberCount:                        e.MemberCount,
		MemberIDsPreview:                   e.MemberIDsPreview,
		ThreadMetadata:                     e.ThreadMetadata,
		Member:                             e.Member,
		DefaultAutoArchiveDuration:         e.DefaultAutoArchiveDuration,
		DefaultThreadRateLimitPerUser:      e.DefaultThreadRateLimitPerUser,
		Permissions:                        e.Permissions,
		Flags:                              e.Flags,
		AvailableTags:                      e.AvailableTags,
		AppliedTags:                        e.AppliedTags,
		DefaultReactionEmoji:               e.DefaultReactionEmoji,
		DefaultForumLayout:                 e.DefaultForumLayout,
		DefaultSortOrder:                   e.DefaultSortOrder,
		DefaultTagSetting:                  e.DefaultTagSetting,
		IconEmoji:                          e.IconEmoji,
		ThemeColor:                         e.ThemeColor,
		IsMessageRequest:                   e.IsMessageRequest,
		IsMessageRequestTimestamp:          e.IsMessageRequestTimestamp,
		IsSpam:                             e.IsSpam,
		Status:                             e.Status,
		HdStreamingUntil:                   e.HdStreamingUntil,
		HdStreamingBuyerID:                 e.HdStreamingBuyerID,
		LinkedLobby:                        e.LinkedLobby,
		IsLinkable:                         e.IsLinkable,
		IsViewableAndWriteableByAllMembers: e.IsViewableAndWriteableByAllMembers,
		Template:                           e.Template,
		Version:                            e.Version,
	}
}

type ChannelSyncEvent struct {
	GuildID        snowflake.ID `json:"guild_id"`
	Channels       []Channel    `json:"channels"`
	IntegrityCheck bool         `json:"integrity_check"`
}

func (e *ChannelSyncEvent) AsChannelSync() ChannelSync {
	return ChannelSync{
		GuildID:        e.GuildID,
		Channels:       e.Channels,
		IntegrityCheck: e.IntegrityCheck,
	}
}

type ChannelUpdatePartialEvent struct {
	ID               snowflake.ID  `json:"id"`
	LastMessageID    *snowflake.ID `json:"last_message_id"`
	LastPinTimestamp *time.Time    `json:"last_pin_timestamp,omitempty"`
}

func (e *ChannelUpdatePartialEvent) AsChannelUnread() ChannelUnread {
	return ChannelUnread{
		ID:               e.ID,
		LastMessageID:    e.LastMessageID,
		LastPinTimestamp: e.LastPinTimestamp,
	}
}

type ChannelInfoEvent struct {
	GuildID  snowflake.ID      `json:"guild_id"`
	Channels []ChannelMetadata `json:"channels"`
}

func (e *ChannelInfoEvent) AsChannelInfo() ChannelInfo {
	return ChannelInfo{
		GuildID:  e.GuildID,
		Channels: e.Channels,
	}
}

type ChannelStatusesEvent struct {
	GuildID  snowflake.ID    `json:"guild_id"`
	Channels []ChannelStatus `json:"channels"`
}

func (e *ChannelStatusesEvent) AsChannelStatuses() ChannelStatuses {
	return ChannelStatuses{
		GuildID:  e.GuildID,
		Channels: e.Channels,
	}
}

type ChannelMemberCountUpdateEvent struct {
	GuildID       snowflake.ID `json:"guild_id"`
	ChannelID     snowflake.ID `json:"channel_id"`
	MemberCount   int64        `json:"member_count"`
	PresenceCount int64        `json:"presence_count"`
}

func (e *ChannelMemberCountUpdateEvent) AsChannelMemberCountUpdate() ChannelMemberCountUpdate {
	return ChannelMemberCountUpdate{
		GuildID:       e.GuildID,
		ChannelID:     e.ChannelID,
		MemberCount:   e.MemberCount,
		PresenceCount: e.PresenceCount,
	}
}

type ChannelUnreadUpdateEvent struct {
	GuildID              snowflake.ID    `json:"guild_id"`
	ChannelUnreadUpdates []ChannelUnread `json:"channel_unread_updates"`
}

func (e *ChannelUnreadUpdateEvent) AsChannelUnreadUpdate() ChannelUnreadUpdate {
	return ChannelUnreadUpdate{
		GuildID:              e.GuildID,
		ChannelUnreadUpdates: e.ChannelUnreadUpdates,
	}
}

type ChannelPinsUpdateEvent struct {
	GuildID          snowflake.ID `json:"guild_id,omitzero"`
	ChannelID        snowflake.ID `json:"channel_id"`
	LastPinTimestamp *time.Time   `json:"last_pin_timestamp,omitempty"`
}

func (e *ChannelPinsUpdateEvent) AsChannelPinsUpdate() ChannelPinsUpdate {
	return ChannelPinsUpdate{
		GuildID:          e.GuildID,
		ChannelID:        e.ChannelID,
		LastPinTimestamp: e.LastPinTimestamp,
	}
}

type ChannelPinsAckEvent struct {
	ChannelID snowflake.ID `json:"channel_id"`
	Timestamp time.Time    `json:"timestamp"`
	Version   int64        `json:"version"`
}

func (e *ChannelPinsAckEvent) AsChannelPinsAck() ChannelPinsAck {
	return ChannelPinsAck{
		ChannelID: e.ChannelID,
		Timestamp: e.Timestamp,
		Version:   e.Version,
	}
}

type ChannelRecipientAddEvent struct {
	ChannelID snowflake.ID `json:"channel_id"`
	User      PartialUser  `json:"user"`
	Nick      string       `json:"nick,omitzero"`
}

func (e *ChannelRecipientAddEvent) AsChannelRecipientAdd() ChannelRecipientAdd {
	return ChannelRecipientAdd{
		ChannelID: e.ChannelID,
		User:      e.User,
		Nick:      e.Nick,
	}
}

type ChannelRecipientRemoveEvent struct {
	ChannelID snowflake.ID `json:"channel_id"`
	User      PartialUser  `json:"user"`
}

func (e *ChannelRecipientRemoveEvent) AsChannelRecipientRemove() ChannelRecipientRemove {
	return ChannelRecipientRemove{
		ChannelID: e.ChannelID,
		User:      e.User,
	}
}

type ConsoleCommandUpdateEvent struct {
	ID     snowflake.ID    `json:"id"`
	Result string          `json:"result"`
	Error  json.RawMessage `json:"error"`
}

func (e *ConsoleCommandUpdateEvent) AsConsoleCommandUpdate() ConsoleCommandUpdate {
	return ConsoleCommandUpdate{
		ID:     e.ID,
		Result: e.Result,
		Error:  e.Error,
	}
}

type ConversationSummaryUpdateEvent struct {
	GuildID   snowflake.ID          `json:"guild_id"`
	ChannelID snowflake.ID          `json:"channel_id"`
	Summaries []ConversationSummary `json:"summaries"`
}

func (e *ConversationSummaryUpdateEvent) AsConversationSummaryUpdate() ConversationSummaryUpdate {
	return ConversationSummaryUpdate{
		GuildID:   e.GuildID,
		ChannelID: e.ChannelID,
		Summaries: e.Summaries,
	}
}

type DeletedEntityIDsEvent struct {
	GuildID  snowflake.ID   `json:"guild_id"`
	Stickers []snowflake.ID `json:"stickers,omitempty"`
	Roles    []snowflake.ID `json:"roles,omitempty"`
	Emojis   []snowflake.ID `json:"emojis,omitempty"`
	Channels []snowflake.ID `json:"channels,omitempty"`
}

func (e *DeletedEntityIDsEvent) AsDeletedEntityIDs() DeletedEntityIDs {
	return DeletedEntityIDs{
		GuildID:  e.GuildID,
		Stickers: e.Stickers,
		Roles:    e.Roles,
		Emojis:   e.Emojis,
		Channels: e.Channels,
	}
}

type DMSettingsUpsellShowEvent struct {
	GuildID snowflake.ID `json:"guild_id"`
}

func (e *DMSettingsUpsellShowEvent) AsDMSettingsUpsellShow() DMSettingsUpsellShow {
	return DMSettingsUpsellShow{
		GuildID: e.GuildID,
	}
}

type ThreadCreateEvent struct {
	ID                                 snowflake.ID          `json:"id"`
	Type                               ChannelType           `json:"type"`
	GuildID                            snowflake.ID          `json:"guild_id,omitzero"`
	Position                           int64                 `json:"position,omitzero"`
	PermissionOverwrites               []PermissionOverwrite `json:"permission_overwrites,omitempty"`
	Name                               *string               `json:"name,omitempty"`
	Topic                              *string               `json:"topic,omitempty"`
	NSFW                               bool                  `json:"nsfw,omitzero"`
	LastMessageID                      *snowflake.ID         `json:"last_message_id,omitempty"`
	LastPinTimestamp                   *time.Time            `json:"last_pin_timestamp,omitempty"`
	Bitrate                            int64                 `json:"bitrate,omitzero"`
	UserLimit                          int64                 `json:"user_limit,omitzero"`
	RateLimitPerUser                   int64                 `json:"rate_limit_per_user,omitzero"`
	Recipients                         []PartialUser         `json:"recipients,omitempty"`
	RecipientFlags                     RecipientFlags        `json:"recipient_flags,omitzero"`
	Icon                               *string               `json:"icon,omitempty"`
	Nicks                              []ChannelNick         `json:"nicks,omitempty"`
	Managed                            bool                  `json:"managed,omitzero"`
	BlockedUserWarningDismissed        bool                  `json:"blocked_user_warning_dismissed,omitzero"`
	SafetyWarnings                     []SafetyWarning       `json:"safety_warnings,omitempty"`
	ApplicationID                      snowflake.ID          `json:"application_id,omitzero"`
	OwnerID                            snowflake.ID          `json:"owner_id,omitzero"`
	Owner                              *GuildMember          `json:"owner,omitempty"`
	ParentID                           *snowflake.ID         `json:"parent_id,omitempty"`
	RTCRegion                          *string               `json:"rtc_region,omitempty"`
	VideoQualityMode                   VideoQualityMode      `json:"video_quality_mode,omitzero"`
	TotalMessageSent                   int64                 `json:"total_message_sent,omitzero"`
	MessageCount                       int64                 `json:"message_count,omitzero"`
	MemberCount                        int64                 `json:"member_count,omitzero"`
	MemberIDsPreview                   []snowflake.ID        `json:"member_ids_preview,omitempty"`
	ThreadMetadata                     ThreadMetadata        `json:"thread_metadata,omitzero"`
	Member                             ThreadMember          `json:"member,omitzero"`
	DefaultAutoArchiveDuration         *int64                `json:"default_auto_archive_duration,omitempty"`
	DefaultThreadRateLimitPerUser      int64                 `json:"default_thread_rate_limit_per_user,omitzero"`
	Permissions                        Permissions           `json:"permissions,omitzero"`
	Flags                              ChannelFlags          `json:"flags,omitzero"`
	AvailableTags                      []json.RawMessage     `json:"available_tags,omitempty"`
	AppliedTags                        []snowflake.ID        `json:"applied_tags,omitempty"`
	DefaultReactionEmoji               *DefaultReaction      `json:"default_reaction_emoji,omitempty"`
	DefaultForumLayout                 ForumLayoutType       `json:"default_forum_layout,omitzero"`
	DefaultSortOrder                   *SortOrderType        `json:"default_sort_order,omitempty"`
	DefaultTagSetting                  SearchTagSetting      `json:"default_tag_setting,omitzero"`
	IconEmoji                          *IconEmoji            `json:"icon_emoji,omitempty"`
	ThemeColor                         *int64                `json:"theme_color,omitempty"`
	IsMessageRequest                   bool                  `json:"is_message_request,omitzero"`
	IsMessageRequestTimestamp          *time.Time            `json:"is_message_request_timestamp,omitempty"`
	IsSpam                             bool                  `json:"is_spam,omitzero"`
	Status                             *string               `json:"status,omitempty"`
	HdStreamingUntil                   *time.Time            `json:"hd_streaming_until,omitempty"`
	HdStreamingBuyerID                 *snowflake.ID         `json:"hd_streaming_buyer_id,omitempty"`
	LinkedLobby                        *LinkedLobby          `json:"linked_lobby,omitempty"`
	IsLinkable                         bool                  `json:"is_linkable,omitzero"`
	IsViewableAndWriteableByAllMembers bool                  `json:"is_viewable_and_writeable_by_all_members,omitzero"`
	Template                           string                `json:"template,omitzero"`
	Version                            string                `json:"version,omitzero"`
	NewlyCreated                       bool                  `json:"newly_created,omitzero"`
}

func (e *ThreadCreateEvent) AsChannel() Channel {
	return Channel{
		ID:                                 e.ID,
		Type:                               e.Type,
		GuildID:                            e.GuildID,
		Position:                           e.Position,
		PermissionOverwrites:               e.PermissionOverwrites,
		Name:                               e.Name,
		Topic:                              e.Topic,
		NSFW:                               e.NSFW,
		LastMessageID:                      e.LastMessageID,
		LastPinTimestamp:                   e.LastPinTimestamp,
		Bitrate:                            e.Bitrate,
		UserLimit:                          e.UserLimit,
		RateLimitPerUser:                   e.RateLimitPerUser,
		Recipients:                         e.Recipients,
		RecipientFlags:                     e.RecipientFlags,
		Icon:                               e.Icon,
		Nicks:                              e.Nicks,
		Managed:                            e.Managed,
		BlockedUserWarningDismissed:        e.BlockedUserWarningDismissed,
		SafetyWarnings:                     e.SafetyWarnings,
		ApplicationID:                      e.ApplicationID,
		OwnerID:                            e.OwnerID,
		Owner:                              e.Owner,
		ParentID:                           e.ParentID,
		RTCRegion:                          e.RTCRegion,
		VideoQualityMode:                   e.VideoQualityMode,
		TotalMessageSent:                   e.TotalMessageSent,
		MessageCount:                       e.MessageCount,
		MemberCount:                        e.MemberCount,
		MemberIDsPreview:                   e.MemberIDsPreview,
		ThreadMetadata:                     e.ThreadMetadata,
		Member:                             e.Member,
		DefaultAutoArchiveDuration:         e.DefaultAutoArchiveDuration,
		DefaultThreadRateLimitPerUser:      e.DefaultThreadRateLimitPerUser,
		Permissions:                        e.Permissions,
		Flags:                              e.Flags,
		AvailableTags:                      e.AvailableTags,
		AppliedTags:                        e.AppliedTags,
		DefaultReactionEmoji:               e.DefaultReactionEmoji,
		DefaultForumLayout:                 e.DefaultForumLayout,
		DefaultSortOrder:                   e.DefaultSortOrder,
		DefaultTagSetting:                  e.DefaultTagSetting,
		IconEmoji:                          e.IconEmoji,
		ThemeColor:                         e.ThemeColor,
		IsMessageRequest:                   e.IsMessageRequest,
		IsMessageRequestTimestamp:          e.IsMessageRequestTimestamp,
		IsSpam:                             e.IsSpam,
		Status:                             e.Status,
		HdStreamingUntil:                   e.HdStreamingUntil,
		HdStreamingBuyerID:                 e.HdStreamingBuyerID,
		LinkedLobby:                        e.LinkedLobby,
		IsLinkable:                         e.IsLinkable,
		IsViewableAndWriteableByAllMembers: e.IsViewableAndWriteableByAllMembers,
		Template:                           e.Template,
		Version:                            e.Version,
	}
}

type ThreadUpdateEvent struct {
	ID                                 snowflake.ID          `json:"id"`
	Type                               ChannelType           `json:"type"`
	GuildID                            snowflake.ID          `json:"guild_id,omitzero"`
	Position                           int64                 `json:"position,omitzero"`
	PermissionOverwrites               []PermissionOverwrite `json:"permission_overwrites,omitempty"`
	Name                               *string               `json:"name,omitempty"`
	Topic                              *string               `json:"topic,omitempty"`
	NSFW                               bool                  `json:"nsfw,omitzero"`
	LastMessageID                      *snowflake.ID         `json:"last_message_id,omitempty"`
	LastPinTimestamp                   *time.Time            `json:"last_pin_timestamp,omitempty"`
	Bitrate                            int64                 `json:"bitrate,omitzero"`
	UserLimit                          int64                 `json:"user_limit,omitzero"`
	RateLimitPerUser                   int64                 `json:"rate_limit_per_user,omitzero"`
	Recipients                         []PartialUser         `json:"recipients,omitempty"`
	RecipientFlags                     RecipientFlags        `json:"recipient_flags,omitzero"`
	Icon                               *string               `json:"icon,omitempty"`
	Nicks                              []ChannelNick         `json:"nicks,omitempty"`
	Managed                            bool                  `json:"managed,omitzero"`
	BlockedUserWarningDismissed        bool                  `json:"blocked_user_warning_dismissed,omitzero"`
	SafetyWarnings                     []SafetyWarning       `json:"safety_warnings,omitempty"`
	ApplicationID                      snowflake.ID          `json:"application_id,omitzero"`
	OwnerID                            snowflake.ID          `json:"owner_id,omitzero"`
	Owner                              *GuildMember          `json:"owner,omitempty"`
	ParentID                           *snowflake.ID         `json:"parent_id,omitempty"`
	RTCRegion                          *string               `json:"rtc_region,omitempty"`
	VideoQualityMode                   VideoQualityMode      `json:"video_quality_mode,omitzero"`
	TotalMessageSent                   int64                 `json:"total_message_sent,omitzero"`
	MessageCount                       int64                 `json:"message_count,omitzero"`
	MemberCount                        int64                 `json:"member_count,omitzero"`
	MemberIDsPreview                   []snowflake.ID        `json:"member_ids_preview,omitempty"`
	ThreadMetadata                     ThreadMetadata        `json:"thread_metadata,omitzero"`
	Member                             ThreadMember          `json:"member,omitzero"`
	DefaultAutoArchiveDuration         *int64                `json:"default_auto_archive_duration,omitempty"`
	DefaultThreadRateLimitPerUser      int64                 `json:"default_thread_rate_limit_per_user,omitzero"`
	Permissions                        Permissions           `json:"permissions,omitzero"`
	Flags                              ChannelFlags          `json:"flags,omitzero"`
	AvailableTags                      []json.RawMessage     `json:"available_tags,omitempty"`
	AppliedTags                        []snowflake.ID        `json:"applied_tags,omitempty"`
	DefaultReactionEmoji               *DefaultReaction      `json:"default_reaction_emoji,omitempty"`
	DefaultForumLayout                 ForumLayoutType       `json:"default_forum_layout,omitzero"`
	DefaultSortOrder                   *SortOrderType        `json:"default_sort_order,omitempty"`
	DefaultTagSetting                  SearchTagSetting      `json:"default_tag_setting,omitzero"`
	IconEmoji                          *IconEmoji            `json:"icon_emoji,omitempty"`
	ThemeColor                         *int64                `json:"theme_color,omitempty"`
	IsMessageRequest                   bool                  `json:"is_message_request,omitzero"`
	IsMessageRequestTimestamp          *time.Time            `json:"is_message_request_timestamp,omitempty"`
	IsSpam                             bool                  `json:"is_spam,omitzero"`
	Status                             *string               `json:"status,omitempty"`
	HdStreamingUntil                   *time.Time            `json:"hd_streaming_until,omitempty"`
	HdStreamingBuyerID                 *snowflake.ID         `json:"hd_streaming_buyer_id,omitempty"`
	LinkedLobby                        *LinkedLobby          `json:"linked_lobby,omitempty"`
	IsLinkable                         bool                  `json:"is_linkable,omitzero"`
	IsViewableAndWriteableByAllMembers bool                  `json:"is_viewable_and_writeable_by_all_members,omitzero"`
	Template                           string                `json:"template,omitzero"`
	Version                            string                `json:"version,omitzero"`
}

func (e *ThreadUpdateEvent) AsChannel() Channel {
	return Channel{
		ID:                                 e.ID,
		Type:                               e.Type,
		GuildID:                            e.GuildID,
		Position:                           e.Position,
		PermissionOverwrites:               e.PermissionOverwrites,
		Name:                               e.Name,
		Topic:                              e.Topic,
		NSFW:                               e.NSFW,
		LastMessageID:                      e.LastMessageID,
		LastPinTimestamp:                   e.LastPinTimestamp,
		Bitrate:                            e.Bitrate,
		UserLimit:                          e.UserLimit,
		RateLimitPerUser:                   e.RateLimitPerUser,
		Recipients:                         e.Recipients,
		RecipientFlags:                     e.RecipientFlags,
		Icon:                               e.Icon,
		Nicks:                              e.Nicks,
		Managed:                            e.Managed,
		BlockedUserWarningDismissed:        e.BlockedUserWarningDismissed,
		SafetyWarnings:                     e.SafetyWarnings,
		ApplicationID:                      e.ApplicationID,
		OwnerID:                            e.OwnerID,
		Owner:                              e.Owner,
		ParentID:                           e.ParentID,
		RTCRegion:                          e.RTCRegion,
		VideoQualityMode:                   e.VideoQualityMode,
		TotalMessageSent:                   e.TotalMessageSent,
		MessageCount:                       e.MessageCount,
		MemberCount:                        e.MemberCount,
		MemberIDsPreview:                   e.MemberIDsPreview,
		ThreadMetadata:                     e.ThreadMetadata,
		Member:                             e.Member,
		DefaultAutoArchiveDuration:         e.DefaultAutoArchiveDuration,
		DefaultThreadRateLimitPerUser:      e.DefaultThreadRateLimitPerUser,
		Permissions:                        e.Permissions,
		Flags:                              e.Flags,
		AvailableTags:                      e.AvailableTags,
		AppliedTags:                        e.AppliedTags,
		DefaultReactionEmoji:               e.DefaultReactionEmoji,
		DefaultForumLayout:                 e.DefaultForumLayout,
		DefaultSortOrder:                   e.DefaultSortOrder,
		DefaultTagSetting:                  e.DefaultTagSetting,
		IconEmoji:                          e.IconEmoji,
		ThemeColor:                         e.ThemeColor,
		IsMessageRequest:                   e.IsMessageRequest,
		IsMessageRequestTimestamp:          e.IsMessageRequestTimestamp,
		IsSpam:                             e.IsSpam,
		Status:                             e.Status,
		HdStreamingUntil:                   e.HdStreamingUntil,
		HdStreamingBuyerID:                 e.HdStreamingBuyerID,
		LinkedLobby:                        e.LinkedLobby,
		IsLinkable:                         e.IsLinkable,
		IsViewableAndWriteableByAllMembers: e.IsViewableAndWriteableByAllMembers,
		Template:                           e.Template,
		Version:                            e.Version,
	}
}

type ThreadDeleteEvent struct {
	ID                                 snowflake.ID          `json:"id"`
	Type                               ChannelType           `json:"type"`
	GuildID                            snowflake.ID          `json:"guild_id,omitzero"`
	Position                           int64                 `json:"position,omitzero"`
	PermissionOverwrites               []PermissionOverwrite `json:"permission_overwrites,omitempty"`
	Name                               *string               `json:"name,omitempty"`
	Topic                              *string               `json:"topic,omitempty"`
	NSFW                               bool                  `json:"nsfw,omitzero"`
	LastMessageID                      *snowflake.ID         `json:"last_message_id,omitempty"`
	LastPinTimestamp                   *time.Time            `json:"last_pin_timestamp,omitempty"`
	Bitrate                            int64                 `json:"bitrate,omitzero"`
	UserLimit                          int64                 `json:"user_limit,omitzero"`
	RateLimitPerUser                   int64                 `json:"rate_limit_per_user,omitzero"`
	Recipients                         []PartialUser         `json:"recipients,omitempty"`
	RecipientFlags                     RecipientFlags        `json:"recipient_flags,omitzero"`
	Icon                               *string               `json:"icon,omitempty"`
	Nicks                              []ChannelNick         `json:"nicks,omitempty"`
	Managed                            bool                  `json:"managed,omitzero"`
	BlockedUserWarningDismissed        bool                  `json:"blocked_user_warning_dismissed,omitzero"`
	SafetyWarnings                     []SafetyWarning       `json:"safety_warnings,omitempty"`
	ApplicationID                      snowflake.ID          `json:"application_id,omitzero"`
	OwnerID                            snowflake.ID          `json:"owner_id,omitzero"`
	Owner                              *GuildMember          `json:"owner,omitempty"`
	ParentID                           *snowflake.ID         `json:"parent_id,omitempty"`
	RTCRegion                          *string               `json:"rtc_region,omitempty"`
	VideoQualityMode                   VideoQualityMode      `json:"video_quality_mode,omitzero"`
	TotalMessageSent                   int64                 `json:"total_message_sent,omitzero"`
	MessageCount                       int64                 `json:"message_count,omitzero"`
	MemberCount                        int64                 `json:"member_count,omitzero"`
	MemberIDsPreview                   []snowflake.ID        `json:"member_ids_preview,omitempty"`
	ThreadMetadata                     ThreadMetadata        `json:"thread_metadata,omitzero"`
	Member                             ThreadMember          `json:"member,omitzero"`
	DefaultAutoArchiveDuration         *int64                `json:"default_auto_archive_duration,omitempty"`
	DefaultThreadRateLimitPerUser      int64                 `json:"default_thread_rate_limit_per_user,omitzero"`
	Permissions                        Permissions           `json:"permissions,omitzero"`
	Flags                              ChannelFlags          `json:"flags,omitzero"`
	AvailableTags                      []json.RawMessage     `json:"available_tags,omitempty"`
	AppliedTags                        []snowflake.ID        `json:"applied_tags,omitempty"`
	DefaultReactionEmoji               *DefaultReaction      `json:"default_reaction_emoji,omitempty"`
	DefaultForumLayout                 ForumLayoutType       `json:"default_forum_layout,omitzero"`
	DefaultSortOrder                   *SortOrderType        `json:"default_sort_order,omitempty"`
	DefaultTagSetting                  SearchTagSetting      `json:"default_tag_setting,omitzero"`
	IconEmoji                          *IconEmoji            `json:"icon_emoji,omitempty"`
	ThemeColor                         *int64                `json:"theme_color,omitempty"`
	IsMessageRequest                   bool                  `json:"is_message_request,omitzero"`
	IsMessageRequestTimestamp          *time.Time            `json:"is_message_request_timestamp,omitempty"`
	IsSpam                             bool                  `json:"is_spam,omitzero"`
	Status                             *string               `json:"status,omitempty"`
	HdStreamingUntil                   *time.Time            `json:"hd_streaming_until,omitempty"`
	HdStreamingBuyerID                 *snowflake.ID         `json:"hd_streaming_buyer_id,omitempty"`
	LinkedLobby                        *LinkedLobby          `json:"linked_lobby,omitempty"`
	IsLinkable                         bool                  `json:"is_linkable,omitzero"`
	IsViewableAndWriteableByAllMembers bool                  `json:"is_viewable_and_writeable_by_all_members,omitzero"`
	Template                           string                `json:"template,omitzero"`
	Version                            string                `json:"version,omitzero"`
}

func (e *ThreadDeleteEvent) AsChannel() Channel {
	return Channel{
		ID:                                 e.ID,
		Type:                               e.Type,
		GuildID:                            e.GuildID,
		Position:                           e.Position,
		PermissionOverwrites:               e.PermissionOverwrites,
		Name:                               e.Name,
		Topic:                              e.Topic,
		NSFW:                               e.NSFW,
		LastMessageID:                      e.LastMessageID,
		LastPinTimestamp:                   e.LastPinTimestamp,
		Bitrate:                            e.Bitrate,
		UserLimit:                          e.UserLimit,
		RateLimitPerUser:                   e.RateLimitPerUser,
		Recipients:                         e.Recipients,
		RecipientFlags:                     e.RecipientFlags,
		Icon:                               e.Icon,
		Nicks:                              e.Nicks,
		Managed:                            e.Managed,
		BlockedUserWarningDismissed:        e.BlockedUserWarningDismissed,
		SafetyWarnings:                     e.SafetyWarnings,
		ApplicationID:                      e.ApplicationID,
		OwnerID:                            e.OwnerID,
		Owner:                              e.Owner,
		ParentID:                           e.ParentID,
		RTCRegion:                          e.RTCRegion,
		VideoQualityMode:                   e.VideoQualityMode,
		TotalMessageSent:                   e.TotalMessageSent,
		MessageCount:                       e.MessageCount,
		MemberCount:                        e.MemberCount,
		MemberIDsPreview:                   e.MemberIDsPreview,
		ThreadMetadata:                     e.ThreadMetadata,
		Member:                             e.Member,
		DefaultAutoArchiveDuration:         e.DefaultAutoArchiveDuration,
		DefaultThreadRateLimitPerUser:      e.DefaultThreadRateLimitPerUser,
		Permissions:                        e.Permissions,
		Flags:                              e.Flags,
		AvailableTags:                      e.AvailableTags,
		AppliedTags:                        e.AppliedTags,
		DefaultReactionEmoji:               e.DefaultReactionEmoji,
		DefaultForumLayout:                 e.DefaultForumLayout,
		DefaultSortOrder:                   e.DefaultSortOrder,
		DefaultTagSetting:                  e.DefaultTagSetting,
		IconEmoji:                          e.IconEmoji,
		ThemeColor:                         e.ThemeColor,
		IsMessageRequest:                   e.IsMessageRequest,
		IsMessageRequestTimestamp:          e.IsMessageRequestTimestamp,
		IsSpam:                             e.IsSpam,
		Status:                             e.Status,
		HdStreamingUntil:                   e.HdStreamingUntil,
		HdStreamingBuyerID:                 e.HdStreamingBuyerID,
		LinkedLobby:                        e.LinkedLobby,
		IsLinkable:                         e.IsLinkable,
		IsViewableAndWriteableByAllMembers: e.IsViewableAndWriteableByAllMembers,
		Template:                           e.Template,
		Version:                            e.Version,
	}
}

type ThreadListSyncEvent struct {
	GuildID    snowflake.ID   `json:"guild_id"`
	ChannelIDs []snowflake.ID `json:"channel_ids,omitempty"`
	Threads    []Channel      `json:"threads"`
	Members    []ThreadMember `json:"members"`
}

func (e *ThreadListSyncEvent) AsThreadListSync() ThreadListSync {
	return ThreadListSync{
		GuildID:    e.GuildID,
		ChannelIDs: e.ChannelIDs,
		Threads:    e.Threads,
		Members:    e.Members,
	}
}

type ThreadMemberUpdateEvent struct {
	ID            snowflake.ID      `json:"id,omitzero"`
	UserID        snowflake.ID      `json:"user_id,omitzero"`
	JoinTimestamp time.Time         `json:"join_timestamp"`
	Flags         ThreadMemberFlags `json:"flags"`
	Muted         bool              `json:"muted,omitzero"`
	MuteConfig    *MuteConfig       `json:"mute_config,omitempty"`
	Member        GuildMember       `json:"member,omitzero"`
	GuildID       snowflake.ID      `json:"guild_id"`
}

func (e *ThreadMemberUpdateEvent) AsThreadMember() ThreadMember {
	return ThreadMember{
		ID:            e.ID,
		UserID:        e.UserID,
		JoinTimestamp: e.JoinTimestamp,
		Flags:         e.Flags,
		Muted:         e.Muted,
		MuteConfig:    e.MuteConfig,
		Member:        e.Member,
	}
}

type ThreadMembersUpdateEvent struct {
	ID               snowflake.ID    `json:"id"`
	GuildID          snowflake.ID    `json:"guild_id"`
	MemberCount      int64           `json:"member_count"`
	AddedMembers     []ThreadMember  `json:"added_members,omitempty"`
	RemovedMemberIDs json.RawMessage `json:"removed_member_ids,omitempty"`
}

func (e *ThreadMembersUpdateEvent) AsThreadMembersUpdate() ThreadMembersUpdate {
	return ThreadMembersUpdate{
		ID:               e.ID,
		GuildID:          e.GuildID,
		MemberCount:      e.MemberCount,
		AddedMembers:     e.AddedMembers,
		RemovedMemberIDs: e.RemovedMemberIDs,
	}
}

type EmbeddedActivityUpdateV2Event struct {
	ApplicationID       snowflake.ID                  `json:"application_id"`
	CompositeInstanceID string                        `json:"composite_instance_id"`
	InstanceID          snowflake.ID                  `json:"instance_id"`
	LaunchID            snowflake.ID                  `json:"launch_id"`
	Location            EmbeddedActivityLocation      `json:"location"`
	Participants        []EmbeddedActivityParticipant `json:"participants"`
}

func (e *EmbeddedActivityUpdateV2Event) AsEmbeddedActivityInstance() EmbeddedActivityInstance {
	return EmbeddedActivityInstance{
		ApplicationID:       e.ApplicationID,
		CompositeInstanceID: e.CompositeInstanceID,
		InstanceID:          e.InstanceID,
		LaunchID:            e.LaunchID,
		Location:            e.Location,
		Participants:        e.Participants,
	}
}

type EntitlementCreateEvent struct {
	ID                snowflake.ID                 `json:"id"`
	Type              EntitlementType              `json:"type"`
	SKUID             snowflake.ID                 `json:"sku_id"`
	ApplicationID     snowflake.ID                 `json:"application_id"`
	UserID            snowflake.ID                 `json:"user_id"`
	User              PartialUser                  `json:"user,omitzero"`
	GuildID           snowflake.ID                 `json:"guild_id,omitzero"`
	ParentID          snowflake.ID                 `json:"parent_id,omitzero"`
	Deleted           bool                         `json:"deleted"`
	Consumed          bool                         `json:"consumed,omitzero"`
	Branches          []snowflake.ID               `json:"branches,omitempty"`
	StartsAt          *time.Time                   `json:"starts_at"`
	EndsAt            *time.Time                   `json:"ends_at"`
	PromotionID       *snowflake.ID                `json:"promotion_id"`
	SubscriptionID    snowflake.ID                 `json:"subscription_id,omitzero"`
	GiftCodeFlags     GiftCodeFlags                `json:"gift_code_flags"`
	GiftCodeBatchID   snowflake.ID                 `json:"gift_code_batch_id,omitzero"`
	GifterUserID      snowflake.ID                 `json:"gifter_user_id,omitzero"`
	GiftStyle         GiftStyle                    `json:"gift_style,omitzero"`
	FulfillmentStatus EntitlementFulfillmentStatus `json:"fulfillment_status,omitzero"`
	FulfilledAt       time.Time                    `json:"fulfilled_at,omitzero"`
	SourceType        EntitlementSourceType        `json:"source_type,omitzero"`
	TenantMetadata    TenantMetadata               `json:"tenant_metadata,omitzero"`
	SKU               SKU                          `json:"sku,omitzero"`
	SubscriptionPlan  PartialSubscriptionPlan      `json:"subscription_plan,omitzero"`
}

func (e *EntitlementCreateEvent) AsEntitlement() Entitlement {
	return Entitlement{
		ID:                e.ID,
		Type:              e.Type,
		SKUID:             e.SKUID,
		ApplicationID:     e.ApplicationID,
		UserID:            e.UserID,
		User:              e.User,
		GuildID:           e.GuildID,
		ParentID:          e.ParentID,
		Deleted:           e.Deleted,
		Consumed:          e.Consumed,
		Branches:          e.Branches,
		StartsAt:          e.StartsAt,
		EndsAt:            e.EndsAt,
		PromotionID:       e.PromotionID,
		SubscriptionID:    e.SubscriptionID,
		GiftCodeFlags:     e.GiftCodeFlags,
		GiftCodeBatchID:   e.GiftCodeBatchID,
		GifterUserID:      e.GifterUserID,
		GiftStyle:         e.GiftStyle,
		FulfillmentStatus: e.FulfillmentStatus,
		FulfilledAt:       e.FulfilledAt,
		SourceType:        e.SourceType,
		TenantMetadata:    e.TenantMetadata,
		SKU:               e.SKU,
		SubscriptionPlan:  e.SubscriptionPlan,
	}
}

type EntitlementUpdateEvent struct {
	ID                snowflake.ID                 `json:"id"`
	Type              EntitlementType              `json:"type"`
	SKUID             snowflake.ID                 `json:"sku_id"`
	ApplicationID     snowflake.ID                 `json:"application_id"`
	UserID            snowflake.ID                 `json:"user_id"`
	User              PartialUser                  `json:"user,omitzero"`
	GuildID           snowflake.ID                 `json:"guild_id,omitzero"`
	ParentID          snowflake.ID                 `json:"parent_id,omitzero"`
	Deleted           bool                         `json:"deleted"`
	Consumed          bool                         `json:"consumed,omitzero"`
	Branches          []snowflake.ID               `json:"branches,omitempty"`
	StartsAt          *time.Time                   `json:"starts_at"`
	EndsAt            *time.Time                   `json:"ends_at"`
	PromotionID       *snowflake.ID                `json:"promotion_id"`
	SubscriptionID    snowflake.ID                 `json:"subscription_id,omitzero"`
	GiftCodeFlags     GiftCodeFlags                `json:"gift_code_flags"`
	GiftCodeBatchID   snowflake.ID                 `json:"gift_code_batch_id,omitzero"`
	GifterUserID      snowflake.ID                 `json:"gifter_user_id,omitzero"`
	GiftStyle         GiftStyle                    `json:"gift_style,omitzero"`
	FulfillmentStatus EntitlementFulfillmentStatus `json:"fulfillment_status,omitzero"`
	FulfilledAt       time.Time                    `json:"fulfilled_at,omitzero"`
	SourceType        EntitlementSourceType        `json:"source_type,omitzero"`
	TenantMetadata    TenantMetadata               `json:"tenant_metadata,omitzero"`
	SKU               SKU                          `json:"sku,omitzero"`
	SubscriptionPlan  PartialSubscriptionPlan      `json:"subscription_plan,omitzero"`
}

func (e *EntitlementUpdateEvent) AsEntitlement() Entitlement {
	return Entitlement{
		ID:                e.ID,
		Type:              e.Type,
		SKUID:             e.SKUID,
		ApplicationID:     e.ApplicationID,
		UserID:            e.UserID,
		User:              e.User,
		GuildID:           e.GuildID,
		ParentID:          e.ParentID,
		Deleted:           e.Deleted,
		Consumed:          e.Consumed,
		Branches:          e.Branches,
		StartsAt:          e.StartsAt,
		EndsAt:            e.EndsAt,
		PromotionID:       e.PromotionID,
		SubscriptionID:    e.SubscriptionID,
		GiftCodeFlags:     e.GiftCodeFlags,
		GiftCodeBatchID:   e.GiftCodeBatchID,
		GifterUserID:      e.GifterUserID,
		GiftStyle:         e.GiftStyle,
		FulfillmentStatus: e.FulfillmentStatus,
		FulfilledAt:       e.FulfilledAt,
		SourceType:        e.SourceType,
		TenantMetadata:    e.TenantMetadata,
		SKU:               e.SKU,
		SubscriptionPlan:  e.SubscriptionPlan,
	}
}

type EntitlementDeleteEvent struct {
	ID                snowflake.ID                 `json:"id"`
	Type              EntitlementType              `json:"type"`
	SKUID             snowflake.ID                 `json:"sku_id"`
	ApplicationID     snowflake.ID                 `json:"application_id"`
	UserID            snowflake.ID                 `json:"user_id"`
	User              PartialUser                  `json:"user,omitzero"`
	GuildID           snowflake.ID                 `json:"guild_id,omitzero"`
	ParentID          snowflake.ID                 `json:"parent_id,omitzero"`
	Deleted           bool                         `json:"deleted"`
	Consumed          bool                         `json:"consumed,omitzero"`
	Branches          []snowflake.ID               `json:"branches,omitempty"`
	StartsAt          *time.Time                   `json:"starts_at"`
	EndsAt            *time.Time                   `json:"ends_at"`
	PromotionID       *snowflake.ID                `json:"promotion_id"`
	SubscriptionID    snowflake.ID                 `json:"subscription_id,omitzero"`
	GiftCodeFlags     GiftCodeFlags                `json:"gift_code_flags"`
	GiftCodeBatchID   snowflake.ID                 `json:"gift_code_batch_id,omitzero"`
	GifterUserID      snowflake.ID                 `json:"gifter_user_id,omitzero"`
	GiftStyle         GiftStyle                    `json:"gift_style,omitzero"`
	FulfillmentStatus EntitlementFulfillmentStatus `json:"fulfillment_status,omitzero"`
	FulfilledAt       time.Time                    `json:"fulfilled_at,omitzero"`
	SourceType        EntitlementSourceType        `json:"source_type,omitzero"`
	TenantMetadata    TenantMetadata               `json:"tenant_metadata,omitzero"`
	SKU               SKU                          `json:"sku,omitzero"`
	SubscriptionPlan  PartialSubscriptionPlan      `json:"subscription_plan,omitzero"`
}

func (e *EntitlementDeleteEvent) AsEntitlement() Entitlement {
	return Entitlement{
		ID:                e.ID,
		Type:              e.Type,
		SKUID:             e.SKUID,
		ApplicationID:     e.ApplicationID,
		UserID:            e.UserID,
		User:              e.User,
		GuildID:           e.GuildID,
		ParentID:          e.ParentID,
		Deleted:           e.Deleted,
		Consumed:          e.Consumed,
		Branches:          e.Branches,
		StartsAt:          e.StartsAt,
		EndsAt:            e.EndsAt,
		PromotionID:       e.PromotionID,
		SubscriptionID:    e.SubscriptionID,
		GiftCodeFlags:     e.GiftCodeFlags,
		GiftCodeBatchID:   e.GiftCodeBatchID,
		GifterUserID:      e.GifterUserID,
		GiftStyle:         e.GiftStyle,
		FulfillmentStatus: e.FulfillmentStatus,
		FulfilledAt:       e.FulfilledAt,
		SourceType:        e.SourceType,
		TenantMetadata:    e.TenantMetadata,
		SKU:               e.SKU,
		SubscriptionPlan:  e.SubscriptionPlan,
	}
}

type ExperimentSessionOverrideCreateEvent struct {
	ExperimentName string `json:"experiment_name"`
	VariantID      int64  `json:"variant_id"`
}

func (e *ExperimentSessionOverrideCreateEvent) AsExperimentSessionOverrideCreate() ExperimentSessionOverrideCreate {
	return ExperimentSessionOverrideCreate{
		ExperimentName: e.ExperimentName,
		VariantID:      e.VariantID,
	}
}

type ExperimentSessionOverrideDeleteEvent struct {
	ExperimentName string `json:"experiment_name"`
}

func (e *ExperimentSessionOverrideDeleteEvent) AsExperimentSessionOverrideDelete() ExperimentSessionOverrideDelete {
	return ExperimentSessionOverrideDelete{
		ExperimentName: e.ExperimentName,
	}
}

type FriendSuggestionCreateEvent struct {
	SuggestedUser             PartialUser              `json:"suggested_user"`
	Reasons                   []FriendSuggestionReason `json:"reasons"`
	FromSuggestedUserContacts bool                     `json:"from_suggested_user_contacts,omitzero"`
	MutualFriendsCount        int64                    `json:"mutual_friends_count,omitzero"`
	ContactNames              []string                 `json:"contact_names,omitempty"`
}

func (e *FriendSuggestionCreateEvent) AsFriendSuggestion() FriendSuggestion {
	return FriendSuggestion{
		SuggestedUser:             e.SuggestedUser,
		Reasons:                   e.Reasons,
		FromSuggestedUserContacts: e.FromSuggestedUserContacts,
		MutualFriendsCount:        e.MutualFriendsCount,
		ContactNames:              e.ContactNames,
	}
}

type FriendSuggestionDeleteEvent struct {
	SuggestedUserID snowflake.ID `json:"suggested_user_id"`
}

func (e *FriendSuggestionDeleteEvent) AsFriendSuggestionDelete() FriendSuggestionDelete {
	return FriendSuggestionDelete{
		SuggestedUserID: e.SuggestedUserID,
	}
}

type GameServerCreateEvent struct {
	GuildID    snowflake.ID `json:"guild_id"`
	GameServer GameServer   `json:"game_server"`
}

func (e *GameServerCreateEvent) AsGameServerCreate() GameServerCreate {
	return GameServerCreate{
		GuildID:    e.GuildID,
		GameServer: e.GameServer,
	}
}

type GameServerUpdateEvent struct {
	GuildID    snowflake.ID `json:"guild_id"`
	GameServer GameServer   `json:"game_server"`
}

func (e *GameServerUpdateEvent) AsGameServerUpdate() GameServerUpdate {
	return GameServerUpdate{
		GuildID:    e.GuildID,
		GameServer: e.GameServer,
	}
}

type GameServerDeleteEvent struct {
	GuildID      snowflake.ID `json:"guild_id"`
	GameServerID snowflake.ID `json:"game_server_id"`
}

func (e *GameServerDeleteEvent) AsGameServerDelete() GameServerDelete {
	return GameServerDelete{
		GuildID:      e.GuildID,
		GameServerID: e.GameServerID,
	}
}

type GiftCodeCreateEvent struct {
	Code      string       `json:"code"`
	SKUID     snowflake.ID `json:"sku_id"`
	Uses      int64        `json:"uses"`
	ChannelID snowflake.ID `json:"channel_id,omitzero"`
	GuildID   snowflake.ID `json:"guild_id,omitzero"`
}

func (e *GiftCodeCreateEvent) AsGatewayGiftCode() GatewayGiftCode {
	return GatewayGiftCode{
		Code:      e.Code,
		SKUID:     e.SKUID,
		Uses:      e.Uses,
		ChannelID: e.ChannelID,
		GuildID:   e.GuildID,
	}
}

type GiftCodeUpdateEvent struct {
	Code      string       `json:"code"`
	SKUID     snowflake.ID `json:"sku_id"`
	Uses      int64        `json:"uses"`
	ChannelID snowflake.ID `json:"channel_id,omitzero"`
	GuildID   snowflake.ID `json:"guild_id,omitzero"`
}

func (e *GiftCodeUpdateEvent) AsGatewayGiftCode() GatewayGiftCode {
	return GatewayGiftCode{
		Code:      e.Code,
		SKUID:     e.SKUID,
		Uses:      e.Uses,
		ChannelID: e.ChannelID,
		GuildID:   e.GuildID,
	}
}

type GuildCreateEvent struct {
	JoinedAt                 time.Time                  `json:"joined_at"`
	Large                    bool                       `json:"large"`
	Unavailable              bool                       `json:"unavailable,omitzero"`
	GeoRestricted            bool                       `json:"geo_restricted,omitzero"`
	MemberCount              int64                      `json:"member_count"`
	Members                  []GuildMember              `json:"members"`
	Channels                 []Channel                  `json:"channels"`
	Threads                  []Channel                  `json:"threads"`
	Presences                []Presence                 `json:"presences"`
	VoiceStates              []VoiceState               `json:"voice_states"`
	ActivityInstances        []EmbeddedActivityInstance `json:"activity_instances"`
	StageInstances           []StageInstance            `json:"stage_instances"`
	GuildScheduledEvents     []GuildScheduledEvent      `json:"guild_scheduled_events"`
	DataMode                 DataMode                   `json:"data_mode"`
	PartialUpdates           PartialGuildUpdates        `json:"partial_updates,omitzero"`
	ChannelUpdates           []ChannelUnread            `json:"channel_updates,omitempty"`
	UnableToSyncDeletes      bool                       `json:"unable_to_sync_deletes,omitzero"`
	Properties               PartialGuild               `json:"properties"`
	Stickers                 []Sticker                  `json:"stickers"`
	Roles                    []Role                     `json:"roles"`
	Emojis                   []Emoji                    `json:"emojis"`
	SoundboardSounds         []SoundboardSound          `json:"soundboard_sounds"`
	PremiumSubscriptionCount int64                      `json:"premium_subscription_count"`
	ID                       snowflake.ID               `json:"id"`
}

func (e *GuildCreateEvent) AsGatewayGuild() GatewayGuild {
	return GatewayGuild{
		JoinedAt:                 e.JoinedAt,
		Large:                    e.Large,
		Unavailable:              e.Unavailable,
		GeoRestricted:            e.GeoRestricted,
		MemberCount:              e.MemberCount,
		Members:                  e.Members,
		Channels:                 e.Channels,
		Threads:                  e.Threads,
		Presences:                e.Presences,
		VoiceStates:              e.VoiceStates,
		ActivityInstances:        e.ActivityInstances,
		StageInstances:           e.StageInstances,
		GuildScheduledEvents:     e.GuildScheduledEvents,
		DataMode:                 e.DataMode,
		PartialUpdates:           e.PartialUpdates,
		ChannelUpdates:           e.ChannelUpdates,
		UnableToSyncDeletes:      e.UnableToSyncDeletes,
		Properties:               e.Properties,
		Stickers:                 e.Stickers,
		Roles:                    e.Roles,
		Emojis:                   e.Emojis,
		SoundboardSounds:         e.SoundboardSounds,
		PremiumSubscriptionCount: e.PremiumSubscriptionCount,
		ID:                       e.ID,
	}
}

type GuildUpdateEvent struct {
	ID                                     snowflake.ID                    `json:"id"`
	Name                                   string                          `json:"name"`
	Icon                                   *string                         `json:"icon"`
	Banner                                 *string                         `json:"banner"`
	HomeHeader                             *string                         `json:"home_header"`
	Splash                                 *string                         `json:"splash"`
	DiscoverySplash                        *string                         `json:"discovery_splash"`
	OwnerID                                snowflake.ID                    `json:"owner_id"`
	ApplicationID                          *snowflake.ID                   `json:"application_id"`
	Description                            *string                         `json:"description"`
	Region                                 *string                         `json:"region,omitempty"`
	AFKChannelID                           *snowflake.ID                   `json:"afk_channel_id"`
	AFKTimeout                             int64                           `json:"afk_timeout"`
	WidgetEnabled                          bool                            `json:"widget_enabled,omitzero"`
	WidgetChannelID                        *snowflake.ID                   `json:"widget_channel_id,omitempty"`
	VerificationLevel                      VerificationLevel               `json:"verification_level"`
	DefaultMessageNotifications            MessageNotificationLevel        `json:"default_message_notifications"`
	ExplicitContentFilter                  GuildExplicitContentFilterLevel `json:"explicit_content_filter"`
	Features                               []string                        `json:"features"`
	Roles                                  []Role                          `json:"roles"`
	Emojis                                 []Emoji                         `json:"emojis"`
	Stickers                               []Sticker                       `json:"stickers"`
	MFALevel                               MFALevel                        `json:"mfa_level"`
	SystemChannelID                        *snowflake.ID                   `json:"system_channel_id"`
	SystemChannelFlags                     SystemChannelFlags              `json:"system_channel_flags"`
	RulesChannelID                         *snowflake.ID                   `json:"rules_channel_id"`
	PublicUpdatesChannelID                 *snowflake.ID                   `json:"public_updates_channel_id"`
	SafetyAlertsChannelID                  *snowflake.ID                   `json:"safety_alerts_channel_id"`
	MaxPresences                           *int64                          `json:"max_presences,omitempty"`
	MaxMembers                             int64                           `json:"max_members,omitzero"`
	VanityURLCode                          *string                         `json:"vanity_url_code"`
	PremiumTier                            PremiumTier                     `json:"premium_tier"`
	PremiumSubscriptionCount               int64                           `json:"premium_subscription_count"`
	PreferredLocale                        string                          `json:"preferred_locale"`
	MaxVideoChannelUsers                   int64                           `json:"max_video_channel_users,omitzero"`
	MaxStageVideoChannelUsers              int64                           `json:"max_stage_video_channel_users,omitzero"`
	NSFW                                   bool                            `json:"nsfw"`
	NSFWLevel                              NSFWLevel                       `json:"nsfw_level"`
	OwnerConfiguredContentLevel            *NSFWLevel                      `json:"owner_configured_content_level"`
	HubType                                *HubType                        `json:"hub_type"`
	PremiumProgressBarEnabled              bool                            `json:"premium_progress_bar_enabled"`
	PremiumProgressBarEnabledUserUpdatedAt *time.Time                      `json:"premium_progress_bar_enabled_user_updated_at"`
	LatestOnboardingQuestionID             *snowflake.ID                   `json:"latest_onboarding_question_id"`
	IncidentsData                          *AutomodIncidentsData           `json:"incidents_data"`
	InventorySettings                      *GuildInventorySettings         `json:"inventory_settings"`
	PremiumFeatures                        *GuildPremiumFeatures           `json:"premium_features"`
	Profile                                *GuildIdentity                  `json:"profile"`
	ApproximateMemberCount                 int64                           `json:"approximate_member_count,omitzero"`
	ApproximatePresenceCount               int64                           `json:"approximate_presence_count,omitzero"`
	OfficialMessageColor                   *int64                          `json:"official_message_color"`
	Version                                string                          `json:"version,omitzero"`
}

func (e *GuildUpdateEvent) AsGuild() Guild {
	return Guild{
		ID:                                     e.ID,
		Name:                                   e.Name,
		Icon:                                   e.Icon,
		Banner:                                 e.Banner,
		HomeHeader:                             e.HomeHeader,
		Splash:                                 e.Splash,
		DiscoverySplash:                        e.DiscoverySplash,
		OwnerID:                                e.OwnerID,
		ApplicationID:                          e.ApplicationID,
		Description:                            e.Description,
		Region:                                 e.Region,
		AFKChannelID:                           e.AFKChannelID,
		AFKTimeout:                             e.AFKTimeout,
		WidgetEnabled:                          e.WidgetEnabled,
		WidgetChannelID:                        e.WidgetChannelID,
		VerificationLevel:                      e.VerificationLevel,
		DefaultMessageNotifications:            e.DefaultMessageNotifications,
		ExplicitContentFilter:                  e.ExplicitContentFilter,
		Features:                               e.Features,
		Roles:                                  e.Roles,
		Emojis:                                 e.Emojis,
		Stickers:                               e.Stickers,
		MFALevel:                               e.MFALevel,
		SystemChannelID:                        e.SystemChannelID,
		SystemChannelFlags:                     e.SystemChannelFlags,
		RulesChannelID:                         e.RulesChannelID,
		PublicUpdatesChannelID:                 e.PublicUpdatesChannelID,
		SafetyAlertsChannelID:                  e.SafetyAlertsChannelID,
		MaxPresences:                           e.MaxPresences,
		MaxMembers:                             e.MaxMembers,
		VanityURLCode:                          e.VanityURLCode,
		PremiumTier:                            e.PremiumTier,
		PremiumSubscriptionCount:               e.PremiumSubscriptionCount,
		PreferredLocale:                        e.PreferredLocale,
		MaxVideoChannelUsers:                   e.MaxVideoChannelUsers,
		MaxStageVideoChannelUsers:              e.MaxStageVideoChannelUsers,
		NSFW:                                   e.NSFW,
		NSFWLevel:                              e.NSFWLevel,
		OwnerConfiguredContentLevel:            e.OwnerConfiguredContentLevel,
		HubType:                                e.HubType,
		PremiumProgressBarEnabled:              e.PremiumProgressBarEnabled,
		PremiumProgressBarEnabledUserUpdatedAt: e.PremiumProgressBarEnabledUserUpdatedAt,
		LatestOnboardingQuestionID:             e.LatestOnboardingQuestionID,
		IncidentsData:                          e.IncidentsData,
		InventorySettings:                      e.InventorySettings,
		PremiumFeatures:                        e.PremiumFeatures,
		Profile:                                e.Profile,
		ApproximateMemberCount:                 e.ApproximateMemberCount,
		ApproximatePresenceCount:               e.ApproximatePresenceCount,
		OfficialMessageColor:                   e.OfficialMessageColor,
		Version:                                e.Version,
	}
}

type GuildDeleteEvent struct {
	ID            snowflake.ID `json:"id"`
	Unavailable   bool         `json:"unavailable,omitzero"`
	GeoRestricted bool         `json:"geo_restricted,omitzero"`
	Name          string       `json:"name,omitzero"`
	Icon          *string      `json:"icon,omitempty"`
}

func (e *GuildDeleteEvent) AsUnavailableGuild() UnavailableGuild {
	return UnavailableGuild{
		ID:            e.ID,
		Unavailable:   e.Unavailable,
		GeoRestricted: e.GeoRestricted,
		Name:          e.Name,
		Icon:          e.Icon,
	}
}

type GuildApplicationCommandIndexUpdateEvent struct {
	GuildID snowflake.ID `json:"guild_id"`
}

func (e *GuildApplicationCommandIndexUpdateEvent) AsGuildApplicationCommandIndexUpdate() GuildApplicationCommandIndexUpdate {
	return GuildApplicationCommandIndexUpdate{
		GuildID: e.GuildID,
	}
}

type GuildAppliedBoostsUpdateEvent struct {
	ID          snowflake.ID `json:"id"`
	GuildID     snowflake.ID `json:"guild_id"`
	UserID      snowflake.ID `json:"user_id"`
	Ended       bool         `json:"ended"`
	EndsAt      time.Time    `json:"ends_at,omitzero"`
	PauseEndsAt *time.Time   `json:"pause_ends_at"`
	User        PartialUser  `json:"user"`
}

func (e *GuildAppliedBoostsUpdateEvent) AsPremiumGuildSubscription() PremiumGuildSubscription {
	return PremiumGuildSubscription{
		ID:          e.ID,
		GuildID:     e.GuildID,
		UserID:      e.UserID,
		Ended:       e.Ended,
		EndsAt:      e.EndsAt,
		PauseEndsAt: e.PauseEndsAt,
		User:        e.User,
	}
}

type GuildAuditLogEntryCreateEvent struct {
	TargetID   *snowflake.ID      `json:"target_id"`
	Changes    []AuditLogChange   `json:"changes,omitempty"`
	UserID     *snowflake.ID      `json:"user_id"`
	ID         snowflake.ID       `json:"id"`
	ActionType AuditLogActionType `json:"action_type"`
	Options    json.RawMessage    `json:"options,omitempty"`
	Reason     string             `json:"reason,omitzero"`
}

func (e *GuildAuditLogEntryCreateEvent) AsAuditLogEntry() AuditLogEntry {
	return AuditLogEntry{
		TargetID:   e.TargetID,
		Changes:    e.Changes,
		UserID:     e.UserID,
		ID:         e.ID,
		ActionType: e.ActionType,
		Options:    e.Options,
		Reason:     e.Reason,
	}
}

type GuildBanAddEvent struct {
	GuildID           snowflake.ID `json:"guild_id"`
	User              PartialUser  `json:"user"`
	DeleteMessageSecs int64        `json:"delete_message_secs"`
}

func (e *GuildBanAddEvent) AsGuildBanAdd() GuildBanAdd {
	return GuildBanAdd{
		GuildID:           e.GuildID,
		User:              e.User,
		DeleteMessageSecs: e.DeleteMessageSecs,
	}
}

type GuildBanRemoveEvent struct {
	GuildID snowflake.ID `json:"guild_id"`
	User    PartialUser  `json:"user"`
}

func (e *GuildBanRemoveEvent) AsGuildBanRemove() GuildBanRemove {
	return GuildBanRemove{
		GuildID: e.GuildID,
		User:    e.User,
	}
}

type GuildBulkBanUpdateEvent struct {
	GuildID snowflake.ID `json:"guild_id"`
	BulkBan GuildBulkBan `json:"bulk_ban"`
}

func (e *GuildBulkBanUpdateEvent) AsGuildBulkBanUpdate() GuildBulkBanUpdate {
	return GuildBulkBanUpdate{
		GuildID: e.GuildID,
		BulkBan: e.BulkBan,
	}
}

type GuildDirectoryEntryCreateEvent struct {
	Type                DirectoryEntryType           `json:"type"`
	DirectoryChannelID  snowflake.ID                 `json:"directory_channel_id"`
	EntityID            snowflake.ID                 `json:"entity_id"`
	CreatedAt           string                       `json:"created_at"`
	PrimaryCategoryID   DirectoryCategory            `json:"primary_category_id,omitzero"`
	Description         *string                      `json:"description"`
	AuthorID            snowflake.ID                 `json:"author_id"`
	Guild               DirectoryGuild               `json:"guild,omitzero"`
	GuildScheduledEvent DirectoryGuildScheduledEvent `json:"guild_scheduled_event,omitzero"`
}

func (e *GuildDirectoryEntryCreateEvent) AsDirectoryEntry() DirectoryEntry {
	return DirectoryEntry{
		Type:                e.Type,
		DirectoryChannelID:  e.DirectoryChannelID,
		EntityID:            e.EntityID,
		CreatedAt:           e.CreatedAt,
		PrimaryCategoryID:   e.PrimaryCategoryID,
		Description:         e.Description,
		AuthorID:            e.AuthorID,
		Guild:               e.Guild,
		GuildScheduledEvent: e.GuildScheduledEvent,
	}
}

type GuildDirectoryEntryUpdateEvent struct {
	Type                DirectoryEntryType           `json:"type"`
	DirectoryChannelID  snowflake.ID                 `json:"directory_channel_id"`
	EntityID            snowflake.ID                 `json:"entity_id"`
	CreatedAt           string                       `json:"created_at"`
	PrimaryCategoryID   DirectoryCategory            `json:"primary_category_id,omitzero"`
	Description         *string                      `json:"description"`
	AuthorID            snowflake.ID                 `json:"author_id"`
	Guild               DirectoryGuild               `json:"guild,omitzero"`
	GuildScheduledEvent DirectoryGuildScheduledEvent `json:"guild_scheduled_event,omitzero"`
}

func (e *GuildDirectoryEntryUpdateEvent) AsDirectoryEntry() DirectoryEntry {
	return DirectoryEntry{
		Type:                e.Type,
		DirectoryChannelID:  e.DirectoryChannelID,
		EntityID:            e.EntityID,
		CreatedAt:           e.CreatedAt,
		PrimaryCategoryID:   e.PrimaryCategoryID,
		Description:         e.Description,
		AuthorID:            e.AuthorID,
		Guild:               e.Guild,
		GuildScheduledEvent: e.GuildScheduledEvent,
	}
}

type GuildDirectoryEntryDeleteEvent struct {
	Type               DirectoryEntryType `json:"type"`
	DirectoryChannelID snowflake.ID       `json:"directory_channel_id"`
	GuildID            snowflake.ID       `json:"guild_id"`
	EntityID           snowflake.ID       `json:"entity_id"`
	CreatedAt          string             `json:"created_at"`
	PrimaryCategoryID  DirectoryCategory  `json:"primary_category_id,omitzero"`
	Description        *string            `json:"description"`
	AuthorID           snowflake.ID       `json:"author_id"`
}

func (e *GuildDirectoryEntryDeleteEvent) AsGuildDirectoryEntryDelete() GuildDirectoryEntryDelete {
	return GuildDirectoryEntryDelete{
		Type:               e.Type,
		DirectoryChannelID: e.DirectoryChannelID,
		GuildID:            e.GuildID,
		EntityID:           e.EntityID,
		CreatedAt:          e.CreatedAt,
		PrimaryCategoryID:  e.PrimaryCategoryID,
		Description:        e.Description,
		AuthorID:           e.AuthorID,
	}
}

type GuildEmojisUpdateEvent struct {
	GuildID snowflake.ID `json:"guild_id"`
	Emojis  []Emoji      `json:"emojis"`
}

func (e *GuildEmojisUpdateEvent) AsGuildEmojisUpdate() GuildEmojisUpdate {
	return GuildEmojisUpdate{
		GuildID: e.GuildID,
		Emojis:  e.Emojis,
	}
}

type GuildStickersUpdateEvent struct {
	GuildID  snowflake.ID `json:"guild_id"`
	Stickers []Sticker    `json:"stickers"`
}

func (e *GuildStickersUpdateEvent) AsGuildStickersUpdate() GuildStickersUpdate {
	return GuildStickersUpdate{
		GuildID:  e.GuildID,
		Stickers: e.Stickers,
	}
}

type GuildFeatureAckEvent struct {
	AckType    ReadStateType `json:"ack_type"`
	ResourceID snowflake.ID  `json:"resource_id"`
	EntityID   snowflake.ID  `json:"entity_id"`
	Version    int64         `json:"version"`
}

func (e *GuildFeatureAckEvent) AsGuildFeatureAck() GuildFeatureAck {
	return GuildFeatureAck{
		AckType:    e.AckType,
		ResourceID: e.ResourceID,
		EntityID:   e.EntityID,
		Version:    e.Version,
	}
}

type GuildJoinRequestCreateEvent struct {
	GuildID snowflake.ID           `json:"guild_id"`
	Request GuildJoinRequest       `json:"request"`
	Status  GuildJoinRequestStatus `json:"status"`
}

func (e *GuildJoinRequestCreateEvent) AsGuildJoinRequestCreate() GuildJoinRequestCreate {
	return GuildJoinRequestCreate{
		GuildID: e.GuildID,
		Request: e.Request,
		Status:  e.Status,
	}
}

type GuildJoinRequestUpdateEvent struct {
	GuildID snowflake.ID           `json:"guild_id"`
	Request GuildJoinRequest       `json:"request"`
	Status  GuildJoinRequestStatus `json:"status"`
}

func (e *GuildJoinRequestUpdateEvent) AsGuildJoinRequestUpdate() GuildJoinRequestUpdate {
	return GuildJoinRequestUpdate{
		GuildID: e.GuildID,
		Request: e.Request,
		Status:  e.Status,
	}
}

type GuildJoinRequestDeleteEvent struct {
	GuildID snowflake.ID `json:"guild_id"`
	ID      snowflake.ID `json:"id"`
	UserID  snowflake.ID `json:"user_id"`
}

func (e *GuildJoinRequestDeleteEvent) AsGuildJoinRequestDelete() GuildJoinRequestDelete {
	return GuildJoinRequestDelete{
		GuildID: e.GuildID,
		ID:      e.ID,
		UserID:  e.UserID,
	}
}

type GuildMemberAddEvent struct {
	User                       PartialUser           `json:"user"`
	Nick                       *string               `json:"nick,omitempty"`
	Avatar                     *string               `json:"avatar,omitempty"`
	AvatarDecorationData       *AvatarDecorationData `json:"avatar_decoration_data,omitempty"`
	Collectibles               *GuildCollectibles    `json:"collectibles,omitempty"`
	DisplayNameStyles          *DisplayNameStyle     `json:"display_name_styles,omitempty"`
	Banner                     *string               `json:"banner,omitempty"`
	Bio                        *string               `json:"bio,omitempty"`
	Roles                      []snowflake.ID        `json:"roles"`
	JoinedAt                   time.Time             `json:"joined_at"`
	PremiumSince               *time.Time            `json:"premium_since,omitempty"`
	Deaf                       bool                  `json:"deaf,omitzero"`
	Mute                       bool                  `json:"mute,omitzero"`
	Pending                    bool                  `json:"pending,omitzero"`
	CommunicationDisabledUntil *time.Time            `json:"communication_disabled_until,omitempty"`
	UnusualDMActivityUntil     *time.Time            `json:"unusual_dm_activity_until,omitempty"`
	Flags                      GuildMemberFlags      `json:"flags"`
	Permissions                Permissions           `json:"permissions,omitzero"`
	UserID                     snowflake.ID          `json:"user_id,omitzero"`
	GuildID                    snowflake.ID          `json:"guild_id"`
}

func (e *GuildMemberAddEvent) AsGuildMember() GuildMember {
	return GuildMember{
		User:                       e.User,
		Nick:                       e.Nick,
		Avatar:                     e.Avatar,
		AvatarDecorationData:       e.AvatarDecorationData,
		Collectibles:               e.Collectibles,
		DisplayNameStyles:          e.DisplayNameStyles,
		Banner:                     e.Banner,
		Bio:                        e.Bio,
		Roles:                      e.Roles,
		JoinedAt:                   e.JoinedAt,
		PremiumSince:               e.PremiumSince,
		Deaf:                       e.Deaf,
		Mute:                       e.Mute,
		Pending:                    e.Pending,
		CommunicationDisabledUntil: e.CommunicationDisabledUntil,
		UnusualDMActivityUntil:     e.UnusualDMActivityUntil,
		Flags:                      e.Flags,
		Permissions:                e.Permissions,
		UserID:                     e.UserID,
	}
}

type GuildMemberUpdateEvent struct {
	GuildID                    snowflake.ID          `json:"guild_id"`
	User                       PartialUser           `json:"user"`
	Nick                       *string               `json:"nick,omitempty"`
	Avatar                     *string               `json:"avatar"`
	AvatarDecorationData       *AvatarDecorationData `json:"avatar_decoration_data,omitempty"`
	Roles                      []snowflake.ID        `json:"roles"`
	JoinedAt                   time.Time             `json:"joined_at"`
	PremiumSince               *time.Time            `json:"premium_since"`
	Deaf                       bool                  `json:"deaf,omitzero"`
	Mute                       bool                  `json:"mute,omitzero"`
	Pending                    bool                  `json:"pending,omitzero"`
	CommunicationDisabledUntil *time.Time            `json:"communication_disabled_until"`
	Flags                      GuildMemberFlags      `json:"flags"`
}

func (e *GuildMemberUpdateEvent) AsGuildMemberUpdate() GuildMemberUpdate {
	return GuildMemberUpdate{
		GuildID:                    e.GuildID,
		User:                       e.User,
		Nick:                       e.Nick,
		Avatar:                     e.Avatar,
		AvatarDecorationData:       e.AvatarDecorationData,
		Roles:                      e.Roles,
		JoinedAt:                   e.JoinedAt,
		PremiumSince:               e.PremiumSince,
		Deaf:                       e.Deaf,
		Mute:                       e.Mute,
		Pending:                    e.Pending,
		CommunicationDisabledUntil: e.CommunicationDisabledUntil,
		Flags:                      e.Flags,
	}
}

type GuildMemberRemoveEvent struct {
	GuildID snowflake.ID `json:"guild_id"`
	User    User         `json:"user"`
}

func (e *GuildMemberRemoveEvent) AsGuildMemberRemove() GuildMemberRemove {
	return GuildMemberRemove{
		GuildID: e.GuildID,
		User:    e.User,
	}
}

type GuildMembersChunkEvent struct {
	GuildID    snowflake.ID    `json:"guild_id"`
	Members    []GuildMember   `json:"members"`
	ChunkIndex int64           `json:"chunk_index"`
	ChunkCount int64           `json:"chunk_count"`
	NotFound   json.RawMessage `json:"not_found,omitempty"`
	Presences  []Presence      `json:"presences,omitempty"`
	Nonce      string          `json:"nonce,omitzero"`
}

func (e *GuildMembersChunkEvent) AsGuildMembersChunk() GuildMembersChunk {
	return GuildMembersChunk{
		GuildID:    e.GuildID,
		Members:    e.Members,
		ChunkIndex: e.ChunkIndex,
		ChunkCount: e.ChunkCount,
		NotFound:   e.NotFound,
		Presences:  e.Presences,
		Nonce:      e.Nonce,
	}
}

type GuildOfficialGameApplicationsUpdateEvent struct {
	GuildID            snowflake.ID   `json:"guild_id"`
	GameApplicationIDs []snowflake.ID `json:"game_application_ids"`
}

func (e *GuildOfficialGameApplicationsUpdateEvent) AsGuildOfficialGameApplicationsUpdate() GuildOfficialGameApplicationsUpdate {
	return GuildOfficialGameApplicationsUpdate{
		GuildID:            e.GuildID,
		GameApplicationIDs: e.GameApplicationIDs,
	}
}

type GuildPowerupEntitlementsCreateEvent struct {
	GuildID      snowflake.ID  `json:"guild_id"`
	Entitlements []Entitlement `json:"entitlements"`
}

func (e *GuildPowerupEntitlementsCreateEvent) AsGuildPowerupEntitlementsCreate() GuildPowerupEntitlementsCreate {
	return GuildPowerupEntitlementsCreate{
		GuildID:      e.GuildID,
		Entitlements: e.Entitlements,
	}
}

type GuildPowerupEntitlementsDeleteEvent struct {
	GuildID      snowflake.ID  `json:"guild_id"`
	Entitlements []Entitlement `json:"entitlements"`
}

func (e *GuildPowerupEntitlementsDeleteEvent) AsGuildPowerupEntitlementsDelete() GuildPowerupEntitlementsDelete {
	return GuildPowerupEntitlementsDelete{
		GuildID:      e.GuildID,
		Entitlements: e.Entitlements,
	}
}

type GuildRoleCreateEvent struct {
	GuildID snowflake.ID `json:"guild_id"`
	Role    Role         `json:"role"`
}

func (e *GuildRoleCreateEvent) AsGuildRoleCreate() GuildRoleCreate {
	return GuildRoleCreate{
		GuildID: e.GuildID,
		Role:    e.Role,
	}
}

type GuildRoleUpdateEvent struct {
	GuildID snowflake.ID `json:"guild_id"`
	Role    Role         `json:"role"`
}

func (e *GuildRoleUpdateEvent) AsGuildRoleUpdate() GuildRoleUpdate {
	return GuildRoleUpdate{
		GuildID: e.GuildID,
		Role:    e.Role,
	}
}

type GuildRoleDeleteEvent struct {
	GuildID snowflake.ID `json:"guild_id"`
	RoleID  snowflake.ID `json:"role_id"`
}

func (e *GuildRoleDeleteEvent) AsGuildRoleDelete() GuildRoleDelete {
	return GuildRoleDelete{
		GuildID: e.GuildID,
		RoleID:  e.RoleID,
	}
}

type GuildScheduledEventCreateEvent struct {
	ID                            snowflake.ID                  `json:"id"`
	GuildID                       snowflake.ID                  `json:"guild_id"`
	ChannelID                     *snowflake.ID                 `json:"channel_id"`
	CreatorID                     *snowflake.ID                 `json:"creator_id,omitempty"`
	Creator                       PartialUser                   `json:"creator,omitzero"`
	Name                          string                        `json:"name"`
	Description                   *string                       `json:"description,omitempty"`
	ScheduledStartTime            time.Time                     `json:"scheduled_start_time"`
	ScheduledEndTime              *time.Time                    `json:"scheduled_end_time"`
	AutoStart                     bool                          `json:"auto_start,omitzero"`
	PrivacyLevel                  PrivacyLevel                  `json:"privacy_level"`
	Status                        GuildScheduledEventStatus     `json:"status"`
	EntityType                    GuildScheduledEventEntityType `json:"entity_type"`
	EntityID                      *snowflake.ID                 `json:"entity_id"`
	EntityMetadata                json.RawMessage               `json:"entity_metadata"`
	UserCount                     int64                         `json:"user_count,omitzero"`
	Image                         *string                       `json:"image,omitempty"`
	RecurrenceRule                json.RawMessage               `json:"recurrence_rule"`
	GuildScheduledEventExceptions []json.RawMessage             `json:"guild_scheduled_event_exceptions"`
}

func (e *GuildScheduledEventCreateEvent) AsGuildScheduledEvent() GuildScheduledEvent {
	return GuildScheduledEvent{
		ID:                            e.ID,
		GuildID:                       e.GuildID,
		ChannelID:                     e.ChannelID,
		CreatorID:                     e.CreatorID,
		Creator:                       e.Creator,
		Name:                          e.Name,
		Description:                   e.Description,
		ScheduledStartTime:            e.ScheduledStartTime,
		ScheduledEndTime:              e.ScheduledEndTime,
		AutoStart:                     e.AutoStart,
		PrivacyLevel:                  e.PrivacyLevel,
		Status:                        e.Status,
		EntityType:                    e.EntityType,
		EntityID:                      e.EntityID,
		EntityMetadata:                e.EntityMetadata,
		UserCount:                     e.UserCount,
		Image:                         e.Image,
		RecurrenceRule:                e.RecurrenceRule,
		GuildScheduledEventExceptions: e.GuildScheduledEventExceptions,
	}
}

type GuildScheduledEventUpdateEvent struct {
	ID                            snowflake.ID                  `json:"id"`
	GuildID                       snowflake.ID                  `json:"guild_id"`
	ChannelID                     *snowflake.ID                 `json:"channel_id"`
	CreatorID                     *snowflake.ID                 `json:"creator_id,omitempty"`
	Creator                       PartialUser                   `json:"creator,omitzero"`
	Name                          string                        `json:"name"`
	Description                   *string                       `json:"description,omitempty"`
	ScheduledStartTime            time.Time                     `json:"scheduled_start_time"`
	ScheduledEndTime              *time.Time                    `json:"scheduled_end_time"`
	AutoStart                     bool                          `json:"auto_start,omitzero"`
	PrivacyLevel                  PrivacyLevel                  `json:"privacy_level"`
	Status                        GuildScheduledEventStatus     `json:"status"`
	EntityType                    GuildScheduledEventEntityType `json:"entity_type"`
	EntityID                      *snowflake.ID                 `json:"entity_id"`
	EntityMetadata                json.RawMessage               `json:"entity_metadata"`
	UserCount                     int64                         `json:"user_count,omitzero"`
	Image                         *string                       `json:"image,omitempty"`
	RecurrenceRule                json.RawMessage               `json:"recurrence_rule"`
	GuildScheduledEventExceptions []json.RawMessage             `json:"guild_scheduled_event_exceptions"`
}

func (e *GuildScheduledEventUpdateEvent) AsGuildScheduledEvent() GuildScheduledEvent {
	return GuildScheduledEvent{
		ID:                            e.ID,
		GuildID:                       e.GuildID,
		ChannelID:                     e.ChannelID,
		CreatorID:                     e.CreatorID,
		Creator:                       e.Creator,
		Name:                          e.Name,
		Description:                   e.Description,
		ScheduledStartTime:            e.ScheduledStartTime,
		ScheduledEndTime:              e.ScheduledEndTime,
		AutoStart:                     e.AutoStart,
		PrivacyLevel:                  e.PrivacyLevel,
		Status:                        e.Status,
		EntityType:                    e.EntityType,
		EntityID:                      e.EntityID,
		EntityMetadata:                e.EntityMetadata,
		UserCount:                     e.UserCount,
		Image:                         e.Image,
		RecurrenceRule:                e.RecurrenceRule,
		GuildScheduledEventExceptions: e.GuildScheduledEventExceptions,
	}
}

type GuildScheduledEventDeleteEvent struct {
	ID                            snowflake.ID                  `json:"id"`
	GuildID                       snowflake.ID                  `json:"guild_id"`
	ChannelID                     *snowflake.ID                 `json:"channel_id"`
	CreatorID                     *snowflake.ID                 `json:"creator_id,omitempty"`
	Creator                       PartialUser                   `json:"creator,omitzero"`
	Name                          string                        `json:"name"`
	Description                   *string                       `json:"description,omitempty"`
	ScheduledStartTime            time.Time                     `json:"scheduled_start_time"`
	ScheduledEndTime              *time.Time                    `json:"scheduled_end_time"`
	AutoStart                     bool                          `json:"auto_start,omitzero"`
	PrivacyLevel                  PrivacyLevel                  `json:"privacy_level"`
	Status                        GuildScheduledEventStatus     `json:"status"`
	EntityType                    GuildScheduledEventEntityType `json:"entity_type"`
	EntityID                      *snowflake.ID                 `json:"entity_id"`
	EntityMetadata                json.RawMessage               `json:"entity_metadata"`
	UserCount                     int64                         `json:"user_count,omitzero"`
	Image                         *string                       `json:"image,omitempty"`
	RecurrenceRule                json.RawMessage               `json:"recurrence_rule"`
	GuildScheduledEventExceptions []json.RawMessage             `json:"guild_scheduled_event_exceptions"`
}

func (e *GuildScheduledEventDeleteEvent) AsGuildScheduledEvent() GuildScheduledEvent {
	return GuildScheduledEvent{
		ID:                            e.ID,
		GuildID:                       e.GuildID,
		ChannelID:                     e.ChannelID,
		CreatorID:                     e.CreatorID,
		Creator:                       e.Creator,
		Name:                          e.Name,
		Description:                   e.Description,
		ScheduledStartTime:            e.ScheduledStartTime,
		ScheduledEndTime:              e.ScheduledEndTime,
		AutoStart:                     e.AutoStart,
		PrivacyLevel:                  e.PrivacyLevel,
		Status:                        e.Status,
		EntityType:                    e.EntityType,
		EntityID:                      e.EntityID,
		EntityMetadata:                e.EntityMetadata,
		UserCount:                     e.UserCount,
		Image:                         e.Image,
		RecurrenceRule:                e.RecurrenceRule,
		GuildScheduledEventExceptions: e.GuildScheduledEventExceptions,
	}
}

type GuildScheduledEventExceptionCreateEvent struct {
	EventID            snowflake.ID `json:"event_id"`
	EventExceptionID   snowflake.ID `json:"event_exception_id"`
	IsCanceled         bool         `json:"is_canceled"`
	ScheduledStartTime *time.Time   `json:"scheduled_start_time"`
	ScheduledEndTime   *time.Time   `json:"scheduled_end_time"`
}

func (e *GuildScheduledEventExceptionCreateEvent) AsGuildScheduledEventException() GuildScheduledEventException {
	return GuildScheduledEventException{
		EventID:            e.EventID,
		EventExceptionID:   e.EventExceptionID,
		IsCanceled:         e.IsCanceled,
		ScheduledStartTime: e.ScheduledStartTime,
		ScheduledEndTime:   e.ScheduledEndTime,
	}
}

type GuildScheduledEventExceptionUpdateEvent struct {
	EventID            snowflake.ID `json:"event_id"`
	EventExceptionID   snowflake.ID `json:"event_exception_id"`
	IsCanceled         bool         `json:"is_canceled"`
	ScheduledStartTime *time.Time   `json:"scheduled_start_time"`
	ScheduledEndTime   *time.Time   `json:"scheduled_end_time"`
}

func (e *GuildScheduledEventExceptionUpdateEvent) AsGuildScheduledEventException() GuildScheduledEventException {
	return GuildScheduledEventException{
		EventID:            e.EventID,
		EventExceptionID:   e.EventExceptionID,
		IsCanceled:         e.IsCanceled,
		ScheduledStartTime: e.ScheduledStartTime,
		ScheduledEndTime:   e.ScheduledEndTime,
	}
}

type GuildScheduledEventExceptionDeleteEvent struct {
	EventID            snowflake.ID `json:"event_id"`
	EventExceptionID   snowflake.ID `json:"event_exception_id"`
	IsCanceled         bool         `json:"is_canceled"`
	ScheduledStartTime *time.Time   `json:"scheduled_start_time"`
	ScheduledEndTime   *time.Time   `json:"scheduled_end_time"`
}

func (e *GuildScheduledEventExceptionDeleteEvent) AsGuildScheduledEventException() GuildScheduledEventException {
	return GuildScheduledEventException{
		EventID:            e.EventID,
		EventExceptionID:   e.EventExceptionID,
		IsCanceled:         e.IsCanceled,
		ScheduledStartTime: e.ScheduledStartTime,
		ScheduledEndTime:   e.ScheduledEndTime,
	}
}

type GuildScheduledEventExceptionsDeleteEvent struct {
	GuildID snowflake.ID `json:"guild_id"`
	EventID snowflake.ID `json:"event_id"`
}

func (e *GuildScheduledEventExceptionsDeleteEvent) AsGuildScheduledEventExceptionsDelete() GuildScheduledEventExceptionsDelete {
	return GuildScheduledEventExceptionsDelete{
		GuildID: e.GuildID,
		EventID: e.EventID,
	}
}

type GuildScheduledEventUserAddEvent struct {
	GuildScheduledEventID          snowflake.ID                    `json:"guild_scheduled_event_id"`
	GuildScheduledEventExceptionID snowflake.ID                    `json:"guild_scheduled_event_exception_id,omitzero"`
	Response                       GuildScheduledEventUserResponse `json:"response"`
	UserID                         snowflake.ID                    `json:"user_id"`
	User                           PartialUser                     `json:"user,omitzero"`
	Member                         GuildMember                     `json:"member,omitzero"`
}

func (e *GuildScheduledEventUserAddEvent) AsGuildScheduledEventUser() GuildScheduledEventUser {
	return GuildScheduledEventUser{
		GuildScheduledEventID:          e.GuildScheduledEventID,
		GuildScheduledEventExceptionID: e.GuildScheduledEventExceptionID,
		Response:                       e.Response,
		UserID:                         e.UserID,
		User:                           e.User,
		Member:                         e.Member,
	}
}

type GuildScheduledEventUserRemoveEvent struct {
	GuildScheduledEventID          snowflake.ID                    `json:"guild_scheduled_event_id"`
	GuildScheduledEventExceptionID snowflake.ID                    `json:"guild_scheduled_event_exception_id,omitzero"`
	Response                       GuildScheduledEventUserResponse `json:"response"`
	UserID                         snowflake.ID                    `json:"user_id"`
	User                           PartialUser                     `json:"user,omitzero"`
	Member                         GuildMember                     `json:"member,omitzero"`
}

func (e *GuildScheduledEventUserRemoveEvent) AsGuildScheduledEventUser() GuildScheduledEventUser {
	return GuildScheduledEventUser{
		GuildScheduledEventID:          e.GuildScheduledEventID,
		GuildScheduledEventExceptionID: e.GuildScheduledEventExceptionID,
		Response:                       e.Response,
		UserID:                         e.UserID,
		User:                           e.User,
		Member:                         e.Member,
	}
}

type GuildSoundboardSoundCreateEvent struct {
	SoundID   snowflake.ID  `json:"sound_id"`
	Name      string        `json:"name"`
	Volume    float64       `json:"volume"`
	EmojiID   *snowflake.ID `json:"emoji_id"`
	EmojiName *string       `json:"emoji_name"`
	GuildID   snowflake.ID  `json:"guild_id,omitzero"`
	Available bool          `json:"available"`
	User      PartialUser   `json:"user,omitzero"`
	UserID    snowflake.ID  `json:"user_id,omitzero"`
}

func (e *GuildSoundboardSoundCreateEvent) AsSoundboardSound() SoundboardSound {
	return SoundboardSound{
		SoundID:   e.SoundID,
		Name:      e.Name,
		Volume:    e.Volume,
		EmojiID:   e.EmojiID,
		EmojiName: e.EmojiName,
		GuildID:   e.GuildID,
		Available: e.Available,
		User:      e.User,
		UserID:    e.UserID,
	}
}

type GuildSoundboardSoundUpdateEvent struct {
	SoundID   snowflake.ID  `json:"sound_id"`
	Name      string        `json:"name"`
	Volume    float64       `json:"volume"`
	EmojiID   *snowflake.ID `json:"emoji_id"`
	EmojiName *string       `json:"emoji_name"`
	GuildID   snowflake.ID  `json:"guild_id,omitzero"`
	Available bool          `json:"available"`
	User      PartialUser   `json:"user,omitzero"`
	UserID    snowflake.ID  `json:"user_id,omitzero"`
}

func (e *GuildSoundboardSoundUpdateEvent) AsSoundboardSound() SoundboardSound {
	return SoundboardSound{
		SoundID:   e.SoundID,
		Name:      e.Name,
		Volume:    e.Volume,
		EmojiID:   e.EmojiID,
		EmojiName: e.EmojiName,
		GuildID:   e.GuildID,
		Available: e.Available,
		User:      e.User,
		UserID:    e.UserID,
	}
}

type GuildSoundboardSoundDeleteEvent struct {
	GuildID snowflake.ID `json:"guild_id"`
	SoundID snowflake.ID `json:"sound_id"`
}

func (e *GuildSoundboardSoundDeleteEvent) AsGuildSoundboardSoundDelete() GuildSoundboardSoundDelete {
	return GuildSoundboardSoundDelete{
		GuildID: e.GuildID,
		SoundID: e.SoundID,
	}
}

type GuildSoundboardSoundsUpdateEvent struct {
	GuildID          snowflake.ID      `json:"guild_id"`
	SoundboardSounds []SoundboardSound `json:"soundboard_sounds"`
}

func (e *GuildSoundboardSoundsUpdateEvent) AsGuildSoundboardSoundsUpdate() GuildSoundboardSoundsUpdate {
	return GuildSoundboardSoundsUpdate{
		GuildID:          e.GuildID,
		SoundboardSounds: e.SoundboardSounds,
	}
}

type GuildIntegrationsUpdateEvent struct {
	GuildID snowflake.ID `json:"guild_id"`
}

func (e *GuildIntegrationsUpdateEvent) AsGuildIntegrationsUpdate() GuildIntegrationsUpdate {
	return GuildIntegrationsUpdate{
		GuildID: e.GuildID,
	}
}

type IntegrationCreateEvent struct {
	ID                      snowflake.ID                        `json:"id"`
	Name                    string                              `json:"name"`
	Type                    IntegrationType                     `json:"type"`
	Enabled                 bool                                `json:"enabled"`
	Account                 IntegrationAccount                  `json:"account"`
	Syncing                 bool                                `json:"syncing,omitzero"`
	RoleID                  *snowflake.ID                       `json:"role_id,omitempty"`
	EnableEmoticons         bool                                `json:"enable_emoticons,omitzero"`
	ExpireBehavior          IntegrationExpireBehavior           `json:"expire_behavior,omitzero"`
	ExpireGracePeriod       int64                               `json:"expire_grace_period,omitzero"`
	SyncedAt                time.Time                           `json:"synced_at,omitzero"`
	SubscriberCount         int64                               `json:"subscriber_count,omitzero"`
	Revoked                 bool                                `json:"revoked,omitzero"`
	Application             IntegrationApplication              `json:"application,omitzero"`
	Scopes                  []string                            `json:"scopes,omitempty"`
	RoleConnectionsMetadata []ApplicationRoleConnectionMetadata `json:"role_connections_metadata"`
	User                    PartialUser                         `json:"user,omitzero"`
	GuildID                 snowflake.ID                        `json:"guild_id"`
}

func (e *IntegrationCreateEvent) AsIntegration() Integration {
	return Integration{
		ID:                      e.ID,
		Name:                    e.Name,
		Type:                    e.Type,
		Enabled:                 e.Enabled,
		Account:                 e.Account,
		Syncing:                 e.Syncing,
		RoleID:                  e.RoleID,
		EnableEmoticons:         e.EnableEmoticons,
		ExpireBehavior:          e.ExpireBehavior,
		ExpireGracePeriod:       e.ExpireGracePeriod,
		SyncedAt:                e.SyncedAt,
		SubscriberCount:         e.SubscriberCount,
		Revoked:                 e.Revoked,
		Application:             e.Application,
		Scopes:                  e.Scopes,
		RoleConnectionsMetadata: e.RoleConnectionsMetadata,
		User:                    e.User,
	}
}

type IntegrationUpdateEvent struct {
	ID                      snowflake.ID                        `json:"id"`
	Name                    string                              `json:"name"`
	Type                    IntegrationType                     `json:"type"`
	Enabled                 bool                                `json:"enabled"`
	Account                 IntegrationAccount                  `json:"account"`
	Syncing                 bool                                `json:"syncing,omitzero"`
	RoleID                  *snowflake.ID                       `json:"role_id,omitempty"`
	EnableEmoticons         bool                                `json:"enable_emoticons,omitzero"`
	ExpireBehavior          IntegrationExpireBehavior           `json:"expire_behavior,omitzero"`
	ExpireGracePeriod       int64                               `json:"expire_grace_period,omitzero"`
	SyncedAt                time.Time                           `json:"synced_at,omitzero"`
	SubscriberCount         int64                               `json:"subscriber_count,omitzero"`
	Revoked                 bool                                `json:"revoked,omitzero"`
	Application             IntegrationApplication              `json:"application,omitzero"`
	Scopes                  []string                            `json:"scopes,omitempty"`
	RoleConnectionsMetadata []ApplicationRoleConnectionMetadata `json:"role_connections_metadata"`
	User                    PartialUser                         `json:"user,omitzero"`
	GuildID                 snowflake.ID                        `json:"guild_id"`
}

func (e *IntegrationUpdateEvent) AsIntegration() Integration {
	return Integration{
		ID:                      e.ID,
		Name:                    e.Name,
		Type:                    e.Type,
		Enabled:                 e.Enabled,
		Account:                 e.Account,
		Syncing:                 e.Syncing,
		RoleID:                  e.RoleID,
		EnableEmoticons:         e.EnableEmoticons,
		ExpireBehavior:          e.ExpireBehavior,
		ExpireGracePeriod:       e.ExpireGracePeriod,
		SyncedAt:                e.SyncedAt,
		SubscriberCount:         e.SubscriberCount,
		Revoked:                 e.Revoked,
		Application:             e.Application,
		Scopes:                  e.Scopes,
		RoleConnectionsMetadata: e.RoleConnectionsMetadata,
		User:                    e.User,
	}
}

type IntegrationDeleteEvent struct {
	ID            snowflake.ID `json:"id"`
	GuildID       snowflake.ID `json:"guild_id"`
	ApplicationID snowflake.ID `json:"application_id,omitzero"`
}

func (e *IntegrationDeleteEvent) AsIntegrationDelete() IntegrationDelete {
	return IntegrationDelete{
		ID:            e.ID,
		GuildID:       e.GuildID,
		ApplicationID: e.ApplicationID,
	}
}

type InteractionCreateEvent struct {
	ID    snowflake.ID `json:"id"`
	Nonce string       `json:"nonce,omitzero"`
}

func (e *InteractionCreateEvent) AsInteractionCreate() InteractionCreate {
	return InteractionCreate{
		ID:    e.ID,
		Nonce: e.Nonce,
	}
}

type InteractionFailureEvent struct {
	ID         snowflake.ID             `json:"id"`
	Nonce      string                   `json:"nonce,omitzero"`
	ReasonCode InteractionFailureReason `json:"reason_code"`
}

func (e *InteractionFailureEvent) AsInteractionFailure() InteractionFailure {
	return InteractionFailure{
		ID:         e.ID,
		Nonce:      e.Nonce,
		ReasonCode: e.ReasonCode,
	}
}

type InteractionSuccessEvent struct {
	ID    snowflake.ID `json:"id"`
	Nonce *string      `json:"nonce"`
}

func (e *InteractionSuccessEvent) AsInteractionSuccess() InteractionSuccess {
	return InteractionSuccess{
		ID:    e.ID,
		Nonce: e.Nonce,
	}
}

type ApplicationCommandAutocompleteResponseEvent struct {
	Choices []ApplicationCommandOptionChoice `json:"choices"`
	Nonce   *string                          `json:"nonce"`
}

func (e *ApplicationCommandAutocompleteResponseEvent) AsApplicationCommandAutocompleteResponse() ApplicationCommandAutocompleteResponse {
	return ApplicationCommandAutocompleteResponse{
		Choices: e.Choices,
		Nonce:   e.Nonce,
	}
}

type InteractionModalCreateEvent struct {
	ID          snowflake.ID           `json:"id"`
	ChannelID   snowflake.ID           `json:"channel_id"`
	CustomID    string                 `json:"custom_id"`
	Application IntegrationApplication `json:"application"`
	Title       string                 `json:"title"`
	Components  Components             `json:"components"`
	Nonce       *string                `json:"nonce"`
	Resolved    ResolvedData           `json:"resolved,omitzero"`
}

func (e *InteractionModalCreateEvent) AsInteractionModalCreate() InteractionModalCreate {
	return InteractionModalCreate{
		ID:          e.ID,
		ChannelID:   e.ChannelID,
		CustomID:    e.CustomID,
		Application: e.Application,
		Title:       e.Title,
		Components:  e.Components,
		Nonce:       e.Nonce,
		Resolved:    e.Resolved,
	}
}

type InteractionIframeModalCreateEvent struct {
	ID          snowflake.ID           `json:"id"`
	ChannelID   snowflake.ID           `json:"channel_id"`
	CustomID    string                 `json:"custom_id"`
	Application IntegrationApplication `json:"application"`
	Title       string                 `json:"title"`
	IframePath  string                 `json:"iframe_path"`
	ModalSize   IframeModalSize        `json:"modal_size"`
	Nonce       *string                `json:"nonce"`
}

func (e *InteractionIframeModalCreateEvent) AsInteractionIframeModalCreate() InteractionIframeModalCreate {
	return InteractionIframeModalCreate{
		ID:          e.ID,
		ChannelID:   e.ChannelID,
		CustomID:    e.CustomID,
		Application: e.Application,
		Title:       e.Title,
		IframePath:  e.IframePath,
		ModalSize:   e.ModalSize,
		Nonce:       e.Nonce,
	}
}

type SocialLayerSKUPurchaseEligibilityResponseEvent struct {
	InteractionID snowflake.ID `json:"interaction_id"`
	ApplicationID snowflake.ID `json:"application_id"`
	SKUID         snowflake.ID `json:"sku_id"`
	RecipientID   snowflake.ID `json:"recipient_id"`
	Eligible      snowflake.ID `json:"eligible"`
}

func (e *SocialLayerSKUPurchaseEligibilityResponseEvent) AsSocialLayerSKUPurchaseEligibilityResponse() SocialLayerSKUPurchaseEligibilityResponse {
	return SocialLayerSKUPurchaseEligibilityResponse{
		InteractionID: e.InteractionID,
		ApplicationID: e.ApplicationID,
		SKUID:         e.SKUID,
		RecipientID:   e.RecipientID,
		Eligible:      e.Eligible,
	}
}

type InviteCreateEvent struct {
	Code              string             `json:"code"`
	Type              InviteType         `json:"type"`
	ChannelID         snowflake.ID       `json:"channel_id"`
	GuildID           snowflake.ID       `json:"guild_id"`
	Inviter           PartialUser        `json:"inviter,omitzero"`
	TargetType        InviteTargetType   `json:"target_type,omitzero"`
	TargetUser        PartialUser        `json:"target_user,omitzero"`
	TargetApplication PartialApplication `json:"target_application,omitzero"`
	RoleIDs           []snowflake.ID     `json:"role_ids,omitempty"`
	ExpiresAt         *time.Time         `json:"expires_at"`
	CreatedAt         time.Time          `json:"created_at"`
	Uses              int64              `json:"uses"`
	MaxUses           int64              `json:"max_uses"`
	MaxAge            int64              `json:"max_age"`
	Temporary         bool               `json:"temporary"`
}

func (e *InviteCreateEvent) AsInviteCreate() InviteCreate {
	return InviteCreate{
		Code:              e.Code,
		Type:              e.Type,
		ChannelID:         e.ChannelID,
		GuildID:           e.GuildID,
		Inviter:           e.Inviter,
		TargetType:        e.TargetType,
		TargetUser:        e.TargetUser,
		TargetApplication: e.TargetApplication,
		RoleIDs:           e.RoleIDs,
		ExpiresAt:         e.ExpiresAt,
		CreatedAt:         e.CreatedAt,
		Uses:              e.Uses,
		MaxUses:           e.MaxUses,
		MaxAge:            e.MaxAge,
		Temporary:         e.Temporary,
	}
}

type InviteDeleteEvent struct {
	Code      string       `json:"code"`
	ChannelID snowflake.ID `json:"channel_id"`
	GuildID   snowflake.ID `json:"guild_id"`
}

func (e *InviteDeleteEvent) AsInviteDelete() InviteDelete {
	return InviteDelete{
		Code:      e.Code,
		ChannelID: e.ChannelID,
		GuildID:   e.GuildID,
	}
}

type MessageCreateEvent struct {
	ID                   snowflake.ID                `json:"id"`
	ChannelID            snowflake.ID                `json:"channel_id"`
	LobbyID              snowflake.ID                `json:"lobby_id,omitzero"`
	Author               PartialUser                 `json:"author"`
	Content              string                      `json:"content"`
	Timestamp            time.Time                   `json:"timestamp"`
	EditedTimestamp      *time.Time                  `json:"edited_timestamp"`
	TTS                  bool                        `json:"tts"`
	MentionEveryone      bool                        `json:"mention_everyone"`
	Mentions             []PartialUser               `json:"mentions"`
	MentionRoles         []snowflake.ID              `json:"mention_roles"`
	MentionChannels      []PartialChannel            `json:"mention_channels,omitempty"`
	Attachments          []Attachment                `json:"attachments"`
	Embeds               []Embed                     `json:"embeds"`
	Reactions            []Reaction                  `json:"reactions,omitempty"`
	Nonce                json.RawMessage             `json:"nonce,omitempty"`
	Pinned               bool                        `json:"pinned"`
	WebhookID            snowflake.ID                `json:"webhook_id,omitzero"`
	Type                 MessageType                 `json:"type"`
	Activity             MessageActivity             `json:"activity,omitzero"`
	Application          IntegrationApplication      `json:"application,omitzero"`
	ApplicationID        snowflake.ID                `json:"application_id,omitzero"`
	Flags                MessageFlags                `json:"flags"`
	MessageReference     MessageReference            `json:"message_reference,omitzero"`
	ReferencedMessage    *Message                    `json:"referenced_message,omitempty"`
	MessageSnapshots     []MessageSnapshot           `json:"message_snapshots,omitempty"`
	Call                 MessageCall                 `json:"call,omitzero"`
	Interaction          PartialMessageInteraction   `json:"interaction,omitzero"`
	InteractionMetadata  MessageInteraction          `json:"interaction_metadata,omitzero"`
	Resolved             ResolvedData                `json:"resolved,omitzero"`
	Thread               Channel                     `json:"thread,omitzero"`
	RoleSubscriptionData MessageRoleSubscription     `json:"role_subscription_data,omitzero"`
	PurchaseNotification MessagePurchaseNotification `json:"purchase_notification,omitzero"`
	GiftInfo             MessageGiftInfo             `json:"gift_info,omitzero"`
	Components           Components                  `json:"components"`
	StickerItems         []StickerItem               `json:"sticker_items,omitempty"`
	Stickers             []Sticker                   `json:"stickers,omitempty"`
	Poll                 Poll                        `json:"poll,omitzero"`
	ChangelogID          snowflake.ID                `json:"changelog_id,omitzero"`
	SoundboardSounds     []SoundboardSound           `json:"soundboard_sounds,omitempty"`
	Potions              []Potion                    `json:"potions,omitempty"`
	SharedClientTheme    SharedClientTheme           `json:"shared_client_theme,omitzero"`
	ChannelType          ChannelType                 `json:"channel_type"`
	GuildID              snowflake.ID                `json:"guild_id,omitzero"`
	Member               GuildMember                 `json:"member,omitzero"`
	Metadata             map[string]string           `json:"metadata,omitempty"`
	ModerationMetadata   map[string]string           `json:"moderation_metadata,omitempty"`
}

func (e *MessageCreateEvent) AsMessage() Message {
	return Message{
		ID:                   e.ID,
		ChannelID:            e.ChannelID,
		LobbyID:              e.LobbyID,
		Author:               e.Author,
		Content:              e.Content,
		Timestamp:            e.Timestamp,
		EditedTimestamp:      e.EditedTimestamp,
		TTS:                  e.TTS,
		MentionEveryone:      e.MentionEveryone,
		Mentions:             e.Mentions,
		MentionRoles:         e.MentionRoles,
		MentionChannels:      e.MentionChannels,
		Attachments:          e.Attachments,
		Embeds:               e.Embeds,
		Reactions:            e.Reactions,
		Nonce:                e.Nonce,
		Pinned:               e.Pinned,
		WebhookID:            e.WebhookID,
		Type:                 e.Type,
		Activity:             e.Activity,
		Application:          e.Application,
		ApplicationID:        e.ApplicationID,
		Flags:                e.Flags,
		MessageReference:     e.MessageReference,
		ReferencedMessage:    e.ReferencedMessage,
		MessageSnapshots:     e.MessageSnapshots,
		Call:                 e.Call,
		Interaction:          e.Interaction,
		InteractionMetadata:  e.InteractionMetadata,
		Resolved:             e.Resolved,
		Thread:               e.Thread,
		RoleSubscriptionData: e.RoleSubscriptionData,
		PurchaseNotification: e.PurchaseNotification,
		GiftInfo:             e.GiftInfo,
		Components:           e.Components,
		StickerItems:         e.StickerItems,
		Stickers:             e.Stickers,
		Poll:                 e.Poll,
		ChangelogID:          e.ChangelogID,
		SoundboardSounds:     e.SoundboardSounds,
		Potions:              e.Potions,
		SharedClientTheme:    e.SharedClientTheme,
	}
}

type MessageUpdateEvent struct {
	ID                   snowflake.ID                `json:"id"`
	ChannelID            snowflake.ID                `json:"channel_id"`
	LobbyID              snowflake.ID                `json:"lobby_id,omitzero"`
	Author               PartialUser                 `json:"author"`
	Content              string                      `json:"content"`
	Timestamp            time.Time                   `json:"timestamp"`
	EditedTimestamp      *time.Time                  `json:"edited_timestamp"`
	TTS                  bool                        `json:"tts"`
	MentionEveryone      bool                        `json:"mention_everyone"`
	Mentions             []PartialUser               `json:"mentions"`
	MentionRoles         []snowflake.ID              `json:"mention_roles"`
	MentionChannels      []PartialChannel            `json:"mention_channels,omitempty"`
	Attachments          []Attachment                `json:"attachments"`
	Embeds               []Embed                     `json:"embeds"`
	Reactions            []Reaction                  `json:"reactions,omitempty"`
	Nonce                json.RawMessage             `json:"nonce,omitempty"`
	Pinned               bool                        `json:"pinned"`
	WebhookID            snowflake.ID                `json:"webhook_id,omitzero"`
	Type                 MessageType                 `json:"type"`
	Activity             MessageActivity             `json:"activity,omitzero"`
	Application          IntegrationApplication      `json:"application,omitzero"`
	ApplicationID        snowflake.ID                `json:"application_id,omitzero"`
	Flags                MessageFlags                `json:"flags"`
	MessageReference     MessageReference            `json:"message_reference,omitzero"`
	ReferencedMessage    *Message                    `json:"referenced_message,omitempty"`
	MessageSnapshots     []MessageSnapshot           `json:"message_snapshots,omitempty"`
	Call                 MessageCall                 `json:"call,omitzero"`
	Interaction          PartialMessageInteraction   `json:"interaction,omitzero"`
	InteractionMetadata  MessageInteraction          `json:"interaction_metadata,omitzero"`
	Resolved             ResolvedData                `json:"resolved,omitzero"`
	Thread               Channel                     `json:"thread,omitzero"`
	RoleSubscriptionData MessageRoleSubscription     `json:"role_subscription_data,omitzero"`
	PurchaseNotification MessagePurchaseNotification `json:"purchase_notification,omitzero"`
	GiftInfo             MessageGiftInfo             `json:"gift_info,omitzero"`
	Components           Components                  `json:"components"`
	StickerItems         []StickerItem               `json:"sticker_items,omitempty"`
	Stickers             []Sticker                   `json:"stickers,omitempty"`
	Poll                 Poll                        `json:"poll,omitzero"`
	ChangelogID          snowflake.ID                `json:"changelog_id,omitzero"`
	SoundboardSounds     []SoundboardSound           `json:"soundboard_sounds,omitempty"`
	Potions              []Potion                    `json:"potions,omitempty"`
	SharedClientTheme    SharedClientTheme           `json:"shared_client_theme,omitzero"`
	ChannelType          ChannelType                 `json:"channel_type"`
	GuildID              snowflake.ID                `json:"guild_id,omitzero"`
	Member               GuildMember                 `json:"member,omitzero"`
	Metadata             map[string]string           `json:"metadata,omitempty"`
	ModerationMetadata   map[string]string           `json:"moderation_metadata,omitempty"`
}

func (e *MessageUpdateEvent) AsMessage() Message {
	return Message{
		ID:                   e.ID,
		ChannelID:            e.ChannelID,
		LobbyID:              e.LobbyID,
		Author:               e.Author,
		Content:              e.Content,
		Timestamp:            e.Timestamp,
		EditedTimestamp:      e.EditedTimestamp,
		TTS:                  e.TTS,
		MentionEveryone:      e.MentionEveryone,
		Mentions:             e.Mentions,
		MentionRoles:         e.MentionRoles,
		MentionChannels:      e.MentionChannels,
		Attachments:          e.Attachments,
		Embeds:               e.Embeds,
		Reactions:            e.Reactions,
		Nonce:                e.Nonce,
		Pinned:               e.Pinned,
		WebhookID:            e.WebhookID,
		Type:                 e.Type,
		Activity:             e.Activity,
		Application:          e.Application,
		ApplicationID:        e.ApplicationID,
		Flags:                e.Flags,
		MessageReference:     e.MessageReference,
		ReferencedMessage:    e.ReferencedMessage,
		MessageSnapshots:     e.MessageSnapshots,
		Call:                 e.Call,
		Interaction:          e.Interaction,
		InteractionMetadata:  e.InteractionMetadata,
		Resolved:             e.Resolved,
		Thread:               e.Thread,
		RoleSubscriptionData: e.RoleSubscriptionData,
		PurchaseNotification: e.PurchaseNotification,
		GiftInfo:             e.GiftInfo,
		Components:           e.Components,
		StickerItems:         e.StickerItems,
		Stickers:             e.Stickers,
		Poll:                 e.Poll,
		ChangelogID:          e.ChangelogID,
		SoundboardSounds:     e.SoundboardSounds,
		Potions:              e.Potions,
		SharedClientTheme:    e.SharedClientTheme,
	}
}

type MessageDeleteEvent struct {
	ID        snowflake.ID `json:"id"`
	ChannelID snowflake.ID `json:"channel_id"`
	GuildID   snowflake.ID `json:"guild_id,omitzero"`
}

func (e *MessageDeleteEvent) AsGatewayEventsMessageDelete() GatewayEventsMessageDelete {
	return GatewayEventsMessageDelete{
		ID:        e.ID,
		ChannelID: e.ChannelID,
		GuildID:   e.GuildID,
	}
}

type MessageDeleteBulkEvent struct {
	IDs       []snowflake.ID `json:"ids"`
	ChannelID snowflake.ID   `json:"channel_id"`
	GuildID   snowflake.ID   `json:"guild_id,omitzero"`
}

func (e *MessageDeleteBulkEvent) AsMessageDeleteBulk() MessageDeleteBulk {
	return MessageDeleteBulk{
		IDs:       e.IDs,
		ChannelID: e.ChannelID,
		GuildID:   e.GuildID,
	}
}

type MessageAckEvent struct {
	AckType      ReadStateType   `json:"ack_type,omitzero"`
	ChannelID    snowflake.ID    `json:"channel_id"`
	MessageID    snowflake.ID    `json:"message_id"`
	Manual       bool            `json:"manual,omitzero"`
	MentionCount int64           `json:"mention_count,omitzero"`
	Flags        *ReadStateFlags `json:"flags,omitempty"`
	LastViewed   *int64          `json:"last_viewed,omitempty"`
	Version      int64           `json:"version"`
}

func (e *MessageAckEvent) AsMessageAck() MessageAck {
	return MessageAck{
		AckType:      e.AckType,
		ChannelID:    e.ChannelID,
		MessageID:    e.MessageID,
		Manual:       e.Manual,
		MentionCount: e.MentionCount,
		Flags:        e.Flags,
		LastViewed:   e.LastViewed,
		Version:      e.Version,
	}
}

type MessagePollVoteAddEvent struct {
	UserID    snowflake.ID `json:"user_id"`
	ChannelID snowflake.ID `json:"channel_id"`
	MessageID snowflake.ID `json:"message_id"`
	GuildID   snowflake.ID `json:"guild_id,omitzero"`
	AnswerID  int64        `json:"answer_id"`
}

func (e *MessagePollVoteAddEvent) AsMessagePollVoteAdd() MessagePollVoteAdd {
	return MessagePollVoteAdd{
		UserID:    e.UserID,
		ChannelID: e.ChannelID,
		MessageID: e.MessageID,
		GuildID:   e.GuildID,
		AnswerID:  e.AnswerID,
	}
}

type MessagePollVoteRemoveEvent struct {
	UserID    snowflake.ID `json:"user_id"`
	ChannelID snowflake.ID `json:"channel_id"`
	MessageID snowflake.ID `json:"message_id"`
	GuildID   snowflake.ID `json:"guild_id,omitzero"`
	AnswerID  int64        `json:"answer_id"`
}

func (e *MessagePollVoteRemoveEvent) AsMessagePollVoteRemove() MessagePollVoteRemove {
	return MessagePollVoteRemove{
		UserID:    e.UserID,
		ChannelID: e.ChannelID,
		MessageID: e.MessageID,
		GuildID:   e.GuildID,
		AnswerID:  e.AnswerID,
	}
}

type MessageReactionAddEvent struct {
	UserID          snowflake.ID    `json:"user_id"`
	ChannelID       snowflake.ID    `json:"channel_id"`
	MessageID       snowflake.ID    `json:"message_id"`
	MessageAuthorID snowflake.ID    `json:"message_author_id"`
	GuildID         snowflake.ID    `json:"guild_id,omitzero"`
	Member          json.RawMessage `json:"member,omitempty"`
	Emoji           Emoji           `json:"emoji"`
	Type            ReactionType    `json:"type"`
	BurstColors     []string        `json:"burst_colors,omitempty"`
}

func (e *MessageReactionAddEvent) AsMessageReactionAdd() MessageReactionAdd {
	return MessageReactionAdd{
		UserID:          e.UserID,
		ChannelID:       e.ChannelID,
		MessageID:       e.MessageID,
		MessageAuthorID: e.MessageAuthorID,
		GuildID:         e.GuildID,
		Member:          e.Member,
		Emoji:           e.Emoji,
		Type:            e.Type,
		BurstColors:     e.BurstColors,
	}
}

type MessageReactionAddManyEvent struct {
	ChannelID snowflake.ID        `json:"channel_id"`
	MessageID snowflake.ID        `json:"message_id"`
	GuildID   snowflake.ID        `json:"guild_id,omitzero"`
	Reactions []DebouncedReaction `json:"reactions"`
}

func (e *MessageReactionAddManyEvent) AsMessageReactionAddMany() MessageReactionAddMany {
	return MessageReactionAddMany{
		ChannelID: e.ChannelID,
		MessageID: e.MessageID,
		GuildID:   e.GuildID,
		Reactions: e.Reactions,
	}
}

type MessageReactionRemoveEvent struct {
	UserID    snowflake.ID    `json:"user_id"`
	ChannelID snowflake.ID    `json:"channel_id"`
	MessageID snowflake.ID    `json:"message_id"`
	GuildID   snowflake.ID    `json:"guild_id,omitzero"`
	Emoji     json.RawMessage `json:"emoji"`
	Type      ReactionType    `json:"type"`
}

func (e *MessageReactionRemoveEvent) AsMessageReactionRemove() MessageReactionRemove {
	return MessageReactionRemove{
		UserID:    e.UserID,
		ChannelID: e.ChannelID,
		MessageID: e.MessageID,
		GuildID:   e.GuildID,
		Emoji:     e.Emoji,
		Type:      e.Type,
	}
}

type MessageReactionRemoveAllEvent struct {
	ChannelID snowflake.ID `json:"channel_id"`
	MessageID snowflake.ID `json:"message_id"`
	GuildID   snowflake.ID `json:"guild_id,omitzero"`
}

func (e *MessageReactionRemoveAllEvent) AsMessageReactionRemoveAll() MessageReactionRemoveAll {
	return MessageReactionRemoveAll{
		ChannelID: e.ChannelID,
		MessageID: e.MessageID,
		GuildID:   e.GuildID,
	}
}

type MessageReactionRemoveEmojiEvent struct {
	ChannelID snowflake.ID `json:"channel_id"`
	MessageID snowflake.ID `json:"message_id"`
	GuildID   snowflake.ID `json:"guild_id,omitzero"`
	Emoji     Emoji        `json:"emoji"`
}

func (e *MessageReactionRemoveEmojiEvent) AsMessageReactionRemoveEmoji() MessageReactionRemoveEmoji {
	return MessageReactionRemoveEmoji{
		ChannelID: e.ChannelID,
		MessageID: e.MessageID,
		GuildID:   e.GuildID,
		Emoji:     e.Emoji,
	}
}

type ReactionNotificationSentEvent struct {
	Message       Message      `json:"message"`
	ReactorUserID snowflake.ID `json:"reactor_user_id"`
	Emoji         Emoji        `json:"emoji"`
	Title         string       `json:"title"`
	Body          string       `json:"body"`
	Icon          string       `json:"icon"`
	Route         string       `json:"route"`
	TrackingType  string       `json:"tracking_type"`
}

func (e *ReactionNotificationSentEvent) AsReactionNotificationSent() ReactionNotificationSent {
	return ReactionNotificationSent{
		Message:       e.Message,
		ReactorUserID: e.ReactorUserID,
		Emoji:         e.Emoji,
		Title:         e.Title,
		Body:          e.Body,
		Icon:          e.Icon,
		Route:         e.Route,
		TrackingType:  e.TrackingType,
	}
}

type RecentMentionDeleteEvent struct {
	MessageID snowflake.ID `json:"message_id"`
}

func (e *RecentMentionDeleteEvent) AsRecentMentionDelete() RecentMentionDelete {
	return RecentMentionDelete{
		MessageID: e.MessageID,
	}
}

type LastMessagesEvent struct {
	GuildID  snowflake.ID `json:"guild_id"`
	Messages []Message    `json:"messages"`
}

func (e *LastMessagesEvent) AsLastMessages() LastMessages {
	return LastMessages{
		GuildID:  e.GuildID,
		Messages: e.Messages,
	}
}

type NotificationCenterItemCreateEvent struct {
	ID                     snowflake.ID                `json:"id"`
	BundleID               string                      `json:"bundle_id"`
	Type                   NotificationCenterItemType  `json:"type"`
	ItemEnum               *NotificationCenterItemEnum `json:"item_enum"`
	Body                   string                      `json:"body"`
	Acked                  bool                        `json:"acked"`
	Deeplink               string                      `json:"deeplink"`
	IconURL                *string                     `json:"icon_url"`
	IconName               *string                     `json:"icon_name"`
	OtherUser              *PartialUser                `json:"other_user"`
	Message                *Message                    `json:"message"`
	Completed              bool                        `json:"completed"`
	GuildID                *snowflake.ID               `json:"guild_id"`
	MessageID              *snowflake.ID               `json:"message_id"`
	MessageChannelID       *snowflake.ID               `json:"message_channel_id"`
	GuildScheduledEventID  *snowflake.ID               `json:"guild_scheduled_event_id"`
	DisableAction          bool                        `json:"disable_action"`
	Callout                *string                     `json:"callout"`
	Application            *PartialApplication         `json:"application"`
	EmojiID                *snowflake.ID               `json:"emoji_id"`
	EmojiName              *string                     `json:"emoji_name"`
	MessageContent         *string                     `json:"message_content"`
	MessageEmbedCount      *int64                      `json:"message_embed_count"`
	MessageAttachmentCount *int64                      `json:"message_attachment_count"`
	MessageStickerCount    *int64                      `json:"message_sticker_count"`
	IsVoiceMessage         bool                        `json:"is_voice_message"`
}

func (e *NotificationCenterItemCreateEvent) AsNotificationCenterItem() NotificationCenterItem {
	return NotificationCenterItem{
		ID:                     e.ID,
		BundleID:               e.BundleID,
		Type:                   e.Type,
		ItemEnum:               e.ItemEnum,
		Body:                   e.Body,
		Acked:                  e.Acked,
		Deeplink:               e.Deeplink,
		IconURL:                e.IconURL,
		IconName:               e.IconName,
		OtherUser:              e.OtherUser,
		Message:                e.Message,
		Completed:              e.Completed,
		GuildID:                e.GuildID,
		MessageID:              e.MessageID,
		MessageChannelID:       e.MessageChannelID,
		GuildScheduledEventID:  e.GuildScheduledEventID,
		DisableAction:          e.DisableAction,
		Callout:                e.Callout,
		Application:            e.Application,
		EmojiID:                e.EmojiID,
		EmojiName:              e.EmojiName,
		MessageContent:         e.MessageContent,
		MessageEmbedCount:      e.MessageEmbedCount,
		MessageAttachmentCount: e.MessageAttachmentCount,
		MessageStickerCount:    e.MessageStickerCount,
		IsVoiceMessage:         e.IsVoiceMessage,
	}
}

type NotificationCenterItemDeleteEvent struct {
	ID     snowflake.ID `json:"id"`
	UserID snowflake.ID `json:"user_id"`
}

func (e *NotificationCenterItemDeleteEvent) AsNotificationCenterItemDelete() NotificationCenterItemDelete {
	return NotificationCenterItemDelete{
		ID:     e.ID,
		UserID: e.UserID,
	}
}

type NotificationCenterItemsAckEvent struct {
	ID snowflake.ID `json:"id"`
}

func (e *NotificationCenterItemsAckEvent) AsNotificationCenterItemsAck() NotificationCenterItemsAck {
	return NotificationCenterItemsAck{
		ID: e.ID,
	}
}

type NotificationCenterItemCompletedEvent struct {
	ItemEnum NotificationCenterItemEnum `json:"item_enum"`
}

func (e *NotificationCenterItemCompletedEvent) AsNotificationCenterItemCompleted() NotificationCenterItemCompleted {
	return NotificationCenterItemCompleted{
		ItemEnum: e.ItemEnum,
	}
}

type NotificationSettingsUpdateEvent struct {
	Flags NotificationSettingsFlags `json:"flags"`
}

func (e *NotificationSettingsUpdateEvent) AsNotificationSettings() NotificationSettings {
	return NotificationSettings{
		Flags: e.Flags,
	}
}

type OAuth2TokenCreateEvent struct {
	ID          snowflake.ID       `json:"id"`
	Scopes      []string           `json:"scopes"`
	Application PartialApplication `json:"application"`
	Disclosures []int64            `json:"disclosures,omitempty"`
}

func (e *OAuth2TokenCreateEvent) AsOAuth2Authorization() OAuth2Authorization {
	return OAuth2Authorization{
		ID:          e.ID,
		Scopes:      e.Scopes,
		Application: e.Application,
		Disclosures: e.Disclosures,
	}
}

type OAuth2TokenDeleteEvent struct {
	ID            snowflake.ID `json:"id"`
	ApplicationID snowflake.ID `json:"application_id"`
}

func (e *OAuth2TokenDeleteEvent) AsOAuth2TokenDelete() OAuth2TokenDelete {
	return OAuth2TokenDelete{
		ID:            e.ID,
		ApplicationID: e.ApplicationID,
	}
}

type OAuth2TokenRevokeEvent struct {
	AccessToken   string       `json:"access_token"`
	ApplicationID snowflake.ID `json:"application_id,omitzero"`
}

func (e *OAuth2TokenRevokeEvent) AsOAuth2TokenRevoke() OAuth2TokenRevoke {
	return OAuth2TokenRevoke{
		AccessToken:   e.AccessToken,
		ApplicationID: e.ApplicationID,
	}
}

type PaymentUpdateEvent struct {
	ID                            snowflake.ID        `json:"id"`
	Amount                        int64               `json:"amount"`
	Tax                           int64               `json:"tax"`
	TaxInclusive                  bool                `json:"tax_inclusive"`
	Currency                      string              `json:"currency"`
	AmountRefunded                int64               `json:"amount_refunded"`
	Description                   string              `json:"description"`
	Status                        PaymentStatus       `json:"status"`
	CreatedAt                     time.Time           `json:"created_at"`
	SKUID                         snowflake.ID        `json:"sku_id,omitzero"`
	SKUPrice                      int64               `json:"sku_price,omitzero"`
	SKUSubscriptionPlanID         snowflake.ID        `json:"sku_subscription_plan_id,omitzero"`
	PaymentGateway                PaymentGateway      `json:"payment_gateway,omitzero"`
	PaymentGatewayPaymentID       string              `json:"payment_gateway_payment_id,omitzero"`
	HasInvoiceURL                 bool                `json:"has_invoice_url,omitzero"`
	HasRefundInvoiceURLs          bool                `json:"has_refund_invoice_urls,omitzero"`
	DownloadableInvoice           string              `json:"downloadable_invoice,omitzero"`
	DownloadableRefundInvoices    []string            `json:"downloadable_refund_invoices,omitempty"`
	RefundDisqualificationReasons []int64             `json:"refund_disqualification_reasons,omitempty"`
	Flags                         PaymentFlags        `json:"flags"`
	SKU                           SKU                 `json:"sku,omitzero"`
	PaymentSource                 PaymentSource       `json:"payment_source,omitzero"`
	Subscription                  PartialSubscription `json:"subscription,omitzero"`
	Metadata                      PaymentMetadata     `json:"metadata"`
}

func (e *PaymentUpdateEvent) AsPayment() Payment {
	return Payment{
		ID:                            e.ID,
		Amount:                        e.Amount,
		Tax:                           e.Tax,
		TaxInclusive:                  e.TaxInclusive,
		Currency:                      e.Currency,
		AmountRefunded:                e.AmountRefunded,
		Description:                   e.Description,
		Status:                        e.Status,
		CreatedAt:                     e.CreatedAt,
		SKUID:                         e.SKUID,
		SKUPrice:                      e.SKUPrice,
		SKUSubscriptionPlanID:         e.SKUSubscriptionPlanID,
		PaymentGateway:                e.PaymentGateway,
		PaymentGatewayPaymentID:       e.PaymentGatewayPaymentID,
		HasInvoiceURL:                 e.HasInvoiceURL,
		HasRefundInvoiceURLs:          e.HasRefundInvoiceURLs,
		DownloadableInvoice:           e.DownloadableInvoice,
		DownloadableRefundInvoices:    e.DownloadableRefundInvoices,
		RefundDisqualificationReasons: e.RefundDisqualificationReasons,
		Flags:                         e.Flags,
		SKU:                           e.SKU,
		PaymentSource:                 e.PaymentSource,
		Subscription:                  e.Subscription,
		Metadata:                      e.Metadata,
	}
}

type PresenceUpdateEvent struct {
	User             PartialUser  `json:"user"`
	GuildID          snowflake.ID `json:"guild_id,omitzero"`
	Status           StatusType   `json:"status"`
	Activities       []Activity   `json:"activities"`
	HiddenActivities []Activity   `json:"hidden_activities,omitempty"`
	ClientStatus     ClientStatus `json:"client_status"`
	HasPlayedGame    bool         `json:"has_played_game,omitzero"`
}

func (e *PresenceUpdateEvent) AsPresence() Presence {
	return Presence{
		User:             e.User,
		GuildID:          e.GuildID,
		Status:           e.Status,
		Activities:       e.Activities,
		HiddenActivities: e.HiddenActivities,
		ClientStatus:     e.ClientStatus,
		HasPlayedGame:    e.HasPlayedGame,
	}
}

type QuestsUserStatusUpdateEvent struct {
	UserStatus QuestUserStatus `json:"user_status"`
}

func (e *QuestsUserStatusUpdateEvent) AsQuestsUserStatusUpdate() QuestsUserStatusUpdate {
	return QuestsUserStatusUpdate{
		UserStatus: e.UserStatus,
	}
}

type QuestsUserCompletionUpdateEvent struct {
	QuestEnrollmentBlockedUntil *time.Time `json:"quest_enrollment_blocked_until"`
}

func (e *QuestsUserCompletionUpdateEvent) AsQuestsUserCompletionUpdate() QuestsUserCompletionUpdate {
	return QuestsUserCompletionUpdate{
		QuestEnrollmentBlockedUntil: e.QuestEnrollmentBlockedUntil,
	}
}

type RelationshipAddEvent struct {
	ID                  string           `json:"id"`
	Type                RelationshipType `json:"type"`
	User                PartialUser      `json:"user"`
	Nickname            *string          `json:"nickname"`
	IsSpamRequest       bool             `json:"is_spam_request,omitzero"`
	StrangerRequest     bool             `json:"stranger_request,omitzero"`
	UserIgnored         bool             `json:"user_ignored"`
	OriginApplicationID *snowflake.ID    `json:"origin_application_id,omitempty"`
	Since               time.Time        `json:"since,omitzero"`
	HasPlayedGame       bool             `json:"has_played_game,omitzero"`
	Note                string           `json:"note,omitzero"`
	ShouldNotify        bool             `json:"should_notify,omitzero"`
}

func (e *RelationshipAddEvent) AsRelationship() Relationship {
	return Relationship{
		ID:                  e.ID,
		Type:                e.Type,
		User:                e.User,
		Nickname:            e.Nickname,
		IsSpamRequest:       e.IsSpamRequest,
		StrangerRequest:     e.StrangerRequest,
		UserIgnored:         e.UserIgnored,
		OriginApplicationID: e.OriginApplicationID,
		Since:               e.Since,
		HasPlayedGame:       e.HasPlayedGame,
		Note:                e.Note,
	}
}

type RelationshipUpdateEvent struct {
	ID              snowflake.ID     `json:"id"`
	Type            RelationshipType `json:"type"`
	Nickname        *string          `json:"nickname"`
	Since           time.Time        `json:"since,omitzero"`
	StrangerRequest bool             `json:"stranger_request,omitzero"`
	UserIgnored     bool             `json:"user_ignored"`
}

func (e *RelationshipUpdateEvent) AsPartialRelationship() PartialRelationship {
	return PartialRelationship{
		ID:              e.ID,
		Type:            e.Type,
		Nickname:        e.Nickname,
		Since:           e.Since,
		StrangerRequest: e.StrangerRequest,
		UserIgnored:     e.UserIgnored,
	}
}

type RelationshipRemoveEvent struct {
	ID              snowflake.ID     `json:"id"`
	Type            RelationshipType `json:"type"`
	Nickname        *string          `json:"nickname"`
	Since           time.Time        `json:"since,omitzero"`
	StrangerRequest bool             `json:"stranger_request,omitzero"`
	UserIgnored     bool             `json:"user_ignored"`
}

func (e *RelationshipRemoveEvent) AsPartialRelationship() PartialRelationship {
	return PartialRelationship{
		ID:              e.ID,
		Type:            e.Type,
		Nickname:        e.Nickname,
		Since:           e.Since,
		StrangerRequest: e.StrangerRequest,
		UserIgnored:     e.UserIgnored,
	}
}

type GameRelationshipAddEvent struct {
	ID            string       `json:"id"`
	ApplicationID snowflake.ID `json:"application_id"`
	Type          int64        `json:"type"`
	User          PartialUser  `json:"user"`
	Since         time.Time    `json:"since"`
	DMAccessType  int64        `json:"dm_access_type"`
	UserID        snowflake.ID `json:"user_id"`
}

func (e *GameRelationshipAddEvent) AsGameRelationship() GameRelationship {
	return GameRelationship{
		ID:            e.ID,
		ApplicationID: e.ApplicationID,
		Type:          e.Type,
		User:          e.User,
		Since:         e.Since,
		DMAccessType:  e.DMAccessType,
		UserID:        e.UserID,
	}
}

type GameRelationshipRemoveEvent struct {
	ID            string       `json:"id"`
	ApplicationID snowflake.ID `json:"application_id"`
	Type          int64        `json:"type"`
	Since         time.Time    `json:"since"`
	DMAccessType  int64        `json:"dm_access_type"`
	UserID        snowflake.ID `json:"user_id"`
}

func (e *GameRelationshipRemoveEvent) AsGameRelationshipRemove() GameRelationshipRemove {
	return GameRelationshipRemove{
		ID:            e.ID,
		ApplicationID: e.ApplicationID,
		Type:          e.Type,
		Since:         e.Since,
		DMAccessType:  e.DMAccessType,
		UserID:        e.UserID,
	}
}

type GameInviteCreateEvent struct {
	InviteID         snowflake.ID   `json:"invite_id"`
	CreatedAt        time.Time      `json:"created_at"`
	TTL              int64          `json:"ttl"`
	InviterID        snowflake.ID   `json:"inviter_id"`
	RecipientID      snowflake.ID   `json:"recipient_id"`
	PlatformType     ConnectionType `json:"platform_type"`
	LaunchParameters string         `json:"launch_parameters"`
	Installed        bool           `json:"installed,omitzero"`
	Joinable         bool           `json:"joinable,omitzero"`
	FallbackURL      *string        `json:"fallback_url"`
	ApplicationAsset string         `json:"application_asset"`
	ApplicationName  string         `json:"application_name"`
}

func (e *GameInviteCreateEvent) AsGameInvite() GameInvite {
	return GameInvite{
		InviteID:         e.InviteID,
		CreatedAt:        e.CreatedAt,
		TTL:              e.TTL,
		InviterID:        e.InviterID,
		RecipientID:      e.RecipientID,
		PlatformType:     e.PlatformType,
		LaunchParameters: e.LaunchParameters,
		Installed:        e.Installed,
		Joinable:         e.Joinable,
		FallbackURL:      e.FallbackURL,
		ApplicationAsset: e.ApplicationAsset,
		ApplicationName:  e.ApplicationName,
	}
}

type GameInviteDeleteEvent struct {
	InviteID snowflake.ID `json:"invite_id"`
}

func (e *GameInviteDeleteEvent) AsGameInviteDelete() GameInviteDelete {
	return GameInviteDelete{
		InviteID: e.InviteID,
	}
}

type GameInviteDeleteManyEvent struct {
	InviteIDs []snowflake.ID `json:"invite_ids"`
}

func (e *GameInviteDeleteManyEvent) AsGameInviteDeleteMany() GameInviteDeleteMany {
	return GameInviteDeleteMany{
		InviteIDs: e.InviteIDs,
	}
}

type GatewayLobbyObjectEvent struct {
	VoiceStates []json.RawMessage `json:"voice_states"`
	Region      string            `json:"region"`
	Metadata    map[string]string `json:"metadata"`
}

type LobbyDeleteEvent struct {
	ID     snowflake.ID        `json:"id"`
	Reason LobbyDeletionReason `json:"reason"`
}

func (e *LobbyDeleteEvent) AsLobbyDelete() LobbyDelete {
	return LobbyDelete{
		ID:     e.ID,
		Reason: e.Reason,
	}
}

type LobbyMemberAddEvent struct {
	Member        LobbyMember  `json:"member"`
	LobbyID       snowflake.ID `json:"lobby_id"`
	ApplicationID snowflake.ID `json:"application_id"`
}

func (e *LobbyMemberAddEvent) AsLobbyMemberAdd() LobbyMemberAdd {
	return LobbyMemberAdd{
		Member:        e.Member,
		LobbyID:       e.LobbyID,
		ApplicationID: e.ApplicationID,
	}
}

type LobbyMemberConnectEvent struct {
	Member  LobbyMember  `json:"member"`
	LobbyID snowflake.ID `json:"lobby_id"`
}

func (e *LobbyMemberConnectEvent) AsLobbyMemberConnect() LobbyMemberConnect {
	return LobbyMemberConnect{
		Member:  e.Member,
		LobbyID: e.LobbyID,
	}
}

type LobbyMemberDisconnectEvent struct {
	Member  LobbyMember  `json:"member"`
	LobbyID snowflake.ID `json:"lobby_id"`
}

func (e *LobbyMemberDisconnectEvent) AsLobbyMemberDisconnect() LobbyMemberDisconnect {
	return LobbyMemberDisconnect{
		Member:  e.Member,
		LobbyID: e.LobbyID,
	}
}

type LobbyMemberUpdateEvent struct {
	Member        LobbyMember  `json:"member"`
	LobbyID       snowflake.ID `json:"lobby_id"`
	ApplicationID snowflake.ID `json:"application_id"`
}

func (e *LobbyMemberUpdateEvent) AsLobbyMemberUpdate() LobbyMemberUpdate {
	return LobbyMemberUpdate{
		Member:        e.Member,
		LobbyID:       e.LobbyID,
		ApplicationID: e.ApplicationID,
	}
}

type LobbyMemberRemoveEvent struct {
	Member        LobbyMember  `json:"member"`
	LobbyID       snowflake.ID `json:"lobby_id"`
	ApplicationID snowflake.ID `json:"application_id"`
}

func (e *LobbyMemberRemoveEvent) AsLobbyMemberRemove() LobbyMemberRemove {
	return LobbyMemberRemove{
		Member:        e.Member,
		LobbyID:       e.LobbyID,
		ApplicationID: e.ApplicationID,
	}
}

type LobbyMessageCreateEvent struct {
	ID                   snowflake.ID                `json:"id"`
	ChannelID            snowflake.ID                `json:"channel_id"`
	LobbyID              snowflake.ID                `json:"lobby_id,omitzero"`
	Author               PartialUser                 `json:"author"`
	Content              string                      `json:"content"`
	Timestamp            time.Time                   `json:"timestamp"`
	EditedTimestamp      *time.Time                  `json:"edited_timestamp"`
	TTS                  bool                        `json:"tts"`
	MentionEveryone      bool                        `json:"mention_everyone"`
	Mentions             []PartialUser               `json:"mentions"`
	MentionRoles         []snowflake.ID              `json:"mention_roles"`
	MentionChannels      []PartialChannel            `json:"mention_channels,omitempty"`
	Attachments          []Attachment                `json:"attachments"`
	Embeds               []Embed                     `json:"embeds"`
	Reactions            []Reaction                  `json:"reactions,omitempty"`
	Nonce                json.RawMessage             `json:"nonce,omitempty"`
	Pinned               bool                        `json:"pinned"`
	WebhookID            snowflake.ID                `json:"webhook_id,omitzero"`
	Type                 MessageType                 `json:"type"`
	Activity             MessageActivity             `json:"activity,omitzero"`
	Application          IntegrationApplication      `json:"application,omitzero"`
	ApplicationID        snowflake.ID                `json:"application_id,omitzero"`
	Flags                MessageFlags                `json:"flags"`
	MessageReference     MessageReference            `json:"message_reference,omitzero"`
	ReferencedMessage    *Message                    `json:"referenced_message,omitempty"`
	MessageSnapshots     []MessageSnapshot           `json:"message_snapshots,omitempty"`
	Call                 MessageCall                 `json:"call,omitzero"`
	Interaction          PartialMessageInteraction   `json:"interaction,omitzero"`
	InteractionMetadata  MessageInteraction          `json:"interaction_metadata,omitzero"`
	Resolved             ResolvedData                `json:"resolved,omitzero"`
	Thread               Channel                     `json:"thread,omitzero"`
	RoleSubscriptionData MessageRoleSubscription     `json:"role_subscription_data,omitzero"`
	PurchaseNotification MessagePurchaseNotification `json:"purchase_notification,omitzero"`
	GiftInfo             MessageGiftInfo             `json:"gift_info,omitzero"`
	Components           Components                  `json:"components"`
	StickerItems         []StickerItem               `json:"sticker_items,omitempty"`
	Stickers             []Sticker                   `json:"stickers,omitempty"`
	Poll                 Poll                        `json:"poll,omitzero"`
	ChangelogID          snowflake.ID                `json:"changelog_id,omitzero"`
	SoundboardSounds     []SoundboardSound           `json:"soundboard_sounds,omitempty"`
	Potions              []Potion                    `json:"potions,omitempty"`
	SharedClientTheme    SharedClientTheme           `json:"shared_client_theme,omitzero"`
	ChannelType          ChannelType                 `json:"channel_type"`
	GuildID              snowflake.ID                `json:"guild_id,omitzero"`
	Member               GuildMember                 `json:"member,omitzero"`
	Metadata             map[string]string           `json:"metadata,omitempty"`
	ModerationMetadata   map[string]string           `json:"moderation_metadata,omitempty"`
}

func (e *LobbyMessageCreateEvent) AsMessage() Message {
	return Message{
		ID:                   e.ID,
		ChannelID:            e.ChannelID,
		LobbyID:              e.LobbyID,
		Author:               e.Author,
		Content:              e.Content,
		Timestamp:            e.Timestamp,
		EditedTimestamp:      e.EditedTimestamp,
		TTS:                  e.TTS,
		MentionEveryone:      e.MentionEveryone,
		Mentions:             e.Mentions,
		MentionRoles:         e.MentionRoles,
		MentionChannels:      e.MentionChannels,
		Attachments:          e.Attachments,
		Embeds:               e.Embeds,
		Reactions:            e.Reactions,
		Nonce:                e.Nonce,
		Pinned:               e.Pinned,
		WebhookID:            e.WebhookID,
		Type:                 e.Type,
		Activity:             e.Activity,
		Application:          e.Application,
		ApplicationID:        e.ApplicationID,
		Flags:                e.Flags,
		MessageReference:     e.MessageReference,
		ReferencedMessage:    e.ReferencedMessage,
		MessageSnapshots:     e.MessageSnapshots,
		Call:                 e.Call,
		Interaction:          e.Interaction,
		InteractionMetadata:  e.InteractionMetadata,
		Resolved:             e.Resolved,
		Thread:               e.Thread,
		RoleSubscriptionData: e.RoleSubscriptionData,
		PurchaseNotification: e.PurchaseNotification,
		GiftInfo:             e.GiftInfo,
		Components:           e.Components,
		StickerItems:         e.StickerItems,
		Stickers:             e.Stickers,
		Poll:                 e.Poll,
		ChangelogID:          e.ChangelogID,
		SoundboardSounds:     e.SoundboardSounds,
		Potions:              e.Potions,
		SharedClientTheme:    e.SharedClientTheme,
	}
}

type LobbyMessageUpdateEvent struct {
	ID                   snowflake.ID                `json:"id"`
	ChannelID            snowflake.ID                `json:"channel_id"`
	LobbyID              snowflake.ID                `json:"lobby_id,omitzero"`
	Author               PartialUser                 `json:"author"`
	Content              string                      `json:"content"`
	Timestamp            time.Time                   `json:"timestamp"`
	EditedTimestamp      *time.Time                  `json:"edited_timestamp"`
	TTS                  bool                        `json:"tts"`
	MentionEveryone      bool                        `json:"mention_everyone"`
	Mentions             []PartialUser               `json:"mentions"`
	MentionRoles         []snowflake.ID              `json:"mention_roles"`
	MentionChannels      []PartialChannel            `json:"mention_channels,omitempty"`
	Attachments          []Attachment                `json:"attachments"`
	Embeds               []Embed                     `json:"embeds"`
	Reactions            []Reaction                  `json:"reactions,omitempty"`
	Nonce                json.RawMessage             `json:"nonce,omitempty"`
	Pinned               bool                        `json:"pinned"`
	WebhookID            snowflake.ID                `json:"webhook_id,omitzero"`
	Type                 MessageType                 `json:"type"`
	Activity             MessageActivity             `json:"activity,omitzero"`
	Application          IntegrationApplication      `json:"application,omitzero"`
	ApplicationID        snowflake.ID                `json:"application_id,omitzero"`
	Flags                MessageFlags                `json:"flags"`
	MessageReference     MessageReference            `json:"message_reference,omitzero"`
	ReferencedMessage    *Message                    `json:"referenced_message,omitempty"`
	MessageSnapshots     []MessageSnapshot           `json:"message_snapshots,omitempty"`
	Call                 MessageCall                 `json:"call,omitzero"`
	Interaction          PartialMessageInteraction   `json:"interaction,omitzero"`
	InteractionMetadata  MessageInteraction          `json:"interaction_metadata,omitzero"`
	Resolved             ResolvedData                `json:"resolved,omitzero"`
	Thread               Channel                     `json:"thread,omitzero"`
	RoleSubscriptionData MessageRoleSubscription     `json:"role_subscription_data,omitzero"`
	PurchaseNotification MessagePurchaseNotification `json:"purchase_notification,omitzero"`
	GiftInfo             MessageGiftInfo             `json:"gift_info,omitzero"`
	Components           Components                  `json:"components"`
	StickerItems         []StickerItem               `json:"sticker_items,omitempty"`
	Stickers             []Sticker                   `json:"stickers,omitempty"`
	Poll                 Poll                        `json:"poll,omitzero"`
	ChangelogID          snowflake.ID                `json:"changelog_id,omitzero"`
	SoundboardSounds     []SoundboardSound           `json:"soundboard_sounds,omitempty"`
	Potions              []Potion                    `json:"potions,omitempty"`
	SharedClientTheme    SharedClientTheme           `json:"shared_client_theme,omitzero"`
	ChannelType          ChannelType                 `json:"channel_type"`
	GuildID              snowflake.ID                `json:"guild_id,omitzero"`
	Member               GuildMember                 `json:"member,omitzero"`
	Metadata             map[string]string           `json:"metadata,omitempty"`
	ModerationMetadata   map[string]string           `json:"moderation_metadata,omitempty"`
}

func (e *LobbyMessageUpdateEvent) AsMessage() Message {
	return Message{
		ID:                   e.ID,
		ChannelID:            e.ChannelID,
		LobbyID:              e.LobbyID,
		Author:               e.Author,
		Content:              e.Content,
		Timestamp:            e.Timestamp,
		EditedTimestamp:      e.EditedTimestamp,
		TTS:                  e.TTS,
		MentionEveryone:      e.MentionEveryone,
		Mentions:             e.Mentions,
		MentionRoles:         e.MentionRoles,
		MentionChannels:      e.MentionChannels,
		Attachments:          e.Attachments,
		Embeds:               e.Embeds,
		Reactions:            e.Reactions,
		Nonce:                e.Nonce,
		Pinned:               e.Pinned,
		WebhookID:            e.WebhookID,
		Type:                 e.Type,
		Activity:             e.Activity,
		Application:          e.Application,
		ApplicationID:        e.ApplicationID,
		Flags:                e.Flags,
		MessageReference:     e.MessageReference,
		ReferencedMessage:    e.ReferencedMessage,
		MessageSnapshots:     e.MessageSnapshots,
		Call:                 e.Call,
		Interaction:          e.Interaction,
		InteractionMetadata:  e.InteractionMetadata,
		Resolved:             e.Resolved,
		Thread:               e.Thread,
		RoleSubscriptionData: e.RoleSubscriptionData,
		PurchaseNotification: e.PurchaseNotification,
		GiftInfo:             e.GiftInfo,
		Components:           e.Components,
		StickerItems:         e.StickerItems,
		Stickers:             e.Stickers,
		Poll:                 e.Poll,
		ChangelogID:          e.ChangelogID,
		SoundboardSounds:     e.SoundboardSounds,
		Potions:              e.Potions,
		SharedClientTheme:    e.SharedClientTheme,
	}
}

type LobbyVoiceServerUpdateEvent struct {
	Token    string       `json:"token"`
	LobbyID  snowflake.ID `json:"lobby_id"`
	Endpoint *string      `json:"endpoint"`
}

func (e *LobbyVoiceServerUpdateEvent) AsLobbyVoiceServerUpdate() LobbyVoiceServerUpdate {
	return LobbyVoiceServerUpdate{
		Token:    e.Token,
		LobbyID:  e.LobbyID,
		Endpoint: e.Endpoint,
	}
}

type PassiveUpdateV1Event struct {
	GuildID     snowflake.ID    `json:"guild_id"`
	Channels    []ChannelUnread `json:"channels"`
	VoiceStates []VoiceState    `json:"voice_states,omitempty"`
	Members     []GuildMember   `json:"members,omitempty"`
}

func (e *PassiveUpdateV1Event) AsPassiveUpdateV1() PassiveUpdateV1 {
	return PassiveUpdateV1{
		GuildID:     e.GuildID,
		Channels:    e.Channels,
		VoiceStates: e.VoiceStates,
		Members:     e.Members,
	}
}

type PassiveUpdateV2Event struct {
	GuildID            snowflake.ID    `json:"guild_id"`
	UpdatedChannels    []ChannelUnread `json:"updated_channels"`
	UpdatedVoiceStates []VoiceState    `json:"updated_voice_states"`
	RemovedVoiceStates []snowflake.ID  `json:"removed_voice_states"`
	UpdatedMembers     []GuildMember   `json:"updated_members"`
}

func (e *PassiveUpdateV2Event) AsPassiveUpdateV2() PassiveUpdateV2 {
	return PassiveUpdateV2{
		GuildID:            e.GuildID,
		UpdatedChannels:    e.UpdatedChannels,
		UpdatedVoiceStates: e.UpdatedVoiceStates,
		RemovedVoiceStates: e.RemovedVoiceStates,
		UpdatedMembers:     e.UpdatedMembers,
	}
}

type SavedMessageCreateEvent struct {
	Message  *Message `json:"message"`
	SaveData SaveData `json:"save_data"`
}

func (e *SavedMessageCreateEvent) AsSavedMessage() SavedMessage {
	return SavedMessage{
		Message:  e.Message,
		SaveData: e.SaveData,
	}
}

type SavedMessageDeleteEvent struct {
	ChannelID snowflake.ID `json:"channel_id"`
	MessageID snowflake.ID `json:"message_id"`
}

func (e *SavedMessageDeleteEvent) AsSavedMessageDelete() SavedMessageDelete {
	return SavedMessageDelete{
		ChannelID: e.ChannelID,
		MessageID: e.MessageID,
	}
}

type SessionsReplaceEvent []Session

type StageInstanceCreateEvent struct {
	ID                    snowflake.ID  `json:"id"`
	GuildID               snowflake.ID  `json:"guild_id"`
	ChannelID             snowflake.ID  `json:"channel_id"`
	Topic                 string        `json:"topic"`
	PrivacyLevel          PrivacyLevel  `json:"privacy_level"`
	InviteCode            *string       `json:"invite_code"`
	DiscoverableDisabled  bool          `json:"discoverable_disabled"`
	GuildScheduledEventID *snowflake.ID `json:"guild_scheduled_event_id"`
}

func (e *StageInstanceCreateEvent) AsStageInstance() StageInstance {
	return StageInstance{
		ID:                    e.ID,
		GuildID:               e.GuildID,
		ChannelID:             e.ChannelID,
		Topic:                 e.Topic,
		PrivacyLevel:          e.PrivacyLevel,
		InviteCode:            e.InviteCode,
		DiscoverableDisabled:  e.DiscoverableDisabled,
		GuildScheduledEventID: e.GuildScheduledEventID,
	}
}

type StageInstanceUpdateEvent struct {
	ID                    snowflake.ID  `json:"id"`
	GuildID               snowflake.ID  `json:"guild_id"`
	ChannelID             snowflake.ID  `json:"channel_id"`
	Topic                 string        `json:"topic"`
	PrivacyLevel          PrivacyLevel  `json:"privacy_level"`
	InviteCode            *string       `json:"invite_code"`
	DiscoverableDisabled  bool          `json:"discoverable_disabled"`
	GuildScheduledEventID *snowflake.ID `json:"guild_scheduled_event_id"`
}

func (e *StageInstanceUpdateEvent) AsStageInstance() StageInstance {
	return StageInstance{
		ID:                    e.ID,
		GuildID:               e.GuildID,
		ChannelID:             e.ChannelID,
		Topic:                 e.Topic,
		PrivacyLevel:          e.PrivacyLevel,
		InviteCode:            e.InviteCode,
		DiscoverableDisabled:  e.DiscoverableDisabled,
		GuildScheduledEventID: e.GuildScheduledEventID,
	}
}

type StageInstanceDeleteEvent struct {
	ID                    snowflake.ID  `json:"id"`
	GuildID               snowflake.ID  `json:"guild_id"`
	ChannelID             snowflake.ID  `json:"channel_id"`
	Topic                 string        `json:"topic"`
	PrivacyLevel          PrivacyLevel  `json:"privacy_level"`
	InviteCode            *string       `json:"invite_code"`
	DiscoverableDisabled  bool          `json:"discoverable_disabled"`
	GuildScheduledEventID *snowflake.ID `json:"guild_scheduled_event_id"`
}

func (e *StageInstanceDeleteEvent) AsStageInstance() StageInstance {
	return StageInstance{
		ID:                    e.ID,
		GuildID:               e.GuildID,
		ChannelID:             e.ChannelID,
		Topic:                 e.Topic,
		PrivacyLevel:          e.PrivacyLevel,
		InviteCode:            e.InviteCode,
		DiscoverableDisabled:  e.DiscoverableDisabled,
		GuildScheduledEventID: e.GuildScheduledEventID,
	}
}

type StreamCreateEvent struct {
	StreamKey    string          `json:"stream_key"`
	RTCServerID  snowflake.ID    `json:"rtc_server_id"`
	RTCChannelID snowflake.ID    `json:"rtc_channel_id"`
	Region       string          `json:"region"`
	ViewerIDs    json.RawMessage `json:"viewer_ids"`
	Paused       bool            `json:"paused"`
}

func (e *StreamCreateEvent) AsStream() Stream {
	return Stream{
		StreamKey:    e.StreamKey,
		RTCServerID:  e.RTCServerID,
		RTCChannelID: e.RTCChannelID,
		Region:       e.Region,
		ViewerIDs:    e.ViewerIDs,
		Paused:       e.Paused,
	}
}

type StreamServerUpdateEvent struct {
	Token     string  `json:"token"`
	StreamKey string  `json:"stream_key"`
	Endpoint  *string `json:"endpoint"`
}

func (e *StreamServerUpdateEvent) AsStreamServerUpdate() StreamServerUpdate {
	return StreamServerUpdate{
		Token:     e.Token,
		StreamKey: e.StreamKey,
		Endpoint:  e.Endpoint,
	}
}

type StreamUpdateEvent struct {
	StreamKey    string          `json:"stream_key"`
	RTCServerID  snowflake.ID    `json:"rtc_server_id"`
	RTCChannelID snowflake.ID    `json:"rtc_channel_id"`
	Region       string          `json:"region"`
	ViewerIDs    json.RawMessage `json:"viewer_ids"`
	Paused       bool            `json:"paused"`
}

func (e *StreamUpdateEvent) AsStream() Stream {
	return Stream{
		StreamKey:    e.StreamKey,
		RTCServerID:  e.RTCServerID,
		RTCChannelID: e.RTCChannelID,
		Region:       e.Region,
		ViewerIDs:    e.ViewerIDs,
		Paused:       e.Paused,
	}
}

type StreamDeleteEvent struct {
	StreamKey   string             `json:"stream_key"`
	Reason      StreamDeleteReason `json:"reason"`
	Unavailable bool               `json:"unavailable,omitzero"`
}

func (e *StreamDeleteEvent) AsStreamDelete() StreamDelete {
	return StreamDelete{
		StreamKey:   e.StreamKey,
		Reason:      e.Reason,
		Unavailable: e.Unavailable,
	}
}

type TypingStartEvent struct {
	ChannelID snowflake.ID    `json:"channel_id"`
	GuildID   snowflake.ID    `json:"guild_id,omitzero"`
	UserID    snowflake.ID    `json:"user_id"`
	Timestamp int64           `json:"timestamp"`
	Member    json.RawMessage `json:"member,omitempty"`
}

func (e *TypingStartEvent) AsTypingStart() TypingStart {
	return TypingStart{
		ChannelID: e.ChannelID,
		GuildID:   e.GuildID,
		UserID:    e.UserID,
		Timestamp: e.Timestamp,
		Member:    e.Member,
	}
}

type UserUpdateEvent struct {
	ID                    snowflake.ID          `json:"id"`
	Username              string                `json:"username"`
	Discriminator         string                `json:"discriminator"`
	GlobalName            *string               `json:"global_name"`
	Avatar                *string               `json:"avatar"`
	AvatarDecorationData  *AvatarDecorationData `json:"avatar_decoration_data"`
	Collectibles          *Collectibles         `json:"collectibles,omitempty"`
	DisplayNameStyles     *DisplayNameStyle     `json:"display_name_styles,omitempty"`
	PrimaryGuild          *PrimaryGuild         `json:"primary_guild,omitempty"`
	LinkedUsers           []LinkedUser          `json:"linked_users"`
	Bot                   bool                  `json:"bot,omitzero"`
	System                bool                  `json:"system,omitzero"`
	MFAEnabled            bool                  `json:"mfa_enabled"`
	NSFWAllowed           *bool                 `json:"nsfw_allowed,omitempty"`
	AgeVerificationStatus AgeVerificationStatus `json:"age_verification_status"`
	Pronouns              string                `json:"pronouns,omitzero"`
	Bio                   string                `json:"bio"`
	Banner                *string               `json:"banner"`
	AccentColor           *int64                `json:"accent_color"`
	Locale                string                `json:"locale,omitzero"`
	Verified              bool                  `json:"verified"`
	Email                 *string               `json:"email"`
	Phone                 *string               `json:"phone,omitempty"`
	Premium               bool                  `json:"premium"`
	PremiumType           PremiumType           `json:"premium_type"`
	PremiumState          PremiumState          `json:"premium_state,omitzero"`
	PersonalConnectionID  snowflake.ID          `json:"personal_connection_id,omitzero"`
	Flags                 UserFlags             `json:"flags"`
	PublicFlags           UserFlags             `json:"public_flags,omitzero"`
	PurchasedFlags        PurchasedFlags        `json:"purchased_flags,omitzero"`
	PremiumUsageFlags     PremiumUsageFlags     `json:"premium_usage_flags,omitzero"`
	Desktop               bool                  `json:"desktop,omitzero"`
	Mobile                bool                  `json:"mobile,omitzero"`
	HasBouncedEmail       bool                  `json:"has_bounced_email,omitzero"`
	AuthenticatorTypes    []int64               `json:"authenticator_types,omitempty"`
	AnalyticsToken        string                `json:"analytics_token"`
}

func (e *UserUpdateEvent) AsUser() User {
	return User{
		ID:                    e.ID,
		Username:              e.Username,
		Discriminator:         e.Discriminator,
		GlobalName:            e.GlobalName,
		Avatar:                e.Avatar,
		AvatarDecorationData:  e.AvatarDecorationData,
		Collectibles:          e.Collectibles,
		DisplayNameStyles:     e.DisplayNameStyles,
		PrimaryGuild:          e.PrimaryGuild,
		LinkedUsers:           e.LinkedUsers,
		Bot:                   e.Bot,
		System:                e.System,
		MFAEnabled:            e.MFAEnabled,
		NSFWAllowed:           e.NSFWAllowed,
		AgeVerificationStatus: e.AgeVerificationStatus,
		Pronouns:              e.Pronouns,
		Bio:                   e.Bio,
		Banner:                e.Banner,
		AccentColor:           e.AccentColor,
		Locale:                e.Locale,
		Verified:              e.Verified,
		Email:                 e.Email,
		Phone:                 e.Phone,
		Premium:               e.Premium,
		PremiumType:           e.PremiumType,
		PremiumState:          e.PremiumState,
		PersonalConnectionID:  e.PersonalConnectionID,
		Flags:                 e.Flags,
		PublicFlags:           e.PublicFlags,
		PurchasedFlags:        e.PurchasedFlags,
		PremiumUsageFlags:     e.PremiumUsageFlags,
		Desktop:               e.Desktop,
		Mobile:                e.Mobile,
		HasBouncedEmail:       e.HasBouncedEmail,
		AuthenticatorTypes:    e.AuthenticatorTypes,
		AnalyticsToken:        e.AnalyticsToken,
	}
}

type UserApplicationRemoveEvent struct {
	ApplicationID snowflake.ID `json:"application_id"`
}

func (e *UserApplicationRemoveEvent) AsGatewayEventsUserApplicationRemove() GatewayEventsUserApplicationRemove {
	return GatewayEventsUserApplicationRemove{
		ApplicationID: e.ApplicationID,
	}
}

type UserApplicationIdentityUpdateEvent struct {
	Username    *string                          `json:"username"`
	Metadata    *string                          `json:"metadata"`
	Data        *UserApplicationProfileData      `json:"data,omitempty"`
	DataTrusted *bool                            `json:"data_trusted,omitempty"`
	ExternalID  UserApplicationProfileExternalID `json:"external_id"`
	AvatarHash  *string                          `json:"avatar_hash"`
}

func (e *UserApplicationIdentityUpdateEvent) AsUserApplicationProfile() UserApplicationProfile {
	return UserApplicationProfile{
		Username:    e.Username,
		Metadata:    e.Metadata,
		Data:        e.Data,
		DataTrusted: e.DataTrusted,
		ExternalID:  e.ExternalID,
		AvatarHash:  e.AvatarHash,
	}
}

type UserApplicationIdentityRemoveEvent struct {
	UserID        snowflake.ID `json:"user_id"`
	ApplicationID snowflake.ID `json:"application_id"`
}

func (e *UserApplicationIdentityRemoveEvent) AsUserApplicationIdentityRemove() UserApplicationIdentityRemove {
	return UserApplicationIdentityRemove{
		UserID:        e.UserID,
		ApplicationID: e.ApplicationID,
	}
}

type UserConnectionsUpdateEvent struct {
	UserID snowflake.ID `json:"user_id"`
}

func (e *UserConnectionsUpdateEvent) AsUserConnectionsUpdate() UserConnectionsUpdate {
	return UserConnectionsUpdate{
		UserID: e.UserID,
	}
}

type UserGuildSettingsUpdateEvent struct {
	ChannelOverrides     []ChannelOverride        `json:"channel_overrides"`
	Flags                UserGuildSettingsFlags   `json:"flags"`
	GuildID              *snowflake.ID            `json:"guild_id"`
	HideMutedChannels    bool                     `json:"hide_muted_channels"`
	MessageNotifications MessageNotificationLevel `json:"message_notifications"`
	MobilePush           bool                     `json:"mobile_push"`
	MuteScheduledEvents  bool                     `json:"mute_scheduled_events"`
	Muted                bool                     `json:"muted"`
	MuteConfig           *MuteConfig              `json:"mute_config"`
	NotifyHighlights     HighlightLevel           `json:"notify_highlights"`
	SuppressEveryone     bool                     `json:"suppress_everyone"`
	SuppressRoles        bool                     `json:"suppress_roles"`
	Version              int64                    `json:"version"`
}

func (e *UserGuildSettingsUpdateEvent) AsUserGuildSettings() UserGuildSettings {
	return UserGuildSettings{
		ChannelOverrides:     e.ChannelOverrides,
		Flags:                e.Flags,
		GuildID:              e.GuildID,
		HideMutedChannels:    e.HideMutedChannels,
		MessageNotifications: e.MessageNotifications,
		MobilePush:           e.MobilePush,
		MuteScheduledEvents:  e.MuteScheduledEvents,
		Muted:                e.Muted,
		MuteConfig:           e.MuteConfig,
		NotifyHighlights:     e.NotifyHighlights,
		SuppressEveryone:     e.SuppressEveryone,
		SuppressRoles:        e.SuppressRoles,
		Version:              e.Version,
	}
}

type UserMergeOperationCompletedEvent struct {
	MergeOperationID snowflake.ID `json:"merge_operation_id"`
	SourceUserID     snowflake.ID `json:"source_user_id"`
}

func (e *UserMergeOperationCompletedEvent) AsUserMergeOperationCompleted() UserMergeOperationCompleted {
	return UserMergeOperationCompleted{
		MergeOperationID: e.MergeOperationID,
		SourceUserID:     e.SourceUserID,
	}
}

type UserNonChannelAckEvent struct {
	AckType    ReadStateType `json:"ack_type"`
	ResourceID snowflake.ID  `json:"resource_id"`
	EntityID   snowflake.ID  `json:"entity_id"`
	Version    int64         `json:"version"`
}

func (e *UserNonChannelAckEvent) AsUserNonChannelAck() UserNonChannelAck {
	return UserNonChannelAck{
		AckType:    e.AckType,
		ResourceID: e.ResourceID,
		EntityID:   e.EntityID,
		Version:    e.Version,
	}
}

type UserNoteUpdateEvent struct {
	ID   snowflake.ID `json:"id"`
	Note string       `json:"note"`
}

func (e *UserNoteUpdateEvent) AsUserNoteUpdate() UserNoteUpdate {
	return UserNoteUpdate{
		ID:   e.ID,
		Note: e.Note,
	}
}

type UserPremiumGuildSubscriptionSlotCreateEvent struct {
	ID                       snowflake.ID              `json:"id"`
	SubscriptionID           snowflake.ID              `json:"subscription_id"`
	Canceled                 bool                      `json:"canceled"`
	CooldownEndsAt           *string                   `json:"cooldown_ends_at"`
	PremiumGuildSubscription *PremiumGuildSubscription `json:"premium_guild_subscription"`
}

func (e *UserPremiumGuildSubscriptionSlotCreateEvent) AsPremiumGuildSubscriptionSlot() PremiumGuildSubscriptionSlot {
	return PremiumGuildSubscriptionSlot{
		ID:                       e.ID,
		SubscriptionID:           e.SubscriptionID,
		Canceled:                 e.Canceled,
		CooldownEndsAt:           e.CooldownEndsAt,
		PremiumGuildSubscription: e.PremiumGuildSubscription,
	}
}

type UserPremiumGuildSubscriptionSlotUpdateEvent struct {
	ID                       snowflake.ID              `json:"id"`
	SubscriptionID           snowflake.ID              `json:"subscription_id"`
	Canceled                 bool                      `json:"canceled"`
	CooldownEndsAt           *string                   `json:"cooldown_ends_at"`
	PremiumGuildSubscription *PremiumGuildSubscription `json:"premium_guild_subscription"`
}

func (e *UserPremiumGuildSubscriptionSlotUpdateEvent) AsPremiumGuildSubscriptionSlot() PremiumGuildSubscriptionSlot {
	return PremiumGuildSubscriptionSlot{
		ID:                       e.ID,
		SubscriptionID:           e.SubscriptionID,
		Canceled:                 e.Canceled,
		CooldownEndsAt:           e.CooldownEndsAt,
		PremiumGuildSubscription: e.PremiumGuildSubscription,
	}
}

type UserPremiumGuildSubscriptionSlotDeleteEvent struct {
	ID                       snowflake.ID              `json:"id"`
	SubscriptionID           snowflake.ID              `json:"subscription_id"`
	Canceled                 bool                      `json:"canceled"`
	CooldownEndsAt           *string                   `json:"cooldown_ends_at"`
	PremiumGuildSubscription *PremiumGuildSubscription `json:"premium_guild_subscription"`
}

func (e *UserPremiumGuildSubscriptionSlotDeleteEvent) AsPremiumGuildSubscriptionSlot() PremiumGuildSubscriptionSlot {
	return PremiumGuildSubscriptionSlot{
		ID:                       e.ID,
		SubscriptionID:           e.SubscriptionID,
		Canceled:                 e.Canceled,
		CooldownEndsAt:           e.CooldownEndsAt,
		PremiumGuildSubscription: e.PremiumGuildSubscription,
	}
}

type UserSettingsProtoUpdateEvent struct {
	Settings UserSettingsProto `json:"settings"`
	Partial  bool              `json:"partial"`
}

func (e *UserSettingsProtoUpdateEvent) AsUserSettingsProtoUpdate() UserSettingsProtoUpdate {
	return UserSettingsProtoUpdate{
		Settings: e.Settings,
		Partial:  e.Partial,
	}
}

type UserSettingsUpdateEvent struct {
	ActivityRestrictedGuildIDs            []snowflake.ID            `json:"activity_restricted_guild_ids"`
	ActivityJoiningRestrictedGuildIDs     []snowflake.ID            `json:"activity_joining_restricted_guild_ids"`
	AFKTimeout                            int64                     `json:"afk_timeout"`
	AllowAccessibilityDetection           bool                      `json:"allow_accessibility_detection"`
	AllowActivityPartyPrivacyFriends      bool                      `json:"allow_activity_party_privacy_friends"`
	AllowActivityPartyPrivacyVoiceChannel bool                      `json:"allow_activity_party_privacy_voice_channel"`
	AnimateEmoji                          bool                      `json:"animate_emoji"`
	AnimateStickers                       StickerAnimationOption    `json:"animate_stickers"`
	ContactSyncEnabled                    bool                      `json:"contact_sync_enabled"`
	ConvertEmoticons                      bool                      `json:"convert_emoticons"`
	CustomStatus                          *CustomStatus             `json:"custom_status"`
	DefaultGuildsRestricted               bool                      `json:"default_guilds_restricted"`
	DetectPlatformAccounts                bool                      `json:"detect_platform_accounts"`
	DeveloperMode                         bool                      `json:"developer_mode"`
	DisableGamesTab                       bool                      `json:"disable_games_tab"`
	EnableTTSCommand                      bool                      `json:"enable_tts_command"`
	ExplicitContentFilter                 ExplicitContentFilter     `json:"explicit_content_filter"`
	FriendDiscoveryFlags                  FriendDiscoveryFlags      `json:"friend_discovery_flags"`
	FriendSourceFlags                     *FriendSourceFlags        `json:"friend_source_flags"`
	GIFAutoPlay                           bool                      `json:"gif_auto_play"`
	GuildFolders                          []GuildFolder             `json:"guild_folders"`
	InlineAttachmentMedia                 bool                      `json:"inline_attachment_media"`
	InlineEmbedMedia                      bool                      `json:"inline_embed_media"`
	Locale                                string                    `json:"locale"`
	MessageDisplayCompact                 bool                      `json:"message_display_compact"`
	NativePhoneIntegrationEnabled         bool                      `json:"native_phone_integration_enabled"`
	Passwordless                          bool                      `json:"passwordless"`
	RenderEmbeds                          bool                      `json:"render_embeds"`
	RenderReactions                       bool                      `json:"render_reactions"`
	RestrictedGuilds                      []snowflake.ID            `json:"restricted_guilds"`
	ShowCurrentGame                       bool                      `json:"show_current_game"`
	SlayerSdkReceiveDMsInGame             SlayerSdkReceiveInGameDMs `json:"slayer_sdk_receive_dms_in_game"`
	SoundboardVolume                      float64                   `json:"soundboard_volume"`
	Status                                StatusType                `json:"status"`
	StreamNotificationsEnabled            bool                      `json:"stream_notifications_enabled"`
	Theme                                 Theme                     `json:"theme"`
	TimezoneOffset                        int64                     `json:"timezone_offset"`
	ViewNSFWCommands                      bool                      `json:"view_nsfw_commands"`
	ViewNSFWGuilds                        bool                      `json:"view_nsfw_guilds"`
}

func (e *UserSettingsUpdateEvent) AsUserSettings() UserSettings {
	return UserSettings{
		ActivityRestrictedGuildIDs:            e.ActivityRestrictedGuildIDs,
		ActivityJoiningRestrictedGuildIDs:     e.ActivityJoiningRestrictedGuildIDs,
		AFKTimeout:                            e.AFKTimeout,
		AllowAccessibilityDetection:           e.AllowAccessibilityDetection,
		AllowActivityPartyPrivacyFriends:      e.AllowActivityPartyPrivacyFriends,
		AllowActivityPartyPrivacyVoiceChannel: e.AllowActivityPartyPrivacyVoiceChannel,
		AnimateEmoji:                          e.AnimateEmoji,
		AnimateStickers:                       e.AnimateStickers,
		ContactSyncEnabled:                    e.ContactSyncEnabled,
		ConvertEmoticons:                      e.ConvertEmoticons,
		CustomStatus:                          e.CustomStatus,
		DefaultGuildsRestricted:               e.DefaultGuildsRestricted,
		DetectPlatformAccounts:                e.DetectPlatformAccounts,
		DeveloperMode:                         e.DeveloperMode,
		DisableGamesTab:                       e.DisableGamesTab,
		EnableTTSCommand:                      e.EnableTTSCommand,
		ExplicitContentFilter:                 e.ExplicitContentFilter,
		FriendDiscoveryFlags:                  e.FriendDiscoveryFlags,
		FriendSourceFlags:                     e.FriendSourceFlags,
		GIFAutoPlay:                           e.GIFAutoPlay,
		GuildFolders:                          e.GuildFolders,
		InlineAttachmentMedia:                 e.InlineAttachmentMedia,
		InlineEmbedMedia:                      e.InlineEmbedMedia,
		Locale:                                e.Locale,
		MessageDisplayCompact:                 e.MessageDisplayCompact,
		NativePhoneIntegrationEnabled:         e.NativePhoneIntegrationEnabled,
		Passwordless:                          e.Passwordless,
		RenderEmbeds:                          e.RenderEmbeds,
		RenderReactions:                       e.RenderReactions,
		RestrictedGuilds:                      e.RestrictedGuilds,
		ShowCurrentGame:                       e.ShowCurrentGame,
		SlayerSdkReceiveDMsInGame:             e.SlayerSdkReceiveDMsInGame,
		SoundboardVolume:                      e.SoundboardVolume,
		Status:                                e.Status,
		StreamNotificationsEnabled:            e.StreamNotificationsEnabled,
		Theme:                                 e.Theme,
		TimezoneOffset:                        e.TimezoneOffset,
		ViewNSFWCommands:                      e.ViewNSFWCommands,
		ViewNSFWGuilds:                        e.ViewNSFWGuilds,
	}
}

type AudioSettingsUpdateEvent struct {
	User   map[snowflake.ID]json.RawMessage `json:"user"`
	Stream map[snowflake.ID]json.RawMessage `json:"stream"`
}

func (e *AudioSettingsUpdateEvent) AsAudioSettingsUpdate() AudioSettingsUpdate {
	return AudioSettingsUpdate{
		User:   e.User,
		Stream: e.Stream,
	}
}

type UserPaymentBrowserCheckoutDoneEvent struct {
	LoadID                *string       `json:"load_id"`
	SKUID                 *snowflake.ID `json:"sku_id"`
	SKUSubscriptionPlanID *snowflake.ID `json:"sku_subscription_plan_id"`
}

func (e *UserPaymentBrowserCheckoutDoneEvent) AsUserPaymentBrowserCheckoutDone() UserPaymentBrowserCheckoutDone {
	return UserPaymentBrowserCheckoutDone{
		LoadID:                e.LoadID,
		SKUID:                 e.SKUID,
		SKUSubscriptionPlanID: e.SKUSubscriptionPlanID,
	}
}

type UserPaymentClientAddEvent struct {
	PurchaseTokenHash string    `json:"purchase_token_hash"`
	ExpiresAt         time.Time `json:"expires_at"`
}

func (e *UserPaymentClientAddEvent) AsUserPaymentClientAdd() UserPaymentClientAdd {
	return UserPaymentClientAdd{
		PurchaseTokenHash: e.PurchaseTokenHash,
		ExpiresAt:         e.ExpiresAt,
	}
}

type UserRequiredActionUpdateEvent struct {
	RequiredAction *RequiredActionType `json:"required_action"`
}

func (e *UserRequiredActionUpdateEvent) AsUserRequiredActionUpdate() UserRequiredActionUpdate {
	return UserRequiredActionUpdate{
		RequiredAction: e.RequiredAction,
	}
}

type VoiceStateUpdateEvent struct {
	GuildID                 *snowflake.ID `json:"guild_id,omitempty"`
	ChannelID               *snowflake.ID `json:"channel_id"`
	LobbyID                 snowflake.ID  `json:"lobby_id,omitzero"`
	UserID                  snowflake.ID  `json:"user_id"`
	Member                  GuildMember   `json:"member,omitzero"`
	SessionID               string        `json:"session_id"`
	ConnectedAt             int64         `json:"connected_at,omitzero"`
	Deaf                    bool          `json:"deaf"`
	Mute                    bool          `json:"mute"`
	SelfDeaf                bool          `json:"self_deaf"`
	SelfMute                bool          `json:"self_mute"`
	SelfStream              bool          `json:"self_stream,omitzero"`
	SelfVideo               bool          `json:"self_video"`
	Suppress                bool          `json:"suppress"`
	RequestToSpeakTimestamp *time.Time    `json:"request_to_speak_timestamp"`
	Discoverable            bool          `json:"discoverable,omitzero"`
	UserVolume              float64       `json:"user_volume,omitzero"`
}

func (e *VoiceStateUpdateEvent) AsVoiceState() VoiceState {
	return VoiceState{
		GuildID:                 e.GuildID,
		ChannelID:               e.ChannelID,
		LobbyID:                 e.LobbyID,
		UserID:                  e.UserID,
		Member:                  e.Member,
		SessionID:               e.SessionID,
		ConnectedAt:             e.ConnectedAt,
		Deaf:                    e.Deaf,
		Mute:                    e.Mute,
		SelfDeaf:                e.SelfDeaf,
		SelfMute:                e.SelfMute,
		SelfStream:              e.SelfStream,
		SelfVideo:               e.SelfVideo,
		Suppress:                e.Suppress,
		RequestToSpeakTimestamp: e.RequestToSpeakTimestamp,
		Discoverable:            e.Discoverable,
		UserVolume:              e.UserVolume,
	}
}

type VoiceServerUpdateEvent struct {
	Token     string        `json:"token"`
	GuildID   *snowflake.ID `json:"guild_id"`
	ChannelID snowflake.ID  `json:"channel_id,omitzero"`
	Endpoint  *string       `json:"endpoint"`
}

func (e *VoiceServerUpdateEvent) AsVoiceServerUpdate() VoiceServerUpdate {
	return VoiceServerUpdate{
		Token:     e.Token,
		GuildID:   e.GuildID,
		ChannelID: e.ChannelID,
		Endpoint:  e.Endpoint,
	}
}

type VoiceChannelEffectSendEvent struct {
	ChannelID     snowflake.ID                     `json:"channel_id"`
	GuildID       snowflake.ID                     `json:"guild_id"`
	UserID        snowflake.ID                     `json:"user_id"`
	AnimationType *VoiceChannelEffectAnimationType `json:"animation_type"`
	AnimationID   int64                            `json:"animation_id"`
	Emoji         *Emoji                           `json:"emoji"`
	SoundID       snowflake.ID                     `json:"sound_id,omitzero"`
	SoundVolume   float64                          `json:"sound_volume,omitzero"`
}

func (e *VoiceChannelEffectSendEvent) AsVoiceChannelEffectSend() VoiceChannelEffectSend {
	return VoiceChannelEffectSend{
		ChannelID:     e.ChannelID,
		GuildID:       e.GuildID,
		UserID:        e.UserID,
		AnimationType: e.AnimationType,
		AnimationID:   e.AnimationID,
		Emoji:         e.Emoji,
		SoundID:       e.SoundID,
		SoundVolume:   e.SoundVolume,
	}
}

type VoiceChannelStartTimeUpdateEvent struct {
	ID             snowflake.ID `json:"id"`
	GuildID        snowflake.ID `json:"guild_id"`
	VoiceStartTime *int64       `json:"voice_start_time"`
}

func (e *VoiceChannelStartTimeUpdateEvent) AsVoiceChannelStartTimeUpdate() VoiceChannelStartTimeUpdate {
	return VoiceChannelStartTimeUpdate{
		ID:             e.ID,
		GuildID:        e.GuildID,
		VoiceStartTime: e.VoiceStartTime,
	}
}

type VoiceChannelStatusUpdateEvent struct {
	ID      snowflake.ID `json:"id"`
	GuildID snowflake.ID `json:"guild_id"`
	Status  *string      `json:"status"`
}

func (e *VoiceChannelStatusUpdateEvent) AsVoiceChannelStatusUpdate() VoiceChannelStatusUpdate {
	return VoiceChannelStatusUpdate{
		ID:      e.ID,
		GuildID: e.GuildID,
		Status:  e.Status,
	}
}

type VirtualCurrencyBalanceUpdateEvent struct {
	Balance int64 `json:"balance"`
}

func (e *VirtualCurrencyBalanceUpdateEvent) AsVirtualCurrencyBalanceUpdate() VirtualCurrencyBalanceUpdate {
	return VirtualCurrencyBalanceUpdate{
		Balance: e.Balance,
	}
}

type WebhooksUpdateEvent struct {
	GuildID   snowflake.ID `json:"guild_id"`
	ChannelID snowflake.ID `json:"channel_id"`
}

func (e *WebhooksUpdateEvent) AsWebhooksUpdate() WebhooksUpdate {
	return WebhooksUpdate{
		GuildID:   e.GuildID,
		ChannelID: e.ChannelID,
	}
}

var Events = map[string]func() any{
	"ACTIVITY_INVITE_CREATE":                         func() any { return new(ActivityInviteCreateEvent) },
	"APPLICATION_COMMAND_AUTOCOMPLETE_RESPONSE":      func() any { return new(ApplicationCommandAutocompleteResponseEvent) },
	"APPLICATION_COMMAND_PERMISSIONS_UPDATE":         func() any { return new(ApplicationCommandPermissionsUpdateEvent) },
	"AUDIO_SETTINGS_UPDATE":                          func() any { return new(AudioSettingsUpdateEvent) },
	"AUTHENTICATOR_CREATE":                           func() any { return new(AuthenticatorCreateEvent) },
	"AUTHENTICATOR_DELETE":                           func() any { return new(AuthenticatorDeleteEvent) },
	"AUTHENTICATOR_UPDATE":                           func() any { return new(AuthenticatorUpdateEvent) },
	"AUTH_SESSION_CHANGE":                            func() any { return new(AuthSessionChangeEvent) },
	"AUTO_MODERATION_ACTION_EXECUTION":               func() any { return new(AutoModerationActionExecutionEvent) },
	"AUTO_MODERATION_MENTION_RAID_DETECTION":         func() any { return new(AutoModerationMentionRaidDetectionEvent) },
	"AUTO_MODERATION_RULE_CREATE":                    func() any { return new(AutoModerationRuleCreateEvent) },
	"AUTO_MODERATION_RULE_DELETE":                    func() any { return new(AutoModerationRuleDeleteEvent) },
	"AUTO_MODERATION_RULE_UPDATE":                    func() any { return new(AutoModerationRuleUpdateEvent) },
	"BILLING_POPUP_BRIDGE_CALLBACK":                  func() any { return new(BillingPopupBridgeCallbackEvent) },
	"CALL_CREATE":                                    func() any { return new(CallCreateEvent) },
	"CALL_DELETE":                                    func() any { return new(CallDeleteEvent) },
	"CALL_UPDATE":                                    func() any { return new(CallUpdateEvent) },
	"CHANNEL_CREATE":                                 func() any { return new(ChannelCreateEvent) },
	"CHANNEL_DELETE":                                 func() any { return new(ChannelDeleteEvent) },
	"CHANNEL_INFO":                                   func() any { return new(ChannelInfoEvent) },
	"CHANNEL_MEMBER_COUNT_UPDATE":                    func() any { return new(ChannelMemberCountUpdateEvent) },
	"CHANNEL_PINS_ACK":                               func() any { return new(ChannelPinsAckEvent) },
	"CHANNEL_PINS_UPDATE":                            func() any { return new(ChannelPinsUpdateEvent) },
	"CHANNEL_RECIPIENT_ADD":                          func() any { return new(ChannelRecipientAddEvent) },
	"CHANNEL_RECIPIENT_REMOVE":                       func() any { return new(ChannelRecipientRemoveEvent) },
	"CHANNEL_STATUSES":                               func() any { return new(ChannelStatusesEvent) },
	"CHANNEL_SYNC":                                   func() any { return new(ChannelSyncEvent) },
	"CHANNEL_UNREAD_UPDATE":                          func() any { return new(ChannelUnreadUpdateEvent) },
	"CHANNEL_UPDATE":                                 func() any { return new(ChannelUpdateEvent) },
	"CHANNEL_UPDATE_PARTIAL":                         func() any { return new(ChannelUpdatePartialEvent) },
	"CONSOLE_COMMAND_UPDATE":                         func() any { return new(ConsoleCommandUpdateEvent) },
	"CONVERSATION_SUMMARY_UPDATE":                    func() any { return new(ConversationSummaryUpdateEvent) },
	"CREATE_SPEED_TEST":                              func() any { return new(CreateSpeedTestEvent) },
	"CREATE_STREAM":                                  func() any { return new(CreateStreamEvent) },
	"DELETED_ENTITY_IDS":                             func() any { return new(DeletedEntityIDsEvent) },
	"DELETE_STREAM":                                  func() any { return new(DeleteStreamEvent) },
	"DM_SETTINGS_UPSELL_SHOW":                        func() any { return new(DMSettingsUpsellShowEvent) },
	"EMBEDDED_ACTIVITY_UPDATE_V2":                    func() any { return new(EmbeddedActivityUpdateV2Event) },
	"ENTITLEMENT_CREATE":                             func() any { return new(EntitlementCreateEvent) },
	"ENTITLEMENT_DELETE":                             func() any { return new(EntitlementDeleteEvent) },
	"ENTITLEMENT_UPDATE":                             func() any { return new(EntitlementUpdateEvent) },
	"EXPERIMENT_SESSION_OVERRIDE_CREATE":             func() any { return new(ExperimentSessionOverrideCreateEvent) },
	"EXPERIMENT_SESSION_OVERRIDE_DELETE":             func() any { return new(ExperimentSessionOverrideDeleteEvent) },
	"FRIEND_SUGGESTION_CREATE":                       func() any { return new(FriendSuggestionCreateEvent) },
	"FRIEND_SUGGESTION_DELETE":                       func() any { return new(FriendSuggestionDeleteEvent) },
	"GAME_INVITE_CREATE":                             func() any { return new(GameInviteCreateEvent) },
	"GAME_INVITE_DELETE":                             func() any { return new(GameInviteDeleteEvent) },
	"GAME_INVITE_DELETE_MANY":                        func() any { return new(GameInviteDeleteManyEvent) },
	"GAME_RELATIONSHIP_ADD":                          func() any { return new(GameRelationshipAddEvent) },
	"GAME_RELATIONSHIP_REMOVE":                       func() any { return new(GameRelationshipRemoveEvent) },
	"GAME_SERVER_CREATE":                             func() any { return new(GameServerCreateEvent) },
	"GAME_SERVER_DELETE":                             func() any { return new(GameServerDeleteEvent) },
	"GAME_SERVER_UPDATE":                             func() any { return new(GameServerUpdateEvent) },
	"GATEWAY_LOBBY_OBJECT":                           func() any { return new(GatewayLobbyObjectEvent) },
	"GET_DELETED_ENTITY_IDS_NOT_MATCHING_HASH":       func() any { return new(GetDeletedEntityIDsNotMatchingHashEvent) },
	"GIFT_CODE_CREATE":                               func() any { return new(GiftCodeCreateEvent) },
	"GIFT_CODE_UPDATE":                               func() any { return new(GiftCodeUpdateEvent) },
	"GUILD_APPLICATION_COMMAND_INDEX_UPDATE":         func() any { return new(GuildApplicationCommandIndexUpdateEvent) },
	"GUILD_APPLIED_BOOSTS_UPDATE":                    func() any { return new(GuildAppliedBoostsUpdateEvent) },
	"GUILD_AUDIT_LOG_ENTRY_CREATE":                   func() any { return new(GuildAuditLogEntryCreateEvent) },
	"GUILD_BAN_ADD":                                  func() any { return new(GuildBanAddEvent) },
	"GUILD_BAN_REMOVE":                               func() any { return new(GuildBanRemoveEvent) },
	"GUILD_BULK_BAN_UPDATE":                          func() any { return new(GuildBulkBanUpdateEvent) },
	"GUILD_CREATE":                                   func() any { return new(GuildCreateEvent) },
	"GUILD_DELETE":                                   func() any { return new(GuildDeleteEvent) },
	"GUILD_DIRECTORY_ENTRY_CREATE":                   func() any { return new(GuildDirectoryEntryCreateEvent) },
	"GUILD_DIRECTORY_ENTRY_DELETE":                   func() any { return new(GuildDirectoryEntryDeleteEvent) },
	"GUILD_DIRECTORY_ENTRY_UPDATE":                   func() any { return new(GuildDirectoryEntryUpdateEvent) },
	"GUILD_EMOJIS_UPDATE":                            func() any { return new(GuildEmojisUpdateEvent) },
	"GUILD_FEATURE_ACK":                              func() any { return new(GuildFeatureAckEvent) },
	"GUILD_INTEGRATIONS_UPDATE":                      func() any { return new(GuildIntegrationsUpdateEvent) },
	"GUILD_JOIN_REQUEST_CREATE":                      func() any { return new(GuildJoinRequestCreateEvent) },
	"GUILD_JOIN_REQUEST_DELETE":                      func() any { return new(GuildJoinRequestDeleteEvent) },
	"GUILD_JOIN_REQUEST_UPDATE":                      func() any { return new(GuildJoinRequestUpdateEvent) },
	"GUILD_MEMBERS_CHUNK":                            func() any { return new(GuildMembersChunkEvent) },
	"GUILD_MEMBER_ADD":                               func() any { return new(GuildMemberAddEvent) },
	"GUILD_MEMBER_REMOVE":                            func() any { return new(GuildMemberRemoveEvent) },
	"GUILD_MEMBER_UPDATE":                            func() any { return new(GuildMemberUpdateEvent) },
	"GUILD_OFFICIAL_GAME_APPLICATIONS_UPDATE":        func() any { return new(GuildOfficialGameApplicationsUpdateEvent) },
	"GUILD_POWERUP_ENTITLEMENTS_CREATE":              func() any { return new(GuildPowerupEntitlementsCreateEvent) },
	"GUILD_POWERUP_ENTITLEMENTS_DELETE":              func() any { return new(GuildPowerupEntitlementsDeleteEvent) },
	"GUILD_ROLE_CREATE":                              func() any { return new(GuildRoleCreateEvent) },
	"GUILD_ROLE_DELETE":                              func() any { return new(GuildRoleDeleteEvent) },
	"GUILD_ROLE_UPDATE":                              func() any { return new(GuildRoleUpdateEvent) },
	"GUILD_SCHEDULED_EVENT_CREATE":                   func() any { return new(GuildScheduledEventCreateEvent) },
	"GUILD_SCHEDULED_EVENT_DELETE":                   func() any { return new(GuildScheduledEventDeleteEvent) },
	"GUILD_SCHEDULED_EVENT_EXCEPTIONS_DELETE":        func() any { return new(GuildScheduledEventExceptionsDeleteEvent) },
	"GUILD_SCHEDULED_EVENT_EXCEPTION_CREATE":         func() any { return new(GuildScheduledEventExceptionCreateEvent) },
	"GUILD_SCHEDULED_EVENT_EXCEPTION_DELETE":         func() any { return new(GuildScheduledEventExceptionDeleteEvent) },
	"GUILD_SCHEDULED_EVENT_EXCEPTION_UPDATE":         func() any { return new(GuildScheduledEventExceptionUpdateEvent) },
	"GUILD_SCHEDULED_EVENT_UPDATE":                   func() any { return new(GuildScheduledEventUpdateEvent) },
	"GUILD_SCHEDULED_EVENT_USER_ADD":                 func() any { return new(GuildScheduledEventUserAddEvent) },
	"GUILD_SCHEDULED_EVENT_USER_REMOVE":              func() any { return new(GuildScheduledEventUserRemoveEvent) },
	"GUILD_SOUNDBOARD_SOUNDS_UPDATE":                 func() any { return new(GuildSoundboardSoundsUpdateEvent) },
	"GUILD_SOUNDBOARD_SOUND_CREATE":                  func() any { return new(GuildSoundboardSoundCreateEvent) },
	"GUILD_SOUNDBOARD_SOUND_DELETE":                  func() any { return new(GuildSoundboardSoundDeleteEvent) },
	"GUILD_SOUNDBOARD_SOUND_UPDATE":                  func() any { return new(GuildSoundboardSoundUpdateEvent) },
	"GUILD_STICKERS_UPDATE":                          func() any { return new(GuildStickersUpdateEvent) },
	"GUILD_UPDATE":                                   func() any { return new(GuildUpdateEvent) },
	"HELLO":                                          func() any { return new(HelloEvent) },
	"IDENTIFY":                                       func() any { return new(IdentifyEvent) },
	"INITIAL_GUILD":                                  func() any { return new(InitialGuildEvent) },
	"INTEGRATION_CREATE":                             func() any { return new(IntegrationCreateEvent) },
	"INTEGRATION_DELETE":                             func() any { return new(IntegrationDeleteEvent) },
	"INTEGRATION_UPDATE":                             func() any { return new(IntegrationUpdateEvent) },
	"INTERACTION_CREATE":                             func() any { return new(InteractionCreateEvent) },
	"INTERACTION_FAILURE":                            func() any { return new(InteractionFailureEvent) },
	"INTERACTION_IFRAME_MODAL_CREATE":                func() any { return new(InteractionIframeModalCreateEvent) },
	"INTERACTION_MODAL_CREATE":                       func() any { return new(InteractionModalCreateEvent) },
	"INTERACTION_SUCCESS":                            func() any { return new(InteractionSuccessEvent) },
	"INVITE_CREATE":                                  func() any { return new(InviteCreateEvent) },
	"INVITE_DELETE":                                  func() any { return new(InviteDeleteEvent) },
	"LAST_MESSAGES":                                  func() any { return new(LastMessagesEvent) },
	"LOBBY_DELETE":                                   func() any { return new(LobbyDeleteEvent) },
	"LOBBY_MEMBER_ADD":                               func() any { return new(LobbyMemberAddEvent) },
	"LOBBY_MEMBER_CONNECT":                           func() any { return new(LobbyMemberConnectEvent) },
	"LOBBY_MEMBER_DISCONNECT":                        func() any { return new(LobbyMemberDisconnectEvent) },
	"LOBBY_MEMBER_REMOVE":                            func() any { return new(LobbyMemberRemoveEvent) },
	"LOBBY_MEMBER_UPDATE":                            func() any { return new(LobbyMemberUpdateEvent) },
	"LOBBY_MESSAGE_CREATE":                           func() any { return new(LobbyMessageCreateEvent) },
	"LOBBY_MESSAGE_UPDATE":                           func() any { return new(LobbyMessageUpdateEvent) },
	"LOBBY_VOICE_SERVER_UPDATE":                      func() any { return new(LobbyVoiceServerUpdateEvent) },
	"MESSAGE_ACK":                                    func() any { return new(MessageAckEvent) },
	"MESSAGE_CREATE":                                 func() any { return new(MessageCreateEvent) },
	"MESSAGE_DELETE":                                 func() any { return new(MessageDeleteEvent) },
	"MESSAGE_DELETE_BULK":                            func() any { return new(MessageDeleteBulkEvent) },
	"MESSAGE_POLL_VOTE_ADD":                          func() any { return new(MessagePollVoteAddEvent) },
	"MESSAGE_POLL_VOTE_REMOVE":                       func() any { return new(MessagePollVoteRemoveEvent) },
	"MESSAGE_REACTION_ADD":                           func() any { return new(MessageReactionAddEvent) },
	"MESSAGE_REACTION_ADD_MANY":                      func() any { return new(MessageReactionAddManyEvent) },
	"MESSAGE_REACTION_REMOVE":                        func() any { return new(MessageReactionRemoveEvent) },
	"MESSAGE_REACTION_REMOVE_ALL":                    func() any { return new(MessageReactionRemoveAllEvent) },
	"MESSAGE_REACTION_REMOVE_EMOJI":                  func() any { return new(MessageReactionRemoveEmojiEvent) },
	"MESSAGE_UPDATE":                                 func() any { return new(MessageUpdateEvent) },
	"NOTIFICATION_CENTER_ITEMS_ACK":                  func() any { return new(NotificationCenterItemsAckEvent) },
	"NOTIFICATION_CENTER_ITEM_COMPLETED":             func() any { return new(NotificationCenterItemCompletedEvent) },
	"NOTIFICATION_CENTER_ITEM_CREATE":                func() any { return new(NotificationCenterItemCreateEvent) },
	"NOTIFICATION_CENTER_ITEM_DELETE":                func() any { return new(NotificationCenterItemDeleteEvent) },
	"NOTIFICATION_SETTINGS_UPDATE":                   func() any { return new(NotificationSettingsUpdateEvent) },
	"OAUTH2_TOKEN_CREATE":                            func() any { return new(OAuth2TokenCreateEvent) },
	"OAUTH2_TOKEN_DELETE":                            func() any { return new(OAuth2TokenDeleteEvent) },
	"OAUTH2_TOKEN_REVOKE":                            func() any { return new(OAuth2TokenRevokeEvent) },
	"PASSIVE_UPDATE_V1":                              func() any { return new(PassiveUpdateV1Event) },
	"PASSIVE_UPDATE_V2":                              func() any { return new(PassiveUpdateV2Event) },
	"PAYMENT_UPDATE":                                 func() any { return new(PaymentUpdateEvent) },
	"PING_LOBBY_VOICE_SERVER":                        func() any { return new(PingLobbyVoiceServerEvent) },
	"PING_STREAM_SERVER":                             func() any { return new(PingStreamServerEvent) },
	"PRESENCE_UPDATE":                                func() any { return new(PresenceUpdateEvent) },
	"QOS_HEARTBEAT":                                  func() any { return new(QoSHeartbeatEvent) },
	"QUESTS_USER_COMPLETION_UPDATE":                  func() any { return new(QuestsUserCompletionUpdateEvent) },
	"QUESTS_USER_STATUS_UPDATE":                      func() any { return new(QuestsUserStatusUpdateEvent) },
	"RATE_LIMITED":                                   func() any { return new(RateLimitedEvent) },
	"REACTION_NOTIFICATION_SENT":                     func() any { return new(ReactionNotificationSentEvent) },
	"READY":                                          func() any { return new(ReadyEvent) },
	"READY_SUPPLEMENTAL":                             func() any { return new(ReadySupplementalEvent) },
	"RECENT_MENTION_DELETE":                          func() any { return new(RecentMentionDeleteEvent) },
	"RELATIONSHIP_ADD":                               func() any { return new(RelationshipAddEvent) },
	"RELATIONSHIP_REMOVE":                            func() any { return new(RelationshipRemoveEvent) },
	"RELATIONSHIP_UPDATE":                            func() any { return new(RelationshipUpdateEvent) },
	"REQUEST_CALL_CONNECT":                           func() any { return new(RequestCallConnectEvent) },
	"REQUEST_CHANNEL_INFO":                           func() any { return new(RequestChannelInfoEvent) },
	"REQUEST_CHANNEL_MEMBER_COUNT":                   func() any { return new(RequestChannelMemberCountEvent) },
	"REQUEST_CHANNEL_STATUSES":                       func() any { return new(RequestChannelStatusesEvent) },
	"REQUEST_GUILD_MEMBERS":                          func() any { return new(RequestGuildMembersEvent) },
	"REQUEST_LAST_MESSAGES":                          func() any { return new(RequestLastMessagesEvent) },
	"REQUEST_SOUNDBOARD_SOUNDS":                      func() any { return new(RequestSoundboardSoundsEvent) },
	"RESUME":                                         func() any { return new(ResumeEvent) },
	"RESUMED":                                        func() any { return new(ResumedEvent) },
	"RESYNC_GUILD_CHANNELS":                          func() any { return new(ResyncGuildChannelsEvent) },
	"SAVED_MESSAGE_CREATE":                           func() any { return new(SavedMessageCreateEvent) },
	"SAVED_MESSAGE_DELETE":                           func() any { return new(SavedMessageDeleteEvent) },
	"SEARCH_RECENT_MEMBERS":                          func() any { return new(SearchRecentMembersEvent) },
	"SESSIONS_REPLACE":                               func() any { return new(SessionsReplaceEvent) },
	"SET_STREAM_PAUSED":                              func() any { return new(SetStreamPausedEvent) },
	"SOCIAL_LAYER_SKU_PURCHASE_ELIGIBILITY_RESPONSE": func() any { return new(SocialLayerSKUPurchaseEligibilityResponseEvent) },
	"STAGE_INSTANCE_CREATE":                          func() any { return new(StageInstanceCreateEvent) },
	"STAGE_INSTANCE_DELETE":                          func() any { return new(StageInstanceDeleteEvent) },
	"STAGE_INSTANCE_UPDATE":                          func() any { return new(StageInstanceUpdateEvent) },
	"STREAM_CREATE":                                  func() any { return new(StreamCreateEvent) },
	"STREAM_DELETE":                                  func() any { return new(StreamDeleteEvent) },
	"STREAM_SERVER_UPDATE":                           func() any { return new(StreamServerUpdateEvent) },
	"STREAM_UPDATE":                                  func() any { return new(StreamUpdateEvent) },
	"THREAD_CREATE":                                  func() any { return new(ThreadCreateEvent) },
	"THREAD_DELETE":                                  func() any { return new(ThreadDeleteEvent) },
	"THREAD_LIST_SYNC":                               func() any { return new(ThreadListSyncEvent) },
	"THREAD_MEMBERS_UPDATE":                          func() any { return new(ThreadMembersUpdateEvent) },
	"THREAD_MEMBER_UPDATE":                           func() any { return new(ThreadMemberUpdateEvent) },
	"THREAD_UPDATE":                                  func() any { return new(ThreadUpdateEvent) },
	"TYPING_START":                                   func() any { return new(TypingStartEvent) },
	"UPDATE_PRESENCE":                                func() any { return new(UpdatePresenceEvent) },
	"UPDATE_VOICE_STATE":                             func() any { return new(UpdateVoiceStateEvent) },
	"USER_APPLICATION_IDENTITY_REMOVE":               func() any { return new(UserApplicationIdentityRemoveEvent) },
	"USER_APPLICATION_IDENTITY_UPDATE":               func() any { return new(UserApplicationIdentityUpdateEvent) },
	"USER_APPLICATION_REMOVE":                        func() any { return new(UserApplicationRemoveEvent) },
	"USER_CONNECTIONS_UPDATE":                        func() any { return new(UserConnectionsUpdateEvent) },
	"USER_GUILD_SETTINGS_UPDATE":                     func() any { return new(UserGuildSettingsUpdateEvent) },
	"USER_MERGE_OPERATION_COMPLETED":                 func() any { return new(UserMergeOperationCompletedEvent) },
	"USER_NON_CHANNEL_ACK":                           func() any { return new(UserNonChannelAckEvent) },
	"USER_NOTE_UPDATE":                               func() any { return new(UserNoteUpdateEvent) },
	"USER_PAYMENT_BROWSER_CHECKOUT_DONE":             func() any { return new(UserPaymentBrowserCheckoutDoneEvent) },
	"USER_PAYMENT_CLIENT_ADD":                        func() any { return new(UserPaymentClientAddEvent) },
	"USER_PREMIUM_GUILD_SUBSCRIPTION_SLOT_CREATE":    func() any { return new(UserPremiumGuildSubscriptionSlotCreateEvent) },
	"USER_PREMIUM_GUILD_SUBSCRIPTION_SLOT_DELETE":    func() any { return new(UserPremiumGuildSubscriptionSlotDeleteEvent) },
	"USER_PREMIUM_GUILD_SUBSCRIPTION_SLOT_UPDATE":    func() any { return new(UserPremiumGuildSubscriptionSlotUpdateEvent) },
	"USER_REQUIRED_ACTION_UPDATE":                    func() any { return new(UserRequiredActionUpdateEvent) },
	"USER_SETTINGS_PROTO_UPDATE":                     func() any { return new(UserSettingsProtoUpdateEvent) },
	"USER_SETTINGS_UPDATE":                           func() any { return new(UserSettingsUpdateEvent) },
	"USER_UPDATE":                                    func() any { return new(UserUpdateEvent) },
	"VIRTUAL_CURRENCY_BALANCE_UPDATE":                func() any { return new(VirtualCurrencyBalanceUpdateEvent) },
	"VOICE_CHANNEL_EFFECT_SEND":                      func() any { return new(VoiceChannelEffectSendEvent) },
	"VOICE_CHANNEL_START_TIME_UPDATE":                func() any { return new(VoiceChannelStartTimeUpdateEvent) },
	"VOICE_CHANNEL_STATUS_UPDATE":                    func() any { return new(VoiceChannelStatusUpdateEvent) },
	"VOICE_SERVER_UPDATE":                            func() any { return new(VoiceServerUpdateEvent) },
	"VOICE_STATE_UPDATE":                             func() any { return new(VoiceStateUpdateEvent) },
	"WEBHOOKS_UPDATE":                                func() any { return new(WebhooksUpdateEvent) },
}

type Component interface {
	ComponentType() ComponentType
}

type UnknownComponent struct {
	Type ComponentType   `json:"type"`
	Raw  json.RawMessage `json:"-"`
}

func (u *UnknownComponent) ComponentType() ComponentType { return u.Type }

func (u *UnknownComponent) MarshalJSON() ([]byte, error) { return u.Raw, nil }

func (*ActionRow) ComponentType() ComponentType { return ComponentTypeActionRow }

func (*Button) ComponentType() ComponentType { return ComponentTypeButton }

func (*StringSelect) ComponentType() ComponentType { return ComponentTypeStringSelect }

func (*TextInput) ComponentType() ComponentType { return ComponentTypeTextInput }

func (*UserSelect) ComponentType() ComponentType { return ComponentTypeUserSelect }

func (*RoleSelect) ComponentType() ComponentType { return ComponentTypeRoleSelect }

func (*MentionableSelect) ComponentType() ComponentType { return ComponentTypeMentionableSelect }

func (*ChannelSelect) ComponentType() ComponentType { return ComponentTypeChannelSelect }

func (*Section) ComponentType() ComponentType { return ComponentTypeSection }

func (*TextDisplay) ComponentType() ComponentType { return ComponentTypeTextDisplay }

func (*Thumbnail) ComponentType() ComponentType { return ComponentTypeThumbnail }

func (*MediaGallery) ComponentType() ComponentType { return ComponentTypeMediaGallery }

func (*File) ComponentType() ComponentType { return ComponentTypeFile }

func (*Separator) ComponentType() ComponentType { return ComponentTypeSeparator }

func (*ContentInventoryEntry) ComponentType() ComponentType {
	return ComponentTypeContentInventoryEntry
}

func (*Container) ComponentType() ComponentType { return ComponentTypeContainer }

func (*Label) ComponentType() ComponentType { return ComponentTypeLabel }

func (*FileUpload) ComponentType() ComponentType { return ComponentTypeFileUpload }

func (*CheckpointCard) ComponentType() ComponentType { return ComponentTypeCheckpointCard }

func (*RadioGroup) ComponentType() ComponentType { return ComponentTypeRadioGroup }

func (*CheckboxGroup) ComponentType() ComponentType { return ComponentTypeCheckboxGroup }

func (*Checkbox) ComponentType() ComponentType { return ComponentTypeCheckbox }

func decodeComponent(raw json.RawMessage) (Component, error) {
	var probe struct {
		Type ComponentType `json:"type"`
	}
	if err := json.Unmarshal(raw, &probe); err != nil {
		return nil, err
	}
	var v Component
	switch probe.Type {
	case ComponentTypeActionRow:
		v = new(ActionRow)
	case ComponentTypeButton:
		v = new(Button)
	case ComponentTypeStringSelect:
		v = new(StringSelect)
	case ComponentTypeTextInput:
		v = new(TextInput)
	case ComponentTypeUserSelect:
		v = new(UserSelect)
	case ComponentTypeRoleSelect:
		v = new(RoleSelect)
	case ComponentTypeMentionableSelect:
		v = new(MentionableSelect)
	case ComponentTypeChannelSelect:
		v = new(ChannelSelect)
	case ComponentTypeSection:
		v = new(Section)
	case ComponentTypeTextDisplay:
		v = new(TextDisplay)
	case ComponentTypeThumbnail:
		v = new(Thumbnail)
	case ComponentTypeMediaGallery:
		v = new(MediaGallery)
	case ComponentTypeFile:
		v = new(File)
	case ComponentTypeSeparator:
		v = new(Separator)
	case ComponentTypeContentInventoryEntry:
		v = new(ContentInventoryEntry)
	case ComponentTypeContainer:
		v = new(Container)
	case ComponentTypeLabel:
		v = new(Label)
	case ComponentTypeFileUpload:
		v = new(FileUpload)
	case ComponentTypeCheckpointCard:
		v = new(CheckpointCard)
	case ComponentTypeRadioGroup:
		v = new(RadioGroup)
	case ComponentTypeCheckboxGroup:
		v = new(CheckboxGroup)
	case ComponentTypeCheckbox:
		v = new(Checkbox)
	default:
		return &UnknownComponent{Type: probe.Type, Raw: raw}, nil
	}
	return v, json.Unmarshal(raw, v)
}

type Components []Component

func (c *Components) UnmarshalJSON(b []byte) error {
	var raws []json.RawMessage
	if err := json.Unmarshal(b, &raws); err != nil {
		return err
	}
	out := make(Components, 0, len(raws))
	for _, raw := range raws {
		v, err := decodeComponent(raw)
		if err != nil {
			return err
		}
		out = append(out, v)
	}
	*c = out
	return nil
}

type AnyComponent struct {
	Component
}

func (a *AnyComponent) UnmarshalJSON(b []byte) error {
	v, err := decodeComponent(b)
	a.Component = v
	return err
}

func (a AnyComponent) MarshalJSON() ([]byte, error) { return json.Marshal(a.Component) }
