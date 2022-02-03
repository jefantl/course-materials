
This code dispatches concurrent workers who scan `scanme.nmap.org` to find how many of its ports are open and closed, as well as printing each ports information.

TODO 1: See `scanner.go`.
TODO 2: Used 0.5 second as the timeout value.
TODO 3: Added closedports, the same type as openports. Also added the logic to add closed ports to the closedports array.
TODO 4: I found 10 gave reliable results.
TODO 5: Prints both open and closed ports, the port and state seperated by a comma.

I modified the code so the results channel passed labeled ports (a port and whether its open or not).

The test for the number of open ports also includes tests for the number of closed ports.

Note: I had to move the open/closed ports arrays inside the PortScanner function since the tests were interfering with each other. The arrays do not reset between tests when they are global.
