// Support TLS-specific URLs, when appropriate.
if (window.location.protocol == "https:") {
  var ws_scheme = "wss://";
} else {
  var ws_scheme = "ws://"
};


var inbox = new ReconnectingWebSocket(ws_scheme + location.host + "/receive");
/*var outbox = new ReconnectingWebSocket(ws_scheme + location.host + "/submit");*/

var initial = false;
console.log(initial);

$(document).ready(function() {
  inbox.onmessage = function(message) {

    console.log(message);
    var data = JSON.parse(message.data);

    if (!initial) {
      initiate(data);
      initial = true;
      console.log('init-update');

    };

    if (initial) {
      update_css(data);
      console.log('update');
    };
    console.log(initial);
  };


  inbox.onclose = function(error) {
    console.log(error.code);
    this.inbox = new WebSocket(inbox.url);
  };
});

/*
$("#chat-text").append("<div class='panel panel-default'><div class='panel-heading'>" + $('<span/>').text(data.handle).html() + "</div><div class='panel-body'>" + $('<span/>').text(data.text).html() + "</div></div>");
$("#chat-text").stop().animate({
  scrollTop: $('#chat-text')[0].scrollHeight
}, 800);*/
/*outbox.onclose = function(){
    console.log('outbox closed');
    this.outbox = new WebSocket(outbox.url);
};

$("#input-form").on("submit", function(event) {
  event.preventDefault();
  var handle = $("#input-handle")[0].value;
  var text   = $("#input-text")[0].value;
  outbox.send(JSON.stringify({ handle: handle, text: text }));
  $("#input-text")[0].value = "";
});*/
