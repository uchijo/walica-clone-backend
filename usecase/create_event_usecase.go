package usecase

func CreateEvent(repo Repository, members []string, eventName string) (string, error) {
	eventId, err := repo.CreateEvent(eventName, members)
	if err != nil {
		return "", err
	}
	return eventId, nil
}
