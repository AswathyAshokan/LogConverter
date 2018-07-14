package model

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)
type LogDetails struct {
	LogTime 	string `json:"logTime" bson:"logTime"`
	LogMessage string  `json:"logMessage" bson:"logMessage"`
	FileName   string `json:"fileName" bson:"fileName"`
	LogFormat  string `json:"logFormat" bson:"logFormat"`
}

func (log  LogDetails )InsertIntoDb(LogData[][]string)bool{
	session,err:=mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("LogDetails").C("logData")
    for i:=0;i<len(LogData);i++{

    	logResult :=[]LogDetails{}
    	log.LogTime=LogData[i][0]
    	log.LogMessage=LogData[i][1]
		log.FileName=LogData[i][2]
		log.LogFormat=LogData[i][3]
		err = c.Find(bson.M{"log_time": LogData[i][0]}).All(&logResult)
		if len(logResult)==0{
			 err = c.Insert(log)
			 if err != nil {
				return false
			}
		}
	}
	return true
}
