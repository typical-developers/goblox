package opencloud

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
)

// -- General
type ExperienceSocialLink struct {
	Title string `json:"title"`
	URI   string `json:"uri"`
}

// -- Univese and Places

// https://create.roblox.com/docs/en-us/cloud/reference/Instance#Instance
type Instance struct {
	Path           string `json:"path"`
	HasChildren    bool   `json:"hasChildren"`
	EngineInstance struct {
		ID      string `json:"Id"`
		Parent  string `json:"Parent"`
		Name    string `json:"Name"`
		Details struct {
			Folder      struct{} `json:"Folder"`
			LocalScript struct {
				Enabled bool `json:"Enabled"`
				// Possible values:
				// "Legacy" | "Server" | "Client" | "Plugin"
				RunContext string `json:"RunContext"`
				Source     string `json:"Source"`
			} `json:"LocalScript"`
			MolduleScript struct {
				Source string `json:"Source"`
			} `json:"ModuleScript"`
			Script struct {
				Enabled bool `json:"Enabled"`
				// Possible values:
				// "Legacy" | "Server" | "Client" | "Plugin"
				RunContext string `json:"RunContext"`
				Source     string `json:"Source"`
			} `json:"Script"`
		} `json:"Details"`
	} `json:"engineInstance"`
}

type ListInstanceQuery struct {
	MaxPageSize *int    `json:"maxPageSize"`
	PageToken   *string `json:"pageToken"`
}

func (query *ListInstanceQuery) ConvertToStringMap() map[string]string {
	result := make(map[string]string)

	if query.MaxPageSize != nil {
		result["maxPageSize"] = strconv.Itoa(*query.MaxPageSize)
	}

	if query.PageToken != nil {
		result["pageToken"] = *query.PageToken
	}

	return result
}

type ListInstance struct {
	Instances     []Instance `json:"instances"`
	NextPageToken string     `json:"nextPageToken"`
}

// https://create.roblox.com/docs/en-us/cloud/reference/Place
type Place struct {
	Path        string `json:"path"`
	CreateTime  string `json:"createTime"`
	UpdateTime  string `json:"updateTime"`
	DisplayName string `json:"displayName"`
	Description string `json:"description"`
	ServerSize  int    `json:"serverSize"`
}

type PlaceUpdate struct {
	DisplayName *string `json:"displayName,omitempty"`
	Description *string `json:"description,omitempty"`
	ServerSize  *int    `json:"serverSize,omitempty"`
}

func (data *PlaceUpdate) SetName(name string) {
	data.DisplayName = &name
}

func (data *PlaceUpdate) SetDescription(name string) {
	data.Description = &name
}

func (data *PlaceUpdate) SetServerSize(size int) {
	data.ServerSize = &size
}

func (data *PlaceUpdate) Validate() error {
	if data.DisplayName == nil && data.Description == nil && data.ServerSize == nil {
		return fmt.Errorf("PlaceUpdate: At least one field must be set")
	}

	if data.ServerSize != nil && (*data.ServerSize < 1 || *data.ServerSize > 200) {
		return fmt.Errorf("PlaceUpdate: ServerSize must be between 1 and 200")
	}

	return nil
}

type Universe struct {
	Path        string  `json:"path"`
	CreateTime  string  `json:"createTime"`
	UpdateTime  string  `json:"updateTime"`
	DisplayName string  `json:"displayName"`
	Description string  `json:"description"`
	User        *string `json:"user"`
	Group       *string `json:"group"`
	// Possible values:
	// "VISBILITY_UNSPECIFIED" | "PUBLIC" | "PRIVATE"
	Visibility         string               `json:"visibility"`
	FacebookSocialLink ExperienceSocialLink `json:"facebookSocialLink"`
	TwitterSocialLink  ExperienceSocialLink `json:"twitterSocialLink"`
	YoutubeSocialLink  ExperienceSocialLink `json:"youtubeSocialLink"`
	TwitchSocialLink   ExperienceSocialLink `json:"twitchSocialLink"`
	DiscordSocialLink  ExperienceSocialLink `json:"discordSocialLink"`
	VoiceChatEnabled   bool                 `json:"voiceChatEnabled"`
	// Possible values:
	// "AGE_RATING_UNSPECIFIED" | "AGE_RATING_ALL" | "AGE_RATING_9_PLUS" | "AGE_RATING_13_PLUS" | "AGE_RATING_17_PLUS"
	AgeRating               string `json:"ageRating"`
	PrivateServerPriceRobux int    `json:"privateServerPriceRobux"`
	DesktopEnabled          bool   `json:"desktopEnabled"`
	MobileEnabled           bool   `json:"mobileEnabled"`
	TabletEnabled           bool   `json:"tabletEnabled"`
	ConsoleEnabled          bool   `json:"consoleEnabled"`
	VREnabled               bool   `json:"vrEnabled"`
}

