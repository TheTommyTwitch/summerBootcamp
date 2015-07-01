var imgs = document.querySelectorAll('img');
var boxes = document.querySelectorAll('.box');

var draggedElement;

for (var i = 0; i < imgs.length; i++) {
	imgs[i].addEventListener('dragstart', function(e) {
		draggedElement = e.target;
	});
};

for (var i = 0; i < boxes.length; i++) {
	boxes[i].addEventListener('drop', function(e) {
		draggedElement.remove();
		e.target.appendChild(draggedElement);
	});

	boxes[i].addEventListener('dragover', function(e) {
		e.preventDefault();
	});
};

