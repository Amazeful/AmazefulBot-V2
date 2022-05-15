package filters

import (
	"unicode"

	"github.com/Amazeful/Amazefulbot/message"
	"github.com/Amazeful/dataful/models"
)

//CapsFilter moderates messages with excessive capitalized characters.
type CapsFilter struct {
	channelConfig  *models.Channel
	filterSettings *models.CapsFilter
}

func NewCapsFilter(channelConfig *models.Channel, filterSettings *models.CapsFilter) *CapsFilter {
	return &CapsFilter{
		channelConfig:  channelConfig,
		filterSettings: filterSettings,
	}
}

//ProcessMessage processes a chat message.
func (cf *CapsFilter) ProcessMessage(message *message.ChatMessage) {
	if cf.shouldExit(message) {
		return
	}

	capsCount := cf.countUpperCase(message.Text)

	if capsCount > cf.filterSettings.MaxCount {

	}

}

//countUpperCase counts number of uppercase characters in a string.
func (cf *CapsFilter) countUpperCase(message string) int {
	result := 0
	for _, c := range message {
		if unicode.IsUpper(c) {
			result++
		}
	}

	return result
}

//getPercantage gets what percentage of message is capitalized.
func (cf *CapsFilter) getPercantage(capsCount, length int) float64 {
	return float64(capsCount) / float64(length)
}

//shouldExit checks if any of filter conditions are not met.
func (cf *CapsFilter) shouldExit(message *message.ChatMessage) bool {
	//do not process if message shorter than min chars
	if message.Length < cf.filterSettings.MinChars {
		return true
	}

	return false
}
