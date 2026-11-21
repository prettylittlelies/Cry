package api

import (
	"context"
	"encoding/json"

	"github.com/prettylittlelies/cry/discord"
	"github.com/prettylittlelies/cry/rest"
	"github.com/prettylittlelies/cry/snowflake"
)

func (a *API) PostExample(ctx context.Context, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me")}
	return raw(ctx, a, req, opts)
}

func (a *API) PutExample(ctx context.Context, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("PUT", "/users/@me")}
	return raw(ctx, a, req, opts)
}

func (a *API) PatchExample(ctx context.Context, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/users/@me")}
	return raw(ctx, a, req, opts)
}

func (a *API) DeleteExample(ctx context.Context, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/users/@me")}
	return raw(ctx, a, req, opts)
}

func (a *API) GetAuthSessions(ctx context.Context, opts ...Option) (*discord.GetAuthSessionsResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/auth/sessions")}
	return call[discord.GetAuthSessionsResponse](ctx, a, req, opts)
}

func (a *API) LogoutAuthSessions(ctx context.Context, params discord.LogoutAuthSessionsParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/auth/sessions/logout")}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) LoginAccount(ctx context.Context, params discord.LoginAccountParams, opts ...Option) (*discord.LoginAccountResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/auth/login")}
	req.JSON = params
	return call[discord.LoginAccountResponse](ctx, a, req, opts)
}

func (a *API) VerifyMFALogin(ctx context.Context, authenticatorType string, params discord.VerifyMFALoginParams, opts ...Option) (*discord.VerifyMFALoginResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/auth/mfa/{authenticator_type}", rest.P("authenticator_type", authenticatorType))}
	req.JSON = params
	return call[discord.VerifyMFALoginResponse](ctx, a, req, opts)
}

func (a *API) StartPasswordlessLogin(ctx context.Context, opts ...Option) (*discord.StartPasswordlessLoginResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/auth/passwordless/start")}
	return call[discord.StartPasswordlessLoginResponse](ctx, a, req, opts)
}

func (a *API) StartConditionalLogin(ctx context.Context, opts ...Option) (*discord.StartConditionalLoginResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/auth/conditional/start")}
	return call[discord.StartConditionalLoginResponse](ctx, a, req, opts)
}

func (a *API) FinishPasswordlessLogin(ctx context.Context, params discord.FinishPasswordlessLoginParams, opts ...Option) (*discord.FinishPasswordlessLoginResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/auth/conditional/finish")}
	req.JSON = params
	return call[discord.FinishPasswordlessLoginResponse](ctx, a, req, opts)
}

func (a *API) AuthorizeIPAddress(ctx context.Context, params discord.AuthorizeIPAddressParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/auth/authorize-ip")}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) CreateHandoffToken(ctx context.Context, params discord.CreateHandoffTokenParams, opts ...Option) (*discord.CreateHandoffTokenResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/auth/handoff")}
	req.JSON = params
	return call[discord.CreateHandoffTokenResponse](ctx, a, req, opts)
}

func (a *API) ExchangeHandoffToken(ctx context.Context, params discord.ExchangeHandoffTokenParams, opts ...Option) (*discord.ExchangeHandoffTokenResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/auth/handoff/exchange")}
	req.JSON = params
	return call[discord.ExchangeHandoffTokenResponse](ctx, a, req, opts)
}

func (a *API) RegisterAccount(ctx context.Context, params discord.RegisterAccountParams, opts ...Option) (*discord.RegisterAccountResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/auth/register")}
	req.JSON = params
	return call[discord.RegisterAccountResponse](ctx, a, req, opts)
}

func (a *API) RegisterAccountWithPhoneNumber(ctx context.Context, params discord.RegisterAccountWithPhoneNumberParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/auth/register/phone")}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) GetLocationMetadata(ctx context.Context, opts ...Option) (*discord.GetLocationMetadataResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/auth/location-metadata")}
	return call[discord.GetLocationMetadataResponse](ctx, a, req, opts)
}

func (a *API) GetPasswordStrength(ctx context.Context, params discord.GetPasswordStrengthParams, opts ...Option) (*discord.GetPasswordStrengthResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/auth/password/validate")}
	req.JSON = params
	return call[discord.GetPasswordStrengthResponse](ctx, a, req, opts)
}

func (a *API) AuthenticationGetUniqueUsernameSuggestions(ctx context.Context, query discord.GetUniqueUsernameSuggestionsQuery, opts ...Option) (*discord.AuthenticationGetUniqueUsernameSuggestionsResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/unique-username/username-suggestions-unauthed")}
	req.Query = rest.Values(query)
	return call[discord.AuthenticationGetUniqueUsernameSuggestionsResponse](ctx, a, req, opts)
}

func (a *API) AuthenticationGetUniqueUsernameEligibility(ctx context.Context, params discord.AuthenticationGetUniqueUsernameEligibilityParams, opts ...Option) (*discord.AuthenticationGetUniqueUsernameEligibilityResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/unique-username/username-attempt-unauthed")}
	req.JSON = params
	return call[discord.AuthenticationGetUniqueUsernameEligibilityResponse](ctx, a, req, opts)
}

func (a *API) Logout(ctx context.Context, params discord.LogoutParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/auth/logout")}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) ForgotPassword(ctx context.Context, params discord.ForgotPasswordParams, opts ...Option) (*discord.ForgotPasswordResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/auth/forgot")}
	req.JSON = params
	return call[discord.ForgotPasswordResponse](ctx, a, req, opts)
}

func (a *API) ResetPassword(ctx context.Context, params discord.ResetPasswordParams, opts ...Option) (*discord.ResetPasswordResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/auth/reset")}
	req.JSON = params
	return call[discord.ResetPasswordResponse](ctx, a, req, opts)
}

func (a *API) OneTimeLogin(ctx context.Context, params discord.OneTimeLoginParams, opts ...Option) (*discord.OneTimeLoginResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/auth/one-time-login")}
	req.JSON = params
	return call[discord.OneTimeLoginResponse](ctx, a, req, opts)
}

func (a *API) RevertAccount(ctx context.Context, params discord.RevertAccountParams, opts ...Option) (*discord.RevertAccountResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/auth/revert")}
	req.JSON = params
	return call[discord.RevertAccountResponse](ctx, a, req, opts)
}

func (a *API) VerifyMFA(ctx context.Context, params discord.VerifyMFAParams, opts ...Option) (*discord.VerifyMFAResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/mfa/finish")}
	req.JSON = params
	return call[discord.VerifyMFAResponse](ctx, a, req, opts)
}

func (a *API) SendMFASMS(ctx context.Context, params discord.SendMFASMSParams, opts ...Option) (*discord.SendMFASMSResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/auth/mfa/sms/send")}
	req.JSON = params
	return call[discord.SendMFASMSResponse](ctx, a, req, opts)
}

func (a *API) GetGateway(ctx context.Context, opts ...Option) (*discord.GetGatewayResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/gateway")}
	return call[discord.GetGatewayResponse](ctx, a, req, opts)
}

func (a *API) GetGatewayBot(ctx context.Context, opts ...Option) (*discord.GetGatewayBotResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/gateway/bot")}
	return call[discord.GetGatewayBotResponse](ctx, a, req, opts)
}

func (a *API) GetApplicationCommandIndex(ctx context.Context, applicationID snowflake.ID, opts ...Option) (*discord.ApplicationCommandIndex, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/applications/{application_id}/application-command-index", rest.P("application_id", applicationID))}
	return call[discord.ApplicationCommandIndex](ctx, a, req, opts)
}

func (a *API) GetChannelApplicationCommandIndex(ctx context.Context, channelID snowflake.ID, opts ...Option) (*discord.ApplicationCommandIndex, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/channels/{channel_id}/application-command-index", rest.P("channel_id", channelID))}
	return call[discord.ApplicationCommandIndex](ctx, a, req, opts)
}

func (a *API) GetGuildApplicationCommandIndex(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.ApplicationCommandIndex, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/application-command-index", rest.P("guild_id", guildID))}
	return call[discord.ApplicationCommandIndex](ctx, a, req, opts)
}

func (a *API) GetUserApplicationCommandIndex(ctx context.Context, query discord.ApplicationCommandsGetUserApplicationCommandIndexQuery, params discord.GetUserApplicationCommandIndexParamsApplicationCommandsXxxx, opts ...Option) (*discord.ApplicationCommandIndex, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/application-command-index")}
	req.Query = rest.Values(query)
	req.JSON = params
	return call[discord.ApplicationCommandIndex](ctx, a, req, opts)
}

func (a *API) CreateInteractionResponse(ctx context.Context, interactionID snowflake.ID, interactionToken string, query discord.CreateInteractionResponseQuery, opts ...Option) (*discord.CreateInteractionResponseResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/interactions/{interaction_id}/{interaction_token}/callback", rest.P("interaction_id", interactionID), rest.P("interaction_token", interactionToken))}
	req.Query = rest.Values(query)
	return call[discord.CreateInteractionResponseResponse](ctx, a, req, opts)
}

func (a *API) GetOriginalInteractionResponse(ctx context.Context, applicationID snowflake.ID, interactionToken string, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/webhooks/{application_id}/{interaction_token}/messages/@original", rest.P("application_id", applicationID), rest.P("interaction_token", interactionToken))}
	return raw(ctx, a, req, opts)
}

func (a *API) ModifyOriginalInteractionResponse(ctx context.Context, applicationID snowflake.ID, interactionToken string, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/webhooks/{application_id}/{interaction_token}/messages/@original", rest.P("application_id", applicationID), rest.P("interaction_token", interactionToken))}
	return raw(ctx, a, req, opts)
}

func (a *API) DeleteOriginalInteractionResponse(ctx context.Context, applicationID snowflake.ID, interactionToken string, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/webhooks/{application_id}/{interaction_token}/messages/@original", rest.P("application_id", applicationID), rest.P("interaction_token", interactionToken))}
	return none(ctx, a, req, opts)
}

func (a *API) CreateFollowupMessage(ctx context.Context, applicationID snowflake.ID, interactionToken string, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/webhooks/{application_id}/{interaction_token}", rest.P("application_id", applicationID), rest.P("interaction_token", interactionToken))}
	return raw(ctx, a, req, opts)
}

func (a *API) GetFollowupMessage(ctx context.Context, applicationID snowflake.ID, interactionToken string, messageID snowflake.ID, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/webhooks/{application_id}/{interaction_token}/messages/{message_id}", rest.P("application_id", applicationID), rest.P("interaction_token", interactionToken), rest.P("message_id", messageID))}
	return raw(ctx, a, req, opts)
}

func (a *API) ModifyFollowupMessage(ctx context.Context, applicationID snowflake.ID, interactionToken string, messageID snowflake.ID, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/webhooks/{application_id}/{interaction_token}/messages/{message_id}", rest.P("application_id", applicationID), rest.P("interaction_token", interactionToken), rest.P("message_id", messageID))}
	return raw(ctx, a, req, opts)
}

func (a *API) DeleteFollowupMessage(ctx context.Context, applicationID snowflake.ID, interactionToken string, messageID snowflake.ID, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/webhooks/{application_id}/{interaction_token}/messages/{message_id}", rest.P("application_id", applicationID), rest.P("interaction_token", interactionToken), rest.P("message_id", messageID))}
	return raw(ctx, a, req, opts)
}

func (a *API) CreateInteraction(ctx context.Context, params discord.CreateInteractionParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/interactions")}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) ExchangeRemoteAuthTicket(ctx context.Context, params discord.ExchangeRemoteAuthTicketParams, opts ...Option) (*discord.ExchangeRemoteAuthTicketResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/remote-auth/login")}
	req.JSON = params
	return call[discord.ExchangeRemoteAuthTicketResponse](ctx, a, req, opts)
}

func (a *API) CreateRemoteAuthSession(ctx context.Context, params discord.CreateRemoteAuthSessionParams, opts ...Option) (*discord.CreateRemoteAuthSessionResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/remote-auth")}
	req.JSON = params
	return call[discord.CreateRemoteAuthSessionResponse](ctx, a, req, opts)
}

func (a *API) FinishRemoteAuth(ctx context.Context, params discord.FinishRemoteAuthParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/remote-auth/finish")}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) CancelRemoteAuth(ctx context.Context, params discord.CancelRemoteAuthParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/remote-auth/cancel")}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) TranslateMessage(ctx context.Context, params discord.TranslateMessageParams, opts ...Option) (*discord.TranslateMessageResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/ai/translate")}
	req.JSON = params
	return call[discord.TranslateMessageResponse](ctx, a, req, opts)
}

func (a *API) ParaphraseThreadTitle(ctx context.Context, params discord.ParaphraseThreadTitleParams, opts ...Option) (*discord.ParaphraseThreadTitleResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/ai/title")}
	req.JSON = params
	return call[discord.ParaphraseThreadTitleResponse](ctx, a, req, opts)
}

func (a *API) FixGrammar(ctx context.Context, params discord.FixGrammarParams, opts ...Option) (*discord.FixGrammarResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/ai/fix-grammar")}
	req.JSON = params
	return call[discord.FixGrammarResponse](ctx, a, req, opts)
}

func (a *API) SummarizeThread(ctx context.Context, channelID snowflake.ID, params discord.SummarizeThreadParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/ai/summarize-thread/{channel_id}", rest.P("channel_id", channelID))}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) ListApplicationDirectoryCategories(ctx context.Context, query discord.ListApplicationDirectoryCategoriesQuery, opts ...Option) ([]discord.ApplicationDirectoryCategory, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/application-directory-static/categories")}
	req.Query = rest.Values(query)
	return list[discord.ApplicationDirectoryCategory](ctx, a, req, opts)
}

func (a *API) ListApplicationDirectoryCollections(ctx context.Context, query discord.ListApplicationDirectoryCollectionsQuery, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/application-directory-static/collections")}
	req.Query = rest.Values(query)
	return raw(ctx, a, req, opts)
}

func (a *API) GetApplicationDirectoryApplication(ctx context.Context, applicationID snowflake.ID, query discord.GetApplicationDirectoryApplicationQuery, opts ...Option) (*discord.Application, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/application-directory-static/applications/{application_id}", rest.P("application_id", applicationID))}
	req.Query = rest.Values(query)
	return call[discord.Application](ctx, a, req, opts)
}

func (a *API) GetApplicationDirectoryApplicationEmbed(ctx context.Context, applicationID snowflake.ID, query discord.GetApplicationDirectoryApplicationEmbedQuery, opts ...Option) (*discord.Application, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/application-directory/applications/{application_id}/embed", rest.P("application_id", applicationID))}
	req.Query = rest.Values(query)
	return call[discord.Application](ctx, a, req, opts)
}

func (a *API) ListApplicationDirectorySimilarApplications(ctx context.Context, applicationID snowflake.ID, query discord.ListApplicationDirectorySimilarApplicationsQuery, opts ...Option) (*discord.ListApplicationDirectorySimilarApplicationsResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/application-directory-static/applications/{application_id}/similar", rest.P("application_id", applicationID))}
	req.Query = rest.Values(query)
	return call[discord.ListApplicationDirectorySimilarApplicationsResponse](ctx, a, req, opts)
}

func (a *API) SearchApplicationsDirectory(ctx context.Context, query discord.SearchApplicationsDirectoryQuery, opts ...Option) (*discord.SearchApplicationsDirectoryResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/application-directory-static/search")}
	req.Query = rest.Values(query)
	return call[discord.SearchApplicationsDirectoryResponse](ctx, a, req, opts)
}

func (a *API) ListApplications(ctx context.Context, query discord.ListApplicationsQuery, opts ...Option) ([]discord.Application, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/applications")}
	req.Query = rest.Values(query)
	return list[discord.Application](ctx, a, req, opts)
}

func (a *API) ListApplicationsWithAssets(ctx context.Context, query discord.ListApplicationsWithAssetsQuery, opts ...Option) (*discord.ListApplicationsWithAssetsResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/applications-with-assets")}
	req.Query = rest.Values(query)
	return call[discord.ListApplicationsWithAssetsResponse](ctx, a, req, opts)
}

func (a *API) CreateApplication(ctx context.Context, params discord.CreateApplicationParams, opts ...Option) (*discord.Application, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/applications")}
	req.JSON = params
	return call[discord.Application](ctx, a, req, opts)
}

func (a *API) GetApplication(ctx context.Context, applicationID snowflake.ID, opts ...Option) (*discord.Application, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/applications/{application_id}", rest.P("application_id", applicationID))}
	return call[discord.Application](ctx, a, req, opts)
}

func (a *API) GetCurrentApplication(ctx context.Context, opts ...Option) (*discord.Application, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/applications/@me")}
	return call[discord.Application](ctx, a, req, opts)
}

func (a *API) ModifyApplication(ctx context.Context, applicationID snowflake.ID, params discord.ModifyApplicationParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/applications/{application_id}", rest.P("application_id", applicationID))}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) ModifyCurrentApplication(ctx context.Context, params discord.ModifyCurrentApplicationParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/applications/@me")}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) DeleteApplication(ctx context.Context, applicationID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/applications/{application_id}/delete", rest.P("application_id", applicationID))}
	return none(ctx, a, req, opts)
}

func (a *API) TransferApplication(ctx context.Context, applicationID snowflake.ID, params discord.TransferApplicationParams, opts ...Option) (*discord.Application, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/applications/{application_id}/transfer", rest.P("application_id", applicationID))}
	req.JSON = params
	return call[discord.Application](ctx, a, req, opts)
}

func (a *API) ResetApplicationSecret(ctx context.Context, applicationID snowflake.ID, opts ...Option) (*discord.ResetApplicationSecretResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/applications/{application_id}/reset", rest.P("application_id", applicationID))}
	return call[discord.ResetApplicationSecretResponse](ctx, a, req, opts)
}

func (a *API) ListApplicationTesters(ctx context.Context, applicationID snowflake.ID, opts ...Option) ([]discord.WhitelistedUser, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/oauth2/applications/{application_id}/allowlist", rest.P("application_id", applicationID))}
	return list[discord.WhitelistedUser](ctx, a, req, opts)
}

func (a *API) AddApplicationTester(ctx context.Context, applicationID snowflake.ID, params discord.AddApplicationTesterParams, opts ...Option) (*discord.WhitelistedUser, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/oauth2/applications/{application_id}/allowlist", rest.P("application_id", applicationID))}
	req.JSON = params
	return call[discord.WhitelistedUser](ctx, a, req, opts)
}

func (a *API) AcceptApplicationTesterInvitation(ctx context.Context, query discord.AcceptApplicationTesterInvitationQuery, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/oauth2/allowlist/accept")}
	req.Query = rest.Values(query)
	return none(ctx, a, req, opts)
}

func (a *API) RemoveApplicationTester(ctx context.Context, applicationID snowflake.ID, userID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/oauth2/applications/{application_id}/allowlist/{user_id}", rest.P("application_id", applicationID), rest.P("user_id", userID))}
	return none(ctx, a, req, opts)
}

func (a *API) CreateApplicationBot(ctx context.Context, applicationID snowflake.ID, opts ...Option) (*discord.CreateApplicationBotResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/applications/{application_id}/bot", rest.P("application_id", applicationID))}
	return call[discord.CreateApplicationBotResponse](ctx, a, req, opts)
}

func (a *API) ModifyApplicationBot(ctx context.Context, applicationID snowflake.ID, params discord.ModifyApplicationBotParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/applications/{application_id}/bot", rest.P("application_id", applicationID))}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) ResetApplicationBotToken(ctx context.Context, applicationID snowflake.ID, opts ...Option) (*discord.ResetApplicationBotTokenResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/applications/{application_id}/bot/reset", rest.P("application_id", applicationID))}
	return call[discord.ResetApplicationBotTokenResponse](ctx, a, req, opts)
}

func (a *API) RequestApplicationGatewayIntents(ctx context.Context, applicationID snowflake.ID, params discord.RequestApplicationGatewayIntentsParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/applications/{application_id}/request-additional-intents", rest.P("application_id", applicationID))}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) GetApplicationDiscoverabilityState(ctx context.Context, applicationID snowflake.ID, opts ...Option) (*discord.GetApplicationDiscoverabilityStateResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/applications/{application_id}/discoverability-state", rest.P("application_id", applicationID))}
	return call[discord.GetApplicationDiscoverabilityStateResponse](ctx, a, req, opts)
}

func (a *API) QueryApplicationTestMode(ctx context.Context, applicationID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("GET", "/activities/{application_id}/test-mode", rest.P("application_id", applicationID))}
	return none(ctx, a, req, opts)
}

func (a *API) ListEmbeddedActivities(ctx context.Context, query discord.ListEmbeddedActivitiesQuery, opts ...Option) (*discord.ListEmbeddedActivitiesResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/activities/shelf")}
	req.Query = rest.Values(query)
	return call[discord.ListEmbeddedActivitiesResponse](ctx, a, req, opts)
}

func (a *API) SetApplicationEmbeddability(ctx context.Context, applicationID snowflake.ID, params discord.SetApplicationEmbeddabilityParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/applications/{application_id}/set-embedded", rest.P("application_id", applicationID))}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) GetEmbeddedActivityConfig(ctx context.Context, applicationID snowflake.ID, opts ...Option) (*discord.EmbeddedActivityConfig, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/applications/{application_id}/embedded-activity-config", rest.P("application_id", applicationID))}
	return call[discord.EmbeddedActivityConfig](ctx, a, req, opts)
}

func (a *API) ModifyEmbeddedActivityConfig(ctx context.Context, applicationID snowflake.ID, params discord.ModifyEmbeddedActivityConfigParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/applications/{application_id}/embedded-activity-config", rest.P("application_id", applicationID))}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) ListEmbeddedActivityInstances(ctx context.Context, applicationID snowflake.ID, channelID snowflake.ID, opts ...Option) (*discord.ListEmbeddedActivityInstancesResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/activities/{application_id}/instances/{channel_id}", rest.P("application_id", applicationID), rest.P("channel_id", channelID))}
	return call[discord.ListEmbeddedActivityInstancesResponse](ctx, a, req, opts)
}

func (a *API) LaunchEmbeddedActivity(ctx context.Context, channelID snowflake.ID, applicationID snowflake.ID, params discord.LaunchEmbeddedActivityParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/activities/{channel_id}/{application_id}", rest.P("channel_id", channelID), rest.P("application_id", applicationID))}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) GetApplicationProxyConfig(ctx context.Context, applicationID snowflake.ID, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/applications/{application_id}/proxy-config", rest.P("application_id", applicationID))}
	return raw(ctx, a, req, opts)
}

func (a *API) ModifyApplicationProxyConfig(ctx context.Context, applicationID snowflake.ID, params discord.ModifyApplicationProxyConfigParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/applications/{application_id}/proxy-config", rest.P("application_id", applicationID))}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) ListOAuth2ApplicationAssets(ctx context.Context, applicationID snowflake.ID, query discord.ListOAuth2ApplicationAssetsQuery, opts ...Option) ([]discord.ApplicationAsset, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/oauth2/applications/{application_id}/assets", rest.P("application_id", applicationID))}
	req.Query = rest.Values(query)
	return list[discord.ApplicationAsset](ctx, a, req, opts)
}

func (a *API) CreateOAuth2ApplicationAsset(ctx context.Context, applicationID snowflake.ID, params discord.CreateOAuth2ApplicationAssetParams, opts ...Option) (*discord.ApplicationAsset, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/oauth2/applications/{application_id}/assets", rest.P("application_id", applicationID))}
	req.JSON = params
	return call[discord.ApplicationAsset](ctx, a, req, opts)
}

func (a *API) DeleteOAuth2ApplicationAsset(ctx context.Context, applicationID snowflake.ID, oAuth2ApplicationAssetID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/oauth2/applications/{application_id}/assets/{oauth2_application_asset_id}", rest.P("application_id", applicationID), rest.P("oauth2_application_asset_id", oAuth2ApplicationAssetID))}
	return none(ctx, a, req, opts)
}

func (a *API) ProxyApplicationAssets(ctx context.Context, applicationID snowflake.ID, params discord.ProxyApplicationAssetsParams, opts ...Option) ([]discord.ExternalAsset, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/applications/{application_id}/external-assets", rest.P("application_id", applicationID))}
	req.JSON = params
	return list[discord.ExternalAsset](ctx, a, req, opts)
}

func (a *API) GetApplicationAssets(ctx context.Context, applicationID snowflake.ID, opts ...Option) ([]discord.ApplicationAsset, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/applications/{application_id}/assets", rest.P("application_id", applicationID))}
	return list[discord.ApplicationAsset](ctx, a, req, opts)
}

func (a *API) UploadApplicationAsset(ctx context.Context, applicationID snowflake.ID, params discord.UploadApplicationAssetParams, opts ...Option) ([]discord.CloudAttachment, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/applications/{application_id}/assets/upload", rest.P("application_id", applicationID))}
	req.JSON = params
	return list[discord.CloudAttachment](ctx, a, req, opts)
}

func (a *API) CreateApplicationAsset(ctx context.Context, applicationID snowflake.ID, params discord.CreateApplicationAssetParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/applications/{application_id}/assets", rest.P("application_id", applicationID))}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) UpdateApplicationAsset(ctx context.Context, applicationID snowflake.ID, applicationAssetKey string, params discord.UpdateApplicationAssetParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/applications/{application_id}/{application_asset_key}", rest.P("application_id", applicationID), rest.P("application_asset_key", applicationAssetKey))}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) DeleteApplicationAsset(ctx context.Context, applicationID snowflake.ID, applicationAssetKey string, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/applications/{application_id}/{application_asset_key}", rest.P("application_id", applicationID), rest.P("application_asset_key", applicationAssetKey))}
	return none(ctx, a, req, opts)
}

func (a *API) CreateApplicationAttachment(ctx context.Context, applicationID snowflake.ID, params discord.CreateApplicationAttachmentParams, opts ...Option) (*discord.CreateApplicationAttachmentResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/applications/{application_id}/attachment", rest.P("application_id", applicationID))}
	req.JSON = params
	return call[discord.CreateApplicationAttachmentResponse](ctx, a, req, opts)
}

func (a *API) GetPartialApplication(ctx context.Context, applicationID snowflake.ID, query discord.GetPartialApplicationQuery, opts ...Option) (*discord.Application, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/applications/{application_id}/public", rest.P("application_id", applicationID))}
	req.Query = rest.Values(query)
	return call[discord.Application](ctx, a, req, opts)
}

func (a *API) ListPartialApplications(ctx context.Context, query discord.ListPartialApplicationsQuery, opts ...Option) ([]discord.Application, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/applications/public")}
	req.Query = rest.Values(query)
	return list[discord.Application](ctx, a, req, opts)
}

func (a *API) GetRPCApplication(ctx context.Context, applicationID snowflake.ID, opts ...Option) (*discord.Application, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/applications/{application_id}/rpc", rest.P("application_id", applicationID))}
	return call[discord.Application](ctx, a, req, opts)
}

func (a *API) ListApplicationDisclosures(ctx context.Context, applicationID snowflake.ID, opts ...Option) (*discord.ListApplicationDisclosuresResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/applications/{application_id}/disclosures", rest.P("application_id", applicationID))}
	return call[discord.ListApplicationDisclosuresResponse](ctx, a, req, opts)
}

func (a *API) AcknowledgeApplicationDisclosures(ctx context.Context, applicationID snowflake.ID, params discord.AcknowledgeApplicationDisclosuresParams, opts ...Option) (*discord.AcknowledgeApplicationDisclosuresResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/applications/{application_id}/disclosures", rest.P("application_id", applicationID))}
	req.JSON = params
	return call[discord.AcknowledgeApplicationDisclosuresResponse](ctx, a, req, opts)
}

func (a *API) ListGuildApplications(ctx context.Context, guildID snowflake.ID, query discord.ListGuildApplicationsQuery, opts ...Option) ([]discord.Application, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/applications", rest.P("guild_id", guildID))}
	req.Query = rest.Values(query)
	return list[discord.Application](ctx, a, req, opts)
}

func (a *API) ReportUnverifiedApplication(ctx context.Context, params discord.ReportUnverifiedApplicationParams, opts ...Option) (*discord.ReportUnverifiedApplicationResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/unverified-applications")}
	req.JSON = params
	return call[discord.ReportUnverifiedApplicationResponse](ctx, a, req, opts)
}

func (a *API) UploadUnverifiedApplicationIcon(ctx context.Context, params discord.UploadUnverifiedApplicationIconParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/unverified-applications/icons")}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) ListUserApplicationRoleConnections(ctx context.Context, params discord.ListUserApplicationRoleConnectionsParams, opts ...Option) ([]discord.ApplicationRoleConnection, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/applications/role-connections")}
	req.JSON = params
	return list[discord.ApplicationRoleConnection](ctx, a, req, opts)
}

func (a *API) ListApplicationManagedLinks(ctx context.Context, applicationID snowflake.ID, opts ...Option) ([]discord.ActivityLink, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/applications/{application_id}/managed-links/", rest.P("application_id", applicationID))}
	return list[discord.ActivityLink](ctx, a, req, opts)
}

func (a *API) CreateApplicationManagedLink(ctx context.Context, applicationID snowflake.ID, params discord.CreateApplicationManagedLinkParams, opts ...Option) (*discord.ActivityLink, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/applications/{application_id}/managed-links/", rest.P("application_id", applicationID))}
	req.JSON = params
	return call[discord.ActivityLink](ctx, a, req, opts)
}

func (a *API) GetApplicationManagedLink(ctx context.Context, applicationID snowflake.ID, activityLinkID snowflake.ID, opts ...Option) (*discord.ActivityLink, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/applications/{application_id}/managed-links/{activity_link_id}", rest.P("application_id", applicationID), rest.P("activity_link_id", activityLinkID))}
	return call[discord.ActivityLink](ctx, a, req, opts)
}

func (a *API) UpdateApplicationManagedLink(ctx context.Context, applicationID snowflake.ID, activityLinkID snowflake.ID, opts ...Option) (*discord.ActivityLink, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/applications/{application_id}/managed-links/{activity_link_id}", rest.P("application_id", applicationID), rest.P("activity_link_id", activityLinkID))}
	return call[discord.ActivityLink](ctx, a, req, opts)
}

func (a *API) DeleteApplicationManagedLink(ctx context.Context, applicationID snowflake.ID, activityLinkID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/applications/{application_id}/managed-links/{activity_link_id}", rest.P("application_id", applicationID), rest.P("activity_link_id", activityLinkID))}
	return none(ctx, a, req, opts)
}

func (a *API) CreateApplicationQuickLink(ctx context.Context, applicationID snowflake.ID, params discord.CreateApplicationQuickLinkParams, opts ...Option) (*discord.ActivityLink, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/applications/{application_id}/quick-links/", rest.P("application_id", applicationID))}
	req.JSON = params
	return call[discord.ActivityLink](ctx, a, req, opts)
}

func (a *API) GetApplicationQuickLink(ctx context.Context, applicationID snowflake.ID, activityLinkID snowflake.ID, opts ...Option) (*discord.ActivityLink, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/applications/{application_id}/quick-links/{activity_link_id}", rest.P("application_id", applicationID), rest.P("activity_link_id", activityLinkID))}
	return call[discord.ActivityLink](ctx, a, req, opts)
}

func (a *API) GetApplicationVerificationEligibility(ctx context.Context, applicationID snowflake.ID, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/applications/{application_id}/verification", rest.P("application_id", applicationID))}
	return raw(ctx, a, req, opts)
}

func (a *API) VerifyApplication(ctx context.Context, applicationID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/applications/{application_id}/auto-verification", rest.P("application_id", applicationID))}
	return none(ctx, a, req, opts)
}

func (a *API) ListSocialSdkReleases(ctx context.Context, opts ...Option) (*discord.ListSocialSdkReleasesResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/social-sdk/releases")}
	return call[discord.ListSocialSdkReleasesResponse](ctx, a, req, opts)
}

