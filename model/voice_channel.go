package model

type VoiceChannel struct {
	Title      string `json:"title" validate:"required"`
	Key        string `json:"key" validate:"required,min=2"`
	BgColor    string `json:"bgColor" validate:"required"`
	IfSingle   int    `json:"ifSingle" validate:"required"`
	IfTTS      int    `json:"ifTTS" validate:"required"`
	TTSRepeats int    `json:"TTSRepeats" validate:"required"`
	MaxHistory int    `json:"maxHistory" validate:"required"`
}
