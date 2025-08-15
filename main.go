package main

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

type RadioInfo struct {
	VideoURL        string `json:"videoUrl"`
	AudioURL        string `json:"audioUrl"`
	AudioMountPoint string `json:"audioMountPoint"`
	StartTime       string `json:"startTime"`
}

var (
	port            = String("PORT", ":8080")
	videoURL        = String("RADIO_VIDEO_URL", "https://restreamer.gewis.nl/memfs/b0d27314-d7cc-477b-9c5a-f22ef7c7ad32.m3u8")
	audioURL        = String("RADIO_AUDIO_URL", "https://radio-icecast.gewis.nl")
	audioMountPoint = String("RADIO_AUDIO_MOUNT_POINT", "/stream")
	radioStartTime  = String("RADIO_START_TIME", "2025-08-18T07:00:00Z")
	token           = String("RADIO_GEWIS_TOKEN", "gewis-radio")
)

func main() {
	chat := NewChat()

	http.HandleFunc("/ws", chat.HandleWS)

	http.HandleFunc("/api/v1/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})

	http.HandleFunc("/api/v1/token", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(token)
	})

	http.HandleFunc("/api/v1/radio", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(RadioInfo{
			VideoURL:        videoURL,
			AudioURL:        audioURL,
			AudioMountPoint: audioMountPoint,
			StartTime:       radioStartTime,
		})
	})

	log.Info().Str("port", port).Msg("Starting server")
	log.Fatal().Err(http.ListenAndServe(port, nil))
}
