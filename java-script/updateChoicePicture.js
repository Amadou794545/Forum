// Sélection de la photo de profil
var selectElement = document.getElementById('profile_picture_choices');
selectElement.addEventListener('click', function(event) {
    if (event.target.tagName === 'IMG') {
        selectProfilePicture(event.target);
    }
});

function selectProfilePicture(element) {
    var selectedPhoto = element.src;

    // Supprimez border de toutes les images
    var images = document.getElementsByTagName('img');
    for (var i = 0; i < images.length; i++) {
        images[i].style.border = 'none';
    }

    // Ajouter border sur l'image sélectionnée
    element.style.border = '2px solid blue';

    // Set the selected picture source in the hidden input field
    document.getElementById('selectedPicture').value = selectedPhoto;

    fetch('/inscriptionPicture', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ pictureToSend: selectedPhoto })
    })
    .then(function(response) {
        if (response.ok) {
            return fetch('/getCurrentPhoto');
        } else {
            throw new Error('Échec de la requête de mise à jour du profil');
        }
    })
    .then(function(response) {
        if (response.ok) {
            return response.json();
        } else {
            throw new Error('Échec de la récupération du chemin de la photo de profil');
        }
    })
    .then(function(data) {
        var imgPath = data.PhotoProfil;
        var imgElement = document.getElementById('profile_picture');
        imgElement.src = imgPath;
        imgElement.alt = "Profile picture";
    })
    .catch(function(error) {
        console.error('Error :', error);
    });
}
