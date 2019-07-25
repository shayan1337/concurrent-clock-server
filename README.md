**CONCURRENT CLOCK SERVER**

This is a TCP server that accepts connections on port 8000 and prints the time every second to the connection. 


The server handles multiple requests at a time using Goroutines.


**Usage** 

~~~~
go run main.go

./main
~~~~


**NOTE:**

I haven't created any script for the client. I have used _netcat_ tool for making TCP connections to my server.