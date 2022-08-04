1 - Faça broadcast anunciar o conjunto atual de clientes a cada nova chegada (para quem entrar). Isso exige que o conjunto "clients" e os canais "entering" e "leaving" registrem o nome do cliente também. - OK

2 - Faça o servidor desconectar clientes ociosos, aqueles que não enviaram nenhuma mensagem nos ultimos cinco minutos.
Dica: chamar "conn.Close()" em outra gorrotina desbloqueia chamadas ativas a "Read", como aquela feita por "input.Scan()"

3 - Altere o protocolo de rede do servidor de modo que cada cliente forneça seu nome na entrada. Use esse nome no lugar do endereço de rede quando prefixar cada mensagem com a identidade de quem está fazendo o envio. - OK

4 - Falha de qualquer programa cliente em ler dados em tempo hábil, em última instância faz todos os clientes ficarem travados. Modifique o broadcaster para que ignore uma mensagem em vez de esperar, caso o write do cliente não esteja pronto para aceitá-la. De modo alternativo, acrescente o uso de buffer no canal de mensagem de saída de cada cliente, para que a maioria das mensagens não seja descartada; o broadcaster deve usar um envio não bloqueante nesse canal.