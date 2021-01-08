# Contest_Notif
A tool built in GO to notify you whenever there is a contest on codeforces.

The tool uses the codeforces API to fetch details of all the contests on codeforces. The tool then filters from this output the list of all upcoming contests and their
starting data and time. If the contest is to be held today it outputs the contest details in a JSON format. 

I have used BASH and slack webhooks to feed the output of this tool to my slack channel where i am notified everyday if there is contest that day on codeforces.

To use this tool use need to have GO installed on your computer.
If GO is installed on your PATH then just write-

**go run fetch.go**
