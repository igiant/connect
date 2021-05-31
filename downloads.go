package connect

type DownloadList []Download

// Download management

// DownloadsRemove - Remove file prepared to download.
// Parameters
//	url - url of file prepared to download
func (s *ServerConnection) DownloadsRemove(url string) error {
	params := struct {
		Url string `json:"url"`
	}{url}
	_, err := s.CallRaw("Downloads.downloadsRemove", params)
	return err
}
