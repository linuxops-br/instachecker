package instachecker

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/linuxops-br/instachecker/pkg/appid"
	"github.com/linuxops-br/instachecker/pkg/util"
)

type User struct {
	ID           string
	Username     string
	Name         string
	Followers    int64
	Following    int64
	LargePicture string
	SmallPicture string
	Biography    string
	Biolinks     []struct {
		Title    string
		LynxURL  string
		URL      string
		LinkType string
	}
	IsBusinessAccount     bool
	IsProfessionalAccount bool
	IsPrivateAccount      bool
	IsVerifiedAccount     bool
	IsSupervisionEnabled  bool
	HiddenLikesAndViews   bool
	ReelCount             int
	IsNewAccount          bool
	responseBody          string
}

type InstagramUser interface {
	GetName() string
	GetResponse() string
	GetUserName() string
	GetBiography() string
	GetID() string
	GetBioLinks() []struct {
		Title    string
		LynxURL  string
		URL      string
		LinkType string
	}
	GetFollowersCount() int64
	GetFollowCount() int64
	GetPicture() struct {
		Small string
		Large string
	}
	GetReelCount() int
	CheckAccountInfo() map[string]bool
}

func NewInstagramUser(username string) InstagramUser {
	api := appid.New()
	str, _ := httpGetInfo(username, api.Get())

	dto, _ := util.StringToJSON(str)

	user := &User{
		ID:           dto.Data.User.ID,
		Username:     dto.Data.User.Username,
		Name:         dto.Data.User.FullName,
		Followers:    dto.Data.User.EdgeFollowedBy.Count,
		Following:    dto.Data.User.EdgeFollow.Count,
		SmallPicture: dto.Data.User.ProfilePicURL,
		LargePicture: dto.Data.User.ProfilePicURLHd,
		Biography:    dto.Data.User.Biography,
		Biolinks: []struct {
			Title    string
			LynxURL  string
			URL      string
			LinkType string
		}(dto.Data.User.BioLinks),
		ReelCount:             dto.Data.User.HighlightReelCount,
		HiddenLikesAndViews:   dto.Data.User.HideLikeAndViewCounts,
		IsProfessionalAccount: dto.Data.User.IsProfessionalAccount,
		IsBusinessAccount:     dto.Data.User.IsBusinessAccount,
		IsNewAccount:          dto.Data.User.IsJoinedRecently,
		IsSupervisionEnabled:  dto.Data.User.IsSupervisionEnabled,
		IsVerifiedAccount:     dto.Data.User.IsVerified,
		IsPrivateAccount:      dto.Data.User.IsPrivate,
		responseBody:          str,
	}

	return user

}

func httpGetInfo(username string, appid string) (string, error) {
	url := fmt.Sprintf("https://www.instagram.com/api/v1/users/web_profile_info/?username=%s", username)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("x-ig-app-id", appid)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(resBody), nil
}

func (u User) GetName() string {
	return u.Name
}

func (u User) GetResponse() string {
	return u.responseBody
}

func (u User) GetID() string {
	return u.ID
}

func (u User) GetPicture() struct {
	Small string
	Large string
} {
	return struct {
		Small string
		Large string
	}{
		Small: u.SmallPicture,
		Large: u.LargePicture,
	}
}

func (u User) GetUserName() string {
	return u.Username
}

func (u User) GetBiography() string {
	return u.Biography
}

func (u User) GetBioLinks() []struct {
	Title    string
	LynxURL  string
	URL      string
	LinkType string
} {
	return u.Biolinks
}

func (u User) GetFollowersCount() int64 {
	return u.Followers
}

func (u User) GetFollowCount() int64 {
	return u.Following
}

func (u User) GetReelCount() int {
	return u.ReelCount
}

func (u User) CheckAccountInfo() map[string]bool {

	return map[string]bool{
		"IsNewAccount":          u.IsNewAccount,
		"IsBusinessAccount":     u.IsBusinessAccount,
		"IsProfessionalAccount": u.IsProfessionalAccount,
		"IsSupervisionEnabled":  u.IsSupervisionEnabled,
		"IsPrivateAccount":      u.IsPrivateAccount,
		"IsVerifiedAccount":     u.IsVerifiedAccount,
		"IsHideLikesAndViews":   u.HiddenLikesAndViews,
	}
}
