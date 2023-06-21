window.addEventListener('DOMContentLoaded', () => {
    // Vérifier la présence du cookie de session
    const sessionCookie = document.cookie.split(';').find(cookie => cookie.trim().startsWith('session'));

    if (sessionCookie) { // Utilisateur connecté
        const unconnectedSection = document.querySelector('.unconnected');
        unconnectedSection.style.display = 'none';
    
        const connectedSection = document.querySelector('.connected');
        connectedSection.style.display = 'block';
    } else { // Utilisateur non connecté
        const connectedSection = document.querySelector('.connected');
        connectedSection.style.display = 'none';
    
        const unconnectedSection = document.querySelector('.unconnected');
        unconnectedSection.style.display = 'block';
    }    
});

function toggleMenu() {
    var menu = document.getElementById('menu');

    if(menu.style.display == "block") { // si visible, cacher menu
        menu.style.display = "none";
    } else { // si caché, montrer menu
        menu.style.display = "block";
    }
}
