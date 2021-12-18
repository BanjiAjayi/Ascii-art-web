# Ascii-Art-Web

ASCII-art-web is a project that creates and maintains a server, in which it's possible for the user to use a web GUI (graphical user interface). to make text string into ASCII art.

## running the project

1. From your teminal going to the ascii-art-web dir.

2. then type `go run .` this will start the server on port 8080

3. in a web browser go to localhost:8080

## Implementation details: algorithm

By using the Package http provides us with a HTTP client and server implementations.

1. The web request is first received at the http.ListenAndServe function inside ascii-art-web.go
2. Then the request is passed to the pathHandler for routing to different endpoints, based on the request's URL Path
3. The initial request should be directed to the homeHandler "/", with the method GET, and it displays the HTML form to the client
4. After the client fills in the form, it is send back to the server through the method POST, which is handled by the asciiArtHandler 
5. The filled in info (the text string and the chosen banner style) is parsed and passed into the ascii-art-generator, which generates a file with the lines of the ASCII art. Note that any previous artwork.txt file will be deleted
6. The generated file is then read line-by-line back into the code, and is displayed to the client through a template


## Authors 

David, Burak, Banji