func (a *API) GetSocialSdkRelease(ctx context.Context, sdkReleaseVersion string, opts ...Option) (*discord.SocialSdkRelease, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/social-sdk/releases/{sdk_release_version}", rest.P("sdk_release_version", sdkReleaseVersion))}
	return call[discord.SocialSdkRelease](ctx, a, req, opts)
}

func (a *API) EnableSocialSdk(ctx context.Context, applicationID snowflake.ID, params discord.EnableSocialSdkParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/applications/{application_id}/social-sdk/enable", rest.P("application_id", applicationID))}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) ListApplicationExternalIdentityProviderConfigurations(ctx context.Context, applicationID snowflake.ID, params discord.ListApplicationExternalIdentityProviderConfigurationsParams, opts ...Option) ([]discord.ApplicationExternalIdentityProviderConfiguration, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/applications/{application_id}/external-identity-provider-configurations", rest.P("application_id", applicationID))}
	req.JSON = params
	return list[discord.ApplicationExternalIdentityProviderConfiguration](ctx, a, req, opts)
}

func (a *API) GetApplicationUndeletableReason(ctx context.Context, applicationID snowflake.ID, opts ...Option) (*discord.GetApplicationUndeletableReasonResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/applications/{application_id}/can-delete", rest.P("application_id", applicationID))}
	return call[discord.GetApplicationUndeletableReasonResponse](ctx, a, req, opts)
}

func (a *API) ListGuildAuditLogEntries(ctx context.Context, guildID snowflake.ID, query discord.ListGuildAuditLogEntriesQuery, opts ...Option) (*discord.ListGuildAuditLogEntriesResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/audit-logs", rest.P("guild_id", guildID))}
	req.Query = rest.Values(query)
	return call[discord.ListGuildAuditLogEntriesResponse](ctx, a, req, opts)
}

func (a *API) ListGuildAutomodRules(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.AutomodRule, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/auto-moderation/rules", rest.P("guild_id", guildID))}
	return list[discord.AutomodRule](ctx, a, req, opts)
}

func (a *API) GetGuildAutomodRule(ctx context.Context, guildID snowflake.ID, automodRuleID snowflake.ID, opts ...Option) (*discord.AutomodRule, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/auto-moderation/rules/{automod_rule_id}", rest.P("guild_id", guildID), rest.P("automod_rule_id", automodRuleID))}
	return call[discord.AutomodRule](ctx, a, req, opts)
}

func (a *API) CreateGuildAutomodRule(ctx context.Context, guildID snowflake.ID, params discord.CreateGuildAutomodRuleParams, opts ...Option) (*discord.AutomodRule, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/guilds/{guild_id}/auto-moderation/rules", rest.P("guild_id", guildID))}
	req.JSON = params
	return call[discord.AutomodRule](ctx, a, req, opts)
}

func (a *API) ValidateGuildAutomodRule(ctx context.Context, guildID snowflake.ID, params discord.AutoModerationValidateGuildAutomodRuleParams, opts ...Option) (*discord.ValidateGuildAutomodRuleResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/guilds/{guild_id}/auto-moderation/rules/validate", rest.P("guild_id", guildID))}
	req.JSON = params
	return call[discord.ValidateGuildAutomodRuleResponse](ctx, a, req, opts)
}

func (a *API) ExecuteAutomodAlertAction(ctx context.Context, guildID snowflake.ID, params discord.ExecuteAutomodAlertActionParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/guilds/{guild_id}/auto-moderation/alert-action", rest.P("guild_id", guildID))}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) ModifyAutomodIncidentActions(ctx context.Context, guildID snowflake.ID, params discord.ModifyAutomodIncidentActionsParams, opts ...Option) (*discord.ModifyAutomodIncidentActionsResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("PUT", "/guilds/{guild_id}/incident-actions", rest.P("guild_id", guildID))}
	req.JSON = params
	return call[discord.ModifyAutomodIncidentActionsResponse](ctx, a, req, opts)
}

func (a *API) ResolveAutomodIncident(ctx context.Context, guildID snowflake.ID, params discord.ResolveAutomodIncidentParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/guilds/{guild_id}/auto-moderation/false-alarm", rest.P("guild_id", guildID))}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) ReportAutomodIncident(ctx context.Context, guildID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/guilds/{guild_id}/auto-moderation/report-raid", rest.P("guild_id", guildID))}
	return none(ctx, a, req, opts)
}

func (a *API) ClearMentionRaidIncident(ctx context.Context, guildID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/guilds/{guild_id}/auto-moderation/clear-mention-raid", rest.P("guild_id", guildID))}
	return none(ctx, a, req, opts)
}

func (a *API) GetBillingCountryCode(ctx context.Context, opts ...Option) (*discord.GetBillingCountryCodeResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/billing/country-code")}
	return call[discord.GetBillingCountryCodeResponse](ctx, a, req, opts)
}

func (a *API) GetBillingLocationInfo(ctx context.Context, opts ...Option) (*discord.GetBillingLocationInfoResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/billing/location-info")}
	return call[discord.GetBillingLocationInfoResponse](ctx, a, req, opts)
}

func (a *API) ListPaymentSources(ctx context.Context, opts ...Option) ([]discord.PaymentSource, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/billing/payment-sources")}
	return list[discord.PaymentSource](ctx, a, req, opts)
}

func (a *API) GetPaymentSource(ctx context.Context, paymentSourceID snowflake.ID, opts ...Option) (*discord.PaymentSource, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/billing/payment-sources/{payment_source_id}", rest.P("payment_source_id", paymentSourceID))}
	return call[discord.PaymentSource](ctx, a, req, opts)
}

func (a *API) GetPaymentSourceCreationContext(ctx context.Context, opts ...Option) (*discord.GetPaymentSourceCreationContextResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/billing/payment-source-creation-context")}
	return call[discord.GetPaymentSourceCreationContextResponse](ctx, a, req, opts)
}

func (a *API) CreatePaymentSource(ctx context.Context, query discord.CreatePaymentSourceQuery, params discord.CreatePaymentSourceParams, opts ...Option) (*discord.PaymentSource, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/billing/payment-sources")}
	req.Query = rest.Values(query)
	req.JSON = params
	return call[discord.PaymentSource](ctx, a, req, opts)
}

func (a *API) ModifyPaymentSource(ctx context.Context, paymentSourceID snowflake.ID, params discord.ModifyPaymentSourceParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/users/@me/billing/payment-sources/{payment_source_id}", rest.P("payment_source_id", paymentSourceID))}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) DeletePaymentSource(ctx context.Context, paymentSourceID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/users/@me/billing/payment-sources/{payment_source_id}", rest.P("payment_source_id", paymentSourceID))}
	return none(ctx, a, req, opts)
}

func (a *API) ValidateBillingAddress(ctx context.Context, opts ...Option) (*discord.ValidateBillingAddressResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/billing/payment-sources/validate-billing-address")}
	return call[discord.ValidateBillingAddressResponse](ctx, a, req, opts)
}

func (a *API) GetStripeSetupIntentSecret(ctx context.Context, params discord.GetStripeSetupIntentSecretParams, opts ...Option) (*discord.GetStripeSetupIntentSecretResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/billing/stripe/setup-intents")}
	req.JSON = params
	return call[discord.GetStripeSetupIntentSecretResponse](ctx, a, req, opts)
}

func (a *API) CreatePaypalBillingAgreementToken(ctx context.Context, params discord.CreatePaypalBillingAgreementTokenParams, opts ...Option) (*discord.CreatePaypalBillingAgreementTokenResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/billing/paypal/billing-agreement-tokens")}
	req.JSON = params
	return call[discord.CreatePaypalBillingAgreementTokenResponse](ctx, a, req, opts)
}

func (a *API) ListAvailableAdyenPaymentMethods(ctx context.Context, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/billing/adyen/payment-methods")}
	return raw(ctx, a, req, opts)
}

func (a *API) CreateBillingPopupBridge(ctx context.Context, paymentSourceType string, opts ...Option) (*discord.CreateBillingPopupBridgeResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/billing/popup-bridge/{payment_source_type}", rest.P("payment_source_type", paymentSourceType))}
	return call[discord.CreateBillingPopupBridgeResponse](ctx, a, req, opts)
}

func (a *API) CreateBillingPopupBridgeRedirect(ctx context.Context, paymentSourceType string, state string, responseType string, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/billing/popup-bridge/{payment_source_type}/callback/{state}/{response_type}", rest.P("payment_source_type", paymentSourceType), rest.P("state", state), rest.P("response_type", responseType))}
	return raw(ctx, a, req, opts)
}

func (a *API) CreateBillingPopupBridgeCallback(ctx context.Context, paymentSourceType string, params discord.CreateBillingPopupBridgeCallbackParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/billing/popup-bridge/{payment_source_type}/callback", rest.P("payment_source_type", paymentSourceType))}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) GetLocalizedPricingPromo(ctx context.Context, opts ...Option) (*discord.GetLocalizedPricingPromoResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/billing/localized-pricing-promo")}
	return call[discord.GetLocalizedPricingPromoResponse](ctx, a, req, opts)
}

func (a *API) GetUserOffer(ctx context.Context, params discord.GetUserOfferParams, opts ...Option) (*discord.GetUserOfferResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/billing/user-offer")}
	req.JSON = params
	return call[discord.GetUserOfferResponse](ctx, a, req, opts)
}

func (a *API) AcknowledgeUserOffer(ctx context.Context, params discord.AcknowledgeUserOfferParams, opts ...Option) (*discord.AcknowledgeUserOfferResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/billing/user-offer/ack")}
	req.JSON = params
	return call[discord.AcknowledgeUserOfferResponse](ctx, a, req, opts)
}

func (a *API) CreateChurnUserOffer(ctx context.Context, opts ...Option) (*discord.CreateChurnUserOfferResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/billing/churn-user-offer")}
	return call[discord.CreateChurnUserOfferResponse](ctx, a, req, opts)
}

func (a *API) GetChurnUserOffer(ctx context.Context, opts ...Option) (*discord.GetChurnUserOfferResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/billing/churn-user-offer")}
	return call[discord.GetChurnUserOfferResponse](ctx, a, req, opts)
}

func (a *API) GetUserTrialOffer(ctx context.Context, opts ...Option) (*discord.UserTrialOffer, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/billing/user-trial-offer")}
	return call[discord.UserTrialOffer](ctx, a, req, opts)
}

func (a *API) AcknowledgeUserTrialOffer(ctx context.Context, userTrialOfferID snowflake.ID, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/billing/user-trial-offer/{user_trial_offer_id}/ack", rest.P("user_trial_offer_id", userTrialOfferID))}
	return raw(ctx, a, req, opts)
}

func (a *API) RedeemUserOffer(ctx context.Context, params discord.RedeemUserOfferParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/billing/user-offer/redeem")}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) GetCheckoutRecovery(ctx context.Context, opts ...Option) (*discord.GetCheckoutRecoveryResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/billing/checkout-recovery")}
	return call[discord.GetCheckoutRecoveryResponse](ctx, a, req, opts)
}

func (a *API) ListPremiumUserAffinities(ctx context.Context, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/billing/nitro-affinity")}
	return raw(ctx, a, req, opts)
}

func (a *API) ListEligibleApplicationSubscriptionGuilds(ctx context.Context, query discord.ListEligibleApplicationSubscriptionGuildsQuery, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/billing/eligible-application-subscription-guilds")}
	req.Query = rest.Values(query)
	return raw(ctx, a, req, opts)
}

func (a *API) ListPrivateChannels(ctx context.Context, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/channels")}
	return raw(ctx, a, req, opts)
}

func (a *API) GetDMChannel(ctx context.Context, userID snowflake.ID, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/dms/{user_id}", rest.P("user_id", userID))}
	return raw(ctx, a, req, opts)
}

func (a *API) CreatePrivateChannel(ctx context.Context, params discord.CreatePrivateChannelParams, opts ...Option) (*discord.Channel, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/channels")}
	req.JSON = params
	return call[discord.Channel](ctx, a, req, opts)
}

func (a *API) CreateGroupDMShell(ctx context.Context, params discord.CreateGroupDMShellParams, opts ...Option) (*discord.Channel, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/channels/group-dm-shell")}
	req.JSON = params
	return call[discord.Channel](ctx, a, req, opts)
}

func (a *API) ListGuildChannels(ctx context.Context, guildID snowflake.ID, query discord.ListGuildChannelsQuery, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/channels", rest.P("guild_id", guildID))}
	req.Query = rest.Values(query)
	return raw(ctx, a, req, opts)
}

func (a *API) ListGuildTopReadChannels(ctx context.Context, guildID snowflake.ID, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/top-read-channels", rest.P("guild_id", guildID))}
	return raw(ctx, a, req, opts)
}

func (a *API) CreateGuildChannel(ctx context.Context, guildID snowflake.ID, params discord.CreateGuildChannelParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/guilds/{guild_id}/channels", rest.P("guild_id", guildID))}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) ModifyGuildChannelPositions(ctx context.Context, guildID snowflake.ID, params discord.ModifyGuildChannelPositionsParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/guilds/{guild_id}/channels", rest.P("guild_id", guildID))}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) GetChannel(ctx context.Context, channelID snowflake.ID, opts ...Option) (*discord.Channel, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/channels/{channel_id}", rest.P("channel_id", channelID))}
	return call[discord.Channel](ctx, a, req, opts)
}

func (a *API) ModifyChannel(ctx context.Context, channelID snowflake.ID, params discord.ModifyChannelParams, opts ...Option) (*discord.Channel, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/channels/{channel_id}", rest.P("channel_id", channelID))}
	req.JSON = params
	return call[discord.Channel](ctx, a, req, opts)
}

func (a *API) DeleteChannel(ctx context.Context, channelID snowflake.ID, query discord.DeleteChannelQuery, opts ...Option) (*discord.Channel, error) {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/channels/{channel_id}", rest.P("channel_id", channelID))}
	req.Query = rest.Values(query)
	return call[discord.Channel](ctx, a, req, opts)
}

func (a *API) BulkLeaveGroupDMs(ctx context.Context, params discord.BulkLeaveGroupDMsParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/channels/bulk-leave")}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) ModifyChannelStatus(ctx context.Context, channelID snowflake.ID, params discord.ModifyChannelStatusParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("PUT", "/channels/{channel_id}/voice-status", rest.P("channel_id", channelID))}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) ModifyChannelPermissions(ctx context.Context, channelID snowflake.ID, overwriteID snowflake.ID, params discord.ModifyChannelPermissionsParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("PUT", "/channels/{channel_id}/permissions/{overwrite_id}", rest.P("channel_id", channelID), rest.P("overwrite_id", overwriteID))}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) DeleteChannelPermission(ctx context.Context, channelID snowflake.ID, overwriteID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/channels/{channel_id}/permissions/{overwrite_id}", rest.P("channel_id", channelID), rest.P("overwrite_id", overwriteID))}
	return none(ctx, a, req, opts)
}

func (a *API) FollowChannel(ctx context.Context, channelID snowflake.ID, params discord.FollowChannelParams, opts ...Option) (*discord.FollowedChannel, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/channels/{channel_id}/followers", rest.P("channel_id", channelID))}
	req.JSON = params
	return call[discord.FollowedChannel](ctx, a, req, opts)
}

func (a *API) TriggerTypingIndicator(ctx context.Context, channelID snowflake.ID, opts ...Option) (*discord.TriggerTypingIndicatorResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/channels/{channel_id}/typing", rest.P("channel_id", channelID))}
	return call[discord.TriggerTypingIndicatorResponse](ctx, a, req, opts)
}

func (a *API) GetCallEligibility(ctx context.Context, channelID snowflake.ID, opts ...Option) (*discord.GetCallEligibilityResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/channels/{channel_id}/call", rest.P("channel_id", channelID))}
	return call[discord.GetCallEligibilityResponse](ctx, a, req, opts)
}

func (a *API) ModifyCall(ctx context.Context, channelID snowflake.ID, params discord.ModifyCallParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/channels/{channel_id}/call", rest.P("channel_id", channelID))}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) RingChannelRecipients(ctx context.Context, channelID snowflake.ID, params discord.RingChannelRecipientsParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/channels/{channel_id}/call/ring", rest.P("channel_id", channelID))}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) StopRingingChannelRecipients(ctx context.Context, channelID snowflake.ID, params discord.StopRingingChannelRecipientsParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/channels/{channel_id}/call/stop-ringing", rest.P("channel_id", channelID))}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) AddChannelRecipient(ctx context.Context, channelID snowflake.ID, userID snowflake.ID, params discord.AddChannelRecipientParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("PUT", "/channels/{channel_id}/recipients/{user_id}", rest.P("channel_id", channelID), rest.P("user_id", userID))}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) RemoveChannelRecipient(ctx context.Context, channelID snowflake.ID, userID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/channels/{channel_id}/recipients/{user_id}", rest.P("channel_id", channelID), rest.P("user_id", userID))}
	return none(ctx, a, req, opts)
}

func (a *API) UpdateMessageRequest(ctx context.Context, channelID snowflake.ID, params discord.UpdateMessageRequestParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("PUT", "/channels/{channel_id}/recipients/@me", rest.P("channel_id", channelID))}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) RejectMessageRequest(ctx context.Context, channelID snowflake.ID, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/channels/{channel_id}/recipients/@me", rest.P("channel_id", channelID))}
	return raw(ctx, a, req, opts)
}

func (a *API) BatchRejectMessageRequests(ctx context.Context, params discord.BatchRejectMessageRequestsParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("PUT", "/channels/recipients/@me/batch-reject")}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) GetSupplementalMessageRequestData(ctx context.Context, query discord.GetSupplementalMessageRequestDataQuery, opts ...Option) ([]discord.SupplementalMessageRequest, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/message-requests/supplemental-data")}
	req.Query = rest.Values(query)
	return list[discord.SupplementalMessageRequest](ctx, a, req, opts)
}

func (a *API) AcknowledgeBlockedUserWarning(ctx context.Context, channelID snowflake.ID, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/channels/{channel_id}/blocked-user-warning-dismissal", rest.P("channel_id", channelID))}
	return raw(ctx, a, req, opts)
}

func (a *API) AcknowledgeSafetyWarnings(ctx context.Context, channelID snowflake.ID, params discord.AcknowledgeSafetyWarningsParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/channels/{channel_id}/safety-warnings/ack", rest.P("channel_id", channelID))}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) AddSafetyWarning(ctx context.Context, channelID snowflake.ID, params discord.AddSafetyWarningParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/channels/{channel_id}/add-safety-warning", rest.P("channel_id", channelID))}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) DeleteSafetyWarnings(ctx context.Context, channelID snowflake.ID, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/channels/{channel_id}/safety-warnings", rest.P("channel_id", channelID))}
	return raw(ctx, a, req, opts)
}

func (a *API) ReportSafetyWarningFalsePositive(ctx context.Context, channelID snowflake.ID, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/channels/{channel_id}/safety-warning/report-false-positive", rest.P("channel_id", channelID))}
	return raw(ctx, a, req, opts)
}

func (a *API) ListGuildActiveThreads(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.ListGuildActiveThreadsResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/threads/active", rest.P("guild_id", guildID))}
	return call[discord.ListGuildActiveThreadsResponse](ctx, a, req, opts)
}

func (a *API) ListActiveThreads(ctx context.Context, channelID snowflake.ID, opts ...Option) (*discord.ListActiveThreadsResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/channels/{channel_id}/threads/active", rest.P("channel_id", channelID))}
	return call[discord.ListActiveThreadsResponse](ctx, a, req, opts)
}

func (a *API) ListPublicArchivedThreads(ctx context.Context, channelID snowflake.ID, query discord.ListPublicArchivedThreadsQuery, opts ...Option) (*discord.ListPublicArchivedThreadsResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/channels/{channel_id}/threads/archived/public", rest.P("channel_id", channelID))}
	req.Query = rest.Values(query)
	return call[discord.ListPublicArchivedThreadsResponse](ctx, a, req, opts)
}

func (a *API) ListPrivateArchivedThreads(ctx context.Context, channelID snowflake.ID, query discord.ListPrivateArchivedThreadsQuery, opts ...Option) (*discord.ListPrivateArchivedThreadsResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/channels/{channel_id}/threads/archived/private", rest.P("channel_id", channelID))}
	req.Query = rest.Values(query)
	return call[discord.ListPrivateArchivedThreadsResponse](ctx, a, req, opts)
}

func (a *API) ListJoinedPrivateArchivedThreads(ctx context.Context, channelID snowflake.ID, query discord.ListJoinedPrivateArchivedThreadsQuery, opts ...Option) (*discord.ListJoinedPrivateArchivedThreadsResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/channels/{channel_id}/users/@me/threads/archived/private", rest.P("channel_id", channelID))}
	req.Query = rest.Values(query)
	return call[discord.ListJoinedPrivateArchivedThreadsResponse](ctx, a, req, opts)
}

func (a *API) SearchThreads(ctx context.Context, channelID snowflake.ID, query discord.SearchThreadsQuery, opts ...Option) (*discord.SearchThreadsResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/channels/{channel_id}/threads/search", rest.P("channel_id", channelID))}
	req.Query = rest.Values(query)
	return call[discord.SearchThreadsResponse](ctx, a, req, opts)
}

func (a *API) CreateThreadFromMessage(ctx context.Context, channelID snowflake.ID, messageID snowflake.ID, params discord.CreateThreadFromMessageParams, opts ...Option) (*discord.Channel, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/channels/{channel_id}/messages/{message_id}/threads", rest.P("channel_id", channelID), rest.P("message_id", messageID))}
	req.JSON = params
	return call[discord.Channel](ctx, a, req, opts)
}

func (a *API) CreateThread(ctx context.Context, channelID snowflake.ID, params discord.CreateThreadParams, opts ...Option) (*discord.Channel, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/channels/{channel_id}/threads", rest.P("channel_id", channelID))}
	req.JSON = params
	return call[discord.Channel](ctx, a, req, opts)
}

func (a *API) GetChannelPostData(ctx context.Context, channelID snowflake.ID, params discord.GetChannelPostDataParams, opts ...Option) (*discord.GetChannelPostDataResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/channels/{channel_id}/post-data", rest.P("channel_id", channelID))}
	req.JSON = params
	return call[discord.GetChannelPostDataResponse](ctx, a, req, opts)
}

func (a *API) ListThreadMembers(ctx context.Context, channelID snowflake.ID, query discord.ListThreadMembersQuery, opts ...Option) ([]discord.ThreadMember, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/channels/{channel_id}/thread-members", rest.P("channel_id", channelID))}
	req.Query = rest.Values(query)
	return list[discord.ThreadMember](ctx, a, req, opts)
}

func (a *API) GetThreadMember(ctx context.Context, channelID snowflake.ID, userID snowflake.ID, query discord.GetThreadMemberQuery, opts ...Option) (*discord.ThreadMember, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/channels/{channel_id}/thread-members/{user_id}", rest.P("channel_id", channelID), rest.P("user_id", userID))}
	req.Query = rest.Values(query)
	return call[discord.ThreadMember](ctx, a, req, opts)
}

func (a *API) JoinThread(ctx context.Context, channelID snowflake.ID, query discord.JoinThreadQuery, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("PUT", "/channels/{channel_id}/thread-members/@me", rest.P("channel_id", channelID))}
	req.Query = rest.Values(query)
	return none(ctx, a, req, opts)
}

func (a *API) AddThreadMember(ctx context.Context, channelID snowflake.ID, userID snowflake.ID, query discord.AddThreadMemberQuery, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("PUT", "/channels/{channel_id}/thread-members/{user_id}", rest.P("channel_id", channelID), rest.P("user_id", userID))}
	req.Query = rest.Values(query)
	return none(ctx, a, req, opts)
}

func (a *API) ModifyThreadSettings(ctx context.Context, channelID snowflake.ID, params discord.ModifyThreadSettingsParams, opts ...Option) (*discord.ThreadMember, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/channels/{channel_id}/thread-members/@me/settings", rest.P("channel_id", channelID))}
	req.JSON = params
	return call[discord.ThreadMember](ctx, a, req, opts)
}

func (a *API) LeaveThread(ctx context.Context, channelID snowflake.ID, query discord.LeaveThreadQuery, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/channels/{channel_id}/thread-members/@me", rest.P("channel_id", channelID))}
	req.Query = rest.Values(query)
	return none(ctx, a, req, opts)
}

func (a *API) RemoveThreadMember(ctx context.Context, channelID snowflake.ID, userID snowflake.ID, query discord.RemoveThreadMemberQuery, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/channels/{channel_id}/thread-members/{user_id}", rest.P("channel_id", channelID), rest.P("user_id", userID))}
	req.Query = rest.Values(query)
	return none(ctx, a, req, opts)
}

func (a *API) CreateChannelTag(ctx context.Context, channelID snowflake.ID, params discord.CreateChannelTagParams, opts ...Option) (*discord.Channel, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/channels/{channel_id}/tags", rest.P("channel_id", channelID))}
	req.JSON = params
	return call[discord.Channel](ctx, a, req, opts)
}

func (a *API) ModifyChannelTag(ctx context.Context, channelID snowflake.ID, forumTagID snowflake.ID, params discord.ModifyChannelTagParams, opts ...Option) (*discord.Channel, error) {
	req := &rest.Request{Route: rest.NewRoute("PUT", "/channels/{channel_id}/tags/{forum_tag_id}", rest.P("channel_id", channelID), rest.P("forum_tag_id", forumTagID))}
	req.JSON = params
	return call[discord.Channel](ctx, a, req, opts)
}

func (a *API) DeleteChannelTag(ctx context.Context, channelID snowflake.ID, forumTagID snowflake.ID, opts ...Option) (*discord.Channel, error) {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/channels/{channel_id}/tags/{forum_tag_id}", rest.P("channel_id", channelID), rest.P("forum_tag_id", forumTagID))}
	return call[discord.Channel](ctx, a, req, opts)
}

func (a *API) GetChannelLinkedAccounts(ctx context.Context, channelID snowflake.ID, query discord.GetChannelLinkedAccountsQuery, opts ...Option) (*discord.GetChannelLinkedAccountsResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/channels/{channel_id}/linked-accounts", rest.P("channel_id", channelID))}
	req.Query = rest.Values(query)
	return call[discord.GetChannelLinkedAccountsResponse](ctx, a, req, opts)
}

func (a *API) RemoveLobbyLink(ctx context.Context, channelID snowflake.ID, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/channels/{channel_id}/linked-lobby", rest.P("channel_id", channelID))}
	return raw(ctx, a, req, opts)
}

func (a *API) GetCheckpoint(ctx context.Context, opts ...Option) (*discord.GetCheckpointResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/checkpoint")}
	return call[discord.GetCheckpointResponse](ctx, a, req, opts)
}

func (a *API) ClaimCheckpointAvatarDecoration(ctx context.Context, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/checkpoint/loot")}
	return none(ctx, a, req, opts)
}

func (a *API) ListCollectiblesCategories(ctx context.Context, query discord.ListCollectiblesCategoriesQuery, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/collectibles-categories")}
	req.Query = rest.Values(query)
	return raw(ctx, a, req, opts)
}

func (a *API) ListCollectiblesCategoriesV2(ctx context.Context, query discord.ListCollectiblesCategoriesV2Query, opts ...Option) (*discord.ListCollectiblesCategoriesV2Response, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/collectibles-categories/v2")}
	req.Query = rest.Values(query)
	return call[discord.ListCollectiblesCategoriesV2Response](ctx, a, req, opts)
}

func (a *API) GetCollectiblesShop(ctx context.Context, query discord.GetCollectiblesShopQuery, opts ...Option) (*discord.GetCollectiblesShopResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/collectibles-shop")}
	req.Query = rest.Values(query)
	return call[discord.GetCollectiblesShopResponse](ctx, a, req, opts)
}

func (a *API) SearchCollectibles(ctx context.Context, query discord.SearchCollectiblesQuery, opts ...Option) (*discord.SearchCollectiblesResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/shop/search")}
	req.Query = rest.Values(query)
	return call[discord.SearchCollectiblesResponse](ctx, a, req, opts)
}

func (a *API) GetCollectiblesProduct(ctx context.Context, skuid snowflake.ID, query discord.GetCollectiblesProductQuery, opts ...Option) (*discord.CollectibleProduct, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/collectibles-products/{sku_id}", rest.P("sku_id", skuid))}
	req.Query = rest.Values(query)
	return call[discord.CollectibleProduct](ctx, a, req, opts)
}

func (a *API) ListUserPurchasedCollectibles(ctx context.Context, query discord.ListUserPurchasedCollectiblesQuery, opts ...Option) ([]discord.CollectibleProduct, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/collectibles-purchases")}
	req.Query = rest.Values(query)
	return list[discord.CollectibleProduct](ctx, a, req, opts)
}

func (a *API) GetValidCollectiblesGiftRecipient(ctx context.Context, query discord.GetValidCollectiblesGiftRecipientQuery, opts ...Option) (*discord.GiftEligibility, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/valid-collectibles-gift-recipient")}
	req.Query = rest.Values(query)
	return call[discord.GiftEligibility](ctx, a, req, opts)
}

func (a *API) GetValidCollectiblesGiftRecipientsBatch(ctx context.Context, query discord.GetValidCollectiblesGiftRecipientsBatchQuery, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/valid-collectibles-gift-recipients-batch")}
	req.Query = rest.Values(query)
	return raw(ctx, a, req, opts)
}

func (a *API) ClaimPremiumCollectiblesProduct(ctx context.Context, params discord.ClaimPremiumCollectiblesProductParams, opts ...Option) ([]discord.CollectibleProduct, error) {
	req := &rest.Request{Route: rest.NewRoute("PUT", "/users/@me/claim-premium-collectibles-product")}
	req.JSON = params
	return list[discord.CollectibleProduct](ctx, a, req, opts)
}

func (a *API) ClaimRewardCategoryProduct(ctx context.Context, params discord.ClaimRewardCategoryProductParams, opts ...Option) ([]discord.CollectibleProduct, error) {
	req := &rest.Request{Route: rest.NewRoute("PUT", "/users/@me/claim-reward-category-product")}
	req.JSON = params
	return list[discord.CollectibleProduct](ctx, a, req, opts)
}

func (a *API) GetCollectiblesMarketing(ctx context.Context, query discord.GetCollectiblesMarketingQuery, opts ...Option) (*discord.GetCollectiblesMarketingResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/collectibles-marketing")}
	req.Query = rest.Values(query)
	return call[discord.GetCollectiblesMarketingResponse](ctx, a, req, opts)
}

func (a *API) GetCollectiblesShopTabLayout(ctx context.Context, collectiblesShopTab string, opts ...Option) (*discord.GetCollectiblesShopTabLayoutResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/collectibles-shop-tab-layouts/{collectibles_shop_tab}", rest.P("collectibles_shop_tab", collectiblesShopTab))}
	return call[discord.GetCollectiblesShopTabLayoutResponse](ctx, a, req, opts)
}

func (a *API) AuthorizeUserConnection(ctx context.Context, connectionType string, query discord.AuthorizeUserConnectionQuery, opts ...Option) (*discord.AuthorizeUserConnectionResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/connections/{connection_type}/authorize", rest.P("connection_type", connectionType))}
	req.Query = rest.Values(query)
	return call[discord.AuthorizeUserConnectionResponse](ctx, a, req, opts)
}

func (a *API) CreateUserConnectionCallback(ctx context.Context, connectionType string, params discord.CreateUserConnectionCallbackParams, opts ...Option) (*discord.Connection, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/connections/{connection_type}/callback", rest.P("connection_type", connectionType))}
	req.JSON = params
	return call[discord.Connection](ctx, a, req, opts)
}

func (a *API) CreateContactSyncConnection(ctx context.Context, connectionID snowflake.ID, params discord.CreateContactSyncConnectionParams, opts ...Option) (*discord.Connection, error) {
	req := &rest.Request{Route: rest.NewRoute("PUT", "/users/@me/connections/contacts/{connection_id}", rest.P("connection_id", connectionID))}
	req.JSON = params
	return call[discord.Connection](ctx, a, req, opts)
}

