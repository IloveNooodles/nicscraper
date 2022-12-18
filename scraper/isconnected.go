package scraper

func (s Scraper) IsConnected() bool {

	/* Test connection. */

	if s.Args.UseTeams {
		_, err := s.GetByNIMTeams("13518035")
		return err == nil
	}

	_, err := s.GetByNIM("13518035")
	return err == nil
}
