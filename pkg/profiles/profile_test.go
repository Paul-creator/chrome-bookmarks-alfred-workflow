package profiles

import (
	"testing"

	"github.com/mdreizin/chrome-bookmarks-alfred-workflow/pkg/browsers"
	"github.com/stretchr/testify/assert"
)

func TestProfile_AvatarIconURL(t *testing.T) {
	test := assert.New(t)
	browser := browsers.Browser{Path: "test", IconURL: "img/chrome.png"}
	profile := Profile{
		AvatarURL:       "chrome://theme/IDR_PROFILE_AVATAR_12",
		CustomAvatarURL: "avatar.png",
		IsDefaultAvatar: false,
	}

	test.Equal("img/chrome.png", profile.AvatarIconURL(&browser, DefaultProfileFolderName))

	profile.IsDefaultAvatar = true

	test.Equal("test/Avatars/avatar_ninja.png", profile.AvatarIconURL(&browser, DefaultProfileFolderName))
}