func (a *API) UpdateExternalFriendListEntries(ctx context.Context, connectionID snowflake.ID, params discord.UpdateExternalFriendListEntriesParams, opts ...Option) (*discord.UpdateExternalFriendListEntriesResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("PUT", "/users/@me/connections/contacts/{connection_id}/external-friend-list-entries", rest.P("connection_id", connectionID))}
	req.JSON = params
	return call[discord.UpdateExternalFriendListEntriesResponse](ctx, a, req, opts)
}

func (a *API) ContactSyncSettings(ctx context.Context, connectionID snowflake.ID, opts ...Option) (*discord.ContactSyncSettingsResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/connections/contacts/{connection_id}/external-friend-list-entries/settings", rest.P("connection_id", connectionID))}
	return call[discord.ContactSyncSettingsResponse](ctx, a, req, opts)
}

func (a *API) CreateDomainConnection(ctx context.Context, connectionID snowflake.ID, opts ...Option) (*discord.Connection, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/connections/domain/{connection_id}", rest.P("connection_id", connectionID))}
	return call[discord.Connection](ctx, a, req, opts)
}

func (a *API) ListUserConnections(ctx context.Context, opts ...Option) ([]discord.Connection, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/connections")}
	return list[discord.Connection](ctx, a, req, opts)
}

func (a *API) GetUserConnectionAccessToken(ctx context.Context, connectionType string, connectionID snowflake.ID, opts ...Option) (*discord.GetUserConnectionAccessTokenResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/connections/{connection_type}/{connection_id}/access-token", rest.P("connection_type", connectionType), rest.P("connection_id", connectionID))}
	return call[discord.GetUserConnectionAccessTokenResponse](ctx, a, req, opts)
}

func (a *API) ListUserConnectionSubreddits(ctx context.Context, connectionID snowflake.ID, opts ...Option) ([]discord.Subreddit, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/connections/reddit/{connection_id}/subreddits", rest.P("connection_id", connectionID))}
	return list[discord.Subreddit](ctx, a, req, opts)
}

func (a *API) RefreshUserConnection(ctx context.Context, connectionType string, connectionID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/connections/{connection_type}/{connection_id}/refresh", rest.P("connection_type", connectionType), rest.P("connection_id", connectionID))}
	return none(ctx, a, req, opts)
}

func (a *API) ModifyUserConnection(ctx context.Context, connectionType string, connectionID snowflake.ID, params discord.ModifyUserConnectionParams, opts ...Option) (*discord.Connection, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/users/@me/connections/{connection_type}/{connection_id}", rest.P("connection_type", connectionType), rest.P("connection_id", connectionID))}
	req.JSON = params
	return call[discord.Connection](ctx, a, req, opts)
}

func (a *API) DeleteUserConnection(ctx context.Context, connectionType string, connectionID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/users/@me/connections/{connection_type}/{connection_id}", rest.P("connection_type", connectionType), rest.P("connection_id", connectionID))}
	return none(ctx, a, req, opts)
}

func (a *API) ListUserLinkedConnections(ctx context.Context, opts ...Option) ([]discord.Connection, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/linked-connections")}
	return list[discord.Connection](ctx, a, req, opts)
}

func (a *API) CreateConsoleConnection(ctx context.Context, params discord.CreateConsoleConnectionParams, opts ...Option) (*discord.CreateConsoleConnectionResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/consoles/connect-request")}
	req.JSON = params
	return call[discord.CreateConsoleConnectionResponse](ctx, a, req, opts)
}

func (a *API) CancelConsoleConnectionRequest(ctx context.Context, nonce string, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/consoles/connect-request/{nonce}", rest.P("nonce", nonce))}
	return none(ctx, a, req, opts)
}

func (a *API) ListConsoleDevices(ctx context.Context, connectionType string, opts ...Option) (*discord.ListConsoleDevicesResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/consoles/{connection_type}/devices", rest.P("connection_type", connectionType))}
	return call[discord.ListConsoleDevicesResponse](ctx, a, req, opts)
}

func (a *API) SendConsoleCommand(ctx context.Context, connectionType string, deviceID snowflake.ID, params discord.SendConsoleCommandParams, opts ...Option) (*discord.SendConsoleCommandResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/consoles/{connection_type}/devices/{device_id}/commands", rest.P("connection_type", connectionType), rest.P("device_id", deviceID))}
	req.JSON = params
	return call[discord.SendConsoleCommandResponse](ctx, a, req, opts)
}

func (a *API) CancelConsoleCommand(ctx context.Context, connectionType string, deviceID snowflake.ID, commandID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/consoles/{connection_type}/devices/{device_id}/commands/{command_id}", rest.P("connection_type", connectionType), rest.P("device_id", deviceID), rest.P("command_id", commandID))}
	return none(ctx, a, req, opts)
}

func (a *API) GetDirectoryCounts(ctx context.Context, channelID snowflake.ID, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/channels/{channel_id}/directory-entries/counts", rest.P("channel_id", channelID))}
	return raw(ctx, a, req, opts)
}

func (a *API) ListDirectoryEntries(ctx context.Context, channelID snowflake.ID, opts ...Option) ([]discord.DirectoryEntry, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/channels/{channel_id}/directory-entries", rest.P("channel_id", channelID))}
	return list[discord.DirectoryEntry](ctx, a, req, opts)
}

func (a *API) ListPartialDirectoryEntries(ctx context.Context, channelID snowflake.ID, opts ...Option) ([]discord.DirectoryEntry, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/channels/{channel_id}/directory-entries/list", rest.P("channel_id", channelID))}
	return list[discord.DirectoryEntry](ctx, a, req, opts)
}

func (a *API) SearchDirectoryEntries(ctx context.Context, channelID snowflake.ID, opts ...Option) ([]discord.DirectoryEntry, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/channels/{channel_id}/directory-entries/search", rest.P("channel_id", channelID))}
	return list[discord.DirectoryEntry](ctx, a, req, opts)
}

func (a *API) GetDirectoryEntry(ctx context.Context, channelID snowflake.ID, entityID snowflake.ID, opts ...Option) (*discord.DirectoryEntry, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/channels/{channel_id}/directory-entry/{entity_id}", rest.P("channel_id", channelID), rest.P("entity_id", entityID))}
	return call[discord.DirectoryEntry](ctx, a, req, opts)
}

func (a *API) CreateDirectoryEntry(ctx context.Context, channelID snowflake.ID, entityID snowflake.ID, params discord.CreateDirectoryEntryParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/channels/{channel_id}/directory-entry/{entity_id}", rest.P("channel_id", channelID), rest.P("entity_id", entityID))}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) ModifyDirectoryEntry(ctx context.Context, channelID snowflake.ID, entityID snowflake.ID, params discord.ModifyDirectoryEntryParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/channels/{channel_id}/directory-entry/{entity_id}", rest.P("channel_id", channelID), rest.P("entity_id", entityID))}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) DeleteDirectoryEntry(ctx context.Context, channelID snowflake.ID, entityID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/channels/{channel_id}/directory-entry/{entity_id}", rest.P("channel_id", channelID), rest.P("entity_id", entityID))}
	return none(ctx, a, req, opts)
}

func (a *API) GetDirectoryBroadcastInfo(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.GetDirectoryBroadcastInfoResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/directory-entries/broadcast", rest.P("guild_id", guildID))}
	return call[discord.GetDirectoryBroadcastInfoResponse](ctx, a, req, opts)
}

func (a *API) ListDiscoverableGuilds(ctx context.Context, query discord.ListDiscoverableGuildsQuery, opts ...Option) (*discord.ListDiscoverableGuildsResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/discoverable-guilds")}
	req.Query = rest.Values(query)
	return call[discord.ListDiscoverableGuildsResponse](ctx, a, req, opts)
}

func (a *API) SearchDiscoverableGuilds(ctx context.Context, query discord.SearchDiscoverableGuildsQuery, opts ...Option) ([]discord.DiscoverableGuild, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/discoverable-guilds/search")}
	req.Query = rest.Values(query)
	return list[discord.DiscoverableGuild](ctx, a, req, opts)
}

func (a *API) SearchPublishedGuilds(ctx context.Context, query discord.SearchPublishedGuildsQuery, opts ...Option) ([]discord.DiscoverableGuild, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/discovery/search")}
	req.Query = rest.Values(query)
	return list[discord.DiscoverableGuild](ctx, a, req, opts)
}

func (a *API) GetDiscoverySlug(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.GetDiscoverySlugResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/discovery/{guild_id}", rest.P("guild_id", guildID))}
	return call[discord.GetDiscoverySlugResponse](ctx, a, req, opts)
}

func (a *API) ListDiscoveryCategories(ctx context.Context, query discord.ListDiscoveryCategoriesQuery, opts ...Option) ([]discord.DiscoveryCategory, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/discovery/categories")}
	req.Query = rest.Values(query)
	return list[discord.DiscoveryCategory](ctx, a, req, opts)
}

func (a *API) ValidateDiscoverySearchTerm(ctx context.Context, query discord.ValidateDiscoverySearchTermQuery, opts ...Option) (*discord.ValidateDiscoverySearchTermResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/discovery/valid-term")}
	req.Query = rest.Values(query)
	return call[discord.ValidateDiscoverySearchTermResponse](ctx, a, req, opts)
}

func (a *API) GetGuildDiscoveryRequirements(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.DiscoveryRequirements, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/discovery-requirements", rest.P("guild_id", guildID))}
	return call[discord.DiscoveryRequirements](ctx, a, req, opts)
}

func (a *API) GetGuildDiscoveryMetadata(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.DiscoveryMetadata, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/discovery-metadata", rest.P("guild_id", guildID))}
	return call[discord.DiscoveryMetadata](ctx, a, req, opts)
}

func (a *API) ModifyGuildDiscoveryMetadata(ctx context.Context, guildID snowflake.ID, params discord.ModifyGuildDiscoveryMetadataParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/guilds/{guild_id}/discovery-metadata", rest.P("guild_id", guildID))}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) AddGuildDiscoverySubcategory(ctx context.Context, guildID snowflake.ID, discoveryCategoryID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("PUT", "/guilds/{guild_id}/discovery-categories/{discovery_category_id}", rest.P("guild_id", guildID), rest.P("discovery_category_id", discoveryCategoryID))}
	return none(ctx, a, req, opts)
}

func (a *API) RemoveGuildDiscoverySubcategory(ctx context.Context, guildID snowflake.ID, discoveryCategoryID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/guilds/{guild_id}/discovery-categories/{discovery_category_id}", rest.P("guild_id", guildID), rest.P("discovery_category_id", discoveryCategoryID))}
	return none(ctx, a, req, opts)
}

func (a *API) GetGuildProfile(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.GuildProfile, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/profile", rest.P("guild_id", guildID))}
	return call[discord.GuildProfile](ctx, a, req, opts)
}

func (a *API) ModifyGuildProfile(ctx context.Context, guildID snowflake.ID, params discord.ModifyGuildProfileParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/guilds/{guild_id}/profile", rest.P("guild_id", guildID))}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) ListGuildEmojis(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.Emoji, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/emojis", rest.P("guild_id", guildID))}
	return list[discord.Emoji](ctx, a, req, opts)
}

func (a *API) GetGuildEmoji(ctx context.Context, guildID snowflake.ID, emojiID snowflake.ID, opts ...Option) (*discord.Emoji, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/emojis/{emoji_id}", rest.P("guild_id", guildID), rest.P("emoji_id", emojiID))}
	return call[discord.Emoji](ctx, a, req, opts)
}

func (a *API) ListGuildTopEmojis(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.ListGuildTopEmojisResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/top-emojis", rest.P("guild_id", guildID))}
	return call[discord.ListGuildTopEmojisResponse](ctx, a, req, opts)
}

func (a *API) GetEmojiGuild(ctx context.Context, emojiID snowflake.ID, opts ...Option) (*discord.DiscoverableGuild, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/emojis/{emoji_id}/guild", rest.P("emoji_id", emojiID))}
	return call[discord.DiscoverableGuild](ctx, a, req, opts)
}

func (a *API) GetEmojiSource(ctx context.Context, emojiID snowflake.ID, opts ...Option) (*discord.GetEmojiSourceResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/emojis/{emoji_id}/source", rest.P("emoji_id", emojiID))}
	return call[discord.GetEmojiSourceResponse](ctx, a, req, opts)
}

func (a *API) CreateGuildEmoji(ctx context.Context, guildID snowflake.ID, params discord.CreateGuildEmojiParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/guilds/{guild_id}/emojis", rest.P("guild_id", guildID))}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) ModifyGuildEmoji(ctx context.Context, guildID snowflake.ID, emojiID snowflake.ID, params discord.ModifyGuildEmojiParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/guilds/{guild_id}/emojis/{emoji_id}", rest.P("guild_id", guildID), rest.P("emoji_id", emojiID))}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) DeleteGuildEmoji(ctx context.Context, guildID snowflake.ID, emojiID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/guilds/{guild_id}/emojis/{emoji_id}", rest.P("guild_id", guildID), rest.P("emoji_id", emojiID))}
	return none(ctx, a, req, opts)
}

func (a *API) ListApplicationEmojis(ctx context.Context, applicationID snowflake.ID, opts ...Option) (*discord.ListApplicationEmojisResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/applications/{application_id}/emojis", rest.P("application_id", applicationID))}
	return call[discord.ListApplicationEmojisResponse](ctx, a, req, opts)
}

func (a *API) GetApplicationEmoji(ctx context.Context, applicationID snowflake.ID, emojiID snowflake.ID, opts ...Option) (*discord.Emoji, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/applications/{application_id}/emojis/{emoji_id}", rest.P("application_id", applicationID), rest.P("emoji_id", emojiID))}
	return call[discord.Emoji](ctx, a, req, opts)
}

func (a *API) CreateApplicationEmoji(ctx context.Context, applicationID snowflake.ID, params discord.CreateApplicationEmojiParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/applications/{application_id}/emojis", rest.P("application_id", applicationID))}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) ModifyApplicationEmoji(ctx context.Context, applicationID snowflake.ID, emojiID snowflake.ID, params discord.ModifyApplicationEmojiParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/applications/{application_id}/emojis/{emoji_id}", rest.P("application_id", applicationID), rest.P("emoji_id", emojiID))}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) DeleteApplicationEmoji(ctx context.Context, applicationID snowflake.ID, emojiID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/applications/{application_id}/emojis/{emoji_id}", rest.P("application_id", applicationID), rest.P("emoji_id", emojiID))}
	return none(ctx, a, req, opts)
}

func (a *API) ListUserEntitlements(ctx context.Context, query discord.ListUserEntitlementsQuery, opts ...Option) ([]discord.Entitlement, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/entitlements")}
	req.Query = rest.Values(query)
	return list[discord.Entitlement](ctx, a, req, opts)
}

func (a *API) ListUserGiftableEntitlements(ctx context.Context, query discord.ListUserGiftableEntitlementsQuery, opts ...Option) ([]discord.Entitlement, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/entitlements/gifts")}
	req.Query = rest.Values(query)
	return list[discord.Entitlement](ctx, a, req, opts)
}

func (a *API) ListGuildEntitlements(ctx context.Context, guildID snowflake.ID, query discord.ListGuildEntitlementsQuery, opts ...Option) ([]discord.Entitlement, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/entitlements", rest.P("guild_id", guildID))}
	req.Query = rest.Values(query)
	return list[discord.Entitlement](ctx, a, req, opts)
}

func (a *API) ListUserApplicationEntitlements(ctx context.Context, applicationID snowflake.ID, query discord.ListUserApplicationEntitlementsQuery, opts ...Option) ([]discord.Entitlement, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/applications/{application_id}/entitlements", rest.P("application_id", applicationID))}
	req.Query = rest.Values(query)
	return list[discord.Entitlement](ctx, a, req, opts)
}

func (a *API) CreateApplicationEntitlement(ctx context.Context, applicationID snowflake.ID, params discord.CreateApplicationEntitlementParams, opts ...Option) (*discord.Entitlement, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/applications/{application_id}/entitlements", rest.P("application_id", applicationID))}
	req.JSON = params
	return call[discord.Entitlement](ctx, a, req, opts)
}

func (a *API) GetGiftCode(ctx context.Context, giftCodeCode string, query discord.GetGiftCodeQuery, opts ...Option) (*discord.GiftCode, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/entitlements/gift-codes/{gift_code_code}", rest.P("gift_code_code", giftCodeCode))}
	req.Query = rest.Values(query)
	return call[discord.GiftCode](ctx, a, req, opts)
}

func (a *API) RedeemGiftCode(ctx context.Context, giftCodeCode string, params discord.RedeemGiftCodeParams, opts ...Option) (*discord.Entitlement, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/entitlements/gift-codes/{gift_code_code}/redeem", rest.P("gift_code_code", giftCodeCode))}
	req.JSON = params
	return call[discord.Entitlement](ctx, a, req, opts)
}

func (a *API) ListUserGiftCodes(ctx context.Context, query discord.ListUserGiftCodesQuery, opts ...Option) ([]discord.GiftCode, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/entitlements/gift-codes")}
	req.Query = rest.Values(query)
	return list[discord.GiftCode](ctx, a, req, opts)
}

func (a *API) CreateUserGiftCode(ctx context.Context, params discord.CreateUserGiftCodeParams, opts ...Option) (*discord.GiftCode, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/entitlements/gift-codes")}
	req.JSON = params
	return call[discord.GiftCode](ctx, a, req, opts)
}

func (a *API) RevokeUserGiftCode(ctx context.Context, giftCodeCode string, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/users/@me/entitlements/gift-codes/{gift_code_code}", rest.P("gift_code_code", giftCodeCode))}
	return none(ctx, a, req, opts)
}

func (a *API) ListApplicationGiftCodeBatches(ctx context.Context, applicationID snowflake.ID, opts ...Option) ([]discord.GiftCodeBatch, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/applications/{application_id}/gift-code-batches", rest.P("application_id", applicationID))}
	return list[discord.GiftCodeBatch](ctx, a, req, opts)
}

func (a *API) CreateApplicationGiftCodeBatch(ctx context.Context, applicationID snowflake.ID, params discord.CreateApplicationGiftCodeBatchParams, opts ...Option) (*discord.GiftCodeBatch, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/applications/{application_id}/gift-code-batches", rest.P("application_id", applicationID))}
	req.JSON = params
	return call[discord.GiftCodeBatch](ctx, a, req, opts)
}

func (a *API) GetApplicationGiftCodeBatch(ctx context.Context, applicationID snowflake.ID, giftCodeBatchID snowflake.ID, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/applications/{application_id}/gift-code-batches/{gift_code_batch_id}", rest.P("application_id", applicationID), rest.P("gift_code_batch_id", giftCodeBatchID))}
	return raw(ctx, a, req, opts)
}

func (a *API) GetFamilyCenterOverview(ctx context.Context, opts ...Option) (*discord.FamilyCenter, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/family-center/@me")}
	return call[discord.FamilyCenter](ctx, a, req, opts)
}

func (a *API) GetLinkCode(ctx context.Context, opts ...Option) (*discord.GetLinkCodeResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/family-center/@me/link-code")}
	return call[discord.GetLinkCodeResponse](ctx, a, req, opts)
}

func (a *API) GetLinkedUsers(ctx context.Context, opts ...Option) (*discord.LinkedUser, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/linked-users")}
	return call[discord.LinkedUser](ctx, a, req, opts)
}

func (a *API) CreateLinkedUserRequest(ctx context.Context, params discord.CreateLinkedUserRequestParams, opts ...Option) (*discord.LinkedUser, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/linked-users")}
	req.JSON = params
	return call[discord.LinkedUser](ctx, a, req, opts)
}

func (a *API) ModifyLinkedUser(ctx context.Context, params discord.ModifyLinkedUserParams, opts ...Option) ([]discord.LinkedUser, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/users/@me/linked-users")}
	req.JSON = params
	return list[discord.LinkedUser](ctx, a, req, opts)
}

func (a *API) RemoveLinkedUser(ctx context.Context, params discord.RemoveLinkedUserParams, opts ...Option) ([]discord.LinkedUser, error) {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/users/@me/linked-users")}
	req.JSON = params
	return list[discord.LinkedUser](ctx, a, req, opts)
}

func (a *API) CreateGameInvite(ctx context.Context, params discord.CreateGameInviteParams, opts ...Option) (*discord.CreateGameInviteResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/game-invite/@me")}
	req.JSON = params
	return call[discord.CreateGameInviteResponse](ctx, a, req, opts)
}

func (a *API) DeleteGameInvite(ctx context.Context, gameInviteInviteID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/game-invite/@me/{game_invite_invite_id}", rest.P("game_invite_invite_id", gameInviteInviteID))}
	return none(ctx, a, req, opts)
}

func (a *API) DeleteGameInvites(ctx context.Context, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/game-invite/@me")}
	return none(ctx, a, req, opts)
}

func (a *API) ListApplicationGameClaims(ctx context.Context, applicationID snowflake.ID, opts ...Option) ([]discord.GameClaim, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/applications/{application_id}/game-claims", rest.P("application_id", applicationID))}
	return list[discord.GameClaim](ctx, a, req, opts)
}

func (a *API) GetApplicationGameClaim(ctx context.Context, applicationID snowflake.ID, gameClaimID snowflake.ID, opts ...Option) (*discord.GameClaim, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/applications/{application_id}/game-claims/{game_claim_id}", rest.P("application_id", applicationID), rest.P("game_claim_id", gameClaimID))}
	return call[discord.GameClaim](ctx, a, req, opts)
}

func (a *API) CreateApplicationGameClaim(ctx context.Context, applicationID snowflake.ID, gameID snowflake.ID, query discord.CreateApplicationGameClaimQuery, opts ...Option) (*discord.GameClaim, error) {
	req := &rest.Request{Route: rest.NewRoute("PUT", "/applications/{application_id}/game-claims/{game_id}", rest.P("application_id", applicationID), rest.P("game_id", gameID))}
	req.Query = rest.Values(query)
	return call[discord.GameClaim](ctx, a, req, opts)
}

func (a *API) ModifyApplicationGameClaim(ctx context.Context, applicationID snowflake.ID, gameClaimID snowflake.ID, params discord.ModifyApplicationGameClaimParams, opts ...Option) (*discord.GameClaim, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/applications/{application_id}/game-claims/{game_claim_id}", rest.P("application_id", applicationID), rest.P("game_claim_id", gameClaimID))}
	req.JSON = params
	return call[discord.GameClaim](ctx, a, req, opts)
}

func (a *API) DeleteApplicationGameClaim(ctx context.Context, applicationID snowflake.ID, gameClaimID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/applications/{application_id}/game-claims/{game_claim_id}", rest.P("application_id", applicationID), rest.P("game_claim_id", gameClaimID))}
	return none(ctx, a, req, opts)
}

func (a *API) ListDetectableNonGameApplications(ctx context.Context, opts ...Option) ([]discord.DetectableApplication, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/applications/non-games/detectable")}
	return list[discord.DetectableApplication](ctx, a, req, opts)
}

func (a *API) ListDetectableGames(ctx context.Context, opts ...Option) ([]discord.DetectableApplication, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/games/detectable")}
	return list[discord.DetectableApplication](ctx, a, req, opts)
}

func (a *API) ListDetectableGameExclusions(ctx context.Context, opts ...Option) (*discord.ListDetectableGameExclusionsResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/games/detectable/exclusions")}
	return call[discord.ListDetectableGameExclusionsResponse](ctx, a, req, opts)
}

func (a *API) ListGames(ctx context.Context, query discord.ListGamesQuery, opts ...Option) ([]discord.Game, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/games")}
	req.Query = rest.Values(query)
	return list[discord.Game](ctx, a, req, opts)
}

func (a *API) GetGame(ctx context.Context, gameID snowflake.ID, query discord.GetGameQuery, opts ...Option) (*discord.Game, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/games/{game_id}", rest.P("game_id", gameID))}
	req.Query = rest.Values(query)
	return call[discord.Game](ctx, a, req, opts)
}

func (a *API) ListGameAnnouncements(ctx context.Context, gameID snowflake.ID, query discord.ListGameAnnouncementsQuery, opts ...Option) (*discord.ListGameAnnouncementsResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/games/{game_id}/announcements", rest.P("game_id", gameID))}
	req.Query = rest.Values(query)
	return call[discord.ListGameAnnouncementsResponse](ctx, a, req, opts)
}

func (a *API) ListGuildGrowthActivationOverview(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.GrowthActivationOverview, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/analytics/growth-activation/overview", rest.P("guild_id", guildID))}
	return list[discord.GrowthActivationOverview](ctx, a, req, opts)
}

func (a *API) ListGuildGrowthActivationJoins(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.GrowthActivationJoins, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/analytics/growth-activation/joins", rest.P("guild_id", guildID))}
	return list[discord.GrowthActivationJoins](ctx, a, req, opts)
}

func (a *API) ListGuildGrowthActivationJoinsByInvite(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.GrowthActivationJoins, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/analytics/growth-activation/joins-by-invite-link", rest.P("guild_id", guildID))}
	return list[discord.GrowthActivationJoins](ctx, a, req, opts)
}

func (a *API) ListGuildGrowthActivationJoinsByReferrer(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.GrowthActivationJoinsByReferrer, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/analytics/growth-activation/joins-by-referrer", rest.P("guild_id", guildID))}
	return list[discord.GrowthActivationJoinsByReferrer](ctx, a, req, opts)
}

func (a *API) ListGuildGrowthActivationJoinsBySources(ctx context.Context, guildID snowflake.ID, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/analytics/growth-activation/joins-by-source", rest.P("guild_id", guildID))}
	return raw(ctx, a, req, opts)
}

func (a *API) ListGuildGrowthActivationLeavers(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.GrowthActivationLeavers, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/analytics/growth-activation/leavers", rest.P("guild_id", guildID))}
	return list[discord.GrowthActivationLeavers](ctx, a, req, opts)
}

func (a *API) ListGuildGrowthActivationPercentages(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.GrowthActivationPercentage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/analytics/growth-activation/activation", rest.P("guild_id", guildID))}
	return list[discord.GrowthActivationPercentage](ctx, a, req, opts)
}

func (a *API) ListGuildGrowthActivationRetention(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.GrowthActivationRetention, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/analytics/growth-activation/retention", rest.P("guild_id", guildID))}
	return list[discord.GrowthActivationRetention](ctx, a, req, opts)
}

func (a *API) ListGuildGrowthActivationMembership(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.GrowthActivationMembership, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/analytics/growth-activation/membership", rest.P("guild_id", guildID))}
	return list[discord.GrowthActivationMembership](ctx, a, req, opts)
}

func (a *API) ListGuildEngagementBase(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.EngagementBase, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/analytics/engagement/base", rest.P("guild_id", guildID))}
	return list[discord.EngagementBase](ctx, a, req, opts)
}

func (a *API) ListGuildEngagementOverview(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.EngagementOverview, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/analytics/engagement/overview", rest.P("guild_id", guildID))}
	return list[discord.EngagementOverview](ctx, a, req, opts)
}

func (a *API) ListGuildEngagementMuters(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.EngagementMuters, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/analytics/engagement/muters", rest.P("guild_id", guildID))}
	return list[discord.EngagementMuters](ctx, a, req, opts)
}

func (a *API) ListGuildEngagementPruneableMembers(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.EngagementPruneableMembers, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/analytics/engagement/pruneable-members", rest.P("guild_id", guildID))}
	return list[discord.EngagementPruneableMembers](ctx, a, req, opts)
}

func (a *API) ListGuildEngagementTextChannels(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.TextChannelEngagement, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/analytics/engagement/text-channels", rest.P("guild_id", guildID))}
	return list[discord.TextChannelEngagement](ctx, a, req, opts)
}

func (a *API) ListGuildEngagementVoiceChannels(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.VoiceChannelEngagement, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/analytics/engagement/voice-channels", rest.P("guild_id", guildID))}
	return list[discord.VoiceChannelEngagement](ctx, a, req, opts)
}

func (a *API) ListGuildAudienceNewMembersByDiscordTenure(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.AudienceNewMembersByDiscordTenure, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/analytics/audience/new-members-by-discord-tenure", rest.P("guild_id", guildID))}
	return list[discord.AudienceNewMembersByDiscordTenure](ctx, a, req, opts)
}

func (a *API) ListGuildAudienceParticipatorsByGuildTenure(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.AudienceParticipatorsByGuildTenure, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/analytics/audience/participators-by-guild-tenure", rest.P("guild_id", guildID))}
	return list[discord.AudienceParticipatorsByGuildTenure](ctx, a, req, opts)
}

func (a *API) ListGuildAudienceParticipatorsByPlatform(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.AudienceParticipatorsByPlatform, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/analytics/audience/participators-by-platform", rest.P("guild_id", guildID))}
	return list[discord.AudienceParticipatorsByPlatform](ctx, a, req, opts)
}

func (a *API) ListGuildAudienceParticipatorsByRegistrationCountry(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.AudienceParticipatorsByRegistrationCountry, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/analytics/audience/participators-by-reg-country", rest.P("guild_id", guildID))}
	return list[discord.AudienceParticipatorsByRegistrationCountry](ctx, a, req, opts)
}

func (a *API) ListGuildChannelFollowingOverview(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.ChannelFollowingOverview, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/analytics/channel-following/overview", rest.P("guild_id", guildID))}
	return list[discord.ChannelFollowingOverview](ctx, a, req, opts)
}

func (a *API) ListGuildChannelFollowingByChannel(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.ChannelFollowingByChannel, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/analytics/channel-following/announcements-by-channel", rest.P("guild_id", guildID))}
	return list[discord.ChannelFollowingByChannel](ctx, a, req, opts)
}

func (a *API) ListGuildChannelFollowingReach(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.ChannelFollowingReach, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/analytics/channel-following/reach", rest.P("guild_id", guildID))}
	return list[discord.ChannelFollowingReach](ctx, a, req, opts)
}

func (a *API) ListGuildChannelFollowingGuildSize(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.ChannelFollowingGuildSize, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/analytics/channel-following/following-guild-size", rest.P("guild_id", guildID))}
	return list[discord.ChannelFollowingGuildSize](ctx, a, req, opts)
}

func (a *API) ListGuildChannelFollowingGuildSizeByChannel(ctx context.Context, guildID snowflake.ID, query discord.ListGuildChannelFollowingGuildSizeByChannelQuery, opts ...Option) ([]discord.ChannelFollowingGuildSizeByChannel, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/analytics/channel-following/following-guild-size-by-channel", rest.P("guild_id", guildID))}
	req.Query = rest.Values(query)
	return list[discord.ChannelFollowingGuildSizeByChannel](ctx, a, req, opts)
}

func (a *API) ListGuildWelcomeScreenFunnel(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.WelcomeScreenFunnel, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/analytics/welcome-screen/funnel", rest.P("guild_id", guildID))}
	return list[discord.WelcomeScreenFunnel](ctx, a, req, opts)
}

func (a *API) ListGuildWelcomeScreenUsers(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.WelcomeScreenUsers, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/analytics/welcome-screen/users", rest.P("guild_id", guildID))}
	return list[discord.WelcomeScreenUsers](ctx, a, req, opts)
}

func (a *API) ListUserGuildScheduledEvents(ctx context.Context, query discord.ListUserGuildScheduledEventsQuery, opts ...Option) ([]discord.GuildScheduledEventUser, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/scheduled-events")}
	req.Query = rest.Values(query)
	return list[discord.GuildScheduledEventUser](ctx, a, req, opts)
}

func (a *API) ListGuildScheduledEvents(ctx context.Context, guildID snowflake.ID, query discord.ListGuildScheduledEventsQuery, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/scheduled-events", rest.P("guild_id", guildID))}
	req.Query = rest.Values(query)
	return raw(ctx, a, req, opts)
}

func (a *API) GetGuildScheduledEvent(ctx context.Context, guildID snowflake.ID, guildScheduledEventID snowflake.ID, query discord.GetGuildScheduledEventQuery, opts ...Option) (*discord.GuildScheduledEvent, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/scheduled-events/{guild_scheduled_event_id}", rest.P("guild_id", guildID), rest.P("guild_scheduled_event_id", guildScheduledEventID))}
	req.Query = rest.Values(query)
	return call[discord.GuildScheduledEvent](ctx, a, req, opts)
}

