local req = require('http')
local json = require('json')
local data = json.unmarshal("{\"author\":\"Raggaer\",\"package\":\"bison\"}")
req.write("Author is " .. data.author .. " and package is " .. data.package)
