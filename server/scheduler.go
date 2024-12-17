package main

// type scheduler struct {
// 	running        bool
// 	executing      bool
// 	last_execution time.Time
// 	ticker         *time.Ticker
// 	done           chan bool
// }

// func (s *scheduler) start(work func(t any) error) {
// 	s.running = true
// 	s.ticker = time.NewTicker(2 * time.Second)
// 	s.done = make(chan bool)
// 	go func() {
// 		for {
// 			select {
// 			case <-s.done:
// 				return
// 			case t := <-s.ticker.C:
// 				s.last_execution = t
// 				work(t)
// 			}
// 		}
// 	}()
// }

// func (s *scheduler) stop() {
// 	s.running = false
// 	s.ticker.Stop()
// 	s.done <- true
// }
