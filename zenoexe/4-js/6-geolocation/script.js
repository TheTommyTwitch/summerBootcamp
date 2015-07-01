navigator.geolocation.getCurrentPosition(function(e){
  lat = e.coords.latitude;
  longi = e.coords.longitude;

  var img = new Image();
  img.alt = 'My map'
  img.src = 'https://maps.google.com/maps/api/staticmap?markers='+ lat +','+ longi + '&zoom=10&size=700x700';
  document.body.appendChild(img);
});




