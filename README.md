# bison

Build websites using lua. 
**Bison** uses [gopher-lua](https://github.com/tul/gopher-lua) to parse http requests and [Go](https://golang.org/) for the http server and the templating
This provides an easy way to create websites (using a simple language like **lua**) with great tools from a more complex language (templating, logging)

## Testing

All the testing is done on the `test` folder, where a simple example of how `bison` works is located. Every lua module is tested, on each test a fasthttp server is created with a random port
Then a request is simulated using `net/http` package