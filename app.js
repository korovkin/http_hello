console.log("Hello");

$(document).ready(function () {
    console.log("Hello");
    now = new Date();
    $('#body').append($('<h1>').append('Hello'));
    var js = $('#body').append($('<div class="js" id="js">'));
    js.append($('<p>').append('' + now.toLocaleString()));
})