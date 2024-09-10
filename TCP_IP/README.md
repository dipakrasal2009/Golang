
demonstration

Let’s try it out. In one terminal, we’ll run the server:

IN

    go build -o tcpupperecho ./tcpupperecho.go
    ./tcpupperecho -p 8080 # run the server, listening on port 8080
    ```

OUT:

    tcpupperecho    2023/09/07 10:13:13 listening at localhost: [::]:8080

Let’s run the client in another terminal and send it a message:

    $ go build -o writetcp ./writetcp.go
    $ ./writetcp -p 8080 # run the client, connecting to localhost:8080
    > writetcp 2023/09/07 10:20:32 connected to 127.0.0.1:8080: will forward stdin
    hello
    writetcp        2023/09/07 10:20:49 sent: hello
    HELLO

And checking in back in the server terminal, we see:

tcpupperecho    2023/09/07 10:20:49 received: hello

