package main

// func (cfg *apiConfig) dbVideoToSignedVideo(video database.Video) (database.Video, error) {
// 	videoParts := strings.Split(*video.VideoURL, ",")
// 	if len(videoParts) != 2 {
// 		return video, errors.New("invalid video URL")
// 	}

// 	newURL, err := auth.GeneratePresignedURL(cfg.s3Client, videoParts[0], videoParts[1], 15*time.Minute)
// 	if err != nil {
// 		return video, err
// 	}

// 	video.VideoURL = &newURL
// 	return video, nil
// }
