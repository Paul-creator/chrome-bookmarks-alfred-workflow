package browsers

import (
	"os/user"
	"path"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	fullPath                 = "~/Library/Application Support/Google/Chrome"
	defaultProfileFolderName = "Default"
)

func pathFor(elem ...string) string {
	paths := append([]string{fullPath}, elem...)

	return path.Join(paths...)
}

func resolvePath(elem ...string) string {
	usr, _ := user.Current()
	dir := usr.HomeDir

	return strings.Replace(pathFor(elem...), tilde, dir+sep, 1)
}

func TestBrowser_JoinPath(t *testing.T) {
	test := assert.New(t)
	browser := &Browser{Path: fullPath}

	test.Equal(path.Join(fullPath, defaultProfileFolderName), browser.JoinPath(defaultProfileFolderName))
}

func TestBrowser_ResolvePath(t *testing.T) {
	test := assert.New(t)
	browser := &Browser{Path: fullPath}

	test.Equal(resolvePath(), browser.ResolvePath())
	test.Equal(resolvePath(defaultProfileFolderName), browser.ResolvePath(defaultProfileFolderName))
}
