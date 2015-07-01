var play = document.querySelector('.play');
var pause = document.querySelector('.pause');
var up = document.querySelector('.volumeUp');
var down = document.querySelector('.volumeDown');
var replay = document.querySelector('.replay');
var video = document.querySelector('video');
var time = document.querySelector('.time');
video.volume = .5;

	(function () {
        var currentTime = 0;
        var durationTime = 0;
        video.addEventListener('loadedmetadata', function () {
            durationTime = parseInt(video.duration, 10);
            console.log(durationTime);
        });
        video.addEventListener('timeupdate', function () {
            currentTime = parseInt(video.currentTime, 10);
            time.textContent = currentTime + ' / ' + durationTime;
        });
    }());



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

