package profiles

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/mdreizin/chrome-bookmarks-alfred-workflow/pkg/browsers"
)

type profileAux struct {
	Root struct {
		Profiles map[string]Profile `json:"info_cache"`
		Name     string             `json:"last_used,omitempty"`
	} `json:"profile"`
}

type JsonProfileRepository struct {
	Filename string
}

func (r *JsonProfileRepository) GetProfiles(browser *browsers.Browser) (ProfileSlice, error) {
	filename := browser.ResolvePath(r.Filename)

	bytes, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	aux := profileAux{}
	err = json.Unmarshal(bytes, &aux)

	if err != nil {
		return nil, err
	}

	profileSlice := ProfileSlice{}

	profileFolderName := browser.ProfileFolderName

	if profileFolderName == "" {
		profileFolderName = aux.Root.Name
	}

	aux.Root.Profiles[AutoProfile.Name] = AutoProfile

	for k, v := range aux.Root.Profiles {
		name := v.Name

		if name == "" {
			name = k
		}

		profileSlice = profileSlice.Add(&Profile{
			Name:            name,
			IsVirtual:       v.IsVirtual,
			IsActive:        strings.EqualFold(name, profileFolderName),
			AvatarURL:       v.AvatarURL,
			IconURL:         v.AvatarIconURL(browser, k),
			CustomAvatarURL: v.CustomAvatarURL,
			IsDefaultAvatar: v.IsDefaultAvatar,
			DisplayName:     v.DisplayName,
			UserName:        v.UserName,
			UserEmail:       v.UserEmail,
		})
	}

	return profileSlice, nil
}
