//quando acontecer o evento identificado como 'submit' no formulário 'nova-publicacao' chame afunção criarPublicacao
$('#nova-publicacao').on('submit', criarPublicacao);//esse é um comando para o jquery #identifica um id
//$('.curtir-publicacao').on('click', curtirPublicacao);//esse é um comando para o jquery ponto identifica quando uma classe é chamada
$(document).on('click', '.curtir-publicacao', curtirPublicacao);
$(document).on('click', '.descurtir-publicacao', descurtirPublicacao);

$('#atualizar-publicacao').on('click', atualizarPublicacao);
$('.deletar-publicacao').on('click', deletarPublicacao);


function criarPublicacao(evento) {
    evento.preventDefault();//retira o comportamento padrão da página

    //fazendo a requisição para uma rota configurada dentro do WebApp passando os dados do formulário
    $.ajax(
        {
            url: "/publicacoes",
            method: "POST",
            data: {
                titulo: $('#titulo').val(),
                conteudo: $('#conteudo').val(),
            }
            //função anonima que será executada caso a requisição retorne um status de sucesso
        }).done(function() {
            //se der tudo certo recarrega a página home
            window.location = "/home";
            //função anônima que será executada caso a requisição retorne um status de falha
        }).fail(function(erro) {
            Swal.fire(
                "Ops...",
                "Erro ao criar publicação!",
                "error"
                );
        });
}

function curtirPublicacao(evento) {
    evento.preventDefault();

    const elementoClicado = $(evento.target);//pega o evento de curtida do botão (coração)
    //pega a div que está mais perto do botão de curtir (coração) na hierarquia acima, fazemos isso para conseguir pegar o id da publicação (publicação-id)
    const publicacaoId = elementoClicado.closest('div').data('publicacao-id');
    //console.log(publicacaoId)

    //desabilita o botão de curtidas enquanto a requisição esta em andameno para evitar de enviar várias requisições ao mesmo tempo evirando que elas se percam
    elementoClicado.prop('disabled', true);

    $.ajax({
        url: `/publicacoes/${publicacaoId}/curtir`,
        method: "POST"
    }).done(function(){
        //pega o numero de curtidas
        const contadorCurtidas = elementoClicado.next('span');
        const quantidadeCurtidas = parseInt(contadorCurtidas.text());

        contadorCurtidas.text(quantidadeCurtidas + 1);

        elementoClicado.addClass('descurtir-publicacao');
        elementoClicado.addClass('text-danger');
        elementoClicado.removeClass('curtir-publicacao');

    }).fail(function(){
        Swal.fire(
            "Ops...",
            "Erro ao curtir publicação!",
            "error"
            );
    }).always(function() {
        //habilita o botão de curtidas novamente pois neste ponto a requisição já terminou
        elementoClicado.prop('disabled', false);
    });
}

function descurtirPublicacao(evento) {
    evento.preventDefault();

    const elementoClicado = $(evento.target);//pega o evento de curtida do botão (coração)
    //pega a div que está mais perto do botão de curtir (coração) na hierarquia acima, fazemos isso para conseguir pegar o id da publicação (publicação-id)
    const publicacaoId = elementoClicado.closest('div').data('publicacao-id');
    //console.log(publicacaoId)

    //desabilita o botão de curtidas enquanto a requisição esta em andameno para evitar de enviar várias requisições ao mesmo tempo evirando que elas se percam
    elementoClicado.prop('disabled', true);
    $.ajax({
        url: `/publicacoes/${publicacaoId}/descurtir`,
        method: "POST"
    }).done(function() {
        //pega o numero de curtidas
        const contadorCurtidas = elementoClicado.next('span');
        const quantidadeCurtidas = parseInt(contadorCurtidas.text());

        contadorCurtidas.text(quantidadeCurtidas - 1);

        elementoClicado.removeClass('descurtir-publicacao');
        elementoClicado.removeClass('text-danger');
        elementoClicado.addClass('curtir-publicacao');

    }).fail(function() {
        Swal.fire(
            "Ops...",
            "Erro ao descurtir publicação!",
            "error"
            );
    }).always(function() {
        //habilita o botão de curtidas novamente pois neste ponto a requisição já terminou
        elementoClicado.prop('disable', false);
    });
}

function atualizarPublicacao(evento) {
    $(this).prop('disabled', true);//desabilitando o botão de atualizar publicação para evitar que o usuário fique clicando varias vezes

    const publicacaoId = $(this).data('publicacao-id');

    $.ajax({
        url: `/publicacoes/${publicacaoId}`,
        method: "PUT",
        data: {
            titulo: $('#titulo').val(),
            conteudo: $('#conteudo').val()
        }
    }).done(function() {
        //utilizando uma ferramenta para lançar mensagens de confirmação na tela
        Swal.fire(
            'Sucesso!',
            'Publicação criada com sucesso',
            'success'
        ).then(function() {//depois que o usuário clicar no ok, direciona para a tela home
            window.location = "/home";
        });
    }).fail(function() {
        Swal.fire(
            "Ops...",
            "Erro ao atualizar publicação!",
            "error"
            );
    }).always(function() {
        //habilita o botão de curtidas novamente pois neste ponto a requisição já terminou
        $('#atualizar-publicacao').prop('disabled', false);
    })
}

function deletarPublicacao(evento) {
    evento.preventDefault();

    Swal.fire({
        titulo: "Atenção!",
        text: "Tem certeza que deseja excluir essa publicação? Essa ação é irreversível",
        showCancelButton: true,
        cancelButtonText: "Cancelar",
        icon: "warning"  
    }).then(function(confirmacao){
        if (!confirmacao.value) return;

        const elementoClicado = $(evento.target);//pega o evento de curtida do botão (coração)
        const publicacao = elementoClicado.closest('div');
        //pega a div que está mais perto do botão de curtir (coração) na hierarquia acima, fazemos isso para conseguir pegar o id da publicação (publicação-id)
        const publicacaoId = publicacao.data('publicacao-id');
        //console.log(publicacaoId)
    
        //desabilita o botão de curtidas enquanto a requisição esta em andameno para evitar de enviar várias requisições ao mesmo tempo evirando que elas se percam
       // elementoClicado.prop('disabled', true);
        
        $.ajax({
            url: `/publicacoes/${publicacaoId}`,
            method: "DELETE"
        }).done(function() {
            //sumindo com o jumbotron da publicação que foi excluida
            publicacao.fadeOut("slow", function() {
                $(this).remove();
            });
        }).fail(function() {
            Swal.fire(
                "Ops...",
                "Erro ao deletar publicação!",
                "error"
                );
        });
    })
}