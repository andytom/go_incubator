var index = 1
var timer = null

function load() {
  timer = setInterval(doUpdate, 1000);
}

function doUpdate() {

  var xhr = new XMLHttpRequest();

  xhr.open('GET', encodeURI('/digit/' + index));

  xhr.onload = function() {
        updateDigit(xhr.responseText);
        index++;
  };

  xhr.send();
}

function updateDigit(newDigit) {
      var piDiv = document.getElementById('pi');
      piDiv.innerHTML = piDiv.innerHTML + newDigit
}