func (a *API) CreateGuildScheduledEvent(ctx context.Context, guildID snowflake.ID, params discord.GuildScheduledEventCreateGuildScheduledEventParams, opts ...Option) (*discord.GuildScheduledEvent, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/guilds/{guild_id}/scheduled-events", rest.P("guild_id", guildID))}
	req.JSON = params
	return call[discord.GuildScheduledEvent](ctx, a, req, opts)
}

func (a *API) GetGuildScheduledEventUserCount(ctx context.Context, guildID snowflake.ID, guildScheduledEventID snowflake.ID, query discord.GetGuildScheduledEventUserCountQuery, opts ...Option) (*discord.GetGuildScheduledEventUserCountResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/scheduled-events/{guild_scheduled_event_id}/users/counts", rest.P("guild_id", guildID), rest.P("guild_scheduled_event_id", guildScheduledEventID))}
	req.Query = rest.Values(query)
	return call[discord.GetGuildScheduledEventUserCountResponse](ctx, a, req, opts)
}

func (a *API) ListGuildScheduledEventUsers(ctx context.Context, guildID snowflake.ID, guildScheduledEventID snowflake.ID, query discord.ListGuildScheduledEventUsersQuery, opts ...Option) (*discord.ListGuildScheduledEventUsersResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/scheduled-events/{guild_scheduled_event_id}/users", rest.P("guild_id", guildID), rest.P("guild_scheduled_event_id", guildScheduledEventID))}
	req.Query = rest.Values(query)
	return call[discord.ListGuildScheduledEventUsersResponse](ctx, a, req, opts)
}

func (a *API) CreateGuildScheduledEventUser(ctx context.Context, guildID snowflake.ID, guildScheduledEventID snowflake.ID, opts ...Option) (*discord.GuildScheduledEventUser, error) {
	req := &rest.Request{Route: rest.NewRoute("PUT", "/guilds/{guild_id}/scheduled-events/{guild_scheduled_event_id}/users/@me", rest.P("guild_id", guildID), rest.P("guild_scheduled_event_id", guildScheduledEventID))}
	return call[discord.GuildScheduledEventUser](ctx, a, req, opts)
}

func (a *API) DeleteGuildScheduledEventUser(ctx context.Context, guildID snowflake.ID, guildScheduledEventID snowflake.ID, query discord.DeleteGuildScheduledEventUserQuery, params discord.DeleteGuildScheduledEventUserParams, opts ...Option) ([]discord.GuildScheduledEventUser, error) {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/guilds/{guild_id}/scheduled-events/{guild_scheduled_event_id}/users/@me", rest.P("guild_id", guildID), rest.P("guild_scheduled_event_id", guildScheduledEventID))}
	req.Query = rest.Values(query)
	req.JSON = params
	return list[discord.GuildScheduledEventUser](ctx, a, req, opts)
}

func (a *API) GetGuildTemplate(ctx context.Context, templateCode string, opts ...Option) (*discord.GuildTemplate, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/templates/{template_code}", rest.P("template_code", templateCode))}
	return call[discord.GuildTemplate](ctx, a, req, opts)
}

func (a *API) UseGuildTemplate(ctx context.Context, templateCode string, params discord.UseGuildTemplateParams, opts ...Option) (*discord.Guild, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/guilds/templates/{template_code}", rest.P("template_code", templateCode))}
	req.JSON = params
	return call[discord.Guild](ctx, a, req, opts)
}

func (a *API) ListGuildTemplates(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.GuildTemplate, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/templates", rest.P("guild_id", guildID))}
	return list[discord.GuildTemplate](ctx, a, req, opts)
}

func (a *API) CreateGuildTemplate(ctx context.Context, guildID snowflake.ID, params discord.CreateGuildTemplateParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/guilds/{guild_id}/templates", rest.P("guild_id", guildID))}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) SyncGuildTemplate(ctx context.Context, guildID snowflake.ID, templateCode string, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("PUT", "/guilds/{guild_id}/templates/{template_code}", rest.P("guild_id", guildID), rest.P("template_code", templateCode))}
	return raw(ctx, a, req, opts)
}

func (a *API) ModifyGuildTemplate(ctx context.Context, guildID snowflake.ID, templateCode string, params discord.ModifyGuildTemplateParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/guilds/{guild_id}/templates/{template_code}", rest.P("guild_id", guildID), rest.P("template_code", templateCode))}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) DeleteGuildTemplate(ctx context.Context, guildID snowflake.ID, templateCode string, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/guilds/{guild_id}/templates/{template_code}", rest.P("guild_id", guildID), rest.P("template_code", templateCode))}
	return raw(ctx, a, req, opts)
}

func (a *API) ListUserGuilds(ctx context.Context, query discord.ListUserGuildsQuery, opts ...Option) ([]discord.UserGuild, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/guilds")}
	req.Query = rest.Values(query)
	return list[discord.UserGuild](ctx, a, req, opts)
}

func (a *API) ListJoinRequestGuilds(ctx context.Context, opts ...Option) ([]discord.Guild, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/join-request-guilds")}
	return list[discord.Guild](ctx, a, req, opts)
}

func (a *API) LeaveGuild(ctx context.Context, guildID snowflake.ID, params discord.LeaveGuildParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/users/@me/guilds/{guild_id}", rest.P("guild_id", guildID))}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) CreateGuild(ctx context.Context, params discord.CreateGuildParams, opts ...Option) (*discord.Guild, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/guilds")}
	req.JSON = params
	return call[discord.Guild](ctx, a, req, opts)
}

func (a *API) GetGuild(ctx context.Context, guildID snowflake.ID, query discord.GetGuildQuery, opts ...Option) (*discord.Guild, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}", rest.P("guild_id", guildID))}
	req.Query = rest.Values(query)
	return call[discord.Guild](ctx, a, req, opts)
}

func (a *API) GetGuildBasic(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.Guild, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/basic", rest.P("guild_id", guildID))}
	return call[discord.Guild](ctx, a, req, opts)
}

func (a *API) GetGuildPreview(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.Guild, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/preview", rest.P("guild_id", guildID))}
	return call[discord.Guild](ctx, a, req, opts)
}

func (a *API) ModifyGuild(ctx context.Context, guildID snowflake.ID, params discord.ModifyGuildParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/guilds/{guild_id}", rest.P("guild_id", guildID))}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) ModifyGuildMFALevel(ctx context.Context, guildID snowflake.ID, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/guilds/{guild_id}/mfa", rest.P("guild_id", guildID))}
	return raw(ctx, a, req, opts)
}

func (a *API) GetGuildOwnershipTransferCode(ctx context.Context, guildID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("PUT", "/guilds/{guild_id}/pincode", rest.P("guild_id", guildID))}
	return none(ctx, a, req, opts)
}

func (a *API) DeleteGuild(ctx context.Context, guildID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/guilds/{guild_id}", rest.P("guild_id", guildID))}
	return none(ctx, a, req, opts)
}

func (a *API) ListGuildMembers(ctx context.Context, guildID snowflake.ID, query discord.ListGuildMembersQuery, opts ...Option) ([]discord.GuildMember, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/members", rest.P("guild_id", guildID))}
	req.Query = rest.Values(query)
	return list[discord.GuildMember](ctx, a, req, opts)
}

func (a *API) QueryGuildMembers(ctx context.Context, guildID snowflake.ID, query discord.QueryGuildMembersQuery, opts ...Option) ([]discord.GuildMember, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/members/search", rest.P("guild_id", guildID))}
	req.Query = rest.Values(query)
	return list[discord.GuildMember](ctx, a, req, opts)
}

func (a *API) SearchGuildMembers(ctx context.Context, guildID snowflake.ID, params discord.SearchGuildMembersParams, opts ...Option) (*discord.SearchGuildMembersResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/guilds/{guild_id}/members-search", rest.P("guild_id", guildID))}
	req.JSON = params
	return call[discord.SearchGuildMembersResponse](ctx, a, req, opts)
}

func (a *API) ListGuildMembersSupplemental(ctx context.Context, guildID snowflake.ID, params discord.ListGuildMembersSupplementalParams, opts ...Option) ([]discord.SupplementalGuildMember, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/guilds/{guild_id}/members/supplemental", rest.P("guild_id", guildID))}
	req.JSON = params
	return list[discord.SupplementalGuildMember](ctx, a, req, opts)
}

func (a *API) ListGuildMembersWithUnusualDMActivity(ctx context.Context, guildID snowflake.ID, query discord.ListGuildMembersWithUnusualDMActivityQuery, opts ...Option) ([]discord.GuildMemberUnusualDMActivity, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/members/unusual-dm-activity", rest.P("guild_id", guildID))}
	req.Query = rest.Values(query)
	return list[discord.GuildMemberUnusualDMActivity](ctx, a, req, opts)
}

func (a *API) GetCurrentGuildMember(ctx context.Context, guildID snowflake.ID, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/guilds/{guild_id}/member", rest.P("guild_id", guildID))}
	return raw(ctx, a, req, opts)
}

func (a *API) GetGuildMember(ctx context.Context, guildID snowflake.ID, userID snowflake.ID, opts ...Option) (*discord.GuildMember, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/members/{user_id}", rest.P("guild_id", guildID), rest.P("user_id", userID))}
	return call[discord.GuildMember](ctx, a, req, opts)
}

func (a *API) JoinGuild(ctx context.Context, guildID snowflake.ID, query discord.JoinGuildQuery, opts ...Option) (*discord.Guild, error) {
	req := &rest.Request{Route: rest.NewRoute("PUT", "/guilds/{guild_id}/members/@me", rest.P("guild_id", guildID))}
	req.Query = rest.Values(query)
	return call[discord.Guild](ctx, a, req, opts)
}

func (a *API) AddGuildMember(ctx context.Context, guildID snowflake.ID, userID snowflake.ID, params discord.AddGuildMemberParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("PUT", "/guilds/{guild_id}/members/{user_id}", rest.P("guild_id", guildID), rest.P("user_id", userID))}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) ModifyGuildMember(ctx context.Context, guildID snowflake.ID, userID snowflake.ID, params discord.ModifyGuildMemberParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/guilds/{guild_id}/members/{user_id}", rest.P("guild_id", guildID), rest.P("user_id", userID))}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) ModifyCurrentGuildMember(ctx context.Context, guildID snowflake.ID, params discord.ModifyCurrentGuildMemberParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/guilds/{guild_id}/members/@me", rest.P("guild_id", guildID))}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) ModifyCurrentGuildMemberNick(ctx context.Context, guildID snowflake.ID, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/guilds/{guild_id}/members/@me/nick", rest.P("guild_id", guildID))}
	return raw(ctx, a, req, opts)
}

func (a *API) ModifyGuildMemberProfile(ctx context.Context, guildID snowflake.ID, params discord.ModifyGuildMemberProfileParams, opts ...Option) (*discord.ProfileMetadata, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/guilds/{guild_id}/profile/@me", rest.P("guild_id", guildID))}
	req.JSON = params
	return call[discord.ProfileMetadata](ctx, a, req, opts)
}

func (a *API) AddGuildMemberRole(ctx context.Context, guildID snowflake.ID, userID snowflake.ID, roleID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("PUT", "/guilds/{guild_id}/members/{user_id}/roles/{role_id}", rest.P("guild_id", guildID), rest.P("user_id", userID), rest.P("role_id", roleID))}
	return none(ctx, a, req, opts)
}

func (a *API) RemoveGuildMemberRole(ctx context.Context, guildID snowflake.ID, userID snowflake.ID, roleID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/guilds/{guild_id}/members/{user_id}/roles/{role_id}", rest.P("guild_id", guildID), rest.P("user_id", userID), rest.P("role_id", roleID))}
	return none(ctx, a, req, opts)
}

func (a *API) RemoveGuildMember(ctx context.Context, guildID snowflake.ID, userID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/guilds/{guild_id}/members/{user_id}", rest.P("guild_id", guildID), rest.P("user_id", userID))}
	return none(ctx, a, req, opts)
}

func (a *API) AcknowledgeDMSettingsUpsellModal(ctx context.Context, guildID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/guilds/{guild_id}/member/ack-dm-upsell-settings", rest.P("guild_id", guildID))}
	return none(ctx, a, req, opts)
}

func (a *API) ListGuildBans(ctx context.Context, guildID snowflake.ID, query discord.ListGuildBansQuery, opts ...Option) ([]discord.Ban, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/bans", rest.P("guild_id", guildID))}
	req.Query = rest.Values(query)
	return list[discord.Ban](ctx, a, req, opts)
}

func (a *API) SearchGuildBans(ctx context.Context, guildID snowflake.ID, query discord.SearchGuildBansQuery, opts ...Option) ([]discord.Ban, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/bans/search", rest.P("guild_id", guildID))}
	req.Query = rest.Values(query)
	return list[discord.Ban](ctx, a, req, opts)
}

func (a *API) GetGuildBan(ctx context.Context, guildID snowflake.ID, userID snowflake.ID, opts ...Option) (*discord.Ban, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/bans/{user_id}", rest.P("guild_id", guildID), rest.P("user_id", userID))}
	return call[discord.Ban](ctx, a, req, opts)
}

func (a *API) CreateGuildBan(ctx context.Context, guildID snowflake.ID, userID snowflake.ID, params discord.CreateGuildBanParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("PUT", "/guilds/{guild_id}/bans/{user_id}", rest.P("guild_id", guildID), rest.P("user_id", userID))}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) BulkGuildBan(ctx context.Context, guildID snowflake.ID, params discord.BulkGuildBanParams, opts ...Option) (*discord.BulkGuildBanResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/guilds/{guild_id}/bulk-ban", rest.P("guild_id", guildID))}
	req.JSON = params
	return call[discord.BulkGuildBanResponse](ctx, a, req, opts)
}

func (a *API) BulkGuildBanV2(ctx context.Context, guildID snowflake.ID, params discord.BulkGuildBanV2Params, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/guilds/{guild_id}/bulk-ban/v2", rest.P("guild_id", guildID))}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) DeleteGuildBan(ctx context.Context, guildID snowflake.ID, userID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/guilds/{guild_id}/bans/{user_id}", rest.P("guild_id", guildID), rest.P("user_id", userID))}
	return none(ctx, a, req, opts)
}

func (a *API) ListGuildRoles(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.Role, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/roles", rest.P("guild_id", guildID))}
	return list[discord.Role](ctx, a, req, opts)
}

func (a *API) GetGuildRole(ctx context.Context, guildID snowflake.ID, roleID snowflake.ID, opts ...Option) (*discord.Role, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/roles/{role_id}", rest.P("guild_id", guildID), rest.P("role_id", roleID))}
	return call[discord.Role](ctx, a, req, opts)
}

func (a *API) GetGuildRoleMemberCounts(ctx context.Context, guildID snowflake.ID, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/roles/member-counts", rest.P("guild_id", guildID))}
	return raw(ctx, a, req, opts)
}

func (a *API) ListGuildRoleMembers(ctx context.Context, guildID snowflake.ID, roleID snowflake.ID, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/roles/{role_id}/member-ids", rest.P("guild_id", guildID), rest.P("role_id", roleID))}
	return raw(ctx, a, req, opts)
}

func (a *API) AddGuildRoleMembers(ctx context.Context, guildID snowflake.ID, roleID snowflake.ID, params discord.AddGuildRoleMembersParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/guilds/{guild_id}/roles/{role_id}/members", rest.P("guild_id", guildID), rest.P("role_id", roleID))}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) ListGuildRoleConnectionsConfigurations(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.RoleConnectionRule, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/roles/connections-configurations", rest.P("guild_id", guildID))}
	return list[discord.RoleConnectionRule](ctx, a, req, opts)
}

func (a *API) GetGuildRoleConnectionConfiguration(ctx context.Context, guildID snowflake.ID, roleID snowflake.ID, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/roles/{role_id}/connections/configuration", rest.P("guild_id", guildID), rest.P("role_id", roleID))}
	return raw(ctx, a, req, opts)
}

func (a *API) ModifyGuildRoleConnectionConfiguration(ctx context.Context, guildID snowflake.ID, roleID snowflake.ID, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("PUT", "/guilds/{guild_id}/roles/{role_id}/connections/configuration", rest.P("guild_id", guildID), rest.P("role_id", roleID))}
	return raw(ctx, a, req, opts)
}

func (a *API) GetGuildRoleConnectionEligibility(ctx context.Context, guildID snowflake.ID, roleID snowflake.ID, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/roles/{role_id}/connections/eligibility", rest.P("guild_id", guildID), rest.P("role_id", roleID))}
	return raw(ctx, a, req, opts)
}

func (a *API) AssignGuildRoleConnection(ctx context.Context, guildID snowflake.ID, roleID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("PUT", "/guilds/{guild_id}/roles/{role_id}/connections/assign", rest.P("guild_id", guildID), rest.P("role_id", roleID))}
	return none(ctx, a, req, opts)
}

func (a *API) UnassignGuildRoleConnection(ctx context.Context, guildID snowflake.ID, roleID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/guilds/{guild_id}/roles/{role_id}/connections/unassign", rest.P("guild_id", guildID), rest.P("role_id", roleID))}
	return none(ctx, a, req, opts)
}

func (a *API) CreateGuildRole(ctx context.Context, guildID snowflake.ID, params discord.CreateGuildRoleParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/guilds/{guild_id}/roles", rest.P("guild_id", guildID))}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) ModifyGuildRolePositions(ctx context.Context, guildID snowflake.ID, params discord.ModifyGuildRolePositionsParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/guilds/{guild_id}/roles", rest.P("guild_id", guildID))}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) ModifyGuildRole(ctx context.Context, guildID snowflake.ID, roleID snowflake.ID, params discord.ModifyGuildRoleParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/guilds/{guild_id}/roles/{role_id}", rest.P("guild_id", guildID), rest.P("role_id", roleID))}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) DeleteGuildRole(ctx context.Context, guildID snowflake.ID, roleID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/guilds/{guild_id}/roles/{role_id}", rest.P("guild_id", guildID), rest.P("role_id", roleID))}
	return none(ctx, a, req, opts)
}

func (a *API) GetGuildPrune(ctx context.Context, guildID snowflake.ID, query discord.GetGuildPruneQuery, opts ...Option) (*discord.GetGuildPruneResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/prune", rest.P("guild_id", guildID))}
	req.Query = rest.Values(query)
	return call[discord.GetGuildPruneResponse](ctx, a, req, opts)
}

func (a *API) PruneGuild(ctx context.Context, guildID snowflake.ID, params discord.PruneGuildParams, opts ...Option) (*discord.PruneGuildResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/guilds/{guild_id}/prune", rest.P("guild_id", guildID))}
	req.JSON = params
	return call[discord.PruneGuildResponse](ctx, a, req, opts)
}

func (a *API) GetGuildWidgetSettings(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.GuildWidgetSettings, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/widget", rest.P("guild_id", guildID))}
	return call[discord.GuildWidgetSettings](ctx, a, req, opts)
}

func (a *API) ModifyGuildWidget(ctx context.Context, guildID snowflake.ID, params discord.ModifyGuildWidgetParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/guilds/{guild_id}/widget", rest.P("guild_id", guildID))}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) GetGuildWidget(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.GuildWidget, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/widget.json", rest.P("guild_id", guildID))}
	return call[discord.GuildWidget](ctx, a, req, opts)
}

func (a *API) GetGuildWidgetImage(ctx context.Context, guildID snowflake.ID, query discord.GetGuildWidgetImageQuery, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/widget.png", rest.P("guild_id", guildID))}
	req.Query = rest.Values(query)
	return raw(ctx, a, req, opts)
}

func (a *API) GetGuildVanityInvite(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.GetGuildVanityInviteResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/vanity-url", rest.P("guild_id", guildID))}
	return call[discord.GetGuildVanityInviteResponse](ctx, a, req, opts)
}

func (a *API) ModifyGuildVanityInvite(ctx context.Context, guildID snowflake.ID, params discord.ModifyGuildVanityInviteParams, opts ...Option) (*discord.ModifyGuildVanityInviteResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/guilds/{guild_id}/vanity-url", rest.P("guild_id", guildID))}
	req.JSON = params
	return call[discord.ModifyGuildVanityInviteResponse](ctx, a, req, opts)
}

func (a *API) GetGuildMemberVerification(ctx context.Context, guildID snowflake.ID, query discord.GetGuildMemberVerificationQuery, opts ...Option) (*discord.MemberVerification, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/member-verification", rest.P("guild_id", guildID))}
	req.Query = rest.Values(query)
	return call[discord.MemberVerification](ctx, a, req, opts)
}

func (a *API) ModifyGuildMemberVerification(ctx context.Context, guildID snowflake.ID, params discord.ModifyGuildMemberVerificationParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/guilds/{guild_id}/member-verification", rest.P("guild_id", guildID))}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) ListGuildJoinRequests(ctx context.Context, guildID snowflake.ID, query discord.ListGuildJoinRequestsQuery, opts ...Option) (*discord.ListGuildJoinRequestsResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/requests", rest.P("guild_id", guildID))}
	req.Query = rest.Values(query)
	return call[discord.ListGuildJoinRequestsResponse](ctx, a, req, opts)
}

func (a *API) GetGuildJoinRequest(ctx context.Context, guildJoinRequestID snowflake.ID, opts ...Option) (*discord.GuildJoinRequest, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/join-requests/{guild_join_request_id}", rest.P("guild_join_request_id", guildJoinRequestID))}
	return call[discord.GuildJoinRequest](ctx, a, req, opts)
}

func (a *API) GetCurrentUserGuildJoinRequest(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.GuildJoinRequest, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/requests/@me", rest.P("guild_id", guildID))}
	return call[discord.GuildJoinRequest](ctx, a, req, opts)
}

func (a *API) GetGuildJoinRequestCooldown(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.GetGuildJoinRequestCooldownResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/requests/@me/cooldown", rest.P("guild_id", guildID))}
	return call[discord.GetGuildJoinRequestCooldownResponse](ctx, a, req, opts)
}

func (a *API) CreateGuildJoinRequest(ctx context.Context, guildID snowflake.ID, params discord.CreateGuildJoinRequestParams, opts ...Option) (*discord.GuildJoinRequest, error) {
	req := &rest.Request{Route: rest.NewRoute("PUT", "/guilds/{guild_id}/requests/@me", rest.P("guild_id", guildID))}
	req.JSON = params
	return call[discord.GuildJoinRequest](ctx, a, req, opts)
}

func (a *API) ResetGuildJoinRequest(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.GuildJoinRequest, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/guilds/{guild_id}/requests/@me", rest.P("guild_id", guildID))}
	return call[discord.GuildJoinRequest](ctx, a, req, opts)
}

func (a *API) AcknowledgeGuildJoinRequest(ctx context.Context, guildID snowflake.ID, guildJoinRequestID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/guilds/{guild_id}/requests/{guild_join_request_id}/ack", rest.P("guild_id", guildID), rest.P("guild_join_request_id", guildJoinRequestID))}
	return none(ctx, a, req, opts)
}

func (a *API) DeleteGuildJoinRequest(ctx context.Context, guildID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/guilds/{guild_id}/requests/@me", rest.P("guild_id", guildID))}
	return none(ctx, a, req, opts)
}

func (a *API) CreateGuildJoinRequestInterview(ctx context.Context, guildJoinRequestID snowflake.ID, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/join-requests/{guild_join_request_id}/interview", rest.P("guild_join_request_id", guildJoinRequestID))}
	return raw(ctx, a, req, opts)
}

func (a *API) ActionGuildJoinRequest(ctx context.Context, guildID snowflake.ID, guildJoinRequestID snowflake.ID, params discord.ActionGuildJoinRequestParams, opts ...Option) (*discord.GuildJoinRequest, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/guilds/{guild_id}/requests/{guild_join_request_id}", rest.P("guild_id", guildID), rest.P("guild_join_request_id", guildJoinRequestID))}
	req.JSON = params
	return call[discord.GuildJoinRequest](ctx, a, req, opts)
}

func (a *API) ActionGuildJoinRequestByID(ctx context.Context, guildID snowflake.ID, guildJoinRequestID snowflake.ID, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/guilds/{guild_id}/requests/id/{guild_join_request_id}", rest.P("guild_id", guildID), rest.P("guild_join_request_id", guildJoinRequestID))}
	return raw(ctx, a, req, opts)
}

func (a *API) BulkActionGuildJoinRequests(ctx context.Context, guildID snowflake.ID, params discord.BulkActionGuildJoinRequestsParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/guilds/{guild_id}/requests", rest.P("guild_id", guildID))}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) ListUserGuildJoinRequests(ctx context.Context, guildID snowflake.ID, userID snowflake.ID, opts ...Option) ([]discord.GuildJoinRequest, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/requests/users/{user_id}", rest.P("guild_id", guildID), rest.P("user_id", userID))}
	return list[discord.GuildJoinRequest](ctx, a, req, opts)
}

func (a *API) GetGuildWelcomeScreen(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.WelcomeScreen, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/welcome-screen", rest.P("guild_id", guildID))}
	return call[discord.WelcomeScreen](ctx, a, req, opts)
}

func (a *API) ModifyGuildWelcomeScreen(ctx context.Context, guildID snowflake.ID, params discord.ModifyGuildWelcomeScreenParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/guilds/{guild_id}/welcome-screen", rest.P("guild_id", guildID))}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) GetGuildOnboarding(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.Onboarding, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/onboarding", rest.P("guild_id", guildID))}
	return call[discord.Onboarding](ctx, a, req, opts)
}

func (a *API) ModifyGuildOnboarding(ctx context.Context, guildID snowflake.ID, params discord.ModifyGuildOnboardingParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("PUT", "/guilds/{guild_id}/onboarding", rest.P("guild_id", guildID))}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) ListGuildOnboardingAllowedApplications(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.ListGuildOnboardingAllowedApplicationsResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/onboarding/allowed-applications", rest.P("guild_id", guildID))}
	return call[discord.ListGuildOnboardingAllowedApplicationsResponse](ctx, a, req, opts)
}

func (a *API) CreateGuildOnboardingResponses(ctx context.Context, guildID snowflake.ID, params discord.CreateGuildOnboardingResponsesParams, opts ...Option) (*discord.OnboardingResponses, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/guilds/{guild_id}/onboarding-responses", rest.P("guild_id", guildID))}
	req.JSON = params
	return call[discord.OnboardingResponses](ctx, a, req, opts)
}

func (a *API) ModifyGuildOnboardingResponses(ctx context.Context, guildID snowflake.ID, params discord.ModifyGuildOnboardingResponsesParams, opts ...Option) (*discord.OnboardingResponses, error) {
	req := &rest.Request{Route: rest.NewRoute("PUT", "/guilds/{guild_id}/onboarding-responses", rest.P("guild_id", guildID))}
	req.JSON = params
	return call[discord.OnboardingResponses](ctx, a, req, opts)
}

func (a *API) GetGuildNewMemberWelcome(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.NewMemberWelcome, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/new-member-welcome", rest.P("guild_id", guildID))}
	return call[discord.NewMemberWelcome](ctx, a, req, opts)
}

func (a *API) ModifyGuildNewMemberWelcome(ctx context.Context, guildID snowflake.ID, params discord.ModifyGuildNewMemberWelcomeParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("PUT", "/guilds/{guild_id}/new-member-welcome", rest.P("guild_id", guildID))}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) ModifyGuildNewMemberAction(ctx context.Context, guildID snowflake.ID, channelID snowflake.ID, params discord.ModifyGuildNewMemberActionParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/guilds/{guild_id}/new-member-actions/{channel_id}", rest.P("guild_id", guildID), rest.P("channel_id", channelID))}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) ModifyGuildResourceChannel(ctx context.Context, guildID snowflake.ID, channelID snowflake.ID, params discord.ModifyGuildResourceChannelParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("PUT", "/guilds/{guild_id}/resource-channels/{channel_id}", rest.P("guild_id", guildID), rest.P("channel_id", channelID))}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) GetGuildNewMemberActions(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.NewMemberActionsProgress, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/new-member-actions", rest.P("guild_id", guildID))}
	return call[discord.NewMemberActionsProgress](ctx, a, req, opts)
}

func (a *API) CompleteGuildNewMemberAction(ctx context.Context, guildID snowflake.ID, channelID snowflake.ID, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/guilds/{guild_id}/new-member-action/{channel_id}", rest.P("guild_id", guildID), rest.P("channel_id", channelID))}
	return raw(ctx, a, req, opts)
}

func (a *API) ListGuildTopGames(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.ListGuildTopGamesResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/top-games", rest.P("guild_id", guildID))}
	return call[discord.ListGuildTopGamesResponse](ctx, a, req, opts)
}

func (a *API) ListPremiumGuildSubscriptions(ctx context.Context, guildID snowflake.ID, query discord.ListPremiumGuildSubscriptionsQuery, opts ...Option) ([]discord.PremiumGuildSubscription, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/premium/subscriptions", rest.P("guild_id", guildID))}
	req.Query = rest.Values(query)
	return list[discord.PremiumGuildSubscription](ctx, a, req, opts)
}

func (a *API) CreatePremiumGuildSubscriptions(ctx context.Context, guildID snowflake.ID, params discord.CreatePremiumGuildSubscriptionsParams, opts ...Option) ([]discord.PremiumGuildSubscription, error) {
	req := &rest.Request{Route: rest.NewRoute("PUT", "/guilds/{guild_id}/premium/subscriptions", rest.P("guild_id", guildID))}
	req.JSON = params
	return list[discord.PremiumGuildSubscription](ctx, a, req, opts)
}

func (a *API) DeletePremiumGuildSubscription(ctx context.Context, guildID snowflake.ID, premiumGuildSubscriptionID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/guilds/{guild_id}/premium/subscriptions/{premium_guild_subscription_id}", rest.P("guild_id", guildID), rest.P("premium_guild_subscription_id", premiumGuildSubscriptionID))}
	return none(ctx, a, req, opts)
}

func (a *API) ListGuildPowerups(ctx context.Context, guildID snowflake.ID, query discord.ListGuildPowerupsQuery, opts ...Option) ([]discord.Entitlement, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/powerups", rest.P("guild_id", guildID))}
	req.Query = rest.Values(query)
	return list[discord.Entitlement](ctx, a, req, opts)
}

func (a *API) AddGuildPowerup(ctx context.Context, guildID snowflake.ID, skuid snowflake.ID, params discord.AddGuildPowerupParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/guilds/{guild_id}/skus/{sku_id}", rest.P("guild_id", guildID), rest.P("sku_id", skuid))}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) ModifyGuildPowerup(ctx context.Context, guildID snowflake.ID, entitlementID snowflake.ID, params discord.ModifyGuildPowerupParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/guilds/{guild_id}/entitlements/{entitlement_id}", rest.P("guild_id", guildID), rest.P("entitlement_id", entitlementID))}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) RemoveGuildPowerup(ctx context.Context, guildID snowflake.ID, skuid snowflake.ID, query discord.RemoveGuildPowerupQuery, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/guilds/{guild_id}/skus/{sku_id}", rest.P("guild_id", guildID), rest.P("sku_id", skuid))}
	req.Query = rest.Values(query)
	return none(ctx, a, req, opts)
}

func (a *API) ListGuildGameServers(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.GameServer, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/game-servers", rest.P("guild_id", guildID))}
	return list[discord.GameServer](ctx, a, req, opts)
}

func (a *API) ListGuildGameServerRegions(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.GameServerRegion, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/game-server-regions", rest.P("guild_id", guildID))}
	return list[discord.GameServerRegion](ctx, a, req, opts)
}

func (a *API) WakeGuildGameServer(ctx context.Context, guildID snowflake.ID, gameServerID snowflake.ID, opts ...Option) (*discord.GameServer, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/guilds/{guild_id}/game-servers/{game_server_id}/wake", rest.P("guild_id", guildID), rest.P("game_server_id", gameServerID))}
	return call[discord.GameServer](ctx, a, req, opts)
}

