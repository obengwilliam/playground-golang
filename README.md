# Notes on learning golang








`race condition`
This is a race condition, a bug that occurs when the output of our software is dependent on the timing and sequence of events that we have no control over. Because we cannot control exactly when each goroutine writes to the results map, we are vulnerable to two goroutines writing to it at the same time.
