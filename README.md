<h1 align="center"> Ferramenta de Criptografia </h1>

Um programa Go que pode ser executado no terminal para criptografar e descriptografar arquivos. Use o comando `encrypt` para a primeira operação, `decrypt` para a segunda, 
e `help` para aprender mais sobre o programa, se necessário. Qualquer tipo de arquivo pode ser criptografado. O projeto original é de autoria de [@AkhilSharma90](https://github.com/AkhilSharma90).

<h3>Lógica de Criptografia</h3>

O algoritmo utilizado no programa é o **AES (Advanced Encryption Standard)**, em seu modo **GCM (Galois/Counter Mode)**. O modo de funcionamento é o seguinte:

1. Após o arquivo ser aberto e seu conteúdo lido, uma chave é gerada a partir da senha fornecida pelo usuário. O algoritmo utilizado para isso é o PBKDF2, que aplica uma função específica a uma senha várias vezes. Para garantir que o processo seja aleatório (a mesma senha irá gerar chaves diferentes em diferentes tentativas), é criado um *nonce*, um número aleatório usado somente uma vez na operação.
2. A seguir, a chave derivada criada é utilizada para criar blocos de cifra (bytes criptografados). O modo GCM é aplicado.
3. Por fim, o processo é finalizado com a função `Seal` para criptografar.
4. O novo arquivo é criado, sobrescrevendo o anterior, e o *nonce* é adicionado em seus dados, para que, no processo de descriptografia, possamos lê-lo e recriar o arquivo.

O processo de descriptografia é basicamente o inverso.

<h3>Execução</h3>

![Captura de tela de 2024-02-08 09-34-42](https://github.com/KalimaraPeleteiro/CLI-FileEncryption/assets/94702837/1ff1ce1f-055c-4782-a79a-446f4ff75e36)

![Captura de tela de 2024-02-08 09-34-58](https://github.com/KalimaraPeleteiro/CLI-FileEncryption/assets/94702837/13fedfeb-46b5-44b9-9606-6746f018512f)

![Captura de tela de 2024-02-08 09-35-13](https://github.com/KalimaraPeleteiro/CLI-FileEncryption/assets/94702837/59055c70-3579-4811-b317-f14ba15cf301)
