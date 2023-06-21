

// Toggle the checkbox state when the label is clicked
var labels = document.querySelectorAll('.checkbox-label');
labels.forEach(function (label) {
label.addEventListener('click', function () {
    var checkbox = this.previousElementSibling;
    checkbox.checked = !checkbox.checked;
    });
});