local req = require('http')
local url = require('url')
req.write(url.pathUnescape('Testing%20bison%20URL%20module'))