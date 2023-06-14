function selectDefaultPicture(element) {
    // Get the source of the clicked picture
    var newPictureSrc = element.src;
  
    // Update the profile picture
    var profilePicture = document.getElementById('profile_picture');
    profilePicture.src = newPictureSrc;
  
    // Update the selected picture value in the hidden input field
    var selectedPictureInput = document.querySelector('input[name="selectedPicture"]');
    selectedPictureInput.value = newPictureSrc;
  }
  
  // Add event listener to the form submit button
  var sendButton = document.getElementById('sendButton');
  sendButton.addEventListener('click', function(e) {
    // Prevent the default form submission
    e.preventDefault();
  
    // Submit the form
    var profileForm = document.getElementById('profileForm');
    profileForm.submit();
  });

function fileUpload(input) {
  var file = input.files[0];
  var reader = new FileReader();

  reader.onload = function(e) {
    // Get the data URL of the uploaded image
    var newPictureSrc = e.target.result;

    // Update the profile picture
    var profilePicture = document.getElementById('profile_picture');
    profilePicture.src = newPictureSrc;

    // Update the selected picture value in the hidden input field
    var selectedPictureInput = document.querySelector('input[name="selectedPicture"]');
    selectedPictureInput.value = newPictureSrc;
  };

  reader.readAsDataURL(file);
}
  