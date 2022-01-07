package main

// func failOnError(err error, msg string) {
// 	if err != nil {
// 		log.Fatalf("%s: %s", msg, err)
// 	}
// }

func (em *EventMediator) AttachEventChannel(ec EventChannel) {
	em.EventChannels = append(em.EventChannels, ec)
}

func (em *EventMediator) AddRecognizedEvent(s string) {
	em.RecognizedEvents = append(em.RecognizedEvents, s)
}
