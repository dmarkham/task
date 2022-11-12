# Task

Create a production ready web service which combines two existing web services. Fetch a random name from https://names.mcquay.me/api/v0 Fetch a random Chuck Norris joke from http://joke.loc8u.com:8888/joke?limitTo=nerdy&firstName=John&lastName=Doe Combine the results and return them to the user.

# Demo

# Usage Instructions
## These instructions assume some familiarity with git, cli, and go

1. Clone the repository [(instructions)](https://docs.github.com/en/repositories/creating-and-managing-repositories/cloning-a-repository)
2. There are two options to spin up the server:
   * Build - 
        1. Open the project directory in a command line interface (CLI)
        2. Enter `go build`, this will create an executable
        3. Double click the executable, this will instantiate the server
   * Run
        1. Open the project directory in a command line interface (CLI) 
        2. Enter `go run task.go const.go handlers.go joke.go name.go`, this will instantiate the server
3. Open a web browser or any other tool that can send and receive get requests
4. Send a Get request to this address `http://localhost:8080/`
5. A response similar to the image below should appear:
