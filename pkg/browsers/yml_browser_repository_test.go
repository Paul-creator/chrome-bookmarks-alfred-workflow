package browsers

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestYmlBrowserRepository_GetBrowsers(t *testing.T) {
	test := assert.New(t)
	browserRepository := YmlBrowserRepository{Filename: "test/browser.yml"}
	browserSlice, err := browserRepository.GetBrowsers()

	test.NoError(err)
	test.Len(browserSlice, 3)
	test.Equal("Profile 1", browserSlice[0].ProfileFolderName)
}

func TestYmlBrowserRepository_GetBrowsers_ReadFileError(t *testing.T) {
	test := assert.New(t)
	browserRepository := YmlBrowserRepository{Filename: "test"}
	browserSlice, err := browserRepository.GetBrowsers()

	test.Error(err)
	test.Len(browserSlice, 0)
}

func TestYmlBrowserRepository_UpdateBrowser(t *testing.T) {
	test := assert.New(t)
	filename := path.Join("test", "marat", "browser.yml")
	dirname := path.Dir(filename)
	bytes, _ := ioutil.ReadFile("test/browser.yml")

	_ = os.MkdirAll(dirname, 0777)
	_ = ioutil.WriteFile(filename, bytes, 0644)

	repository := YmlBrowserRepository{Filename: filename}
	browsersSlice, _ := repository.GetBrowsers()
	browser, _ := browsersSlice.FindByName("chrome")

	browser.ProfileFolderName = defaultProfileFolderName

	err := repository.UpdateBrowser(browser)

	test.NoError(err)

	browsersSlice, _ = repository.GetBrowsers()
	browser, _ = browsersSlice.FindByName("chrome")

	test.Equal(defaultProfileFolderName, browser.ProfileFolderName)

	_ = os.RemoveAll(dirname)
}
