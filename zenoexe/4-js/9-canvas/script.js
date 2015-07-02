var canvas = document.querySelector('canvas');
var context = canvas.getContext("2d");


//context.arc(300, 300, 100, 0, Math.PI*2, true);
//context.arc(300, 300, 75, 0, Math.PI, false);
///context.arc(280, 280, 20, 0, Math.PI*2, true);
context.moveTo(520, 280);
context.arc(320, 280, 200, 0, Math.PI*2, true);
context.moveTo(275, 200);
context.arc(250, 200, 25, 0, Math.PI*2, true);
context.moveTo(405, 200);
context.arc(380, 200, 25, 0, Math.PI*2, true);
context.moveTo(495, 280);
context.arc(320, 280, 175, 0, Math.PI, false);


context.stroke();

