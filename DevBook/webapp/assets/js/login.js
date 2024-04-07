//No formulário login quando a ação submit for executada, chame a função fazerLogin
$('#login').on('submit', fazerLogin);

function fazerLogin(evento){
    evento.preventDefault();//setar o compartamento default

    $.ajax(
        {
            url: "/login",
            method: "POST",
            data: {
                email: $('#email').val(),
                senha: $('#senha').val()
            }   
    }).done(function(){
        window.location = "/home";
    }).fail(function() {
        Swal.fire(
            "Ops...",
            "Usuário ou senha incorreta",
            "error"
            );
    });
}