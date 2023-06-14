document.forms['register-form'].onsubmit = function(event) {
  // username
  if (this.username.value.trim() === "") {
    document.querySelector(".username-error").innerHTML = "Veuillez saisir un nom d’utilisateur";
    document.querySelector(".username-error").style.display = "block";
    event.preventDefault();
    return false;
  }
  // email
  if (this.email.value.trim() === "") {
    document.querySelector(".email-error").innerHTML = "Veuillez saisir un e-mail";
    document.querySelector(".email-error").style.display = "block";
    event.preventDefault();
    return false;
  }
  if (!this.email.value.includes("@") || !this.email.value.includes(".")) {
      document.querySelector(".email-error").innerHTML = "Veuillez saisir une adresse e-mail valide";
      document.querySelector(".email-error").style.display = "block";
      event.preventDefault();
      return false;
  }
  // password
  var password = this.password.value.trim();
  if (password === "") {
    document.querySelector(".password-error").innerHTML = "Veuillez entrer un mot de passe";
    document.querySelector(".password-error").style.display = "block";
    event.preventDefault();
    return false;
  }
  // Check password requirements
  if (!/\d/.test(password)) {
    document.querySelector(".password-error").innerHTML = "Le mot de passe doit contenir au moins un chiffre";
    document.querySelector(".password-error").style.display = "block";
    event.preventDefault();
    return false;
  }
  if (!/[A-Z]/.test(password)) {
    document.querySelector(".password-error").innerHTML = "Le mot de passe doit contenir au moins une lettre majuscule";
    document.querySelector(".password-error").style.display = "block";
    event.preventDefault();
    return false;
  }
  if (!/[!@#$%^&*.]/.test(password)) {
    document.querySelector(".password-error").innerHTML = "Le mot de passe doit contenir au moins un caractère spécial";
    document.querySelector(".password-error").style.display = "block";
    event.preventDefault();
    return false;
  }
  // confirm password
  var confirmPassword = this['confirm-password'].value.trim();
  if (confirmPassword === "") {
    document.querySelector(".confirm-password-error").innerHTML = "Veuillez réécrire le mot de passe";
    document.querySelector(".confirm-password-error").style.display = "block";
    event.preventDefault();
    return false;
  }
  
  if (password !== confirmPassword) {
    document.querySelector(".confirm-password-error").innerHTML = "Les mots de passe ne correspondent pas";
    document.querySelector(".confirm-password-error").style.display = "block";
    event.preventDefault();
    return false;
  }
  // All data is valid, allow form submission
  return true;
};