package model

type SourceDetails struct {
	Id            int
	Name          string
	SourceType    string
	SourceConnStr string
	Path          string
	ConnId        int
}
type DestinationDetails struct {
	Id          int
	Name        string
	DestType    string
	DestConnStr string
	Path        string
	ConnId      int
}
type TransformDetails struct {
	Id              int
	Name            string
	SourceDelimiter string
	DestDelimiter   string
}

type TaskDetails struct {
	Id          int
	TaskName    string
	Schedule    string
	Source      SourceDetails
	Destination DestinationDetails
	Transform   TransformDetails
}
