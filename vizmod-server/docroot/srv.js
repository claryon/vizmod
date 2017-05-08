var fs = require('fs');
var http = require('http');

//Lets define a port we want to listen to
const PORT=8080; 

//We need a function which handles requests and send response
function handleRequest(request, response){
    response.setHeader('Connection', 'Transfer-Encoding');
    response.setHeader('Content-Type', 'text/html; charset=utf-8');
    response.setHeader('Transfer-Encoding', 'chunked');

    data = fs.readFileSync('./test.html', 'utf8');
    response.write(data);

    var length = 1000;
    for(var i = 0; i < length; i++) {
         
        response.write(content());
        sleep(200);
    }
        
    setTimeout(function() {
        response.write(' world!');
        response.end();
    }, 1000);

        //response.write('It Works!! Path Hit: ' + request.url);
}

//Create a server
var server = http.createServer(handleRequest);

//Lets start our server
server.listen(PORT, function(){
    //Callback triggered when server is successfully listening. Hurray!
    console.log("Server listening on: http://localhost:%s", PORT);
});

function content() {
        var r = getRandomInt(0,255);
        var g = getRandomInt(0,255);
        var b = getRandomInt(0,255);
        return "X <script>set_color("+r+", "+g+", "+b+");</script>";

}

// Gibt eine Zufallszahl zwischen min (inklusive) und max (inklusive) zurück 
// Die Verwendung von Math.round() erzeugt keine gleichmäßige Verteilung! 
function getRandomInt(min, max) {
  min = Math.ceil(min);
  max = Math.floor(max);
  return Math.floor(Math.random() * (max - min +1)) + min; 
}

function sleep(milliseconds) {
  var start = new Date().getTime();
  for (var i = 0; i < 1e7; i++) {
    if ((new Date().getTime() - start) > milliseconds){
      break;
    }
  }
}
