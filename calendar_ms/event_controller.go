package main

func createEvent(event Event) error {
	bd, err := getDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("INSERT INTO event (description, link, start_time, end_time, status) VALUES (?, ?, ?, ?, ?)", event.Description, event.Link, event.Start_time, event.End_time, event.Status)
	return err
}

func deleteEvent(id int64) error {

	bd, err := getDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("DELETE FROM event WHERE id = ?", id)
	return err
}

// It takes the ID to make the update
func updateEvent(event Event) error {
	bd, err := getDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("UPDATE event SET description = ?, link = ?, start_time = ?, end_time = ?, status = ? WHERE id = ?", event.Description, event.Link, event.Start_time, event.End_time, event.Status, event.Id)
	return err
}

func getEvents() ([]Event, error) {
	//Declare an array because if there's error, we return it empty
	events := []Event{}
	bd, err := getDB()
	if err != nil {
		return events, err
	}
	// Get rows so we can iterate them
	rows, err := bd.Query("SELECT id, COALESCE(description, '') as description, COALESCE(link, '') as link, start_time, end_time, status FROM event")
	if err != nil {
		return events, err
	}
	// Iterate rows...
	for rows.Next() {
		// In each step, scan one row
		var event Event
		err = rows.Scan(&event.Id, &event.Description, &event.Link, &event.Start_time, &event.End_time, &event.Status)
		if err != nil {
			return events, err
		}
		// and append it to the array
		events = append(events, event)
	}
	return events, nil
}

func getEventById(id int64) (Event, error) {
	var event Event
	bd, err := getDB()
	if err != nil {
		return event, err
	}
	row := bd.QueryRow("SELECT id, description, link, start_time, end_time, status FROM event WHERE id = ?", id)
	err = row.Scan(&event.Id, &event.Description, &event.Link, &event.Start_time, &event.End_time, &event.Status)
	if err != nil {
		return event, err
	}
	// Success!
	return event, nil
}
