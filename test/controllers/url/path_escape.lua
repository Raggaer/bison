local req = require('http')
local url = require('url')
req.write(url.pathEscape('Testing bison URL module'))