func (a *API) GetAdminCommunityEligibility(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.GetAdminCommunityEligibilityResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/admin-server-eligibility", rest.P("guild_id", guildID))}
	return call[discord.GetAdminCommunityEligibilityResponse](ctx, a, req, opts)
}

func (a *API) JoinAdminCommunity(ctx context.Context, guildID snowflake.ID, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/guilds/{guild_id}/join-admin-server", rest.P("guild_id", guildID))}
	return raw(ctx, a, req, opts)
}

func (a *API) MigratePinPermission(ctx context.Context, guildID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/guilds/{guild_id}/migrate-pin-permission", rest.P("guild_id", guildID))}
	return none(ctx, a, req, opts)
}

func (a *API) MigrateBypassSlowmodePermission(ctx context.Context, guildID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/guilds/{guild_id}/migrate-bypass-slowmode-permission", rest.P("guild_id", guildID))}
	return none(ctx, a, req, opts)
}

func (a *API) QueryStudentHubs(ctx context.Context, params discord.QueryStudentHubsParams, opts ...Option) (*discord.QueryStudentHubsResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/guilds/automations/email-domain-lookup")}
	req.JSON = params
	return call[discord.QueryStudentHubsResponse](ctx, a, req, opts)
}

func (a *API) JoinStudentHub(ctx context.Context, params discord.JoinStudentHubParams, opts ...Option) (*discord.JoinStudentHubResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/guilds/automations/email-domain-lookup/verify-code")}
	req.JSON = params
	return call[discord.JoinStudentHubResponse](ctx, a, req, opts)
}

func (a *API) JoinStudentHubWaitlist(ctx context.Context, params discord.JoinStudentHubWaitlistParams, opts ...Option) (*discord.JoinStudentHubWaitlistResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/hub-waitlist/signup")}
	req.JSON = params
	return call[discord.JoinStudentHubWaitlistResponse](ctx, a, req, opts)
}

func (a *API) GetGuildRoleSubscriptionsSettings(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.RoleSubscriptionSettings, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/role-subscriptions/settings", rest.P("guild_id", guildID))}
	return call[discord.RoleSubscriptionSettings](ctx, a, req, opts)
}

func (a *API) ModifyGuildRoleSubscriptionsSettings(ctx context.Context, guildID snowflake.ID, params discord.ModifyGuildRoleSubscriptionsSettingsParams, opts ...Option) (*discord.RoleSubscriptionSettings, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/guilds/{guild_id}/role-subscriptions/settings", rest.P("guild_id", guildID))}
	req.JSON = params
	return call[discord.RoleSubscriptionSettings](ctx, a, req, opts)
}

func (a *API) CreateGuildRoleSubscriptionGroupListing(ctx context.Context, guildID snowflake.ID, params discord.CreateGuildRoleSubscriptionGroupListingParams, opts ...Option) (*discord.RoleSubscriptionGroupListing, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/guilds/{guild_id}/role-subscriptions/group-listings", rest.P("guild_id", guildID))}
	req.JSON = params
	return call[discord.RoleSubscriptionGroupListing](ctx, a, req, opts)
}

func (a *API) GetGuildRoleSubscriptionGroupListing(ctx context.Context, guildID snowflake.ID, subscriptionGroupListingID snowflake.ID, opts ...Option) (*discord.RoleSubscriptionGroupListing, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/role-subscriptions/group-listings/{subscription_group_listing_id}", rest.P("guild_id", guildID), rest.P("subscription_group_listing_id", subscriptionGroupListingID))}
	return call[discord.RoleSubscriptionGroupListing](ctx, a, req, opts)
}

func (a *API) ListGuildRoleSubscriptionGroupListings(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.RoleSubscriptionGroupListing, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/role-subscriptions/group-listings", rest.P("guild_id", guildID))}
	return list[discord.RoleSubscriptionGroupListing](ctx, a, req, opts)
}

func (a *API) ModifyGuildRoleSubscriptionGroupListing(ctx context.Context, guildID snowflake.ID, subscriptionGroupListingID snowflake.ID, params discord.ModifyGuildRoleSubscriptionGroupListingParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/guilds/{guild_id}/role-subscriptions/group-listings/{subscription_group_listing_id}", rest.P("guild_id", guildID), rest.P("subscription_group_listing_id", subscriptionGroupListingID))}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) DeleteGuildRoleSubscriptionGroupListing(ctx context.Context, guildID snowflake.ID, subscriptionGroupListingID snowflake.ID, params discord.DeleteGuildRoleSubscriptionGroupListingParams, opts ...Option) (*discord.RoleSubscriptionListing, error) {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/guilds/{guild_id}/role-subscriptions/group-listings/{subscription_group_listing_id}", rest.P("guild_id", guildID), rest.P("subscription_group_listing_id", subscriptionGroupListingID))}
	req.JSON = params
	return call[discord.RoleSubscriptionListing](ctx, a, req, opts)
}

func (a *API) ListGuildRoleSubscriptionTrials(ctx context.Context, guildID snowflake.ID, params discord.ListGuildRoleSubscriptionTrialsParams, opts ...Option) ([]discord.RoleSubscriptionTrial, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/role-subscriptions/trials", rest.P("guild_id", guildID))}
	req.JSON = params
	return list[discord.RoleSubscriptionTrial](ctx, a, req, opts)
}

func (a *API) ListGuildRoleSubscriptionListingTemplates(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.ListGuildRoleSubscriptionListingTemplatesResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/role-subscriptions/templates", rest.P("guild_id", guildID))}
	return call[discord.ListGuildRoleSubscriptionListingTemplatesResponse](ctx, a, req, opts)
}

func (a *API) GetCreatorMonetizationEligibility(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.GetCreatorMonetizationEligibilityResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/creator-monetization/requirements", rest.P("guild_id", guildID))}
	return call[discord.GetCreatorMonetizationEligibilityResponse](ctx, a, req, opts)
}

func (a *API) CreateCreatorMonetizationEnableRequest(ctx context.Context, guildID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/guilds/{guild_id}/creator-monetization/enable-requests", rest.P("guild_id", guildID))}
	return none(ctx, a, req, opts)
}

func (a *API) AcceptCreatorMonetizationTerms(ctx context.Context, guildID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/guilds/{guild_id}/creator-monetization/accept-terms", rest.P("guild_id", guildID))}
	return none(ctx, a, req, opts)
}

func (a *API) GetCreatorMonetizationRestrictions(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.GetCreatorMonetizationRestrictionsResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/creator-monetization/restrictions", rest.P("guild_id", guildID))}
	return call[discord.GetCreatorMonetizationRestrictionsResponse](ctx, a, req, opts)
}

func (a *API) GetCreatorMonetizationMarketingOnboarding(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.GetCreatorMonetizationMarketingOnboardingResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/creator-monetization/marketing/onboarding", rest.P("guild_id", guildID))}
	return call[discord.GetCreatorMonetizationMarketingOnboardingResponse](ctx, a, req, opts)
}

func (a *API) UpdateCreatorMonetizationTeam(ctx context.Context, guildID snowflake.ID, params discord.UpdateCreatorMonetizationTeamParams, opts ...Option) (*discord.UpdateCreatorMonetizationTeamResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/guilds/{guild_id}/creator-monetization/ownership-transfer/onboarding", rest.P("guild_id", guildID))}
	req.JSON = params
	return call[discord.UpdateCreatorMonetizationTeamResponse](ctx, a, req, opts)
}

func (a *API) AcceptCreatorMonetizationNewTermsDemonetized(ctx context.Context, guildID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/guilds/{guild_id}/creator-monetization/accept-new-terms-demonetized", rest.P("guild_id", guildID))}
	return none(ctx, a, req, opts)
}

func (a *API) RemoveCreatorMonetization(ctx context.Context, guildID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/guilds/{guild_id}/creator-monetization/remove-monetization", rest.P("guild_id", guildID))}
	return none(ctx, a, req, opts)
}

func (a *API) CreateGuildProductListing(ctx context.Context, guildID snowflake.ID, params discord.CreateGuildProductListingParams, opts ...Option) (*discord.GuildProductListing, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/guilds/{guild_id}/products", rest.P("guild_id", guildID))}
	req.JSON = params
	return call[discord.GuildProductListing](ctx, a, req, opts)
}

func (a *API) ListGuildProductListings(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.ListGuildProductListingsResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/products/listings", rest.P("guild_id", guildID))}
	return call[discord.ListGuildProductListingsResponse](ctx, a, req, opts)
}

func (a *API) GetGuildProductListing(ctx context.Context, guildID snowflake.ID, productListingID snowflake.ID, opts ...Option) (*discord.GuildProductListing, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/products/listings/{product_listing_id}", rest.P("guild_id", guildID), rest.P("product_listing_id", productListingID))}
	return call[discord.GuildProductListing](ctx, a, req, opts)
}

func (a *API) ModifyGuildProductListing(ctx context.Context, guildID snowflake.ID, productListingID snowflake.ID, params discord.ModifyGuildProductListingParams, opts ...Option) (*discord.GuildProductListing, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/guilds/{guild_id}/products/listings/{product_listing_id}", rest.P("guild_id", guildID), rest.P("product_listing_id", productListingID))}
	req.JSON = params
	return call[discord.GuildProductListing](ctx, a, req, opts)
}

func (a *API) DeleteGuildProductListing(ctx context.Context, guildID snowflake.ID, productListingID snowflake.ID, opts ...Option) (*discord.DeleteGuildProductListingResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/guilds/{guild_id}/products/listings/{product_listing_id}", rest.P("guild_id", guildID), rest.P("product_listing_id", productListingID))}
	return call[discord.DeleteGuildProductListingResponse](ctx, a, req, opts)
}

func (a *API) ListGuildIntegrations(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.Integration, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/integrations", rest.P("guild_id", guildID))}
	return list[discord.Integration](ctx, a, req, opts)
}

func (a *API) CreateGuildIntegration(ctx context.Context, guildID snowflake.ID, params discord.CreateGuildIntegrationParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/guilds/{guild_id}/integrations", rest.P("guild_id", guildID))}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) SyncGuildIntegration(ctx context.Context, guildID snowflake.ID, integrationID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/guilds/{guild_id}/integrations/{integration_id}/sync", rest.P("guild_id", guildID), rest.P("integration_id", integrationID))}
	return none(ctx, a, req, opts)
}

func (a *API) ModifyGuildIntegration(ctx context.Context, guildID snowflake.ID, integrationID snowflake.ID, params discord.ModifyGuildIntegrationParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/guilds/{guild_id}/integrations/{integration_id}", rest.P("guild_id", guildID), rest.P("integration_id", integrationID))}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) DeleteGuildIntegration(ctx context.Context, guildID snowflake.ID, integrationID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/guilds/{guild_id}/integrations/{integration_id}", rest.P("guild_id", guildID), rest.P("integration_id", integrationID))}
	return none(ctx, a, req, opts)
}

func (a *API) MigrateGuildCommandScope(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.MigrateGuildCommandScopeResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/guilds/{guild_id}/migrate-command-scope", rest.P("guild_id", guildID))}
	return call[discord.MigrateGuildCommandScopeResponse](ctx, a, req, opts)
}

func (a *API) GetGuildIntegrationApplicationIDs(ctx context.Context, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/guilds/integration-application-ids")}
	return raw(ctx, a, req, opts)
}

func (a *API) ListChannelIntegrations(ctx context.Context, channelID snowflake.ID, opts ...Option) ([]discord.Integration, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/channels/{channel_id}/integrations", rest.P("channel_id", channelID))}
	return list[discord.Integration](ctx, a, req, opts)
}

func (a *API) DeleteChannelIntegration(ctx context.Context, channelID snowflake.ID, integrationID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/channels/{channel_id}/integrations/{integration_id}", rest.P("channel_id", channelID), rest.P("integration_id", integrationID))}
	return none(ctx, a, req, opts)
}

func (a *API) JoinIntegrationGuild(ctx context.Context, integrationID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/integrations/{integration_id}/join", rest.P("integration_id", integrationID))}
	return none(ctx, a, req, opts)
}

func (a *API) ListTrendingGIFSearchTerms(ctx context.Context, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/gifs/trending-search")}
	return raw(ctx, a, req, opts)
}

func (a *API) ListSuggestedGIFSearchTerms(ctx context.Context, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/gifs/suggest")}
	return raw(ctx, a, req, opts)
}

func (a *API) SearchGifs(ctx context.Context, opts ...Option) ([]discord.GIF, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/gifs/search")}
	return list[discord.GIF](ctx, a, req, opts)
}

func (a *API) ListTrendingGIFCategories(ctx context.Context, opts ...Option) (*discord.ListTrendingGIFCategoriesResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/gifs/trending")}
	return call[discord.ListTrendingGIFCategoriesResponse](ctx, a, req, opts)
}

func (a *API) ListTrendingGifs(ctx context.Context, opts ...Option) ([]discord.GIF, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/gifs/trending-gifs")}
	return list[discord.GIF](ctx, a, req, opts)
}

func (a *API) TrackSelectedGIF(ctx context.Context, params discord.TrackSelectedGIFParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/gifs/select")}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) GetInvite(ctx context.Context, inviteCode string, query discord.GetInviteQuery, opts ...Option) (*discord.Invite, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/invites/{invite_code}", rest.P("invite_code", inviteCode))}
	req.Query = rest.Values(query)
	return call[discord.Invite](ctx, a, req, opts)
}

func (a *API) ListInviteFriendMembers(ctx context.Context, inviteCode string, opts ...Option) (*discord.ListInviteFriendMembersResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/invites/{invite_code}/friend-members", rest.P("invite_code", inviteCode))}
	return call[discord.ListInviteFriendMembersResponse](ctx, a, req, opts)
}

func (a *API) ListInviteTargetUsers(ctx context.Context, inviteCode string, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/invites/{invite_code}/target-users", rest.P("invite_code", inviteCode))}
	return raw(ctx, a, req, opts)
}

func (a *API) UpdateInviteTargetUsers(ctx context.Context, inviteCode string, params discord.UpdateInviteTargetUsersParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("PUT", "/invites/{invite_code}/target-users", rest.P("invite_code", inviteCode))}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) GetInviteTargetUsersJobStatus(ctx context.Context, inviteCode string, opts ...Option) (*discord.InviteTargetUsersJob, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/invites/{invite_code}/target-users/job-status", rest.P("invite_code", inviteCode))}
	return call[discord.InviteTargetUsersJob](ctx, a, req, opts)
}

func (a *API) AcceptInvite(ctx context.Context, inviteCode string, params discord.AcceptInviteParams, opts ...Option) (*discord.Invite, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/invites/{invite_code}", rest.P("invite_code", inviteCode))}
	req.JSON = params
	return call[discord.Invite](ctx, a, req, opts)
}

func (a *API) DeleteInvite(ctx context.Context, inviteCode string, opts ...Option) (*discord.Invite, error) {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/invites/{invite_code}", rest.P("invite_code", inviteCode))}
	return call[discord.Invite](ctx, a, req, opts)
}

func (a *API) ListGuildInvites(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.Invite, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/invites", rest.P("guild_id", guildID))}
	return list[discord.Invite](ctx, a, req, opts)
}

func (a *API) ListChannelInvites(ctx context.Context, channelID snowflake.ID, opts ...Option) ([]discord.Invite, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/channels/{channel_id}/invites", rest.P("channel_id", channelID))}
	return list[discord.Invite](ctx, a, req, opts)
}

func (a *API) CreateChannelInvite(ctx context.Context, channelID snowflake.ID, params discord.CreateChannelInviteParams, opts ...Option) (*discord.Invite, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/channels/{channel_id}/invites", rest.P("channel_id", channelID))}
	req.JSON = params
	return call[discord.Invite](ctx, a, req, opts)
}

func (a *API) ListUserInvites(ctx context.Context, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/invites")}
	return raw(ctx, a, req, opts)
}

func (a *API) CreateUserInvite(ctx context.Context, params discord.CreateUserInviteParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/invites")}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) RevokeUserInvites(ctx context.Context, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/users/@me/invites")}
	return raw(ctx, a, req, opts)
}

func (a *API) CreateLobby(ctx context.Context, params discord.CreateLobbyParams, opts ...Option) (*discord.Lobby, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/lobbies")}
	req.JSON = params
	return call[discord.Lobby](ctx, a, req, opts)
}

func (a *API) JoinOrCreateLobby(ctx context.Context, params discord.JoinOrCreateLobbyParams, opts ...Option) (*discord.Lobby, error) {
	req := &rest.Request{Route: rest.NewRoute("PUT", "/lobbies")}
	req.JSON = params
	return call[discord.Lobby](ctx, a, req, opts)
}

func (a *API) GetLobby(ctx context.Context, lobbyID snowflake.ID, opts ...Option) (*discord.Lobby, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/lobbies/{lobby_id}", rest.P("lobby_id", lobbyID))}
	return call[discord.Lobby](ctx, a, req, opts)
}

func (a *API) ModifyLobby(ctx context.Context, lobbyID snowflake.ID, params discord.ModifyLobbyParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/lobbies/{lobby_id}", rest.P("lobby_id", lobbyID))}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) DeleteLobby(ctx context.Context, lobbyID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/lobbies/{lobby_id}", rest.P("lobby_id", lobbyID))}
	return none(ctx, a, req, opts)
}

func (a *API) AddLobbyMember(ctx context.Context, lobbyID snowflake.ID, userID snowflake.ID, params discord.AddLobbyMemberParams, opts ...Option) (*discord.LobbyMember, error) {
	req := &rest.Request{Route: rest.NewRoute("PUT", "/lobbies/{lobby_id}/members/{user_id}", rest.P("lobby_id", lobbyID), rest.P("user_id", userID))}
	req.JSON = params
	return call[discord.LobbyMember](ctx, a, req, opts)
}

func (a *API) RemoveLobbyMember(ctx context.Context, lobbyID snowflake.ID, userID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/lobbies/{lobby_id}/members/{user_id}", rest.P("lobby_id", lobbyID), rest.P("user_id", userID))}
	return none(ctx, a, req, opts)
}

func (a *API) LeaveLobby(ctx context.Context, lobbyID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/lobbies/{lobby_id}/members/@me", rest.P("lobby_id", lobbyID))}
	return none(ctx, a, req, opts)
}

func (a *API) BulkUpdateLobbyMembers(ctx context.Context, lobbyID snowflake.ID, opts ...Option) ([]discord.LobbyMember, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/lobbies/{lobby_id}/members/bulk", rest.P("lobby_id", lobbyID))}
	return list[discord.LobbyMember](ctx, a, req, opts)
}

func (a *API) CreateLobbyInviteForCurrentUser(ctx context.Context, lobbyID snowflake.ID, opts ...Option) (*discord.CreateLobbyInviteForCurrentUserResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/lobbies/{lobby_id}/members/@me/invites", rest.P("lobby_id", lobbyID))}
	return call[discord.CreateLobbyInviteForCurrentUserResponse](ctx, a, req, opts)
}

func (a *API) CreateLobbyInvite(ctx context.Context, lobbyID snowflake.ID, userID snowflake.ID, opts ...Option) (*discord.CreateLobbyInviteResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/lobbies/{lobby_id}/members/{user_id}/invites", rest.P("lobby_id", lobbyID), rest.P("user_id", userID))}
	return call[discord.CreateLobbyInviteResponse](ctx, a, req, opts)
}

func (a *API) ModifyLobbyLinkedChannel(ctx context.Context, lobbyID snowflake.ID, params discord.ModifyLobbyLinkedChannelParams, opts ...Option) (*discord.Lobby, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/lobbies/{lobby_id}/channel-linking", rest.P("lobby_id", lobbyID))}
	req.JSON = params
	return call[discord.Lobby](ctx, a, req, opts)
}

func (a *API) ListLobbyMessages(ctx context.Context, lobbyID snowflake.ID, query discord.ListLobbyMessagesQuery, opts ...Option) ([]discord.Message, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/lobbies/{lobby_id}/messages", rest.P("lobby_id", lobbyID))}
	req.Query = rest.Values(query)
	return list[discord.Message](ctx, a, req, opts)
}

func (a *API) CreateLobbyMessage(ctx context.Context, lobbyID snowflake.ID, opts ...Option) (*discord.Message, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/lobbies/{lobby_id}/messages", rest.P("lobby_id", lobbyID))}
	return call[discord.Message](ctx, a, req, opts)
}

func (a *API) UpdateLobbyMessageModerationMetadata(ctx context.Context, lobbyID snowflake.ID, messageID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("PUT", "/lobbies/{lobby_id}/messages/{message_id}/moderation-metadata", rest.P("lobby_id", lobbyID), rest.P("message_id", messageID))}
	return none(ctx, a, req, opts)
}

func (a *API) ListMessages(ctx context.Context, channelID snowflake.ID, query discord.ListMessagesQuery, opts ...Option) ([]discord.Message, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/channels/{channel_id}/messages", rest.P("channel_id", channelID))}
	req.Query = rest.Values(query)
	return list[discord.Message](ctx, a, req, opts)
}

func (a *API) ListDMMessages(ctx context.Context, userID snowflake.ID, query discord.ListDMMessagesQuery, opts ...Option) ([]discord.Message, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/{user_id}/messages", rest.P("user_id", userID))}
	req.Query = rest.Values(query)
	return list[discord.Message](ctx, a, req, opts)
}

func (a *API) PreloadMessages(ctx context.Context, opts ...Option) ([]discord.Message, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/channels/preload-messages")}
	return list[discord.Message](ctx, a, req, opts)
}

func (a *API) SearchGuildMessages(ctx context.Context, guildID snowflake.ID, query discord.SearchGuildMessagesQuery, opts ...Option) (*discord.SearchGuildMessagesResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/messages/search", rest.P("guild_id", guildID))}
	req.Query = rest.Values(query)
	return call[discord.SearchGuildMessagesResponse](ctx, a, req, opts)
}

func (a *API) SearchChannelMessages(ctx context.Context, channelID snowflake.ID, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/channels/{channel_id}/messages/search", rest.P("channel_id", channelID))}
	return raw(ctx, a, req, opts)
}

func (a *API) SearchGuildMessagesByTab(ctx context.Context, guildID snowflake.ID, params discord.SearchGuildMessagesByTabParams, opts ...Option) (*discord.SearchGuildMessagesByTabResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/guilds/{guild_id}/messages/search/tabs", rest.P("guild_id", guildID))}
	req.JSON = params
	return call[discord.SearchGuildMessagesByTabResponse](ctx, a, req, opts)
}

func (a *API) SearchChannelMessagesByTab(ctx context.Context, channelID snowflake.ID, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/channels/{channel_id}/messages/search/tabs", rest.P("channel_id", channelID))}
	return raw(ctx, a, req, opts)
}

func (a *API) SearchUserMessagesByTab(ctx context.Context, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/messages/search/tabs")}
	return raw(ctx, a, req, opts)
}

func (a *API) GetMessage(ctx context.Context, channelID snowflake.ID, messageID snowflake.ID, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/channels/{channel_id}/messages/{message_id}", rest.P("channel_id", channelID), rest.P("message_id", messageID))}
	return raw(ctx, a, req, opts)
}

func (a *API) CreateMessage(ctx context.Context, channelID snowflake.ID, params discord.CreateMessageParams, opts ...Option) (*discord.Message, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/channels/{channel_id}/messages", rest.P("channel_id", channelID))}
	req.JSON = params
	return call[discord.Message](ctx, a, req, opts)
}

func (a *API) CreateGreetMessage(ctx context.Context, channelID snowflake.ID, params discord.CreateGreetMessageParams, opts ...Option) (*discord.Message, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/channels/{channel_id}/greet", rest.P("channel_id", channelID))}
	req.JSON = params
	return call[discord.Message](ctx, a, req, opts)
}

func (a *API) ScanExplicitMedia(ctx context.Context, channelID snowflake.ID, params discord.ScanExplicitMediaParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/channels/{channel_id}/explicit-media", rest.P("channel_id", channelID))}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) BulkScanExplicitMedia(ctx context.Context, params discord.BulkScanExplicitMediaParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/messages/explicit-media")}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) ReportSentExplicitContentFalsePositive(ctx context.Context, params discord.ReportSentExplicitContentFalsePositiveParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/attachments/sender-report-false-positive")}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) ReportExplicitContentFalsePositive(ctx context.Context, params discord.ReportExplicitContentFalsePositiveParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/attachments/report-false-positive")}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) CrosspostMessage(ctx context.Context, channelID snowflake.ID, messageID snowflake.ID, opts ...Option) (*discord.Message, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/channels/{channel_id}/messages/{message_id}/crosspost", rest.P("channel_id", channelID), rest.P("message_id", messageID))}
	return call[discord.Message](ctx, a, req, opts)
}

func (a *API) HideMessageFromGuildFeed(ctx context.Context, channelID snowflake.ID, messageID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/channels/{channel_id}/messages/{message_id}/hide-guild-feed", rest.P("channel_id", channelID), rest.P("message_id", messageID))}
	return none(ctx, a, req, opts)
}

func (a *API) UnhideMessageFromGuildFeed(ctx context.Context, channelID snowflake.ID, messageID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/channels/{channel_id}/messages/{message_id}/hide-guild-feed", rest.P("channel_id", channelID), rest.P("message_id", messageID))}
	return none(ctx, a, req, opts)
}

func (a *API) ListReactions(ctx context.Context, channelID snowflake.ID, messageID snowflake.ID, emoji string, query discord.ListReactionsQuery, opts ...Option) ([]discord.User, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/channels/{channel_id}/messages/{message_id}/reactions/{emoji}", rest.P("channel_id", channelID), rest.P("message_id", messageID), rest.P("emoji", emoji))}
	req.Query = rest.Values(query)
	return list[discord.User](ctx, a, req, opts)
}

func (a *API) CreateReaction(ctx context.Context, channelID snowflake.ID, messageID snowflake.ID, emoji string, query discord.CreateReactionQuery, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("PUT", "/channels/{channel_id}/messages/{message_id}/reactions/{emoji}/@me", rest.P("channel_id", channelID), rest.P("message_id", messageID), rest.P("emoji", emoji))}
	req.Query = rest.Values(query)
	return none(ctx, a, req, opts)
}

func (a *API) DeleteReactionEmoji(ctx context.Context, channelID snowflake.ID, messageID snowflake.ID, emoji string, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/channels/{channel_id}/messages/{message_id}/reactions/{emoji}", rest.P("channel_id", channelID), rest.P("message_id", messageID), rest.P("emoji", emoji))}
	return none(ctx, a, req, opts)
}

func (a *API) DeleteAllReactions(ctx context.Context, channelID snowflake.ID, messageID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/channels/{channel_id}/messages/{message_id}/reactions", rest.P("channel_id", channelID), rest.P("message_id", messageID))}
	return none(ctx, a, req, opts)
}

func (a *API) GetMessageInteractionData(ctx context.Context, channelID snowflake.ID, messageID snowflake.ID, opts ...Option) (*discord.GetMessageInteractionDataResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/channels/{channel_id}/messages/{message_id}/interaction-data", rest.P("channel_id", channelID), rest.P("message_id", messageID))}
	return call[discord.GetMessageInteractionDataResponse](ctx, a, req, opts)
}

func (a *API) EditMessage(ctx context.Context, channelID snowflake.ID, messageID snowflake.ID, params discord.EditMessageParams, opts ...Option) (*discord.Message, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/channels/{channel_id}/messages/{message_id}", rest.P("channel_id", channelID), rest.P("message_id", messageID))}
	req.JSON = params
	return call[discord.Message](ctx, a, req, opts)
}

func (a *API) EditDMMessage(ctx context.Context, userID snowflake.ID, messageID snowflake.ID, params discord.EditDMMessageParams, opts ...Option) (*discord.Message, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/users/{user_id}/messages/{message_id}", rest.P("user_id", userID), rest.P("message_id", messageID))}
	req.JSON = params
	return call[discord.Message](ctx, a, req, opts)
}

func (a *API) DeleteMessage(ctx context.Context, channelID snowflake.ID, messageID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/channels/{channel_id}/messages/{message_id}", rest.P("channel_id", channelID), rest.P("message_id", messageID))}
	return none(ctx, a, req, opts)
}

func (a *API) DeleteDMMessage(ctx context.Context, userID snowflake.ID, messageID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/users/{user_id}/messages/{message_id}", rest.P("user_id", userID), rest.P("message_id", messageID))}
	return none(ctx, a, req, opts)
}

func (a *API) BulkDeleteMessages(ctx context.Context, channelID snowflake.ID, params discord.BulkDeleteMessagesParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/channels/{channel_id}/messages/bulk-delete", rest.P("channel_id", channelID))}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) ListMessagePins(ctx context.Context, channelID snowflake.ID, query discord.ListMessagePinsQuery, opts ...Option) (*discord.ListMessagePinsResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/channels/{channel_id}/messages/pins", rest.P("channel_id", channelID))}
	req.Query = rest.Values(query)
	return call[discord.ListMessagePinsResponse](ctx, a, req, opts)
}

func (a *API) ListPinnedMessages(ctx context.Context, channelID snowflake.ID, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/channels/{channel_id}/pins", rest.P("channel_id", channelID))}
	return raw(ctx, a, req, opts)
}

func (a *API) PinMessage(ctx context.Context, channelID snowflake.ID, messageID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("PUT", "/channels/{channel_id}/messages/pins/{message_id}", rest.P("channel_id", channelID), rest.P("message_id", messageID))}
	return none(ctx, a, req, opts)
}

func (a *API) UnpinMessage(ctx context.Context, channelID snowflake.ID, messageID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/channels/{channel_id}/messages/pins/{message_id}", rest.P("channel_id", channelID), rest.P("message_id", messageID))}
	return none(ctx, a, req, opts)
}

func (a *API) GetChannelMediaPreview(ctx context.Context, channelID snowflake.ID, opts ...Option) (*discord.GetChannelMediaPreviewResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/channels/{channel_id}/media-post-preview", rest.P("channel_id", channelID))}
	return call[discord.GetChannelMediaPreviewResponse](ctx, a, req, opts)
}

func (a *API) UnfurlEmbed(ctx context.Context, params discord.UnfurlEmbedParams, opts ...Option) (*discord.UnfurlEmbedResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/unfurler/unfurl")}
	req.JSON = params
	return call[discord.UnfurlEmbedResponse](ctx, a, req, opts)
}

func (a *API) UnfurlEmbeds(ctx context.Context, params discord.UnfurlEmbedsParams, opts ...Option) (*discord.UnfurlEmbedsResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/unfurler/embed-urls")}
	req.JSON = params
	return call[discord.UnfurlEmbedsResponse](ctx, a, req, opts)
}

func (a *API) ListPollAnswerVoters(ctx context.Context, channelID snowflake.ID, messageID snowflake.ID, pollAnswerID snowflake.ID, query discord.ListPollAnswerVotersQuery, opts ...Option) (*discord.ListPollAnswerVotersResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/channels/{channel_id}/polls/{message_id}/answers/{poll_answer_id}", rest.P("channel_id", channelID), rest.P("message_id", messageID), rest.P("poll_answer_id", pollAnswerID))}
	req.Query = rest.Values(query)
	return call[discord.ListPollAnswerVotersResponse](ctx, a, req, opts)
}

func (a *API) EndPoll(ctx context.Context, channelID snowflake.ID, messageID snowflake.ID, opts ...Option) (*discord.Message, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/channels/{channel_id}/polls/{message_id}/expire", rest.P("channel_id", channelID), rest.P("message_id", messageID))}
	return call[discord.Message](ctx, a, req, opts)
}

func (a *API) CreatePollVote(ctx context.Context, channelID snowflake.ID, messageID snowflake.ID, params discord.CreatePollVoteParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("PUT", "/channels/{channel_id}/polls/{message_id}/answers/@me", rest.P("channel_id", channelID), rest.P("message_id", messageID))}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) ListConversationSummaries(ctx context.Context, channelID snowflake.ID, opts ...Option) (*discord.ListConversationSummariesResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/channels/{channel_id}/summaries", rest.P("channel_id", channelID))}
	return call[discord.ListConversationSummariesResponse](ctx, a, req, opts)
}

