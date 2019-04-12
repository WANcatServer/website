var lastId = "i";
function popUp(newId) {
    document.getElementById(lastId).style.display = "none";
    if (newId !== lastId) {
        document.getElementById(newId).style.display = "block";
    }
    if (lastId == newId) { lastId = "i"; } else { lastId = newId; }
}
window.onclick = function (e) {
    if (!e.target.matches(".popUp") && !e.target.matches(".animate *") && !e.target.matches(".animate")) {
        document.getElementById(lastId).style.display = "none";
        lastId = "i"; 
    }
}
function validate() {
    if (document.getElementById("psw1").value == document.getElementById("psw2").value) {
        document.getElementById("match").innerHTML = "密碼相符";
        document.getElementById("signUp-submit").disabled = false;
    }
    else {
        document.getElementById("match").innerHTML = "密碼不相符";
        document.getElementById("signUp-submit").disabled = true;
    }
}

function logIn() {
    document.getElementById("logged-out").style.display = "none";
    document.getElementById("logged-in").style.display = "block";    
}

function logOut() {
    document.getElementById("logged-in").style.display = "none";
    document.getElementById("logged-out").style.display = "block";
}