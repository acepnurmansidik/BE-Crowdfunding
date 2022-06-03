package campaign

import (
	"strconv"
	"strings"
)

type CampaignsFormatter struct {
	ID               int    `json:"id"`
	Title            string `json:"title"`
	ShortDescription string `json:"short_description"`
	ImageUrl         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	UserID           int    `json:"user_id"`
	Slug             string `json:"slug"`
}

type CampaignDetailFormatter struct {
	ID               int                   `json:"id"`
	Title            string                `json:"title"`
	ShortDescription string                `json:"short_description"`
	ImageUrl         string                `json:"image_url"`
	GoalAmount       int                   `json:"goal_amount"`
	CurrentAmount    int                   `json:"current_amount"`
	UserID           int                   `json:"user_id"`
	Slug             string                `json:"slug"`
	Perks            []string              `json:"perks"`
	User             CampaignUserFormatter `json:"user"`
}

type CampaignUserFormatter struct {
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

func FormatCampaign(campaign Campaign) CampaignsFormatter {
	id, _ := strconv.Atoi(campaign.ID)
	formatter := CampaignsFormatter{
		ID:               id,
		Title:            campaign.Name,
		ShortDescription: campaign.ShortDescription,
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		Slug:             campaign.Slug,
		ImageUrl:         "",
	}

	if len(campaign.CampaignImages) > 0 {
		formatter.ImageUrl = campaign.CampaignImages[0].FileName
	}

	return formatter
}

func FormatCampaigns(campaigns []Campaign) []CampaignsFormatter {
	// buat slice utk menampung datanya
	campaignsFormatter := []CampaignsFormatter{}
	// looping setiap datanya
	for _, campaign := range campaigns {
		campaignFormatter := FormatCampaign(campaign)
		// simpan ke dalam slice
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}

	return campaignsFormatter
}

func FormatCampaignDetail(campaign Campaign) CampaignDetailFormatter {
	id, _ := strconv.Atoi(campaign.ID)
	formatter := CampaignDetailFormatter{
		ID:               id,
		Title:            campaign.Name,
		ShortDescription: campaign.ShortDescription,
		ImageUrl:         "",
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		UserID:           campaign.UserID,
		Slug:             campaign.Slug,
	}

	if len(campaign.CampaignImages) > 0 {
		formatter.ImageUrl = campaign.CampaignImages[0].FileName
	}

	// buat variabel utk menampung slice of string
	var perks []string

	// looping setiap karakter yang sudah dipisan
	for _, perk := range strings.Split(campaign.Perks, ",") {
		// msukan ke dalam slice
		perks = append(perks, perk)
	}

	formatter.Perks = perks

	// user formatter
	user := campaign.User

	campaignUserFormatter := CampaignUserFormatter{}
	campaignUserFormatter.Name = user.Name
	campaignUserFormatter.ImageUrl = user.AvatarFileName

	// simpan ke detailFormatter
	formatter.User = campaignUserFormatter

	return formatter
}
