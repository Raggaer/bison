local req = require('http')
local url = require('url')
req.write(url.queryEscape('Testing bison URL module'))