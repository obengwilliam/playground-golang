package socialmedia

import (
	"time"
)

//go:generate stringier -type=Moodstate
type MoodState int

const (
	MoodStateSad MoodState = iota
	MoodStateHappy
	MoodStateNeutral
	MoodStateAngry
	MoodStateHopeful
	MoodStateShy
	MoodStateOnCloudNine
	MoodStateThrilled
	MoodStateComical
	MoodStateBored
)

// Embedded type
type AuditableContent struct {
	createdTime time.Time
	updatedTime time.Time
	createdBy   string
	updatedBy   string
}
type Post struct {
	auditableContent AuditableContent
	caption          string
	messageBody      string
	url              string
	keywords         []string
	imageURL         string
	thumbnailURL     string
	likers           []string
	authorMood       MoodState
}

// a container to contain all modes and their values
var Moods map[string]MoodState

// initilize mood state here
func init() {
	Moods = map[string]MoodState{
		"cloudNine": MoodStateOnCloudNine,
		"happy":     MoodStateHappy,
		"bored":     MoodStateAngry,
	}
}

// function for creating a new post
func NewPost(username string, imageURL string, thumbnailURL string, caption string, url string, messageBody string, mood MoodState, keywords []string, likers []string) *Post {

	auditableContent := AuditableContent{createdBy: username, createdTime: time.Now()}

	return &Post{
		auditableContent: auditableContent,
		caption:          caption,
		url:              url,
		imageURL:         imageURL,
		thumbnailURL:     thumbnailURL,
		messageBody:      messageBody,
		authorMood:       mood,
		keywords:         keywords,
		likers:           likers,
	}
}
