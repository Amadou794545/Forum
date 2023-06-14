
$(document).ready(function() {
    // Écouteur d'événement pour les modifications des cases à cocher
    $('#filters input[type="checkbox"]').on('change', function() {
        var selectedFilters = [];
        // Parcourir les cases à cocher et récupérer les filtres sélectionnés
        $('#filters input[type="checkbox"]:checked').each(function() {
            selectedFilters.push($(this).val());
        });
        // Cacher tous les éléments de post par défaut
        $('#container .post').hide();
        // Afficher les éléments correspondant aux filtres sélectionnés
        if (selectedFilters.length > 0) {
            // Afficher les éléments correspondant aux filtres sélectionnés
            for (var i = 0; i < selectedFilters.length; i++) {
                var filter = selectedFilters[i];
                $('#container .post.' + filter).show();
            }
        } else {
            // Afficher tous les éléments si aucun filtre n'est sélectionné
            $('#container .post').show();
        }
    });
});
