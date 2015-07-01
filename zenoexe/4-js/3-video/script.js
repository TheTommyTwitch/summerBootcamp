var play = document.querySelector('.play');
var pause = document.querySelector('.pause');
var up = document.querySelector('.volumeUp');
var down = document.querySelector('.volumeDown');
var down = document.querySelector('.replay');
var video = document.querySelector('video');
video.volume = .5;

video.addEventListener('timeupdate', function() {
    currentTimeDisplay = document.createElement('p');
    durationDisplay = document.createElement('p');
	currentTimeDisplay.innerHTMl = (video.timeupdate)
	durationDisplay.innerHTML += (video.duration);
	document.querySelector('.container').appendChild(durationDisplay);
	document.querySelector('.container').appendChild(currentTimeDisplay);
});



play.addEventListener('click', function(e){
	video.play();
});

pause.addEventListener('click', function(e){
	video.pause();
});

up.addEventListener('click', function(e){
	if(audio.volume < 1){
	video.volume += .1;
	}
});

down.addEventListener('click', function(e){
	if(audio.volume > 0.1){
	video.volume -= .1;
	}
});

replay.addEventListener('click', function(e){
	video.currentTime = 0;
	video.play();
});

