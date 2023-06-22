package util

import (
	"encoding/json"
)

type UserDTO struct {
	Status string `json:"status"`
	Data   struct {
		User struct {
			Biography string `json:"biography"`
			BioLinks  []struct {
				Title    string `json:"title"`
				LynxURL  string `json:"lynx_url"`
				URL      string `json:"url"`
				LinkType string `json:"link_type"`
			} `json:"bio_links"`
			BiographyWithEntities struct {
				RawText  string `json:"raw_text"`
				Entities []any  `json:"entities"`
			} `json:"biography_with_entities"`
			BlockedByViewer        bool   `json:"blocked_by_viewer"`
			CountryBlock           bool   `json:"country_block"`
			ExternalURL            string `json:"external_url"`
			ExternalURLLinkshimmed string `json:"external_url_linkshimmed"`
			EdgeFollowedBy         struct {
				Count int64 `json:"count"`
			} `json:"edge_followed_by"`
			EimuID     string `json:"eimu_id"`
			Fbid       string `json:"fbid"`
			EdgeFollow struct {
				Count int64 `json:"count"`
			} `json:"edge_follow"`
			FullName              string `json:"full_name"`
			HasArEffects          bool   `json:"has_ar_effects"`
			HasClips              bool   `json:"has_clips"`
			HasGuides             bool   `json:"has_guides"`
			HasChannel            bool   `json:"has_channel"`
			HighlightReelCount    int    `json:"highlight_reel_count"`
			HasRequestedViewer    bool   `json:"has_requested_viewer"`
			HideLikeAndViewCounts bool   `json:"hide_like_and_view_counts"`
			ID                    string `json:"id"`
			IsBusinessAccount     bool   `json:"is_business_account"`
			IsProfessionalAccount bool   `json:"is_professional_account"`
			IsSupervisionEnabled  bool   `json:"is_supervision_enabled"`
			IsGuardianOfViewer    bool   `json:"is_guardian_of_viewer"`
			IsSupervisedByViewer  bool   `json:"is_supervised_by_viewer"`
			IsSupervisedUser      bool   `json:"is_supervised_user"`
			IsEmbedsDisabled      bool   `json:"is_embeds_disabled"`
			IsJoinedRecently      bool   `json:"is_joined_recently"`
			IsPrivate             bool   `json:"is_private"`
			IsVerified            bool   `json:"is_verified"`
			ProfilePicURL         string `json:"profile_pic_url"`
			ProfilePicURLHd       string `json:"profile_pic_url_hd"`
			Username              string `json:"username"`
			ConnectedFbPage       any    `json:"connected_fb_page"`
			Pronouns              []any  `json:"pronouns"`
		} `json:"user"`
	} `json:"data"`
}

func StringToJSON(str string) (UserJSON UserDTO, err error) {
	err = json.Unmarshal([]byte(str), &UserJSON)
	return
}
