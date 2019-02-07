# bison

Build websites using lua. 
**bison** uses [gopher-lua](https://github.com/yuin/gopher-lua) to parse http requests and [Go](https://golang.org/) for the http server and the templating
This provides an easy way to create websites (using a simple language like **lua**) with great tools from a more complex language (templating, logging)

## Config

To handle all your configuration values `bison` uses a `config.lua` file (should be located at `app/config`). This file is a simple lua table:

```lua
local config = {
  address = ':0',
  devMode = false,
  myTestKey = 'testing-bison'
}

-- Always return the config table
return config
```
There are some **mandatory** fields that need to be present on your configuration file:

- `address`: The address where the http server will listen on
- `devMode`: Development mode boolean, if enabled some features (like hot-reload lua files) will be activated, to make development easier

After the mandatory fields are all setup you can declare any field (even use tables) and access them later using the `config` module 
on your lua files

## Router

`bison` uses a `router.lua` file (should be located at `app/router`) to handle all the request routing. The router file follows this simple structure:

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

All the testing is done on the `test` folder, where a simple example of how `bison` works is located. Every lua module is tested, on each test a fasthttp server is created with a random port.
Then a request is simulated using `net/http` package

Since the testing package is basically a `bison` application, it can be used as a base/example on how should your directory structure look like