func (a *API) DeleteConversationSummary(ctx context.Context, channelID snowflake.ID, summaryID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/channels/{channel_id}/summaries/{summary_id}", rest.P("channel_id", channelID), rest.P("summary_id", summaryID))}
	return none(ctx, a, req, opts)
}

func (a *API) ListUserMessageSummaries(ctx context.Context, opts ...Option) ([]discord.UserMessageSummary, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/partner-sdk/users/@me/channels")}
	return list[discord.UserMessageSummary](ctx, a, req, opts)
}

func (a *API) ListNotificationCenterItems(ctx context.Context, query discord.ListNotificationCenterItemsQuery, opts ...Option) (*discord.ListNotificationCenterItemsResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/notification-center/items")}
	req.Query = rest.Values(query)
	return call[discord.ListNotificationCenterItemsResponse](ctx, a, req, opts)
}

func (a *API) DeleteNotificationCenterItem(ctx context.Context, notificationCenterItemID snowflake.ID, params discord.DeleteNotificationCenterItemParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/users/@me/notification-center/items/{notification_center_item_id}", rest.P("notification_center_item_id", notificationCenterItemID))}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) AcknowledgeNotificationCenterItem(ctx context.Context, notificationCenterItemID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/notification-center/items/{notification_center_item_id}/ack", rest.P("notification_center_item_id", notificationCenterItemID))}
	return none(ctx, a, req, opts)
}

func (a *API) BulkAcknowledgeNotificationCenterItems(ctx context.Context, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/notification-center/items/bulk-ack")}
	return none(ctx, a, req, opts)
}

func (a *API) VerifyPurchaseRequest(ctx context.Context, params discord.VerifyPurchaseRequestParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/billing/verify-purchase-request")}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) ResendPaymentVerificationEmail(ctx context.Context, params discord.ResendPaymentVerificationEmailParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/store/email/resend-payment-verification")}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) ListPayments(ctx context.Context, opts ...Option) ([]discord.Payment, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/billing/payments")}
	return list[discord.Payment](ctx, a, req, opts)
}

func (a *API) GetPayment(ctx context.Context, paymentID snowflake.ID, opts ...Option) (*discord.Payment, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/billing/payments/{payment_id}", rest.P("payment_id", paymentID))}
	return call[discord.Payment](ctx, a, req, opts)
}

func (a *API) VoidPayment(ctx context.Context, paymentID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/billing/payments/{payment_id}/void", rest.P("payment_id", paymentID))}
	return none(ctx, a, req, opts)
}

func (a *API) GetPaymentInvoiceBreakdown(ctx context.Context, query discord.GetPaymentInvoiceBreakdownQuery, opts ...Option) (*discord.GetPaymentInvoiceBreakdownResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/billing/invoice/breakdown")}
	req.Query = rest.Values(query)
	return call[discord.GetPaymentInvoiceBreakdownResponse](ctx, a, req, opts)
}

func (a *API) GetPremiumReferral(ctx context.Context, referralID snowflake.ID, opts ...Option) (*discord.PremiumReferral, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/referrals/{referral_id}", rest.P("referral_id", referralID))}
	return call[discord.PremiumReferral](ctx, a, req, opts)
}

func (a *API) GetPremiumReferralEligibility(ctx context.Context, opts ...Option) (*discord.PremiumReferralEligibility, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/referrals/eligibility")}
	return call[discord.PremiumReferralEligibility](ctx, a, req, opts)
}

func (a *API) GetPremiumReferralIncentiveEligibility(ctx context.Context, opts ...Option) (*discord.GetPremiumReferralIncentiveEligibilityResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/referrals/incentive-eligibility")}
	return call[discord.GetPremiumReferralIncentiveEligibilityResponse](ctx, a, req, opts)
}

func (a *API) ListPremiumReferralEligibleUsers(ctx context.Context, params discord.ListPremiumReferralEligibleUsersParams, opts ...Option) (*discord.ListPremiumReferralEligibleUsersResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/referrals/eligible-users")}
	req.JSON = params
	return call[discord.ListPremiumReferralEligibleUsersResponse](ctx, a, req, opts)
}

func (a *API) CreatePremiumReferral(ctx context.Context, userID snowflake.ID, opts ...Option) (*discord.PremiumReferral, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/referrals/{user_id}", rest.P("user_id", userID))}
	return call[discord.PremiumReferral](ctx, a, req, opts)
}

func (a *API) ListPresences(ctx context.Context, opts ...Option) (*discord.ListPresencesResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/presences")}
	return call[discord.ListPresencesResponse](ctx, a, req, opts)
}

func (a *API) ListPresencesForXbox(ctx context.Context, opts ...Option) (*discord.ListPresencesForXboxResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/consoles/xbox/presences")}
	return call[discord.ListPresencesForXboxResponse](ctx, a, req, opts)
}

func (a *API) UpdatePresence(ctx context.Context, params discord.UpdatePresenceParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/presences")}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) ListGlobalActivityStatistics(ctx context.Context, query discord.ListGlobalActivityStatisticsQuery, opts ...Option) ([]discord.GlobalActivityStatistics, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/activities")}
	req.Query = rest.Values(query)
	return list[discord.GlobalActivityStatistics](ctx, a, req, opts)
}

func (a *API) ListApplicationActivityStatistics(ctx context.Context, applicationID snowflake.ID, opts ...Option) ([]discord.ApplicationActivityStatistics, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/activities/statistics/applications/{application_id}", rest.P("application_id", applicationID))}
	return list[discord.ApplicationActivityStatistics](ctx, a, req, opts)
}

func (a *API) ListUserApplicationActivityStatistics(ctx context.Context, opts ...Option) ([]discord.UserApplicationActivityStatistics, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/activities/statistics/applications")}
	return list[discord.UserApplicationActivityStatistics](ctx, a, req, opts)
}

func (a *API) UpdateActivitySession(ctx context.Context, params discord.UpdateActivitySessionParams, opts ...Option) (*discord.UpdateActivitySessionResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/activities")}
	req.JSON = params
	return call[discord.UpdateActivitySessionResponse](ctx, a, req, opts)
}

func (a *API) CreateHeadlessSession(ctx context.Context, params discord.CreateHeadlessSessionParams, opts ...Option) (*discord.CreateHeadlessSessionResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/headless-sessions")}
	req.JSON = params
	return call[discord.CreateHeadlessSessionResponse](ctx, a, req, opts)
}

func (a *API) DeleteHeadlessSession(ctx context.Context, params discord.DeleteHeadlessSessionParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/headless-sessions/delete")}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) GetActivityMetadata(ctx context.Context, userID snowflake.ID, sessionID snowflake.ID, applicationID snowflake.ID, query discord.GetActivityMetadataQuery, opts ...Option) (*discord.GetActivityMetadataResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/{user_id}/sessions/{session_id}/activities/{application_id}/metadata", rest.P("user_id", userID), rest.P("session_id", sessionID), rest.P("application_id", applicationID))}
	req.Query = rest.Values(query)
	return call[discord.GetActivityMetadataResponse](ctx, a, req, opts)
}

func (a *API) UpdateActivitySubscriptions(ctx context.Context, params discord.UpdateActivitySubscriptionsParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/activities/subscribe")}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) ListOutboundPromotions(ctx context.Context, query discord.ListOutboundPromotionsQuery, opts ...Option) ([]discord.Promotion, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/outbound-promotions")}
	req.Query = rest.Values(query)
	return list[discord.Promotion](ctx, a, req, opts)
}

func (a *API) ListPromotions(ctx context.Context, query discord.ListPromotionsQuery, opts ...Option) ([]discord.Promotion, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/promotions")}
	req.Query = rest.Values(query)
	return list[discord.Promotion](ctx, a, req, opts)
}

func (a *API) ListBogoPromotions(ctx context.Context, query discord.ListBogoPromotionsQuery, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/bogo-promotions")}
	req.Query = rest.Values(query)
	return raw(ctx, a, req, opts)
}

func (a *API) ListClaimedPromotions(ctx context.Context, query discord.ListClaimedPromotionsQuery, opts ...Option) ([]discord.ClaimedPromotion, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/outbound-promotions/codes")}
	req.Query = rest.Values(query)
	return list[discord.ClaimedPromotion](ctx, a, req, opts)
}

func (a *API) ClaimPromotion(ctx context.Context, promotionID snowflake.ID, opts ...Option) (*discord.ClaimedPromotion, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/outbound-promotions/{promotion_id}/claim", rest.P("promotion_id", promotionID))}
	return call[discord.ClaimedPromotion](ctx, a, req, opts)
}

func (a *API) ListCurrentUserQuests(ctx context.Context, opts ...Option) (*discord.ListCurrentUserQuestsResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/quests/@me")}
	return call[discord.ListCurrentUserQuestsResponse](ctx, a, req, opts)
}

func (a *API) ListClaimedQuests(ctx context.Context, opts ...Option) (*discord.ListClaimedQuestsResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/quests/@me/claimed")}
	return call[discord.ListClaimedQuestsResponse](ctx, a, req, opts)
}

func (a *API) GetQuestConfig(ctx context.Context, questID snowflake.ID, opts ...Option) (*discord.QuestConfig, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/quests/{quest_id}", rest.P("quest_id", questID))}
	return call[discord.QuestConfig](ctx, a, req, opts)
}

func (a *API) GetQuestPreview(ctx context.Context, questID snowflake.ID, opts ...Option) (*discord.Quest, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/quests/{quest_id}/preview", rest.P("quest_id", questID))}
	return call[discord.Quest](ctx, a, req, opts)
}

func (a *API) GetQuestPlacement(ctx context.Context, query discord.GetQuestPlacementQuery, opts ...Option) (*discord.GetQuestPlacementResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/quests/decision")}
	req.Query = rest.Values(query)
	return call[discord.GetQuestPlacementResponse](ctx, a, req, opts)
}

func (a *API) GetEarnedQuestPlacement(ctx context.Context, query discord.GetEarnedQuestPlacementQuery, opts ...Option) (*discord.GetEarnedQuestPlacementResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/quests/earned-decision")}
	req.Query = rest.Values(query)
	return call[discord.GetEarnedQuestPlacementResponse](ctx, a, req, opts)
}

func (a *API) GetQuestDecisions(ctx context.Context, query discord.GetQuestDecisionsQuery, opts ...Option) (*discord.GetQuestDecisionsResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/quests/get-decisions")}
	req.Query = rest.Values(query)
	return call[discord.GetQuestDecisionsResponse](ctx, a, req, opts)
}

func (a *API) GetQuestCreativePreview(ctx context.Context, query discord.GetQuestCreativePreviewQuery, opts ...Option) (*discord.GetQuestCreativePreviewResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/quests/preview")}
	req.Query = rest.Values(query)
	return call[discord.GetQuestCreativePreviewResponse](ctx, a, req, opts)
}

func (a *API) AcceptQuest(ctx context.Context, questID snowflake.ID, params discord.AcceptQuestParams, opts ...Option) (*discord.QuestUserStatus, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/quests/{quest_id}/enroll", rest.P("quest_id", questID))}
	req.JSON = params
	return call[discord.QuestUserStatus](ctx, a, req, opts)
}

func (a *API) ClaimQuestReward(ctx context.Context, questID snowflake.ID, params discord.ClaimQuestRewardParams, opts ...Option) (*discord.ClaimQuestRewardResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/quests/{quest_id}/claim-reward", rest.P("quest_id", questID))}
	req.JSON = params
	return call[discord.ClaimQuestRewardResponse](ctx, a, req, opts)
}

func (a *API) ClaimQuestCreativeReward(ctx context.Context, adCreativeID snowflake.ID, params discord.ClaimQuestCreativeRewardParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/quests/creatives/{ad_creative_id}/claim-reward", rest.P("ad_creative_id", adCreativeID))}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) GetQuestRewardCode(ctx context.Context, questID snowflake.ID, opts ...Option) (*discord.QuestRewardCode, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/quests/{quest_id}/reward-code", rest.P("quest_id", questID))}
	return call[discord.QuestRewardCode](ctx, a, req, opts)
}

func (a *API) SendQuestHeartbeat(ctx context.Context, questID snowflake.ID, params discord.SendQuestHeartbeatParams, opts ...Option) (*discord.QuestUserStatus, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/quests/{quest_id}/heartbeat", rest.P("quest_id", questID))}
	req.JSON = params
	return call[discord.QuestUserStatus](ctx, a, req, opts)
}

func (a *API) SendQuestVideoProgress(ctx context.Context, questID snowflake.ID, params discord.SendQuestVideoProgressParams, opts ...Option) (*discord.QuestUserStatus, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/quests/{quest_id}/video-progress", rest.P("quest_id", questID))}
	req.JSON = params
	return call[discord.QuestUserStatus](ctx, a, req, opts)
}

func (a *API) StartConsoleQuest(ctx context.Context, questID snowflake.ID, query discord.StartConsoleQuestQuery, opts ...Option) (*discord.StartConsoleQuestResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/quests/{quest_id}/console/start", rest.P("quest_id", questID))}
	req.Query = rest.Values(query)
	return call[discord.StartConsoleQuestResponse](ctx, a, req, opts)
}

func (a *API) StopConsoleQuest(ctx context.Context, questID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/quests/{quest_id}/console/stop", rest.P("quest_id", questID))}
	return none(ctx, a, req, opts)
}

func (a *API) CompleteQuest(ctx context.Context, questID snowflake.ID, opts ...Option) (*discord.QuestUserStatus, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/quests/{quest_id}/preview/complete", rest.P("quest_id", questID))}
	return call[discord.QuestUserStatus](ctx, a, req, opts)
}

func (a *API) ResetQuest(ctx context.Context, questID snowflake.ID, opts ...Option) (*discord.QuestUserStatus, error) {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/quests/{quest_id}/preview/status", rest.P("quest_id", questID))}
	return call[discord.QuestUserStatus](ctx, a, req, opts)
}

func (a *API) ResetRecentQuestCompletions(ctx context.Context, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/quests/reset-recent-quest-completions")}
	return none(ctx, a, req, opts)
}

func (a *API) DismissQuestContent(ctx context.Context, questID snowflake.ID, questContentType string, opts ...Option) (*discord.QuestUserStatus, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/quests/{quest_id}/dismissible-content/{quest_content_type}/dismiss", rest.P("quest_id", questID), rest.P("quest_content_type", questContentType))}
	return call[discord.QuestUserStatus](ctx, a, req, opts)
}

func (a *API) ResetQuestDismissibility(ctx context.Context, questID snowflake.ID, opts ...Option) (*discord.QuestUserStatus, error) {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/quests/{quest_id}/preview/dismissibility", rest.P("quest_id", questID))}
	return call[discord.QuestUserStatus](ctx, a, req, opts)
}

func (a *API) ListRelationships(ctx context.Context, opts ...Option) ([]discord.Relationship, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/relationships")}
	return list[discord.Relationship](ctx, a, req, opts)
}

func (a *API) SendFriendRequest(ctx context.Context, params discord.SendFriendRequestParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/relationships")}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) CreateRelationship(ctx context.Context, userID snowflake.ID, params discord.CreateRelationshipParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("PUT", "/users/@me/relationships/{user_id}", rest.P("user_id", userID))}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) IgnoreUser(ctx context.Context, userID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("PUT", "/users/@me/relationships/{user_id}/ignore", rest.P("user_id", userID))}
	return none(ctx, a, req, opts)
}

func (a *API) UnignoreUser(ctx context.Context, userID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/users/@me/relationships/{user_id}/ignore", rest.P("user_id", userID))}
	return none(ctx, a, req, opts)
}

func (a *API) ModifyRelationship(ctx context.Context, userID snowflake.ID, params discord.ModifyRelationshipParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/users/@me/relationships/{user_id}", rest.P("user_id", userID))}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) RemoveRelationship(ctx context.Context, userID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/users/@me/relationships/{user_id}", rest.P("user_id", userID))}
	return none(ctx, a, req, opts)
}

func (a *API) BulkRemoveRelationships(ctx context.Context, query discord.BulkRemoveRelationshipsQuery, params discord.BulkRemoveRelationshipsParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/users/@me/relationships")}
	req.Query = rest.Values(query)
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) BulkAddRelationships(ctx context.Context, params discord.BulkAddRelationshipsParams, opts ...Option) (*discord.BulkAddRelationshipsResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/relationships/bulk")}
	req.JSON = params
	return call[discord.BulkAddRelationshipsResponse](ctx, a, req, opts)
}

func (a *API) ListGameRelationships(ctx context.Context, opts ...Option) ([]discord.GameRelationship, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/game-relationships")}
	return list[discord.GameRelationship](ctx, a, req, opts)
}

func (a *API) SendGameFriendRequest(ctx context.Context, params discord.SendGameFriendRequestParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/game-relationships")}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) CreateGameRelationship(ctx context.Context, userID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("PUT", "/users/@me/game-relationships/{user_id}", rest.P("user_id", userID))}
	return none(ctx, a, req, opts)
}

func (a *API) CreateGameRelationshipByApplication(ctx context.Context, userID snowflake.ID, applicationID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("PUT", "/users/@me/game-relationships/{user_id}/{application_id}", rest.P("user_id", userID), rest.P("application_id", applicationID))}
	return none(ctx, a, req, opts)
}

func (a *API) RemoveGameRelationship(ctx context.Context, userID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/users/@me/game-relationships/{user_id}", rest.P("user_id", userID))}
	return none(ctx, a, req, opts)
}

func (a *API) RemoveGameRelationshipByApplication(ctx context.Context, userID snowflake.ID, applicationID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/users/@me/game-relationships/{user_id}/{application_id}", rest.P("user_id", userID), rest.P("application_id", applicationID))}
	return none(ctx, a, req, opts)
}

func (a *API) ListFriendSuggestions(ctx context.Context, opts ...Option) ([]discord.FriendSuggestion, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/friend-suggestions")}
	return list[discord.FriendSuggestion](ctx, a, req, opts)
}

func (a *API) RemoveFriendSuggestion(ctx context.Context, userID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/friend-suggestions/{user_id}", rest.P("user_id", userID))}
	return none(ctx, a, req, opts)
}

func (a *API) GetUserSafetyHub(ctx context.Context, opts ...Option) (*discord.SafetyHub, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/safety-hub/@me")}
	return call[discord.SafetyHub](ctx, a, req, opts)
}

func (a *API) GetSuspendedUserSafetyHub(ctx context.Context, params discord.GetSuspendedUserSafetyHubParams, opts ...Option) (*discord.SafetyHub, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/safety-hub/suspended/@me")}
	req.JSON = params
	return call[discord.SafetyHub](ctx, a, req, opts)
}

func (a *API) RequestClassificationReview(ctx context.Context, classificationID snowflake.ID, params discord.RequestClassificationReviewParams, opts ...Option) (*discord.RequestClassificationReviewResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("PUT", "/safety-hub/request-review/{classification_id}", rest.P("classification_id", classificationID))}
	req.JSON = params
	return call[discord.RequestClassificationReviewResponse](ctx, a, req, opts)
}

func (a *API) RequestClassificationReviewForSuspendedUser(ctx context.Context, classificationID snowflake.ID, params discord.RequestClassificationReviewForSuspendedUserParams, opts ...Option) (*discord.RequestClassificationReviewForSuspendedUserResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("PUT", "/safety-hub/suspended/request-review/{classification_id}", rest.P("classification_id", classificationID))}
	req.JSON = params
	return call[discord.RequestClassificationReviewForSuspendedUserResponse](ctx, a, req, opts)
}

func (a *API) CheckAgeVerificationForSuspendedUser(ctx context.Context, params discord.CheckAgeVerificationForSuspendedUserParams, opts ...Option) (*discord.CheckAgeVerificationForSuspendedUserResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/safety-hub/suspended/check-verification")}
	req.JSON = params
	return call[discord.CheckAgeVerificationForSuspendedUserResponse](ctx, a, req, opts)
}

func (a *API) RequestAgeVerificationForSuspendedUser(ctx context.Context, params discord.RequestAgeVerificationForSuspendedUserParams, opts ...Option) (*discord.RequestAgeVerificationForSuspendedUserResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/safety-hub/suspended/request-verification")}
	req.JSON = params
	return call[discord.RequestAgeVerificationForSuspendedUserResponse](ctx, a, req, opts)
}

func (a *API) ListDefaultSoundboardSounds(ctx context.Context, opts ...Option) ([]discord.SoundboardSound, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/soundboard-default-sounds")}
	return list[discord.SoundboardSound](ctx, a, req, opts)
}

func (a *API) ListGuildSoundboardSounds(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.ListGuildSoundboardSoundsResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/soundboard-sounds", rest.P("guild_id", guildID))}
	return call[discord.ListGuildSoundboardSoundsResponse](ctx, a, req, opts)
}

func (a *API) GetGuildSoundboardSound(ctx context.Context, guildID snowflake.ID, soundID snowflake.ID, opts ...Option) (*discord.SoundboardSound, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/soundboard-sounds/{sound_id}", rest.P("guild_id", guildID), rest.P("sound_id", soundID))}
	return call[discord.SoundboardSound](ctx, a, req, opts)
}

func (a *API) CreateGuildSoundboardSound(ctx context.Context, guildID snowflake.ID, params discord.CreateGuildSoundboardSoundParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/guilds/{guild_id}/soundboard-sounds", rest.P("guild_id", guildID))}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) ModifyGuildSoundboardSound(ctx context.Context, guildID snowflake.ID, soundID snowflake.ID, params discord.ModifyGuildSoundboardSoundParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/guilds/{guild_id}/soundboard-sounds/{sound_id}", rest.P("guild_id", guildID), rest.P("sound_id", soundID))}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) DeleteGuildSoundboardSound(ctx context.Context, guildID snowflake.ID, soundID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/guilds/{guild_id}/soundboard-sounds/{sound_id}", rest.P("guild_id", guildID), rest.P("sound_id", soundID))}
	return none(ctx, a, req, opts)
}

func (a *API) GetSoundboardSoundGuild(ctx context.Context, soundID snowflake.ID, guildID snowflake.ID, opts ...Option) (*discord.DiscoverableGuild, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/soundboard-sounds/{sound_id}/guild/{guild_id}", rest.P("sound_id", soundID), rest.P("guild_id", guildID))}
	return call[discord.DiscoverableGuild](ctx, a, req, opts)
}

func (a *API) SendSoundboardSound(ctx context.Context, channelID snowflake.ID, params discord.SendSoundboardSoundParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/channels/{channel_id}/send-soundboard-sound", rest.P("channel_id", channelID))}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) CreateStageInstance(ctx context.Context, params discord.CreateStageInstanceParams, opts ...Option) (*discord.StageInstance, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/stage-instances")}
	req.JSON = params
	return call[discord.StageInstance](ctx, a, req, opts)
}

func (a *API) GetStageInstance(ctx context.Context, channelID snowflake.ID, opts ...Option) (*discord.StageInstance, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/stage-instances/{channel_id}", rest.P("channel_id", channelID))}
	return call[discord.StageInstance](ctx, a, req, opts)
}

func (a *API) ModifyStageInstance(ctx context.Context, channelID snowflake.ID, params discord.ModifyStageInstanceParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/stage-instances/{channel_id}", rest.P("channel_id", channelID))}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) DeleteStageInstance(ctx context.Context, channelID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/stage-instances/{channel_id}", rest.P("channel_id", channelID))}
	return none(ctx, a, req, opts)
}

func (a *API) ListStickerPacks(ctx context.Context, opts ...Option) (*discord.ListStickerPacksResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/sticker-packs")}
	return call[discord.ListStickerPacksResponse](ctx, a, req, opts)
}

func (a *API) GetStickerPack(ctx context.Context, stickerPackID snowflake.ID, opts ...Option) (*discord.StickerPack, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/sticker-packs/{sticker_pack_id}", rest.P("sticker_pack_id", stickerPackID))}
	return call[discord.StickerPack](ctx, a, req, opts)
}

func (a *API) GetSticker(ctx context.Context, stickerID snowflake.ID, opts ...Option) (*discord.Sticker, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/stickers/{sticker_id}", rest.P("sticker_id", stickerID))}
	return call[discord.Sticker](ctx, a, req, opts)
}

func (a *API) GetStickerGuild(ctx context.Context, stickerID snowflake.ID, opts ...Option) (*discord.DiscoverableGuild, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/stickers/{sticker_id}/guild", rest.P("sticker_id", stickerID))}
	return call[discord.DiscoverableGuild](ctx, a, req, opts)
}

func (a *API) ListGuildStickers(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.Sticker, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/stickers", rest.P("guild_id", guildID))}
	return list[discord.Sticker](ctx, a, req, opts)
}

func (a *API) GetGuildSticker(ctx context.Context, guildID snowflake.ID, stickerID snowflake.ID, opts ...Option) (*discord.Sticker, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/stickers/{sticker_id}", rest.P("guild_id", guildID), rest.P("sticker_id", stickerID))}
	return call[discord.Sticker](ctx, a, req, opts)
}

func (a *API) CreateGuildSticker(ctx context.Context, guildID snowflake.ID, params discord.CreateGuildStickerParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/guilds/{guild_id}/stickers", rest.P("guild_id", guildID))}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) ModifyGuildSticker(ctx context.Context, guildID snowflake.ID, stickerID snowflake.ID, params discord.ModifyGuildStickerParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/guilds/{guild_id}/stickers/{sticker_id}", rest.P("guild_id", guildID), rest.P("sticker_id", stickerID))}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) DeleteGuildSticker(ctx context.Context, guildID snowflake.ID, stickerID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/guilds/{guild_id}/stickers/{sticker_id}", rest.P("guild_id", guildID), rest.P("sticker_id", stickerID))}
	return none(ctx, a, req, opts)
}

func (a *API) CreateSKU(ctx context.Context, params discord.CreateSKUParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/store/skus")}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) GetSKU(ctx context.Context, skuid snowflake.ID, query discord.GetSKUQuery, opts ...Option) (*discord.SKU, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/store/skus/{sku_id}", rest.P("sku_id", skuid))}
	req.Query = rest.Values(query)
	return call[discord.SKU](ctx, a, req, opts)
}

func (a *API) ModifySKU(ctx context.Context, skuid snowflake.ID, params discord.ModifySKUParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/store/skus/{sku_id}", rest.P("sku_id", skuid))}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) ListSKUStoreListings(ctx context.Context, skuid snowflake.ID, query discord.ListSKUStoreListingsQuery, opts ...Option) ([]discord.StoreListing, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/store/skus/{sku_id}/listings", rest.P("sku_id", skuid))}
	req.Query = rest.Values(query)
	return list[discord.StoreListing](ctx, a, req, opts)
}

func (a *API) CreateStoreListing(ctx context.Context, query discord.CreateStoreListingQuery, params discord.StoreCreateStoreListingParams, opts ...Option) (*discord.StoreListing, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/store/listings")}
	req.Query = rest.Values(query)
	req.JSON = params
	return call[discord.StoreListing](ctx, a, req, opts)
}

func (a *API) ListApplicationPublishedStoreListings(ctx context.Context, query discord.ListApplicationPublishedStoreListingsQuery, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/store/published-listings/skus")}
	req.Query = rest.Values(query)
	return raw(ctx, a, req, opts)
}

func (a *API) GetApplicationPrimaryStoreListing(ctx context.Context, applicationID snowflake.ID, query discord.GetApplicationPrimaryStoreListingQuery, opts ...Option) (*discord.StoreListing, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/store/published-listings/applications/{application_id}", rest.P("application_id", applicationID))}
	req.Query = rest.Values(query)
	return call[discord.StoreListing](ctx, a, req, opts)
}

func (a *API) ListBulkApplicationPrimaryStoreListing(ctx context.Context, query discord.ListBulkApplicationPrimaryStoreListingQuery, opts ...Option) ([]discord.StoreListing, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/store/published-listings/applications")}
	req.Query = rest.Values(query)
	return list[discord.StoreListing](ctx, a, req, opts)
}

func (a *API) GetSKUPublishedStoreListing(ctx context.Context, skuid snowflake.ID, query discord.GetSKUPublishedStoreListingQuery, opts ...Option) (*discord.StoreListing, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/store/published-listings/skus/{sku_id}", rest.P("sku_id", skuid))}
	req.Query = rest.Values(query)
	return call[discord.StoreListing](ctx, a, req, opts)
}

func (a *API) ListSubscriptionPlans(ctx context.Context, skuid snowflake.ID, opts ...Option) ([]discord.SubscriptionPlan, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/store/skus/{sku_id}/plans", rest.P("sku_id", skuid))}
	return list[discord.SubscriptionPlan](ctx, a, req, opts)
}

func (a *API) ListPublishedSubscriptionPlans(ctx context.Context, skuid snowflake.ID, query discord.ListPublishedSubscriptionPlansQuery, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/store/published-listings/skus/{sku_id}/subscription-plans", rest.P("sku_id", skuid))}
	req.Query = rest.Values(query)
	return raw(ctx, a, req, opts)
}

func (a *API) ListBulkPublishedSubscriptionPlans(ctx context.Context, query discord.ListBulkPublishedSubscriptionPlansQuery, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/store/published-listings/skus/subscription-plans")}
	req.Query = rest.Values(query)
	return raw(ctx, a, req, opts)
}

func (a *API) GetSubscriptionGroupListingBySubscriptionPlan(ctx context.Context, subscriptionPlanID snowflake.ID, opts ...Option) (*discord.GetSubscriptionGroupListingBySubscriptionPlanResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/subscription-plans/{subscription_plan_id}/subscription-group-listing", rest.P("subscription_plan_id", subscriptionPlanID))}
	return call[discord.GetSubscriptionGroupListingBySubscriptionPlanResponse](ctx, a, req, opts)
}

func (a *API) GetSKUPurchasePreview(ctx context.Context, skuid snowflake.ID, query discord.GetSKUPurchasePreviewQuery, opts ...Option) (*discord.Invoice, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/store/skus/{sku_id}/purchase", rest.P("sku_id", skuid))}
	req.Query = rest.Values(query)
	return call[discord.Invoice](ctx, a, req, opts)
}

func (a *API) CreateSKUPurchase(ctx context.Context, skuid snowflake.ID, params discord.StoreCreateSKUPurchaseParams, opts ...Option) (*discord.CreateSKUPurchaseResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/store/skus/{sku_id}/purchase", rest.P("sku_id", skuid))}
	req.JSON = params
	return call[discord.CreateSKUPurchaseResponse](ctx, a, req, opts)
}

func (a *API) ListStorePriceTiers(ctx context.Context, query discord.ListStorePriceTiersQuery, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/store/price-tiers")}
	req.Query = rest.Values(query)
	return raw(ctx, a, req, opts)
}

func (a *API) GetStorePriceTier(ctx context.Context, priceTier string, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/store/price-tiers/{price_tier}", rest.P("price_tier", priceTier))}
	return raw(ctx, a, req, opts)
}

func (a *API) GetEula(ctx context.Context, eulaID snowflake.ID, opts ...Option) (*discord.Eula, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/store/eulas/{eula_id}", rest.P("eula_id", eulaID))}
	return call[discord.Eula](ctx, a, req, opts)
}

func (a *API) GetApplicationStoreLayout(ctx context.Context, applicationID snowflake.ID, opts ...Option) (*discord.GetApplicationStoreLayoutResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/applications/{application_id}/store-layout", rest.P("application_id", applicationID))}
	return call[discord.GetApplicationStoreLayoutResponse](ctx, a, req, opts)
}

func (a *API) ModifyApplicationStorefrontPublishStates(ctx context.Context, applicationID snowflake.ID, params discord.ModifyApplicationStorefrontPublishStatesParams, opts ...Option) (*discord.ModifyApplicationStorefrontPublishStatesResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/applications/{application_id}/storefront/publish", rest.P("application_id", applicationID))}
	req.JSON = params
	return call[discord.ModifyApplicationStorefrontPublishStatesResponse](ctx, a, req, opts)
}

