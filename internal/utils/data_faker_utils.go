package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

// Create a global random source and generator, protected by a sync.Mutex for thread safety
var (
	random   *rand.Rand
	randomMu sync.Mutex
)

func init() {
	// Initialize the global random generator with a single source
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// GenerateBool generates a bool true or false
func GenerateBool() bool {
	randomMu.Lock()
	defer randomMu.Unlock()

	randomBool := random.Intn(2) == 0

	return randomBool
}

// GenerateInt generates a int between a min int and a max int
func GenerateInt(min int, max int) int {
	randomMu.Lock()
	defer randomMu.Unlock()

	// Generate a random int in the range [0, max-min] and add min to shift it into the [min, max] range
	return min + random.Intn(max-min+1)
}

// GenerateFloat64 generates a float64 between a min float64 and a max float64
func GenerateFloat64(min float64, max float64) float64 {
	randomMu.Lock()
	defer randomMu.Unlock()

	// Generate a random float64 in the range [min, max]
	return min + (random.Float64() * (max - min))
}

// GenerateDate generates a date between a minDate and a maxDate
func GenerateDate(minDate time.Time, maxDate time.Time) time.Time {
	randomMu.Lock()
	defer randomMu.Unlock()

	// Calculate the duration range in seconds
	minUnix := minDate.Unix()
	maxUnix := maxDate.Unix()

	// Generate a random Unix timestamp in the range [minUnix, maxUnix]
	randomUnix := random.Int63n(maxUnix-minUnix+1) + minUnix

	// Convert the random Unix timestamp to a time.Time
	return time.Unix(randomUnix, 0)
}

// GenerateRandomNumbersString generates a string made of n numbers from 0 to 9
func GenerateRandomNumbersString(n int) string {
	randomMu.Lock()
	defer randomMu.Unlock()

	// Use a strings.Builder for efficient string concatenation
	var sb strings.Builder

	for i := 0; i < n; i++ {
		// Append a random digit (0-9) to the string
		sb.WriteByte('0' + byte(random.Intn(10)))
	}

	return sb.String()
}

func FundNameGenerator() string {
	randomMu.Lock()
	defer randomMu.Unlock()

	// Define word lists for different parts of the fund name
	nouns := []string{"Growth", "Value", "Income", "Equity", "Bond", "Index", "Global", "Emerging", "Tactical", "Balanced"}
	adjectives := []string{"Strong", "Dynamic", "Strategic", "Diversified", "Focused", "Core", "Enhanced", "Managed", "Absolute", "High"}
	locations := []string{"US", "International", "Global", "Emerging Markets", "Asia", "Europe", "Africa", "Latin America", "Pacific", "China", "India", "Japan", "Brazil", "Russia", "Canada", "Australia", "UK", "Germany", "France", "Italy", "Spain", "Netherlands", "Switzerland", "Sweden", "Norway", "Denmark", "Finland", "Poland", "Turkey", "South Africa", "Nigeria", "Egypt", "Saudi Arabia", "UAE", "Qatar", "Kuwait", "Bahrain", "Oman", "Jordan", "Lebanon"}

	// Randomly select words from each list
	noun := nouns[random.Intn(len(nouns))]
	adjective := adjectives[random.Intn(len(adjectives))]
	location := locations[random.Intn(len(locations))]

	// Combine words to form the fund name
	fundName := fmt.Sprintf("%s %s %s Fund", adjective, noun, location)

	return fundName
}

func SettlementNameGenerator() string {
	randomMu.Lock()
	defer randomMu.Unlock()

	// Define word lists for different parts of the receivable name
	months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	types := []string{"Receivable", "Invoice", "Payment", "Balance Due", "Settlement", "Outstanding"}
	clients := []string{"Deutsche Bank", "Aberdeen Standard", "Columbia Threadneedle", "Neuberger Berman", "Franklin Templeton", "Northern Trust", "Previ", "Petros", "Funcef", "Postalis", "Vale Previdência", "Montblanc Previdência", "Bradesco Previdência", "Itaú Previdência", "BTG Pactual"}
	transactionTypes := []string{"Product Sale", "Service Payment", "Loan Repayment", "Subscription", "Interest Payment", "Wholesale"}
	priority := []string{"High-Priority", "Standard", "Urgent", "Regular", "Late"}

	// Randomly select elements from each list
	month := months[rand.Intn(len(months))]
	transactionType := transactionTypes[rand.Intn(len(transactionTypes))]
	client := clients[rand.Intn(len(clients))]
	typeOfReceivable := types[rand.Intn(len(types))]
	priorityLevel := priority[rand.Intn(len(priority))]

	// Combine selected words into a receivable name
	var name strings.Builder
	name.WriteString(fmt.Sprintf("%s %s %s - %s %s", priorityLevel, transactionType, month, client, typeOfReceivable))

	// Return the generated name
	return name.String()
}

// GenerateFakeNotificationTitle generates a random fake notification title
func GenerateFakeNotificationTitle() string {
	randomMu.Lock()
	defer randomMu.Unlock()

	// NotificationType represents the type of notification
	type NotificationType string

	const (
		SystemAlert  NotificationType = "System Alert"
		Promotion    NotificationType = "Promotion"
		SocialUpdate NotificationType = "Social Update"
	)

	// Combine types and texts
	types := []NotificationType{SystemAlert, Promotion, SocialUpdate}
	randomType := types[rand.Intn(len(types))]

	return string(randomType)
}

// GenerateFakeNotificationText generates a random fake notification text
func GenerateFakeNotificationText() string {
	randomMu.Lock()
	defer randomMu.Unlock()

	// NotificationType represents the type of notification
	type NotificationType string

	const (
		SystemAlert  NotificationType = "System Alert"
		Promotion    NotificationType = "Promotion"
		SocialUpdate NotificationType = "Social Update"
	)

	// Sample notification texts
	systemAlerts := []string{
		"System update available. Please restart your device.",
		"Your account login was detected from a new device.",
		"Low disk space. Consider cleaning up unnecessary files.",
	}

	promotions := []string{
		"Flash Sale: Get up to 50% off on electronics!",
		"Subscribe now and get a free trial for 30 days.",
		"Special offer: Buy 1 Get 1 Free on all clothing items.",
	}

	socialUpdates := []string{
		"John Doe liked your post.",
		"You have a new friend request from Jane Smith.",
		"Your photo received 20 new likes.",
	}

	// Combine types and texts
	types := []NotificationType{SystemAlert, Promotion, SocialUpdate}
	randomType := types[rand.Intn(len(types))]

	var message string
	switch randomType {
	case SystemAlert:
		message = systemAlerts[rand.Intn(len(systemAlerts))]
	case Promotion:
		message = promotions[rand.Intn(len(promotions))]
	case SocialUpdate:
		message = socialUpdates[rand.Intn(len(socialUpdates))]
	}

	return fmt.Sprintf("[%s] %s", randomType, message)
}
