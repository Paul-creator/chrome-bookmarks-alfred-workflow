package browsers

import (
	"io/ioutil"
	"regexp"
	"strings"

	"gopkg.in/yaml.v2"
)

type YmlBrowserRepository struct {
	Filename string
}

func (r *YmlBrowserRepository) GetBrowsers() (BrowserSlice, error) {
	bytes, err := ioutil.ReadFile(r.Filename)

	if err != nil {
		return nil, err
	}

	browserSlice := BrowserSlice{}
	err = yaml.Unmarshal(bytes, &browserSlice)

	if err != nil {
		return nil, err
	}

	re := regexp.MustCompile(`(\\ )`)

	for i, v := range browserSlice {
		browserSlice[i].ProfileFolderName = re.ReplaceAllString(v.ProfileFolderName, " ")
	}

	return browserSlice, nil
}

func (r *YmlBrowserRepository) UpdateBrowser(b *Browser) error {
	browserSlice, err := r.GetBrowsers()

	if err != nil {
		return err
	}

	i := browserSlice.FindIndex(func(v *Browser) bool { return strings.EqualFold(b.Name, v.Name) })

	if i >= 0 {
		browserSlice[i] = *b

		bytes, err := yaml.Marshal(browserSlice)

		if err != nil {
			return err
		}

		err = ioutil.WriteFile(r.Filename, bytes, 0644)
	}

	return err
}
