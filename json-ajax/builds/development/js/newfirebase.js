var ref = new Firebase("https://popping-torch-8077.firebaseio.com/");
var outputNode = document.querySelector('#artists');
var dataNode = document.querySelector('#addArtist');
var outputHTML = '';
var formNode = document.querySelector('#formDiv');

dataNode.addEventListener('click', function(e) {
  e.preventDefault();
  if (e.target.id === 'addArtist') {
    var myArtist = {};
    myArtist.name = document.querySelector('#nameid').value;
    myArtist.reknown = document.querySelector('#reknownid').value;
    myArtist.shortname = document.querySelector('#shortnameid').value;
    myArtist.bio = document.querySelector('#bioid').value;
    ref.push(myArtist);
  }
});

outputNode.addEventListener('click', function(e) {
    var deleteRef = new Firebase("https://popping-torch-8077.firebaseio.com/" + e.target.id);
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
    outputHTML += '<p class="bio">' + bio + '</p>';
    outputHTML += '</li>';
  }
  outputHTML += '</ul>';
  outputNode.innerHTML = outputHTML;
});

