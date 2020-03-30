package slack

import (
	"fmt"
	"net/url"
	"time"

	"github.com/CptOfEvilMinions/AskJeevesSecBot/pkg/brokers"
	"github.com/CptOfEvilMinions/AskJeevesSecBot/pkg/config"
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

// generateUserInformationBlocks input:
// generateUserInformationBlocks output:
func generateUserInformationBlocks(vpnEntry brokers.VPNdata, location string, vpnHash string) []*slack.TextBlockObject {
	// User information fields array
	var userInformationBlocks []*slack.TextBlockObject
	userInformationBlocks = append(userInformationBlocks, slack.NewTextBlockObject("mrkdwn", "*EventID*\n123abc", false, false))                                                      // EventID
	userInformationBlocks = append(userInformationBlocks, slack.NewTextBlockObject("mrkdwn", fmt.Sprintf("*Username*\n%s", vpnEntry.Username), false, false))                         // Username
	userInformationBlocks = append(userInformationBlocks, slack.NewTextBlockObject("mrkdwn", fmt.Sprintf("*Timestamp*\n%s", vpnEntry.Timestamp.Format(time.RubyDate)), false, false)) // Timestamp
	userInformationBlocks = append(userInformationBlocks, slack.NewTextBlockObject("mrkdwn", fmt.Sprintf("*Location*\n%s", location), false, false))                                  // Location
	userInformationBlocks = append(userInformationBlocks, slack.NewTextBlockObject("mrkdwn", fmt.Sprintf("*IPaddress*\n%s", vpnEntry.SrcIP), false, false))                           // IP address
	userInformationBlocks = append(userInformationBlocks, slack.NewTextBlockObject("mrkdwn", fmt.Sprintf("*VPNhash*\n%s", vpnHash), false, false))                                    // VPNhash
	userInformationBlocks = append(userInformationBlocks, slack.NewTextBlockObject("mrkdwn", "*Device*\nmacOS", false, false))                                                        // Device
	userInformationBlocks = append(userInformationBlocks, slack.NewTextBlockObject("mrkdwn", "*Hostname*\nstarwars", false, false))                                                   // Hostname
	return userInformationBlocks
}

// SendUserMessage input:
// SendUserMessage output:
// https://godoc.org/github.com/nlopes/slack#Client.SendMessage
// https://github.com/slack-go/slack/issues/603
func SendUserMessage(cfg *config.Config, slackAPI *slack.Client, vpnEntry brokers.VPNdata, location string, vpnHash string) error {
	// Divider between sections
	divSection := slack.NewDividerBlock()

	// Header Section
	headerText := slack.NewTextBlockObject("mrkdwn", ":red_circle: New VPN login", false, false)
	headerSection := slack.NewSectionBlock(headerText, nil, nil)

	// Google Map static image
	imageText := slack.NewTextBlockObject("plain_text", "Location", false, false)
	imageBlock := slack.NewImageBlock(fmt.Sprintf("https://hackinglab.beer/askjeeves/GoogleMaps?location=%s", url.QueryEscape(location)), "Marker", "test", imageText)

	// Generate user information fields
	userInformationBlocks := generateUserInformationBlocks(vpnEntry, location, vpnHash)
	userInformationSecion := slack.NewSectionBlock(nil, userInformationBlocks, nil)

	//////////////////////////////////////// User selection section ////////////////////////////////////////
	//var userSelectorBtns []*slack.BlockElements

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

	//userSelectorBtns = append(userSelectorBtns, legitimateBtn)
	//userSelectorBtns = append(userSelectorBtns, unauthorizedBtn)

	//userSelectorActionBlock := slack.NewActionBlock("", userSelectorBtns)

	params := slack.MsgOptionBlocks(
		headerSection,
		imageBlock,
		divSection,
		userInformationSecion,
		divSection,
		//userSelectorActionBlock,
		legitimateActionBlock,
		unauthorizedActionBlock,
	)

	slackUsername := "@" + vpnEntry.Username // Prepend username with @

	if _, _, err := slackAPI.PostMessage(slackUsername, slack.MsgOptionText("", false), params); err != nil {
		return fmt.Errorf("failed to post message: %s", err)
	}
	return nil
}
