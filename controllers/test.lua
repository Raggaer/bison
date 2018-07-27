local c = require('config')
local req = require('http')
print(c.get('myCustomKey'))
print(req.param('name'))