package main

import (
	"log"
	"server/db"
)

func newRun() {
	training_run := db.CreateTrainingRun("v0.2 6x64 Random start")
	training_run.Active = true
	training_run.TrainParameters = `["--randomize", "-n"]`
	err := db.GetDB().Save(&training_run).Error
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	db.Init(true)

  /*
	training_run := db.TrainingRun{
		Active: true,
	}
	err := db.GetDB().Where(&training_run).First(&training_run).Error
	if err != nil {
		log.Fatal(err)
	}
	training_run.Active = false
	err = db.GetDB().Save(&training_run).Error
	if err != nil {
		log.Fatal(err)
	}
  */
  newRun()

	defer db.Close()
}
