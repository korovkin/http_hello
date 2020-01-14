console.log("Hello");

$(document).ready(function () {
    console.log("Hello");
    now = new Date();
    $('#js').append($('<p>').append('' + now.toLocaleString()));
})