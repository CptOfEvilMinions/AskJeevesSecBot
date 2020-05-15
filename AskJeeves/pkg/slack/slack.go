package slack

import (
	"fmt"
	"net/url"
	"time"

	"github.com/CptOfEvilMinions/AskJeevesSecBot/pkg/config"
	"github.com/CptOfEvilMinions/AskJeevesSecBot/pkg/database"
	"github.com/slack-go/slack"
)

// InitConnector input:
// InitConnector outout:
func InitConnector(cfg *config.Config) *slack.Client {
	fmt.Println("=============== Slack ===============")
	fmt.Println("Token:", cfg.Slack.Token)
	// Init Slack API
	return slack.New(cfg.Slack.Token)
}


// generateUserInformationBlocks input: vpnEntry about login, location, and vpnHash
// generateUserInformationBlocks output: Return Slack information text block based on VPN login
func generateUserInformationBlocks(userVPNLog database.UserVPNLog) []*slack.TextBlockObject {
	// User information fields array
	var userInformationBlocks []*slack.TextBlockObject
	userInformationBlocks = append(userInformationBlocks, slack.NewTextBlockObject("mrkdwn", fmt.Sprintf("*EventID*\n%s", userVPNLog.EventID), false, false))                           // EventID
	userInformationBlocks = append(userInformationBlocks, slack.NewTextBlockObject("mrkdwn", fmt.Sprintf("*Username*\n%s", userVPNLog.Username), false, false))                         // Username
	userInformationBlocks = append(userInformationBlocks, slack.NewTextBlockObject("mrkdwn", fmt.Sprintf("*Timestamp*\n%s", userVPNLog.CreatedAt.Format(time.RubyDate)), false, false)) // Timestamp
	userInformationBlocks = append(userInformationBlocks, slack.NewTextBlockObject("mrkdwn", fmt.Sprintf("*Location*\n%s", userVPNLog.Location), false, false))                         // Location
	userInformationBlocks = append(userInformationBlocks, slack.NewTextBlockObject("mrkdwn", fmt.Sprintf("*IPaddress*\n%s", userVPNLog.IPaddr), false, false))                          // IP address
	userInformationBlocks = append(userInformationBlocks, slack.NewTextBlockObject("mrkdwn", fmt.Sprintf("*VPNhash*\n%s", userVPNLog.VpnHash), false, false))                           // VPNhash
	userInformationBlocks = append(userInformationBlocks, slack.NewTextBlockObject("mrkdwn", fmt.Sprintf("*Device*\n%s", userVPNLog.Device), false, false))                             // Device
	userInformationBlocks = append(userInformationBlocks, slack.NewTextBlockObject("mrkdwn", fmt.Sprintf("*Hostname*\n%s", userVPNLog.Hostname), false, false))                         // Hostname
	return userInformationBlocks
}


// SendUserMessage input: config, slack connector, vpnEntry, location, and vpnHash
// SendUserMessage output: Return error if there is one
// https://godoc.org/github.com/nlopes/slack#Client.SendMessage
// https://github.com/slack-go/slack/issues/603
func SendUserMessage(cfg *config.Config, slackAPI *slack.Client, userVPNLog database.UserVPNLog) error {
	// Divider between sections
	divSection := slack.NewDividerBlock()

	// Header Section
	headerText := slack.NewTextBlockObject("mrkdwn", ":red_circle: New VPN login", false, false)
	headerSection := slack.NewSectionBlock(headerText, nil, nil)

	// Google Map static image
	imageText := slack.NewTextBlockObject("plain_text", "Location", false, false)
	imageBlock := slack.NewImageBlock(fmt.Sprintf("%s/askjeeves/GoogleMaps?location=%s", cfg.ButlingButler.URL, url.QueryEscape(userVPNLog.Location)), "Marker", "test", imageText)

	// Generate user information fields
	userInformationBlocks := generateUserInformationBlocks(userVPNLog)
	userInformationSecion := slack.NewSectionBlock(nil, userInformationBlocks, nil)

	//////////////////////////////////////// User selection section ////////////////////////////////////////

	// Legititmate login
	legitimateBtnTxt := slack.NewTextBlockObject("plain_text", "This was me", false, false)
	legitimateBtn := slack.NewButtonBlockElement("", "legitimate_login", legitimateBtnTxt)
	legitimateBtn.WithStyle("primary")
	legitimateActionBlock := slack.NewActionBlock("", legitimateBtn)

	// UNauthorized login
	unauthorizedBtnTxt := slack.NewTextBlockObject("plain_text", "This was NOT me", false, false)
	unauthorizedBtn := slack.NewButtonBlockElement("", "unauthorized_login", unauthorizedBtnTxt)
	unauthorizedBtn.WithStyle("danger")
	unauthorizedActionBlock := slack.NewActionBlock("", unauthorizedBtn)

	params := slack.MsgOptionBlocks(
		headerSection,
		imageBlock,
		divSection,
		userInformationSecion,
		divSection,
		legitimateActionBlock,
		unauthorizedActionBlock,
	)

	slackUsername := "@" + userVPNLog.Username // Prepend username with @

	if _, _, err := slackAPI.PostMessage(slackUsername, slack.MsgOptionText("", false), params); err != nil {
		return fmt.Errorf("failed to post message: %s", err)
	}
	
	return nil
}
