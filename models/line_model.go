package models

type ChengduRealModel struct{
	Status int64 `json:"status"`
	ResultDes string `json:"resultdes"`
	BeginStation string `json:"beginstation"`
	EndStation string `json:"endstation"`
	BeginTime string `json:"begintime"`
	EndTime string `json:"endtime"`
	BusNumber string `json:"busnumber"`
	LineType int64 `json:"linetype"`
	List []ChengduStationModel `json:"list"`
}

type ChengduStationModel struct{
	Station string `json:"station"`
	GpsNumber int64 `json:"gpsnumber"`
	Lon string `json:"lon"`
	Lat string `json:"lat"`
}

type DbLineModel struct{
	LineName string `json:"line_name"`
	LineId string `json:"line_id"`
}

type LineStationModel struct{
	LineResults0 LineResultsModel `json:"lineResults0"`
	LineResults1 LineResultsModel `json:"lineResults1"`
}

type LineResultsModel struct{
	Direction string `json:"direction"`
	Stops []Stop `json:"stops"`
}

type Stop struct{
	Zdmc string `json:"zdmc"`
	Id string `json:"id"`
}

type LineModel struct{
   Value string `json:"-value"`
   Name string `json:"-name"`
}

type CrawLineSonModel struct{
	Line []LineModel `json:"line"`
	Version string `json:"-version"`
}

type CrawLineModel struct{
	Lines CrawLineSonModel `json:"lines"`
}

type UpdownModel struct{
	EndEarlytime string `json:"end_earlytime"`
	EndLatetime string `json:"end_latetime"`
	EndStop string `json:"end_stop"`
	LineId string `json:"line_id"`
	LineName string `json:"line_name"`
	StartEarlytime string `json:"start_earlytime"`
	StartLatetime string `json:"start_latetime"`
	StartStop string `json:"start_stop"`
}


