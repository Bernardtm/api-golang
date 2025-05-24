package emails

func MapAddresses(emails []string) []AddressName {
	var addresses []AddressName
	for _, email := range emails {
		addresses = append(addresses, AddressName{
			Email: email,
			Name:  "", // Add name if available
		})
	}
	return addresses
}
