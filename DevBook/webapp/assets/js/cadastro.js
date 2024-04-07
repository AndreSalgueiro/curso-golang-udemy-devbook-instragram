//quando acontecer o evento identificado como 'submit' no 'formulario-cadastro' chame afunção criarusuario
$('#formulario-cadastro').on('submit', criarUsuario);//esse é um comando para o jquery

function criarUsuario(evento) {
    evento.preventDefault();//retira o comportamento padrão da página
    console.log("Dentro da função usuário");

    if($('#senha').val() != $('#confirmar-senha').val()) {
        Swal.fire(
            "Ops...",
            "As senhas não coincidem",
            "error"
            );
        return;
    }

    //fazendo a requisição para uma rota configurada dentro do WebApp passando os dados do formulário
    $.ajax(
        {
            url: "/usuarios",
            method: "POST",
            data: {
                nome: $('#nome').val(),
                email: $('#email').val(),
                nick: $('#nick').val(),
                senha: $('#senha').val()
            }
            //função anonima que será executada caso a requisição retorne um status de sucesso
        }).done(function() {
            Swal.fire(
                "Sucesso!",
                "Usuário cadastrado com sucesso!",
                "success"
                ).then(function() {
                    $.ajax(
                        {
                            url: "/login",
                            method: "POST",
                            data: {
                                email: $('#email').val(),
                                senha: $('#senha').val()
                            }
                        }).done(function() {
                           window.location = "/home"; 
                        }).fail(function() {
                            Swal.fire(
                                "Ops...",
                                "Erro ao autenticar o usuário!",
                                "error"
                                );
                        })
                })
            //função anônima que será executada caso a requisição retorne um status de falha
        }).fail(function(erro) {
            Swal.fire(
                "Ops...",
                "Erro ao cadastrar o usuário!",
                "error"
                );
        });
}