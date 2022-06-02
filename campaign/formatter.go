package campaign

type CampaignsFormatter struct {
	ID int `json:"id"`
	Title string `json:"title"`
	ShortDescription string `json:"short_description"`
	ImageUrl string `json:"image_url"`
	GoalAmount int `json:"goal_amount"`
	CurrentAmount int `json:"current_amount"`
	UserID int `json:"user_id"`
	Slug string `json:"slug"`
}

func FormatCampaign(campaign Campaign) CampaignsFormatter {
	formatter := CampaignsFormatter{
		ID: campaign.UserID,
		Title: campaign.Name,
		ShortDescription: campaign.ShortDescription,
		GoalAmount: campaign.GoalAmount,
		CurrentAmount: campaign.CurrentAmount,
		Slug: campaign.Slug,
		ImageUrl: "",
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