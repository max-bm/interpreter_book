# Chapter 1: Lexing

## 1.1 Lexical Analysis

- We need to represent source code in forms that are easier to work with - we're going to change it twice before evaluating it: `source code -> tokens -> abstract syntax tree`.
- The first transformation, from source code to tokens, is called "lexical analysis" (or "lexing"). It's done by a lexer (or tokeniser, or scanner).
- Tokens are small, easily catergorisable data structures that are fed to the parser, which does the second transformation, from tokens to an abstract syntax tree.
- Example: the input to the lexer might be

    ```monkey
    let x = 5 + 5;
    ```

    with the output looking somewhat like

    ```tokens
    [
        LET,
        IDENTIFIER("x"),
        EQUAL_SIGN,
        INTEGER(5),
        PLUS_SIGN,
        INTEGER(5),
        SEMICOLON
    ]
    ```

- What exactly constitutes a "token" is dependent on the lexer implementation, e.g. some lexers would convert the "5" to an integer (as above), whilst some wouldn't convert it until the parsing stage, not when constructing tokens.
- Whitespace characters are not considered when lexing the Monkey language (above), but are for some languages (e.g. Python).
- A production-ready lexer might also attach the line number, column number and filename to a token - for example, this can make error messages more useful.

## 1.2 Defining our Tokens

- We have to define the tokens our lexer is going to output - we'll start with just a few and add more when extending the lexer. The subset of the Monkey language we'll start with looks like

    ```monkey
    let five = 5;
    let ten = 10;

    let add = fn(x, y) {
        x + y;
    };

    let result = add(five, ten);
    ```

- Breaking this down: which types of tokes does the example use?
  - Numbers, `5` and `10`
  - Variable **identifiers**, `x`, `y`, `add` and `result`
  - Language **keywords**, `let` and `fn`
  - **Special characters**, `(`, `)`, `{`, `}`, `=`, `,`, `;`
- Our `Token` data structure needs a "type" attribute to distinguish between "integers" and "right bracket", for example. It also needs a field to hold the literal value of the token to reuse later.
- We defined the `TokenType` to be a string because it's simple to understand, but using an `int` or a `byte` could lead to better performance.
- We define the possible `TokenType`s as constants, with two special types: `ILLEGAL` and `EOF`.

## 1.3 The Lexer

- The lexer we write will take source code as input and output the tokens that represent the code.
- It will go through its inputs and output the next token it recognises - it doesn't need to buffer or save tokens, since there will only be one method called `NextToken`.
- So we'll initialise the lexer with the source code and then repeatedly call `NextToken()` on it to go through the source code, token by token, character by character.
- We make things simpler again by using `string` as the type for the source code.
- Again, in production it's better to attach filenames and line numbers to tokens, so it would be better to initialise the lexer with an `io.Reader` and the filename.
- The reason for two "pointers", `position` and `readPosition`, pointing into our input string is our need to "peek" further into the input whilst looking after the current character to see what comes next. `readPosition` always points to the "next" character in the input. `position` points to the character in the input that corresponds to the `ch` byte.
- Our lexer only supports ASCII characters instead of the full Unicode range for simplicity - supporting Unicode and UTF-8 would require changing `l.ch` from 'byte' to 'rune' and the way we read characters as they can now be multiple bytes wide.
- To turn identifiers and keywords into tokens, our lexer needs to recognise if the current character is a letter - if so, it needs to read the rest of the identifier/keyword until it encounters a non-letter character.
- The `default` branch of the `switch` statement lets us check for identifiers whenever `l.ch` is not a special character, including handling `ILLEGAL` tokens.
- The `isLetter` function has a significant impact on the language our interpreter can parse for such a small function. For example, we treat `_` as a letter and allow it in identifiers and keywords.
- When reading identifiers, we need to be able to tell user-defined identifiers apart from language keywords, i.e. we need a function that returns the correct `TokenType` - we define it in the `token` package

## 1.4 Extending our Token Set and Lexer
