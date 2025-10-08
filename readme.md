# Encontros API: Organização de Eventos Informais

![Badge de Status](https://img.shields.io/badge/status-em%20desenvolvimento-yellow)
![Badge de Licença](https://img.shields.io/badge/license-MIT-blue)

Uma API flexível e poderosa para simplificar a organização de encontros informais, desde o futebol semanal com os amigos até o churrasco de fim de ano. Chega de confusão em grupos de WhatsApp!

## 🎯 O Problema que Resolvemos

Organizar qualquer evento em grupo, por menor que seja, geralmente vira uma bagunça em chats:
- Quem vai? A lista de confirmados se perde no meio das conversas.
- O que cada um leva? A famosa pergunta "Precisa levar algo?" gera dezenas de mensagens e ninguém sabe quem ficou responsável pelo quê.
- Sobrou ou faltou gente? Controlar o número de participantes para um jogo ou evento com vagas limitadas é um desafio.

A **Encontros API** centraliza tudo isso, fornecendo uma fonte única de verdade para seu evento.

## ✨ Funcionalidades Principais

- **Criação de Eventos Genéricos**: Crie qualquer tipo de encontro, seja um "Churrasco", "Futebol", "Happy Hour" ou "Trilha no Fim de Semana".
- **Sistema de Confirmação**: Participantes podem confirmar presença de forma simples, preenchendo as vagas disponíveis (se houver um limite).
- **Lista de Atribuições Colaborativa**: Para eventos como churrascos, o organizador pode criar uma lista de itens/tarefas ("Carne", "Bebida", "Carvão") e os participantes podem "reivindicar" a responsabilidade por cada um.
- **Consulta Centralizada**: Um único endpoint para ver todos os detalhes do evento: quem vai e o que cada um levará.
- **Flexibilidade**: O sistema se adapta ao seu evento. Para um futebol, use apenas a confirmação de presença. Para um churrasco, use a confirmação + a lista de atribuições.

## 🔧 Modelo de Dados

A API é estruturada em torno de quatro entidades principais:

| Tabela        | Descrição                            |
|---------------|--------------------------------------|
| `Event`       | Armazena as informações centrais de cada encontro. |
| `User`        | Armazena os usuários.                |
| `Participant` | Uma ponte entre o usuário e eventos. |

## 🚀 Documentação da API (Endpoints)

A seguir estão os principais endpoints para interagir com a API.

---

### Eventos

#### `POST /events/new`
Cria um novo evento.
<details>
  <summary><strong>Exemplo de Requisição</strong></summary>
  
  ```json
  {
    "title": "Churrasco de Fim de Ano",
    "description": "Churrasco para fechar o ano na casa do Bruno!",
    "location": "Rua Fictícia, 123",
    "date_and_time": "2025-12-20T13:00:00",
    "participant_limit": 30
  }
  ```
</details>

#### `GET /events/`
Lista todos os eventos futuros.

#### `GET /evento/{id}`
Busca os detalhes completos de um evento específico, incluindo a lista de confirmados e as atribuições.
<details>
  <summary><strong>Exemplo de Resposta</strong></summary>
  
  ```json
  {
    "event": {
        "id": "01998cb1-2446-7680-bec9-35262ea02638",
        "title": "test",
        "description": "testando",
        "location": "Bar do zé",
        "date_and_time": "2026-02-01T13:40:00-03:00",
        "participant_limit": 1,
        "created_at": "2025-09-27T16:40:43.462467-03:00",
        "updated_at": "2025-09-27T16:40:43.462467-03:00"
    }
}
  ```
</details>

---

### Participação (TODO)

#### `POST /events/{id}/confirm`
Confirma a presença de um participante em um evento.
<details>
  <summary><strong>Exemplo de Requisição</strong></summary>
  
  ```json
  {
    "nome_participante": "Carla"
  }
  ```
</details>

---

### Atribuições (Itens/Tarefas) (TODO)

#### `POST /events/{id}/assignments`
Adiciona um novo item ou tarefa a ser feita para um evento.
<details>
  <summary><strong>Exemplo de Requisição</strong></summary>
  
  ```json
  {
    "descricao": "Levar o som",
    "quantidade_necessaria": 1
  }
  ```
</details>

#### `POST /assignments/{id_atribuicao}/claim`
Um participante assume a responsabilidade por um item/tarefa.
<details>
  <summary><strong>Exemplo de Requisição</strong></summary>
  
  ```json
  {
    "id_participante": 2 // ID da Mariana
  }
  ```
</details>

#### `POST /assignments/{id_atribuicao}/release`
Um participante libera uma atribuição que havia pego, tornando-a disponível novamente.

## 🛠️ Tecnologias Sugeridas

Este projeto pode ser construído com qualquer stack de backend. Uma sugestão popular e moderna seria:

- **Linguagem:** **GoLang**
- **Framework:** **Fiber**
- **Banco de Dados:** **PostgreSQL**
- **ORM:** **Gorm** para uma interação segura e tipada com o banco de dados.
- **Validação:** **validator** para validar os dados de entrada da API.
- **Containerização:** **Docker** para facilitar o setup do ambiente de desenvolvimento.

## 🤝 Como Contribuir

Contribuições são muito bem-vindas! Se você tem ideias para novas funcionalidades, melhorias ou encontrou algum bug, sinta-se à vontade para:

1.  Fazer um "Fork" do projeto.
2.  Criar uma nova "Branch" (`git checkout -b feature/sua-feature`).
3.  Fazer o "Commit" das suas mudanças (`git commit -m 'Adiciona nova feature'`).
4.  Fazer o "Push" para a Branch (`git push origin feature/sua-feature`).
5.  Abrir um "Pull Request".

## 📜 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.
