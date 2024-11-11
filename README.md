# [SQUAD MAGNETO] Hackathon 01 - Currency Converter

### Requisitos Funcionais

1. **Adicionar Taxa de Câmbio**
   - Criar um endpoint para adicionar uma taxa de câmbio entre duas moedas.
   - **Endpoint**: `POST /exchange-rate`
   - **Parâmetros**:
     - `from_currency`: moeda de origem (e.g., USD).
     - `to_currency`: moeda de destino (e.g., BRL).
     - `rate`: taxa de câmbio (float).
   - **Exemplo de Request**:
     ```json
     POST /exchange-rate
     {
       "from_currency": "USD",
       "to_currency": "BRL",
       "rate": 5.25
     }
     ```

2. **Consultar Taxa de Câmbio**
   - Criar um endpoint para consultar a taxa de câmbio atual entre duas moedas.
   - **Endpoint**: `GET /exchange-rate`
   - **Parâmetros de Query**:
     - `from`: moeda de origem.
     - `to`: moeda de destino.
   - **Exemplo de Request**:
     ```json
     GET /exchange-rate?from=USD&to=BRL
     ```
   - **Resposta esperada**:
     ```json
     {
       "rate": 5.25
     }
     ```

3. **Converter Valores**
   - Criar um endpoint para converter um valor com base na taxa de câmbio entre duas moedas.
   - **Endpoint**: `GET /convert`
   - **Parâmetros de Query**:
     - `from`: moeda de origem.
     - `to`: moeda de destino.
     - `amount`: valor a ser convertido.
   - **Exemplo de Request**:
     ```json
     GET /convert?from=USD&to=BRL&amount=100
     ```
   - **Resposta esperada**:
     ```json
     {
       "converted_amount": 525.0
     }
     ```

4. **Error Handler**
   - Implementar tratamento de erros com mensagens claras para os seguintes casos:
     - Tentativa de conversão com moedas não registradas.
     - Taxa de câmbio inexistente entre duas moedas.
     - Dados de entrada inválidos.

---

### Requisitos Técnicos

1. **Banco de Dados**: Utilizar uma base de dados SQLite local para armazenar as taxas de câmbio entre moedas.
   - Estrutura básica da tabela `exchange_rates`:
     - `id`: identificador único.
     - `from_currency`: moeda de origem.
     - `to_currency`: moeda de destino.
     - `rate`: taxa de câmbio.

2. **Logs de Erros**: Capturar e salvar logs de erros gerados nas operações, usando um middleware ou função centralizada de logging.

3. **Execução Local**: O projeto deve ser executável em ambiente local, usando `Go` e o framework `Gin`.
