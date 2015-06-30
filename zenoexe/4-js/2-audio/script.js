var play = document.querySelector('.play');
var pause = document.querySelector('.pause');
var up = document.querySelector('.volumeUp');
var down = document.querySelector('.volumeDown');
var audio = document.querySelector('audio');
audio.volume = .5;

play.addEventListener('click', function(e){
	audio.play();
});

pause.addEventListener('click', function(e){
	audio.pause();
});

up.addEventListener('click', function(e){
	if(audio.volume < 1){
	audio.volume += .1;
	}
});

down.addEventListener('click', function(e){
	if(audio.volume > 0.1){
	audio.volume -= .1;
	}
});