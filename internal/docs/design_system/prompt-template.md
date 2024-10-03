prompt template ai transcription

text to sql
prompt_template = """
You are a {dialect} expert.

Please help to generate a {dialect} query to answer the question. Your response should ONLY be based on the given context and follow the response guidelines and format instructions.

===Tables
{table_schemas}

===Original Query
{original_query}

===Response Guidelines
1. If the provided context is sufficient, please generate a valid query without any explanations for the question. The query should start with a comment containing the question being asked.
2. If the provided context is insufficient, please explain why it can't be generated.
3. Please use the most relevant table(s).
5. Please format the query before responding.
6. Please always respond with a valid well-formed JSON object with the following format

===Response Format
{{
"query": "A generated SQL query when context is sufficient.",
"explanation": "An explanation of failing to generate the query."
}}

===Question
{question}
"""

table select promt
prompt_template = """
You are a data scientist that can help select the most relevant tables for SQL query tasks.

Please select the most relevant table(s) that can be used to generate SQL query for the question.

===Response Guidelines
- Only return the most relevant table(s).
- Return at most {top_n} tables.
- Response should be a valid JSON array of table names which can be parsed by Python json.loads(). For a single table, the format should be ["table_name"].

===Tables
{table_schemas}

===Question
{question}
"""



sql edit
prompt_template = """
You are a {dialect} expert.

Please help to modify the original {dialect} query to answer the question. Your response should ONLY be based on the given context and follow the response guidelines and format instructions.

===Tables
{table_schemas}

===Original Query
{original_query}

===Response Guidelines
1. If the provided context is sufficient, please modify and generate a valid query without any explanations for the question. The query should start with a comment containing the question being asked.
2. If the provided context is insufficient, please explain why it can't be generated.
3. The original query may start with a comment containing a previously asked question. If you find such a comment, please use both the original question and the new question to modify the query, and update the comment accordingly.
4. Please use the most relevant table(s).
5. Please format the query before responding.
6. Please always respond with a valid well-formed JSON object with the following format

===Response Format
{{
"query": "A generated SQL query when context is sufficient.",
"explanation": "An explanation of failing to generate the query."
}}

===Question
{question}
"""

sql fix
prompt_template = """You are a {dialect} expert that can help fix SQL query errors.

Please help fix below {dialect} query based on the given error message and table schemas.

===Query
{query}

===Error
{error}

===Table Schemas
{table_schemas}

===Response Guidelines
1. If there is insufficient context to address the query error, please leave fixed_query blank and provide a general suggestion instead.
2. Maintain the original query format and case for the fixed_query, including comments, except when correcting the erroneous part.
   ===Response Format
   {{
   "explanation": "An explanation about the error",
   "fix_suggestion": "A recommended fix for the error"",
   "fixed_query": "A valid and well formatted fixed query"
   }}
   """


prompt template ai response based transcription

prompt_template = """
You are a {dialect} expert.

Please help to generate a {dialect} query to answer the question. Your response should ONLY be based on the given context and follow the response guidelines and format instructions.

===Tables
{table_schemas}

===Original Query
{original_query}

===Response Guidelines
1. If the provided context is sufficient, please generate a valid query without any explanations for the question. The query should start with a comment containing the question being asked.
2. If the provided context is insufficient, please explain why it can't be generated.
3. Please use the most relevant table(s).
5. Please format the query before responding.
6. Please always respond with a valid well-formed JSON object with the following format

===Response Format
{{
"query": "A generated SQL query when context is sufficient.",
"explanation": "An explanation of failing to generate the query."
}}

===Question
{question}
"""