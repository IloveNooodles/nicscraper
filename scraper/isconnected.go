package scraper

func (s Scraper) IsConnected() bool {
	// Test connection.
	_, err := s.GetByNIM("13518035")
	return err == nil
}
