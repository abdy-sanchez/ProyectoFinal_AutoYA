let ButtonLogin = document.querySelector('[id="logeate-a"]');
let ButtonRegister = document.querySelector('[id="registrate-a"]');

let registerForm = document.querySelector('[id="Register-form"]');
let loginForm = document.querySelector('[id="login-form"]');

let badCreds = document.querySelector('[id="bad-creds"]');

let usrRegistered = document.querySelector('[id="user-registered"]');


//Botones de Login! y Registrame!

let logeame = document.querySelector('[id="loginButton"]');

let registrame = document.querySelector('[id="RegisterButton"]');

//Credenciales ingresadas por el usuario

//Register

let registro_nombreUsuario = document.querySelector('[id="Register-form"] [class="input-name"]');

let registro_emailUsuario = document.querySelector('[id="Register-form"] [class="input-email"]');

let registro_passwordUsuario = document.querySelector('[id="Register-form"] [class="input-password"]');

//Login

let login_emailUsuario = document.querySelector('[id="login-form"] [class="input-email"]');

let login_passwordUsuario = document.querySelector('[id="login-form"] [class="input-password"]');



registerForm.style.display = 'none'; //Se oculta inicialmente el formulario de registro de usuario

badCreds.style.display = 'none'; //Se oculta el letrero de malas credenciales

usrRegistered.style.display = 'none';



//Funciones para cambiar entre Login y Registro

ButtonLogin.addEventListener("click",()=>{

   loginForm.style.display = 'block';
   registerForm.style.display = 'none'

})

ButtonRegister.addEventListener("click",()=>{

    registerForm.style.display = 'block';
    loginForm.style.display = 'none';

})


//Funciones para el logeo y registro [capturar y enviar la data]


logeame.addEventListener("click",()=>{

    
    loginData = {

        email : login_emailUsuario.value,
        password : login_passwordUsuario.value

    }

    //console.log(loginData);
    
    LogeoUsuario(loginData);

})

registrame.addEventListener("click",()=>{

    registroData = {

        email : registro_emailUsuario.value,
        nombreUsuario: registro_nombreUsuario.value,
        password : registro_passwordUsuario.value

    }

    //console.log(registroData);

    RegistroUsuario(registroData);
    
})


//Funcion para hacer la peticion de Logeo

async function LoginQuery( loginOBJCT ){


    try {

        const response = await axios.post('http://localhost:8080/login', loginOBJCT);

        return response.data;
        
    } catch (error) {
        
        console.error(`API Error: ${error}`);

    }

};


async function LogeoUsuario(loginOBJCT) {

    const loginAttempt =  await LoginQuery(loginOBJCT);

    if(!loginAttempt){

        badCreds.style.display = 'block';

        Swal.fire({
            title: "Algo salió mal...",
            text: `Verifica las credenciales.`,
            icon: "error"
        });

    }else{

        badCreds.style.display = 'none';


        sessionStorage.setItem('email',loginAttempt.email);
        sessionStorage.setItem('nombreUsuario',loginAttempt.nombreUsuario);
        sessionStorage.setItem('reservas',loginAttempt.reservas);

        Swal.fire({
            title: "Bienvenido,",
            text: `${loginAttempt.nombreUsuario}!`,
            icon: "success"
        }).then((result)=>{

            if (result.isConfirmed) {

                window.location.assign('home.html');

              }

        })

        

        
    }
    
};



//Funcion para hacer la peticion de Registro


async function RegisterQuery( registerOBJCT ){


    try {

        const response = await axios.post('http://localhost:8080/register', registerOBJCT);

        return response.data;
        
    } catch (error) {
        
        console.error(`API Error: ${error}`);

    }

};


async function RegistroUsuario(registerOBJCT) {

    const registerAttempt =  await RegisterQuery(registroData);

    if(registerAttempt){

        usrRegistered.style.display = 'block';      //Se se creó bien, muestra un texto que lo dice

        Swal.fire({
            title: "Usuario correctamente registrado!",
            text: `Ya puedes logearte, ${registerAttempt.nombreUsuario}.`,
            icon: "success"
        });
    }
    
};


