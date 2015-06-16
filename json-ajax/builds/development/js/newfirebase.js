var ref = new Firebase("https://popping-torch-8077.firebaseio.com");
var outputNode = document.querySelector('#artists');
var dataNode = document.querySelector('#addArtist');
var outputHTML = '';
var formNode = document.querySelector('#fornDiv');

dataNode.addEventListener('click', function(e) {

  if (e.target.id === 'addArtist') {
    ref.push({
      name: "Richard Tweed",
      reknown: "Vietname Vet, Award Winning Architect",
      shortname: "Richard_Tweed",
      bio: "Art is Richard’s third career. He is a Vietnam veteran who proudly served his country for 15 years after the war. As an adult, Richard returned to college to pursue a degree in Architecture, and was a practicing and award-winning architect for another 30 years. His architectural designs awakened within him a deep longing for painting and drawing, and he began slowly by painting incredibly unique renditions of his successful architectural structures. Seeing peoples’ response to his paintings, brought an excitement that Richard had never known, and upon retiring, Richard enrolled in the Roux Academy of Art, Media, and Design to learn all he could about painting."
    });
  }
});

outputNode.addEventListener('click', function(e) {
    var deleteRef = new Firebase("https://popping-torch-8077.firebaseio.com" + e.target.id);
    var onDeleteError = function(error) {
      if (error) {
        console.log('Synchronization failed');
      } else {
        console.log('Synchronization succeeded');
      }
    };
    deleteRef.remove(onDeleteError);
});

ref.on('value', function(snapshot) {
  var name, reknown, shortname, bio;
  var data = snapshot.val();

  outputHTML = '<ul class="artists">';
  for (var key in data) {
    name = data[key].name;
    reknown = data[key].reknown;
    shortname = data[key].shortname;
    bio = data[key].bio;

    outputHTML += '<li class="item">';
    outputHTML += '<h1 class="name">' + name + ' <button class="deleteArtist" id="' + key + '">-</button></h1>';
    outputHTML += '<h2 class="reknown">' + reknown + '</h2>';
    outputHTML += '<p class="bio"><img src="http://planetoftheweb.com/i/artists/' + shortname + '_tn.jpg" alt="' + name + ' photo">' + bio + '</p>';
    outputHTML += '</li>';
  }
  outputHTML += '</ul>';
  outputNode.innerHTML = outputHTML;
});

