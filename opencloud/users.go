package opencloud

import (
	"context"
	"net/http"

	"github.com/luckfire/goblox"
	"github.com/luckfire/goblox/auth"
)

var (
	ListAssetQuota = goblox.NewEndpoint(
		http.MethodGet,
		"/cloud/v2/users/{userId}/asset-quotas",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("asset:read")

	ListInventoryItems = goblox.NewEndpoint(
		http.MethodGet,
		"/cloud/v2/users/{userId}/inventory-items",
		auth.KindAPIKey, auth.KindOAuth2,
	) // user.inventory-item:read is only needed for private inventories

	GetUser = goblox.NewEndpoint(
		http.MethodGet,
		"/cloud/v2/users/{userId}",
		auth.KindAPIKey, auth.KindOAuth2,
	) // user.advanced:read and user.social:read add optional fields

	GenerateUserThumbnail = goblox.NewEndpoint(
		http.MethodGet,
		"/cloud/v2/users/{userId}:generateThumbnail",
		auth.KindAPIKey, auth.KindOAuth2,
	)

	CreateUserNotification = goblox.NewEndpoint(
		http.MethodPost,
		"/cloud/v2/users/{userId}/notifications",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("user.user-notification:write")

	ListUserRestrictions = goblox.NewEndpoint(
		http.MethodGet,
		"/cloud/v2/universes/{universeId}/user-restrictions",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("universe.user-restriction:read")

	ListUserRestrictionsPlace = goblox.NewEndpoint(
		http.MethodGet,
		"/cloud/v2/universes/{universeId}/places/{placeId}/user-restrictions",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("universe.user-restriction:read")

	GetUserRestriction = goblox.NewEndpoint(
		http.MethodGet,
		"/cloud/v2/universes/{universeId}/user-restrictions/{userId}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("universe.user-restriction:read")

	GetUserRestrictionPlace = goblox.NewEndpoint(
		http.MethodGet,
		"/cloud/v2/universes/{universeId}/places/{placeId}/user-restrictions/{userId}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("universe.user-restriction:read")

	UpdateUserRestriction = goblox.NewEndpoint(
		http.MethodPatch,
		"/cloud/v2/universes/{universeId}/user-restrictions/{userId}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("universe.user-restriction:write")

	UpdateUserRestrictionPlace = goblox.NewEndpoint(
		http.MethodPatch,
		"/cloud/v2/universes/{universeId}/places/{placeId}/user-restrictions/{userId}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("universe.user-restriction:write")

	ListUserRestrictionLogs = goblox.NewEndpoint(
		http.MethodGet,
		"/cloud/v2/universes/{universeId}/user-restrictions:listLogs",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("universe.user-restriction:read")

	ListUserRestrictionLogsPlace = goblox.NewEndpoint(
		http.MethodGet,
		"/cloud/v2/universes/{universeId}/places/{placeId}/user-restrictions:listLogs",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("universe.user-restriction:read")
)

// UsersService handles communication with user endpoints.
type UsersService service

type QuotaType string

const (
	QuotaTypeUnspecified                     QuotaType = "QUOTA_TYPE_UNSPECIFIED"
	QuotaTypeRateLimitUpload                 QuotaType = "RATE_LIMIT_UPLOAD"
	QuotaTypeRateLimitCreatorStoreDistribute QuotaType = "RATE_LIMIT_CREATOR_STORE_DISTRIBUTE"
)

type AssetQuotaType string

const (
	AssetQuotaTypeUnspecified               AssetQuotaType = "ASSET_TYPE_UNSPECIFIED"
	AssetQuotaTypeImage                     AssetQuotaType = "IMAGE"
	AssetQuotaTypeTShirt                    AssetQuotaType = "TSHIRT"
	AssetQuotaTypeAudio                     AssetQuotaType = "AUDIO"
	AssetQuotaTypeMesh                      AssetQuotaType = "MESH"
	AssetQuotaTypeLua                       AssetQuotaType = "LUA"
	AssetQuotaTypeHat                       AssetQuotaType = "HAT"
	AssetQuotaTypePlace                     AssetQuotaType = "PLACE"
	AssetQuotaTypeModel                     AssetQuotaType = "MODEL"
	AssetQuotaTypeShirt                     AssetQuotaType = "SHIRT"
	AssetQuotaTypePants                     AssetQuotaType = "PANTS"
	AssetQuotaTypeDecal                     AssetQuotaType = "DECAL"
	AssetQuotaTypeHead                      AssetQuotaType = "HEAD"
	AssetQuotaTypeFace                      AssetQuotaType = "FACE"
	AssetQuotaTypeGear                      AssetQuotaType = "GEAR"
	AssetQuotaTypeAnimation                 AssetQuotaType = "ANIMATION"
	AssetQuotaTypeTorso                     AssetQuotaType = "TORSO"
	AssetQuotaTypeRightArm                  AssetQuotaType = "RIGHT_ARM"
	AssetQuotaTypeLeftArm                   AssetQuotaType = "LEFT_ARM"
	AssetQuotaTypeLeftLeg                   AssetQuotaType = "LEFT_LEG"
	AssetQuotaTypeRightLeg                  AssetQuotaType = "RIGHT_LEG"
	AssetQuotaTypeYouTubeVideo              AssetQuotaType = "YOUTUBE_VIDEO"
	AssetQuotaTypeApp                       AssetQuotaType = "APP"
	AssetQuotaTypeCode                      AssetQuotaType = "CODE"
	AssetQuotaTypePlugin                    AssetQuotaType = "PLUGIN"
	AssetQuotaTypeSolidModel                AssetQuotaType = "SOLID_MODEL"
	AssetQuotaTypeMeshPart                  AssetQuotaType = "MESH_PART"
	AssetQuotaTypeHairAccessory             AssetQuotaType = "HAIR_ACCESSORY"
	AssetQuotaTypeFaceAccessory             AssetQuotaType = "FACE_ACCESSORY"
	AssetQuotaTypeNeckAccessory             AssetQuotaType = "NECK_ACCESSORY"
	AssetQuotaTypeShoulderAccessory         AssetQuotaType = "SHOULDER_ACCESSORY"
	AssetQuotaTypeFrontAccessory            AssetQuotaType = "FRONT_ACCESSORY"
	AssetQuotaTypeBackAccessory             AssetQuotaType = "BACK_ACCESSORY"
	AssetQuotaTypeWaistAccessory            AssetQuotaType = "WAIST_ACCESSORY"
	AssetQuotaTypeClimbAnimation            AssetQuotaType = "CLIMB_ANIMATION"
	AssetQuotaTypeDeathAnimation            AssetQuotaType = "DEATH_ANIMATION"
	AssetQuotaTypeFallAnimation             AssetQuotaType = "FALL_ANIMATION"
	AssetQuotaTypeIdleAnimation             AssetQuotaType = "IDLE_ANIMATION"
	AssetQuotaTypeJumpAnimation             AssetQuotaType = "JUMP_ANIMATION"
	AssetQuotaTypeRunAnimation              AssetQuotaType = "RUN_ANIMATION"
	AssetQuotaTypeSwimAnimation             AssetQuotaType = "SWIM_ANIMATION"
	AssetQuotaTypeWalkAnimation             AssetQuotaType = "WALK_ANIMATION"
	AssetQuotaTypePoseAnimation             AssetQuotaType = "POSE_ANIMATION"
	AssetQuotaTypeLocalizationTableManifest AssetQuotaType = "LOCALIZATION_TABLE_MANIFEST"
	AssetQuotaTypeEmoteAnimation            AssetQuotaType = "EMOTE_ANIMATION"
	AssetQuotaTypeVideo                     AssetQuotaType = "VIDEO"
	AssetQuotaTypeTexturePack               AssetQuotaType = "TEXTURE_PACK"
	AssetQuotaTypeTShirtAccessory           AssetQuotaType = "TSHIRT_ACCESSORY"
	AssetQuotaTypeShirtAccessory            AssetQuotaType = "SHIRT_ACCESSORY"
	AssetQuotaTypePantsAccessory            AssetQuotaType = "PANTS_ACCESSORY"
	AssetQuotaTypeJacketAccessory           AssetQuotaType = "JACKET_ACCESSORY"
	AssetQuotaTypeSweaterAccessory          AssetQuotaType = "SWEATER_ACCESSORY"
	AssetQuotaTypeShortsAccessory           AssetQuotaType = "SHORTS_ACCESSORY"
	AssetQuotaTypeLeftShoeAccessory         AssetQuotaType = "LEFT_SHOE_ACCESSORY"
	AssetQuotaTypeRightShoeAccessory        AssetQuotaType = "RIGHT_SHOE_ACCESSORY"
	AssetQuotaTypeDressSkirtAccessory       AssetQuotaType = "DRESS_SKIRT_ACCESSORY"
	AssetQuotaTypeFontFamily                AssetQuotaType = "FONT_FAMILY"
	AssetQuotaTypeFontFace                  AssetQuotaType = "FONT_FACE"
	AssetQuotaTypeMeshHiddenSurfaceRemoval  AssetQuotaType = "MESH_HIDDEN_SURFACE_REMOVAL"
	AssetQuotaTypeEyebrowAccessory          AssetQuotaType = "EYEBROW_ACCESSORY"
	AssetQuotaTypeEyelashAccessory          AssetQuotaType = "EYELASH_ACCESSORY"
	AssetQuotaTypeMoodAnimation             AssetQuotaType = "MOOD_ANIMATION"
	AssetQuotaTypeDynamicHead               AssetQuotaType = "DYNAMIC_HEAD"
	AssetQuotaTypeCodeSnippet               AssetQuotaType = "CODE_SNIPPET"
	AssetQuotaTypeAdsVideo                  AssetQuotaType = "ADS_VIDEO"
)

type QuotaPeriod string

const (
	QuotaPeriodUnspecified QuotaPeriod = "PERIOD_UNSPECIFIED"
	QuotaPeriodMonth       QuotaPeriod = "MONTH"
	QuotaPeriodDay         QuotaPeriod = "DAY"
)

type AssetQuota struct {
	Path                string         `json:"path"`
	QuotaType           QuotaType      `json:"quotaType"`
	AssetType           AssetQuotaType `json:"assetType"`
	Usage               int            `json:"usage"`
	Capacity            int            `json:"capacity"`
	Period              QuotaPeriod    `json:"period"`
	UsageResetTimestamp string         `json:"usageResetTime"`
}

type AssetQuotas struct {
	AssetQuotas   []AssetQuota `json:"assetQuotas"`
	NextPageToken string       `json:"nextPageToken"`
}

// ListAssetQuota will return information regarding the user's asset quota uploads.
//
// Required scopes: asset:read
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/AssetQuota#Cloud_ListAssetQuotas
//
// [GET] /cloud/v2/users/{user_id}/asset-quotas
func (s *UsersService) ListAssetQuota(ctx context.Context, userId string, opts *OptionsWithFilter, reqOpts ...goblox.RequestOpt) (*AssetQuotas, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	assetQuotas := new(AssetQuotas)
	resp, err := s.client.Call(ctx, ListAssetQuota.Compile(q, userId), nil, assetQuotas, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return assetQuotas, resp, nil
}

type InventoryItemAssetType string

const (
	InventoryItemAssetTypeUnspecified InventoryItemAssetType = "INVENTORY_ITEM_ASSET_TYPE_UNSPECIFIED"
	ClassicTShirt                     InventoryItemAssetType = "CLASSIC_TSHIRT"
	Audio                             InventoryItemAssetType = "AUDIO"
	Hat                               InventoryItemAssetType = "HAT"
	Model                             InventoryItemAssetType = "MODEL"
	ClassicShirt                      InventoryItemAssetType = "CLASSIC_SHIRT"
	ClassicPants                      InventoryItemAssetType = "CLASSIC_PANTS"
	Decal                             InventoryItemAssetType = "DECAL"
	ClassicHead                       InventoryItemAssetType = "CLASSIC_HEAD"
	Face                              InventoryItemAssetType = "FACE"
	Gear                              InventoryItemAssetType = "GEAR"
	Animation                         InventoryItemAssetType = "ANIMATION"
	Torso                             InventoryItemAssetType = "TORSO"
	RightArm                          InventoryItemAssetType = "RIGHT_ARM"
	LeftArm                           InventoryItemAssetType = "LEFT_ARM"
	LeftLeg                           InventoryItemAssetType = "LEFT_LEG"
	RightLeg                          InventoryItemAssetType = "RIGHT_LEG"
	Package                           InventoryItemAssetType = "PACKAGE"
	Plugin                            InventoryItemAssetType = "PLUGIN"
	MeshPart                          InventoryItemAssetType = "MESH_PART"
	HairAccessory                     InventoryItemAssetType = "HAIR_ACCESSORY"
	FaceAccessory                     InventoryItemAssetType = "FACE_ACCESSORY"
	NeckAccessory                     InventoryItemAssetType = "NECK_ACCESSORY"
	ShoulderAccessory                 InventoryItemAssetType = "SHOULDER_ACCESSORY"
	FrontAccessory                    InventoryItemAssetType = "FRONT_ACCESSORY"
	BackAccessory                     InventoryItemAssetType = "BACK_ACCESSORY"
	WaistAccessory                    InventoryItemAssetType = "WAIST_ACCESSORY"
	ClimbAnimation                    InventoryItemAssetType = "CLIMB_ANIMATION"
	DeathAnimation                    InventoryItemAssetType = "DEATH_ANIMATION"
	FallAnimation                     InventoryItemAssetType = "FALL_ANIMATION"
	IdleAnimation                     InventoryItemAssetType = "IDLE_ANIMATION"
	JumpAnimation                     InventoryItemAssetType = "JUMP_ANIMATION"
	RunAnimation                      InventoryItemAssetType = "RUN_ANIMATION"
	SwimAnimation                     InventoryItemAssetType = "SWIM_ANIMATION"
	WalkAnimation                     InventoryItemAssetType = "WALK_ANIMATION"
	PoseAnimation                     InventoryItemAssetType = "POSE_ANIMATION"
	EmoteAnimation                    InventoryItemAssetType = "EMOTE_ANIMATION"
	Video                             InventoryItemAssetType = "VIDEO"
	TShirtAccessory                   InventoryItemAssetType = "TSHIRT_ACCESSORY"
	ShirtAccessory                    InventoryItemAssetType = "SHIRT_ACCESSORY"
	PantsAccessory                    InventoryItemAssetType = "PANTS_ACCESSORY"
	JacketAccessory                   InventoryItemAssetType = "JACKET_ACCESSORY"
	SweaterAccessory                  InventoryItemAssetType = "SWEATER_ACCESSORY"
	ShortsAccessory                   InventoryItemAssetType = "SHORTS_ACCESSORY"
	LeftShoeAccessory                 InventoryItemAssetType = "LEFT_SHOE_ACCESSORY"
	RightShoeAccessory                InventoryItemAssetType = "RIGHT_SHOE_ACCESSORY"
	DressSkirtAccessory               InventoryItemAssetType = "DRESS_SKIRT_ACCESSORY"
	EyebrowAccessory                  InventoryItemAssetType = "EYEBROW_ACCESSORY"
	EyelashAccessory                  InventoryItemAssetType = "EYELASH_ACCESSORY"
	MoodAnimation                     InventoryItemAssetType = "MOOD_ANIMATION"
	DynamicHead                       InventoryItemAssetType = "DYNAMIC_HEAD"
	CreatedPlace                      InventoryItemAssetType = "CREATED_PLACE"
	PurchasedPlace                    InventoryItemAssetType = "PURCHASED_PLACE"
)

type InventoryItemInstanceState string

const (
	InventoryItemInstanceStateUnspecified InventoryItemInstanceState = "COLLECTIBLE_ITEM_INSTANCE_STATE_UNSPECIFIED"
	InventoryItemInstanceStateAvailable   InventoryItemInstanceState = "AVAILABLE"
	InventoryItemInstanceStateHold        InventoryItemInstanceState = "HOLD"
)

type InventoryItemCollectibleDetails struct {
	ItemID        string                     `json:"itemId"`
	InstanceID    string                     `json:"instanceId"`
	InstanceState InventoryItemInstanceState `json:"instanceState"`
	SerialNumber  int                        `json:"serialNumber"`
}

type InventoryItemAssetDetails struct {
	AssetID                string                          `json:"assetId"`
	InventoryItemAssetType InventoryItemAssetType          `json:"inventoryItemAssetType"`
	InstanceID             string                          `json:"instanceId"`
	CollectibleDetails     InventoryItemCollectibleDetails `json:"collectibleDetails"`
}

type InventoryItemBadgeDetils struct {
	BadgeID string `json:"badgeId"`
}

type InventoryItemGamePassDetails struct {
	GamePassID string `json:"gamePassId"`
}

type InventoryItemPrivateServerDetails struct {
	PrivateServerID string `json:"privateServerId"`
}

type InventoryItem struct {
	Path                 string                             `json:"path"`
	AssetDetails         *InventoryItemAssetDetails         `json:"assetDetails"`
	BadgeDetails         *InventoryItemBadgeDetils          `json:"badgeDetails"`
	GamePassDetails      *InventoryItemGamePassDetails      `json:"gamePassDetails"`
	PrivateServerDetails *InventoryItemPrivateServerDetails `json:"privateServerDetails"`
	AddTime              string                             `json:"addTime"`
}

type InventoryItems struct {
	InventoryItems []InventoryItem `json:"inventoryItems"`
	NextPageToken  string          `json:"nextPageToken"`
}

// ListInventoryItems will fetch inventory items for a specificed user.
//
// Required scopes:
//
// - none (Inventory Public)
//
// - user.inventory-item:read (Inventory Private)
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/InventoryItem#Cloud_ListInventoryItems
//
// [GET] /cloud/v2/users/{user_id}/inventory-items
func (s *UsersService) ListInventoryItems(ctx context.Context, userId string, opts *OptionsWithFilter, reqOpts ...goblox.RequestOpt) (*InventoryItems, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	inventoryItems := new(InventoryItems)
	resp, err := s.client.Call(ctx, ListInventoryItems.Compile(q, userId), nil, inventoryItems, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return inventoryItems, resp, nil
}

type UserSocialLinksVisibility string

const (
	UserSocialLinksVisibilityUnspecified                  UserSocialLinksVisibility = "SOCIAL_NETWORK_VISIBILITY_UNSPECIFIED"
	UserSocialLinksVisibilityNoOne                        UserSocialLinksVisibility = "NO_ONE"
	UserSocialLinksVisibilityFriends                      UserSocialLinksVisibility = "FRIENDS"
	UserSocialLinksVisibilityFriendsAndFollowing          UserSocialLinksVisibility = "FRIENDS_AND_FOLLOWING"
	UserSocialLinksVisibilityFriendsFollowingAndFollowers UserSocialLinksVisibility = "FRIENDS_FOLLOWING_AND_FOLLOWERS"
	UserSocialLinksVisibilityEveryone                     UserSocialLinksVisibility = "EVERYONE"
)

type UserSocialLinks struct {
	Facebook   string                    `json:"facebook"`
	Twitter    string                    `json:"twitter"`
	YouTube    string                    `json:"youtube"`
	Twitch     string                    `json:"twitch"`
	Guilded    string                    `json:"guilded"`
	Visibility UserSocialLinksVisibility `json:"visibility"`
}

type User struct {
	Path                  string          `json:"path"`
	CreateTime            string          `json:"createTime"`
	ID                    string          `json:"id"`
	DisplayName           string          `json:"displayName"`
	About                 string          `json:"about"`
	Locale                string          `json:"locale"`
	Premium               bool            `json:"premium"`
	IDVerified            bool            `json:"idVerified"`
	SocialNetworkProfiles UserSocialLinks `json:"socialNetworkProfiles"`
}

// GetUser will fetch information on a specificed user.
//
// Required scopes:
//
// - none (Public Information)
//
// - user.advanced:read (Verification Status)
//
// - user.social:read (Social Information)
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/User#Cloud_GetUser
//
// [GET] /cloud/v2/users/{user_id}
func (s *UsersService) GetUser(ctx context.Context, userId string, reqOpts ...goblox.RequestOpt) (*User, *Response, error) {
	user := new(User)
	resp, err := s.client.Call(ctx, GetUser.Compile(nil, userId), nil, user, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return user, resp, nil
}

type UserThumbnail struct {
	ImageURI string `json:"imageUri"`
}

type UserThumbnailSize int

const (
	UserThumbnailSize48  UserThumbnailSize = 48
	UserThumbnailSize50  UserThumbnailSize = 50
	UserThumbnailSize60  UserThumbnailSize = 60
	UserThumbnailSize75  UserThumbnailSize = 75
	UserThumbnailSize100 UserThumbnailSize = 100
	UserThumbnailSize110 UserThumbnailSize = 110
	UserThumbnailSize150 UserThumbnailSize = 150
	UserThumbnailSize180 UserThumbnailSize = 180
	UserThumbnailSize352 UserThumbnailSize = 352
	UserThumbnailSize420 UserThumbnailSize = 420
	UserThumbnailSize720 UserThumbnailSize = 720
)

type UserThumbnailFormat string

const (
	UserThumbnailFormatUnspecified UserThumbnailFormat = "FORMAT_UNSPECIFIED"
	UserThumbnailFormatJPEG        UserThumbnailFormat = "JPEG"
	UserThumbnailFormatPNG         UserThumbnailFormat = "PNG"
)

type UserThumbnailShape string

const (
	UserThumbnailShapeUnspecified UserThumbnailShape = "SHAPE_UNSPECIFIED"
	UserThumbnailShapeRound       UserThumbnailShape = "ROUND"
	UserThumbnailShapeSquare      UserThumbnailShape = "SQUARE"
)

type UserThumbnailOptions struct {
	Size   UserThumbnailSize   `url:"size,omitempty"`
	Format UserThumbnailFormat `url:"format,omitempty"`
	Shape  UserThumbnailShape  `url:"shape,omitempty"`
}

// GenerateUserThumbnail will generate and return a url for a user's avatar.
//
// Required scopes: none
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/User#Cloud_GenerateUserThumbnail
//
// [GET] /cloud/v2/users/{user_id}:generateThumbnail
func (s *UsersService) GenerateUserThumbnail(ctx context.Context, userId string, opts *UserThumbnailOptions, reqOpts ...goblox.RequestOpt) (*UserThumbnail, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	userThumbnail := new(UserThumbnail)
	resp, err := s.client.Call(ctx, GenerateUserThumbnail.Compile(q, userId), nil, userThumbnail, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return userThumbnail, resp, nil
}

type UserNotificationSource struct {
	Universe string `json:"universe"`
}

type UserNotificationType string

const (
	UserNotificationTypeUnspecified UserNotificationType = "TYPE_UNSPECIFIED"
	UserNotificationTypeMoment      UserNotificationType = "MOMENT"
)

type UserNotificationJoinExperience struct {
	LaunchData string `json:"launchData"`
}

type UserNotificationAnalyticsData struct {
	Category string `json:"category"`
}

type UserNotificationPayload struct {
	Type           UserNotificationType           `json:"type"`
	MessageID      string                         `json:"messageId"`
	Parameters     map[string]any                 `json:"parameters"`
	JoinExperience UserNotificationJoinExperience `json:"joinExperience"`
	AnalyticsData  UserNotificationAnalyticsData  `json:"analyticsData"`
}

type UserNotification struct {
	Path    string                  `json:"path"`
	ID      string                  `json:"id"`
	Source  UserNotificationSource  `json:"source"`
	Payload UserNotificationPayload `json:"payload"`
}

type UserNotificationCreate struct {
	Source  *UserNotificationSource  `json:"source,omitempty"`
	Payload *UserNotificationPayload `json:"payload,omitempty"`
}

// CreateUserNotification will send a notification to a user.
//
// Required scopes: user.user-notification:write
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/UserNotification#Cloud_CreateUserNotification
//
// [POST] /cloud/v2/users/{user_id}/notifications
func (s *UsersService) CreateUserNotification(ctx context.Context, userId string, data UserNotificationCreate, reqOpts ...goblox.RequestOpt) (*UserNotification, *Response, error) {
	userNotification := new(UserNotification)
	resp, err := s.client.Call(ctx, CreateUserNotification.Compile(nil, userId), data, userNotification, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return userNotification, resp, nil
}

type GameJoinRestriction struct {
	Active             bool    `json:"active"`
	StartTime          string  `json:"startTime"`
	Duration           *string `json:"duration,omitempty"`
	PrivateReason      string  `json:"privateReason"`
	DisplayReason      string  `json:"displayReason"`
	ExcludeAltAccounts bool    `json:"excludeAltAccounts"`
	Inherited          bool    `json:"inherited"`
}

type UserRestriction struct {
	Path                string              `json:"path"`
	UpdateTime          string              `json:"updateTime"`
	User                string              `json:"user"`
	GameJoinRestriction GameJoinRestriction `json:"gameJoinRestriction"`
}

type UserRestrictions struct {
	UserRestrictions []UserRestriction `json:"userRestrictions"`
	NextPageToken    string            `json:"nextPageToken"`
}

// ListUserRestrictions will list restrictions for users thave have been restricted for the specified universe or place.
//
// Required scopes: universe.user-restriction:read
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/UserRestriction#Cloud_ListUserRestrictions
//
// [GET] /cloud/v2/universes/{universe_id}/user-restrictions
//
// [GET] /cloud/v2/universes/{universe_id}/places/{place_id}/user-restrictions
func (s *UsersService) ListUserRestrictions(ctx context.Context, universeId string, placeId *string, opts *Options, reqOpts ...goblox.RequestOpt) (*UserRestrictions, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	e := ListUserRestrictions.Compile(q, universeId)
	if placeId != nil {
		e = ListUserRestrictionsPlace.Compile(q, universeId, *placeId)
	}

	userRestrictions := new(UserRestrictions)
	resp, err := s.client.Call(ctx, e, nil, userRestrictions, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return userRestrictions, resp, nil
}

// GetUserRestriction will get the user restriction for the specified user under the specified universe or place.
//
// Required scopes: universe.user-restriction:read
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/UserRestriction#Cloud_GetUserRestriction
//
// [GET] /cloud/v2/universes/{universe_id}/user-restrictions/{user_id}
//
// [GET] /cloud/v2/universes/{universe_id}/places/{place_id}/user-restrictions/{user_id}
func (s *UsersService) GetUserRestriction(ctx context.Context, universeId string, placeId *string, userId string, reqOpts ...goblox.RequestOpt) (*UserRestriction, *Response, error) {
	e := GetUserRestriction.Compile(nil, universeId, userId)
	if placeId != nil {
		e = GetUserRestrictionPlace.Compile(nil, universeId, *placeId, userId)
	}

	userRestriction := new(UserRestriction)
	resp, err := s.client.Call(ctx, e, nil, userRestriction, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return userRestriction, resp, nil
}

type GameJoinRestrictionUpdate struct {
	Active             *bool   `json:"active,omitempty"`
	Duration           *string `json:"duration,omitempty"`
	PrivateReason      *string `json:"privateReason,omitempty"`
	DisplayReason      *string `json:"displayReason,omitempty"`
	ExcludeAltAccounts *bool   `json:"excludeAltAccounts,omitempty"`
}

type UserRestrictionUpdate struct {
	GameJoinRestriction *GameJoinRestrictionUpdate `json:"gameJoinRestriction,omitempty"`
}

type UserRestrictionUpdateOptions struct {
	UpdateMask           *string `url:"updateMask,omitempty"`
	IdempotencyKey       *string `url:"idempotencyKey.key,omitempty"`
	IdempotencyFirstSent *string `url:"idempotencyKey.firstSent,omitempty"`
}

// UpdateUserRestriction will update the current user restriction for the specified user under the specified universe or place.
//
// Required scopes: universe.user-restriction:write
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/UserRestriction#Cloud_UpdateUserRestriction
//
// [PATCH] /cloud/v2/universes/{universe_id}/user-restrictions/{user_id}
//
// [PATCH] /cloud/v2/universes/{universe_id}/places/{place_id}/user-restrictions/{user_id}
func (s *UsersService) UpdateUserRestriction(ctx context.Context, universeId string, placeId *string, userId string, data UserRestrictionUpdate, opts *UserRestrictionUpdateOptions, reqOpts ...goblox.RequestOpt) (*UserRestriction, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	e := UpdateUserRestriction.Compile(q, universeId, userId)
	if placeId != nil {
		e = UpdateUserRestrictionPlace.Compile(q, universeId, *placeId, userId)
	}

	userRestriction := new(UserRestriction)
	resp, err := s.client.Call(ctx, e, data, userRestriction, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return userRestriction, resp, nil
}

type UserRestrictionModerator struct {
	RobloxUser string `json:"robloxUser"`
}

type UserRestrictionLog struct {
	User               string                   `json:"user"`
	Place              string                   `json:"place"`
	Moderator          UserRestrictionModerator `json:"moderator"`
	CreateTime         string                   `json:"createTime"`
	Active             bool                     `json:"active"`
	StartTime          string                   `json:"startTime"`
	Duration           string                   `json:"duration"`
	PrivateReason      string                   `json:"privateReason"`
	DisplayReason      string                   `json:"displayReason"`
	RestrictionType    GameJoinRestriction      `json:"restrictionType"`
	ExcludeAltAccounts bool                     `json:"excludeAltAccounts"`
}

type UserRestrictionLogs struct {
	Logs          []UserRestrictionLog `json:"logs"`
	NextPageToken string               `json:"nextPageToken"`
}

// ListUserRestrictionLogs will list logs of user restrictions for the specified universe or place.
//
// Required scopes: universe.user-restriction:read
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/en-us/cloud/reference/UserRestriction#Cloud_ListUserRestrictionLogs
//
// [GET] /cloud/v2/universes/{universe_id}/user-restrictions:listLogs
//
// [GET] /cloud/v2/universes/{universe_id}/places/{place_id}/user-restrictions:listLogs
func (s *UsersService) ListUserRestrictionLogs(ctx context.Context, universeId string, placeId *string, opts *OptionsWithFilter, reqOpts ...goblox.RequestOpt) (*UserRestrictionLogs, *Response, error) {
	q, err := queryValues(opts)
	if err != nil {
		return nil, nil, err
	}

	e := ListUserRestrictionLogs.Compile(q, universeId)
	if placeId != nil {
		e = ListUserRestrictionLogsPlace.Compile(q, universeId, *placeId)
	}

	restrictionLogs := new(UserRestrictionLogs)
	resp, err := s.client.Call(ctx, e, nil, restrictionLogs, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return restrictionLogs, resp, nil
}

var (
	GetUserThumbnailGenerationOperation = goblox.NewEndpoint(
		http.MethodGet,
		"/cloud/v2/users/{userId}/operations/{operationId}",
		auth.KindAPIKey, auth.KindOAuth2,
	)

	UnfollowUniverse = goblox.NewEndpoint(
		http.MethodDelete,
		"/legacy-followings/v1/users/{userId}/universes/{universeId}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-universe.following:write")

	FollowUniverse = goblox.NewEndpoint(
		http.MethodPost,
		"/legacy-followings/v1/users/{userId}/universes/{universeId}",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-universe.following:write")

	GetUniverseFollowStatus = goblox.NewEndpoint(
		http.MethodGet,
		"/legacy-followings/v1/users/{userId}/universes/{universeId}/status",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-universe.following:read")

	ListUserFollowedUniverses = goblox.NewEndpoint(
		http.MethodGet,
		"/legacy-followings/v2/users/{userId}/universes",
		auth.KindAPIKey, auth.KindOAuth2,
	).WithScopes("legacy-universe.following:read")
)

// GetUserThumbnailGenerationOperation will get the status and result of a user thumbnail generation operation.
//
// Required scopes: none
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/avatars#Cloud_GetUserThumbnailGenerationOperation
//
// [GET] /cloud/v2/users/{user_id}/operations/{operation_id}
func (s *UsersService) GetUserThumbnailGenerationOperation(ctx context.Context, userId, operationId string, reqOpts ...goblox.RequestOpt) (*Operation, *Response, error) {
	operation := new(Operation)
	resp, err := s.client.Call(ctx, GetUserThumbnailGenerationOperation.Compile(nil, userId, operationId), nil, operation, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return operation, resp, nil
}

type UserFollowingUniverseResponse struct {
	UniverseID int64 `json:"universeId"`
	UserID     int64 `json:"userId"`
}

// UnfollowUniverse will delete the following between a user with userId and universe with universeId.
//
// Required scopes: legacy-universe.following:write
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/interactions#delete_legacy_followings_v1_users__userId__universes__universeId_
//
// [DELETE] /legacy-followings/v1/users/{userId}/universes/{universeId}
func (s *UsersService) UnfollowUniverse(ctx context.Context, userId, universeId string, reqOpts ...goblox.RequestOpt) (*UserFollowingUniverseResponse, *Response, error) {
	userFollowingUniverseResponse := new(UserFollowingUniverseResponse)
	resp, err := s.client.Call(ctx, UnfollowUniverse.Compile(nil, userId, universeId), nil, userFollowingUniverseResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return userFollowingUniverseResponse, resp, nil
}

// FollowUniverse will create the following between a user with userId and universe with universeId.
//
// Required scopes: legacy-universe.following:write
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/interactions#post_legacy_followings_v1_users__userId__universes__universeId_
//
// [POST] /legacy-followings/v1/users/{userId}/universes/{universeId}
func (s *UsersService) FollowUniverse(ctx context.Context, userId, universeId string, reqOpts ...goblox.RequestOpt) (*UserFollowingUniverseResponse, *Response, error) {
	userFollowingUniverseResponse := new(UserFollowingUniverseResponse)
	resp, err := s.client.Call(ctx, FollowUniverse.Compile(nil, userId, universeId), nil, userFollowingUniverseResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return userFollowingUniverseResponse, resp, nil
}

type UserFollowingUniverseStatusResponse struct {
	UniverseID           int64 `json:"UniverseId"`
	UserID               int64 `json:"UserId"`
	CanFollow            bool  `json:"CanFollow"`
	IsFollowing          bool  `json:"IsFollowing"`
	FollowingCountByType int   `json:"FollowingCountByType"`
	FollowingLimitByType int   `json:"FollowingLimitByType"`
}

// GetUniverseFollowStatus will get the status of a following relationship between a user and a universe.
//
// Required scopes: legacy-universe.following:read
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/interactions#get_legacy_followings_v1_users__userId__universes__universeId__status
//
// [GET] /legacy-followings/v1/users/{userId}/universes/{universeId}/status
func (s *UsersService) GetUniverseFollowStatus(ctx context.Context, userId, universeId string, reqOpts ...goblox.RequestOpt) (*UserFollowingUniverseStatusResponse, *Response, error) {
	userFollowingUniverseStatusResponse := new(UserFollowingUniverseStatusResponse)
	resp, err := s.client.Call(ctx, GetUniverseFollowStatus.Compile(nil, userId, universeId), nil, userFollowingUniverseStatusResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return userFollowingUniverseStatusResponse, resp, nil
}

type FollowsByTypeResponseFollowerType int

type FollowsByTypeResponseSourceType int

type FollowsByTypeResponse struct {
	FollowerType    FollowsByTypeResponseFollowerType `json:"followerType"`
	FollowerID      int64                             `json:"followerId"`
	SourceType      FollowsByTypeResponseSourceType   `json:"sourceType"`
	FollowedSources map[string]string                 `json:"followedSources"`
}

// ListUserFollowedUniverses will get all universes followed by a user.
//
// Required scopes: legacy-universe.following:read
//
// Roblox Opencloud API Docs: https://create.roblox.com/docs/cloud/reference/features/interactions#get_legacy_followings_v2_users__userId__universes
//
// [GET] /legacy-followings/v2/users/{userId}/universes
func (s *UsersService) ListUserFollowedUniverses(ctx context.Context, userId string, reqOpts ...goblox.RequestOpt) (*FollowsByTypeResponse, *Response, error) {
	followsByTypeResponse := new(FollowsByTypeResponse)
	resp, err := s.client.Call(ctx, ListUserFollowedUniverses.Compile(nil, userId), nil, followsByTypeResponse, reqOpts...)
	if err != nil {
		return nil, resp, err
	}

	return followsByTypeResponse, resp, nil
}