func (a *API) GetConsumableSKUPricing(ctx context.Context, skuid snowflake.ID, opts ...Option) (*discord.GetConsumableSKUPricingResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/store/consumable/pricing/{sku_id}", rest.P("sku_id", skuid))}
	return call[discord.GetConsumableSKUPricingResponse](ctx, a, req, opts)
}

func (a *API) GetHdStreamingConsumable(ctx context.Context, opts ...Option) (*discord.GetHdStreamingConsumableResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/consumable/hd-streaming")}
	return call[discord.GetHdStreamingConsumableResponse](ctx, a, req, opts)
}

func (a *API) ApplyHdStreamingConsumable(ctx context.Context, params discord.ApplyHdStreamingConsumableParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/consumable/hd-streaming")}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) GetConfettiConsumable(ctx context.Context, opts ...Option) (*discord.GetConfettiConsumableResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/consumable/confetti")}
	return call[discord.GetConfettiConsumableResponse](ctx, a, req, opts)
}

func (a *API) ApplyConfettiConsumable(ctx context.Context, params discord.ApplyConfettiConsumableParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/consumable/confetti")}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) GetVirtualCurrencyBalance(ctx context.Context, opts ...Option) (*discord.GetVirtualCurrencyBalanceResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/virtual-currency/balance")}
	return call[discord.GetVirtualCurrencyBalanceResponse](ctx, a, req, opts)
}

func (a *API) RedeemVirtualCurrency(ctx context.Context, skuid snowflake.ID, params discord.RedeemVirtualCurrencyParams, opts ...Option) ([]discord.Entitlement, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/virtual-currency/skus/{sku_id}/redeem", rest.P("sku_id", skuid))}
	req.JSON = params
	return list[discord.Entitlement](ctx, a, req, opts)
}

func (a *API) GetApplicationStorefrontForPremiumButton(ctx context.Context, storeListingID snowflake.ID, opts ...Option) (*discord.GetApplicationStorefrontForPremiumButtonResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/applications/storefront/interactions/premium-button/{store_listing_id}", rest.P("store_listing_id", storeListingID))}
	return call[discord.GetApplicationStorefrontForPremiumButtonResponse](ctx, a, req, opts)
}

func (a *API) GetStorefrontCollection(ctx context.Context, collectionID snowflake.ID, query discord.GetStorefrontCollectionQuery, opts ...Option) (*discord.GetStorefrontCollectionResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/storefront/collections/{collection_id}", rest.P("collection_id", collectionID))}
	req.Query = rest.Values(query)
	return call[discord.GetStorefrontCollectionResponse](ctx, a, req, opts)
}

func (a *API) GetStorefrontProduct(ctx context.Context, productID snowflake.ID, opts ...Option) (*discord.StorefrontProduct, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/storefront/products/{product_id}", rest.P("product_id", productID))}
	return call[discord.StorefrontProduct](ctx, a, req, opts)
}

func (a *API) GetStorefrontProductBySKUID(ctx context.Context, skuid snowflake.ID, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/storefront/products/sku/{sku_id}", rest.P("sku_id", skuid))}
	return raw(ctx, a, req, opts)
}

func (a *API) ListStorefrontProductsBySKUID(ctx context.Context, query discord.ListStorefrontProductsBySKUIDQuery, opts ...Option) (*discord.ListStorefrontProductsBySKUIDResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/storefront/products/skus")}
	req.Query = rest.Values(query)
	return call[discord.ListStorefrontProductsBySKUIDResponse](ctx, a, req, opts)
}

func (a *API) ListStorefrontSKUPrices(ctx context.Context, query discord.ListStorefrontSKUPricesQuery, opts ...Option) (*discord.ListStorefrontSKUPricesResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/storefront/skus/prices")}
	req.Query = rest.Values(query)
	return call[discord.ListStorefrontSKUPricesResponse](ctx, a, req, opts)
}

func (a *API) GetGuildApplicationStorefront(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.Storefront, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/partner-sdk/guilds/{guild_id}/application-storefront", rest.P("guild_id", guildID))}
	return call[discord.Storefront](ctx, a, req, opts)
}

func (a *API) ListSocialLayerSKUs(ctx context.Context, applicationID snowflake.ID, opts ...Option) ([]discord.SKU, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/partner-sdk/application/{application_id}/skus", rest.P("application_id", applicationID))}
	return list[discord.SKU](ctx, a, req, opts)
}

func (a *API) CreateSocialLayerSKU(ctx context.Context, applicationID snowflake.ID, params discord.CreateSocialLayerSKUParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/partner-sdk/application/{application_id}/skus", rest.P("application_id", applicationID))}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) GetApplicationStorefront(ctx context.Context, applicationID snowflake.ID, opts ...Option) (*discord.Storefront, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/partner-sdk/applications/{application_id}/storefront", rest.P("application_id", applicationID))}
	return call[discord.Storefront](ctx, a, req, opts)
}

func (a *API) ModifyApplicationStorefront(ctx context.Context, applicationID snowflake.ID, params discord.ModifyApplicationStorefrontParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("PUT", "/partner-sdk/applications/{application_id}/storefront", rest.P("application_id", applicationID))}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) DeleteApplicationStorefront(ctx context.Context, applicationID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/partner-sdk/applications/{application_id}/storefront", rest.P("application_id", applicationID))}
	return none(ctx, a, req, opts)
}

func (a *API) GetGuildApplicationSKUStorefront(ctx context.Context, guildID snowflake.ID, skuid snowflake.ID, opts ...Option) (*discord.GetGuildApplicationSKUStorefrontResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/partner-sdk/guilds/{guild_id}/application-storefront/skus/{sku_id}", rest.P("guild_id", guildID), rest.P("sku_id", skuid))}
	return call[discord.GetGuildApplicationSKUStorefrontResponse](ctx, a, req, opts)
}

func (a *API) GetApplicationSKUStorefront(ctx context.Context, applicationID snowflake.ID, skuid snowflake.ID, opts ...Option) (*discord.GetApplicationSKUStorefrontResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/partner-sdk/applications/{application_id}/storefront/skus/{sku_id}", rest.P("application_id", applicationID), rest.P("sku_id", skuid))}
	return call[discord.GetApplicationSKUStorefrontResponse](ctx, a, req, opts)
}

func (a *API) GetGuildApplicationStorefrontAnnouncement(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.GetGuildApplicationStorefrontAnnouncementResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/partner-sdk/guilds/{guild_id}/application-storefront/announcement", rest.P("guild_id", guildID))}
	return call[discord.GetGuildApplicationStorefrontAnnouncementResponse](ctx, a, req, opts)
}

func (a *API) CheckSocialLayerSKUPurchaseEligibility(ctx context.Context, applicationID snowflake.ID, skuid snowflake.ID, opts ...Option) (*discord.CheckSocialLayerSKUPurchaseEligibilityResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/partner-sdk/applications/{application_id}/storefront/skus/{sku_id}/eligibility", rest.P("application_id", applicationID), rest.P("sku_id", skuid))}
	return call[discord.CheckSocialLayerSKUPurchaseEligibilityResponse](ctx, a, req, opts)
}

func (a *API) GetSocialLayerStorefrontConfig(ctx context.Context, opts ...Option) (*discord.GetSocialLayerStorefrontConfigResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/partner-sdk/storefront-config")}
	return call[discord.GetSocialLayerStorefrontConfigResponse](ctx, a, req, opts)
}

func (a *API) GetSocialLayerStorefrontEligibilities(ctx context.Context, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/partner-sdk/storefront-eligibility")}
	return raw(ctx, a, req, opts)
}

func (a *API) ListSubscriptions(ctx context.Context, query discord.ListSubscriptionsQuery, opts ...Option) ([]discord.Subscription, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/billing/subscriptions")}
	req.Query = rest.Values(query)
	return list[discord.Subscription](ctx, a, req, opts)
}

func (a *API) GetSubscription(ctx context.Context, subscriptionID snowflake.ID, opts ...Option) (*discord.Subscription, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/billing/subscriptions/{subscription_id}", rest.P("subscription_id", subscriptionID))}
	return call[discord.Subscription](ctx, a, req, opts)
}

func (a *API) CreateSubscription(ctx context.Context, params discord.CreateSubscriptionParams, opts ...Option) (*discord.Subscription, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/billing/subscriptions")}
	req.JSON = params
	return call[discord.Subscription](ctx, a, req, opts)
}

func (a *API) ModifySubscription(ctx context.Context, subscriptionID snowflake.ID, query discord.ModifySubscriptionQuery, params discord.ModifySubscriptionParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/users/@me/billing/subscriptions/{subscription_id}", rest.P("subscription_id", subscriptionID))}
	req.Query = rest.Values(query)
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) DeleteSubscription(ctx context.Context, subscriptionID snowflake.ID, query discord.DeleteSubscriptionQuery, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/users/@me/billing/subscriptions/{subscription_id}", rest.P("subscription_id", subscriptionID))}
	req.Query = rest.Values(query)
	return none(ctx, a, req, opts)
}

func (a *API) CreateSubscriptionPreview(ctx context.Context, params discord.CreateSubscriptionPreviewParams, opts ...Option) (*discord.Invoice, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/billing/subscriptions/preview")}
	req.JSON = params
	return call[discord.Invoice](ctx, a, req, opts)
}

func (a *API) GetSubscriptionPreview(ctx context.Context, subscriptionID snowflake.ID, opts ...Option) (*discord.Invoice, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/billing/subscriptions/{subscription_id}/preview", rest.P("subscription_id", subscriptionID))}
	return call[discord.Invoice](ctx, a, req, opts)
}

func (a *API) ModifySubscriptionPreview(ctx context.Context, subscriptionID snowflake.ID, params discord.ModifySubscriptionPreviewParams, opts ...Option) (*discord.Invoice, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/users/@me/billing/subscriptions/{subscription_id}/preview", rest.P("subscription_id", subscriptionID))}
	req.JSON = params
	return call[discord.Invoice](ctx, a, req, opts)
}

func (a *API) ListSubscriptionInvoices(ctx context.Context, subscriptionID snowflake.ID, params discord.ListSubscriptionInvoicesParams, opts ...Option) ([]discord.Invoice, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/billing/subscriptions/{subscription_id}/invoices", rest.P("subscription_id", subscriptionID))}
	req.JSON = params
	return list[discord.Invoice](ctx, a, req, opts)
}

func (a *API) ClaimSubscriptionPromotionReward(ctx context.Context, subscriptionID snowflake.ID, opts ...Option) (*discord.ClaimSubscriptionPromotionRewardResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/billing/subscriptions/{subscription_id}/promotion-reward", rest.P("subscription_id", subscriptionID))}
	return call[discord.ClaimSubscriptionPromotionRewardResponse](ctx, a, req, opts)
}

func (a *API) GetPremiumGuildSubscriptionCooldown(ctx context.Context, opts ...Option) (*discord.GetPremiumGuildSubscriptionCooldownResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/guilds/premium/subscriptions/cooldown")}
	return call[discord.GetPremiumGuildSubscriptionCooldownResponse](ctx, a, req, opts)
}

func (a *API) ListAppliedPremiumGuildSubscriptions(ctx context.Context, query discord.ListAppliedPremiumGuildSubscriptionsQuery, opts ...Option) ([]discord.PremiumGuildSubscription, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/guilds/premium/subscriptions")}
	req.Query = rest.Values(query)
	return list[discord.PremiumGuildSubscription](ctx, a, req, opts)
}

func (a *API) ListPremiumGuildSubscriptionSlots(ctx context.Context, opts ...Option) ([]discord.PremiumGuildSubscriptionSlot, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/guilds/premium/subscription-slots")}
	return list[discord.PremiumGuildSubscriptionSlot](ctx, a, req, opts)
}

func (a *API) ListTeams(ctx context.Context, query discord.ListTeamsQuery, opts ...Option) ([]discord.Team, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/teams")}
	req.Query = rest.Values(query)
	return list[discord.Team](ctx, a, req, opts)
}

func (a *API) CreateTeam(ctx context.Context, params discord.CreateTeamParams, opts ...Option) (*discord.Team, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/teams")}
	req.JSON = params
	return call[discord.Team](ctx, a, req, opts)
}

func (a *API) GetTeam(ctx context.Context, teamID snowflake.ID, opts ...Option) (*discord.Team, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/teams/{team_id}", rest.P("team_id", teamID))}
	return call[discord.Team](ctx, a, req, opts)
}

func (a *API) ModifyTeam(ctx context.Context, teamID snowflake.ID, params discord.ModifyTeamParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/teams/{team_id}", rest.P("team_id", teamID))}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) DeleteTeam(ctx context.Context, teamID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/teams/{team_id}/delete", rest.P("team_id", teamID))}
	return none(ctx, a, req, opts)
}

func (a *API) AcceptTeamInvite(ctx context.Context, params discord.AcceptTeamInviteParams, opts ...Option) (*discord.Team, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/teams/invite/accept")}
	req.JSON = params
	return call[discord.Team](ctx, a, req, opts)
}

func (a *API) ListTeamMembers(ctx context.Context, teamID snowflake.ID, opts ...Option) ([]discord.TeamMember, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/teams/{team_id}/members", rest.P("team_id", teamID))}
	return list[discord.TeamMember](ctx, a, req, opts)
}

func (a *API) AddTeamMember(ctx context.Context, teamID snowflake.ID, params discord.AddTeamMemberParams, opts ...Option) (*discord.TeamMember, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/teams/{team_id}/members", rest.P("team_id", teamID))}
	req.JSON = params
	return call[discord.TeamMember](ctx, a, req, opts)
}

func (a *API) ModifyTeamMember(ctx context.Context, teamID snowflake.ID, userID snowflake.ID, params discord.ModifyTeamMemberParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/teams/{team_id}/members/{user_id}", rest.P("team_id", teamID), rest.P("user_id", userID))}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) RemoveTeamMember(ctx context.Context, teamID snowflake.ID, userID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/teams/{team_id}/members/{user_id}", rest.P("team_id", teamID), rest.P("user_id", userID))}
	return none(ctx, a, req, opts)
}

func (a *API) ListTeamApplications(ctx context.Context, teamID snowflake.ID, opts ...Option) ([]discord.Application, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/teams/{team_id}/applications", rest.P("team_id", teamID))}
	return list[discord.Application](ctx, a, req, opts)
}

func (a *API) GetTeamStripeConnectURL(ctx context.Context, teamID snowflake.ID, params discord.GetTeamStripeConnectURLParams, opts ...Option) (*discord.GetTeamStripeConnectURLResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/teams/{team_id}/stripe/connect/redirect-url", rest.P("team_id", teamID))}
	req.JSON = params
	return call[discord.GetTeamStripeConnectURLResponse](ctx, a, req, opts)
}

func (a *API) GetTeamPayoutOnboarding(ctx context.Context, teamID snowflake.ID, opts ...Option) (*discord.GetTeamPayoutOnboardingResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/teams/{team_id}/payouts/onboarding", rest.P("team_id", teamID))}
	return call[discord.GetTeamPayoutOnboardingResponse](ctx, a, req, opts)
}

func (a *API) ListTeamPayouts(ctx context.Context, teamID snowflake.ID, query discord.ListTeamPayoutsQuery, opts ...Option) ([]discord.TeamPayout, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/teams/{team_id}/payouts", rest.P("team_id", teamID))}
	req.Query = rest.Values(query)
	return list[discord.TeamPayout](ctx, a, req, opts)
}

func (a *API) GetTeamPayoutReport(ctx context.Context, teamID snowflake.ID, payoutID snowflake.ID, query discord.GetTeamPayoutReportQuery, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/teams/{team_id}/payouts/{payout_id}/report", rest.P("team_id", teamID), rest.P("payout_id", payoutID))}
	req.Query = rest.Values(query)
	return raw(ctx, a, req, opts)
}

func (a *API) SearchCompanies(ctx context.Context, query discord.SearchCompaniesQuery, opts ...Option) ([]discord.TeamCompany, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/companies")}
	req.Query = rest.Values(query)
	return list[discord.TeamCompany](ctx, a, req, opts)
}

func (a *API) GetCompany(ctx context.Context, companyID snowflake.ID, opts ...Option) (*discord.TeamCompany, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/company/{company_id}", rest.P("company_id", companyID))}
	return call[discord.TeamCompany](ctx, a, req, opts)
}

func (a *API) CreateCompany(ctx context.Context, teamID snowflake.ID, params discord.CreateCompanyParams, opts ...Option) (*discord.TeamCompany, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/teams/{team_id}/companies", rest.P("team_id", teamID))}
	req.JSON = params
	return call[discord.TeamCompany](ctx, a, req, opts)
}

func (a *API) CreateTeamIdentityVerification(ctx context.Context, teamID snowflake.ID, params discord.CreateTeamIdentityVerificationParams, opts ...Option) (*discord.UserIdentityVerification, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/teams/{team_id}/identity/verification", rest.P("team_id", teamID))}
	req.JSON = params
	return call[discord.UserIdentityVerification](ctx, a, req, opts)
}

func (a *API) GetTeamIdentityVerification(ctx context.Context, teamID snowflake.ID, opts ...Option) (*discord.UserIdentityVerification, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/teams/{team_id}/identity/verification", rest.P("team_id", teamID))}
	return call[discord.UserIdentityVerification](ctx, a, req, opts)
}

func (a *API) GetUserSettingsProto(ctx context.Context, userSettingsProtoType string, opts ...Option) (*discord.GetUserSettingsProtoResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/settings-proto/{user_settings_proto_type}", rest.P("user_settings_proto_type", userSettingsProtoType))}
	return call[discord.GetUserSettingsProtoResponse](ctx, a, req, opts)
}

func (a *API) ModifyUserSettingsProto(ctx context.Context, userSettingsProtoType string, params discord.ModifyUserSettingsProtoParams, opts ...Option) (*discord.ModifyUserSettingsProtoResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/users/@me/settings-proto/{user_settings_proto_type}", rest.P("user_settings_proto_type", userSettingsProtoType))}
	req.JSON = params
	return call[discord.ModifyUserSettingsProtoResponse](ctx, a, req, opts)
}

func (a *API) GetUserSettings(ctx context.Context, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/settings")}
	return raw(ctx, a, req, opts)
}

func (a *API) ModifyUserSettings(ctx context.Context, params discord.ModifyUserSettingsParams, opts ...Option) (*discord.UserSettings, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/users/@me/settings")}
	req.JSON = params
	return call[discord.UserSettings](ctx, a, req, opts)
}

func (a *API) GetUserConsents(ctx context.Context, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/consent")}
	return raw(ctx, a, req, opts)
}

func (a *API) ModifyUserConsents(ctx context.Context, params discord.ModifyUserConsentsParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/consent")}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) GetEmailSettings(ctx context.Context, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/email-settings")}
	return raw(ctx, a, req, opts)
}

func (a *API) ModifyEmailSettings(ctx context.Context, params discord.ModifyEmailSettingsParams, opts ...Option) (*discord.EmailSettings, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/users/@me/email-settings")}
	req.JSON = params
	return call[discord.EmailSettings](ctx, a, req, opts)
}

func (a *API) ModifyNotificationSettings(ctx context.Context, params discord.ModifyNotificationSettingsParams, opts ...Option) (*discord.NotificationSettings, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/users/@me/notification-settings")}
	req.JSON = params
	return call[discord.NotificationSettings](ctx, a, req, opts)
}

func (a *API) ListNotificationSettingsSnapshots(ctx context.Context, opts ...Option) ([]discord.NotificationSettingsSnapshot, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/notification-settings/snapshots")}
	return list[discord.NotificationSettingsSnapshot](ctx, a, req, opts)
}

func (a *API) CreateNotificationSettingsSnapshot(ctx context.Context, params discord.CreateNotificationSettingsSnapshotParams, opts ...Option) ([]discord.NotificationSettingsSnapshot, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/notification-settings/snapshots")}
	req.JSON = params
	return list[discord.NotificationSettingsSnapshot](ctx, a, req, opts)
}

func (a *API) DeleteNotificationSettingsSnapshot(ctx context.Context, notificationSettingsSnapshotID snowflake.ID, opts ...Option) ([]discord.NotificationSettingsSnapshot, error) {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/users/@me/notification-settings/snapshots/{notification_settings_snapshot_id}", rest.P("notification_settings_snapshot_id", notificationSettingsSnapshotID))}
	return list[discord.NotificationSettingsSnapshot](ctx, a, req, opts)
}

func (a *API) ModifyUserGuildSettings(ctx context.Context, guildID snowflake.ID, opts ...Option) (*discord.UserGuildSettings, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/users/@me/guilds/{guild_id}/settings", rest.P("guild_id", guildID))}
	return call[discord.UserGuildSettings](ctx, a, req, opts)
}

func (a *API) BulkModifyUserGuildSettings(ctx context.Context, params discord.BulkModifyUserGuildSettingsParams, opts ...Option) ([]discord.UserGuildSettings, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/users/@me/guilds/settings")}
	req.JSON = params
	return list[discord.UserGuildSettings](ctx, a, req, opts)
}

func (a *API) ModifyAudioSettings(ctx context.Context, audioContextType string, userID snowflake.ID, params discord.ModifyAudioSettingsParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/users/@me/audio-settings/{audio_context_type}/{user_id}", rest.P("audio_context_type", audioContextType), rest.P("user_id", userID))}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) ListVideoFilterAssets(ctx context.Context, opts ...Option) ([]discord.VideoFilterAsset, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/video-filters/assets")}
	return list[discord.VideoFilterAsset](ctx, a, req, opts)
}

func (a *API) CreateVideoFilterAsset(ctx context.Context, params discord.CreateVideoFilterAssetParams, opts ...Option) (*discord.VideoFilterAsset, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/video-filters/assets")}
	req.JSON = params
	return call[discord.VideoFilterAsset](ctx, a, req, opts)
}

func (a *API) DeleteVideoFilterAsset(ctx context.Context, videoFilterAssetID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/users/@me/video-filters/assets/{video_filter_asset_id}", rest.P("video_filter_asset_id", videoFilterAssetID))}
	return none(ctx, a, req, opts)
}

func (a *API) MarkVideoFilterAssetLastUsed(ctx context.Context, videoFilterAssetID snowflake.ID, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/video-filters/assets/{video_filter_asset_id}/last-used", rest.P("video_filter_asset_id", videoFilterAssetID))}
	return raw(ctx, a, req, opts)
}

func (a *API) GetCurrentUser(ctx context.Context, query discord.GetCurrentUserQuery, opts ...Option) (*discord.User, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me")}
	req.Query = rest.Values(query)
	return call[discord.User](ctx, a, req, opts)
}

func (a *API) ModifyCurrentUser(ctx context.Context, params discord.ModifyCurrentUserParams, opts ...Option) (*discord.User, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/users/@me")}
	req.JSON = params
	return call[discord.User](ctx, a, req, opts)
}

func (a *API) ModifyCurrentUserAccount(ctx context.Context, params discord.ModifyCurrentUserAccountParams, opts ...Option) (*discord.User, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/users/@me/account")}
	req.JSON = params
	return call[discord.User](ctx, a, req, opts)
}

func (a *API) ReportMeaningfullyOnline(ctx context.Context, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/meaningfully-online")}
	return none(ctx, a, req, opts)
}

func (a *API) ListRecentAvatars(ctx context.Context, opts ...Option) (*discord.ListRecentAvatarsResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/avatars")}
	return call[discord.ListRecentAvatarsResponse](ctx, a, req, opts)
}

func (a *API) DeleteRecentAvatar(ctx context.Context, avatarID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/users/@me/avatars/{avatar_id}", rest.P("avatar_id", avatarID))}
	return none(ctx, a, req, opts)
}

func (a *API) GetUser(ctx context.Context, userID snowflake.ID, opts ...Option) (*discord.User, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/{user_id}", rest.P("user_id", userID))}
	return call[discord.User](ctx, a, req, opts)
}

func (a *API) GetUserProfile(ctx context.Context, userID snowflake.ID, query discord.GetUserProfileQuery, opts ...Option) (*discord.GetUserProfileResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/{user_id}/profile", rest.P("user_id", userID))}
	req.Query = rest.Values(query)
	return call[discord.GetUserProfileResponse](ctx, a, req, opts)
}

func (a *API) ModifyUserProfile(ctx context.Context, params discord.ModifyUserProfileParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/users/@me/profile")}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) ListMutualRelationships(ctx context.Context, userID snowflake.ID, opts ...Option) ([]discord.User, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/{user_id}/relationships", rest.P("user_id", userID))}
	return list[discord.User](ctx, a, req, opts)
}

func (a *API) EnableTotpMFA(ctx context.Context, params discord.EnableTotpMFAParams, opts ...Option) (*discord.EnableTotpMFAResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/mfa/totp/enable")}
	req.JSON = params
	return call[discord.EnableTotpMFAResponse](ctx, a, req, opts)
}

func (a *API) DisableTotpMFA(ctx context.Context, opts ...Option) (*discord.DisableTotpMFAResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/mfa/totp/disable")}
	return call[discord.DisableTotpMFAResponse](ctx, a, req, opts)
}

func (a *API) EnableSMSMFA(ctx context.Context, params discord.EnableSMSMFAParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/mfa/sms/enable")}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) DisableSMSMFA(ctx context.Context, params discord.DisableSMSMFAParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/mfa/sms/disable")}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) ListWebauthnAuthenticators(ctx context.Context, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/mfa/webauthn/credentials")}
	return raw(ctx, a, req, opts)
}

func (a *API) CreateWebauthnAuthenticator(ctx context.Context, params discord.CreateWebauthnAuthenticatorParams, opts ...Option) (*discord.CreateWebauthnAuthenticatorResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/mfa/webauthn/credentials")}
	req.JSON = params
	return call[discord.CreateWebauthnAuthenticatorResponse](ctx, a, req, opts)
}

func (a *API) ModifyWebauthnAuthenticator(ctx context.Context, authenticatorID snowflake.ID, params discord.ModifyWebauthnAuthenticatorParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/users/@me/mfa/webauthn/credentials/{authenticator_id}", rest.P("authenticator_id", authenticatorID))}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) DeleteWebauthnAuthenticator(ctx context.Context, authenticatorID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/users/@me/mfa/webauthn/credentials/{authenticator_id}", rest.P("authenticator_id", authenticatorID))}
	return none(ctx, a, req, opts)
}

func (a *API) SendBackupCodesChallenge(ctx context.Context, params discord.SendBackupCodesChallengeParams, opts ...Option) (*discord.SendBackupCodesChallengeResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/auth/verify/view-backup-codes-challenge")}
	req.JSON = params
	return call[discord.SendBackupCodesChallengeResponse](ctx, a, req, opts)
}

func (a *API) GetBackupCodes(ctx context.Context, params discord.GetBackupCodesParams, opts ...Option) (*discord.GetBackupCodesResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/mfa/codes-verification")}
	req.JSON = params
	return call[discord.GetBackupCodesResponse](ctx, a, req, opts)
}

func (a *API) DisableUserAccount(ctx context.Context, params discord.DisableUserAccountParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/disable")}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) DeleteUserAccount(ctx context.Context, params discord.DeleteUserAccountParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/delete")}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) VerifyUserCaptcha(ctx context.Context, params discord.VerifyUserCaptchaParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/captcha/verify")}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) ModifyUserAgreements(ctx context.Context, params discord.ModifyUserAgreementsParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/users/@me/agreements")}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) UserGetUniqueUsernameSuggestions(ctx context.Context, opts ...Option) (*discord.GetUniqueUsernameSuggestionsResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/pomelo-suggestions")}
	return call[discord.GetUniqueUsernameSuggestionsResponse](ctx, a, req, opts)
}

func (a *API) UserGetUniqueUsernameEligibility(ctx context.Context, params discord.GetUniqueUsernameEligibilityParams, opts ...Option) (*discord.GetUniqueUsernameEligibilityResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/pomelo-attempt")}
	req.JSON = params
	return call[discord.GetUniqueUsernameEligibilityResponse](ctx, a, req, opts)
}

func (a *API) CreateUniqueUsername(ctx context.Context, params discord.CreateUniqueUsernameParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/pomelo")}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) SetGuildIdentity(ctx context.Context, params discord.SetGuildIdentityParams, opts ...Option) (*discord.User, error) {
	req := &rest.Request{Route: rest.NewRoute("PUT", "/users/@me/clan")}
	req.JSON = params
	return call[discord.User](ctx, a, req, opts)
}

func (a *API) ListRecentMentions(ctx context.Context, query discord.ListRecentMentionsQuery, opts ...Option) ([]discord.Message, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/mentions")}
	req.Query = rest.Values(query)
	return list[discord.Message](ctx, a, req, opts)
}

func (a *API) DeleteRecentMention(ctx context.Context, messageID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/users/@me/mentions/{message_id}", rest.P("message_id", messageID))}
	return none(ctx, a, req, opts)
}

func (a *API) GetUserHarvest(ctx context.Context, opts ...Option) (*discord.Harvest, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/harvest")}
	return call[discord.Harvest](ctx, a, req, opts)
}

func (a *API) CreateUserHarvest(ctx context.Context, params discord.CreateUserHarvestParams, opts ...Option) (*discord.Harvest, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/harvest")}
	req.JSON = params
	return call[discord.Harvest](ctx, a, req, opts)
}

func (a *API) GetUserSurvey(ctx context.Context, query discord.GetUserSurveyQuery, opts ...Option) (*discord.GetUserSurveyResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/survey")}
	req.Query = rest.Values(query)
	return call[discord.GetUserSurveyResponse](ctx, a, req, opts)
}

func (a *API) AcknowledgeUserSurvey(ctx context.Context, surveyID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/survey/{survey_id}/seen", rest.P("survey_id", surveyID))}
	return none(ctx, a, req, opts)
}

func (a *API) GetUserNotes(ctx context.Context, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/notes")}
	return raw(ctx, a, req, opts)
}

func (a *API) GetUserNote(ctx context.Context, userID snowflake.ID, opts ...Option) (*discord.GetUserNoteResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/notes/{user_id}", rest.P("user_id", userID))}
	return call[discord.GetUserNoteResponse](ctx, a, req, opts)
}

func (a *API) ModifyUserNote(ctx context.Context, userID snowflake.ID, params discord.ModifyUserNoteParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("PUT", "/users/@me/notes/{user_id}", rest.P("user_id", userID))}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) ListUserAffinities(ctx context.Context, opts ...Option) (*discord.ListUserAffinitiesResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/affinities/users")}
	return call[discord.ListUserAffinitiesResponse](ctx, a, req, opts)
}

func (a *API) ListUserAffinitiesV2(ctx context.Context, opts ...Option) (*discord.ListUserAffinitiesV2Response, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/affinities/v2/users")}
	return call[discord.ListUserAffinitiesV2Response](ctx, a, req, opts)
}

func (a *API) ListGuildAffinities(ctx context.Context, opts ...Option) (*discord.ListGuildAffinitiesResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/affinities/guilds")}
	return call[discord.ListGuildAffinitiesResponse](ctx, a, req, opts)
}

func (a *API) ListChannelAffinities(ctx context.Context, opts ...Option) (*discord.ListChannelAffinitiesResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/affinities/channels")}
	return call[discord.ListChannelAffinitiesResponse](ctx, a, req, opts)
}

func (a *API) GetTutorial(ctx context.Context, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("GET", "/tutorial")}
	return none(ctx, a, req, opts)
}

func (a *API) ConfirmTutorialIndicator(ctx context.Context, indicator string, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("PUT", "/tutorial/indicators/{indicator}", rest.P("indicator", indicator))}
	return none(ctx, a, req, opts)
}

func (a *API) SuppressTutorial(ctx context.Context, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/tutorial/indicators/suppress")}
	return none(ctx, a, req, opts)
}

func (a *API) JoinHypesquadOnline(ctx context.Context, params discord.JoinHypesquadOnlineParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/hypesquad/online")}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) LeaveHypesquadOnline(ctx context.Context, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/hypesquad/online")}
	return none(ctx, a, req, opts)
}

