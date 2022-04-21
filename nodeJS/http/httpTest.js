const { on } = require("events");
var http=require("http");

http.createServer(a).listen(8888);
function a(request ,response){

response.end("hello_mxd111111\n");
}
console.log('server runningg at http://127.0.0.1:8888')