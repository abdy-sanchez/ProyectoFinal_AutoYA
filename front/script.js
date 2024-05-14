let ButtonLogin = document.querySelector('[id="logeate-a"]');
let ButtonRegister = document.querySelector('[id="registrate-a"]');

let registerForm = document.querySelector('[id="Register-form"]');
let loginForm = document.querySelector('[id="login-form"]');

registerForm.style.display = 'none'; //Se oculta inicialmente el formulario de registro de usuario


//Funciones para cambiar entre Login y Registro

ButtonLogin.addEventListener("click",()=>{

   loginForm.style.display = 'block';
   registerForm.style.display = 'none'

})

ButtonRegister.addEventListener("click",()=>{

    registerForm.style.display = 'block';
    loginForm.style.display = 'none';

})