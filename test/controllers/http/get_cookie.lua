local req = require('http')
local c = req.getCookie("bison-test-cookie")
req.write(c)