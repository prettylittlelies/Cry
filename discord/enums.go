package discord

import "strconv"

type AccessibilityFeatureFlags uint64

const (
	AccessibilityFeatureFlagsScreenreader                   AccessibilityFeatureFlags = 1 << 0
	AccessibilityFeatureFlagsReducedMotion                  AccessibilityFeatureFlags = 1 << 1
	AccessibilityFeatureFlagsReducedTransparency            AccessibilityFeatureFlags = 1 << 2
	AccessibilityFeatureFlagsHighContrast                   AccessibilityFeatureFlags = 1 << 3
	AccessibilityFeatureFlagsBoldText                       AccessibilityFeatureFlags = 1 << 4
	AccessibilityFeatureFlagsGrayscale                      AccessibilityFeatureFlags = 1 << 5
	AccessibilityFeatureFlagsInvertColors                   AccessibilityFeatureFlags = 1 << 6
	AccessibilityFeatureFlagsPrefersColorSchemeLight        AccessibilityFeatureFlags = 1 << 7
	AccessibilityFeatureFlagsPrefersColorSchemeDark         AccessibilityFeatureFlags = 1 << 8
	AccessibilityFeatureFlagsChatFontScaleIncreased         AccessibilityFeatureFlags = 1 << 9
	AccessibilityFeatureFlagsChatFontScaleDecreased         AccessibilityFeatureFlags = 1 << 10
	AccessibilityFeatureFlagsZoomLevelIncreased             AccessibilityFeatureFlags = 1 << 11
	AccessibilityFeatureFlagsZoomLevelDecreased             AccessibilityFeatureFlags = 1 << 12
	AccessibilityFeatureFlagsMessageGroupSpacingIncreased   AccessibilityFeatureFlags = 1 << 13
	AccessibilityFeatureFlagsMessageGroupSpacingDecreased   AccessibilityFeatureFlags = 1 << 14
	AccessibilityFeatureFlagsDarkSidebar                    AccessibilityFeatureFlags = 1 << 15
	AccessibilityFeatureFlagsReducedMotionFromUserSettings  AccessibilityFeatureFlags = 1 << 16
	AccessibilityFeatureFlagsSaturationLevelDecreased       AccessibilityFeatureFlags = 1 << 17
	AccessibilityFeatureFlagsForcedColors                   AccessibilityFeatureFlags = 1 << 18
	AccessibilityFeatureFlagsForcedColorsFromUserSettings   AccessibilityFeatureFlags = 1 << 19
	AccessibilityFeatureFlagsRoleStyleAdjusted              AccessibilityFeatureFlags = 1 << 20
	AccessibilityFeatureFlagsSyncProfileThemeWithUserTheme  AccessibilityFeatureFlags = 1 << 21
	AccessibilityFeatureFlagsReducedMotionPrefersCrossfades AccessibilityFeatureFlags = 1 << 22
	AccessibilityFeatureFlagsContrastLevelIncreased         AccessibilityFeatureFlags = 1 << 23
	AccessibilityFeatureFlagsContrastLevelDecreased         AccessibilityFeatureFlags = 1 << 24
)

func (v AccessibilityFeatureFlags) String() string {
	switch v {
	case AccessibilityFeatureFlagsScreenreader:
		return "SCREENREADER"
	case AccessibilityFeatureFlagsReducedMotion:
		return "REDUCED_MOTION"
	case AccessibilityFeatureFlagsReducedTransparency:
		return "REDUCED_TRANSPARENCY"
	case AccessibilityFeatureFlagsHighContrast:
		return "HIGH_CONTRAST"
	case AccessibilityFeatureFlagsBoldText:
		return "BOLD_TEXT"
	case AccessibilityFeatureFlagsGrayscale:
		return "GRAYSCALE"
	case AccessibilityFeatureFlagsInvertColors:
		return "INVERT_COLORS"
	case AccessibilityFeatureFlagsPrefersColorSchemeLight:
		return "PREFERS_COLOR_SCHEME_LIGHT"
	case AccessibilityFeatureFlagsPrefersColorSchemeDark:
		return "PREFERS_COLOR_SCHEME_DARK"
	case AccessibilityFeatureFlagsChatFontScaleIncreased:
		return "CHAT_FONT_SCALE_INCREASED"
	case AccessibilityFeatureFlagsChatFontScaleDecreased:
		return "CHAT_FONT_SCALE_DECREASED"
	case AccessibilityFeatureFlagsZoomLevelIncreased:
		return "ZOOM_LEVEL_INCREASED"
	case AccessibilityFeatureFlagsZoomLevelDecreased:
		return "ZOOM_LEVEL_DECREASED"
	case AccessibilityFeatureFlagsMessageGroupSpacingIncreased:
		return "MESSAGE_GROUP_SPACING_INCREASED"
	case AccessibilityFeatureFlagsMessageGroupSpacingDecreased:
		return "MESSAGE_GROUP_SPACING_DECREASED"
	case AccessibilityFeatureFlagsDarkSidebar:
		return "DARK_SIDEBAR"
	case AccessibilityFeatureFlagsReducedMotionFromUserSettings:
		return "REDUCED_MOTION_FROM_USER_SETTINGS"
	case AccessibilityFeatureFlagsSaturationLevelDecreased:
		return "SATURATION_LEVEL_DECREASED"
	case AccessibilityFeatureFlagsForcedColors:
		return "FORCED_COLORS"
	case AccessibilityFeatureFlagsForcedColorsFromUserSettings:
		return "FORCED_COLORS_FROM_USER_SETTINGS"
	case AccessibilityFeatureFlagsRoleStyleAdjusted:
		return "ROLE_STYLE_ADJUSTED"
	case AccessibilityFeatureFlagsSyncProfileThemeWithUserTheme:
		return "SYNC_PROFILE_THEME_WITH_USER_THEME"
	case AccessibilityFeatureFlagsReducedMotionPrefersCrossfades:
		return "REDUCED_MOTION_PREFERS_CROSSFADES"
	case AccessibilityFeatureFlagsContrastLevelIncreased:
		return "CONTRAST_LEVEL_INCREASED"
	case AccessibilityFeatureFlagsContrastLevelDecreased:
		return "CONTRAST_LEVEL_DECREASED"
	}
	return strconv.FormatUint(uint64(v), 10)
}

type AccountStandingState int64

const (
	AccountStandingStateAllGood     AccountStandingState = 100
	AccountStandingStateLimited     AccountStandingState = 200
	AccountStandingStateVeryLimited AccountStandingState = 300
	AccountStandingStateAtRisk      AccountStandingState = 400
	AccountStandingStateSuspended   AccountStandingState = 500
)

func (v AccountStandingState) String() string {
	switch v {
	case AccountStandingStateAllGood:
		return "ALL_GOOD"
	case AccountStandingStateLimited:
		return "LIMITED"
	case AccountStandingStateVeryLimited:
		return "VERY_LIMITED"
	case AccountStandingStateAtRisk:
		return "AT_RISK"
	case AccountStandingStateSuspended:
		return "SUSPENDED"
	}
	return strconv.FormatInt(int64(v), 10)
}

type ActionType int64

const (
	ActionTypeBan                       ActionType = 0
	ActionTypeTempBan                   ActionType = 1
	ActionTypeGlobalQuarantine          ActionType = 2
	ActionTypeRequireVerification       ActionType = 3
	ActionTypeUserWarning               ActionType = 4
	ActionTypeUserSpammer               ActionType = 5
	ActionTypeChannelSpam               ActionType = 6
	ActionTypeMessageSpam               ActionType = 7
	ActionTypeDisableSuspiciousActivity ActionType = 8
	ActionTypeLimitedAccess             ActionType = 9
	ActionTypeChannelScheduleDelete     ActionType = 10
	ActionTypeMessageContentRemoval     ActionType = 11
	ActionTypeGuildDisableInvite        ActionType = 12
	ActionTypeUserContentRemoval        ActionType = 13
	ActionTypeUserUsernameMangled       ActionType = 14
	ActionTypeGuildLimitedAccess        ActionType = 15
	ActionTypeUserMessageRemoval        ActionType = 16
	ActionTypeGuildDelete               ActionType = 20
	ActionTypeUserProfileMangled        ActionType = 22
)

func (v ActionType) String() string {
	switch v {
	case ActionTypeBan:
		return "BAN"
	case ActionTypeTempBan:
		return "TEMP_BAN"
	case ActionTypeGlobalQuarantine:
		return "GLOBAL_QUARANTINE"
	case ActionTypeRequireVerification:
		return "REQUIRE_VERIFICATION"
	case ActionTypeUserWarning:
		return "USER_WARNING"
	case ActionTypeUserSpammer:
		return "USER_SPAMMER"
	case ActionTypeChannelSpam:
		return "CHANNEL_SPAM"
	case ActionTypeMessageSpam:
		return "MESSAGE_SPAM"
	case ActionTypeDisableSuspiciousActivity:
		return "DISABLE_SUSPICIOUS_ACTIVITY"
	case ActionTypeLimitedAccess:
		return "LIMITED_ACCESS"
	case ActionTypeChannelScheduleDelete:
		return "CHANNEL_SCHEDULE_DELETE"
	case ActionTypeMessageContentRemoval:
		return "MESSAGE_CONTENT_REMOVAL"
	case ActionTypeGuildDisableInvite:
		return "GUILD_DISABLE_INVITE"
	case ActionTypeUserContentRemoval:
		return "USER_CONTENT_REMOVAL"
	case ActionTypeUserUsernameMangled:
		return "USER_USERNAME_MANGLED"
	case ActionTypeGuildLimitedAccess:
		return "GUILD_LIMITED_ACCESS"
	case ActionTypeUserMessageRemoval:
		return "USER_MESSAGE_REMOVAL"
	case ActionTypeGuildDelete:
		return "GUILD_DELETE"
	case ActionTypeUserProfileMangled:
		return "USER_PROFILE_MANGLED"
	}
	return strconv.FormatInt(int64(v), 10)
}

type ActivityActionType int64

const (
	ActivityActionTypeJoin        ActivityActionType = 1
	ActivityActionTypeSpectate    ActivityActionType = 2
	ActivityActionTypeListen      ActivityActionType = 3
	ActivityActionTypeWatch       ActivityActionType = 4
	ActivityActionTypeJoinRequest ActivityActionType = 5
)

func (v ActivityActionType) String() string {
	switch v {
	case ActivityActionTypeJoin:
		return "JOIN"
	case ActivityActionTypeSpectate:
		return "SPECTATE"
	case ActivityActionTypeListen:
		return "LISTEN"
	case ActivityActionTypeWatch:
		return "WATCH"
	case ActivityActionTypeJoinRequest:
		return "JOIN_REQUEST"
	}
	return strconv.FormatInt(int64(v), 10)
}

type ActivityFlags uint64

const (
	ActivityFlagsInstance                 ActivityFlags = 1 << 0
	ActivityFlagsJoin                     ActivityFlags = 1 << 1
	ActivityFlagsSpectate                 ActivityFlags = 1 << 2
	ActivityFlagsJoinRequest              ActivityFlags = 1 << 3
	ActivityFlagsSync                     ActivityFlags = 1 << 4
	ActivityFlagsPlay                     ActivityFlags = 1 << 5
	ActivityFlagsPartyPrivacyFriends      ActivityFlags = 1 << 6
	ActivityFlagsPartyPrivacyVoiceChannel ActivityFlags = 1 << 7
	ActivityFlagsEmbedded                 ActivityFlags = 1 << 8
	ActivityFlagsContextless              ActivityFlags = 1 << 9
)

func (v ActivityFlags) String() string {
	switch v {
	case ActivityFlagsInstance:
		return "INSTANCE"
	case ActivityFlagsJoin:
		return "JOIN"
	case ActivityFlagsSpectate:
		return "SPECTATE"
	case ActivityFlagsJoinRequest:
		return "JOIN_REQUEST"
	case ActivityFlagsSync:
		return "SYNC"
	case ActivityFlagsPlay:
		return "PLAY"
	case ActivityFlagsPartyPrivacyFriends:
		return "PARTY_PRIVACY_FRIENDS"
	case ActivityFlagsPartyPrivacyVoiceChannel:
		return "PARTY_PRIVACY_VOICE_CHANNEL"
	case ActivityFlagsEmbedded:
		return "EMBEDDED"
	case ActivityFlagsContextless:
		return "CONTEXTLESS"
	}
	return strconv.FormatUint(uint64(v), 10)
}

type ActivityLinkType int64

const (
	ActivityLinkTypeManagedLink ActivityLinkType = 0
	ActivityLinkTypeQuickLink   ActivityLinkType = 1
)

func (v ActivityLinkType) String() string {
	switch v {
	case ActivityLinkTypeManagedLink:
		return "MANAGED_LINK"
	case ActivityLinkTypeQuickLink:
		return "QUICK_LINK"
	}
	return strconv.FormatInt(int64(v), 10)
}

type ActivityPartyPrivacy int64

const (
	ActivityPartyPrivacyPrivate ActivityPartyPrivacy = 0
	ActivityPartyPrivacyPublic  ActivityPartyPrivacy = 1
)

func (v ActivityPartyPrivacy) String() string {
	switch v {
	case ActivityPartyPrivacyPrivate:
		return "PRIVATE"
	case ActivityPartyPrivacyPublic:
		return "PUBLIC"
	}
	return strconv.FormatInt(int64(v), 10)
}

type ActivityPlatformType string

const (
	ActivityPlatformTypeDesktop   ActivityPlatformType = "desktop"
	ActivityPlatformTypeXbox      ActivityPlatformType = "xbox"
	ActivityPlatformTypeSamsung   ActivityPlatformType = "samsung"
	ActivityPlatformTypeIos       ActivityPlatformType = "ios"
	ActivityPlatformTypeAndroid   ActivityPlatformType = "android"
	ActivityPlatformTypeEmbedded  ActivityPlatformType = "embedded"
	ActivityPlatformTypePs4       ActivityPlatformType = "ps4"
	ActivityPlatformTypePs5       ActivityPlatformType = "ps5"
	ActivityPlatformTypeMetaQuest ActivityPlatformType = "meta_quest"
)

type ActivityType int64

const (
	ActivityTypePlaying   ActivityType = 0
	ActivityTypeStreaming ActivityType = 1
	ActivityTypeListening ActivityType = 2
	ActivityTypeWatching  ActivityType = 3
	ActivityTypeCustom    ActivityType = 4
	ActivityTypeCompeting ActivityType = 5
	ActivityTypeHang      ActivityType = 6
)

func (v ActivityType) String() string {
	switch v {
	case ActivityTypePlaying:
		return "PLAYING"
	case ActivityTypeStreaming:
		return "STREAMING"
	case ActivityTypeListening:
		return "LISTENING"
	case ActivityTypeWatching:
		return "WATCHING"
	case ActivityTypeCustom:
		return "CUSTOM"
	case ActivityTypeCompeting:
		return "COMPETING"
	case ActivityTypeHang:
		return "HANG"
	}
	return strconv.FormatInt(int64(v), 10)
}

type AdCreativeType int64

const (
	AdCreativeTypeQuest         AdCreativeType = 1
	AdCreativeTypeQuestHomeHero AdCreativeType = 2
	AdCreativeTypeBounty        AdCreativeType = 3
)

func (v AdCreativeType) String() string {
	switch v {
	case AdCreativeTypeQuest:
		return "QUEST"
	case AdCreativeTypeQuestHomeHero:
		return "QUEST_HOME_HERO"
	case AdCreativeTypeBounty:
		return "BOUNTY"
	}
	return strconv.FormatInt(int64(v), 10)
}

type AgeGatedFeatures uint64

const (
	AgeGatedFeaturesAgeGatedSpaces                   AgeGatedFeatures = 1 << 0
	AgeGatedFeaturesSensitiveContentShowSetting      AgeGatedFeatures = 1 << 1
	AgeGatedFeaturesDMPrivacySettings                AgeGatedFeatures = 1 << 2
	AgeGatedFeaturesMessageRequestRestrictionsToggle AgeGatedFeatures = 1 << 3
	AgeGatedFeaturesLargeServerAccess                AgeGatedFeatures = 1 << 4
	AgeGatedFeaturesCommandsToggle                   AgeGatedFeatures = 1 << 5
	AgeGatedFeaturesReactiveCheck                    AgeGatedFeatures = 1 << 6
	AgeGatedFeaturesStageSpeaking                    AgeGatedFeatures = 1 << 7
)

func (v AgeGatedFeatures) String() string {
	switch v {
	case AgeGatedFeaturesAgeGatedSpaces:
		return "AGE_GATED_SPACES"
	case AgeGatedFeaturesSensitiveContentShowSetting:
		return "SENSITIVE_CONTENT_SHOW_SETTING"
	case AgeGatedFeaturesDMPrivacySettings:
		return "DM_PRIVACY_SETTINGS"
	case AgeGatedFeaturesMessageRequestRestrictionsToggle:
		return "MESSAGE_REQUEST_RESTRICTIONS_TOGGLE"
	case AgeGatedFeaturesLargeServerAccess:
		return "LARGE_SERVER_ACCESS"
	case AgeGatedFeaturesCommandsToggle:
		return "COMMANDS_TOGGLE"
	case AgeGatedFeaturesReactiveCheck:
		return "REACTIVE_CHECK"
	case AgeGatedFeaturesStageSpeaking:
		return "STAGE_SPEAKING"
	}
	return strconv.FormatUint(uint64(v), 10)
}

type AgeVerificationStatus int64

const (
	AgeVerificationStatusUnverified    AgeVerificationStatus = 1
	AgeVerificationStatusVerifiedTeen  AgeVerificationStatus = 2
	AgeVerificationStatusVerifiedAdult AgeVerificationStatus = 3
	AgeVerificationStatusInferredAdult AgeVerificationStatus = 4
)

func (v AgeVerificationStatus) String() string {
	switch v {
	case AgeVerificationStatusUnverified:
		return "UNVERIFIED"
	case AgeVerificationStatusVerifiedTeen:
		return "VERIFIED_TEEN"
	case AgeVerificationStatusVerifiedAdult:
		return "VERIFIED_ADULT"
	case AgeVerificationStatusInferredAdult:
		return "INFERRED_ADULT"
	}
	return strconv.FormatInt(int64(v), 10)
}

type AggregationInterval int64

const (
	AggregationIntervalHourly  AggregationInterval = 0
	AggregationIntervalDaily   AggregationInterval = 1
	AggregationIntervalWeekly  AggregationInterval = 2
	AggregationIntervalMonthly AggregationInterval = 3
)

func (v AggregationInterval) String() string {
	switch v {
	case AggregationIntervalHourly:
		return "HOURLY"
	case AggregationIntervalDaily:
		return "DAILY"
	case AggregationIntervalWeekly:
		return "WEEKLY"
	case AggregationIntervalMonthly:
		return "MONTHLY"
	}
	return strconv.FormatInt(int64(v), 10)
}

type AllowedMentionTypes string

const (
	AllowedMentionTypesRoles    AllowedMentionTypes = "roles"
	AllowedMentionTypesUsers    AllowedMentionTypes = "users"
	AllowedMentionTypesEveryone AllowedMentionTypes = "everyone"
)

type ApexExperimentFlags uint64

const (
	ApexExperimentFlagsIsOverride              ApexExperimentFlags = 1 << 0
	ApexExperimentFlagsExposureTrackingEnabled ApexExperimentFlags = 1 << 1
	ApexExperimentFlagsDependentExperiment     ApexExperimentFlags = 1 << 2
	ApexExperimentFlagsUseAsEligibility        ApexExperimentFlags = 1 << 3
)

func (v ApexExperimentFlags) String() string {
	switch v {
	case ApexExperimentFlagsIsOverride:
		return "IS_OVERRIDE"
	case ApexExperimentFlagsExposureTrackingEnabled:
		return "EXPOSURE_TRACKING_ENABLED"
	case ApexExperimentFlagsDependentExperiment:
		return "DEPENDENT_EXPERIMENT"
	case ApexExperimentFlagsUseAsEligibility:
		return "USE_AS_ELIGIBILITY"
	}
	return strconv.FormatUint(uint64(v), 10)
}

type ApexExperimentSurface int64

const (
	ApexExperimentSurfaceAPI             ApexExperimentSurface = 1
	ApexExperimentSurfaceApp             ApexExperimentSurface = 2
	ApexExperimentSurfaceDeveloperPortal ApexExperimentSurface = 3
	ApexExperimentSurfaceAdminPanel      ApexExperimentSurface = 4
	ApexExperimentSurfaceAdsBudgetAb     ApexExperimentSurface = 5
)

func (v ApexExperimentSurface) String() string {
	switch v {
	case ApexExperimentSurfaceAPI:
		return "API"
	case ApexExperimentSurfaceApp:
		return "APP"
	case ApexExperimentSurfaceDeveloperPortal:
		return "DEVELOPER_PORTAL"
	case ApexExperimentSurfaceAdminPanel:
		return "ADMIN_PANEL"
	case ApexExperimentSurfaceAdsBudgetAb:
		return "ADS_BUDGET_AB"
	}
	return strconv.FormatInt(int64(v), 10)
}

type ApexExperimentUnitType int64

const (
	ApexExperimentUnitTypeUser         ApexExperimentUnitType = 1
	ApexExperimentUnitTypeInstallation ApexExperimentUnitType = 2
	ApexExperimentUnitTypeGuild        ApexExperimentUnitType = 3
	ApexExperimentUnitTypeCustom       ApexExperimentUnitType = 4
)

func (v ApexExperimentUnitType) String() string {
	switch v {
	case ApexExperimentUnitTypeUser:
		return "USER"
	case ApexExperimentUnitTypeInstallation:
		return "INSTALLATION"
	case ApexExperimentUnitTypeGuild:
		return "GUILD"
	case ApexExperimentUnitTypeCustom:
		return "CUSTOM"
	}
	return strconv.FormatInt(int64(v), 10)
}

type ApexExperimentVariantType int64

const (
	ApexExperimentVariantTypeActive    ApexExperimentVariantType = 1
	ApexExperimentVariantTypeUnused    ApexExperimentVariantType = 2
	ApexExperimentVariantTypeBurned    ApexExperimentVariantType = 3
	ApexExperimentVariantTypePreserved ApexExperimentVariantType = 4
)

func (v ApexExperimentVariantType) String() string {
	switch v {
	case ApexExperimentVariantTypeActive:
		return "ACTIVE"
	case ApexExperimentVariantTypeUnused:
		return "UNUSED"
	case ApexExperimentVariantTypeBurned:
		return "BURNED"
	case ApexExperimentVariantTypePreserved:
		return "PRESERVED"
	}
	return strconv.FormatInt(int64(v), 10)
}

type AppealEligibility int64

const (
	AppealEligibilityDsaEligible       AppealEligibility = 1
	AppealEligibilityInAppEligible     AppealEligibility = 2
	AppealEligibilityAgeVerifyEligible AppealEligibility = 3
)

func (v AppealEligibility) String() string {
	switch v {
	case AppealEligibilityDsaEligible:
		return "DSA_ELIGIBLE"
	case AppealEligibilityInAppEligible:
		return "IN_APP_ELIGIBLE"
	case AppealEligibilityAgeVerifyEligible:
		return "AGE_VERIFY_ELIGIBLE"
	}
	return strconv.FormatInt(int64(v), 10)
}

type AppealIngestionSignal int64

const (
	AppealIngestionSignalDidntViolatePolicy AppealIngestionSignal = 0
	AppealIngestionSignalTooStrictUnfair    AppealIngestionSignal = 1
	AppealIngestionSignalDontAgreePenalty   AppealIngestionSignal = 2
	AppealIngestionSignalSomethingElse      AppealIngestionSignal = 3
)

func (v AppealIngestionSignal) String() string {
	switch v {
	case AppealIngestionSignalDidntViolatePolicy:
		return "DIDNT_VIOLATE_POLICY"
	case AppealIngestionSignalTooStrictUnfair:
		return "TOO_STRICT_UNFAIR"
	case AppealIngestionSignalDontAgreePenalty:
		return "DONT_AGREE_PENALTY"
	case AppealIngestionSignalSomethingElse:
		return "SOMETHING_ELSE"
	}
	return strconv.FormatInt(int64(v), 10)
}

type AppealIngestionType int64

const (
	AppealIngestionTypeWebform   AppealIngestionType = 0
	AppealIngestionTypeAgeVerify AppealIngestionType = 1
	AppealIngestionTypeInApp     AppealIngestionType = 2
)

func (v AppealIngestionType) String() string {
	switch v {
	case AppealIngestionTypeWebform:
		return "WEBFORM"
	case AppealIngestionTypeAgeVerify:
		return "AGE_VERIFY"
	case AppealIngestionTypeInApp:
		return "IN_APP"
	}
	return strconv.FormatInt(int64(v), 10)
}

type ApplicationAssetType string

const (
	ApplicationAssetTypeImage ApplicationAssetType = "image"
)

type ApplicationAssetVisibility string

const (
	ApplicationAssetVisibilityPrivate ApplicationAssetVisibility = "private"
	ApplicationAssetVisibilityPublic  ApplicationAssetVisibility = "public"
)

type ApplicationCommandHandlerType int64

const (
	ApplicationCommandHandlerTypeAppHandler               ApplicationCommandHandlerType = 1
	ApplicationCommandHandlerTypeDiscordLaunchActivity    ApplicationCommandHandlerType = 2
	ApplicationCommandHandlerTypeAppHandlerLaunchActivity ApplicationCommandHandlerType = 3
)

func (v ApplicationCommandHandlerType) String() string {
	switch v {
	case ApplicationCommandHandlerTypeAppHandler:
		return "APP_HANDLER"
	case ApplicationCommandHandlerTypeDiscordLaunchActivity:
		return "DISCORD_LAUNCH_ACTIVITY"
	case ApplicationCommandHandlerTypeAppHandlerLaunchActivity:
		return "APP_HANDLER_LAUNCH_ACTIVITY"
	}
	return strconv.FormatInt(int64(v), 10)
}

type ApplicationCommandOptionType int64

const (
	ApplicationCommandOptionTypeSubCommand      ApplicationCommandOptionType = 1
	ApplicationCommandOptionTypeSubCommandGroup ApplicationCommandOptionType = 2
	ApplicationCommandOptionTypeString          ApplicationCommandOptionType = 3
	ApplicationCommandOptionTypeInteger         ApplicationCommandOptionType = 4
	ApplicationCommandOptionTypeBoolean         ApplicationCommandOptionType = 5
	ApplicationCommandOptionTypeUser            ApplicationCommandOptionType = 6
	ApplicationCommandOptionTypeChannel         ApplicationCommandOptionType = 7
	ApplicationCommandOptionTypeRole            ApplicationCommandOptionType = 8
	ApplicationCommandOptionTypeMentionable     ApplicationCommandOptionType = 9
	ApplicationCommandOptionTypeNumber          ApplicationCommandOptionType = 10
	ApplicationCommandOptionTypeAttachment      ApplicationCommandOptionType = 11
)

func (v ApplicationCommandOptionType) String() string {
	switch v {
	case ApplicationCommandOptionTypeSubCommand:
		return "SUB_COMMAND"
	case ApplicationCommandOptionTypeSubCommandGroup:
		return "SUB_COMMAND_GROUP"
	case ApplicationCommandOptionTypeString:
		return "STRING"
	case ApplicationCommandOptionTypeInteger:
		return "INTEGER"
	case ApplicationCommandOptionTypeBoolean:
		return "BOOLEAN"
	case ApplicationCommandOptionTypeUser:
		return "USER"
	case ApplicationCommandOptionTypeChannel:
		return "CHANNEL"
	case ApplicationCommandOptionTypeRole:
		return "ROLE"
	case ApplicationCommandOptionTypeMentionable:
		return "MENTIONABLE"
	case ApplicationCommandOptionTypeNumber:
		return "NUMBER"
	case ApplicationCommandOptionTypeAttachment:
		return "ATTACHMENT"
	}
	return strconv.FormatInt(int64(v), 10)
}

type ApplicationCommandPermissionType int64

const (
	ApplicationCommandPermissionTypeRole    ApplicationCommandPermissionType = 1
	ApplicationCommandPermissionTypeUser    ApplicationCommandPermissionType = 2
	ApplicationCommandPermissionTypeChannel ApplicationCommandPermissionType = 3
)

func (v ApplicationCommandPermissionType) String() string {
	switch v {
	case ApplicationCommandPermissionTypeRole:
		return "ROLE"
	case ApplicationCommandPermissionTypeUser:
		return "USER"
	case ApplicationCommandPermissionTypeChannel:
		return "CHANNEL"
	}
	return strconv.FormatInt(int64(v), 10)
}

type ApplicationCommandType int64

const (
	ApplicationCommandTypeChatInput         ApplicationCommandType = 1
	ApplicationCommandTypeUser              ApplicationCommandType = 2
	ApplicationCommandTypeMessage           ApplicationCommandType = 3
	ApplicationCommandTypePrimaryEntryPoint ApplicationCommandType = 4
)

func (v ApplicationCommandType) String() string {
	switch v {
	case ApplicationCommandTypeChatInput:
		return "CHAT_INPUT"
	case ApplicationCommandTypeUser:
		return "USER"
	case ApplicationCommandTypeMessage:
		return "MESSAGE"
	case ApplicationCommandTypePrimaryEntryPoint:
		return "PRIMARY_ENTRY_POINT"
	}
	return strconv.FormatInt(int64(v), 10)
}

type ApplicationDirectoryActiveStateType int64

const (
	ApplicationDirectoryActiveStateTypePreview ApplicationDirectoryActiveStateType = 0
	ApplicationDirectoryActiveStateTypeActive  ApplicationDirectoryActiveStateType = 1
)

func (v ApplicationDirectoryActiveStateType) String() string {
	switch v {
	case ApplicationDirectoryActiveStateTypePreview:
		return "PREVIEW"
	case ApplicationDirectoryActiveStateTypeActive:
		return "ACTIVE"
	}
	return strconv.FormatInt(int64(v), 10)
}

type ApplicationDirectoryCarouselItemType int64

const (
	ApplicationDirectoryCarouselItemTypeImage ApplicationDirectoryCarouselItemType = 1
	ApplicationDirectoryCarouselItemTypeVideo ApplicationDirectoryCarouselItemType = 2
)

func (v ApplicationDirectoryCarouselItemType) String() string {
	switch v {
	case ApplicationDirectoryCarouselItemTypeImage:
		return "IMAGE"
	case ApplicationDirectoryCarouselItemTypeVideo:
		return "VIDEO"
	}
	return strconv.FormatInt(int64(v), 10)
}

type ApplicationDirectoryItemType int64

const (
	ApplicationDirectoryItemTypeApplication       ApplicationDirectoryItemType = 1
	ApplicationDirectoryItemTypeLink              ApplicationDirectoryItemType = 2
	ApplicationDirectoryItemTypeApplicationBanner ApplicationDirectoryItemType = 3
)

func (v ApplicationDirectoryItemType) String() string {
	switch v {
	case ApplicationDirectoryItemTypeApplication:
		return "APPLICATION"
	case ApplicationDirectoryItemTypeLink:
		return "LINK"
	case ApplicationDirectoryItemTypeApplicationBanner:
		return "APPLICATION_BANNER"
	}
	return strconv.FormatInt(int64(v), 10)
}

type ApplicationDirectoryPlatformFlags uint64

const (
	ApplicationDirectoryPlatformFlagsIos     ApplicationDirectoryPlatformFlags = 1 << 0
	ApplicationDirectoryPlatformFlagsAndroid ApplicationDirectoryPlatformFlags = 1 << 1
	ApplicationDirectoryPlatformFlagsWeb     ApplicationDirectoryPlatformFlags = 1 << 2
)

func (v ApplicationDirectoryPlatformFlags) String() string {
	switch v {
	case ApplicationDirectoryPlatformFlagsIos:
		return "IOS"
	case ApplicationDirectoryPlatformFlagsAndroid:
		return "ANDROID"
	case ApplicationDirectoryPlatformFlagsWeb:
		return "WEB"
	}
	return strconv.FormatUint(uint64(v), 10)
}

type ApplicationDirectoryRequestSurface int64

const (
	ApplicationDirectoryRequestSurfaceApplicationDirectory     ApplicationDirectoryRequestSurface = 1
	ApplicationDirectoryRequestSurfaceAppLauncherInText        ApplicationDirectoryRequestSurface = 2
	ApplicationDirectoryRequestSurfaceAppLauncherInVoiceBanner ApplicationDirectoryRequestSurface = 3
)

func (v ApplicationDirectoryRequestSurface) String() string {
	switch v {
	case ApplicationDirectoryRequestSurfaceApplicationDirectory:
		return "APPLICATION_DIRECTORY"
	case ApplicationDirectoryRequestSurfaceAppLauncherInText:
		return "APP_LAUNCHER_IN_TEXT"
	case ApplicationDirectoryRequestSurfaceAppLauncherInVoiceBanner:
		return "APP_LAUNCHER_IN_VOICE_BANNER"
	}
	return strconv.FormatInt(int64(v), 10)
}

type ApplicationDisclosureType int64

const (
	ApplicationDisclosureTypeUnspecifiedDisclosure        ApplicationDisclosureType = 0
	ApplicationDisclosureTypeIPLocation                   ApplicationDisclosureType = 1
	ApplicationDisclosureTypeDisplaysAdvertisements       ApplicationDisclosureType = 2
	ApplicationDisclosureTypePartnerSdkDataSharingMessage ApplicationDisclosureType = 3
)

func (v ApplicationDisclosureType) String() string {
	switch v {
	case ApplicationDisclosureTypeUnspecifiedDisclosure:
		return "UNSPECIFIED_DISCLOSURE"
	case ApplicationDisclosureTypeIPLocation:
		return "IP_LOCATION"
	case ApplicationDisclosureTypeDisplaysAdvertisements:
		return "DISPLAYS_ADVERTISEMENTS"
	case ApplicationDisclosureTypePartnerSdkDataSharingMessage:
		return "PARTNER_SDK_DATA_SHARING_MESSAGE"
	}
	return strconv.FormatInt(int64(v), 10)
}

type ApplicationDiscoverabilityState int64

const (
	ApplicationDiscoverabilityStateIneligible      ApplicationDiscoverabilityState = 1
	ApplicationDiscoverabilityStateNotDiscoverable ApplicationDiscoverabilityState = 2
	ApplicationDiscoverabilityStateDiscoverable    ApplicationDiscoverabilityState = 3
	ApplicationDiscoverabilityStateFeatureable     ApplicationDiscoverabilityState = 4
	ApplicationDiscoverabilityStateBlocked         ApplicationDiscoverabilityState = 5
)

func (v ApplicationDiscoverabilityState) String() string {
	switch v {
	case ApplicationDiscoverabilityStateIneligible:
		return "INELIGIBLE"
	case ApplicationDiscoverabilityStateNotDiscoverable:
		return "NOT_DISCOVERABLE"
	case ApplicationDiscoverabilityStateDiscoverable:
		return "DISCOVERABLE"
	case ApplicationDiscoverabilityStateFeatureable:
		return "FEATUREABLE"
	case ApplicationDiscoverabilityStateBlocked:
		return "BLOCKED"
	}
	return strconv.FormatInt(int64(v), 10)
}

type ApplicationDiscoveryEligibilityFlags uint64

const (
	ApplicationDiscoveryEligibilityFlagsVerified              ApplicationDiscoveryEligibilityFlags = 1 << 0
	ApplicationDiscoveryEligibilityFlagsTag                   ApplicationDiscoveryEligibilityFlags = 1 << 1
	ApplicationDiscoveryEligibilityFlagsDescription           ApplicationDiscoveryEligibilityFlags = 1 << 2
	ApplicationDiscoveryEligibilityFlagsTermsOfService        ApplicationDiscoveryEligibilityFlags = 1 << 3
	ApplicationDiscoveryEligibilityFlagsPrivacyPolicy         ApplicationDiscoveryEligibilityFlags = 1 << 4
	ApplicationDiscoveryEligibilityFlagsInstallParams         ApplicationDiscoveryEligibilityFlags = 1 << 5
	ApplicationDiscoveryEligibilityFlagsSafeName              ApplicationDiscoveryEligibilityFlags = 1 << 6
	ApplicationDiscoveryEligibilityFlagsSafeDescription       ApplicationDiscoveryEligibilityFlags = 1 << 7
	ApplicationDiscoveryEligibilityFlagsApprovedCommands      ApplicationDiscoveryEligibilityFlags = 1 << 8
	ApplicationDiscoveryEligibilityFlagsSupportGuild          ApplicationDiscoveryEligibilityFlags = 1 << 9
	ApplicationDiscoveryEligibilityFlagsSafeCommands          ApplicationDiscoveryEligibilityFlags = 1 << 10
	ApplicationDiscoveryEligibilityFlagsMFA                   ApplicationDiscoveryEligibilityFlags = 1 << 11
	ApplicationDiscoveryEligibilityFlagsSafeDirectoryOverview ApplicationDiscoveryEligibilityFlags = 1 << 12
	ApplicationDiscoveryEligibilityFlagsSupportedLocales      ApplicationDiscoveryEligibilityFlags = 1 << 13
	ApplicationDiscoveryEligibilityFlagsSafeShortDescription  ApplicationDiscoveryEligibilityFlags = 1 << 14
	ApplicationDiscoveryEligibilityFlagsSafeRoleConnections   ApplicationDiscoveryEligibilityFlags = 1 << 15
)

func (v ApplicationDiscoveryEligibilityFlags) String() string {
	switch v {
	case ApplicationDiscoveryEligibilityFlagsVerified:
		return "VERIFIED"
	case ApplicationDiscoveryEligibilityFlagsTag:
		return "TAG"
	case ApplicationDiscoveryEligibilityFlagsDescription:
		return "DESCRIPTION"
	case ApplicationDiscoveryEligibilityFlagsTermsOfService:
		return "TERMS_OF_SERVICE"
	case ApplicationDiscoveryEligibilityFlagsPrivacyPolicy:
		return "PRIVACY_POLICY"
	case ApplicationDiscoveryEligibilityFlagsInstallParams:
		return "INSTALL_PARAMS"
	case ApplicationDiscoveryEligibilityFlagsSafeName:
		return "SAFE_NAME"
	case ApplicationDiscoveryEligibilityFlagsSafeDescription:
		return "SAFE_DESCRIPTION"
	case ApplicationDiscoveryEligibilityFlagsApprovedCommands:
		return "APPROVED_COMMANDS"
	case ApplicationDiscoveryEligibilityFlagsSupportGuild:
		return "SUPPORT_GUILD"
	case ApplicationDiscoveryEligibilityFlagsSafeCommands:
		return "SAFE_COMMANDS"
	case ApplicationDiscoveryEligibilityFlagsMFA:
		return "MFA"
	case ApplicationDiscoveryEligibilityFlagsSafeDirectoryOverview:
		return "SAFE_DIRECTORY_OVERVIEW"
	case ApplicationDiscoveryEligibilityFlagsSupportedLocales:
		return "SUPPORTED_LOCALES"
	case ApplicationDiscoveryEligibilityFlagsSafeShortDescription:
		return "SAFE_SHORT_DESCRIPTION"
	case ApplicationDiscoveryEligibilityFlagsSafeRoleConnections:
		return "SAFE_ROLE_CONNECTIONS"
	}
	return strconv.FormatUint(uint64(v), 10)
}

type ApplicationFlags uint64

const (
	ApplicationFlagsEmbeddedReleased               ApplicationFlags = 1 << 1
	ApplicationFlagsManagedEmoji                   ApplicationFlags = 1 << 2
	ApplicationFlagsEmbeddedIap                    ApplicationFlags = 1 << 3
	ApplicationFlagsGroupDMCreate                  ApplicationFlags = 1 << 4
	ApplicationFlagsRPCPrivateBeta                 ApplicationFlags = 1 << 5
	ApplicationFlagsAutoModerationRuleCreateBadge  ApplicationFlags = 1 << 6
	ApplicationFlagsGameProfileDisabled            ApplicationFlags = 1 << 7
	ApplicationFlagsPublicOAuth2Client             ApplicationFlags = 1 << 8
	ApplicationFlagsContextlessActivity            ApplicationFlags = 1 << 9
	ApplicationFlagsSocialLayerIntegrationLimited  ApplicationFlags = 1 << 10
	ApplicationFlagsCloudGamingDemo                ApplicationFlags = 1 << 11
	ApplicationFlagsGatewayPresence                ApplicationFlags = 1 << 12
	ApplicationFlagsGatewayPresenceLimited         ApplicationFlags = 1 << 13
	ApplicationFlagsGatewayGuildMembers            ApplicationFlags = 1 << 14
	ApplicationFlagsGatewayGuildMembersLimited     ApplicationFlags = 1 << 15
	ApplicationFlagsVerificationPendingGuildLimit  ApplicationFlags = 1 << 16
	ApplicationFlagsEmbedded                       ApplicationFlags = 1 << 17
	ApplicationFlagsGatewayMessageContent          ApplicationFlags = 1 << 18
	ApplicationFlagsGatewayMessageContentLimited   ApplicationFlags = 1 << 19
	ApplicationFlagsEmbeddedFirstParty             ApplicationFlags = 1 << 20
	ApplicationFlagsApplicationCommandMigrated     ApplicationFlags = 1 << 21
	ApplicationFlagsApplicationCommandBadge        ApplicationFlags = 1 << 23
	ApplicationFlagsActive                         ApplicationFlags = 1 << 24
	ApplicationFlagsActiveGracePeriod              ApplicationFlags = 1 << 25
	ApplicationFlagsIframeModal                    ApplicationFlags = 1 << 26
	ApplicationFlagsSocialLayerIntegration         ApplicationFlags = 1 << 27
	ApplicationFlagsPromoted                       ApplicationFlags = 1 << 29
	ApplicationFlagsPartner                        ApplicationFlags = 1 << 30
	ApplicationFlagsParent                         ApplicationFlags = 1 << 33
	ApplicationFlagsDisableRelationshipAccess      ApplicationFlags = 1 << 34
	ApplicationFlagsStorefrontEligible             ApplicationFlags = 1 << 35
	ApplicationFlagsAllowAssets                    ApplicationFlags = 1 << 8
	ApplicationFlagsAllowActivityActionSpectate    ApplicationFlags = 1 << 9
	ApplicationFlagsAllowActivityActionJoinRequest ApplicationFlags = 1 << 10
	ApplicationFlagsRPCHasConnected                ApplicationFlags = 1 << 11
)

func (v ApplicationFlags) String() string {
	switch v {
	case ApplicationFlagsEmbeddedReleased:
		return "EMBEDDED_RELEASED"
	case ApplicationFlagsManagedEmoji:
		return "MANAGED_EMOJI"
	case ApplicationFlagsEmbeddedIap:
		return "EMBEDDED_IAP"
	case ApplicationFlagsGroupDMCreate:
		return "GROUP_DM_CREATE"
	case ApplicationFlagsRPCPrivateBeta:
		return "RPC_PRIVATE_BETA"
	case ApplicationFlagsAutoModerationRuleCreateBadge:
		return "AUTO_MODERATION_RULE_CREATE_BADGE"
	case ApplicationFlagsGameProfileDisabled:
		return "GAME_PROFILE_DISABLED"
	case ApplicationFlagsPublicOAuth2Client:
		return "PUBLIC_OAUTH2_CLIENT"
	case ApplicationFlagsContextlessActivity:
		return "CONTEXTLESS_ACTIVITY"
	case ApplicationFlagsSocialLayerIntegrationLimited:
		return "SOCIAL_LAYER_INTEGRATION_LIMITED"
	case ApplicationFlagsCloudGamingDemo:
		return "CLOUD_GAMING_DEMO"
	case ApplicationFlagsGatewayPresence:
		return "GATEWAY_PRESENCE"
	case ApplicationFlagsGatewayPresenceLimited:
		return "GATEWAY_PRESENCE_LIMITED"
	case ApplicationFlagsGatewayGuildMembers:
		return "GATEWAY_GUILD_MEMBERS"
	case ApplicationFlagsGatewayGuildMembersLimited:
		return "GATEWAY_GUILD_MEMBERS_LIMITED"
	case ApplicationFlagsVerificationPendingGuildLimit:
		return "VERIFICATION_PENDING_GUILD_LIMIT"
	case ApplicationFlagsEmbedded:
		return "EMBEDDED"
	case ApplicationFlagsGatewayMessageContent:
		return "GATEWAY_MESSAGE_CONTENT"
	case ApplicationFlagsGatewayMessageContentLimited:
		return "GATEWAY_MESSAGE_CONTENT_LIMITED"
	case ApplicationFlagsEmbeddedFirstParty:
		return "EMBEDDED_FIRST_PARTY"
	case ApplicationFlagsApplicationCommandMigrated:
		return "APPLICATION_COMMAND_MIGRATED"
	case ApplicationFlagsApplicationCommandBadge:
		return "APPLICATION_COMMAND_BADGE"
	case ApplicationFlagsActive:
		return "ACTIVE"
	case ApplicationFlagsActiveGracePeriod:
		return "ACTIVE_GRACE_PERIOD"
	case ApplicationFlagsIframeModal:
		return "IFRAME_MODAL"
	case ApplicationFlagsSocialLayerIntegration:
		return "SOCIAL_LAYER_INTEGRATION"
	case ApplicationFlagsPromoted:
		return "PROMOTED"
	case ApplicationFlagsPartner:
		return "PARTNER"
	case ApplicationFlagsParent:
		return "PARENT"
	case ApplicationFlagsDisableRelationshipAccess:
		return "DISABLE_RELATIONSHIP_ACCESS"
	case ApplicationFlagsStorefrontEligible:
		return "STOREFRONT_ELIGIBLE"
	}
	return strconv.FormatUint(uint64(v), 10)
}

type ApplicationIdentityProviderType string

const (
	ApplicationIdentityProviderTypeOidc               ApplicationIdentityProviderType = "OIDC"
	ApplicationIdentityProviderTypeEpicOnlineServices ApplicationIdentityProviderType = "EPIC_ONLINE_SERVICES"
	ApplicationIdentityProviderTypeSteam              ApplicationIdentityProviderType = "STEAM"
	ApplicationIdentityProviderTypeUnity              ApplicationIdentityProviderType = "UNITY"
	ApplicationIdentityProviderTypeDiscordBot         ApplicationIdentityProviderType = "DISCORD_BOT"
	ApplicationIdentityProviderTypeApple              ApplicationIdentityProviderType = "APPLE"
	ApplicationIdentityProviderTypePlaystationNetwork ApplicationIdentityProviderType = "PLAYSTATION_NETWORK"
)

type ApplicationIntegrationType int64

const (
	ApplicationIntegrationTypeGuildInstall ApplicationIntegrationType = 0
	ApplicationIntegrationTypeUserInstall  ApplicationIntegrationType = 1
)

func (v ApplicationIntegrationType) String() string {
	switch v {
	case ApplicationIntegrationTypeGuildInstall:
		return "GUILD_INSTALL"
	case ApplicationIntegrationTypeUserInstall:
		return "USER_INSTALL"
	}
	return strconv.FormatInt(int64(v), 10)
}

type ApplicationInteractionsVersion int64

const (
	ApplicationInteractionsVersionVersion1 ApplicationInteractionsVersion = 1
	ApplicationInteractionsVersionVersion2 ApplicationInteractionsVersion = 2
)

func (v ApplicationInteractionsVersion) String() string {
	switch v {
	case ApplicationInteractionsVersionVersion1:
		return "VERSION_1"
	case ApplicationInteractionsVersionVersion2:
		return "VERSION_2"
	}
	return strconv.FormatInt(int64(v), 10)
}

type ApplicationMembershipState int64

const (
	ApplicationMembershipStateInvited  ApplicationMembershipState = 1
	ApplicationMembershipStateAccepted ApplicationMembershipState = 2
)

func (v ApplicationMembershipState) String() string {
	switch v {
	case ApplicationMembershipStateInvited:
		return "INVITED"
	case ApplicationMembershipStateAccepted:
		return "ACCEPTED"
	}
	return strconv.FormatInt(int64(v), 10)
}

type ApplicationMissingDataType string

const (
	ApplicationMissingDataTypeIcon ApplicationMissingDataType = "icon"
)

type ApplicationMonetizationEligibilityFlags uint64

const (
	ApplicationMonetizationEligibilityFlagsVerified                 ApplicationMonetizationEligibilityFlags = 1 << 0
	ApplicationMonetizationEligibilityFlagsHasTeam                  ApplicationMonetizationEligibilityFlags = 1 << 1
	ApplicationMonetizationEligibilityFlagsApprovedCommands         ApplicationMonetizationEligibilityFlags = 1 << 2
	ApplicationMonetizationEligibilityFlagsTermsOfService           ApplicationMonetizationEligibilityFlags = 1 << 3
	ApplicationMonetizationEligibilityFlagsPrivacyPolicy            ApplicationMonetizationEligibilityFlags = 1 << 4
	ApplicationMonetizationEligibilityFlagsSafeName                 ApplicationMonetizationEligibilityFlags = 1 << 5
	ApplicationMonetizationEligibilityFlagsSafeDescription          ApplicationMonetizationEligibilityFlags = 1 << 6
	ApplicationMonetizationEligibilityFlagsSafeRoleConnections      ApplicationMonetizationEligibilityFlags = 1 << 7
	ApplicationMonetizationEligibilityFlagsUserIsTeamOwner          ApplicationMonetizationEligibilityFlags = 1 << 8
	ApplicationMonetizationEligibilityFlagsNotQuarantined           ApplicationMonetizationEligibilityFlags = 1 << 9
	ApplicationMonetizationEligibilityFlagsUserLocaleSupported      ApplicationMonetizationEligibilityFlags = 1 << 10
	ApplicationMonetizationEligibilityFlagsUserAgeSupported         ApplicationMonetizationEligibilityFlags = 1 << 11
	ApplicationMonetizationEligibilityFlagsUserDateOfBirthDefined   ApplicationMonetizationEligibilityFlags = 1 << 12
	ApplicationMonetizationEligibilityFlagsUserMFAEnabled           ApplicationMonetizationEligibilityFlags = 1 << 13
	ApplicationMonetizationEligibilityFlagsUserEmailVerified        ApplicationMonetizationEligibilityFlags = 1 << 14
	ApplicationMonetizationEligibilityFlagsTeamMembersEmailVerified ApplicationMonetizationEligibilityFlags = 1 << 15
	ApplicationMonetizationEligibilityFlagsTeamMembersMFAEnabled    ApplicationMonetizationEligibilityFlags = 1 << 16
	ApplicationMonetizationEligibilityFlagsNoBlockingIssues         ApplicationMonetizationEligibilityFlags = 1 << 17
	ApplicationMonetizationEligibilityFlagsValidPayoutStatus        ApplicationMonetizationEligibilityFlags = 1 << 18
)

func (v ApplicationMonetizationEligibilityFlags) String() string {
	switch v {
	case ApplicationMonetizationEligibilityFlagsVerified:
		return "VERIFIED"
	case ApplicationMonetizationEligibilityFlagsHasTeam:
		return "HAS_TEAM"
	case ApplicationMonetizationEligibilityFlagsApprovedCommands:
		return "APPROVED_COMMANDS"
	case ApplicationMonetizationEligibilityFlagsTermsOfService:
		return "TERMS_OF_SERVICE"
	case ApplicationMonetizationEligibilityFlagsPrivacyPolicy:
		return "PRIVACY_POLICY"
	case ApplicationMonetizationEligibilityFlagsSafeName:
		return "SAFE_NAME"
	case ApplicationMonetizationEligibilityFlagsSafeDescription:
		return "SAFE_DESCRIPTION"
	case ApplicationMonetizationEligibilityFlagsSafeRoleConnections:
		return "SAFE_ROLE_CONNECTIONS"
	case ApplicationMonetizationEligibilityFlagsUserIsTeamOwner:
		return "USER_IS_TEAM_OWNER"
	case ApplicationMonetizationEligibilityFlagsNotQuarantined:
		return "NOT_QUARANTINED"
	case ApplicationMonetizationEligibilityFlagsUserLocaleSupported:
		return "USER_LOCALE_SUPPORTED"
	case ApplicationMonetizationEligibilityFlagsUserAgeSupported:
		return "USER_AGE_SUPPORTED"
	case ApplicationMonetizationEligibilityFlagsUserDateOfBirthDefined:
		return "USER_DATE_OF_BIRTH_DEFINED"
	case ApplicationMonetizationEligibilityFlagsUserMFAEnabled:
		return "USER_MFA_ENABLED"
	case ApplicationMonetizationEligibilityFlagsUserEmailVerified:
		return "USER_EMAIL_VERIFIED"
	case ApplicationMonetizationEligibilityFlagsTeamMembersEmailVerified:
		return "TEAM_MEMBERS_EMAIL_VERIFIED"
	case ApplicationMonetizationEligibilityFlagsTeamMembersMFAEnabled:
		return "TEAM_MEMBERS_MFA_ENABLED"
	case ApplicationMonetizationEligibilityFlagsNoBlockingIssues:
		return "NO_BLOCKING_ISSUES"
	case ApplicationMonetizationEligibilityFlagsValidPayoutStatus:
		return "VALID_PAYOUT_STATUS"
	}
	return strconv.FormatUint(uint64(v), 10)
}

type ApplicationMonetizationState int64

const (
	ApplicationMonetizationStateNone    ApplicationMonetizationState = 1
	ApplicationMonetizationStateEnabled ApplicationMonetizationState = 2
	ApplicationMonetizationStateBlocked ApplicationMonetizationState = 3
)

func (v ApplicationMonetizationState) String() string {
	switch v {
	case ApplicationMonetizationStateNone:
		return "NONE"
	case ApplicationMonetizationStateEnabled:
		return "ENABLED"
	case ApplicationMonetizationStateBlocked:
		return "BLOCKED"
	}
	return strconv.FormatInt(int64(v), 10)
}

type ApplicationOAuth2AssetType int64

const (
	ApplicationOAuth2AssetTypeSmall ApplicationOAuth2AssetType = 1
	ApplicationOAuth2AssetTypeLarge ApplicationOAuth2AssetType = 2
)

func (v ApplicationOAuth2AssetType) String() string {
	switch v {
	case ApplicationOAuth2AssetTypeSmall:
		return "SMALL"
	case ApplicationOAuth2AssetTypeLarge:
		return "LARGE"
	}
	return strconv.FormatInt(int64(v), 10)
}

type ApplicationType int64

const (
	ApplicationTypeDeprecatedGame      ApplicationType = 1
	ApplicationTypeMusic               ApplicationType = 2
	ApplicationTypeTicketedEvents      ApplicationType = 3
	ApplicationTypeCreatorMonetization ApplicationType = 4
	ApplicationTypeGame                ApplicationType = 5
)

func (v ApplicationType) String() string {
	switch v {
	case ApplicationTypeDeprecatedGame:
		return "DEPRECATED_GAME"
	case ApplicationTypeMusic:
		return "MUSIC"
	case ApplicationTypeTicketedEvents:
		return "TICKETED_EVENTS"
	case ApplicationTypeCreatorMonetization:
		return "CREATOR_MONETIZATION"
	case ApplicationTypeGame:
		return "GAME"
	}
	return strconv.FormatInt(int64(v), 10)
}

type ApplicationUndeletableReason int64

const (
	ApplicationUndeletableReasonUnknown                    ApplicationUndeletableReason = 0
	ApplicationUndeletableReasonUserThresholdExceeded      ApplicationUndeletableReason = 1
	ApplicationUndeletableReasonSocialSdkAppDeletionError  ApplicationUndeletableReason = 2
	ApplicationUndeletableReasonParentHasChildApplications ApplicationUndeletableReason = 3
)

func (v ApplicationUndeletableReason) String() string {
	switch v {
	case ApplicationUndeletableReasonUnknown:
		return "UNKNOWN"
	case ApplicationUndeletableReasonUserThresholdExceeded:
		return "USER_THRESHOLD_EXCEEDED"
	case ApplicationUndeletableReasonSocialSdkAppDeletionError:
		return "SOCIAL_SDK_APP_DELETION_ERROR"
	case ApplicationUndeletableReasonParentHasChildApplications:
		return "PARENT_HAS_CHILD_APPLICATIONS"
	}
	return strconv.FormatInt(int64(v), 10)
}

type ApplicationVerificationState int64

const (
	ApplicationVerificationStateIneligible            ApplicationVerificationState = 1
	ApplicationVerificationStateUnsubmitted           ApplicationVerificationState = 2
	ApplicationVerificationStateSubmitted             ApplicationVerificationState = 3
	ApplicationVerificationStateApprovedManually      ApplicationVerificationState = 4
	ApplicationVerificationStateBlocked               ApplicationVerificationState = 5
	ApplicationVerificationStateApprovedAutomatically ApplicationVerificationState = 6
)

func (v ApplicationVerificationState) String() string {
	switch v {
	case ApplicationVerificationStateIneligible:
		return "INELIGIBLE"
	case ApplicationVerificationStateUnsubmitted:
		return "UNSUBMITTED"
	case ApplicationVerificationStateSubmitted:
		return "SUBMITTED"
	case ApplicationVerificationStateApprovedManually:
		return "APPROVED_MANUALLY"
	case ApplicationVerificationStateBlocked:
		return "BLOCKED"
	case ApplicationVerificationStateApprovedAutomatically:
		return "APPROVED_AUTOMATICALLY"
	}
	return strconv.FormatInt(int64(v), 10)
}

type ApprovableConsoleType int64

const (
	ApprovableConsoleTypeXbox         ApprovableConsoleType = 1
	ApprovableConsoleTypePlaystation5 ApprovableConsoleType = 2
	ApprovableConsoleTypePlaystation4 ApprovableConsoleType = 3
)

func (v ApprovableConsoleType) String() string {
	switch v {
	case ApprovableConsoleTypeXbox:
		return "XBOX"
	case ApprovableConsoleTypePlaystation5:
		return "PLAYSTATION_5"
	case ApprovableConsoleTypePlaystation4:
		return "PLAYSTATION_4"
	}
	return strconv.FormatInt(int64(v), 10)
}

type AttachmentFlags uint64

const (
	AttachmentFlagsIsClip                  AttachmentFlags = 1 << 0
	AttachmentFlagsIsThumbnail             AttachmentFlags = 1 << 1
	AttachmentFlagsIsRemix                 AttachmentFlags = 1 << 2
	AttachmentFlagsIsSpoiler               AttachmentFlags = 1 << 3
	AttachmentFlagsContainsExplicitMedia   AttachmentFlags = 1 << 4
	AttachmentFlagsIsAnimated              AttachmentFlags = 1 << 5
	AttachmentFlagsContainsGoreContent     AttachmentFlags = 1 << 6
	AttachmentFlagsContainsSelfHarmContent AttachmentFlags = 1 << 7
)

func (v AttachmentFlags) String() string {
	switch v {
	case AttachmentFlagsIsClip:
		return "IS_CLIP"
	case AttachmentFlagsIsThumbnail:
		return "IS_THUMBNAIL"
	case AttachmentFlagsIsRemix:
		return "IS_REMIX"
	case AttachmentFlagsIsSpoiler:
		return "IS_SPOILER"
	case AttachmentFlagsContainsExplicitMedia:
		return "CONTAINS_EXPLICIT_MEDIA"
	case AttachmentFlagsIsAnimated:
		return "IS_ANIMATED"
	case AttachmentFlagsContainsGoreContent:
		return "CONTAINS_GORE_CONTENT"
	case AttachmentFlagsContainsSelfHarmContent:
		return "CONTAINS_SELF_HARM_CONTENT"
	}
	return strconv.FormatUint(uint64(v), 10)
}

type AttachmentTextVariant int64

const (
	AttachmentTextVariantSentAnImage AttachmentTextVariant = 1
	AttachmentTextVariantSentImages  AttachmentTextVariant = 2
	AttachmentTextVariantSentAGIF    AttachmentTextVariant = 3
	AttachmentTextVariantSentAVideo  AttachmentTextVariant = 4
	AttachmentTextVariantSentVideos  AttachmentTextVariant = 5
)

func (v AttachmentTextVariant) String() string {
	switch v {
	case AttachmentTextVariantSentAnImage:
		return "SENT_AN_IMAGE"
	case AttachmentTextVariantSentImages:
		return "SENT_IMAGES"
	case AttachmentTextVariantSentAGIF:
		return "SENT_A_GIF"
	case AttachmentTextVariantSentAVideo:
		return "SENT_A_VIDEO"
	case AttachmentTextVariantSentVideos:
		return "SENT_VIDEOS"
	}
	return strconv.FormatInt(int64(v), 10)
}

type AudioContextType string

const (
	AudioContextTypeUser   AudioContextType = "user"
	AudioContextTypeStream AudioContextType = "stream"
)

type AuditLogActionType int64

const (
	AuditLogActionTypeGuildUpdate                             AuditLogActionType = 1
	AuditLogActionTypeChannelCreate                           AuditLogActionType = 10
	AuditLogActionTypeChannelUpdate                           AuditLogActionType = 11
	AuditLogActionTypeChannelDelete                           AuditLogActionType = 12
	AuditLogActionTypeChannelOverwriteCreate                  AuditLogActionType = 13
	AuditLogActionTypeChannelOverwriteUpdate                  AuditLogActionType = 14
	AuditLogActionTypeChannelOverwriteDelete                  AuditLogActionType = 15
	AuditLogActionTypeMemberKick                              AuditLogActionType = 20
	AuditLogActionTypeMemberPrune                             AuditLogActionType = 21
	AuditLogActionTypeMemberBanAdd                            AuditLogActionType = 22
	AuditLogActionTypeMemberBanRemove                         AuditLogActionType = 23
	AuditLogActionTypeMemberUpdate                            AuditLogActionType = 24
	AuditLogActionTypeMemberRoleUpdate                        AuditLogActionType = 25
	AuditLogActionTypeMemberMove                              AuditLogActionType = 26
	AuditLogActionTypeMemberDisconnect                        AuditLogActionType = 27
	AuditLogActionTypeBotAdd                                  AuditLogActionType = 28
	AuditLogActionTypeRoleCreate                              AuditLogActionType = 30
	AuditLogActionTypeRoleUpdate                              AuditLogActionType = 31
	AuditLogActionTypeRoleDelete                              AuditLogActionType = 32
	AuditLogActionTypeInviteCreate                            AuditLogActionType = 40
	AuditLogActionTypeInviteUpdate                            AuditLogActionType = 41
	AuditLogActionTypeInviteDelete                            AuditLogActionType = 42
	AuditLogActionTypeWebhookCreate                           AuditLogActionType = 50
	AuditLogActionTypeWebhookUpdate                           AuditLogActionType = 51
	AuditLogActionTypeWebhookDelete                           AuditLogActionType = 52
	AuditLogActionTypeEmojiCreate                             AuditLogActionType = 60
	AuditLogActionTypeEmojiUpdate                             AuditLogActionType = 61
	AuditLogActionTypeEmojiDelete                             AuditLogActionType = 62
	AuditLogActionTypeMessageDelete                           AuditLogActionType = 72
	AuditLogActionTypeMessageBulkDelete                       AuditLogActionType = 73
	AuditLogActionTypeMessagePin                              AuditLogActionType = 74
	AuditLogActionTypeMessageUnpin                            AuditLogActionType = 75
	AuditLogActionTypeIntegrationCreate                       AuditLogActionType = 80
	AuditLogActionTypeIntegrationUpdate                       AuditLogActionType = 81
	AuditLogActionTypeIntegrationDelete                       AuditLogActionType = 82
	AuditLogActionTypeStageInstanceCreate                     AuditLogActionType = 83
	AuditLogActionTypeStageInstanceUpdate                     AuditLogActionType = 84
	AuditLogActionTypeStageInstanceDelete                     AuditLogActionType = 85
	AuditLogActionTypeStickerCreate                           AuditLogActionType = 90
	AuditLogActionTypeStickerUpdate                           AuditLogActionType = 91
	AuditLogActionTypeStickerDelete                           AuditLogActionType = 92
	AuditLogActionTypeGuildScheduledEventCreate               AuditLogActionType = 100
	AuditLogActionTypeGuildScheduledEventUpdate               AuditLogActionType = 101
	AuditLogActionTypeGuildScheduledEventDelete               AuditLogActionType = 102
	AuditLogActionTypeThreadCreate                            AuditLogActionType = 110
	AuditLogActionTypeThreadUpdate                            AuditLogActionType = 111
	AuditLogActionTypeThreadDelete                            AuditLogActionType = 112
	AuditLogActionTypeApplicationCommandPermissionUpdate      AuditLogActionType = 121
	AuditLogActionTypeSoundboardSoundCreate                   AuditLogActionType = 130
	AuditLogActionTypeSoundboardSoundUpdate                   AuditLogActionType = 131
	AuditLogActionTypeSoundboardSoundDelete                   AuditLogActionType = 132
	AuditLogActionTypeAutoModerationRuleCreate                AuditLogActionType = 140
	AuditLogActionTypeAutoModerationRuleUpdate                AuditLogActionType = 141
	AuditLogActionTypeAutoModerationRuleDelete                AuditLogActionType = 142
	AuditLogActionTypeAutoModerationBlockMessage              AuditLogActionType = 143
	AuditLogActionTypeAutoModerationFlagToChannel             AuditLogActionType = 144
	AuditLogActionTypeAutoModerationUserCommunicationDisabled AuditLogActionType = 145
	AuditLogActionTypeAutoModerationQuarantineUser            AuditLogActionType = 146
	AuditLogActionTypeCreatorMonetizationRequestCreated       AuditLogActionType = 150
	AuditLogActionTypeCreatorMonetizationTermsAccepted        AuditLogActionType = 151
	AuditLogActionTypeOnboardingPromptCreate                  AuditLogActionType = 163
	AuditLogActionTypeOnboardingPromptUpdate                  AuditLogActionType = 164
	AuditLogActionTypeOnboardingPromptDelete                  AuditLogActionType = 165
	AuditLogActionTypeOnboardingCreate                        AuditLogActionType = 166
	AuditLogActionTypeOnboardingUpdate                        AuditLogActionType = 167
	AuditLogActionTypeGuildHomeFeatureItem                    AuditLogActionType = 171
	AuditLogActionTypeGuildHomeRemoveItem                     AuditLogActionType = 172
	AuditLogActionTypeHarmfulLinksBlockedMessage              AuditLogActionType = 180
	AuditLogActionTypeHomeSettingsCreate                      AuditLogActionType = 190
	AuditLogActionTypeHomeSettingsUpdate                      AuditLogActionType = 191
	AuditLogActionTypeVoiceChannelStatusCreate                AuditLogActionType = 192
	AuditLogActionTypeVoiceChannelStatusDelete                AuditLogActionType = 193
	AuditLogActionTypeClydeAiProfileUpdate                    AuditLogActionType = 194
	AuditLogActionTypeGuildScheduledEventExceptionCreate      AuditLogActionType = 200
	AuditLogActionTypeGuildScheduledEventExceptionUpdate      AuditLogActionType = 201
	AuditLogActionTypeGuildScheduledEventExceptionDelete      AuditLogActionType = 202
	AuditLogActionTypeGuildMemberVerificationUpdate           AuditLogActionType = 210
	AuditLogActionTypeGuildProfileUpdate                      AuditLogActionType = 211
	AuditLogActionTypeGuildMigratePinPermission               AuditLogActionType = 212
	AuditLogActionTypeGuildMigrateBypassSlowmodePermission    AuditLogActionType = 213
)

func (v AuditLogActionType) String() string {
	switch v {
	case AuditLogActionTypeGuildUpdate:
		return "GUILD_UPDATE"
	case AuditLogActionTypeChannelCreate:
		return "CHANNEL_CREATE"
	case AuditLogActionTypeChannelUpdate:
		return "CHANNEL_UPDATE"
	case AuditLogActionTypeChannelDelete:
		return "CHANNEL_DELETE"
	case AuditLogActionTypeChannelOverwriteCreate:
		return "CHANNEL_OVERWRITE_CREATE"
	case AuditLogActionTypeChannelOverwriteUpdate:
		return "CHANNEL_OVERWRITE_UPDATE"
	case AuditLogActionTypeChannelOverwriteDelete:
		return "CHANNEL_OVERWRITE_DELETE"
	case AuditLogActionTypeMemberKick:
		return "MEMBER_KICK"
	case AuditLogActionTypeMemberPrune:
		return "MEMBER_PRUNE"
	case AuditLogActionTypeMemberBanAdd:
		return "MEMBER_BAN_ADD"
	case AuditLogActionTypeMemberBanRemove:
		return "MEMBER_BAN_REMOVE"
	case AuditLogActionTypeMemberUpdate:
		return "MEMBER_UPDATE"
	case AuditLogActionTypeMemberRoleUpdate:
		return "MEMBER_ROLE_UPDATE"
	case AuditLogActionTypeMemberMove:
		return "MEMBER_MOVE"
	case AuditLogActionTypeMemberDisconnect:
		return "MEMBER_DISCONNECT"
	case AuditLogActionTypeBotAdd:
		return "BOT_ADD"
	case AuditLogActionTypeRoleCreate:
		return "ROLE_CREATE"
	case AuditLogActionTypeRoleUpdate:
		return "ROLE_UPDATE"
	case AuditLogActionTypeRoleDelete:
		return "ROLE_DELETE"
	case AuditLogActionTypeInviteCreate:
		return "INVITE_CREATE"
	case AuditLogActionTypeInviteUpdate:
		return "INVITE_UPDATE"
	case AuditLogActionTypeInviteDelete:
		return "INVITE_DELETE"
	case AuditLogActionTypeWebhookCreate:
		return "WEBHOOK_CREATE"
	case AuditLogActionTypeWebhookUpdate:
		return "WEBHOOK_UPDATE"
	case AuditLogActionTypeWebhookDelete:
		return "WEBHOOK_DELETE"
	case AuditLogActionTypeEmojiCreate:
		return "EMOJI_CREATE"
	case AuditLogActionTypeEmojiUpdate:
		return "EMOJI_UPDATE"
	case AuditLogActionTypeEmojiDelete:
		return "EMOJI_DELETE"
	case AuditLogActionTypeMessageDelete:
		return "MESSAGE_DELETE"
	case AuditLogActionTypeMessageBulkDelete:
		return "MESSAGE_BULK_DELETE"
	case AuditLogActionTypeMessagePin:
		return "MESSAGE_PIN"
	case AuditLogActionTypeMessageUnpin:
		return "MESSAGE_UNPIN"
	case AuditLogActionTypeIntegrationCreate:
		return "INTEGRATION_CREATE"
	case AuditLogActionTypeIntegrationUpdate:
		return "INTEGRATION_UPDATE"
	case AuditLogActionTypeIntegrationDelete:
		return "INTEGRATION_DELETE"
	case AuditLogActionTypeStageInstanceCreate:
		return "STAGE_INSTANCE_CREATE"
	case AuditLogActionTypeStageInstanceUpdate:
		return "STAGE_INSTANCE_UPDATE"
	case AuditLogActionTypeStageInstanceDelete:
		return "STAGE_INSTANCE_DELETE"
	case AuditLogActionTypeStickerCreate:
		return "STICKER_CREATE"
	case AuditLogActionTypeStickerUpdate:
		return "STICKER_UPDATE"
	case AuditLogActionTypeStickerDelete:
		return "STICKER_DELETE"
	case AuditLogActionTypeGuildScheduledEventCreate:
		return "GUILD_SCHEDULED_EVENT_CREATE"
	case AuditLogActionTypeGuildScheduledEventUpdate:
		return "GUILD_SCHEDULED_EVENT_UPDATE"
	case AuditLogActionTypeGuildScheduledEventDelete:
		return "GUILD_SCHEDULED_EVENT_DELETE"
	case AuditLogActionTypeThreadCreate:
		return "THREAD_CREATE"
	case AuditLogActionTypeThreadUpdate:
		return "THREAD_UPDATE"
	case AuditLogActionTypeThreadDelete:
		return "THREAD_DELETE"
	case AuditLogActionTypeApplicationCommandPermissionUpdate:
		return "APPLICATION_COMMAND_PERMISSION_UPDATE"
	case AuditLogActionTypeSoundboardSoundCreate:
		return "SOUNDBOARD_SOUND_CREATE"
	case AuditLogActionTypeSoundboardSoundUpdate:
		return "SOUNDBOARD_SOUND_UPDATE"
	case AuditLogActionTypeSoundboardSoundDelete:
		return "SOUNDBOARD_SOUND_DELETE"
	case AuditLogActionTypeAutoModerationRuleCreate:
		return "AUTO_MODERATION_RULE_CREATE"
	case AuditLogActionTypeAutoModerationRuleUpdate:
		return "AUTO_MODERATION_RULE_UPDATE"
	case AuditLogActionTypeAutoModerationRuleDelete:
		return "AUTO_MODERATION_RULE_DELETE"
	case AuditLogActionTypeAutoModerationBlockMessage:
		return "AUTO_MODERATION_BLOCK_MESSAGE"
	case AuditLogActionTypeAutoModerationFlagToChannel:
		return "AUTO_MODERATION_FLAG_TO_CHANNEL"
	case AuditLogActionTypeAutoModerationUserCommunicationDisabled:
		return "AUTO_MODERATION_USER_COMMUNICATION_DISABLED"
	case AuditLogActionTypeAutoModerationQuarantineUser:
		return "AUTO_MODERATION_QUARANTINE_USER"
	case AuditLogActionTypeCreatorMonetizationRequestCreated:
		return "CREATOR_MONETIZATION_REQUEST_CREATED"
	case AuditLogActionTypeCreatorMonetizationTermsAccepted:
		return "CREATOR_MONETIZATION_TERMS_ACCEPTED"
	case AuditLogActionTypeOnboardingPromptCreate:
		return "ONBOARDING_PROMPT_CREATE"
	case AuditLogActionTypeOnboardingPromptUpdate:
		return "ONBOARDING_PROMPT_UPDATE"
	case AuditLogActionTypeOnboardingPromptDelete:
		return "ONBOARDING_PROMPT_DELETE"
	case AuditLogActionTypeOnboardingCreate:
		return "ONBOARDING_CREATE"
	case AuditLogActionTypeOnboardingUpdate:
		return "ONBOARDING_UPDATE"
	case AuditLogActionTypeGuildHomeFeatureItem:
		return "GUILD_HOME_FEATURE_ITEM"
	case AuditLogActionTypeGuildHomeRemoveItem:
		return "GUILD_HOME_REMOVE_ITEM"
	case AuditLogActionTypeHarmfulLinksBlockedMessage:
		return "HARMFUL_LINKS_BLOCKED_MESSAGE"
	case AuditLogActionTypeHomeSettingsCreate:
		return "HOME_SETTINGS_CREATE"
	case AuditLogActionTypeHomeSettingsUpdate:
		return "HOME_SETTINGS_UPDATE"
	case AuditLogActionTypeVoiceChannelStatusCreate:
		return "VOICE_CHANNEL_STATUS_CREATE"
	case AuditLogActionTypeVoiceChannelStatusDelete:
		return "VOICE_CHANNEL_STATUS_DELETE"
	case AuditLogActionTypeClydeAiProfileUpdate:
		return "CLYDE_AI_PROFILE_UPDATE"
	case AuditLogActionTypeGuildScheduledEventExceptionCreate:
		return "GUILD_SCHEDULED_EVENT_EXCEPTION_CREATE"
	case AuditLogActionTypeGuildScheduledEventExceptionUpdate:
		return "GUILD_SCHEDULED_EVENT_EXCEPTION_UPDATE"
	case AuditLogActionTypeGuildScheduledEventExceptionDelete:
		return "GUILD_SCHEDULED_EVENT_EXCEPTION_DELETE"
	case AuditLogActionTypeGuildMemberVerificationUpdate:
		return "GUILD_MEMBER_VERIFICATION_UPDATE"
	case AuditLogActionTypeGuildProfileUpdate:
		return "GUILD_PROFILE_UPDATE"
	case AuditLogActionTypeGuildMigratePinPermission:
		return "GUILD_MIGRATE_PIN_PERMISSION"
	case AuditLogActionTypeGuildMigrateBypassSlowmodePermission:
		return "GUILD_MIGRATE_BYPASS_SLOWMODE_PERMISSION"
	}
	return strconv.FormatInt(int64(v), 10)
}

type AuthenticatorType string

const (
	AuthenticatorTypeTotp     AuthenticatorType = "totp"
	AuthenticatorTypeSMS      AuthenticatorType = "sms"
	AuthenticatorTypeBackup   AuthenticatorType = "backup"
	AuthenticatorTypeWebauthn AuthenticatorType = "webauthn"
	AuthenticatorTypePassword AuthenticatorType = "password"
)

type AuthorType string

const (
	AuthorTypeUser    AuthorType = "user"
	AuthorTypeBot     AuthorType = "bot"
	AuthorTypeWebhook AuthorType = "webhook"
)

type AutomodActionType int64

const (
	AutomodActionTypeBlockMessage     AutomodActionType = 1
	AutomodActionTypeSendAlertMessage AutomodActionType = 2
	AutomodActionTypeTimeoutUser      AutomodActionType = 3
	AutomodActionTypeQuarantineUser   AutomodActionType = 4
)

func (v AutomodActionType) String() string {
	switch v {
	case AutomodActionTypeBlockMessage:
		return "BLOCK_MESSAGE"
	case AutomodActionTypeSendAlertMessage:
		return "SEND_ALERT_MESSAGE"
	case AutomodActionTypeTimeoutUser:
		return "TIMEOUT_USER"
	case AutomodActionTypeQuarantineUser:
		return "QUARANTINE_USER"
	}
	return strconv.FormatInt(int64(v), 10)
}

type AutomodAlertActionType int64

const (
	AutomodAlertActionTypeSetCompleted      AutomodAlertActionType = 1
	AutomodAlertActionTypeUnsetCompleted    AutomodAlertActionType = 2
	AutomodAlertActionTypeDeleteUserMessage AutomodAlertActionType = 3
	AutomodAlertActionTypeSubmitFeedback    AutomodAlertActionType = 4
)

func (v AutomodAlertActionType) String() string {
	switch v {
	case AutomodAlertActionTypeSetCompleted:
		return "SET_COMPLETED"
	case AutomodAlertActionTypeUnsetCompleted:
		return "UNSET_COMPLETED"
	case AutomodAlertActionTypeDeleteUserMessage:
		return "DELETE_USER_MESSAGE"
	case AutomodAlertActionTypeSubmitFeedback:
		return "SUBMIT_FEEDBACK"
	}
	return strconv.FormatInt(int64(v), 10)
}

type AutomodDecisionOutcome string

const (
	AutomodDecisionOutcomeFlagged AutomodDecisionOutcome = "flagged"
	AutomodDecisionOutcomeBlocked AutomodDecisionOutcome = "blocked"
)

type AutomodEventType int64

const (
	AutomodEventTypeMessageSend      AutomodEventType = 1
	AutomodEventTypeGuildMemberEvent AutomodEventType = 2
)

func (v AutomodEventType) String() string {
	switch v {
	case AutomodEventTypeMessageSend:
		return "MESSAGE_SEND"
	case AutomodEventTypeGuildMemberEvent:
		return "GUILD_MEMBER_EVENT"
	}
	return strconv.FormatInt(int64(v), 10)
}

type AutomodIncidentNotificationType string

const (
	AutomodIncidentNotificationTypeActivityAlertsEnabled AutomodIncidentNotificationType = "activity_alerts_enabled"
	AutomodIncidentNotificationTypeRaid                  AutomodIncidentNotificationType = "raid"
	AutomodIncidentNotificationTypeMentionRaid           AutomodIncidentNotificationType = "mention_raid"
	AutomodIncidentNotificationTypeInteractionBlocked    AutomodIncidentNotificationType = "interaction_blocked"
)

type AutomodInteractionCallbackType string

const (
	AutomodInteractionCallbackTypeModal AutomodInteractionCallbackType = "modal"
)

type AutomodKeywordPresetType int64

const (
	AutomodKeywordPresetTypeProfanity     AutomodKeywordPresetType = 1
	AutomodKeywordPresetTypeSexualContent AutomodKeywordPresetType = 2
	AutomodKeywordPresetTypeSlurs         AutomodKeywordPresetType = 3
)

func (v AutomodKeywordPresetType) String() string {
	switch v {
	case AutomodKeywordPresetTypeProfanity:
		return "PROFANITY"
	case AutomodKeywordPresetTypeSexualContent:
		return "SEXUAL_CONTENT"
	case AutomodKeywordPresetTypeSlurs:
		return "SLURS"
	}
	return strconv.FormatInt(int64(v), 10)
}

type AutomodProfileUpdateType string

const (
	AutomodProfileUpdateTypeNicknameUpdate AutomodProfileUpdateType = "nickname_update"
	AutomodProfileUpdateTypeNicknameReset  AutomodProfileUpdateType = "nickname_reset"
)

type AutomodQuarantineEventType string

const (
	AutomodQuarantineEventTypeGuildJoin      AutomodQuarantineEventType = "guild_join"
	AutomodQuarantineEventTypeMessageSend    AutomodQuarantineEventType = "message_send"
	AutomodQuarantineEventTypeUsernameUpdate AutomodQuarantineEventType = "username_update"
	AutomodQuarantineEventTypeClanTagUpdate  AutomodQuarantineEventType = "clan_tag_update"
)

type AutomodQuarantineUserAction string

const (
	AutomodQuarantineUserActionBlockGuestJoin     AutomodQuarantineUserAction = "block_guest_join"
	AutomodQuarantineUserActionBlockProfileUpdate AutomodQuarantineUserAction = "block_profile_update"
	AutomodQuarantineUserActionQuarantineUser     AutomodQuarantineUserAction = "quarantine_user"
)

type AutomodQuarantineUserReason string

const (
	AutomodQuarantineUserReasonUsername    AutomodQuarantineUserReason = "username"
	AutomodQuarantineUserReasonDisplayName AutomodQuarantineUserReason = "display_name"
	AutomodQuarantineUserReasonBio         AutomodQuarantineUserReason = "bio"
	AutomodQuarantineUserReasonNickname    AutomodQuarantineUserReason = "nickname"
	AutomodQuarantineUserReasonClanTag     AutomodQuarantineUserReason = "clan_tag"
)

type AutomodRaidResolutionReason string

const (
	AutomodRaidResolutionReasonLegitimateActivity AutomodRaidResolutionReason = "LEGITIMATE_ACTIVITY"
	AutomodRaidResolutionReasonLegitimateAccounts AutomodRaidResolutionReason = "LEGITIMATE_ACCOUNTS"
	AutomodRaidResolutionReasonLegitimateDMs      AutomodRaidResolutionReason = "LEGITIMATE_DMS"
	AutomodRaidResolutionReasonDMSpam             AutomodRaidResolutionReason = "DM_SPAM"
	AutomodRaidResolutionReasonJoinRaid           AutomodRaidResolutionReason = "JOIN_RAID"
	AutomodRaidResolutionReasonOther              AutomodRaidResolutionReason = "OTHER"
)

type AutomodRaidType string

const (
	AutomodRaidTypeJoinRaid    AutomodRaidType = "JOIN_RAID"
	AutomodRaidTypeMentionRaid AutomodRaidType = "MENTION_RAID"
)

type AutomodTriggerType int64

const (
	AutomodTriggerTypeKeyword       AutomodTriggerType = 1
	AutomodTriggerTypeHarmfulLink   AutomodTriggerType = 2
	AutomodTriggerTypeSpam          AutomodTriggerType = 3
	AutomodTriggerTypeKeywordPreset AutomodTriggerType = 4
	AutomodTriggerTypeMentionSpam   AutomodTriggerType = 5
	AutomodTriggerTypeUserProfile   AutomodTriggerType = 6
	AutomodTriggerTypeGuildPolicy   AutomodTriggerType = 7
)

func (v AutomodTriggerType) String() string {
	switch v {
	case AutomodTriggerTypeKeyword:
		return "KEYWORD"
	case AutomodTriggerTypeHarmfulLink:
		return "HARMFUL_LINK"
	case AutomodTriggerTypeSpam:
		return "SPAM"
	case AutomodTriggerTypeKeywordPreset:
		return "KEYWORD_PRESET"
	case AutomodTriggerTypeMentionSpam:
		return "MENTION_SPAM"
	case AutomodTriggerTypeUserProfile:
		return "USER_PROFILE"
	case AutomodTriggerTypeGuildPolicy:
		return "GUILD_POLICY"
	}
	return strconv.FormatInt(int64(v), 10)
}

type ButtonStyle int64

const (
	ButtonStylePrimary   ButtonStyle = 1
	ButtonStyleSecondary ButtonStyle = 2
	ButtonStyleSuccess   ButtonStyle = 3
	ButtonStyleDanger    ButtonStyle = 4
	ButtonStyleLink      ButtonStyle = 5
	ButtonStylePremium   ButtonStyle = 6
)

func (v ButtonStyle) String() string {
	switch v {
	case ButtonStylePrimary:
		return "PRIMARY"
	case ButtonStyleSecondary:
		return "SECONDARY"
	case ButtonStyleSuccess:
		return "SUCCESS"
	case ButtonStyleDanger:
		return "DANGER"
	case ButtonStyleLink:
		return "LINK"
	case ButtonStylePremium:
		return "PREMIUM"
	}
	return strconv.FormatInt(int64(v), 10)
}

type CaptchaService string

const (
	CaptchaServiceHcaptcha            CaptchaService = "hcaptcha"
	CaptchaServiceRecaptcha           CaptchaService = "recaptcha"
	CaptchaServiceRecaptchaEnterprise CaptchaService = "recaptcha_enterprise"
)

type CertifiedDeviceType string

const (
	CertifiedDeviceTypeAudioinput  CertifiedDeviceType = "audioinput"
	CertifiedDeviceTypeAudiooutput CertifiedDeviceType = "audiooutput"
	CertifiedDeviceTypeVideoinput  CertifiedDeviceType = "videoinput"
)

type ChangePhoneReason string

const (
	ChangePhoneReasonUserActionRequired ChangePhoneReason = "user_action_required"
	ChangePhoneReasonUserSettingsUpdate ChangePhoneReason = "user_settings_update"
	ChangePhoneReasonGuildPhoneRequired ChangePhoneReason = "guild_phone_required"
	ChangePhoneReasonMFAPhoneUpdate     ChangePhoneReason = "mfa_phone_update"
	ChangePhoneReasonContactSync        ChangePhoneReason = "contact_sync"
)

type ChannelConsentStatus int64

const (
	ChannelConsentStatusUnspecified ChannelConsentStatus = 0
	ChannelConsentStatusPending     ChannelConsentStatus = 1
	ChannelConsentStatusAccepted    ChannelConsentStatus = 2
	ChannelConsentStatusRejected    ChannelConsentStatus = 3
)

func (v ChannelConsentStatus) String() string {
	switch v {
	case ChannelConsentStatusUnspecified:
		return "UNSPECIFIED"
	case ChannelConsentStatusPending:
		return "PENDING"
	case ChannelConsentStatusAccepted:
		return "ACCEPTED"
	case ChannelConsentStatusRejected:
		return "REJECTED"
	}
	return strconv.FormatInt(int64(v), 10)
}

type ChannelFlags uint64

const (
	ChannelFlagsGuildFeedRemoved                         ChannelFlags = 1 << 0
	ChannelFlagsPinned                                   ChannelFlags = 1 << 1
	ChannelFlagsActiveChannelsRemoved                    ChannelFlags = 1 << 2
	ChannelFlagsRequireTag                               ChannelFlags = 1 << 4
	ChannelFlagsIsSpam                                   ChannelFlags = 1 << 5
	ChannelFlagsIsGuildResourceChannel                   ChannelFlags = 1 << 7
	ChannelFlagsClydeAi                                  ChannelFlags = 1 << 8
	ChannelFlagsIsScheduledForDeletion                   ChannelFlags = 1 << 9
	ChannelFlagsIsMediaChannel                           ChannelFlags = 1 << 10
	ChannelFlagsSummariesDisabled                        ChannelFlags = 1 << 11
	ChannelFlagsApplicationShelfConsent                  ChannelFlags = 1 << 12
	ChannelFlagsIsRoleSubscriptionTemplatePreviewChannel ChannelFlags = 1 << 13
	ChannelFlagsIsBroadcasting                           ChannelFlags = 1 << 14
	ChannelFlagsHideMediaDownloadOptions                 ChannelFlags = 1 << 15
	ChannelFlagsIsJoinRequestInterviewChannel            ChannelFlags = 1 << 16
	ChannelFlagsObfuscated                               ChannelFlags = 1 << 17
	ChannelFlagsIsModeratorReportChannel                 ChannelFlags = 1 << 19
	ChannelFlagsIsSpoilerChannel                         ChannelFlags = 1 << 21
)

func (v ChannelFlags) String() string {
	switch v {
	case ChannelFlagsGuildFeedRemoved:
		return "GUILD_FEED_REMOVED"
	case ChannelFlagsPinned:
		return "PINNED"
	case ChannelFlagsActiveChannelsRemoved:
		return "ACTIVE_CHANNELS_REMOVED"
	case ChannelFlagsRequireTag:
		return "REQUIRE_TAG"
	case ChannelFlagsIsSpam:
		return "IS_SPAM"
	case ChannelFlagsIsGuildResourceChannel:
		return "IS_GUILD_RESOURCE_CHANNEL"
	case ChannelFlagsClydeAi:
		return "CLYDE_AI"
	case ChannelFlagsIsScheduledForDeletion:
		return "IS_SCHEDULED_FOR_DELETION"
	case ChannelFlagsIsMediaChannel:
		return "IS_MEDIA_CHANNEL"
	case ChannelFlagsSummariesDisabled:
		return "SUMMARIES_DISABLED"
	case ChannelFlagsApplicationShelfConsent:
		return "APPLICATION_SHELF_CONSENT"
	case ChannelFlagsIsRoleSubscriptionTemplatePreviewChannel:
		return "IS_ROLE_SUBSCRIPTION_TEMPLATE_PREVIEW_CHANNEL"
	case ChannelFlagsIsBroadcasting:
		return "IS_BROADCASTING"
	case ChannelFlagsHideMediaDownloadOptions:
		return "HIDE_MEDIA_DOWNLOAD_OPTIONS"
	case ChannelFlagsIsJoinRequestInterviewChannel:
		return "IS_JOIN_REQUEST_INTERVIEW_CHANNEL"
	case ChannelFlagsObfuscated:
		return "OBFUSCATED"
	case ChannelFlagsIsModeratorReportChannel:
		return "IS_MODERATOR_REPORT_CHANNEL"
	case ChannelFlagsIsSpoilerChannel:
		return "IS_SPOILER_CHANNEL"
	}
	return strconv.FormatUint(uint64(v), 10)
}

type ChannelOverrideFlags uint64

const (
	ChannelOverrideFlagsUnreadsOnlyMentions ChannelOverrideFlags = 1 << 9
	ChannelOverrideFlagsUnreadsAllMessages  ChannelOverrideFlags = 1 << 10
	ChannelOverrideFlagsFavorited           ChannelOverrideFlags = 1 << 11
	ChannelOverrideFlagsOptInEnabled        ChannelOverrideFlags = 1 << 12
	ChannelOverrideFlagsNewForumThreadsOff  ChannelOverrideFlags = 1 << 13
	ChannelOverrideFlagsNewForumThreadsOn   ChannelOverrideFlags = 1 << 14
)

func (v ChannelOverrideFlags) String() string {
	switch v {
	case ChannelOverrideFlagsUnreadsOnlyMentions:
		return "UNREADS_ONLY_MENTIONS"
	case ChannelOverrideFlagsUnreadsAllMessages:
		return "UNREADS_ALL_MESSAGES"
	case ChannelOverrideFlagsFavorited:
		return "FAVORITED"
	case ChannelOverrideFlagsOptInEnabled:
		return "OPT_IN_ENABLED"
	case ChannelOverrideFlagsNewForumThreadsOff:
		return "NEW_FORUM_THREADS_OFF"
	case ChannelOverrideFlagsNewForumThreadsOn:
		return "NEW_FORUM_THREADS_ON"
	}
	return strconv.FormatUint(uint64(v), 10)
}

type ChannelType int64

const (
	ChannelTypeGuildText       ChannelType = 0
	ChannelTypeDM              ChannelType = 1
	ChannelTypeGuildVoice      ChannelType = 2
	ChannelTypeGroupDM         ChannelType = 3
	ChannelTypeGuildCategory   ChannelType = 4
	ChannelTypeGuildNews       ChannelType = 5
	ChannelTypeGuildStore      ChannelType = 6
	ChannelTypeGuildLfg        ChannelType = 7
	ChannelTypeLfgGroupDM      ChannelType = 8
	ChannelTypeThreadAlpha     ChannelType = 9
	ChannelTypeNewsThread      ChannelType = 10
	ChannelTypePublicThread    ChannelType = 11
	ChannelTypePrivateThread   ChannelType = 12
	ChannelTypeGuildStageVoice ChannelType = 13
	ChannelTypeGuildDirectory  ChannelType = 14
	ChannelTypeGuildForum      ChannelType = 15
	ChannelTypeGuildMedia      ChannelType = 16
	ChannelTypeLobby           ChannelType = 17
	ChannelTypeEphemeralDM     ChannelType = 18
)

func (v ChannelType) String() string {
	switch v {
	case ChannelTypeGuildText:
		return "GUILD_TEXT"
	case ChannelTypeDM:
		return "DM"
	case ChannelTypeGuildVoice:
		return "GUILD_VOICE"
	case ChannelTypeGroupDM:
		return "GROUP_DM"
	case ChannelTypeGuildCategory:
		return "GUILD_CATEGORY"
	case ChannelTypeGuildNews:
		return "GUILD_NEWS"
	case ChannelTypeGuildStore:
		return "GUILD_STORE"
	case ChannelTypeGuildLfg:
		return "GUILD_LFG"
	case ChannelTypeLfgGroupDM:
		return "LFG_GROUP_DM"
	case ChannelTypeThreadAlpha:
		return "THREAD_ALPHA"
	case ChannelTypeNewsThread:
		return "NEWS_THREAD"
	case ChannelTypePublicThread:
		return "PUBLIC_THREAD"
	case ChannelTypePrivateThread:
		return "PRIVATE_THREAD"
	case ChannelTypeGuildStageVoice:
		return "GUILD_STAGE_VOICE"
	case ChannelTypeGuildDirectory:
		return "GUILD_DIRECTORY"
	case ChannelTypeGuildForum:
		return "GUILD_FORUM"
	case ChannelTypeGuildMedia:
		return "GUILD_MEDIA"
	case ChannelTypeLobby:
		return "LOBBY"
	case ChannelTypeEphemeralDM:
		return "EPHEMERAL_DM"
	}
	return strconv.FormatInt(int64(v), 10)
}

type CheckpointCardType int64

const (
	CheckpointCardTypePlant              CheckpointCardType = 0
	CheckpointCardTypeDonut              CheckpointCardType = 1
	CheckpointCardTypeDogInSwimRing      CheckpointCardType = 2
	CheckpointCardTypeDiscoBall          CheckpointCardType = 3
	CheckpointCardTypeOrigamiPaperCrane  CheckpointCardType = 4
	CheckpointCardTypeSnail              CheckpointCardType = 5
	CheckpointCardTypeDuckWithSunglasses CheckpointCardType = 6
	CheckpointCardTypeBanana             CheckpointCardType = 7
	CheckpointCardTypeCat                CheckpointCardType = 8
	CheckpointCardTypeCartridge          CheckpointCardType = 9
)

func (v CheckpointCardType) String() string {
	switch v {
	case CheckpointCardTypePlant:
		return "PLANT"
	case CheckpointCardTypeDonut:
		return "DONUT"
	case CheckpointCardTypeDogInSwimRing:
		return "DOG_IN_SWIM_RING"
	case CheckpointCardTypeDiscoBall:
		return "DISCO_BALL"
	case CheckpointCardTypeOrigamiPaperCrane:
		return "ORIGAMI_PAPER_CRANE"
	case CheckpointCardTypeSnail:
		return "SNAIL"
	case CheckpointCardTypeDuckWithSunglasses:
		return "DUCK_WITH_SUNGLASSES"
	case CheckpointCardTypeBanana:
		return "BANANA"
	case CheckpointCardTypeCat:
		return "CAT"
	case CheckpointCardTypeCartridge:
		return "CARTRIDGE"
	}
	return strconv.FormatInt(int64(v), 10)
}

type ClassificationType int64

const (
	ClassificationTypeUnknown                          ClassificationType = 1
	ClassificationTypeUnsolicitedPornography           ClassificationType = 100
	ClassificationTypeNonconsensualPornography         ClassificationType = 200
	ClassificationTypeGlorifyingViolence               ClassificationType = 210
	ClassificationTypeHateSpeech                       ClassificationType = 220
	ClassificationTypeCrackedAccounts                  ClassificationType = 230
	ClassificationTypeIllicitGoods                     ClassificationType = 240
	ClassificationTypeSocialEngineering                ClassificationType = 250
	ClassificationTypeChildSafety                      ClassificationType = 280
	ClassificationTypeHarrasmentAndBullying            ClassificationType = 290
	ClassificationTypeHarrasmentAndBullying2           ClassificationType = 310
	ClassificationTypeHatefulConduct                   ClassificationType = 320
	ClassificationTypeHarrasmentAndBullying3           ClassificationType = 390
	ClassificationTypeChildSafety2                     ClassificationType = 600
	ClassificationTypeChildSafety3                     ClassificationType = 650
	ClassificationTypeImpersonation                    ClassificationType = 711
	ClassificationTypeBanEvasion                       ClassificationType = 720
	ClassificationTypeMaliciousConduct                 ClassificationType = 3010
	ClassificationTypeSpam                             ClassificationType = 3030
	ClassificationTypeNonconsensualAdultContent        ClassificationType = 4000
	ClassificationTypeFraud                            ClassificationType = 4010
	ClassificationTypeDoxxingGuildOwner                ClassificationType = 4130
	ClassificationTypeCopyrightInfringementGuildOwner  ClassificationType = 4140
	ClassificationTypeChildSafety4                     ClassificationType = 5010
	ClassificationTypeChildSelfEndangerment            ClassificationType = 5090
	ClassificationTypeHarassmentAndBullyingGuildMember ClassificationType = 5245
	ClassificationTypeDoxxingGuildMember               ClassificationType = 5305
	ClassificationTypeUnderage                         ClassificationType = 5411
	ClassificationTypeCopyrightInfringementGuildMember ClassificationType = 5440
	ClassificationTypeCopyrightInfringement3           ClassificationType = 5485
)

func (v ClassificationType) String() string {
	switch v {
	case ClassificationTypeUnknown:
		return "UNKNOWN"
	case ClassificationTypeUnsolicitedPornography:
		return "UNSOLICITED_PORNOGRAPHY"
	case ClassificationTypeNonconsensualPornography:
		return "NONCONSENSUAL_PORNOGRAPHY"
	case ClassificationTypeGlorifyingViolence:
		return "GLORIFYING_VIOLENCE"
	case ClassificationTypeHateSpeech:
		return "HATE_SPEECH"
	case ClassificationTypeCrackedAccounts:
		return "CRACKED_ACCOUNTS"
	case ClassificationTypeIllicitGoods:
		return "ILLICIT_GOODS"
	case ClassificationTypeSocialEngineering:
		return "SOCIAL_ENGINEERING"
	case ClassificationTypeChildSafety:
		return "CHILD_SAFETY"
	case ClassificationTypeHarrasmentAndBullying:
		return "HARRASMENT_AND_BULLYING"
	case ClassificationTypeHarrasmentAndBullying2:
		return "HARRASMENT_AND_BULLYING_2"
	case ClassificationTypeHatefulConduct:
		return "HATEFUL_CONDUCT"
	case ClassificationTypeHarrasmentAndBullying3:
		return "HARRASMENT_AND_BULLYING_3"
	case ClassificationTypeChildSafety2:
		return "CHILD_SAFETY_2"
	case ClassificationTypeChildSafety3:
		return "CHILD_SAFETY_3"
	case ClassificationTypeImpersonation:
		return "IMPERSONATION"
	case ClassificationTypeBanEvasion:
		return "BAN_EVASION"
	case ClassificationTypeMaliciousConduct:
		return "MALICIOUS_CONDUCT"
	case ClassificationTypeSpam:
		return "SPAM"
	case ClassificationTypeNonconsensualAdultContent:
		return "NONCONSENSUAL_ADULT_CONTENT"
	case ClassificationTypeFraud:
		return "FRAUD"
	case ClassificationTypeDoxxingGuildOwner:
		return "DOXXING_GUILD_OWNER"
	case ClassificationTypeCopyrightInfringementGuildOwner:
		return "COPYRIGHT_INFRINGEMENT_GUILD_OWNER"
	case ClassificationTypeChildSafety4:
		return "CHILD_SAFETY_4"
	case ClassificationTypeChildSelfEndangerment:
		return "CHILD_SELF_ENDANGERMENT"
	case ClassificationTypeHarassmentAndBullyingGuildMember:
		return "HARASSMENT_AND_BULLYING_GUILD_MEMBER"
	case ClassificationTypeDoxxingGuildMember:
		return "DOXXING_GUILD_MEMBER"
	case ClassificationTypeUnderage:
		return "UNDERAGE"
	case ClassificationTypeCopyrightInfringementGuildMember:
		return "COPYRIGHT_INFRINGEMENT_GUILD_MEMBER"
	case ClassificationTypeCopyrightInfringement3:
		return "COPYRIGHT_INFRINGEMENT_3"
	}
	return strconv.FormatInt(int64(v), 10)
}

type ClientAppState string

const (
	ClientAppStateFocused    ClientAppState = "focused"
	ClientAppStateUnfocused  ClientAppState = "unfocused"
	ClientAppStateActive     ClientAppState = "active"
	ClientAppStateInactive   ClientAppState = "inactive"
	ClientAppStateBackground ClientAppState = "background"
)

type ClientEventSource string

const (
	ClientEventSourceOverlay ClientEventSource = "OVERLAY"
)

type ClientType string

const (
	ClientTypeDesktop ClientType = "desktop"
	ClientTypeWeb     ClientType = "web"
	ClientTypeMobile  ClientType = "mobile"
	ClientTypeVR      ClientType = "vr"
	ClientTypeUnknown ClientType = "unknown"
)

type CollectibleProductType int64

const (
	CollectibleProductTypeAvatarDecoration CollectibleProductType = 0
	CollectibleProductTypeProfileEffect    CollectibleProductType = 1
	CollectibleProductTypeNameplate        CollectibleProductType = 2
	CollectibleProductTypeProfileFrame     CollectibleProductType = 3
	CollectibleProductTypeBundle           CollectibleProductType = 1000
	CollectibleProductTypeVariantsGroup    CollectibleProductType = 2000
	CollectibleProductTypeExternalSKU      CollectibleProductType = 3000
)

func (v CollectibleProductType) String() string {
	switch v {
	case CollectibleProductTypeAvatarDecoration:
		return "AVATAR_DECORATION"
	case CollectibleProductTypeProfileEffect:
		return "PROFILE_EFFECT"
	case CollectibleProductTypeNameplate:
		return "NAMEPLATE"
	case CollectibleProductTypeProfileFrame:
		return "PROFILE_FRAME"
	case CollectibleProductTypeBundle:
		return "BUNDLE"
	case CollectibleProductTypeVariantsGroup:
		return "VARIANTS_GROUP"
	case CollectibleProductTypeExternalSKU:
		return "EXTERNAL_SKU"
	}
	return strconv.FormatInt(int64(v), 10)
}

type CollectiblePurchaseType int64

const (
	CollectiblePurchaseTypePurchased    CollectiblePurchaseType = 1
	CollectiblePurchaseTypePromotional  CollectiblePurchaseType = 5
	CollectiblePurchaseTypeGifted       CollectiblePurchaseType = 6
	CollectiblePurchaseTypeSubscription CollectiblePurchaseType = 7
	CollectiblePurchaseTypeQuest        CollectiblePurchaseType = 10
)

func (v CollectiblePurchaseType) String() string {
	switch v {
	case CollectiblePurchaseTypePurchased:
		return "PURCHASED"
	case CollectiblePurchaseTypePromotional:
		return "PROMOTIONAL"
	case CollectiblePurchaseTypeGifted:
		return "GIFTED"
	case CollectiblePurchaseTypeSubscription:
		return "SUBSCRIPTION"
	case CollectiblePurchaseTypeQuest:
		return "QUEST"
	}
	return strconv.FormatInt(int64(v), 10)
}

type CollectiblesMarketingPlatformType int64

const (
	CollectiblesMarketingPlatformTypeDesktop CollectiblesMarketingPlatformType = 0
	CollectiblesMarketingPlatformTypeMobile  CollectiblesMarketingPlatformType = 1
)

func (v CollectiblesMarketingPlatformType) String() string {
	switch v {
	case CollectiblesMarketingPlatformTypeDesktop:
		return "DESKTOP"
	case CollectiblesMarketingPlatformTypeMobile:
		return "MOBILE"
	}
	return strconv.FormatInt(int64(v), 10)
}

type CollectiblesMarketingReleaseType int64

const (
	CollectiblesMarketingReleaseTypeProd CollectiblesMarketingReleaseType = 0
	CollectiblesMarketingReleaseTypeBeta CollectiblesMarketingReleaseType = 1
)

func (v CollectiblesMarketingReleaseType) String() string {
	switch v {
	case CollectiblesMarketingReleaseTypeProd:
		return "PROD"
	case CollectiblesMarketingReleaseTypeBeta:
		return "BETA"
	}
	return strconv.FormatInt(int64(v), 10)
}

type CollectiblesMarketingSurfaceType int64

const (
	CollectiblesMarketingSurfaceTypeDesktopShopButton   CollectiblesMarketingSurfaceType = 0
	CollectiblesMarketingSurfaceTypeMobileShopButton    CollectiblesMarketingSurfaceType = 1
	CollectiblesMarketingSurfaceTypeEditProfileSettings CollectiblesMarketingSurfaceType = 2
)

func (v CollectiblesMarketingSurfaceType) String() string {
	switch v {
	case CollectiblesMarketingSurfaceTypeDesktopShopButton:
		return "DESKTOP_SHOP_BUTTON"
	case CollectiblesMarketingSurfaceTypeMobileShopButton:
		return "MOBILE_SHOP_BUTTON"
	case CollectiblesMarketingSurfaceTypeEditProfileSettings:
		return "EDIT_PROFILE_SETTINGS"
	}
	return strconv.FormatInt(int64(v), 10)
}

type CollectiblesMarketingType int64

const (
	CollectiblesMarketingTypeCoachtip   CollectiblesMarketingType = 0
	CollectiblesMarketingTypeBadge      CollectiblesMarketingType = 1
	CollectiblesMarketingTypeBanner     CollectiblesMarketingType = 2
	CollectiblesMarketingTypeCoachmark  CollectiblesMarketingType = 3
	CollectiblesMarketingTypeTabTooltip CollectiblesMarketingType = 4
)

func (v CollectiblesMarketingType) String() string {
	switch v {
	case CollectiblesMarketingTypeCoachtip:
		return "COACHTIP"
	case CollectiblesMarketingTypeBadge:
		return "BADGE"
	case CollectiblesMarketingTypeBanner:
		return "BANNER"
	case CollectiblesMarketingTypeCoachmark:
		return "COACHMARK"
	case CollectiblesMarketingTypeTabTooltip:
		return "TAB_TOOLTIP"
	}
	return strconv.FormatInt(int64(v), 10)
}

type CollectiblesShopHomeConfig string

const (
	CollectiblesShopHomeConfigDefault             CollectiblesShopHomeConfig = "default"
	CollectiblesShopHomeConfigDefaultWithOrbShelf CollectiblesShopHomeConfig = "default_with_orb_shelf"
	CollectiblesShopHomeConfigOrbTab              CollectiblesShopHomeConfig = "orb_tab"
	CollectiblesShopHomeConfigSummerSaleTakeover  CollectiblesShopHomeConfig = "summer_sale_takeover"
)

type CollectiblesShopHomeTab string

const (
	CollectiblesShopHomeTabHome              CollectiblesShopHomeTab = "home"
	CollectiblesShopHomeTabCatalog           CollectiblesShopHomeTab = "catalog"
	CollectiblesShopHomeTabOrbs              CollectiblesShopHomeTab = "orbs"
	CollectiblesShopHomeTabAvatarDecorations CollectiblesShopHomeTab = "avatar-decorations"
	CollectiblesShopHomeTabProfileEffects    CollectiblesShopHomeTab = "profile-effects"
	CollectiblesShopHomeTabNameplates        CollectiblesShopHomeTab = "nameplates"
	CollectiblesShopHomeTabProfileFrames     CollectiblesShopHomeTab = "profile-frames"
	CollectiblesShopHomeTabBundles           CollectiblesShopHomeTab = "bundles"
	CollectiblesShopHomeTabLayout            CollectiblesShopHomeTab = "layout"
	CollectiblesShopHomeTabCollectionIndex   CollectiblesShopHomeTab = "collection-index"
	CollectiblesShopHomeTabGameShops         CollectiblesShopHomeTab = "game-shops"
)

type ComponentType int64

const (
	ComponentTypeActionRow             ComponentType = 1
	ComponentTypeButton                ComponentType = 2
	ComponentTypeStringSelect          ComponentType = 3
	ComponentTypeTextInput             ComponentType = 4
	ComponentTypeUserSelect            ComponentType = 5
	ComponentTypeRoleSelect            ComponentType = 6
	ComponentTypeMentionableSelect     ComponentType = 7
	ComponentTypeChannelSelect         ComponentType = 8
	ComponentTypeSection               ComponentType = 9
	ComponentTypeTextDisplay           ComponentType = 10
	ComponentTypeThumbnail             ComponentType = 11
	ComponentTypeMediaGallery          ComponentType = 12
	ComponentTypeFile                  ComponentType = 13
	ComponentTypeSeparator             ComponentType = 14
	ComponentTypeContentInventoryEntry ComponentType = 16
	ComponentTypeContainer             ComponentType = 17
	ComponentTypeLabel                 ComponentType = 18
	ComponentTypeFileUpload            ComponentType = 19
	ComponentTypeCheckpointCard        ComponentType = 20
	ComponentTypeRadioGroup            ComponentType = 21
	ComponentTypeCheckboxGroup         ComponentType = 22
	ComponentTypeCheckbox              ComponentType = 23
)

func (v ComponentType) String() string {
	switch v {
	case ComponentTypeActionRow:
		return "ACTION_ROW"
	case ComponentTypeButton:
		return "BUTTON"
	case ComponentTypeStringSelect:
		return "STRING_SELECT"
	case ComponentTypeTextInput:
		return "TEXT_INPUT"
	case ComponentTypeUserSelect:
		return "USER_SELECT"
	case ComponentTypeRoleSelect:
		return "ROLE_SELECT"
	case ComponentTypeMentionableSelect:
		return "MENTIONABLE_SELECT"
	case ComponentTypeChannelSelect:
		return "CHANNEL_SELECT"
	case ComponentTypeSection:
		return "SECTION"
	case ComponentTypeTextDisplay:
		return "TEXT_DISPLAY"
	case ComponentTypeThumbnail:
		return "THUMBNAIL"
	case ComponentTypeMediaGallery:
		return "MEDIA_GALLERY"
	case ComponentTypeFile:
		return "FILE"
	case ComponentTypeSeparator:
		return "SEPARATOR"
	case ComponentTypeContentInventoryEntry:
		return "CONTENT_INVENTORY_ENTRY"
	case ComponentTypeContainer:
		return "CONTAINER"
	case ComponentTypeLabel:
		return "LABEL"
	case ComponentTypeFileUpload:
		return "FILE_UPLOAD"
	case ComponentTypeCheckpointCard:
		return "CHECKPOINT_CARD"
	case ComponentTypeRadioGroup:
		return "RADIO_GROUP"
	case ComponentTypeCheckboxGroup:
		return "CHECKBOX_GROUP"
	case ComponentTypeCheckbox:
		return "CHECKBOX"
	}
	return strconv.FormatInt(int64(v), 10)
}

type ConnectionType string

const (
	ConnectionTypeAmazonMusic               ConnectionType = "amazon-music"
	ConnectionTypeBattleNet                 ConnectionType = "battlenet"
	ConnectionTypeBluesky                   ConnectionType = "bluesky"
	ConnectionTypeBungieNet                 ConnectionType = "bungie"
	ConnectionTypeContactSync               ConnectionType = "contacts"
	ConnectionTypeCrunchyroll               ConnectionType = "crunchyroll"
	ConnectionTypeDomain                    ConnectionType = "domain"
	ConnectionTypeEbay                      ConnectionType = "ebay"
	ConnectionTypeEpicGames                 ConnectionType = "epicgames"
	ConnectionTypeFacebook                  ConnectionType = "facebook"
	ConnectionTypeGithub                    ConnectionType = "github"
	ConnectionTypeInstagram                 ConnectionType = "instagram"
	ConnectionTypeLeagueOfLegends           ConnectionType = "leagueoflegends"
	ConnectionTypeMastodon                  ConnectionType = "mastodon"
	ConnectionTypePaypal                    ConnectionType = "paypal"
	ConnectionTypePlaystationNetwork        ConnectionType = "playstation"
	ConnectionTypePlaystationNetworkStaging ConnectionType = "playstation-stg"
	ConnectionTypeReddit                    ConnectionType = "reddit"
	ConnectionTypeRoblox                    ConnectionType = "roblox"
	ConnectionTypeRiotGames                 ConnectionType = "riotgames"
	ConnectionTypeSamsungGalaxy             ConnectionType = "samsung"
	ConnectionTypeSoundcloud                ConnectionType = "soundcloud"
	ConnectionTypeSpotify                   ConnectionType = "spotify"
	ConnectionTypeSkype                     ConnectionType = "skype"
	ConnectionTypeSteam                     ConnectionType = "steam"
	ConnectionTypeTiktok                    ConnectionType = "tiktok"
	ConnectionTypeTwitch                    ConnectionType = "twitch"
	ConnectionTypeTwitter                   ConnectionType = "twitter"
	ConnectionTypeXbox                      ConnectionType = "xbox"
	ConnectionTypeYoutube                   ConnectionType = "youtube"
)

type ConsentType string

const (
	ConsentTypePersonalization ConsentType = "personalization"
	ConsentTypeUsageStatistics ConsentType = "usage_statistics"
)

type ConsoleCommandType string

const (
	ConsoleCommandTypeConnectVoice ConsoleCommandType = "connect_voice"
)

type ConsoleHandoffType string

const (
	ConsoleHandoffTypeCreateNewCall        ConsoleHandoffType = "CREATE_NEW_CALL"
	ConsoleHandoffTypeTransferExistingCall ConsoleHandoffType = "TRANSFER_EXISTING_CALL"
)

type ContactSyncSuggestionsSetting int64

const (
	ContactSyncSuggestionsSettingMutualContactInfoOnly ContactSyncSuggestionsSetting = 1
	ContactSyncSuggestionsSettingAnyoneWithContactInfo ContactSyncSuggestionsSetting = 2
)

func (v ContactSyncSuggestionsSetting) String() string {
	switch v {
	case ContactSyncSuggestionsSettingMutualContactInfoOnly:
		return "MUTUAL_CONTACT_INFO_ONLY"
	case ContactSyncSuggestionsSettingAnyoneWithContactInfo:
		return "ANYONE_WITH_CONTACT_INFO"
	}
	return strconv.FormatInt(int64(v), 10)
}

type ContentInventoryAggregateRangeType int64

const (
	ContentInventoryAggregateRangeTypeAggregateRangeUnspecified ContentInventoryAggregateRangeType = 0
	ContentInventoryAggregateRangeTypeWeek                      ContentInventoryAggregateRangeType = 1
)

func (v ContentInventoryAggregateRangeType) String() string {
	switch v {
	case ContentInventoryAggregateRangeTypeAggregateRangeUnspecified:
		return "AGGREGATE_RANGE_UNSPECIFIED"
	case ContentInventoryAggregateRangeTypeWeek:
		return "WEEK"
	}
	return strconv.FormatInt(int64(v), 10)
}

type ContentInventoryAuthorType int64

const (
	ContentInventoryAuthorTypeAuthorTypeUnspecified ContentInventoryAuthorType = 0
	ContentInventoryAuthorTypeUser                  ContentInventoryAuthorType = 1
)

func (v ContentInventoryAuthorType) String() string {
	switch v {
	case ContentInventoryAuthorTypeAuthorTypeUnspecified:
		return "AUTHOR_TYPE_UNSPECIFIED"
	case ContentInventoryAuthorTypeUser:
		return "USER"
	}
	return strconv.FormatInt(int64(v), 10)
}

type ContentInventoryContentType int64

const (
	ContentInventoryContentTypeContentTypeUnspecified ContentInventoryContentType = 0
	ContentInventoryContentTypePlayedGame             ContentInventoryContentType = 1
	ContentInventoryContentTypeWatchedMedia           ContentInventoryContentType = 2
	ContentInventoryContentTypeTopGame                ContentInventoryContentType = 3
	ContentInventoryContentTypeListenedMedia          ContentInventoryContentType = 4
	ContentInventoryContentTypeListenedSession        ContentInventoryContentType = 5
	ContentInventoryContentTypeTopArtist              ContentInventoryContentType = 6
	ContentInventoryContentTypeCustomStatus           ContentInventoryContentType = 7
	ContentInventoryContentTypeLaunchedActivity       ContentInventoryContentType = 8
	ContentInventoryContentTypeLeaderboard            ContentInventoryContentType = 9
)

func (v ContentInventoryContentType) String() string {
	switch v {
	case ContentInventoryContentTypeContentTypeUnspecified:
		return "CONTENT_TYPE_UNSPECIFIED"
	case ContentInventoryContentTypePlayedGame:
		return "PLAYED_GAME"
	case ContentInventoryContentTypeWatchedMedia:
		return "WATCHED_MEDIA"
	case ContentInventoryContentTypeTopGame:
		return "TOP_GAME"
	case ContentInventoryContentTypeListenedMedia:
		return "LISTENED_MEDIA"
	case ContentInventoryContentTypeListenedSession:
		return "LISTENED_SESSION"
	case ContentInventoryContentTypeTopArtist:
		return "TOP_ARTIST"
	case ContentInventoryContentTypeCustomStatus:
		return "CUSTOM_STATUS"
	case ContentInventoryContentTypeLaunchedActivity:
		return "LAUNCHED_ACTIVITY"
	case ContentInventoryContentTypeLeaderboard:
		return "LEADERBOARD"
	}
	return strconv.FormatInt(int64(v), 10)
}

type ContentInventoryMediaType int64

const (
	ContentInventoryMediaTypeTopArtist ContentInventoryMediaType = 0
	ContentInventoryMediaTypeTrack     ContentInventoryMediaType = 1
)

func (v ContentInventoryMediaType) String() string {
	switch v {
	case ContentInventoryMediaTypeTopArtist:
		return "TOP_ARTIST"
	case ContentInventoryMediaTypeTrack:
		return "TRACK"
	}
	return strconv.FormatInt(int64(v), 10)
}

type ContentInventoryMetadataType string

const (
	ContentInventoryMetadataTypePlayedGameExtra       ContentInventoryMetadataType = "played_game_extra"
	ContentInventoryMetadataTypeListenedSessionExtra  ContentInventoryMetadataType = "listened_session_extra"
	ContentInventoryMetadataTypeListenedMediaExtra    ContentInventoryMetadataType = "listened_media_extra"
	ContentInventoryMetadataTypeTopArtistExtra        ContentInventoryMetadataType = "top_artist_extra"
	ContentInventoryMetadataTypeWatchedMediaExtra     ContentInventoryMetadataType = "watched_media_extra"
	ContentInventoryMetadataTypeLaunchedActivityExtra ContentInventoryMetadataType = "launched_activity_extra"
)

type ContentInventoryPlatformType int64

const (
	ContentInventoryPlatformTypeDesktop     ContentInventoryPlatformType = 0
	ContentInventoryPlatformTypeXbox        ContentInventoryPlatformType = 1
	ContentInventoryPlatformTypePlaystation ContentInventoryPlatformType = 2
	ContentInventoryPlatformTypeIos         ContentInventoryPlatformType = 3
	ContentInventoryPlatformTypeAndroid     ContentInventoryPlatformType = 4
	ContentInventoryPlatformTypeNintendo    ContentInventoryPlatformType = 5
	ContentInventoryPlatformTypeLinux       ContentInventoryPlatformType = 6
	ContentInventoryPlatformTypeMacos       ContentInventoryPlatformType = 7
)

func (v ContentInventoryPlatformType) String() string {
	switch v {
	case ContentInventoryPlatformTypeDesktop:
		return "DESKTOP"
	case ContentInventoryPlatformTypeXbox:
		return "XBOX"
	case ContentInventoryPlatformTypePlaystation:
		return "PLAYSTATION"
	case ContentInventoryPlatformTypeIos:
		return "IOS"
	case ContentInventoryPlatformTypeAndroid:
		return "ANDROID"
	case ContentInventoryPlatformTypeNintendo:
		return "NINTENDO"
	case ContentInventoryPlatformTypeLinux:
		return "LINUX"
	case ContentInventoryPlatformTypeMacos:
		return "MACOS"
	}
	return strconv.FormatInt(int64(v), 10)
}

type ContentInventoryProviderType int64

const (
	ContentInventoryProviderTypeProviderUnspecified ContentInventoryProviderType = 0
	ContentInventoryProviderTypeSpotify             ContentInventoryProviderType = 1
)

func (v ContentInventoryProviderType) String() string {
	switch v {
	case ContentInventoryProviderTypeProviderUnspecified:
		return "PROVIDER_UNSPECIFIED"
	case ContentInventoryProviderTypeSpotify:
		return "SPOTIFY"
	}
	return strconv.FormatInt(int64(v), 10)
}

type ContentInventoryTraitType int64

const (
	ContentInventoryTraitTypeTraitTypeUnspecified    ContentInventoryTraitType = 0
	ContentInventoryTraitTypeFirstTime               ContentInventoryTraitType = 1
	ContentInventoryTraitTypeDurationSeconds         ContentInventoryTraitType = 2
	ContentInventoryTraitTypeIsLive                  ContentInventoryTraitType = 3
	ContentInventoryTraitTypeAggregateRange          ContentInventoryTraitType = 4
	ContentInventoryTraitTypeResurrected             ContentInventoryTraitType = 5
	ContentInventoryTraitTypeMarathon                ContentInventoryTraitType = 6
	ContentInventoryTraitTypeNewRelease              ContentInventoryTraitType = 7
	ContentInventoryTraitTypeStreakDays              ContentInventoryTraitType = 8
	ContentInventoryTraitTypeTrendingContent         ContentInventoryTraitType = 9
	ContentInventoryTraitTypeTopItemTotalCount       ContentInventoryTraitType = 10
	ContentInventoryTraitTypeTopParentItemTotalCount ContentInventoryTraitType = 11
	ContentInventoryTraitTypeAggregateCount          ContentInventoryTraitType = 12
)

func (v ContentInventoryTraitType) String() string {
	switch v {
	case ContentInventoryTraitTypeTraitTypeUnspecified:
		return "TRAIT_TYPE_UNSPECIFIED"
	case ContentInventoryTraitTypeFirstTime:
		return "FIRST_TIME"
	case ContentInventoryTraitTypeDurationSeconds:
		return "DURATION_SECONDS"
	case ContentInventoryTraitTypeIsLive:
		return "IS_LIVE"
	case ContentInventoryTraitTypeAggregateRange:
		return "AGGREGATE_RANGE"
	case ContentInventoryTraitTypeResurrected:
		return "RESURRECTED"
	case ContentInventoryTraitTypeMarathon:
		return "MARATHON"
	case ContentInventoryTraitTypeNewRelease:
		return "NEW_RELEASE"
	case ContentInventoryTraitTypeStreakDays:
		return "STREAK_DAYS"
	case ContentInventoryTraitTypeTrendingContent:
		return "TRENDING_CONTENT"
	case ContentInventoryTraitTypeTopItemTotalCount:
		return "TOP_ITEM_TOTAL_COUNT"
	case ContentInventoryTraitTypeTopParentItemTotalCount:
		return "TOP_PARENT_ITEM_TOTAL_COUNT"
	case ContentInventoryTraitTypeAggregateCount:
		return "AGGREGATE_COUNT"
	}
	return strconv.FormatInt(int64(v), 10)
}

type ContentInventoryTrendingType int64

const (
	ContentInventoryTrendingTypeTrendingTypeUnspecified ContentInventoryTrendingType = 0
	ContentInventoryTrendingTypeGlobal                  ContentInventoryTrendingType = 1
)

func (v ContentInventoryTrendingType) String() string {
	switch v {
	case ContentInventoryTrendingTypeTrendingTypeUnspecified:
		return "TRENDING_TYPE_UNSPECIFIED"
	case ContentInventoryTrendingTypeGlobal:
		return "GLOBAL"
	}
	return strconv.FormatInt(int64(v), 10)
}

type ContentRatingAgency int64

const (
	ContentRatingAgencyEsrb ContentRatingAgency = 1
	ContentRatingAgencyPegi ContentRatingAgency = 2
)

func (v ContentRatingAgency) String() string {
	switch v {
	case ContentRatingAgencyEsrb:
		return "ESRB"
	case ContentRatingAgencyPegi:
		return "PEGI"
	}
	return strconv.FormatInt(int64(v), 10)
}

type ContentScanFlags uint64

const (
	ContentScanFlagsExplicit ContentScanFlags = 1 << 0
	ContentScanFlagsGore     ContentScanFlags = 1 << 1
)

func (v ContentScanFlags) String() string {
	switch v {
	case ContentScanFlagsExplicit:
		return "EXPLICIT"
	case ContentScanFlagsGore:
		return "GORE"
	}
	return strconv.FormatUint(uint64(v), 10)
}

type CreatorMonetizationMarketingOnboardingSectionType string

const (
	CreatorMonetizationMarketingOnboardingSectionTypeHowItWorks    CreatorMonetizationMarketingOnboardingSectionType = "HOW_IT_WORKS"
	CreatorMonetizationMarketingOnboardingSectionTypeBenefits      CreatorMonetizationMarketingOnboardingSectionType = "BENEFITS"
	CreatorMonetizationMarketingOnboardingSectionTypeOtherCreators CreatorMonetizationMarketingOnboardingSectionType = "OTHER_CREATORS"
)

type CreatorMonetizationRequestState string

const (
	CreatorMonetizationRequestStateOpen           CreatorMonetizationRequestState = "OPEN"
	CreatorMonetizationRequestStateRejected       CreatorMonetizationRequestState = "REJECTED"
	CreatorMonetizationRequestStateApproved       CreatorMonetizationRequestState = "APPROVED"
	CreatorMonetizationRequestStateActionRequired CreatorMonetizationRequestState = "ACTION_REQUIRED"
)

type CreatorMonetizationRestrictionType string

const (
	CreatorMonetizationRestrictionTypeNewPurchasesDisabled             CreatorMonetizationRestrictionType = "NEW_PURCHASES_DISABLED"
	CreatorMonetizationRestrictionTypeReapplicationDisabled            CreatorMonetizationRestrictionType = "REAPPLICATION_DISABLED"
	CreatorMonetizationRestrictionTypeSettingsReadOnly                 CreatorMonetizationRestrictionType = "SETTINGS_READ_ONLY"
	CreatorMonetizationRestrictionTypeSubscriptionsEndedFullRefund     CreatorMonetizationRestrictionType = "SUBSCRIPTIONS_ENDED_FULL_REFUND"
	CreatorMonetizationRestrictionTypeSubscriptionsEndedProratedRefund CreatorMonetizationRestrictionType = "SUBSCRIPTIONS_ENDED_PRORATED_REFUND"
)

type CustomStatusPushNotificationType int64

const (
	CustomStatusPushNotificationTypeStatusPushUnset    CustomStatusPushNotificationType = 0
	CustomStatusPushNotificationTypeStatusPushEnabled  CustomStatusPushNotificationType = 1
	CustomStatusPushNotificationTypeStatusPushDisabled CustomStatusPushNotificationType = 2
)

func (v CustomStatusPushNotificationType) String() string {
	switch v {
	case CustomStatusPushNotificationTypeStatusPushUnset:
		return "STATUS_PUSH_UNSET"
	case CustomStatusPushNotificationTypeStatusPushEnabled:
		return "STATUS_PUSH_ENABLED"
	case CustomStatusPushNotificationTypeStatusPushDisabled:
		return "STATUS_PUSH_DISABLED"
	}
	return strconv.FormatInt(int64(v), 10)
}

type DMSpamFilter int64

const (
	DMSpamFilterDisabled             DMSpamFilter = 0
	DMSpamFilterNonFriends           DMSpamFilter = 1
	DMSpamFilterFriendsAndNonFriends DMSpamFilter = 2
)

func (v DMSpamFilter) String() string {
	switch v {
	case DMSpamFilterDisabled:
		return "DISABLED"
	case DMSpamFilterNonFriends:
		return "NON_FRIENDS"
	case DMSpamFilterFriendsAndNonFriends:
		return "FRIENDS_AND_NON_FRIENDS"
	}
	return strconv.FormatInt(int64(v), 10)
}

type DMSpamFilterV2 int64

const (
	DMSpamFilterV2Unset                DMSpamFilterV2 = 0
	DMSpamFilterV2Disabled             DMSpamFilterV2 = 1
	DMSpamFilterV2NonFriends           DMSpamFilterV2 = 2
	DMSpamFilterV2FriendsAndNonFriends DMSpamFilterV2 = 3
)

func (v DMSpamFilterV2) String() string {
	switch v {
	case DMSpamFilterV2Unset:
		return "UNSET"
	case DMSpamFilterV2Disabled:
		return "DISABLED"
	case DMSpamFilterV2NonFriends:
		return "NON_FRIENDS"
	case DMSpamFilterV2FriendsAndNonFriends:
		return "FRIENDS_AND_NON_FRIENDS"
	}
	return strconv.FormatInt(int64(v), 10)
}

type DataMode string

const (
	DataModeFull        DataMode = "full"
	DataModePartial     DataMode = "partial"
	DataModeUnavailable DataMode = "unavailable"
)

type Debugging string

const (
	DebuggingCanary Debugging = "canary"
	DebuggingTrace  Debugging = "trace"
)

type DeepLinkType string

const (
	DeepLinkTypeUserSettings         DeepLinkType = "USER_SETTINGS"
	DeepLinkTypeChangelog            DeepLinkType = "CHANGELOG"
	DeepLinkTypeLibrary              DeepLinkType = "LIBRARY"
	DeepLinkTypeStoreHome            DeepLinkType = "STORE_HOME"
	DeepLinkTypeStoreListing         DeepLinkType = "STORE_LISTING"
	DeepLinkTypeChannel              DeepLinkType = "CHANNEL"
	DeepLinkTypeGameShop             DeepLinkType = "GAME_SHOP"
	DeepLinkTypePickGuildSettings    DeepLinkType = "PICK_GUILD_SETTINGS"
	DeepLinkTypeQuestHome            DeepLinkType = "QUEST_HOME"
	DeepLinkTypeDiscoveryGameResults DeepLinkType = "DISCOVERY_GAME_RESULTS"
	DeepLinkTypeOAuth2               DeepLinkType = "OAUTH2"
	DeepLinkTypeFeatures             DeepLinkType = "FEATURES"
	DeepLinkTypeShop                 DeepLinkType = "SHOP"
	DeepLinkTypeActivities           DeepLinkType = "ACTIVITIES"
	DeepLinkTypeQuestPreviewTool     DeepLinkType = "QUEST_PREVIEW_TOOL"
	DeepLinkTypeOneTimeLogin         DeepLinkType = "ONE_TIME_LOGIN"
)

type DesignID int64

const (
	DesignIDClassicIa    DesignID = 0
	DesignIDDesignIa     DesignID = 1
	DesignIDDesignTabsIa DesignID = 2
	DesignIDYouBarIa     DesignID = 3
)

func (v DesignID) String() string {
	switch v {
	case DesignIDClassicIa:
		return "CLASSIC_IA"
	case DesignIDDesignIa:
		return "DESIGN_IA"
	case DesignIDDesignTabsIa:
		return "DESIGN_TABS_IA"
	case DesignIDYouBarIa:
		return "YOU_BAR_IA"
	}
	return strconv.FormatInt(int64(v), 10)
}

type DesktopArchitectureType string

const (
	DesktopArchitectureTypeX86   DesktopArchitectureType = "x86"
	DesktopArchitectureTypeX64   DesktopArchitectureType = "x64"
	DesktopArchitectureTypeArm64 DesktopArchitectureType = "arm64"
)

type DesktopExecutableFormat string

const (
	DesktopExecutableFormatDeb       DesktopExecutableFormat = "deb"
	DesktopExecutableFormatTarGz     DesktopExecutableFormat = "tar.gz"
	DesktopExecutableFormatRpm       DesktopExecutableFormat = "rpm"
	DesktopExecutableFormatPkgTarZst DesktopExecutableFormat = "pkg.tar.zst"
)

type DesktopPlatformType string

const (
	DesktopPlatformTypeWin   DesktopPlatformType = "win"
	DesktopPlatformTypeOsx   DesktopPlatformType = "osx"
	DesktopPlatformTypeLinux DesktopPlatformType = "linux"
)

type DesktopReleaseChannel string

const (
	DesktopReleaseChannelStable      DesktopReleaseChannel = "stable"
	DesktopReleaseChannelPtb         DesktopReleaseChannel = "ptb"
	DesktopReleaseChannelCanary      DesktopReleaseChannel = "canary"
	DesktopReleaseChannelDevelopment DesktopReleaseChannel = "development"
)

type DirectoryCategory int64

const (
	DirectoryCategoryUncategorized DirectoryCategory = 0
	DirectoryCategorySchoolClub    DirectoryCategory = 1
	DirectoryCategoryClass         DirectoryCategory = 2
	DirectoryCategoryStudySocial   DirectoryCategory = 3
	DirectoryCategorySubjectMajor  DirectoryCategory = 4
	DirectoryCategoryMisc          DirectoryCategory = 5
)

func (v DirectoryCategory) String() string {
	switch v {
	case DirectoryCategoryUncategorized:
		return "UNCATEGORIZED"
	case DirectoryCategorySchoolClub:
		return "SCHOOL_CLUB"
	case DirectoryCategoryClass:
		return "CLASS"
	case DirectoryCategoryStudySocial:
		return "STUDY_SOCIAL"
	case DirectoryCategorySubjectMajor:
		return "SUBJECT_MAJOR"
	case DirectoryCategoryMisc:
		return "MISC"
	}
	return strconv.FormatInt(int64(v), 10)
}

type DirectoryEntryType int64

const (
	DirectoryEntryTypeGuild               DirectoryEntryType = 0
	DirectoryEntryTypeGuildScheduledEvent DirectoryEntryType = 1
)

func (v DirectoryEntryType) String() string {
	switch v {
	case DirectoryEntryTypeGuild:
		return "GUILD"
	case DirectoryEntryTypeGuildScheduledEvent:
		return "GUILD_SCHEDULED_EVENT"
	}
	return strconv.FormatInt(int64(v), 10)
}

type DismissibleQuestContentFlags uint64

const (
	DismissibleQuestContentFlagsGiftInventorySettingsBadge DismissibleQuestContentFlags = 1 << 0
	DismissibleQuestContentFlagsQuestBar                   DismissibleQuestContentFlags = 1 << 1
	DismissibleQuestContentFlagsActivityPanel              DismissibleQuestContentFlags = 1 << 2
	DismissibleQuestContentFlagsQuestLiveStream            DismissibleQuestContentFlags = 1 << 3
)

func (v DismissibleQuestContentFlags) String() string {
	switch v {
	case DismissibleQuestContentFlagsGiftInventorySettingsBadge:
		return "GIFT_INVENTORY_SETTINGS_BADGE"
	case DismissibleQuestContentFlagsQuestBar:
		return "QUEST_BAR"
	case DismissibleQuestContentFlagsActivityPanel:
		return "ACTIVITY_PANEL"
	case DismissibleQuestContentFlagsQuestLiveStream:
		return "QUEST_LIVE_STREAM"
	}
	return strconv.FormatUint(uint64(v), 10)
}

type DisplayNameEffect int64

const (
	DisplayNameEffectSolid    DisplayNameEffect = 1
	DisplayNameEffectGradient DisplayNameEffect = 2
	DisplayNameEffectNeon     DisplayNameEffect = 3
	DisplayNameEffectToon     DisplayNameEffect = 4
	DisplayNameEffectPop      DisplayNameEffect = 5
	DisplayNameEffectGlow     DisplayNameEffect = 6
	DisplayNameEffectPrism    DisplayNameEffect = 7
	DisplayNameEffectGummy    DisplayNameEffect = 8
	DisplayNameEffectTest1    DisplayNameEffect = 1001
	DisplayNameEffectTest2    DisplayNameEffect = 1002
	DisplayNameEffectTest3    DisplayNameEffect = 1003
	DisplayNameEffectTest4    DisplayNameEffect = 1004
)

func (v DisplayNameEffect) String() string {
	switch v {
	case DisplayNameEffectSolid:
		return "SOLID"
	case DisplayNameEffectGradient:
		return "GRADIENT"
	case DisplayNameEffectNeon:
		return "NEON"
	case DisplayNameEffectToon:
		return "TOON"
	case DisplayNameEffectPop:
		return "POP"
	case DisplayNameEffectGlow:
		return "GLOW"
	case DisplayNameEffectPrism:
		return "PRISM"
	case DisplayNameEffectGummy:
		return "GUMMY"
	case DisplayNameEffectTest1:
		return "TEST_1"
	case DisplayNameEffectTest2:
		return "TEST_2"
	case DisplayNameEffectTest3:
		return "TEST_3"
	case DisplayNameEffectTest4:
		return "TEST_4"
	}
	return strconv.FormatInt(int64(v), 10)
}

type DisplayNameFont int64

const (
	DisplayNameFontDefault      DisplayNameFont = 11
	DisplayNameFontBangers      DisplayNameFont = 1
	DisplayNameFontBioRhyme     DisplayNameFont = 2
	DisplayNameFontCherryBomb   DisplayNameFont = 3
	DisplayNameFontChicle       DisplayNameFont = 4
	DisplayNameFontCompagnon    DisplayNameFont = 5
	DisplayNameFontMuseoModerno DisplayNameFont = 6
	DisplayNameFontNeoCastel    DisplayNameFont = 7
	DisplayNameFontPixelify     DisplayNameFont = 8
	DisplayNameFontRibes        DisplayNameFont = 9
	DisplayNameFontSinistre     DisplayNameFont = 10
	DisplayNameFontZillaSlab    DisplayNameFont = 12
	DisplayNameFontPlaypenSans  DisplayNameFont = 13
	DisplayNameFontOrbitron     DisplayNameFont = 14
	DisplayNameFontNewRocker    DisplayNameFont = 15
	DisplayNameFontKalam        DisplayNameFont = 16
)

func (v DisplayNameFont) String() string {
	switch v {
	case DisplayNameFontDefault:
		return "DEFAULT"
	case DisplayNameFontBangers:
		return "BANGERS"
	case DisplayNameFontBioRhyme:
		return "BIO_RHYME"
	case DisplayNameFontCherryBomb:
		return "CHERRY_BOMB"
	case DisplayNameFontChicle:
		return "CHICLE"
	case DisplayNameFontCompagnon:
		return "COMPAGNON"
	case DisplayNameFontMuseoModerno:
		return "MUSEO_MODERNO"
	case DisplayNameFontNeoCastel:
		return "NEO_CASTEL"
	case DisplayNameFontPixelify:
		return "PIXELIFY"
	case DisplayNameFontRibes:
		return "RIBES"
	case DisplayNameFontSinistre:
		return "SINISTRE"
	case DisplayNameFontZillaSlab:
		return "ZILLA_SLAB"
	case DisplayNameFontPlaypenSans:
		return "PLAYPEN_SANS"
	case DisplayNameFontOrbitron:
		return "ORBITRON"
	case DisplayNameFontNewRocker:
		return "NEW_ROCKER"
	case DisplayNameFontKalam:
		return "KALAM"
	}
	return strconv.FormatInt(int64(v), 10)
}

type DistributorType string

const (
	DistributorTypeDiscord      DistributorType = "discord"
	DistributorTypeSteam        DistributorType = "steam"
	DistributorTypeTwitch       DistributorType = "twitch"
	DistributorTypeUplay        DistributorType = "uplay"
	DistributorTypeBattlenet    DistributorType = "battlenet"
	DistributorTypeOrigin       DistributorType = "origin"
	DistributorTypeGog          DistributorType = "gog"
	DistributorTypeEpic         DistributorType = "epic"
	DistributorTypeMicrosoft    DistributorType = "microsoft"
	DistributorTypeIgdb         DistributorType = "igdb"
	DistributorTypeGlyph        DistributorType = "glyph"
	DistributorTypeGooglePlay   DistributorType = "google_play"
	DistributorTypeNvidiaGdnApp DistributorType = "nvidia_gdn_app"
	DistributorTypeGop          DistributorType = "gop"
	DistributorTypeRoblox       DistributorType = "roblox"
	DistributorTypeGdco         DistributorType = "gdco"
	DistributorTypeXbox         DistributorType = "xbox"
	DistributorTypeXboxTitle    DistributorType = "xbox_title"
	DistributorTypeXboxGamePass DistributorType = "xbox_game_pass"
	DistributorTypePlaystation  DistributorType = "playstation"
	DistributorTypeOpencritic   DistributorType = "opencritic"
)

type EmailSettingsCategory string

const (
	EmailSettingsCategoryCommunication            EmailSettingsCategory = "communication"
	EmailSettingsCategorySocial                   EmailSettingsCategory = "social"
	EmailSettingsCategoryRecommendationsAndEvents EmailSettingsCategory = "recommendations_and_events"
	EmailSettingsCategoryTips                     EmailSettingsCategory = "tips"
	EmailSettingsCategoryUpdatesAndAnnouncements  EmailSettingsCategory = "updates_and_announcements"
	EmailSettingsCategoryFamilyCenterDigest       EmailSettingsCategory = "family_center_digest"
)

type EmbedFlags uint64

const (
	EmbedFlagsContainsExplicitMedia   EmbedFlags = 1 << 4
	EmbedFlagsContentInventoryEntry   EmbedFlags = 1 << 5
	EmbedFlagsContainsGoreContent     EmbedFlags = 1 << 6
	EmbedFlagsContainsSelfHarmContent EmbedFlags = 1 << 7
)

func (v EmbedFlags) String() string {
	switch v {
	case EmbedFlagsContainsExplicitMedia:
		return "CONTAINS_EXPLICIT_MEDIA"
	case EmbedFlagsContentInventoryEntry:
		return "CONTENT_INVENTORY_ENTRY"
	case EmbedFlagsContainsGoreContent:
		return "CONTAINS_GORE_CONTENT"
	case EmbedFlagsContainsSelfHarmContent:
		return "CONTAINS_SELF_HARM_CONTENT"
	}
	return strconv.FormatUint(uint64(v), 10)
}

type EmbeddedActivityLabelType int64

const (
	EmbeddedActivityLabelTypeNone    EmbeddedActivityLabelType = 0
	EmbeddedActivityLabelTypeNew     EmbeddedActivityLabelType = 1
	EmbeddedActivityLabelTypeUpdated EmbeddedActivityLabelType = 2
)

func (v EmbeddedActivityLabelType) String() string {
	switch v {
	case EmbeddedActivityLabelTypeNone:
		return "NONE"
	case EmbeddedActivityLabelTypeNew:
		return "NEW"
	case EmbeddedActivityLabelTypeUpdated:
		return "UPDATED"
	}
	return strconv.FormatInt(int64(v), 10)
}

type EmbeddedActivityLocationType string

const (
	EmbeddedActivityLocationTypeGc EmbeddedActivityLocationType = "gc"
	EmbeddedActivityLocationTypePC EmbeddedActivityLocationType = "pc"
)

type EmbeddedActivityOrientationLockStateType int64

const (
	EmbeddedActivityOrientationLockStateTypeUnlocked  EmbeddedActivityOrientationLockStateType = 1
	EmbeddedActivityOrientationLockStateTypePortrait  EmbeddedActivityOrientationLockStateType = 2
	EmbeddedActivityOrientationLockStateTypeLandscape EmbeddedActivityOrientationLockStateType = 3
)

func (v EmbeddedActivityOrientationLockStateType) String() string {
	switch v {
	case EmbeddedActivityOrientationLockStateTypeUnlocked:
		return "UNLOCKED"
	case EmbeddedActivityOrientationLockStateTypePortrait:
		return "PORTRAIT"
	case EmbeddedActivityOrientationLockStateTypeLandscape:
		return "LANDSCAPE"
	}
	return strconv.FormatInt(int64(v), 10)
}

type EmbeddedActivityPlatformType string

const (
	EmbeddedActivityPlatformTypeWeb     EmbeddedActivityPlatformType = "web"
	EmbeddedActivityPlatformTypeAndroid EmbeddedActivityPlatformType = "android"
	EmbeddedActivityPlatformTypeIos     EmbeddedActivityPlatformType = "ios"
)

type EmbeddedActivityReleasePhase string

const (
	EmbeddedActivityReleasePhaseInDevelopment      EmbeddedActivityReleasePhase = "in_development"
	EmbeddedActivityReleasePhaseActivitiesTeam     EmbeddedActivityReleasePhase = "activities_team"
	EmbeddedActivityReleasePhaseEmployeeRelease    EmbeddedActivityReleasePhase = "employee_release"
	EmbeddedActivityReleasePhaseSoftLaunch         EmbeddedActivityReleasePhase = "soft_launch"
	EmbeddedActivityReleasePhaseSoftLaunchMultiGeo EmbeddedActivityReleasePhase = "soft_launch_multi_geo"
	EmbeddedActivityReleasePhaseGlobalLaunch       EmbeddedActivityReleasePhase = "global_launch"
)

type EmbeddedActivitySurface string

const (
	EmbeddedActivitySurfaceVoiceLauncher EmbeddedActivitySurface = "voice_launcher"
	EmbeddedActivitySurfaceTextLauncher  EmbeddedActivitySurface = "text_launcher"
)

type EmojiPickerSection string

const (
	EmojiPickerSectionFavorites     EmojiPickerSection = "FAVORITES"
	EmojiPickerSectionTopGuildEmoji EmojiPickerSection = "TOP_GUILD_EMOJI"
	EmojiPickerSectionRecent        EmojiPickerSection = "RECENT"
	EmojiPickerSectionPeople        EmojiPickerSection = "people"
	EmojiPickerSectionNature        EmojiPickerSection = "nature"
	EmojiPickerSectionFood          EmojiPickerSection = "food"
	EmojiPickerSectionActivity      EmojiPickerSection = "activity"
	EmojiPickerSectionTravel        EmojiPickerSection = "travel"
	EmojiPickerSectionObjects       EmojiPickerSection = "objects"
	EmojiPickerSectionSymbols       EmojiPickerSection = "symbols"
	EmojiPickerSectionFlags         EmojiPickerSection = "flags"
)

type EmojiSourceType string

const (
	EmojiSourceTypeGuild       EmojiSourceType = "GUILD"
	EmojiSourceTypeApplication EmojiSourceType = "APPLICATION"
)

type EntitlementFulfillmentStatus int64

const (
	EntitlementFulfillmentStatusUnknown                   EntitlementFulfillmentStatus = 0
	EntitlementFulfillmentStatusFulfillmentNotNeeded      EntitlementFulfillmentStatus = 1
	EntitlementFulfillmentStatusFulfillmentNeeded         EntitlementFulfillmentStatus = 2
	EntitlementFulfillmentStatusFulfilled                 EntitlementFulfillmentStatus = 3
	EntitlementFulfillmentStatusFulfillmentFailed         EntitlementFulfillmentStatus = 4
	EntitlementFulfillmentStatusUnfulfillmentNeeded       EntitlementFulfillmentStatus = 5
	EntitlementFulfillmentStatusUnfulfilled               EntitlementFulfillmentStatus = 6
	EntitlementFulfillmentStatusUnfulfillmentFailed       EntitlementFulfillmentStatus = 7
	EntitlementFulfillmentStatusUnfulfillmentNeededManual EntitlementFulfillmentStatus = 8
)

func (v EntitlementFulfillmentStatus) String() string {
	switch v {
	case EntitlementFulfillmentStatusUnknown:
		return "UNKNOWN"
	case EntitlementFulfillmentStatusFulfillmentNotNeeded:
		return "FULFILLMENT_NOT_NEEDED"
	case EntitlementFulfillmentStatusFulfillmentNeeded:
		return "FULFILLMENT_NEEDED"
	case EntitlementFulfillmentStatusFulfilled:
		return "FULFILLED"
	case EntitlementFulfillmentStatusFulfillmentFailed:
		return "FULFILLMENT_FAILED"
	case EntitlementFulfillmentStatusUnfulfillmentNeeded:
		return "UNFULFILLMENT_NEEDED"
	case EntitlementFulfillmentStatusUnfulfilled:
		return "UNFULFILLED"
	case EntitlementFulfillmentStatusUnfulfillmentFailed:
		return "UNFULFILLMENT_FAILED"
	case EntitlementFulfillmentStatusUnfulfillmentNeededManual:
		return "UNFULFILLMENT_NEEDED_MANUAL"
	}
	return strconv.FormatInt(int64(v), 10)
}

type EntitlementOwnerType int64

const (
	EntitlementOwnerTypeGuild EntitlementOwnerType = 1
	EntitlementOwnerTypeUser  EntitlementOwnerType = 2
)

func (v EntitlementOwnerType) String() string {
	switch v {
	case EntitlementOwnerTypeGuild:
		return "GUILD"
	case EntitlementOwnerTypeUser:
		return "USER"
	}
	return strconv.FormatInt(int64(v), 10)
}

type EntitlementSourceType int64

const (
	EntitlementSourceTypeQuestReward               EntitlementSourceType = 1
	EntitlementSourceTypeDeveloperGift             EntitlementSourceType = 2
	EntitlementSourceTypeInvoice                   EntitlementSourceType = 3
	EntitlementSourceTypeReverseTrial              EntitlementSourceType = 4
	EntitlementSourceTypeUserGift                  EntitlementSourceType = 5
	EntitlementSourceTypeGuildPowerup              EntitlementSourceType = 6
	EntitlementSourceTypeHolidayPromotion          EntitlementSourceType = 7
	EntitlementSourceTypeFractionalPremiumGiveback EntitlementSourceType = 8
	EntitlementSourceTypeSubscription              EntitlementSourceType = 9
	EntitlementSourceTypeSubscriptionMember        EntitlementSourceType = 11
)

func (v EntitlementSourceType) String() string {
	switch v {
	case EntitlementSourceTypeQuestReward:
		return "QUEST_REWARD"
	case EntitlementSourceTypeDeveloperGift:
		return "DEVELOPER_GIFT"
	case EntitlementSourceTypeInvoice:
		return "INVOICE"
	case EntitlementSourceTypeReverseTrial:
		return "REVERSE_TRIAL"
	case EntitlementSourceTypeUserGift:
		return "USER_GIFT"
	case EntitlementSourceTypeGuildPowerup:
		return "GUILD_POWERUP"
	case EntitlementSourceTypeHolidayPromotion:
		return "HOLIDAY_PROMOTION"
	case EntitlementSourceTypeFractionalPremiumGiveback:
		return "FRACTIONAL_PREMIUM_GIVEBACK"
	case EntitlementSourceTypeSubscription:
		return "SUBSCRIPTION"
	case EntitlementSourceTypeSubscriptionMember:
		return "SUBSCRIPTION_MEMBER"
	}
	return strconv.FormatInt(int64(v), 10)
}

type EntitlementType int64

const (
	EntitlementTypePurchase                  EntitlementType = 1
	EntitlementTypePremiumSubscription       EntitlementType = 2
	EntitlementTypeDeveloperGift             EntitlementType = 3
	EntitlementTypeTestModePurchase          EntitlementType = 4
	EntitlementTypeFreePurchase              EntitlementType = 5
	EntitlementTypeUserGift                  EntitlementType = 6
	EntitlementTypePremiumPurchase           EntitlementType = 7
	EntitlementTypeApplicationSubscription   EntitlementType = 8
	EntitlementTypeFreeStaffPurchase         EntitlementType = 9
	EntitlementTypeQuestReward               EntitlementType = 10
	EntitlementTypeFractionalRedemption      EntitlementType = 11
	EntitlementTypeVirtualCurrencyRedemption EntitlementType = 12
	EntitlementTypeGuildPowerup              EntitlementType = 13
)

func (v EntitlementType) String() string {
	switch v {
	case EntitlementTypePurchase:
		return "PURCHASE"
	case EntitlementTypePremiumSubscription:
		return "PREMIUM_SUBSCRIPTION"
	case EntitlementTypeDeveloperGift:
		return "DEVELOPER_GIFT"
	case EntitlementTypeTestModePurchase:
		return "TEST_MODE_PURCHASE"
	case EntitlementTypeFreePurchase:
		return "FREE_PURCHASE"
	case EntitlementTypeUserGift:
		return "USER_GIFT"
	case EntitlementTypePremiumPurchase:
		return "PREMIUM_PURCHASE"
	case EntitlementTypeApplicationSubscription:
		return "APPLICATION_SUBSCRIPTION"
	case EntitlementTypeFreeStaffPurchase:
		return "FREE_STAFF_PURCHASE"
	case EntitlementTypeQuestReward:
		return "QUEST_REWARD"
	case EntitlementTypeFractionalRedemption:
		return "FRACTIONAL_REDEMPTION"
	case EntitlementTypeVirtualCurrencyRedemption:
		return "VIRTUAL_CURRENCY_REDEMPTION"
	case EntitlementTypeGuildPowerup:
		return "GUILD_POWERUP"
	}
	return strconv.FormatInt(int64(v), 10)
}

type EphemeralityReason int64

const (
	EphemeralityReasonNone                  EphemeralityReason = 0
	EphemeralityReasonFeatureLimited        EphemeralityReason = 1
	EphemeralityReasonGuildFeatureLimited   EphemeralityReason = 2
	EphemeralityReasonUserFeatureLimited    EphemeralityReason = 3
	EphemeralityReasonSlowmode              EphemeralityReason = 4
	EphemeralityReasonRateLimit             EphemeralityReason = 5
	EphemeralityReasonCannotMessageUser     EphemeralityReason = 6
	EphemeralityReasonUserVerificationLevel EphemeralityReason = 7
	EphemeralityReasonCannotUnarchiveThread EphemeralityReason = 8
	EphemeralityReasonCannotJoinThread      EphemeralityReason = 9
	EphemeralityReasonMissingPermissions    EphemeralityReason = 10
	EphemeralityReasonCannotSendAttachments EphemeralityReason = 11
	EphemeralityReasonCannotSendEmbeds      EphemeralityReason = 12
	EphemeralityReasonCannotSendStickers    EphemeralityReason = 13
	EphemeralityReasonAutomodBlocked        EphemeralityReason = 14
	EphemeralityReasonHarmfulLink           EphemeralityReason = 15
	EphemeralityReasonCannotUseCommand      EphemeralityReason = 16
	EphemeralityReasonBetaGuildSize         EphemeralityReason = 17
	EphemeralityReasonCannotUseExternalApps EphemeralityReason = 18
)

func (v EphemeralityReason) String() string {
	switch v {
	case EphemeralityReasonNone:
		return "NONE"
	case EphemeralityReasonFeatureLimited:
		return "FEATURE_LIMITED"
	case EphemeralityReasonGuildFeatureLimited:
		return "GUILD_FEATURE_LIMITED"
	case EphemeralityReasonUserFeatureLimited:
		return "USER_FEATURE_LIMITED"
	case EphemeralityReasonSlowmode:
		return "SLOWMODE"
	case EphemeralityReasonRateLimit:
		return "RATE_LIMIT"
	case EphemeralityReasonCannotMessageUser:
		return "CANNOT_MESSAGE_USER"
	case EphemeralityReasonUserVerificationLevel:
		return "USER_VERIFICATION_LEVEL"
	case EphemeralityReasonCannotUnarchiveThread:
		return "CANNOT_UNARCHIVE_THREAD"
	case EphemeralityReasonCannotJoinThread:
		return "CANNOT_JOIN_THREAD"
	case EphemeralityReasonMissingPermissions:
		return "MISSING_PERMISSIONS"
	case EphemeralityReasonCannotSendAttachments:
		return "CANNOT_SEND_ATTACHMENTS"
	case EphemeralityReasonCannotSendEmbeds:
		return "CANNOT_SEND_EMBEDS"
	case EphemeralityReasonCannotSendStickers:
		return "CANNOT_SEND_STICKERS"
	case EphemeralityReasonAutomodBlocked:
		return "AUTOMOD_BLOCKED"
	case EphemeralityReasonHarmfulLink:
		return "HARMFUL_LINK"
	case EphemeralityReasonCannotUseCommand:
		return "CANNOT_USE_COMMAND"
	case EphemeralityReasonBetaGuildSize:
		return "BETA_GUILD_SIZE"
	case EphemeralityReasonCannotUseExternalApps:
		return "CANNOT_USE_EXTERNAL_APPS"
	}
	return strconv.FormatInt(int64(v), 10)
}

type EsrbContentDescriptor int64

const (
	EsrbContentDescriptorAlcoholReference     EsrbContentDescriptor = 1
	EsrbContentDescriptorAnimatedBlood        EsrbContentDescriptor = 2
	EsrbContentDescriptorBlood                EsrbContentDescriptor = 3
	EsrbContentDescriptorBloodAndGore         EsrbContentDescriptor = 4
	EsrbContentDescriptorCartoonViolence      EsrbContentDescriptor = 5
	EsrbContentDescriptorComicMischief        EsrbContentDescriptor = 6
	EsrbContentDescriptorCrudeHumor           EsrbContentDescriptor = 7
	EsrbContentDescriptorDrugReference        EsrbContentDescriptor = 8
	EsrbContentDescriptorFantasyViolence      EsrbContentDescriptor = 9
	EsrbContentDescriptorIntenseViolence      EsrbContentDescriptor = 10
	EsrbContentDescriptorLanguage             EsrbContentDescriptor = 11
	EsrbContentDescriptorLyrics               EsrbContentDescriptor = 12
	EsrbContentDescriptorMatureHumor          EsrbContentDescriptor = 13
	EsrbContentDescriptorNudity               EsrbContentDescriptor = 14
	EsrbContentDescriptorPartialNudity        EsrbContentDescriptor = 15
	EsrbContentDescriptorRealGambling         EsrbContentDescriptor = 16
	EsrbContentDescriptorSexualContent        EsrbContentDescriptor = 17
	EsrbContentDescriptorSexualThemes         EsrbContentDescriptor = 18
	EsrbContentDescriptorSexualViolence       EsrbContentDescriptor = 19
	EsrbContentDescriptorSimulatedGambling    EsrbContentDescriptor = 20
	EsrbContentDescriptorStrongLanguage       EsrbContentDescriptor = 21
	EsrbContentDescriptorStrongLyrics         EsrbContentDescriptor = 22
	EsrbContentDescriptorStrongSexualContent  EsrbContentDescriptor = 23
	EsrbContentDescriptorSuggestiveThemes     EsrbContentDescriptor = 24
	EsrbContentDescriptorTobaccoReference     EsrbContentDescriptor = 25
	EsrbContentDescriptorUseOfAlcohol         EsrbContentDescriptor = 26
	EsrbContentDescriptorUseOfDrugs           EsrbContentDescriptor = 27
	EsrbContentDescriptorUseOfTobacco         EsrbContentDescriptor = 28
	EsrbContentDescriptorViolence             EsrbContentDescriptor = 29
	EsrbContentDescriptorViolentReferences    EsrbContentDescriptor = 30
	EsrbContentDescriptorInGamePurchases      EsrbContentDescriptor = 31
	EsrbContentDescriptorUsersInteract        EsrbContentDescriptor = 32
	EsrbContentDescriptorSharesLocation       EsrbContentDescriptor = 33
	EsrbContentDescriptorUnrestrictedInternet EsrbContentDescriptor = 34
	EsrbContentDescriptorMildBlood            EsrbContentDescriptor = 35
	EsrbContentDescriptorMildCartoonViolence  EsrbContentDescriptor = 36
	EsrbContentDescriptorMildFantasyViolence  EsrbContentDescriptor = 37
	EsrbContentDescriptorMildLanguage         EsrbContentDescriptor = 38
	EsrbContentDescriptorMildLyrics           EsrbContentDescriptor = 39
	EsrbContentDescriptorMildSexualThemes     EsrbContentDescriptor = 40
	EsrbContentDescriptorMildSuggestiveThemes EsrbContentDescriptor = 41
	EsrbContentDescriptorMildViolence         EsrbContentDescriptor = 42
	EsrbContentDescriptorAnimatedViolence     EsrbContentDescriptor = 43
)

func (v EsrbContentDescriptor) String() string {
	switch v {
	case EsrbContentDescriptorAlcoholReference:
		return "ALCOHOL_REFERENCE"
	case EsrbContentDescriptorAnimatedBlood:
		return "ANIMATED_BLOOD"
	case EsrbContentDescriptorBlood:
		return "BLOOD"
	case EsrbContentDescriptorBloodAndGore:
		return "BLOOD_AND_GORE"
	case EsrbContentDescriptorCartoonViolence:
		return "CARTOON_VIOLENCE"
	case EsrbContentDescriptorComicMischief:
		return "COMIC_MISCHIEF"
	case EsrbContentDescriptorCrudeHumor:
		return "CRUDE_HUMOR"
	case EsrbContentDescriptorDrugReference:
		return "DRUG_REFERENCE"
	case EsrbContentDescriptorFantasyViolence:
		return "FANTASY_VIOLENCE"
	case EsrbContentDescriptorIntenseViolence:
		return "INTENSE_VIOLENCE"
	case EsrbContentDescriptorLanguage:
		return "LANGUAGE"
	case EsrbContentDescriptorLyrics:
		return "LYRICS"
	case EsrbContentDescriptorMatureHumor:
		return "MATURE_HUMOR"
	case EsrbContentDescriptorNudity:
		return "NUDITY"
	case EsrbContentDescriptorPartialNudity:
		return "PARTIAL_NUDITY"
	case EsrbContentDescriptorRealGambling:
		return "REAL_GAMBLING"
	case EsrbContentDescriptorSexualContent:
		return "SEXUAL_CONTENT"
	case EsrbContentDescriptorSexualThemes:
		return "SEXUAL_THEMES"
	case EsrbContentDescriptorSexualViolence:
		return "SEXUAL_VIOLENCE"
	case EsrbContentDescriptorSimulatedGambling:
		return "SIMULATED_GAMBLING"
	case EsrbContentDescriptorStrongLanguage:
		return "STRONG_LANGUAGE"
	case EsrbContentDescriptorStrongLyrics:
		return "STRONG_LYRICS"
	case EsrbContentDescriptorStrongSexualContent:
		return "STRONG_SEXUAL_CONTENT"
	case EsrbContentDescriptorSuggestiveThemes:
		return "SUGGESTIVE_THEMES"
	case EsrbContentDescriptorTobaccoReference:
		return "TOBACCO_REFERENCE"
	case EsrbContentDescriptorUseOfAlcohol:
		return "USE_OF_ALCOHOL"
	case EsrbContentDescriptorUseOfDrugs:
		return "USE_OF_DRUGS"
	case EsrbContentDescriptorUseOfTobacco:
		return "USE_OF_TOBACCO"
	case EsrbContentDescriptorViolence:
		return "VIOLENCE"
	case EsrbContentDescriptorViolentReferences:
		return "VIOLENT_REFERENCES"
	case EsrbContentDescriptorInGamePurchases:
		return "IN_GAME_PURCHASES"
	case EsrbContentDescriptorUsersInteract:
		return "USERS_INTERACT"
	case EsrbContentDescriptorSharesLocation:
		return "SHARES_LOCATION"
	case EsrbContentDescriptorUnrestrictedInternet:
		return "UNRESTRICTED_INTERNET"
	case EsrbContentDescriptorMildBlood:
		return "MILD_BLOOD"
	case EsrbContentDescriptorMildCartoonViolence:
		return "MILD_CARTOON_VIOLENCE"
	case EsrbContentDescriptorMildFantasyViolence:
		return "MILD_FANTASY_VIOLENCE"
	case EsrbContentDescriptorMildLanguage:
		return "MILD_LANGUAGE"
	case EsrbContentDescriptorMildLyrics:
		return "MILD_LYRICS"
	case EsrbContentDescriptorMildSexualThemes:
		return "MILD_SEXUAL_THEMES"
	case EsrbContentDescriptorMildSuggestiveThemes:
		return "MILD_SUGGESTIVE_THEMES"
	case EsrbContentDescriptorMildViolence:
		return "MILD_VIOLENCE"
	case EsrbContentDescriptorAnimatedViolence:
		return "ANIMATED_VIOLENCE"
	}
	return strconv.FormatInt(int64(v), 10)
}

type EsrbContentRating int64

const (
	EsrbContentRatingEveryone        EsrbContentRating = 1
	EsrbContentRatingEveryoneTenPlus EsrbContentRating = 2
	EsrbContentRatingTeen            EsrbContentRating = 3
	EsrbContentRatingMature          EsrbContentRating = 4
	EsrbContentRatingAdultsOnly      EsrbContentRating = 5
	EsrbContentRatingRatingPending   EsrbContentRating = 6
)

func (v EsrbContentRating) String() string {
	switch v {
	case EsrbContentRatingEveryone:
		return "EVERYONE"
	case EsrbContentRatingEveryoneTenPlus:
		return "EVERYONE_TEN_PLUS"
	case EsrbContentRatingTeen:
		return "TEEN"
	case EsrbContentRatingMature:
		return "MATURE"
	case EsrbContentRatingAdultsOnly:
		return "ADULTS_ONLY"
	case EsrbContentRatingRatingPending:
		return "RATING_PENDING"
	}
	return strconv.FormatInt(int64(v), 10)
}

type EventWebhooksStatus int64

const (
	EventWebhooksStatusDisabled EventWebhooksStatus = 1
	EventWebhooksStatusEnabled  EventWebhooksStatus = 2
)

func (v EventWebhooksStatus) String() string {
	switch v {
	case EventWebhooksStatusDisabled:
		return "DISABLED"
	case EventWebhooksStatusEnabled:
		return "ENABLED"
	}
	return strconv.FormatInt(int64(v), 10)
}

type EventWebhooksType string

const (
	EventWebhooksTypeApplicationAuthorized   EventWebhooksType = "APPLICATION_AUTHORIZED"
	EventWebhooksTypeApplicationDeauthorized EventWebhooksType = "APPLICATION_DEAUTHORIZED"
	EventWebhooksTypeEntitlementCreate       EventWebhooksType = "ENTITLEMENT_CREATE"
	EventWebhooksTypeEntitlementUpdate       EventWebhooksType = "ENTITLEMENT_UPDATE"
	EventWebhooksTypeEntitlementDelete       EventWebhooksType = "ENTITLEMENT_DELETE"
	EventWebhooksTypeQuestUserEnrollment     EventWebhooksType = "QUEST_USER_ENROLLMENT"
	EventWebhooksTypeLobbyMessageCreate      EventWebhooksType = "LOBBY_MESSAGE_CREATE"
	EventWebhooksTypeLobbyMessageUpdate      EventWebhooksType = "LOBBY_MESSAGE_UPDATE"
	EventWebhooksTypeLobbyMessageDelete      EventWebhooksType = "LOBBY_MESSAGE_DELETE"
	EventWebhooksTypeGameDirectMessageCreate EventWebhooksType = "GAME_DIRECT_MESSAGE_CREATE"
	EventWebhooksTypeGameDirectMessageUpdate EventWebhooksType = "GAME_DIRECT_MESSAGE_UPDATE"
	EventWebhooksTypeGameDirectMessageDelete EventWebhooksType = "GAME_DIRECT_MESSAGE_DELETE"
)

type ExperimentPlatform string

const (
	ExperimentPlatformDeveloperPortal ExperimentPlatform = "DEVELOPER_PORTAL"
)

type ExplicitContentFilter int64

const (
	ExplicitContentFilterDisabled    ExplicitContentFilter = 0
	ExplicitContentFilterNonFriends  ExplicitContentFilter = 1
	ExplicitContentFilterAllMessages ExplicitContentFilter = 2
)

func (v ExplicitContentFilter) String() string {
	switch v {
	case ExplicitContentFilterDisabled:
		return "DISABLED"
	case ExplicitContentFilterNonFriends:
		return "NON_FRIENDS"
	case ExplicitContentFilterAllMessages:
		return "ALL_MESSAGES"
	}
	return strconv.FormatInt(int64(v), 10)
}

type ExplicitContentFilterLevel int64

const (
	ExplicitContentFilterLevelInherit ExplicitContentFilterLevel = 0
	ExplicitContentFilterLevelAlways  ExplicitContentFilterLevel = 1
)

func (v ExplicitContentFilterLevel) String() string {
	switch v {
	case ExplicitContentFilterLevelInherit:
		return "INHERIT"
	case ExplicitContentFilterLevelAlways:
		return "ALWAYS"
	}
	return strconv.FormatInt(int64(v), 10)
}

type ExplicitContentRedaction int64

const (
	ExplicitContentRedactionUnset ExplicitContentRedaction = 0
	ExplicitContentRedactionShow  ExplicitContentRedaction = 1
	ExplicitContentRedactionBlur  ExplicitContentRedaction = 2
	ExplicitContentRedactionBlock ExplicitContentRedaction = 3
)

func (v ExplicitContentRedaction) String() string {
	switch v {
	case ExplicitContentRedactionUnset:
		return "UNSET"
	case ExplicitContentRedactionShow:
		return "SHOW"
	case ExplicitContentRedactionBlur:
		return "BLUR"
	case ExplicitContentRedactionBlock:
		return "BLOCK"
	}
	return strconv.FormatInt(int64(v), 10)
}

type ExternalProviderAuthenticationType string

const (
	ExternalProviderAuthenticationTypeOidc                          ExternalProviderAuthenticationType = "OIDC"
	ExternalProviderAuthenticationTypeEpicOnlineServicesAccessToken ExternalProviderAuthenticationType = "EPIC_ONLINE_SERVICES_ACCESS_TOKEN"
	ExternalProviderAuthenticationTypeEpicOnlineServicesIDToken     ExternalProviderAuthenticationType = "EPIC_ONLINE_SERVICES_ID_TOKEN"
	ExternalProviderAuthenticationTypeSteamSessionTicket            ExternalProviderAuthenticationType = "STEAM_SESSION_TICKET"
	ExternalProviderAuthenticationTypeUnityServicesIDToken          ExternalProviderAuthenticationType = "UNITY_SERVICES_ID_TOKEN"
	ExternalProviderAuthenticationTypeDiscordBotIssuedAccessToken   ExternalProviderAuthenticationType = "DISCORD_BOT_ISSUED_ACCESS_TOKEN"
	ExternalProviderAuthenticationTypeAppleIDToken                  ExternalProviderAuthenticationType = "APPLE_ID_TOKEN"
	ExternalProviderAuthenticationTypePlaystationNetworkIDToken     ExternalProviderAuthenticationType = "PLAYSTATION_NETWORK_ID_TOKEN"
)

type ExternalSKUStrategyType int64

const (
	ExternalSKUStrategyTypeConstant     ExternalSKUStrategyType = 1
	ExternalSKUStrategyTypeAppleSticker ExternalSKUStrategyType = 2
)

func (v ExternalSKUStrategyType) String() string {
	switch v {
	case ExternalSKUStrategyTypeConstant:
		return "CONSTANT"
	case ExternalSKUStrategyTypeAppleSticker:
		return "APPLE_STICKER"
	}
	return strconv.FormatInt(int64(v), 10)
}

type FavoriteChannelType string

const (
	FavoriteChannelTypeUnset             FavoriteChannelType = "UNSET"
	FavoriteChannelTypeReferenceOriginal FavoriteChannelType = "REFERENCE_ORIGINAL"
	FavoriteChannelTypeCategory          FavoriteChannelType = "CATEGORY"
)

type FileTypes string

const (
	FileTypesImage         FileTypes = "image"
	FileTypesVideo         FileTypes = "video"
	FileTypesAudio         FileTypes = "audio"
	FileTypesFileExtension FileTypes = ".{file_extension}"
)

type ForLaterTab int64

const (
	ForLaterTabUnspecified ForLaterTab = 0
	ForLaterTabAll         ForLaterTab = 1
	ForLaterTabBookmarks   ForLaterTab = 2
	ForLaterTabReminders   ForLaterTab = 3
)

func (v ForLaterTab) String() string {
	switch v {
	case ForLaterTabUnspecified:
		return "UNSPECIFIED"
	case ForLaterTabAll:
		return "ALL"
	case ForLaterTabBookmarks:
		return "BOOKMARKS"
	case ForLaterTabReminders:
		return "REMINDERS"
	}
	return strconv.FormatInt(int64(v), 10)
}

type ForumLayoutType int64

const (
	ForumLayoutTypeDefault ForumLayoutType = 0
	ForumLayoutTypeList    ForumLayoutType = 1
	ForumLayoutTypeGrid    ForumLayoutType = 2
)

func (v ForumLayoutType) String() string {
	switch v {
	case ForumLayoutTypeDefault:
		return "DEFAULT"
	case ForumLayoutTypeList:
		return "LIST"
	case ForumLayoutTypeGrid:
		return "GRID"
	}
	return strconv.FormatInt(int64(v), 10)
}

type FriendDiscoveryFlags uint64

const (
	FriendDiscoveryFlagsFindByPhone FriendDiscoveryFlags = 1 << 1
	FriendDiscoveryFlagsFindByEmail FriendDiscoveryFlags = 1 << 2
)

func (v FriendDiscoveryFlags) String() string {
	switch v {
	case FriendDiscoveryFlagsFindByPhone:
		return "FIND_BY_PHONE"
	case FriendDiscoveryFlagsFindByEmail:
		return "FIND_BY_EMAIL"
	}
	return strconv.FormatUint(uint64(v), 10)
}

type FriendSuggestionReasonType int64

const (
	FriendSuggestionReasonTypeExternalFriend FriendSuggestionReasonType = 1
)

func (v FriendSuggestionReasonType) String() string {
	switch v {
	case FriendSuggestionReasonTypeExternalFriend:
		return "EXTERNAL_FRIEND"
	}
	return strconv.FormatInt(int64(v), 10)
}

type GIFMediaFormat string

const (
	GIFMediaFormatMp4       GIFMediaFormat = "mp4"
	GIFMediaFormatTinymp4   GIFMediaFormat = "tinymp4"
	GIFMediaFormatNanomp4   GIFMediaFormat = "nanomp4"
	GIFMediaFormatLoopedmp4 GIFMediaFormat = "loopedmp4"
	GIFMediaFormatWebm      GIFMediaFormat = "webm"
	GIFMediaFormatTinywebm  GIFMediaFormat = "tinywebm"
	GIFMediaFormatNanowebm  GIFMediaFormat = "nanowebm"
	GIFMediaFormatGIF       GIFMediaFormat = "gif"
	GIFMediaFormatMediumgif GIFMediaFormat = "mediumgif"
	GIFMediaFormatTinygif   GIFMediaFormat = "tinygif"
	GIFMediaFormatNanogif   GIFMediaFormat = "nanogif"
	GIFMediaFormatWebp      GIFMediaFormat = "webp"
	GIFMediaFormatTinywebp  GIFMediaFormat = "tinywebp"
	GIFMediaFormatNanowebp  GIFMediaFormat = "nanowebp"
)

type GIFType int64

const (
	GIFTypeNone  GIFType = 0
	GIFTypeImage GIFType = 1
	GIFTypeVideo GIFType = 2
)

func (v GIFType) String() string {
	switch v {
	case GIFTypeNone:
		return "NONE"
	case GIFTypeImage:
		return "IMAGE"
	case GIFTypeVideo:
		return "VIDEO"
	}
	return strconv.FormatInt(int64(v), 10)
}

type GameActivityNotificationType int64

const (
	GameActivityNotificationTypeActivityNotificationsUnset    GameActivityNotificationType = 0
	GameActivityNotificationTypeActivityNotificationsDisabled GameActivityNotificationType = 1
	GameActivityNotificationTypeActivityNotificationsEnabled  GameActivityNotificationType = 2
	GameActivityNotificationTypeOnlyGamesPlayed               GameActivityNotificationType = 3
)

func (v GameActivityNotificationType) String() string {
	switch v {
	case GameActivityNotificationTypeActivityNotificationsUnset:
		return "ACTIVITY_NOTIFICATIONS_UNSET"
	case GameActivityNotificationTypeActivityNotificationsDisabled:
		return "ACTIVITY_NOTIFICATIONS_DISABLED"
	case GameActivityNotificationTypeActivityNotificationsEnabled:
		return "ACTIVITY_NOTIFICATIONS_ENABLED"
	case GameActivityNotificationTypeOnlyGamesPlayed:
		return "ONLY_GAMES_PLAYED"
	}
	return strconv.FormatInt(int64(v), 10)
}

type GameClaimStatus string

const (
	GameClaimStatusNotStarted                                 GameClaimStatus = "not_started"
	GameClaimStatusReadyToProcess                             GameClaimStatus = "ready_to_process"
	GameClaimStatusFetchingOfficialGuild                      GameClaimStatus = "fetching_official_guild"
	GameClaimStatusHasFetchedOfficialGuild                    GameClaimStatus = "has_fetched_official_guild"
	GameClaimStatusAwaitingVerificationCode                   GameClaimStatus = "awaiting_verification_code"
	GameClaimStatusVerificationPassed                         GameClaimStatus = "verification_passed"
	GameClaimStatusClaimApproved                              GameClaimStatus = "claim_approved"
	GameClaimStatusClaimRejected                              GameClaimStatus = "claim_rejected"
	GameClaimStatusClaimRevoked                               GameClaimStatus = "claim_revoked"
	GameClaimStatusFailedNoGame                               GameClaimStatus = "failed_no_game"
	GameClaimStatusFailedNoApplication                        GameClaimStatus = "failed_no_application"
	GameClaimStatusFailedNoTeam                               GameClaimStatus = "failed_no_team"
	GameClaimStatusFailedToFetchOfficialGuildFromThirdParties GameClaimStatus = "failed_to_fetch_official_guild_from_third_parties"
	GameClaimStatusFailedNoOfficialGuildFound                 GameClaimStatus = "failed_no_official_guild_found"
	GameClaimStatusFailedMultipleOfficialGuildsFound          GameClaimStatus = "failed_multiple_official_guilds_found"
	GameClaimStatusFailedInviteProofMissing                   GameClaimStatus = "failed_invite_proof_missing"
	GameClaimStatusFailedInviteProofInvalid                   GameClaimStatus = "failed_invite_proof_invalid"
	GameClaimStatusFailedOfficialGuildOwnerNotFound           GameClaimStatus = "failed_official_guild_owner_not_found"
	GameClaimStatusFailedOfficialGuildOwnerNotOnTeam          GameClaimStatus = "failed_official_guild_owner_not_on_team"
	GameClaimStatusFailedOfficialGuildOwnerMFANotEnabled      GameClaimStatus = "failed_official_guild_owner_mfa_not_enabled"
	GameClaimStatusFailedVerificationCodeSentTooRecently      GameClaimStatus = "failed_verification_code_sent_too_recently"
	GameClaimStatusFailedVerificationCodeNotSent              GameClaimStatus = "failed_verification_code_not_sent"
	GameClaimStatusFailedVerificationCodeInvalid              GameClaimStatus = "failed_verification_code_invalid"
	GameClaimStatusFailedUnknown                              GameClaimStatus = "failed_unknown"
)

type GameFlags uint64

const (
	GameFlagsGameProfileDisabled GameFlags = 1 << 0
)

func (v GameFlags) String() string {
	switch v {
	case GameFlagsGameProfileDisabled:
		return "GAME_PROFILE_DISABLED"
	}
	return strconv.FormatUint(uint64(v), 10)
}

type GamePlatformType int64

const (
	GamePlatformTypeDesktop     GamePlatformType = 0
	GamePlatformTypeXbox        GamePlatformType = 1
	GamePlatformTypePlaystation GamePlatformType = 2
	GamePlatformTypeIos         GamePlatformType = 3
	GamePlatformTypeAndroid     GamePlatformType = 4
	GamePlatformTypeNintendo    GamePlatformType = 5
	GamePlatformTypeLinux       GamePlatformType = 6
	GamePlatformTypeMacos       GamePlatformType = 7
)

func (v GamePlatformType) String() string {
	switch v {
	case GamePlatformTypeDesktop:
		return "DESKTOP"
	case GamePlatformTypeXbox:
		return "XBOX"
	case GamePlatformTypePlaystation:
		return "PLAYSTATION"
	case GamePlatformTypeIos:
		return "IOS"
	case GamePlatformTypeAndroid:
		return "ANDROID"
	case GamePlatformTypeNintendo:
		return "NINTENDO"
	case GamePlatformTypeLinux:
		return "LINUX"
	case GamePlatformTypeMacos:
		return "MACOS"
	}
	return strconv.FormatInt(int64(v), 10)
}

type GameServerProviderType string

const (
	GameServerProviderTypeShockbyte GameServerProviderType = "shockbyte"
)

type GameServerStatus string

const (
	GameServerStatusStarting        GameServerStatus = "starting"
	GameServerStatusStartupFailed   GameServerStatus = "startup_failed"
	GameServerStatusMissingStock    GameServerStatus = "missing_stock"
	GameServerStatusSleeping        GameServerStatus = "sleeping"
	GameServerStatusOffline         GameServerStatus = "offline"
	GameServerStatusOnline          GameServerStatus = "online"
	GameServerStatusDeleted         GameServerStatus = "deleted"
	GameServerStatusProviderErrored GameServerStatus = "provider_errored"
)

type GameThemeType int64

const (
	GameThemeTypeThriller       GameThemeType = 0
	GameThemeTypeScienceFiction GameThemeType = 1
	GameThemeTypeAction         GameThemeType = 2
	GameThemeTypeHorror         GameThemeType = 3
	GameThemeTypeSurvival       GameThemeType = 4
	GameThemeTypeFantasy        GameThemeType = 5
	GameThemeTypeHistorical     GameThemeType = 6
	GameThemeTypeStealth        GameThemeType = 7
	GameThemeTypeComedy         GameThemeType = 8
	GameThemeTypeBusiness       GameThemeType = 9
	GameThemeTypeDrama          GameThemeType = 10
	GameThemeTypeNonFiction     GameThemeType = 11
	GameThemeTypeKids           GameThemeType = 12
	GameThemeTypeSandbox        GameThemeType = 13
	GameThemeTypeOpenWorld      GameThemeType = 14
	GameThemeTypeWarfare        GameThemeType = 15
	GameThemeTypeEducational    GameThemeType = 16
	GameThemeTypeMystery        GameThemeType = 17
	GameThemeTypeParty          GameThemeType = 18
	GameThemeTypeRomance        GameThemeType = 19
	GameThemeTypeErotic         GameThemeType = 20
)

func (v GameThemeType) String() string {
	switch v {
	case GameThemeTypeThriller:
		return "THRILLER"
	case GameThemeTypeScienceFiction:
		return "SCIENCE_FICTION"
	case GameThemeTypeAction:
		return "ACTION"
	case GameThemeTypeHorror:
		return "HORROR"
	case GameThemeTypeSurvival:
		return "SURVIVAL"
	case GameThemeTypeFantasy:
		return "FANTASY"
	case GameThemeTypeHistorical:
		return "HISTORICAL"
	case GameThemeTypeStealth:
		return "STEALTH"
	case GameThemeTypeComedy:
		return "COMEDY"
	case GameThemeTypeBusiness:
		return "BUSINESS"
	case GameThemeTypeDrama:
		return "DRAMA"
	case GameThemeTypeNonFiction:
		return "NON_FICTION"
	case GameThemeTypeKids:
		return "KIDS"
	case GameThemeTypeSandbox:
		return "SANDBOX"
	case GameThemeTypeOpenWorld:
		return "OPEN_WORLD"
	case GameThemeTypeWarfare:
		return "WARFARE"
	case GameThemeTypeEducational:
		return "EDUCATIONAL"
	case GameThemeTypeMystery:
		return "MYSTERY"
	case GameThemeTypeParty:
		return "PARTY"
	case GameThemeTypeRomance:
		return "ROMANCE"
	case GameThemeTypeErotic:
		return "EROTIC"
	}
	return strconv.FormatInt(int64(v), 10)
}

type GameWebsiteCategory int64

const (
	GameWebsiteCategoryOfficial  GameWebsiteCategory = 1
	GameWebsiteCategoryWikia     GameWebsiteCategory = 2
	GameWebsiteCategoryWikipedia GameWebsiteCategory = 3
	GameWebsiteCategoryFacebook  GameWebsiteCategory = 4
	GameWebsiteCategoryTwitter   GameWebsiteCategory = 5
	GameWebsiteCategoryTwitch    GameWebsiteCategory = 6
	GameWebsiteCategoryInstagram GameWebsiteCategory = 8
	GameWebsiteCategoryYoutube   GameWebsiteCategory = 9
	GameWebsiteCategoryIphone    GameWebsiteCategory = 10
	GameWebsiteCategoryIpad      GameWebsiteCategory = 11
	GameWebsiteCategoryAndroid   GameWebsiteCategory = 12
	GameWebsiteCategorySteam     GameWebsiteCategory = 13
	GameWebsiteCategoryReddit    GameWebsiteCategory = 14
	GameWebsiteCategoryItch      GameWebsiteCategory = 15
	GameWebsiteCategoryEpicgames GameWebsiteCategory = 16
	GameWebsiteCategoryGog       GameWebsiteCategory = 17
	GameWebsiteCategoryDiscord   GameWebsiteCategory = 18
	GameWebsiteCategoryBluesky   GameWebsiteCategory = 19
	GameWebsiteCategoryBattlenet GameWebsiteCategory = 20
	GameWebsiteCategoryRiot      GameWebsiteCategory = 21
	GameWebsiteCategoryRoblox    GameWebsiteCategory = 22
	GameWebsiteCategoryMinecraft GameWebsiteCategory = 23
)

func (v GameWebsiteCategory) String() string {
	switch v {
	case GameWebsiteCategoryOfficial:
		return "OFFICIAL"
	case GameWebsiteCategoryWikia:
		return "WIKIA"
	case GameWebsiteCategoryWikipedia:
		return "WIKIPEDIA"
	case GameWebsiteCategoryFacebook:
		return "FACEBOOK"
	case GameWebsiteCategoryTwitter:
		return "TWITTER"
	case GameWebsiteCategoryTwitch:
		return "TWITCH"
	case GameWebsiteCategoryInstagram:
		return "INSTAGRAM"
	case GameWebsiteCategoryYoutube:
		return "YOUTUBE"
	case GameWebsiteCategoryIphone:
		return "IPHONE"
	case GameWebsiteCategoryIpad:
		return "IPAD"
	case GameWebsiteCategoryAndroid:
		return "ANDROID"
	case GameWebsiteCategorySteam:
		return "STEAM"
	case GameWebsiteCategoryReddit:
		return "REDDIT"
	case GameWebsiteCategoryItch:
		return "ITCH"
	case GameWebsiteCategoryEpicgames:
		return "EPICGAMES"
	case GameWebsiteCategoryGog:
		return "GOG"
	case GameWebsiteCategoryDiscord:
		return "DISCORD"
	case GameWebsiteCategoryBluesky:
		return "BLUESKY"
	case GameWebsiteCategoryBattlenet:
		return "BATTLENET"
	case GameWebsiteCategoryRiot:
		return "RIOT"
	case GameWebsiteCategoryRoblox:
		return "ROBLOX"
	case GameWebsiteCategoryMinecraft:
		return "MINECRAFT"
	}
	return strconv.FormatInt(int64(v), 10)
}

type GameWidgetTag string

const (
	GameWidgetTagNoob             GameWidgetTag = "noob"
	GameWidgetTagLearningTheRopes GameWidgetTag = "learning_the_ropes"
	GameWidgetTagCasual           GameWidgetTag = "casual"
	GameWidgetTagGettingGood      GameWidgetTag = "getting_good"
	GameWidgetTagIntermediate     GameWidgetTag = "intermediate"
	GameWidgetTagExpert           GameWidgetTag = "expert"
	GameWidgetTagBetterThanYou    GameWidgetTag = "better_than_you"
	GameWidgetTagObsessed         GameWidgetTag = "obsessed"
	GameWidgetTagLoveIt           GameWidgetTag = "love_it"
	GameWidgetTagKindOfLoveIt     GameWidgetTag = "kind_of_love_it"
	GameWidgetTagKindOfHateIt     GameWidgetTag = "kind_of_hate_it"
	GameWidgetTagRageQuitting     GameWidgetTag = "rage_quitting"
	GameWidgetTagLikeIt           GameWidgetTag = "like_it"
	GameWidgetTagFrustrated       GameWidgetTag = "frustrated"
	GameWidgetTagTooEasy          GameWidgetTag = "too_easy"
	GameWidgetTagLookingForGroup  GameWidgetTag = "looking_for_group"
	GameWidgetTagOpenToPlay       GameWidgetTag = "open_to_play"
	GameWidgetTagLookingForTips   GameWidgetTag = "looking_for_tips"
	GameWidgetTagOpenToTeach      GameWidgetTag = "open_to_teach"
	GameWidgetTagLookingToDiscuss GameWidgetTag = "looking_to_discuss"
)

type GameWidgetType string

const (
	GameWidgetTypeFavoriteGames   GameWidgetType = "favorite_games"
	GameWidgetTypePlayedGames     GameWidgetType = "played_games"
	GameWidgetTypeCurrentGames    GameWidgetType = "current_games"
	GameWidgetTypeWantToPlayGames GameWidgetType = "want_to_play_games"
	GameWidgetTypeApplication     GameWidgetType = "application"
)

type GiftCodeFlags uint64

const (
	GiftCodeFlagsPaymentSourceRequired          GiftCodeFlags = 1 << 0
	GiftCodeFlagsExistingSubscriptionDisallowed GiftCodeFlags = 1 << 1
	GiftCodeFlagsNotSelfRedeemable              GiftCodeFlags = 1 << 2
	GiftCodeFlagsPromotion                      GiftCodeFlags = 1 << 3
)

func (v GiftCodeFlags) String() string {
	switch v {
	case GiftCodeFlagsPaymentSourceRequired:
		return "PAYMENT_SOURCE_REQUIRED"
	case GiftCodeFlagsExistingSubscriptionDisallowed:
		return "EXISTING_SUBSCRIPTION_DISALLOWED"
	case GiftCodeFlagsNotSelfRedeemable:
		return "NOT_SELF_REDEEMABLE"
	case GiftCodeFlagsPromotion:
		return "PROMOTION"
	}
	return strconv.FormatUint(uint64(v), 10)
}

type GiftStyle int64

const (
	GiftStyleSnowglobe           GiftStyle = 1
	GiftStyleBox                 GiftStyle = 2
	GiftStyleCup                 GiftStyle = 3
	GiftStyleStandardBox         GiftStyle = 4
	GiftStyleCake                GiftStyle = 5
	GiftStyleChest               GiftStyle = 6
	GiftStyleCoffee              GiftStyle = 7
	GiftStyleSeasonalStandardBox GiftStyle = 8
	GiftStyleSeasonalCake        GiftStyle = 9
	GiftStyleSeasonalChest       GiftStyle = 10
	GiftStyleSeasonalCoffee      GiftStyle = 11
	GiftStyleNitroweenStandard   GiftStyle = 12
)

func (v GiftStyle) String() string {
	switch v {
	case GiftStyleSnowglobe:
		return "SNOWGLOBE"
	case GiftStyleBox:
		return "BOX"
	case GiftStyleCup:
		return "CUP"
	case GiftStyleStandardBox:
		return "STANDARD_BOX"
	case GiftStyleCake:
		return "CAKE"
	case GiftStyleChest:
		return "CHEST"
	case GiftStyleCoffee:
		return "COFFEE"
	case GiftStyleSeasonalStandardBox:
		return "SEASONAL_STANDARD_BOX"
	case GiftStyleSeasonalCake:
		return "SEASONAL_CAKE"
	case GiftStyleSeasonalChest:
		return "SEASONAL_CHEST"
	case GiftStyleSeasonalCoffee:
		return "SEASONAL_COFFEE"
	case GiftStyleNitroweenStandard:
		return "NITROWEEN_STANDARD"
	}
	return strconv.FormatInt(int64(v), 10)
}

type GuildActivityStatusRestrictionDefault int64

const (
	GuildActivityStatusRestrictionDefaultOff              GuildActivityStatusRestrictionDefault = 0
	GuildActivityStatusRestrictionDefaultOnForLargeGuilds GuildActivityStatusRestrictionDefault = 1
	GuildActivityStatusRestrictionDefaultOn               GuildActivityStatusRestrictionDefault = 2
)

func (v GuildActivityStatusRestrictionDefault) String() string {
	switch v {
	case GuildActivityStatusRestrictionDefaultOff:
		return "OFF"
	case GuildActivityStatusRestrictionDefaultOnForLargeGuilds:
		return "ON_FOR_LARGE_GUILDS"
	case GuildActivityStatusRestrictionDefaultOn:
		return "ON"
	}
	return strconv.FormatInt(int64(v), 10)
}

type GuildActivityStatusRestrictionDefaultV2 int64

const (
	GuildActivityStatusRestrictionDefaultV2ActivityStatusUnset            GuildActivityStatusRestrictionDefaultV2 = 0
	GuildActivityStatusRestrictionDefaultV2ActivityStatusOff              GuildActivityStatusRestrictionDefaultV2 = 1
	GuildActivityStatusRestrictionDefaultV2ActivityStatusOnForLargeGuilds GuildActivityStatusRestrictionDefaultV2 = 2
	GuildActivityStatusRestrictionDefaultV2ActivityStatusOn               GuildActivityStatusRestrictionDefaultV2 = 3
)

func (v GuildActivityStatusRestrictionDefaultV2) String() string {
	switch v {
	case GuildActivityStatusRestrictionDefaultV2ActivityStatusUnset:
		return "ACTIVITY_STATUS_UNSET"
	case GuildActivityStatusRestrictionDefaultV2ActivityStatusOff:
		return "ACTIVITY_STATUS_OFF"
	case GuildActivityStatusRestrictionDefaultV2ActivityStatusOnForLargeGuilds:
		return "ACTIVITY_STATUS_ON_FOR_LARGE_GUILDS"
	case GuildActivityStatusRestrictionDefaultV2ActivityStatusOn:
		return "ACTIVITY_STATUS_ON"
	}
	return strconv.FormatInt(int64(v), 10)
}

type GuildBadgeType int64

const (
	GuildBadgeTypeSword        GuildBadgeType = 0
	GuildBadgeTypeWaterDrop    GuildBadgeType = 1
	GuildBadgeTypeSkull        GuildBadgeType = 2
	GuildBadgeTypeToadstool    GuildBadgeType = 3
	GuildBadgeTypeMoon         GuildBadgeType = 4
	GuildBadgeTypeLightning    GuildBadgeType = 5
	GuildBadgeTypeLeaf         GuildBadgeType = 6
	GuildBadgeTypeHeart        GuildBadgeType = 7
	GuildBadgeTypeFire         GuildBadgeType = 8
	GuildBadgeTypeCompass      GuildBadgeType = 9
	GuildBadgeTypeCrosshairs   GuildBadgeType = 10
	GuildBadgeTypeFlower       GuildBadgeType = 11
	GuildBadgeTypeForce        GuildBadgeType = 12
	GuildBadgeTypeGem          GuildBadgeType = 13
	GuildBadgeTypeLava         GuildBadgeType = 14
	GuildBadgeTypePsychic      GuildBadgeType = 15
	GuildBadgeTypeSmoke        GuildBadgeType = 16
	GuildBadgeTypeSnow         GuildBadgeType = 17
	GuildBadgeTypeSound        GuildBadgeType = 18
	GuildBadgeTypeSun          GuildBadgeType = 19
	GuildBadgeTypeWind         GuildBadgeType = 20
	GuildBadgeTypeBunny        GuildBadgeType = 21
	GuildBadgeTypeDog          GuildBadgeType = 22
	GuildBadgeTypeFrog         GuildBadgeType = 23
	GuildBadgeTypeGoat         GuildBadgeType = 24
	GuildBadgeTypeCat          GuildBadgeType = 25
	GuildBadgeTypeDiamond      GuildBadgeType = 26
	GuildBadgeTypeCrown        GuildBadgeType = 27
	GuildBadgeTypeTrophy       GuildBadgeType = 28
	GuildBadgeTypeMoneyBag     GuildBadgeType = 29
	GuildBadgeTypeDollarSign   GuildBadgeType = 30
	GuildBadgeTypeClover       GuildBadgeType = 31
	GuildBadgeTypeBlossom      GuildBadgeType = 32
	GuildBadgeTypePottedPlant  GuildBadgeType = 33
	GuildBadgeTypeMaple        GuildBadgeType = 34
	GuildBadgeTypeWiltedFlower GuildBadgeType = 35
	GuildBadgeTypeButterfly    GuildBadgeType = 36
	GuildBadgeTypeSnail        GuildBadgeType = 37
	GuildBadgeTypeCaterpillar  GuildBadgeType = 38
	GuildBadgeTypeSpider       GuildBadgeType = 39
	GuildBadgeTypeBee          GuildBadgeType = 40
)

func (v GuildBadgeType) String() string {
	switch v {
	case GuildBadgeTypeSword:
		return "SWORD"
	case GuildBadgeTypeWaterDrop:
		return "WATER_DROP"
	case GuildBadgeTypeSkull:
		return "SKULL"
	case GuildBadgeTypeToadstool:
		return "TOADSTOOL"
	case GuildBadgeTypeMoon:
		return "MOON"
	case GuildBadgeTypeLightning:
		return "LIGHTNING"
	case GuildBadgeTypeLeaf:
		return "LEAF"
	case GuildBadgeTypeHeart:
		return "HEART"
	case GuildBadgeTypeFire:
		return "FIRE"
	case GuildBadgeTypeCompass:
		return "COMPASS"
	case GuildBadgeTypeCrosshairs:
		return "CROSSHAIRS"
	case GuildBadgeTypeFlower:
		return "FLOWER"
	case GuildBadgeTypeForce:
		return "FORCE"
	case GuildBadgeTypeGem:
		return "GEM"
	case GuildBadgeTypeLava:
		return "LAVA"
	case GuildBadgeTypePsychic:
		return "PSYCHIC"
	case GuildBadgeTypeSmoke:
		return "SMOKE"
	case GuildBadgeTypeSnow:
		return "SNOW"
	case GuildBadgeTypeSound:
		return "SOUND"
	case GuildBadgeTypeSun:
		return "SUN"
	case GuildBadgeTypeWind:
		return "WIND"
	case GuildBadgeTypeBunny:
		return "BUNNY"
	case GuildBadgeTypeDog:
		return "DOG"
	case GuildBadgeTypeFrog:
		return "FROG"
	case GuildBadgeTypeGoat:
		return "GOAT"
	case GuildBadgeTypeCat:
		return "CAT"
	case GuildBadgeTypeDiamond:
		return "DIAMOND"
	case GuildBadgeTypeCrown:
		return "CROWN"
	case GuildBadgeTypeTrophy:
		return "TROPHY"
	case GuildBadgeTypeMoneyBag:
		return "MONEY_BAG"
	case GuildBadgeTypeDollarSign:
		return "DOLLAR_SIGN"
	case GuildBadgeTypeClover:
		return "CLOVER"
	case GuildBadgeTypeBlossom:
		return "BLOSSOM"
	case GuildBadgeTypePottedPlant:
		return "POTTED_PLANT"
	case GuildBadgeTypeMaple:
		return "MAPLE"
	case GuildBadgeTypeWiltedFlower:
		return "WILTED_FLOWER"
	case GuildBadgeTypeButterfly:
		return "BUTTERFLY"
	case GuildBadgeTypeSnail:
		return "SNAIL"
	case GuildBadgeTypeCaterpillar:
		return "CATERPILLAR"
	case GuildBadgeTypeSpider:
		return "SPIDER"
	case GuildBadgeTypeBee:
		return "BEE"
	}
	return strconv.FormatInt(int64(v), 10)
}

type GuildExplicitContentFilterLevel int64

const (
	GuildExplicitContentFilterLevelDisabled            GuildExplicitContentFilterLevel = 0
	GuildExplicitContentFilterLevelMembersWithoutRoles GuildExplicitContentFilterLevel = 1
	GuildExplicitContentFilterLevelAllMembers          GuildExplicitContentFilterLevel = 2
)

func (v GuildExplicitContentFilterLevel) String() string {
	switch v {
	case GuildExplicitContentFilterLevelDisabled:
		return "DISABLED"
	case GuildExplicitContentFilterLevelMembersWithoutRoles:
		return "MEMBERS_WITHOUT_ROLES"
	case GuildExplicitContentFilterLevelAllMembers:
		return "ALL_MEMBERS"
	}
	return strconv.FormatInt(int64(v), 10)
}

type GuildFeatures string

const (
	GuildFeaturesActivitiesAlpha                             GuildFeatures = "ACTIVITIES_ALPHA"
	GuildFeaturesActivitiesEmployee                          GuildFeatures = "ACTIVITIES_EMPLOYEE"
	GuildFeaturesActivitiesInternalDev                       GuildFeatures = "ACTIVITIES_INTERNAL_DEV"
	GuildFeaturesActivityFeedDisabledByUser                  GuildFeatures = "ACTIVITY_FEED_DISABLED_BY_USER"
	GuildFeaturesActivityFeedEnabledByUser                   GuildFeatures = "ACTIVITY_FEED_ENABLED_BY_USER"
	GuildFeaturesAgeVerificationLargeGuild                   GuildFeatures = "AGE_VERIFICATION_LARGE_GUILD"
	GuildFeaturesAnimatedBanner                              GuildFeatures = "ANIMATED_BANNER"
	GuildFeaturesAnimatedIcon                                GuildFeatures = "ANIMATED_ICON"
	GuildFeaturesApplicationCommandPermissionsV2             GuildFeatures = "APPLICATION_COMMAND_PERMISSIONS_V2"
	GuildFeaturesAudioBitrate128Kbps                         GuildFeatures = "AUDIO_BITRATE_128_KBPS"
	GuildFeaturesAudioBitrate256Kbps                         GuildFeatures = "AUDIO_BITRATE_256_KBPS"
	GuildFeaturesAudioBitrate384Kbps                         GuildFeatures = "AUDIO_BITRATE_384_KBPS"
	GuildFeaturesAutoModeration                              GuildFeatures = "AUTO_MODERATION"
	GuildFeaturesAutomodTriggerKeywordFilter                 GuildFeatures = "AUTOMOD_TRIGGER_KEYWORD_FILTER"
	GuildFeaturesAutomodTriggerMlSpamFilter                  GuildFeatures = "AUTOMOD_TRIGGER_ML_SPAM_FILTER"
	GuildFeaturesAutomodTriggerSpamLinkFilter                GuildFeatures = "AUTOMOD_TRIGGER_SPAM_LINK_FILTER"
	GuildFeaturesAutomodTriggerUserProfile                   GuildFeatures = "AUTOMOD_TRIGGER_USER_PROFILE"
	GuildFeaturesBanner                                      GuildFeatures = "BANNER"
	GuildFeaturesBfg                                         GuildFeatures = "BFG"
	GuildFeaturesBoostingTiersExperimentMediumGuild          GuildFeatures = "BOOSTING_TIERS_EXPERIMENT_MEDIUM_GUILD"
	GuildFeaturesBoostingTiersExperimentSmallGuild           GuildFeatures = "BOOSTING_TIERS_EXPERIMENT_SMALL_GUILD"
	GuildFeaturesBotDeveloperEarlyAccess                     GuildFeatures = "BOT_DEVELOPER_EARLY_ACCESS"
	GuildFeaturesBurstReactions                              GuildFeatures = "BURST_REACTIONS"
	GuildFeaturesBypassSlowmodePermissionMigrationComplete   GuildFeatures = "BYPASS_SLOWMODE_PERMISSION_MIGRATION_COMPLETE"
	GuildFeaturesChannelBanner                               GuildFeatures = "CHANNEL_BANNER"
	GuildFeaturesChannelEmojisGenerated                      GuildFeatures = "CHANNEL_EMOJIS_GENERATED"
	GuildFeaturesChannelIconEmojisGenerated                  GuildFeatures = "CHANNEL_ICON_EMOJIS_GENERATED"
	GuildFeaturesChannelHighlights                           GuildFeatures = "CHANNEL_HIGHLIGHTS"
	GuildFeaturesClan                                        GuildFeatures = "CLAN"
	GuildFeaturesClanDiscoveryDisabled                       GuildFeatures = "CLAN_DISCOVERY_DISABLED"
	GuildFeaturesClanPilotGenshin                            GuildFeatures = "CLAN_PILOT_GENSHIN"
	GuildFeaturesClanPilotValorant                           GuildFeatures = "CLAN_PILOT_VALORANT"
	GuildFeaturesClydeDisabled                               GuildFeatures = "CLYDE_DISABLED"
	GuildFeaturesClydeEnabled                                GuildFeatures = "CLYDE_ENABLED"
	GuildFeaturesClydeExperimentEnabled                      GuildFeatures = "CLYDE_EXPERIMENT_ENABLED"
	GuildFeaturesCommerce                                    GuildFeatures = "COMMERCE"
	GuildFeaturesCommunity                                   GuildFeatures = "COMMUNITY"
	GuildFeaturesCommunityCanary                             GuildFeatures = "COMMUNITY_CANARY"
	GuildFeaturesConference                                  GuildFeatures = "CONFERENCE"
	GuildFeaturesConsideredExternallyDiscoverable            GuildFeatures = "CONSIDERED_EXTERNALLY_DISCOVERABLE"
	GuildFeaturesCreatorAcceptedNewTerms                     GuildFeatures = "CREATOR_ACCEPTED_NEW_TERMS"
	GuildFeaturesCreatorMonetizable                          GuildFeatures = "CREATOR_MONETIZABLE"
	GuildFeaturesCreatorMonetizableDisabled                  GuildFeatures = "CREATOR_MONETIZABLE_DISABLED"
	GuildFeaturesCreatorMonetizablePendingNewOwnerOnboarding GuildFeatures = "CREATOR_MONETIZABLE_PENDING_NEW_OWNER_ONBOARDING"
	GuildFeaturesCreatorMonetizableProvisional               GuildFeatures = "CREATOR_MONETIZABLE_PROVISIONAL"
	GuildFeaturesCreatorMonetizableRestricted                GuildFeatures = "CREATOR_MONETIZABLE_RESTRICTED"
	GuildFeaturesCreatorMonetizableWhiteglove                GuildFeatures = "CREATOR_MONETIZABLE_WHITEGLOVE"
	GuildFeaturesCreatorStorePage                            GuildFeatures = "CREATOR_STORE_PAGE"
	GuildFeaturesDeveloperSupportServer                      GuildFeatures = "DEVELOPER_SUPPORT_SERVER"
	GuildFeaturesDiscoverable                                GuildFeatures = "DISCOVERABLE"
	GuildFeaturesDiscoverableDisabled                        GuildFeatures = "DISCOVERABLE_DISABLED"
	GuildFeaturesEnabledDiscoverableBefore                   GuildFeatures = "ENABLED_DISCOVERABLE_BEFORE"
	GuildFeaturesEnabledModerationExperienceForNonCommunity  GuildFeatures = "ENABLED_MODERATION_EXPERIENCE_FOR_NON_COMMUNITY"
	GuildFeaturesEnhancedRoleColors                          GuildFeatures = "ENHANCED_ROLE_COLORS"
	GuildFeaturesExposedToActivitiesWtpExperiment            GuildFeatures = "EXPOSED_TO_ACTIVITIES_WTP_EXPERIMENT"
	GuildFeaturesExposedToBoostingTiersExperiment            GuildFeatures = "EXPOSED_TO_BOOSTING_TIERS_EXPERIMENT"
	GuildFeaturesFeaturable                                  GuildFeatures = "FEATURABLE"
	GuildFeaturesForceRelay                                  GuildFeatures = "FORCE_RELAY"
	GuildFeaturesForwardingDisabled                          GuildFeatures = "FORWARDING_DISABLED"
	GuildFeaturesGameServerHosting                           GuildFeatures = "GAME_SERVER_HOSTING"
	GuildFeaturesGameServers                                 GuildFeatures = "GAME_SERVERS"
	GuildFeaturesGenshinL30                                  GuildFeatures = "GENSHIN_L30"
	GuildFeaturesGuestsEnabled                               GuildFeatures = "GUESTS_ENABLED"
	GuildFeaturesGuildAutomodDefaultList                     GuildFeatures = "GUILD_AUTOMOD_DEFAULT_LIST"
	GuildFeaturesGuildCommunicationDisabledGuilds            GuildFeatures = "GUILD_COMMUNICATION_DISABLED_GUILDS"
	GuildFeaturesGuildHomeDeprecationOverride                GuildFeatures = "GUILD_HOME_DEPRECATION_OVERRIDE"
	GuildFeaturesGuildHomeOverride                           GuildFeatures = "GUILD_HOME_OVERRIDE"
	GuildFeaturesGuildHomeTest                               GuildFeatures = "GUILD_HOME_TEST"
	GuildFeaturesGuildMemberVerificationExperiment           GuildFeatures = "GUILD_MEMBER_VERIFICATION_EXPERIMENT"
	GuildFeaturesGuildOnboarding                             GuildFeatures = "GUILD_ONBOARDING"
	GuildFeaturesGuildOnboardingAdminOnly                    GuildFeatures = "GUILD_ONBOARDING_ADMIN_ONLY"
	GuildFeaturesGuildOnboardingEverEnabled                  GuildFeatures = "GUILD_ONBOARDING_EVER_ENABLED"
	GuildFeaturesGuildOnboardingHasPrompts                   GuildFeatures = "GUILD_ONBOARDING_HAS_PROMPTS"
	GuildFeaturesGuildProducts                               GuildFeatures = "GUILD_PRODUCTS"
	GuildFeaturesGuildProductsAllowArchivedFile              GuildFeatures = "GUILD_PRODUCTS_ALLOW_ARCHIVED_FILE"
	GuildFeaturesGuildRoleSubscriptions                      GuildFeatures = "GUILD_ROLE_SUBSCRIPTIONS"
	GuildFeaturesGuildRoleSubscriptionPurchaseFeedbackLoop   GuildFeatures = "GUILD_ROLE_SUBSCRIPTION_PURCHASE_FEEDBACK_LOOP"
	GuildFeaturesGuildRoleSubscriptionTierTemplate           GuildFeatures = "GUILD_ROLE_SUBSCRIPTION_TIER_TEMPLATE"
	GuildFeaturesGuildRoleSubscriptionTrials                 GuildFeatures = "GUILD_ROLE_SUBSCRIPTION_TRIALS"
	GuildFeaturesGuildServerGuide                            GuildFeatures = "GUILD_SERVER_GUIDE"
	GuildFeaturesGuildTags                                   GuildFeatures = "GUILD_TAGS"
	GuildFeaturesGuildTagsBadgePackCreepyCrawlies            GuildFeatures = "GUILD_TAGS_BADGE_PACK_CREEPY_CRAWLIES"
	GuildFeaturesGuildTagsBadgePackFlex                      GuildFeatures = "GUILD_TAGS_BADGE_PACK_FLEX"
	GuildFeaturesGuildTagsBadgePackPets                      GuildFeatures = "GUILD_TAGS_BADGE_PACK_PETS"
	GuildFeaturesGuildTagsBadgePackPlant                     GuildFeatures = "GUILD_TAGS_BADGE_PACK_PLANT"
	GuildFeaturesGuildWebPageVanityURL                       GuildFeatures = "GUILD_WEB_PAGE_VANITY_URL"
	GuildFeaturesHadEarlyActivitiesAccess                    GuildFeatures = "HAD_EARLY_ACTIVITIES_ACCESS"
	GuildFeaturesHasDirectoryEntry                           GuildFeatures = "HAS_DIRECTORY_ENTRY"
	GuildFeaturesHub                                         GuildFeatures = "HUB"
	GuildFeaturesIncreasedThreadLimit                        GuildFeatures = "INCREASED_THREAD_LIMIT"
	GuildFeaturesInternalEmployeeOnly                        GuildFeatures = "INTERNAL_EMPLOYEE_ONLY"
	GuildFeaturesInviteSplash                                GuildFeatures = "INVITE_SPLASH"
	GuildFeaturesInvitesDisabled                             GuildFeatures = "INVITES_DISABLED"
	GuildFeaturesLeaderboardEnabled                          GuildFeatures = "LEADERBOARD_ENABLED"
	GuildFeaturesLinkedToHub                                 GuildFeatures = "LINKED_TO_HUB"
	GuildFeaturesLurkable                                    GuildFeatures = "LURKABLE"
	GuildFeaturesMarketplacesConnectionRoles                 GuildFeatures = "MARKETPLACES_CONNECTION_ROLES"
	GuildFeaturesMaxFileSize50Mb                             GuildFeatures = "MAX_FILE_SIZE_50_MB"
	GuildFeaturesMaxFileSize100Mb                            GuildFeatures = "MAX_FILE_SIZE_100_MB"
	GuildFeaturesMaxFileSize250Mb                            GuildFeatures = "MAX_FILE_SIZE_250_MB"
	GuildFeaturesMediaChannelAlpha                           GuildFeatures = "MEDIA_CHANNEL_ALPHA"
	GuildFeaturesMemberListDisabled                          GuildFeatures = "MEMBER_LIST_DISABLED"
	GuildFeaturesMemberProfiles                              GuildFeatures = "MEMBER_PROFILES"
	GuildFeaturesMemberSafetyPageRollout                     GuildFeatures = "MEMBER_SAFETY_PAGE_ROLLOUT"
	GuildFeaturesMemberVerificationGateEnabled               GuildFeatures = "MEMBER_VERIFICATION_GATE_ENABLED"
	GuildFeaturesMemberVerificationManualApproval            GuildFeatures = "MEMBER_VERIFICATION_MANUAL_APPROVAL"
	GuildFeaturesMemberVerificationRolloutTest               GuildFeatures = "MEMBER_VERIFICATION_ROLLOUT_TEST"
	GuildFeaturesMobileWebRoleSubscriptionPurchasePage       GuildFeatures = "MOBILE_WEB_ROLE_SUBSCRIPTION_PURCHASE_PAGE"
	GuildFeaturesMonetizationEnabled                         GuildFeatures = "MONETIZATION_ENABLED"
	GuildFeaturesMoreEmoji                                   GuildFeatures = "MORE_EMOJI"
	GuildFeaturesMoreSoundboard                              GuildFeatures = "MORE_SOUNDBOARD"
	GuildFeaturesMoreStickers                                GuildFeatures = "MORE_STICKERS"
	GuildFeaturesNews                                        GuildFeatures = "NEWS"
	GuildFeaturesNewThreadPermissions                        GuildFeatures = "NEW_THREAD_PERMISSIONS"
	GuildFeaturesNonCommunityRaidAlerts                      GuildFeatures = "NON_COMMUNITY_RAID_ALERTS"
	GuildFeaturesOfficialGameGuild                           GuildFeatures = "OFFICIAL_GAME_GUILD"
	GuildFeaturesPartnered                                   GuildFeatures = "PARTNERED"
	GuildFeaturesPinPermissionMigrationComplete              GuildFeatures = "PIN_PERMISSION_MIGRATION_COMPLETE"
	GuildFeaturesPowerupBetaFeatures                         GuildFeatures = "POWERUP_BETA_FEATURES"
	GuildFeaturesPremiumTier3Override                        GuildFeatures = "PREMIUM_TIER_3_OVERRIDE"
	GuildFeaturesPreviewEnabled                              GuildFeatures = "PREVIEW_ENABLED"
	GuildFeaturesPrivateThreads                              GuildFeatures = "PRIVATE_THREADS"
	GuildFeaturesProductsAvailableForPurchase                GuildFeatures = "PRODUCTS_AVAILABLE_FOR_PURCHASE"
	GuildFeaturesPruneRequiresAdmin                          GuildFeatures = "PRUNE_REQUIRES_ADMIN"
	GuildFeaturesPublic                                      GuildFeatures = "PUBLIC"
	GuildFeaturesPublicDisabled                              GuildFeatures = "PUBLIC_DISABLED"
	GuildFeaturesRaidAlertsDisabled                          GuildFeatures = "RAID_ALERTS_DISABLED"
	GuildFeaturesRaidAlertsEnabled                           GuildFeatures = "RAID_ALERTS_ENABLED"
	GuildFeaturesRapidashTest                                GuildFeatures = "RAPIDASH_TEST"
	GuildFeaturesRapidashTestRebirth                         GuildFeatures = "RAPIDASH_TEST_REBIRTH"
	GuildFeaturesRelayEnabled                                GuildFeatures = "RELAY_ENABLED"
	GuildFeaturesReportToModPilot                            GuildFeatures = "REPORT_TO_MOD_PILOT"
	GuildFeaturesReportToModSurvey                           GuildFeatures = "REPORT_TO_MOD_SURVEY"
	GuildFeaturesRestrictSpamRiskGuilds                      GuildFeatures = "RESTRICT_SPAM_RISK_GUILDS"
	GuildFeaturesRoleIcons                                   GuildFeatures = "ROLE_ICONS"
	GuildFeaturesRoleSubscriptionsAvailableForPurchase       GuildFeatures = "ROLE_SUBSCRIPTIONS_AVAILABLE_FOR_PURCHASE"
	GuildFeaturesRoleSubscriptionsEnabled                    GuildFeatures = "ROLE_SUBSCRIPTIONS_ENABLED"
	GuildFeaturesRoleSubscriptionsEnabledForPurchase         GuildFeatures = "ROLE_SUBSCRIPTIONS_ENABLED_FOR_PURCHASE"
	GuildFeaturesSevenDayThreadArchive                       GuildFeatures = "SEVEN_DAY_THREAD_ARCHIVE"
	GuildFeaturesServerProfilesTest                          GuildFeatures = "SERVER_PROFILES_TEST"
	GuildFeaturesShard                                       GuildFeatures = "SHARD"
	GuildFeaturesSharedCanvasFriendsAndFamilyTest            GuildFeatures = "SHARED_CANVAS_FRIENDS_AND_FAMILY_TEST"
	GuildFeaturesSocialLayerStorefront                       GuildFeatures = "SOCIAL_LAYER_STOREFRONT"
	GuildFeaturesSoundboard                                  GuildFeatures = "SOUNDBOARD"
	GuildFeaturesStageChannelViewers50                       GuildFeatures = "STAGE_CHANNEL_VIEWERS_50"
	GuildFeaturesStageChannelViewers150                      GuildFeatures = "STAGE_CHANNEL_VIEWERS_150"
	GuildFeaturesStageChannelViewers300                      GuildFeatures = "STAGE_CHANNEL_VIEWERS_300"
	GuildFeaturesSummariesEnabled                            GuildFeatures = "SUMMARIES_ENABLED"
	GuildFeaturesSummariesEnabledGa                          GuildFeatures = "SUMMARIES_ENABLED_GA"
	GuildFeaturesSummariesDisabledByUser                     GuildFeatures = "SUMMARIES_DISABLED_BY_USER"
	GuildFeaturesSummariesEnabledByUser                      GuildFeatures = "SUMMARIES_ENABLED_BY_USER"
	GuildFeaturesSummariesLongLookback                       GuildFeatures = "SUMMARIES_LONG_LOOKBACK"
	GuildFeaturesSummariesOptOutExperience                   GuildFeatures = "SUMMARIES_OPT_OUT_EXPERIENCE"
	GuildFeaturesStaffLevelCollaboratorRequired              GuildFeatures = "STAFF_LEVEL_COLLABORATOR_REQUIRED"
	GuildFeaturesStaffLevelRestrictedCollaboratorRequired    GuildFeatures = "STAFF_LEVEL_RESTRICTED_COLLABORATOR_REQUIRED"
	GuildFeaturesTextInStageEnabled                          GuildFeatures = "TEXT_IN_STAGE_ENABLED"
	GuildFeaturesTextInVoiceEnabled                          GuildFeatures = "TEXT_IN_VOICE_ENABLED"
	GuildFeaturesThreadsEnabled                              GuildFeatures = "THREADS_ENABLED"
	GuildFeaturesThreadsEnabledTesting                       GuildFeatures = "THREADS_ENABLED_TESTING"
	GuildFeaturesThreadsOnlyChannel                          GuildFeatures = "THREADS_ONLY_CHANNEL"
	GuildFeaturesThreeDayThreadArchive                       GuildFeatures = "THREE_DAY_THREAD_ARCHIVE"
	GuildFeaturesTicketedEventsEnabled                       GuildFeatures = "TICKETED_EVENTS_ENABLED"
	GuildFeaturesTicketingEnabled                            GuildFeatures = "TICKETING_ENABLED"
	GuildFeaturesTierlessBoosting                            GuildFeatures = "TIERLESS_BOOSTING"
	GuildFeaturesTierlessBoostingClientTest                  GuildFeatures = "TIERLESS_BOOSTING_CLIENT_TEST"
	GuildFeaturesTierlessBoostingSystemMessage               GuildFeatures = "TIERLESS_BOOSTING_SYSTEM_MESSAGE"
	GuildFeaturesTierlessBoostingTest                        GuildFeatures = "TIERLESS_BOOSTING_TEST"
	GuildFeaturesValorantL30                                 GuildFeatures = "VALORANT_L30"
	GuildFeaturesVanityURL                                   GuildFeatures = "VANITY_URL"
	GuildFeaturesVerified                                    GuildFeatures = "VERIFIED"
	GuildFeaturesVideoBitrateEnhanced                        GuildFeatures = "VIDEO_BITRATE_ENHANCED"
	GuildFeaturesVideoQuality72060fps                        GuildFeatures = "VIDEO_QUALITY_720_60FPS"
	GuildFeaturesVideoQuality108060fps                       GuildFeatures = "VIDEO_QUALITY_1080_60FPS"
	GuildFeaturesVipRegions                                  GuildFeatures = "VIP_REGIONS"
	GuildFeaturesVoiceChannelEffects                         GuildFeatures = "VOICE_CHANNEL_EFFECTS"
	GuildFeaturesVoiceInThreads                              GuildFeatures = "VOICE_IN_THREADS"
	GuildFeaturesWelcomeScreenEnabled                        GuildFeatures = "WELCOME_SCREEN_ENABLED"
)

type GuildJoinRequestStatus string

const (
	GuildJoinRequestStatusStarted   GuildJoinRequestStatus = "STARTED"
	GuildJoinRequestStatusSubmitted GuildJoinRequestStatus = "SUBMITTED"
	GuildJoinRequestStatusRejected  GuildJoinRequestStatus = "REJECTED"
	GuildJoinRequestStatusApproved  GuildJoinRequestStatus = "APPROVED"
)

type GuildMemberFlags uint64

const (
	GuildMemberFlagsDidRejoin                    GuildMemberFlags = 1 << 0
	GuildMemberFlagsCompletedOnboarding          GuildMemberFlags = 1 << 1
	GuildMemberFlagsBypassesVerification         GuildMemberFlags = 1 << 2
	GuildMemberFlagsStartedOnboarding            GuildMemberFlags = 1 << 3
	GuildMemberFlagsIsGuest                      GuildMemberFlags = 1 << 4
	GuildMemberFlagsStartedHomeActions           GuildMemberFlags = 1 << 5
	GuildMemberFlagsCompletedHomeActions         GuildMemberFlags = 1 << 6
	GuildMemberFlagsAutomodQuarantinedName       GuildMemberFlags = 1 << 7
	GuildMemberFlagsAutomodQuarantinedBio        GuildMemberFlags = 1 << 8
	GuildMemberFlagsDMSettingsUpsellAcknowledged GuildMemberFlags = 1 << 9
	GuildMemberFlagsAutomodQuarantinedGuildTag   GuildMemberFlags = 1 << 10
)

func (v GuildMemberFlags) String() string {
	switch v {
	case GuildMemberFlagsDidRejoin:
		return "DID_REJOIN"
	case GuildMemberFlagsCompletedOnboarding:
		return "COMPLETED_ONBOARDING"
	case GuildMemberFlagsBypassesVerification:
		return "BYPASSES_VERIFICATION"
	case GuildMemberFlagsStartedOnboarding:
		return "STARTED_ONBOARDING"
	case GuildMemberFlagsIsGuest:
		return "IS_GUEST"
	case GuildMemberFlagsStartedHomeActions:
		return "STARTED_HOME_ACTIONS"
	case GuildMemberFlagsCompletedHomeActions:
		return "COMPLETED_HOME_ACTIONS"
	case GuildMemberFlagsAutomodQuarantinedName:
		return "AUTOMOD_QUARANTINED_NAME"
	case GuildMemberFlagsAutomodQuarantinedBio:
		return "AUTOMOD_QUARANTINED_BIO"
	case GuildMemberFlagsDMSettingsUpsellAcknowledged:
		return "DM_SETTINGS_UPSELL_ACKNOWLEDGED"
	case GuildMemberFlagsAutomodQuarantinedGuildTag:
		return "AUTOMOD_QUARANTINED_GUILD_TAG"
	}
	return strconv.FormatUint(uint64(v), 10)
}

type GuildMemberType int64

const (
	GuildMemberTypeOwner  GuildMemberType = 1
	GuildMemberTypeMember GuildMemberType = 2
)

func (v GuildMemberType) String() string {
	switch v {
	case GuildMemberTypeOwner:
		return "OWNER"
	case GuildMemberTypeMember:
		return "MEMBER"
	}
	return strconv.FormatInt(int64(v), 10)
}

type GuildOnboardingProgressFlags uint64

const (
	GuildOnboardingProgressFlagsNoticeShown   GuildOnboardingProgressFlags = 1 << 0
	GuildOnboardingProgressFlagsNoticeCleared GuildOnboardingProgressFlags = 1 << 1
)

func (v GuildOnboardingProgressFlags) String() string {
	switch v {
	case GuildOnboardingProgressFlagsNoticeShown:
		return "NOTICE_SHOWN"
	case GuildOnboardingProgressFlagsNoticeCleared:
		return "NOTICE_CLEARED"
	}
	return strconv.FormatUint(uint64(v), 10)
}

type GuildPowerupCategoryType string

const (
	GuildPowerupCategoryTypeLevel      GuildPowerupCategoryType = "level"
	GuildPowerupCategoryTypePerk       GuildPowerupCategoryType = "perk"
	GuildPowerupCategoryTypeGameServer GuildPowerupCategoryType = "game_server"
)

type GuildProductShopTabsOrderType int64

const (
	GuildProductShopTabsOrderTypeSubsFirst     GuildProductShopTabsOrderType = 1
	GuildProductShopTabsOrderTypeProductsFirst GuildProductShopTabsOrderType = 2
	GuildProductShopTabsOrderTypeUnknown3      GuildProductShopTabsOrderType = 3
)

func (v GuildProductShopTabsOrderType) String() string {
	switch v {
	case GuildProductShopTabsOrderTypeSubsFirst:
		return "SUBS_FIRST"
	case GuildProductShopTabsOrderTypeProductsFirst:
		return "PRODUCTS_FIRST"
	case GuildProductShopTabsOrderTypeUnknown3:
		return "UNKNOWN_3"
	}
	return strconv.FormatInt(int64(v), 10)
}

type GuildProductSortOrderType int64

const (
	GuildProductSortOrderTypeName           GuildProductSortOrderType = 1
	GuildProductSortOrderTypePriceAsc       GuildProductSortOrderType = 2
	GuildProductSortOrderTypePriceDesc      GuildProductSortOrderType = 3
	GuildProductSortOrderTypeNewestArrivals GuildProductSortOrderType = 4
)

func (v GuildProductSortOrderType) String() string {
	switch v {
	case GuildProductSortOrderTypeName:
		return "NAME"
	case GuildProductSortOrderTypePriceAsc:
		return "PRICE_ASC"
	case GuildProductSortOrderTypePriceDesc:
		return "PRICE_DESC"
	case GuildProductSortOrderTypeNewestArrivals:
		return "NEWEST_ARRIVALS"
	}
	return strconv.FormatInt(int64(v), 10)
}

type GuildScheduledEventEntityType int64

const (
	GuildScheduledEventEntityTypeStageInstance GuildScheduledEventEntityType = 1
	GuildScheduledEventEntityTypeVoice         GuildScheduledEventEntityType = 2
	GuildScheduledEventEntityTypeExternal      GuildScheduledEventEntityType = 3
	GuildScheduledEventEntityTypePrimeTime     GuildScheduledEventEntityType = 4
)

func (v GuildScheduledEventEntityType) String() string {
	switch v {
	case GuildScheduledEventEntityTypeStageInstance:
		return "STAGE_INSTANCE"
	case GuildScheduledEventEntityTypeVoice:
		return "VOICE"
	case GuildScheduledEventEntityTypeExternal:
		return "EXTERNAL"
	case GuildScheduledEventEntityTypePrimeTime:
		return "PRIME_TIME"
	}
	return strconv.FormatInt(int64(v), 10)
}

type GuildScheduledEventRecurrenceRuleFrequency int64

const (
	GuildScheduledEventRecurrenceRuleFrequencyYearly  GuildScheduledEventRecurrenceRuleFrequency = 0
	GuildScheduledEventRecurrenceRuleFrequencyMonthly GuildScheduledEventRecurrenceRuleFrequency = 1
	GuildScheduledEventRecurrenceRuleFrequencyWeekly  GuildScheduledEventRecurrenceRuleFrequency = 2
	GuildScheduledEventRecurrenceRuleFrequencyDaily   GuildScheduledEventRecurrenceRuleFrequency = 3
)

func (v GuildScheduledEventRecurrenceRuleFrequency) String() string {
	switch v {
	case GuildScheduledEventRecurrenceRuleFrequencyYearly:
		return "YEARLY"
	case GuildScheduledEventRecurrenceRuleFrequencyMonthly:
		return "MONTHLY"
	case GuildScheduledEventRecurrenceRuleFrequencyWeekly:
		return "WEEKLY"
	case GuildScheduledEventRecurrenceRuleFrequencyDaily:
		return "DAILY"
	}
	return strconv.FormatInt(int64(v), 10)
}

type GuildScheduledEventRecurrenceRuleMonth int64

const (
	GuildScheduledEventRecurrenceRuleMonthJanuary   GuildScheduledEventRecurrenceRuleMonth = 1
	GuildScheduledEventRecurrenceRuleMonthFebruary  GuildScheduledEventRecurrenceRuleMonth = 2
	GuildScheduledEventRecurrenceRuleMonthMarch     GuildScheduledEventRecurrenceRuleMonth = 3
	GuildScheduledEventRecurrenceRuleMonthApril     GuildScheduledEventRecurrenceRuleMonth = 4
	GuildScheduledEventRecurrenceRuleMonthMay       GuildScheduledEventRecurrenceRuleMonth = 5
	GuildScheduledEventRecurrenceRuleMonthJune      GuildScheduledEventRecurrenceRuleMonth = 6
	GuildScheduledEventRecurrenceRuleMonthJuly      GuildScheduledEventRecurrenceRuleMonth = 7
	GuildScheduledEventRecurrenceRuleMonthAugust    GuildScheduledEventRecurrenceRuleMonth = 8
	GuildScheduledEventRecurrenceRuleMonthSeptember GuildScheduledEventRecurrenceRuleMonth = 9
	GuildScheduledEventRecurrenceRuleMonthOctober   GuildScheduledEventRecurrenceRuleMonth = 10
	GuildScheduledEventRecurrenceRuleMonthNovember  GuildScheduledEventRecurrenceRuleMonth = 11
	GuildScheduledEventRecurrenceRuleMonthDecember  GuildScheduledEventRecurrenceRuleMonth = 12
)

func (v GuildScheduledEventRecurrenceRuleMonth) String() string {
	switch v {
	case GuildScheduledEventRecurrenceRuleMonthJanuary:
		return "JANUARY"
	case GuildScheduledEventRecurrenceRuleMonthFebruary:
		return "FEBRUARY"
	case GuildScheduledEventRecurrenceRuleMonthMarch:
		return "MARCH"
	case GuildScheduledEventRecurrenceRuleMonthApril:
		return "APRIL"
	case GuildScheduledEventRecurrenceRuleMonthMay:
		return "MAY"
	case GuildScheduledEventRecurrenceRuleMonthJune:
		return "JUNE"
	case GuildScheduledEventRecurrenceRuleMonthJuly:
		return "JULY"
	case GuildScheduledEventRecurrenceRuleMonthAugust:
		return "AUGUST"
	case GuildScheduledEventRecurrenceRuleMonthSeptember:
		return "SEPTEMBER"
	case GuildScheduledEventRecurrenceRuleMonthOctober:
		return "OCTOBER"
	case GuildScheduledEventRecurrenceRuleMonthNovember:
		return "NOVEMBER"
	case GuildScheduledEventRecurrenceRuleMonthDecember:
		return "DECEMBER"
	}
	return strconv.FormatInt(int64(v), 10)
}

type GuildScheduledEventRecurrenceRuleWeekday int64

const (
	GuildScheduledEventRecurrenceRuleWeekdayMonday    GuildScheduledEventRecurrenceRuleWeekday = 0
	GuildScheduledEventRecurrenceRuleWeekdayTuesday   GuildScheduledEventRecurrenceRuleWeekday = 1
	GuildScheduledEventRecurrenceRuleWeekdayWednesday GuildScheduledEventRecurrenceRuleWeekday = 2
	GuildScheduledEventRecurrenceRuleWeekdayThursday  GuildScheduledEventRecurrenceRuleWeekday = 3
	GuildScheduledEventRecurrenceRuleWeekdayFriday    GuildScheduledEventRecurrenceRuleWeekday = 4
	GuildScheduledEventRecurrenceRuleWeekdaySaturday  GuildScheduledEventRecurrenceRuleWeekday = 5
	GuildScheduledEventRecurrenceRuleWeekdaySunday    GuildScheduledEventRecurrenceRuleWeekday = 6
)

func (v GuildScheduledEventRecurrenceRuleWeekday) String() string {
	switch v {
	case GuildScheduledEventRecurrenceRuleWeekdayMonday:
		return "MONDAY"
	case GuildScheduledEventRecurrenceRuleWeekdayTuesday:
		return "TUESDAY"
	case GuildScheduledEventRecurrenceRuleWeekdayWednesday:
		return "WEDNESDAY"
	case GuildScheduledEventRecurrenceRuleWeekdayThursday:
		return "THURSDAY"
	case GuildScheduledEventRecurrenceRuleWeekdayFriday:
		return "FRIDAY"
	case GuildScheduledEventRecurrenceRuleWeekdaySaturday:
		return "SATURDAY"
	case GuildScheduledEventRecurrenceRuleWeekdaySunday:
		return "SUNDAY"
	}
	return strconv.FormatInt(int64(v), 10)
}

type GuildScheduledEventStatus int64

const (
	GuildScheduledEventStatusScheduled GuildScheduledEventStatus = 1
	GuildScheduledEventStatusActive    GuildScheduledEventStatus = 2
	GuildScheduledEventStatusCompleted GuildScheduledEventStatus = 3
	GuildScheduledEventStatusCanceled  GuildScheduledEventStatus = 4
)

func (v GuildScheduledEventStatus) String() string {
	switch v {
	case GuildScheduledEventStatusScheduled:
		return "SCHEDULED"
	case GuildScheduledEventStatusActive:
		return "ACTIVE"
	case GuildScheduledEventStatusCompleted:
		return "COMPLETED"
	case GuildScheduledEventStatusCanceled:
		return "CANCELED"
	}
	return strconv.FormatInt(int64(v), 10)
}

type GuildScheduledEventUserResponse int64

const (
	GuildScheduledEventUserResponseUninterested GuildScheduledEventUserResponse = 0
	GuildScheduledEventUserResponseInterested   GuildScheduledEventUserResponse = 1
)

func (v GuildScheduledEventUserResponse) String() string {
	switch v {
	case GuildScheduledEventUserResponseUninterested:
		return "UNINTERESTED"
	case GuildScheduledEventUserResponseInterested:
		return "INTERESTED"
	}
	return strconv.FormatInt(int64(v), 10)
}

type GuildThemeSourcePreference int64

const (
	GuildThemeSourcePreferenceUnspecified GuildThemeSourcePreference = 0
	GuildThemeSourcePreferencePersonal    GuildThemeSourcePreference = 1
	GuildThemeSourcePreferenceGuild       GuildThemeSourcePreference = 2
)

func (v GuildThemeSourcePreference) String() string {
	switch v {
	case GuildThemeSourcePreferenceUnspecified:
		return "UNSPECIFIED"
	case GuildThemeSourcePreferencePersonal:
		return "PERSONAL"
	case GuildThemeSourcePreferenceGuild:
		return "GUILD"
	}
	return strconv.FormatInt(int64(v), 10)
}

type GuildVisibility int64

const (
	GuildVisibilityPublic                GuildVisibility = 1
	GuildVisibilityRestricted            GuildVisibility = 2
	GuildVisibilityPublicWithRecruitment GuildVisibility = 3
)

func (v GuildVisibility) String() string {
	switch v {
	case GuildVisibilityPublic:
		return "PUBLIC"
	case GuildVisibilityRestricted:
		return "RESTRICTED"
	case GuildVisibilityPublicWithRecruitment:
		return "PUBLIC_WITH_RECRUITMENT"
	}
	return strconv.FormatInt(int64(v), 10)
}

type GuildWidgetImageStyleOption string

const (
	GuildWidgetImageStyleOptionShield  GuildWidgetImageStyleOption = "shield"
	GuildWidgetImageStyleOptionBanner1 GuildWidgetImageStyleOption = "banner1"
	GuildWidgetImageStyleOptionBanner2 GuildWidgetImageStyleOption = "banner2"
	GuildWidgetImageStyleOptionBanner3 GuildWidgetImageStyleOption = "banner3"
	GuildWidgetImageStyleOptionBanner4 GuildWidgetImageStyleOption = "banner4"
)

type GuildsLeaderboardOptOutDefault int64

const (
	GuildsLeaderboardOptOutDefaultOffForNewGuilds GuildsLeaderboardOptOutDefault = 0
	GuildsLeaderboardOptOutDefaultOnForNewGuilds  GuildsLeaderboardOptOutDefault = 1
)

func (v GuildsLeaderboardOptOutDefault) String() string {
	switch v {
	case GuildsLeaderboardOptOutDefaultOffForNewGuilds:
		return "OFF_FOR_NEW_GUILDS"
	case GuildsLeaderboardOptOutDefaultOnForNewGuilds:
		return "ON_FOR_NEW_GUILDS"
	}
	return strconv.FormatInt(int64(v), 10)
}

type HarvestBackendInternalType string

const (
	HarvestBackendInternalTypeUsers       HarvestBackendInternalType = "users"
	HarvestBackendInternalTypeAnalytics   HarvestBackendInternalType = "analytics"
	HarvestBackendInternalTypeActivitiesE HarvestBackendInternalType = "activities_e"
	HarvestBackendInternalTypeActivitiesW HarvestBackendInternalType = "activities_w"
	HarvestBackendInternalTypeMessages    HarvestBackendInternalType = "messages"
	HarvestBackendInternalTypeHubspot     HarvestBackendInternalType = "hubspot"
	HarvestBackendInternalTypeGuilds      HarvestBackendInternalType = "guilds"
	HarvestBackendInternalTypeAds         HarvestBackendInternalType = "ads"
	HarvestBackendInternalTypeZendesk     HarvestBackendInternalType = "zendesk"
)

type HarvestBackendState string

const (
	HarvestBackendStateInitial   HarvestBackendState = "INITIAL"
	HarvestBackendStateRunning   HarvestBackendState = "RUNNING"
	HarvestBackendStateExtracted HarvestBackendState = "EXTRACTED"
)

type HarvestBackendType string

const (
	HarvestBackendTypeAccounts   HarvestBackendType = "Accounts"
	HarvestBackendTypeAds        HarvestBackendType = "Ads"
	HarvestBackendTypeAnalytics  HarvestBackendType = "Analytics"
	HarvestBackendTypeActivities HarvestBackendType = "Activities"
	HarvestBackendTypeMessages   HarvestBackendType = "Messages"
	HarvestBackendTypePrograms   HarvestBackendType = "Programs"
	HarvestBackendTypeServers    HarvestBackendType = "Servers"
	HarvestBackendTypeZendesk    HarvestBackendType = "Zendesk"
)

type HarvestState string

const (
	HarvestStateIncomplete HarvestState = "INCOMPLETE"
	HarvestStateDelivered  HarvestState = "DELIVERED"
	HarvestStateCancelled  HarvestState = "CANCELLED"
)

type HarvestStatus int64

const (
	HarvestStatusQueued    HarvestStatus = 0
	HarvestStatusRunning   HarvestStatus = 1
	HarvestStatusFailed    HarvestStatus = 2
	HarvestStatusCompleted HarvestStatus = 3
	HarvestStatusCancelled HarvestStatus = 4
)

func (v HarvestStatus) String() string {
	switch v {
	case HarvestStatusQueued:
		return "QUEUED"
	case HarvestStatusRunning:
		return "RUNNING"
	case HarvestStatusFailed:
		return "FAILED"
	case HarvestStatusCompleted:
		return "COMPLETED"
	case HarvestStatusCancelled:
		return "CANCELLED"
	}
	return strconv.FormatInt(int64(v), 10)
}

type HighlightLevel int64

const (
	HighlightLevelDefault  HighlightLevel = 0
	HighlightLevelDisabled HighlightLevel = 1
	HighlightLevelEnabled  HighlightLevel = 2
)

func (v HighlightLevel) String() string {
	switch v {
	case HighlightLevelDefault:
		return "DEFAULT"
	case HighlightLevelDisabled:
		return "DISABLED"
	case HighlightLevelEnabled:
		return "ENABLED"
	}
	return strconv.FormatInt(int64(v), 10)
}

type HubProgressFlags uint64

const (
	HubProgressFlagsJoinGuild   HubProgressFlags = 1 << 0
	HubProgressFlagsInviteUser  HubProgressFlags = 1 << 1
	HubProgressFlagsContactSync HubProgressFlags = 1 << 2
)

func (v HubProgressFlags) String() string {
	switch v {
	case HubProgressFlagsJoinGuild:
		return "JOIN_GUILD"
	case HubProgressFlagsInviteUser:
		return "INVITE_USER"
	case HubProgressFlagsContactSync:
		return "CONTACT_SYNC"
	}
	return strconv.FormatUint(uint64(v), 10)
}

type HubType int64

const (
	HubTypeDefault    HubType = 0
	HubTypeHighSchool HubType = 1
	HubTypeCollege    HubType = 2
)

func (v HubType) String() string {
	switch v {
	case HubTypeDefault:
		return "DEFAULT"
	case HubTypeHighSchool:
		return "HIGH_SCHOOL"
	case HubTypeCollege:
		return "COLLEGE"
	}
	return strconv.FormatInt(int64(v), 10)
}

type HypesquadHouse int64

const (
	HypesquadHouseN1 HypesquadHouse = 1
	HypesquadHouseN2 HypesquadHouse = 2
	HypesquadHouseN3 HypesquadHouse = 3
)

func (v HypesquadHouse) String() string {
	switch v {
	case HypesquadHouseN1:
		return "1"
	case HypesquadHouseN2:
		return "2"
	case HypesquadHouseN3:
		return "3"
	}
	return strconv.FormatInt(int64(v), 10)
}

type IframeModalSize int64

const (
	IframeModalSizeSmall  IframeModalSize = 1
	IframeModalSizeNormal IframeModalSize = 2
	IframeModalSizeBig    IframeModalSize = 3
)

func (v IframeModalSize) String() string {
	switch v {
	case IframeModalSizeSmall:
		return "SMALL"
	case IframeModalSizeNormal:
		return "NORMAL"
	case IframeModalSizeBig:
		return "BIG"
	}
	return strconv.FormatInt(int64(v), 10)
}

type ImageQuality string

const (
	ImageQualityLossless ImageQuality = "lossless"
	ImageQualityHigh     ImageQuality = "high"
	ImageQualityLow      ImageQuality = "low"
)

type ImageType string

const (
	ImageTypeUser ImageType = "user"
)

type InboxTab int64

const (
	InboxTabUnspecified InboxTab = 0
	InboxTabMentions    InboxTab = 1
	InboxTabUnreads     InboxTab = 2
	InboxTabTodos       InboxTab = 3
	InboxTabForYou      InboxTab = 4
	InboxTabGameInvites InboxTab = 5
	InboxTabBookmarks   InboxTab = 6
	InboxTabScheduled   InboxTab = 7
)

func (v InboxTab) String() string {
	switch v {
	case InboxTabUnspecified:
		return "UNSPECIFIED"
	case InboxTabMentions:
		return "MENTIONS"
	case InboxTabUnreads:
		return "UNREADS"
	case InboxTabTodos:
		return "TODOS"
	case InboxTabForYou:
		return "FOR_YOU"
	case InboxTabGameInvites:
		return "GAME_INVITES"
	case InboxTabBookmarks:
		return "BOOKMARKS"
	case InboxTabScheduled:
		return "SCHEDULED"
	}
	return strconv.FormatInt(int64(v), 10)
}

type IntegrationExpireBehavior int64

const (
	IntegrationExpireBehaviorRemoveRole IntegrationExpireBehavior = 0
	IntegrationExpireBehaviorKick       IntegrationExpireBehavior = 1
)

func (v IntegrationExpireBehavior) String() string {
	switch v {
	case IntegrationExpireBehaviorRemoveRole:
		return "REMOVE_ROLE"
	case IntegrationExpireBehaviorKick:
		return "KICK"
	}
	return strconv.FormatInt(int64(v), 10)
}

type IntegrationType string

const (
	IntegrationTypeTwitchIntegration                   IntegrationType = "twitch"
	IntegrationTypeYoutubeIntegration                  IntegrationType = "youtube"
	IntegrationTypeDiscordApplicationIntegration       IntegrationType = "discord"
	IntegrationTypeInternalRoleSubscriptionIntegration IntegrationType = "guild_subscription"
)

type InteractionCallbackType int64

const (
	InteractionCallbackTypePong                                 InteractionCallbackType = 1
	InteractionCallbackTypeAcknowledge                          InteractionCallbackType = 2
	InteractionCallbackTypeChannelMessage                       InteractionCallbackType = 3
	InteractionCallbackTypeChannelMessageWithSource             InteractionCallbackType = 4
	InteractionCallbackTypeDeferredChannelMessageWithSource     InteractionCallbackType = 5
	InteractionCallbackTypeDeferredUpdateMessage                InteractionCallbackType = 6
	InteractionCallbackTypeUpdateMessage                        InteractionCallbackType = 7
	InteractionCallbackTypeApplicationCommandAutocompleteResult InteractionCallbackType = 8
	InteractionCallbackTypeModal                                InteractionCallbackType = 9
	InteractionCallbackTypePremiumRequired                      InteractionCallbackType = 10
	InteractionCallbackTypeIframeModal                          InteractionCallbackType = 11
	InteractionCallbackTypeLaunchActivity                       InteractionCallbackType = 12
	InteractionCallbackTypeSocialLayerSKUPurchaseEligibility    InteractionCallbackType = 13
)

func (v InteractionCallbackType) String() string {
	switch v {
	case InteractionCallbackTypePong:
		return "PONG"
	case InteractionCallbackTypeAcknowledge:
		return "ACKNOWLEDGE"
	case InteractionCallbackTypeChannelMessage:
		return "CHANNEL_MESSAGE"
	case InteractionCallbackTypeChannelMessageWithSource:
		return "CHANNEL_MESSAGE_WITH_SOURCE"
	case InteractionCallbackTypeDeferredChannelMessageWithSource:
		return "DEFERRED_CHANNEL_MESSAGE_WITH_SOURCE"
	case InteractionCallbackTypeDeferredUpdateMessage:
		return "DEFERRED_UPDATE_MESSAGE"
	case InteractionCallbackTypeUpdateMessage:
		return "UPDATE_MESSAGE"
	case InteractionCallbackTypeApplicationCommandAutocompleteResult:
		return "APPLICATION_COMMAND_AUTOCOMPLETE_RESULT"
	case InteractionCallbackTypeModal:
		return "MODAL"
	case InteractionCallbackTypePremiumRequired:
		return "PREMIUM_REQUIRED"
	case InteractionCallbackTypeIframeModal:
		return "IFRAME_MODAL"
	case InteractionCallbackTypeLaunchActivity:
		return "LAUNCH_ACTIVITY"
	case InteractionCallbackTypeSocialLayerSKUPurchaseEligibility:
		return "SOCIAL_LAYER_SKU_PURCHASE_ELIGIBILITY"
	}
	return strconv.FormatInt(int64(v), 10)
}

type InteractionContextType int64

const (
	InteractionContextTypeGuild          InteractionContextType = 0
	InteractionContextTypeBotDM          InteractionContextType = 1
	InteractionContextTypePrivateChannel InteractionContextType = 2
)

func (v InteractionContextType) String() string {
	switch v {
	case InteractionContextTypeGuild:
		return "GUILD"
	case InteractionContextTypeBotDM:
		return "BOT_DM"
	case InteractionContextTypePrivateChannel:
		return "PRIVATE_CHANNEL"
	}
	return strconv.FormatInt(int64(v), 10)
}

type InteractionCreationLocation string

const (
	InteractionCreationLocationDiscovery                          InteractionCreationLocation = "discovery"
	InteractionCreationLocationSuggestion                         InteractionCreationLocation = "suggestion"
	InteractionCreationLocationMention                            InteractionCreationLocation = "mention"
	InteractionCreationLocationPaste                              InteractionCreationLocation = "paste"
	InteractionCreationLocationRecall                             InteractionCreationLocation = "recall"
	InteractionCreationLocationPopularCommands                    InteractionCreationLocation = "popular_commands"
	InteractionCreationLocationMjChatBar                          InteractionCreationLocation = "mj_chat_bar"
	InteractionCreationLocationQuery                              InteractionCreationLocation = "query"
	InteractionCreationLocationSlashUI                            InteractionCreationLocation = "slash_ui"
	InteractionCreationLocationAppLauncher                        InteractionCreationLocation = "app_launcher"
	InteractionCreationLocationAppLauncherHome                    InteractionCreationLocation = "app_launcher_home"
	InteractionCreationLocationAppLauncherHomeSearch              InteractionCreationLocation = "app_launcher_home_search"
	InteractionCreationLocationAppLauncherListViewAll             InteractionCreationLocation = "app_launcher_list_view_all"
	InteractionCreationLocationAppLauncherApplicationView         InteractionCreationLocation = "app_launcher_application_view"
	InteractionCreationLocationAppLauncherApplicationViewFrecent  InteractionCreationLocation = "app_launcher_application_view_frecent"
	InteractionCreationLocationAppLauncherApplicationViewMoreMenu InteractionCreationLocation = "app_launcher_application_view_more_menu"
	InteractionCreationLocationAppLauncherSlashSearch             InteractionCreationLocation = "app_launcher_slash_search"
	InteractionCreationLocationAppLauncherFrecentsViewAll         InteractionCreationLocation = "app_launcher_frecents_view_all"
	InteractionCreationLocationImageRecsMenu                      InteractionCreationLocation = "image_recs_menu"
	InteractionCreationLocationImageRecsSubmenu                   InteractionCreationLocation = "image_recs_submenu"
	InteractionCreationLocationActivityInstanceEmbed              InteractionCreationLocation = "activity_instance_embed"
	InteractionCreationLocationActivityBookmarkEmbed              InteractionCreationLocation = "activity_bookmark_embed"
	InteractionCreationLocationActivitiesMiniShelf                InteractionCreationLocation = "activities_mini_shelf"
	InteractionCreationLocationVcTileActivitySuggestion           InteractionCreationLocation = "vc_tile_activity_suggestion"
	InteractionCreationLocationAppDMsEntryPointCommandButton      InteractionCreationLocation = "app_dms_entry_point_command_button"
)

type InteractionCreationSectionName string

const (
	InteractionCreationSectionNameSearch           InteractionCreationSectionName = "search"
	InteractionCreationSectionNameRecent           InteractionCreationSectionName = "recent"
	InteractionCreationSectionNameInstalled        InteractionCreationSectionName = "installed"
	InteractionCreationSectionNameActivities       InteractionCreationSectionName = "activities"
	InteractionCreationSectionNameRecentApps       InteractionCreationSectionName = "recent_apps"
	InteractionCreationSectionNameRecentCommands   InteractionCreationSectionName = "recent_commands"
	InteractionCreationSectionNameNewToApps        InteractionCreationSectionName = "new_to_apps"
	InteractionCreationSectionNameAppsInThisServer InteractionCreationSectionName = "apps_in_this_server"
)

type InteractionCreationSourceType string

const (
	InteractionCreationSourceTypeNone  InteractionCreationSourceType = "NONE"
	InteractionCreationSourceTypeText  InteractionCreationSourceType = "TEXT"
	InteractionCreationSourceTypeVoice InteractionCreationSourceType = "VOICE"
)

type InteractionFailureReason int64

const (
	InteractionFailureReasonUnknown                                                InteractionFailureReason = 1
	InteractionFailureReasonTimeout                                                InteractionFailureReason = 2
	InteractionFailureReasonActivityLaunchUnknownApplication                       InteractionFailureReason = 3
	InteractionFailureReasonActivityLaunchUnknownChannel                           InteractionFailureReason = 4
	InteractionFailureReasonActivityLaunchUnknownGuild                             InteractionFailureReason = 5
	InteractionFailureReasonActivityLaunchInvalidPlatform                          InteractionFailureReason = 6
	InteractionFailureReasonActivityLaunchNotInExperiment                          InteractionFailureReason = 7
	InteractionFailureReasonActivityLaunchInvalidChannelType                       InteractionFailureReason = 8
	InteractionFailureReasonActivityLaunchInvalidChannelNoAFK                      InteractionFailureReason = 9
	InteractionFailureReasonActivityLaunchInvalidDevPreviewGuildSize               InteractionFailureReason = 10
	InteractionFailureReasonActivityLaunchInvalidUserAgeGate                       InteractionFailureReason = 11
	InteractionFailureReasonActivityLaunchInvalidUserVerificationLevel             InteractionFailureReason = 12
	InteractionFailureReasonActivityLaunchInvalidUserPermissions                   InteractionFailureReason = 13
	InteractionFailureReasonActivityLaunchInvalidConfigurationNotEmbedded          InteractionFailureReason = 14
	InteractionFailureReasonActivityLaunchInvalidConfigurationPlatformNotSupported InteractionFailureReason = 15
	InteractionFailureReasonActivityLaunchInvalidConfigurationPlatformNotReleased  InteractionFailureReason = 16
	InteractionFailureReasonActivityLaunchFailedToLaunch                           InteractionFailureReason = 17
	InteractionFailureReasonActivityLaunchInvalidUserNoAccessToActivity            InteractionFailureReason = 18
	InteractionFailureReasonActivityLaunchInvalidLocationType                      InteractionFailureReason = 19
	InteractionFailureReasonActivityLaunchInvalidUserRegionForApplication          InteractionFailureReason = 20
)

func (v InteractionFailureReason) String() string {
	switch v {
	case InteractionFailureReasonUnknown:
		return "UNKNOWN"
	case InteractionFailureReasonTimeout:
		return "TIMEOUT"
	case InteractionFailureReasonActivityLaunchUnknownApplication:
		return "ACTIVITY_LAUNCH_UNKNOWN_APPLICATION"
	case InteractionFailureReasonActivityLaunchUnknownChannel:
		return "ACTIVITY_LAUNCH_UNKNOWN_CHANNEL"
	case InteractionFailureReasonActivityLaunchUnknownGuild:
		return "ACTIVITY_LAUNCH_UNKNOWN_GUILD"
	case InteractionFailureReasonActivityLaunchInvalidPlatform:
		return "ACTIVITY_LAUNCH_INVALID_PLATFORM"
	case InteractionFailureReasonActivityLaunchNotInExperiment:
		return "ACTIVITY_LAUNCH_NOT_IN_EXPERIMENT"
	case InteractionFailureReasonActivityLaunchInvalidChannelType:
		return "ACTIVITY_LAUNCH_INVALID_CHANNEL_TYPE"
	case InteractionFailureReasonActivityLaunchInvalidChannelNoAFK:
		return "ACTIVITY_LAUNCH_INVALID_CHANNEL_NO_AFK"
	case InteractionFailureReasonActivityLaunchInvalidDevPreviewGuildSize:
		return "ACTIVITY_LAUNCH_INVALID_DEV_PREVIEW_GUILD_SIZE"
	case InteractionFailureReasonActivityLaunchInvalidUserAgeGate:
		return "ACTIVITY_LAUNCH_INVALID_USER_AGE_GATE"
	case InteractionFailureReasonActivityLaunchInvalidUserVerificationLevel:
		return "ACTIVITY_LAUNCH_INVALID_USER_VERIFICATION_LEVEL"
	case InteractionFailureReasonActivityLaunchInvalidUserPermissions:
		return "ACTIVITY_LAUNCH_INVALID_USER_PERMISSIONS"
	case InteractionFailureReasonActivityLaunchInvalidConfigurationNotEmbedded:
		return "ACTIVITY_LAUNCH_INVALID_CONFIGURATION_NOT_EMBEDDED"
	case InteractionFailureReasonActivityLaunchInvalidConfigurationPlatformNotSupported:
		return "ACTIVITY_LAUNCH_INVALID_CONFIGURATION_PLATFORM_NOT_SUPPORTED"
	case InteractionFailureReasonActivityLaunchInvalidConfigurationPlatformNotReleased:
		return "ACTIVITY_LAUNCH_INVALID_CONFIGURATION_PLATFORM_NOT_RELEASED"
	case InteractionFailureReasonActivityLaunchFailedToLaunch:
		return "ACTIVITY_LAUNCH_FAILED_TO_LAUNCH"
	case InteractionFailureReasonActivityLaunchInvalidUserNoAccessToActivity:
		return "ACTIVITY_LAUNCH_INVALID_USER_NO_ACCESS_TO_ACTIVITY"
	case InteractionFailureReasonActivityLaunchInvalidLocationType:
		return "ACTIVITY_LAUNCH_INVALID_LOCATION_TYPE"
	case InteractionFailureReasonActivityLaunchInvalidUserRegionForApplication:
		return "ACTIVITY_LAUNCH_INVALID_USER_REGION_FOR_APPLICATION"
	}
	return strconv.FormatInt(int64(v), 10)
}

type InteractionType int64

const (
	InteractionTypePing                              InteractionType = 1
	InteractionTypeApplicationCommand                InteractionType = 2
	InteractionTypeMessageComponent                  InteractionType = 3
	InteractionTypeApplicationCommandAutocomplete    InteractionType = 4
	InteractionTypeModalSubmit                       InteractionType = 5
	InteractionTypeSocialLayerSKUPurchaseEligibility InteractionType = 6
)

func (v InteractionType) String() string {
	switch v {
	case InteractionTypePing:
		return "PING"
	case InteractionTypeApplicationCommand:
		return "APPLICATION_COMMAND"
	case InteractionTypeMessageComponent:
		return "MESSAGE_COMPONENT"
	case InteractionTypeApplicationCommandAutocomplete:
		return "APPLICATION_COMMAND_AUTOCOMPLETE"
	case InteractionTypeModalSubmit:
		return "MODAL_SUBMIT"
	case InteractionTypeSocialLayerSKUPurchaseEligibility:
		return "SOCIAL_LAYER_SKU_PURCHASE_ELIGIBILITY"
	}
	return strconv.FormatInt(int64(v), 10)
}

type InternalGuildRestriction int64

const (
	InternalGuildRestrictionJoinAll          InternalGuildRestriction = 1
	InternalGuildRestrictionJoinExternalOnly InternalGuildRestriction = 2
	InternalGuildRestrictionJoinInternalOnly InternalGuildRestriction = 3
)

func (v InternalGuildRestriction) String() string {
	switch v {
	case InternalGuildRestrictionJoinAll:
		return "JOIN_ALL"
	case InternalGuildRestrictionJoinExternalOnly:
		return "JOIN_EXTERNAL_ONLY"
	case InternalGuildRestrictionJoinInternalOnly:
		return "JOIN_INTERNAL_ONLY"
	}
	return strconv.FormatInt(int64(v), 10)
}

type InviteFlags uint64

const (
	InviteFlagsIsGuestInvite       InviteFlags = 1 << 0
	InviteFlagsIsViewed            InviteFlags = 1 << 1
	InviteFlagsIsEnhanced          InviteFlags = 1 << 2
	InviteFlagsIsApplicationBypass InviteFlags = 1 << 3
)

func (v InviteFlags) String() string {
	switch v {
	case InviteFlagsIsGuestInvite:
		return "IS_GUEST_INVITE"
	case InviteFlagsIsViewed:
		return "IS_VIEWED"
	case InviteFlagsIsEnhanced:
		return "IS_ENHANCED"
	case InviteFlagsIsApplicationBypass:
		return "IS_APPLICATION_BYPASS"
	}
	return strconv.FormatUint(uint64(v), 10)
}

type InviteTargetType int64

const (
	InviteTargetTypeStream              InviteTargetType = 1
	InviteTargetTypeEmbeddedApplication InviteTargetType = 2
	InviteTargetTypeRoleSubscriptions   InviteTargetType = 3
	InviteTargetTypeCreatorPage         InviteTargetType = 4
	InviteTargetTypeLobby               InviteTargetType = 5
)

func (v InviteTargetType) String() string {
	switch v {
	case InviteTargetTypeStream:
		return "STREAM"
	case InviteTargetTypeEmbeddedApplication:
		return "EMBEDDED_APPLICATION"
	case InviteTargetTypeRoleSubscriptions:
		return "ROLE_SUBSCRIPTIONS"
	case InviteTargetTypeCreatorPage:
		return "CREATOR_PAGE"
	case InviteTargetTypeLobby:
		return "LOBBY"
	}
	return strconv.FormatInt(int64(v), 10)
}

type InviteTargetUsersJobStatus int64

const (
	InviteTargetUsersJobStatusUnspecified InviteTargetUsersJobStatus = 0
	InviteTargetUsersJobStatusProcessing  InviteTargetUsersJobStatus = 1
	InviteTargetUsersJobStatusCompleted   InviteTargetUsersJobStatus = 2
	InviteTargetUsersJobStatusFailed      InviteTargetUsersJobStatus = 3
)

func (v InviteTargetUsersJobStatus) String() string {
	switch v {
	case InviteTargetUsersJobStatusUnspecified:
		return "UNSPECIFIED"
	case InviteTargetUsersJobStatusProcessing:
		return "PROCESSING"
	case InviteTargetUsersJobStatusCompleted:
		return "COMPLETED"
	case InviteTargetUsersJobStatusFailed:
		return "FAILED"
	}
	return strconv.FormatInt(int64(v), 10)
}

type InviteTitleVariant int64

const (
	InviteTitleVariantSenderOnly InviteTitleVariant = 1
	InviteTitleVariantEmoji      InviteTitleVariant = 2
)

func (v InviteTitleVariant) String() string {
	switch v {
	case InviteTitleVariantSenderOnly:
		return "SENDER_ONLY"
	case InviteTitleVariantEmoji:
		return "EMOJI"
	}
	return strconv.FormatInt(int64(v), 10)
}

type InviteType int64

const (
	InviteTypeGuild   InviteType = 0
	InviteTypeGroupDM InviteType = 1
	InviteTypeFriend  InviteType = 2
)

func (v InviteType) String() string {
	switch v {
	case InviteTypeGuild:
		return "GUILD"
	case InviteTypeGroupDM:
		return "GROUP_DM"
	case InviteTypeFriend:
		return "FRIEND"
	}
	return strconv.FormatInt(int64(v), 10)
}

type InvoiceDiscountType int64

const (
	InvoiceDiscountTypeSubscriptionPlan              InvoiceDiscountType = 1
	InvoiceDiscountTypeEntitlement                   InvoiceDiscountType = 2
	InvoiceDiscountTypePremiumLegacyUpgradePromotion InvoiceDiscountType = 3
	InvoiceDiscountTypePremiumTrial                  InvoiceDiscountType = 4
	InvoiceDiscountTypeDefault                       InvoiceDiscountType = 5
)

func (v InvoiceDiscountType) String() string {
	switch v {
	case InvoiceDiscountTypeSubscriptionPlan:
		return "SUBSCRIPTION_PLAN"
	case InvoiceDiscountTypeEntitlement:
		return "ENTITLEMENT"
	case InvoiceDiscountTypePremiumLegacyUpgradePromotion:
		return "PREMIUM_LEGACY_UPGRADE_PROMOTION"
	case InvoiceDiscountTypePremiumTrial:
		return "PREMIUM_TRIAL"
	case InvoiceDiscountTypeDefault:
		return "DEFAULT"
	}
	return strconv.FormatInt(int64(v), 10)
}

type InvoiceStatus int64

const (
	InvoiceStatusOpen          InvoiceStatus = 1
	InvoiceStatusPaid          InvoiceStatus = 2
	InvoiceStatusVoid          InvoiceStatus = 3
	InvoiceStatusUncollectible InvoiceStatus = 4
)

func (v InvoiceStatus) String() string {
	switch v {
	case InvoiceStatusOpen:
		return "OPEN"
	case InvoiceStatusPaid:
		return "PAID"
	case InvoiceStatusVoid:
		return "VOID"
	case InvoiceStatusUncollectible:
		return "UNCOLLECTIBLE"
	}
	return strconv.FormatInt(int64(v), 10)
}

type IpcOpcodes int64

const (
	IpcOpcodesHandshake IpcOpcodes = 0
	IpcOpcodesFrame     IpcOpcodes = 1
	IpcOpcodesClose     IpcOpcodes = 2
	IpcOpcodesPing      IpcOpcodes = 3
	IpcOpcodesPong      IpcOpcodes = 4
)

func (v IpcOpcodes) String() string {
	switch v {
	case IpcOpcodesHandshake:
		return "HANDSHAKE"
	case IpcOpcodesFrame:
		return "FRAME"
	case IpcOpcodesClose:
		return "CLOSE"
	case IpcOpcodesPing:
		return "PING"
	case IpcOpcodesPong:
		return "PONG"
	}
	return strconv.FormatInt(int64(v), 10)
}

type JoinIntent int64

const (
	JoinIntentPlay     JoinIntent = 0
	JoinIntentSpectate JoinIntent = 1
)

func (v JoinIntent) String() string {
	switch v {
	case JoinIntentPlay:
		return "PLAY"
	case JoinIntentSpectate:
		return "SPECTATE"
	}
	return strconv.FormatInt(int64(v), 10)
}

type JoinSourceType int64

const (
	JoinSourceTypeUnspecified                         JoinSourceType = 0
	JoinSourceTypeBot                                 JoinSourceType = 1
	JoinSourceTypeIntegration                         JoinSourceType = 2
	JoinSourceTypeDiscovery                           JoinSourceType = 3
	JoinSourceTypeHub                                 JoinSourceType = 4
	JoinSourceTypeInvite                              JoinSourceType = 5
	JoinSourceTypeVanityURL                           JoinSourceType = 6
	JoinSourceTypeManualMemberVerification            JoinSourceType = 7
	JoinSourceTypeSocialLayerIntegrationLinkedChannel JoinSourceType = 8
)

func (v JoinSourceType) String() string {
	switch v {
	case JoinSourceTypeUnspecified:
		return "UNSPECIFIED"
	case JoinSourceTypeBot:
		return "BOT"
	case JoinSourceTypeIntegration:
		return "INTEGRATION"
	case JoinSourceTypeDiscovery:
		return "DISCOVERY"
	case JoinSourceTypeHub:
		return "HUB"
	case JoinSourceTypeInvite:
		return "INVITE"
	case JoinSourceTypeVanityURL:
		return "VANITY_URL"
	case JoinSourceTypeManualMemberVerification:
		return "MANUAL_MEMBER_VERIFICATION"
	case JoinSourceTypeSocialLayerIntegrationLinkedChannel:
		return "SOCIAL_LAYER_INTEGRATION_LINKED_CHANNEL"
	}
	return strconv.FormatInt(int64(v), 10)
}

type LaunchPadMode int64

const (
	LaunchPadModeDisabled          LaunchPadMode = 0
	LaunchPadModeGestureFullScreen LaunchPadMode = 1
	LaunchPadModeGestureRightEdge  LaunchPadMode = 2
	LaunchPadModePullTab           LaunchPadMode = 3
)

func (v LaunchPadMode) String() string {
	switch v {
	case LaunchPadModeDisabled:
		return "DISABLED"
	case LaunchPadModeGestureFullScreen:
		return "GESTURE_FULL_SCREEN"
	case LaunchPadModeGestureRightEdge:
		return "GESTURE_RIGHT_EDGE"
	case LaunchPadModePullTab:
		return "PULL_TAB"
	}
	return strconv.FormatInt(int64(v), 10)
}

type LayoutDefinitionKey string

const (
	LayoutDefinitionKeyActivityAccessoryStat     LayoutDefinitionKey = "activity_accessory_stat"
	LayoutDefinitionKeyAddWidgetPreviewContained LayoutDefinitionKey = "add_widget_preview_contained"
	LayoutDefinitionKeyAddWidgetPreviewHero      LayoutDefinitionKey = "add_widget_preview_hero"
	LayoutDefinitionKeyMiniProfileContainedStat  LayoutDefinitionKey = "mini_profile_contained_stat"
	LayoutDefinitionKeyMiniProfileHeroStat       LayoutDefinitionKey = "mini_profile_hero_stat"
	LayoutDefinitionKeyWidgetBottomCollection    LayoutDefinitionKey = "widget_bottom_collection"
	LayoutDefinitionKeyWidgetBottomProgress      LayoutDefinitionKey = "widget_bottom_progress"
	LayoutDefinitionKeyWidgetBottomStats         LayoutDefinitionKey = "widget_bottom_stats"
	LayoutDefinitionKeyWidgetTopContained        LayoutDefinitionKey = "widget_top_contained"
	LayoutDefinitionKeyWidgetTopHero             LayoutDefinitionKey = "widget_top_hero"
)

type LayoutMode int64

const (
	LayoutModeFocused LayoutMode = 0
	LayoutModePip     LayoutMode = 1
	LayoutModeGrid    LayoutMode = 2
)

func (v LayoutMode) String() string {
	switch v {
	case LayoutModeFocused:
		return "FOCUSED"
	case LayoutModePip:
		return "PIP"
	case LayoutModeGrid:
		return "GRID"
	}
	return strconv.FormatInt(int64(v), 10)
}

type LinkStatus int64

const (
	LinkStatusN1 LinkStatus = 1
	LinkStatusN2 LinkStatus = 2
	LinkStatusN3 LinkStatus = 3
	LinkStatusN4 LinkStatus = 4
	LinkStatusN5 LinkStatus = 5
)

func (v LinkStatus) String() string {
	switch v {
	case LinkStatusN1:
		return "1"
	case LinkStatusN2:
		return "2"
	case LinkStatusN3:
		return "3"
	case LinkStatusN4:
		return "4"
	case LinkStatusN5:
		return "5"
	}
	return strconv.FormatInt(int64(v), 10)
}

type LinkType int64

const (
	LinkTypeN1 LinkType = 1
	LinkTypeN2 LinkType = 2
)

func (v LinkType) String() string {
	switch v {
	case LinkTypeN1:
		return "1"
	case LinkTypeN2:
		return "2"
	}
	return strconv.FormatInt(int64(v), 10)
}

type LinkedApplicationType int64

const (
	LinkedApplicationTypeLinked   LinkedApplicationType = 1
	LinkedApplicationTypeOfficial LinkedApplicationType = 2
	LinkedApplicationTypeNvidia   LinkedApplicationType = 3
)

func (v LinkedApplicationType) String() string {
	switch v {
	case LinkedApplicationTypeLinked:
		return "LINKED"
	case LinkedApplicationTypeOfficial:
		return "OFFICIAL"
	case LinkedApplicationTypeNvidia:
		return "NVIDIA"
	}
	return strconv.FormatInt(int64(v), 10)
}

type ListOfCapabilities uint64

const (
	ListOfCapabilitiesLazyUserNotes                      ListOfCapabilities = 1 << 0
	ListOfCapabilitiesNoAffineUserIDs                    ListOfCapabilities = 1 << 1
	ListOfCapabilitiesVersionedReadStates                ListOfCapabilities = 1 << 2
	ListOfCapabilitiesVersionedUserGuildSettings         ListOfCapabilities = 1 << 3
	ListOfCapabilitiesDedupeUserObjects                  ListOfCapabilities = 1 << 4
	ListOfCapabilitiesPrioritizedReadyPayload            ListOfCapabilities = 1 << 5
	ListOfCapabilitiesMultipleGuildExperimentPopulations ListOfCapabilities = 1 << 6
	ListOfCapabilitiesNonChannelReadStates               ListOfCapabilities = 1 << 7
	ListOfCapabilitiesAuthTokenRefresh                   ListOfCapabilities = 1 << 8
	ListOfCapabilitiesUserSettingsProto                  ListOfCapabilities = 1 << 9
	ListOfCapabilitiesClientStateV2                      ListOfCapabilities = 1 << 10
	ListOfCapabilitiesPassiveGuildUpdate                 ListOfCapabilities = 1 << 11
	ListOfCapabilitiesAutoCallConnect                    ListOfCapabilities = 1 << 12
	ListOfCapabilitiesDebounceMessageReactions           ListOfCapabilities = 1 << 13
	ListOfCapabilitiesPassiveGuildUpdateV2               ListOfCapabilities = 1 << 14
	ListOfCapabilitiesChannelObfuscation                 ListOfCapabilities = 1 << 15
	ListOfCapabilitiesAutoLobbyConnect                   ListOfCapabilities = 1 << 16
)

func (v ListOfCapabilities) String() string {
	switch v {
	case ListOfCapabilitiesLazyUserNotes:
		return "LAZY_USER_NOTES"
	case ListOfCapabilitiesNoAffineUserIDs:
		return "NO_AFFINE_USER_IDS"
	case ListOfCapabilitiesVersionedReadStates:
		return "VERSIONED_READ_STATES"
	case ListOfCapabilitiesVersionedUserGuildSettings:
		return "VERSIONED_USER_GUILD_SETTINGS"
	case ListOfCapabilitiesDedupeUserObjects:
		return "DEDUPE_USER_OBJECTS"
	case ListOfCapabilitiesPrioritizedReadyPayload:
		return "PRIORITIZED_READY_PAYLOAD"
	case ListOfCapabilitiesMultipleGuildExperimentPopulations:
		return "MULTIPLE_GUILD_EXPERIMENT_POPULATIONS"
	case ListOfCapabilitiesNonChannelReadStates:
		return "NON_CHANNEL_READ_STATES"
	case ListOfCapabilitiesAuthTokenRefresh:
		return "AUTH_TOKEN_REFRESH"
	case ListOfCapabilitiesUserSettingsProto:
		return "USER_SETTINGS_PROTO"
	case ListOfCapabilitiesClientStateV2:
		return "CLIENT_STATE_V2"
	case ListOfCapabilitiesPassiveGuildUpdate:
		return "PASSIVE_GUILD_UPDATE"
	case ListOfCapabilitiesAutoCallConnect:
		return "AUTO_CALL_CONNECT"
	case ListOfCapabilitiesDebounceMessageReactions:
		return "DEBOUNCE_MESSAGE_REACTIONS"
	case ListOfCapabilitiesPassiveGuildUpdateV2:
		return "PASSIVE_GUILD_UPDATE_V2"
	case ListOfCapabilitiesChannelObfuscation:
		return "CHANNEL_OBFUSCATION"
	case ListOfCapabilitiesAutoLobbyConnect:
		return "AUTO_LOBBY_CONNECT"
	}
	return strconv.FormatUint(uint64(v), 10)
}

type LobbyDeletionReason string

const (
	LobbyDeletionReasonDeleted LobbyDeletionReason = "deleted"
	LobbyDeletionReasonRemoved LobbyDeletionReason = "removed"
)

type LobbyFlags uint64

const (
	LobbyFlagsRequireApplicationAuthorization LobbyFlags = 1 << 0
)

func (v LobbyFlags) String() string {
	switch v {
	case LobbyFlagsRequireApplicationAuthorization:
		return "REQUIRE_APPLICATION_AUTHORIZATION"
	}
	return strconv.FormatUint(uint64(v), 10)
}

type LobbyMemberFlags uint64

const (
	LobbyMemberFlagsCanLinkLobby LobbyMemberFlags = 1 << 0
)

func (v LobbyMemberFlags) String() string {
	switch v {
	case LobbyMemberFlagsCanLinkLobby:
		return "CAN_LINK_LOBBY"
	}
	return strconv.FormatUint(uint64(v), 10)
}

type LogLevel string

const (
	LogLevelLog   LogLevel = "log"
	LogLevelWarn  LogLevel = "warn"
	LogLevelDebug LogLevel = "debug"
	LogLevelInfo  LogLevel = "info"
	LogLevelError LogLevel = "error"
)

type LoginRequiredActionType string

const (
	LoginRequiredActionTypeUpdatePassword LoginRequiredActionType = "update_password"
)

type LoginSource string

const (
	LoginSourceGift                    LoginSource = "gift"
	LoginSourceGuildTemplate           LoginSource = "guild_template"
	LoginSourceGuildInvite             LoginSource = "guild_invite"
	LoginSourceDMInvite                LoginSource = "dm_invite"
	LoginSourceFriendInvite            LoginSource = "friend_invite"
	LoginSourceRoleSubscription        LoginSource = "role_subscription"
	LoginSourceRoleSubscriptionSetting LoginSource = "role_subscription_setting"
)

type MFALevel int64

const (
	MFALevelNone     MFALevel = 0
	MFALevelElevated MFALevel = 1
)

func (v MFALevel) String() string {
	switch v {
	case MFALevelNone:
		return "NONE"
	case MFALevelElevated:
		return "ELEVATED"
	}
	return strconv.FormatInt(int64(v), 10)
}

type MediaType string

const (
	MediaTypeVideo  MediaType = "video"
	MediaTypeScreen MediaType = "screen"
	MediaTypeTest   MediaType = "test"
)

type MemberSortType int64

const (
	MemberSortTypeJoinedAtDesc MemberSortType = 1
	MemberSortTypeJoinedAtAsc  MemberSortType = 2
	MemberSortTypeUserIDDesc   MemberSortType = 3
	MemberSortTypeUserIDAsc    MemberSortType = 4
)

func (v MemberSortType) String() string {
	switch v {
	case MemberSortTypeJoinedAtDesc:
		return "JOINED_AT_DESC"
	case MemberSortTypeJoinedAtAsc:
		return "JOINED_AT_ASC"
	case MemberSortTypeUserIDDesc:
		return "USER_ID_DESC"
	case MemberSortTypeUserIDAsc:
		return "USER_ID_ASC"
	}
	return strconv.FormatInt(int64(v), 10)
}

type MemberVerificationFormFieldType string

const (
	MemberVerificationFormFieldTypeTerms          MemberVerificationFormFieldType = "TERMS"
	MemberVerificationFormFieldTypeTextInput      MemberVerificationFormFieldType = "TEXT_INPUT"
	MemberVerificationFormFieldTypeParagraph      MemberVerificationFormFieldType = "PARAGRAPH"
	MemberVerificationFormFieldTypeMultipleChoice MemberVerificationFormFieldType = "MULTIPLE_CHOICE"
	MemberVerificationFormFieldTypeVerification   MemberVerificationFormFieldType = "VERIFICATION"
)

type MessageFlags uint64

const (
	MessageFlagsCrossposted                      MessageFlags = 1 << 0
	MessageFlagsIsCrosspost                      MessageFlags = 1 << 1
	MessageFlagsSuppressEmbeds                   MessageFlags = 1 << 2
	MessageFlagsSourceMessageDeleted             MessageFlags = 1 << 3
	MessageFlagsUrgent                           MessageFlags = 1 << 4
	MessageFlagsHasThread                        MessageFlags = 1 << 5
	MessageFlagsEphemeral                        MessageFlags = 1 << 6
	MessageFlagsLoading                          MessageFlags = 1 << 7
	MessageFlagsFailedToMentionSomeRolesInThread MessageFlags = 1 << 8
	MessageFlagsGuildFeedHidden                  MessageFlags = 1 << 9
	MessageFlagsShouldShowLinkNotDiscordWarning  MessageFlags = 1 << 10
	MessageFlagsSuppressNotifications            MessageFlags = 1 << 12
	MessageFlagsIsVoiceMessage                   MessageFlags = 1 << 13
	MessageFlagsHasSnapshot                      MessageFlags = 1 << 14
	MessageFlagsIsComponentsV2                   MessageFlags = 1 << 15
	MessageFlagsSentBySocialLayerIntegration     MessageFlags = 1 << 16
	MessageFlagsHiddenSuspendedUser              MessageFlags = 1 << 17
	MessageFlagsIsFirstBooster                   MessageFlags = 1 << 18
	MessageFlagsIsGuildOfficial                  MessageFlags = 1 << 19
)

func (v MessageFlags) String() string {
	switch v {
	case MessageFlagsCrossposted:
		return "CROSSPOSTED"
	case MessageFlagsIsCrosspost:
		return "IS_CROSSPOST"
	case MessageFlagsSuppressEmbeds:
		return "SUPPRESS_EMBEDS"
	case MessageFlagsSourceMessageDeleted:
		return "SOURCE_MESSAGE_DELETED"
	case MessageFlagsUrgent:
		return "URGENT"
	case MessageFlagsHasThread:
		return "HAS_THREAD"
	case MessageFlagsEphemeral:
		return "EPHEMERAL"
	case MessageFlagsLoading:
		return "LOADING"
	case MessageFlagsFailedToMentionSomeRolesInThread:
		return "FAILED_TO_MENTION_SOME_ROLES_IN_THREAD"
	case MessageFlagsGuildFeedHidden:
		return "GUILD_FEED_HIDDEN"
	case MessageFlagsShouldShowLinkNotDiscordWarning:
		return "SHOULD_SHOW_LINK_NOT_DISCORD_WARNING"
	case MessageFlagsSuppressNotifications:
		return "SUPPRESS_NOTIFICATIONS"
	case MessageFlagsIsVoiceMessage:
		return "IS_VOICE_MESSAGE"
	case MessageFlagsHasSnapshot:
		return "HAS_SNAPSHOT"
	case MessageFlagsIsComponentsV2:
		return "IS_COMPONENTS_V2"
	case MessageFlagsSentBySocialLayerIntegration:
		return "SENT_BY_SOCIAL_LAYER_INTEGRATION"
	case MessageFlagsHiddenSuspendedUser:
		return "HIDDEN_SUSPENDED_USER"
	case MessageFlagsIsFirstBooster:
		return "IS_FIRST_BOOSTER"
	case MessageFlagsIsGuildOfficial:
		return "IS_GUILD_OFFICIAL"
	}
	return strconv.FormatUint(uint64(v), 10)
}

type MessageNotificationLevel int64

const (
	MessageNotificationLevelAllMessages  MessageNotificationLevel = 0
	MessageNotificationLevelOnlyMentions MessageNotificationLevel = 1
	MessageNotificationLevelNoMessages   MessageNotificationLevel = 2
	MessageNotificationLevelInherit      MessageNotificationLevel = 3
)

func (v MessageNotificationLevel) String() string {
	switch v {
	case MessageNotificationLevelAllMessages:
		return "ALL_MESSAGES"
	case MessageNotificationLevelOnlyMentions:
		return "ONLY_MENTIONS"
	case MessageNotificationLevelNoMessages:
		return "NO_MESSAGES"
	case MessageNotificationLevelInherit:
		return "INHERIT"
	}
	return strconv.FormatInt(int64(v), 10)
}

type MessagePurchaseNotificationType int64

const (
	MessagePurchaseNotificationTypeGuildProduct MessagePurchaseNotificationType = 0
)

func (v MessagePurchaseNotificationType) String() string {
	switch v {
	case MessagePurchaseNotificationTypeGuildProduct:
		return "GUILD_PRODUCT"
	}
	return strconv.FormatInt(int64(v), 10)
}

type MessageReferenceType int64

const (
	MessageReferenceTypeDefault MessageReferenceType = 0
	MessageReferenceTypeForward MessageReferenceType = 1
)

func (v MessageReferenceType) String() string {
	switch v {
	case MessageReferenceTypeDefault:
		return "DEFAULT"
	case MessageReferenceTypeForward:
		return "FORWARD"
	}
	return strconv.FormatInt(int64(v), 10)
}

type MessageType int64

const (
	MessageTypeDefault                                 MessageType = 0
	MessageTypeRecipientAdd                            MessageType = 1
	MessageTypeRecipientRemove                         MessageType = 2
	MessageTypeCall                                    MessageType = 3
	MessageTypeChannelNameChange                       MessageType = 4
	MessageTypeChannelIconChange                       MessageType = 5
	MessageTypeChannelPinnedMessage                    MessageType = 6
	MessageTypeUserJoin                                MessageType = 7
	MessageTypePremiumGuildSubscription                MessageType = 8
	MessageTypePremiumGuildSubscriptionTier1           MessageType = 9
	MessageTypePremiumGuildSubscriptionTier2           MessageType = 10
	MessageTypePremiumGuildSubscriptionTier3           MessageType = 11
	MessageTypeChannelFollowAdd                        MessageType = 12
	MessageTypeGuildStream                             MessageType = 13
	MessageTypeGuildDiscoveryDisqualified              MessageType = 14
	MessageTypeGuildDiscoveryRequalified               MessageType = 15
	MessageTypeGuildDiscoveryGracePeriodInitialWarning MessageType = 16
	MessageTypeGuildDiscoveryGracePeriodFinalWarning   MessageType = 17
	MessageTypeThreadCreated                           MessageType = 18
	MessageTypeReply                                   MessageType = 19
	MessageTypeChatInputCommand                        MessageType = 20
	MessageTypeThreadStarterMessage                    MessageType = 21
	MessageTypeGuildInviteReminder                     MessageType = 22
	MessageTypeContextMenuCommand                      MessageType = 23
	MessageTypeAutoModerationAction                    MessageType = 24
	MessageTypeRoleSubscriptionPurchase                MessageType = 25
	MessageTypeInteractionPremiumUpsell                MessageType = 26
	MessageTypeStageStart                              MessageType = 27
	MessageTypeStageEnd                                MessageType = 28
	MessageTypeStageSpeaker                            MessageType = 29
	MessageTypeStageRaiseHand                          MessageType = 30
	MessageTypeStageTopic                              MessageType = 31
	MessageTypeGuildApplicationPremiumSubscription     MessageType = 32
	MessageTypePrivateChannelIntegrationAdded          MessageType = 33
	MessageTypePrivateChannelIntegrationRemoved        MessageType = 34
	MessageTypePremiumReferral                         MessageType = 35
	MessageTypeGuildIncidentAlertModeEnabled           MessageType = 36
	MessageTypeGuildIncidentAlertModeDisabled          MessageType = 37
	MessageTypeGuildIncidentReportRaid                 MessageType = 38
	MessageTypeGuildIncidentReportFalseAlarm           MessageType = 39
	MessageTypeGuildDeadchatRevivePrompt               MessageType = 40
	MessageTypeCustomGift                              MessageType = 41
	MessageTypeGuildGamingStatsPrompt                  MessageType = 42
	MessageTypePoll                                    MessageType = 43
	MessageTypePurchaseNotification                    MessageType = 44
	MessageTypeVoiceHangoutInvite                      MessageType = 45
	MessageTypePollResult                              MessageType = 46
	MessageTypeChangelog                               MessageType = 47
	MessageTypeNitroNotification                       MessageType = 48
	MessageTypeChannelLinkedToLobby                    MessageType = 49
	MessageTypeGiftingPrompt                           MessageType = 50
	MessageTypeInGameMessageNux                        MessageType = 51
	MessageTypeGuildJoinRequestAcceptNotification      MessageType = 52
	MessageTypeGuildJoinRequestRejectNotification      MessageType = 53
	MessageTypeGuildJoinRequestWithdrawnNotification   MessageType = 54
	MessageTypeHdStreamingUpgraded                     MessageType = 55
	MessageTypeChatWallpaperSet                        MessageType = 56
	MessageTypeChatWallpaperRemove                     MessageType = 57
	MessageTypeReportToModDeletedMessage               MessageType = 58
	MessageTypeReportToModTimeoutUser                  MessageType = 59
	MessageTypeReportToModKickUser                     MessageType = 60
	MessageTypeReportToModBanUser                      MessageType = 61
	MessageTypeReportToModClosedReport                 MessageType = 62
	MessageTypeEmojiAdded                              MessageType = 63
	MessageTypePremiumGroupInvite                      MessageType = 64
	MessageTypeVoiceSession                            MessageType = 65
	MessageTypeGuildBoostUpsell                        MessageType = 66
	MessageTypeFriendRequestAccepted                   MessageType = 67
	MessageTypeMediaMentionMessage                     MessageType = 68
)

func (v MessageType) String() string {
	switch v {
	case MessageTypeDefault:
		return "DEFAULT"
	case MessageTypeRecipientAdd:
		return "RECIPIENT_ADD"
	case MessageTypeRecipientRemove:
		return "RECIPIENT_REMOVE"
	case MessageTypeCall:
		return "CALL"
	case MessageTypeChannelNameChange:
		return "CHANNEL_NAME_CHANGE"
	case MessageTypeChannelIconChange:
		return "CHANNEL_ICON_CHANGE"
	case MessageTypeChannelPinnedMessage:
		return "CHANNEL_PINNED_MESSAGE"
	case MessageTypeUserJoin:
		return "USER_JOIN"
	case MessageTypePremiumGuildSubscription:
		return "PREMIUM_GUILD_SUBSCRIPTION"
	case MessageTypePremiumGuildSubscriptionTier1:
		return "PREMIUM_GUILD_SUBSCRIPTION_TIER_1"
	case MessageTypePremiumGuildSubscriptionTier2:
		return "PREMIUM_GUILD_SUBSCRIPTION_TIER_2"
	case MessageTypePremiumGuildSubscriptionTier3:
		return "PREMIUM_GUILD_SUBSCRIPTION_TIER_3"
	case MessageTypeChannelFollowAdd:
		return "CHANNEL_FOLLOW_ADD"
	case MessageTypeGuildStream:
		return "GUILD_STREAM"
	case MessageTypeGuildDiscoveryDisqualified:
		return "GUILD_DISCOVERY_DISQUALIFIED"
	case MessageTypeGuildDiscoveryRequalified:
		return "GUILD_DISCOVERY_REQUALIFIED"
	case MessageTypeGuildDiscoveryGracePeriodInitialWarning:
		return "GUILD_DISCOVERY_GRACE_PERIOD_INITIAL_WARNING"
	case MessageTypeGuildDiscoveryGracePeriodFinalWarning:
		return "GUILD_DISCOVERY_GRACE_PERIOD_FINAL_WARNING"
	case MessageTypeThreadCreated:
		return "THREAD_CREATED"
	case MessageTypeReply:
		return "REPLY"
	case MessageTypeChatInputCommand:
		return "CHAT_INPUT_COMMAND"
	case MessageTypeThreadStarterMessage:
		return "THREAD_STARTER_MESSAGE"
	case MessageTypeGuildInviteReminder:
		return "GUILD_INVITE_REMINDER"
	case MessageTypeContextMenuCommand:
		return "CONTEXT_MENU_COMMAND"
	case MessageTypeAutoModerationAction:
		return "AUTO_MODERATION_ACTION"
	case MessageTypeRoleSubscriptionPurchase:
		return "ROLE_SUBSCRIPTION_PURCHASE"
	case MessageTypeInteractionPremiumUpsell:
		return "INTERACTION_PREMIUM_UPSELL"
	case MessageTypeStageStart:
		return "STAGE_START"
	case MessageTypeStageEnd:
		return "STAGE_END"
	case MessageTypeStageSpeaker:
		return "STAGE_SPEAKER"
	case MessageTypeStageRaiseHand:
		return "STAGE_RAISE_HAND"
	case MessageTypeStageTopic:
		return "STAGE_TOPIC"
	case MessageTypeGuildApplicationPremiumSubscription:
		return "GUILD_APPLICATION_PREMIUM_SUBSCRIPTION"
	case MessageTypePrivateChannelIntegrationAdded:
		return "PRIVATE_CHANNEL_INTEGRATION_ADDED"
	case MessageTypePrivateChannelIntegrationRemoved:
		return "PRIVATE_CHANNEL_INTEGRATION_REMOVED"
	case MessageTypePremiumReferral:
		return "PREMIUM_REFERRAL"
	case MessageTypeGuildIncidentAlertModeEnabled:
		return "GUILD_INCIDENT_ALERT_MODE_ENABLED"
	case MessageTypeGuildIncidentAlertModeDisabled:
		return "GUILD_INCIDENT_ALERT_MODE_DISABLED"
	case MessageTypeGuildIncidentReportRaid:
		return "GUILD_INCIDENT_REPORT_RAID"
	case MessageTypeGuildIncidentReportFalseAlarm:
		return "GUILD_INCIDENT_REPORT_FALSE_ALARM"
	case MessageTypeGuildDeadchatRevivePrompt:
		return "GUILD_DEADCHAT_REVIVE_PROMPT"
	case MessageTypeCustomGift:
		return "CUSTOM_GIFT"
	case MessageTypeGuildGamingStatsPrompt:
		return "GUILD_GAMING_STATS_PROMPT"
	case MessageTypePoll:
		return "POLL"
	case MessageTypePurchaseNotification:
		return "PURCHASE_NOTIFICATION"
	case MessageTypeVoiceHangoutInvite:
		return "VOICE_HANGOUT_INVITE"
	case MessageTypePollResult:
		return "POLL_RESULT"
	case MessageTypeChangelog:
		return "CHANGELOG"
	case MessageTypeNitroNotification:
		return "NITRO_NOTIFICATION"
	case MessageTypeChannelLinkedToLobby:
		return "CHANNEL_LINKED_TO_LOBBY"
	case MessageTypeGiftingPrompt:
		return "GIFTING_PROMPT"
	case MessageTypeInGameMessageNux:
		return "IN_GAME_MESSAGE_NUX"
	case MessageTypeGuildJoinRequestAcceptNotification:
		return "GUILD_JOIN_REQUEST_ACCEPT_NOTIFICATION"
	case MessageTypeGuildJoinRequestRejectNotification:
		return "GUILD_JOIN_REQUEST_REJECT_NOTIFICATION"
	case MessageTypeGuildJoinRequestWithdrawnNotification:
		return "GUILD_JOIN_REQUEST_WITHDRAWN_NOTIFICATION"
	case MessageTypeHdStreamingUpgraded:
		return "HD_STREAMING_UPGRADED"
	case MessageTypeChatWallpaperSet:
		return "CHAT_WALLPAPER_SET"
	case MessageTypeChatWallpaperRemove:
		return "CHAT_WALLPAPER_REMOVE"
	case MessageTypeReportToModDeletedMessage:
		return "REPORT_TO_MOD_DELETED_MESSAGE"
	case MessageTypeReportToModTimeoutUser:
		return "REPORT_TO_MOD_TIMEOUT_USER"
	case MessageTypeReportToModKickUser:
		return "REPORT_TO_MOD_KICK_USER"
	case MessageTypeReportToModBanUser:
		return "REPORT_TO_MOD_BAN_USER"
	case MessageTypeReportToModClosedReport:
		return "REPORT_TO_MOD_CLOSED_REPORT"
	case MessageTypeEmojiAdded:
		return "EMOJI_ADDED"
	case MessageTypePremiumGroupInvite:
		return "PREMIUM_GROUP_INVITE"
	case MessageTypeVoiceSession:
		return "VOICE_SESSION"
	case MessageTypeGuildBoostUpsell:
		return "GUILD_BOOST_UPSELL"
	case MessageTypeFriendRequestAccepted:
		return "FRIEND_REQUEST_ACCEPTED"
	case MessageTypeMediaMentionMessage:
		return "MEDIA_MENTION_MESSAGE"
	}
	return strconv.FormatInt(int64(v), 10)
}

type NSFWLevel int64

const (
	NSFWLevelDefault       NSFWLevel = 0
	NSFWLevelExplicit      NSFWLevel = 1
	NSFWLevelSafe          NSFWLevel = 2
	NSFWLevelAgeRestricted NSFWLevel = 3
)

func (v NSFWLevel) String() string {
	switch v {
	case NSFWLevelDefault:
		return "DEFAULT"
	case NSFWLevelExplicit:
		return "EXPLICIT"
	case NSFWLevelSafe:
		return "SAFE"
	case NSFWLevelAgeRestricted:
		return "AGE_RESTRICTED"
	}
	return strconv.FormatInt(int64(v), 10)
}

type NameplateColorPalette string

const (
	NameplateColorPaletteNone      NameplateColorPalette = "none"
	NameplateColorPaletteCrimson   NameplateColorPalette = "crimson"
	NameplateColorPaletteBerry     NameplateColorPalette = "berry"
	NameplateColorPaletteSky       NameplateColorPalette = "sky"
	NameplateColorPaletteTeal      NameplateColorPalette = "teal"
	NameplateColorPaletteForest    NameplateColorPalette = "forest"
	NameplateColorPaletteBubbleGum NameplateColorPalette = "bubble_gum"
	NameplateColorPaletteViolet    NameplateColorPalette = "violet"
	NameplateColorPaletteCobalt    NameplateColorPalette = "cobalt"
	NameplateColorPaletteClover    NameplateColorPalette = "clover"
	NameplateColorPaletteLemon     NameplateColorPalette = "lemon"
	NameplateColorPaletteWhite     NameplateColorPalette = "white"
)

type NewMemberActionType int64

const (
	NewMemberActionTypeView NewMemberActionType = 0
	NewMemberActionTypeChat NewMemberActionType = 1
)

func (v NewMemberActionType) String() string {
	switch v {
	case NewMemberActionTypeView:
		return "VIEW"
	case NewMemberActionTypeChat:
		return "CHAT"
	}
	return strconv.FormatInt(int64(v), 10)
}

type NotificationCenterItemDeletionType string

const (
	NotificationCenterItemDeletionTypeMention NotificationCenterItemDeletionType = "mention"
	NotificationCenterItemDeletionTypeRegular NotificationCenterItemDeletionType = "regular"
)

type NotificationCenterItemEnum int64

const (
	NotificationCenterItemEnumUpdateProfile NotificationCenterItemEnum = 0
	NotificationCenterItemEnumFindFriends   NotificationCenterItemEnum = 1
	NotificationCenterItemEnumAddFriend     NotificationCenterItemEnum = 2
	NotificationCenterItemEnumFirstMessage  NotificationCenterItemEnum = 3
)

func (v NotificationCenterItemEnum) String() string {
	switch v {
	case NotificationCenterItemEnumUpdateProfile:
		return "UPDATE_PROFILE"
	case NotificationCenterItemEnumFindFriends:
		return "FIND_FRIENDS"
	case NotificationCenterItemEnumAddFriend:
		return "ADD_FRIEND"
	case NotificationCenterItemEnumFirstMessage:
		return "FIRST_MESSAGE"
	}
	return strconv.FormatInt(int64(v), 10)
}

type NotificationCenterItemType string

const (
	NotificationCenterItemTypeGoLivePush                        NotificationCenterItemType = "go_live_push"
	NotificationCenterItemTypeFriendRequestAccepted             NotificationCenterItemType = "friend_request_accepted"
	NotificationCenterItemTypeFriendRequestPending              NotificationCenterItemType = "friend_request_pending"
	NotificationCenterItemTypeFriendSuggestionCreated           NotificationCenterItemType = "friend_suggestion_created"
	NotificationCenterItemTypeFriendRequestReminder             NotificationCenterItemType = "friend_request_reminder"
	NotificationCenterItemTypeDMFriendNudge                     NotificationCenterItemType = "dm_friend_nudge"
	NotificationCenterItemTypeRecentMention                     NotificationCenterItemType = "recent_mention"
	NotificationCenterItemTypeReplyMention                      NotificationCenterItemType = "reply_mention"
	NotificationCenterItemTypeScheduledGuildEventStarted        NotificationCenterItemType = "scheduled_guild_event_started"
	NotificationCenterItemTypeSystemDemo                        NotificationCenterItemType = "system_demo"
	NotificationCenterItemTypeMissedMessages                    NotificationCenterItemType = "missed_messages"
	NotificationCenterItemTypeTopMessages                       NotificationCenterItemType = "top_messages"
	NotificationCenterItemTypeLifecycleItem                     NotificationCenterItemType = "lifecycle_item"
	NotificationCenterItemTypeTrendingContent                   NotificationCenterItemType = "trending_content"
	NotificationCenterItemTypeReferralProgramEntrypointReminder NotificationCenterItemType = "referral_program_entrypoint_reminder"
	NotificationCenterItemTypePollEnded                         NotificationCenterItemType = "poll_ended"
	NotificationCenterItemTypeGameFriendRequestAccepted         NotificationCenterItemType = "game_friend_request_accepted"
	NotificationCenterItemTypeReactionSent                      NotificationCenterItemType = "reaction_sent"
)

type NotificationChannelGroup string

const (
	NotificationChannelGroupN111Realtime NotificationChannelGroup = "111_realtime"
	NotificationChannelGroupN222Social   NotificationChannelGroup = "222_social"
	NotificationChannelGroupN333Server   NotificationChannelGroup = "333_server"
	NotificationChannelGroupN444Other    NotificationChannelGroup = "444_other"
)

type NotificationSettingsFlags uint64

const (
	NotificationSettingsFlagsUseNewNotifications  NotificationSettingsFlags = 1 << 4
	NotificationSettingsFlagsMentionOnAllMessages NotificationSettingsFlags = 1 << 5
)

func (v NotificationSettingsFlags) String() string {
	switch v {
	case NotificationSettingsFlagsUseNewNotifications:
		return "USE_NEW_NOTIFICATIONS"
	case NotificationSettingsFlagsMentionOnAllMessages:
		return "MENTION_ON_ALL_MESSAGES"
	}
	return strconv.FormatUint(uint64(v), 10)
}

type NotificationSoundPack string

const (
	NotificationSoundPackClassic       NotificationSoundPack = "classic"
	NotificationSoundPackRetro         NotificationSoundPack = "retro"
	NotificationSoundPackBop           NotificationSoundPack = "bop"
	NotificationSoundPackDucky         NotificationSoundPack = "ducky"
	NotificationSoundPackLofi          NotificationSoundPack = "lofi"
	NotificationSoundPackAsmr          NotificationSoundPack = "asmr"
	NotificationSoundPackDiscodo       NotificationSoundPack = "discodo"
	NotificationSoundPackHalloween     NotificationSoundPack = "halloween"
	NotificationSoundPackWinterHoliday NotificationSoundPack = "winter_holiday"
)

type NotificationType int64

const (
	NotificationTypeMessageCreate                                NotificationType = 1
	NotificationTypeCallRing                                     NotificationType = 2
	NotificationTypeRelationshipAdd                              NotificationType = 3
	NotificationTypeFriendSuggestionCreate                       NotificationType = 4
	NotificationTypeApplicationLibraryInstallComplete            NotificationType = 5
	NotificationTypeGuildStreamStart                             NotificationType = 6
	NotificationTypeStageInstanceCreate                          NotificationType = 7
	NotificationTypeGuildScheduledEventUpdate                    NotificationType = 8
	NotificationTypeForumThreadCreated                           NotificationType = 9
	NotificationTypeTopMessagesPush                              NotificationType = 10
	NotificationTypeMissedMessagePush                            NotificationType = 11
	NotificationTypeFriendRequestReminderPush                    NotificationType = 13
	NotificationTypeHomeLifecyclePush                            NotificationType = 14
	NotificationTypePollEndedPush                                NotificationType = 15
	NotificationTypeNudgeNewFriendDMPush                         NotificationType = 16
	NotificationTypeFamilyCenterRequestSend                      NotificationType = 17
	NotificationTypeFamilyCenterRequestAccepted                  NotificationType = 18
	NotificationTypeFamilyCenterRequestDeclined                  NotificationType = 19
	NotificationTypeFamilyCenterDisconnected                     NotificationType = 20
	NotificationTypeGuildJoinRequestApproved                     NotificationType = 21
	NotificationTypeGuildJoinRequestRejected                     NotificationType = 22
	NotificationTypeReactionsPushNotification                    NotificationType = 23
	NotificationTypeRaidDetected                                 NotificationType = 24
	NotificationTypeMentionRaidDetected                          NotificationType = 25
	NotificationTypeDMSpamDetected                               NotificationType = 26
	NotificationTypeSuspiciousSession                            NotificationType = 27
	NotificationTypeNewUserSession                               NotificationType = 28
	NotificationTypeMessageReminderDue                           NotificationType = 29
	NotificationTypeIcymiSummary                                 NotificationType = 30
	NotificationTypeChannelPromptDeadchat                        NotificationType = 31
	NotificationTypeReactionTrendingPushNotification             NotificationType = 32
	NotificationTypeSummonsDirect                                NotificationType = 33
	NotificationTypeAdminNotificationPush                        NotificationType = 34
	NotificationTypeVoiceChannelActivity                         NotificationType = 35
	NotificationTypeMissedMessageEmail                           NotificationType = 36
	NotificationTypeCustomStatusUpdate                           NotificationType = 37
	NotificationTypeGoLiveNotification                           NotificationType = 38
	NotificationTypeFriendGamingActivityPush                     NotificationType = 39
	NotificationTypeUserResurrectionNotification                 NotificationType = 40
	NotificationTypeFriendOnlinePush                             NotificationType = 41
	NotificationTypeInviteReminderPush                           NotificationType = 42
	NotificationTypeServerTrendingNotification                   NotificationType = 43
	NotificationTypeMessagePin                                   NotificationType = 44
	NotificationTypeProfileUpdatesNotification                   NotificationType = 45
	NotificationTypeSummaryReminder                              NotificationType = 46
	NotificationTypeTrialForAllReminder                          NotificationType = 47
	NotificationTypeFriendsPlayingGame                           NotificationType = 48
	NotificationTypeReferralProgramPushNotifEntrypointReminder   NotificationType = 49
	NotificationTypeReferralProgramNotifCenterEntrypointReminder NotificationType = 50
	NotificationTypeGameUpdate                                   NotificationType = 51
	NotificationTypeVoiceChannelInvite                           NotificationType = 52
	NotificationTypeGuildScheduledEventUpcoming                  NotificationType = 53
	NotificationTypeFamilyCenterRestrictedScheduleUpdated        NotificationType = 54
	NotificationTypeVoiceChannelActivityPeakAffinity             NotificationType = 55
	NotificationTypeParentalConsentFinalWarning                  NotificationType = 56
	NotificationTypeMessageRequest                               NotificationType = 57
	NotificationTypeFriendsPlayingTrendingGame                   NotificationType = 58
)

func (v NotificationType) String() string {
	switch v {
	case NotificationTypeMessageCreate:
		return "MESSAGE_CREATE"
	case NotificationTypeCallRing:
		return "CALL_RING"
	case NotificationTypeRelationshipAdd:
		return "RELATIONSHIP_ADD"
	case NotificationTypeFriendSuggestionCreate:
		return "FRIEND_SUGGESTION_CREATE"
	case NotificationTypeApplicationLibraryInstallComplete:
		return "APPLICATION_LIBRARY_INSTALL_COMPLETE"
	case NotificationTypeGuildStreamStart:
		return "GUILD_STREAM_START"
	case NotificationTypeStageInstanceCreate:
		return "STAGE_INSTANCE_CREATE"
	case NotificationTypeGuildScheduledEventUpdate:
		return "GUILD_SCHEDULED_EVENT_UPDATE"
	case NotificationTypeForumThreadCreated:
		return "FORUM_THREAD_CREATED"
	case NotificationTypeTopMessagesPush:
		return "TOP_MESSAGES_PUSH"
	case NotificationTypeMissedMessagePush:
		return "MISSED_MESSAGE_PUSH"
	case NotificationTypeFriendRequestReminderPush:
		return "FRIEND_REQUEST_REMINDER_PUSH"
	case NotificationTypeHomeLifecyclePush:
		return "HOME_LIFECYCLE_PUSH"
	case NotificationTypePollEndedPush:
		return "POLL_ENDED_PUSH"
	case NotificationTypeNudgeNewFriendDMPush:
		return "NUDGE_NEW_FRIEND_DM_PUSH"
	case NotificationTypeFamilyCenterRequestSend:
		return "FAMILY_CENTER_REQUEST_SEND"
	case NotificationTypeFamilyCenterRequestAccepted:
		return "FAMILY_CENTER_REQUEST_ACCEPTED"
	case NotificationTypeFamilyCenterRequestDeclined:
		return "FAMILY_CENTER_REQUEST_DECLINED"
	case NotificationTypeFamilyCenterDisconnected:
		return "FAMILY_CENTER_DISCONNECTED"
	case NotificationTypeGuildJoinRequestApproved:
		return "GUILD_JOIN_REQUEST_APPROVED"
	case NotificationTypeGuildJoinRequestRejected:
		return "GUILD_JOIN_REQUEST_REJECTED"
	case NotificationTypeReactionsPushNotification:
		return "REACTIONS_PUSH_NOTIFICATION"
	case NotificationTypeRaidDetected:
		return "RAID_DETECTED"
	case NotificationTypeMentionRaidDetected:
		return "MENTION_RAID_DETECTED"
	case NotificationTypeDMSpamDetected:
		return "DM_SPAM_DETECTED"
	case NotificationTypeSuspiciousSession:
		return "SUSPICIOUS_SESSION"
	case NotificationTypeNewUserSession:
		return "NEW_USER_SESSION"
	case NotificationTypeMessageReminderDue:
		return "MESSAGE_REMINDER_DUE"
	case NotificationTypeIcymiSummary:
		return "ICYMI_SUMMARY"
	case NotificationTypeChannelPromptDeadchat:
		return "CHANNEL_PROMPT_DEADCHAT"
	case NotificationTypeReactionTrendingPushNotification:
		return "REACTION_TRENDING_PUSH_NOTIFICATION"
	case NotificationTypeSummonsDirect:
		return "SUMMONS_DIRECT"
	case NotificationTypeAdminNotificationPush:
		return "ADMIN_NOTIFICATION_PUSH"
	case NotificationTypeVoiceChannelActivity:
		return "VOICE_CHANNEL_ACTIVITY"
	case NotificationTypeMissedMessageEmail:
		return "MISSED_MESSAGE_EMAIL"
	case NotificationTypeCustomStatusUpdate:
		return "CUSTOM_STATUS_UPDATE"
	case NotificationTypeGoLiveNotification:
		return "GO_LIVE_NOTIFICATION"
	case NotificationTypeFriendGamingActivityPush:
		return "FRIEND_GAMING_ACTIVITY_PUSH"
	case NotificationTypeUserResurrectionNotification:
		return "USER_RESURRECTION_NOTIFICATION"
	case NotificationTypeFriendOnlinePush:
		return "FRIEND_ONLINE_PUSH"
	case NotificationTypeInviteReminderPush:
		return "INVITE_REMINDER_PUSH"
	case NotificationTypeServerTrendingNotification:
		return "SERVER_TRENDING_NOTIFICATION"
	case NotificationTypeMessagePin:
		return "MESSAGE_PIN"
	case NotificationTypeProfileUpdatesNotification:
		return "PROFILE_UPDATES_NOTIFICATION"
	case NotificationTypeSummaryReminder:
		return "SUMMARY_REMINDER"
	case NotificationTypeTrialForAllReminder:
		return "TRIAL_FOR_ALL_REMINDER"
	case NotificationTypeFriendsPlayingGame:
		return "FRIENDS_PLAYING_GAME"
	case NotificationTypeReferralProgramPushNotifEntrypointReminder:
		return "REFERRAL_PROGRAM_PUSH_NOTIF_ENTRYPOINT_REMINDER"
	case NotificationTypeReferralProgramNotifCenterEntrypointReminder:
		return "REFERRAL_PROGRAM_NOTIF_CENTER_ENTRYPOINT_REMINDER"
	case NotificationTypeGameUpdate:
		return "GAME_UPDATE"
	case NotificationTypeVoiceChannelInvite:
		return "VOICE_CHANNEL_INVITE"
	case NotificationTypeGuildScheduledEventUpcoming:
		return "GUILD_SCHEDULED_EVENT_UPCOMING"
	case NotificationTypeFamilyCenterRestrictedScheduleUpdated:
		return "FAMILY_CENTER_RESTRICTED_SCHEDULE_UPDATED"
	case NotificationTypeVoiceChannelActivityPeakAffinity:
		return "VOICE_CHANNEL_ACTIVITY_PEAK_AFFINITY"
	case NotificationTypeParentalConsentFinalWarning:
		return "PARENTAL_CONSENT_FINAL_WARNING"
	case NotificationTypeMessageRequest:
		return "MESSAGE_REQUEST"
	case NotificationTypeFriendsPlayingTrendingGame:
		return "FRIENDS_PLAYING_TRENDING_GAME"
	}
	return strconv.FormatInt(int64(v), 10)
}

type OAuth2DeviceFlowStatus string

const (
	OAuth2DeviceFlowStatusPending OAuth2DeviceFlowStatus = "pending"
	OAuth2DeviceFlowStatusGranted OAuth2DeviceFlowStatus = "granted"
	OAuth2DeviceFlowStatusDenied  OAuth2DeviceFlowStatus = "denied"
)

type OAuth2GrantType string

const (
	OAuth2GrantTypeAuthorizationCode                     OAuth2GrantType = "authorization_code"
	OAuth2GrantTypeRefreshToken                          OAuth2GrantType = "refresh_token"
	OAuth2GrantTypeClientCredentials                     OAuth2GrantType = "client_credentials"
	OAuth2GrantTypeUrnIetfParamsOauthGrantTypeDeviceCode OAuth2GrantType = "urn:ietf:params:oauth:grant-type:device_code"
)

type OAuth2Scopes string

const (
	OAuth2ScopesAccountGlobalNameUpdate               OAuth2Scopes = "account.global_name.update"
	OAuth2ScopesActivitiesInvitesWrite                OAuth2Scopes = "activities.invites.write"
	OAuth2ScopesActivitiesRead                        OAuth2Scopes = "activities.read"
	OAuth2ScopesActivitiesWrite                       OAuth2Scopes = "activities.write"
	OAuth2ScopesApplicationsBuildsRead                OAuth2Scopes = "applications.builds.read"
	OAuth2ScopesApplicationsBuildsUpload              OAuth2Scopes = "applications.builds.upload"
	OAuth2ScopesApplicationsCommands                  OAuth2Scopes = "applications.commands"
	OAuth2ScopesApplicationsCommandsPermissionsUpdate OAuth2Scopes = "applications.commands.permissions.update"
	OAuth2ScopesApplicationsCommandsUpdate            OAuth2Scopes = "applications.commands.update"
	OAuth2ScopesApplicationsEntitlements              OAuth2Scopes = "applications.entitlements"
	OAuth2ScopesApplicationsStoreUpdate               OAuth2Scopes = "applications.store.update"
	OAuth2ScopesApplicationIdentitiesWrite            OAuth2Scopes = "application_identities.write"
	OAuth2ScopesBot                                   OAuth2Scopes = "bot"
	OAuth2ScopesConnections                           OAuth2Scopes = "connections"
	OAuth2ScopesDMChannelsRead                        OAuth2Scopes = "dm_channels.read"
	OAuth2ScopesDMChannelsMessagesRead                OAuth2Scopes = "dm_channels.messages.read"
	OAuth2ScopesDMChannelsMessagesWrite               OAuth2Scopes = "dm_channels.messages.write"
	OAuth2ScopesEmail                                 OAuth2Scopes = "email"
	OAuth2ScopesGatewayConnect                        OAuth2Scopes = "gateway.connect"
	OAuth2ScopesGdmJoin                               OAuth2Scopes = "gdm.join"
	OAuth2ScopesGuilds                                OAuth2Scopes = "guilds"
	OAuth2ScopesGuildsChannelsRead                    OAuth2Scopes = "guilds.channels.read"
	OAuth2ScopesGuildsJoin                            OAuth2Scopes = "guilds.join"
	OAuth2ScopesGuildsMembersRead                     OAuth2Scopes = "guilds.members.read"
	OAuth2ScopesIdentify                              OAuth2Scopes = "identify"
	OAuth2ScopesLobbiesWrite                          OAuth2Scopes = "lobbies.write"
	OAuth2ScopesMessagesRead                          OAuth2Scopes = "messages.read"
	OAuth2ScopesOpenid                                OAuth2Scopes = "openid"
	OAuth2ScopesPaymentSourcesCountryCode             OAuth2Scopes = "payment_sources.country_code"
	OAuth2ScopesPresencesRead                         OAuth2Scopes = "presences.read"
	OAuth2ScopesPresencesWrite                        OAuth2Scopes = "presences.write"
	OAuth2ScopesRelationshipsRead                     OAuth2Scopes = "relationships.read"
	OAuth2ScopesRelationshipsWrite                    OAuth2Scopes = "relationships.write"
	OAuth2ScopesRoleConnectionsWrite                  OAuth2Scopes = "role_connections.write"
	OAuth2ScopesRPC                                   OAuth2Scopes = "rpc"
	OAuth2ScopesRPCActivitiesWrite                    OAuth2Scopes = "rpc.activities.write"
	OAuth2ScopesRPCAPI                                OAuth2Scopes = "rpc.api"
	OAuth2ScopesRPCNotificationsRead                  OAuth2Scopes = "rpc.notifications.read"
	OAuth2ScopesRPCScreenshareRead                    OAuth2Scopes = "rpc.screenshare.read"
	OAuth2ScopesRPCScreenshareWrite                   OAuth2Scopes = "rpc.screenshare.write"
	OAuth2ScopesRPCVideoRead                          OAuth2Scopes = "rpc.video.read"
	OAuth2ScopesRPCVideoWrite                         OAuth2Scopes = "rpc.video.write"
	OAuth2ScopesRPCVoiceRead                          OAuth2Scopes = "rpc.voice.read"
	OAuth2ScopesRPCVoiceWrite                         OAuth2Scopes = "rpc.voice.write"
	OAuth2ScopesVoice                                 OAuth2Scopes = "voice"
	OAuth2ScopesWebhookIncoming                       OAuth2Scopes = "webhook.incoming"
)

type OnboardingConnectionType int64

const (
	OnboardingConnectionTypeApplication              OnboardingConnectionType = 0
	OnboardingConnectionTypeProviderConnectedAccount OnboardingConnectionType = 1
)

func (v OnboardingConnectionType) String() string {
	switch v {
	case OnboardingConnectionTypeApplication:
		return "APPLICATION"
	case OnboardingConnectionTypeProviderConnectedAccount:
		return "PROVIDER_CONNECTED_ACCOUNT"
	}
	return strconv.FormatInt(int64(v), 10)
}

type OnboardingMode int64

const (
	OnboardingModeOnboardingDefault  OnboardingMode = 0
	OnboardingModeOnboardingAdvanced OnboardingMode = 1
)

func (v OnboardingMode) String() string {
	switch v {
	case OnboardingModeOnboardingDefault:
		return "ONBOARDING_DEFAULT"
	case OnboardingModeOnboardingAdvanced:
		return "ONBOARDING_ADVANCED"
	}
	return strconv.FormatInt(int64(v), 10)
}

type OnboardingPromptType int64

const (
	OnboardingPromptTypeMultipleChoice OnboardingPromptType = 0
	OnboardingPromptTypeDropdown       OnboardingPromptType = 1
)

func (v OnboardingPromptType) String() string {
	switch v {
	case OnboardingPromptTypeMultipleChoice:
		return "MULTIPLE_CHOICE"
	case OnboardingPromptTypeDropdown:
		return "DROPDOWN"
	}
	return strconv.FormatInt(int64(v), 10)
}

type OpencriticTier int64

const (
	OpencriticTierMighty OpencriticTier = 1
	OpencriticTierStrong OpencriticTier = 2
	OpencriticTierFair   OpencriticTier = 3
	OpencriticTierWeak   OpencriticTier = 4
)

func (v OpencriticTier) String() string {
	switch v {
	case OpencriticTierMighty:
		return "MIGHTY"
	case OpencriticTierStrong:
		return "STRONG"
	case OpencriticTierFair:
		return "FAIR"
	case OpencriticTierWeak:
		return "WEAK"
	}
	return strconv.FormatInt(int64(v), 10)
}

type OperatingSystem string

const (
	OperatingSystemWin32  OperatingSystem = "win32"
	OperatingSystemDarwin OperatingSystem = "darwin"
	OperatingSystemLinux  OperatingSystem = "linux"
)

type OperatingSystemType string

const (
	OperatingSystemTypeWindows     OperatingSystemType = "windows"
	OperatingSystemTypeOsx         OperatingSystemType = "osx"
	OperatingSystemTypeLinux       OperatingSystemType = "linux"
	OperatingSystemTypeAndroid     OperatingSystemType = "android"
	OperatingSystemTypeIos         OperatingSystemType = "ios"
	OperatingSystemTypePlaystation OperatingSystemType = "playstation"
	OperatingSystemTypeXbox        OperatingSystemType = "xbox"
	OperatingSystemTypeUnknown     OperatingSystemType = "unknown"
)

type OrientationState int64

const (
	OrientationStateUnlocked  OrientationState = 1
	OrientationStatePortrait  OrientationState = 2
	OrientationStateLandscape OrientationState = 3
)

func (v OrientationState) String() string {
	switch v {
	case OrientationStateUnlocked:
		return "UNLOCKED"
	case OrientationStatePortrait:
		return "PORTRAIT"
	case OrientationStateLandscape:
		return "LANDSCAPE"
	}
	return strconv.FormatInt(int64(v), 10)
}

type OverlayMethodFlags uint64

const (
	OverlayMethodFlagsOutOfProcess OverlayMethodFlags = 1 << 0
)

func (v OverlayMethodFlags) String() string {
	switch v {
	case OverlayMethodFlagsOutOfProcess:
		return "OUT_OF_PROCESS"
	}
	return strconv.FormatUint(uint64(v), 10)
}

type PasswordResetMethod string

const (
	PasswordResetMethodPasswordReset PasswordResetMethod = "password_reset"
	PasswordResetMethodOneTimeLogin  PasswordResetMethod = "one_time_login"
)

type PaymentFlags uint64

const (
	PaymentFlagsGift                   PaymentFlags = 1 << 0
	PaymentFlagsUserRefunded           PaymentFlags = 1 << 2
	PaymentFlagsPreorder               PaymentFlags = 1 << 3
	PaymentFlagsPending                PaymentFlags = 1 << 4
	PaymentFlagsTemporaryAuthorization PaymentFlags = 1 << 5
)

func (v PaymentFlags) String() string {
	switch v {
	case PaymentFlagsGift:
		return "GIFT"
	case PaymentFlagsUserRefunded:
		return "USER_REFUNDED"
	case PaymentFlagsPreorder:
		return "PREORDER"
	case PaymentFlagsPending:
		return "PENDING"
	case PaymentFlagsTemporaryAuthorization:
		return "TEMPORARY_AUTHORIZATION"
	}
	return strconv.FormatUint(uint64(v), 10)
}

type PaymentGateway int64

const (
	PaymentGatewayStripe                PaymentGateway = 1
	PaymentGatewayBraintree             PaymentGateway = 2
	PaymentGatewayApple                 PaymentGateway = 3
	PaymentGatewayGoogle                PaymentGateway = 4
	PaymentGatewayAdyen                 PaymentGateway = 5
	PaymentGatewayApplePartner          PaymentGateway = 6
	PaymentGatewayAmazon                PaymentGateway = 7
	PaymentGatewayVirtualCurrency       PaymentGateway = 8
	PaymentGatewayAppleAdvancedCommerce PaymentGateway = 9
	PaymentGatewayTds                   PaymentGateway = 10
)

func (v PaymentGateway) String() string {
	switch v {
	case PaymentGatewayStripe:
		return "STRIPE"
	case PaymentGatewayBraintree:
		return "BRAINTREE"
	case PaymentGatewayApple:
		return "APPLE"
	case PaymentGatewayGoogle:
		return "GOOGLE"
	case PaymentGatewayAdyen:
		return "ADYEN"
	case PaymentGatewayApplePartner:
		return "APPLE_PARTNER"
	case PaymentGatewayAmazon:
		return "AMAZON"
	case PaymentGatewayVirtualCurrency:
		return "VIRTUAL_CURRENCY"
	case PaymentGatewayAppleAdvancedCommerce:
		return "APPLE_ADVANCED_COMMERCE"
	case PaymentGatewayTds:
		return "TDS"
	}
	return strconv.FormatInt(int64(v), 10)
}

type PaymentSourceFlags uint64

const (
	PaymentSourceFlagsNew               PaymentSourceFlags = 1 << 0
	PaymentSourceFlagsSuccessfulPayment PaymentSourceFlags = 1 << 1
)

func (v PaymentSourceFlags) String() string {
	switch v {
	case PaymentSourceFlagsNew:
		return "NEW"
	case PaymentSourceFlagsSuccessfulPayment:
		return "SUCCESSFUL_PAYMENT"
	}
	return strconv.FormatUint(uint64(v), 10)
}

type PaymentSourceType int64

const (
	PaymentSourceTypeCard        PaymentSourceType = 1
	PaymentSourceTypePaypal      PaymentSourceType = 2
	PaymentSourceTypeGiropay     PaymentSourceType = 3
	PaymentSourceTypeSofort      PaymentSourceType = 4
	PaymentSourceTypePrzelewy24  PaymentSourceType = 5
	PaymentSourceTypeSepaDebit   PaymentSourceType = 6
	PaymentSourceTypePaysafeCard PaymentSourceType = 7
	PaymentSourceTypeGcash       PaymentSourceType = 8
	PaymentSourceTypeGrabpayMy   PaymentSourceType = 9
	PaymentSourceTypeMomoWallet  PaymentSourceType = 10
	PaymentSourceTypeVenmo       PaymentSourceType = 11
	PaymentSourceTypeGopayWallet PaymentSourceType = 12
	PaymentSourceTypeKakaopay    PaymentSourceType = 13
	PaymentSourceTypeBancontact  PaymentSourceType = 14
	PaymentSourceTypeEps         PaymentSourceType = 15
	PaymentSourceTypeIdeal       PaymentSourceType = 16
	PaymentSourceTypeCashApp     PaymentSourceType = 17
	PaymentSourceTypeApple       PaymentSourceType = 18
	PaymentSourceTypeTdsWallet   PaymentSourceType = 19
)

func (v PaymentSourceType) String() string {
	switch v {
	case PaymentSourceTypeCard:
		return "CARD"
	case PaymentSourceTypePaypal:
		return "PAYPAL"
	case PaymentSourceTypeGiropay:
		return "GIROPAY"
	case PaymentSourceTypeSofort:
		return "SOFORT"
	case PaymentSourceTypePrzelewy24:
		return "PRZELEWY24"
	case PaymentSourceTypeSepaDebit:
		return "SEPA_DEBIT"
	case PaymentSourceTypePaysafeCard:
		return "PAYSAFE_CARD"
	case PaymentSourceTypeGcash:
		return "GCASH"
	case PaymentSourceTypeGrabpayMy:
		return "GRABPAY_MY"
	case PaymentSourceTypeMomoWallet:
		return "MOMO_WALLET"
	case PaymentSourceTypeVenmo:
		return "VENMO"
	case PaymentSourceTypeGopayWallet:
		return "GOPAY_WALLET"
	case PaymentSourceTypeKakaopay:
		return "KAKAOPAY"
	case PaymentSourceTypeBancontact:
		return "BANCONTACT"
	case PaymentSourceTypeEps:
		return "EPS"
	case PaymentSourceTypeIdeal:
		return "IDEAL"
	case PaymentSourceTypeCashApp:
		return "CASH_APP"
	case PaymentSourceTypeApple:
		return "APPLE"
	case PaymentSourceTypeTdsWallet:
		return "TDS_WALLET"
	}
	return strconv.FormatInt(int64(v), 10)
}

type PaymentStatus int64

const (
	PaymentStatusPending   PaymentStatus = 0
	PaymentStatusCompleted PaymentStatus = 1
	PaymentStatusFailed    PaymentStatus = 2
	PaymentStatusReversed  PaymentStatus = 3
	PaymentStatusRefunded  PaymentStatus = 4
	PaymentStatusCanceled  PaymentStatus = 5
)

func (v PaymentStatus) String() string {
	switch v {
	case PaymentStatusPending:
		return "PENDING"
	case PaymentStatusCompleted:
		return "COMPLETED"
	case PaymentStatusFailed:
		return "FAILED"
	case PaymentStatusReversed:
		return "REVERSED"
	case PaymentStatusRefunded:
		return "REFUNDED"
	case PaymentStatusCanceled:
		return "CANCELED"
	}
	return strconv.FormatInt(int64(v), 10)
}

type PegiContentDescriptor int64

const (
	PegiContentDescriptorViolence       PegiContentDescriptor = 1
	PegiContentDescriptorBadLanguage    PegiContentDescriptor = 2
	PegiContentDescriptorFear           PegiContentDescriptor = 3
	PegiContentDescriptorGambling       PegiContentDescriptor = 4
	PegiContentDescriptorSex            PegiContentDescriptor = 5
	PegiContentDescriptorDrugs          PegiContentDescriptor = 6
	PegiContentDescriptorDiscrimination PegiContentDescriptor = 7
)

func (v PegiContentDescriptor) String() string {
	switch v {
	case PegiContentDescriptorViolence:
		return "VIOLENCE"
	case PegiContentDescriptorBadLanguage:
		return "BAD_LANGUAGE"
	case PegiContentDescriptorFear:
		return "FEAR"
	case PegiContentDescriptorGambling:
		return "GAMBLING"
	case PegiContentDescriptorSex:
		return "SEX"
	case PegiContentDescriptorDrugs:
		return "DRUGS"
	case PegiContentDescriptorDiscrimination:
		return "DISCRIMINATION"
	}
	return strconv.FormatInt(int64(v), 10)
}

type PegiContentRating int64

const (
	PegiContentRatingThree    PegiContentRating = 1
	PegiContentRatingSeven    PegiContentRating = 2
	PegiContentRatingTwelve   PegiContentRating = 3
	PegiContentRatingSixteen  PegiContentRating = 4
	PegiContentRatingEighteen PegiContentRating = 5
)

func (v PegiContentRating) String() string {
	switch v {
	case PegiContentRatingThree:
		return "THREE"
	case PegiContentRatingSeven:
		return "SEVEN"
	case PegiContentRatingTwelve:
		return "TWELVE"
	case PegiContentRatingSixteen:
		return "SIXTEEN"
	case PegiContentRatingEighteen:
		return "EIGHTEEN"
	}
	return strconv.FormatInt(int64(v), 10)
}

type Permission uint64

const (
	PermissionCreateInstantInvite              Permission = 1 << 0
	PermissionKickMembers                      Permission = 1 << 1
	PermissionBanMembers                       Permission = 1 << 2
	PermissionAdministrator                    Permission = 1 << 3
	PermissionManageChannels                   Permission = 1 << 4
	PermissionManageGuild                      Permission = 1 << 5
	PermissionAddReactions                     Permission = 1 << 6
	PermissionViewAuditLog                     Permission = 1 << 7
	PermissionPrioritySpeaker                  Permission = 1 << 8
	PermissionStream                           Permission = 1 << 9
	PermissionViewChannel                      Permission = 1 << 10
	PermissionSendMessages                     Permission = 1 << 11
	PermissionSendTTSMessages                  Permission = 1 << 12
	PermissionManageMessages                   Permission = 1 << 13
	PermissionEmbedLinks                       Permission = 1 << 14
	PermissionAttachFiles                      Permission = 1 << 15
	PermissionReadMessageHistory               Permission = 1 << 16
	PermissionMentionEveryone                  Permission = 1 << 17
	PermissionUseExternalEmojis                Permission = 1 << 18
	PermissionViewGuildInsights                Permission = 1 << 19
	PermissionConnect                          Permission = 1 << 20
	PermissionSpeak                            Permission = 1 << 21
	PermissionMuteMembers                      Permission = 1 << 22
	PermissionDeafenMembers                    Permission = 1 << 23
	PermissionMoveMembers                      Permission = 1 << 24
	PermissionUseVad                           Permission = 1 << 25
	PermissionChangeNickname                   Permission = 1 << 26
	PermissionManageNicknames                  Permission = 1 << 27
	PermissionManageRoles                      Permission = 1 << 28
	PermissionManageWebhooks                   Permission = 1 << 29
	PermissionManageExpressions                Permission = 1 << 30
	PermissionUseApplicationCommands           Permission = 1 << 31
	PermissionRequestToSpeak                   Permission = 1 << 32
	PermissionManageEvents                     Permission = 1 << 33
	PermissionManageThreads                    Permission = 1 << 34
	PermissionCreatePublicThreads              Permission = 1 << 35
	PermissionCreatePrivateThreads             Permission = 1 << 36
	PermissionUseExternalStickers              Permission = 1 << 37
	PermissionSendMessagesInThreads            Permission = 1 << 38
	PermissionUseEmbeddedActivities            Permission = 1 << 39
	PermissionModerateMembers                  Permission = 1 << 40
	PermissionViewCreatorMonetizationAnalytics Permission = 1 << 41
	PermissionUseSoundboard                    Permission = 1 << 42
	PermissionCreateExpressions                Permission = 1 << 43
	PermissionCreateEvents                     Permission = 1 << 44
	PermissionUseExternalSounds                Permission = 1 << 45
	PermissionSendVoiceMessages                Permission = 1 << 46
	PermissionUseClydeAi                       Permission = 1 << 47
	PermissionSetVoiceChannelStatus            Permission = 1 << 48
	PermissionSendPolls                        Permission = 1 << 49
	PermissionUseExternalApps                  Permission = 1 << 50
	PermissionPinMessages                      Permission = 1 << 51
	PermissionBypassSlowmode                   Permission = 1 << 52
	PermissionManageOfficialMessages           Permission = 1 << 53
)

func (v Permission) String() string {
	switch v {
	case PermissionCreateInstantInvite:
		return "CREATE_INSTANT_INVITE"
	case PermissionKickMembers:
		return "KICK_MEMBERS"
	case PermissionBanMembers:
		return "BAN_MEMBERS"
	case PermissionAdministrator:
		return "ADMINISTRATOR"
	case PermissionManageChannels:
		return "MANAGE_CHANNELS"
	case PermissionManageGuild:
		return "MANAGE_GUILD"
	case PermissionAddReactions:
		return "ADD_REACTIONS"
	case PermissionViewAuditLog:
		return "VIEW_AUDIT_LOG"
	case PermissionPrioritySpeaker:
		return "PRIORITY_SPEAKER"
	case PermissionStream:
		return "STREAM"
	case PermissionViewChannel:
		return "VIEW_CHANNEL"
	case PermissionSendMessages:
		return "SEND_MESSAGES"
	case PermissionSendTTSMessages:
		return "SEND_TTS_MESSAGES"
	case PermissionManageMessages:
		return "MANAGE_MESSAGES"
	case PermissionEmbedLinks:
		return "EMBED_LINKS"
	case PermissionAttachFiles:
		return "ATTACH_FILES"
	case PermissionReadMessageHistory:
		return "READ_MESSAGE_HISTORY"
	case PermissionMentionEveryone:
		return "MENTION_EVERYONE"
	case PermissionUseExternalEmojis:
		return "USE_EXTERNAL_EMOJIS"
	case PermissionViewGuildInsights:
		return "VIEW_GUILD_INSIGHTS"
	case PermissionConnect:
		return "CONNECT"
	case PermissionSpeak:
		return "SPEAK"
	case PermissionMuteMembers:
		return "MUTE_MEMBERS"
	case PermissionDeafenMembers:
		return "DEAFEN_MEMBERS"
	case PermissionMoveMembers:
		return "MOVE_MEMBERS"
	case PermissionUseVad:
		return "USE_VAD"
	case PermissionChangeNickname:
		return "CHANGE_NICKNAME"
	case PermissionManageNicknames:
		return "MANAGE_NICKNAMES"
	case PermissionManageRoles:
		return "MANAGE_ROLES"
	case PermissionManageWebhooks:
		return "MANAGE_WEBHOOKS"
	case PermissionManageExpressions:
		return "MANAGE_EXPRESSIONS"
	case PermissionUseApplicationCommands:
		return "USE_APPLICATION_COMMANDS"
	case PermissionRequestToSpeak:
		return "REQUEST_TO_SPEAK"
	case PermissionManageEvents:
		return "MANAGE_EVENTS"
	case PermissionManageThreads:
		return "MANAGE_THREADS"
	case PermissionCreatePublicThreads:
		return "CREATE_PUBLIC_THREADS"
	case PermissionCreatePrivateThreads:
		return "CREATE_PRIVATE_THREADS"
	case PermissionUseExternalStickers:
		return "USE_EXTERNAL_STICKERS"
	case PermissionSendMessagesInThreads:
		return "SEND_MESSAGES_IN_THREADS"
	case PermissionUseEmbeddedActivities:
		return "USE_EMBEDDED_ACTIVITIES"
	case PermissionModerateMembers:
		return "MODERATE_MEMBERS"
	case PermissionViewCreatorMonetizationAnalytics:
		return "VIEW_CREATOR_MONETIZATION_ANALYTICS"
	case PermissionUseSoundboard:
		return "USE_SOUNDBOARD"
	case PermissionCreateExpressions:
		return "CREATE_EXPRESSIONS"
	case PermissionCreateEvents:
		return "CREATE_EVENTS"
	case PermissionUseExternalSounds:
		return "USE_EXTERNAL_SOUNDS"
	case PermissionSendVoiceMessages:
		return "SEND_VOICE_MESSAGES"
	case PermissionUseClydeAi:
		return "USE_CLYDE_AI"
	case PermissionSetVoiceChannelStatus:
		return "SET_VOICE_CHANNEL_STATUS"
	case PermissionSendPolls:
		return "SEND_POLLS"
	case PermissionUseExternalApps:
		return "USE_EXTERNAL_APPS"
	case PermissionPinMessages:
		return "PIN_MESSAGES"
	case PermissionBypassSlowmode:
		return "BYPASS_SLOWMODE"
	case PermissionManageOfficialMessages:
		return "MANAGE_OFFICIAL_MESSAGES"
	}
	return strconv.FormatUint(uint64(v), 10)
}

type PermissionOverwriteType int64

const (
	PermissionOverwriteTypeRole   PermissionOverwriteType = 0
	PermissionOverwriteTypeMember PermissionOverwriteType = 1
)

func (v PermissionOverwriteType) String() string {
	switch v {
	case PermissionOverwriteTypeRole:
		return "role"
	case PermissionOverwriteTypeMember:
		return "member"
	}
	return strconv.FormatInt(int64(v), 10)
}

type PollLayoutType int64

const (
	PollLayoutTypeDefault          PollLayoutType = 1
	PollLayoutTypeImageOnlyAnswers PollLayoutType = 2
)

func (v PollLayoutType) String() string {
	switch v {
	case PollLayoutTypeDefault:
		return "DEFAULT"
	case PollLayoutTypeImageOnlyAnswers:
		return "IMAGE_ONLY_ANSWERS"
	}
	return strconv.FormatInt(int64(v), 10)
}

type PotionType int64

const (
	PotionTypeConfetti PotionType = 0
)

func (v PotionType) String() string {
	switch v {
	case PotionTypeConfetti:
		return "CONFETTI"
	}
	return strconv.FormatInt(int64(v), 10)
}

type PremiumReferralIncentiveStatus int64

const (
	PremiumReferralIncentiveStatusNotEligible PremiumReferralIncentiveStatus = 0
	PremiumReferralIncentiveStatusEligible    PremiumReferralIncentiveStatus = 1
	PremiumReferralIncentiveStatusQualified   PremiumReferralIncentiveStatus = 2
	PremiumReferralIncentiveStatusCooldown    PremiumReferralIncentiveStatus = 3
	PremiumReferralIncentiveStatusUnapplied   PremiumReferralIncentiveStatus = 4
)

func (v PremiumReferralIncentiveStatus) String() string {
	switch v {
	case PremiumReferralIncentiveStatusNotEligible:
		return "NOT_ELIGIBLE"
	case PremiumReferralIncentiveStatusEligible:
		return "ELIGIBLE"
	case PremiumReferralIncentiveStatusQualified:
		return "QUALIFIED"
	case PremiumReferralIncentiveStatusCooldown:
		return "COOLDOWN"
	case PremiumReferralIncentiveStatusUnapplied:
		return "UNAPPLIED"
	}
	return strconv.FormatInt(int64(v), 10)
}

type PremiumReferralRecipientStatus int64

const (
	PremiumReferralRecipientStatusRedeemed PremiumReferralRecipientStatus = 1
	PremiumReferralRecipientStatusPending  PremiumReferralRecipientStatus = 2
)

func (v PremiumReferralRecipientStatus) String() string {
	switch v {
	case PremiumReferralRecipientStatusRedeemed:
		return "REDEEMED"
	case PremiumReferralRecipientStatusPending:
		return "PENDING"
	}
	return strconv.FormatInt(int64(v), 10)
}

type PremiumSource int64

const (
	PremiumSourceSubscription      PremiumSource = 1
	PremiumSourceFractionalPremium PremiumSource = 2
	PremiumSourceReverseTrial      PremiumSource = 3
	PremiumSourceSubscriptionGroup PremiumSource = 4
)

func (v PremiumSource) String() string {
	switch v {
	case PremiumSourceSubscription:
		return "SUBSCRIPTION"
	case PremiumSourceFractionalPremium:
		return "FRACTIONAL_PREMIUM"
	case PremiumSourceReverseTrial:
		return "REVERSE_TRIAL"
	case PremiumSourceSubscriptionGroup:
		return "SUBSCRIPTION_GROUP"
	}
	return strconv.FormatInt(int64(v), 10)
}

type PremiumSubscriptionGroupRole int64

const (
	PremiumSubscriptionGroupRolePrimary PremiumSubscriptionGroupRole = 1
	PremiumSubscriptionGroupRoleMember  PremiumSubscriptionGroupRole = 2
)

func (v PremiumSubscriptionGroupRole) String() string {
	switch v {
	case PremiumSubscriptionGroupRolePrimary:
		return "PRIMARY"
	case PremiumSubscriptionGroupRoleMember:
		return "MEMBER"
	}
	return strconv.FormatInt(int64(v), 10)
}

type PremiumSubscriptionType int64

const (
	PremiumSubscriptionTypeBoostOnly PremiumSubscriptionType = 1
	PremiumSubscriptionTypeTier0     PremiumSubscriptionType = 2
	PremiumSubscriptionTypeTier1     PremiumSubscriptionType = 3
	PremiumSubscriptionTypeTier2     PremiumSubscriptionType = 4
)

func (v PremiumSubscriptionType) String() string {
	switch v {
	case PremiumSubscriptionTypeBoostOnly:
		return "BOOST_ONLY"
	case PremiumSubscriptionTypeTier0:
		return "TIER_0"
	case PremiumSubscriptionTypeTier1:
		return "TIER_1"
	case PremiumSubscriptionTypeTier2:
		return "TIER_2"
	}
	return strconv.FormatInt(int64(v), 10)
}

type PremiumTier int64

const (
	PremiumTierNone  PremiumTier = 0
	PremiumTierTier1 PremiumTier = 1
	PremiumTierTier2 PremiumTier = 2
	PremiumTierTier3 PremiumTier = 3
)

func (v PremiumTier) String() string {
	switch v {
	case PremiumTierNone:
		return "NONE"
	case PremiumTierTier1:
		return "TIER_1"
	case PremiumTierTier2:
		return "TIER_2"
	case PremiumTierTier3:
		return "TIER_3"
	}
	return strconv.FormatInt(int64(v), 10)
}

type PremiumType int64

const (
	PremiumTypeNone  PremiumType = 0
	PremiumTypeTier1 PremiumType = 1
	PremiumTypeTier2 PremiumType = 2
	PremiumTypeTier3 PremiumType = 3
)

func (v PremiumType) String() string {
	switch v {
	case PremiumTypeNone:
		return "NONE"
	case PremiumTypeTier1:
		return "TIER_1"
	case PremiumTypeTier2:
		return "TIER_2"
	case PremiumTypeTier3:
		return "TIER_3"
	}
	return strconv.FormatInt(int64(v), 10)
}

type PremiumUsageFlags uint64

const (
	PremiumUsageFlagsPremiumDiscriminator PremiumUsageFlags = 1 << 0
	PremiumUsageFlagsAnimatedAvatar       PremiumUsageFlags = 1 << 1
	PremiumUsageFlagsProfileBanner        PremiumUsageFlags = 1 << 2
)

func (v PremiumUsageFlags) String() string {
	switch v {
	case PremiumUsageFlagsPremiumDiscriminator:
		return "PREMIUM_DISCRIMINATOR"
	case PremiumUsageFlagsAnimatedAvatar:
		return "ANIMATED_AVATAR"
	case PremiumUsageFlagsProfileBanner:
		return "PROFILE_BANNER"
	}
	return strconv.FormatUint(uint64(v), 10)
}

type PresenceOperatingSystemType string

const (
	PresenceOperatingSystemTypeWin32  PresenceOperatingSystemType = "win32"
	PresenceOperatingSystemTypeDarwin PresenceOperatingSystemType = "darwin"
	PresenceOperatingSystemTypeLinux  PresenceOperatingSystemType = "linux"
)

type PresenceUpdateType string

const (
	PresenceUpdateTypeStart  PresenceUpdateType = "START"
	PresenceUpdateTypeUpdate PresenceUpdateType = "UPDATE"
	PresenceUpdateTypeStop   PresenceUpdateType = "STOP"
)

type PresentationType string

const (
	PresentationTypeImage    PresentationType = "image"
	PresentationTypeNumber   PresentationType = "number"
	PresentationTypeText     PresentationType = "text"
	PresentationTypeDuration PresentationType = "duration"
)

type PriceSetAssignmentPurchaseType int64

const (
	PriceSetAssignmentPurchaseTypeBase               PriceSetAssignmentPurchaseType = 0
	PriceSetAssignmentPurchaseTypePremiumTier0       PriceSetAssignmentPurchaseType = 1
	PriceSetAssignmentPurchaseTypePremiumTier1       PriceSetAssignmentPurchaseType = 2
	PriceSetAssignmentPurchaseTypePremiumTier2       PriceSetAssignmentPurchaseType = 3
	PriceSetAssignmentPurchaseTypeMobile             PriceSetAssignmentPurchaseType = 4
	PriceSetAssignmentPurchaseTypeMobilePremiumTier2 PriceSetAssignmentPurchaseType = 5
	PriceSetAssignmentPurchaseTypeGift               PriceSetAssignmentPurchaseType = 6
)

func (v PriceSetAssignmentPurchaseType) String() string {
	switch v {
	case PriceSetAssignmentPurchaseTypeBase:
		return "BASE"
	case PriceSetAssignmentPurchaseTypePremiumTier0:
		return "PREMIUM_TIER_0"
	case PriceSetAssignmentPurchaseTypePremiumTier1:
		return "PREMIUM_TIER_1"
	case PriceSetAssignmentPurchaseTypePremiumTier2:
		return "PREMIUM_TIER_2"
	case PriceSetAssignmentPurchaseTypeMobile:
		return "MOBILE"
	case PriceSetAssignmentPurchaseTypeMobilePremiumTier2:
		return "MOBILE_PREMIUM_TIER_2"
	case PriceSetAssignmentPurchaseTypeGift:
		return "GIFT"
	}
	return strconv.FormatInt(int64(v), 10)
}

type PriceTierType int64

const (
	PriceTierTypeGuildRoleSubscriptions PriceTierType = 1
	PriceTierTypeGuildProducts          PriceTierType = 2
)

func (v PriceTierType) String() string {
	switch v {
	case PriceTierTypeGuildRoleSubscriptions:
		return "GUILD_ROLE_SUBSCRIPTIONS"
	case PriceTierTypeGuildProducts:
		return "GUILD_PRODUCTS"
	}
	return strconv.FormatInt(int64(v), 10)
}

type PricingLocalizationStrategy string

const (
	PricingLocalizationStrategyLocalizedPriceSets PricingLocalizationStrategy = "localized_price_sets"
)

type PrivacyLevel int64

const (
	PrivacyLevelPublic    PrivacyLevel = 1
	PrivacyLevelGuildOnly PrivacyLevel = 2
)

func (v PrivacyLevel) String() string {
	switch v {
	case PrivacyLevelPublic:
		return "PUBLIC"
	case PrivacyLevelGuildOnly:
		return "GUILD_ONLY"
	}
	return strconv.FormatInt(int64(v), 10)
}

type ProfileEffectAnimationType int64

const (
	ProfileEffectAnimationTypeUnspecified  ProfileEffectAnimationType = 0
	ProfileEffectAnimationTypePersistent   ProfileEffectAnimationType = 1
	ProfileEffectAnimationTypeIntermittent ProfileEffectAnimationType = 2
)

func (v ProfileEffectAnimationType) String() string {
	switch v {
	case ProfileEffectAnimationTypeUnspecified:
		return "UNSPECIFIED"
	case ProfileEffectAnimationTypePersistent:
		return "PERSISTENT"
	case ProfileEffectAnimationTypeIntermittent:
		return "INTERMITTENT"
	}
	return strconv.FormatInt(int64(v), 10)
}

type ProfileRequestType string

const (
	ProfileRequestTypePopout        ProfileRequestType = "popout"
	ProfileRequestTypeModal         ProfileRequestType = "modal"
	ProfileRequestTypeSidebar       ProfileRequestType = "sidebar"
	ProfileRequestTypeAccountPopout ProfileRequestType = "account_popout"
	ProfileRequestTypeActionSheet   ProfileRequestType = "action_sheet"
	ProfileRequestTypeYouScreen     ProfileRequestType = "you_screen"
)

type ProfileVisibility int64

const (
	ProfileVisibilityUnset                 ProfileVisibility = 0
	ProfileVisibilityFriendsOnly           ProfileVisibility = 1
	ProfileVisibilityFriendsAndSmallGuilds ProfileVisibility = 2
	ProfileVisibilityFriendsAndAllGuilds   ProfileVisibility = 3
)

func (v ProfileVisibility) String() string {
	switch v {
	case ProfileVisibilityUnset:
		return "UNSET"
	case ProfileVisibilityFriendsOnly:
		return "FRIENDS_ONLY"
	case ProfileVisibilityFriendsAndSmallGuilds:
		return "FRIENDS_AND_SMALL_GUILDS"
	case ProfileVisibilityFriendsAndAllGuilds:
		return "FRIENDS_AND_ALL_GUILDS"
	}
	return strconv.FormatInt(int64(v), 10)
}

type PromotionFlags uint64

const (
	PromotionFlagsIsBlockedIos                     PromotionFlags = 1 << 5
	PromotionFlagsIsOutboundRedeemableByTrialUsers PromotionFlags = 1 << 6
	PromotionFlagsSuppressNotification             PromotionFlags = 1 << 7
)

func (v PromotionFlags) String() string {
	switch v {
	case PromotionFlagsIsBlockedIos:
		return "IS_BLOCKED_IOS"
	case PromotionFlagsIsOutboundRedeemableByTrialUsers:
		return "IS_OUTBOUND_REDEEMABLE_BY_TRIAL_USERS"
	case PromotionFlagsSuppressNotification:
		return "SUPPRESS_NOTIFICATION"
	}
	return strconv.FormatUint(uint64(v), 10)
}

type PromotionMarketingComponentType int64

const (
	PromotionMarketingComponentTypeAnnouncementModal   PromotionMarketingComponentType = 0
	PromotionMarketingComponentTypePremiumTab          PromotionMarketingComponentType = 1
	PromotionMarketingComponentTypeMarketingPageBanner PromotionMarketingComponentType = 2
	PromotionMarketingComponentTypePaymentModalBanner  PromotionMarketingComponentType = 3
	PromotionMarketingComponentTypeMobileBottomSheet   PromotionMarketingComponentType = 4
)

func (v PromotionMarketingComponentType) String() string {
	switch v {
	case PromotionMarketingComponentTypeAnnouncementModal:
		return "ANNOUNCEMENT_MODAL"
	case PromotionMarketingComponentTypePremiumTab:
		return "PREMIUM_TAB"
	case PromotionMarketingComponentTypeMarketingPageBanner:
		return "MARKETING_PAGE_BANNER"
	case PromotionMarketingComponentTypePaymentModalBanner:
		return "PAYMENT_MODAL_BANNER"
	case PromotionMarketingComponentTypeMobileBottomSheet:
		return "MOBILE_BOTTOM_SHEET"
	}
	return strconv.FormatInt(int64(v), 10)
}

type PromotionPartnerType string

const (
	PromotionPartnerTypeSteelseries  PromotionPartnerType = "steelseries"
	PromotionPartnerTypeKontrolfreek PromotionPartnerType = "kontrolfreek"
)

type PromotionPlatformType int64

const (
	PromotionPlatformTypeDesktop PromotionPlatformType = 0
	PromotionPlatformTypeMobile  PromotionPlatformType = 1
)

func (v PromotionPlatformType) String() string {
	switch v {
	case PromotionPlatformTypeDesktop:
		return "DESKTOP"
	case PromotionPlatformTypeMobile:
		return "MOBILE"
	}
	return strconv.FormatInt(int64(v), 10)
}

type PromotionType int64

const (
	PromotionTypeThirdParty         PromotionType = 0
	PromotionTypeBogo               PromotionType = 1
	PromotionTypeThirdPartyInbound  PromotionType = 3
	PromotionTypeThirdPartyOutbound PromotionType = 4
	PromotionTypeMarketingMoment    PromotionType = 5
	PromotionTypeGiftPromotion      PromotionType = 6
)

func (v PromotionType) String() string {
	switch v {
	case PromotionTypeThirdParty:
		return "THIRD_PARTY"
	case PromotionTypeBogo:
		return "BOGO"
	case PromotionTypeThirdPartyInbound:
		return "THIRD_PARTY_INBOUND"
	case PromotionTypeThirdPartyOutbound:
		return "THIRD_PARTY_OUTBOUND"
	case PromotionTypeMarketingMoment:
		return "MARKETING_MOMENT"
	case PromotionTypeGiftPromotion:
		return "GIFT_PROMOTION"
	}
	return strconv.FormatInt(int64(v), 10)
}

type PromptBehavior string

const (
	PromptBehaviorNone    PromptBehavior = "none"
	PromptBehaviorConsent PromptBehavior = "consent"
)

type ProtocolType string

const (
	ProtocolTypeUdp       ProtocolType = "udp"
	ProtocolTypeWebrtc    ProtocolType = "webrtc"
	ProtocolTypeWebrtcP2p ProtocolType = "webrtc-p2p"
)

type PurchasedFlags uint64

const (
	PurchasedFlagsNitroClassic   PurchasedFlags = 1 << 0
	PurchasedFlagsNitro          PurchasedFlags = 1 << 1
	PurchasedFlagsGuildBoost     PurchasedFlags = 1 << 2
	PurchasedFlagsNitroBasic     PurchasedFlags = 1 << 3
	PurchasedFlagsOnReverseTrial PurchasedFlags = 1 << 4
)

func (v PurchasedFlags) String() string {
	switch v {
	case PurchasedFlagsNitroClassic:
		return "NITRO_CLASSIC"
	case PurchasedFlagsNitro:
		return "NITRO"
	case PurchasedFlagsGuildBoost:
		return "GUILD_BOOST"
	case PurchasedFlagsNitroBasic:
		return "NITRO_BASIC"
	case PurchasedFlagsOnReverseTrial:
		return "ON_REVERSE_TRIAL"
	}
	return strconv.FormatUint(uint64(v), 10)
}

type PushNotificationProvider string

const (
	PushNotificationProviderGcm              PushNotificationProvider = "gcm"
	PushNotificationProviderMetaHorizon      PushNotificationProvider = "meta_horizon"
	PushNotificationProviderApns             PushNotificationProvider = "apns"
	PushNotificationProviderApnsInternal     PushNotificationProvider = "apns_internal"
	PushNotificationProviderApnsLocal        PushNotificationProvider = "apns_local"
	PushNotificationProviderApnsVoip         PushNotificationProvider = "apns_voip"
	PushNotificationProviderApnsInternalVoip PushNotificationProvider = "apns_internal_voip"
	PushNotificationProviderApnsLocalVoip    PushNotificationProvider = "apns_local_voip"
)

type PushNotificationType string

const (
	PushNotificationTypeMessageCreate                     PushNotificationType = "MESSAGE_CREATE"
	PushNotificationTypeForumThreadCreated                PushNotificationType = "FORUM_THREAD_CREATED"
	PushNotificationTypeRelationshipAdd                   PushNotificationType = "RELATIONSHIP_ADD"
	PushNotificationTypeFriendSuggestionCreate            PushNotificationType = "FRIEND_SUGGESTION_CREATE"
	PushNotificationTypeCallRing                          PushNotificationType = "CALL_RING"
	PushNotificationTypeCallRingEnd                       PushNotificationType = "CALL_RING_END"
	PushNotificationTypeCallAck                           PushNotificationType = "CALL_ACK"
	PushNotificationTypeCallConnect                       PushNotificationType = "CALL_CONNECT"
	PushNotificationTypeChannelAck                        PushNotificationType = "CHANNEL_ACK"
	PushNotificationTypeActivityStart                     PushNotificationType = "ACTIVITY_START"
	PushNotificationTypeApplicationLibraryInstallComplete PushNotificationType = "APPLICATION_LIBRARY_INSTALL_COMPLETE"
	PushNotificationTypeStageInstanceCreate               PushNotificationType = "STAGE_INSTANCE_CREATE"
	PushNotificationTypeGuildScheduledEventUpdate         PushNotificationType = "GUILD_SCHEDULED_EVENT_UPDATE"
	PushNotificationTypeGuildStreamStart                  PushNotificationType = "GUILD_STREAM_START"
	PushNotificationTypeGenericPushNotificationSent       PushNotificationType = "GENERIC_PUSH_NOTIFICATION_SENT"
)

type QuestContentType int64

const (
	QuestContentTypeGiftInventorySettingsBadge        QuestContentType = 0
	QuestContentTypeQuestBar                          QuestContentType = 1
	QuestContentTypeQuestInventoryCard                QuestContentType = 2
	QuestContentTypeQuestsEmbed                       QuestContentType = 3
	QuestContentTypeActivityPanel                     QuestContentType = 4
	QuestContentTypeQuestLiveStream                   QuestContentType = 5
	QuestContentTypeMembersList                       QuestContentType = 6
	QuestContentTypeQuestBadge                        QuestContentType = 7
	QuestContentTypeGiftInventoryForYou               QuestContentType = 8
	QuestContentTypeGiftInventoryOther                QuestContentType = 9
	QuestContentTypeQuestBarV2                        QuestContentType = 10
	QuestContentTypeQuestHomeDesktop                  QuestContentType = 11
	QuestContentTypeQuestHomeMobile                   QuestContentType = 12
	QuestContentTypeQuestBarMobile                    QuestContentType = 13
	QuestContentTypeThirdPartyApp                     QuestContentType = 14
	QuestContentTypeQuestBottomSheet                  QuestContentType = 15
	QuestContentTypeQuestEmbedMobile                  QuestContentType = 16
	QuestContentTypeQuestHomeMoveCallout              QuestContentType = 17
	QuestContentTypeDiscoverySidebar                  QuestContentType = 18
	QuestContentTypeQuestShareLink                    QuestContentType = 19
	QuestContentTypeConnectionsModal                  QuestContentType = 20
	QuestContentTypeDiscoveryCompass                  QuestContentType = 21
	QuestContentTypeTrophyCaseCard                    QuestContentType = 22
	QuestContentTypeVideoModal                        QuestContentType = 23
	QuestContentTypeVideoModalEndCard                 QuestContentType = 24
	QuestContentTypeRewardModal                       QuestContentType = 25
	QuestContentTypeExcludedQuestEmbed                QuestContentType = 26
	QuestContentTypeVideoModalMobile                  QuestContentType = 27
	QuestContentTypeOrbsAnnouncementModal             QuestContentType = 28
	QuestContentTypeOrbsBalanceMenu                   QuestContentType = 29
	QuestContentTypeQuestEnrollmentBlockedBottomSheet QuestContentType = 30
	QuestContentTypeOrbsShopHeroCta                   QuestContentType = 31
	QuestContentTypeQuestEnrollmentBlockedModal       QuestContentType = 32
	QuestContentTypeInternalPreviewTool               QuestContentType = 33
	QuestContentTypeOrbsReheatCoachmarkCta            QuestContentType = 34
	QuestContentTypeInvalidQuestEmbed                 QuestContentType = 35
	QuestContentTypeNotShareableQuestEmbed            QuestContentType = 36
	QuestContentTypeQuestHomeMoveCalloutDiscover      QuestContentType = 37
	QuestContentTypeSponsoredQuestSheet               QuestContentType = 38
	QuestContentTypeMobileOrbsOnboardingDc            QuestContentType = 39
	QuestContentTypeRunningActivity                   QuestContentType = 40
	QuestContentTypeVideoModalPrimaryCta              QuestContentType = 41
	QuestContentTypeQuestHomeTakeover                 QuestContentType = 42
	QuestContentTypeUserProfileActivity               QuestContentType = 43
	QuestContentTypeMembersListCard                   QuestContentType = 44
	QuestContentTypeAppLauncher                       QuestContentType = 45
	QuestContentTypeActivitySuggestion                QuestContentType = 46
	QuestContentTypeQuestHomeEntrypoint               QuestContentType = 47
	QuestContentTypeQuestHomeEntrypointThemed         QuestContentType = 48
	QuestContentTypeQuestActivityUnenrolledModal      QuestContentType = 49
	QuestContentTypeQuestHomeHero                     QuestContentType = 50
	QuestContentTypeQuestActivityHeader               QuestContentType = 51
	QuestContentTypeUserProfileHeader                 QuestContentType = 52
	QuestContentTypeUserSettings                      QuestContentType = 53
	QuestContentTypeNitroHomePerkCard                 QuestContentType = 54
	QuestContentTypeQuestHomeMobileCarousel           QuestContentType = 55
	QuestContentTypeQuestHomeHeroShelf                QuestContentType = 56
	QuestContentTypeVideoModalIconEndCard             QuestContentType = 57
	QuestContentTypeAchievementInGameModal            QuestContentType = 58
	QuestContentTypeQuestHomeFeaturedSection          QuestContentType = 59
	QuestContentTypeQuestHomeInProgressSection        QuestContentType = 60
	QuestContentTypeQuestHomeEndingSoonSection        QuestContentType = 61
	QuestContentTypeQuestHomeOrbSection               QuestContentType = 62
	QuestContentTypeQuestHomeDiscoveredSection        QuestContentType = 63
	QuestContentTypeQuestHomeSearchResult             QuestContentType = 64
	QuestContentTypePlayQuestModal                    QuestContentType = 65
	QuestContentTypeVideoModalMobileFooter            QuestContentType = 66
	QuestContentTypeQuestHomeEntrypointMobile         QuestContentType = 67
	QuestContentTypeBountiesEndInterstitial           QuestContentType = 68
	QuestContentTypeQuestHomeExpiredSection           QuestContentType = 69
	QuestContentTypeQuestHomePreviewSection           QuestContentType = 70
	QuestContentTypeSocialLayerStorefront             QuestContentType = 71
	QuestContentTypeQuestHomeSpecialQuestsSection     QuestContentType = 72
)

func (v QuestContentType) String() string {
	switch v {
	case QuestContentTypeGiftInventorySettingsBadge:
		return "GIFT_INVENTORY_SETTINGS_BADGE"
	case QuestContentTypeQuestBar:
		return "QUEST_BAR"
	case QuestContentTypeQuestInventoryCard:
		return "QUEST_INVENTORY_CARD"
	case QuestContentTypeQuestsEmbed:
		return "QUESTS_EMBED"
	case QuestContentTypeActivityPanel:
		return "ACTIVITY_PANEL"
	case QuestContentTypeQuestLiveStream:
		return "QUEST_LIVE_STREAM"
	case QuestContentTypeMembersList:
		return "MEMBERS_LIST"
	case QuestContentTypeQuestBadge:
		return "QUEST_BADGE"
	case QuestContentTypeGiftInventoryForYou:
		return "GIFT_INVENTORY_FOR_YOU"
	case QuestContentTypeGiftInventoryOther:
		return "GIFT_INVENTORY_OTHER"
	case QuestContentTypeQuestBarV2:
		return "QUEST_BAR_V2"
	case QuestContentTypeQuestHomeDesktop:
		return "QUEST_HOME_DESKTOP"
	case QuestContentTypeQuestHomeMobile:
		return "QUEST_HOME_MOBILE"
	case QuestContentTypeQuestBarMobile:
		return "QUEST_BAR_MOBILE"
	case QuestContentTypeThirdPartyApp:
		return "THIRD_PARTY_APP"
	case QuestContentTypeQuestBottomSheet:
		return "QUEST_BOTTOM_SHEET"
	case QuestContentTypeQuestEmbedMobile:
		return "QUEST_EMBED_MOBILE"
	case QuestContentTypeQuestHomeMoveCallout:
		return "QUEST_HOME_MOVE_CALLOUT"
	case QuestContentTypeDiscoverySidebar:
		return "DISCOVERY_SIDEBAR"
	case QuestContentTypeQuestShareLink:
		return "QUEST_SHARE_LINK"
	case QuestContentTypeConnectionsModal:
		return "CONNECTIONS_MODAL"
	case QuestContentTypeDiscoveryCompass:
		return "DISCOVERY_COMPASS"
	case QuestContentTypeTrophyCaseCard:
		return "TROPHY_CASE_CARD"
	case QuestContentTypeVideoModal:
		return "VIDEO_MODAL"
	case QuestContentTypeVideoModalEndCard:
		return "VIDEO_MODAL_END_CARD"
	case QuestContentTypeRewardModal:
		return "REWARD_MODAL"
	case QuestContentTypeExcludedQuestEmbed:
		return "EXCLUDED_QUEST_EMBED"
	case QuestContentTypeVideoModalMobile:
		return "VIDEO_MODAL_MOBILE"
	case QuestContentTypeOrbsAnnouncementModal:
		return "ORBS_ANNOUNCEMENT_MODAL"
	case QuestContentTypeOrbsBalanceMenu:
		return "ORBS_BALANCE_MENU"
	case QuestContentTypeQuestEnrollmentBlockedBottomSheet:
		return "QUEST_ENROLLMENT_BLOCKED_BOTTOM_SHEET"
	case QuestContentTypeOrbsShopHeroCta:
		return "ORBS_SHOP_HERO_CTA"
	case QuestContentTypeQuestEnrollmentBlockedModal:
		return "QUEST_ENROLLMENT_BLOCKED_MODAL"
	case QuestContentTypeInternalPreviewTool:
		return "INTERNAL_PREVIEW_TOOL"
	case QuestContentTypeOrbsReheatCoachmarkCta:
		return "ORBS_REHEAT_COACHMARK_CTA"
	case QuestContentTypeInvalidQuestEmbed:
		return "INVALID_QUEST_EMBED"
	case QuestContentTypeNotShareableQuestEmbed:
		return "NOT_SHAREABLE_QUEST_EMBED"
	case QuestContentTypeQuestHomeMoveCalloutDiscover:
		return "QUEST_HOME_MOVE_CALLOUT_DISCOVER"
	case QuestContentTypeSponsoredQuestSheet:
		return "SPONSORED_QUEST_SHEET"
	case QuestContentTypeMobileOrbsOnboardingDc:
		return "MOBILE_ORBS_ONBOARDING_DC"
	case QuestContentTypeRunningActivity:
		return "RUNNING_ACTIVITY"
	case QuestContentTypeVideoModalPrimaryCta:
		return "VIDEO_MODAL_PRIMARY_CTA"
	case QuestContentTypeQuestHomeTakeover:
		return "QUEST_HOME_TAKEOVER"
	case QuestContentTypeUserProfileActivity:
		return "USER_PROFILE_ACTIVITY"
	case QuestContentTypeMembersListCard:
		return "MEMBERS_LIST_CARD"
	case QuestContentTypeAppLauncher:
		return "APP_LAUNCHER"
	case QuestContentTypeActivitySuggestion:
		return "ACTIVITY_SUGGESTION"
	case QuestContentTypeQuestHomeEntrypoint:
		return "QUEST_HOME_ENTRYPOINT"
	case QuestContentTypeQuestHomeEntrypointThemed:
		return "QUEST_HOME_ENTRYPOINT_THEMED"
	case QuestContentTypeQuestActivityUnenrolledModal:
		return "QUEST_ACTIVITY_UNENROLLED_MODAL"
	case QuestContentTypeQuestHomeHero:
		return "QUEST_HOME_HERO"
	case QuestContentTypeQuestActivityHeader:
		return "QUEST_ACTIVITY_HEADER"
	case QuestContentTypeUserProfileHeader:
		return "USER_PROFILE_HEADER"
	case QuestContentTypeUserSettings:
		return "USER_SETTINGS"
	case QuestContentTypeNitroHomePerkCard:
		return "NITRO_HOME_PERK_CARD"
	case QuestContentTypeQuestHomeMobileCarousel:
		return "QUEST_HOME_MOBILE_CAROUSEL"
	case QuestContentTypeQuestHomeHeroShelf:
		return "QUEST_HOME_HERO_SHELF"
	case QuestContentTypeVideoModalIconEndCard:
		return "VIDEO_MODAL_ICON_END_CARD"
	case QuestContentTypeAchievementInGameModal:
		return "ACHIEVEMENT_IN_GAME_MODAL"
	case QuestContentTypeQuestHomeFeaturedSection:
		return "QUEST_HOME_FEATURED_SECTION"
	case QuestContentTypeQuestHomeInProgressSection:
		return "QUEST_HOME_IN_PROGRESS_SECTION"
	case QuestContentTypeQuestHomeEndingSoonSection:
		return "QUEST_HOME_ENDING_SOON_SECTION"
	case QuestContentTypeQuestHomeOrbSection:
		return "QUEST_HOME_ORB_SECTION"
	case QuestContentTypeQuestHomeDiscoveredSection:
		return "QUEST_HOME_DISCOVERED_SECTION"
	case QuestContentTypeQuestHomeSearchResult:
		return "QUEST_HOME_SEARCH_RESULT"
	case QuestContentTypePlayQuestModal:
		return "PLAY_QUEST_MODAL"
	case QuestContentTypeVideoModalMobileFooter:
		return "VIDEO_MODAL_MOBILE_FOOTER"
	case QuestContentTypeQuestHomeEntrypointMobile:
		return "QUEST_HOME_ENTRYPOINT_MOBILE"
	case QuestContentTypeBountiesEndInterstitial:
		return "BOUNTIES_END_INTERSTITIAL"
	case QuestContentTypeQuestHomeExpiredSection:
		return "QUEST_HOME_EXPIRED_SECTION"
	case QuestContentTypeQuestHomePreviewSection:
		return "QUEST_HOME_PREVIEW_SECTION"
	case QuestContentTypeSocialLayerStorefront:
		return "SOCIAL_LAYER_STOREFRONT"
	case QuestContentTypeQuestHomeSpecialQuestsSection:
		return "QUEST_HOME_SPECIAL_QUESTS_SECTION"
	}
	return strconv.FormatInt(int64(v), 10)
}

type QuestFeature int64

const (
	QuestFeaturePostEnrollmentCta                QuestFeature = 1
	QuestFeaturePlaytimeCriteria                 QuestFeature = 2
	QuestFeatureQuestBarV2                       QuestFeature = 3
	QuestFeatureExcludeMinors                    QuestFeature = 4
	QuestFeatureExcludeRussia                    QuestFeature = 5
	QuestFeatureInHouseConsoleQuest              QuestFeature = 6
	QuestFeatureMobileConsoleQuest               QuestFeature = 7
	QuestFeatureStartQuestCta                    QuestFeature = 8
	QuestFeatureRewardHighlighting               QuestFeature = 9
	QuestFeatureFractionsQuest                   QuestFeature = 10
	QuestFeatureAdditionalRedemptionInstructions QuestFeature = 11
	QuestFeaturePacingV2                         QuestFeature = 12
	QuestFeatureDismissalSurvey                  QuestFeature = 13
	QuestFeatureMobileQuestDock                  QuestFeature = 14
	QuestFeatureQuestsCDN                        QuestFeature = 15
	QuestFeaturePacingController                 QuestFeature = 16
	QuestFeatureQuestHomeForceStaticImage        QuestFeature = 17
	QuestFeatureVideoQuestForceHlsVideo          QuestFeature = 18
	QuestFeatureVideoQuestForceEndCardCtaSwap    QuestFeature = 19
	QuestFeatureExperimentalTargetingTraits      QuestFeature = 20
	QuestFeatureDoNotDisplay                     QuestFeature = 21
	QuestFeatureExternalDialog                   QuestFeature = 22
	QuestFeatureMobileOnlyQuestPushToMobile      QuestFeature = 23
	QuestFeatureManualHeartbeatInitialization    QuestFeature = 24
	QuestFeatureCloudGamingActivity              QuestFeature = 25
	QuestFeatureNonGamingPlayQuest               QuestFeature = 26
	QuestFeatureActivityQuestAutoEnrollment      QuestFeature = 27
	QuestFeaturePackageActionAdventure           QuestFeature = 28
	QuestFeaturePackageRpgMmo                    QuestFeature = 29
	QuestFeaturePackageRacingSports              QuestFeature = 30
	QuestFeaturePackageSandboxCreative           QuestFeature = 31
	QuestFeaturePackageFamilyFriendly            QuestFeature = 32
	QuestFeaturePackageHolidaySeason             QuestFeature = 33
	QuestFeaturePackageNewYears                  QuestFeature = 34
	QuestFeatureFullEpisodeVideoQuest            QuestFeature = 35
	QuestFeatureMobileActivityQuest              QuestFeature = 36
	QuestFeatureQuestBarUnfurl                   QuestFeature = 37
	QuestFeatureNoPremiumOrbsPerk                QuestFeature = 38
	QuestFeatureNitroControlCta                  QuestFeature = 39
	QuestFeatureNitro2Point0Cta                  QuestFeature = 40
	QuestFeatureOrbsMultiplierQuest              QuestFeature = 41
	QuestFeatureXboxGamePassQuest                QuestFeature = 42
)

func (v QuestFeature) String() string {
	switch v {
	case QuestFeaturePostEnrollmentCta:
		return "POST_ENROLLMENT_CTA"
	case QuestFeaturePlaytimeCriteria:
		return "PLAYTIME_CRITERIA"
	case QuestFeatureQuestBarV2:
		return "QUEST_BAR_V2"
	case QuestFeatureExcludeMinors:
		return "EXCLUDE_MINORS"
	case QuestFeatureExcludeRussia:
		return "EXCLUDE_RUSSIA"
	case QuestFeatureInHouseConsoleQuest:
		return "IN_HOUSE_CONSOLE_QUEST"
	case QuestFeatureMobileConsoleQuest:
		return "MOBILE_CONSOLE_QUEST"
	case QuestFeatureStartQuestCta:
		return "START_QUEST_CTA"
	case QuestFeatureRewardHighlighting:
		return "REWARD_HIGHLIGHTING"
	case QuestFeatureFractionsQuest:
		return "FRACTIONS_QUEST"
	case QuestFeatureAdditionalRedemptionInstructions:
		return "ADDITIONAL_REDEMPTION_INSTRUCTIONS"
	case QuestFeaturePacingV2:
		return "PACING_V2"
	case QuestFeatureDismissalSurvey:
		return "DISMISSAL_SURVEY"
	case QuestFeatureMobileQuestDock:
		return "MOBILE_QUEST_DOCK"
	case QuestFeatureQuestsCDN:
		return "QUESTS_CDN"
	case QuestFeaturePacingController:
		return "PACING_CONTROLLER"
	case QuestFeatureQuestHomeForceStaticImage:
		return "QUEST_HOME_FORCE_STATIC_IMAGE"
	case QuestFeatureVideoQuestForceHlsVideo:
		return "VIDEO_QUEST_FORCE_HLS_VIDEO"
	case QuestFeatureVideoQuestForceEndCardCtaSwap:
		return "VIDEO_QUEST_FORCE_END_CARD_CTA_SWAP"
	case QuestFeatureExperimentalTargetingTraits:
		return "EXPERIMENTAL_TARGETING_TRAITS"
	case QuestFeatureDoNotDisplay:
		return "DO_NOT_DISPLAY"
	case QuestFeatureExternalDialog:
		return "EXTERNAL_DIALOG"
	case QuestFeatureMobileOnlyQuestPushToMobile:
		return "MOBILE_ONLY_QUEST_PUSH_TO_MOBILE"
	case QuestFeatureManualHeartbeatInitialization:
		return "MANUAL_HEARTBEAT_INITIALIZATION"
	case QuestFeatureCloudGamingActivity:
		return "CLOUD_GAMING_ACTIVITY"
	case QuestFeatureNonGamingPlayQuest:
		return "NON_GAMING_PLAY_QUEST"
	case QuestFeatureActivityQuestAutoEnrollment:
		return "ACTIVITY_QUEST_AUTO_ENROLLMENT"
	case QuestFeaturePackageActionAdventure:
		return "PACKAGE_ACTION_ADVENTURE"
	case QuestFeaturePackageRpgMmo:
		return "PACKAGE_RPG_MMO"
	case QuestFeaturePackageRacingSports:
		return "PACKAGE_RACING_SPORTS"
	case QuestFeaturePackageSandboxCreative:
		return "PACKAGE_SANDBOX_CREATIVE"
	case QuestFeaturePackageFamilyFriendly:
		return "PACKAGE_FAMILY_FRIENDLY"
	case QuestFeaturePackageHolidaySeason:
		return "PACKAGE_HOLIDAY_SEASON"
	case QuestFeaturePackageNewYears:
		return "PACKAGE_NEW_YEARS"
	case QuestFeatureFullEpisodeVideoQuest:
		return "FULL_EPISODE_VIDEO_QUEST"
	case QuestFeatureMobileActivityQuest:
		return "MOBILE_ACTIVITY_QUEST"
	case QuestFeatureQuestBarUnfurl:
		return "QUEST_BAR_UNFURL"
	case QuestFeatureNoPremiumOrbsPerk:
		return "NO_PREMIUM_ORBS_PERK"
	case QuestFeatureNitroControlCta:
		return "NITRO_CONTROL_CTA"
	case QuestFeatureNitro2Point0Cta:
		return "NITRO_2_POINT_0_CTA"
	case QuestFeatureOrbsMultiplierQuest:
		return "ORBS_MULTIPLIER_QUEST"
	case QuestFeatureXboxGamePassQuest:
		return "XBOX_GAME_PASS_QUEST"
	}
	return strconv.FormatInt(int64(v), 10)
}

type QuestHomeFilterType string

const (
	QuestHomeFilterTypeRewardVirtualCurrency QuestHomeFilterType = "reward_virtual_currency"
	QuestHomeFilterTypeRewardCollectible     QuestHomeFilterType = "reward_collectible"
	QuestHomeFilterTypeRewardInGame          QuestHomeFilterType = "reward_in_game"
	QuestHomeFilterTypeTaskPlay              QuestHomeFilterType = "task_play"
	QuestHomeFilterTypeTaskVideo             QuestHomeFilterType = "task_video"
)

type QuestHomeSortType string

const (
	QuestHomeSortTypeSuggested        QuestHomeSortType = "suggested"
	QuestHomeSortTypeMostRecent       QuestHomeSortType = "most_recent"
	QuestHomeSortTypeExpiringSoon     QuestHomeSortType = "expiring_soon"
	QuestHomeSortTypeRecentlyEnrolled QuestHomeSortType = "recently_enrolled"
)

type QuestHomeTab string

const (
	QuestHomeTabAll     QuestHomeTab = "all"
	QuestHomeTabClaimed QuestHomeTab = "claimed"
)

type QuestPlacementArea int64

const (
	QuestPlacementAreaDesktopAccountPanelArea QuestPlacementArea = 1
	QuestPlacementAreaMobileHomeDockArea      QuestPlacementArea = 2
	QuestPlacementAreaQuestHomeBannerDesktop  QuestPlacementArea = 3
	QuestPlacementAreaQuestHomeMobileCarousel QuestPlacementArea = 4
	QuestPlacementAreaVideoModalMobile        QuestPlacementArea = 5
)

func (v QuestPlacementArea) String() string {
	switch v {
	case QuestPlacementAreaDesktopAccountPanelArea:
		return "DESKTOP_ACCOUNT_PANEL_AREA"
	case QuestPlacementAreaMobileHomeDockArea:
		return "MOBILE_HOME_DOCK_AREA"
	case QuestPlacementAreaQuestHomeBannerDesktop:
		return "QUEST_HOME_BANNER_DESKTOP"
	case QuestPlacementAreaQuestHomeMobileCarousel:
		return "QUEST_HOME_MOBILE_CAROUSEL"
	case QuestPlacementAreaVideoModalMobile:
		return "VIDEO_MODAL_MOBILE"
	}
	return strconv.FormatInt(int64(v), 10)
}

type QuestPlatformType int64

const (
	QuestPlatformTypeCrossPlatform QuestPlatformType = 0
	QuestPlatformTypeXbox          QuestPlatformType = 1
	QuestPlatformTypePlaystation   QuestPlatformType = 2
	QuestPlatformTypeSwitch        QuestPlatformType = 3
	QuestPlatformTypePC            QuestPlatformType = 4
)

func (v QuestPlatformType) String() string {
	switch v {
	case QuestPlatformTypeCrossPlatform:
		return "CROSS_PLATFORM"
	case QuestPlatformTypeXbox:
		return "XBOX"
	case QuestPlatformTypePlaystation:
		return "PLAYSTATION"
	case QuestPlatformTypeSwitch:
		return "SWITCH"
	case QuestPlatformTypePC:
		return "PC"
	}
	return strconv.FormatInt(int64(v), 10)
}

type QuestRewardAssignmentMethod int64

const (
	QuestRewardAssignmentMethodAll    QuestRewardAssignmentMethod = 1
	QuestRewardAssignmentMethodTiered QuestRewardAssignmentMethod = 2
)

func (v QuestRewardAssignmentMethod) String() string {
	switch v {
	case QuestRewardAssignmentMethodAll:
		return "ALL"
	case QuestRewardAssignmentMethodTiered:
		return "TIERED"
	}
	return strconv.FormatInt(int64(v), 10)
}

type QuestRewardExpirationMode int64

const (
	QuestRewardExpirationModeNormal           QuestRewardExpirationMode = 1
	QuestRewardExpirationModePremiumExtension QuestRewardExpirationMode = 2
	QuestRewardExpirationModePremiumPermanent QuestRewardExpirationMode = 3
)

func (v QuestRewardExpirationMode) String() string {
	switch v {
	case QuestRewardExpirationModeNormal:
		return "NORMAL"
	case QuestRewardExpirationModePremiumExtension:
		return "PREMIUM_EXTENSION"
	case QuestRewardExpirationModePremiumPermanent:
		return "PREMIUM_PERMANENT"
	}
	return strconv.FormatInt(int64(v), 10)
}

type QuestRewardType int64

const (
	QuestRewardTypeRewardCode        QuestRewardType = 1
	QuestRewardTypeInGame            QuestRewardType = 2
	QuestRewardTypeCollectible       QuestRewardType = 3
	QuestRewardTypeVirtualCurrency   QuestRewardType = 4
	QuestRewardTypeFractionalPremium QuestRewardType = 5
)

func (v QuestRewardType) String() string {
	switch v {
	case QuestRewardTypeRewardCode:
		return "REWARD_CODE"
	case QuestRewardTypeInGame:
		return "IN_GAME"
	case QuestRewardTypeCollectible:
		return "COLLECTIBLE"
	case QuestRewardTypeVirtualCurrency:
		return "VIRTUAL_CURRENCY"
	case QuestRewardTypeFractionalPremium:
		return "FRACTIONAL_PREMIUM"
	}
	return strconv.FormatInt(int64(v), 10)
}

type QuestSharePolicy string

const (
	QuestSharePolicyShareableEverywhere QuestSharePolicy = "shareable_everywhere"
	QuestSharePolicyNotShareable        QuestSharePolicy = "not_shareable"
)

type QuestTaskType string

const (
	QuestTaskTypeStreamOnDesktop       QuestTaskType = "STREAM_ON_DESKTOP"
	QuestTaskTypePlayOnDesktop         QuestTaskType = "PLAY_ON_DESKTOP"
	QuestTaskTypePlayOnDesktopV2       QuestTaskType = "PLAY_ON_DESKTOP_V2"
	QuestTaskTypePlayOnXbox            QuestTaskType = "PLAY_ON_XBOX"
	QuestTaskTypePlayOnPlaystation     QuestTaskType = "PLAY_ON_PLAYSTATION"
	QuestTaskTypeWatchVideo            QuestTaskType = "WATCH_VIDEO"
	QuestTaskTypeWatchVideoOnMobile    QuestTaskType = "WATCH_VIDEO_ON_MOBILE"
	QuestTaskTypePlayActivity          QuestTaskType = "PLAY_ACTIVITY"
	QuestTaskTypeAchievementInGame     QuestTaskType = "ACHIEVEMENT_IN_GAME"
	QuestTaskTypeAchievementInActivity QuestTaskType = "ACHIEVEMENT_IN_ACTIVITY"
)

type RPCApplicationState int64

const (
	RPCApplicationStateDisabled    RPCApplicationState = 0
	RPCApplicationStateUnsubmitted RPCApplicationState = 1
	RPCApplicationStateSubmitted   RPCApplicationState = 2
	RPCApplicationStateApproved    RPCApplicationState = 3
	RPCApplicationStateRejected    RPCApplicationState = 4
)

func (v RPCApplicationState) String() string {
	switch v {
	case RPCApplicationStateDisabled:
		return "DISABLED"
	case RPCApplicationStateUnsubmitted:
		return "UNSUBMITTED"
	case RPCApplicationStateSubmitted:
		return "SUBMITTED"
	case RPCApplicationStateApproved:
		return "APPROVED"
	case RPCApplicationStateRejected:
		return "REJECTED"
	}
	return strconv.FormatInt(int64(v), 10)
}

type RPCGuildTemplateState string

const (
	RPCGuildTemplateStateResolving RPCGuildTemplateState = "RESOLVING"
	RPCGuildTemplateStateResolved  RPCGuildTemplateState = "RESOLVED"
	RPCGuildTemplateStateExpired   RPCGuildTemplateState = "EXPIRED"
	RPCGuildTemplateStateAccepted  RPCGuildTemplateState = "ACCEPTED"
	RPCGuildTemplateStateAccepting RPCGuildTemplateState = "ACCEPTING"
)

type RPCScopes string

const (
	RPCScopesRPCAuthenticated  RPCScopes = "rpc.authenticated"
	RPCScopesRPCLocal          RPCScopes = "rpc.local"
	RPCScopesRPCPrivate        RPCScopes = "rpc.private"
	RPCScopesRPCPrivateLimited RPCScopes = "rpc.private.limited"
	RPCScopesRPCEmbeddedApp    RPCScopes = "rpc.embedded_app"
)

type RPCVoiceSettingsModeType string

const (
	RPCVoiceSettingsModeTypePushToTalk    RPCVoiceSettingsModeType = "PUSH_TO_TALK"
	RPCVoiceSettingsModeTypeVoiceActivity RPCVoiceSettingsModeType = "VOICE_ACTIVITY"
)

type ReactionNotificationType int64

const (
	ReactionNotificationTypeNotificationsEnabled  ReactionNotificationType = 0
	ReactionNotificationTypeOnlyDMs               ReactionNotificationType = 1
	ReactionNotificationTypeNotificationsDisabled ReactionNotificationType = 2
)

func (v ReactionNotificationType) String() string {
	switch v {
	case ReactionNotificationTypeNotificationsEnabled:
		return "NOTIFICATIONS_ENABLED"
	case ReactionNotificationTypeOnlyDMs:
		return "ONLY_DMS"
	case ReactionNotificationTypeNotificationsDisabled:
		return "NOTIFICATIONS_DISABLED"
	}
	return strconv.FormatInt(int64(v), 10)
}

type ReactionType int64

const (
	ReactionTypeNormal ReactionType = 0
	ReactionTypeBurst  ReactionType = 1
)

func (v ReactionType) String() string {
	switch v {
	case ReactionTypeNormal:
		return "NORMAL"
	case ReactionTypeBurst:
		return "BURST"
	}
	return strconv.FormatInt(int64(v), 10)
}

type ReadStateFlags uint64

const (
	ReadStateFlagsIsGuildChannel         ReadStateFlags = 1 << 0
	ReadStateFlagsIsThread               ReadStateFlags = 1 << 1
	ReadStateFlagsIsMentionLowImportance ReadStateFlags = 1 << 2
)

func (v ReadStateFlags) String() string {
	switch v {
	case ReadStateFlagsIsGuildChannel:
		return "IS_GUILD_CHANNEL"
	case ReadStateFlagsIsThread:
		return "IS_THREAD"
	case ReadStateFlagsIsMentionLowImportance:
		return "IS_MENTION_LOW_IMPORTANCE"
	}
	return strconv.FormatUint(uint64(v), 10)
}

type ReadStateType int64

const (
	ReadStateTypeChannel                 ReadStateType = 0
	ReadStateTypeGuildEvent              ReadStateType = 1
	ReadStateTypeNotificationCenter      ReadStateType = 2
	ReadStateTypeGuildHome               ReadStateType = 3
	ReadStateTypeGuildOnboardingQuestion ReadStateType = 4
	ReadStateTypeMessageRequests         ReadStateType = 5
)

func (v ReadStateType) String() string {
	switch v {
	case ReadStateTypeChannel:
		return "CHANNEL"
	case ReadStateTypeGuildEvent:
		return "GUILD_EVENT"
	case ReadStateTypeNotificationCenter:
		return "NOTIFICATION_CENTER"
	case ReadStateTypeGuildHome:
		return "GUILD_HOME"
	case ReadStateTypeGuildOnboardingQuestion:
		return "GUILD_ONBOARDING_QUESTION"
	case ReadStateTypeMessageRequests:
		return "MESSAGE_REQUESTS"
	}
	return strconv.FormatInt(int64(v), 10)
}

type RecipientFlags uint64

const (
	RecipientFlagsDismissedInGameMessageNux RecipientFlags = 1 << 0
)

func (v RecipientFlags) String() string {
	switch v {
	case RecipientFlagsDismissedInGameMessageNux:
		return "DISMISSED_IN_GAME_MESSAGE_NUX"
	}
	return strconv.FormatUint(uint64(v), 10)
}

type RefundDisqualificationReason int64

const (
	RefundDisqualificationReasonOther                       RefundDisqualificationReason = 0
	RefundDisqualificationReasonAlreadyRefunded             RefundDisqualificationReason = 1
	RefundDisqualificationReasonNotUserRefundableType       RefundDisqualificationReason = 2
	RefundDisqualificationReasonPastRefundableDate          RefundDisqualificationReason = 3
	RefundDisqualificationReasonEntitlementAlreadyConsumed  RefundDisqualificationReason = 4
	RefundDisqualificationReasonAlreadyRefundedPremium      RefundDisqualificationReason = 5
	RefundDisqualificationReasonAlreadyRefundedPremiumGuild RefundDisqualificationReason = 6
)

func (v RefundDisqualificationReason) String() string {
	switch v {
	case RefundDisqualificationReasonOther:
		return "OTHER"
	case RefundDisqualificationReasonAlreadyRefunded:
		return "ALREADY_REFUNDED"
	case RefundDisqualificationReasonNotUserRefundableType:
		return "NOT_USER_REFUNDABLE_TYPE"
	case RefundDisqualificationReasonPastRefundableDate:
		return "PAST_REFUNDABLE_DATE"
	case RefundDisqualificationReasonEntitlementAlreadyConsumed:
		return "ENTITLEMENT_ALREADY_CONSUMED"
	case RefundDisqualificationReasonAlreadyRefundedPremium:
		return "ALREADY_REFUNDED_PREMIUM"
	case RefundDisqualificationReasonAlreadyRefundedPremiumGuild:
		return "ALREADY_REFUNDED_PREMIUM_GUILD"
	}
	return strconv.FormatInt(int64(v), 10)
}

type RelationshipRemovalFilter int64

const (
	RelationshipRemovalFilterSpam    RelationshipRemovalFilter = 1
	RelationshipRemovalFilterIgnored RelationshipRemovalFilter = 2
)

func (v RelationshipRemovalFilter) String() string {
	switch v {
	case RelationshipRemovalFilterSpam:
		return "SPAM"
	case RelationshipRemovalFilterIgnored:
		return "IGNORED"
	}
	return strconv.FormatInt(int64(v), 10)
}

type RelationshipType int64

const (
	RelationshipTypeNone            RelationshipType = 0
	RelationshipTypeFriend          RelationshipType = 1
	RelationshipTypeBlocked         RelationshipType = 2
	RelationshipTypeIncomingRequest RelationshipType = 3
	RelationshipTypeOutgoingRequest RelationshipType = 4
	RelationshipTypeImplicit        RelationshipType = 5
	RelationshipTypeSuggestion      RelationshipType = 6
)

func (v RelationshipType) String() string {
	switch v {
	case RelationshipTypeNone:
		return "NONE"
	case RelationshipTypeFriend:
		return "FRIEND"
	case RelationshipTypeBlocked:
		return "BLOCKED"
	case RelationshipTypeIncomingRequest:
		return "INCOMING_REQUEST"
	case RelationshipTypeOutgoingRequest:
		return "OUTGOING_REQUEST"
	case RelationshipTypeImplicit:
		return "IMPLICIT"
	case RelationshipTypeSuggestion:
		return "SUGGESTION"
	}
	return strconv.FormatInt(int64(v), 10)
}

type ReportButtonType string

const (
	ReportButtonTypeNext   ReportButtonType = "next"
	ReportButtonTypeSubmit ReportButtonType = "submit"
	ReportButtonTypeDone   ReportButtonType = "done"
	ReportButtonTypeCancel ReportButtonType = "cancel"
)

type ReportElementType string

const (
	ReportElementTypeCheckbox                   ReportElementType = "checkbox"
	ReportElementTypeDropdown                   ReportElementType = "dropdown"
	ReportElementTypeFreeText                   ReportElementType = "free_text"
	ReportElementTypeRadio                      ReportElementType = "radio"
	ReportElementTypeSuccess                    ReportElementType = "success"
	ReportElementTypeFail                       ReportElementType = "fail"
	ReportElementTypeBreadcrumbs                ReportElementType = "breadcrumbs"
	ReportElementTypeTextLineResource           ReportElementType = "text_line_resource"
	ReportElementTypeText                       ReportElementType = "text"
	ReportElementTypeExternalLink               ReportElementType = "external_link"
	ReportElementTypeMoreYouCanDo               ReportElementType = "more_you_can_do"
	ReportElementTypeBlockUsers                 ReportElementType = "block_users"
	ReportElementTypeIgnoreUsers                ReportElementType = "ignore_users"
	ReportElementTypeMuteUsers                  ReportElementType = "mute_users"
	ReportElementTypeDeleteMessage              ReportElementType = "delete_message"
	ReportElementTypeLeaveGuild                 ReportElementType = "leave_guild"
	ReportElementTypeDeauthorizeApp             ReportElementType = "deauthorize_app"
	ReportElementTypeShareWithParents           ReportElementType = "share_with_parents"
	ReportElementTypeSettingsUpsells            ReportElementType = "settings_upsells"
	ReportElementTypeGuildPreview               ReportElementType = "guild_preview"
	ReportElementTypeGuildDiscoveryPreview      ReportElementType = "guild_discovery_preview"
	ReportElementTypeGuildDirectoryEntryPreview ReportElementType = "guild_directory_entry_preview"
	ReportElementTypeGuildScheduledEventPreview ReportElementType = "guild_scheduled_event_preview"
	ReportElementTypeMessagePreview             ReportElementType = "message_preview"
	ReportElementTypeChannelPreview             ReportElementType = "channel_preview"
	ReportElementTypeUserPreview                ReportElementType = "user_preview"
	ReportElementTypeAppPreview                 ReportElementType = "app_preview"
	ReportElementTypeWidgetPreview              ReportElementType = "widget_preview"
)

type ReportMenuType string

const (
	ReportMenuTypeGuild               ReportMenuType = "guild"
	ReportMenuTypeGuildDiscovery      ReportMenuType = "guild_discovery"
	ReportMenuTypeGuildDirectoryEntry ReportMenuType = "guild_directory_entry"
	ReportMenuTypeGuildScheduledEvent ReportMenuType = "guild_scheduled_event"
	ReportMenuTypeMessage             ReportMenuType = "message"
	ReportMenuTypeStageChannel        ReportMenuType = "stage_channel"
	ReportMenuTypeFirstDM             ReportMenuType = "first_dm"
	ReportMenuTypeUser                ReportMenuType = "user"
	ReportMenuTypeApplication         ReportMenuType = "application"
	ReportMenuTypeWidget              ReportMenuType = "widget"
)

type RequiredActionType string

const (
	RequiredActionTypeAgreements                              RequiredActionType = "AGREEMENTS"
	RequiredActionTypeRequireCaptcha                          RequiredActionType = "REQUIRE_CAPTCHA"
	RequiredActionTypeRequireVerifiedEmail                    RequiredActionType = "REQUIRE_VERIFIED_EMAIL"
	RequiredActionTypeRequireReverifiedEmail                  RequiredActionType = "REQUIRE_REVERIFIED_EMAIL"
	RequiredActionTypeRequireVerifiedPhone                    RequiredActionType = "REQUIRE_VERIFIED_PHONE"
	RequiredActionTypeRequireReverifiedPhone                  RequiredActionType = "REQUIRE_REVERIFIED_PHONE"
	RequiredActionTypeRequireVerifiedPhoneThenEmail           RequiredActionType = "REQUIRE_VERIFIED_PHONE_THEN_EMAIL"
	RequiredActionTypeRequireVerifiedEmailOrVerifiedPhone     RequiredActionType = "REQUIRE_VERIFIED_EMAIL_OR_VERIFIED_PHONE"
	RequiredActionTypeRequireReverifiedEmailOrVerifiedPhone   RequiredActionType = "REQUIRE_REVERIFIED_EMAIL_OR_VERIFIED_PHONE"
	RequiredActionTypeRequireVerifiedEmailOrReverifiedPhone   RequiredActionType = "REQUIRE_VERIFIED_EMAIL_OR_REVERIFIED_PHONE"
	RequiredActionTypeRequireReverifiedEmailOrReverifiedPhone RequiredActionType = "REQUIRE_REVERIFIED_EMAIL_OR_REVERIFIED_PHONE"
)

type ResolutionType string

const (
	ResolutionTypeFixed  ResolutionType = "fixed"
	ResolutionTypeSource ResolutionType = "source"
)

type ResponseType string

const (
	ResponseTypeCode  ResponseType = "code"
	ResponseTypeToken ResponseType = "token"
)

type RevenueSurface int64

const (
	RevenueSurfaceDiscovery RevenueSurface = 0
	RevenueSurfaceCheckout  RevenueSurface = 1
)

func (v RevenueSurface) String() string {
	switch v {
	case RevenueSurfaceDiscovery:
		return "DISCOVERY"
	case RevenueSurfaceCheckout:
		return "CHECKOUT"
	}
	return strconv.FormatInt(int64(v), 10)
}

type RoleConnectionOperatorType int64

const (
	RoleConnectionOperatorTypeIntegerLessThanOrEqual     RoleConnectionOperatorType = 1
	RoleConnectionOperatorTypeIntegerGreaterThanOrEqual  RoleConnectionOperatorType = 2
	RoleConnectionOperatorTypeIntegerEqual               RoleConnectionOperatorType = 3
	RoleConnectionOperatorTypeIntegerNotEqual            RoleConnectionOperatorType = 4
	RoleConnectionOperatorTypeDatetimeLessThanOrEqual    RoleConnectionOperatorType = 5
	RoleConnectionOperatorTypeDatetimeGreaterThanOrEqual RoleConnectionOperatorType = 6
	RoleConnectionOperatorTypeBooleanEqual               RoleConnectionOperatorType = 7
	RoleConnectionOperatorTypeBooleanNotEqual            RoleConnectionOperatorType = 8
)

func (v RoleConnectionOperatorType) String() string {
	switch v {
	case RoleConnectionOperatorTypeIntegerLessThanOrEqual:
		return "INTEGER_LESS_THAN_OR_EQUAL"
	case RoleConnectionOperatorTypeIntegerGreaterThanOrEqual:
		return "INTEGER_GREATER_THAN_OR_EQUAL"
	case RoleConnectionOperatorTypeIntegerEqual:
		return "INTEGER_EQUAL"
	case RoleConnectionOperatorTypeIntegerNotEqual:
		return "INTEGER_NOT_EQUAL"
	case RoleConnectionOperatorTypeDatetimeLessThanOrEqual:
		return "DATETIME_LESS_THAN_OR_EQUAL"
	case RoleConnectionOperatorTypeDatetimeGreaterThanOrEqual:
		return "DATETIME_GREATER_THAN_OR_EQUAL"
	case RoleConnectionOperatorTypeBooleanEqual:
		return "BOOLEAN_EQUAL"
	case RoleConnectionOperatorTypeBooleanNotEqual:
		return "BOOLEAN_NOT_EQUAL"
	}
	return strconv.FormatInt(int64(v), 10)
}

type RoleFlags uint64

const (
	RoleFlagsInPrompt RoleFlags = 1 << 0
)

func (v RoleFlags) String() string {
	switch v {
	case RoleFlagsInPrompt:
		return "IN_PROMPT"
	}
	return strconv.FormatUint(uint64(v), 10)
}

type RoleSubscriptionBenefitType int64

const (
	RoleSubscriptionBenefitTypeChannel    RoleSubscriptionBenefitType = 1
	RoleSubscriptionBenefitTypeIntangible RoleSubscriptionBenefitType = 2
)

func (v RoleSubscriptionBenefitType) String() string {
	switch v {
	case RoleSubscriptionBenefitTypeChannel:
		return "CHANNEL"
	case RoleSubscriptionBenefitTypeIntangible:
		return "INTANGIBLE"
	}
	return strconv.FormatInt(int64(v), 10)
}

type RuntimeEnvironment string

const (
	RuntimeEnvironmentNative   RuntimeEnvironment = "native"
	RuntimeEnvironmentFlatpak  RuntimeEnvironment = "flatpak"
	RuntimeEnvironmentSnap     RuntimeEnvironment = "snap"
	RuntimeEnvironmentAppimage RuntimeEnvironment = "appimage"
)

type SKUAccessType int64

const (
	SKUAccessTypeFull        SKUAccessType = 1
	SKUAccessTypeEarlyAccess SKUAccessType = 2
	SKUAccessTypeVipAccess   SKUAccessType = 3
)

func (v SKUAccessType) String() string {
	switch v {
	case SKUAccessTypeFull:
		return "FULL"
	case SKUAccessTypeEarlyAccess:
		return "EARLY_ACCESS"
	case SKUAccessTypeVipAccess:
		return "VIP_ACCESS"
	}
	return strconv.FormatInt(int64(v), 10)
}

type SKUFeature int64

const (
	SKUFeatureSinglePlayer       SKUFeature = 1
	SKUFeatureOnlineMultiplayer  SKUFeature = 2
	SKUFeatureLocalMultiplayer   SKUFeature = 3
	SKUFeaturePvp                SKUFeature = 4
	SKUFeatureLocalCoop          SKUFeature = 5
	SKUFeatureCrossPlatform      SKUFeature = 6
	SKUFeatureRichPresence       SKUFeature = 7
	SKUFeatureDiscordGameInvites SKUFeature = 8
	SKUFeatureSpectatorMode      SKUFeature = 9
	SKUFeatureControllerSupport  SKUFeature = 10
	SKUFeatureCloudSaves         SKUFeature = 11
	SKUFeatureOnlineCoop         SKUFeature = 12
	SKUFeatureSecureNetworking   SKUFeature = 13
)

func (v SKUFeature) String() string {
	switch v {
	case SKUFeatureSinglePlayer:
		return "SINGLE_PLAYER"
	case SKUFeatureOnlineMultiplayer:
		return "ONLINE_MULTIPLAYER"
	case SKUFeatureLocalMultiplayer:
		return "LOCAL_MULTIPLAYER"
	case SKUFeaturePvp:
		return "PVP"
	case SKUFeatureLocalCoop:
		return "LOCAL_COOP"
	case SKUFeatureCrossPlatform:
		return "CROSS_PLATFORM"
	case SKUFeatureRichPresence:
		return "RICH_PRESENCE"
	case SKUFeatureDiscordGameInvites:
		return "DISCORD_GAME_INVITES"
	case SKUFeatureSpectatorMode:
		return "SPECTATOR_MODE"
	case SKUFeatureControllerSupport:
		return "CONTROLLER_SUPPORT"
	case SKUFeatureCloudSaves:
		return "CLOUD_SAVES"
	case SKUFeatureOnlineCoop:
		return "ONLINE_COOP"
	case SKUFeatureSecureNetworking:
		return "SECURE_NETWORKING"
	}
	return strconv.FormatInt(int64(v), 10)
}

type SKUFlags uint64

const (
	SKUFlagsPremiumPurchase                 SKUFlags = 1 << 0
	SKUFlagsHasFreePremiumContent           SKUFlags = 1 << 1
	SKUFlagsAvailable                       SKUFlags = 1 << 2
	SKUFlagsPremiumAndDistribution          SKUFlags = 1 << 3
	SKUFlagsSticker                         SKUFlags = 1 << 4
	SKUFlagsGuildRole                       SKUFlags = 1 << 5
	SKUFlagsAvailableForSubscriptionGifting SKUFlags = 1 << 6
	SKUFlagsApplicationGuildSubscription    SKUFlags = 1 << 7
	SKUFlagsApplicationUserSubscription     SKUFlags = 1 << 8
	SKUFlagsCreatorMonetization             SKUFlags = 1 << 9
	SKUFlagsGuildProduct                    SKUFlags = 1 << 10
	SKUFlagsAvailableForApplicationGifting  SKUFlags = 1 << 11
)

func (v SKUFlags) String() string {
	switch v {
	case SKUFlagsPremiumPurchase:
		return "PREMIUM_PURCHASE"
	case SKUFlagsHasFreePremiumContent:
		return "HAS_FREE_PREMIUM_CONTENT"
	case SKUFlagsAvailable:
		return "AVAILABLE"
	case SKUFlagsPremiumAndDistribution:
		return "PREMIUM_AND_DISTRIBUTION"
	case SKUFlagsSticker:
		return "STICKER"
	case SKUFlagsGuildRole:
		return "GUILD_ROLE"
	case SKUFlagsAvailableForSubscriptionGifting:
		return "AVAILABLE_FOR_SUBSCRIPTION_GIFTING"
	case SKUFlagsApplicationGuildSubscription:
		return "APPLICATION_GUILD_SUBSCRIPTION"
	case SKUFlagsApplicationUserSubscription:
		return "APPLICATION_USER_SUBSCRIPTION"
	case SKUFlagsCreatorMonetization:
		return "CREATOR_MONETIZATION"
	case SKUFlagsGuildProduct:
		return "GUILD_PRODUCT"
	case SKUFlagsAvailableForApplicationGifting:
		return "AVAILABLE_FOR_APPLICATION_GIFTING"
	}
	return strconv.FormatUint(uint64(v), 10)
}

type SKUGenre int64

const (
	SKUGenreAction               SKUGenre = 1
	SKUGenreActionRpg            SKUGenre = 2
	SKUGenreBrawler              SKUGenre = 3
	SKUGenreHackAndSlash         SKUGenre = 4
	SKUGenrePlatformer           SKUGenre = 5
	SKUGenreStealth              SKUGenre = 6
	SKUGenreSurvival             SKUGenre = 7
	SKUGenreAdventure            SKUGenre = 8
	SKUGenreActionAdventure      SKUGenre = 9
	SKUGenreMetroidvania         SKUGenre = 10
	SKUGenreOpenWorld            SKUGenre = 11
	SKUGenrePsychologicalHorror  SKUGenre = 12
	SKUGenreSandbox              SKUGenre = 13
	SKUGenreSurvivalHorror       SKUGenre = 14
	SKUGenreVisualNovel          SKUGenre = 15
	SKUGenreDrivingRacing        SKUGenre = 16
	SKUGenreVehicularCombat      SKUGenre = 17
	SKUGenreMassivelyMultiplayer SKUGenre = 18
	SKUGenreMmorpg               SKUGenre = 19
	SKUGenreRolePlaying          SKUGenre = 20
	SKUGenreDungeonCrawler       SKUGenre = 21
	SKUGenreRoguelike            SKUGenre = 22
	SKUGenreShooter              SKUGenre = 23
	SKUGenreLightGun             SKUGenre = 24
	SKUGenreShootEmUp            SKUGenre = 25
	SKUGenreFps                  SKUGenre = 26
	SKUGenreDualJoystickShooter  SKUGenre = 27
	SKUGenreSimulation           SKUGenre = 28
	SKUGenreFlightSimulator      SKUGenre = 29
	SKUGenreTrainSimulator       SKUGenre = 30
	SKUGenreLifeSimulator        SKUGenre = 31
	SKUGenreFishing              SKUGenre = 32
	SKUGenreSports               SKUGenre = 33
	SKUGenreBaseball             SKUGenre = 34
	SKUGenreBasketball           SKUGenre = 35
	SKUGenreBilliards            SKUGenre = 36
	SKUGenreBowling              SKUGenre = 37
	SKUGenreBoxing               SKUGenre = 38
	SKUGenreFootball             SKUGenre = 39
	SKUGenreGolf                 SKUGenre = 40
	SKUGenreHockey               SKUGenre = 41
	SKUGenreSkateboardingSkating SKUGenre = 42
	SKUGenreSnowboardingSkiing   SKUGenre = 43
	SKUGenreSoccer               SKUGenre = 44
	SKUGenreTrackField           SKUGenre = 45
	SKUGenreSurfingWakeboarding  SKUGenre = 46
	SKUGenreWrestling            SKUGenre = 47
	SKUGenreStrategy             SKUGenre = 48
	SKUGenreFourX                SKUGenre = 49
	SKUGenreArtillery            SKUGenre = 50
	SKUGenreRts                  SKUGenre = 51
	SKUGenreTowerDefense         SKUGenre = 52
	SKUGenreTurnBasedStrategy    SKUGenre = 53
	SKUGenreWargame              SKUGenre = 54
	SKUGenreMoba                 SKUGenre = 55
	SKUGenreFighting             SKUGenre = 56
	SKUGenrePuzzle               SKUGenre = 57
	SKUGenreCardGame             SKUGenre = 58
	SKUGenreEducation            SKUGenre = 59
	SKUGenreFitness              SKUGenre = 60
	SKUGenreGambling             SKUGenre = 61
	SKUGenreMusicRhythm          SKUGenre = 62
	SKUGenrePartyMiniGame        SKUGenre = 63
	SKUGenrePinball              SKUGenre = 64
	SKUGenreTriviaBoardGame      SKUGenre = 65
	SKUGenreTactical             SKUGenre = 66
	SKUGenreIndie                SKUGenre = 67
	SKUGenreArcade               SKUGenre = 68
	SKUGenrePointAndClick        SKUGenre = 69
)

func (v SKUGenre) String() string {
	switch v {
	case SKUGenreAction:
		return "ACTION"
	case SKUGenreActionRpg:
		return "ACTION_RPG"
	case SKUGenreBrawler:
		return "BRAWLER"
	case SKUGenreHackAndSlash:
		return "HACK_AND_SLASH"
	case SKUGenrePlatformer:
		return "PLATFORMER"
	case SKUGenreStealth:
		return "STEALTH"
	case SKUGenreSurvival:
		return "SURVIVAL"
	case SKUGenreAdventure:
		return "ADVENTURE"
	case SKUGenreActionAdventure:
		return "ACTION_ADVENTURE"
	case SKUGenreMetroidvania:
		return "METROIDVANIA"
	case SKUGenreOpenWorld:
		return "OPEN_WORLD"
	case SKUGenrePsychologicalHorror:
		return "PSYCHOLOGICAL_HORROR"
	case SKUGenreSandbox:
		return "SANDBOX"
	case SKUGenreSurvivalHorror:
		return "SURVIVAL_HORROR"
	case SKUGenreVisualNovel:
		return "VISUAL_NOVEL"
	case SKUGenreDrivingRacing:
		return "DRIVING_RACING"
	case SKUGenreVehicularCombat:
		return "VEHICULAR_COMBAT"
	case SKUGenreMassivelyMultiplayer:
		return "MASSIVELY_MULTIPLAYER"
	case SKUGenreMmorpg:
		return "MMORPG"
	case SKUGenreRolePlaying:
		return "ROLE_PLAYING"
	case SKUGenreDungeonCrawler:
		return "DUNGEON_CRAWLER"
	case SKUGenreRoguelike:
		return "ROGUELIKE"
	case SKUGenreShooter:
		return "SHOOTER"
	case SKUGenreLightGun:
		return "LIGHT_GUN"
	case SKUGenreShootEmUp:
		return "SHOOT_EM_UP"
	case SKUGenreFps:
		return "FPS"
	case SKUGenreDualJoystickShooter:
		return "DUAL_JOYSTICK_SHOOTER"
	case SKUGenreSimulation:
		return "SIMULATION"
	case SKUGenreFlightSimulator:
		return "FLIGHT_SIMULATOR"
	case SKUGenreTrainSimulator:
		return "TRAIN_SIMULATOR"
	case SKUGenreLifeSimulator:
		return "LIFE_SIMULATOR"
	case SKUGenreFishing:
		return "FISHING"
	case SKUGenreSports:
		return "SPORTS"
	case SKUGenreBaseball:
		return "BASEBALL"
	case SKUGenreBasketball:
		return "BASKETBALL"
	case SKUGenreBilliards:
		return "BILLIARDS"
	case SKUGenreBowling:
		return "BOWLING"
	case SKUGenreBoxing:
		return "BOXING"
	case SKUGenreFootball:
		return "FOOTBALL"
	case SKUGenreGolf:
		return "GOLF"
	case SKUGenreHockey:
		return "HOCKEY"
	case SKUGenreSkateboardingSkating:
		return "SKATEBOARDING_SKATING"
	case SKUGenreSnowboardingSkiing:
		return "SNOWBOARDING_SKIING"
	case SKUGenreSoccer:
		return "SOCCER"
	case SKUGenreTrackField:
		return "TRACK_FIELD"
	case SKUGenreSurfingWakeboarding:
		return "SURFING_WAKEBOARDING"
	case SKUGenreWrestling:
		return "WRESTLING"
	case SKUGenreStrategy:
		return "STRATEGY"
	case SKUGenreFourX:
		return "FOUR_X"
	case SKUGenreArtillery:
		return "ARTILLERY"
	case SKUGenreRts:
		return "RTS"
	case SKUGenreTowerDefense:
		return "TOWER_DEFENSE"
	case SKUGenreTurnBasedStrategy:
		return "TURN_BASED_STRATEGY"
	case SKUGenreWargame:
		return "WARGAME"
	case SKUGenreMoba:
		return "MOBA"
	case SKUGenreFighting:
		return "FIGHTING"
	case SKUGenrePuzzle:
		return "PUZZLE"
	case SKUGenreCardGame:
		return "CARD_GAME"
	case SKUGenreEducation:
		return "EDUCATION"
	case SKUGenreFitness:
		return "FITNESS"
	case SKUGenreGambling:
		return "GAMBLING"
	case SKUGenreMusicRhythm:
		return "MUSIC_RHYTHM"
	case SKUGenrePartyMiniGame:
		return "PARTY_MINI_GAME"
	case SKUGenrePinball:
		return "PINBALL"
	case SKUGenreTriviaBoardGame:
		return "TRIVIA_BOARD_GAME"
	case SKUGenreTactical:
		return "TACTICAL"
	case SKUGenreIndie:
		return "INDIE"
	case SKUGenreArcade:
		return "ARCADE"
	case SKUGenrePointAndClick:
		return "POINT_AND_CLICK"
	}
	return strconv.FormatInt(int64(v), 10)
}

type SKUProductLine int64

const (
	SKUProductLinePremium             SKUProductLine = 1
	SKUProductLinePremiumGuild        SKUProductLine = 2
	SKUProductLineActivityIap         SKUProductLine = 3
	SKUProductLineGuildRole           SKUProductLine = 4
	SKUProductLineGuildProduct        SKUProductLine = 5
	SKUProductLineApplication         SKUProductLine = 6
	SKUProductLineCollectibles        SKUProductLine = 7
	SKUProductLineTenureReward        SKUProductLine = 8
	SKUProductLineQuestInGameReward   SKUProductLine = 9
	SKUProductLineQuestRewardCode     SKUProductLine = 10
	SKUProductLineFractionalPremium   SKUProductLine = 11
	SKUProductLineVirtualCurrency     SKUProductLine = 12
	SKUProductLineGuildPowerup        SKUProductLine = 13
	SKUProductLineSocialLayerGameItem SKUProductLine = 14
)

func (v SKUProductLine) String() string {
	switch v {
	case SKUProductLinePremium:
		return "PREMIUM"
	case SKUProductLinePremiumGuild:
		return "PREMIUM_GUILD"
	case SKUProductLineActivityIap:
		return "ACTIVITY_IAP"
	case SKUProductLineGuildRole:
		return "GUILD_ROLE"
	case SKUProductLineGuildProduct:
		return "GUILD_PRODUCT"
	case SKUProductLineApplication:
		return "APPLICATION"
	case SKUProductLineCollectibles:
		return "COLLECTIBLES"
	case SKUProductLineTenureReward:
		return "TENURE_REWARD"
	case SKUProductLineQuestInGameReward:
		return "QUEST_IN_GAME_REWARD"
	case SKUProductLineQuestRewardCode:
		return "QUEST_REWARD_CODE"
	case SKUProductLineFractionalPremium:
		return "FRACTIONAL_PREMIUM"
	case SKUProductLineVirtualCurrency:
		return "VIRTUAL_CURRENCY"
	case SKUProductLineGuildPowerup:
		return "GUILD_POWERUP"
	case SKUProductLineSocialLayerGameItem:
		return "SOCIAL_LAYER_GAME_ITEM"
	}
	return strconv.FormatInt(int64(v), 10)
}

type SKUType int64

const (
	SKUTypeDurablePrimary    SKUType = 1
	SKUTypeDurable           SKUType = 2
	SKUTypeConsumable        SKUType = 3
	SKUTypeBundle            SKUType = 4
	SKUTypeSubscription      SKUType = 5
	SKUTypeSubscriptionGroup SKUType = 6
)

func (v SKUType) String() string {
	switch v {
	case SKUTypeDurablePrimary:
		return "DURABLE_PRIMARY"
	case SKUTypeDurable:
		return "DURABLE"
	case SKUTypeConsumable:
		return "CONSUMABLE"
	case SKUTypeBundle:
		return "BUNDLE"
	case SKUTypeSubscription:
		return "SUBSCRIPTION"
	case SKUTypeSubscriptionGroup:
		return "SUBSCRIPTION_GROUP"
	}
	return strconv.FormatInt(int64(v), 10)
}

type SafetyHubAppealStatus int64

const (
	SafetyHubAppealStatusReviewPending             SafetyHubAppealStatus = 1
	SafetyHubAppealStatusClassificationUpheld      SafetyHubAppealStatus = 2
	SafetyHubAppealStatusClassificationInvalidated SafetyHubAppealStatus = 3
)

func (v SafetyHubAppealStatus) String() string {
	switch v {
	case SafetyHubAppealStatusReviewPending:
		return "REVIEW_PENDING"
	case SafetyHubAppealStatusClassificationUpheld:
		return "CLASSIFICATION_UPHELD"
	case SafetyHubAppealStatusClassificationInvalidated:
		return "CLASSIFICATION_INVALIDATED"
	}
	return strconv.FormatInt(int64(v), 10)
}

type SafetySettingsPresetType int64

const (
	SafetySettingsPresetTypeUnset    SafetySettingsPresetType = 0
	SafetySettingsPresetTypeBalanced SafetySettingsPresetType = 1
	SafetySettingsPresetTypeStrict   SafetySettingsPresetType = 2
	SafetySettingsPresetTypeRelaxed  SafetySettingsPresetType = 3
	SafetySettingsPresetTypeCustom   SafetySettingsPresetType = 4
)

func (v SafetySettingsPresetType) String() string {
	switch v {
	case SafetySettingsPresetTypeUnset:
		return "UNSET"
	case SafetySettingsPresetTypeBalanced:
		return "BALANCED"
	case SafetySettingsPresetTypeStrict:
		return "STRICT"
	case SafetySettingsPresetTypeRelaxed:
		return "RELAXED"
	case SafetySettingsPresetTypeCustom:
		return "CUSTOM"
	}
	return strconv.FormatInt(int64(v), 10)
}

type SafetySystemNotificationCtaType string

const (
	SafetySystemNotificationCtaTypeLearnMoreLink         SafetySystemNotificationCtaType = "learn_more_link"
	SafetySystemNotificationCtaTypePolicyViolationDetail SafetySystemNotificationCtaType = "policy_violation_detail"
)

type SafetyWarningType int64

const (
	SafetyWarningTypeStrangerDanger                 SafetyWarningType = 1
	SafetyWarningTypeInappropriateConversationTier1 SafetyWarningType = 2
	SafetyWarningTypeInappropriateConversationTier2 SafetyWarningType = 3
	SafetyWarningTypeLikelyAto                      SafetyWarningType = 4
)

func (v SafetyWarningType) String() string {
	switch v {
	case SafetyWarningTypeStrangerDanger:
		return "STRANGER_DANGER"
	case SafetyWarningTypeInappropriateConversationTier1:
		return "INAPPROPRIATE_CONVERSATION_TIER_1"
	case SafetyWarningTypeInappropriateConversationTier2:
		return "INAPPROPRIATE_CONVERSATION_TIER_2"
	case SafetyWarningTypeLikelyAto:
		return "LIKELY_ATO"
	}
	return strconv.FormatInt(int64(v), 10)
}

type SearchCursorType string

const (
	SearchCursorTypeTimestamp SearchCursorType = "timestamp"
	SearchCursorTypeScore     SearchCursorType = "score"
)

type SearchEmbedType string

const (
	SearchEmbedTypeImage   SearchEmbedType = "image"
	SearchEmbedTypeVideo   SearchEmbedType = "video"
	SearchEmbedTypeGIF     SearchEmbedType = "gif"
	SearchEmbedTypeSound   SearchEmbedType = "sound"
	SearchEmbedTypeArticle SearchEmbedType = "article"
)

type SearchEngine string

const (
	SearchEngineGoogle     SearchEngine = "google"
	SearchEngineBing       SearchEngine = "bing"
	SearchEngineYahoo      SearchEngine = "yahoo"
	SearchEngineDuckduckgo SearchEngine = "duckduckgo"
)

type SearchHasType string

const (
	SearchHasTypeImage    SearchHasType = "image"
	SearchHasTypeSound    SearchHasType = "sound"
	SearchHasTypeVideo    SearchHasType = "video"
	SearchHasTypeFile     SearchHasType = "file"
	SearchHasTypeSticker  SearchHasType = "sticker"
	SearchHasTypeEmbed    SearchHasType = "embed"
	SearchHasTypeLink     SearchHasType = "link"
	SearchHasTypePoll     SearchHasType = "poll"
	SearchHasTypeSnapshot SearchHasType = "snapshot"
)

type SearchProvider int64

const (
	SearchProviderUnset      SearchProvider = 0
	SearchProviderGoogle     SearchProvider = 1
	SearchProviderBing       SearchProvider = 2
	SearchProviderDuckduckgo SearchProvider = 3
	SearchProviderCustom     SearchProvider = 4
)

func (v SearchProvider) String() string {
	switch v {
	case SearchProviderUnset:
		return "UNSET"
	case SearchProviderGoogle:
		return "GOOGLE"
	case SearchProviderBing:
		return "BING"
	case SearchProviderDuckduckgo:
		return "DUCKDUCKGO"
	case SearchProviderCustom:
		return "CUSTOM"
	}
	return strconv.FormatInt(int64(v), 10)
}

type SearchSortMode string

const (
	SearchSortModeTimestamp SearchSortMode = "timestamp"
	SearchSortModeRelevance SearchSortMode = "relevance"
)

type SearchTabType string

const (
	SearchTabTypeMessages SearchTabType = "messages"
	SearchTabTypeLinks    SearchTabType = "links"
	SearchTabTypeMedia    SearchTabType = "media"
	SearchTabTypeFiles    SearchTabType = "files"
	SearchTabTypePins     SearchTabType = "pins"
)

type SearchTagSetting string

const (
	SearchTagSettingMatchSome SearchTagSetting = "match_some"
	SearchTagSettingMatchAll  SearchTagSetting = "match_all"
)

type SeparatorSpacingType int64

const (
	SeparatorSpacingTypeSmall SeparatorSpacingType = 1
	SeparatorSpacingTypeLarge SeparatorSpacingType = 2
)

func (v SeparatorSpacingType) String() string {
	switch v {
	case SeparatorSpacingTypeSmall:
		return "SMALL"
	case SeparatorSpacingTypeLarge:
		return "LARGE"
	}
	return strconv.FormatInt(int64(v), 10)
}

type SessionType string

const (
	SessionTypeNormal SessionType = "normal"
	SessionTypeOauth  SessionType = "oauth"
)

type ShopBlockType int64

const (
	ShopBlockTypeHero                                   ShopBlockType = 0
	ShopBlockTypeFeatured                               ShopBlockType = 1
	ShopBlockTypeFeed                                   ShopBlockType = 2
	ShopBlockTypeWideBanner                             ShopBlockType = 3
	ShopBlockTypeShelf                                  ShopBlockType = 4
	ShopBlockTypeCountdownTimer                         ShopBlockType = 5
	ShopBlockTypeImmersiveBanner                        ShopBlockType = 6
	ShopBlockTypeRewardHero                             ShopBlockType = 7
	ShopBlockTypeSocialLayerStorefrontPromotionalBanner ShopBlockType = 9
	ShopBlockTypeFramesBanner                           ShopBlockType = 10
	ShopBlockTypeFramesProductShelf                     ShopBlockType = 11
)

func (v ShopBlockType) String() string {
	switch v {
	case ShopBlockTypeHero:
		return "HERO"
	case ShopBlockTypeFeatured:
		return "FEATURED"
	case ShopBlockTypeFeed:
		return "FEED"
	case ShopBlockTypeWideBanner:
		return "WIDE_BANNER"
	case ShopBlockTypeShelf:
		return "SHELF"
	case ShopBlockTypeCountdownTimer:
		return "COUNTDOWN_TIMER"
	case ShopBlockTypeImmersiveBanner:
		return "IMMERSIVE_BANNER"
	case ShopBlockTypeRewardHero:
		return "REWARD_HERO"
	case ShopBlockTypeSocialLayerStorefrontPromotionalBanner:
		return "SOCIAL_LAYER_STOREFRONT_PROMOTIONAL_BANNER"
	case ShopBlockTypeFramesBanner:
		return "FRAMES_BANNER"
	case ShopBlockTypeFramesProductShelf:
		return "FRAMES_PRODUCT_SHELF"
	}
	return strconv.FormatInt(int64(v), 10)
}

type ShopItemType string

const (
	ShopItemTypeAll              ShopItemType = "ALL"
	ShopItemTypeAvatarDecoration ShopItemType = "AVATAR_DECORATION"
	ShopItemTypeProfileEffect    ShopItemType = "PROFILE_EFFECT"
	ShopItemTypeNameplate        ShopItemType = "NAMEPLATE"
	ShopItemTypeProfileFrame     ShopItemType = "PROFILE_FRAME"
	ShopItemTypeBundle           ShopItemType = "BUNDLE"
)

type ShopProductColor string

const (
	ShopProductColorCollectiblesColorBlue   ShopProductColor = "COLLECTIBLES_COLOR_BLUE"
	ShopProductColorCollectiblesColorGreen  ShopProductColor = "COLLECTIBLES_COLOR_GREEN"
	ShopProductColorCollectiblesColorPink   ShopProductColor = "COLLECTIBLES_COLOR_PINK"
	ShopProductColorCollectiblesColorRed    ShopProductColor = "COLLECTIBLES_COLOR_RED"
	ShopProductColorCollectiblesColorYellow ShopProductColor = "COLLECTIBLES_COLOR_YELLOW"
	ShopProductColorCollectiblesColorOrange ShopProductColor = "COLLECTIBLES_COLOR_ORANGE"
	ShopProductColorCollectiblesColorPurple ShopProductColor = "COLLECTIBLES_COLOR_PURPLE"
	ShopProductColorCollectiblesColorBrown  ShopProductColor = "COLLECTIBLES_COLOR_BROWN"
	ShopProductColorCollectiblesColorBlack  ShopProductColor = "COLLECTIBLES_COLOR_BLACK"
	ShopProductColorCollectiblesColorWhite  ShopProductColor = "COLLECTIBLES_COLOR_WHITE"
)

type ShopProductTheme string

const (
	ShopProductThemeCollectiblesThemeAnime         ShopProductTheme = "COLLECTIBLES_THEME_ANIME"
	ShopProductThemeCollectiblesThemeGaming        ShopProductTheme = "COLLECTIBLES_THEME_GAMING"
	ShopProductThemeCollectiblesThemeCuteCozy      ShopProductTheme = "COLLECTIBLES_THEME_CUTE_COZY"
	ShopProductThemeCollectiblesThemeFoodDrinks    ShopProductTheme = "COLLECTIBLES_THEME_FOOD_DRINKS"
	ShopProductThemeCollectiblesThemeAnimalsPets   ShopProductTheme = "COLLECTIBLES_THEME_ANIMALS_PETS"
	ShopProductThemeCollectiblesThemeMoviesTvShows ShopProductTheme = "COLLECTIBLES_THEME_MOVIES_TV_SHOWS"
	ShopProductThemeCollectiblesThemeFantasy       ShopProductTheme = "COLLECTIBLES_THEME_FANTASY"
	ShopProductThemeCollectiblesThemeDarkMoody     ShopProductTheme = "COLLECTIBLES_THEME_DARK_MOODY"
	ShopProductThemeCollectiblesThemeNature        ShopProductTheme = "COLLECTIBLES_THEME_NATURE"
	ShopProductThemeCollectiblesThemeSciFi         ShopProductTheme = "COLLECTIBLES_THEME_SCI_FI"
)

type ShopSortType string

const (
	ShopSortTypeRelevance    ShopSortType = "relevance"
	ShopSortTypePrice        ShopSortType = "price"
	ShopSortTypeAlphabetical ShopSortType = "alphabetical"
	ShopSortTypeRecency      ShopSortType = "recency"
	ShopSortTypePopularity   ShopSortType = "popularity"
)

type ShortcutKeyComboType int64

const (
	ShortcutKeyComboTypeKeyboardKey         ShortcutKeyComboType = 0
	ShortcutKeyComboTypeMouseButton         ShortcutKeyComboType = 1
	ShortcutKeyComboTypeKeyboardModifierKey ShortcutKeyComboType = 2
	ShortcutKeyComboTypeGamepadButton       ShortcutKeyComboType = 3
)

func (v ShortcutKeyComboType) String() string {
	switch v {
	case ShortcutKeyComboTypeKeyboardKey:
		return "KEYBOARD_KEY"
	case ShortcutKeyComboTypeMouseButton:
		return "MOUSE_BUTTON"
	case ShortcutKeyComboTypeKeyboardModifierKey:
		return "KEYBOARD_MODIFIER_KEY"
	case ShortcutKeyComboTypeGamepadButton:
		return "GAMEPAD_BUTTON"
	}
	return strconv.FormatInt(int64(v), 10)
}

type SlayerSdkReceiveInGameDMs int64

const (
	SlayerSdkReceiveInGameDMsUnset         SlayerSdkReceiveInGameDMs = 0
	SlayerSdkReceiveInGameDMsAll           SlayerSdkReceiveInGameDMs = 1
	SlayerSdkReceiveInGameDMsUsersWithGame SlayerSdkReceiveInGameDMs = 2
	SlayerSdkReceiveInGameDMsNone          SlayerSdkReceiveInGameDMs = 3
)

func (v SlayerSdkReceiveInGameDMs) String() string {
	switch v {
	case SlayerSdkReceiveInGameDMsUnset:
		return "UNSET"
	case SlayerSdkReceiveInGameDMsAll:
		return "ALL"
	case SlayerSdkReceiveInGameDMsUsersWithGame:
		return "USERS_WITH_GAME"
	case SlayerSdkReceiveInGameDMsNone:
		return "NONE"
	}
	return strconv.FormatInt(int64(v), 10)
}

type SortOrderType int64

const (
	SortOrderTypeLatestActivity SortOrderType = 0
	SortOrderTypeCreationTime   SortOrderType = 1
)

func (v SortOrderType) String() string {
	switch v {
	case SortOrderTypeLatestActivity:
		return "LATEST_ACTIVITY"
	case SortOrderTypeCreationTime:
		return "CREATION_TIME"
	}
	return strconv.FormatInt(int64(v), 10)
}

type SoundboardPickerSection int64

const (
	SoundboardPickerSectionN0 SoundboardPickerSection = 0
	SoundboardPickerSectionN4 SoundboardPickerSection = 4
)

func (v SoundboardPickerSection) String() string {
	switch v {
	case SoundboardPickerSectionN0:
		return "0"
	case SoundboardPickerSectionN4:
		return "4"
	}
	return strconv.FormatInt(int64(v), 10)
}

type SpeakingFlags uint64

const (
	SpeakingFlagsVoice      SpeakingFlags = 1 << 0
	SpeakingFlagsSoundshare SpeakingFlags = 1 << 1
	SpeakingFlagsPriority   SpeakingFlags = 1 << 2
)

func (v SpeakingFlags) String() string {
	switch v {
	case SpeakingFlagsVoice:
		return "VOICE"
	case SpeakingFlagsSoundshare:
		return "SOUNDSHARE"
	case SpeakingFlagsPriority:
		return "PRIORITY"
	}
	return strconv.FormatUint(uint64(v), 10)
}

type SpoilerRenderOptions string

const (
	SpoilerRenderOptionsAlways      SpoilerRenderOptions = "ALWAYS"
	SpoilerRenderOptionsOnClick     SpoilerRenderOptions = "ON_CLICK"
	SpoilerRenderOptionsIfModerator SpoilerRenderOptions = "IF_MODERATOR"
)

type StatusDisplayType int64

const (
	StatusDisplayTypeName    StatusDisplayType = 0
	StatusDisplayTypeState   StatusDisplayType = 1
	StatusDisplayTypeDetails StatusDisplayType = 2
)

func (v StatusDisplayType) String() string {
	switch v {
	case StatusDisplayTypeName:
		return "NAME"
	case StatusDisplayTypeState:
		return "STATE"
	case StatusDisplayTypeDetails:
		return "DETAILS"
	}
	return strconv.FormatInt(int64(v), 10)
}

type StatusType string

const (
	StatusTypeOnline    StatusType = "online"
	StatusTypeIdle      StatusType = "idle"
	StatusTypeDnd       StatusType = "dnd"
	StatusTypeInvisible StatusType = "invisible"
	StatusTypeOffline   StatusType = "offline"
	StatusTypeUnknown   StatusType = "unknown"
)

type SteamReleaseStatus int64

const (
	SteamReleaseStatusPreRelease       SteamReleaseStatus = 1
	SteamReleaseStatusDayOfRelease     SteamReleaseStatus = 2
	SteamReleaseStatusPostRelease      SteamReleaseStatus = 3
	SteamReleaseStatusRetiredAbandoned SteamReleaseStatus = 4
	SteamReleaseStatusChildApp         SteamReleaseStatus = 6
)

func (v SteamReleaseStatus) String() string {
	switch v {
	case SteamReleaseStatusPreRelease:
		return "PRE_RELEASE"
	case SteamReleaseStatusDayOfRelease:
		return "DAY_OF_RELEASE"
	case SteamReleaseStatusPostRelease:
		return "POST_RELEASE"
	case SteamReleaseStatusRetiredAbandoned:
		return "RETIRED_ABANDONED"
	case SteamReleaseStatusChildApp:
		return "CHILD_APP"
	}
	return strconv.FormatInt(int64(v), 10)
}

type StickerAnimationOption int64

const (
	StickerAnimationOptionAlwaysAnimate        StickerAnimationOption = 0
	StickerAnimationOptionAnimateOnInteraction StickerAnimationOption = 1
	StickerAnimationOptionNeverAnimate         StickerAnimationOption = 2
)

func (v StickerAnimationOption) String() string {
	switch v {
	case StickerAnimationOptionAlwaysAnimate:
		return "ALWAYS_ANIMATE"
	case StickerAnimationOptionAnimateOnInteraction:
		return "ANIMATE_ON_INTERACTION"
	case StickerAnimationOptionNeverAnimate:
		return "NEVER_ANIMATE"
	}
	return strconv.FormatInt(int64(v), 10)
}

type StickerFormatTypes int64

const (
	StickerFormatTypesPNG    StickerFormatTypes = 1
	StickerFormatTypesApng   StickerFormatTypes = 2
	StickerFormatTypesLottie StickerFormatTypes = 3
	StickerFormatTypesGIF    StickerFormatTypes = 4
)

func (v StickerFormatTypes) String() string {
	switch v {
	case StickerFormatTypesPNG:
		return "PNG"
	case StickerFormatTypesApng:
		return "APNG"
	case StickerFormatTypesLottie:
		return "LOTTIE"
	case StickerFormatTypesGIF:
		return "GIF"
	}
	return strconv.FormatInt(int64(v), 10)
}

type StickerPickerSection string

const (
	StickerPickerSectionFavorite StickerPickerSection = "FAVORITE"
	StickerPickerSectionRecent   StickerPickerSection = "RECENT"
)

type StickerTypes int64

const (
	StickerTypesStandard StickerTypes = 1
	StickerTypesGuild    StickerTypes = 2
)

func (v StickerTypes) String() string {
	switch v {
	case StickerTypesStandard:
		return "STANDARD"
	case StickerTypesGuild:
		return "GUILD"
	}
	return strconv.FormatInt(int64(v), 10)
}

type StoreApplicationState int64

const (
	StoreApplicationStateNone      StoreApplicationState = 1
	StoreApplicationStatePaid      StoreApplicationState = 2
	StoreApplicationStateSubmitted StoreApplicationState = 3
	StoreApplicationStateApproved  StoreApplicationState = 4
	StoreApplicationStateRejected  StoreApplicationState = 5
)

func (v StoreApplicationState) String() string {
	switch v {
	case StoreApplicationStateNone:
		return "NONE"
	case StoreApplicationStatePaid:
		return "PAID"
	case StoreApplicationStateSubmitted:
		return "SUBMITTED"
	case StoreApplicationStateApproved:
		return "APPROVED"
	case StoreApplicationStateRejected:
		return "REJECTED"
	}
	return strconv.FormatInt(int64(v), 10)
}

type StoreListingIconType int64

const (
	StoreListingIconTypeStoreAsset StoreListingIconType = 1
	StoreListingIconTypeEmoji      StoreListingIconType = 2
)

func (v StoreListingIconType) String() string {
	switch v {
	case StoreListingIconTypeStoreAsset:
		return "STORE_ASSET"
	case StoreListingIconTypeEmoji:
		return "EMOJI"
	}
	return strconv.FormatInt(int64(v), 10)
}

type StoreOperatingSystem int64

const (
	StoreOperatingSystemWindows StoreOperatingSystem = 1
	StoreOperatingSystemMacos   StoreOperatingSystem = 2
	StoreOperatingSystemLinux   StoreOperatingSystem = 3
)

func (v StoreOperatingSystem) String() string {
	switch v {
	case StoreOperatingSystemWindows:
		return "WINDOWS"
	case StoreOperatingSystemMacos:
		return "MACOS"
	case StoreOperatingSystemLinux:
		return "LINUX"
	}
	return strconv.FormatInt(int64(v), 10)
}

type StorefrontPlatformType string

const (
	StorefrontPlatformTypeDesktop     StorefrontPlatformType = "desktop"
	StorefrontPlatformTypeXbox        StorefrontPlatformType = "xbox"
	StorefrontPlatformTypePlaystation StorefrontPlatformType = "playstation"
)

type StorefrontPriceVariantType int64

const (
	StorefrontPriceVariantTypeNormal     StorefrontPriceVariantType = 0
	StorefrontPriceVariantTypeDiscounted StorefrontPriceVariantType = 1
)

func (v StorefrontPriceVariantType) String() string {
	switch v {
	case StorefrontPriceVariantTypeNormal:
		return "NORMAL"
	case StorefrontPriceVariantTypeDiscounted:
		return "DISCOUNTED"
	}
	return strconv.FormatInt(int64(v), 10)
}

type StorefrontPromotionRewardType int64

const (
	StorefrontPromotionRewardTypeDiscount StorefrontPromotionRewardType = 1
)

func (v StorefrontPromotionRewardType) String() string {
	switch v {
	case StorefrontPromotionRewardTypeDiscount:
		return "DISCOUNT"
	}
	return strconv.FormatInt(int64(v), 10)
}

type StorefrontPublishState int64

const (
	StorefrontPublishStateUnpublished         StorefrontPublishState = 1
	StorefrontPublishStatePublishedHidden     StorefrontPublishState = 2
	StorefrontPublishStatePublishedStorefront StorefrontPublishState = 3
)

func (v StorefrontPublishState) String() string {
	switch v {
	case StorefrontPublishStateUnpublished:
		return "UNPUBLISHED"
	case StorefrontPublishStatePublishedHidden:
		return "PUBLISHED_HIDDEN"
	case StorefrontPublishStatePublishedStorefront:
		return "PUBLISHED_STOREFRONT"
	}
	return strconv.FormatInt(int64(v), 10)
}

type StorefrontPurchaseType int64

const (
	StorefrontPurchaseTypeSelfPurchase StorefrontPurchaseType = 0
	StorefrontPurchaseTypeGift         StorefrontPurchaseType = 1
)

func (v StorefrontPurchaseType) String() string {
	switch v {
	case StorefrontPurchaseTypeSelfPurchase:
		return "SELF_PURCHASE"
	case StorefrontPurchaseTypeGift:
		return "GIFT"
	}
	return strconv.FormatInt(int64(v), 10)
}

type StreamDeleteReason string

const (
	StreamDeleteReasonUserRequested          StreamDeleteReason = "user_requested"
	StreamDeleteReasonStreamEnded            StreamDeleteReason = "stream_ended"
	StreamDeleteReasonStreamFull             StreamDeleteReason = "stream_full"
	StreamDeleteReasonUnauthorized           StreamDeleteReason = "unauthorized"
	StreamDeleteReasonSafetyGuildRateLimited StreamDeleteReason = "safety_guild_rate_limited"
	StreamDeleteReasonParseFailed            StreamDeleteReason = "parse_failed"
	StreamDeleteReasonInvalidChannel         StreamDeleteReason = "invalid_channel"
)

type StreamType string

const (
	StreamTypeGuild StreamType = "guild"
	StreamTypeCall  StreamType = "call"
	StreamTypeTest  StreamType = "test"
)

type SubblockType int64

const (
	SubblockTypeCategory SubblockType = 0
)

func (v SubblockType) String() string {
	switch v {
	case SubblockTypeCategory:
		return "CATEGORY"
	}
	return strconv.FormatInt(int64(v), 10)
}

type SubscriptionInterval int64

const (
	SubscriptionIntervalMonth SubscriptionInterval = 1
	SubscriptionIntervalYear  SubscriptionInterval = 2
	SubscriptionIntervalDay   SubscriptionInterval = 3
)

func (v SubscriptionInterval) String() string {
	switch v {
	case SubscriptionIntervalMonth:
		return "MONTH"
	case SubscriptionIntervalYear:
		return "YEAR"
	case SubscriptionIntervalDay:
		return "DAY"
	}
	return strconv.FormatInt(int64(v), 10)
}

type SubscriptionPlanPurchaseType int64

const (
	SubscriptionPlanPurchaseTypeDefault            SubscriptionPlanPurchaseType = 0
	SubscriptionPlanPurchaseTypeGift               SubscriptionPlanPurchaseType = 1
	SubscriptionPlanPurchaseTypeSale               SubscriptionPlanPurchaseType = 2
	SubscriptionPlanPurchaseTypePremiumTier1       SubscriptionPlanPurchaseType = 3
	SubscriptionPlanPurchaseTypePremiumTier2       SubscriptionPlanPurchaseType = 4
	SubscriptionPlanPurchaseTypeMobile             SubscriptionPlanPurchaseType = 5
	SubscriptionPlanPurchaseTypePremiumTier0       SubscriptionPlanPurchaseType = 6
	SubscriptionPlanPurchaseTypeMobilePremiumTier2 SubscriptionPlanPurchaseType = 7
)

func (v SubscriptionPlanPurchaseType) String() string {
	switch v {
	case SubscriptionPlanPurchaseTypeDefault:
		return "DEFAULT"
	case SubscriptionPlanPurchaseTypeGift:
		return "GIFT"
	case SubscriptionPlanPurchaseTypeSale:
		return "SALE"
	case SubscriptionPlanPurchaseTypePremiumTier1:
		return "PREMIUM_TIER_1"
	case SubscriptionPlanPurchaseTypePremiumTier2:
		return "PREMIUM_TIER_2"
	case SubscriptionPlanPurchaseTypeMobile:
		return "MOBILE"
	case SubscriptionPlanPurchaseTypePremiumTier0:
		return "PREMIUM_TIER_0"
	case SubscriptionPlanPurchaseTypeMobilePremiumTier2:
		return "MOBILE_PREMIUM_TIER_2"
	}
	return strconv.FormatInt(int64(v), 10)
}

type SubscriptionPromotionRewardEligibilityReason string

const (
	SubscriptionPromotionRewardEligibilityReasonUserEligibleForReward        SubscriptionPromotionRewardEligibilityReason = "user_eligible_for_reward"
	SubscriptionPromotionRewardEligibilityReasonUserNotEligibleForExperiment SubscriptionPromotionRewardEligibilityReason = "user_not_eligible_for_experiment"
)

type SubscriptionStatus int64

const (
	SubscriptionStatusUnpaid       SubscriptionStatus = 0
	SubscriptionStatusActive       SubscriptionStatus = 1
	SubscriptionStatusPastDue      SubscriptionStatus = 2
	SubscriptionStatusCanceled     SubscriptionStatus = 3
	SubscriptionStatusEnded        SubscriptionStatus = 4
	SubscriptionStatusAccountHold  SubscriptionStatus = 6
	SubscriptionStatusBillingRetry SubscriptionStatus = 7
	SubscriptionStatusPaused       SubscriptionStatus = 8
	SubscriptionStatusPausePending SubscriptionStatus = 9
)

func (v SubscriptionStatus) String() string {
	switch v {
	case SubscriptionStatusUnpaid:
		return "UNPAID"
	case SubscriptionStatusActive:
		return "ACTIVE"
	case SubscriptionStatusPastDue:
		return "PAST_DUE"
	case SubscriptionStatusCanceled:
		return "CANCELED"
	case SubscriptionStatusEnded:
		return "ENDED"
	case SubscriptionStatusAccountHold:
		return "ACCOUNT_HOLD"
	case SubscriptionStatusBillingRetry:
		return "BILLING_RETRY"
	case SubscriptionStatusPaused:
		return "PAUSED"
	case SubscriptionStatusPausePending:
		return "PAUSE_PENDING"
	}
	return strconv.FormatInt(int64(v), 10)
}

type SubscriptionType int64

const (
	SubscriptionTypePremium     SubscriptionType = 1
	SubscriptionTypeGuild       SubscriptionType = 2
	SubscriptionTypeApplication SubscriptionType = 3
)

func (v SubscriptionType) String() string {
	switch v {
	case SubscriptionTypePremium:
		return "PREMIUM"
	case SubscriptionTypeGuild:
		return "GUILD"
	case SubscriptionTypeApplication:
		return "APPLICATION"
	}
	return strconv.FormatInt(int64(v), 10)
}

type SummarySource int64

const (
	SummarySourceSource0 SummarySource = 0
	SummarySourceSource1 SummarySource = 1
	SummarySourceSource2 SummarySource = 2
)

func (v SummarySource) String() string {
	switch v {
	case SummarySourceSource0:
		return "SOURCE_0"
	case SummarySourceSource1:
		return "SOURCE_1"
	case SummarySourceSource2:
		return "SOURCE_2"
	}
	return strconv.FormatInt(int64(v), 10)
}

type SummaryType int64

const (
	SummaryTypeUnset   SummaryType = 0
	SummaryTypeSource1 SummaryType = 1
	SummaryTypeSource2 SummaryType = 2
	SummaryTypeUnknown SummaryType = 3
)

func (v SummaryType) String() string {
	switch v {
	case SummaryTypeUnset:
		return "UNSET"
	case SummaryTypeSource1:
		return "SOURCE_1"
	case SummaryTypeSource2:
		return "SOURCE_2"
	case SummaryTypeUnknown:
		return "UNKNOWN"
	}
	return strconv.FormatInt(int64(v), 10)
}

type SurfaceType string

const (
	SurfaceTypeWidgetTop         SurfaceType = "widget_top"
	SurfaceTypeWidgetBottom      SurfaceType = "widget_bottom"
	SurfaceTypeAddWidgetPreview  SurfaceType = "add_widget_preview"
	SurfaceTypeMiniProfile       SurfaceType = "mini_profile"
	SurfaceTypeActivityAccessory SurfaceType = "activity_accessory"
)

type SurveyRequirementType string

const (
	SurveyRequirementTypeIsOwner          SurveyRequirementType = "IS_OWNER"
	SurveyRequirementTypeIsAdmin          SurveyRequirementType = "IS_ADMIN"
	SurveyRequirementTypeIsCommunity      SurveyRequirementType = "IS_COMMUNITY"
	SurveyRequirementTypeGuildSize        SurveyRequirementType = "GUILD_SIZE"
	SurveyRequirementTypeGuildSizeAll     SurveyRequirementType = "GUILD_SIZE_ALL"
	SurveyRequirementTypeIsHub            SurveyRequirementType = "IS_HUB"
	SurveyRequirementTypeIsViewing        SurveyRequirementType = "IS_VIEWING"
	SurveyRequirementTypeGuildPermissions SurveyRequirementType = "GUILD_PERMISSIONS"
)

type SwipeRightToLeftMode int64

const (
	SwipeRightToLeftModeUnset          SwipeRightToLeftMode = 0
	SwipeRightToLeftModeChannelDetails SwipeRightToLeftMode = 1
	SwipeRightToLeftModeReply          SwipeRightToLeftMode = 2
)

func (v SwipeRightToLeftMode) String() string {
	switch v {
	case SwipeRightToLeftModeUnset:
		return "UNSET"
	case SwipeRightToLeftModeChannelDetails:
		return "CHANNEL_DETAILS"
	case SwipeRightToLeftModeReply:
		return "REPLY"
	}
	return strconv.FormatInt(int64(v), 10)
}

type SystemChannelFlags uint64

const (
	SystemChannelFlagsSuppressJoinNotifications                           SystemChannelFlags = 1 << 0
	SystemChannelFlagsSuppressPremiumSubscriptions                        SystemChannelFlags = 1 << 1
	SystemChannelFlagsSuppressGuildReminderNotifications                  SystemChannelFlags = 1 << 2
	SystemChannelFlagsSuppressJoinNotificationReplies                     SystemChannelFlags = 1 << 3
	SystemChannelFlagsSuppressRoleSubscriptionPurchaseNotifications       SystemChannelFlags = 1 << 4
	SystemChannelFlagsSuppressRoleSubscriptionPurchaseNotificationReplies SystemChannelFlags = 1 << 5
	SystemChannelFlagsSuppressChannelPromptDeadchat                       SystemChannelFlags = 1 << 7
	SystemChannelFlagsSuppressUgcAddedNotifications                       SystemChannelFlags = 1 << 8
)

func (v SystemChannelFlags) String() string {
	switch v {
	case SystemChannelFlagsSuppressJoinNotifications:
		return "SUPPRESS_JOIN_NOTIFICATIONS"
	case SystemChannelFlagsSuppressPremiumSubscriptions:
		return "SUPPRESS_PREMIUM_SUBSCRIPTIONS"
	case SystemChannelFlagsSuppressGuildReminderNotifications:
		return "SUPPRESS_GUILD_REMINDER_NOTIFICATIONS"
	case SystemChannelFlagsSuppressJoinNotificationReplies:
		return "SUPPRESS_JOIN_NOTIFICATION_REPLIES"
	case SystemChannelFlagsSuppressRoleSubscriptionPurchaseNotifications:
		return "SUPPRESS_ROLE_SUBSCRIPTION_PURCHASE_NOTIFICATIONS"
	case SystemChannelFlagsSuppressRoleSubscriptionPurchaseNotificationReplies:
		return "SUPPRESS_ROLE_SUBSCRIPTION_PURCHASE_NOTIFICATION_REPLIES"
	case SystemChannelFlagsSuppressChannelPromptDeadchat:
		return "SUPPRESS_CHANNEL_PROMPT_DEADCHAT"
	case SystemChannelFlagsSuppressUgcAddedNotifications:
		return "SUPPRESS_UGC_ADDED_NOTIFICATIONS"
	}
	return strconv.FormatUint(uint64(v), 10)
}

type TeamMemberRoleTypes string

const (
	TeamMemberRoleTypesAdmin     TeamMemberRoleTypes = "admin"
	TeamMemberRoleTypesDeveloper TeamMemberRoleTypes = "developer"
	TeamMemberRoleTypesReadOnly  TeamMemberRoleTypes = "read_only"
)

type TeamMembershipState int64

const (
	TeamMembershipStateInvited  TeamMembershipState = 1
	TeamMembershipStateAccepted TeamMembershipState = 2
)

func (v TeamMembershipState) String() string {
	switch v {
	case TeamMembershipStateInvited:
		return "INVITED"
	case TeamMembershipStateAccepted:
		return "ACCEPTED"
	}
	return strconv.FormatInt(int64(v), 10)
}

type TeamPayoutAccountStatus int64

const (
	TeamPayoutAccountStatusUnsubmitted    TeamPayoutAccountStatus = 1
	TeamPayoutAccountStatusPending        TeamPayoutAccountStatus = 2
	TeamPayoutAccountStatusActionRequired TeamPayoutAccountStatus = 3
	TeamPayoutAccountStatusActive         TeamPayoutAccountStatus = 4
	TeamPayoutAccountStatusBlocked        TeamPayoutAccountStatus = 5
	TeamPayoutAccountStatusSuspended      TeamPayoutAccountStatus = 6
)

func (v TeamPayoutAccountStatus) String() string {
	switch v {
	case TeamPayoutAccountStatusUnsubmitted:
		return "UNSUBMITTED"
	case TeamPayoutAccountStatusPending:
		return "PENDING"
	case TeamPayoutAccountStatusActionRequired:
		return "ACTION_REQUIRED"
	case TeamPayoutAccountStatusActive:
		return "ACTIVE"
	case TeamPayoutAccountStatusBlocked:
		return "BLOCKED"
	case TeamPayoutAccountStatusSuspended:
		return "SUSPENDED"
	}
	return strconv.FormatInt(int64(v), 10)
}

type TeamPayoutGateway int64

const (
	TeamPayoutGatewayStripeTopup   TeamPayoutGateway = 1
	TeamPayoutGatewayTipalti       TeamPayoutGateway = 2
	TeamPayoutGatewayStripePrimary TeamPayoutGateway = 3
)

func (v TeamPayoutGateway) String() string {
	switch v {
	case TeamPayoutGatewayStripeTopup:
		return "STRIPE_TOPUP"
	case TeamPayoutGatewayTipalti:
		return "TIPALTI"
	case TeamPayoutGatewayStripePrimary:
		return "STRIPE_PRIMARY"
	}
	return strconv.FormatInt(int64(v), 10)
}

type TeamPayoutReportType string

const (
	TeamPayoutReportTypeSKU         TeamPayoutReportType = "sku"
	TeamPayoutReportTypeTransaction TeamPayoutReportType = "transaction"
)

type TeamPayoutStatus int64

const (
	TeamPayoutStatusOpen             TeamPayoutStatus = 1
	TeamPayoutStatusPaid             TeamPayoutStatus = 2
	TeamPayoutStatusPending          TeamPayoutStatus = 3
	TeamPayoutStatusManual           TeamPayoutStatus = 4
	TeamPayoutStatusCancelled        TeamPayoutStatus = 5
	TeamPayoutStatusDeferred         TeamPayoutStatus = 6
	TeamPayoutStatusDeferredInternal TeamPayoutStatus = 7
	TeamPayoutStatusProcessing       TeamPayoutStatus = 8
	TeamPayoutStatusError            TeamPayoutStatus = 9
	TeamPayoutStatusRejected         TeamPayoutStatus = 10
	TeamPayoutStatusRiskReview       TeamPayoutStatus = 11
	TeamPayoutStatusSubmitted        TeamPayoutStatus = 12
	TeamPayoutStatusPendingFunds     TeamPayoutStatus = 13
)

func (v TeamPayoutStatus) String() string {
	switch v {
	case TeamPayoutStatusOpen:
		return "OPEN"
	case TeamPayoutStatusPaid:
		return "PAID"
	case TeamPayoutStatusPending:
		return "PENDING"
	case TeamPayoutStatusManual:
		return "MANUAL"
	case TeamPayoutStatusCancelled:
		return "CANCELLED"
	case TeamPayoutStatusDeferred:
		return "DEFERRED"
	case TeamPayoutStatusDeferredInternal:
		return "DEFERRED_INTERNAL"
	case TeamPayoutStatusProcessing:
		return "PROCESSING"
	case TeamPayoutStatusError:
		return "ERROR"
	case TeamPayoutStatusRejected:
		return "REJECTED"
	case TeamPayoutStatusRiskReview:
		return "RISK_REVIEW"
	case TeamPayoutStatusSubmitted:
		return "SUBMITTED"
	case TeamPayoutStatusPendingFunds:
		return "PENDING_FUNDS"
	}
	return strconv.FormatInt(int64(v), 10)
}

type TeenActionDisplayType int64

const (
	TeenActionDisplayTypeUserAdd           TeenActionDisplayType = 1
	TeenActionDisplayTypeGuildAdd          TeenActionDisplayType = 2
	TeenActionDisplayTypeUserInteraction   TeenActionDisplayType = 3
	TeenActionDisplayTypeGuildInteraction  TeenActionDisplayType = 4
	TeenActionDisplayTypeUserCalled        TeenActionDisplayType = 5
	TeenActionDisplayTypeTotalVoiceMinutes TeenActionDisplayType = 6
	TeenActionDisplayTypePurchases         TeenActionDisplayType = 7
	TeenActionDisplayTypeGifts             TeenActionDisplayType = 8
)

func (v TeenActionDisplayType) String() string {
	switch v {
	case TeenActionDisplayTypeUserAdd:
		return "USER_ADD"
	case TeenActionDisplayTypeGuildAdd:
		return "GUILD_ADD"
	case TeenActionDisplayTypeUserInteraction:
		return "USER_INTERACTION"
	case TeenActionDisplayTypeGuildInteraction:
		return "GUILD_INTERACTION"
	case TeenActionDisplayTypeUserCalled:
		return "USER_CALLED"
	case TeenActionDisplayTypeTotalVoiceMinutes:
		return "TOTAL_VOICE_MINUTES"
	case TeenActionDisplayTypePurchases:
		return "PURCHASES"
	case TeenActionDisplayTypeGifts:
		return "GIFTS"
	}
	return strconv.FormatInt(int64(v), 10)
}

type TeenActionType int64

const (
	TeenActionTypeDMMessageSend          TeenActionType = 1
	TeenActionTypeGdmMessageSend         TeenActionType = 2
	TeenActionTypeMessageReact           TeenActionType = 3
	TeenActionTypeAddFriend              TeenActionType = 4
	TeenActionTypeSendCall               TeenActionType = 5
	TeenActionTypeCallJoin               TeenActionType = 6
	TeenActionTypeGuildJoin              TeenActionType = 7
	TeenActionTypeGuildMessageSend       TeenActionType = 8
	TeenActionTypeGuildVcJoin            TeenActionType = 9
	TeenActionTypeGuildVoiceChannelLeave TeenActionType = 10
	TeenActionTypeCallLeave              TeenActionType = 11
	TeenActionTypeCallStart              TeenActionType = 12
	TeenActionTypeInvoiceComplete        TeenActionType = 13
)

func (v TeenActionType) String() string {
	switch v {
	case TeenActionTypeDMMessageSend:
		return "DM_MESSAGE_SEND"
	case TeenActionTypeGdmMessageSend:
		return "GDM_MESSAGE_SEND"
	case TeenActionTypeMessageReact:
		return "MESSAGE_REACT"
	case TeenActionTypeAddFriend:
		return "ADD_FRIEND"
	case TeenActionTypeSendCall:
		return "SEND_CALL"
	case TeenActionTypeCallJoin:
		return "CALL_JOIN"
	case TeenActionTypeGuildJoin:
		return "GUILD_JOIN"
	case TeenActionTypeGuildMessageSend:
		return "GUILD_MESSAGE_SEND"
	case TeenActionTypeGuildVcJoin:
		return "GUILD_VC_JOIN"
	case TeenActionTypeGuildVoiceChannelLeave:
		return "GUILD_VOICE_CHANNEL_LEAVE"
	case TeenActionTypeCallLeave:
		return "CALL_LEAVE"
	case TeenActionTypeCallStart:
		return "CALL_START"
	case TeenActionTypeInvoiceComplete:
		return "INVOICE_COMPLETE"
	}
	return strconv.FormatInt(int64(v), 10)
}

type TeenByDefaultSettings uint64

const (
	TeenByDefaultSettingsSensitiveContent                  TeenByDefaultSettings = 1 << 0
	TeenByDefaultSettingsFriendRequestStrangerConfirmation TeenByDefaultSettings = 1 << 1
	TeenByDefaultSettingsMessageRequestRestrictions        TeenByDefaultSettings = 1 << 2
	TeenByDefaultSettingsGuildActivityStatus               TeenByDefaultSettings = 1 << 3
	TeenByDefaultSettingsSpamFilters                       TeenByDefaultSettings = 1 << 4
)

func (v TeenByDefaultSettings) String() string {
	switch v {
	case TeenByDefaultSettingsSensitiveContent:
		return "SENSITIVE_CONTENT"
	case TeenByDefaultSettingsFriendRequestStrangerConfirmation:
		return "FRIEND_REQUEST_STRANGER_CONFIRMATION"
	case TeenByDefaultSettingsMessageRequestRestrictions:
		return "MESSAGE_REQUEST_RESTRICTIONS"
	case TeenByDefaultSettingsGuildActivityStatus:
		return "GUILD_ACTIVITY_STATUS"
	case TeenByDefaultSettingsSpamFilters:
		return "SPAM_FILTERS"
	}
	return strconv.FormatUint(uint64(v), 10)
}

type TextInputStyle int64

const (
	TextInputStyleSmall     TextInputStyle = 1
	TextInputStyleParagraph TextInputStyle = 2
)

func (v TextInputStyle) String() string {
	switch v {
	case TextInputStyleSmall:
		return "SMALL"
	case TextInputStyleParagraph:
		return "PARAGRAPH"
	}
	return strconv.FormatInt(int64(v), 10)
}

type Theme string

const (
	ThemeDark     Theme = "dark"
	ThemeLight    Theme = "light"
	ThemeDarker   Theme = "darker"
	ThemeMidnight Theme = "midnight"
)

type ThermalState int64

const (
	ThermalStateNominal  ThermalState = 0
	ThermalStateFair     ThermalState = 1
	ThermalStateSerious  ThermalState = 2
	ThermalStateCritical ThermalState = 3
)

func (v ThermalState) String() string {
	switch v {
	case ThermalStateNominal:
		return "NOMINAL"
	case ThermalStateFair:
		return "FAIR"
	case ThermalStateSerious:
		return "SERIOUS"
	case ThermalStateCritical:
		return "CRITICAL"
	}
	return strconv.FormatInt(int64(v), 10)
}

type ThreadMemberFlags uint64

const (
	ThreadMemberFlagsHasInteracted ThreadMemberFlags = 1 << 0
	ThreadMemberFlagsAllMessages   ThreadMemberFlags = 1 << 1
	ThreadMemberFlagsOnlyMentions  ThreadMemberFlags = 1 << 2
	ThreadMemberFlagsNoMessages    ThreadMemberFlags = 1 << 3
)

func (v ThreadMemberFlags) String() string {
	switch v {
	case ThreadMemberFlagsHasInteracted:
		return "HAS_INTERACTED"
	case ThreadMemberFlagsAllMessages:
		return "ALL_MESSAGES"
	case ThreadMemberFlagsOnlyMentions:
		return "ONLY_MENTIONS"
	case ThreadMemberFlagsNoMessages:
		return "NO_MESSAGES"
	}
	return strconv.FormatUint(uint64(v), 10)
}

type ThreadSortMode string

const (
	ThreadSortModeLastMessageTime ThreadSortMode = "last_message_time"
	ThreadSortModeArchiveTime     ThreadSortMode = "archive_time"
	ThreadSortModeRelevance       ThreadSortMode = "relevance"
	ThreadSortModeCreationTime    ThreadSortMode = "creation_time"
)

type TimestampHourCycle int64

const (
	TimestampHourCycleAuto TimestampHourCycle = 0
	TimestampHourCycleH12  TimestampHourCycle = 1
	TimestampHourCycleH23  TimestampHourCycle = 2
)

func (v TimestampHourCycle) String() string {
	switch v {
	case TimestampHourCycleAuto:
		return "AUTO"
	case TimestampHourCycleH12:
		return "H12"
	case TimestampHourCycleH23:
		return "H23"
	}
	return strconv.FormatInt(int64(v), 10)
}

type TrackingType string

const (
	TrackingTypeGenericFriendRequestReminder TrackingType = "generic_friend_request_reminder"
	TrackingTypeGenericHomeFeaturedMessage   TrackingType = "generic_home_featured_message"
	TrackingTypeGenericMissedMessage         TrackingType = "generic_missed_message"
	TrackingTypeGuildStreamStart             TrackingType = "GUILD_STREAM_START"
	TrackingTypeHomeLifecyclePush            TrackingType = "home_lifecycle_push"
	TrackingTypeNewUserSession               TrackingType = "new_user_session"
	TrackingTypeNudgeNewFriendDMPush         TrackingType = "nudge_new_friend_dm_push"
	TrackingTypePollEnded                    TrackingType = "poll_ended"
	TrackingTypeReactionsPushNotification    TrackingType = "reactions_push_notification"
	TrackingTypeReminder                     TrackingType = "reminder"
	TrackingTypeSuspiciousSession            TrackingType = "suspicious_session"
	TrackingTypeTopMessagesPush              TrackingType = "top_messages_push"
	TrackingTypeTrendingContentPush          TrackingType = "trending_content_push"
)

type TransportEncryptionMode string

const (
	TransportEncryptionModeAeadAes256GcmRtpSize         TransportEncryptionMode = "aead_aes256_gcm_rtpsize"
	TransportEncryptionModeAeadXchacha20Poly1305RtpSize TransportEncryptionMode = "aead_xchacha20_poly1305_rtpsize"
	TransportEncryptionModeXsalsa20Poly1305LiteRtpSize  TransportEncryptionMode = "xsalsa20_poly1305_lite_rtpsize"
	TransportEncryptionModeAeadAes256Gcm                TransportEncryptionMode = "aead_aes256_gcm"
	TransportEncryptionModeXsalsa20Poly1305             TransportEncryptionMode = "xsalsa20_poly1305"
	TransportEncryptionModeXsalsa20Poly1305Suffix       TransportEncryptionMode = "xsalsa20_poly1305_suffix"
	TransportEncryptionModeXsalsa20Poly1305Lite         TransportEncryptionMode = "xsalsa20_poly1305_lite"
)

type TwoWayLinkType string

const (
	TwoWayLinkTypeWeb     TwoWayLinkType = "web"
	TwoWayLinkTypeMobile  TwoWayLinkType = "mobile"
	TwoWayLinkTypeDesktop TwoWayLinkType = "desktop"
)

type UIDensity int64

const (
	UIDensityUnset      UIDensity = 0
	UIDensityCompact    UIDensity = 1
	UIDensityCozy       UIDensity = 2
	UIDensityResponsive UIDensity = 3
	UIDensityDefault    UIDensity = 4
)

func (v UIDensity) String() string {
	switch v {
	case UIDensityUnset:
		return "UNSET"
	case UIDensityCompact:
		return "COMPACT"
	case UIDensityCozy:
		return "COZY"
	case UIDensityResponsive:
		return "RESPONSIVE"
	case UIDensityDefault:
		return "DEFAULT"
	}
	return strconv.FormatInt(int64(v), 10)
}

type UmbrellaOAuth2Scopes string

const (
	UmbrellaOAuth2ScopesSdkSocialLayerPresence UmbrellaOAuth2Scopes = "sdk.social_layer_presence"
	UmbrellaOAuth2ScopesSdkSocialLayer         UmbrellaOAuth2Scopes = "sdk.social_layer"
)

type UnfurledMediaItemLoadingState int64

const (
	UnfurledMediaItemLoadingStateUnknown        UnfurledMediaItemLoadingState = 0
	UnfurledMediaItemLoadingStateLoading        UnfurledMediaItemLoadingState = 1
	UnfurledMediaItemLoadingStateLoadedSuccess  UnfurledMediaItemLoadingState = 2
	UnfurledMediaItemLoadingStateLoadedNotFound UnfurledMediaItemLoadingState = 3
)

func (v UnfurledMediaItemLoadingState) String() string {
	switch v {
	case UnfurledMediaItemLoadingStateUnknown:
		return "UNKNOWN"
	case UnfurledMediaItemLoadingStateLoading:
		return "LOADING"
	case UnfurledMediaItemLoadingStateLoadedSuccess:
		return "LOADED_SUCCESS"
	case UnfurledMediaItemLoadingStateLoadedNotFound:
		return "LOADED_NOT_FOUND"
	}
	return strconv.FormatInt(int64(v), 10)
}

type UserApplicationProfileDynamicFieldType int64

const (
	UserApplicationProfileDynamicFieldTypeText   UserApplicationProfileDynamicFieldType = 1
	UserApplicationProfileDynamicFieldTypeNumber UserApplicationProfileDynamicFieldType = 2
	UserApplicationProfileDynamicFieldTypeImage  UserApplicationProfileDynamicFieldType = 3
)

func (v UserApplicationProfileDynamicFieldType) String() string {
	switch v {
	case UserApplicationProfileDynamicFieldTypeText:
		return "TEXT"
	case UserApplicationProfileDynamicFieldTypeNumber:
		return "NUMBER"
	case UserApplicationProfileDynamicFieldTypeImage:
		return "IMAGE"
	}
	return strconv.FormatInt(int64(v), 10)
}

type UserAuthenticatorType int64

const (
	UserAuthenticatorTypeWebauthn UserAuthenticatorType = 1
	UserAuthenticatorTypeTotp     UserAuthenticatorType = 2
	UserAuthenticatorTypeSMS      UserAuthenticatorType = 3
)

func (v UserAuthenticatorType) String() string {
	switch v {
	case UserAuthenticatorTypeWebauthn:
		return "WEBAUTHN"
	case UserAuthenticatorTypeTotp:
		return "TOTP"
	case UserAuthenticatorTypeSMS:
		return "SMS"
	}
	return strconv.FormatInt(int64(v), 10)
}

type UserFlags uint64

const (
	UserFlagsStaff                            UserFlags = 1 << 0
	UserFlagsPartner                          UserFlags = 1 << 1
	UserFlagsHypesquad                        UserFlags = 1 << 2
	UserFlagsBugHunterLevel1                  UserFlags = 1 << 3
	UserFlagsMFASMS                           UserFlags = 1 << 4
	UserFlagsPremiumPromoDismissed            UserFlags = 1 << 5
	UserFlagsHypesquadOnlineHouse1            UserFlags = 1 << 6
	UserFlagsHypesquadOnlineHouse2            UserFlags = 1 << 7
	UserFlagsHypesquadOnlineHouse3            UserFlags = 1 << 8
	UserFlagsPremiumEarlySupporter            UserFlags = 1 << 9
	UserFlagsTeamPseudoUser                   UserFlags = 1 << 10
	UserFlagsIsHubspotContact                 UserFlags = 1 << 11
	UserFlagsSystem                           UserFlags = 1 << 12
	UserFlagsHasUnreadUrgentMessages          UserFlags = 1 << 13
	UserFlagsBugHunterLevel2                  UserFlags = 1 << 14
	UserFlagsUnderageDeleted                  UserFlags = 1 << 15
	UserFlagsVerifiedBot                      UserFlags = 1 << 16
	UserFlagsVerifiedDeveloper                UserFlags = 1 << 17
	UserFlagsCertifiedModerator               UserFlags = 1 << 18
	UserFlagsBotHTTPInteractions              UserFlags = 1 << 19
	UserFlagsSpammer                          UserFlags = 1 << 20
	UserFlagsDisablePremium                   UserFlags = 1 << 21
	UserFlagsActiveDeveloper                  UserFlags = 1 << 22
	UserFlagsProvisionalAccount               UserFlags = 1 << 23
	UserFlagsHighGlobalRateLimit              UserFlags = 1 << 33
	UserFlagsDeleted                          UserFlags = 1 << 34
	UserFlagsDisabledSuspiciousActivity       UserFlags = 1 << 35
	UserFlagsSelfDeleted                      UserFlags = 1 << 36
	UserFlagsPremiumDiscriminator             UserFlags = 1 << 37
	UserFlagsUsedDesktopClient                UserFlags = 1 << 38
	UserFlagsUsedWebClient                    UserFlags = 1 << 39
	UserFlagsUsedMobileClient                 UserFlags = 1 << 40
	UserFlagsDisabled                         UserFlags = 1 << 41
	UserFlagsHasSessionStarted                UserFlags = 1 << 43
	UserFlagsQuarantined                      UserFlags = 1 << 44
	UserFlagsPremiumEligibleForUniqueUsername UserFlags = 1 << 47
	UserFlagsCollaborator                     UserFlags = 1 << 50
	UserFlagsRestrictedCollaborator           UserFlags = 1 << 51
)

func (v UserFlags) String() string {
	switch v {
	case UserFlagsStaff:
		return "STAFF"
	case UserFlagsPartner:
		return "PARTNER"
	case UserFlagsHypesquad:
		return "HYPESQUAD"
	case UserFlagsBugHunterLevel1:
		return "BUG_HUNTER_LEVEL_1"
	case UserFlagsMFASMS:
		return "MFA_SMS"
	case UserFlagsPremiumPromoDismissed:
		return "PREMIUM_PROMO_DISMISSED"
	case UserFlagsHypesquadOnlineHouse1:
		return "HYPESQUAD_ONLINE_HOUSE_1"
	case UserFlagsHypesquadOnlineHouse2:
		return "HYPESQUAD_ONLINE_HOUSE_2"
	case UserFlagsHypesquadOnlineHouse3:
		return "HYPESQUAD_ONLINE_HOUSE_3"
	case UserFlagsPremiumEarlySupporter:
		return "PREMIUM_EARLY_SUPPORTER"
	case UserFlagsTeamPseudoUser:
		return "TEAM_PSEUDO_USER"
	case UserFlagsIsHubspotContact:
		return "IS_HUBSPOT_CONTACT"
	case UserFlagsSystem:
		return "SYSTEM"
	case UserFlagsHasUnreadUrgentMessages:
		return "HAS_UNREAD_URGENT_MESSAGES"
	case UserFlagsBugHunterLevel2:
		return "BUG_HUNTER_LEVEL_2"
	case UserFlagsUnderageDeleted:
		return "UNDERAGE_DELETED"
	case UserFlagsVerifiedBot:
		return "VERIFIED_BOT"
	case UserFlagsVerifiedDeveloper:
		return "VERIFIED_DEVELOPER"
	case UserFlagsCertifiedModerator:
		return "CERTIFIED_MODERATOR"
	case UserFlagsBotHTTPInteractions:
		return "BOT_HTTP_INTERACTIONS"
	case UserFlagsSpammer:
		return "SPAMMER"
	case UserFlagsDisablePremium:
		return "DISABLE_PREMIUM"
	case UserFlagsActiveDeveloper:
		return "ACTIVE_DEVELOPER"
	case UserFlagsProvisionalAccount:
		return "PROVISIONAL_ACCOUNT"
	case UserFlagsHighGlobalRateLimit:
		return "HIGH_GLOBAL_RATE_LIMIT"
	case UserFlagsDeleted:
		return "DELETED"
	case UserFlagsDisabledSuspiciousActivity:
		return "DISABLED_SUSPICIOUS_ACTIVITY"
	case UserFlagsSelfDeleted:
		return "SELF_DELETED"
	case UserFlagsPremiumDiscriminator:
		return "PREMIUM_DISCRIMINATOR"
	case UserFlagsUsedDesktopClient:
		return "USED_DESKTOP_CLIENT"
	case UserFlagsUsedWebClient:
		return "USED_WEB_CLIENT"
	case UserFlagsUsedMobileClient:
		return "USED_MOBILE_CLIENT"
	case UserFlagsDisabled:
		return "DISABLED"
	case UserFlagsHasSessionStarted:
		return "HAS_SESSION_STARTED"
	case UserFlagsQuarantined:
		return "QUARANTINED"
	case UserFlagsPremiumEligibleForUniqueUsername:
		return "PREMIUM_ELIGIBLE_FOR_UNIQUE_USERNAME"
	case UserFlagsCollaborator:
		return "COLLABORATOR"
	case UserFlagsRestrictedCollaborator:
		return "RESTRICTED_COLLABORATOR"
	}
	return strconv.FormatUint(uint64(v), 10)
}

type UserGuildSettingsFlags uint64

const (
	UserGuildSettingsFlagsUnreadsAllMessages  UserGuildSettingsFlags = 1 << 11
	UserGuildSettingsFlagsUnreadsOnlyMentions UserGuildSettingsFlags = 1 << 12
	UserGuildSettingsFlagsOptInChannelsOff    UserGuildSettingsFlags = 1 << 13
	UserGuildSettingsFlagsOptInChannelsOn     UserGuildSettingsFlags = 1 << 14
)

func (v UserGuildSettingsFlags) String() string {
	switch v {
	case UserGuildSettingsFlagsUnreadsAllMessages:
		return "UNREADS_ALL_MESSAGES"
	case UserGuildSettingsFlagsUnreadsOnlyMentions:
		return "UNREADS_ONLY_MENTIONS"
	case UserGuildSettingsFlagsOptInChannelsOff:
		return "OPT_IN_CHANNELS_OFF"
	case UserGuildSettingsFlagsOptInChannelsOn:
		return "OPT_IN_CHANNELS_ON"
	}
	return strconv.FormatUint(uint64(v), 10)
}

type UserIdentityVerificationError int64

const (
	UserIdentityVerificationErrorConsentDeclined                       UserIdentityVerificationError = 1
	UserIdentityVerificationErrorUnverified                            UserIdentityVerificationError = 2
	UserIdentityVerificationErrorDeviceUnsupported                     UserIdentityVerificationError = 3
	UserIdentityVerificationErrorVerificationDocumentExpired           UserIdentityVerificationError = 4
	UserIdentityVerificationErrorVerificationDocumentInvalid           UserIdentityVerificationError = 5
	UserIdentityVerificationErrorVerificationUnexpectedDocumentCountry UserIdentityVerificationError = 6
	UserIdentityVerificationErrorVerificationUnexpectedDocumentType    UserIdentityVerificationError = 7
	UserIdentityVerificationErrorVerificationScanNotReadable           UserIdentityVerificationError = 8
	UserIdentityVerificationErrorVerificationScanMissingBack           UserIdentityVerificationError = 9
	UserIdentityVerificationErrorVerificationScanIDTypeNotSupported    UserIdentityVerificationError = 10
	UserIdentityVerificationErrorVerificationScanCorrupt               UserIdentityVerificationError = 11
	UserIdentityVerificationErrorVerificationScanFailedCopy            UserIdentityVerificationError = 12
	UserIdentityVerificationErrorVerificationScanManipulatedDocument   UserIdentityVerificationError = 13
	UserIdentityVerificationErrorVerificationScanFailedGrayscale       UserIdentityVerificationError = 14
	UserIdentityVerificationErrorVerificationUnderSupportedAge         UserIdentityVerificationError = 15
)

func (v UserIdentityVerificationError) String() string {
	switch v {
	case UserIdentityVerificationErrorConsentDeclined:
		return "CONSENT_DECLINED"
	case UserIdentityVerificationErrorUnverified:
		return "UNVERIFIED"
	case UserIdentityVerificationErrorDeviceUnsupported:
		return "DEVICE_UNSUPPORTED"
	case UserIdentityVerificationErrorVerificationDocumentExpired:
		return "VERIFICATION_DOCUMENT_EXPIRED"
	case UserIdentityVerificationErrorVerificationDocumentInvalid:
		return "VERIFICATION_DOCUMENT_INVALID"
	case UserIdentityVerificationErrorVerificationUnexpectedDocumentCountry:
		return "VERIFICATION_UNEXPECTED_DOCUMENT_COUNTRY"
	case UserIdentityVerificationErrorVerificationUnexpectedDocumentType:
		return "VERIFICATION_UNEXPECTED_DOCUMENT_TYPE"
	case UserIdentityVerificationErrorVerificationScanNotReadable:
		return "VERIFICATION_SCAN_NOT_READABLE"
	case UserIdentityVerificationErrorVerificationScanMissingBack:
		return "VERIFICATION_SCAN_MISSING_BACK"
	case UserIdentityVerificationErrorVerificationScanIDTypeNotSupported:
		return "VERIFICATION_SCAN_ID_TYPE_NOT_SUPPORTED"
	case UserIdentityVerificationErrorVerificationScanCorrupt:
		return "VERIFICATION_SCAN_CORRUPT"
	case UserIdentityVerificationErrorVerificationScanFailedCopy:
		return "VERIFICATION_SCAN_FAILED_COPY"
	case UserIdentityVerificationErrorVerificationScanManipulatedDocument:
		return "VERIFICATION_SCAN_MANIPULATED_DOCUMENT"
	case UserIdentityVerificationErrorVerificationScanFailedGrayscale:
		return "VERIFICATION_SCAN_FAILED_GRAYSCALE"
	case UserIdentityVerificationErrorVerificationUnderSupportedAge:
		return "VERIFICATION_UNDER_SUPPORTED_AGE"
	}
	return strconv.FormatInt(int64(v), 10)
}

type UserIdentityVerificationStatus int64

const (
	UserIdentityVerificationStatusRequiresAction       UserIdentityVerificationStatus = 1
	UserIdentityVerificationStatusProcessing           UserIdentityVerificationStatus = 2
	UserIdentityVerificationStatusCanceled             UserIdentityVerificationStatus = 3
	UserIdentityVerificationStatusSucceeded            UserIdentityVerificationStatus = 4
	UserIdentityVerificationStatusManuallySucceeded    UserIdentityVerificationStatus = 5
	UserIdentityVerificationStatusDeleted              UserIdentityVerificationStatus = 6
	UserIdentityVerificationStatusSucceededGracePeriod UserIdentityVerificationStatus = 7
)

func (v UserIdentityVerificationStatus) String() string {
	switch v {
	case UserIdentityVerificationStatusRequiresAction:
		return "REQUIRES_ACTION"
	case UserIdentityVerificationStatusProcessing:
		return "PROCESSING"
	case UserIdentityVerificationStatusCanceled:
		return "CANCELED"
	case UserIdentityVerificationStatusSucceeded:
		return "SUCCEEDED"
	case UserIdentityVerificationStatusManuallySucceeded:
		return "MANUALLY_SUCCEEDED"
	case UserIdentityVerificationStatusDeleted:
		return "DELETED"
	case UserIdentityVerificationStatusSucceededGracePeriod:
		return "SUCCEEDED_GRACE_PERIOD"
	}
	return strconv.FormatInt(int64(v), 10)
}

type UserLazyPerkSyncLevel int64

const (
	UserLazyPerkSyncLevelNone               UserLazyPerkSyncLevel = 0
	UserLazyPerkSyncLevelAddPerksIfDetected UserLazyPerkSyncLevel = 1
	UserLazyPerkSyncLevelFullResync         UserLazyPerkSyncLevel = 2
)

func (v UserLazyPerkSyncLevel) String() string {
	switch v {
	case UserLazyPerkSyncLevelNone:
		return "NONE"
	case UserLazyPerkSyncLevelAddPerksIfDetected:
		return "ADD_PERKS_IF_DETECTED"
	case UserLazyPerkSyncLevelFullResync:
		return "FULL_RESYNC"
	}
	return strconv.FormatInt(int64(v), 10)
}

type UserSegmentType string

const (
	UserSegmentTypeHfuMau    UserSegmentType = "HFU_MAU"
	UserSegmentTypeNonHfuMau UserSegmentType = "NON_HFU_MAU"
	UserSegmentTypeNonMau    UserSegmentType = "NON_MAU"
)

type UserSettingsProtoFriendSourceFlags uint64

const (
	UserSettingsProtoFriendSourceFlagsMutualFriends UserSettingsProtoFriendSourceFlags = 1 << 1
	UserSettingsProtoFriendSourceFlagsMutualGuilds  UserSettingsProtoFriendSourceFlags = 1 << 2
	UserSettingsProtoFriendSourceFlagsNoRelation    UserSettingsProtoFriendSourceFlags = 1 << 3
)

func (v UserSettingsProtoFriendSourceFlags) String() string {
	switch v {
	case UserSettingsProtoFriendSourceFlagsMutualFriends:
		return "MUTUAL_FRIENDS"
	case UserSettingsProtoFriendSourceFlagsMutualGuilds:
		return "MUTUAL_GUILDS"
	case UserSettingsProtoFriendSourceFlagsNoRelation:
		return "NO_RELATION"
	}
	return strconv.FormatUint(uint64(v), 10)
}

type UserSettingsProtoTheme int64

const (
	UserSettingsProtoThemeUnset    UserSettingsProtoTheme = 0
	UserSettingsProtoThemeDark     UserSettingsProtoTheme = 1
	UserSettingsProtoThemeLight    UserSettingsProtoTheme = 2
	UserSettingsProtoThemeDarker   UserSettingsProtoTheme = 3
	UserSettingsProtoThemeMidnight UserSettingsProtoTheme = 4
)

func (v UserSettingsProtoTheme) String() string {
	switch v {
	case UserSettingsProtoThemeUnset:
		return "UNSET"
	case UserSettingsProtoThemeDark:
		return "DARK"
	case UserSettingsProtoThemeLight:
		return "LIGHT"
	case UserSettingsProtoThemeDarker:
		return "DARKER"
	case UserSettingsProtoThemeMidnight:
		return "MIDNIGHT"
	}
	return strconv.FormatInt(int64(v), 10)
}

type UserSettingsProtoType int64

const (
	UserSettingsProtoTypePreloaded    UserSettingsProtoType = 1
	UserSettingsProtoTypeFrecency     UserSettingsProtoType = 2
	UserSettingsProtoTypeTestSettings UserSettingsProtoType = 3
)

func (v UserSettingsProtoType) String() string {
	switch v {
	case UserSettingsProtoTypePreloaded:
		return "PRELOADED"
	case UserSettingsProtoTypeFrecency:
		return "FRECENCY"
	case UserSettingsProtoTypeTestSettings:
		return "TEST_SETTINGS"
	}
	return strconv.FormatInt(int64(v), 10)
}

type VariantsReturnStyle int64

const (
	VariantsReturnStyleIndividualProducts VariantsReturnStyle = 1
	VariantsReturnStyleVariantsGroup      VariantsReturnStyle = 2
)

func (v VariantsReturnStyle) String() string {
	switch v {
	case VariantsReturnStyleIndividualProducts:
		return "INDIVIDUAL_PRODUCTS"
	case VariantsReturnStyleVariantsGroup:
		return "VARIANTS_GROUP"
	}
	return strconv.FormatInt(int64(v), 10)
}

type VerificationLevel int64

const (
	VerificationLevelNone     VerificationLevel = 0
	VerificationLevelLow      VerificationLevel = 1
	VerificationLevelMedium   VerificationLevel = 2
	VerificationLevelHigh     VerificationLevel = 3
	VerificationLevelVeryHigh VerificationLevel = 4
)

func (v VerificationLevel) String() string {
	switch v {
	case VerificationLevelNone:
		return "NONE"
	case VerificationLevelLow:
		return "LOW"
	case VerificationLevelMedium:
		return "MEDIUM"
	case VerificationLevelHigh:
		return "HIGH"
	case VerificationLevelVeryHigh:
		return "VERY_HIGH"
	}
	return strconv.FormatInt(int64(v), 10)
}

type VideoFilterBackgroundPreset int64

const (
	VideoFilterBackgroundPresetN0  VideoFilterBackgroundPreset = 0
	VideoFilterBackgroundPresetN1  VideoFilterBackgroundPreset = 1
	VideoFilterBackgroundPresetN2  VideoFilterBackgroundPreset = 2
	VideoFilterBackgroundPresetN3  VideoFilterBackgroundPreset = 3
	VideoFilterBackgroundPresetN4  VideoFilterBackgroundPreset = 4
	VideoFilterBackgroundPresetN7  VideoFilterBackgroundPreset = 7
	VideoFilterBackgroundPresetN8  VideoFilterBackgroundPreset = 8
	VideoFilterBackgroundPresetN9  VideoFilterBackgroundPreset = 9
	VideoFilterBackgroundPresetN10 VideoFilterBackgroundPreset = 10
)

func (v VideoFilterBackgroundPreset) String() string {
	switch v {
	case VideoFilterBackgroundPresetN0:
		return "0"
	case VideoFilterBackgroundPresetN1:
		return "1"
	case VideoFilterBackgroundPresetN2:
		return "2"
	case VideoFilterBackgroundPresetN3:
		return "3"
	case VideoFilterBackgroundPresetN4:
		return "4"
	case VideoFilterBackgroundPresetN7:
		return "7"
	case VideoFilterBackgroundPresetN8:
		return "8"
	case VideoFilterBackgroundPresetN9:
		return "9"
	case VideoFilterBackgroundPresetN10:
		return "10"
	}
	return strconv.FormatInt(int64(v), 10)
}

type VideoFilterType int64

const (
	VideoFilterTypeBackground VideoFilterType = 0
)

func (v VideoFilterType) String() string {
	switch v {
	case VideoFilterTypeBackground:
		return "BACKGROUND"
	}
	return strconv.FormatInt(int64(v), 10)
}

type VideoQualityMode int64

const (
	VideoQualityModeAuto VideoQualityMode = 1
	VideoQualityModeFull VideoQualityMode = 2
)

func (v VideoQualityMode) String() string {
	switch v {
	case VideoQualityModeAuto:
		return "AUTO"
	case VideoQualityModeFull:
		return "FULL"
	}
	return strconv.FormatInt(int64(v), 10)
}

type VisibilityType int64

const (
	VisibilityTypeNone     VisibilityType = 0
	VisibilityTypeEveryone VisibilityType = 1
)

func (v VisibilityType) String() string {
	switch v {
	case VisibilityTypeNone:
		return "NONE"
	case VisibilityTypeEveryone:
		return "EVERYONE"
	}
	return strconv.FormatInt(int64(v), 10)
}

type VoiceChannelEffectAnimationType int64

const (
	VoiceChannelEffectAnimationTypePremium VoiceChannelEffectAnimationType = 0
	VoiceChannelEffectAnimationTypeBasic   VoiceChannelEffectAnimationType = 1
)

func (v VoiceChannelEffectAnimationType) String() string {
	switch v {
	case VoiceChannelEffectAnimationTypePremium:
		return "PREMIUM"
	case VoiceChannelEffectAnimationTypeBasic:
		return "BASIC"
	}
	return strconv.FormatInt(int64(v), 10)
}

type VoiceConnectionState string

const (
	VoiceConnectionStateDisconnected      VoiceConnectionState = "DISCONNECTED"
	VoiceConnectionStateAwaitingEndpoint  VoiceConnectionState = "AWAITING_ENDPOINT"
	VoiceConnectionStateAuthenticating    VoiceConnectionState = "AUTHENTICATING"
	VoiceConnectionStateConnecting        VoiceConnectionState = "CONNECTING"
	VoiceConnectionStateVoiceDisconnected VoiceConnectionState = "VOICE_DISCONNECTED"
	VoiceConnectionStateVoiceConnecting   VoiceConnectionState = "VOICE_CONNECTING"
	VoiceConnectionStateVoiceConnected    VoiceConnectionState = "VOICE_CONNECTED"
	VoiceConnectionStateNoRoute           VoiceConnectionState = "NO_ROUTE"
	VoiceConnectionStateIceChecking       VoiceConnectionState = "ICE_CHECKING"
	VoiceConnectionStateDtlsConnecting    VoiceConnectionState = "DTLS_CONNECTING"
)

type VoiceFlags uint64

const (
	VoiceFlagsClipsEnabled        VoiceFlags = 1 << 0
	VoiceFlagsAllowVoiceRecording VoiceFlags = 1 << 1
	VoiceFlagsAllowAnyViewerClips VoiceFlags = 1 << 2
)

func (v VoiceFlags) String() string {
	switch v {
	case VoiceFlagsClipsEnabled:
		return "CLIPS_ENABLED"
	case VoiceFlagsAllowVoiceRecording:
		return "ALLOW_VOICE_RECORDING"
	case VoiceFlagsAllowAnyViewerClips:
		return "ALLOW_ANY_VIEWER_CLIPS"
	}
	return strconv.FormatUint(uint64(v), 10)
}

type VoicePlatform int64

const (
	VoicePlatformDesktop     VoicePlatform = 0
	VoicePlatformMobile      VoicePlatform = 1
	VoicePlatformXbox        VoicePlatform = 2
	VoicePlatformPlaystation VoicePlatform = 3
)

func (v VoicePlatform) String() string {
	switch v {
	case VoicePlatformDesktop:
		return "DESKTOP"
	case VoicePlatformMobile:
		return "MOBILE"
	case VoicePlatformXbox:
		return "XBOX"
	case VoicePlatformPlaystation:
		return "PLAYSTATION"
	}
	return strconv.FormatInt(int64(v), 10)
}

type VoipPushNotificationType string

const (
	VoipPushNotificationTypeCallRing    VoipPushNotificationType = "CALL_RING"
	VoipPushNotificationTypeCallRingEnd VoipPushNotificationType = "CALL_RING_END"
)

type WebhookTypes int64

const (
	WebhookTypesIncoming        WebhookTypes = 1
	WebhookTypesChannelFollower WebhookTypes = 2
	WebhookTypesApplication     WebhookTypes = 3
)

func (v WebhookTypes) String() string {
	switch v {
	case WebhookTypesIncoming:
		return "INCOMING"
	case WebhookTypesChannelFollower:
		return "CHANNEL_FOLLOWER"
	case WebhookTypesApplication:
		return "APPLICATION"
	}
	return strconv.FormatInt(int64(v), 10)
}

type WidgetConfigStatus string

const (
	WidgetConfigStatusPublished WidgetConfigStatus = "published"
	WidgetConfigStatusDraft     WidgetConfigStatus = "draft"
)

type WidgetSurfaceValueType string

const (
	WidgetSurfaceValueTypeData                       WidgetSurfaceValueType = "data"
	WidgetSurfaceValueTypeCustomString               WidgetSurfaceValueType = "custom_string"
	WidgetSurfaceValueTypeApplicationAsset           WidgetSurfaceValueType = "application_asset"
	WidgetSurfaceValueTypeApplicationLocalizedString WidgetSurfaceValueType = "application_localized_string"
)