type UniverseUpdate struct {
	FacebookSocialLink      ExperienceSocialLink `json:"facebookSocialLink"`
	TwitterSocialLink       ExperienceSocialLink `json:"twitterSocialLink"`
	YoutubeSocialLink       ExperienceSocialLink `json:"youtubeSocialLink"`
	TwitchSocialLink        ExperienceSocialLink `json:"twitchSocialLink"`
	DiscordSocialLink       ExperienceSocialLink `json:"discordSocialLink"`
	VoiceChatEnabled        bool                 `json:"voiceChatEnabled"`
	PrivateServerPriceRobux int                  `json:"privateServerPriceRobux"`
	DesktopEnabled          bool                 `json:"desktopEnabled"`
	MobileEnabled           bool                 `json:"mobileEnabled"`
	TabletEnabled           bool                 `json:"tabletEnabled"`
	ConsoleEnabled          bool                 `json:"consoleEnabled"`
	VREnabled               bool                 `json:"vrEnabled"`
}

type UniverseMessage struct {
	Topic   string `json:"topic"`
	Message string `json:"message"`
}

func (m *UniverseMessage) Validate(topic string) error {
	if len(topic) < 1 || len(topic) > 80 {
		return fmt.Errorf("UniverseMessage: Topic must be between 1 and 80 characters.")
	}

	if len([]byte(m.Message)) > 1024 {
		return errors.New("UniverseMessage: Message can't be more than 1kB (1024 bytes).")
	}

	return nil
}

// -- Users and Groups

// https://create.roblox.com/docs/en-us/cloud/reference/User#User
type User struct {
	Path                  string `json:"path"`
	CreateTime            string `json:"createTime"`
	ID                    string `json:"id"`
	Name                  string `json:"name"`
	DisplayName           string `json:"displayName"`
	About                 string `json:"about"`
	Locale                string `json:"locale"`
	Premium               bool   `json:"premium"`
	IDVerified            bool   `json:"idVerified"`
	SocialNetworkProfiles struct {
		Facebook string `json:"facebook"`
		Twitter  string `json:"twitter"`
		YouTube  string `json:"youtube"`
		Twitch   string `json:"twitch"`
		Guilded  string `json:"guilded"`
		// Possible values:
		// "SOCIAL_NETWORK_VISIBILITY_UNSPECIFIED" | "NO_ONE" | "FRIENDS" | "FRIENDS_AND_FOLLOWING" | "FRIENDS_FOLLOWING_AND_FOLLOWERS" | "EVERYONE"
		Visbility string `json:"visibility"`
	} `json:"socialNetworkProfiles"`
}

type UserThumbnail struct {
	ImageURI string `json:"imageUri"`
}

type GenerateUserThumbnailQuery struct {
	Size   *int    `json:"size"`
	Format *string `json:"format"`
	Shape  *string `json:"shape"`
}

func (query *GenerateUserThumbnailQuery) Validate() error {
	supportedSizes := []int{
		48, 50, 60, 75, 100, 110, 150, 180, 352, 420, 720,
	}

	if !slices.Contains(supportedSizes, *query.Size) {
		return fmt.Errorf("GenerateUserThumbnailQuery: Size must be one of the following: %v", supportedSizes)
	}

	return nil
}

func (query *GenerateUserThumbnailQuery) ConvertToStringMap() map[string]string {
	result := make(map[string]string)

	if query == nil {
		return result
	}

	if query.Size != nil {
		result["size"] = strconv.Itoa(*query.Size)
	}

	if query.Format != nil {
		result["format"] = *query.Format
	}

	if query.Shape != nil {
		result["shape"] = *query.Shape
	}

	return result
}

type GameJoinRestriction struct {
	Active             bool   `json:"active"`
	StartTime          string `json:"startTime"`
	Duration           int    `json:"duration"`
	PrivateReason      string `json:"privateReason"`
	DisplayReason      string `json:"displayReason"`
	ExcludeAltAccounts bool   `json:"excludeAltAccounts"`
	Inherited          bool   `json:"inherited"`
}

type UserRestriction struct {
	Path                 string              `json:"path"`
	UpdateTime           string              `json:"updateTime"`
	User                 string              `json:"user"`
	GameJoinRestrictions GameJoinRestriction `json:"gameJoinRestriction"`
}

type UserRestrictionsList struct {
	UserRestrictions []UserRestriction `json:"userRestrictions"`
	NextPageToken    string            `json:"nextPageToken"`
}
