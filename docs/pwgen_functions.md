# Template Functions

- [`word`](#word)
- [`words`](#words)
- [`wordsWithNum`](#wordswithnum-wordswithnumber)
- [`num`](#num-number-numeric)
- [`alpha`](#alpha)
- [`alphaNum`](#alphaNum)
- [`ascii`](#ascii)
- [`binary`](#binary)
- [`shuffle`](#shuffle)

## `word`

Outputs a random word from the wordlist. For title case, the output can be piped to `title`.

### Examples
- Lowercase:
  ```gotemplate
  {{ word }}
  ```
- Title case:
  ```gotemplate
  {{ title word }}
  ```

## `words`

Outputs a slice of random words from the wordlist. The output will be a slice, which can be joined using `join`. For title case, the output can be piped to `title`.

### Examples
- Lowercase, joined with `-`:
  ```gotemplate
  {{ words 3 | join "-" }}
  ```
- Title case, joined with `-`:
  ```gotemplate
  {{ words 3 | join "-" | title }}
  ```

## `wordsWithNum`, `wordsWithNumber`

Behaves similarly to [`words`](#words), but will append a random number to one of the words.

### Examples
```gotemplate
{{ wordsWithNumber 3 | join "-" }}
```

## `num`, `number`, `numeric`

A random number will be generated with the number of digits determined by the parameter.

### Examples
- Generate a number from 0-9:
  ```gotemplate
  {{ num 1 }}
  ```
- Generate a number from 10-99:
  ```gotemplate
  {{ num 2 }}
  ```

## `alpha`

Random letters will be generated with the length determined by the parameter.

### Examples
- Generate a random string of letters:
  ```gotemplate
  {{ alpha 32 }}
  ```

## `alphaNum`

Random letters and numbers will be generated with the length determined by the parameter.

### Examples
- Generate a random string of letters and numbers:
  ```gotemplate
  {{ alphaNum 32 }}
  ```

## `ascii`

Random letters, numbers, and symbols will be generated with the length determined by the parameter.

### Examples
- Generate a random string of letters, numbers, and symbols:
  ```gotemplate
  {{ ascii 32 }}
  ```

## `binary`

Random binary data will be generated with number of bytes determined by the parameter. Useful with [`b64enc`](https://masterminds.github.io/sprig/encoding.html).

### Examples
- Generate a random base64 string with 32 bytes of data:
  ```gotemplate
  {{ binary 32 | b64enc }}
  ```

## `shuffle`

Randomly shuffles a slice/list.

### Examples
- Shuffles a list of numbers:
  ```gotemplate
  {{ list 1 2 3 4 5 6 | shuffle | join "" }}
  ```
