# Guia de Contribuição

Obrigado por considerar contribuir para o projeto PortuScript! Sua contribuição é valiosa e ajuda a melhorar a linguagem para todos os usuários. Antes de começar, por favor, leia este guia para entender como você pode contribuir de maneira eficaz.

## Como Contribuir

1.  **Crie uma Issue:** Antes de começar a trabalhar em uma nova funcionalidade ou correção, crie uma issue para discutir sua proposta. Isso ajuda a evitar que você gaste tempo em algo que pode não ser aceito ou que já esteja em andamento.
    
2.  **Fork do Repositório:** Faça um fork do repositório PortuScript para sua própria conta do GitHub. Isso permitirá que você faça alterações e envie solicitações de pull.
    
3.  **Clone o Repositório Forked:** Clone o repositório forked para sua máquina local usando o seguinte comando:
    
    ```bash
    git clone https://github.com/seu-usuario/portuscript.git
    ``` 
    
4.  **Crie um Branch:** Antes de começar a trabalhar em uma nova funcionalidade ou correção, crie um branch para sua tarefa específica:
    
    ```bash
    git checkout -b nome-da-sua-tarefa
    ```
    
5.  **Faça as Alterações:** Faça as alterações necessárias no código. Certifique-se de seguir as diretrizes de estilo e manter o código limpo.
    
6.  **Teste suas Alterações:** Execute os testes para garantir que suas alterações não quebraram nada. Se possível, adicione novos testes para cobrir sua funcionalidade.
    
7.  **Envie as Alterações:** Quando você estiver satisfeito com suas alterações, faça commit e envie o branch para o seu fork:
    
    ```bash
    git add .
    git commit -m "feat: Descrição concisa das alterações"
    git push origin nome-da-sua-tarefa
    ```
    
8.  **Crie uma Solicitação de Pull (Pull Request):** Vá para a página do seu fork no GitHub e crie uma Pull Request. Certifique-se de descrever suas alterações e vincular a issue relevante.
    

## Padrão de Commits

Este projeto segue o padrão de commits [Conventional Commits](https://www.conventionalcommits.org/pt-br/).

Os tipos de commit que você pode usar são:

-   **feat:** (nova funcionalidade)
-   **fix:** (correção de bug)
-   **chore:** (tarefas de manutenção)
-   **docs:** (atualização de documentação)
-   **style:** (formatação, ponto e vírgula ausente, etc.)
-   **refactor:** (refatoração de código)
-   **test:** (adição ou modificação de testes)
-   **ci:** (alterações nos scripts de CI/CD)
-   **perf:** (melhorias de desempenho)
-   **exemplo:** (adição de exemplo)
-   **merge:** (merge entre branchs)
-   **config:** (adição/atualização de configurações)

Por exemplo, ao adicionar uma nova funcionalidade, você pode fazer um commit da seguinte forma:

```bash
git commit -m "feat: Adiciona nova funcionalidade incrível"
```

Por favor, siga este padrão ao fazer seus commits.

## Diretrizes de Contribuição

-   Mantenha o código compatível com as versões anteriores, a menos que haja uma razão convincente para a mudança.
-   Siga as diretrizes de estilo do código existente.
-   Documente qualquer nova funcionalidade ou mudança no código.
-   Se estiver adicionando uma nova funcionalidade, considere adicionar testes correspondentes.

## Agradecimentos

O projeto Portuscript agradece a todos os contribuidores pela ajuda e apoio. Suas contribuições fazem deste projeto algo especial.

Se você tiver alguma dúvida ou precisar de assistência, sinta-se à vontade para abrir uma issue ou entrar em contato com o [desenvolvedor](https://github.com/natanfeitosa/).

Obrigado por contribuir para o Portuscript!
