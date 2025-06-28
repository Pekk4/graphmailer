package graphmailer

// RemoveSublist removes all elements in sublist from the main list
func RemoveSublist(mainlist, sublist []string) []string {
	result := []string{}
	skip := make(map[string]bool)

	// Create a map of emails to skip
	for _, email := range sublist {
		skip[email] = true
	}

	// Add non wanted emails to the result
	for _, email := range mainlist {
		if !skip[email] {
			result = append(result, email)
		}
	}

	return result
}
