function selectDefaultPicture(filename) {
  var newPictureSrc = "pictures/Profil/" + filename;
  // Update 'profile picture'
  var profilePicture = document.getElementById('profile_picture');
  profilePicture.src = newPictureSrc;

  // Update 'uploadInput'
  document.getElementById("selectedPicture").value = filename;
}

function fileUpload(input) {
  if (input.files && input.files[0]) {
      var reader = new FileReader();

      reader.onload = function(e) {
      // MAJ source de img de profil
      document.getElementById('profile_picture').src = e.target.result;
      // Update 'selectedPicture' -> indique que img perso uploadée
      document.getElementById("selectedPicture").value = "";
      };

      reader.readAsDataURL(input.files[0]);
  }
}