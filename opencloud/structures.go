package opencloud

import (
	"errors"
	"fmt"
	"strconv"
)

// -- General
type ExperienceOwner struct {
}

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
