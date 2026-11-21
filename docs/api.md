# api

```go
import "github.com/prettylittlelies/cry/api"
```

## Index

- [type API](#type-api)
- [type Option](#type-option)

## Types

### type API

```go
type API struct {
	// contains filtered or unexported fields
}
```

#### func New

```go
func New(c *rest.Client) *API
```

#### func (*API) AcceptApplicationTesterInvitation

```go
func (a *API) AcceptApplicationTesterInvitation(ctx context.Context, query discord.AcceptApplicationTesterInvitationQuery, opts ...Option) error
```

#### func (*API) AcceptCreatorMonetizationNewTermsDemonetized

```go
func (a *API) AcceptCreatorMonetizationNewTermsDemonetized(ctx context.Context, guildID snowflake.ID, opts ...Option) error
```

#### func (*API) AcceptCreatorMonetizationTerms

```go
func (a *API) AcceptCreatorMonetizationTerms(ctx context.Context, guildID snowflake.ID, opts ...Option) error
```

#### func (*API) AcceptInvite

```go
func (a *API) AcceptInvite(ctx context.Context, inviteCode string, params discord.AcceptInviteParams, opts ...Option) (*discord.Invite, error)
```

#### func (*API) AcceptQuest

```go
func (a *API) AcceptQuest(ctx context.Context, questID snowflake.ID, params discord.AcceptQuestParams, opts ...Option) (*discord.QuestUserStatus, error)
```

#### func (*API) AcceptTeamInvite

```go
func (a *API) AcceptTeamInvite(ctx context.Context, params discord.AcceptTeamInviteParams, opts ...Option) (*discord.Team, error)
```

#### func (*API) AcknowledgeApplicationDisclosures

```go
func (a *API) AcknowledgeApplicationDisclosures(ctx context.Context, applicationID snowflake.ID, params discord.AcknowledgeApplicationDisclosuresParams, opts ...Option) (*discord.AcknowledgeApplicationDisclosuresResponse, error)
```

#### func (*API) AcknowledgeBlockedUserWarning

```go
func (a *API) AcknowledgeBlockedUserWarning(ctx context.Context, channelID snowflake.ID, opts ...Option) (json.RawMessage, error)
```

#### func (*API) AcknowledgeDMSettingsUpsellModal

```go
func (a *API) AcknowledgeDMSettingsUpsellModal(ctx context.Context, guildID snowflake.ID, opts ...Option) error
```

#### func (*API) AcknowledgeGuild

```go
func (a *API) AcknowledgeGuild(ctx context.Context, guildID snowflake.ID, opts ...Option) (json.RawMessage, error)
```

#### func (*API) AcknowledgeGuildFeature

```go
func (a *API) AcknowledgeGuildFeature(ctx context.Context, guildID snowflake.ID, readStateType string, entityID snowflake.ID, params discord.AcknowledgeGuildFeatureParams, opts ...Option) (*discord.AcknowledgeGuildFeatureResponse, error)
```

#### func (*API) AcknowledgeGuildJoinRequest

```go
func (a *API) AcknowledgeGuildJoinRequest(ctx context.Context, guildID snowflake.ID, guildJoinRequestID snowflake.ID, opts ...Option) error
```

#### func (*API) AcknowledgeMessage

```go
func (a *API) AcknowledgeMessage(ctx context.Context, channelID snowflake.ID, messageID snowflake.ID, params discord.AcknowledgeMessageParams, opts ...Option) (*discord.AcknowledgeMessageResponse, error)
```

#### func (*API) AcknowledgeNotificationCenterItem

```go
func (a *API) AcknowledgeNotificationCenterItem(ctx context.Context, notificationCenterItemID snowflake.ID, opts ...Option) error
```

#### func (*API) AcknowledgePinnedMessages

```go
func (a *API) AcknowledgePinnedMessages(ctx context.Context, channelID snowflake.ID, opts ...Option) error
```

#### func (*API) AcknowledgeSafetyWarnings

```go
func (a *API) AcknowledgeSafetyWarnings(ctx context.Context, channelID snowflake.ID, params discord.AcknowledgeSafetyWarningsParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) AcknowledgeUserFeature

```go
func (a *API) AcknowledgeUserFeature(ctx context.Context, readStateType string, entityID snowflake.ID, params discord.AcknowledgeUserFeatureParams, opts ...Option) (*discord.AcknowledgeUserFeatureResponse, error)
```

#### func (*API) AcknowledgeUserOffer

```go
func (a *API) AcknowledgeUserOffer(ctx context.Context, params discord.AcknowledgeUserOfferParams, opts ...Option) (*discord.AcknowledgeUserOfferResponse, error)
```

#### func (*API) AcknowledgeUserSurvey

```go
func (a *API) AcknowledgeUserSurvey(ctx context.Context, surveyID snowflake.ID, opts ...Option) error
```

#### func (*API) AcknowledgeUserTrialOffer

```go
func (a *API) AcknowledgeUserTrialOffer(ctx context.Context, userTrialOfferID snowflake.ID, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ActionGuildJoinRequest

```go
func (a *API) ActionGuildJoinRequest(ctx context.Context, guildID snowflake.ID, guildJoinRequestID snowflake.ID, params discord.ActionGuildJoinRequestParams, opts ...Option) (*discord.GuildJoinRequest, error)
```

#### func (*API) ActionGuildJoinRequestByID

```go
func (a *API) ActionGuildJoinRequestByID(ctx context.Context, guildID snowflake.ID, guildJoinRequestID snowflake.ID, opts ...Option) (json.RawMessage, error)
```

#### func (*API) AddApplicationTester

```go
func (a *API) AddApplicationTester(ctx context.Context, applicationID snowflake.ID, params discord.AddApplicationTesterParams, opts ...Option) (*discord.WhitelistedUser, error)
```

#### func (*API) AddChannelRecipient

```go
func (a *API) AddChannelRecipient(ctx context.Context, channelID snowflake.ID, userID snowflake.ID, params discord.AddChannelRecipientParams, opts ...Option) error
```

#### func (*API) AddGuildDiscoverySubcategory

```go
func (a *API) AddGuildDiscoverySubcategory(ctx context.Context, guildID snowflake.ID, discoveryCategoryID snowflake.ID, opts ...Option) error
```

#### func (*API) AddGuildMember

```go
func (a *API) AddGuildMember(ctx context.Context, guildID snowflake.ID, userID snowflake.ID, params discord.AddGuildMemberParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) AddGuildMemberRole

```go
func (a *API) AddGuildMemberRole(ctx context.Context, guildID snowflake.ID, userID snowflake.ID, roleID snowflake.ID, opts ...Option) error
```

#### func (*API) AddGuildPowerup

```go
func (a *API) AddGuildPowerup(ctx context.Context, guildID snowflake.ID, skuid snowflake.ID, params discord.AddGuildPowerupParams, opts ...Option) error
```

#### func (*API) AddGuildRoleMembers

```go
func (a *API) AddGuildRoleMembers(ctx context.Context, guildID snowflake.ID, roleID snowflake.ID, params discord.AddGuildRoleMembersParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) AddLobbyMember

```go
func (a *API) AddLobbyMember(ctx context.Context, lobbyID snowflake.ID, userID snowflake.ID, params discord.AddLobbyMemberParams, opts ...Option) (*discord.LobbyMember, error)
```

#### func (*API) AddPhoneNumber

```go
func (a *API) AddPhoneNumber(ctx context.Context, params discord.AddPhoneNumberParams, opts ...Option) error
```

#### func (*API) AddSafetyWarning

```go
func (a *API) AddSafetyWarning(ctx context.Context, channelID snowflake.ID, params discord.AddSafetyWarningParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) AddTeamMember

```go
func (a *API) AddTeamMember(ctx context.Context, teamID snowflake.ID, params discord.AddTeamMemberParams, opts ...Option) (*discord.TeamMember, error)
```

#### func (*API) AddThreadMember

```go
func (a *API) AddThreadMember(ctx context.Context, channelID snowflake.ID, userID snowflake.ID, query discord.AddThreadMemberQuery, opts ...Option) error
```

#### func (*API) ApplyConfettiConsumable

```go
func (a *API) ApplyConfettiConsumable(ctx context.Context, params discord.ApplyConfettiConsumableParams, opts ...Option) error
```

#### func (*API) ApplyHdStreamingConsumable

```go
func (a *API) ApplyHdStreamingConsumable(ctx context.Context, params discord.ApplyHdStreamingConsumableParams, opts ...Option) error
```

#### func (*API) AssignGuildRoleConnection

```go
func (a *API) AssignGuildRoleConnection(ctx context.Context, guildID snowflake.ID, roleID snowflake.ID, opts ...Option) error
```

#### func (*API) AuthenticationGetUniqueUsernameEligibility

```go
func (a *API) AuthenticationGetUniqueUsernameEligibility(ctx context.Context, params discord.AuthenticationGetUniqueUsernameEligibilityParams, opts ...Option) (*discord.AuthenticationGetUniqueUsernameEligibilityResponse, error)
```

#### func (*API) AuthenticationGetUniqueUsernameSuggestions

```go
func (a *API) AuthenticationGetUniqueUsernameSuggestions(ctx context.Context, query discord.GetUniqueUsernameSuggestionsQuery, opts ...Option) (*discord.AuthenticationGetUniqueUsernameSuggestionsResponse, error)
```

#### func (*API) AuthorizeIPAddress

```go
func (a *API) AuthorizeIPAddress(ctx context.Context, params discord.AuthorizeIPAddressParams, opts ...Option) error
```

#### func (*API) AuthorizeUserConnection

```go
func (a *API) AuthorizeUserConnection(ctx context.Context, connectionType string, query discord.AuthorizeUserConnectionQuery, opts ...Option) (*discord.AuthorizeUserConnectionResponse, error)
```

#### func (*API) BatchRejectMessageRequests

```go
func (a *API) BatchRejectMessageRequests(ctx context.Context, params discord.BatchRejectMessageRequestsParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) BroadcastStreamNotification

```go
func (a *API) BroadcastStreamNotification(ctx context.Context, streamKey string, opts ...Option) error
```

#### func (*API) BulkAcknowledgeNotificationCenterItems

```go
func (a *API) BulkAcknowledgeNotificationCenterItems(ctx context.Context, opts ...Option) error
```

#### func (*API) BulkActionGuildJoinRequests

```go
func (a *API) BulkActionGuildJoinRequests(ctx context.Context, guildID snowflake.ID, params discord.BulkActionGuildJoinRequestsParams, opts ...Option) error
```

#### func (*API) BulkAddRelationships

```go
func (a *API) BulkAddRelationships(ctx context.Context, params discord.BulkAddRelationshipsParams, opts ...Option) (*discord.BulkAddRelationshipsResponse, error)
```

#### func (*API) BulkDeleteMessages

```go
func (a *API) BulkDeleteMessages(ctx context.Context, channelID snowflake.ID, params discord.BulkDeleteMessagesParams, opts ...Option) error
```

#### func (*API) BulkGuildBan

```go
func (a *API) BulkGuildBan(ctx context.Context, guildID snowflake.ID, params discord.BulkGuildBanParams, opts ...Option) (*discord.BulkGuildBanResponse, error)
```

#### func (*API) BulkGuildBanV2

```go
func (a *API) BulkGuildBanV2(ctx context.Context, guildID snowflake.ID, params discord.BulkGuildBanV2Params, opts ...Option) error
```

#### func (*API) BulkLeaveGroupDMs

```go
func (a *API) BulkLeaveGroupDMs(ctx context.Context, params discord.BulkLeaveGroupDMsParams, opts ...Option) error
```

#### func (*API) BulkModifyUserGuildSettings

```go
func (a *API) BulkModifyUserGuildSettings(ctx context.Context, params discord.BulkModifyUserGuildSettingsParams, opts ...Option) ([]discord.UserGuildSettings, error)
```

#### func (*API) BulkRemoveRelationships

```go
func (a *API) BulkRemoveRelationships(ctx context.Context, query discord.BulkRemoveRelationshipsQuery, params discord.BulkRemoveRelationshipsParams, opts ...Option) error
```

#### func (*API) BulkScanExplicitMedia

```go
func (a *API) BulkScanExplicitMedia(ctx context.Context, params discord.BulkScanExplicitMediaParams, opts ...Option) error
```

#### func (*API) BulkUpdateLobbyMembers

```go
func (a *API) BulkUpdateLobbyMembers(ctx context.Context, lobbyID snowflake.ID, opts ...Option) ([]discord.LobbyMember, error)
```

#### func (*API) BulkUpdateReadStates

```go
func (a *API) BulkUpdateReadStates(ctx context.Context, params discord.BulkUpdateReadStatesParams, opts ...Option) error
```

#### func (*API) CancelConsoleCommand

```go
func (a *API) CancelConsoleCommand(ctx context.Context, connectionType string, deviceID snowflake.ID, commandID snowflake.ID, opts ...Option) error
```

#### func (*API) CancelConsoleConnectionRequest

```go
func (a *API) CancelConsoleConnectionRequest(ctx context.Context, nonce string, opts ...Option) error
```

#### func (*API) CancelRemoteAuth

```go
func (a *API) CancelRemoteAuth(ctx context.Context, params discord.CancelRemoteAuthParams, opts ...Option) error
```

#### func (*API) CheckAgeVerificationForSuspendedUser

```go
func (a *API) CheckAgeVerificationForSuspendedUser(ctx context.Context, params discord.CheckAgeVerificationForSuspendedUserParams, opts ...Option) (*discord.CheckAgeVerificationForSuspendedUserResponse, error)
```

#### func (*API) CheckSocialLayerSKUPurchaseEligibility

```go
func (a *API) CheckSocialLayerSKUPurchaseEligibility(ctx context.Context, applicationID snowflake.ID, skuid snowflake.ID, opts ...Option) (*discord.CheckSocialLayerSKUPurchaseEligibilityResponse, error)
```

#### func (*API) ClaimCheckpointAvatarDecoration

```go
func (a *API) ClaimCheckpointAvatarDecoration(ctx context.Context, opts ...Option) error
```

#### func (*API) ClaimPremiumCollectiblesProduct

```go
func (a *API) ClaimPremiumCollectiblesProduct(ctx context.Context, params discord.ClaimPremiumCollectiblesProductParams, opts ...Option) ([]discord.CollectibleProduct, error)
```

#### func (*API) ClaimPromotion

```go
func (a *API) ClaimPromotion(ctx context.Context, promotionID snowflake.ID, opts ...Option) (*discord.ClaimedPromotion, error)
```

#### func (*API) ClaimQuestCreativeReward

```go
func (a *API) ClaimQuestCreativeReward(ctx context.Context, adCreativeID snowflake.ID, params discord.ClaimQuestCreativeRewardParams, opts ...Option) error
```

#### func (*API) ClaimQuestReward

```go
func (a *API) ClaimQuestReward(ctx context.Context, questID snowflake.ID, params discord.ClaimQuestRewardParams, opts ...Option) (*discord.ClaimQuestRewardResponse, error)
```

#### func (*API) ClaimRewardCategoryProduct

```go
func (a *API) ClaimRewardCategoryProduct(ctx context.Context, params discord.ClaimRewardCategoryProductParams, opts ...Option) ([]discord.CollectibleProduct, error)
```

#### func (*API) ClaimSubscriptionPromotionReward

```go
func (a *API) ClaimSubscriptionPromotionReward(ctx context.Context, subscriptionID snowflake.ID, opts ...Option) (*discord.ClaimSubscriptionPromotionRewardResponse, error)
```

#### func (*API) ClearMentionRaidIncident

```go
func (a *API) ClearMentionRaidIncident(ctx context.Context, guildID snowflake.ID, opts ...Option) error
```

#### func (*API) CompleteGuildNewMemberAction

```go
func (a *API) CompleteGuildNewMemberAction(ctx context.Context, guildID snowflake.ID, channelID snowflake.ID, opts ...Option) (json.RawMessage, error)
```

#### func (*API) CompleteQuest

```go
func (a *API) CompleteQuest(ctx context.Context, questID snowflake.ID, opts ...Option) (*discord.QuestUserStatus, error)
```

#### func (*API) ConfirmTutorialIndicator

```go
func (a *API) ConfirmTutorialIndicator(ctx context.Context, indicator string, opts ...Option) error
```

#### func (*API) ContactSyncSettings

```go
func (a *API) ContactSyncSettings(ctx context.Context, connectionID snowflake.ID, opts ...Option) (*discord.ContactSyncSettingsResponse, error)
```

#### func (*API) CreateApplication

```go
func (a *API) CreateApplication(ctx context.Context, params discord.CreateApplicationParams, opts ...Option) (*discord.Application, error)
```

#### func (*API) CreateApplicationAsset

```go
func (a *API) CreateApplicationAsset(ctx context.Context, applicationID snowflake.ID, params discord.CreateApplicationAssetParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) CreateApplicationAttachment

```go
func (a *API) CreateApplicationAttachment(ctx context.Context, applicationID snowflake.ID, params discord.CreateApplicationAttachmentParams, opts ...Option) (*discord.CreateApplicationAttachmentResponse, error)
```

#### func (*API) CreateApplicationBot

```go
func (a *API) CreateApplicationBot(ctx context.Context, applicationID snowflake.ID, opts ...Option) (*discord.CreateApplicationBotResponse, error)
```

#### func (*API) CreateApplicationEmoji

```go
func (a *API) CreateApplicationEmoji(ctx context.Context, applicationID snowflake.ID, params discord.CreateApplicationEmojiParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) CreateApplicationEntitlement

```go
func (a *API) CreateApplicationEntitlement(ctx context.Context, applicationID snowflake.ID, params discord.CreateApplicationEntitlementParams, opts ...Option) (*discord.Entitlement, error)
```

#### func (*API) CreateApplicationGameClaim

```go
func (a *API) CreateApplicationGameClaim(ctx context.Context, applicationID snowflake.ID, gameID snowflake.ID, query discord.CreateApplicationGameClaimQuery, opts ...Option) (*discord.GameClaim, error)
```

#### func (*API) CreateApplicationGiftCodeBatch

```go
func (a *API) CreateApplicationGiftCodeBatch(ctx context.Context, applicationID snowflake.ID, params discord.CreateApplicationGiftCodeBatchParams, opts ...Option) (*discord.GiftCodeBatch, error)
```

#### func (*API) CreateApplicationManagedLink

```go
func (a *API) CreateApplicationManagedLink(ctx context.Context, applicationID snowflake.ID, params discord.CreateApplicationManagedLinkParams, opts ...Option) (*discord.ActivityLink, error)
```

#### func (*API) CreateApplicationQuickLink

```go
func (a *API) CreateApplicationQuickLink(ctx context.Context, applicationID snowflake.ID, params discord.CreateApplicationQuickLinkParams, opts ...Option) (*discord.ActivityLink, error)
```

#### func (*API) CreateApplicationWidgetConfig

```go
func (a *API) CreateApplicationWidgetConfig(ctx context.Context, applicationID snowflake.ID, params discord.CreateApplicationWidgetConfigParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) CreateBillingPopupBridge

```go
func (a *API) CreateBillingPopupBridge(ctx context.Context, paymentSourceType string, opts ...Option) (*discord.CreateBillingPopupBridgeResponse, error)
```

#### func (*API) CreateBillingPopupBridgeCallback

```go
func (a *API) CreateBillingPopupBridgeCallback(ctx context.Context, paymentSourceType string, params discord.CreateBillingPopupBridgeCallbackParams, opts ...Option) error
```

#### func (*API) CreateBillingPopupBridgeRedirect

```go
func (a *API) CreateBillingPopupBridgeRedirect(ctx context.Context, paymentSourceType string, state string, responseType string, opts ...Option) (json.RawMessage, error)
```

#### func (*API) CreateChannelInvite

```go
func (a *API) CreateChannelInvite(ctx context.Context, channelID snowflake.ID, params discord.CreateChannelInviteParams, opts ...Option) (*discord.Invite, error)
```

#### func (*API) CreateChannelTag

```go
func (a *API) CreateChannelTag(ctx context.Context, channelID snowflake.ID, params discord.CreateChannelTagParams, opts ...Option) (*discord.Channel, error)
```

#### func (*API) CreateChurnUserOffer

```go
func (a *API) CreateChurnUserOffer(ctx context.Context, opts ...Option) (*discord.CreateChurnUserOfferResponse, error)
```

#### func (*API) CreateCompany

```go
func (a *API) CreateCompany(ctx context.Context, teamID snowflake.ID, params discord.CreateCompanyParams, opts ...Option) (*discord.TeamCompany, error)
```

#### func (*API) CreateConsoleConnection

```go
func (a *API) CreateConsoleConnection(ctx context.Context, params discord.CreateConsoleConnectionParams, opts ...Option) (*discord.CreateConsoleConnectionResponse, error)
```

#### func (*API) CreateContactSyncConnection

```go
func (a *API) CreateContactSyncConnection(ctx context.Context, connectionID snowflake.ID, params discord.CreateContactSyncConnectionParams, opts ...Option) (*discord.Connection, error)
```

#### func (*API) CreateCreatorMonetizationEnableRequest

```go
func (a *API) CreateCreatorMonetizationEnableRequest(ctx context.Context, guildID snowflake.ID, opts ...Option) error
```

#### func (*API) CreateDirectoryEntry

```go
func (a *API) CreateDirectoryEntry(ctx context.Context, channelID snowflake.ID, entityID snowflake.ID, params discord.CreateDirectoryEntryParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) CreateDomainConnection

```go
func (a *API) CreateDomainConnection(ctx context.Context, connectionID snowflake.ID, opts ...Option) (*discord.Connection, error)
```

#### func (*API) CreateFingerprint

```go
func (a *API) CreateFingerprint(ctx context.Context, opts ...Option) (*discord.CreateFingerprintResponse, error)
```

#### func (*API) CreateFollowupMessage

```go
func (a *API) CreateFollowupMessage(ctx context.Context, applicationID snowflake.ID, interactionToken string, opts ...Option) (json.RawMessage, error)
```

#### func (*API) CreateGameInvite

```go
func (a *API) CreateGameInvite(ctx context.Context, params discord.CreateGameInviteParams, opts ...Option) (*discord.CreateGameInviteResponse, error)
```

#### func (*API) CreateGameRelationship

```go
func (a *API) CreateGameRelationship(ctx context.Context, userID snowflake.ID, opts ...Option) error
```

#### func (*API) CreateGameRelationshipByApplication

```go
func (a *API) CreateGameRelationshipByApplication(ctx context.Context, userID snowflake.ID, applicationID snowflake.ID, opts ...Option) error
```

#### func (*API) CreateGravityAttachments

```go
func (a *API) CreateGravityAttachments(ctx context.Context, params discord.CreateGravityAttachmentsParams, opts ...Option) ([]discord.CloudAttachment, error)
```

#### func (*API) CreateGreetMessage

```go
func (a *API) CreateGreetMessage(ctx context.Context, channelID snowflake.ID, params discord.CreateGreetMessageParams, opts ...Option) (*discord.Message, error)
```

#### func (*API) CreateGroupDMShell

```go
func (a *API) CreateGroupDMShell(ctx context.Context, params discord.CreateGroupDMShellParams, opts ...Option) (*discord.Channel, error)
```

#### func (*API) CreateGuild

```go
func (a *API) CreateGuild(ctx context.Context, params discord.CreateGuildParams, opts ...Option) (*discord.Guild, error)
```

#### func (*API) CreateGuildAutomodRule

```go
func (a *API) CreateGuildAutomodRule(ctx context.Context, guildID snowflake.ID, params discord.CreateGuildAutomodRuleParams, opts ...Option) (*discord.AutomodRule, error)
```

#### func (*API) CreateGuildBan

```go
func (a *API) CreateGuildBan(ctx context.Context, guildID snowflake.ID, userID snowflake.ID, params discord.CreateGuildBanParams, opts ...Option) error
```

#### func (*API) CreateGuildChannel

```go
func (a *API) CreateGuildChannel(ctx context.Context, guildID snowflake.ID, params discord.CreateGuildChannelParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) CreateGuildEmoji

```go
func (a *API) CreateGuildEmoji(ctx context.Context, guildID snowflake.ID, params discord.CreateGuildEmojiParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) CreateGuildIntegration

```go
func (a *API) CreateGuildIntegration(ctx context.Context, guildID snowflake.ID, params discord.CreateGuildIntegrationParams, opts ...Option) error
```

#### func (*API) CreateGuildJoinRequest

```go
func (a *API) CreateGuildJoinRequest(ctx context.Context, guildID snowflake.ID, params discord.CreateGuildJoinRequestParams, opts ...Option) (*discord.GuildJoinRequest, error)
```

#### func (*API) CreateGuildJoinRequestInterview

```go
func (a *API) CreateGuildJoinRequestInterview(ctx context.Context, guildJoinRequestID snowflake.ID, opts ...Option) (json.RawMessage, error)
```

#### func (*API) CreateGuildOnboardingResponses

```go
func (a *API) CreateGuildOnboardingResponses(ctx context.Context, guildID snowflake.ID, params discord.CreateGuildOnboardingResponsesParams, opts ...Option) (*discord.OnboardingResponses, error)
```

#### func (*API) CreateGuildProductAttachments

```go
func (a *API) CreateGuildProductAttachments(ctx context.Context, channelID snowflake.ID, params discord.CreateGuildProductAttachmentsParams, opts ...Option) ([]discord.CloudAttachment, error)
```

#### func (*API) CreateGuildProductListing

```go
func (a *API) CreateGuildProductListing(ctx context.Context, guildID snowflake.ID, params discord.CreateGuildProductListingParams, opts ...Option) (*discord.GuildProductListing, error)
```

#### func (*API) CreateGuildRole

```go
func (a *API) CreateGuildRole(ctx context.Context, guildID snowflake.ID, params discord.CreateGuildRoleParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) CreateGuildRoleSubscriptionGroupListing

```go
func (a *API) CreateGuildRoleSubscriptionGroupListing(ctx context.Context, guildID snowflake.ID, params discord.CreateGuildRoleSubscriptionGroupListingParams, opts ...Option) (*discord.RoleSubscriptionGroupListing, error)
```

#### func (*API) CreateGuildScheduledEvent

```go
func (a *API) CreateGuildScheduledEvent(ctx context.Context, guildID snowflake.ID, params discord.GuildScheduledEventCreateGuildScheduledEventParams, opts ...Option) (*discord.GuildScheduledEvent, error)
```

#### func (*API) CreateGuildScheduledEventUser

```go
func (a *API) CreateGuildScheduledEventUser(ctx context.Context, guildID snowflake.ID, guildScheduledEventID snowflake.ID, opts ...Option) (*discord.GuildScheduledEventUser, error)
```

#### func (*API) CreateGuildSoundboardSound

```go
func (a *API) CreateGuildSoundboardSound(ctx context.Context, guildID snowflake.ID, params discord.CreateGuildSoundboardSoundParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) CreateGuildSticker

```go
func (a *API) CreateGuildSticker(ctx context.Context, guildID snowflake.ID, params discord.CreateGuildStickerParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) CreateGuildTemplate

```go
func (a *API) CreateGuildTemplate(ctx context.Context, guildID snowflake.ID, params discord.CreateGuildTemplateParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) CreateHandoffToken

```go
func (a *API) CreateHandoffToken(ctx context.Context, params discord.CreateHandoffTokenParams, opts ...Option) (*discord.CreateHandoffTokenResponse, error)
```

#### func (*API) CreateHeadlessSession

```go
func (a *API) CreateHeadlessSession(ctx context.Context, params discord.CreateHeadlessSessionParams, opts ...Option) (*discord.CreateHeadlessSessionResponse, error)
```

#### func (*API) CreateInteraction

```go
func (a *API) CreateInteraction(ctx context.Context, params discord.CreateInteractionParams, opts ...Option) error
```

#### func (*API) CreateInteractionResponse

```go
func (a *API) CreateInteractionResponse(ctx context.Context, interactionID snowflake.ID, interactionToken string, query discord.CreateInteractionResponseQuery, opts ...Option) (*discord.CreateInteractionResponseResponse, error)
```

#### func (*API) CreateLinkedUserRequest

```go
func (a *API) CreateLinkedUserRequest(ctx context.Context, params discord.CreateLinkedUserRequestParams, opts ...Option) (*discord.LinkedUser, error)
```

#### func (*API) CreateLobby

```go
func (a *API) CreateLobby(ctx context.Context, params discord.CreateLobbyParams, opts ...Option) (*discord.Lobby, error)
```

#### func (*API) CreateLobbyInvite

```go
func (a *API) CreateLobbyInvite(ctx context.Context, lobbyID snowflake.ID, userID snowflake.ID, opts ...Option) (*discord.CreateLobbyInviteResponse, error)
```

#### func (*API) CreateLobbyInviteForCurrentUser

```go
func (a *API) CreateLobbyInviteForCurrentUser(ctx context.Context, lobbyID snowflake.ID, opts ...Option) (*discord.CreateLobbyInviteForCurrentUserResponse, error)
```

#### func (*API) CreateLobbyMessage

```go
func (a *API) CreateLobbyMessage(ctx context.Context, lobbyID snowflake.ID, opts ...Option) (*discord.Message, error)
```

#### func (*API) CreateMessage

```go
func (a *API) CreateMessage(ctx context.Context, channelID snowflake.ID, params discord.CreateMessageParams, opts ...Option) (*discord.Message, error)
```

#### func (*API) CreateMessageAttachments

```go
func (a *API) CreateMessageAttachments(ctx context.Context, channelID snowflake.ID, params discord.CreateMessageAttachmentsParams, opts ...Option) ([]discord.CloudAttachment, error)
```

#### func (*API) CreateNotificationSettingsSnapshot

```go
func (a *API) CreateNotificationSettingsSnapshot(ctx context.Context, params discord.CreateNotificationSettingsSnapshotParams, opts ...Option) ([]discord.NotificationSettingsSnapshot, error)
```

#### func (*API) CreateOAuth2ApplicationAsset

```go
func (a *API) CreateOAuth2ApplicationAsset(ctx context.Context, applicationID snowflake.ID, params discord.CreateOAuth2ApplicationAssetParams, opts ...Option) (*discord.ApplicationAsset, error)
```

#### func (*API) CreateOAuth2Authorization

```go
func (a *API) CreateOAuth2Authorization(ctx context.Context, query discord.CreateOAuth2AuthorizationQuery, params discord.CreateOAuth2AuthorizationParams, opts ...Option) (*discord.CreateOAuth2AuthorizationResponse, error)
```

#### func (*API) CreatePaymentSource

```go
func (a *API) CreatePaymentSource(ctx context.Context, query discord.CreatePaymentSourceQuery, params discord.CreatePaymentSourceParams, opts ...Option) (*discord.PaymentSource, error)
```

#### func (*API) CreatePaypalBillingAgreementToken

```go
func (a *API) CreatePaypalBillingAgreementToken(ctx context.Context, params discord.CreatePaypalBillingAgreementTokenParams, opts ...Option) (*discord.CreatePaypalBillingAgreementTokenResponse, error)
```

#### func (*API) CreatePollVote

```go
func (a *API) CreatePollVote(ctx context.Context, channelID snowflake.ID, messageID snowflake.ID, params discord.CreatePollVoteParams, opts ...Option) error
```

#### func (*API) CreatePremiumGuildSubscriptions

```go
func (a *API) CreatePremiumGuildSubscriptions(ctx context.Context, guildID snowflake.ID, params discord.CreatePremiumGuildSubscriptionsParams, opts ...Option) ([]discord.PremiumGuildSubscription, error)
```

#### func (*API) CreatePremiumReferral

```go
func (a *API) CreatePremiumReferral(ctx context.Context, userID snowflake.ID, opts ...Option) (*discord.PremiumReferral, error)
```

#### func (*API) CreatePrivateChannel

```go
func (a *API) CreatePrivateChannel(ctx context.Context, params discord.CreatePrivateChannelParams, opts ...Option) (*discord.Channel, error)
```

#### func (*API) CreateReaction

```go
func (a *API) CreateReaction(ctx context.Context, channelID snowflake.ID, messageID snowflake.ID, emoji string, query discord.CreateReactionQuery, opts ...Option) error
```

#### func (*API) CreateRelationship

```go
func (a *API) CreateRelationship(ctx context.Context, userID snowflake.ID, params discord.CreateRelationshipParams, opts ...Option) error
```

#### func (*API) CreateRemoteAuthSession

```go
func (a *API) CreateRemoteAuthSession(ctx context.Context, params discord.CreateRemoteAuthSessionParams, opts ...Option) (*discord.CreateRemoteAuthSessionResponse, error)
```

#### func (*API) CreateReport

```go
func (a *API) CreateReport(ctx context.Context, params discord.CreateReportParams, opts ...Option) (*discord.CreateReportResponse, error)
```

#### func (*API) CreateSKU

```go
func (a *API) CreateSKU(ctx context.Context, params discord.CreateSKUParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) CreateSKUPurchase

```go
func (a *API) CreateSKUPurchase(ctx context.Context, skuid snowflake.ID, params discord.StoreCreateSKUPurchaseParams, opts ...Option) (*discord.CreateSKUPurchaseResponse, error)
```

#### func (*API) CreateSocialLayerSKU

```go
func (a *API) CreateSocialLayerSKU(ctx context.Context, applicationID snowflake.ID, params discord.CreateSocialLayerSKUParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) CreateStageInstance

```go
func (a *API) CreateStageInstance(ctx context.Context, params discord.CreateStageInstanceParams, opts ...Option) (*discord.StageInstance, error)
```

#### func (*API) CreateStagedReport

```go
func (a *API) CreateStagedReport(ctx context.Context, params discord.CreateStagedReportParams, opts ...Option) error
```

#### func (*API) CreateStoreListing

```go
func (a *API) CreateStoreListing(ctx context.Context, query discord.CreateStoreListingQuery, params discord.StoreCreateStoreListingParams, opts ...Option) (*discord.StoreListing, error)
```

#### func (*API) CreateSubscription

```go
func (a *API) CreateSubscription(ctx context.Context, params discord.CreateSubscriptionParams, opts ...Option) (*discord.Subscription, error)
```

#### func (*API) CreateSubscriptionPreview

```go
func (a *API) CreateSubscriptionPreview(ctx context.Context, params discord.CreateSubscriptionPreviewParams, opts ...Option) (*discord.Invoice, error)
```

#### func (*API) CreateTeam

```go
func (a *API) CreateTeam(ctx context.Context, params discord.CreateTeamParams, opts ...Option) (*discord.Team, error)
```

#### func (*API) CreateTeamIdentityVerification

```go
func (a *API) CreateTeamIdentityVerification(ctx context.Context, teamID snowflake.ID, params discord.CreateTeamIdentityVerificationParams, opts ...Option) (*discord.UserIdentityVerification, error)
```

#### func (*API) CreateThread

```go
func (a *API) CreateThread(ctx context.Context, channelID snowflake.ID, params discord.CreateThreadParams, opts ...Option) (*discord.Channel, error)
```

#### func (*API) CreateThreadFromMessage

```go
func (a *API) CreateThreadFromMessage(ctx context.Context, channelID snowflake.ID, messageID snowflake.ID, params discord.CreateThreadFromMessageParams, opts ...Option) (*discord.Channel, error)
```

#### func (*API) CreateUniqueUsername

```go
func (a *API) CreateUniqueUsername(ctx context.Context, params discord.CreateUniqueUsernameParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) CreateUserConnectionCallback

```go
func (a *API) CreateUserConnectionCallback(ctx context.Context, connectionType string, params discord.CreateUserConnectionCallbackParams, opts ...Option) (*discord.Connection, error)
```

#### func (*API) CreateUserGiftCode

```go
func (a *API) CreateUserGiftCode(ctx context.Context, params discord.CreateUserGiftCodeParams, opts ...Option) (*discord.GiftCode, error)
```

#### func (*API) CreateUserHarvest

```go
func (a *API) CreateUserHarvest(ctx context.Context, params discord.CreateUserHarvestParams, opts ...Option) (*discord.Harvest, error)
```

#### func (*API) CreateUserIdentityVerification

```go
func (a *API) CreateUserIdentityVerification(ctx context.Context, params discord.CreateUserIdentityVerificationParams, opts ...Option) (*discord.UserIdentityVerification, error)
```

#### func (*API) CreateUserInvite

```go
func (a *API) CreateUserInvite(ctx context.Context, params discord.CreateUserInviteParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) CreateVideoFilterAsset

```go
func (a *API) CreateVideoFilterAsset(ctx context.Context, params discord.CreateVideoFilterAssetParams, opts ...Option) (*discord.VideoFilterAsset, error)
```

#### func (*API) CreateWebauthnAuthenticator

```go
func (a *API) CreateWebauthnAuthenticator(ctx context.Context, params discord.CreateWebauthnAuthenticatorParams, opts ...Option) (*discord.CreateWebauthnAuthenticatorResponse, error)
```

#### func (*API) CreateWebhook

```go
func (a *API) CreateWebhook(ctx context.Context, channelID snowflake.ID, params discord.CreateWebhookParams, opts ...Option) (*discord.Webhook, error)
```

#### func (*API) CrosspostMessage

```go
func (a *API) CrosspostMessage(ctx context.Context, channelID snowflake.ID, messageID snowflake.ID, opts ...Option) (*discord.Message, error)
```

#### func (*API) DeleteAllReactions

```go
func (a *API) DeleteAllReactions(ctx context.Context, channelID snowflake.ID, messageID snowflake.ID, opts ...Option) error
```

#### func (*API) DeleteApplication

```go
func (a *API) DeleteApplication(ctx context.Context, applicationID snowflake.ID, opts ...Option) error
```

#### func (*API) DeleteApplicationAsset

```go
func (a *API) DeleteApplicationAsset(ctx context.Context, applicationID snowflake.ID, applicationAssetKey string, opts ...Option) error
```

#### func (*API) DeleteApplicationEmoji

```go
func (a *API) DeleteApplicationEmoji(ctx context.Context, applicationID snowflake.ID, emojiID snowflake.ID, opts ...Option) error
```

#### func (*API) DeleteApplicationGameClaim

```go
func (a *API) DeleteApplicationGameClaim(ctx context.Context, applicationID snowflake.ID, gameClaimID snowflake.ID, opts ...Option) error
```

#### func (*API) DeleteApplicationManagedLink

```go
func (a *API) DeleteApplicationManagedLink(ctx context.Context, applicationID snowflake.ID, activityLinkID snowflake.ID, opts ...Option) error
```

#### func (*API) DeleteApplicationStorefront

```go
func (a *API) DeleteApplicationStorefront(ctx context.Context, applicationID snowflake.ID, opts ...Option) error
```

#### func (*API) DeleteApplicationWidgetConfig

```go
func (a *API) DeleteApplicationWidgetConfig(ctx context.Context, applicationID snowflake.ID, widgetConfigID snowflake.ID, opts ...Option) error
```

#### func (*API) DeleteAttachment

```go
func (a *API) DeleteAttachment(ctx context.Context, cloudAttachmentUploadFilename string, opts ...Option) error
```

#### func (*API) DeleteChannel

```go
func (a *API) DeleteChannel(ctx context.Context, channelID snowflake.ID, query discord.DeleteChannelQuery, opts ...Option) (*discord.Channel, error)
```

#### func (*API) DeleteChannelIntegration

```go
func (a *API) DeleteChannelIntegration(ctx context.Context, channelID snowflake.ID, integrationID snowflake.ID, opts ...Option) error
```

#### func (*API) DeleteChannelPermission

```go
func (a *API) DeleteChannelPermission(ctx context.Context, channelID snowflake.ID, overwriteID snowflake.ID, opts ...Option) error
```

#### func (*API) DeleteChannelTag

```go
func (a *API) DeleteChannelTag(ctx context.Context, channelID snowflake.ID, forumTagID snowflake.ID, opts ...Option) (*discord.Channel, error)
```

#### func (*API) DeleteConversationSummary

```go
func (a *API) DeleteConversationSummary(ctx context.Context, channelID snowflake.ID, summaryID snowflake.ID, opts ...Option) error
```

#### func (*API) DeleteDMMessage

```go
func (a *API) DeleteDMMessage(ctx context.Context, userID snowflake.ID, messageID snowflake.ID, opts ...Option) error
```

#### func (*API) DeleteDirectoryEntry

```go
func (a *API) DeleteDirectoryEntry(ctx context.Context, channelID snowflake.ID, entityID snowflake.ID, opts ...Option) error
```

#### func (*API) DeleteExample

```go
func (a *API) DeleteExample(ctx context.Context, opts ...Option) (json.RawMessage, error)
```

#### func (*API) DeleteFollowupMessage

```go
func (a *API) DeleteFollowupMessage(ctx context.Context, applicationID snowflake.ID, interactionToken string, messageID snowflake.ID, opts ...Option) (json.RawMessage, error)
```

#### func (*API) DeleteGameInvite

```go
func (a *API) DeleteGameInvite(ctx context.Context, gameInviteInviteID snowflake.ID, opts ...Option) error
```

#### func (*API) DeleteGameInvites

```go
func (a *API) DeleteGameInvites(ctx context.Context, opts ...Option) error
```

#### func (*API) DeleteGuild

```go
func (a *API) DeleteGuild(ctx context.Context, guildID snowflake.ID, opts ...Option) error
```

#### func (*API) DeleteGuildBan

```go
func (a *API) DeleteGuildBan(ctx context.Context, guildID snowflake.ID, userID snowflake.ID, opts ...Option) error
```

#### func (*API) DeleteGuildEmoji

```go
func (a *API) DeleteGuildEmoji(ctx context.Context, guildID snowflake.ID, emojiID snowflake.ID, opts ...Option) error
```

#### func (*API) DeleteGuildIntegration

```go
func (a *API) DeleteGuildIntegration(ctx context.Context, guildID snowflake.ID, integrationID snowflake.ID, opts ...Option) error
```

#### func (*API) DeleteGuildJoinRequest

```go
func (a *API) DeleteGuildJoinRequest(ctx context.Context, guildID snowflake.ID, opts ...Option) error
```

#### func (*API) DeleteGuildProductListing

```go
func (a *API) DeleteGuildProductListing(ctx context.Context, guildID snowflake.ID, productListingID snowflake.ID, opts ...Option) (*discord.DeleteGuildProductListingResponse, error)
```

#### func (*API) DeleteGuildRole

```go
func (a *API) DeleteGuildRole(ctx context.Context, guildID snowflake.ID, roleID snowflake.ID, opts ...Option) error
```

#### func (*API) DeleteGuildRoleSubscriptionGroupListing

```go
func (a *API) DeleteGuildRoleSubscriptionGroupListing(ctx context.Context, guildID snowflake.ID, subscriptionGroupListingID snowflake.ID, params discord.DeleteGuildRoleSubscriptionGroupListingParams, opts ...Option) (*discord.RoleSubscriptionListing, error)
```

#### func (*API) DeleteGuildScheduledEventUser

```go
func (a *API) DeleteGuildScheduledEventUser(ctx context.Context, guildID snowflake.ID, guildScheduledEventID snowflake.ID, query discord.DeleteGuildScheduledEventUserQuery, params discord.DeleteGuildScheduledEventUserParams, opts ...Option) ([]discord.GuildScheduledEventUser, error)
```

#### func (*API) DeleteGuildSoundboardSound

```go
func (a *API) DeleteGuildSoundboardSound(ctx context.Context, guildID snowflake.ID, soundID snowflake.ID, opts ...Option) error
```

#### func (*API) DeleteGuildSticker

```go
func (a *API) DeleteGuildSticker(ctx context.Context, guildID snowflake.ID, stickerID snowflake.ID, opts ...Option) error
```

#### func (*API) DeleteGuildTemplate

```go
func (a *API) DeleteGuildTemplate(ctx context.Context, guildID snowflake.ID, templateCode string, opts ...Option) (json.RawMessage, error)
```

#### func (*API) DeleteHeadlessSession

```go
func (a *API) DeleteHeadlessSession(ctx context.Context, params discord.DeleteHeadlessSessionParams, opts ...Option) error
```

#### func (*API) DeleteInvite

```go
func (a *API) DeleteInvite(ctx context.Context, inviteCode string, opts ...Option) (*discord.Invite, error)
```

#### func (*API) DeleteLobby

```go
func (a *API) DeleteLobby(ctx context.Context, lobbyID snowflake.ID, opts ...Option) error
```

#### func (*API) DeleteMessage

```go
func (a *API) DeleteMessage(ctx context.Context, channelID snowflake.ID, messageID snowflake.ID, opts ...Option) error
```

#### func (*API) DeleteNotificationCenterItem

```go
func (a *API) DeleteNotificationCenterItem(ctx context.Context, notificationCenterItemID snowflake.ID, params discord.DeleteNotificationCenterItemParams, opts ...Option) error
```

#### func (*API) DeleteNotificationSettingsSnapshot

```go
func (a *API) DeleteNotificationSettingsSnapshot(ctx context.Context, notificationSettingsSnapshotID snowflake.ID, opts ...Option) ([]discord.NotificationSettingsSnapshot, error)
```

#### func (*API) DeleteOAuth2ApplicationAsset

```go
func (a *API) DeleteOAuth2ApplicationAsset(ctx context.Context, applicationID snowflake.ID, oAuth2ApplicationAssetID snowflake.ID, opts ...Option) error
```

#### func (*API) DeleteOAuth2Authorization

```go
func (a *API) DeleteOAuth2Authorization(ctx context.Context, tokenID snowflake.ID, opts ...Option) error
```

#### func (*API) DeleteOriginalInteractionResponse

```go
func (a *API) DeleteOriginalInteractionResponse(ctx context.Context, applicationID snowflake.ID, interactionToken string, opts ...Option) error
```

#### func (*API) DeleteOwnReaction

```go
func (a *API) DeleteOwnReaction(ctx context.Context, channelID, messageID snowflake.ID, emoji string, opts ...Option) error
```

#### func (*API) DeletePaymentSource

```go
func (a *API) DeletePaymentSource(ctx context.Context, paymentSourceID snowflake.ID, opts ...Option) error
```

#### func (*API) DeletePremiumGuildSubscription

```go
func (a *API) DeletePremiumGuildSubscription(ctx context.Context, guildID snowflake.ID, premiumGuildSubscriptionID snowflake.ID, opts ...Option) error
```

#### func (*API) DeleteReactionEmoji

```go
func (a *API) DeleteReactionEmoji(ctx context.Context, channelID snowflake.ID, messageID snowflake.ID, emoji string, opts ...Option) error
```

#### func (*API) DeleteReadState

```go
func (a *API) DeleteReadState(ctx context.Context, readStateID snowflake.ID, params discord.DeleteReadStateParams, opts ...Option) error
```

#### func (*API) DeleteRecentAvatar

```go
func (a *API) DeleteRecentAvatar(ctx context.Context, avatarID snowflake.ID, opts ...Option) error
```

#### func (*API) DeleteRecentMention

```go
func (a *API) DeleteRecentMention(ctx context.Context, messageID snowflake.ID, opts ...Option) error
```

#### func (*API) DeleteSafetyWarnings

```go
func (a *API) DeleteSafetyWarnings(ctx context.Context, channelID snowflake.ID, opts ...Option) (json.RawMessage, error)
```

#### func (*API) DeleteStageInstance

```go
func (a *API) DeleteStageInstance(ctx context.Context, channelID snowflake.ID, opts ...Option) error
```

#### func (*API) DeleteSubscription

```go
func (a *API) DeleteSubscription(ctx context.Context, subscriptionID snowflake.ID, query discord.DeleteSubscriptionQuery, opts ...Option) error
```

#### func (*API) DeleteTeam

```go
func (a *API) DeleteTeam(ctx context.Context, teamID snowflake.ID, opts ...Option) error
```

#### func (*API) DeleteUserAccount

```go
func (a *API) DeleteUserAccount(ctx context.Context, params discord.DeleteUserAccountParams, opts ...Option) error
```

#### func (*API) DeleteUserConnection

```go
func (a *API) DeleteUserConnection(ctx context.Context, connectionType string, connectionID snowflake.ID, opts ...Option) error
```

#### func (*API) DeleteUserReaction

```go
func (a *API) DeleteUserReaction(ctx context.Context, channelID, messageID snowflake.ID, emoji string, userID snowflake.ID, opts ...Option) error
```

#### func (*API) DeleteVideoFilterAsset

```go
func (a *API) DeleteVideoFilterAsset(ctx context.Context, videoFilterAssetID snowflake.ID, opts ...Option) error
```

#### func (*API) DeleteWebauthnAuthenticator

```go
func (a *API) DeleteWebauthnAuthenticator(ctx context.Context, authenticatorID snowflake.ID, opts ...Option) error
```

#### func (*API) DeleteWebhook

```go
func (a *API) DeleteWebhook(ctx context.Context, webhookID snowflake.ID, opts ...Option) error
```

#### func (*API) DeleteWebhookMessage

```go
func (a *API) DeleteWebhookMessage(ctx context.Context, webhookID snowflake.ID, webhookToken string, messageID snowflake.ID, opts ...Option) error
```

#### func (*API) DeleteWebhookWithToken

```go
func (a *API) DeleteWebhookWithToken(ctx context.Context, webhookID snowflake.ID, webhookToken string, opts ...Option) (json.RawMessage, error)
```

#### func (*API) DisableSMSMFA

```go
func (a *API) DisableSMSMFA(ctx context.Context, params discord.DisableSMSMFAParams, opts ...Option) error
```

#### func (*API) DisableTotpMFA

```go
func (a *API) DisableTotpMFA(ctx context.Context, opts ...Option) (*discord.DisableTotpMFAResponse, error)
```

#### func (*API) DisableUserAccount

```go
func (a *API) DisableUserAccount(ctx context.Context, params discord.DisableUserAccountParams, opts ...Option) error
```

#### func (*API) DismissQuestContent

```go
func (a *API) DismissQuestContent(ctx context.Context, questID snowflake.ID, questContentType string, opts ...Option) (*discord.QuestUserStatus, error)
```

#### func (*API) EditDMMessage

```go
func (a *API) EditDMMessage(ctx context.Context, userID snowflake.ID, messageID snowflake.ID, params discord.EditDMMessageParams, opts ...Option) (*discord.Message, error)
```

#### func (*API) EditMessage

```go
func (a *API) EditMessage(ctx context.Context, channelID snowflake.ID, messageID snowflake.ID, params discord.EditMessageParams, opts ...Option) (*discord.Message, error)
```

#### func (*API) EditWebhookMessage

```go
func (a *API) EditWebhookMessage(ctx context.Context, webhookID snowflake.ID, webhookToken string, messageID snowflake.ID, params discord.EditWebhookMessageParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) EnableSMSMFA

```go
func (a *API) EnableSMSMFA(ctx context.Context, params discord.EnableSMSMFAParams, opts ...Option) error
```

#### func (*API) EnableSocialSdk

```go
func (a *API) EnableSocialSdk(ctx context.Context, applicationID snowflake.ID, params discord.EnableSocialSdkParams, opts ...Option) error
```

#### func (*API) EnableTotpMFA

```go
func (a *API) EnableTotpMFA(ctx context.Context, params discord.EnableTotpMFAParams, opts ...Option) (*discord.EnableTotpMFAResponse, error)
```

#### func (*API) EndPoll

```go
func (a *API) EndPoll(ctx context.Context, channelID snowflake.ID, messageID snowflake.ID, opts ...Option) (*discord.Message, error)
```

#### func (*API) ExchangeHandoffToken

```go
func (a *API) ExchangeHandoffToken(ctx context.Context, params discord.ExchangeHandoffTokenParams, opts ...Option) (*discord.ExchangeHandoffTokenResponse, error)
```

#### func (*API) ExchangeProvisionalAccountChildToken

```go
func (a *API) ExchangeProvisionalAccountChildToken(ctx context.Context, params discord.ExchangeProvisionalAccountChildTokenParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ExchangeRemoteAuthTicket

```go
func (a *API) ExchangeRemoteAuthTicket(ctx context.Context, params discord.ExchangeRemoteAuthTicketParams, opts ...Option) (*discord.ExchangeRemoteAuthTicketResponse, error)
```

#### func (*API) ExecuteAutomodAlertAction

```go
func (a *API) ExecuteAutomodAlertAction(ctx context.Context, guildID snowflake.ID, params discord.ExecuteAutomodAlertActionParams, opts ...Option) error
```

#### func (*API) ExecuteServiceWebhook

```go
func (a *API) ExecuteServiceWebhook(ctx context.Context, webhookID snowflake.ID, webhookToken string, webhookService string, query discord.ExecuteServiceWebhookQuery, opts ...Option) (*discord.Message, error)
```

#### func (*API) ExecuteWebhook

```go
func (a *API) ExecuteWebhook(ctx context.Context, webhookID snowflake.ID, webhookToken string, query discord.ExecuteWebhookQuery, params discord.ExecuteWebhookParams, opts ...Option) (*discord.Message, error)
```

#### func (*API) FinishOAuth2DeviceFlow

```go
func (a *API) FinishOAuth2DeviceFlow(ctx context.Context, params discord.FinishOAuth2DeviceFlowParams, opts ...Option) error
```

#### func (*API) FinishPasswordlessLogin

```go
func (a *API) FinishPasswordlessLogin(ctx context.Context, params discord.FinishPasswordlessLoginParams, opts ...Option) (*discord.FinishPasswordlessLoginResponse, error)
```

#### func (*API) FinishRemoteAuth

```go
func (a *API) FinishRemoteAuth(ctx context.Context, params discord.FinishRemoteAuthParams, opts ...Option) error
```

#### func (*API) FixGrammar

```go
func (a *API) FixGrammar(ctx context.Context, params discord.FixGrammarParams, opts ...Option) (*discord.FixGrammarResponse, error)
```

#### func (*API) FollowChannel

```go
func (a *API) FollowChannel(ctx context.Context, channelID snowflake.ID, params discord.FollowChannelParams, opts ...Option) (*discord.FollowedChannel, error)
```

#### func (*API) ForgotPassword

```go
func (a *API) ForgotPassword(ctx context.Context, params discord.ForgotPasswordParams, opts ...Option) (*discord.ForgotPasswordResponse, error)
```

#### func (*API) GetActivityMetadata

```go
func (a *API) GetActivityMetadata(ctx context.Context, userID snowflake.ID, sessionID snowflake.ID, applicationID snowflake.ID, query discord.GetActivityMetadataQuery, opts ...Option) (*discord.GetActivityMetadataResponse, error)
```

#### func (*API) GetAdminCommunityEligibility

```go
func (a *API) GetAdminCommunityEligibility(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.GetAdminCommunityEligibilityResponse, error)
```

#### func (*API) GetApexExperimentAssignments

```go
func (a *API) GetApexExperimentAssignments(ctx context.Context, opts ...Option) (*discord.ApexExperiments, error)
```

#### func (*API) GetApplication

```go
func (a *API) GetApplication(ctx context.Context, applicationID snowflake.ID, opts ...Option) (*discord.Application, error)
```

#### func (*API) GetApplicationAssets

```go
func (a *API) GetApplicationAssets(ctx context.Context, applicationID snowflake.ID, opts ...Option) ([]discord.ApplicationAsset, error)
```

#### func (*API) GetApplicationCommandIndex

```go
func (a *API) GetApplicationCommandIndex(ctx context.Context, applicationID snowflake.ID, opts ...Option) (*discord.ApplicationCommandIndex, error)
```

#### func (*API) GetApplicationDirectoryApplication

```go
func (a *API) GetApplicationDirectoryApplication(ctx context.Context, applicationID snowflake.ID, query discord.GetApplicationDirectoryApplicationQuery, opts ...Option) (*discord.Application, error)
```

#### func (*API) GetApplicationDirectoryApplicationEmbed

```go
func (a *API) GetApplicationDirectoryApplicationEmbed(ctx context.Context, applicationID snowflake.ID, query discord.GetApplicationDirectoryApplicationEmbedQuery, opts ...Option) (*discord.Application, error)
```

#### func (*API) GetApplicationDiscoverabilityState

```go
func (a *API) GetApplicationDiscoverabilityState(ctx context.Context, applicationID snowflake.ID, opts ...Option) (*discord.GetApplicationDiscoverabilityStateResponse, error)
```

#### func (*API) GetApplicationEmoji

```go
func (a *API) GetApplicationEmoji(ctx context.Context, applicationID snowflake.ID, emojiID snowflake.ID, opts ...Option) (*discord.Emoji, error)
```

#### func (*API) GetApplicationGameClaim

```go
func (a *API) GetApplicationGameClaim(ctx context.Context, applicationID snowflake.ID, gameClaimID snowflake.ID, opts ...Option) (*discord.GameClaim, error)
```

#### func (*API) GetApplicationGiftCodeBatch

```go
func (a *API) GetApplicationGiftCodeBatch(ctx context.Context, applicationID snowflake.ID, giftCodeBatchID snowflake.ID, opts ...Option) (json.RawMessage, error)
```

#### func (*API) GetApplicationManagedLink

```go
func (a *API) GetApplicationManagedLink(ctx context.Context, applicationID snowflake.ID, activityLinkID snowflake.ID, opts ...Option) (*discord.ActivityLink, error)
```

#### func (*API) GetApplicationOAuth2Authorizations

```go
func (a *API) GetApplicationOAuth2Authorizations(ctx context.Context, applicationID snowflake.ID, opts ...Option) ([]discord.OAuth2Authorization, error)
```

#### func (*API) GetApplicationPrimaryStoreListing

```go
func (a *API) GetApplicationPrimaryStoreListing(ctx context.Context, applicationID snowflake.ID, query discord.GetApplicationPrimaryStoreListingQuery, opts ...Option) (*discord.StoreListing, error)
```

#### func (*API) GetApplicationProxyConfig

```go
func (a *API) GetApplicationProxyConfig(ctx context.Context, applicationID snowflake.ID, opts ...Option) (json.RawMessage, error)
```

#### func (*API) GetApplicationQuickLink

```go
func (a *API) GetApplicationQuickLink(ctx context.Context, applicationID snowflake.ID, activityLinkID snowflake.ID, opts ...Option) (*discord.ActivityLink, error)
```

#### func (*API) GetApplicationSKUStorefront

```go
func (a *API) GetApplicationSKUStorefront(ctx context.Context, applicationID snowflake.ID, skuid snowflake.ID, opts ...Option) (*discord.GetApplicationSKUStorefrontResponse, error)
```

#### func (*API) GetApplicationStoreLayout

```go
func (a *API) GetApplicationStoreLayout(ctx context.Context, applicationID snowflake.ID, opts ...Option) (*discord.GetApplicationStoreLayoutResponse, error)
```

#### func (*API) GetApplicationStorefront

```go
func (a *API) GetApplicationStorefront(ctx context.Context, applicationID snowflake.ID, opts ...Option) (*discord.Storefront, error)
```

#### func (*API) GetApplicationStorefrontForPremiumButton

```go
func (a *API) GetApplicationStorefrontForPremiumButton(ctx context.Context, storeListingID snowflake.ID, opts ...Option) (*discord.GetApplicationStorefrontForPremiumButtonResponse, error)
```

#### func (*API) GetApplicationUndeletableReason

```go
func (a *API) GetApplicationUndeletableReason(ctx context.Context, applicationID snowflake.ID, opts ...Option) (*discord.GetApplicationUndeletableReasonResponse, error)
```

#### func (*API) GetApplicationUpdates

```go
func (a *API) GetApplicationUpdates(ctx context.Context, releaseChannel string, query discord.GetApplicationUpdatesQuery, opts ...Option) (*discord.GetApplicationUpdatesResponse, error)
```

#### func (*API) GetApplicationVerificationEligibility

```go
func (a *API) GetApplicationVerificationEligibility(ctx context.Context, applicationID snowflake.ID, opts ...Option) (json.RawMessage, error)
```

#### func (*API) GetAuthSessions

```go
func (a *API) GetAuthSessions(ctx context.Context, opts ...Option) (*discord.GetAuthSessionsResponse, error)
```

#### func (*API) GetBackupCodes

```go
func (a *API) GetBackupCodes(ctx context.Context, params discord.GetBackupCodesParams, opts ...Option) (*discord.GetBackupCodesResponse, error)
```

#### func (*API) GetBillingCountryCode

```go
func (a *API) GetBillingCountryCode(ctx context.Context, opts ...Option) (*discord.GetBillingCountryCodeResponse, error)
```

#### func (*API) GetBillingLocationInfo

```go
func (a *API) GetBillingLocationInfo(ctx context.Context, opts ...Option) (*discord.GetBillingLocationInfoResponse, error)
```

#### func (*API) GetCallEligibility

```go
func (a *API) GetCallEligibility(ctx context.Context, channelID snowflake.ID, opts ...Option) (*discord.GetCallEligibilityResponse, error)
```

#### func (*API) GetChannel

```go
func (a *API) GetChannel(ctx context.Context, channelID snowflake.ID, opts ...Option) (*discord.Channel, error)
```

#### func (*API) GetChannelApplicationCommandIndex

```go
func (a *API) GetChannelApplicationCommandIndex(ctx context.Context, channelID snowflake.ID, opts ...Option) (*discord.ApplicationCommandIndex, error)
```

#### func (*API) GetChannelLinkedAccounts

```go
func (a *API) GetChannelLinkedAccounts(ctx context.Context, channelID snowflake.ID, query discord.GetChannelLinkedAccountsQuery, opts ...Option) (*discord.GetChannelLinkedAccountsResponse, error)
```

#### func (*API) GetChannelMediaPreview

```go
func (a *API) GetChannelMediaPreview(ctx context.Context, channelID snowflake.ID, opts ...Option) (*discord.GetChannelMediaPreviewResponse, error)
```

#### func (*API) GetChannelPostData

```go
func (a *API) GetChannelPostData(ctx context.Context, channelID snowflake.ID, params discord.GetChannelPostDataParams, opts ...Option) (*discord.GetChannelPostDataResponse, error)
```

#### func (*API) GetCheckoutRecovery

```go
func (a *API) GetCheckoutRecovery(ctx context.Context, opts ...Option) (*discord.GetCheckoutRecoveryResponse, error)
```

#### func (*API) GetCheckpoint

```go
func (a *API) GetCheckpoint(ctx context.Context, opts ...Option) (*discord.GetCheckpointResponse, error)
```

#### func (*API) GetChurnUserOffer

```go
func (a *API) GetChurnUserOffer(ctx context.Context, opts ...Option) (*discord.GetChurnUserOfferResponse, error)
```

#### func (*API) GetCollectiblesMarketing

```go
func (a *API) GetCollectiblesMarketing(ctx context.Context, query discord.GetCollectiblesMarketingQuery, opts ...Option) (*discord.GetCollectiblesMarketingResponse, error)
```

#### func (*API) GetCollectiblesProduct

```go
func (a *API) GetCollectiblesProduct(ctx context.Context, skuid snowflake.ID, query discord.GetCollectiblesProductQuery, opts ...Option) (*discord.CollectibleProduct, error)
```

#### func (*API) GetCollectiblesShop

```go
func (a *API) GetCollectiblesShop(ctx context.Context, query discord.GetCollectiblesShopQuery, opts ...Option) (*discord.GetCollectiblesShopResponse, error)
```

#### func (*API) GetCollectiblesShopTabLayout

```go
func (a *API) GetCollectiblesShopTabLayout(ctx context.Context, collectiblesShopTab string, opts ...Option) (*discord.GetCollectiblesShopTabLayoutResponse, error)
```

#### func (*API) GetCompany

```go
func (a *API) GetCompany(ctx context.Context, companyID snowflake.ID, opts ...Option) (*discord.TeamCompany, error)
```

#### func (*API) GetConfettiConsumable

```go
func (a *API) GetConfettiConsumable(ctx context.Context, opts ...Option) (*discord.GetConfettiConsumableResponse, error)
```

#### func (*API) GetConsumableSKUPricing

```go
func (a *API) GetConsumableSKUPricing(ctx context.Context, skuid snowflake.ID, opts ...Option) (*discord.GetConsumableSKUPricingResponse, error)
```

#### func (*API) GetCreatorMonetizationEligibility

```go
func (a *API) GetCreatorMonetizationEligibility(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.GetCreatorMonetizationEligibilityResponse, error)
```

#### func (*API) GetCreatorMonetizationMarketingOnboarding

```go
func (a *API) GetCreatorMonetizationMarketingOnboarding(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.GetCreatorMonetizationMarketingOnboardingResponse, error)
```

#### func (*API) GetCreatorMonetizationRestrictions

```go
func (a *API) GetCreatorMonetizationRestrictions(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.GetCreatorMonetizationRestrictionsResponse, error)
```

#### func (*API) GetCurrentApplication

```go
func (a *API) GetCurrentApplication(ctx context.Context, opts ...Option) (*discord.Application, error)
```

#### func (*API) GetCurrentAuthorizationInformation

```go
func (a *API) GetCurrentAuthorizationInformation(ctx context.Context, opts ...Option) (*discord.GetCurrentAuthorizationInformationResponse, error)
```

#### func (*API) GetCurrentGuildMember

```go
func (a *API) GetCurrentGuildMember(ctx context.Context, guildID snowflake.ID, opts ...Option) (json.RawMessage, error)
```

#### func (*API) GetCurrentUser

```go
func (a *API) GetCurrentUser(ctx context.Context, query discord.GetCurrentUserQuery, opts ...Option) (*discord.User, error)
```

#### func (*API) GetCurrentUserGuildJoinRequest

```go
func (a *API) GetCurrentUserGuildJoinRequest(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.GuildJoinRequest, error)
```

#### func (*API) GetCurrentUserVoiceState

```go
func (a *API) GetCurrentUserVoiceState(ctx context.Context, guildID snowflake.ID, opts ...Option) (json.RawMessage, error)
```

#### func (*API) GetDMChannel

```go
func (a *API) GetDMChannel(ctx context.Context, userID snowflake.ID, opts ...Option) (json.RawMessage, error)
```

#### func (*API) GetDeveloperApplicationWidgetConfigs

```go
func (a *API) GetDeveloperApplicationWidgetConfigs(ctx context.Context, opts ...Option) (*discord.GetDeveloperApplicationWidgetConfigsResponse, error)
```

#### func (*API) GetDeviceSyncToken

```go
func (a *API) GetDeviceSyncToken(ctx context.Context, opts ...Option) (*discord.GetDeviceSyncTokenResponse, error)
```

#### func (*API) GetDirectoryBroadcastInfo

```go
func (a *API) GetDirectoryBroadcastInfo(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.GetDirectoryBroadcastInfoResponse, error)
```

#### func (*API) GetDirectoryCounts

```go
func (a *API) GetDirectoryCounts(ctx context.Context, channelID snowflake.ID, opts ...Option) (json.RawMessage, error)
```

#### func (*API) GetDirectoryEntry

```go
func (a *API) GetDirectoryEntry(ctx context.Context, channelID snowflake.ID, entityID snowflake.ID, opts ...Option) (*discord.DirectoryEntry, error)
```

#### func (*API) GetDiscoverySlug

```go
func (a *API) GetDiscoverySlug(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.GetDiscoverySlugResponse, error)
```

#### func (*API) GetEarnedQuestPlacement

```go
func (a *API) GetEarnedQuestPlacement(ctx context.Context, query discord.GetEarnedQuestPlacementQuery, opts ...Option) (*discord.GetEarnedQuestPlacementResponse, error)
```

#### func (*API) GetEmailSettings

```go
func (a *API) GetEmailSettings(ctx context.Context, opts ...Option) (json.RawMessage, error)
```

#### func (*API) GetEmbeddedActivityConfig

```go
func (a *API) GetEmbeddedActivityConfig(ctx context.Context, applicationID snowflake.ID, opts ...Option) (*discord.EmbeddedActivityConfig, error)
```

#### func (*API) GetEmojiGuild

```go
func (a *API) GetEmojiGuild(ctx context.Context, emojiID snowflake.ID, opts ...Option) (*discord.DiscoverableGuild, error)
```

#### func (*API) GetEmojiSource

```go
func (a *API) GetEmojiSource(ctx context.Context, emojiID snowflake.ID, opts ...Option) (*discord.GetEmojiSourceResponse, error)
```

#### func (*API) GetEula

```go
func (a *API) GetEula(ctx context.Context, eulaID snowflake.ID, opts ...Option) (*discord.Eula, error)
```

#### func (*API) GetExperimentAssignments

```go
func (a *API) GetExperimentAssignments(ctx context.Context, opts ...Option) (*discord.GetExperimentAssignmentsResponse, error)
```

#### func (*API) GetFamilyCenterOverview

```go
func (a *API) GetFamilyCenterOverview(ctx context.Context, opts ...Option) (*discord.FamilyCenter, error)
```

#### func (*API) GetFollowupMessage

```go
func (a *API) GetFollowupMessage(ctx context.Context, applicationID snowflake.ID, interactionToken string, messageID snowflake.ID, opts ...Option) (json.RawMessage, error)
```

#### func (*API) GetGame

```go
func (a *API) GetGame(ctx context.Context, gameID snowflake.ID, query discord.GetGameQuery, opts ...Option) (*discord.Game, error)
```

#### func (*API) GetGateway

```go
func (a *API) GetGateway(ctx context.Context, opts ...Option) (*discord.GetGatewayResponse, error)
```

#### func (*API) GetGatewayBot

```go
func (a *API) GetGatewayBot(ctx context.Context, opts ...Option) (*discord.GetGatewayBotResponse, error)
```

#### func (*API) GetGiftCode

```go
func (a *API) GetGiftCode(ctx context.Context, giftCodeCode string, query discord.GetGiftCodeQuery, opts ...Option) (*discord.GiftCode, error)
```

#### func (*API) GetGuild

```go
func (a *API) GetGuild(ctx context.Context, guildID snowflake.ID, query discord.GetGuildQuery, opts ...Option) (*discord.Guild, error)
```

#### func (*API) GetGuildApplicationCommandIndex

```go
func (a *API) GetGuildApplicationCommandIndex(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.ApplicationCommandIndex, error)
```

#### func (*API) GetGuildApplicationSKUStorefront

```go
func (a *API) GetGuildApplicationSKUStorefront(ctx context.Context, guildID snowflake.ID, skuid snowflake.ID, opts ...Option) (*discord.GetGuildApplicationSKUStorefrontResponse, error)
```

#### func (*API) GetGuildApplicationStorefront

```go
func (a *API) GetGuildApplicationStorefront(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.Storefront, error)
```

#### func (*API) GetGuildApplicationStorefrontAnnouncement

```go
func (a *API) GetGuildApplicationStorefrontAnnouncement(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.GetGuildApplicationStorefrontAnnouncementResponse, error)
```

#### func (*API) GetGuildAutomodRule

```go
func (a *API) GetGuildAutomodRule(ctx context.Context, guildID snowflake.ID, automodRuleID snowflake.ID, opts ...Option) (*discord.AutomodRule, error)
```

#### func (*API) GetGuildBan

```go
func (a *API) GetGuildBan(ctx context.Context, guildID snowflake.ID, userID snowflake.ID, opts ...Option) (*discord.Ban, error)
```

#### func (*API) GetGuildBasic

```go
func (a *API) GetGuildBasic(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.Guild, error)
```

#### func (*API) GetGuildDiscoveryMetadata

```go
func (a *API) GetGuildDiscoveryMetadata(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.DiscoveryMetadata, error)
```

#### func (*API) GetGuildDiscoveryRequirements

```go
func (a *API) GetGuildDiscoveryRequirements(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.DiscoveryRequirements, error)
```

#### func (*API) GetGuildEmoji

```go
func (a *API) GetGuildEmoji(ctx context.Context, guildID snowflake.ID, emojiID snowflake.ID, opts ...Option) (*discord.Emoji, error)
```

#### func (*API) GetGuildIntegrationApplicationIDs

```go
func (a *API) GetGuildIntegrationApplicationIDs(ctx context.Context, opts ...Option) (json.RawMessage, error)
```

#### func (*API) GetGuildJoinRequest

```go
func (a *API) GetGuildJoinRequest(ctx context.Context, guildJoinRequestID snowflake.ID, opts ...Option) (*discord.GuildJoinRequest, error)
```

#### func (*API) GetGuildJoinRequestCooldown

```go
func (a *API) GetGuildJoinRequestCooldown(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.GetGuildJoinRequestCooldownResponse, error)
```

#### func (*API) GetGuildMember

```go
func (a *API) GetGuildMember(ctx context.Context, guildID snowflake.ID, userID snowflake.ID, opts ...Option) (*discord.GuildMember, error)
```

#### func (*API) GetGuildMemberVerification

```go
func (a *API) GetGuildMemberVerification(ctx context.Context, guildID snowflake.ID, query discord.GetGuildMemberVerificationQuery, opts ...Option) (*discord.MemberVerification, error)
```

#### func (*API) GetGuildNewMemberActions

```go
func (a *API) GetGuildNewMemberActions(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.NewMemberActionsProgress, error)
```

#### func (*API) GetGuildNewMemberWelcome

```go
func (a *API) GetGuildNewMemberWelcome(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.NewMemberWelcome, error)
```

#### func (*API) GetGuildOnboarding

```go
func (a *API) GetGuildOnboarding(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.Onboarding, error)
```

#### func (*API) GetGuildOwnershipTransferCode

```go
func (a *API) GetGuildOwnershipTransferCode(ctx context.Context, guildID snowflake.ID, opts ...Option) error
```

#### func (*API) GetGuildPreview

```go
func (a *API) GetGuildPreview(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.Guild, error)
```

#### func (*API) GetGuildProductListing

```go
func (a *API) GetGuildProductListing(ctx context.Context, guildID snowflake.ID, productListingID snowflake.ID, opts ...Option) (*discord.GuildProductListing, error)
```

#### func (*API) GetGuildProfile

```go
func (a *API) GetGuildProfile(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.GuildProfile, error)
```

#### func (*API) GetGuildPrune

```go
func (a *API) GetGuildPrune(ctx context.Context, guildID snowflake.ID, query discord.GetGuildPruneQuery, opts ...Option) (*discord.GetGuildPruneResponse, error)
```

#### func (*API) GetGuildRole

```go
func (a *API) GetGuildRole(ctx context.Context, guildID snowflake.ID, roleID snowflake.ID, opts ...Option) (*discord.Role, error)
```

#### func (*API) GetGuildRoleConnectionConfiguration

```go
func (a *API) GetGuildRoleConnectionConfiguration(ctx context.Context, guildID snowflake.ID, roleID snowflake.ID, opts ...Option) (json.RawMessage, error)
```

#### func (*API) GetGuildRoleConnectionEligibility

```go
func (a *API) GetGuildRoleConnectionEligibility(ctx context.Context, guildID snowflake.ID, roleID snowflake.ID, opts ...Option) (json.RawMessage, error)
```

#### func (*API) GetGuildRoleMemberCounts

```go
func (a *API) GetGuildRoleMemberCounts(ctx context.Context, guildID snowflake.ID, opts ...Option) (json.RawMessage, error)
```

#### func (*API) GetGuildRoleSubscriptionGroupListing

```go
func (a *API) GetGuildRoleSubscriptionGroupListing(ctx context.Context, guildID snowflake.ID, subscriptionGroupListingID snowflake.ID, opts ...Option) (*discord.RoleSubscriptionGroupListing, error)
```

#### func (*API) GetGuildRoleSubscriptionsSettings

```go
func (a *API) GetGuildRoleSubscriptionsSettings(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.RoleSubscriptionSettings, error)
```

#### func (*API) GetGuildScheduledEvent

```go
func (a *API) GetGuildScheduledEvent(ctx context.Context, guildID snowflake.ID, guildScheduledEventID snowflake.ID, query discord.GetGuildScheduledEventQuery, opts ...Option) (*discord.GuildScheduledEvent, error)
```

#### func (*API) GetGuildScheduledEventUserCount

```go
func (a *API) GetGuildScheduledEventUserCount(ctx context.Context, guildID snowflake.ID, guildScheduledEventID snowflake.ID, query discord.GetGuildScheduledEventUserCountQuery, opts ...Option) (*discord.GetGuildScheduledEventUserCountResponse, error)
```

#### func (*API) GetGuildSoundboardSound

```go
func (a *API) GetGuildSoundboardSound(ctx context.Context, guildID snowflake.ID, soundID snowflake.ID, opts ...Option) (*discord.SoundboardSound, error)
```

#### func (*API) GetGuildSticker

```go
func (a *API) GetGuildSticker(ctx context.Context, guildID snowflake.ID, stickerID snowflake.ID, opts ...Option) (*discord.Sticker, error)
```

#### func (*API) GetGuildTemplate

```go
func (a *API) GetGuildTemplate(ctx context.Context, templateCode string, opts ...Option) (*discord.GuildTemplate, error)
```

#### func (*API) GetGuildVanityInvite

```go
func (a *API) GetGuildVanityInvite(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.GetGuildVanityInviteResponse, error)
```

#### func (*API) GetGuildWelcomeScreen

```go
func (a *API) GetGuildWelcomeScreen(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.WelcomeScreen, error)
```

#### func (*API) GetGuildWidget

```go
func (a *API) GetGuildWidget(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.GuildWidget, error)
```

#### func (*API) GetGuildWidgetImage

```go
func (a *API) GetGuildWidgetImage(ctx context.Context, guildID snowflake.ID, query discord.GetGuildWidgetImageQuery, opts ...Option) (json.RawMessage, error)
```

#### func (*API) GetGuildWidgetSettings

```go
func (a *API) GetGuildWidgetSettings(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.GuildWidgetSettings, error)
```

#### func (*API) GetHdStreamingConsumable

```go
func (a *API) GetHdStreamingConsumable(ctx context.Context, opts ...Option) (*discord.GetHdStreamingConsumableResponse, error)
```

#### func (*API) GetInvite

```go
func (a *API) GetInvite(ctx context.Context, inviteCode string, query discord.GetInviteQuery, opts ...Option) (*discord.Invite, error)
```

#### func (*API) GetInviteTargetUsersJobStatus

```go
func (a *API) GetInviteTargetUsersJobStatus(ctx context.Context, inviteCode string, opts ...Option) (*discord.InviteTargetUsersJob, error)
```

#### func (*API) GetLatestApplicationInstaller

```go
func (a *API) GetLatestApplicationInstaller(ctx context.Context, releaseChannel string, query discord.GetLatestApplicationInstallerQuery, opts ...Option) (json.RawMessage, error)
```

#### func (*API) GetLatestDistributedApplicationInstaller

```go
func (a *API) GetLatestDistributedApplicationInstaller(ctx context.Context, query discord.GetLatestDistributedApplicationInstallerQuery, opts ...Option) (json.RawMessage, error)
```

#### func (*API) GetLatestDistributedApplicationManifest

```go
func (a *API) GetLatestDistributedApplicationManifest(ctx context.Context, query discord.GetLatestDistributedApplicationManifestQuery, opts ...Option) (*discord.GetLatestDistributedApplicationManifestResponse, error)
```

#### func (*API) GetLayoutDefinitions

```go
func (a *API) GetLayoutDefinitions(ctx context.Context, opts ...Option) (*discord.GetLayoutDefinitionsResponse, error)
```

#### func (*API) GetLinkCode

```go
func (a *API) GetLinkCode(ctx context.Context, opts ...Option) (*discord.GetLinkCodeResponse, error)
```

#### func (*API) GetLinkedUsers

```go
func (a *API) GetLinkedUsers(ctx context.Context, opts ...Option) (*discord.LinkedUser, error)
```

#### func (*API) GetLobby

```go
func (a *API) GetLobby(ctx context.Context, lobbyID snowflake.ID, opts ...Option) (*discord.Lobby, error)
```

#### func (*API) GetLocalizedPricingPromo

```go
func (a *API) GetLocalizedPricingPromo(ctx context.Context, opts ...Option) (*discord.GetLocalizedPricingPromoResponse, error)
```

#### func (*API) GetLocationMetadata

```go
func (a *API) GetLocationMetadata(ctx context.Context, opts ...Option) (*discord.GetLocationMetadataResponse, error)
```

#### func (*API) GetMessage

```go
func (a *API) GetMessage(ctx context.Context, channelID snowflake.ID, messageID snowflake.ID, opts ...Option) (json.RawMessage, error)
```

#### func (*API) GetMessageInteractionData

```go
func (a *API) GetMessageInteractionData(ctx context.Context, channelID snowflake.ID, messageID snowflake.ID, opts ...Option) (*discord.GetMessageInteractionDataResponse, error)
```

#### func (*API) GetMetadataForApexExperiments

```go
func (a *API) GetMetadataForApexExperiments(ctx context.Context, opts ...Option) (*discord.GetMetadataForApexExperimentsResponse, error)
```

#### func (*API) GetNativeModule

```go
func (a *API) GetNativeModule(ctx context.Context, releaseChannel string, moduleName string, moduleVersion string, query discord.GetNativeModuleQuery, opts ...Option) (json.RawMessage, error)
```

#### func (*API) GetNativeModuleVersions

```go
func (a *API) GetNativeModuleVersions(ctx context.Context, releaseChannel string, query discord.GetNativeModuleVersionsQuery, opts ...Option) (json.RawMessage, error)
```

#### func (*API) GetOAuth2Authorization

```go
func (a *API) GetOAuth2Authorization(ctx context.Context, tokenID snowflake.ID, opts ...Option) (*discord.OAuth2Authorization, error)
```

#### func (*API) GetOAuth2Authorizations

```go
func (a *API) GetOAuth2Authorizations(ctx context.Context, query discord.GetOAuth2AuthorizationsQuery, opts ...Option) ([]discord.OAuth2Authorization, error)
```

#### func (*API) GetOAuth2DeviceCode

```go
func (a *API) GetOAuth2DeviceCode(ctx context.Context, params discord.GetOAuth2DeviceCodeParams, opts ...Option) (*discord.GetOAuth2DeviceCodeResponse, error)
```

#### func (*API) GetOAuth2DeviceFlow

```go
func (a *API) GetOAuth2DeviceFlow(ctx context.Context, params discord.GetOAuth2DeviceFlowParams, opts ...Option) (*discord.GetOAuth2DeviceFlowResponse, error)
```

#### func (*API) GetOAuth2Token

```go
func (a *API) GetOAuth2Token(ctx context.Context, params discord.GetOAuth2TokenParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) GetOpenidConnectKeys

```go
func (a *API) GetOpenidConnectKeys(ctx context.Context, opts ...Option) (json.RawMessage, error)
```

#### func (*API) GetOpenidUserInformation

```go
func (a *API) GetOpenidUserInformation(ctx context.Context, opts ...Option) (*discord.GetOpenidUserInformationResponse, error)
```

#### func (*API) GetOriginalInteractionResponse

```go
func (a *API) GetOriginalInteractionResponse(ctx context.Context, applicationID snowflake.ID, interactionToken string, opts ...Option) (json.RawMessage, error)
```

#### func (*API) GetPartialApplication

```go
func (a *API) GetPartialApplication(ctx context.Context, applicationID snowflake.ID, query discord.GetPartialApplicationQuery, opts ...Option) (*discord.Application, error)
```

#### func (*API) GetPasswordStrength

```go
func (a *API) GetPasswordStrength(ctx context.Context, params discord.GetPasswordStrengthParams, opts ...Option) (*discord.GetPasswordStrengthResponse, error)
```

#### func (*API) GetPayment

```go
func (a *API) GetPayment(ctx context.Context, paymentID snowflake.ID, opts ...Option) (*discord.Payment, error)
```

#### func (*API) GetPaymentInvoiceBreakdown

```go
func (a *API) GetPaymentInvoiceBreakdown(ctx context.Context, query discord.GetPaymentInvoiceBreakdownQuery, opts ...Option) (*discord.GetPaymentInvoiceBreakdownResponse, error)
```

#### func (*API) GetPaymentSource

```go
func (a *API) GetPaymentSource(ctx context.Context, paymentSourceID snowflake.ID, opts ...Option) (*discord.PaymentSource, error)
```

#### func (*API) GetPaymentSourceCreationContext

```go
func (a *API) GetPaymentSourceCreationContext(ctx context.Context, opts ...Option) (*discord.GetPaymentSourceCreationContextResponse, error)
```

#### func (*API) GetPremiumGuildSubscriptionCooldown

```go
func (a *API) GetPremiumGuildSubscriptionCooldown(ctx context.Context, opts ...Option) (*discord.GetPremiumGuildSubscriptionCooldownResponse, error)
```

#### func (*API) GetPremiumReferral

```go
func (a *API) GetPremiumReferral(ctx context.Context, referralID snowflake.ID, opts ...Option) (*discord.PremiumReferral, error)
```

#### func (*API) GetPremiumReferralEligibility

```go
func (a *API) GetPremiumReferralEligibility(ctx context.Context, opts ...Option) (*discord.PremiumReferralEligibility, error)
```

#### func (*API) GetPremiumReferralIncentiveEligibility

```go
func (a *API) GetPremiumReferralIncentiveEligibility(ctx context.Context, opts ...Option) (*discord.GetPremiumReferralIncentiveEligibilityResponse, error)
```

#### func (*API) GetProfileWidgetsSuggestedGames

```go
func (a *API) GetProfileWidgetsSuggestedGames(ctx context.Context, opts ...Option) (*discord.GetProfileWidgetsSuggestedGamesResponse, error)
```

#### func (*API) GetProvisionalAccountToken

```go
func (a *API) GetProvisionalAccountToken(ctx context.Context, params discord.GetProvisionalAccountTokenParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) GetProvisionalAccountTokenWithBot

```go
func (a *API) GetProvisionalAccountTokenWithBot(ctx context.Context, params discord.GetProvisionalAccountTokenWithBotParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) GetQuestConfig

```go
func (a *API) GetQuestConfig(ctx context.Context, questID snowflake.ID, opts ...Option) (*discord.QuestConfig, error)
```

#### func (*API) GetQuestCreativePreview

```go
func (a *API) GetQuestCreativePreview(ctx context.Context, query discord.GetQuestCreativePreviewQuery, opts ...Option) (*discord.GetQuestCreativePreviewResponse, error)
```

#### func (*API) GetQuestDecisions

```go
func (a *API) GetQuestDecisions(ctx context.Context, query discord.GetQuestDecisionsQuery, opts ...Option) (*discord.GetQuestDecisionsResponse, error)
```

#### func (*API) GetQuestPlacement

```go
func (a *API) GetQuestPlacement(ctx context.Context, query discord.GetQuestPlacementQuery, opts ...Option) (*discord.GetQuestPlacementResponse, error)
```

#### func (*API) GetQuestPreview

```go
func (a *API) GetQuestPreview(ctx context.Context, questID snowflake.ID, opts ...Option) (*discord.Quest, error)
```

#### func (*API) GetQuestRewardCode

```go
func (a *API) GetQuestRewardCode(ctx context.Context, questID snowflake.ID, opts ...Option) (*discord.QuestRewardCode, error)
```

#### func (*API) GetRPCApplication

```go
func (a *API) GetRPCApplication(ctx context.Context, applicationID snowflake.ID, opts ...Option) (*discord.Application, error)
```

#### func (*API) GetReportMenu

```go
func (a *API) GetReportMenu(ctx context.Context, typeArg string, query discord.GetReportMenuQuery, opts ...Option) (*discord.ReportMenu, error)
```

#### func (*API) GetReportOptions

```go
func (a *API) GetReportOptions(ctx context.Context, opts ...Option) ([]discord.ReportOption, error)
```

#### func (*API) GetReportReasons

```go
func (a *API) GetReportReasons(ctx context.Context, query discord.GetReportReasonsQuery, opts ...Option) ([]discord.ReportReason, error)
```

#### func (*API) GetSKU

```go
func (a *API) GetSKU(ctx context.Context, skuid snowflake.ID, query discord.GetSKUQuery, opts ...Option) (*discord.SKU, error)
```

#### func (*API) GetSKUPublishedStoreListing

```go
func (a *API) GetSKUPublishedStoreListing(ctx context.Context, skuid snowflake.ID, query discord.GetSKUPublishedStoreListingQuery, opts ...Option) (*discord.StoreListing, error)
```

#### func (*API) GetSKUPurchasePreview

```go
func (a *API) GetSKUPurchasePreview(ctx context.Context, skuid snowflake.ID, query discord.GetSKUPurchasePreviewQuery, opts ...Option) (*discord.Invoice, error)
```

#### func (*API) GetServiceWebhook

```go
func (a *API) GetServiceWebhook(ctx context.Context, webhookID snowflake.ID, webhookToken string, webhookService string, opts ...Option) (json.RawMessage, error)
```

#### func (*API) GetSocialLayerStorefrontConfig

```go
func (a *API) GetSocialLayerStorefrontConfig(ctx context.Context, opts ...Option) (*discord.GetSocialLayerStorefrontConfigResponse, error)
```

#### func (*API) GetSocialLayerStorefrontEligibilities

```go
func (a *API) GetSocialLayerStorefrontEligibilities(ctx context.Context, opts ...Option) (json.RawMessage, error)
```

#### func (*API) GetSocialSdkRelease

```go
func (a *API) GetSocialSdkRelease(ctx context.Context, sdkReleaseVersion string, opts ...Option) (*discord.SocialSdkRelease, error)
```

#### func (*API) GetSoundboardSoundGuild

```go
func (a *API) GetSoundboardSoundGuild(ctx context.Context, soundID snowflake.ID, guildID snowflake.ID, opts ...Option) (*discord.DiscoverableGuild, error)
```

#### func (*API) GetStageInstance

```go
func (a *API) GetStageInstance(ctx context.Context, channelID snowflake.ID, opts ...Option) (*discord.StageInstance, error)
```

#### func (*API) GetSticker

```go
func (a *API) GetSticker(ctx context.Context, stickerID snowflake.ID, opts ...Option) (*discord.Sticker, error)
```

#### func (*API) GetStickerGuild

```go
func (a *API) GetStickerGuild(ctx context.Context, stickerID snowflake.ID, opts ...Option) (*discord.DiscoverableGuild, error)
```

#### func (*API) GetStickerPack

```go
func (a *API) GetStickerPack(ctx context.Context, stickerPackID snowflake.ID, opts ...Option) (*discord.StickerPack, error)
```

#### func (*API) GetStorePriceTier

```go
func (a *API) GetStorePriceTier(ctx context.Context, priceTier string, opts ...Option) (json.RawMessage, error)
```

#### func (*API) GetStorefrontCollection

```go
func (a *API) GetStorefrontCollection(ctx context.Context, collectionID snowflake.ID, query discord.GetStorefrontCollectionQuery, opts ...Option) (*discord.GetStorefrontCollectionResponse, error)
```

#### func (*API) GetStorefrontProduct

```go
func (a *API) GetStorefrontProduct(ctx context.Context, productID snowflake.ID, opts ...Option) (*discord.StorefrontProduct, error)
```

#### func (*API) GetStorefrontProductBySKUID

```go
func (a *API) GetStorefrontProductBySKUID(ctx context.Context, skuid snowflake.ID, opts ...Option) (json.RawMessage, error)
```

#### func (*API) GetStreamPreview

```go
func (a *API) GetStreamPreview(ctx context.Context, streamKey string, opts ...Option) (*discord.GetStreamPreviewResponse, error)
```

#### func (*API) GetStripeSetupIntentSecret

```go
func (a *API) GetStripeSetupIntentSecret(ctx context.Context, params discord.GetStripeSetupIntentSecretParams, opts ...Option) (*discord.GetStripeSetupIntentSecretResponse, error)
```

#### func (*API) GetSubscription

```go
func (a *API) GetSubscription(ctx context.Context, subscriptionID snowflake.ID, opts ...Option) (*discord.Subscription, error)
```

#### func (*API) GetSubscriptionGroupListingBySubscriptionPlan

```go
func (a *API) GetSubscriptionGroupListingBySubscriptionPlan(ctx context.Context, subscriptionPlanID snowflake.ID, opts ...Option) (*discord.GetSubscriptionGroupListingBySubscriptionPlanResponse, error)
```

#### func (*API) GetSubscriptionPreview

```go
func (a *API) GetSubscriptionPreview(ctx context.Context, subscriptionID snowflake.ID, opts ...Option) (*discord.Invoice, error)
```

#### func (*API) GetSupplementalMessageRequestData

```go
func (a *API) GetSupplementalMessageRequestData(ctx context.Context, query discord.GetSupplementalMessageRequestDataQuery, opts ...Option) ([]discord.SupplementalMessageRequest, error)
```

#### func (*API) GetSuspendedUserSafetyHub

```go
func (a *API) GetSuspendedUserSafetyHub(ctx context.Context, params discord.GetSuspendedUserSafetyHubParams, opts ...Option) (*discord.SafetyHub, error)
```

#### func (*API) GetTeam

```go
func (a *API) GetTeam(ctx context.Context, teamID snowflake.ID, opts ...Option) (*discord.Team, error)
```

#### func (*API) GetTeamIdentityVerification

```go
func (a *API) GetTeamIdentityVerification(ctx context.Context, teamID snowflake.ID, opts ...Option) (*discord.UserIdentityVerification, error)
```

#### func (*API) GetTeamPayoutOnboarding

```go
func (a *API) GetTeamPayoutOnboarding(ctx context.Context, teamID snowflake.ID, opts ...Option) (*discord.GetTeamPayoutOnboardingResponse, error)
```

#### func (*API) GetTeamPayoutReport

```go
func (a *API) GetTeamPayoutReport(ctx context.Context, teamID snowflake.ID, payoutID snowflake.ID, query discord.GetTeamPayoutReportQuery, opts ...Option) (json.RawMessage, error)
```

#### func (*API) GetTeamStripeConnectURL

```go
func (a *API) GetTeamStripeConnectURL(ctx context.Context, teamID snowflake.ID, params discord.GetTeamStripeConnectURLParams, opts ...Option) (*discord.GetTeamStripeConnectURLResponse, error)
```

#### func (*API) GetThreadMember

```go
func (a *API) GetThreadMember(ctx context.Context, channelID snowflake.ID, userID snowflake.ID, query discord.GetThreadMemberQuery, opts ...Option) (*discord.ThreadMember, error)
```

#### func (*API) GetTutorial

```go
func (a *API) GetTutorial(ctx context.Context, opts ...Option) error
```

#### func (*API) GetUnauthenticatedReportCapabilities

```go
func (a *API) GetUnauthenticatedReportCapabilities(ctx context.Context, opts ...Option) (*discord.GetUnauthenticatedReportCapabilitiesResponse, error)
```

#### func (*API) GetUnauthenticatedReportMenu

```go
func (a *API) GetUnauthenticatedReportMenu(ctx context.Context, typeArg string, query discord.GetUnauthenticatedReportMenuQuery, opts ...Option) (*discord.ReportMenu, error)
```

#### func (*API) GetUnauthenticatedReportVerificationCode

```go
func (a *API) GetUnauthenticatedReportVerificationCode(ctx context.Context, typeArg string, params discord.GetUnauthenticatedReportVerificationCodeParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) GetUser

```go
func (a *API) GetUser(ctx context.Context, userID snowflake.ID, opts ...Option) (*discord.User, error)
```

#### func (*API) GetUserApplicationCommandIndex

```go
func (a *API) GetUserApplicationCommandIndex(ctx context.Context, query discord.ApplicationCommandsGetUserApplicationCommandIndexQuery, params discord.GetUserApplicationCommandIndexParamsApplicationCommandsXxxx, opts ...Option) (*discord.ApplicationCommandIndex, error)
```

#### func (*API) GetUserConnectionAccessToken

```go
func (a *API) GetUserConnectionAccessToken(ctx context.Context, connectionType string, connectionID snowflake.ID, opts ...Option) (*discord.GetUserConnectionAccessTokenResponse, error)
```

#### func (*API) GetUserConsents

```go
func (a *API) GetUserConsents(ctx context.Context, opts ...Option) (json.RawMessage, error)
```

#### func (*API) GetUserHarvest

```go
func (a *API) GetUserHarvest(ctx context.Context, opts ...Option) (*discord.Harvest, error)
```

#### func (*API) GetUserIdentityVerification

```go
func (a *API) GetUserIdentityVerification(ctx context.Context, opts ...Option) (*discord.UserIdentityVerification, error)
```

#### func (*API) GetUserNote

```go
func (a *API) GetUserNote(ctx context.Context, userID snowflake.ID, opts ...Option) (*discord.GetUserNoteResponse, error)
```

#### func (*API) GetUserNotes

```go
func (a *API) GetUserNotes(ctx context.Context, opts ...Option) (json.RawMessage, error)
```

#### func (*API) GetUserOffer

```go
func (a *API) GetUserOffer(ctx context.Context, params discord.GetUserOfferParams, opts ...Option) (*discord.GetUserOfferResponse, error)
```

#### func (*API) GetUserPremiumUsage

```go
func (a *API) GetUserPremiumUsage(ctx context.Context, opts ...Option) (*discord.GetUserPremiumUsageResponse, error)
```

#### func (*API) GetUserProfile

```go
func (a *API) GetUserProfile(ctx context.Context, userID snowflake.ID, query discord.GetUserProfileQuery, opts ...Option) (*discord.GetUserProfileResponse, error)
```

#### func (*API) GetUserSafetyHub

```go
func (a *API) GetUserSafetyHub(ctx context.Context, opts ...Option) (*discord.SafetyHub, error)
```

#### func (*API) GetUserSettings

```go
func (a *API) GetUserSettings(ctx context.Context, opts ...Option) (json.RawMessage, error)
```

#### func (*API) GetUserSettingsProto

```go
func (a *API) GetUserSettingsProto(ctx context.Context, userSettingsProtoType string, opts ...Option) (*discord.GetUserSettingsProtoResponse, error)
```

#### func (*API) GetUserSurvey

```go
func (a *API) GetUserSurvey(ctx context.Context, query discord.GetUserSurveyQuery, opts ...Option) (*discord.GetUserSurveyResponse, error)
```

#### func (*API) GetUserTrialOffer

```go
func (a *API) GetUserTrialOffer(ctx context.Context, opts ...Option) (*discord.UserTrialOffer, error)
```

#### func (*API) GetUserVoiceState

```go
func (a *API) GetUserVoiceState(ctx context.Context, guildID snowflake.ID, userID snowflake.ID, opts ...Option) (json.RawMessage, error)
```

#### func (*API) GetValidCollectiblesGiftRecipient

```go
func (a *API) GetValidCollectiblesGiftRecipient(ctx context.Context, query discord.GetValidCollectiblesGiftRecipientQuery, opts ...Option) (*discord.GiftEligibility, error)
```

#### func (*API) GetValidCollectiblesGiftRecipientsBatch

```go
func (a *API) GetValidCollectiblesGiftRecipientsBatch(ctx context.Context, query discord.GetValidCollectiblesGiftRecipientsBatchQuery, opts ...Option) (json.RawMessage, error)
```

#### func (*API) GetVirtualCurrencyBalance

```go
func (a *API) GetVirtualCurrencyBalance(ctx context.Context, opts ...Option) (*discord.GetVirtualCurrencyBalanceResponse, error)
```

#### func (*API) GetWebhook

```go
func (a *API) GetWebhook(ctx context.Context, webhookID snowflake.ID, opts ...Option) (*discord.Webhook, error)
```

#### func (*API) GetWebhookMessage

```go
func (a *API) GetWebhookMessage(ctx context.Context, webhookID snowflake.ID, webhookToken string, messageID snowflake.ID, opts ...Option) (json.RawMessage, error)
```

#### func (*API) GetWebhookWithToken

```go
func (a *API) GetWebhookWithToken(ctx context.Context, webhookID snowflake.ID, webhookToken string, opts ...Option) (json.RawMessage, error)
```

#### func (*API) HideMessageFromGuildFeed

```go
func (a *API) HideMessageFromGuildFeed(ctx context.Context, channelID snowflake.ID, messageID snowflake.ID, opts ...Option) error
```

#### func (*API) IgnoreUser

```go
func (a *API) IgnoreUser(ctx context.Context, userID snowflake.ID, opts ...Option) error
```

#### func (*API) JoinAdminCommunity

```go
func (a *API) JoinAdminCommunity(ctx context.Context, guildID snowflake.ID, opts ...Option) (json.RawMessage, error)
```

#### func (*API) JoinGuild

```go
func (a *API) JoinGuild(ctx context.Context, guildID snowflake.ID, query discord.JoinGuildQuery, opts ...Option) (*discord.Guild, error)
```

#### func (*API) JoinHypesquadOnline

```go
func (a *API) JoinHypesquadOnline(ctx context.Context, params discord.JoinHypesquadOnlineParams, opts ...Option) error
```

#### func (*API) JoinIntegrationGuild

```go
func (a *API) JoinIntegrationGuild(ctx context.Context, integrationID snowflake.ID, opts ...Option) error
```

#### func (*API) JoinOrCreateLobby

```go
func (a *API) JoinOrCreateLobby(ctx context.Context, params discord.JoinOrCreateLobbyParams, opts ...Option) (*discord.Lobby, error)
```

#### func (*API) JoinStudentHub

```go
func (a *API) JoinStudentHub(ctx context.Context, params discord.JoinStudentHubParams, opts ...Option) (*discord.JoinStudentHubResponse, error)
```

#### func (*API) JoinStudentHubWaitlist

```go
func (a *API) JoinStudentHubWaitlist(ctx context.Context, params discord.JoinStudentHubWaitlistParams, opts ...Option) (*discord.JoinStudentHubWaitlistResponse, error)
```

#### func (*API) JoinThread

```go
func (a *API) JoinThread(ctx context.Context, channelID snowflake.ID, query discord.JoinThreadQuery, opts ...Option) error
```

#### func (*API) LaunchEmbeddedActivity

```go
func (a *API) LaunchEmbeddedActivity(ctx context.Context, channelID snowflake.ID, applicationID snowflake.ID, params discord.LaunchEmbeddedActivityParams, opts ...Option) error
```

#### func (*API) LeaveGuild

```go
func (a *API) LeaveGuild(ctx context.Context, guildID snowflake.ID, params discord.LeaveGuildParams, opts ...Option) error
```

#### func (*API) LeaveHypesquadOnline

```go
func (a *API) LeaveHypesquadOnline(ctx context.Context, opts ...Option) error
```

#### func (*API) LeaveLobby

```go
func (a *API) LeaveLobby(ctx context.Context, lobbyID snowflake.ID, opts ...Option) error
```

#### func (*API) LeaveThread

```go
func (a *API) LeaveThread(ctx context.Context, channelID snowflake.ID, query discord.LeaveThreadQuery, opts ...Option) error
```

#### func (*API) ListActiveThreads

```go
func (a *API) ListActiveThreads(ctx context.Context, channelID snowflake.ID, opts ...Option) (*discord.ListActiveThreadsResponse, error)
```

#### func (*API) ListApplicationActivityStatistics

```go
func (a *API) ListApplicationActivityStatistics(ctx context.Context, applicationID snowflake.ID, opts ...Option) ([]discord.ApplicationActivityStatistics, error)
```

#### func (*API) ListApplicationDirectoryCategories

```go
func (a *API) ListApplicationDirectoryCategories(ctx context.Context, query discord.ListApplicationDirectoryCategoriesQuery, opts ...Option) ([]discord.ApplicationDirectoryCategory, error)
```

#### func (*API) ListApplicationDirectoryCollections

```go
func (a *API) ListApplicationDirectoryCollections(ctx context.Context, query discord.ListApplicationDirectoryCollectionsQuery, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ListApplicationDirectorySimilarApplications

```go
func (a *API) ListApplicationDirectorySimilarApplications(ctx context.Context, applicationID snowflake.ID, query discord.ListApplicationDirectorySimilarApplicationsQuery, opts ...Option) (*discord.ListApplicationDirectorySimilarApplicationsResponse, error)
```

#### func (*API) ListApplicationDisclosures

```go
func (a *API) ListApplicationDisclosures(ctx context.Context, applicationID snowflake.ID, opts ...Option) (*discord.ListApplicationDisclosuresResponse, error)
```

#### func (*API) ListApplicationEmojis

```go
func (a *API) ListApplicationEmojis(ctx context.Context, applicationID snowflake.ID, opts ...Option) (*discord.ListApplicationEmojisResponse, error)
```

#### func (*API) ListApplicationExternalIdentityProviderConfigurations

```go
func (a *API) ListApplicationExternalIdentityProviderConfigurations(ctx context.Context, applicationID snowflake.ID, params discord.ListApplicationExternalIdentityProviderConfigurationsParams, opts ...Option) ([]discord.ApplicationExternalIdentityProviderConfiguration, error)
```

#### func (*API) ListApplicationGameClaims

```go
func (a *API) ListApplicationGameClaims(ctx context.Context, applicationID snowflake.ID, opts ...Option) ([]discord.GameClaim, error)
```

#### func (*API) ListApplicationGiftCodeBatches

```go
func (a *API) ListApplicationGiftCodeBatches(ctx context.Context, applicationID snowflake.ID, opts ...Option) ([]discord.GiftCodeBatch, error)
```

#### func (*API) ListApplicationManagedLinks

```go
func (a *API) ListApplicationManagedLinks(ctx context.Context, applicationID snowflake.ID, opts ...Option) ([]discord.ActivityLink, error)
```

#### func (*API) ListApplicationPublishedStoreListings

```go
func (a *API) ListApplicationPublishedStoreListings(ctx context.Context, query discord.ListApplicationPublishedStoreListingsQuery, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ListApplicationTesters

```go
func (a *API) ListApplicationTesters(ctx context.Context, applicationID snowflake.ID, opts ...Option) ([]discord.WhitelistedUser, error)
```

#### func (*API) ListApplicationWidgetConfigs

```go
func (a *API) ListApplicationWidgetConfigs(ctx context.Context, applicationID snowflake.ID, opts ...Option) ([]discord.WidgetConfig, error)
```

#### func (*API) ListApplications

```go
func (a *API) ListApplications(ctx context.Context, query discord.ListApplicationsQuery, opts ...Option) ([]discord.Application, error)
```

#### func (*API) ListApplicationsWithAssets

```go
func (a *API) ListApplicationsWithAssets(ctx context.Context, query discord.ListApplicationsWithAssetsQuery, opts ...Option) (*discord.ListApplicationsWithAssetsResponse, error)
```

#### func (*API) ListAppliedPremiumGuildSubscriptions

```go
func (a *API) ListAppliedPremiumGuildSubscriptions(ctx context.Context, query discord.ListAppliedPremiumGuildSubscriptionsQuery, opts ...Option) ([]discord.PremiumGuildSubscription, error)
```

#### func (*API) ListAvailableAdyenPaymentMethods

```go
func (a *API) ListAvailableAdyenPaymentMethods(ctx context.Context, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ListBogoPromotions

```go
func (a *API) ListBogoPromotions(ctx context.Context, query discord.ListBogoPromotionsQuery, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ListBulkApplicationIdentities

```go
func (a *API) ListBulkApplicationIdentities(ctx context.Context, params discord.WidgetsListBulkApplicationIdentitiesParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ListBulkApplicationPrimaryStoreListing

```go
func (a *API) ListBulkApplicationPrimaryStoreListing(ctx context.Context, query discord.ListBulkApplicationPrimaryStoreListingQuery, opts ...Option) ([]discord.StoreListing, error)
```

#### func (*API) ListBulkPublishedSubscriptionPlans

```go
func (a *API) ListBulkPublishedSubscriptionPlans(ctx context.Context, query discord.ListBulkPublishedSubscriptionPlansQuery, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ListChannelAffinities

```go
func (a *API) ListChannelAffinities(ctx context.Context, opts ...Option) (*discord.ListChannelAffinitiesResponse, error)
```

#### func (*API) ListChannelIntegrations

```go
func (a *API) ListChannelIntegrations(ctx context.Context, channelID snowflake.ID, opts ...Option) ([]discord.Integration, error)
```

#### func (*API) ListChannelInvites

```go
func (a *API) ListChannelInvites(ctx context.Context, channelID snowflake.ID, opts ...Option) ([]discord.Invite, error)
```

#### func (*API) ListChannelWebhooks

```go
func (a *API) ListChannelWebhooks(ctx context.Context, channelID snowflake.ID, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ListClaimedPromotions

```go
func (a *API) ListClaimedPromotions(ctx context.Context, query discord.ListClaimedPromotionsQuery, opts ...Option) ([]discord.ClaimedPromotion, error)
```

#### func (*API) ListClaimedQuests

```go
func (a *API) ListClaimedQuests(ctx context.Context, opts ...Option) (*discord.ListClaimedQuestsResponse, error)
```

#### func (*API) ListCollectiblesCategories

```go
func (a *API) ListCollectiblesCategories(ctx context.Context, query discord.ListCollectiblesCategoriesQuery, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ListCollectiblesCategoriesV2

```go
func (a *API) ListCollectiblesCategoriesV2(ctx context.Context, query discord.ListCollectiblesCategoriesV2Query, opts ...Option) (*discord.ListCollectiblesCategoriesV2Response, error)
```

#### func (*API) ListConsoleDevices

```go
func (a *API) ListConsoleDevices(ctx context.Context, connectionType string, opts ...Option) (*discord.ListConsoleDevicesResponse, error)
```

#### func (*API) ListConversationSummaries

```go
func (a *API) ListConversationSummaries(ctx context.Context, channelID snowflake.ID, opts ...Option) (*discord.ListConversationSummariesResponse, error)
```

#### func (*API) ListCurrentUserQuests

```go
func (a *API) ListCurrentUserQuests(ctx context.Context, opts ...Option) (*discord.ListCurrentUserQuestsResponse, error)
```

#### func (*API) ListDMMessages

```go
func (a *API) ListDMMessages(ctx context.Context, userID snowflake.ID, query discord.ListDMMessagesQuery, opts ...Option) ([]discord.Message, error)
```

#### func (*API) ListDefaultSoundboardSounds

```go
func (a *API) ListDefaultSoundboardSounds(ctx context.Context, opts ...Option) ([]discord.SoundboardSound, error)
```

#### func (*API) ListDetectableGameExclusions

```go
func (a *API) ListDetectableGameExclusions(ctx context.Context, opts ...Option) (*discord.ListDetectableGameExclusionsResponse, error)
```

#### func (*API) ListDetectableGames

```go
func (a *API) ListDetectableGames(ctx context.Context, opts ...Option) ([]discord.DetectableApplication, error)
```

#### func (*API) ListDetectableNonGameApplications

```go
func (a *API) ListDetectableNonGameApplications(ctx context.Context, opts ...Option) ([]discord.DetectableApplication, error)
```

#### func (*API) ListDirectoryEntries

```go
func (a *API) ListDirectoryEntries(ctx context.Context, channelID snowflake.ID, opts ...Option) ([]discord.DirectoryEntry, error)
```

#### func (*API) ListDiscoverableGuilds

```go
func (a *API) ListDiscoverableGuilds(ctx context.Context, query discord.ListDiscoverableGuildsQuery, opts ...Option) (*discord.ListDiscoverableGuildsResponse, error)
```

#### func (*API) ListDiscoveryCategories

```go
func (a *API) ListDiscoveryCategories(ctx context.Context, query discord.ListDiscoveryCategoriesQuery, opts ...Option) ([]discord.DiscoveryCategory, error)
```

#### func (*API) ListEligibleApplicationSubscriptionGuilds

```go
func (a *API) ListEligibleApplicationSubscriptionGuilds(ctx context.Context, query discord.ListEligibleApplicationSubscriptionGuildsQuery, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ListEmbeddedActivities

```go
func (a *API) ListEmbeddedActivities(ctx context.Context, query discord.ListEmbeddedActivitiesQuery, opts ...Option) (*discord.ListEmbeddedActivitiesResponse, error)
```

#### func (*API) ListEmbeddedActivityInstances

```go
func (a *API) ListEmbeddedActivityInstances(ctx context.Context, applicationID snowflake.ID, channelID snowflake.ID, opts ...Option) (*discord.ListEmbeddedActivityInstancesResponse, error)
```

#### func (*API) ListFeaturedApplicationWidgetConfigs

```go
func (a *API) ListFeaturedApplicationWidgetConfigs(ctx context.Context, opts ...Option) (*discord.ListFeaturedApplicationWidgetConfigsResponse, error)
```

#### func (*API) ListFriendSuggestions

```go
func (a *API) ListFriendSuggestions(ctx context.Context, opts ...Option) ([]discord.FriendSuggestion, error)
```

#### func (*API) ListGameAnnouncements

```go
func (a *API) ListGameAnnouncements(ctx context.Context, gameID snowflake.ID, query discord.ListGameAnnouncementsQuery, opts ...Option) (*discord.ListGameAnnouncementsResponse, error)
```

#### func (*API) ListGameRelationships

```go
func (a *API) ListGameRelationships(ctx context.Context, opts ...Option) ([]discord.GameRelationship, error)
```

#### func (*API) ListGames

```go
func (a *API) ListGames(ctx context.Context, query discord.ListGamesQuery, opts ...Option) ([]discord.Game, error)
```

#### func (*API) ListGlobalActivityStatistics

```go
func (a *API) ListGlobalActivityStatistics(ctx context.Context, query discord.ListGlobalActivityStatisticsQuery, opts ...Option) ([]discord.GlobalActivityStatistics, error)
```

#### func (*API) ListGuildActiveThreads

```go
func (a *API) ListGuildActiveThreads(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.ListGuildActiveThreadsResponse, error)
```

#### func (*API) ListGuildAffinities

```go
func (a *API) ListGuildAffinities(ctx context.Context, opts ...Option) (*discord.ListGuildAffinitiesResponse, error)
```

#### func (*API) ListGuildApplications

```go
func (a *API) ListGuildApplications(ctx context.Context, guildID snowflake.ID, query discord.ListGuildApplicationsQuery, opts ...Option) ([]discord.Application, error)
```

#### func (*API) ListGuildAudienceNewMembersByDiscordTenure

```go
func (a *API) ListGuildAudienceNewMembersByDiscordTenure(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.AudienceNewMembersByDiscordTenure, error)
```

#### func (*API) ListGuildAudienceParticipatorsByGuildTenure

```go
func (a *API) ListGuildAudienceParticipatorsByGuildTenure(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.AudienceParticipatorsByGuildTenure, error)
```

#### func (*API) ListGuildAudienceParticipatorsByPlatform

```go
func (a *API) ListGuildAudienceParticipatorsByPlatform(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.AudienceParticipatorsByPlatform, error)
```

#### func (*API) ListGuildAudienceParticipatorsByRegistrationCountry

```go
func (a *API) ListGuildAudienceParticipatorsByRegistrationCountry(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.AudienceParticipatorsByRegistrationCountry, error)
```

#### func (*API) ListGuildAuditLogEntries

```go
func (a *API) ListGuildAuditLogEntries(ctx context.Context, guildID snowflake.ID, query discord.ListGuildAuditLogEntriesQuery, opts ...Option) (*discord.ListGuildAuditLogEntriesResponse, error)
```

#### func (*API) ListGuildAutomodRules

```go
func (a *API) ListGuildAutomodRules(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.AutomodRule, error)
```

#### func (*API) ListGuildBans

```go
func (a *API) ListGuildBans(ctx context.Context, guildID snowflake.ID, query discord.ListGuildBansQuery, opts ...Option) ([]discord.Ban, error)
```

#### func (*API) ListGuildChannelFollowingByChannel

```go
func (a *API) ListGuildChannelFollowingByChannel(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.ChannelFollowingByChannel, error)
```

#### func (*API) ListGuildChannelFollowingGuildSize

```go
func (a *API) ListGuildChannelFollowingGuildSize(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.ChannelFollowingGuildSize, error)
```

#### func (*API) ListGuildChannelFollowingGuildSizeByChannel

```go
func (a *API) ListGuildChannelFollowingGuildSizeByChannel(ctx context.Context, guildID snowflake.ID, query discord.ListGuildChannelFollowingGuildSizeByChannelQuery, opts ...Option) ([]discord.ChannelFollowingGuildSizeByChannel, error)
```

#### func (*API) ListGuildChannelFollowingOverview

```go
func (a *API) ListGuildChannelFollowingOverview(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.ChannelFollowingOverview, error)
```

#### func (*API) ListGuildChannelFollowingReach

```go
func (a *API) ListGuildChannelFollowingReach(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.ChannelFollowingReach, error)
```

#### func (*API) ListGuildChannels

```go
func (a *API) ListGuildChannels(ctx context.Context, guildID snowflake.ID, query discord.ListGuildChannelsQuery, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ListGuildEmojis

```go
func (a *API) ListGuildEmojis(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.Emoji, error)
```

#### func (*API) ListGuildEngagementBase

```go
func (a *API) ListGuildEngagementBase(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.EngagementBase, error)
```

#### func (*API) ListGuildEngagementMuters

```go
func (a *API) ListGuildEngagementMuters(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.EngagementMuters, error)
```

#### func (*API) ListGuildEngagementOverview

```go
func (a *API) ListGuildEngagementOverview(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.EngagementOverview, error)
```

#### func (*API) ListGuildEngagementPruneableMembers

```go
func (a *API) ListGuildEngagementPruneableMembers(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.EngagementPruneableMembers, error)
```

#### func (*API) ListGuildEngagementTextChannels

```go
func (a *API) ListGuildEngagementTextChannels(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.TextChannelEngagement, error)
```

#### func (*API) ListGuildEngagementVoiceChannels

```go
func (a *API) ListGuildEngagementVoiceChannels(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.VoiceChannelEngagement, error)
```

#### func (*API) ListGuildEntitlements

```go
func (a *API) ListGuildEntitlements(ctx context.Context, guildID snowflake.ID, query discord.ListGuildEntitlementsQuery, opts ...Option) ([]discord.Entitlement, error)
```

#### func (*API) ListGuildGameServerRegions

```go
func (a *API) ListGuildGameServerRegions(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.GameServerRegion, error)
```

#### func (*API) ListGuildGameServers

```go
func (a *API) ListGuildGameServers(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.GameServer, error)
```

#### func (*API) ListGuildGrowthActivationJoins

```go
func (a *API) ListGuildGrowthActivationJoins(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.GrowthActivationJoins, error)
```

#### func (*API) ListGuildGrowthActivationJoinsByInvite

```go
func (a *API) ListGuildGrowthActivationJoinsByInvite(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.GrowthActivationJoins, error)
```

#### func (*API) ListGuildGrowthActivationJoinsByReferrer

```go
func (a *API) ListGuildGrowthActivationJoinsByReferrer(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.GrowthActivationJoinsByReferrer, error)
```

#### func (*API) ListGuildGrowthActivationJoinsBySources

```go
func (a *API) ListGuildGrowthActivationJoinsBySources(ctx context.Context, guildID snowflake.ID, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ListGuildGrowthActivationLeavers

```go
func (a *API) ListGuildGrowthActivationLeavers(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.GrowthActivationLeavers, error)
```

#### func (*API) ListGuildGrowthActivationMembership

```go
func (a *API) ListGuildGrowthActivationMembership(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.GrowthActivationMembership, error)
```

#### func (*API) ListGuildGrowthActivationOverview

```go
func (a *API) ListGuildGrowthActivationOverview(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.GrowthActivationOverview, error)
```

#### func (*API) ListGuildGrowthActivationPercentages

```go
func (a *API) ListGuildGrowthActivationPercentages(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.GrowthActivationPercentage, error)
```

#### func (*API) ListGuildGrowthActivationRetention

```go
func (a *API) ListGuildGrowthActivationRetention(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.GrowthActivationRetention, error)
```

#### func (*API) ListGuildIntegrations

```go
func (a *API) ListGuildIntegrations(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.Integration, error)
```

#### func (*API) ListGuildInvites

```go
func (a *API) ListGuildInvites(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.Invite, error)
```

#### func (*API) ListGuildJoinRequests

```go
func (a *API) ListGuildJoinRequests(ctx context.Context, guildID snowflake.ID, query discord.ListGuildJoinRequestsQuery, opts ...Option) (*discord.ListGuildJoinRequestsResponse, error)
```

#### func (*API) ListGuildMembers

```go
func (a *API) ListGuildMembers(ctx context.Context, guildID snowflake.ID, query discord.ListGuildMembersQuery, opts ...Option) ([]discord.GuildMember, error)
```

#### func (*API) ListGuildMembersSupplemental

```go
func (a *API) ListGuildMembersSupplemental(ctx context.Context, guildID snowflake.ID, params discord.ListGuildMembersSupplementalParams, opts ...Option) ([]discord.SupplementalGuildMember, error)
```

#### func (*API) ListGuildMembersWithUnusualDMActivity

```go
func (a *API) ListGuildMembersWithUnusualDMActivity(ctx context.Context, guildID snowflake.ID, query discord.ListGuildMembersWithUnusualDMActivityQuery, opts ...Option) ([]discord.GuildMemberUnusualDMActivity, error)
```

#### func (*API) ListGuildOnboardingAllowedApplications

```go
func (a *API) ListGuildOnboardingAllowedApplications(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.ListGuildOnboardingAllowedApplicationsResponse, error)
```

#### func (*API) ListGuildPowerups

```go
func (a *API) ListGuildPowerups(ctx context.Context, guildID snowflake.ID, query discord.ListGuildPowerupsQuery, opts ...Option) ([]discord.Entitlement, error)
```

#### func (*API) ListGuildProductListings

```go
func (a *API) ListGuildProductListings(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.ListGuildProductListingsResponse, error)
```

#### func (*API) ListGuildRoleConnectionsConfigurations

```go
func (a *API) ListGuildRoleConnectionsConfigurations(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.RoleConnectionRule, error)
```

#### func (*API) ListGuildRoleMembers

```go
func (a *API) ListGuildRoleMembers(ctx context.Context, guildID snowflake.ID, roleID snowflake.ID, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ListGuildRoleSubscriptionGroupListings

```go
func (a *API) ListGuildRoleSubscriptionGroupListings(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.RoleSubscriptionGroupListing, error)
```

#### func (*API) ListGuildRoleSubscriptionListingTemplates

```go
func (a *API) ListGuildRoleSubscriptionListingTemplates(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.ListGuildRoleSubscriptionListingTemplatesResponse, error)
```

#### func (*API) ListGuildRoleSubscriptionTrials

```go
func (a *API) ListGuildRoleSubscriptionTrials(ctx context.Context, guildID snowflake.ID, params discord.ListGuildRoleSubscriptionTrialsParams, opts ...Option) ([]discord.RoleSubscriptionTrial, error)
```

#### func (*API) ListGuildRoles

```go
func (a *API) ListGuildRoles(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.Role, error)
```

#### func (*API) ListGuildScheduledEventUsers

```go
func (a *API) ListGuildScheduledEventUsers(ctx context.Context, guildID snowflake.ID, guildScheduledEventID snowflake.ID, query discord.ListGuildScheduledEventUsersQuery, opts ...Option) (*discord.ListGuildScheduledEventUsersResponse, error)
```

#### func (*API) ListGuildScheduledEvents

```go
func (a *API) ListGuildScheduledEvents(ctx context.Context, guildID snowflake.ID, query discord.ListGuildScheduledEventsQuery, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ListGuildSoundboardSounds

```go
func (a *API) ListGuildSoundboardSounds(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.ListGuildSoundboardSoundsResponse, error)
```

#### func (*API) ListGuildStickers

```go
func (a *API) ListGuildStickers(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.Sticker, error)
```

#### func (*API) ListGuildTemplates

```go
func (a *API) ListGuildTemplates(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.GuildTemplate, error)
```

#### func (*API) ListGuildTopEmojis

```go
func (a *API) ListGuildTopEmojis(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.ListGuildTopEmojisResponse, error)
```

#### func (*API) ListGuildTopGames

```go
func (a *API) ListGuildTopGames(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.ListGuildTopGamesResponse, error)
```

#### func (*API) ListGuildTopReadChannels

```go
func (a *API) ListGuildTopReadChannels(ctx context.Context, guildID snowflake.ID, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ListGuildVoiceRegions

```go
func (a *API) ListGuildVoiceRegions(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.VoiceRegion, error)
```

#### func (*API) ListGuildWebhooks

```go
func (a *API) ListGuildWebhooks(ctx context.Context, guildID snowflake.ID, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ListGuildWelcomeScreenFunnel

```go
func (a *API) ListGuildWelcomeScreenFunnel(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.WelcomeScreenFunnel, error)
```

#### func (*API) ListGuildWelcomeScreenUsers

```go
func (a *API) ListGuildWelcomeScreenUsers(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.WelcomeScreenUsers, error)
```

#### func (*API) ListInviteFriendMembers

```go
func (a *API) ListInviteFriendMembers(ctx context.Context, inviteCode string, opts ...Option) (*discord.ListInviteFriendMembersResponse, error)
```

#### func (*API) ListInviteTargetUsers

```go
func (a *API) ListInviteTargetUsers(ctx context.Context, inviteCode string, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ListJoinRequestGuilds

```go
func (a *API) ListJoinRequestGuilds(ctx context.Context, opts ...Option) ([]discord.Guild, error)
```

#### func (*API) ListJoinedPrivateArchivedThreads

```go
func (a *API) ListJoinedPrivateArchivedThreads(ctx context.Context, channelID snowflake.ID, query discord.ListJoinedPrivateArchivedThreadsQuery, opts ...Option) (*discord.ListJoinedPrivateArchivedThreadsResponse, error)
```

#### func (*API) ListLobbyMessages

```go
func (a *API) ListLobbyMessages(ctx context.Context, lobbyID snowflake.ID, query discord.ListLobbyMessagesQuery, opts ...Option) ([]discord.Message, error)
```

#### func (*API) ListMessagePins

```go
func (a *API) ListMessagePins(ctx context.Context, channelID snowflake.ID, query discord.ListMessagePinsQuery, opts ...Option) (*discord.ListMessagePinsResponse, error)
```

#### func (*API) ListMessages

```go
func (a *API) ListMessages(ctx context.Context, channelID snowflake.ID, query discord.ListMessagesQuery, opts ...Option) ([]discord.Message, error)
```

#### func (*API) ListMutualRelationships

```go
func (a *API) ListMutualRelationships(ctx context.Context, userID snowflake.ID, opts ...Option) ([]discord.User, error)
```

#### func (*API) ListNotificationCenterItems

```go
func (a *API) ListNotificationCenterItems(ctx context.Context, query discord.ListNotificationCenterItemsQuery, opts ...Option) (*discord.ListNotificationCenterItemsResponse, error)
```

#### func (*API) ListNotificationSettingsSnapshots

```go
func (a *API) ListNotificationSettingsSnapshots(ctx context.Context, opts ...Option) ([]discord.NotificationSettingsSnapshot, error)
```

#### func (*API) ListOAuth2ApplicationAssets

```go
func (a *API) ListOAuth2ApplicationAssets(ctx context.Context, applicationID snowflake.ID, query discord.ListOAuth2ApplicationAssetsQuery, opts ...Option) ([]discord.ApplicationAsset, error)
```

#### func (*API) ListOutboundPromotions

```go
func (a *API) ListOutboundPromotions(ctx context.Context, query discord.ListOutboundPromotionsQuery, opts ...Option) ([]discord.Promotion, error)
```

#### func (*API) ListPartialApplications

```go
func (a *API) ListPartialApplications(ctx context.Context, query discord.ListPartialApplicationsQuery, opts ...Option) ([]discord.Application, error)
```

#### func (*API) ListPartialDirectoryEntries

```go
func (a *API) ListPartialDirectoryEntries(ctx context.Context, channelID snowflake.ID, opts ...Option) ([]discord.DirectoryEntry, error)
```

#### func (*API) ListPaymentSources

```go
func (a *API) ListPaymentSources(ctx context.Context, opts ...Option) ([]discord.PaymentSource, error)
```

#### func (*API) ListPayments

```go
func (a *API) ListPayments(ctx context.Context, opts ...Option) ([]discord.Payment, error)
```

#### func (*API) ListPinnedMessages

```go
func (a *API) ListPinnedMessages(ctx context.Context, channelID snowflake.ID, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ListPollAnswerVoters

```go
func (a *API) ListPollAnswerVoters(ctx context.Context, channelID snowflake.ID, messageID snowflake.ID, pollAnswerID snowflake.ID, query discord.ListPollAnswerVotersQuery, opts ...Option) (*discord.ListPollAnswerVotersResponse, error)
```

#### func (*API) ListPremiumGuildSubscriptionSlots

```go
func (a *API) ListPremiumGuildSubscriptionSlots(ctx context.Context, opts ...Option) ([]discord.PremiumGuildSubscriptionSlot, error)
```

#### func (*API) ListPremiumGuildSubscriptions

```go
func (a *API) ListPremiumGuildSubscriptions(ctx context.Context, guildID snowflake.ID, query discord.ListPremiumGuildSubscriptionsQuery, opts ...Option) ([]discord.PremiumGuildSubscription, error)
```

#### func (*API) ListPremiumReferralEligibleUsers

```go
func (a *API) ListPremiumReferralEligibleUsers(ctx context.Context, params discord.ListPremiumReferralEligibleUsersParams, opts ...Option) (*discord.ListPremiumReferralEligibleUsersResponse, error)
```

#### func (*API) ListPremiumUserAffinities

```go
func (a *API) ListPremiumUserAffinities(ctx context.Context, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ListPresences

```go
func (a *API) ListPresences(ctx context.Context, opts ...Option) (*discord.ListPresencesResponse, error)
```

#### func (*API) ListPresencesForXbox

```go
func (a *API) ListPresencesForXbox(ctx context.Context, opts ...Option) (*discord.ListPresencesForXboxResponse, error)
```

#### func (*API) ListPrivateArchivedThreads

```go
func (a *API) ListPrivateArchivedThreads(ctx context.Context, channelID snowflake.ID, query discord.ListPrivateArchivedThreadsQuery, opts ...Option) (*discord.ListPrivateArchivedThreadsResponse, error)
```

#### func (*API) ListPrivateChannels

```go
func (a *API) ListPrivateChannels(ctx context.Context, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ListPromotions

```go
func (a *API) ListPromotions(ctx context.Context, query discord.ListPromotionsQuery, opts ...Option) ([]discord.Promotion, error)
```

#### func (*API) ListPublicArchivedThreads

```go
func (a *API) ListPublicArchivedThreads(ctx context.Context, channelID snowflake.ID, query discord.ListPublicArchivedThreadsQuery, opts ...Option) (*discord.ListPublicArchivedThreadsResponse, error)
```

#### func (*API) ListPublishedSubscriptionPlans

```go
func (a *API) ListPublishedSubscriptionPlans(ctx context.Context, skuid snowflake.ID, query discord.ListPublishedSubscriptionPlansQuery, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ListReactions

```go
func (a *API) ListReactions(ctx context.Context, channelID snowflake.ID, messageID snowflake.ID, emoji string, query discord.ListReactionsQuery, opts ...Option) ([]discord.User, error)
```

#### func (*API) ListRecentAvatars

```go
func (a *API) ListRecentAvatars(ctx context.Context, opts ...Option) (*discord.ListRecentAvatarsResponse, error)
```

#### func (*API) ListRecentMentions

```go
func (a *API) ListRecentMentions(ctx context.Context, query discord.ListRecentMentionsQuery, opts ...Option) ([]discord.Message, error)
```

#### func (*API) ListRelationships

```go
func (a *API) ListRelationships(ctx context.Context, opts ...Option) ([]discord.Relationship, error)
```

#### func (*API) ListSKUStoreListings

```go
func (a *API) ListSKUStoreListings(ctx context.Context, skuid snowflake.ID, query discord.ListSKUStoreListingsQuery, opts ...Option) ([]discord.StoreListing, error)
```

#### func (*API) ListSavedMessages

```go
func (a *API) ListSavedMessages(ctx context.Context, opts ...Option) (*discord.ListSavedMessagesResponse, error)
```

#### func (*API) ListSocialLayerSKUs

```go
func (a *API) ListSocialLayerSKUs(ctx context.Context, applicationID snowflake.ID, opts ...Option) ([]discord.SKU, error)
```

#### func (*API) ListSocialSdkReleases

```go
func (a *API) ListSocialSdkReleases(ctx context.Context, opts ...Option) (*discord.ListSocialSdkReleasesResponse, error)
```

#### func (*API) ListStickerPacks

```go
func (a *API) ListStickerPacks(ctx context.Context, opts ...Option) (*discord.ListStickerPacksResponse, error)
```

#### func (*API) ListStorePriceTiers

```go
func (a *API) ListStorePriceTiers(ctx context.Context, query discord.ListStorePriceTiersQuery, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ListStorefrontProductsBySKUID

```go
func (a *API) ListStorefrontProductsBySKUID(ctx context.Context, query discord.ListStorefrontProductsBySKUIDQuery, opts ...Option) (*discord.ListStorefrontProductsBySKUIDResponse, error)
```

#### func (*API) ListStorefrontSKUPrices

```go
func (a *API) ListStorefrontSKUPrices(ctx context.Context, query discord.ListStorefrontSKUPricesQuery, opts ...Option) (*discord.ListStorefrontSKUPricesResponse, error)
```

#### func (*API) ListSubscriptionInvoices

```go
func (a *API) ListSubscriptionInvoices(ctx context.Context, subscriptionID snowflake.ID, params discord.ListSubscriptionInvoicesParams, opts ...Option) ([]discord.Invoice, error)
```

#### func (*API) ListSubscriptionPlans

```go
func (a *API) ListSubscriptionPlans(ctx context.Context, skuid snowflake.ID, opts ...Option) ([]discord.SubscriptionPlan, error)
```

#### func (*API) ListSubscriptions

```go
func (a *API) ListSubscriptions(ctx context.Context, query discord.ListSubscriptionsQuery, opts ...Option) ([]discord.Subscription, error)
```

#### func (*API) ListSuggestedGIFSearchTerms

```go
func (a *API) ListSuggestedGIFSearchTerms(ctx context.Context, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ListTeamApplications

```go
func (a *API) ListTeamApplications(ctx context.Context, teamID snowflake.ID, opts ...Option) ([]discord.Application, error)
```

#### func (*API) ListTeamMembers

```go
func (a *API) ListTeamMembers(ctx context.Context, teamID snowflake.ID, opts ...Option) ([]discord.TeamMember, error)
```

#### func (*API) ListTeamPayouts

```go
func (a *API) ListTeamPayouts(ctx context.Context, teamID snowflake.ID, query discord.ListTeamPayoutsQuery, opts ...Option) ([]discord.TeamPayout, error)
```

#### func (*API) ListTeams

```go
func (a *API) ListTeams(ctx context.Context, query discord.ListTeamsQuery, opts ...Option) ([]discord.Team, error)
```

#### func (*API) ListThreadMembers

```go
func (a *API) ListThreadMembers(ctx context.Context, channelID snowflake.ID, query discord.ListThreadMembersQuery, opts ...Option) ([]discord.ThreadMember, error)
```

#### func (*API) ListTrendingGIFCategories

```go
func (a *API) ListTrendingGIFCategories(ctx context.Context, opts ...Option) (*discord.ListTrendingGIFCategoriesResponse, error)
```

#### func (*API) ListTrendingGIFSearchTerms

```go
func (a *API) ListTrendingGIFSearchTerms(ctx context.Context, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ListTrendingGifs

```go
func (a *API) ListTrendingGifs(ctx context.Context, opts ...Option) ([]discord.GIF, error)
```

#### func (*API) ListUserAffinities

```go
func (a *API) ListUserAffinities(ctx context.Context, opts ...Option) (*discord.ListUserAffinitiesResponse, error)
```

#### func (*API) ListUserAffinitiesV2

```go
func (a *API) ListUserAffinitiesV2(ctx context.Context, opts ...Option) (*discord.ListUserAffinitiesV2Response, error)
```

#### func (*API) ListUserApplicationActivityStatistics

```go
func (a *API) ListUserApplicationActivityStatistics(ctx context.Context, opts ...Option) ([]discord.UserApplicationActivityStatistics, error)
```

#### func (*API) ListUserApplicationEntitlements

```go
func (a *API) ListUserApplicationEntitlements(ctx context.Context, applicationID snowflake.ID, query discord.ListUserApplicationEntitlementsQuery, opts ...Option) ([]discord.Entitlement, error)
```

#### func (*API) ListUserApplicationIdentities

```go
func (a *API) ListUserApplicationIdentities(ctx context.Context, userID snowflake.ID, query discord.ListUserApplicationIdentitiesQuery, opts ...Option) (*discord.ListUserApplicationIdentitiesResponse, error)
```

#### func (*API) ListUserApplicationRoleConnections

```go
func (a *API) ListUserApplicationRoleConnections(ctx context.Context, params discord.ListUserApplicationRoleConnectionsParams, opts ...Option) ([]discord.ApplicationRoleConnection, error)
```

#### func (*API) ListUserConnectionSubreddits

```go
func (a *API) ListUserConnectionSubreddits(ctx context.Context, connectionID snowflake.ID, opts ...Option) ([]discord.Subreddit, error)
```

#### func (*API) ListUserConnections

```go
func (a *API) ListUserConnections(ctx context.Context, opts ...Option) ([]discord.Connection, error)
```

#### func (*API) ListUserEntitlements

```go
func (a *API) ListUserEntitlements(ctx context.Context, query discord.ListUserEntitlementsQuery, opts ...Option) ([]discord.Entitlement, error)
```

#### func (*API) ListUserGiftCodes

```go
func (a *API) ListUserGiftCodes(ctx context.Context, query discord.ListUserGiftCodesQuery, opts ...Option) ([]discord.GiftCode, error)
```

#### func (*API) ListUserGiftableEntitlements

```go
func (a *API) ListUserGiftableEntitlements(ctx context.Context, query discord.ListUserGiftableEntitlementsQuery, opts ...Option) ([]discord.Entitlement, error)
```

#### func (*API) ListUserGuildJoinRequests

```go
func (a *API) ListUserGuildJoinRequests(ctx context.Context, guildID snowflake.ID, userID snowflake.ID, opts ...Option) ([]discord.GuildJoinRequest, error)
```

#### func (*API) ListUserGuildScheduledEvents

```go
func (a *API) ListUserGuildScheduledEvents(ctx context.Context, query discord.ListUserGuildScheduledEventsQuery, opts ...Option) ([]discord.GuildScheduledEventUser, error)
```

#### func (*API) ListUserGuilds

```go
func (a *API) ListUserGuilds(ctx context.Context, query discord.ListUserGuildsQuery, opts ...Option) ([]discord.UserGuild, error)
```

#### func (*API) ListUserInvites

```go
func (a *API) ListUserInvites(ctx context.Context, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ListUserLinkedConnections

```go
func (a *API) ListUserLinkedConnections(ctx context.Context, opts ...Option) ([]discord.Connection, error)
```

#### func (*API) ListUserMessageSummaries

```go
func (a *API) ListUserMessageSummaries(ctx context.Context, opts ...Option) ([]discord.UserMessageSummary, error)
```

#### func (*API) ListUserPurchasedCollectibles

```go
func (a *API) ListUserPurchasedCollectibles(ctx context.Context, query discord.ListUserPurchasedCollectiblesQuery, opts ...Option) ([]discord.CollectibleProduct, error)
```

#### func (*API) ListVideoFilterAssets

```go
func (a *API) ListVideoFilterAssets(ctx context.Context, opts ...Option) ([]discord.VideoFilterAsset, error)
```

#### func (*API) ListVoiceRegions

```go
func (a *API) ListVoiceRegions(ctx context.Context, opts ...Option) ([]discord.VoiceRegion, error)
```

#### func (*API) ListWebauthnAuthenticators

```go
func (a *API) ListWebauthnAuthenticators(ctx context.Context, opts ...Option) (json.RawMessage, error)
```

#### func (*API) LoginAccount

```go
func (a *API) LoginAccount(ctx context.Context, params discord.LoginAccountParams, opts ...Option) (*discord.LoginAccountResponse, error)
```

#### func (*API) Logout

```go
func (a *API) Logout(ctx context.Context, params discord.LogoutParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) LogoutAuthSessions

```go
func (a *API) LogoutAuthSessions(ctx context.Context, params discord.LogoutAuthSessionsParams, opts ...Option) error
```

#### func (*API) MarkVideoFilterAssetLastUsed

```go
func (a *API) MarkVideoFilterAssetLastUsed(ctx context.Context, videoFilterAssetID snowflake.ID, opts ...Option) (json.RawMessage, error)
```

#### func (*API) MigrateBypassSlowmodePermission

```go
func (a *API) MigrateBypassSlowmodePermission(ctx context.Context, guildID snowflake.ID, opts ...Option) error
```

#### func (*API) MigrateGuildCommandScope

```go
func (a *API) MigrateGuildCommandScope(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.MigrateGuildCommandScopeResponse, error)
```

#### func (*API) MigratePinPermission

```go
func (a *API) MigratePinPermission(ctx context.Context, guildID snowflake.ID, opts ...Option) error
```

#### func (*API) ModifyApplication

```go
func (a *API) ModifyApplication(ctx context.Context, applicationID snowflake.ID, params discord.ModifyApplicationParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ModifyApplicationBot

```go
func (a *API) ModifyApplicationBot(ctx context.Context, applicationID snowflake.ID, params discord.ModifyApplicationBotParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ModifyApplicationEmoji

```go
func (a *API) ModifyApplicationEmoji(ctx context.Context, applicationID snowflake.ID, emojiID snowflake.ID, params discord.ModifyApplicationEmojiParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ModifyApplicationGameClaim

```go
func (a *API) ModifyApplicationGameClaim(ctx context.Context, applicationID snowflake.ID, gameClaimID snowflake.ID, params discord.ModifyApplicationGameClaimParams, opts ...Option) (*discord.GameClaim, error)
```

#### func (*API) ModifyApplicationProxyConfig

```go
func (a *API) ModifyApplicationProxyConfig(ctx context.Context, applicationID snowflake.ID, params discord.ModifyApplicationProxyConfigParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ModifyApplicationStorefront

```go
func (a *API) ModifyApplicationStorefront(ctx context.Context, applicationID snowflake.ID, params discord.ModifyApplicationStorefrontParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ModifyApplicationStorefrontPublishStates

```go
func (a *API) ModifyApplicationStorefrontPublishStates(ctx context.Context, applicationID snowflake.ID, params discord.ModifyApplicationStorefrontPublishStatesParams, opts ...Option) (*discord.ModifyApplicationStorefrontPublishStatesResponse, error)
```

#### func (*API) ModifyApplicationWidgetConfig

```go
func (a *API) ModifyApplicationWidgetConfig(ctx context.Context, applicationID snowflake.ID, widgetConfigID snowflake.ID, params discord.ModifyApplicationWidgetConfigParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ModifyAudioSettings

```go
func (a *API) ModifyAudioSettings(ctx context.Context, audioContextType string, userID snowflake.ID, params discord.ModifyAudioSettingsParams, opts ...Option) error
```

#### func (*API) ModifyAutomodIncidentActions

```go
func (a *API) ModifyAutomodIncidentActions(ctx context.Context, guildID snowflake.ID, params discord.ModifyAutomodIncidentActionsParams, opts ...Option) (*discord.ModifyAutomodIncidentActionsResponse, error)
```

#### func (*API) ModifyCall

```go
func (a *API) ModifyCall(ctx context.Context, channelID snowflake.ID, params discord.ModifyCallParams, opts ...Option) error
```

#### func (*API) ModifyChannel

```go
func (a *API) ModifyChannel(ctx context.Context, channelID snowflake.ID, params discord.ModifyChannelParams, opts ...Option) (*discord.Channel, error)
```

#### func (*API) ModifyChannelPermissions

```go
func (a *API) ModifyChannelPermissions(ctx context.Context, channelID snowflake.ID, overwriteID snowflake.ID, params discord.ModifyChannelPermissionsParams, opts ...Option) error
```

#### func (*API) ModifyChannelStatus

```go
func (a *API) ModifyChannelStatus(ctx context.Context, channelID snowflake.ID, params discord.ModifyChannelStatusParams, opts ...Option) error
```

#### func (*API) ModifyChannelTag

```go
func (a *API) ModifyChannelTag(ctx context.Context, channelID snowflake.ID, forumTagID snowflake.ID, params discord.ModifyChannelTagParams, opts ...Option) (*discord.Channel, error)
```

#### func (*API) ModifyCurrentApplication

```go
func (a *API) ModifyCurrentApplication(ctx context.Context, params discord.ModifyCurrentApplicationParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ModifyCurrentGuildMember

```go
func (a *API) ModifyCurrentGuildMember(ctx context.Context, guildID snowflake.ID, params discord.ModifyCurrentGuildMemberParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ModifyCurrentGuildMemberNick

```go
func (a *API) ModifyCurrentGuildMemberNick(ctx context.Context, guildID snowflake.ID, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ModifyCurrentUser

```go
func (a *API) ModifyCurrentUser(ctx context.Context, params discord.ModifyCurrentUserParams, opts ...Option) (*discord.User, error)
```

#### func (*API) ModifyCurrentUserAccount

```go
func (a *API) ModifyCurrentUserAccount(ctx context.Context, params discord.ModifyCurrentUserAccountParams, opts ...Option) (*discord.User, error)
```

#### func (*API) ModifyCurrentUserVoiceState

```go
func (a *API) ModifyCurrentUserVoiceState(ctx context.Context, guildID snowflake.ID, params discord.ModifyCurrentUserVoiceStateParams, opts ...Option) error
```

#### func (*API) ModifyDirectoryEntry

```go
func (a *API) ModifyDirectoryEntry(ctx context.Context, channelID snowflake.ID, entityID snowflake.ID, params discord.ModifyDirectoryEntryParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ModifyEmailSettings

```go
func (a *API) ModifyEmailSettings(ctx context.Context, params discord.ModifyEmailSettingsParams, opts ...Option) (*discord.EmailSettings, error)
```

#### func (*API) ModifyEmbeddedActivityConfig

```go
func (a *API) ModifyEmbeddedActivityConfig(ctx context.Context, applicationID snowflake.ID, params discord.ModifyEmbeddedActivityConfigParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ModifyFollowupMessage

```go
func (a *API) ModifyFollowupMessage(ctx context.Context, applicationID snowflake.ID, interactionToken string, messageID snowflake.ID, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ModifyGuild

```go
func (a *API) ModifyGuild(ctx context.Context, guildID snowflake.ID, params discord.ModifyGuildParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ModifyGuildChannelPositions

```go
func (a *API) ModifyGuildChannelPositions(ctx context.Context, guildID snowflake.ID, params discord.ModifyGuildChannelPositionsParams, opts ...Option) error
```

#### func (*API) ModifyGuildDiscoveryMetadata

```go
func (a *API) ModifyGuildDiscoveryMetadata(ctx context.Context, guildID snowflake.ID, params discord.ModifyGuildDiscoveryMetadataParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ModifyGuildEmoji

```go
func (a *API) ModifyGuildEmoji(ctx context.Context, guildID snowflake.ID, emojiID snowflake.ID, params discord.ModifyGuildEmojiParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ModifyGuildIntegration

```go
func (a *API) ModifyGuildIntegration(ctx context.Context, guildID snowflake.ID, integrationID snowflake.ID, params discord.ModifyGuildIntegrationParams, opts ...Option) error
```

#### func (*API) ModifyGuildMFALevel

```go
func (a *API) ModifyGuildMFALevel(ctx context.Context, guildID snowflake.ID, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ModifyGuildMember

```go
func (a *API) ModifyGuildMember(ctx context.Context, guildID snowflake.ID, userID snowflake.ID, params discord.ModifyGuildMemberParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ModifyGuildMemberProfile

```go
func (a *API) ModifyGuildMemberProfile(ctx context.Context, guildID snowflake.ID, params discord.ModifyGuildMemberProfileParams, opts ...Option) (*discord.ProfileMetadata, error)
```

#### func (*API) ModifyGuildMemberVerification

```go
func (a *API) ModifyGuildMemberVerification(ctx context.Context, guildID snowflake.ID, params discord.ModifyGuildMemberVerificationParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ModifyGuildNewMemberAction

```go
func (a *API) ModifyGuildNewMemberAction(ctx context.Context, guildID snowflake.ID, channelID snowflake.ID, params discord.ModifyGuildNewMemberActionParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ModifyGuildNewMemberWelcome

```go
func (a *API) ModifyGuildNewMemberWelcome(ctx context.Context, guildID snowflake.ID, params discord.ModifyGuildNewMemberWelcomeParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ModifyGuildOnboarding

```go
func (a *API) ModifyGuildOnboarding(ctx context.Context, guildID snowflake.ID, params discord.ModifyGuildOnboardingParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ModifyGuildOnboardingResponses

```go
func (a *API) ModifyGuildOnboardingResponses(ctx context.Context, guildID snowflake.ID, params discord.ModifyGuildOnboardingResponsesParams, opts ...Option) (*discord.OnboardingResponses, error)
```

#### func (*API) ModifyGuildPowerup

```go
func (a *API) ModifyGuildPowerup(ctx context.Context, guildID snowflake.ID, entitlementID snowflake.ID, params discord.ModifyGuildPowerupParams, opts ...Option) error
```

#### func (*API) ModifyGuildProductListing

```go
func (a *API) ModifyGuildProductListing(ctx context.Context, guildID snowflake.ID, productListingID snowflake.ID, params discord.ModifyGuildProductListingParams, opts ...Option) (*discord.GuildProductListing, error)
```

#### func (*API) ModifyGuildProfile

```go
func (a *API) ModifyGuildProfile(ctx context.Context, guildID snowflake.ID, params discord.ModifyGuildProfileParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ModifyGuildResourceChannel

```go
func (a *API) ModifyGuildResourceChannel(ctx context.Context, guildID snowflake.ID, channelID snowflake.ID, params discord.ModifyGuildResourceChannelParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ModifyGuildRole

```go
func (a *API) ModifyGuildRole(ctx context.Context, guildID snowflake.ID, roleID snowflake.ID, params discord.ModifyGuildRoleParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ModifyGuildRoleConnectionConfiguration

```go
func (a *API) ModifyGuildRoleConnectionConfiguration(ctx context.Context, guildID snowflake.ID, roleID snowflake.ID, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ModifyGuildRolePositions

```go
func (a *API) ModifyGuildRolePositions(ctx context.Context, guildID snowflake.ID, params discord.ModifyGuildRolePositionsParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ModifyGuildRoleSubscriptionGroupListing

```go
func (a *API) ModifyGuildRoleSubscriptionGroupListing(ctx context.Context, guildID snowflake.ID, subscriptionGroupListingID snowflake.ID, params discord.ModifyGuildRoleSubscriptionGroupListingParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ModifyGuildRoleSubscriptionsSettings

```go
func (a *API) ModifyGuildRoleSubscriptionsSettings(ctx context.Context, guildID snowflake.ID, params discord.ModifyGuildRoleSubscriptionsSettingsParams, opts ...Option) (*discord.RoleSubscriptionSettings, error)
```

#### func (*API) ModifyGuildSoundboardSound

```go
func (a *API) ModifyGuildSoundboardSound(ctx context.Context, guildID snowflake.ID, soundID snowflake.ID, params discord.ModifyGuildSoundboardSoundParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ModifyGuildSticker

```go
func (a *API) ModifyGuildSticker(ctx context.Context, guildID snowflake.ID, stickerID snowflake.ID, params discord.ModifyGuildStickerParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ModifyGuildTemplate

```go
func (a *API) ModifyGuildTemplate(ctx context.Context, guildID snowflake.ID, templateCode string, params discord.ModifyGuildTemplateParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ModifyGuildVanityInvite

```go
func (a *API) ModifyGuildVanityInvite(ctx context.Context, guildID snowflake.ID, params discord.ModifyGuildVanityInviteParams, opts ...Option) (*discord.ModifyGuildVanityInviteResponse, error)
```

#### func (*API) ModifyGuildWelcomeScreen

```go
func (a *API) ModifyGuildWelcomeScreen(ctx context.Context, guildID snowflake.ID, params discord.ModifyGuildWelcomeScreenParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ModifyGuildWidget

```go
func (a *API) ModifyGuildWidget(ctx context.Context, guildID snowflake.ID, params discord.ModifyGuildWidgetParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ModifyLinkedUser

```go
func (a *API) ModifyLinkedUser(ctx context.Context, params discord.ModifyLinkedUserParams, opts ...Option) ([]discord.LinkedUser, error)
```

#### func (*API) ModifyLobby

```go
func (a *API) ModifyLobby(ctx context.Context, lobbyID snowflake.ID, params discord.ModifyLobbyParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ModifyLobbyLinkedChannel

```go
func (a *API) ModifyLobbyLinkedChannel(ctx context.Context, lobbyID snowflake.ID, params discord.ModifyLobbyLinkedChannelParams, opts ...Option) (*discord.Lobby, error)
```

#### func (*API) ModifyNotificationSettings

```go
func (a *API) ModifyNotificationSettings(ctx context.Context, params discord.ModifyNotificationSettingsParams, opts ...Option) (*discord.NotificationSettings, error)
```

#### func (*API) ModifyOriginalInteractionResponse

```go
func (a *API) ModifyOriginalInteractionResponse(ctx context.Context, applicationID snowflake.ID, interactionToken string, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ModifyPaymentSource

```go
func (a *API) ModifyPaymentSource(ctx context.Context, paymentSourceID snowflake.ID, params discord.ModifyPaymentSourceParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ModifyProfileWidgets

```go
func (a *API) ModifyProfileWidgets(ctx context.Context, params discord.ModifyProfileWidgetsParams, opts ...Option) ([]discord.GameWidget, error)
```

#### func (*API) ModifyRelationship

```go
func (a *API) ModifyRelationship(ctx context.Context, userID snowflake.ID, params discord.ModifyRelationshipParams, opts ...Option) error
```

#### func (*API) ModifySKU

```go
func (a *API) ModifySKU(ctx context.Context, skuid snowflake.ID, params discord.ModifySKUParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ModifyStageInstance

```go
func (a *API) ModifyStageInstance(ctx context.Context, channelID snowflake.ID, params discord.ModifyStageInstanceParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ModifyStream

```go
func (a *API) ModifyStream(ctx context.Context, streamKey string, params discord.ModifyStreamParams, opts ...Option) error
```

#### func (*API) ModifySubscription

```go
func (a *API) ModifySubscription(ctx context.Context, subscriptionID snowflake.ID, query discord.ModifySubscriptionQuery, params discord.ModifySubscriptionParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ModifySubscriptionPreview

```go
func (a *API) ModifySubscriptionPreview(ctx context.Context, subscriptionID snowflake.ID, params discord.ModifySubscriptionPreviewParams, opts ...Option) (*discord.Invoice, error)
```

#### func (*API) ModifyTeam

```go
func (a *API) ModifyTeam(ctx context.Context, teamID snowflake.ID, params discord.ModifyTeamParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ModifyTeamMember

```go
func (a *API) ModifyTeamMember(ctx context.Context, teamID snowflake.ID, userID snowflake.ID, params discord.ModifyTeamMemberParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ModifyThreadSettings

```go
func (a *API) ModifyThreadSettings(ctx context.Context, channelID snowflake.ID, params discord.ModifyThreadSettingsParams, opts ...Option) (*discord.ThreadMember, error)
```

#### func (*API) ModifyUserAgreements

```go
func (a *API) ModifyUserAgreements(ctx context.Context, params discord.ModifyUserAgreementsParams, opts ...Option) error
```

#### func (*API) ModifyUserApplicationProfileConfig

```go
func (a *API) ModifyUserApplicationProfileConfig(ctx context.Context, applicationID snowflake.ID, externalUserID snowflake.ID, params discord.ModifyUserApplicationProfileConfigParams, opts ...Option) error
```

#### func (*API) ModifyUserConnection

```go
func (a *API) ModifyUserConnection(ctx context.Context, connectionType string, connectionID snowflake.ID, params discord.ModifyUserConnectionParams, opts ...Option) (*discord.Connection, error)
```

#### func (*API) ModifyUserConsents

```go
func (a *API) ModifyUserConsents(ctx context.Context, params discord.ModifyUserConsentsParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ModifyUserGuildSettings

```go
func (a *API) ModifyUserGuildSettings(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.UserGuildSettings, error)
```

#### func (*API) ModifyUserNote

```go
func (a *API) ModifyUserNote(ctx context.Context, userID snowflake.ID, params discord.ModifyUserNoteParams, opts ...Option) error
```

#### func (*API) ModifyUserProfile

```go
func (a *API) ModifyUserProfile(ctx context.Context, params discord.ModifyUserProfileParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ModifyUserSettings

```go
func (a *API) ModifyUserSettings(ctx context.Context, params discord.ModifyUserSettingsParams, opts ...Option) (*discord.UserSettings, error)
```

#### func (*API) ModifyUserSettingsProto

```go
func (a *API) ModifyUserSettingsProto(ctx context.Context, userSettingsProtoType string, params discord.ModifyUserSettingsProtoParams, opts ...Option) (*discord.ModifyUserSettingsProtoResponse, error)
```

#### func (*API) ModifyUserVoiceState

```go
func (a *API) ModifyUserVoiceState(ctx context.Context, guildID snowflake.ID, userID snowflake.ID, params discord.ModifyUserVoiceStateParams, opts ...Option) error
```

#### func (*API) ModifyWebauthnAuthenticator

```go
func (a *API) ModifyWebauthnAuthenticator(ctx context.Context, authenticatorID snowflake.ID, params discord.ModifyWebauthnAuthenticatorParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ModifyWebhook

```go
func (a *API) ModifyWebhook(ctx context.Context, webhookID snowflake.ID, params discord.ModifyWebhookParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ModifyWebhookWithToken

```go
func (a *API) ModifyWebhookWithToken(ctx context.Context, webhookID snowflake.ID, webhookToken string, params discord.ModifyWebhookWithTokenParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) OneTimeLogin

```go
func (a *API) OneTimeLogin(ctx context.Context, params discord.OneTimeLoginParams, opts ...Option) (*discord.OneTimeLoginResponse, error)
```

#### func (*API) ParaphraseThreadTitle

```go
func (a *API) ParaphraseThreadTitle(ctx context.Context, params discord.ParaphraseThreadTitleParams, opts ...Option) (*discord.ParaphraseThreadTitleResponse, error)
```

#### func (*API) PatchExample

```go
func (a *API) PatchExample(ctx context.Context, opts ...Option) (json.RawMessage, error)
```

#### func (*API) PinMessage

```go
func (a *API) PinMessage(ctx context.Context, channelID snowflake.ID, messageID snowflake.ID, opts ...Option) error
```

#### func (*API) PostExample

```go
func (a *API) PostExample(ctx context.Context, opts ...Option) (json.RawMessage, error)
```

#### func (*API) PreloadMessages

```go
func (a *API) PreloadMessages(ctx context.Context, opts ...Option) ([]discord.Message, error)
```

#### func (*API) PreviewOAuth2Authorization

```go
func (a *API) PreviewOAuth2Authorization(ctx context.Context, query discord.PreviewOAuth2AuthorizationQuery, opts ...Option) (*discord.PreviewOAuth2AuthorizationResponse, error)
```

#### func (*API) ProxyApplicationAssets

```go
func (a *API) ProxyApplicationAssets(ctx context.Context, applicationID snowflake.ID, params discord.ProxyApplicationAssetsParams, opts ...Option) ([]discord.ExternalAsset, error)
```

#### func (*API) PruneGuild

```go
func (a *API) PruneGuild(ctx context.Context, guildID snowflake.ID, params discord.PruneGuildParams, opts ...Option) (*discord.PruneGuildResponse, error)
```

#### func (*API) PublishApplicationWidgetConfig

```go
func (a *API) PublishApplicationWidgetConfig(ctx context.Context, applicationID snowflake.ID, widgetConfigID snowflake.ID, opts ...Option) (*discord.WidgetConfig, error)
```

#### func (*API) PutExample

```go
func (a *API) PutExample(ctx context.Context, opts ...Option) (json.RawMessage, error)
```

#### func (*API) QueryApplicationTestMode

```go
func (a *API) QueryApplicationTestMode(ctx context.Context, applicationID snowflake.ID, opts ...Option) error
```

#### func (*API) QueryGuildMembers

```go
func (a *API) QueryGuildMembers(ctx context.Context, guildID snowflake.ID, query discord.QueryGuildMembersQuery, opts ...Option) ([]discord.GuildMember, error)
```

#### func (*API) QueryStudentHubs

```go
func (a *API) QueryStudentHubs(ctx context.Context, params discord.QueryStudentHubsParams, opts ...Option) (*discord.QueryStudentHubsResponse, error)
```

#### func (*API) QueryUnauthenticatedReportEligibility

```go
func (a *API) QueryUnauthenticatedReportEligibility(ctx context.Context, opts ...Option) (json.RawMessage, error)
```

#### func (*API) RedeemGiftCode

```go
func (a *API) RedeemGiftCode(ctx context.Context, giftCodeCode string, params discord.RedeemGiftCodeParams, opts ...Option) (*discord.Entitlement, error)
```

#### func (*API) RedeemUserOffer

```go
func (a *API) RedeemUserOffer(ctx context.Context, params discord.RedeemUserOfferParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) RedeemVirtualCurrency

```go
func (a *API) RedeemVirtualCurrency(ctx context.Context, skuid snowflake.ID, params discord.RedeemVirtualCurrencyParams, opts ...Option) ([]discord.Entitlement, error)
```

#### func (*API) RefreshAttachmentURLs

```go
func (a *API) RefreshAttachmentURLs(ctx context.Context, params discord.RefreshAttachmentURLsParams, opts ...Option) (*discord.RefreshAttachmentURLsResponse, error)
```

#### func (*API) RefreshUserConnection

```go
func (a *API) RefreshUserConnection(ctx context.Context, connectionType string, connectionID snowflake.ID, opts ...Option) error
```

#### func (*API) RegisterAccount

```go
func (a *API) RegisterAccount(ctx context.Context, params discord.RegisterAccountParams, opts ...Option) (*discord.RegisterAccountResponse, error)
```

#### func (*API) RegisterAccountWithPhoneNumber

```go
func (a *API) RegisterAccountWithPhoneNumber(ctx context.Context, params discord.RegisterAccountWithPhoneNumberParams, opts ...Option) error
```

#### func (*API) RegisterDevice

```go
func (a *API) RegisterDevice(ctx context.Context, params discord.RegisterDeviceParams, opts ...Option) error
```

#### func (*API) RejectMessageRequest

```go
func (a *API) RejectMessageRequest(ctx context.Context, channelID snowflake.ID, opts ...Option) (json.RawMessage, error)
```

#### func (*API) RemoveApplicationTester

```go
func (a *API) RemoveApplicationTester(ctx context.Context, applicationID snowflake.ID, userID snowflake.ID, opts ...Option) error
```

#### func (*API) RemoveChannelRecipient

```go
func (a *API) RemoveChannelRecipient(ctx context.Context, channelID snowflake.ID, userID snowflake.ID, opts ...Option) error
```

#### func (*API) RemoveCreatorMonetization

```go
func (a *API) RemoveCreatorMonetization(ctx context.Context, guildID snowflake.ID, opts ...Option) error
```

#### func (*API) RemoveFriendSuggestion

```go
func (a *API) RemoveFriendSuggestion(ctx context.Context, userID snowflake.ID, opts ...Option) error
```

#### func (*API) RemoveGameRelationship

```go
func (a *API) RemoveGameRelationship(ctx context.Context, userID snowflake.ID, opts ...Option) error
```

#### func (*API) RemoveGameRelationshipByApplication

```go
func (a *API) RemoveGameRelationshipByApplication(ctx context.Context, userID snowflake.ID, applicationID snowflake.ID, opts ...Option) error
```

#### func (*API) RemoveGuildDiscoverySubcategory

```go
func (a *API) RemoveGuildDiscoverySubcategory(ctx context.Context, guildID snowflake.ID, discoveryCategoryID snowflake.ID, opts ...Option) error
```

#### func (*API) RemoveGuildMember

```go
func (a *API) RemoveGuildMember(ctx context.Context, guildID snowflake.ID, userID snowflake.ID, opts ...Option) error
```

#### func (*API) RemoveGuildMemberRole

```go
func (a *API) RemoveGuildMemberRole(ctx context.Context, guildID snowflake.ID, userID snowflake.ID, roleID snowflake.ID, opts ...Option) error
```

#### func (*API) RemoveGuildPowerup

```go
func (a *API) RemoveGuildPowerup(ctx context.Context, guildID snowflake.ID, skuid snowflake.ID, query discord.RemoveGuildPowerupQuery, opts ...Option) error
```

#### func (*API) RemoveLinkedUser

```go
func (a *API) RemoveLinkedUser(ctx context.Context, params discord.RemoveLinkedUserParams, opts ...Option) ([]discord.LinkedUser, error)
```

#### func (*API) RemoveLobbyLink

```go
func (a *API) RemoveLobbyLink(ctx context.Context, channelID snowflake.ID, opts ...Option) (json.RawMessage, error)
```

#### func (*API) RemoveLobbyMember

```go
func (a *API) RemoveLobbyMember(ctx context.Context, lobbyID snowflake.ID, userID snowflake.ID, opts ...Option) error
```

#### func (*API) RemovePhoneNumber

```go
func (a *API) RemovePhoneNumber(ctx context.Context, params discord.RemovePhoneNumberParams, opts ...Option) error
```

#### func (*API) RemoveRelationship

```go
func (a *API) RemoveRelationship(ctx context.Context, userID snowflake.ID, opts ...Option) error
```

#### func (*API) RemoveTeamMember

```go
func (a *API) RemoveTeamMember(ctx context.Context, teamID snowflake.ID, userID snowflake.ID, opts ...Option) error
```

#### func (*API) RemoveThreadMember

```go
func (a *API) RemoveThreadMember(ctx context.Context, channelID snowflake.ID, userID snowflake.ID, query discord.RemoveThreadMemberQuery, opts ...Option) error
```

#### func (*API) ReportAutomodIncident

```go
func (a *API) ReportAutomodIncident(ctx context.Context, guildID snowflake.ID, opts ...Option) error
```

#### func (*API) ReportExplicitContentFalsePositive

```go
func (a *API) ReportExplicitContentFalsePositive(ctx context.Context, params discord.ReportExplicitContentFalsePositiveParams, opts ...Option) error
```

#### func (*API) ReportMeaningfullyOnline

```go
func (a *API) ReportMeaningfullyOnline(ctx context.Context, opts ...Option) error
```

#### func (*API) ReportSafetyWarningFalsePositive

```go
func (a *API) ReportSafetyWarningFalsePositive(ctx context.Context, channelID snowflake.ID, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ReportSentExplicitContentFalsePositive

```go
func (a *API) ReportSentExplicitContentFalsePositive(ctx context.Context, params discord.ReportSentExplicitContentFalsePositiveParams, opts ...Option) error
```

#### func (*API) ReportUnverifiedApplication

```go
func (a *API) ReportUnverifiedApplication(ctx context.Context, params discord.ReportUnverifiedApplicationParams, opts ...Option) (*discord.ReportUnverifiedApplicationResponse, error)
```

#### func (*API) RequestAgeVerificationForSuspendedUser

```go
func (a *API) RequestAgeVerificationForSuspendedUser(ctx context.Context, params discord.RequestAgeVerificationForSuspendedUserParams, opts ...Option) (*discord.RequestAgeVerificationForSuspendedUserResponse, error)
```

#### func (*API) RequestApplicationGatewayIntents

```go
func (a *API) RequestApplicationGatewayIntents(ctx context.Context, applicationID snowflake.ID, params discord.RequestApplicationGatewayIntentsParams, opts ...Option) error
```

#### func (*API) RequestClassificationReview

```go
func (a *API) RequestClassificationReview(ctx context.Context, classificationID snowflake.ID, params discord.RequestClassificationReviewParams, opts ...Option) (*discord.RequestClassificationReviewResponse, error)
```

#### func (*API) RequestClassificationReviewForSuspendedUser

```go
func (a *API) RequestClassificationReviewForSuspendedUser(ctx context.Context, classificationID snowflake.ID, params discord.RequestClassificationReviewForSuspendedUserParams, opts ...Option) (*discord.RequestClassificationReviewForSuspendedUserResponse, error)
```

#### func (*API) RequestReportReview

```go
func (a *API) RequestReportReview(ctx context.Context, params discord.RequestReportReviewParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ResendPaymentVerificationEmail

```go
func (a *API) ResendPaymentVerificationEmail(ctx context.Context, params discord.ResendPaymentVerificationEmailParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) ResendVerificationCode

```go
func (a *API) ResendVerificationCode(ctx context.Context, params discord.ResendVerificationCodeParams, opts ...Option) error
```

#### func (*API) ResendVerificationEmail

```go
func (a *API) ResendVerificationEmail(ctx context.Context, opts ...Option) error
```

#### func (*API) ResetApplicationBotToken

```go
func (a *API) ResetApplicationBotToken(ctx context.Context, applicationID snowflake.ID, opts ...Option) (*discord.ResetApplicationBotTokenResponse, error)
```

#### func (*API) ResetApplicationSecret

```go
func (a *API) ResetApplicationSecret(ctx context.Context, applicationID snowflake.ID, opts ...Option) (*discord.ResetApplicationSecretResponse, error)
```

#### func (*API) ResetGuildJoinRequest

```go
func (a *API) ResetGuildJoinRequest(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.GuildJoinRequest, error)
```

#### func (*API) ResetPassword

```go
func (a *API) ResetPassword(ctx context.Context, params discord.ResetPasswordParams, opts ...Option) (*discord.ResetPasswordResponse, error)
```

#### func (*API) ResetQuest

```go
func (a *API) ResetQuest(ctx context.Context, questID snowflake.ID, opts ...Option) (*discord.QuestUserStatus, error)
```

#### func (*API) ResetQuestDismissibility

```go
func (a *API) ResetQuestDismissibility(ctx context.Context, questID snowflake.ID, opts ...Option) (*discord.QuestUserStatus, error)
```

#### func (*API) ResetRecentQuestCompletions

```go
func (a *API) ResetRecentQuestCompletions(ctx context.Context, opts ...Option) error
```

#### func (*API) ResolveAutomodIncident

```go
func (a *API) ResolveAutomodIncident(ctx context.Context, guildID snowflake.ID, params discord.ResolveAutomodIncidentParams, opts ...Option) error
```

#### func (*API) ReverifyPhoneNumber

```go
func (a *API) ReverifyPhoneNumber(ctx context.Context, params discord.ReverifyPhoneNumberParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) RevertAccount

```go
func (a *API) RevertAccount(ctx context.Context, params discord.RevertAccountParams, opts ...Option) (*discord.RevertAccountResponse, error)
```

#### func (*API) RevokeOAuth2Token

```go
func (a *API) RevokeOAuth2Token(ctx context.Context, params discord.RevokeOAuth2TokenParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) RevokeUserGiftCode

```go
func (a *API) RevokeUserGiftCode(ctx context.Context, giftCodeCode string, opts ...Option) error
```

#### func (*API) RevokeUserInvites

```go
func (a *API) RevokeUserInvites(ctx context.Context, opts ...Option) (json.RawMessage, error)
```

#### func (*API) RingChannelRecipients

```go
func (a *API) RingChannelRecipients(ctx context.Context, channelID snowflake.ID, params discord.RingChannelRecipientsParams, opts ...Option) error
```

#### func (*API) SaveMessage

```go
func (a *API) SaveMessage(ctx context.Context, channelID snowflake.ID, messageID snowflake.ID, params discord.SaveMessageParams, opts ...Option) (*discord.SavedMessage, error)
```

#### func (*API) ScanExplicitMedia

```go
func (a *API) ScanExplicitMedia(ctx context.Context, channelID snowflake.ID, params discord.ScanExplicitMediaParams, opts ...Option) error
```

#### func (*API) SearchApplicationsDirectory

```go
func (a *API) SearchApplicationsDirectory(ctx context.Context, query discord.SearchApplicationsDirectoryQuery, opts ...Option) (*discord.SearchApplicationsDirectoryResponse, error)
```

#### func (*API) SearchChannelMessages

```go
func (a *API) SearchChannelMessages(ctx context.Context, channelID snowflake.ID, opts ...Option) (json.RawMessage, error)
```

#### func (*API) SearchChannelMessagesByTab

```go
func (a *API) SearchChannelMessagesByTab(ctx context.Context, channelID snowflake.ID, opts ...Option) (json.RawMessage, error)
```

#### func (*API) SearchCollectibles

```go
func (a *API) SearchCollectibles(ctx context.Context, query discord.SearchCollectiblesQuery, opts ...Option) (*discord.SearchCollectiblesResponse, error)
```

#### func (*API) SearchCompanies

```go
func (a *API) SearchCompanies(ctx context.Context, query discord.SearchCompaniesQuery, opts ...Option) ([]discord.TeamCompany, error)
```

#### func (*API) SearchDirectoryEntries

```go
func (a *API) SearchDirectoryEntries(ctx context.Context, channelID snowflake.ID, opts ...Option) ([]discord.DirectoryEntry, error)
```

#### func (*API) SearchDiscoverableGuilds

```go
func (a *API) SearchDiscoverableGuilds(ctx context.Context, query discord.SearchDiscoverableGuildsQuery, opts ...Option) ([]discord.DiscoverableGuild, error)
```

#### func (*API) SearchGifs

```go
func (a *API) SearchGifs(ctx context.Context, opts ...Option) ([]discord.GIF, error)
```

#### func (*API) SearchGuildBans

```go
func (a *API) SearchGuildBans(ctx context.Context, guildID snowflake.ID, query discord.SearchGuildBansQuery, opts ...Option) ([]discord.Ban, error)
```

#### func (*API) SearchGuildMembers

```go
func (a *API) SearchGuildMembers(ctx context.Context, guildID snowflake.ID, params discord.SearchGuildMembersParams, opts ...Option) (*discord.SearchGuildMembersResponse, error)
```

#### func (*API) SearchGuildMessages

```go
func (a *API) SearchGuildMessages(ctx context.Context, guildID snowflake.ID, query discord.SearchGuildMessagesQuery, opts ...Option) (*discord.SearchGuildMessagesResponse, error)
```

#### func (*API) SearchGuildMessagesByTab

```go
func (a *API) SearchGuildMessagesByTab(ctx context.Context, guildID snowflake.ID, params discord.SearchGuildMessagesByTabParams, opts ...Option) (*discord.SearchGuildMessagesByTabResponse, error)
```

#### func (*API) SearchPublishedGuilds

```go
func (a *API) SearchPublishedGuilds(ctx context.Context, query discord.SearchPublishedGuildsQuery, opts ...Option) ([]discord.DiscoverableGuild, error)
```

#### func (*API) SearchThreads

```go
func (a *API) SearchThreads(ctx context.Context, channelID snowflake.ID, query discord.SearchThreadsQuery, opts ...Option) (*discord.SearchThreadsResponse, error)
```

#### func (*API) SearchUserMessagesByTab

```go
func (a *API) SearchUserMessagesByTab(ctx context.Context, opts ...Option) (json.RawMessage, error)
```

#### func (*API) SendBackupCodesChallenge

```go
func (a *API) SendBackupCodesChallenge(ctx context.Context, params discord.SendBackupCodesChallengeParams, opts ...Option) (*discord.SendBackupCodesChallengeResponse, error)
```

#### func (*API) SendConsoleCommand

```go
func (a *API) SendConsoleCommand(ctx context.Context, connectionType string, deviceID snowflake.ID, params discord.SendConsoleCommandParams, opts ...Option) (*discord.SendConsoleCommandResponse, error)
```

#### func (*API) SendCustomCallSound

```go
func (a *API) SendCustomCallSound(ctx context.Context, channelID snowflake.ID, params discord.SendCustomCallSoundParams, opts ...Option) error
```

#### func (*API) SendEmailChangeChallenge

```go
func (a *API) SendEmailChangeChallenge(ctx context.Context, opts ...Option) error
```

#### func (*API) SendFriendRequest

```go
func (a *API) SendFriendRequest(ctx context.Context, params discord.SendFriendRequestParams, opts ...Option) error
```

#### func (*API) SendGameFriendRequest

```go
func (a *API) SendGameFriendRequest(ctx context.Context, params discord.SendGameFriendRequestParams, opts ...Option) error
```

#### func (*API) SendMFASMS

```go
func (a *API) SendMFASMS(ctx context.Context, params discord.SendMFASMSParams, opts ...Option) (*discord.SendMFASMSResponse, error)
```

#### func (*API) SendQuestHeartbeat

```go
func (a *API) SendQuestHeartbeat(ctx context.Context, questID snowflake.ID, params discord.SendQuestHeartbeatParams, opts ...Option) (*discord.QuestUserStatus, error)
```

#### func (*API) SendQuestVideoProgress

```go
func (a *API) SendQuestVideoProgress(ctx context.Context, questID snowflake.ID, params discord.SendQuestVideoProgressParams, opts ...Option) (*discord.QuestUserStatus, error)
```

#### func (*API) SendSoundboardSound

```go
func (a *API) SendSoundboardSound(ctx context.Context, channelID snowflake.ID, params discord.SendSoundboardSoundParams, opts ...Option) error
```

#### func (*API) SendVoiceChannelEffect

```go
func (a *API) SendVoiceChannelEffect(ctx context.Context, channelID snowflake.ID, params discord.SendVoiceChannelEffectParams, opts ...Option) error
```

#### func (*API) SetApplicationEmbeddability

```go
func (a *API) SetApplicationEmbeddability(ctx context.Context, applicationID snowflake.ID, params discord.SetApplicationEmbeddabilityParams, opts ...Option) error
```

#### func (*API) SetGuildIdentity

```go
func (a *API) SetGuildIdentity(ctx context.Context, params discord.SetGuildIdentityParams, opts ...Option) (*discord.User, error)
```

#### func (*API) StageReport

```go
func (a *API) StageReport(ctx context.Context, channelID snowflake.ID, messageID snowflake.ID, opts ...Option) (*discord.StageReportResponse, error)
```

#### func (*API) StartConditionalLogin

```go
func (a *API) StartConditionalLogin(ctx context.Context, opts ...Option) (*discord.StartConditionalLoginResponse, error)
```

#### func (*API) StartConsoleQuest

```go
func (a *API) StartConsoleQuest(ctx context.Context, questID snowflake.ID, query discord.StartConsoleQuestQuery, opts ...Option) (*discord.StartConsoleQuestResponse, error)
```

#### func (*API) StartPasswordlessLogin

```go
func (a *API) StartPasswordlessLogin(ctx context.Context, opts ...Option) (*discord.StartPasswordlessLoginResponse, error)
```

#### func (*API) StopConsoleQuest

```go
func (a *API) StopConsoleQuest(ctx context.Context, questID snowflake.ID, opts ...Option) error
```

#### func (*API) StopRingingChannelRecipients

```go
func (a *API) StopRingingChannelRecipients(ctx context.Context, channelID snowflake.ID, params discord.StopRingingChannelRecipientsParams, opts ...Option) error
```

#### func (*API) SubmitDeveloperPortalCsatSurvey

```go
func (a *API) SubmitDeveloperPortalCsatSurvey(ctx context.Context, params discord.SubmitDeveloperPortalCsatSurveyParams, opts ...Option) error
```

#### func (*API) SubmitReportMenu

```go
func (a *API) SubmitReportMenu(ctx context.Context, typeArg string, params discord.SubmitReportMenuParams, opts ...Option) (*discord.SubmitReportMenuResponse, error)
```

#### func (*API) SubmitUnauthenticatedReportMenu

```go
func (a *API) SubmitUnauthenticatedReportMenu(ctx context.Context, typeArg string, params discord.SubmitUnauthenticatedReportMenuParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) SummarizeThread

```go
func (a *API) SummarizeThread(ctx context.Context, channelID snowflake.ID, params discord.SummarizeThreadParams, opts ...Option) error
```

#### func (*API) SuppressTutorial

```go
func (a *API) SuppressTutorial(ctx context.Context, opts ...Option) error
```

#### func (*API) SyncDevices

```go
func (a *API) SyncDevices(ctx context.Context, params discord.SyncDevicesParams, opts ...Option) (*discord.SyncDevicesResponse, error)
```

#### func (*API) SyncGuildIntegration

```go
func (a *API) SyncGuildIntegration(ctx context.Context, guildID snowflake.ID, integrationID snowflake.ID, opts ...Option) error
```

#### func (*API) SyncGuildTemplate

```go
func (a *API) SyncGuildTemplate(ctx context.Context, guildID snowflake.ID, templateCode string, opts ...Option) (json.RawMessage, error)
```

#### func (*API) TrackSelectedGIF

```go
func (a *API) TrackSelectedGIF(ctx context.Context, params discord.TrackSelectedGIFParams, opts ...Option) error
```

#### func (*API) TransferApplication

```go
func (a *API) TransferApplication(ctx context.Context, applicationID snowflake.ID, params discord.TransferApplicationParams, opts ...Option) (*discord.Application, error)
```

#### func (*API) TranslateMessage

```go
func (a *API) TranslateMessage(ctx context.Context, params discord.TranslateMessageParams, opts ...Option) (*discord.TranslateMessageResponse, error)
```

#### func (*API) TriggerTypingIndicator

```go
func (a *API) TriggerTypingIndicator(ctx context.Context, channelID snowflake.ID, opts ...Option) (*discord.TriggerTypingIndicatorResponse, error)
```

#### func (*API) UnassignGuildRoleConnection

```go
func (a *API) UnassignGuildRoleConnection(ctx context.Context, guildID snowflake.ID, roleID snowflake.ID, opts ...Option) error
```

#### func (*API) UnfurlEmbed

```go
func (a *API) UnfurlEmbed(ctx context.Context, params discord.UnfurlEmbedParams, opts ...Option) (*discord.UnfurlEmbedResponse, error)
```

#### func (*API) UnfurlEmbeds

```go
func (a *API) UnfurlEmbeds(ctx context.Context, params discord.UnfurlEmbedsParams, opts ...Option) (*discord.UnfurlEmbedsResponse, error)
```

#### func (*API) UnhideMessageFromGuildFeed

```go
func (a *API) UnhideMessageFromGuildFeed(ctx context.Context, channelID snowflake.ID, messageID snowflake.ID, opts ...Option) error
```

#### func (*API) UnignoreUser

```go
func (a *API) UnignoreUser(ctx context.Context, userID snowflake.ID, opts ...Option) error
```

#### func (*API) UnmergeProvisionalAccount

```go
func (a *API) UnmergeProvisionalAccount(ctx context.Context, params discord.UnmergeProvisionalAccountParams, opts ...Option) error
```

#### func (*API) UnmergeProvisionalAccountWithBot

```go
func (a *API) UnmergeProvisionalAccountWithBot(ctx context.Context, params discord.UnmergeProvisionalAccountWithBotParams, opts ...Option) error
```

#### func (*API) UnpinMessage

```go
func (a *API) UnpinMessage(ctx context.Context, channelID snowflake.ID, messageID snowflake.ID, opts ...Option) error
```

#### func (*API) UnpublishApplicationWidgetConfig

```go
func (a *API) UnpublishApplicationWidgetConfig(ctx context.Context, applicationID snowflake.ID, widgetConfigID snowflake.ID, opts ...Option) (*discord.WidgetConfig, error)
```

#### func (*API) UnregisterDevice

```go
func (a *API) UnregisterDevice(ctx context.Context, params discord.UnregisterDeviceParams, opts ...Option) error
```

#### func (*API) UnsaveMessage

```go
func (a *API) UnsaveMessage(ctx context.Context, channelID snowflake.ID, messageID snowflake.ID, opts ...Option) error
```

#### func (*API) UpdateActivitySession

```go
func (a *API) UpdateActivitySession(ctx context.Context, params discord.UpdateActivitySessionParams, opts ...Option) (*discord.UpdateActivitySessionResponse, error)
```

#### func (*API) UpdateActivitySubscriptions

```go
func (a *API) UpdateActivitySubscriptions(ctx context.Context, params discord.UpdateActivitySubscriptionsParams, opts ...Option) error
```

#### func (*API) UpdateApplicationAsset

```go
func (a *API) UpdateApplicationAsset(ctx context.Context, applicationID snowflake.ID, applicationAssetKey string, params discord.UpdateApplicationAssetParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) UpdateApplicationManagedLink

```go
func (a *API) UpdateApplicationManagedLink(ctx context.Context, applicationID snowflake.ID, activityLinkID snowflake.ID, opts ...Option) (*discord.ActivityLink, error)
```

#### func (*API) UpdateCreatorMonetizationTeam

```go
func (a *API) UpdateCreatorMonetizationTeam(ctx context.Context, guildID snowflake.ID, params discord.UpdateCreatorMonetizationTeamParams, opts ...Option) (*discord.UpdateCreatorMonetizationTeamResponse, error)
```

#### func (*API) UpdateExternalFriendListEntries

```go
func (a *API) UpdateExternalFriendListEntries(ctx context.Context, connectionID snowflake.ID, params discord.UpdateExternalFriendListEntriesParams, opts ...Option) (*discord.UpdateExternalFriendListEntriesResponse, error)
```

#### func (*API) UpdateInviteTargetUsers

```go
func (a *API) UpdateInviteTargetUsers(ctx context.Context, inviteCode string, params discord.UpdateInviteTargetUsersParams, opts ...Option) error
```

#### func (*API) UpdateLobbyMessageModerationMetadata

```go
func (a *API) UpdateLobbyMessageModerationMetadata(ctx context.Context, lobbyID snowflake.ID, messageID snowflake.ID, opts ...Option) error
```

#### func (*API) UpdateMessageRequest

```go
func (a *API) UpdateMessageRequest(ctx context.Context, channelID snowflake.ID, params discord.UpdateMessageRequestParams, opts ...Option) (json.RawMessage, error)
```

#### func (*API) UpdatePresence

```go
func (a *API) UpdatePresence(ctx context.Context, params discord.UpdatePresenceParams, opts ...Option) error
```

#### func (*API) UploadApplicationAsset

```go
func (a *API) UploadApplicationAsset(ctx context.Context, applicationID snowflake.ID, params discord.UploadApplicationAssetParams, opts ...Option) ([]discord.CloudAttachment, error)
```

#### func (*API) UploadStreamPreview

```go
func (a *API) UploadStreamPreview(ctx context.Context, streamKey string, params discord.UploadStreamPreviewParams, opts ...Option) error
```

#### func (*API) UploadUnverifiedApplicationIcon

```go
func (a *API) UploadUnverifiedApplicationIcon(ctx context.Context, params discord.UploadUnverifiedApplicationIconParams, opts ...Option) error
```

#### func (*API) UploadVideoStreamPreview

```go
func (a *API) UploadVideoStreamPreview(ctx context.Context, streamKey string, params discord.UploadVideoStreamPreviewParams, opts ...Option) error
```

#### func (*API) UploadVoicePublicKey

```go
func (a *API) UploadVoicePublicKey(ctx context.Context, opts ...Option) error
```

#### func (*API) UseGuildTemplate

```go
func (a *API) UseGuildTemplate(ctx context.Context, templateCode string, params discord.UseGuildTemplateParams, opts ...Option) (*discord.Guild, error)
```

#### func (*API) UserGetUniqueUsernameEligibility

```go
func (a *API) UserGetUniqueUsernameEligibility(ctx context.Context, params discord.GetUniqueUsernameEligibilityParams, opts ...Option) (*discord.GetUniqueUsernameEligibilityResponse, error)
```

#### func (*API) UserGetUniqueUsernameSuggestions

```go
func (a *API) UserGetUniqueUsernameSuggestions(ctx context.Context, opts ...Option) (*discord.GetUniqueUsernameSuggestionsResponse, error)
```

#### func (*API) ValidateBillingAddress

```go
func (a *API) ValidateBillingAddress(ctx context.Context, opts ...Option) (*discord.ValidateBillingAddressResponse, error)
```

#### func (*API) ValidateDiscoverySearchTerm

```go
func (a *API) ValidateDiscoverySearchTerm(ctx context.Context, query discord.ValidateDiscoverySearchTermQuery, opts ...Option) (*discord.ValidateDiscoverySearchTermResponse, error)
```

#### func (*API) ValidateGuildAutomodRule

```go
func (a *API) ValidateGuildAutomodRule(ctx context.Context, guildID snowflake.ID, params discord.AutoModerationValidateGuildAutomodRuleParams, opts ...Option) (*discord.ValidateGuildAutomodRuleResponse, error)
```

#### func (*API) VerifyAge

```go
func (a *API) VerifyAge(ctx context.Context, opts ...Option) (*discord.VerifyAgeResponse, error)
```

#### func (*API) VerifyApplication

```go
func (a *API) VerifyApplication(ctx context.Context, applicationID snowflake.ID, opts ...Option) error
```

#### func (*API) VerifyEmailAddress

```go
func (a *API) VerifyEmailAddress(ctx context.Context, params discord.VerifyEmailAddressParams, opts ...Option) (*discord.VerifyEmailAddressResponse, error)
```

#### func (*API) VerifyEmailChangeCode

```go
func (a *API) VerifyEmailChangeCode(ctx context.Context, params discord.VerifyEmailChangeCodeParams, opts ...Option) (*discord.VerifyEmailChangeCodeResponse, error)
```

#### func (*API) VerifyMFA

```go
func (a *API) VerifyMFA(ctx context.Context, params discord.VerifyMFAParams, opts ...Option) (*discord.VerifyMFAResponse, error)
```

#### func (*API) VerifyMFALogin

```go
func (a *API) VerifyMFALogin(ctx context.Context, authenticatorType string, params discord.VerifyMFALoginParams, opts ...Option) (*discord.VerifyMFALoginResponse, error)
```

#### func (*API) VerifyPhoneNumber

```go
func (a *API) VerifyPhoneNumber(ctx context.Context, params discord.VerifyPhoneNumberParams, opts ...Option) (*discord.VerifyPhoneNumberResponse, error)
```

#### func (*API) VerifyPurchaseRequest

```go
func (a *API) VerifyPurchaseRequest(ctx context.Context, params discord.VerifyPurchaseRequestParams, opts ...Option) error
```

#### func (*API) VerifyUnauthenticatedReport

```go
func (a *API) VerifyUnauthenticatedReport(ctx context.Context, typeArg string, params discord.VerifyUnauthenticatedReportParams, opts ...Option) (*discord.VerifyUnauthenticatedReportResponse, error)
```

#### func (*API) VerifyUserCaptcha

```go
func (a *API) VerifyUserCaptcha(ctx context.Context, params discord.VerifyUserCaptchaParams, opts ...Option) error
```

#### func (*API) VerifyVoicePublicKey

```go
func (a *API) VerifyVoicePublicKey(ctx context.Context, userID snowflake.ID, params discord.VerifyVoicePublicKeyParams, opts ...Option) (*discord.VerifyVoicePublicKeyResponse, error)
```

#### func (*API) VoidPayment

```go
func (a *API) VoidPayment(ctx context.Context, paymentID snowflake.ID, opts ...Option) error
```

#### func (*API) WakeGuildGameServer

```go
func (a *API) WakeGuildGameServer(ctx context.Context, guildID snowflake.ID, gameServerID snowflake.ID, opts ...Option) (*discord.GameServer, error)
```

### type Option

```go
type Option func(*rest.Request)
```

#### func WithContextProperties

```go
func WithContextProperties(props any) Option
```

#### func WithHeader

```go
func WithHeader(key, value string) Option
```

#### func WithReason

```go
func WithReason(reason string) Option
```

#### func WithoutAuth

```go
func WithoutAuth() Option
```

