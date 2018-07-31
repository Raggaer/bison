# bison

Build websites using lua. 
**bison** uses [gopher-lua](https://github.com/tul/gopher-lua) to parse http requests and [Go](https://golang.org/) for the http server and the templating
This provides an easy way to create websites (using a simple language like **lua**) with great tools from a more complex language (templating, logging)

## Router

`bison` uses a `router.lua` file (should be located at `app/router`). The router file follows this simple structure:

```lua
local router = {
  ['/test/:name'] = {
    get = 'test.lua',
    post = 'post_test.lua'
  }
}

-- Always return the router table
return router
```

You first specify a route (you can use named parammeters like `:name` on the example), 
each route has a table with the http method and the controller (lua files inside the `controllers` directory)
that should be executed

## Testing

All the testing is done on the `test` folder, where a simple example of how `bison` works is located. Every lua module is tested, on each test a fasthttp server is created with a random port
Then a request is simulated using `net/http` package