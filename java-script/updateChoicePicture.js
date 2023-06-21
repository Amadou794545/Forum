function selectDefaultPicture(filename) {
  var newPictureSrc = "Pictures/Profil/" + filename;
  // Update 'profile picture'
  var profilePicture = document.getElementById('profile_picture');
  profilePicture.src = newPictureSrc;

  // Update 'uploadInput'
  document.getElementById("selectedPicture").value = filename;
}
