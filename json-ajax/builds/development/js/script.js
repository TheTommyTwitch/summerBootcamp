  "use strict";
  var request;
  request=new XMLHttpRequest();
  request.open('GET', 'js/data.json');
  request.onreadystatechange = function() {
    if ((request.status === 200) &&
      (request.readyState === 4)) {
        var data = JSON.parse(request.responseText);
        var template = document.querySelector('#speakerstpl').innerHTML;
        var html = Mustache.to_html(template, data);
        document.querySelector('#speakers').innerHTML=html;
    } //ready
  }; //event
  request.send();

  var socialmediarequest;
  socialmediarequest=new XMLHttpRequest();
  socialmediarequest.open('GET', 'js/socialmediarequest.json');
  socialmediarequest.onreadystatechange = function() {
    if ((socialmediarequest.status === 200) &&
      (socialmediarequest.readyState === 4)) {
        var data = JSON.parse(socialmediarequest.responseText);
        var template = document.querySelector('#socialmediatpl').innerHTML;
        var sm = Mustache.to_html(template, data);
        document.querySelector('#socialmedia').innerHTML=sm;
    } //ready
  }; //event
  socialmediarequest.send();