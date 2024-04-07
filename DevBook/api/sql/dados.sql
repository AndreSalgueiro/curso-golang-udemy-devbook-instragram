insert into usuarios (nome, nick, email, senha)
values
("Elisa", "lisa","elisa@gmail.com", "$2a$10$hF829AwDIRPdYh71HQ8BUeLz/2Nd5z7aZD8BnsdGQr5fLOmg/iKEO"),
("Erika", "kel","erika@gmail.com", "$2a$10$hF829AwDIRPdYh71HQ8BUeLz/2Nd5z7aZD8BnsdGQr5fLOmg/iKEO"),
("André", "deco","andre@gmail.com", "$2a$10$hF829AwDIRPdYh71HQ8BUeLz/2Nd5z7aZD8BnsdGQr5fLOmg/iKEO")

insert into seguidores (usuario_id, seguidor_id)
values
(1,2),#le-se usuário 1 é seguido pelo usuário2
(3,1),
(1,3)

insert into publicacoes(titulo, conteudo, autor_id)
values
("Publicacao do usuário 1", "Essa é a publicação do usuário 1!" Oba!, 1),
("Publicacao do usuário 2", "Essa é a publicação do usuário 2!" Oba!, 2),
("Publicacao do usuário 3", "Essa é a publicação do usuário 3!" Oba!, 3),