package models

type Results struct {
	Results []Result
}

type Result struct {
	Track string
	TrackResults []TrackResult
}

type TrackResult struct {
	Name string
	Rank string
	Time string
}
