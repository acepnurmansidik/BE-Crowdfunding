package campaign

import (
	"strings"
)

type CampaignsFormatter struct {
	ID               int    `json:"id"`
	Name            string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageUrl         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	UserID           int    `json:"user_id"`
	Slug             string `json:"slug"`
}

type CampaignDetailFormatter struct {
	ID               int                   `json:"id"`
	Name            string                `json:"name"`
	ShortDescription string                `json:"short_description"`
	ImageUrl         string                `json:"image_url"`
	GoalAmount       int                   `json:"goal_amount"`
	CurrentAmount    int                   `json:"current_amount"`
	UserID           int                   `json:"user_id"`
	Slug             string                `json:"slug"`
	Description 	string				 `json:"description"`
	Perks            []string              `json:"perks"`
	User             CampaignUserFormatter `json:"user"`
	Images []CampaignImageFormatter `json:"images"`
}

type CampaignUserFormatter struct {
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

type CampaignImageFormatter struct {
	ImageUrl string `json:"image_url"`
	IsPrimary bool `json:"is_primary"`
}

func FormatCampaign(campaign Campaign) CampaignsFormatter {
	formatter := CampaignsFormatter{
		ID:               campaign.ID,
		Name:            campaign.Name,
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
	formatter := CampaignDetailFormatter{}
	formatter.ID = campaign.ID
	formatter.Name = campaign.Name
	formatter.ShortDescription = campaign.ShortDescription
	formatter.Description = campaign.Description
	formatter.ImageUrl = ""
	formatter.GoalAmount = campaign.GoalAmount
	formatter.CurrentAmount = campaign.CurrentAmount
	formatter.UserID = campaign.UserID
	formatter.Slug = campaign.Slug

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

	// 
	images := []CampaignImageFormatter{}

	for _, image := range campaign.CampaignImages {
		campaignImageFormatter := CampaignImageFormatter{}
		campaignImageFormatter.ImageUrl = image.FileName

		isPrimary := false

		if image.IsPrimary == 1 {
			isPrimary = true
		}
		campaignImageFormatter.IsPrimary = isPrimary

		images = append(images, campaignImageFormatter)
	}

	formatter.Images = images

	return formatter
}
