package scraper

func (s Scraper) IsConnected() bool {

	/* Test connection. */
	_, err := s.GetByNIM("13518035")
	return err == nil
}

func (s TeamsScrapper) IsConnected() bool {
  _, err := s.GetByNIMTeams("13518035")
  return err == nil
}
