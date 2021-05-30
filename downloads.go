package connect

type DownloadList []Download

// Download management

// DownloadsRemove - Remove file prepared to download.
// Parameters
//	url - url of file prepared to download
func (c *ServerConnection) DownloadsRemove(url string) error {
	params := struct {
		Url string `json:"url"`
	}{url}
	_, err := c.CallRaw("Downloads.downloadsRemove", params)
	return err
}
