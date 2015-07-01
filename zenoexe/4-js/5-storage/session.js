// Get the text field that we're going to track
 var ol = document.querySelector("ol");
 var clear = document.querySelector('#clear');
 
 // Listen for changes in the text field
 
 ol.addEventListener("input", function(e) {
    // And save the results into the session storage object
    sessionStorage.setItem("autosave", ol.innerHTML);
 });
 
 clear.addEventListener("click", function(e){
 	ol.value = null;
 });


// See if we have an autosave value
// (this will only happen if the page is accidentally refreshed)
 
 if (sessionStorage.getItem("autosave")) {
    // Restore the contents of the text field
    ol.value = sessionStorage.getItem("autosave");
 }