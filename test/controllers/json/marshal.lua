local req = require('http')
local json = require('json')
local data = {
	author = 'Raggaer',
	package = 'bison'
}
req.write(json.marshal(data))
