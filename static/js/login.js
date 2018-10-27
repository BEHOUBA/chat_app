let logForm = document.getElementById("login");
let logEmail = document.getElementById("log-email");
let logPassword = document.getElementById("log-password");
let info = document.getElementById("info");




logForm.addEventListener('submit', e => {
    e.preventDefault();
    sendLogin();
});


function sendLogin() {
    if(logEmail.value === "" || logPassword.value === ""){
        alert("invalid input !!!");
        return;
    }
    var logData = {
        "email": logEmail.value,
        "password": logPassword.value
    }
    console.log(logData)

    axios.post("/login", logData)
        .then(res => {
            console.log(res)
            location.href = "/"
        })
        .catch(err => {
            console.log(err)
            switch (err.response.status){
                case 401:
                info.innerText = "Ce compte n'existe pas !";
                break;
                case 400:
                info.innerText = "Mot de passe incorrect !";
                break;
                default:
                info.innerText = "error interne serveur"
            }
        })
}