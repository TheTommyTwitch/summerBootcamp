navigator.getUserMedia = navigator.getUserMedia || navigator.webkitGetUserMedia;
video = document.querySelector('video');

var options = {
	audio: false,
	video: true
};

navigator.getUserMedia(options, success, error);

function success(result){
	console.log(result);
	video.src = window.URL.createObjectURL(result);
    video.play();
};
function error(result){
	console.log(result);
};