window.addEventListener('DOMContentLoaded', () => {
    const profileIcon = document.querySelector('.profile-icon');
    const userSection = document.getElementById('user-section');

    // Vérifier la présence du cookie de session
    const sessionCookie = document.cookie.split(';').find(cookie => cookie.trim().startsWith('session'));

    if (sessionCookie) { // Utilisateur connecté
        const profileImage = document.createElement('img');
        profileImage.src = "picture/Profil/anonyme.jpg"; //TODO get user img by API
        profileIcon.appendChild(profileImage);
    } else { // Utilisateur non connecté
        const registerButton = document.createElement('a');
        registerButton.href = '/inscription';
        registerButton.classList.add('button');
        registerButton.textContent = "S'inscrire";
        userSection.appendChild(registerButton);

        const loginButton = document.createElement('a');
        loginButton.href = '/login';
        loginButton.classList.add('button');
        loginButton.textContent = 'Se connecter';
        userSection.appendChild(loginButton);
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
