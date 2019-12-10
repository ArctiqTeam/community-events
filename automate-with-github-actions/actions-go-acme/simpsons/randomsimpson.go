package randomsimpson

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Randomsimpsonrequest struct {
	Data string
	Logo string
	Svc  string
}

type randomsimpson struct {
	Data struct {
		Type             string `json:"type"`
		ID               string `json:"id"`
		URL              string `json:"url"`
		Slug             string `json:"slug"`
		BitlyGifURL      string `json:"bitly_gif_url"`
		BitlyURL         string `json:"bitly_url"`
		EmbedURL         string `json:"embed_url"`
		Username         string `json:"username"`
		Source           string `json:"source"`
		Title            string `json:"title"`
		ContentURL       string `json:"content_url"`
		SourceTld        string `json:"source_tld"`
		SourcePostURL    string `json:"source_post_url"`
		IsSticker        int    `json:"is_sticker"`
		ImportDatetime   string `json:"import_datetime"`
		TrendingDatetime string `json:"trending_datetime"`
		Images           struct {
			DownsizedLarge struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				URL    string `json:"url"`
				Width  string `json:"width"`
			} `json:"downsized_large"`
			FixedHeightSmallStill struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				URL    string `json:"url"`
				Width  string `json:"width"`
			} `json:"fixed_height_small_still"`
			Original struct {
				Frames   string `json:"frames"`
				Hash     string `json:"hash"`
				Height   string `json:"height"`
				Mp4      string `json:"mp4"`
				Mp4Size  string `json:"mp4_size"`
				Size     string `json:"size"`
				URL      string `json:"url"`
				Webp     string `json:"webp"`
				WebpSize string `json:"webp_size"`
				Width    string `json:"width"`
			} `json:"original"`
			FixedHeightDownsampled struct {
				Height   string `json:"height"`
				Size     string `json:"size"`
				URL      string `json:"url"`
				Webp     string `json:"webp"`
				WebpSize string `json:"webp_size"`
				Width    string `json:"width"`
			} `json:"fixed_height_downsampled"`
			DownsizedStill struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				URL    string `json:"url"`
				Width  string `json:"width"`
			} `json:"downsized_still"`
			FixedHeightStill struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				URL    string `json:"url"`
				Width  string `json:"width"`
			} `json:"fixed_height_still"`
			DownsizedMedium struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				URL    string `json:"url"`
				Width  string `json:"width"`
			} `json:"downsized_medium"`
			Downsized struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				URL    string `json:"url"`
				Width  string `json:"width"`
			} `json:"downsized"`
			PreviewWebp struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				URL    string `json:"url"`
				Width  string `json:"width"`
			} `json:"preview_webp"`
			OriginalMp4 struct {
				Height  string `json:"height"`
				Mp4     string `json:"mp4"`
				Mp4Size string `json:"mp4_size"`
				Width   string `json:"width"`
			} `json:"original_mp4"`
			FixedHeightSmall struct {
				Height   string `json:"height"`
				Mp4      string `json:"mp4"`
				Mp4Size  string `json:"mp4_size"`
				Size     string `json:"size"`
				URL      string `json:"url"`
				Webp     string `json:"webp"`
				WebpSize string `json:"webp_size"`
				Width    string `json:"width"`
			} `json:"fixed_height_small"`
			FixedHeight struct {
				Height   string `json:"height"`
				Mp4      string `json:"mp4"`
				Mp4Size  string `json:"mp4_size"`
				Size     string `json:"size"`
				URL      string `json:"url"`
				Webp     string `json:"webp"`
				WebpSize string `json:"webp_size"`
				Width    string `json:"width"`
			} `json:"fixed_height"`
			DownsizedSmall struct {
				Height  string `json:"height"`
				Mp4     string `json:"mp4"`
				Mp4Size string `json:"mp4_size"`
				Width   string `json:"width"`
			} `json:"downsized_small"`
			Preview struct {
				Height  string `json:"height"`
				Mp4     string `json:"mp4"`
				Mp4Size string `json:"mp4_size"`
				Width   string `json:"width"`
			} `json:"preview"`
			FixedWidthDownsampled struct {
				Height   string `json:"height"`
				Size     string `json:"size"`
				URL      string `json:"url"`
				Webp     string `json:"webp"`
				WebpSize string `json:"webp_size"`
				Width    string `json:"width"`
			} `json:"fixed_width_downsampled"`
			FixedWidthSmallStill struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				URL    string `json:"url"`
				Width  string `json:"width"`
			} `json:"fixed_width_small_still"`
			FixedWidthSmall struct {
				Height   string `json:"height"`
				Mp4      string `json:"mp4"`
				Mp4Size  string `json:"mp4_size"`
				Size     string `json:"size"`
				URL      string `json:"url"`
				Webp     string `json:"webp"`
				WebpSize string `json:"webp_size"`
				Width    string `json:"width"`
			} `json:"fixed_width_small"`
			OriginalStill struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				URL    string `json:"url"`
				Width  string `json:"width"`
			} `json:"original_still"`
			FixedWidthStill struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				URL    string `json:"url"`
				Width  string `json:"width"`
			} `json:"fixed_width_still"`
			Looping struct {
				Mp4     string `json:"mp4"`
				Mp4Size string `json:"mp4_size"`
			} `json:"looping"`
			FixedWidth struct {
				Height   string `json:"height"`
				Mp4      string `json:"mp4"`
				Mp4Size  string `json:"mp4_size"`
				Size     string `json:"size"`
				URL      string `json:"url"`
				Webp     string `json:"webp"`
				WebpSize string `json:"webp_size"`
				Width    string `json:"width"`
			} `json:"fixed_width"`
			PreviewGif struct {
				Height string `json:"height"`
				Size   string `json:"size"`
				URL    string `json:"url"`
				Width  string `json:"width"`
			} `json:"preview_gif"`
			Four80WStill struct {
				URL    string `json:"url"`
				Width  string `json:"width"`
				Height string `json:"height"`
			} `json:"480w_still"`
		} `json:"images"`
		User struct {
			AvatarURL   string `json:"avatar_url"`
			BannerImage string `json:"banner_image"`
			BannerURL   string `json:"banner_url"`
			ProfileURL  string `json:"profile_url"`
			Username    string `json:"username"`
			DisplayName string `json:"display_name"`
			IsVerified  bool   `json:"is_verified"`
		} `json:"user"`
		ImageOriginalURL             string `json:"image_original_url"`
		ImageURL                     string `json:"image_url"`
		ImageMp4URL                  string `json:"image_mp4_url"`
		ImageFrames                  string `json:"image_frames"`
		ImageWidth                   string `json:"image_width"`
		ImageHeight                  string `json:"image_height"`
		FixedHeightDownsampledURL    string `json:"fixed_height_downsampled_url"`
		FixedHeightDownsampledWidth  string `json:"fixed_height_downsampled_width"`
		FixedHeightDownsampledHeight string `json:"fixed_height_downsampled_height"`
		FixedWidthDownsampledURL     string `json:"fixed_width_downsampled_url"`
		FixedWidthDownsampledWidth   string `json:"fixed_width_downsampled_width"`
		FixedWidthDownsampledHeight  string `json:"fixed_width_downsampled_height"`
		FixedHeightSmallURL          string `json:"fixed_height_small_url"`
		FixedHeightSmallStillURL     string `json:"fixed_height_small_still_url"`
		FixedHeightSmallWidth        string `json:"fixed_height_small_width"`
		FixedHeightSmallHeight       string `json:"fixed_height_small_height"`
		FixedWidthSmallURL           string `json:"fixed_width_small_url"`
		FixedWidthSmallStillURL      string `json:"fixed_width_small_still_url"`
		FixedWidthSmallWidth         string `json:"fixed_width_small_width"`
		FixedWidthSmallHeight        string `json:"fixed_width_small_height"`
		Caption                      string `json:"caption"`
	} `json:"data"`
	Meta struct {
		Status     int    `json:"status"`
		Msg        string `json:"msg"`
		ResponseID string `json:"response_id"`
	} `json:"meta"`
}

func (mybeer *Randomsimpsonrequest) GetData() {
	//this function will return a dad joke
	req, err := http.NewRequest("GET", "http://api.giphy.com/v1/gifs/random?tag=simpsonr&api_key=4Z7XmRZCziCDF0q6rLwAKO7BjgGuF74u", nil)
	if err != nil {
		// handle err
	}
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	simp := randomsimpson{}

	jsonErr := json.Unmarshal(body, &simp)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	defer resp.Body.Close()
	yoursimp := simp.Data.Images.Original.URL
	mybeer.Data = yoursimp
	mybeer.Logo = "https://media2.giphy.com/media/3o6gbbuLW76jkt8vIc/source.gif"
	mybeer.Svc = "Simpson Giphy"

}
