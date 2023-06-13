// Sélection de la photo de profil
var selectElement = document.getElementById('profile_picture_select')
selectElement.addEventListener('change', function() {
    var selectedPhoto = selectElement.value
    
    fetch('/inscriptionPicture', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ photo_profil: selectedPhoto })

    })

    .then(function(response) {
        if (response.ok) {
            //récupérer path img
            return fetch('/getCurrentPhoto')
        } else {
            throw new Error('Échec de la requête de mise à jour du profil')
        }
    })

    .then(function(response) {
        if (response.ok) {
            // optenir données JSON du path de photo de profil actuelle
            return response.json()
        } else {
            throw new Error('Échec de la récupération du chemin de la photo de profil')
        }
    })

    .then(function(data) {
        // Mettre à jour le chemin de la photo de profil dans le code HTML
        var imgPath = data.PhotoProfil
        var imgElement = document.getElementById('profile_picture')
        imgElement.src = imgPath
        imgElement.alt = "Profile picture"
    })
    
    .catch(function(error) {
        console.error('Error :', error)
    })
})
