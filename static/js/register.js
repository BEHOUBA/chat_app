let regName = document.getElementById("reg-name");
let regEmail = document.getElementById("reg-email");
let regPass1 = document.getElementById("reg-password1");
let regPass2 = document.getElementById("reg-password2");
let info = document.getElementById("info");








// sendRegistration check if form is in valid form, then make an object from user's
// inputed data and then send this data to server with ajax request
function sendRegistration(){
    if ((regPass1.value !== regPass2.value) && regPass1.value !== ""){
        alert("password didn't match")
        return
    }
    var regData = {
        "name": regName.value,
        "email": regEmail.value,
        "password": regPass1.value
    }

    axios.post("/register", regData)
        .then(res => {
            if(res.status === 200){
                location.href = "/"
            }
        })
        .catch(err => {
            switch (err.response.status){
                case 400:
                info.innerText = "Donnee fourni incorrect verifiez que vous avez fournis des informations correctes";
                break;
                case 409:
                info.innerText = "Compte existe deja merci de vous connecter";
                break;
                default:
                info.innerText = "error interne serveur"
            }
        })
}


