// Sélection de la photo de profil
var selectElement = document.getElementById('photo_profil_select')
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
            throw new Error('Échec de la requête de mise à jour du profil.')
        }
    })

    .then(function(response) {
        if (response.ok) {
            // optenir données JSON du path de pdp actuelle
            return response.json()
        } else {
            throw new Error('Échec de la récupération du chemin de la photo de profil actuelle.')
        }
    })

    .then(function(data) {
        // Mettre à jour le chemin de la photo de profil dans le code HTML
        var imgPath = data.PhotoProfil
        var imgElement = document.getElementById('photo_profil')
        imgElement.src = imgPath
        imgElement.alt = "Photo de profil"
    })
    
    .catch(function(error) {
        console.error('Une erreur s\'est produite:', error)
    })
})