func (a *API) SubmitDeveloperPortalCsatSurvey(ctx context.Context, params discord.SubmitDeveloperPortalCsatSurveyParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/dev-portal-csat-survey-response")}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) GetUserPremiumUsage(ctx context.Context, opts ...Option) (*discord.GetUserPremiumUsageResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/premium-usage")}
	return call[discord.GetUserPremiumUsageResponse](ctx, a, req, opts)
}

func (a *API) ListSavedMessages(ctx context.Context, opts ...Option) (*discord.ListSavedMessagesResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/saved-messages")}
	return call[discord.ListSavedMessagesResponse](ctx, a, req, opts)
}

func (a *API) SaveMessage(ctx context.Context, channelID snowflake.ID, messageID snowflake.ID, params discord.SaveMessageParams, opts ...Option) (*discord.SavedMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("PUT", "/users/@me/saved-messages/{channel_id}/{message_id}", rest.P("channel_id", channelID), rest.P("message_id", messageID))}
	req.JSON = params
	return call[discord.SavedMessage](ctx, a, req, opts)
}

func (a *API) UnsaveMessage(ctx context.Context, channelID snowflake.ID, messageID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/users/@me/saved-messages/{channel_id}/{message_id}", rest.P("channel_id", channelID), rest.P("message_id", messageID))}
	return none(ctx, a, req, opts)
}

func (a *API) VerifyAge(ctx context.Context, opts ...Option) (*discord.VerifyAgeResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/age-verification/verify")}
	return call[discord.VerifyAgeResponse](ctx, a, req, opts)
}

func (a *API) CreateUserIdentityVerification(ctx context.Context, params discord.CreateUserIdentityVerificationParams, opts ...Option) (*discord.UserIdentityVerification, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/identity/verification")}
	req.JSON = params
	return call[discord.UserIdentityVerification](ctx, a, req, opts)
}

func (a *API) GetUserIdentityVerification(ctx context.Context, opts ...Option) (*discord.UserIdentityVerification, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/identity/verification")}
	return call[discord.UserIdentityVerification](ctx, a, req, opts)
}

func (a *API) ListVoiceRegions(ctx context.Context, opts ...Option) ([]discord.VoiceRegion, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/voice/regions")}
	return list[discord.VoiceRegion](ctx, a, req, opts)
}

func (a *API) ListGuildVoiceRegions(ctx context.Context, guildID snowflake.ID, opts ...Option) ([]discord.VoiceRegion, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/regions", rest.P("guild_id", guildID))}
	return list[discord.VoiceRegion](ctx, a, req, opts)
}

func (a *API) UploadVoicePublicKey(ctx context.Context, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("PUT", "/voice/public-keys")}
	return none(ctx, a, req, opts)
}

func (a *API) VerifyVoicePublicKey(ctx context.Context, userID snowflake.ID, params discord.VerifyVoicePublicKeyParams, opts ...Option) (*discord.VerifyVoicePublicKeyResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/voice/{user_id}/match-public-key", rest.P("user_id", userID))}
	req.JSON = params
	return call[discord.VerifyVoicePublicKeyResponse](ctx, a, req, opts)
}

func (a *API) GetCurrentUserVoiceState(ctx context.Context, guildID snowflake.ID, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/voice-states/@me", rest.P("guild_id", guildID))}
	return raw(ctx, a, req, opts)
}

func (a *API) GetUserVoiceState(ctx context.Context, guildID snowflake.ID, userID snowflake.ID, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/voice-states/{user_id}", rest.P("guild_id", guildID), rest.P("user_id", userID))}
	return raw(ctx, a, req, opts)
}

func (a *API) ModifyCurrentUserVoiceState(ctx context.Context, guildID snowflake.ID, params discord.ModifyCurrentUserVoiceStateParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/guilds/{guild_id}/voice-states/@me", rest.P("guild_id", guildID))}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) ModifyUserVoiceState(ctx context.Context, guildID snowflake.ID, userID snowflake.ID, params discord.ModifyUserVoiceStateParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/guilds/{guild_id}/voice-states/{user_id}", rest.P("guild_id", guildID), rest.P("user_id", userID))}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) SendVoiceChannelEffect(ctx context.Context, channelID snowflake.ID, params discord.SendVoiceChannelEffectParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/channels/{channel_id}/voice-channel-effects", rest.P("channel_id", channelID))}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) SendCustomCallSound(ctx context.Context, channelID snowflake.ID, params discord.SendCustomCallSoundParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/channels/{channel_id}/custom-call-sounds", rest.P("channel_id", channelID))}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) ModifyStream(ctx context.Context, streamKey string, params discord.ModifyStreamParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/streams/{stream_key}/stream", rest.P("stream_key", streamKey))}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) GetStreamPreview(ctx context.Context, streamKey string, opts ...Option) (*discord.GetStreamPreviewResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/streams/{stream_key}/preview", rest.P("stream_key", streamKey))}
	return call[discord.GetStreamPreviewResponse](ctx, a, req, opts)
}

func (a *API) UploadStreamPreview(ctx context.Context, streamKey string, params discord.UploadStreamPreviewParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/streams/{stream_key}/preview", rest.P("stream_key", streamKey))}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) UploadVideoStreamPreview(ctx context.Context, streamKey string, params discord.UploadVideoStreamPreviewParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/streams/{stream_key}/preview/video", rest.P("stream_key", streamKey))}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) BroadcastStreamNotification(ctx context.Context, streamKey string, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/streams/{stream_key}/notify", rest.P("stream_key", streamKey))}
	return none(ctx, a, req, opts)
}

func (a *API) CreateWebhook(ctx context.Context, channelID snowflake.ID, params discord.CreateWebhookParams, opts ...Option) (*discord.Webhook, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/channels/{channel_id}/webhooks", rest.P("channel_id", channelID))}
	req.JSON = params
	return call[discord.Webhook](ctx, a, req, opts)
}

func (a *API) ListChannelWebhooks(ctx context.Context, channelID snowflake.ID, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/channels/{channel_id}/webhooks", rest.P("channel_id", channelID))}
	return raw(ctx, a, req, opts)
}

func (a *API) ListGuildWebhooks(ctx context.Context, guildID snowflake.ID, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/guilds/{guild_id}/webhooks", rest.P("guild_id", guildID))}
	return raw(ctx, a, req, opts)
}

func (a *API) GetWebhook(ctx context.Context, webhookID snowflake.ID, opts ...Option) (*discord.Webhook, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/webhooks/{webhook_id}", rest.P("webhook_id", webhookID))}
	return call[discord.Webhook](ctx, a, req, opts)
}

func (a *API) GetWebhookWithToken(ctx context.Context, webhookID snowflake.ID, webhookToken string, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/webhooks/{webhook_id}/{webhook_token}", rest.P("webhook_id", webhookID), rest.P("webhook_token", webhookToken))}
	return raw(ctx, a, req, opts)
}

func (a *API) ModifyWebhook(ctx context.Context, webhookID snowflake.ID, params discord.ModifyWebhookParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/webhooks/{webhook_id}", rest.P("webhook_id", webhookID))}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) ModifyWebhookWithToken(ctx context.Context, webhookID snowflake.ID, webhookToken string, params discord.ModifyWebhookWithTokenParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/webhooks/{webhook_id}/{webhook_token}", rest.P("webhook_id", webhookID), rest.P("webhook_token", webhookToken))}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) DeleteWebhook(ctx context.Context, webhookID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/webhooks/{webhook_id}", rest.P("webhook_id", webhookID))}
	return none(ctx, a, req, opts)
}

func (a *API) DeleteWebhookWithToken(ctx context.Context, webhookID snowflake.ID, webhookToken string, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/webhooks/{webhook_id}/{webhook_token}", rest.P("webhook_id", webhookID), rest.P("webhook_token", webhookToken))}
	return raw(ctx, a, req, opts)
}

func (a *API) ExecuteWebhook(ctx context.Context, webhookID snowflake.ID, webhookToken string, query discord.ExecuteWebhookQuery, params discord.ExecuteWebhookParams, opts ...Option) (*discord.Message, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/webhooks/{webhook_id}/{webhook_token}", rest.P("webhook_id", webhookID), rest.P("webhook_token", webhookToken))}
	req.Query = rest.Values(query)
	req.JSON = params
	return call[discord.Message](ctx, a, req, opts)
}

func (a *API) GetServiceWebhook(ctx context.Context, webhookID snowflake.ID, webhookToken string, webhookService string, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/webhooks/{webhook_id}/{webhook_token}/{webhook_service}", rest.P("webhook_id", webhookID), rest.P("webhook_token", webhookToken), rest.P("webhook_service", webhookService))}
	return raw(ctx, a, req, opts)
}

func (a *API) ExecuteServiceWebhook(ctx context.Context, webhookID snowflake.ID, webhookToken string, webhookService string, query discord.ExecuteServiceWebhookQuery, opts ...Option) (*discord.Message, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/webhooks/{webhook_id}/{webhook_token}/{webhook_service}", rest.P("webhook_id", webhookID), rest.P("webhook_token", webhookToken), rest.P("webhook_service", webhookService))}
	req.Query = rest.Values(query)
	return call[discord.Message](ctx, a, req, opts)
}

func (a *API) GetWebhookMessage(ctx context.Context, webhookID snowflake.ID, webhookToken string, messageID snowflake.ID, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/webhooks/{webhook_id}/{webhook_token}/messages/{message_id}", rest.P("webhook_id", webhookID), rest.P("webhook_token", webhookToken), rest.P("message_id", messageID))}
	return raw(ctx, a, req, opts)
}

func (a *API) EditWebhookMessage(ctx context.Context, webhookID snowflake.ID, webhookToken string, messageID snowflake.ID, params discord.EditWebhookMessageParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/webhooks/{webhook_id}/{webhook_token}/messages/{message_id}", rest.P("webhook_id", webhookID), rest.P("webhook_token", webhookToken), rest.P("message_id", messageID))}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) DeleteWebhookMessage(ctx context.Context, webhookID snowflake.ID, webhookToken string, messageID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/webhooks/{webhook_id}/{webhook_token}/messages/{message_id}", rest.P("webhook_id", webhookID), rest.P("webhook_token", webhookToken), rest.P("message_id", messageID))}
	return none(ctx, a, req, opts)
}

func (a *API) ModifyProfileWidgets(ctx context.Context, params discord.ModifyProfileWidgetsParams, opts ...Option) ([]discord.GameWidget, error) {
	req := &rest.Request{Route: rest.NewRoute("PUT", "/users/@me/widgets")}
	req.JSON = params
	return list[discord.GameWidget](ctx, a, req, opts)
}

func (a *API) GetProfileWidgetsSuggestedGames(ctx context.Context, opts ...Option) (*discord.GetProfileWidgetsSuggestedGamesResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/widgets/suggested-games")}
	return call[discord.GetProfileWidgetsSuggestedGamesResponse](ctx, a, req, opts)
}

func (a *API) ListBulkApplicationIdentities(ctx context.Context, params discord.WidgetsListBulkApplicationIdentitiesParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/application-identities")}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) ListUserApplicationIdentities(ctx context.Context, userID snowflake.ID, query discord.ListUserApplicationIdentitiesQuery, opts ...Option) (*discord.ListUserApplicationIdentitiesResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/{user_id}/application-identities", rest.P("user_id", userID))}
	req.Query = rest.Values(query)
	return call[discord.ListUserApplicationIdentitiesResponse](ctx, a, req, opts)
}

func (a *API) ModifyUserApplicationProfileConfig(ctx context.Context, applicationID snowflake.ID, externalUserID snowflake.ID, params discord.ModifyUserApplicationProfileConfigParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/users/@me/application-identities/{application_id}/{external_user_id}/config", rest.P("application_id", applicationID), rest.P("external_user_id", externalUserID))}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) ListFeaturedApplicationWidgetConfigs(ctx context.Context, opts ...Option) (*discord.ListFeaturedApplicationWidgetConfigsResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/widget-configs/featured")}
	return call[discord.ListFeaturedApplicationWidgetConfigsResponse](ctx, a, req, opts)
}

func (a *API) GetDeveloperApplicationWidgetConfigs(ctx context.Context, opts ...Option) (*discord.GetDeveloperApplicationWidgetConfigsResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/widget-configs/developer")}
	return call[discord.GetDeveloperApplicationWidgetConfigsResponse](ctx, a, req, opts)
}

func (a *API) GetLayoutDefinitions(ctx context.Context, opts ...Option) (*discord.GetLayoutDefinitionsResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/widget-configs/layout-definitions")}
	return call[discord.GetLayoutDefinitionsResponse](ctx, a, req, opts)
}

func (a *API) ListApplicationWidgetConfigs(ctx context.Context, applicationID snowflake.ID, opts ...Option) ([]discord.WidgetConfig, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/applications/{application_id}/widget-configs", rest.P("application_id", applicationID))}
	return list[discord.WidgetConfig](ctx, a, req, opts)
}

func (a *API) CreateApplicationWidgetConfig(ctx context.Context, applicationID snowflake.ID, params discord.CreateApplicationWidgetConfigParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/applications/{application_id}/widget-configs", rest.P("application_id", applicationID))}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) DeleteApplicationWidgetConfig(ctx context.Context, applicationID snowflake.ID, widgetConfigID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/applications/{application_id}/widget-configs/{widget_config_id}", rest.P("application_id", applicationID), rest.P("widget_config_id", widgetConfigID))}
	return none(ctx, a, req, opts)
}

func (a *API) ModifyApplicationWidgetConfig(ctx context.Context, applicationID snowflake.ID, widgetConfigID snowflake.ID, params discord.ModifyApplicationWidgetConfigParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("PATCH", "/applications/{application_id}/widget-configs/{widget_config_id}", rest.P("application_id", applicationID), rest.P("widget_config_id", widgetConfigID))}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) PublishApplicationWidgetConfig(ctx context.Context, applicationID snowflake.ID, widgetConfigID snowflake.ID, opts ...Option) (*discord.WidgetConfig, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/applications/{application_id}/widget-configs/{widget_config_id}/publish", rest.P("application_id", applicationID), rest.P("widget_config_id", widgetConfigID))}
	return call[discord.WidgetConfig](ctx, a, req, opts)
}

func (a *API) UnpublishApplicationWidgetConfig(ctx context.Context, applicationID snowflake.ID, widgetConfigID snowflake.ID, opts ...Option) (*discord.WidgetConfig, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/applications/{application_id}/widget-configs/{widget_config_id}/unpublish", rest.P("application_id", applicationID), rest.P("widget_config_id", widgetConfigID))}
	return call[discord.WidgetConfig](ctx, a, req, opts)
}

func (a *API) GetLatestApplicationInstaller(ctx context.Context, releaseChannel string, query discord.GetLatestApplicationInstallerQuery, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/download/{release_channel}", rest.P("release_channel", releaseChannel))}
	req.Query = rest.Values(query)
	return raw(ctx, a, req, opts)
}

func (a *API) GetApplicationUpdates(ctx context.Context, releaseChannel string, query discord.GetApplicationUpdatesQuery, opts ...Option) (*discord.GetApplicationUpdatesResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/updates/{release_channel}", rest.P("release_channel", releaseChannel))}
	req.Query = rest.Values(query)
	return call[discord.GetApplicationUpdatesResponse](ctx, a, req, opts)
}

func (a *API) GetNativeModuleVersions(ctx context.Context, releaseChannel string, query discord.GetNativeModuleVersionsQuery, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/modules/{release_channel}/versions.json", rest.P("release_channel", releaseChannel))}
	req.Query = rest.Values(query)
	return raw(ctx, a, req, opts)
}

func (a *API) GetNativeModule(ctx context.Context, releaseChannel string, moduleName string, moduleVersion string, query discord.GetNativeModuleQuery, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/modules/{release_channel}/{module_name}/{module_version}", rest.P("release_channel", releaseChannel), rest.P("module_name", moduleName), rest.P("module_version", moduleVersion))}
	req.Query = rest.Values(query)
	return raw(ctx, a, req, opts)
}

func (a *API) GetLatestDistributedApplicationInstaller(ctx context.Context, query discord.GetLatestDistributedApplicationInstallerQuery, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/downloads/distributions/app/installers/latest")}
	req.Query = rest.Values(query)
	return raw(ctx, a, req, opts)
}

func (a *API) GetLatestDistributedApplicationManifest(ctx context.Context, query discord.GetLatestDistributedApplicationManifestQuery, opts ...Option) (*discord.GetLatestDistributedApplicationManifestResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/updates/distributions/app/manifests/latest")}
	req.Query = rest.Values(query)
	return call[discord.GetLatestDistributedApplicationManifestResponse](ctx, a, req, opts)
}

func (a *API) CreateMessageAttachments(ctx context.Context, channelID snowflake.ID, params discord.CreateMessageAttachmentsParams, opts ...Option) ([]discord.CloudAttachment, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/channels/{channel_id}/attachments", rest.P("channel_id", channelID))}
	req.JSON = params
	return list[discord.CloudAttachment](ctx, a, req, opts)
}

func (a *API) CreateGuildProductAttachments(ctx context.Context, channelID snowflake.ID, params discord.CreateGuildProductAttachmentsParams, opts ...Option) ([]discord.CloudAttachment, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/guilds/{channel_id}/products/attachments", rest.P("channel_id", channelID))}
	req.JSON = params
	return list[discord.CloudAttachment](ctx, a, req, opts)
}

func (a *API) CreateGravityAttachments(ctx context.Context, params discord.CreateGravityAttachmentsParams, opts ...Option) ([]discord.CloudAttachment, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/gravity-attachments")}
	req.JSON = params
	return list[discord.CloudAttachment](ctx, a, req, opts)
}

func (a *API) DeleteAttachment(ctx context.Context, cloudAttachmentUploadFilename string, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/attachments/{cloud_attachment_upload_filename}", rest.Raw("cloud_attachment_upload_filename", cloudAttachmentUploadFilename))}
	return none(ctx, a, req, opts)
}

func (a *API) RefreshAttachmentURLs(ctx context.Context, params discord.RefreshAttachmentURLsParams, opts ...Option) (*discord.RefreshAttachmentURLsResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/attachments/refresh-urls")}
	req.JSON = params
	return call[discord.RefreshAttachmentURLsResponse](ctx, a, req, opts)
}

func (a *API) VerifyEmailAddress(ctx context.Context, params discord.VerifyEmailAddressParams, opts ...Option) (*discord.VerifyEmailAddressResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/auth/verify")}
	req.JSON = params
	return call[discord.VerifyEmailAddressResponse](ctx, a, req, opts)
}

func (a *API) ResendVerificationEmail(ctx context.Context, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/auth/verify/resend")}
	return none(ctx, a, req, opts)
}

func (a *API) SendEmailChangeChallenge(ctx context.Context, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("PUT", "/users/@me/email")}
	return none(ctx, a, req, opts)
}

func (a *API) VerifyEmailChangeCode(ctx context.Context, params discord.VerifyEmailChangeCodeParams, opts ...Option) (*discord.VerifyEmailChangeCodeResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/email/verify-code")}
	req.JSON = params
	return call[discord.VerifyEmailChangeCodeResponse](ctx, a, req, opts)
}

func (a *API) GetExperimentAssignments(ctx context.Context, opts ...Option) (*discord.GetExperimentAssignmentsResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/experiments")}
	return call[discord.GetExperimentAssignmentsResponse](ctx, a, req, opts)
}

func (a *API) CreateFingerprint(ctx context.Context, opts ...Option) (*discord.CreateFingerprintResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/auth/fingerprint")}
	return call[discord.CreateFingerprintResponse](ctx, a, req, opts)
}

func (a *API) GetApexExperimentAssignments(ctx context.Context, opts ...Option) (*discord.ApexExperiments, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/apex/experiments")}
	return call[discord.ApexExperiments](ctx, a, req, opts)
}

func (a *API) GetMetadataForApexExperiments(ctx context.Context, opts ...Option) (*discord.GetMetadataForApexExperimentsResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/apex/experiments/metadata")}
	return call[discord.GetMetadataForApexExperimentsResponse](ctx, a, req, opts)
}

func (a *API) GetCurrentAuthorizationInformation(ctx context.Context, opts ...Option) (*discord.GetCurrentAuthorizationInformationResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/oauth2/@me")}
	return call[discord.GetCurrentAuthorizationInformationResponse](ctx, a, req, opts)
}

func (a *API) GetOpenidConnectKeys(ctx context.Context, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/oauth2/keys")}
	return raw(ctx, a, req, opts)
}

func (a *API) GetOpenidUserInformation(ctx context.Context, opts ...Option) (*discord.GetOpenidUserInformationResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/oauth2/userinfo")}
	return call[discord.GetOpenidUserInformationResponse](ctx, a, req, opts)
}

func (a *API) GetOAuth2DeviceCode(ctx context.Context, params discord.GetOAuth2DeviceCodeParams, opts ...Option) (*discord.GetOAuth2DeviceCodeResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/oauth2/device/authorize")}
	req.JSON = params
	return call[discord.GetOAuth2DeviceCodeResponse](ctx, a, req, opts)
}

func (a *API) GetOAuth2DeviceFlow(ctx context.Context, params discord.GetOAuth2DeviceFlowParams, opts ...Option) (*discord.GetOAuth2DeviceFlowResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/oauth2/device/verify")}
	req.JSON = params
	return call[discord.GetOAuth2DeviceFlowResponse](ctx, a, req, opts)
}

func (a *API) FinishOAuth2DeviceFlow(ctx context.Context, params discord.FinishOAuth2DeviceFlowParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/oauth2/device/finish")}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) GetOAuth2Token(ctx context.Context, params discord.GetOAuth2TokenParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/oauth2/token")}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) RevokeOAuth2Token(ctx context.Context, params discord.RevokeOAuth2TokenParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/oauth2/token/revoke")}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) GetProvisionalAccountToken(ctx context.Context, params discord.GetProvisionalAccountTokenParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/partner-sdk/token")}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) ExchangeProvisionalAccountChildToken(ctx context.Context, params discord.ExchangeProvisionalAccountChildTokenParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/partner-sdk/child-token")}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) GetProvisionalAccountTokenWithBot(ctx context.Context, params discord.GetProvisionalAccountTokenWithBotParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/partner-sdk/token/bot")}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) UnmergeProvisionalAccount(ctx context.Context, params discord.UnmergeProvisionalAccountParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/partner-sdk/provisional-accounts/unmerge")}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) UnmergeProvisionalAccountWithBot(ctx context.Context, params discord.UnmergeProvisionalAccountWithBotParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/partner-sdk/provisional-accounts/unmerge/bot")}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) GetOAuth2Authorizations(ctx context.Context, query discord.GetOAuth2AuthorizationsQuery, opts ...Option) ([]discord.OAuth2Authorization, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/oauth2/tokens")}
	req.Query = rest.Values(query)
	return list[discord.OAuth2Authorization](ctx, a, req, opts)
}

func (a *API) GetApplicationOAuth2Authorizations(ctx context.Context, applicationID snowflake.ID, opts ...Option) ([]discord.OAuth2Authorization, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/oauth2/applications/{application_id}/tokens", rest.P("application_id", applicationID))}
	return list[discord.OAuth2Authorization](ctx, a, req, opts)
}

func (a *API) PreviewOAuth2Authorization(ctx context.Context, query discord.PreviewOAuth2AuthorizationQuery, opts ...Option) (*discord.PreviewOAuth2AuthorizationResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/oauth2/authorize")}
	req.Query = rest.Values(query)
	return call[discord.PreviewOAuth2AuthorizationResponse](ctx, a, req, opts)
}

func (a *API) CreateOAuth2Authorization(ctx context.Context, query discord.CreateOAuth2AuthorizationQuery, params discord.CreateOAuth2AuthorizationParams, opts ...Option) (*discord.CreateOAuth2AuthorizationResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/oauth2/authorize")}
	req.Query = rest.Values(query)
	req.JSON = params
	return call[discord.CreateOAuth2AuthorizationResponse](ctx, a, req, opts)
}

func (a *API) GetOAuth2Authorization(ctx context.Context, tokenID snowflake.ID, opts ...Option) (*discord.OAuth2Authorization, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/oauth2/tokens/{token_id}", rest.P("token_id", tokenID))}
	return call[discord.OAuth2Authorization](ctx, a, req, opts)
}

func (a *API) DeleteOAuth2Authorization(ctx context.Context, tokenID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/oauth2/tokens/{token_id}", rest.P("token_id", tokenID))}
	return none(ctx, a, req, opts)
}

func (a *API) VerifyPhoneNumber(ctx context.Context, params discord.VerifyPhoneNumberParams, opts ...Option) (*discord.VerifyPhoneNumberResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/phone-verifications/verify")}
	req.JSON = params
	return call[discord.VerifyPhoneNumberResponse](ctx, a, req, opts)
}

func (a *API) ResendVerificationCode(ctx context.Context, params discord.ResendVerificationCodeParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/phone-verifications/resend")}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) AddPhoneNumber(ctx context.Context, params discord.AddPhoneNumberParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/phone")}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) ReverifyPhoneNumber(ctx context.Context, params discord.ReverifyPhoneNumberParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/phone/reverify")}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) RemovePhoneNumber(ctx context.Context, params discord.RemovePhoneNumberParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/users/@me/phone")}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) RegisterDevice(ctx context.Context, params discord.RegisterDeviceParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/devices")}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) UnregisterDevice(ctx context.Context, params discord.UnregisterDeviceParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/users/@me/devices")}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) GetDeviceSyncToken(ctx context.Context, opts ...Option) (*discord.GetDeviceSyncTokenResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/users/@me/devices/sync-token")}
	return call[discord.GetDeviceSyncTokenResponse](ctx, a, req, opts)
}

func (a *API) SyncDevices(ctx context.Context, params discord.SyncDevicesParams, opts ...Option) (*discord.SyncDevicesResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("PUT", "/users/@me/devices/sync")}
	req.JSON = params
	return call[discord.SyncDevicesResponse](ctx, a, req, opts)
}

func (a *API) AcknowledgeMessage(ctx context.Context, channelID snowflake.ID, messageID snowflake.ID, params discord.AcknowledgeMessageParams, opts ...Option) (*discord.AcknowledgeMessageResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/channels/{channel_id}/messages/{message_id}/ack", rest.P("channel_id", channelID), rest.P("message_id", messageID))}
	req.JSON = params
	return call[discord.AcknowledgeMessageResponse](ctx, a, req, opts)
}

func (a *API) AcknowledgePinnedMessages(ctx context.Context, channelID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/channels/{channel_id}/pins/ack", rest.P("channel_id", channelID))}
	return none(ctx, a, req, opts)
}

func (a *API) AcknowledgeGuild(ctx context.Context, guildID snowflake.ID, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/guilds/{guild_id}/ack", rest.P("guild_id", guildID))}
	return raw(ctx, a, req, opts)
}

func (a *API) AcknowledgeGuildFeature(ctx context.Context, guildID snowflake.ID, readStateType string, entityID snowflake.ID, params discord.AcknowledgeGuildFeatureParams, opts ...Option) (*discord.AcknowledgeGuildFeatureResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/guilds/{guild_id}/ack/{read_state_type}/{entity_id}", rest.P("guild_id", guildID), rest.P("read_state_type", readStateType), rest.P("entity_id", entityID))}
	req.JSON = params
	return call[discord.AcknowledgeGuildFeatureResponse](ctx, a, req, opts)
}

func (a *API) AcknowledgeUserFeature(ctx context.Context, readStateType string, entityID snowflake.ID, params discord.AcknowledgeUserFeatureParams, opts ...Option) (*discord.AcknowledgeUserFeatureResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/users/@me/{read_state_type}/{entity_id}/ack", rest.P("read_state_type", readStateType), rest.P("entity_id", entityID))}
	req.JSON = params
	return call[discord.AcknowledgeUserFeatureResponse](ctx, a, req, opts)
}

func (a *API) BulkUpdateReadStates(ctx context.Context, params discord.BulkUpdateReadStatesParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/read-states/ack-bulk")}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) DeleteReadState(ctx context.Context, readStateID snowflake.ID, params discord.DeleteReadStateParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/channels/{read_state_id}/messages/ack", rest.P("read_state_id", readStateID))}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) GetReportReasons(ctx context.Context, query discord.GetReportReasonsQuery, opts ...Option) ([]discord.ReportReason, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/report")}
	req.Query = rest.Values(query)
	return list[discord.ReportReason](ctx, a, req, opts)
}

func (a *API) CreateReport(ctx context.Context, params discord.CreateReportParams, opts ...Option) (*discord.CreateReportResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/report")}
	req.JSON = params
	return call[discord.CreateReportResponse](ctx, a, req, opts)
}

func (a *API) GetReportOptions(ctx context.Context, opts ...Option) ([]discord.ReportOption, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/report/options")}
	return list[discord.ReportOption](ctx, a, req, opts)
}

func (a *API) StageReport(ctx context.Context, channelID snowflake.ID, messageID snowflake.ID, opts ...Option) (*discord.StageReportResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/reports/channels/{channel_id}/messages/{message_id}", rest.P("channel_id", channelID), rest.P("message_id", messageID))}
	return call[discord.StageReportResponse](ctx, a, req, opts)
}

func (a *API) CreateStagedReport(ctx context.Context, params discord.CreateStagedReportParams, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("POST", "/reports")}
	req.JSON = params
	return none(ctx, a, req, opts)
}

func (a *API) GetReportMenu(ctx context.Context, typeArg string, query discord.GetReportMenuQuery, opts ...Option) (*discord.ReportMenu, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/reporting/menu/{type}", rest.P("type", typeArg))}
	req.Query = rest.Values(query)
	return call[discord.ReportMenu](ctx, a, req, opts)
}

func (a *API) SubmitReportMenu(ctx context.Context, typeArg string, params discord.SubmitReportMenuParams, opts ...Option) (*discord.SubmitReportMenuResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/reporting/{type}", rest.P("type", typeArg))}
	req.JSON = params
	return call[discord.SubmitReportMenuResponse](ctx, a, req, opts)
}

func (a *API) QueryUnauthenticatedReportEligibility(ctx context.Context, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/reporting/unauthenticated/experiment")}
	return raw(ctx, a, req, opts)
}

func (a *API) GetUnauthenticatedReportCapabilities(ctx context.Context, opts ...Option) (*discord.GetUnauthenticatedReportCapabilitiesResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/reporting/unauthenticated/capabilities")}
	return call[discord.GetUnauthenticatedReportCapabilitiesResponse](ctx, a, req, opts)
}

func (a *API) GetUnauthenticatedReportVerificationCode(ctx context.Context, typeArg string, params discord.GetUnauthenticatedReportVerificationCodeParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/reporting/unauthenticated/{type}/code", rest.P("type", typeArg))}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) VerifyUnauthenticatedReport(ctx context.Context, typeArg string, params discord.VerifyUnauthenticatedReportParams, opts ...Option) (*discord.VerifyUnauthenticatedReportResponse, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/reporting/unauthenticated/{type}/verify", rest.P("type", typeArg))}
	req.JSON = params
	return call[discord.VerifyUnauthenticatedReportResponse](ctx, a, req, opts)
}

func (a *API) GetUnauthenticatedReportMenu(ctx context.Context, typeArg string, query discord.GetUnauthenticatedReportMenuQuery, opts ...Option) (*discord.ReportMenu, error) {
	req := &rest.Request{Route: rest.NewRoute("GET", "/reporting/unauthenticated/menu/{type}", rest.P("type", typeArg))}
	req.Query = rest.Values(query)
	return call[discord.ReportMenu](ctx, a, req, opts)
}

func (a *API) SubmitUnauthenticatedReportMenu(ctx context.Context, typeArg string, params discord.SubmitUnauthenticatedReportMenuParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/reporting/unauthenticated/{type}", rest.P("type", typeArg))}
	req.JSON = params
	return raw(ctx, a, req, opts)
}

func (a *API) RequestReportReview(ctx context.Context, params discord.RequestReportReviewParams, opts ...Option) (json.RawMessage, error) {
	req := &rest.Request{Route: rest.NewRoute("POST", "/reporting/review")}
	req.JSON = params
	return raw(ctx, a, req, opts)
}
