// Get the text field that we're going to track
 var field = document.querySelector("#field");
 
 // Listen for changes in the text field
 
 field.addEventListener("input", function(e) {
    // And save the results into the session storage object
    localStorage.setItem("autosave", field.value);
 });


// See if we have an autosave value
// (this will only happen if the page is accidentally refreshed)
 
 if (localStorage.getItem("autosave")) {
    // Restore the contents of the text field
    field.value = localStorage.getItem("autosave");
 }