function toggleMenu() {
    var menu = document.getElementById('menu');

    if(menu.style.display == "block") { // si visible, cacher menu
        menu.style.display = "none";
    } else { // si caché, montrer menu
        menu.style.display = "block";
    }
}