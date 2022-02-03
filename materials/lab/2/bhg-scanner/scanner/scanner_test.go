package scanner

import (
	"testing"
)

// THESE TESTS ARE LIKELY TO FAIL IF YOU DO NOT CHANGE HOW the worker connects (e.g., you should use DialTimeout)
func TestOpenPort(t *testing.T) {

	openCount, closedCount := PortScanner() // Currently function returns only number of open ports
	wantOpen := 2                           // default value when passing in 1024 TO scanme; also only works because currently PortScanner only returns
	wantClosed := 1022
	//consider what would happen if you parameterize the portscanner address and ports to scan

	if openCount != wantOpen {
		t.Errorf("got %d open ports, wanted %d", openCount, wantOpen)
	}
	if closedCount != wantClosed {
		t.Errorf("got %d closed ports, wanted %d", closedCount, wantClosed)
	}
}

func TestTotalPortsScanned(t *testing.T) {
	// THIS TEST WILL FAIL - YOU MUST MODIFY THE OUTPUT OF PortScanner()

	openCount, closedCount := PortScanner() // Currently function returns only number of open ports
	total := openCount + closedCount
	want := 1024 // default value; consider what would happen if you parameterize the portscanner ports to scan

	if total != want {
		t.Errorf("got %d, wanted %d", total, want)
	}
}
